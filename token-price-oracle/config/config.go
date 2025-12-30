package config

import (
	"fmt"
	"morph-l2/bindings/predeploys"
	"strconv"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/urfave/cli"
	"morph-l2/token-price-oracle/flags"
)

const (
	// MaxPriceThresholdBPS is the maximum allowed price threshold in basis points (100% = 10000 bps)
	MaxPriceThresholdBPS = 10000
)

// PriceFeedType represents the type of price feed source
type PriceFeedType string

const (
	PriceFeedTypeBitget  PriceFeedType = "bitget"
	PriceFeedTypeBinance PriceFeedType = "binance"
)

// ValidPriceFeedTypes returns all valid price feed types
func ValidPriceFeedTypes() []PriceFeedType {
	return []PriceFeedType{
		PriceFeedTypeBitget,
		// PriceFeedTypeBinance,  // TODO: Add back when Binance price feed is implemented
	}
}

// IsValidPriceFeedType checks if a string is a valid price feed type
func IsValidPriceFeedType(s string) bool {
	feedType := PriceFeedType(s)
	for _, valid := range ValidPriceFeedTypes() {
		if feedType == valid {
			return true
		}
	}
	return false
}

// String returns the string representation of PriceFeedType
func (p PriceFeedType) String() string {
	return string(p)
}

// Config contains all service configurations
type Config struct {
	// RPC endpoints
	L2RPC string
	// Contract addresses
	L2TokenRegistryAddr common.Address
	// Private key
	PrivateKey string
	// Price update parameters
	PriceUpdateInterval time.Duration                       // Price update interval
	PriceThreshold      uint64                              // Price change threshold percentage to trigger update
	PriceFeedPriority   []PriceFeedType                     // Price feed types in priority order (fallback mechanism)
	TokenMappings       map[PriceFeedType]map[uint16]string // Token ID to trading pair mappings for each price feed type
	BitgetAPIBaseURL    string                              // Bitget API base URL
	BinanceAPIBaseURL   string                              // Binance API base URL

	// External sign
	ExternalSign        bool
	ExternalSignAddress string
	ExternalSignAppid   string
	ExternalSignChain   string
	ExternalSignUrl     string
	ExternalSignRsaPriv string

	// Metrics
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64

	// Logging
	LogLevel       string
	LogTerminal    bool
	LogFilename    string
	LogFileMaxSize int
	LogFileMaxAge  int
	LogCompress    bool
}

