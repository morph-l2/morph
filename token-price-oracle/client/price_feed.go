package client

import (
	"context"
	"math/big"

	"github.com/morph-l2/go-ethereum/log"
)

// TokenPrice represents token price information
type TokenPrice struct {
	TokenID       uint16
	Symbol        string
	TokenPriceUSD *big.Float // Token price in USD
	EthPriceUSD   *big.Float // ETH price in USD (for reference)
}

// PriceFeed represents a price feed interface
type PriceFeed interface {
	// GetTokenPrice returns token price in USD
	GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error)

	// GetBatchTokenPrices returns token prices in USD for multiple tokens
	GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error)
}

// FallbackPriceFeed implements fallback mechanism for multiple price feeds
type FallbackPriceFeed struct {
	feeds []PriceFeed
	names []string // Feed names for logging
	log   log.Logger
}

// NewFallbackPriceFeed creates a price feed with fallback support
// feeds: price feeds in priority order (first = highest priority)
// names: corresponding names for logging
func NewFallbackPriceFeed(feeds []PriceFeed, names []string) *FallbackPriceFeed {
	return &FallbackPriceFeed{
		feeds: feeds,
		names: names,
		log:   log.New("component", "fallback_price_feed"),
	}
}

// GetTokenPrice tries to get token price from feeds in priority order
func (f *FallbackPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	var lastErr error

	for i, feed := range f.feeds {
		feedName := "unknown"
		if i < len(f.names) {
			feedName = f.names[i]
		}

		price, err := feed.GetTokenPrice(ctx, tokenID)
		if err == nil {
			f.log.Info("Successfully fetched price from feed",
				"source", feedName,
				"token_id", tokenID,
				"symbol", price.Symbol,
				"priority", i,
				"token_price_usd", price.TokenPriceUSD.String(),
				"eth_price_usd", price.EthPriceUSD.String())
			return price, nil
		}

		f.log.Warn("Failed to fetch price from feed, trying next",
			"token_id", tokenID,
			"feed", feedName,
			"priority", i,
			"error", err.Error())
		lastErr = err
	}

	return nil, lastErr
}

// GetBatchTokenPrices tries to get batch token prices from feeds in priority order
func (f *FallbackPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	var lastErr error

	for i, feed := range f.feeds {
		feedName := "unknown"
		if i < len(f.names) {
			feedName = f.names[i]
		}

		prices, err := feed.GetBatchTokenPrices(ctx, tokenIDs)
		if err == nil {
			f.log.Info("Successfully fetched batch prices from feed",
				"token_count", len(tokenIDs),
				"feed", feedName,
				"priority", i)
			return prices, nil
		}

		f.log.Warn("Failed to fetch batch prices from feed, trying next",
			"token_count", len(tokenIDs),
			"feed", feedName,
			"priority", i,
			"error", err.Error())
		lastErr = err
	}

	return nil, lastErr
}

