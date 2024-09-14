package miner

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"0chain.net/chaincore/client"
	"0chain.net/chaincore/threshold/bls"
	"0chain.net/chaincore/transaction"
	"go.uber.org/zap"

	"0chain.net/chaincore/block"
	"0chain.net/chaincore/chain"
	"0chain.net/chaincore/node"
	"0chain.net/chaincore/round"
	"0chain.net/chaincore/state"
	"0chain.net/core/cache"
	"0chain.net/core/common"
	"0chain.net/core/datastore"
	"0chain.net/core/memorystore"
	"github.com/0chain/common/core/logging"
)

const (
	// RoundMismatch - to indicate an error where the current round and the
	// given round don't match.
	RoundMismatch = "round_mismatch"
	// RRSMismatch -- to indicate an error when the current round RRS is
	// different than the block that is generated. Typically happens when
	// timeout count changes while a block is being made.
	RRSMismatch = "rrs_mismatch"
	// RoundTimeout - to indicate an error where the round timeout has happened.
	RoundTimeout = "round_timeout"
)

var (
	// ErrRoundMismatch - an error object for mismatched round error.
	ErrRoundMismatch = common.NewError(RoundMismatch, "Current round number"+
		" of the chain doesn't match the block generation round")
	// ErrRRSMismatch - and error when rrs mismatch happens.
	ErrRRSMismatch = common.NewError(RRSMismatch, "RRS for current round"+
		" of the chain doesn't match the block rrs")
	// ErrRoundTimeout - an error object for round timeout error.
	ErrRoundTimeout = common.NewError(RoundTimeout, "round timed out")

	minerChain = &Chain{}

	mcGuard sync.RWMutex
)

/*SetupMinerChain - setup the miner's chain */
func SetupMinerChain(c *chain.Chain) {
	mcGuard.Lock()
	defer mcGuard.Unlock()

	minerChain.Chain = c
	minerChain.Chain.OnBlockAdded = func(b *block.Block) {
	}
	minerChain.ChainConfig = c.ChainConfig

	minerChain.blockMessageChannel = make(chan *BlockMessage, 128)
	c.SetFetchedNotarizedBlockHandler(minerChain)
	c.SetViewChanger(minerChain)
	c.RoundF = MinerRoundFactory{}
	// view change / DKG
	minerChain.viewChangeProcess.init(minerChain)
	// restart round event
	minerChain.subRestartRoundEventChannel = make(chan chan struct{})
	minerChain.unsubRestartRoundEventChannel = make(chan chan struct{})
	minerChain.restartRoundEventChannel = make(chan struct{})
	minerChain.restartRoundEventWorkerIsDoneChannel = make(chan struct{})
	minerChain.nbpMutex = &sync.Mutex{}
	minerChain.notarizationBlockProcessMap = make(map[string]struct{})
	minerChain.notarizationBlockProcessC = make(chan *Notarization, 10)
	minerChain.blockVerifyC = make(chan *block.Block, 10) // the channel buffer size need to be adjusted
	minerChain.manualViewChangeC = make(chan *ViewChangeEvent, 1)
	minerChain.validateTxnsWithContext = common.NewWithContextFunc(1)
	minerChain.notarizingBlocksTasks = make(map[string]chan struct{})
	minerChain.notarizingBlocksResults = cache.NewLRUCache[string, bool](1000)
	minerChain.nbmMutex = &sync.Mutex{}
	minerChain.verifyBlockNotarizationWorker = common.NewWithContextFunc(4)
	minerChain.mergeBlockVRFSharesWorker = common.NewWithContextFunc(1)
	minerChain.verifyCachedVRFSharesWorker = common.NewWithContextFunc(1)
	minerChain.generateBlockWorker = common.NewWithContextFunc(1)
}

/*GetMinerChain - get the miner's chain */
func GetMinerChain() *Chain {
	mcGuard.RLock()
	defer mcGuard.RUnlock()
	return minerChain
}

type StartChain struct {
	datastore.IDField
	Start bool
}

var startChainEntityMetadata *datastore.EntityMetadataImpl

func (sc *StartChain) GetEntityMetadata() datastore.EntityMetadata {
	return startChainEntityMetadata
}

func StartChainProvider() datastore.Entity {
	sc := &StartChain{}
	return sc
}

