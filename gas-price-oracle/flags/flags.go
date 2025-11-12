package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "GAS_ORACLE_"

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
	L1BeaconRPCFlag = cli.StringFlag{
		Name:     "l1-beacon-rpc",
		Usage:    "HTTP provider URL for L1 Beacon Chain",
		Required: true,
		EnvVar:   prefixEnvVar("L1_BEACON_RPC"),
	}

	L1RollupAddressFlag = cli.StringFlag{
		Name:     "l1-rollup-address",
		Usage:    "Address of the L1 Rollup contract",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ROLLUP"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key to use for sending transactions to L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_PRIVATE_KEY"),
	}

	/* Optional Flags */

	L2GasPriceOracleAddressFlag = cli.StringFlag{
		Name:   "l2-gas-price-oracle-address",
		Usage:  "Address of the L2 GasPriceOracle contract",
		Value:  "0x5300000000000000000000000000000000000002",
		EnvVar: prefixEnvVar("L2_GAS_PRICE_ORACLE"),
	}

	GasThresholdFlag = cli.Uint64Flag{
		Name:   "gas-threshold",
		Usage:  "Percentage threshold to trigger updates",
		Value:  10,
		EnvVar: prefixEnvVar("GAS_THRESHOLD"),
	}

	IntervalFlag = cli.DurationFlag{
		Name:   "interval",
		Usage:  "Base fee update interval",
		Value:  6 * time.Second,
		EnvVar: prefixEnvVar("INTERVAL"),
	}

	OverheadIntervalFlag = cli.Uint64Flag{
		Name:   "overhead-interval",
		Usage:  "Scalar update frequency (every N base fee updates)",
		Value:  10,
		EnvVar: prefixEnvVar("OVERHEAD_INTERVAL"),
	}

	TxnPerBatchFlag = cli.Uint64Flag{
		Name:   "txn-per-batch",
		Usage:  "Expected transactions per batch",
		Value:  50,
		EnvVar: prefixEnvVar("TXN_PER_BATCH"),
	}

	// Logging flags
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}

	LogFilenameFlag = cli.StringFlag{
		Name:   "log-filename",
		Usage:  "The target file for writing logs",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}

	LogFileMaxSizeFlag = cli.IntFlag{
		Name:   "log-file-max-size",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}

	LogFileMaxAgeFlag = cli.IntFlag{
		Name:   "log-file-max-age",
		Usage:  "The maximum number of days to retain old log files",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}

	LogCompressFlag = cli.BoolFlag{
		Name:   "log-compress",
		Usage:  "Whether to compress rotated log files using gzip",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	// Metrics flags
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
	L1BeaconRPCFlag,
	L1RollupAddressFlag,
	PrivateKeyFlag,
}

var optionalFlags = []cli.Flag{
	L2GasPriceOracleAddressFlag,
	GasThresholdFlag,
	IntervalFlag,
	OverheadIntervalFlag,
	TxnPerBatchFlag,

	LogLevelFlag,
	LogFilenameFlag,
	LogFileMaxSizeFlag,
	LogFileMaxAgeFlag,
	LogCompressFlag,

	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
