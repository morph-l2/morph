package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "CHALLENGER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	L1EthRPCFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}

	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "L1 Rollup contract address",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "Hex private key used when external sign is disabled",
		EnvVar: prefixEnvVar("PRIVATE_KEY"),
	}

	HTTPHostnameFlag = cli.StringFlag{
		Name:   "http-hostname",
		Usage:  "Hostname of the challenge HTTP server",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("HTTP_HOSTNAME"),
	}

	HTTPPortFlag = cli.Uint64Flag{
		Name:   "http-port",
		Usage:  "Port of the challenge HTTP server",
		Value:  8080,
		EnvVar: prefixEnvVar("HTTP_PORT"),
	}

	HTTPAuthTokenFlag = cli.StringFlag{
		Name:   "http-auth-token",
		Usage:  "Bearer token required by POST /challenge. Empty leaves the endpoint open",
		EnvVar: prefixEnvVar("HTTP_AUTH_TOKEN"),
	}

	ReceiptTimeoutFlag = cli.DurationFlag{
		Name:   "receipt-timeout",
		Usage:  "How long to wait for the challenge transaction receipt",
		Value:  2 * time.Minute,
		EnvVar: prefixEnvVar("RECEIPT_TIMEOUT"),
	}

	ExternalSignFlag = cli.BoolFlag{
		Name:   "external-sign",
		Usage:  "Sign challengeState transactions with the remote signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN"),
	}

	ExternalSignAddressFlag = cli.StringFlag{
		Name:   "external-sign-address",
		Usage:  "Address of the external signer. This address must be an L1 challenger",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_ADDRESS"),
	}

	ExternalSignAppidFlag = cli.StringFlag{
		Name:   "external-sign-appid",
		Usage:  "App id for the external signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_APPID"),
	}

	ExternalSignChainFlag = cli.StringFlag{
		Name:   "external-sign-chain",
		Usage:  "Chain identifier for the external signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_CHAIN"),
	}

	ExternalSignUrlFlag = cli.StringFlag{
		Name:   "external-sign-url",
		Usage:  "URL of the external signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_URL"),
	}

	ExternalSignRsaPrivFlag = cli.StringFlag{
		Name:   "external-sign-rsa-priv",
		Usage:  "RSA private key for the external signer",
		EnvVar: prefixEnvVar("EXTERNAL_SIGN_RSA_PRIV"),
	}

	MetricsServerEnableFlag = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Run the metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}

	MetricsHostnameFlag = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "Hostname of the metrics server",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}

	MetricsPortFlag = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "Port of the metrics server",
		Value:  6060,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "Lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}

	LogTerminalFlag = cli.BoolFlag{
		Name:   "log-terminal",
		Usage:  "Print logs in terminal format instead of JSON",
		EnvVar: prefixEnvVar("LOG_TERMINAL"),
	}

	LogFilenameFlag = cli.StringFlag{
		Name:   "log-filename",
		Usage:  "Target file for writing logs",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}

	LogFileMaxSizeFlag = cli.IntFlag{
		Name:   "log-file-max-size",
		Usage:  "Maximum size in megabytes of the log file before it gets rotated",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}

	LogFileMaxAgeFlag = cli.IntFlag{
		Name:   "log-file-max-age",
		Usage:  "Maximum number of days to retain old log files",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}

	LogCompressFlag = cli.BoolFlag{
		Name:   "log-compress",
		Usage:  "Compress rotated log files using gzip",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	GasFeeCapFlag = cli.Uint64Flag{
		Name:   "gas-fee-cap",
		Usage:  "Max gas fee cap in wei. Actual fee is min(dynamic, this value)",
		EnvVar: prefixEnvVar("GAS_FEE_CAP"),
	}

	GasTipCapFlag = cli.Uint64Flag{
		Name:   "gas-tip-cap",
		Usage:  "Max gas tip cap in wei. Actual tip is min(dynamic, this value)",
		EnvVar: prefixEnvVar("GAS_TIP_CAP"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRPCFlag,
	RollupAddressFlag,
}

var optionalFlags = []cli.Flag{
	PrivateKeyFlag,
	HTTPHostnameFlag,
	HTTPPortFlag,
	HTTPAuthTokenFlag,
	ReceiptTimeoutFlag,

	ExternalSignFlag,
	ExternalSignAddressFlag,
	ExternalSignAppidFlag,
	ExternalSignChainFlag,
	ExternalSignUrlFlag,
	ExternalSignRsaPrivFlag,

	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,

	LogLevelFlag,
	LogTerminalFlag,
	LogFilenameFlag,
	LogFileMaxSizeFlag,
	LogFileMaxAgeFlag,
	LogCompressFlag,

	GasFeeCapFlag,
	GasTipCapFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
