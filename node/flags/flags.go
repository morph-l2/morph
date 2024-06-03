package flags

import "github.com/urfave/cli"

const envVarPrefix = "MORPH_NODE_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	Home = cli.StringFlag{
		Name:     "home",
		Usage:    "home directory for morph-node",
		EnvVar:   prefixEnvVar("HOME"),
		Required: false,
	}

	L2EthAddr = cli.StringFlag{
		Name:   "l2.eth",
		Usage:  "Address of L2 Engine JSON-RPC endpoints to use (eth namespace required)",
		EnvVar: prefixEnvVar("L2_ETH_RPC"),
	}

	L2EngineAddr = cli.StringFlag{
		Name:   "l2.engine",
		Usage:  "Address of L2 Engine JSON-RPC endpoints to use (engine namespace required)",
		EnvVar: prefixEnvVar("L2_ENGINE_RPC"),
	}

	L2EngineJWTSecret = cli.StringFlag{
		Name:        "l2.jwt-secret",
		Usage:       "Path to JWT secret key. Keys are 32 bytes, hex encoded in a file. A new key will be generated if left empty.",
		EnvVar:      prefixEnvVar("L2_ENGINE_AUTH"),
		Value:       "",
		Destination: new(string),
	}

	MaxL1MessageNumPerBlock = cli.Uint64Flag{
		Name:   "maxL1MessageNumPerBlock",
		Usage:  "The max number allowed for L1 message type transactions to involve in one block",
		EnvVar: prefixEnvVar("MAX_L1_MESSAGE_NUM_PER_BLOCK"),
	}

	L2CrossDomainMessengerContractAddr = cli.StringFlag{
		Name:   "l2CDMContractAddr",
		Usage:  "L2CrossDomainMessenger contract address",
		EnvVar: prefixEnvVar("L2_CDM_CONTRACT_ADDRESS"),
	}

	L2SequencerAddr = cli.StringFlag{
		Name:   "l2SequencerContractAddr",
		Usage:  "l2sequencer contract address",
		EnvVar: prefixEnvVar("L2_SEQUENCER_CONTRACT_ADDRESS"),
	}

	GovAddr = cli.StringFlag{
		Name:   "govContractAddr",
		Usage:  "gov contract address",
		EnvVar: prefixEnvVar("GOV_CONTRACT_ADDRESS"),
	}

	L1NodeAddr = cli.StringFlag{
		Name:   "l1.rpc",
		Usage:  "Address of L1 User JSON-RPC endpoint to use (eth namespace required)",
		EnvVar: prefixEnvVar("L1_ETH_RPC"),
	}

	L1BeaconAddr = cli.StringFlag{
		Name:   "l1.beaconrpc",
		Usage:  "Address of L1 Beacon JSON-RPC endpoint to use (eth namespace required)",
		EnvVar: prefixEnvVar("L1_ETH_BEACON_RPC"),
	}

	L1ChainID = cli.Uint64Flag{
		Name:   "l1.chain-id",
		Usage:  "L1 Chain ID",
		EnvVar: prefixEnvVar("L1_CHAIN_ID"),
	}

	L1Confirmations = cli.Int64Flag{
		Name:   "l1.confirmations",
		Usage:  "Number of confirmations on L1 needed for finalization",
		EnvVar: prefixEnvVar("L1_CONFIRMATIONS"),
	}

	// Flags for syncer
	SyncDepositContractAddr = cli.StringFlag{
		Name:   "sync.depositContractAddr",
		Usage:  "Contract address deployed on layer one",
		EnvVar: prefixEnvVar("SYNC_DEPOSIT_CONTRACT_ADDRESS"),
	}

	SyncStartHeight = cli.Uint64Flag{
		Name:   "sync.startHeight",
		Usage:  "Block height where syncer start to fetch",
		EnvVar: prefixEnvVar("SYNC_START_HEIGHT"),
	}

	SyncPollInterval = cli.DurationFlag{
		Name:   "sync.pollInterval",
		Usage:  "Frequency at which we query for new L1 messages",
		EnvVar: prefixEnvVar("SYNC_POLL_INTERVAL"),
	}

	SyncLogProgressInterval = cli.DurationFlag{
		Name:   "sync.logProgressInterval",
		Usage:  "frequency at which we log progress",
		EnvVar: prefixEnvVar("SYNC_LOG_PROGRESS_INTERVAL"),
	}

	SyncFetchBlockRange = cli.Uint64Flag{
		Name:   "sync.fetchBlockRange",
		Usage:  "Number of blocks that we collect in a single eth_getLogs query",
		EnvVar: prefixEnvVar("SYNC_FETCH_BLOCK_RANGE"),
	}

	// db options
	DBDataDir = cli.StringFlag{
		Name:   "db.dir",
		Usage:  "Directory of the data",
		EnvVar: prefixEnvVar("DB_DIR"),
	}

	DBNamespace = cli.StringFlag{
		Name:   "db.namespace",
		Usage:  "Database namespace",
		EnvVar: prefixEnvVar("DB_NAMESPACE"),
	}

	DBHandles = cli.IntFlag{
		Name:   "db.handles",
		Usage:  "Database handles",
		EnvVar: prefixEnvVar("DB_HANDLES"),
	}

	DBCache = cli.IntFlag{
		Name:   "db.cache",
		Usage:  "Database cache",
		EnvVar: prefixEnvVar("DB_CACHE"),
	}

	DBFreezer = cli.StringFlag{
		Name:   "db.freezer",
		Usage:  "Database freezer",
		EnvVar: prefixEnvVar("DB_FREEZER"),
	}

	TendermintConfigPath = &cli.StringFlag{
		Name:   "tdm.config",
		Usage:  "Directory of tendermint config file",
		EnvVar: prefixEnvVar("TDM_CONFIG"),
	}

	DevSequencer = &cli.BoolFlag{
		Name:   "dev-sequencer",
		Usage:  "explicitly specify that running as a sequencer. If it enables, the tendermint validator/batch params won't be changed no matter what happens to staking/gov contacts. Only use in dev/test mode",
		EnvVar: prefixEnvVar("DEV_SEQUENCER"),
	}

	MockEnabled = &cli.BoolFlag{
		Name:   "mock",
		Usage:  "Enable mock; If enabled, we start a simulated sequencer to manage the block production, just for dev/test use",
		EnvVar: prefixEnvVar("MOCK_SEQUENCER"),
	}

	ValidatorEnable = cli.BoolFlag{
		Name:   "validator",
		Usage:  "Enable the validator mode",
		EnvVar: prefixEnvVar("VALIDATOR"),
	}

	// validator
	ValidatorPrivateKey = cli.StringFlag{
		Name:   "validator.privateKey",
		Usage:  "Private Key corresponding to SUBSIDY Owner",
		EnvVar: prefixEnvVar("VALIDATOR_PRIVATE_KEY"),
	}

	// derivation
	RollupContractAddress = cli.StringFlag{
		Name:   "derivation.rollupAddress",
		Usage:  "Address of rollup contract",
		EnvVar: prefixEnvVar("ROLLUP_ADDRESS"),
	}

	DerivationStartHeight = cli.Uint64Flag{
		Name:   "derivation.startHeight",
		Usage:  "L1 block height where derivation start to fetch",
		EnvVar: prefixEnvVar("DERIVATION_START_HEIGHT"),
	}

	DerivationPollInterval = cli.DurationFlag{
		Name:   "derivation.pollInterval",
		Usage:  "Frequency at which we query for rollup data",
		EnvVar: prefixEnvVar("DERIVATION_POLL_INTERVAL"),
	}

	DerivationLogProgressInterval = cli.DurationFlag{
		Name:   "derivation.logProgressInterval",
		Usage:  "frequency at which we log progress",
		EnvVar: prefixEnvVar("DERIVATION_LOG_PROGRESS_INTERVAL"),
	}

	DerivationFetchBlockRange = cli.Uint64Flag{
		Name:   "derivation.fetchBlockRange",
		Usage:  "Number of blocks that we collect in a single eth_getLogs query",
		EnvVar: prefixEnvVar("DERIVATION_FETCH_BLOCK_RANGE"),
	}
	// Logger
	LogLevel = &cli.StringFlag{
		Name:   "log.level",
		Usage:  "log level: debug, info(default), error, none",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}

	LogFormat = &cli.StringFlag{
		Name:   "log.format",
		Usage:  "log format: plain(default), json",
		EnvVar: prefixEnvVar("LOG_FORMAT"),
	}

	LogFilename = cli.StringFlag{
		Name:   "log.filename",
		Usage:  "The target file for writing logs, backup log files will be retained in the same directory.",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}
	LogFileMaxSize = cli.IntFlag{
		Name:   "log.maxsize",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes. It is used only when log.filename is provided.",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}
	LogFileMaxAge = cli.IntFlag{
		Name:   "log.maxage",
		Usage:  "The maximum number of days to retain old log files based on the timestamp encoded in their filename. It defaults to 30 days. It is used only when log.filename is provided.",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}
	LogCompress = cli.BoolFlag{
		Name:   "log.compress",
		Usage:  "Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression. It is used only when log.filename is provided.",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	// metrics
	MetricsServerEnable = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Whether or not to run the embedded metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostname = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "The hostname of the metrics server",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPort = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "The port of the metrics server",
		Value:  26660,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}
)

var Flags = []cli.Flag{
	Home,
	L1NodeAddr,
	L1ChainID,
	L1Confirmations,
	L2EthAddr,
	L2EngineAddr,
	L2EngineJWTSecret,
	MaxL1MessageNumPerBlock,
	L2CrossDomainMessengerContractAddr,
	L2SequencerAddr,
	GovAddr,

	// sync optioins
	SyncDepositContractAddr,
	SyncStartHeight,
	SyncPollInterval,
	SyncLogProgressInterval,
	SyncFetchBlockRange,

	// db options
	DBDataDir,
	DBNamespace,
	DBHandles,
	DBCache,
	DBFreezer,

	DevSequencer,
	TendermintConfigPath,
	MockEnabled,
	ValidatorEnable,

	// validator
	ValidatorPrivateKey,

	// derivation
	RollupContractAddress,
	DerivationStartHeight,
	DerivationPollInterval,
	DerivationLogProgressInterval,
	DerivationFetchBlockRange,
	L1BeaconAddr,

	// logger
	LogLevel,
	LogFormat,
	LogFilename,
	LogFileMaxSize,
	LogFileMaxAge,
	LogCompress,

	// metrics
	MetricsServerEnable,
	MetricsPort,
	MetricsHostname,
}
