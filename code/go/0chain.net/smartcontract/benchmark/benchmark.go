package benchmark

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"0chain.net/core/common"

	"0chain.net/smartcontract/dbs/event"

	"0chain.net/chaincore/chain/state"
	"0chain.net/chaincore/transaction"
	"0chain.net/core/encryption"
	"github.com/spf13/viper"
)

type Source int

const (
	Storage = Source(iota)
	StorageRest
	Miner
	MinerRest
	Faucet
	FaucetRest
	Vesting
	VestingRest
	MultiSig
	ZCNSCBridge
	ZCNSCBridgeRest
	Control
	EventDatabase
	EventDatabaseEvents
	EventDatabaseAggregates
	NumberOdfBenchmarkSources
)

var (
	SourceNames = []string{
		"storage",
		"storage_rest",
		"miner",
		"miner_rest",
		"faucet",
		"faucet_rest",
		"vesting",
		"vesting_rest",
		"multi_sig",
		"zcnscbridge",
		"zcnscbridge_rest",
		"control",
		"event_db",
		"event_db_events",
		"event_db_aggregates",
	}

	SourceCode = map[string]Source{
		SourceNames[Storage]:                 Storage,
		SourceNames[StorageRest]:             StorageRest,
		SourceNames[Miner]:                   Miner,
		SourceNames[MinerRest]:               MinerRest,
		SourceNames[Faucet]:                  Faucet,
		SourceNames[FaucetRest]:              FaucetRest,
		SourceNames[Vesting]:                 Vesting,
		SourceNames[VestingRest]:             VestingRest,
		SourceNames[MultiSig]:                MultiSig,
		SourceNames[ZCNSCBridge]:             ZCNSCBridge,
		SourceNames[ZCNSCBridgeRest]:         ZCNSCBridgeRest,
		SourceNames[Control]:                 Control,
		SourceNames[EventDatabase]:           EventDatabase,
		SourceNames[EventDatabaseEvents]:     EventDatabaseEvents,
		SourceNames[EventDatabaseAggregates]: EventDatabaseAggregates,
	}
)

type SimulatorParameter int

const (
	SimulationNumClients SimulatorParameter = iota
	SimulationActiveNumClients
	SimulationNumMiners
	SimulationNumActiveMiners
	SimulationNumSharders
	SimulationNumActiveSharders
	SimulationNumAllocations
	SimulationNumBlobbersPerAllocation
	SimulationNumBlobbers
	SimulationNumAllocationPayerPools

	SimulationNumAllocationPayer
	SimulationNumBlobberDelegates
	SimulationNumValidators
	SimulationNumFreeStorageAssigners
	SimulationNumMinerDelegates

	SimulationNumSharderDelegates
	SimulationNumVestingDestinationsClient
	SimulationNumChallengesBlobber
	SimulationNumAuthorizers
	SimulationNumRewardPartitionBlobber

	SimulationNumBlocks
	SimulationNumTransactionsPerBlock
	SimulationNumWriteRedeemAllocation
	SimulationNumReadMarkersAllocation
	SimulationNumRoundsBetweenWrites
	SimulationNumRoundsBetweenReads

	NumberSimulationParameters
)

