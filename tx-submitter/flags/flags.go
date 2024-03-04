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

	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "The private key to use for sending to the rollup contract",
		EnvVar: prefixEnvVar("L1_PRIVATE_KEY"),
	}

	TxTimeoutFlag = cli.DurationFlag{
		Name:     "tx-timeout",
		Usage:    "Timeout for transaction submission",
		Value:    10,
		EnvVar:   prefixEnvVar("TX_TIMEOUT"),
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

	// L2 contract
	SubmitterAddressFlag = cli.StringFlag{
		Name:     "submitter-address",
		Usage:    "Address of the submitter contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SUBMITTER_ADDRESS"),
	}
	L2SequencerAddressFlag = cli.StringFlag{
		Name:     "l2-sequencer",
		Usage:    "Address of the sequencer contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SEQUENCER_ADDRESS"),
	}

	/* Optional Flags */
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
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
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	L1EthRpcFlag,
	L2EthRpcsFlag,
	RollupAddressFlag,
	PollIntervalFlag,
	TxTimeoutFlag,
	FinalizeFlag,
	MaxFinalizeNumFlag,
	PriorityRollupFlag,
	SubmitterAddressFlag,
	L2SequencerAddressFlag,
	PrivateKeyFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	MetricsServerEnable,
	MetricsHostname,
	MetricsPort,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
