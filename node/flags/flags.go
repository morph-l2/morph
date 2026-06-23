package flags

import (
	"time"

	"github.com/urfave/cli"
)

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

	DerivationBaseHeight = cli.Uint64Flag{
		Name:   "derivation.baseHeight",
		Usage:  "The starting height of l2 derive, usually the node snapshot or other trusted starting height, before which stateRoot will not be checked",
		EnvVar: prefixEnvVar("DERIVATION_BASE_HEIGHT"),
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

	// L1 Sequencer options
	L1SequencerContractAddr = cli.StringFlag{
		Name:   "l1.sequencerContract",
		Usage:  "L1 Sequencer contract address for signature verification",
		EnvVar: prefixEnvVar("L1_SEQUENCER_CONTRACT"),
	}

	// SequencerUpgradeTime overrides the PBFT->single-sequencer upgrade timestamp.
	// Unit: Unix milliseconds (matches block.Time.UnixMilli()). If unset, --mainnet / --hoodi
	// selects the corresponding built-in network default; without a network flag, the upgrade
	// package's existing default is left unchanged.
	SequencerUpgradeTime = cli.Int64Flag{
		Name:   "sequencerUpgradeTime",
		Usage:  "Unix timestamp (milliseconds) at which consensus switches to sequencer mode",
		EnvVar: prefixEnvVar("SEQUENCER_UPGRADE_TIME"),
	}

	L1SyncLagThreshold = cli.DurationFlag{
		Name:   "l1.syncLagThreshold",
		Usage:  "L1 sync lag threshold for warning logs",
		EnvVar: prefixEnvVar("L1_SYNC_LAG_THRESHOLD"),
		Value:  5 * time.Minute,
	}

	// Sequencer private key for block signing (hex encoded, without 0x prefix)
	SequencerPrivateKey = cli.StringFlag{
		Name:   "sequencer.privateKey",
		Usage:  "Sequencer private key for block signing (hex encoded)",
		EnvVar: prefixEnvVar("SEQUENCER_PRIVATE_KEY"),
	}

	// Vsock address of the Nitro Enclave signer. Accepts both
	// `CID:port` (legacy) and `vsock:CID:port` (matches ops-cli's
	// --addr convention so the same string works in either tool).
	// Mutually exclusive with sequencer.privateKey — with this set,
	// the node never sees the plaintext key, signing happens via the
	// enclave. The signer's EVM address is fetched at startup via
	// GetPubkey.
	SequencerEnclaveSignerAddr = cli.StringFlag{
		Name:   "sequencer.enclaveSignerAddr",
		Usage:  "Vsock address of the enclave signer: `CID:port` or `vsock:CID:port`. Mutually exclusive with sequencer.privateKey.",
		EnvVar: prefixEnvVar("SEQUENCER_ENCLAVE_SIGNER_ADDR"),
	}

	// Sequencer HA flags (all prefixed with ha.)
	SequencerHAEnabled = cli.BoolFlag{
		Name:   "ha.enabled",
		Usage:  "Enable sequencer HA mode (overrides config file).",
		EnvVar: prefixEnvVar("HA_ENABLED"),
	}
	SequencerHAConfig = cli.StringFlag{
		Name:   "ha.config",
		Usage:  "Path to sequencer HA config file (TOML). If not set, HA is disabled.",
		EnvVar: prefixEnvVar("HA_CONFIG"),
	}
	SequencerHABootstrap = cli.BoolFlag{
		Name:   "ha.bootstrap",
		Usage:  "Bootstrap a new Raft cluster as leader (overrides config file).",
		EnvVar: prefixEnvVar("HA_BOOTSTRAP"),
	}
	SequencerHAJoin = cli.StringSliceFlag{
		Name:   "ha.join",
		Usage:  "Management RPC addresses of existing cluster nodes to join (comma-separated, overrides config file).",
		EnvVar: prefixEnvVar("HA_JOIN"),
	}
	SequencerHAServerID = cli.StringFlag{
		Name:   "ha.server-id",
		Usage:  "Unique server ID for this node (overrides config file; defaults to hostname).",
		EnvVar: prefixEnvVar("HA_SERVER_ID"),
	}
	SequencerHAAdvertisedAddr = cli.StringFlag{
		Name:   "ha.advertised-addr",
		Usage:  "Raft advertised address (host:port). Supports hostname (e.g. node-0:9400) or IP. Auto-detected if not set.",
		EnvVar: prefixEnvVar("HA_ADVERTISED_ADDR"),
	}
	SequencerHARPCToken = cli.StringFlag{
		Name:   "ha.rpc-token",
		Usage:  "Auth token for HAKeeper RPC write APIs. If empty, auth is disabled.",
		EnvVar: prefixEnvVar("HA_RPC_TOKEN"),
	}

	MainnetFlag = cli.BoolFlag{
		Name:  "mainnet",
		Usage: "Morph mainnet",
	}

	HoodiFlag = cli.BoolFlag{
		Name:  "hoodi",
		Usage: "Morph Hoodi testnet",
	}

	DerivationConfirmations = cli.Int64Flag{
		Name:   "derivation.confirmations",
		Usage:  "The number of confirmations needed on L1 for finalization. If not set, the default value is l1.confirmations",
		EnvVar: prefixEnvVar("DERIVATION_CONFIRMATIONS"),
	}

	DerivationVerifyMode = cli.StringFlag{
		Name:   "derivation.verify-mode",
		Usage:  `Batch verification mode (SPEC-005 §4.2). "layer1" pulls beacon blob, decodes, and derives blocks via engine. "local" (default) rebuilds blob bytes from local L2 blocks and compares versioned hashes against L1 (no beacon fetch on the happy path); on versioned hash mismatch the verifier is designed to self-heal by pulling the real blob and re-deriving the batch — currently TODO, blocked on EL number-continuity check relaxation in morph-reth/go-ethereum (separate spec). Selected at startup; not switchable at runtime.`,
		EnvVar: prefixEnvVar("DERIVATION_VERIFY_MODE"),
		Value:  "local",
	}

	DerivationReorgCheckDepth = cli.Uint64Flag{
		Name:   "derivation.reorg-check-depth",
		Usage:  "Number of recent L1 blocks to check for reorgs (SPEC-005 §4.7.6). The scan is a no-op when --derivation.confirmations=finalized (L1 finalized doesn't reorg) and load-bearing when set lower; the gate is intentionally absent so behavior is uniform across configs. Default 64.",
		EnvVar: prefixEnvVar("DERIVATION_REORG_CHECK_DEPTH"),
		Value:  64,
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
)

var Flags = []cli.Flag{
	Home,
	L1NodeAddr,
	L1Confirmations,
	L2EthAddr,
	L2EngineAddr,
	L2EngineJWTSecret,
	MaxL1MessageNumPerBlock,
	L2CrossDomainMessengerContractAddr,
	L2SequencerAddr,

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

	// derivation
	RollupContractAddress,
	DerivationStartHeight,
	DerivationBaseHeight,
	DerivationPollInterval,
	DerivationLogProgressInterval,
	DerivationFetchBlockRange,
	DerivationConfirmations,
	DerivationVerifyMode,
	DerivationReorgCheckDepth,
	L1BeaconAddr,

	// L1 Sequencer options
	L1SequencerContractAddr,
	L1SyncLagThreshold,
	SequencerUpgradeTime,
	SequencerPrivateKey,
	SequencerEnclaveSignerAddr,
	SequencerHAEnabled,
	SequencerHAConfig,
	SequencerHABootstrap,
	SequencerHAJoin,
	SequencerHAServerID,
	SequencerHAAdvertisedAddr,
	SequencerHARPCToken,

	MainnetFlag,
	HoodiFlag,

	// logger
	LogLevel,
	LogFormat,
	LogFilename,
	LogFileMaxSize,
	LogFileMaxAge,
	LogCompress,
}