func SetupStartChainEntity() {
	startChainEntityMetadata = datastore.MetadataProvider()
	startChainEntityMetadata.Name = "start_chain"
	startChainEntityMetadata.Provider = StartChainProvider
	startChainEntityMetadata.IDColumnName = "id"
	datastore.RegisterEntityMetadata("start_chain", startChainEntityMetadata)
}

// MinerRoundFactory -
type MinerRoundFactory struct{}

// CreateRoundF this returns an interface{} of type *miner.Round
func (mrf MinerRoundFactory) CreateRoundF(roundNum int64) round.RoundI {
	mc := GetMinerChain()
	r := round.NewRound(roundNum)
	return mc.CreateRound(r)
}

// Chain - a miner chain to manage the miner activities.
type Chain struct {
	*chain.Chain
	blockMessageChannel chan *BlockMessage
	// muDKG               *sync.RWMutex
	// roundDkg            round.RoundStorage
	discoverClients bool
	started         uint32

	// view change process control
	viewChangeProcess
	manualViewChangeC chan *ViewChangeEvent // TODO: process the StoreDKG and magic block to DB

	// restart round event (rre)
	subRestartRoundEventChannel          chan chan struct{} // subscribe for rre
	unsubRestartRoundEventChannel        chan chan struct{} // unsubscribe rre
	restartRoundEventChannel             chan struct{}      // trigger rre
	restartRoundEventWorkerIsDoneChannel chan struct{}      // rre worker closed
	nbpMutex                             *sync.Mutex
	notarizationBlockProcessMap          map[string]struct{}
	notarizationBlockProcessC            chan *Notarization
	blockVerifyC                         chan *block.Block
	validateTxnsWithContext              *common.WithContextFunc
	notarizingBlocksTasks                map[string]chan struct{}
	notarizingBlocksResults              *cache.LRU[string, bool]
	nbmMutex                             *sync.Mutex
	verifyBlockNotarizationWorker        *common.WithContextFunc
	mergeBlockVRFSharesWorker            *common.WithContextFunc
	verifyCachedVRFSharesWorker          *common.WithContextFunc
	generateBlockWorker                  *common.WithContextFunc
}

type ViewChangeEvent struct {
	MagicBlock *block.MagicBlock
}

func (mc *Chain) sendRestartRoundEvent(ctx context.Context) {
	select {
	case <-ctx.Done(): // caller context is done
	case <-mc.restartRoundEventWorkerIsDoneChannel: // worker context is done
	case mc.restartRoundEventChannel <- struct{}{}:
	}
}

func (mc *Chain) subRestartRoundEvent() (subq chan struct{}) {
	subq = make(chan struct{}, 1)
	select {
	case <-mc.restartRoundEventWorkerIsDoneChannel: // worker context is done
	case mc.subRestartRoundEventChannel <- subq:
	}
	return
}

func (mc *Chain) unsubRestartRoundEvent(subq chan struct{}) {
	select {
	case <-mc.restartRoundEventWorkerIsDoneChannel: // worker context is done
	case mc.unsubRestartRoundEventChannel <- subq:
	}
}

// SetDiscoverClients set the discover clients parameter
func (mc *Chain) SetDiscoverClients(b bool) {
	mc.discoverClients = b
}

// PushBlockMessageChannel pushes the block message to the process channel
func (mc *Chain) PushBlockMessageChannel(bm *BlockMessage) {
	go func() {
		select {
		case mc.blockMessageChannel <- bm:
		case <-time.After(3 * time.Second):
			logging.Logger.Warn("push block message to channel timeout",
				zap.Int("message type", bm.Type))
		}
	}()
}

// SetupGenesisBlock - setup the genesis block for this chain.
func (mc *Chain) SetupGenesisBlock(hash string, magicBlock *block.MagicBlock, initStates *state.InitStates) *block.Block {
	gr, gb := mc.GenerateGenesisBlock(hash, magicBlock, initStates)
	rr, ok := gr.(*round.Round)
	if !ok {
		panic("Genesis round cannot convert to *round.Round")
	}
	mgr := mc.CreateRound(rr)
	mc.AddRound(mgr)
	mc.AddGenesisBlock(gb)
	for _, sharder := range gb.Sharders.Nodes {
		sharder.SetStatus(node.NodeStatusInactive)
	}
	return gb
}

