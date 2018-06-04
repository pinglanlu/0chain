package miner

import (
	"context"
	"fmt"
	"math/rand"

	"0chain.net/common"
	"0chain.net/datastore"
	"0chain.net/memorystore"
	"0chain.net/node"
	"0chain.net/round"
)

/*SetupWorkers - Setup the miner's workers */
func SetupWorkers() {
	ctx := common.GetRootContext()
	go GetMinerChain().BlockWorker(ctx)
}

/*BlockWorker - a job that does all the work related to blocks in each round */
func (mc *Chain) BlockWorker(ctx context.Context) {
	for true {
		select {
		case <-ctx.Done():
			return
		case msg := <-mc.GetBlockMessageChannel():
			if msg.Sender != nil {
				fmt.Printf("received message: %v from %v(%v)\n", msg.Type, msg.Sender.SetIndex, msg.Sender.GetKey())
			} else {
				fmt.Printf("received message: %v\n", msg.Type)
			}
			switch msg.Type {
			case MessageStartRound:
				mc.HandleStartRound(ctx, msg)
			case MessageVerify:
				mc.HandleVerifyBlockMessage(ctx, msg)
			case MessageVerificationTicket:
				mc.HandleVerificationTicketMessage(ctx, msg)
			case MessageConsensus:
				mc.HandleConsensusMessage(ctx, msg)
			}
		}
	}
}

/*HandleStartRound - handles the start round message */
func (mc *Chain) HandleStartRound(ctx context.Context, msg *BlockMessage) {
	r := msg.Round
	mc.startNewRound(ctx, r)
}

func (mc *Chain) startNewRound(ctx context.Context, r *round.Round) {
	pr := mc.GetRound(r.Number - 1)
	//TODO: If for some reason the server is lagging behind (like network outage) we need to fetch the previous round info
	// before proceeding
	if pr == nil {
		return
	}
	if !mc.AddRound(r) {
		return
	}
	self := node.GetSelfNode(ctx)
	rank := r.GetRank(self.SetIndex)
	fmt.Printf("*** Starting round (%v) with (set index=%v, round rank=%v)\n", r.Number, self.SetIndex, rank)
	//TODO: For now, if the rank happens to be in the bottom half, we assume no need to generate block
	if 2*rank > mc.Miners.Size() {
		return
	}
	txnEntityMetadata := datastore.GetEntityMetadata("txn")
	ctx = memorystore.WithEntityConnection(ctx, txnEntityMetadata)
	defer memorystore.Close(ctx)
	mc.GenerateRoundBlock(ctx, r)
}

/*HandleVerifyBlockMessage - handles the verify block message */
func (mc *Chain) HandleVerifyBlockMessage(ctx context.Context, msg *BlockMessage) {
	bvt, err := mc.VerifyBlock(ctx, msg.Block)
	if err != nil {
		return
	}
	r := mc.GetRound(msg.Block.Round)
	if r != nil {
		if r.Block == nil {
			r.Block = msg.Block
		}
	}
	mc.AddBlock(msg.Block)
	mc.SendVerificationTicket(ctx, msg.Block, bvt)
}

/*HandleVerificationTicketMessage - handles the verification ticket message */
func (mc *Chain) HandleVerificationTicketMessage(ctx context.Context, msg *BlockMessage) {
	if mc.ReachedConsensus(ctx, msg.Block) {
		return
	}
	err := mc.VerifyTicket(ctx, msg.Block, msg.BlockVerificationTicket)
	if err != nil {
		return
	}
	if mc.AddVerificationTicket(ctx, msg.Block, &msg.BlockVerificationTicket.VerificationTicket) {
		if mc.ReachedConsensus(ctx, msg.Block) {
			consensus := datastore.GetEntityMetadata("block_consensus").Instance().(*Consensus)
			consensus.BlockID = msg.Block.Hash
			consensus.VerificationTickets = msg.Block.VerificationTickets
			mc.SendConsensus(ctx, consensus)
			r := mc.GetRound(msg.Block.Round)
			r.Block = msg.Block
			if mc.GetRound(r.Number+1) == nil {
				nr := datastore.GetEntityMetadata("round").Instance().(*round.Round)
				nr.Number = r.Number + 1
				nr.RandomSeed = rand.New(rand.NewSource(r.RandomSeed)).Int63()
				go mc.startNewRound(ctx, nr)
				mc.Miners.SendAll(RoundStartSender(nr))
			}
			//TODO: Finalize previous round
		}
	}
}

/*HandleConsensusMessage - handles the block consensus message */
func (mc *Chain) HandleConsensusMessage(ctx context.Context, msg *BlockMessage) {
	//TODO: Finalize previous round
}