// LoadConfig loads configuration from cli.Context
func LoadConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{
		L2RPC:      ctx.String(flags.L2EthRPCFlag.Name),
		PrivateKey: ctx.String(flags.PrivateKeyFlag.Name),

		// External sign
		ExternalSign:        ctx.Bool(flags.ExternalSignFlag.Name),
		ExternalSignAddress: ctx.String(flags.ExternalSignAddressFlag.Name),
		ExternalSignAppid:   ctx.String(flags.ExternalSignAppidFlag.Name),
		ExternalSignChain:   ctx.String(flags.ExternalSignChainFlag.Name),
		ExternalSignUrl:     ctx.String(flags.ExternalSignUrlFlag.Name),
		ExternalSignRsaPriv: ctx.String(flags.ExternalSignRsaPrivFlag.Name),

		MetricsServerEnable: ctx.Bool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.String(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.Uint64(flags.MetricsPortFlag.Name),

		LogLevel:       ctx.String(flags.LogLevelFlag.Name),
		LogFilename:    ctx.String(flags.LogFilenameFlag.Name),
		LogFileMaxSize: ctx.Int(flags.LogFileMaxSizeFlag.Name),
		LogFileMaxAge:  ctx.Int(flags.LogFileMaxAgeFlag.Name),
		LogCompress:    ctx.Bool(flags.LogCompressFlag.Name),
	}

	// Parse token registry address (optional)
	cfg.L2TokenRegistryAddr = predeploys.L2TokenRegistryAddr

	// Parse price update interval
	cfg.PriceUpdateInterval = ctx.Duration(flags.PriceUpdateIntervalFlag.Name)

	cfg.PriceThreshold = ctx.Uint64(flags.PriceThresholdFlag.Name)

	// Validate price threshold is reasonable (basis points should be 0-MaxPriceThresholdBPS)
	if cfg.PriceThreshold > MaxPriceThresholdBPS {
		return nil, fmt.Errorf("price threshold %d is too large (should be 0-%d basis points, where %d bps = 100%%)", 
			cfg.PriceThreshold, MaxPriceThresholdBPS, MaxPriceThresholdBPS)
	}

	// Parse and validate price feed priority list
	priorityStr := ctx.String(flags.PriceFeedPriorityFlag.Name)
	if priorityStr == "" {
		return nil, fmt.Errorf("price feed priority list cannot be empty")
	}

	priorityParts := strings.Split(priorityStr, ",")
	cfg.PriceFeedPriority = make([]PriceFeedType, 0, len(priorityParts))
	seenTypes := make(map[PriceFeedType]bool)

	for _, part := range priorityParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if !IsValidPriceFeedType(part) {
			validTypes := make([]string, len(ValidPriceFeedTypes()))
			for i, t := range ValidPriceFeedTypes() {
				validTypes[i] = t.String()
			}
			return nil, fmt.Errorf("invalid price feed type '%s' in priority list (must be one of: %s)", part, strings.Join(validTypes, ", "))
		}
		feedType := PriceFeedType(part)
		if seenTypes[feedType] {
			return nil, fmt.Errorf("duplicate price feed type '%s' in priority list", part)
		}
		seenTypes[feedType] = true
		cfg.PriceFeedPriority = append(cfg.PriceFeedPriority, feedType)
	}

	if len(cfg.PriceFeedPriority) == 0 {
		return nil, fmt.Errorf("price feed priority list cannot be empty after parsing")
	}

	// Helper function to parse token mapping
	parseTokenMapping := func(mappingStr string) (map[uint16]string, error) {
		mapping := make(map[uint16]string)
		if mappingStr == "" {
			return mapping, nil
		}
		pairs := strings.Split(mappingStr, ",")
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
			mapping[uint16(tokenID)] = symbol
		}
		return mapping, nil
	}

	// Parse all token mappings for different price feed types
	cfg.TokenMappings = make(map[PriceFeedType]map[uint16]string)

	bitgetMapping, err := parseTokenMapping(ctx.String(flags.TokenMappingBitgetFlag.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to parse bitget token mapping: %w", err)
	}
	if len(bitgetMapping) > 0 {
		cfg.TokenMappings[PriceFeedTypeBitget] = bitgetMapping
	}

	binanceMapping, err := parseTokenMapping(ctx.String(flags.TokenMappingBinanceFlag.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to parse binance token mapping: %w", err)
	}
	if len(binanceMapping) > 0 {
		cfg.TokenMappings[PriceFeedTypeBinance] = binanceMapping
	}

	// Parse API base URLs
	cfg.BitgetAPIBaseURL = ctx.String(flags.BitgetAPIBaseURLFlag.Name)
	cfg.BinanceAPIBaseURL = ctx.String(flags.BinanceAPIBaseURLFlag.Name)

	// Validate API URLs for configured feeds (non-empty check only)
	for _, feedType := range cfg.PriceFeedPriority {
		switch feedType {
		case PriceFeedTypeBitget:
			if cfg.BitgetAPIBaseURL == "" {
				return nil, fmt.Errorf("bitget feed is configured but --bitget-api-base-url is not set")
			}

		case PriceFeedTypeBinance:
			if cfg.BinanceAPIBaseURL == "" {
				return nil, fmt.Errorf("binance feed is configured but --binance-api-base-url is not set")
			}
		}
	}

	// Validate external sign config
	if cfg.ExternalSign {
		if cfg.ExternalSignAddress == "" || cfg.ExternalSignUrl == "" ||
			cfg.ExternalSignAppid == "" || cfg.ExternalSignChain == "" ||
			cfg.ExternalSignRsaPriv == "" {
			return nil, fmt.Errorf("external sign is enabled but missing required config: address=%s, url=%s, appid=%s, chain=%s, rsa_priv_set=%t",
				cfg.ExternalSignAddress, cfg.ExternalSignUrl, cfg.ExternalSignAppid, cfg.ExternalSignChain, cfg.ExternalSignRsaPriv != "")
		}

		// Validate address format
		if !common.IsHexAddress(cfg.ExternalSignAddress) {
			return nil, fmt.Errorf("invalid external sign address format: %s", cfg.ExternalSignAddress)
		}
	} else {
		// If not using external sign, private key is required
		if cfg.PrivateKey == "" {
			return nil, fmt.Errorf("private key is required when external sign is not enabled")
		}
	}

	return cfg, nil
}

// MetricAddress returns the metrics server address
func (c *Config) MetricAddress() string {
	return fmt.Sprintf("%s:%d", c.MetricsHostname, c.MetricsPort)
}
