package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"bitget/pkg/client/v2"

	"github.com/sirupsen/logrus"
)

// BitgetSDKPriceFeed uses official Bitget SDK to fetch prices
type BitgetSDKPriceFeed struct {
	client   *v2.SpotMarketClient
	tokenMap map[uint16]string
	ethPrice *big.Float
	log      *logrus.Entry
}

// BitgetV2Response represents Bitget V2 API response
type BitgetV2Response struct {
	Code        string           `json:"code"`
	Msg         string           `json:"msg"`
	RequestTime int64            `json:"requestTime"`
	Data        []BitgetV2Ticker `json:"data"`
}

// BitgetV2Ticker represents V2 ticker data
type BitgetV2Ticker struct {
	Symbol      string `json:"symbol"`
	LastPr      string `json:"lastPr"`
	High24h     string `json:"high24h"`
	Low24h      string `json:"low24h"`
	Change24h   string `json:"change24h"`
	BaseVolume  string `json:"baseVolume"`
	QuoteVolume string `json:"quoteVolume"`
}

// NewBitgetSDKPriceFeed creates a new Bitget SDK price feed
func NewBitgetSDKPriceFeed(tokenMap map[uint16]string) *BitgetSDKPriceFeed {
	client := new(v2.SpotMarketClient).Init()

	return &BitgetSDKPriceFeed{
		client:   client,
		tokenMap: tokenMap,
		ethPrice: big.NewFloat(0),
		log:      logrus.WithField("component", "bitget_sdk_price_feed"),
	}
}

// GetPriceRatio returns price ratio for tokenID
func (b *BitgetSDKPriceFeed) GetPriceRatio(ctx context.Context, tokenID uint16) (*big.Int, error) {
	symbol, exists := b.tokenMap[tokenID]
	if !exists {
		return nil, fmt.Errorf("token ID %d not mapped to trading pair", tokenID)
	}

	if b.ethPrice.Cmp(big.NewFloat(0)) == 0 {
		if err := b.updateETHPrice(ctx); err != nil {
			return nil, fmt.Errorf("failed to update ETH price: %w", err)
		}
	}

	tokenPrice, err := b.fetchPrice(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price for %s: %w", symbol, err)
	}

	if b.ethPrice.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("ETH price is zero")
	}

	ratio := new(big.Float).Quo(tokenPrice, b.ethPrice)
	ratio.Mul(ratio, big.NewFloat(1e18))

	priceRatio, _ := ratio.Int(nil)

	b.log.WithFields(logrus.Fields{
		"token_id":    tokenID,
		"symbol":      symbol,
		"token_price": tokenPrice.String(),
		"eth_price":   b.ethPrice.String(),
		"price_ratio": priceRatio.String(),
	}).Debug("Calculated price ratio")

	return priceRatio, nil
}

// GetBatchPriceRatios returns batch price ratios
func (b *BitgetSDKPriceFeed) GetBatchPriceRatios(ctx context.Context, tokenIDs []uint16) (map[uint16]*big.Int, error) {
	if err := b.updateETHPrice(ctx); err != nil {
		return nil, fmt.Errorf("failed to update ETH price: %w", err)
	}

	prices := make(map[uint16]*big.Int)

	for _, tokenID := range tokenIDs {
		price, err := b.GetPriceRatio(ctx, tokenID)
		if err != nil {
			b.log.WithFields(logrus.Fields{
				"token_id": tokenID,
				"error":    err,
			}).Warn("Failed to get price for token, skipping")
			continue
		}
		prices[tokenID] = price
	}

	return prices, nil
}

// updateETHPrice updates ETH price
func (b *BitgetSDKPriceFeed) updateETHPrice(ctx context.Context) error {
	price, err := b.fetchPrice(ctx, "ETHUSDT")
	if err != nil {
		return fmt.Errorf("failed to fetch ETH price: %w", err)
	}

	b.ethPrice = price
	b.log.WithField("eth_price", price.String()).Debug("Updated ETH price")

	return nil
}

// fetchPrice fetches price with retry
func (b *BitgetSDKPriceFeed) fetchPrice(ctx context.Context, symbol string) (*big.Float, error) {
	maxRetries := 3
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt) * time.Second
			b.log.WithFields(logrus.Fields{
				"symbol":  symbol,
				"attempt": attempt + 1,
				"backoff": backoff,
			}).Debug("Retrying fetch price")

			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		price, err := b.fetchPriceOnce(ctx, symbol)
		if err == nil {
			return price, nil
		}

		lastErr = err
		b.log.WithFields(logrus.Fields{
			"symbol":  symbol,
			"attempt": attempt + 1,
			"error":   err,
		}).Warn("Failed to fetch price, will retry")
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}

// fetchPriceOnce fetches price once using official Bitget SDK
func (b *BitgetSDKPriceFeed) fetchPriceOnce(ctx context.Context, symbol string) (*big.Float, error) {
	params := make(map[string]string)
	params["symbol"] = symbol

	resp, err := b.client.Tickers(params)
	if err != nil {
		return nil, fmt.Errorf("SDK request failed: %w", err)
	}

	var apiResp BitgetV2Response
	if err := json.Unmarshal([]byte(resp), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse SDK response: %w", err)
	}

	if apiResp.Code != "00000" {
		return nil, fmt.Errorf("API error: %s - %s", apiResp.Code, apiResp.Msg)
	}

	if len(apiResp.Data) == 0 {
		return nil, fmt.Errorf("no data returned for symbol %s", symbol)
	}

	lastPriceStr := apiResp.Data[0].LastPr
	if lastPriceStr == "" {
		return nil, fmt.Errorf("no price data returned for symbol %s", symbol)
	}

	lastPrice, err := strconv.ParseFloat(lastPriceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price '%s': %w", lastPriceStr, err)
	}

	b.log.WithFields(logrus.Fields{
		"symbol": symbol,
		"price":  lastPrice,
	}).Debug("Fetched price using Bitget SDK")

	return big.NewFloat(lastPrice), nil
}

// UpdateTokenMap updates token mapping
func (b *BitgetSDKPriceFeed) UpdateTokenMap(tokenMap map[uint16]string) {
	b.tokenMap = tokenMap
	b.log.WithField("token_map", tokenMap).Info("Updated token map")
}

// GetSupportedTokens returns list of supported token IDs
func (b *BitgetSDKPriceFeed) GetSupportedTokens() []uint16 {
	tokenIDs := make([]uint16, 0, len(b.tokenMap))
	for tokenID := range b.tokenMap {
		tokenIDs = append(tokenIDs, tokenID)
	}
	return tokenIDs
}
