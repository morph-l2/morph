package config

import (
	"fmt"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/morph/gas-price-oracle/flags"
	"github.com/urfave/cli"
)

// Config contains all service configurations
type Config struct {
	// RPC endpoints
	L1RPC       string
	L2RPC       string
	L1BeaconRPC string

	// Contract addresses
	L1RollupAddress      common.Address
	L2GasPriceOracleAddr common.Address

	// Private key
	PrivateKey string

	// Update parameters
	GasThreshold     uint64        // Percentage threshold to trigger updates
	Interval         time.Duration // Base fee update interval
	OverheadInterval uint64        // Scalar update frequency (every N base fee updates)
	TxnPerBatch      uint64        // Expected transactions per batch

	// Metrics
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64

	// Logging
	LogLevel       string
	LogFilename    string
	LogFileMaxSize int
	LogFileMaxAge  int
	LogCompress    bool
}

// LoadConfig loads configuration from cli.Context
func LoadConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{
		L1RPC:       ctx.String(flags.L1EthRPCFlag.Name),
		L2RPC:       ctx.String(flags.L2EthRPCFlag.Name),
		L1BeaconRPC: ctx.String(flags.L1BeaconRPCFlag.Name),
		PrivateKey:  ctx.String(flags.PrivateKeyFlag.Name),

		GasThreshold:     ctx.Uint64(flags.GasThresholdFlag.Name),
		Interval:         ctx.Duration(flags.IntervalFlag.Name),
		OverheadInterval: ctx.Uint64(flags.OverheadIntervalFlag.Name),
		TxnPerBatch:      ctx.Uint64(flags.TxnPerBatchFlag.Name),

		MetricsServerEnable: ctx.Bool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.String(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.Uint64(flags.MetricsPortFlag.Name),

		LogLevel:       ctx.String(flags.LogLevelFlag.Name),
		LogFilename:    ctx.String(flags.LogFilenameFlag.Name),
		LogFileMaxSize: ctx.Int(flags.LogFileMaxSizeFlag.Name),
		LogFileMaxAge:  ctx.Int(flags.LogFileMaxAgeFlag.Name),
		LogCompress:    ctx.Bool(flags.LogCompressFlag.Name),
	}

	// Parse contract addresses
	rollupAddr := ctx.String(flags.L1RollupAddressFlag.Name)
	if !common.IsHexAddress(rollupAddr) {
		return nil, fmt.Errorf("invalid L1_ROLLUP address: %s", rollupAddr)
	}
	cfg.L1RollupAddress = common.HexToAddress(rollupAddr)

	oracleAddr := ctx.String(flags.L2GasPriceOracleAddressFlag.Name)
	if !common.IsHexAddress(oracleAddr) {
		return nil, fmt.Errorf("invalid L2_GAS_PRICE_ORACLE address: %s", oracleAddr)
	}
	cfg.L2GasPriceOracleAddr = common.HexToAddress(oracleAddr)

	return cfg, nil
}

// MetricAddress returns the metrics server address
func (c *Config) MetricAddress() string {
	return fmt.Sprintf("%s:%d", c.MetricsHostname, c.MetricsPort)
}
