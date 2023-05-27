package event

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"0chain.net/chaincore/config"
	"golang.org/x/net/context"

	"0chain.net/smartcontract/dbs"

	"go.uber.org/zap"

	"github.com/0chain/common/core/logging"
)

var ErrInvalidEventData = errors.New("invalid event data")

type (
	ProcessEventsOptions struct {
		CommitNow bool
	}
	ProcessEventsOptionsFunc func(peo *ProcessEventsOptions)
)

func CommitNow() ProcessEventsOptionsFunc {
	return func(peo *ProcessEventsOptions) {
		peo.CommitNow = true
	}
}

// CommitOrRollbackFunc represents the callback function to do commit
// or rollback.
type CommitOrRollbackFunc func(rollback bool) error

// ProcessEvents - process events and return commit function or error if any
// The commit function can be called to commit the events changes when needed
func (edb *EventDb) ProcessEvents(
	ctx context.Context,
	events []Event,
	round int64,
	block string,
	blockSize int,
	opts ...ProcessEventsOptionsFunc,
) (*EventDb, error) {
	ts := time.Now()
	es, err := mergeEvents(round, block, events)
	if err != nil {
		return nil, err
	}

	pdu := time.Since(ts)

	event := blockEvents{
		events:    es,
		round:     round,
		block:     block,
		blockSize: blockSize,
		txC:       make(chan *EventDb, 1),
	}

	select {
	case edb.eventsChannel <- event:
	case <-ctx.Done():
		logging.Logger.Warn("process events - context done",
			zap.Error(ctx.Err()),
			zap.Int64("round", round),
			zap.String("block", block),
			zap.Int("block size", blockSize))
		return nil, fmt.Errorf("process events - push to process channel context done: %v", ctx.Err())
	}

	select {
	case eTx := <-event.txC:
		du := time.Since(ts)
		if du.Milliseconds() > 200 {
			logging.Logger.Warn("process events slow",
				zap.Duration("duration", du),
				zap.Duration("merge events duration", pdu),
				zap.Int64("round", round),
				zap.String("block", block),
				zap.Int("block size", blockSize))
		}
		var opt ProcessEventsOptions
		for _, f := range opts {
			f(&opt)
		}

		if opt.CommitNow {
			return nil, eTx.Commit()
		}

		return eTx, nil
	case <-ctx.Done():
		du := time.Since(ts)
		logging.Logger.Warn("process events - context done",
			zap.Error(ctx.Err()),
			zap.Duration("duration", du),
			zap.Int64("round", round),
			zap.String("block", block),
			zap.Int("block size", blockSize))
		return nil, ctx.Err()
	}
}

func mergeEvents(round int64, block string, events []Event) ([]Event, error) {
	var (
		mergers = []eventsMerger{
			mergeAddUsersEvents(),
			mergeAddProviderEvents[Miner](TagAddMiner, withUniqueEventOverwrite()),
			mergeAddProviderEvents[Sharder](TagAddSharder, withUniqueEventOverwrite()),
			mergeAddProviderEvents[Blobber](TagAddBlobber, withUniqueEventOverwrite()),
			mergeAddProviderEvents[Blobber](TagUpdateBlobber, withUniqueEventOverwrite()),
			mergeAddProviderEvents[Validator](TagAddOrOverwiteValidator, withUniqueEventOverwrite()),
			mergeAddProviderEvents[dbs.ProviderID](TagShutdownProvider, withUniqueEventOverwrite()),
			mergeAddProviderEvents[dbs.ProviderID](TagKillProvider, withUniqueEventOverwrite()),

			mergeAddAllocationEvents(),
			mergeUpdateAllocEvents(),
			mergeUpdateAllocStatsEvents(),
			mergeUpdateAllocBlobbersTermsEvents(),
			mergeAddOrOverwriteAllocBlobbersTermsEvents(),
			mergeDeleteAllocBlobbersTermsEvents(),

			mergeAddChallengesEvents(),
			mergeAddChallengesToAllocsEvents(),

			mergeUpdateChallengesEvents(),
			mergeAddChallengePoolsEvents(),

			mergeUpdateBlobberChallengesEvents(),
			mergeAddChallengesToBlobberEvents(),
			mergeUpdateAllocChallengesEvents(),

			mergeUpdateBlobbersEvents(),
			mergeUpdateBlobberTotalStakesEvents(),
			mergeUpdateBlobberTotalOffersEvents(),
			mergeStakePoolRewardsEvents(),
			mergeStakePoolPenaltyEvents(),
			mergeAddDelegatePoolsEvents(),

			mergeUpdateMinerTotalStakesEvents(),
			mergeUpdateSharderTotalStakesEvents(),
			mergeUpdateAuthorizerTotalStakesEvents(),

			mergeAddTransactionsEvents(),
			mergeAddWriteMarkerEvents(),
			mergeAddReadMarkerEvents(),
			mergeAllocationStatsEvents(),
			mergeUpdateBlobberStatsEvents(),
			mergeUpdateValidatorsEvents(),
			mergeUpdateValidatorStakesEvents(),

			mergeMinerHealthCheckEvents(),
			mergeSharderHealthCheckEvents(),
			mergeBlobberHealthCheckEvents(),
			mergeAuthorizerHealthCheckEvents(),
			mergeValidatorHealthCheckEvents(),

			mergeAddBurnTicket(),

			mergeUpdateUserCollectedRewardsEvents(),
			mergeUserStakeEvents(),
			mergeUserUnstakeEvents(),
			mergeUserReadPoolLockEvents(),
			mergeUserReadPoolUnlockEvents(),
			mergeUserWritePoolLockEvents(),
			mergeUserWritePoolUnlockEvents(),
			mergeUpdateUserPayedFeesEvents(),
			mergeAuthorizerBurnEvents(),
			mergeAddBridgeMintEvents(),
		}

		others = make([]Event, 0, len(events))
	)

	for _, e := range events {
		if e.Type == TypeChain || e.Tag == TagUniqueAddress {
			others = append(others, e)
			continue
		}
		if e.Type != TypeStats {
			continue
		}

		var matched bool
		for _, em := range mergers {
			if em.filter(e) {
				matched = true
				break
			}
		}

		if matched {
			continue
		}

		others = append(others, e)
	}

	mergedEvents := make([]Event, 0, len(mergers))
	for _, em := range mergers {
		e, err := em.merge(round, block)
		if err != nil {
			return nil, err
		}

		if e != nil {
			mergedEvents = append(mergedEvents, *e)
		}
	}

	return append(mergedEvents, others...), nil
}