const (
	Simulation    = "simulation."
	Options       = "options."
	Internal      = "internal."
	SmartContract = "smart_contracts."
	MinerSc       = "minersc."
	StorageSc     = "storagesc."
	FaucetSc      = "faucetsc."
	VestingSc     = "vestingsc."
	ZcnSc         = "zcnsc."
	DbsEvents     = "dbs.events."
	DbSettings    = "dbs.settings."

	BlockReward = "block_reward."

	Fas = "free_allocation_settings."

	AvailableKeys           = Internal + "available_keys"
	InternalT               = Internal + "t"
	InternalSignatureScheme = Internal + "signature_scheme"
	StartTokens             = Internal + "start_tokens"
	Bad                     = Internal + "bad"
	Worry                   = Internal + "worry"
	Satisfactory            = Internal + "satisfactory"
	TimeUnit                = Internal + "time_unit"
	Colour                  = Internal + "colour"
	ControlM                = Internal + "control_m"
	ControlN                = Internal + "control_n"
	MptRoot                 = Internal + "mpt_root"
	ShowOutput              = Internal + "show_output"
	MptCreationTime         = Internal + "mpt_creation_time"
	BenchDataListLength     = Internal + "bench_data_list_length"

	OptionVerbose                  = Options + "verbose"
	OptionTestSuites               = Options + "test_suites"
	OptionOmittedTests             = Options + "omitted_tests"
	OptionLoadPath                 = Options + "load_path"
	OptionSavePath                 = Options + "save_path"
	OptionsLoadConcurrency         = Options + "load_concurrency"
	OptionsSmartContractBenchmarks = Options + "smart_contract_benchmarks"
	OptionsQueryBenchmarks         = Options + "query_benchmarks"
	OptionsEventDatabaseBenchmarks = Options + "event_database_benchmarks"
	OptionsSmartContractEventFile  = Options + "smart_contract_event_file"
	OptionsEventDatabaseEventFile  = Options + "event_database_event_file"
	OptionVerifyBurnedTokens       = Options + "verify_burned_tokens"

	MinerMOwner       = SmartContract + MinerSc + "owner_id"
	MinerMaxDelegates = SmartContract + MinerSc + "max_delegates"
	MinerMaxCharge    = SmartContract + MinerSc + "max_charge"
	MinerMinStake     = SmartContract + MinerSc + "min_stake"
	MinerMaxStake     = SmartContract + MinerSc + "max_stake"

	StorageOwner                        = SmartContract + StorageSc + "owner_id"
	StorageMinAllocSize                 = SmartContract + StorageSc + "min_alloc_size"
	StorageMaxReadPrice                 = SmartContract + StorageSc + "max_read_price"
	StorageMaxWritePrice                = SmartContract + StorageSc + "max_write_price"
	StorageMaxFileSize                  = SmartContract + StorageSc + "max_file_size"
	StorageMaxChallengeCompletionRounds = SmartContract + StorageSc + "max_challenge_completion_rounds"
	StorageMinBlobberCapacity           = SmartContract + StorageSc + "min_blobber_capacity"
	StorageMaxCharge                    = SmartContract + StorageSc + "max_charge"
	StorageMinStake                     = SmartContract + StorageSc + "min_stake"
	StorageMaxStake                     = SmartContract + StorageSc + "max_stake"
	StorageMinStakePerDelegates         = SmartContract + StorageSc + "min_stake_per_delegate"
	StorageMaxDelegates                 = SmartContract + StorageSc + "max_delegates"
	StorageDiverseBlobbers              = SmartContract + StorageSc + "diverse_blobbers"
	StorageReadPoolMinLock              = SmartContract + StorageSc + "readpool.min_lock"
	StorageWritePoolMinLock             = SmartContract + StorageSc + "writepool.min_lock"
	StorageStakePoolMinLock             = SmartContract + StorageSc + "stakepool.min_lock"
	StorageChallengeEnabled             = SmartContract + StorageSc + "challenge_enabled"
	StorageChallengeGenerationGap       = SmartContract + StorageSc + "challenge_generation_gap"
	StorageMaxTotalFreeAllocation       = SmartContract + StorageSc + "max_total_free_allocation"
	StorageMaxIndividualFreeAllocation  = SmartContract + StorageSc + "max_individual_free_allocation"
	StorageFasDataShards                = SmartContract + StorageSc + Fas + "data_shards"
	StorageFasParityShards              = SmartContract + StorageSc + Fas + "parity_shards"
	StorageFasSize                      = SmartContract + StorageSc + Fas + "size"
	StorageFasDuration                  = SmartContract + StorageSc + Fas + "duration"
	StorageFasReadPriceMin              = SmartContract + StorageSc + Fas + "read_price_range.min"
	StorageFasReadPriceMax              = SmartContract + StorageSc + Fas + "read_price_range.max"
	StorageFasWritePriceMin             = SmartContract + StorageSc + Fas + "write_price_range.min"
	StorageFasWritePriceMax             = SmartContract + StorageSc + Fas + "write_price_range.max"
	StorageFasReadPoolFraction          = SmartContract + StorageSc + Fas + "read_pool_fraction"
	StorageValidatorsPerChallenge       = SmartContract + StorageSc + "validators_per_challenge"
	StorageMaxBlobberSelectForChallenge = SmartContract + StorageSc + "max_blobber_select_for_challenge"
	StorageMaxBlobbersPerAllocation     = SmartContract + StorageSc + "max_blobbers_per_allocation"
	StorageNumValidatorsRewarded        = SmartContract + StorageSc + "num_validators_rewarded"

	StorageBlockReward                = SmartContract + StorageSc + BlockReward + "block_reward"
	StorageBlockRewardTriggerPeriod   = SmartContract + StorageSc + BlockReward + "trigger_period"
	StorageBlockRewardChangePeriod    = SmartContract + StorageSc + BlockReward + "block_reward_change_period"
	StorageBlockRewardChangeRatio     = SmartContract + StorageSc + BlockReward + "block_reward_change_ratio"
	StorageBlockRewardBlobberRatio    = SmartContract + StorageSc + BlockReward + "blobber_ratio"
	StorageBlockRewardMinerRatio      = SmartContract + StorageSc + BlockReward + "miner_ratio"
	StorageBlockRewardSharderRatio    = SmartContract + StorageSc + BlockReward + "sharder_ratio"
	StorageBlockRewardQualifyingStake = SmartContract + StorageSc + BlockReward + "qualifying_stake"
	StorageBlockRewardGammaAlpha      = SmartContract + StorageSc + BlockReward + "gamma.alpha"
	StorageBlockRewardGammaA          = SmartContract + StorageSc + BlockReward + "gamma.a"
	StorageBlockRewardGammaB          = SmartContract + StorageSc + BlockReward + "gamma.b"
	StorageBlockRewardZetaI           = SmartContract + StorageSc + BlockReward + "zeta.i"
	StorageBlockRewardZetaK           = SmartContract + StorageSc + BlockReward + "zeta.k"
	StorageBlockRewardZetaMu          = SmartContract + StorageSc + BlockReward + "zeta.mu"

	VestingPoolOwner            = SmartContract + VestingSc + "owner_id"
	VestingMinLock              = SmartContract + VestingSc + "min_lock"
	VestingMaxDestinations      = SmartContract + VestingSc + "max_destinations"
	VestingMinDuration          = SmartContract + VestingSc + "min_duration"
	VestingMaxDuration          = SmartContract + VestingSc + "max_duration"
	VestingMaxDescriptionLength = SmartContract + VestingSc + "max_description_length"

	FaucetOwner = SmartContract + FaucetSc + "owner_id"

	ZcnOwner              = SmartContract + ZcnSc + "owner_id"
	ZcnMinMintAmount      = SmartContract + ZcnSc + "min_mint"
	ZcnMinBurnAmount      = SmartContract + ZcnSc + "min_burn"
	ZcnMinStakeAmount     = SmartContract + ZcnSc + "min_stake"
	ZcnMaxStakeAmount     = SmartContract + ZcnSc + "max_stake"
	ZcnMinLockAmount      = SmartContract + ZcnSc + "min_lock"
	ZcnMaxFee             = SmartContract + ZcnSc + "max_fee"
	ZcnPercentAuthorizers = SmartContract + ZcnSc + "percent_authorizers"
	ZcnMinAuthorizers     = SmartContract + ZcnSc + "min_authorizers"
	ZcnBurnAddress        = SmartContract + ZcnSc + "burn_address"
	ZcnMaxDelegates       = SmartContract + ZcnSc + "max_delegates"
	HealthCheckPeriod     = SmartContract + ZcnSc + "health_check_period"

	EventDbEnabled         = DbsEvents + "enabled"
	EventDbName            = DbsEvents + "name"
	EventDbUser            = DbsEvents + "user"
	EventDbPassword        = DbsEvents + "password"
	EventDbHost            = DbsEvents + "host"
	EventDbPort            = DbsEvents + "port"
	EventDbMaxIdleConns    = DbsEvents + "max_idle_conns"
	EventDbOpenConns       = DbsEvents + "max_open_conns"
	EventDbConnMaxLifetime = DbsEvents + "conn_max_lifetime"
	EventDbSlowTableSpace  = DbsEvents + "slowtablespace"

	EventDbDebug                          = DbSettings + "debug"
	EventDbAggregatePeriod                = DbSettings + "aggregate_period"
	EventDbPartitionChangePeriod          = DbSettings + "partition_change_period"
	EventDbPartitionKeepCount             = DbSettings + "partition_keep_count"
	EventDbPermanentPartitionChangePeriod = DbSettings + "permanent_partition_change_period"
	EventDbPermanentPartitionKeepCount    = DbSettings + "permanent_partition_keep_count"
	EventDbPageLimit                      = DbSettings + "page_limit"
)