// CreateRound - create a round.
func (mc *Chain) CreateRound(r *round.Round) *Round {
	var mr Round
	mr.Round = r
	mr.blocksToVerifyChannel = make(chan *block.Block, mc.GetGeneratorsNumOfRound(r.GetRoundNumber()))
	mr.verificationTickets = make(map[string]*block.BlockVerificationTicket)
	mr.vrfSharesCache = newVRFSharesCache()
	return &mr
}

// SetLatestFinalizedBlock - sets the latest finalized block.
func (mc *Chain) SetLatestFinalizedBlock(ctx context.Context, b *block.Block) {
	var r = round.NewRound(b.Round)
	mr := mc.CreateRound(r)
	mr = mc.AddRound(mr).(*Round)
	mc.SetRandomSeed(mr, b.GetRoundRandomSeed())
	b = mc.AddRoundBlock(mr, b)
	// mr.SetFinalized()
	mr.Finalize(b)
	mc.AddNotarizedBlock(mr, b)
	mc.Chain.SetLatestFinalizedBlock(b)
	if b.IsStateComputed() {
		if err := mc.SaveChanges(ctx, b); err != nil {
			logging.Logger.Error("set lfb save changes failed",
				zap.Error(err),
				zap.Int64("round", b.Round),
				zap.String("block", b.Hash))
		}
	}
}

// LoadLatestBlocksFromStore loads LFB and LFMB from store and sets them
// to corresponding fields of the sharder's Chain.
func (mc *Chain) LoadLatestBlocksFromStore(ctx context.Context) error {
	// var bl *blocksLoaded
	lfbr, err := mc.LoadLFBRound()
	if err != nil {
		return fmt.Errorf("load_lfb - could not load lfb from state DB, err: %v", err)
	}

	logging.Logger.Debug("load_lfb - load from stateDB",
		zap.Int64("round", lfbr.Round),
		zap.String("block", lfbr.Hash))

	// fetch from sharders
	retry := 3 // retry 3 times, each time wait for about 5 seconds
	var b *block.Block
	for i := 0; i < retry; i++ {
		b, err = mc.GetNotarizedBlockFromSharders(ctx, lfbr.Hash, lfbr.Round)
		if err != nil {
			logging.Logger.Error("load_lfb - could not fetch block from sharders, wait for retry...",
				zap.Int64("round", lfbr.Round), zap.String("block", lfbr.Hash), zap.Error(err))
			// try fetch from miners
			time.Sleep(5 * time.Second)
			// b, err = mc.GetNotarizedBlockFromMiners(ctx, lfbr.Hash, lfbr.Round, true)
			// if err != nil {
			// 	logging.Logger.Error("load_lfb - could not fetch block from miners",
			// 		zap.Int64("round", lfbr.Round), zap.String("block", lfbr.Hash), zap.Error(err))
			// 	return fmt.Errorf("load_lfb - could not fetch block from miners, round: %d, err: %v", lfbr.Round, err)
			// }
			continue
		}
		break
	}

	if b == nil {
		return fmt.Errorf("load_lfb - could not fetch block from sharders, round: %d", lfbr.Round)
	}

	b.SetStateStatus(block.StateSuccessful)
	if err = mc.InitBlockState(b); err != nil {
		b.SetStateStatus(0)
		logging.Logger.Error("load_lfb -- can't initialize stored block state",
			zap.Error(err))
		return fmt.Errorf("can't init block state: %v", err) // fatal
	}

	mc.SetLatestFinalizedBlock(ctx, b)

	// if lfmb.StartingRound > 0 {
	// 	for i := 0; i < retry; i++ {
	// 		mb, err := mc.GetNotarizedBlockFromSharders(ctx, "", lfmb.StartingRound)
	// 		if err != nil {
	// 			logging.Logger.Error("load_lfb - could not fetch latest finalized magic block from sharders",
	// 				zap.Int64("mb_starting_round", lfmb.StartingRound), zap.Error(err))
	// 			time.Sleep(5 * time.Second)
	// 			continue
	// 		}
	// 		mc.updateMagicBlocks(mb)
	// 		break
	// 	}
	// }

	logging.Logger.Info("load_lfb setup LFB from store",
		zap.String("block", b.Hash),
		zap.Int64("round", b.Round),
		zap.Int64("lf_round", mc.GetLatestFinalizedBlock().Round))

	return nil
}

