package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "STAKING_ORACLE_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	L1EthRPCFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRPCFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPC"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key to use for sending to the rollup contract",
		EnvVar:   prefixEnvVar("RECORD_PRIVATE_KEY"),
		Required: true,
	}
	TendermintFlag = cli.StringFlag{
		Name:     "l2-tendermint-rpc",
		Usage:    "HTTP provider Tendermint URL for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_TENDERMINT_RPC"),
	}
	WSEndpointFlag = cli.StringFlag{
		Name:     "l2-ws-endpoint",
		Usage:    "HTTP provider ws-endpoint URL for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_WS_ENDPOINT"),
	}

	RollupAddress = cli.StringFlag{
		Name:   "rollup-address",
		Usage:  "Address of the rollup",
		EnvVar: prefixEnvVar("ROLLUP"),
	}

	MaxHeaderBatchSizeFlag = cli.Uint64Flag{
		Name:   "max-header-batch-size",
		Usage:  "The maximum number of headers to request as a batch",
		Value:  2000,
		EnvVar: prefixEnvVar("MAX_HEADER_BATCH_SIZE"),
	}

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar("LOG_TERMINAL"),
	}
	LogFilenameFlag = cli.StringFlag{
		Name:   "log.filename",
		Usage:  "The target file for writing logs, backup log files will be retained in the same directory.",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}
	LogFileMaxSizeFlag = cli.IntFlag{
		Name:   "log.maxsize",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes. It is used only when log.filename is provided.",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}
	LogFileMaxAgeFlag = cli.IntFlag{
		Name:   "log.maxage",
		Usage:  "The maximum number of days to retain old log files based on the timestamp encoded in their filename. It defaults to 30 days. It is used only when log.filename is provided.",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}
	LogCompressFlag = cli.BoolFlag{
		Name:   "log.compress",
		Usage:  "Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression. It is used only when log.filename is provided.",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}
	MetricsServerEnableFlag = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Whether or not to run the embedded metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostnameFlag = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "The hostname of the metrics server",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPortFlag = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "The port of the metrics server",
		Value:  6060,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRPCFlag,
	L2EthRPCFlag,
	PrivateKeyFlag,
	TendermintFlag,
	WSEndpointFlag,
	RollupAddress,
}

var optionalFlags = []cli.Flag{

	LogLevelFlag,
	LogTerminalFlag,
	LogFilenameFlag,
	LogFileMaxSizeFlag,
	LogFileMaxAgeFlag,
	LogCompressFlag,
	MaxHeaderBatchSizeFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
