package updater

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/morph/bindings/bindings"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/config"
)

// CreatePriceUpdater creates price updater if conditions are met
func CreatePriceUpdater(
	cfg *config.Config,
	l2Client *client.L2Client,
	txManager *TxManager,
) (*PriceUpdater, error) {
	if cfg.L2TokenRegistryAddr == (common.Address{}) {
		return nil, fmt.Errorf("price update enabled but token registry address not set")
	}

	// Create registry contract
	registryContract, err := bindings.NewL2TokenRegistry(cfg.L2TokenRegistryAddr, l2Client.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create TokenRegistry contract: %w", err)
	}
	log.Info("TokenRegistry contract bound", "address", cfg.L2TokenRegistryAddr.Hex())

	// Create price feeds with fallback support
	priceFeed, err := createFallbackPriceFeed(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create price feed: %w", err)
	}

	// Collect all token mappings from configured feeds
	allTokenMappings := make(map[uint16]string)
	for _, feedType := range cfg.PriceFeedPriority {
		if mapping, exists := cfg.TokenMappings[feedType]; exists {
			for tokenID, symbol := range mapping {
				// Use first mapping found (highest priority)
				if _, alreadyMapped := allTokenMappings[tokenID]; !alreadyMapped {
					allTokenMappings[tokenID] = symbol
				}
			}
		}
	}

	// Create price updater
	priceUpdater := NewPriceUpdater(
		l2Client,
		registryContract,
		priceFeed,
		txManager,
		cfg.TokenIDs,
		allTokenMappings,
		cfg.PriceUpdateInterval,
		cfg.PriceThreshold,
	)

	log.Info("Price updater configured",
		"price_feed_priority", cfg.PriceFeedPriority,
		"token_ids", cfg.TokenIDs,
		"token_mappings", allTokenMappings,
		"interval", cfg.PriceUpdateInterval,
		"threshold", cfg.PriceThreshold)

	return priceUpdater, nil
}

// createFallbackPriceFeed creates price feed with fallback support
func createFallbackPriceFeed(cfg *config.Config) (client.PriceFeed, error) {
	if len(cfg.PriceFeedPriority) == 0 {
		return nil, fmt.Errorf("no price feeds configured in priority list")
	}

	var feeds []client.PriceFeed
	var feedNames []string

	for _, feedType := range cfg.PriceFeedPriority {
		feed, name, err := createSinglePriceFeed(feedType, cfg)
		if err != nil {
			log.Warn("Failed to create price feed, skipping",
				"feed_type", feedType,
				"error", err.Error())
			continue
		}
		feeds = append(feeds, feed)
		feedNames = append(feedNames, name)
	}

	if len(feeds) == 0 {
		return nil, fmt.Errorf("no valid price feeds could be created")
	}

	if len(feeds) == 1 {
		log.Info("Single price feed configured (no fallback)", "feed", feedNames[0])
		return feeds[0], nil
	}

	log.Info("Fallback price feed configured with multiple sources",
		"feeds", feedNames,
		"priority", "first to last")

	return client.NewFallbackPriceFeed(feeds, feedNames), nil
}

// createSinglePriceFeed creates a single price feed instance
func createSinglePriceFeed(feedType config.PriceFeedType, cfg *config.Config) (client.PriceFeed, string, error) {
	switch feedType {
	case config.PriceFeedTypeBitget:
		mapping, exists := cfg.TokenMappings[config.PriceFeedTypeBitget]
		if !exists || len(mapping) == 0 {
			return nil, "", fmt.Errorf("bitget price feed requires token mapping, please configure --token-mapping-bitget")
		}
		feed := client.NewBitgetSDKPriceFeed(mapping)
		log.Info("Bitget price feed created",
			"type", "bitget",
			"mapping", mapping)
		return feed, "bitget", nil

	case config.PriceFeedTypeBinance:
		mapping, exists := cfg.TokenMappings[config.PriceFeedTypeBinance]
		if !exists || len(mapping) == 0 {
			return nil, "", fmt.Errorf("binance price feed requires token mapping, please configure --token-mapping-binance")
		}
		// TODO: Implement Binance price feed when ready
		return nil, "", fmt.Errorf("binance price feed not yet implemented")

	default:
		return nil, "", fmt.Errorf("unsupported price feed type: %s", feedType)
	}
}

// CreateTxManager creates transaction manager
func CreateTxManager(l2Client *client.L2Client) *TxManager {
	return NewTxManager(l2Client)
}
