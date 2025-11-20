package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "TOKEN_PRICE_ORACLE_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */
	L2EthRPCFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPC"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key to use for sending transactions to L2",
		Required: true,
		EnvVar:   prefixEnvVar("PRIVATE_KEY"),
	}

	/* Optional Flags */

	TxnPerBatchFlag = cli.Uint64Flag{
		Name:   "txn-per-batch",
		Usage:  "Expected transactions per batch",
		Value:  50,
		EnvVar: prefixEnvVar("TXN_PER_BATCH"),
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

	PriceThresholdFlag = cli.Uint64Flag{
		Name:   "price-threshold",
		Usage:  "Price change threshold percentage to trigger update (e.g. 5 for 5%)",
		Value:  0,
		EnvVar: prefixEnvVar("PRICE_THRESHOLD"),
	}

	PriceFeedPriorityFlag = cli.StringFlag{
		Name:   "price-feed-priority",
		Usage:  "Comma-separated list of price feed types in priority order (e.g. \"bitget,binance\")",
		Value:  "bitget",
		EnvVar: prefixEnvVar("PRICE_FEED_PRIORITY"),
	}

	TokenMappingBitgetFlag = cli.StringFlag{
		Name:   "token-mapping-bitget",
		Usage:  "Token ID to Bitget trading pair mapping (e.g. \"1:BTCUSDT,2:ETHUSDT\")",
		Value:  "",
		EnvVar: prefixEnvVar("TOKEN_MAPPING_BITGET"),
	}

	TokenMappingBinanceFlag = cli.StringFlag{
		Name:   "token-mapping-binance",
		Usage:  "Token ID to Binance trading pair mapping (e.g. \"1:BTCUSDT,2:ETHUSDT\")",
		Value:  "",
		EnvVar: prefixEnvVar("TOKEN_MAPPING_BINANCE"),
	}

	BitgetAPIBaseURLFlag = cli.StringFlag{
		Name:   "bitget-api-base-url",
		Usage:  "Bitget API base URL (required if bitget feed is enabled)",
		Value:  "",
		EnvVar: prefixEnvVar("BITGET_API_BASE_URL"),
	}

	BinanceAPIBaseURLFlag = cli.StringFlag{
		Name:   "binance-api-base-url",
		Usage:  "Binance API base URL (required if binance feed is enabled)",
		Value:  "",
		EnvVar: prefixEnvVar("BINANCE_API_BASE_URL"),
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
	L2EthRPCFlag,
	PrivateKeyFlag,
}

var optionalFlags = []cli.Flag{
	TxnPerBatchFlag,
	PriceUpdateIntervalFlag,
	TokenIDsFlag,
	PriceThresholdFlag,
	PriceFeedPriorityFlag,
	TokenMappingBitgetFlag,
	TokenMappingBinanceFlag,
	BitgetAPIBaseURLFlag,
	BinanceAPIBaseURLFlag,

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
