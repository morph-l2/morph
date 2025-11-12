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

	L2TokenRegistryAddressFlag = cli.StringFlag{
		Name:   "l2-token-registry-address",
		Usage:  "Address of the L2 TokenRegistry contract",
		Value:  "",
		EnvVar: prefixEnvVar("L2_TOKEN_REGISTRY"),
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

	// Updater enable/disable flags
	BaseFeeUpdateEnabledFlag = cli.BoolFlag{
		Name:   "basefee-update-enabled",
		Usage:  "Whether to enable base fee updates",
		EnvVar: prefixEnvVar("BASEFEE_UPDATE_ENABLED"),
	}

	ScalarUpdateEnabledFlag = cli.BoolFlag{
		Name:   "scalar-update-enabled",
		Usage:  "Whether to enable scalar updates",
		EnvVar: prefixEnvVar("SCALAR_UPDATE_ENABLED"),
	}

	PriceUpdateEnabledFlag = cli.BoolFlag{
		Name:   "price-update-enabled",
		Usage:  "Whether to enable token price updates",
		EnvVar: prefixEnvVar("PRICE_UPDATE_ENABLED"),
	}

	PriceUpdateIntervalFlag = cli.DurationFlag{
		Name:   "price-update-interval",
		Usage:  "Token price update interval",
		Value:  60 * time.Second,
		EnvVar: prefixEnvVar("PRICE_UPDATE_INTERVAL"),
	}

	TokenIDsFlag = cli.StringFlag{
		Name:   "token-ids",
		Usage:  "Comma-separated token IDs to update prices for (e.g. \"1,2,3\")",
		Value:  "",
		EnvVar: prefixEnvVar("TOKEN_IDS"),
	}

	BasePriceFlag = cli.StringFlag{
		Name:   "base-price",
		Usage:  "Base price ratio for mock price feed (wei, e.g. 1000000000000000000 for 1:1 ratio)",
		Value:  "1000000000000000000",
		EnvVar: prefixEnvVar("BASE_PRICE"),
	}

	PriceVariationFlag = cli.Float64Flag{
		Name:   "price-variation",
		Usage:  "Price variation percentage for mock feed (e.g. 0.05 for ±5%)",
		Value:  0.05,
		EnvVar: prefixEnvVar("PRICE_VARIATION"),
	}

	PriceThresholdFlag = cli.Uint64Flag{
		Name:   "price-threshold",
		Usage:  "Price change threshold percentage to trigger update (e.g. 5 for 5%)",
		Value:  5,
		EnvVar: prefixEnvVar("PRICE_THRESHOLD"),
	}

	PriceFeedTypeFlag = cli.StringFlag{
		Name:   "price-feed-type",
		Usage:  "Price feed type: mock, bitget",
		Value:  "mock",
		EnvVar: prefixEnvVar("PRICE_FEED_TYPE"),
	}

	TokenMappingFlag = cli.StringFlag{
		Name:   "token-mapping",
		Usage:  "Token ID to trading pair mapping (e.g. \"1:BTCUSDT,2:ETHUSDT,3:BNBUSDT\")",
		Value:  "",
		EnvVar: prefixEnvVar("TOKEN_MAPPING"),
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
	L2TokenRegistryAddressFlag,
	GasThresholdFlag,
	IntervalFlag,
	OverheadIntervalFlag,
	TxnPerBatchFlag,

	BaseFeeUpdateEnabledFlag,
	ScalarUpdateEnabledFlag,
	PriceUpdateEnabledFlag,
	PriceUpdateIntervalFlag,
	TokenIDsFlag,
	BasePriceFlag,
	PriceVariationFlag,
	PriceThresholdFlag,
	PriceFeedTypeFlag,
	TokenMappingFlag,

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