func (mc *Chain) loadLatestFinalizedMagicBlockFromStore(ctx context.Context) {
	lfmb := mc.GetLatestMagicBlock()
	// load the latest N magic blocks
	n := int64(5)
	retry := 3

	if lfmb.MagicBlockNumber <= 1 {
		return
	}

	// magic block number start from 1, the genesis block
	startNum := int64(2) // 1 is the genesis block, we have it locally, so don't need to fetch from remote
	if lfmb.MagicBlockNumber < startNum {
		// genesis block, return
		return
	}

	newStart := lfmb.MagicBlockNumber - n
	if newStart > startNum {
		startNum = newStart
	}

	for i := startNum; i <= lfmb.MagicBlockNumber; i++ {
		// load MB from local store
		mbStr := strconv.FormatInt(i, 10)
		mb, err := LoadMagicBlock(ctx, mbStr)
		if err != nil {
			logging.Logger.Panic("load_latest_mb", zap.Error(err), zap.Int64("mb number", i))
		}

		logging.Logger.Info("[mvc] load MB by magic bock number", zap.Int64("mb number", i))
		for j := 0; j < retry; j++ {
			bmb, err := mc.GetNotarizedBlockFromSharders(ctx, "", mb.StartingRound)
			if err != nil {
				logging.Logger.Error("load_lfb - could not fetch latest finalized magic block from sharders",
					zap.Int64("mb_starting_round", lfmb.StartingRound), zap.Error(err))
				time.Sleep(3 * time.Second)
				continue
			}
			mc.updateMagicBlocks(bmb)
			break
		}
	}
}

func (mc *Chain) deleteTxns(txns []datastore.Entity) error {
	transactionMetadataProvider := datastore.GetEntityMetadata("txn")
	ctx := memorystore.WithEntityConnection(common.GetRootContext(), transactionMetadataProvider)
	defer memorystore.Close(ctx)
	txnHashes := make([]string, len(txns))
	for i, txn := range txns {
		txnHashes[i] = txn.(*transaction.Transaction).Hash
	}
	logging.Logger.Debug("delete txns", zap.Any("txns", txnHashes))
	return transactionMetadataProvider.GetStore().MultiDelete(ctx, transactionMetadataProvider, txns)
}

// SetPreviousBlock - set the previous block.
func (mc *Chain) SetPreviousBlock(r round.RoundI, b *block.Block, pb *block.Block) {
	b.SetPreviousBlock(pb)
	mc.SetRoundRank(r, b)
}

// GetMinerRound - get the miner's version of the round.
func (mc *Chain) GetMinerRound(roundNumber int64) *Round {
	r := mc.GetRound(roundNumber)
	if r == nil {
		return nil
	}
	mr, ok := r.(*Round)
	if !ok {
		return nil
	}
	return mr
}

// SaveClients - save clients from the block.
func (mc *Chain) SaveClients(clients []*client.Client) error {
	var err error
	clientKeys := make([]datastore.Key, len(clients))
	for idx, c := range clients {
		clientKeys[idx] = c.GetKey()
	}
	clientEntityMetadata := datastore.GetEntityMetadata("client")
	cEntities := datastore.AllocateEntities(len(clients), clientEntityMetadata)
	ctx := memorystore.WithEntityConnection(common.GetRootContext(), clientEntityMetadata)
	defer memorystore.Close(ctx)
	err = clientEntityMetadata.GetStore().MultiRead(ctx, clientEntityMetadata, clientKeys, cEntities)
	if err != nil {
		return err
	}
	ctx = datastore.WithAsyncChannel(ctx, client.ClientEntityChannel)
	for idx, c := range clients {
		if !datastore.IsEmpty(cEntities[idx].GetKey()) {
			continue
		}
		_, cerr := client.PutClient(ctx, c)
		if cerr != nil {
			err = cerr
		}
	}
	return err
}

