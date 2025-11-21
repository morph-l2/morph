package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/log"
)

const (
	bitgetTickerPath = "/api/v2/spot/market/tickers"
)

// BitgetSDKPriceFeed uses Bitget REST API to fetch prices
// This type is safe for concurrent use by multiple goroutines
type BitgetSDKPriceFeed struct {
	httpClient *http.Client
	mu         sync.RWMutex          // protects tokenMap and ethPrice
	tokenMap   map[uint16]string     // guarded by mu
	ethPrice   *big.Float            // guarded by mu
	log        log.Logger
	baseURL    string
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

// NewBitgetSDKPriceFeed creates a new Bitget price feed using REST API
func NewBitgetSDKPriceFeed(tokenMap map[uint16]string, baseURL string) *BitgetSDKPriceFeed {
	return &BitgetSDKPriceFeed{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		tokenMap: tokenMap,
		ethPrice: big.NewFloat(0),
		log:      log.New("component", "bitget_price_feed"),
		baseURL:  baseURL,
	}
}

// GetTokenPrice returns token price in USD
// Note: Caller should ensure ETH price is updated via GetBatchTokenPrices for batch operations
func (b *BitgetSDKPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	b.mu.RLock()
	symbol, exists := b.tokenMap[tokenID]
	ethPrice := new(big.Float).Copy(b.ethPrice)
	b.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("token ID %d not mapped to trading pair", tokenID)
	}

	// Fetch token price
	tokenPrice, err := b.fetchPrice(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price for %s: %w", symbol, err)
	}

	// Use cached ETH price (should be updated by GetBatchTokenPrices)
	if ethPrice.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("ETH price not initialized, please call GetBatchTokenPrices first")
	}

	b.log.Info("Fetched price from Bitget",
		"source", "bitget",
		"token_id", tokenID,
		"symbol", symbol,
		"token_price_usd", tokenPrice.String(),
		"eth_price_usd", ethPrice.String())

	return &TokenPrice{
		TokenID:       tokenID,
		Symbol:        symbol,
		TokenPriceUSD: tokenPrice,
		EthPriceUSD:   ethPrice,
	}, nil
}

// GetBatchTokenPrices returns batch token prices in USD
func (b *BitgetSDKPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	// Update ETH price first (this will acquire write lock)
	if err := b.updateETHPrice(ctx); err != nil {
		return nil, fmt.Errorf("failed to update ETH price: %w", err)
	}

	prices := make(map[uint16]*TokenPrice)

	for _, tokenID := range tokenIDs {
		price, err := b.GetTokenPrice(ctx, tokenID)
		if err != nil {
			b.log.Warn("Failed to get price for token, skipping",
				"token_id", tokenID,
				"error", err)
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

	b.mu.Lock()
	b.ethPrice = price
	b.mu.Unlock()

	b.log.Info("Fetched ETH price from Bitget",
		"source", "bitget",
		"symbol", "ETHUSDT",
		"eth_price_usd", price.String())

	return nil
}

// fetchPrice fetches price with retry
func (b *BitgetSDKPriceFeed) fetchPrice(ctx context.Context, symbol string) (*big.Float, error) {
	maxRetries := 3
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt) * time.Second
			b.log.Debug("Retrying fetch price",
				"symbol", symbol,
				"attempt", attempt+1,
				"backoff", backoff)

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
		b.log.Warn("Failed to fetch price, will retry",
			"symbol", symbol,
			"attempt", attempt+1,
			"error", err)
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}

// fetchPriceOnce fetches price once using Bitget REST API
func (b *BitgetSDKPriceFeed) fetchPriceOnce(ctx context.Context, symbol string) (*big.Float, error) {
	// Build request URL
	url := fmt.Sprintf("%s%s?symbol=%s", b.baseURL, bitgetTickerPath, symbol)

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var apiResp BitgetV2Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Check API response code
	if apiResp.Code != "00000" {
		return nil, fmt.Errorf("API error: %s - %s", apiResp.Code, apiResp.Msg)
	}

	// Check if data exists
	if len(apiResp.Data) == 0 {
		return nil, fmt.Errorf("no data returned for symbol %s", symbol)
	}

	// Parse price
	lastPriceStr := apiResp.Data[0].LastPr
	if lastPriceStr == "" {
		return nil, fmt.Errorf("no price data returned for symbol %s", symbol)
	}

	lastPrice, err := strconv.ParseFloat(lastPriceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price '%s': %w", lastPriceStr, err)
	}

	b.log.Debug("Fetched price from Bitget API",
		"symbol", symbol,
		"price", lastPrice)

	return big.NewFloat(lastPrice), nil
}

// UpdateTokenMap updates token mapping
// This method is safe to call concurrently with other methods
// The input map is copied to prevent external modifications
func (b *BitgetSDKPriceFeed) UpdateTokenMap(tokenMap map[uint16]string) {
	b.mu.Lock()
	// Create a defensive copy to prevent external modifications
	copied := make(map[uint16]string, len(tokenMap))
	for k, v := range tokenMap {
		copied[k] = v
	}
	b.tokenMap = copied
	b.mu.Unlock()
	b.log.Info("Updated token map", "token_map", copied)
}

// GetSupportedTokens returns list of supported token IDs
func (b *BitgetSDKPriceFeed) GetSupportedTokens() []uint16 {
	b.mu.RLock()
	tokenIDs := make([]uint16, 0, len(b.tokenMap))
	for tokenID := range b.tokenMap {
		tokenIDs = append(tokenIDs, tokenID)
	}
	b.mu.RUnlock()
	return tokenIDs
}