func (s Source) String() string {
	i := int(s)
	switch {
	case i <= int(NumberOdfBenchmarkSources):
		return SourceNames[i]
	default:
		return strconv.Itoa(i)
	}
}

var parameterName []string

func init() {
	initParameterNames()
	initParameterConstants()
}

func initParameterNames() {
	parameterName = make([]string, NumberSimulationParameters)
	parameterName[SimulationNumClients] = "num_clients"
	parameterName[SimulationActiveNumClients] = "num_active_clients"
	parameterName[SimulationNumMiners] = "num_miners"
	parameterName[SimulationNumActiveMiners] = "num_active_miners"
	parameterName[SimulationNumSharders] = "nun_sharders"
	parameterName[SimulationNumActiveSharders] = "nun_active_sharders"
	parameterName[SimulationNumAllocations] = "num_allocations"
	parameterName[SimulationNumBlobbersPerAllocation] = "num_blobbers_per_Allocation"
	parameterName[SimulationNumBlobbers] = "num_blobbers"
	parameterName[SimulationNumAllocationPayerPools] = "num_allocation_payers_pools"

	parameterName[SimulationNumAllocationPayer] = "num_allocation_payers"
	parameterName[SimulationNumBlobberDelegates] = "num_blobber_delegates"
	parameterName[SimulationNumValidators] = "num_validators"
	parameterName[SimulationNumFreeStorageAssigners] = "num_free_storage_assigners"
	parameterName[SimulationNumMinerDelegates] = "num_miner_delegates"

	parameterName[SimulationNumSharderDelegates] = "num_sharder_delegates"
	parameterName[SimulationNumVestingDestinationsClient] = "num_vesting_destinations_client"
	parameterName[SimulationNumChallengesBlobber] = "num_challenges_blobber"
	parameterName[SimulationNumAuthorizers] = "num_authorizers"
	parameterName[SimulationNumRewardPartitionBlobber] = "num_reward_partition_blobber"

	parameterName[SimulationNumBlocks] = "num_blocks"
	parameterName[SimulationNumTransactionsPerBlock] = "num_transactions_per_block"
	parameterName[SimulationNumWriteRedeemAllocation] = "num_write_redeem_allocation"
	parameterName[SimulationNumReadMarkersAllocation] = "num_read_markers_allocation"
	parameterName[SimulationNumRoundsBetweenWrites] = "num_rounds_between_writes"
	parameterName[SimulationNumRoundsBetweenReads] = "num_rounds_between_reads"
}