func (edb *EventDb) addEventsWorker(ctx context.Context) {
	var gs *Snapshot
	p := int64(-1)
	edb.managePartitions(0)

	for {
		es := <-edb.eventsChannel

		s, err := edb.work(ctx, gs, es, &p)
		if err != nil {
			if config.Development() { //panic in case of development
				log.Panic(err)
			}
		}
		if s != nil {
			gs = s
		}
	}
}

func (edb *EventDb) work(ctx context.Context,
	gs *Snapshot, es blockEvents, currentPartition *int64) (*Snapshot, error) {
	tx, err := edb.Begin()
	if err != nil {
		logging.Logger.Error("error starting transaction", zap.Error(err))
		return nil, err
	}

	var commit bool
	defer func() {
		if commit {
			es.txC <- tx
		} else {
			es.txC <- nil
		}
	}()

	if *currentPartition < es.round/edb.settings.PartitionChangePeriod {
		tx.managePartitions(es.round)
		*currentPartition = es.round / edb.settings.PartitionChangePeriod
	}

	if err = tx.addEvents(ctx, es); err != nil {
		logging.Logger.Error("error saving events",
			zap.Int64("round", es.round),
			zap.Error(err))

		if rerr := tx.Rollback(); rerr != nil {
			return nil, rerr
		}
		return nil, err
	}

	tse := time.Now()
	tags := make([]string, 0, len(es.events))
	for _, event := range es.events {
		tags, err = tx.processEvent(event, tags, es.round, es.block, es.blockSize)
		if err != nil {
			logging.Logger.Error("error processing event",
				zap.Int64("round", event.BlockNumber),
				zap.Any("tag", event.Tag),
				zap.Error(err))

			if rerr := tx.Rollback(); rerr != nil {
				logging.Logger.Error("commit error", zap.Error(rerr))
				return nil, rerr
			}

			return nil, err
		}
	}

	// process snapshot for none adding block events only
	if isNotAddBlockEvent(es) {
		gs, err = updateSnapshots(gs, es, tx)
		if err != nil {
			logging.Logger.Error("snapshot could not be processed",
				zap.Int64("round", es.round),
				zap.String("block", es.block),
				zap.Int("block size", es.blockSize),
				zap.Error(err),
			)
		}
		err = tx.updateUserAggregates(&es)
		if err != nil {
			logging.Logger.Error("user aggregate could not be processed",
				zap.Error(err),
			)
		}
	}

	commit = true

	due := time.Since(tse)
	if due.Milliseconds() > 200 {
		logging.Logger.Warn("event db work slow",
			zap.Duration("duration", due),
			zap.Int("events number", len(es.events)),
			zap.Strings("tags", tags),
			zap.Int64("round", es.round),
			zap.String("block", es.block),
			zap.Int("block size", es.blockSize))
	}
	return gs, nil
}