// ViewChange on finalized (!) block. Miners check magic blocks during
// generation and notarization. A finalized block should be trusted.
func (mc *Chain) ViewChange(ctx context.Context, b *block.Block) (err error) {
	// if !mc.ChainConfig.IsViewChangeEnabled() {
	// 	return
	// }

	if b.MagicBlock == nil {
		return nil
	}

	mb := b.MagicBlock

	// persist the MB whenever see it. Miners could restart with lfb with this mb_number
	if err = StoreMagicBlock(ctx, mb); err != nil {
		return common.NewErrorf("view_change", "saving MB data: %v", err)
	}

	if !mb.Miners.HasNode(node.Self.Underlying().GetKey()) {
		logging.Logger.Error("[mvc] view change, magic miners does not have self node")
		return // node leaves BC, don't do anything here
	}

	if mc.isSyncingBlocks() {
		// Just store the magic block and return
		// if err = StoreMagicBlock(ctx, mb); err != nil {
		// 	logging.Logger.Panic("[mvc] failed to store magic block", zap.Error(err))
		// }
		return nil
	}

	mc.viewChangeProcess.Lock()
	defer mc.viewChangeProcess.Unlock()
	var (
		// mpks = mc.viewChangeProcess.mpks.GetMpks()
		vcd         = mc.viewChangeProcess.viewChangeDKG
		selfNodeKey = node.Self.Underlying().GetKey()
		vcdkg       = bls.MakeDKG(mb.T, mb.N, selfNodeKey)
	)

	// if len(mpks) == 0 {
	// 	// the miner may just start up, the viewChangeProcess is not set yet
	// 	return nil
	// }

	for key, share := range mb.GetShareOrSigns().GetShares() {
		if key == selfNodeKey {
			continue // skip self
		}
		var myShare, ok = share.ShareOrSigns[selfNodeKey]
		if ok && myShare.Share != "" {
			var share bls.Key
			if err := share.SetHexString(myShare.Share); err != nil {
				return err
			}
			// lmpks, err := bls.ConvertStringToMpk(mpks[key].Mpk)
			// if err != nil {
			// 	return err
			// }

			// var validShare = vcdkg.ValidateShare(lmpks, share)
			// if !validShare {
			// 	continue
			// }
			err = vcdkg.AddSecretShare(bls.ComputeIDdkg(key), myShare.Share,
				true)
			if err != nil {
				return common.NewErrorf("view_change", "adding secret share: %v", err)
			}
		}
	}

	// var miners []string
	// for key := range mc.viewChangeProcess.mpks.GetMpks() {
	// 	if _, ok := mb.Mpks.Mpks[key]; !ok {
	// 		miners = append(miners, key)
	// 	}
	// }
	// vcdkg.DeleteFromSet(miners)
	mpkMap, err := mb.Mpks.GetMpkMap()
	if err != nil {
		return err
	}
	if err := vcdkg.AggregatePublicKeyShares(mpkMap); err != nil {
		return err
	}

	// get self dkg share
	pid := bls.ComputeIDdkg(selfNodeKey)
	share, ok := vcd.GetSecretShare(selfNodeKey)
	if !ok {
		// load from store
		dkgKey, err := LoadDKGKey(ctx, mb.MagicBlockNumber)
		if err != nil {
			logging.Logger.Debug("[mvc] view_change failed to load dkg key",
				zap.Error(err),
				zap.Int64("mb number", mb.MagicBlockNumber))
		}
		logging.Logger.Debug("[mvc] view_change, load dkg key from store",
			zap.String("dkg key", dkgKey.KeyShare),
			zap.Int64("mb number", mb.MagicBlockNumber))
		if err := vcdkg.AddSecretShare(pid, dkgKey.KeyShare, true); err != nil {
			logging.Logger.Error("[mvc] view_change, failed to add self dkg share from store", zap.Error(err))
			return err
		}
	} else {
		if err := vcdkg.AddSecretShare(pid, share.GetHexString(), true); err != nil {
			logging.Logger.Error("[mvc] view_change failed to add self dkg share", zap.Error(err))
			return err
		}
	}

	vcdkg.AggregateSecretKeyShares()
	vcdkg.StartingRound = mb.StartingRound
	vcdkg.MagicBlockNumber = mb.MagicBlockNumber
	// set T and N from the magic block
	vcdkg.T = mb.T
	vcdkg.N = mb.N

	// save DKG and MB
	if err = StoreDKGSummary(ctx, vcdkg.GetDKGSummary()); err != nil {
		return common.NewErrorf("view_change", "saving DKG summary: %v", err)
	}

	if err := SetDKG(ctx, mb); err != nil {
		logging.Logger.Debug("[mvc] view change set dkg failed",
			zap.Int64("mb number", mb.MagicBlockNumber),
			zap.Int64("mb sr", mb.StartingRound),
			zap.Error(err))
	}

	// logging.Logger.Debug("[mvc] set next view change round")
	// if nvc, err = mc.NextViewChangeOfBlock(b); err != nil {
	// 	return common.NewErrorf("view_change", "getting nvc: %v", err)
	// }

	// // set / update the next view change of protocol view change (RAM)
	// //
	// // note: this approach works where a miners is active and finalizes blocks
	// //       but for inactive miners we have to set next view change based on
	// //       blocks fetched from sharders
	// mc.SetNextViewChange(nvc)

	// // next view change is expected, but not given;
	// // it means the MB rejected by Miner SC
	// if b.Round == nvc && mb == nil {
	// 	return // no MB no VC
	// }

	// // just skip the block if it hasn't a MB
	// if mb == nil {
	// 	return // no MB, no VC
	// }

	// // view change

	// if err = mc.UpdateMagicBlock(mb); err != nil {
	// 	return common.NewErrorf("view_change", "updating MB: %v", err)
	// }

	// mc.SetLatestFinalizedMagicBlock(b)

	// go mc.PruneRoundStorage(mc.getPruneCountRoundStorage(),
	// 	mc.GetRoundDkg(), mc.MagicBlockStorage)

	// // set DKG if this node is miner of new MB (it have to have the DKG)
	// var selfNodeKey = node.Self.Underlying().GetKey()

	// if !mb.Miners.HasNode(selfNodeKey) {
	// 	return // ok, all done
	// }

	// // this must be ok, if not -- return error
	// if err = mc.SetDKGSFromStore(ctx, mb); err != nil {
	// 	logging.Logger.Error("view_change - set DKG failed",
	// 		zap.Int64("mb_starting_round", mb.StartingRound),
	// 		zap.Error(err))
	// 	return
	// }

	// new miners has no previous round, and LFB, this block becomes
	// LFB and new miners have to get it from miners or sharders to
	// join BC; now we have to kick them to don't wait for while and
	// get the block from sharders and join BC; anyway the new miners
	// will pool LFB (501) from sharders and join; but the kick used
	// to skip a waiting

	// to mine 503 round (new round with new nodes and new MB)
	// the new miners need:
	//    - 501 block with MB, corresponding DKG saved locally
	//    - 502 block and round (for previous round random seed)

	// the flow:
	//
	//  - 501 - finalized
	//  - 502 - finalize round (finalize 501 block)
	//  - 503 - verify round blocks
	//  - 504 - generate round (new MB/DKG can be used, but slower, use old)
	//  - 505 - generate block (use new MB/DKG)

	return
}