func (w SimulatorParameter) String() string {
	if w < 0 || w >= NumberSimulationParameters {
		return fmt.Sprintf("unrecognised paramter %d", w)
	}

	return parameterName[w]
}

var (
	NumClients                   string
	NumActiveClients             string
	NumMiners                    string
	NumActiveMiners              string
	NumSharders                  string
	NumActiveSharders            string
	NumAllocations               string
	NumBlobbersPerAllocation     string
	NumBlobbers                  string
	NumAllocationPayerPools      string
	NumAllocationPayer           string
	NumBlobberDelegates          string
	NumValidators                string
	NumFreeStorageAssigners      string
	NumMinerDelegates            string
	NumSharderDelegates          string
	NumVestingDestinationsClient string
	NumChallengesBlobber         string
	NumAuthorizers               string
	NumRewardPartitionBlobber    string
	NumBlocks                    string
	NumWriteRedeemAllocation     string
	NumReadMarkersAllocation     string
	NumRoundsBetweenWrites       string
	NumRoundsBetweenReads        string
	NumTransactionPerBlock       string
)

func initParameterConstants() {
	NumClients = Simulation + SimulationNumClients.String()
	NumActiveClients = Simulation + SimulationActiveNumClients.String()
	NumMiners = Simulation + SimulationNumMiners.String()
	NumActiveMiners = Simulation + SimulationNumActiveMiners.String()
	NumSharders = Simulation + SimulationNumSharders.String()
	NumActiveSharders = Simulation + SimulationNumActiveSharders.String()
	NumAllocations = Simulation + SimulationNumAllocations.String()
	NumBlobbersPerAllocation = Simulation + SimulationNumBlobbersPerAllocation.String()
	NumBlobbers = Simulation + SimulationNumBlobbers.String()
	NumAllocationPayerPools = Simulation + SimulationNumAllocationPayerPools.String()
	NumAllocationPayer = Simulation + SimulationNumAllocationPayer.String()
	NumBlobberDelegates = Simulation + SimulationNumBlobberDelegates.String()
	NumValidators = Simulation + SimulationNumValidators.String()
	NumFreeStorageAssigners = Simulation + SimulationNumFreeStorageAssigners.String()
	NumMinerDelegates = Simulation + SimulationNumMinerDelegates.String()
	NumSharderDelegates = Simulation + SimulationNumSharderDelegates.String()
	NumVestingDestinationsClient = Simulation + SimulationNumVestingDestinationsClient.String()
	NumChallengesBlobber = Simulation + SimulationNumChallengesBlobber.String()
	NumAuthorizers = Simulation + SimulationNumAuthorizers.String()
	NumRewardPartitionBlobber = Simulation + SimulationNumRewardPartitionBlobber.String()
	NumBlocks = Simulation + SimulationNumBlocks.String()
	NumWriteRedeemAllocation = Simulation + SimulationNumWriteRedeemAllocation.String()
	NumReadMarkersAllocation = Simulation + SimulationNumReadMarkersAllocation.String()
	NumRoundsBetweenWrites = Simulation + SimulationNumRoundsBetweenWrites.String()
	NumRoundsBetweenReads = Simulation + SimulationNumRoundsBetweenReads.String()
	NumTransactionPerBlock = Simulation + SimulationNumTransactionsPerBlock.String()

}

