package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/urfave/cli"

	"morph-l2/challenger/flags"
)

// Config contains the challenger service configuration.
type Config struct {
	L1RPC         string
	RollupAddress common.Address
	PrivateKey    string

	HTTPHostname   string
	HTTPPort       uint64
	HTTPAuthToken  string
	ReceiptTimeout time.Duration

	ExternalSign        bool
	ExternalSignAddress string
	ExternalSignAppid   string
	ExternalSignChain   string
	ExternalSignUrl     string
	ExternalSignRsaPriv string

	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64

	LogLevel       string
	LogTerminal    bool
	LogFilename    string
	LogFileMaxSize int
	LogFileMaxAge  int
	LogCompress    bool

	GasFeeCap *uint64
	GasTipCap *uint64
}

// LoadConfig loads configuration from CLI flags and environment variables.
func LoadConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{
		L1RPC:      ctx.String(flags.L1EthRPCFlag.Name),
		PrivateKey: strings.TrimSpace(ctx.String(flags.PrivateKeyFlag.Name)),

		HTTPHostname:   ctx.String(flags.HTTPHostnameFlag.Name),
		HTTPPort:       ctx.Uint64(flags.HTTPPortFlag.Name),
		HTTPAuthToken:  strings.TrimSpace(ctx.String(flags.HTTPAuthTokenFlag.Name)),
		ReceiptTimeout: ctx.Duration(flags.ReceiptTimeoutFlag.Name),

		ExternalSign:        ctx.Bool(flags.ExternalSignFlag.Name),
		ExternalSignAddress: strings.TrimSpace(ctx.String(flags.ExternalSignAddressFlag.Name)),
		ExternalSignAppid:   strings.TrimSpace(ctx.String(flags.ExternalSignAppidFlag.Name)),
		ExternalSignChain:   strings.TrimSpace(ctx.String(flags.ExternalSignChainFlag.Name)),
		ExternalSignUrl:     strings.TrimSpace(ctx.String(flags.ExternalSignUrlFlag.Name)),
		ExternalSignRsaPriv: strings.TrimSpace(ctx.String(flags.ExternalSignRsaPrivFlag.Name)),

		MetricsServerEnable: ctx.Bool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.String(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.Uint64(flags.MetricsPortFlag.Name),

		LogLevel:       ctx.String(flags.LogLevelFlag.Name),
		LogTerminal:    ctx.Bool(flags.LogTerminalFlag.Name),
		LogFilename:    ctx.String(flags.LogFilenameFlag.Name),
		LogFileMaxSize: ctx.Int(flags.LogFileMaxSizeFlag.Name),
		LogFileMaxAge:  ctx.Int(flags.LogFileMaxAgeFlag.Name),
		LogCompress:    ctx.Bool(flags.LogCompressFlag.Name),
	}

	if ctx.IsSet(flags.GasFeeCapFlag.Name) {
		v := ctx.Uint64(flags.GasFeeCapFlag.Name)
		cfg.GasFeeCap = &v
	}
	if ctx.IsSet(flags.GasTipCapFlag.Name) {
		v := ctx.Uint64(flags.GasTipCapFlag.Name)
		cfg.GasTipCap = &v
	}
	if cfg.GasFeeCap != nil && cfg.GasTipCap != nil && *cfg.GasFeeCap < *cfg.GasTipCap {
		return nil, fmt.Errorf("--gas-fee-cap (%d) must be >= --gas-tip-cap (%d)", *cfg.GasFeeCap, *cfg.GasTipCap)
	}

	rollupAddress := strings.TrimSpace(ctx.String(flags.RollupAddressFlag.Name))
	if !common.IsHexAddress(rollupAddress) {
		return nil, fmt.Errorf("invalid rollup address: %s", rollupAddress)
	}
	cfg.RollupAddress = common.HexToAddress(rollupAddress)

	if cfg.L1RPC == "" {
		return nil, fmt.Errorf("l1 rpc is required")
	}
	if cfg.HTTPPort == 0 || cfg.HTTPPort > 65535 {
		return nil, fmt.Errorf("invalid http port: %d", cfg.HTTPPort)
	}
	if cfg.MetricsServerEnable && (cfg.MetricsPort == 0 || cfg.MetricsPort > 65535) {
		return nil, fmt.Errorf("invalid metrics port: %d", cfg.MetricsPort)
	}
	if cfg.ReceiptTimeout <= 0 {
		return nil, fmt.Errorf("receipt timeout must be positive")
	}

	if cfg.ExternalSign {
		if cfg.ExternalSignAddress == "" || cfg.ExternalSignUrl == "" ||
			cfg.ExternalSignAppid == "" || cfg.ExternalSignChain == "" ||
			cfg.ExternalSignRsaPriv == "" {
			return nil, fmt.Errorf("external sign is enabled but missing required config: address=%s, url=%s, appid=%s, chain=%s, rsa_priv_set=%t",
				cfg.ExternalSignAddress, cfg.ExternalSignUrl, cfg.ExternalSignAppid, cfg.ExternalSignChain, cfg.ExternalSignRsaPriv != "")
		}
		if !common.IsHexAddress(cfg.ExternalSignAddress) {
			return nil, fmt.Errorf("invalid external sign address: %s", cfg.ExternalSignAddress)
		}
	} else if cfg.PrivateKey == "" {
		return nil, fmt.Errorf("private key is required when external sign is not enabled")
	}

	return cfg, nil
}

// HTTPAddress returns the challenge API listen address.
func (c *Config) HTTPAddress() string {
	return fmt.Sprintf("%s:%d", c.HTTPHostname, c.HTTPPort)
}

// MetricAddress returns the metrics server listen address.
func (c *Config) MetricAddress() string {
	return fmt.Sprintf("%s:%d", c.MetricsHostname, c.MetricsPort)
}