func (edb *EventDb) managePartitions(round int64) {
	logging.Logger.Info("managing partitions", zap.Int64("round", round))
	if err := edb.addPartition(round, "events"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "events"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "snapshots"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "snapshots"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "blobber_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "blobber_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "miner_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "miner_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "sharder_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "sharder_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "validator_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "validator_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "authorizer_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "authorizer_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
	if err := edb.addPartition(round, "user_aggregates"); err != nil {
		logging.Logger.Error("error creating partition", zap.Error(err))
	}
	if err := edb.dropPartition(round, "user_aggregates"); err != nil {
		logging.Logger.Error("error dropping partition", zap.Error(err))
	}
}

func isNotAddBlockEvent(es blockEvents) bool {
	return !(len(es.events) == 1 && es.events[0].Type == TypeChain && es.events[0].Tag == TagAddBlock)
}

func updateSnapshots(gs *Snapshot, es blockEvents, tx *EventDb) (*Snapshot, error) {
	if gs != nil {
		return tx.updateSnapshots(es, gs)
	}

	if es.round == 0 {
		return tx.updateSnapshots(es, &Snapshot{Round: 0})
	}

	g, err := tx.GetGlobal()
	if err != nil {
		logging.Logger.Panic("can't load snapshot for", zap.Int64("round", es.round), zap.Error(err))
	}
	gs = &g

	return tx.updateSnapshots(es, gs)
}

func (edb *EventDb) processEvent(event Event, tags []string, round int64, block string, blockSize int) ([]string, error) {
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("panic recovered in processEvent",
				zap.Any("r", r),
				zap.Any("event", event))
		}
	}()
	var err error = nil
	switch event.Type {
	case TypeStats:
		tags = append(tags, event.Tag.String())
		ts := time.Now()
		err = edb.addStat(event)
		if err != nil {
			logging.Logger.Error("addStat typeStats error",
				zap.Int64("round", round),
				zap.String("block", block),
				zap.Int("block size", blockSize),
				zap.Any("event type", event.Type),
				zap.Any("event tag", event.Tag),
				zap.Error(err),
			)
		}
		du := time.Since(ts)
		if du.Milliseconds() > 50 {
			logging.Logger.Warn("event db save slow - addStat",
				zap.Duration("duration", du),
				zap.String("event tag", event.Tag.String()),
				zap.Int64("round", round),
				zap.String("block", block),
				zap.Int("block size", blockSize),
			)
		}
	case TypeChain:
		tags = append(tags, event.Tag.String())
		ts := time.Now()
		err = edb.addStat(event)
		du := time.Since(ts)
		if du.Milliseconds() > 50 {
			logging.Logger.Warn("event db save slow - addchain",
				zap.Duration("duration", du),
				zap.String("event tag", event.Tag.String()),
				zap.Int64("round", round),
				zap.String("block", block),
				zap.Int("block size", blockSize),
			)
		}
	case TypeError:
		err = edb.addError(Error{
			TransactionID: event.TxHash,
			Error:         fmt.Sprintf("%v", event.Data),
		})
	default:
	}
	if err != nil {
		logging.Logger.Error("event could not be processed",
			zap.Int64("round", round),
			zap.String("block", block),
			zap.Int("block size", blockSize),
			zap.Any("event type", event.Type),
			zap.Any("event tag", event.Tag),
			zap.Error(err),
		)
		return tags, err
	}
	return tags, nil
}

func (edb *EventDb) updateSnapshots(e blockEvents, s *Snapshot) (*Snapshot, error) {
	round := e.round
	var events []Event
	for _, ev := range e.events { //filter out round events
		if ev.Type == TypeStats || (ev.Type == TypeChain && ev.Tag == TagFinalizeBlock) {
			events = append(events, ev)
		}
	}
	if len(events) == 0 {
		return s, nil
	}

	logging.Logger.Debug("getting blobber aggregate ids", zap.Any("snapshot_before", s))

	edb.updateBlobberAggregate(round, edb.AggregatePeriod(), s)
	edb.updateMinerAggregate(round, edb.AggregatePeriod(), s)
	edb.updateSharderAggregate(round, edb.AggregatePeriod(), s)
	edb.updateAuthorizerAggregate(round, edb.AggregatePeriod(), s)
	edb.updateValidatorAggregate(round, edb.AggregatePeriod(), s)
	s.update(events)

	s.Round = round
	if err := edb.addSnapshot(*s); err != nil {
		logging.Logger.Error(fmt.Sprintf("saving snapshot %v for round %v", s, round), zap.Error(err))
	}

	return s, nil
}

func (edb *EventDb) addStat(event Event) (err error) {
	return edb.addStatMain(event)
}

func fromEvent[T any](eventData interface{}) (*T, bool) {
	if eventData == nil {
		return nil, false
	}

	t, ok := eventData.(T)
	if ok {
		return &t, true
	}

	t2, ok := eventData.(*T)
	if ok {
		return t2, true
	}

	logging.Logger.Error("fromEvent invalid data type",
		zap.Any("expect", reflect.TypeOf(new(T))),
		zap.Any("got", reflect.TypeOf(eventData)))
	return nil, false
}

func setEventData[T any](e *Event, data interface{}) error {
	if data == nil {
		return nil
	}

	_, ok := e.Data.(T)
	if ok {
		e.Data = data
		return nil
	}

	tp, ok := e.Data.(*T)
	if ok {
		*(tp) = data.(T)
		return nil
	}

	return ErrInvalidEventData
}