func StartChainRequestHandler(_ context.Context, req *http.Request) (interface{}, error) {
	nodeID := req.Header.Get(node.HeaderNodeID)
	mc := GetMinerChain()

	r, err := strconv.Atoi(req.FormValue("round"))
	if err != nil {
		logging.Logger.Error("failed to send start chain", zap.Error(err))
		return nil, err
	}

	mb := mc.GetMagicBlock(int64(r))
	if mb == nil || !mb.Miners.HasNode(nodeID) {
		logging.Logger.Error("failed to send start chain", zap.String("id", nodeID))
		return nil, common.NewError("failed to send start chain", "miner is not in active set")
	}

	if mc.GetCurrentRound() != int64(r) {
		logging.Logger.Error("failed to send start chain -- different rounds", zap.Int64("current_round", mc.GetCurrentRound()), zap.Int("requested_round", r))
		return nil, common.NewError("failed to send start chain", fmt.Sprintf("differt_rounds -- current_round: %v, requested_round: %v", mc.GetCurrentRound(), r))
	}
	message := datastore.GetEntityMetadata("start_chain").Instance().(*StartChain)
	message.Start = !mc.isStarted()
	message.ID = req.FormValue("round")
	return message, nil
}

func (mc *Chain) SetStarted() {
	if !atomic.CompareAndSwapUint32(&mc.started, 0, 1) {
		logging.Logger.Warn("chain already started")
	}
}

func (mc *Chain) isStarted() bool {
	return atomic.LoadUint32(&mc.started) == 1
}

// SaveMagicBlock returns nil.
func (mc *Chain) SaveMagicBlock() chain.MagicBlockSaveFunc {
	return nil
}

func mbRoundOffset(rn int64) int64 {
	if rn < chain.ViewChangeOffset+1 {
		return rn // the same
	}
	return rn - chain.ViewChangeOffset // MB offset
}

func (mc *Chain) RejectNotarizedBlock(_ string) bool {
	return false
}
