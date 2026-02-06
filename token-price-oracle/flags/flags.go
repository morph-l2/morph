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
		Name:   "private-key",
		Usage:  "The private key to use for sending transactions to L2 (not required if external-sign is enabled)",
		EnvVar: prefixEnvVar("PRIVATE_KEY"),
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

	PriceThresholdFlag = cli.Uint64Flag{
		Name:   "price-threshold",
		Usage:  "Price change threshold in basis points (bps) to trigger update (e.g. 100 for 1%, 10 for 0.1%, 1 for 0.01%)",
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

	// External sign flags
	ExternalSignFlag = cli.BoolFlag{
		Name:   "external-sign",
		Usage:  "Enable external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN"),
	}

	ExternalSignAddressFlag = cli.StringFlag{
		Name:   "external-sign-address",
		Usage:  "The address of the external signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_ADDRESS"),
	}

	ExternalSignAppidFlag = cli.StringFlag{
		Name:   "external-sign-appid",
		Usage:  "The appid for external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_APPID"),
	}

	ExternalSignChainFlag = cli.StringFlag{
		Name:   "external-sign-chain",
		Usage:  "The chain identifier for external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_CHAIN"),
	}

	ExternalSignUrlFlag = cli.StringFlag{
		Name:   "external-sign-url",
		Usage:  "The URL of the external sign service",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_URL"),
	}

	ExternalSignRsaPrivFlag = cli.StringFlag{
		Name:   "external-sign-rsa-priv",
		Usage:  "The RSA private key for external sign",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_RSA_PRIV"),
	}

	// Gas fee flags (optional - if set, use fixed values instead of dynamic)
	GasFeeCapFlag = cli.Uint64Flag{
		Name:   "gas-fee-cap",
		Usage:  "Fixed gas fee cap in wei (if set, overrides dynamic gas price)",
		Value:  0,
		EnvVar: prefixEnvVar("GAS_FEE_CAP"),
	}

	GasTipCapFlag = cli.Uint64Flag{
		Name:   "gas-tip-cap",
		Usage:  "Fixed gas tip cap in wei (if set, overrides dynamic gas tip)",
		Value:  0,
		EnvVar: prefixEnvVar("GAS_TIP_CAP"),
	}
)

var requiredFlags = []cli.Flag{
	L2EthRPCFlag,
	PrivateKeyFlag,
}

var optionalFlags = []cli.Flag{
	TxnPerBatchFlag,
	PriceUpdateIntervalFlag,
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

	// External sign
	ExternalSignFlag,
	ExternalSignAddressFlag,
	ExternalSignAppidFlag,
	ExternalSignChainFlag,
	ExternalSignUrlFlag,
	ExternalSignRsaPrivFlag,

	// Gas fee
	GasFeeCapFlag,
	GasTipCapFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
