package utils

import (
	"time"

	"github.com/urfave/cli"

	"morph-l2/tx-submitter/flags"
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
	// StakingAddress
	L1StakingAddress string

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// LogLevel is the lowest log level that will be output.
	LogLevel string

	// PrivateKey the private key of the wallet used to submit
	// transactions to the rollup contract.
	PrivateKey string
	TxTimeout  time.Duration

	// finalize
	// if start finalize
	Finalize bool
	// L2 contract
	L2SequencerAddress string
	L2GovAddress       string

	// metrics
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64

	// decentralized
	PriorityRollup bool

	// tx fee limit
	TxFeeLimit uint64

	LogFilename    string
	LogFileMaxSize int
	LogFileMaxAge  int
	LogCompress    bool

	// rollup interval
	RollupInterval time.Duration
	// finalize interval
	FinalizeInterval time.Duration
	// tx process interval
	TxProcessInterval time.Duration

	// rollup gas base
	RollupTxGasBase uint64
	// rollup gas per l1 msg
	RollupTxGasPerL1Msg uint64
	// gas limit buffer
	GasLimitBuffer uint64

	// journal file path
	JournalFilePath string
	// calldata fee bump
	CalldataFeeBump uint64
	//max txs in pendingpool
	MaxTxsInPendingPool uint64

	// external sign
	ExternalSign        bool
	ExternalSignAddress string
	ExternalSignAppid   string
	ExternalSignChain   string
	ExternalSignUrl     string
	ExternalSignRsaPriv string
	// rough estimate gas switch
	RoughEstimateGas bool
	// rotator interval buffer
	RotatorBuffer int64
}

// NewConfig parses the DriverConfig from the provided flags or environment variables.
// This method fails if ValidateConfig deems the configuration to be malformed.
func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		/* Required Flags */
		BuildEnv:   ctx.GlobalString(flags.BuildEnvFlag.Name),
		L1EthRpc:   ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpcs:  ctx.GlobalStringSlice(flags.L2EthRpcsFlag.Name),
		LogLevel:   ctx.GlobalString(flags.LogLevelFlag.Name),
		PrivateKey: ctx.GlobalString(flags.PrivateKeyFlag.Name),
		TxTimeout:  ctx.GlobalDuration(flags.TxTimeoutFlag.Name),
		// finalize
		Finalize: ctx.GlobalBool(flags.FinalizeFlag.Name),
		// L1 contract
		L1StakingAddress: ctx.GlobalString(flags.L1StakingAddressFlag.Name),
		RollupAddress:    ctx.GlobalString(flags.RollupAddressFlag.Name),
		// L2 contract
		L2SequencerAddress: ctx.GlobalString(flags.L2SequencerAddressFlag.Name),
		L2GovAddress:       ctx.GlobalString(flags.L2GovAddressFlag.Name),
		// metrics
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnable.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostname.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPort.Name),
		// decentralized
		PriorityRollup: ctx.GlobalBool(flags.PriorityRollupFlag.Name),
		// tx config
		TxFeeLimit: ctx.GlobalUint64(flags.TxFeeLimitFlag.Name),

		LogFilename:    ctx.GlobalString(flags.LogFilename.Name),
		LogFileMaxSize: ctx.GlobalInt(flags.LogFileMaxSize.Name),
		LogFileMaxAge:  ctx.GlobalInt(flags.LogFileMaxAge.Name),
		LogCompress:    ctx.GlobalBool(flags.LogCompress.Name),

		RollupInterval:    ctx.GlobalDuration(flags.RollupInterval.Name),
		FinalizeInterval:  ctx.GlobalDuration(flags.FinalizeInterval.Name),
		TxProcessInterval: ctx.GlobalDuration(flags.TxProcessInterval.Name),

		RollupTxGasBase:     ctx.GlobalUint64(flags.RollupTxGasBase.Name),
		RollupTxGasPerL1Msg: ctx.GlobalUint64(flags.RollupTxGasPerL1Msg.Name),

		GasLimitBuffer: ctx.GlobalUint64(flags.GasLimitBuffer.Name),

		JournalFilePath: ctx.GlobalString(flags.JournalFlag.Name),
		// calldata fee bump
		CalldataFeeBump:     ctx.GlobalUint64(flags.CalldataFeeBumpFlag.Name),
		MaxTxsInPendingPool: ctx.GlobalUint64(flags.MaxTxsInPendingPoolFlag.Name),

		// external sign
		ExternalSign:        ctx.GlobalBool(flags.ExternalSign.Name),
		ExternalSignAppid:   ctx.GlobalString(flags.ExternalSignAppid.Name),
		ExternalSignAddress: ctx.GlobalString(flags.ExternalSignAddress.Name),
		ExternalSignChain:   ctx.GlobalString(flags.ExternalSignChain.Name),
		ExternalSignUrl:     ctx.GlobalString(flags.ExternalSignUrl.Name),
		ExternalSignRsaPriv: ctx.GlobalString(flags.ExternalSignRsaPriv.Name),
		// rough estimate gas switch
		RoughEstimateGas: ctx.GlobalBool(flags.RoughEstimateGasFlag.Name),
		// rotator interval buffer
		RotatorBuffer: ctx.GlobalInt64(flags.RotatorBufferFlag.Name),
	}

	return cfg, nil
}
