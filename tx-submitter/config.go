package tx_summitter

import (
	"time"

	"github.com/morph-l2/tx-submitter/flags"
	"github.com/urfave/cli"
)

type Config struct {
	/* Required Params */

	// BuildEnv identifies the environment this binary is intended for, i.e.
	// production, development, etc.
	BuildEnv string

	// EthNetworkName identifies the intended Ethereum network.
	EthNetworkName string

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for L1.
	L2EthRpcs []string

	// RollupAddress is the Rollup contract address.
	RollupAddress string

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// LogLevel is the lowest log level that will be output.
	LogLevel string

	// PrivateKey the private key of the wallet used to submit
	// transactions to the rollup contract.
	PrivateKey string
	TxTimeout  time.Duration

	// MaxBlock is the max block number to handle
	MaxBlock uint64
	// MinBlock is the min block number to handle
	MinBlock uint64
	// finalize
	// if start finalize
	Finalize       bool
	MaxFinalizeNum uint64
	// L2 contract
	SubmitterAddress string

	// metrics
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64

	// decentralized
	PriorityRollup bool
}

// NewConfig parses the DriverConfig from the provided flags or environment variables.
// This method fails if ValidateConfig deems the configuration to be malformed.
func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		/* Required Flags */
		BuildEnv:      ctx.GlobalString(flags.BuildEnvFlag.Name),
		L1EthRpc:      ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpcs:     ctx.GlobalStringSlice(flags.L2EthRpcsFlag.Name),
		RollupAddress: ctx.GlobalString(flags.RollupAddressFlag.Name),
		PollInterval:  ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		LogLevel:      ctx.GlobalString(flags.LogLevelFlag.Name),
		PrivateKey:    ctx.GlobalString(flags.PrivateKeyFlag.Name),
		TxTimeout:     ctx.GlobalDuration(flags.TxTimeoutFlag.Name),
		// finalize
		Finalize:       ctx.GlobalBool(flags.FinalizeFlag.Name),
		MaxFinalizeNum: ctx.GlobalUint64(flags.MaxFinalizeNumFlag.Name),
		// L2 contract
		SubmitterAddress: ctx.GlobalString(flags.SubmitterAddressFlag.Name),
		// metrics
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnable.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostname.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPort.Name),
		// decentralized
		PriorityRollup: ctx.GlobalBool(flags.PriorityRollupFlag.Name),
	}

	return cfg, nil
}