type BenchTestI interface {
	Name() string
	Transaction() *transaction.Transaction
	Run(state.TimedQueryStateContext, *testing.B) error
}

type WithTimings interface {
	Timings() map[string]time.Duration
}

type SignatureScheme interface {
	encryption.SignatureScheme
	SetPrivateKey(privateKey string)
	GetPrivateKey() string
}

type TestSuite struct {
	Source     Source
	Benchmarks []BenchTestI
	ReadOnly   bool
}

func (ts *TestSuite) RemoveBenchmarks(listToRemove []string) {
	if len(ts.Benchmarks) == 0 {
		return
	}
	var name = ts.Benchmarks[0].Name()
	var prefix = name[:strings.IndexByte(name, '.')]
	for _, testName := range listToRemove {
		if len(testName) > len(prefix) && prefix == testName[:len(prefix)] {
			ts.removeBenchmark(testName)
		}
		if len(ts.Benchmarks) == 0 {
			return
		}
	}
}

func (ts *TestSuite) removeBenchmark(benchToRemove string) bool {
	for i, bks := range ts.Benchmarks {
		if bks.Name() == benchToRemove {
			ts.Benchmarks[i] = ts.Benchmarks[len(ts.Benchmarks)-1]
			ts.Benchmarks = ts.Benchmarks[:len(ts.Benchmarks)-1]
			return true
		}
	}
	return false
}

type BenchData struct {
	BenchDataMpt
	EventDb *event.EventDb
}

func (bd *BenchData) Encode() (b []byte) {
	var err error
	if b, err = json.Marshal(bd); err != nil {
		log.Fatal(err)
	}
	return
}

// Decode from []byte
func (bd *BenchData) Decode(input []byte) error {
	return json.Unmarshal(input, bd)
}

var MockBenchData = BenchData{
	BenchDataMpt: BenchDataMpt{
		Clients:              make([]string, 100),
		PublicKeys:           make([]string, 100),
		PrivateKeys:          make([]string, 100),
		Miners:               make([]string, 100),
		Sharders:             make([]string, 100),
		SharderKeys:          make([]string, 100),
		ValidatorIds:         make([]string, 100),
		ValidatorPrivateKeys: make([]string, 100),
		ValidatorPublicKeys:  make([]string, 100),
		Now:                  common.Now(),
	},
}

func GetOldestPermanentAggregateRound() int64 {
	var (
		permanentPeriod = viper.GetInt(EventDbPermanentPartitionChangePeriod)
		permanentKeep   = viper.GetInt(EventDbPermanentPartitionKeepCount)
		blocks          = viper.GetInt(NumBlocks)
		oldestRoundKept = int64((blocks/permanentPeriod - permanentKeep + 1) * permanentPeriod)
	)
	if oldestRoundKept < 0 {
		return 0
	} else {
		return oldestRoundKept
	}
}

func GetOldestAggregateRound() int64 {
	var (
		period          = viper.GetInt(EventDbPartitionChangePeriod)
		keep            = viper.GetInt(EventDbPartitionKeepCount)
		blocks          = viper.GetInt(NumBlocks)
		oldestRoundKept = int64((blocks/period - keep + 1) * period)
	)
	if oldestRoundKept < 0 {
		return 0
	} else {
		return oldestRoundKept
	}
}
