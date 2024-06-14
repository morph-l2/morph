package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "TX_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	BuildEnvFlag = cli.StringFlag{
		Name: "BUILD_ENV",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   prefixEnvVar("BUILD_ENV"),
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "L1_ETH_RPC",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRpcsFlag = cli.StringSliceFlag{
		Name:     "L2_ETH_RPCS",
		Usage:    "HTTP provider URLs for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPCS"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:     "L1_PRIVATE_KEY",
		Usage:    "The private key to use for sending to the rollup contract",
		EnvVar:   prefixEnvVar("L1_PRIVATE_KEY"),
		Required: true,
	}

	TxTimeoutFlag = cli.DurationFlag{
		Name:     "TX_TIMEOUT",
		Usage:    "Timeout for transaction submission",
		Value:    10,
		EnvVar:   prefixEnvVar("TX_TIMEOUT"),
		Required: true,
	}

	// L1 Address
	RollupAddressFlag = cli.StringFlag{
		Name:     "ROLLUP_ADDRESS",
		Usage:    "Address of the rollup contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}
	L1StakingAddressFlag = cli.StringFlag{
		Name:     "L1_STAKING_ADDRESS",
		Usage:    "Address of the staking contract",
		Required: true,
		EnvVar:   prefixEnvVar("L1_STAKING_ADDRESS"),
	}

	// finalize flags
	FinalizeFlag = cli.BoolFlag{
		Name:     "FINALIZE",
		Usage:    "Enable finalize",
		EnvVar:   prefixEnvVar("FINALIZE"),
		Required: true,
	}

	// decentralize config
	PriorityRollupFlag = cli.BoolFlag{
		Name:     "PRIORITY_ROLLUP",
		Usage:    "Enable priority rollup",
		EnvVar:   prefixEnvVar("PRIORITY_ROLLUP"),
		Required: true,
	}

	// L2 contract
	SubmitterAddressFlag = cli.StringFlag{
		Name:     "L2_SUBMITTER_ADDRESS",
		Usage:    "Address of the submitter contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SUBMITTER_ADDRESS"),
	}
	L2SequencerAddressFlag = cli.StringFlag{
		Name:     "L2_SEQUENCER_ADDRESS",
		Usage:    "Address of the sequencer contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SEQUENCER_ADDRESS"),
	}
	L2GovAddressFlag = cli.StringFlag{
		Name:     "L2_GOV_ADDRESS",
		Usage:    "Address of the gov contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_GOV_ADDRESS"),
	}

	/* Optional Flags */
	LogLevelFlag = cli.StringFlag{
		Name:   "LOG_LEVEL",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	// metrics
	MetricsServerEnable = cli.BoolFlag{
		Name:   "METRICS_SERVER_ENABLE",
		Usage:  "Enable metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostname = cli.StringFlag{
		Name:   "METRICS_HOSTNAME",
		Usage:  "Hostname at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPort = cli.Uint64Flag{
		Name:   "METRICS_PORT",
		Usage:  "Port at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}

	// tx fee limit
	TxFeeLimitFlag = cli.Uint64Flag{
		Name:     "TX_FEE_LIMIT",
		Usage:    "The maximum fee for a transaction",
		Value:    5e17, //0.5eth
		EnvVar:   prefixEnvVar("TX_FEE_LIMIT"),
		Required: true,
	}

	// log to file
	LogFilename = cli.StringFlag{
		Name:   "LOG_FILENAME",
		Usage:  "The target file for writing logs, backup log files will be retained in the same directory.",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}
	LogFileMaxSize = cli.IntFlag{
		Name:   "LOG_FILE_MAX_SIZE",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes. It is used only when log.filename is provided.",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}
	LogFileMaxAge = cli.IntFlag{
		Name:   "LOG_FILE_MAX_AGE",
		Usage:  "The maximum number of days to retain old log files based on the timestamp encoded in their filename. It defaults to 30 days. It is used only when log.filename is provided.",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}
	LogCompress = cli.BoolFlag{
		Name:   "LOG_COMPRESS",
		Usage:  "Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression. It is used only when log.filename is provided.",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	// rollup interval
	RollupInterval = cli.DurationFlag{
		Name:   "ROLLUP_INTERVAL",
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
		Name:   "TX_PROCESS_INTERVAL",
		Usage:  "Interval for tx process",
		Value:  2 * time.Second,
		EnvVar: prefixEnvVar("TX_PROCESS_INTERVAL"),
	}

	// rollup tx gas base
	RollupTxGasBase = cli.Uint64Flag{
		Name:   "ROLLUP_TX_GAS_BASE",
		Usage:  "The base fee for a rollup transaction",
		Value:  400000,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_BASE"),
	}
	// rollup tx gas per l1msg
	RollupTxGasPerL1Msg = cli.Uint64Flag{
		Name:   "ROLLUP_TX_GAS_PER_L1_MSG",
		Usage:  "The gas cost for each L1 message included in a rollup transaction",
		Value:  4200,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_PER_L1_MSG"),
	}

	GasLimitBuffer = cli.Uint64Flag{
		Name:   "GAS_LIMIT_BUFFER",
		Usage:  "The gas limit buffer for a transaction",
		Value:  120,
		EnvVar: prefixEnvVar("GAS_LIMIT_BUFFER"),
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
	SubmitterAddressFlag,
	L2SequencerAddressFlag,
	PrivateKeyFlag,
	TxFeeLimitFlag,
	L1StakingAddressFlag,
	L2GovAddressFlag,
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
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
