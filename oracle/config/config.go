package config

import (
	"fmt"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"

	"morph-l2/oracle/flags"
)

type Config struct {
	/* Required Params */

	// ChainID identifies the chain being indexed.
	ChainID uint64

	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for L1.
	L2EthRpc string

	PrivateKey string

	TendermintRpc string

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

	RollupAddr common.Address

	MaxSize    uint64
	MinSize    uint64
	StartBlock uint64
	PrivKey    string

	// external sign
	ExternalSign        bool
	ExternalSignAddress string
	ExternalSignAppid   string
	ExternalSignChain   string
	ExternalSignUrl     string
	ExternalSignRsaPriv string
}

// NewConfig parses the Config from the provided flags or environment variables.
// This method fails if ValidateConfig deems the configuration to be malformed.
func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		/* Required Flags */
		L1EthRpc:      ctx.GlobalString(flags.L1EthRPCFlag.Name),
		L2EthRpc:      ctx.GlobalString(flags.L2EthRPCFlag.Name),
		PrivateKey:    ctx.GlobalString(flags.PrivateKeyFlag.Name),
		TendermintRpc: ctx.GlobalString(flags.TendermintFlag.Name),
		RollupAddr:    common.HexToAddress(ctx.GlobalString(flags.RollupAddress.Name)),
		PrivKey:       ctx.GlobalString(flags.PrivateKeyFlag.Name),
		/* Optional Flags */
		MaxSize:             ctx.GlobalUint64(flags.MaxHeaderBatchSizeFlag.Name),
		MinSize:             ctx.GlobalUint64(flags.MinHeaderBatchSizeFlag.Name),
		LogLevel:            ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:         ctx.GlobalBool(flags.LogTerminalFlag.Name),
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPortFlag.Name),
		// external sign
		ExternalSign:        ctx.GlobalBool(flags.ExternalSign.Name),
		ExternalSignAppid:   ctx.GlobalString(flags.ExternalSignAppid.Name),
		ExternalSignAddress: ctx.GlobalString(flags.ExternalSignAddress.Name),
		ExternalSignChain:   ctx.GlobalString(flags.ExternalSignChain.Name),
		ExternalSignUrl:     ctx.GlobalString(flags.ExternalSignUrl.Name),
		ExternalSignRsaPriv: ctx.GlobalString(flags.ExternalSignRsaPriv.Name),
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

	if (cfg.RollupAddr == common.Address{}) {
		return fmt.Errorf(
			"invalied address,RollupAddress:%v",
			cfg.RollupAddr.String(),
		)
	}
	if cfg.ExternalSign &&
		(cfg.ExternalSignAddress == "" ||
			cfg.ExternalSignUrl == "" ||
			cfg.ExternalSignAppid == "" ||
			cfg.ExternalSignChain == "" ||
			cfg.ExternalSignRsaPriv == "") {
		return fmt.Errorf("invalid external sign config,ExternalSignAddress:%v,ExternalSignUrl:%v,ExternalSignAppid:%v,ExternalSignChain:%vExternalSignRsaPriv:%v",
			cfg.ExternalSignAddress,
			cfg.ExternalSignUrl,
			cfg.ExternalSignAppid,
			cfg.ExternalSignChain,
			cfg.ExternalSignRsaPriv,
		)
	}
	return nil
}
