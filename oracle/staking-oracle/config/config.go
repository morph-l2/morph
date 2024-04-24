package indexer

import (
	"fmt"
	"time"

	"github.com/morph-l2/morph/oracle/staking-oracle/flags"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli"
)

type Config struct {
	/* Required Params */

	// ChainID identifies the chain being indexed.
	ChainID uint64

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for L1.
	L2EthRpc string

	TendermintRpc string
	WsEndpoint    string

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	/* Optional Params */

	// LogLevel is the lowest log level that will be output.
	LogLevel string

	// LogTerminal if true, prints to stdout in terminal format, otherwise
	// prints using JSON. If SentryEnable is true this flag is ignored, and logs
	// are printed using JSON.
	LogTerminal bool

	LogFilename string

	LogFileMaxSize int

	LogFileMaxAge int

	LogCompress bool

	// MetricsServerEnable if true, will create a metrics client and log to
	// Prometheus.
	MetricsServerEnable bool

	// MetricsHostname is the hostname at which the metrics server is running.
	MetricsHostname string

	// MetricsPort is the port at which the metrics server is running.
	MetricsPort uint64

	RollupAddress common.Address

	StakingAddress common.Address
}

// NewConfig parses the Config from the provided flags or environment variables.
// This method fails if ValidateConfig deems the configuration to be malformed.
func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		/* Required Flags */
		ChainID:        ctx.GlobalUint64(flags.ChainIDFlag.Name),
		L1EthRpc:       ctx.GlobalString(flags.L1EthRPCFlag.Name),
		L2EthRpc:       ctx.GlobalString(flags.L2EthRPCFlag.Name),
		TendermintRpc:  ctx.GlobalString(flags.TendermintFlag.Name),
		WsEndpoint:     ctx.GlobalString(flags.WSEndpointFlag.Name),
		RollupAddress:  common.HexToAddress(ctx.GlobalString(flags.RollupAddress.Name)),
		StakingAddress: common.HexToAddress(ctx.GlobalString(flags.StakingAddress.Name)),
		/* Optional Flags */

		LogLevel:            ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:         ctx.GlobalBool(flags.LogTerminalFlag.Name),
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPortFlag.Name),
	}

	if ctx.GlobalIsSet(flags.LogFilenameFlag.Name) {
		cfg.LogFilename = ctx.GlobalString(flags.LogFilenameFlag.Name)

		maxSize := ctx.GlobalInt(flags.LogFileMaxSizeFlag.Name)
		if maxSize < 1 {
			return Config{}, fmt.Errorf("wrong log.maxsize set: %d", maxSize)
		}
		cfg.LogFileMaxSize = maxSize
		maxAge := ctx.GlobalInt(flags.LogFileMaxAgeFlag.Name)
		if maxAge < 1 {
			return Config{}, fmt.Errorf("wrong log.maxage set: %d", maxAge)
		}
		cfg.LogFileMaxAge = maxAge
		cfg.LogCompress = ctx.GlobalBool(flags.LogCompressFlag.Name)
	}

	err := ValidateConfig(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// ValidateConfig ensures additional constraints on the parsed configuration to
// ensure that it is well-formed.
func ValidateConfig(cfg *Config) error {
	// Sanity check log level.
	if cfg.LogLevel == "" {
		cfg.LogLevel = "debug"
	}

	_, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return err
	}

	if (cfg.RollupAddress == common.Address{} ||
		cfg.StakingAddress == common.Address{}) {
		return fmt.Errorf(
			"invalied address,RollupAddress:%v,StakingAddress:%v",
			cfg.RollupAddress.String(),
			cfg.StakingAddress.String(),
		)
	}

	return nil
}
