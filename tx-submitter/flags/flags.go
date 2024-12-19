package flags

import (
	"time"

	"morph-l2/bindings/predeploys"

	"github.com/urfave/cli"
)

const envVarPrefix = "TX_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	BuildEnvFlag = cli.StringFlag{
		Name: "build_env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   prefixEnvVar("BUILD_ENV"),
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1_eth_rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRpcsFlag = cli.StringSliceFlag{
		Name:     "l2_eth_rpcs",
		Usage:    "HTTP provider URLs for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPCS"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:     "l1_private_key",
		Usage:    "The private key to use for sending to the rollup contract",
		EnvVar:   prefixEnvVar("L1_PRIVATE_KEY"),
		Required: false,
	}

	TxTimeoutFlag = cli.DurationFlag{
		Name:     "tx_timeout",
		Usage:    "Timeout for transaction submission",
		Value:    10,
		EnvVar:   prefixEnvVar("TX_TIMEOUT"),
		Required: true,
	}

	// L1 Address
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup_address",
		Usage:    "Address of the rollup contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}
	L1StakingAddressFlag = cli.StringFlag{
		Name:     "l1_staking_address",
		Usage:    "Address of the staking contract",
		Required: true,
		EnvVar:   prefixEnvVar("L1_STAKING_ADDRESS"),
	}

	// finalize flags
	FinalizeFlag = cli.BoolFlag{
		Name:     "finalize",
		Usage:    "Enable finalize",
		EnvVar:   prefixEnvVar("FINALIZE"),
		Required: true,
	}

	// decentralize config
	PriorityRollupFlag = cli.BoolFlag{
		Name:     "priority_rollup",
		Usage:    "Enable priority rollup",
		EnvVar:   prefixEnvVar("PRIORITY_ROLLUP"),
		Required: true,
	}

	// L2 contract
	L2SequencerAddressFlag = cli.StringFlag{
		Name:     "l2_sequencer_address",
		Usage:    "Address of the sequencer contract",
		Required: false,
		EnvVar:   prefixEnvVar("L2_SEQUENCER_ADDRESS"),
		Value:    predeploys.Sequencer,
	}
	L2GovAddressFlag = cli.StringFlag{
		Name:     "l2_gov_address",
		Usage:    "Address of the gov contract",
		Required: false,
		EnvVar:   prefixEnvVar("L2_GOV_ADDRESS"),
		Value:    predeploys.Gov,
	}

	/* Optional Flags */
	LogLevelFlag = cli.StringFlag{
		Name:   "log_level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	// metrics
	MetricsServerEnable = cli.BoolFlag{
		Name:   "metrics_server_enable",
		Usage:  "Enable metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostname = cli.StringFlag{
		Name:   "metrics_hostname",
		Usage:  "Hostname at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPort = cli.Uint64Flag{
		Name:   "metrics_port",
		Usage:  "Port at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}

	// tx fee limit
	TxFeeLimitFlag = cli.Uint64Flag{
		Name:     "tx_fee_limit",
		Usage:    "The maximum fee for a transaction",
		Value:    5e17, //0.5eth
		EnvVar:   prefixEnvVar("TX_FEE_LIMIT"),
		Required: true,
	}

	// log to file
	LogFilename = cli.StringFlag{
		Name:   "log_filename",
		Usage:  "The target file for writing logs, backup log files will be retained in the same directory.",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}
	LogFileMaxSize = cli.IntFlag{
		Name:   "log_file_max_size",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes. It is used only when log.filename is provided.",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}
	LogFileMaxAge = cli.IntFlag{
		Name:   "log_file_max_age",
		Usage:  "The maximum number of days to retain old log files based on the timestamp encoded in their filename. It defaults to 30 days. It is used only when log.filename is provided.",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}
	LogCompress = cli.BoolFlag{
		Name:   "log_compress",
		Usage:  "Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression. It is used only when log.filename is provided.",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	// rollup interval
	RollupInterval = cli.DurationFlag{
		Name:   "rollup_interval",
		Usage:  "Interval for rollup",
		Value:  500 * time.Millisecond,
		EnvVar: prefixEnvVar("ROLLUP_INTERVAL"),
	}
	// finalize interval
	FinalizeInterval = cli.DurationFlag{
		Name:   "FINALIZE_INTERVAL",
		Usage:  "Interval for finalize",
		Value:  2 * time.Second,
		EnvVar: prefixEnvVar("FINALIZE_INTERVAL"),
	}
	// tx process interval
	TxProcessInterval = cli.DurationFlag{
		Name:   "tx_process_interval",
		Usage:  "Interval for tx process",
		Value:  2 * time.Second,
		EnvVar: prefixEnvVar("TX_PROCESS_INTERVAL"),
	}

	// rollup tx gas base
	RollupTxGasBase = cli.Uint64Flag{
		Name:   "rollup_tx_gas_base",
		Usage:  "The base fee for a rollup transaction",
		Value:  530000,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_BASE"),
	}
	// rollup tx gas per l1msg
	RollupTxGasPerL1Msg = cli.Uint64Flag{
		Name:   "rollup_tx_gas_per_l1_msg",
		Usage:  "The gas cost for each L1 message included in a rollup transaction",
		Value:  4200,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_PER_L1_MSG"),
	}

	GasLimitBuffer = cli.Uint64Flag{
		Name:   "gas_limit_buffer",
		Usage:  "The gas limit buffer for a transaction",
		Value:  20, // add 20%
		EnvVar: prefixEnvVar("GAS_LIMIT_BUFFER"),
	}

	// journal path
	JournalFlag = cli.StringFlag{
		Name:   "journal_file_path",
		Usage:  "The path of the journal file",
		EnvVar: prefixEnvVar("JOURNAL_FILE_PATH"),
		Value:  "journal.rlp",
	}
	// listener processed block record path
	StakingEventStoreFileFlag = cli.StringFlag{
		Name:   "staking_event_store_filename",
		Usage:  "The file name of the storage",
		EnvVar: prefixEnvVar("STAKING_EVENT_STORE_FILENAME"),
		Value:  "StakingEventStore.json",
	}

	TipFeeBumpFlag = cli.Uint64Flag{
		Name:   "TIP_FEE_BUMP",
		Usage:  "The fee bump for tip",
		Value:  120, //bumpTip = tip * TipFeeBump/100
		EnvVar: prefixEnvVar("TIP_FEE_BUMP"),
	}
	MaxTipFlag = cli.Uint64Flag{
		Name:   "max_tip",
		Usage:  "The maximum tip for a transaction",
		Value:  10e9, //10gwei
		EnvVar: prefixEnvVar("MAX_TIP"),
	}
	MinTipFlag = cli.Uint64Flag{
		Name:   "min_tip",
		Usage:  "The minimum tip for a transaction",
		Value:  5e8, //0.5gwei
		EnvVar: prefixEnvVar("MIN_TIP"),
	}
	MaxBaseFeeFlag = cli.Uint64Flag{
		Name:   "max_base_fee",
		Usage:  "The maximum base fee for a transaction",
		Value:  100e9, //100gwei
		EnvVar: prefixEnvVar("MAX_BASE_FEE"),
	}

	MaxTxsInPendingPoolFlag = cli.Uint64Flag{
		Name:   "max_txs_in_pending_pool",
		Usage:  "The maximum number of transactions in the pending pool",
		Value:  12,
		EnvVar: prefixEnvVar("MAX_TXS_IN_PENDING_POOL"),
	}

	// external sign
	ExternalSign = cli.BoolFlag{
		Name:   "external_sign",
		Usage:  "Enable external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN"),
	}

	// address
	ExternalSignAddress = cli.StringFlag{
		Name:   "external_sign_address",
		Usage:  "The address of the external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_ADDRESS"),
	}
	// appid
	ExternalSignAppid = cli.StringFlag{
		Name:   "external_sign_appid",
		Usage:  "The appid of the external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_APPID"),
	}
	// chain
	ExternalSignChain = cli.StringFlag{
		Name:   "external_sign_chain",
		Usage:  "The chain of the external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_CHAIN"),
	}
	// url
	ExternalSignUrl = cli.StringFlag{
		Name:   "external_sign_url",
		Usage:  "The url of the external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_URL"),
	}
	ExternalSignRsaPriv = cli.StringFlag{
		Name:   "external_rsa_priv",
		Usage:  "The rsa private key of the external sign",
		EnvVar: "SEQUENCER_EXTERNAL_SIGN_RSA_PRIV", // use sequencer rsa from xxx
	}
	RoughEstimateGasFlag = cli.BoolFlag{
		Name:   "rough_estimate_gas",
		Usage:  "Whether to use rough estimate gas",
		EnvVar: prefixEnvVar("ROUGH_ESTIMATE_GAS"),
	}

	RotatorBufferFlag = cli.Int64Flag{
		Name:   "rotator_buffer",
		Usage:  "rotation interval buffer",
		Value:  15,
		EnvVar: prefixEnvVar("ROTATOR_BUFFER"),
	}

	// l1 staking deployed blocknum
	L1StakingDeployedBlocknumFlag = cli.Uint64Flag{
		Name:     "l1_staking_deployed_blocknum",
		Usage:    "The deployed block number of L1Staking",
		EnvVar:   prefixEnvVar("L1_STAKING_DEPLOYED_BLOCKNUM"),
		Required: true,
	}

	// event indexer
	EventIndexStepFlag = cli.Uint64Flag{
		Name:   "event_index_step",
		Usage:  "The step size for event indexing",
		Value:  100,
		EnvVar: prefixEnvVar("EVENT_INDEX_STEP"),
	}
	LeveldbPathNameFlag = cli.StringFlag{
		Name:   "leveldb_path_name",
		Usage:  "The path name of the leveldb",
		EnvVar: prefixEnvVar("LEVELDB_PATH_NAME"),
		Value:  "submitter-leveldb",
	}

	// l1 block not incremented threshold
	BlockNotIncreasedThreshold = cli.Int64Flag{
		Name:   "block_not_increased_threshold",
		Usage:  "The threshold for block not incremented",
		Value:  5,
		EnvVar: prefixEnvVar("BLOCK_NOT_INCREASED_THRESHOLD"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	L1EthRpcFlag,
	L2EthRpcsFlag,
	RollupAddressFlag,
	TxTimeoutFlag,
	FinalizeFlag,
	PriorityRollupFlag,
	TxFeeLimitFlag,
	L1StakingAddressFlag,
	L1StakingDeployedBlocknumFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	MetricsServerEnable,
	MetricsHostname,
	MetricsPort,

	LogFilename,
	LogFileMaxSize,
	LogFileMaxAge,
	LogCompress,

	RollupInterval,
	FinalizeInterval,
	TxProcessInterval,

	RollupTxGasBase,
	RollupTxGasPerL1Msg,
	GasLimitBuffer,

	JournalFlag,

	PrivateKeyFlag,
	L2SequencerAddressFlag,
	L2GovAddressFlag,
	TipFeeBumpFlag,
	MaxTxsInPendingPoolFlag,

	// external sign
	ExternalSign,
	ExternalSignAddress,
	ExternalSignAppid,
	ExternalSignChain,
	ExternalSignUrl,
	ExternalSignRsaPriv,
	RoughEstimateGasFlag,
	RotatorBufferFlag,
	StakingEventStoreFileFlag,
	EventIndexStepFlag,
	LeveldbPathNameFlag,
	BlockNotIncreasedThreshold,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
