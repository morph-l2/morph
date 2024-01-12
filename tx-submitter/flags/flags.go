package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "TX_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   prefixEnvVar("BUILD_ENV"),
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRpcsFlag = cli.StringSliceFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URLs for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPCS"),
	}
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the rollup contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}

	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("POLL_INTERVAL"),
	}

	SafeMinimumEtherBalanceFlag = cli.Uint64Flag{
		Name:     "safe-minimum-ether-balance",
		Usage:    "Safe minimum amount of ether the batch submitter key should hold before it starts to log errors",
		Required: true,
		EnvVar:   prefixEnvVar("SAFE_MINIMUM_ETHER_BALANCE"),
	}

	/* Optional Flags */

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "The private key to use for sending to the rollup contract",
		EnvVar: prefixEnvVar("L1_PRIVATE_KEY"),
	}

	NetworkTimeoutFlag = cli.DurationFlag{
		Name:     "network-timeout",
		Usage:    "Timeout for network requests",
		Value:    10,
		EnvVar:   prefixEnvVar("NETWORK_TIMEOUT"),
		Required: true,
	}
	TxTimeoutFlag = cli.DurationFlag{
		Name:     "tx-timeout",
		Usage:    "Timeout for transaction submission",
		Value:    10,
		EnvVar:   prefixEnvVar("TX_TIMEOUT"),
		Required: true,
	}
	BatchBuildTimeoutFlag = cli.DurationFlag{
		Name: "batch-build-timeout",
		Usage: "Maximum amount of time that we will wait before " +
			"submitting an under-sized batch",
		Value:    60,
		EnvVar:   prefixEnvVar("MAX_BATCH_BUILD_TIME"),
		Required: true,
	}
	MaxTxSizeFlag = cli.Uint64Flag{
		Name:     "max-tx-size",
		Usage:    "Maximum byte of data that can be submitted in a single transaction",
		Value:    123 * 1014,
		EnvVar:   prefixEnvVar("MAX_TX_SIZE"),
		Required: true,
	}

	// finalize flags
	FinalizeFlag = cli.BoolFlag{
		Name:     "finalize",
		Usage:    "Enable finalize",
		EnvVar:   prefixEnvVar("FINALIZE"),
		Required: true,
	}

	MaxFinalizeNumFlag = cli.Uint64Flag{
		Name:     "max-finalize-number",
		Usage:    "Maximum number of finalize",
		EnvVar:   prefixEnvVar("MAX_FINALIZE_NUM"),
		Required: true,
	}

	// decentralize config
	PriorityRollupFlag = cli.BoolFlag{
		Name:     "priority-rollup",
		Usage:    "Enable priority rollup",
		EnvVar:   prefixEnvVar("PRIORITY_ROLLUP"),
		Required: true,
	}

	// metrics
	MetricsServerEnable = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Enable metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostname = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "Hostname at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPort = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "Port at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}

	// L2 contract
	SubmitterAddressFlag = cli.StringFlag{
		Name:     "submitter-address",
		Usage:    "Address of the submitter contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SUBMITTER_ADDRESS"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	L1EthRpcFlag,
	L2EthRpcsFlag,
	RollupAddressFlag,
	PollIntervalFlag,
	SafeMinimumEtherBalanceFlag,
	NetworkTimeoutFlag,
	TxTimeoutFlag,
	BatchBuildTimeoutFlag,
	MaxTxSizeFlag,
	SubmitterAddressFlag,
	FinalizeFlag,
	MaxFinalizeNumFlag,
	PriorityRollupFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	PrivateKeyFlag,
	MetricsServerEnable,
	MetricsHostname,
	MetricsPort,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
