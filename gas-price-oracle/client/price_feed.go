package client

import (
	"context"
	"math/big"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// PriceFeed represents a price feed interface
type PriceFeed interface {
	// GetPriceRatio returns price ratio for tokenID
	// priceRatio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
	GetPriceRatio(ctx context.Context, tokenID uint16) (*big.Int, error)

	// GetBatchPriceRatios returns price ratios for multiple tokens
	GetBatchPriceRatios(ctx context.Context, tokenIDs []uint16) (map[uint16]*big.Int, error)
}

// MockPriceFeed simulates a price feed (for testing/development)
type MockPriceFeed struct {
	basePrice *big.Int
	variation float64
	log       *logrus.Entry
	rng       *rand.Rand
}

// NewMockPriceFeed creates a mock price feed
// basePrice: base price ratio (e.g. 1e18 means 1:1 ratio with ETH)
// variation: price variation percentage (e.g. 0.05 means ±5%)
func NewMockPriceFeed(basePrice *big.Int, variation float64) *MockPriceFeed {
	return &MockPriceFeed{
		basePrice: basePrice,
		variation: variation,
		log:       logrus.WithField("component", "mock_price_feed"),
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetPriceRatio returns a simulated price ratio
func (m *MockPriceFeed) GetPriceRatio(ctx context.Context, tokenID uint16) (*big.Int, error) {
	// Generate random variation
	randomFactor := 1.0 + (m.rng.Float64()*2-1)*m.variation

	// Calculate price with variation
	priceFloat := new(big.Float).SetInt(m.basePrice)
	priceFloat.Mul(priceFloat, big.NewFloat(randomFactor))

	price, _ := priceFloat.Int(nil)

	m.log.WithFields(logrus.Fields{
		"token_id":    tokenID,
		"price_ratio": price.String(),
		"variation":   randomFactor,
	}).Debug("Generated mock price")

	return price, nil
}

// GetBatchPriceRatios returns simulated price ratios for multiple tokens
func (m *MockPriceFeed) GetBatchPriceRatios(ctx context.Context, tokenIDs []uint16) (map[uint16]*big.Int, error) {
	prices := make(map[uint16]*big.Int)

	for _, tokenID := range tokenIDs {
		price, err := m.GetPriceRatio(ctx, tokenID)
		if err != nil {
			return nil, err
		}
		prices[tokenID] = price
	}

	return prices, nil
}

// SetBasePrice updates the base price
func (m *MockPriceFeed) SetBasePrice(basePrice *big.Int) {
	m.basePrice = basePrice
	m.log.WithField("base_price", basePrice.String()).Info("Updated base price")
}

// SetVariation updates the price variation
func (m *MockPriceFeed) SetVariation(variation float64) {
	m.variation = variation
	m.log.WithField("variation", variation).Info("Updated price variation")
}

