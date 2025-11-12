package config

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
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
	L2TokenRegistryAddr  common.Address

	// Private key
	PrivateKey string

	// Update parameters
	GasThreshold     uint64        // Percentage threshold to trigger updates
	Interval         time.Duration // Base fee update interval
	OverheadInterval uint64        // Scalar update frequency (every N base fee updates)
	TxnPerBatch      uint64        // Expected transactions per batch

	// Updater enable/disable switches
	BaseFeeUpdateEnabled bool // Whether base fee updates are enabled
	ScalarUpdateEnabled  bool // Whether scalar updates are enabled
	PriceUpdateEnabled   bool // Whether price updates are enabled

	// Price update parameters
	PriceUpdateInterval time.Duration     // Price update interval
	TokenIDs            []uint16          // Token IDs to update
	BasePrice           *big.Int          // Base price for mock feed
	PriceVariation      float64           // Price variation for mock feed
	PriceThreshold      uint64            // Price change threshold percentage to trigger update
	PriceFeedType       string            // Price feed type (mock, bitget)
	TokenMapping        map[uint16]string // Token ID to trading pair mapping

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

		BaseFeeUpdateEnabled: ctx.Bool(flags.BaseFeeUpdateEnabledFlag.Name),
		ScalarUpdateEnabled:  ctx.Bool(flags.ScalarUpdateEnabledFlag.Name),
		PriceUpdateEnabled:   ctx.Bool(flags.PriceUpdateEnabledFlag.Name),

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

	// Parse token registry address (optional)
	registryAddr := ctx.String(flags.L2TokenRegistryAddressFlag.Name)
	if registryAddr != "" {
		if !common.IsHexAddress(registryAddr) {
			return nil, fmt.Errorf("invalid L2_TOKEN_REGISTRY address: %s", registryAddr)
		}
		cfg.L2TokenRegistryAddr = common.HexToAddress(registryAddr)
	}

	// Parse price update interval
	cfg.PriceUpdateInterval = ctx.Duration(flags.PriceUpdateIntervalFlag.Name)

	// Parse token IDs
	tokenIDsStr := ctx.String(flags.TokenIDsFlag.Name)
	if tokenIDsStr != "" {
		parts := strings.Split(tokenIDsStr, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			id, err := strconv.ParseUint(part, 10, 16)
			if err != nil {
				return nil, fmt.Errorf("invalid token ID '%s': %w", part, err)
			}
			cfg.TokenIDs = append(cfg.TokenIDs, uint16(id))
		}
	}

	// Parse base price
	basePriceStr := ctx.String(flags.BasePriceFlag.Name)
	cfg.BasePrice = new(big.Int)
	if _, ok := cfg.BasePrice.SetString(basePriceStr, 10); !ok {
		return nil, fmt.Errorf("invalid base price: %s", basePriceStr)
	}

	cfg.PriceVariation = ctx.Float64(flags.PriceVariationFlag.Name)
	cfg.PriceThreshold = ctx.Uint64(flags.PriceThresholdFlag.Name)

	// Parse price feed type
	cfg.PriceFeedType = ctx.String(flags.PriceFeedTypeFlag.Name)
	if cfg.PriceFeedType != "mock" && cfg.PriceFeedType != "bitget" {
		return nil, fmt.Errorf("invalid price feed type: %s (must be 'mock' or 'bitget')", cfg.PriceFeedType)
	}

	// Parse token mapping for exchanges
	tokenMappingStr := ctx.String(flags.TokenMappingFlag.Name)
	cfg.TokenMapping = make(map[uint16]string)
	if tokenMappingStr != "" {
		pairs := strings.Split(tokenMappingStr, ",")
		for _, pair := range pairs {
			pair = strings.TrimSpace(pair)
			if pair == "" {
				continue
			}
			parts := strings.Split(pair, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid token mapping pair '%s' (expected format: tokenID:symbol)", pair)
			}
			tokenID, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 16)
			if err != nil {
				return nil, fmt.Errorf("invalid token ID in mapping '%s': %w", parts[0], err)
			}
			symbol := strings.TrimSpace(parts[1])
			cfg.TokenMapping[uint16(tokenID)] = symbol
		}
	}

	return cfg, nil
}

// MetricAddress returns the metrics server address
func (c *Config) MetricAddress() string {
	return fmt.Sprintf("%s:%d", c.MetricsHostname, c.MetricsPort)
}
