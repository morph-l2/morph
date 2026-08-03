package client

import (
	"context"
	"fmt"
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
			// Validate returned price to prevent nil pointer panics
			if price == nil || price.TokenPriceUSD == nil || price.EthPriceUSD == nil {
				f.log.Warn("Feed returned nil price or components, treating as failure",
					"token_id", tokenID,
					"feed", feedName,
					"priority", i)
				lastErr = fmt.Errorf("feed %s returned incomplete price for token %d", feedName, tokenID)
				continue
			}

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

// GetBatchTokenPrices resolves prices across feeds in priority order, passing each
// feed only the tokens still unresolved. Discarding a whole response because one
// token was missing meant a provider that covers only part of the active set could
// never contribute: with any active token absent from the Chainlink or Pyth mapping,
// those feeds failed every cycle and every token silently came from a CEX instead.
//
// A token no feed can price is reported to the caller by its absence from the
// returned map. Only a cycle that resolves nothing at all is an error.
func (f *FallbackPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	resolved := make(map[uint16]*TokenPrice, len(tokenIDs))
	pending := make([]uint16, len(tokenIDs))
	copy(pending, tokenIDs)
	var lastErr error

	for i, feed := range f.feeds {
		if len(pending) == 0 {
			break
		}

		feedName := "unknown"
		if i < len(f.names) {
			feedName = f.names[i]
		}

		prices, err := feed.GetBatchTokenPrices(ctx, pending)
		if err != nil {
			f.log.Warn("Failed to fetch batch prices from feed, trying next",
				"token_count", len(pending),
				"feed", feedName,
				"priority", i,
				"error", err.Error())
			lastErr = err
			continue
		}

		stillPending := make([]uint16, 0, len(pending))
		for _, tokenID := range pending {
			price, exists := prices[tokenID]
			if !exists || price == nil || price.TokenPriceUSD == nil || price.EthPriceUSD == nil {
				stillPending = append(stillPending, tokenID)
				continue
			}
			resolved[tokenID] = price
		}

		if resolvedHere := len(pending) - len(stillPending); resolvedHere > 0 {
			f.log.Info("Fetched batch prices from feed",
				"resolved_count", resolvedHere,
				"requested_count", len(pending),
				"feed", feedName,
				"priority", i)
		}
		if len(stillPending) > 0 {
			lastErr = fmt.Errorf("feed %s did not return prices for tokens %v", feedName, stillPending)
		}
		pending = stillPending
	}

	if len(resolved) == 0 {
		if lastErr != nil {
			return nil, fmt.Errorf("no price feed returned any of the %d requested tokens: %w", len(tokenIDs), lastErr)
		}
		return nil, fmt.Errorf("no price feed returned any of the %d requested tokens", len(tokenIDs))
	}

	if len(pending) > 0 {
		f.log.Warn("No price feed could resolve some tokens",
			"unresolved_token_ids", pending,
			"resolved_count", len(resolved),
			"requested_count", len(tokenIDs))
	}

	return resolved, nil
}
