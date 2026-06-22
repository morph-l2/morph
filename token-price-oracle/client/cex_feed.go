package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/log"
)

const (
	binanceTickerPath = "/api/v3/ticker/price"
	okxTickerPath     = "/api/v5/market/ticker"
)

type cexPriceFetcher func(ctx context.Context, httpClient *http.Client, baseURL string, symbol string) (*big.Float, error)

// CEXPriceFeed fetches token prices from a centralized exchange REST API.
type CEXPriceFeed struct {
	httpClient *http.Client
	mu         sync.RWMutex
	tokenMap   map[uint16]string
	ethSymbol  string
	ethPrice   *big.Float
	source     string
	log        log.Logger
	baseURL    string
	fetcher    cexPriceFetcher
}

// NewBinancePriceFeed creates a Binance REST price feed.
func NewBinancePriceFeed(tokenMap map[uint16]string, baseURL string) *CEXPriceFeed {
	return newCEXPriceFeed("binance", tokenMap, baseURL, "ETHUSDT", fetchBinancePrice)
}

// NewOKXPriceFeed creates an OKX REST price feed.
func NewOKXPriceFeed(tokenMap map[uint16]string, baseURL string) *CEXPriceFeed {
	return newCEXPriceFeed("okx", tokenMap, baseURL, "ETH-USDT", fetchOKXPrice)
}

func newCEXPriceFeed(source string, tokenMap map[uint16]string, baseURL string, ethSymbol string, fetcher cexPriceFetcher) *CEXPriceFeed {
	return &CEXPriceFeed{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		tokenMap:   tokenMap,
		ethSymbol:  ethSymbol,
		ethPrice:   big.NewFloat(0),
		source:     source,
		log:        log.New("component", source+"_price_feed"),
		baseURL:    baseURL,
		fetcher:    fetcher,
	}
}

// GetTokenPrice returns token price in USD.
func (f *CEXPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	f.mu.RLock()
	symbol, exists := f.tokenMap[tokenID]
	ethPrice := new(big.Float).Copy(f.ethPrice)
	f.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("token ID %d not mapped to %s trading pair", tokenID, f.source)
	}
	if ethPrice.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("ETH price not initialized, please call GetBatchTokenPrices first")
	}

	tokenPrice, err := f.fetchMappedPrice(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s price for %s: %w", f.source, symbol, err)
	}

	f.log.Info("Fetched price from CEX",
		"source", f.source,
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

// GetBatchTokenPrices returns batch token prices in USD.
func (f *CEXPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	if err := f.updateETHPrice(ctx); err != nil {
		return nil, fmt.Errorf("failed to update ETH price: %w", err)
	}

	prices := make(map[uint16]*TokenPrice, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		price, err := f.GetTokenPrice(ctx, tokenID)
		if err != nil {
			f.log.Warn("Failed to get price for token, skipping",
				"source", f.source,
				"token_id", tokenID,
				"error", err)
			continue
		}
		prices[tokenID] = price
	}
	return prices, nil
}

func (f *CEXPriceFeed) updateETHPrice(ctx context.Context) error {
	price, err := f.fetcher(ctx, f.httpClient, f.baseURL, f.ethSymbol)
	if err != nil {
		return fmt.Errorf("failed to fetch ETH price from %s: %w", f.source, err)
	}

	f.mu.Lock()
	f.ethPrice = price
	f.mu.Unlock()

	f.log.Info("Fetched ETH price from CEX",
		"source", f.source,
		"symbol", f.ethSymbol,
		"eth_price_usd", price.String())
	return nil
}

func (f *CEXPriceFeed) fetchMappedPrice(ctx context.Context, symbol string) (*big.Float, error) {
	if strings.HasPrefix(symbol, StablecoinPrefix) {
		return parseFixedStablecoinPrice(symbol)
	}
	return f.fetcher(ctx, f.httpClient, f.baseURL, symbol)
}

func parseFixedStablecoinPrice(symbol string) (*big.Float, error) {
	priceStr := strings.TrimPrefix(symbol, StablecoinPrefix)
	fixedPrice, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid stablecoin price format '%s': %w", symbol, err)
	}
	if fixedPrice <= 0 {
		return nil, fmt.Errorf("stablecoin price must be positive, got '%s'", symbol)
	}
	return big.NewFloat(fixedPrice), nil
}

type binanceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func fetchBinancePrice(ctx context.Context, httpClient *http.Client, baseURL string, symbol string) (*big.Float, error) {
	requestURL := fmt.Sprintf("%s%s?symbol=%s", strings.TrimRight(baseURL, "/"), binanceTickerPath, url.QueryEscape(symbol))
	body, err := getJSON(ctx, httpClient, requestURL)
	if err != nil {
		return nil, err
	}

	var resp binanceTickerResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse Binance JSON response: %w", err)
	}
	return parsePositiveFloat(resp.Price, symbol)
}

type okxTickerResponse struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []okxTickerRecord `json:"data"`
}

type okxTickerRecord struct {
	InstID string `json:"instId"`
	Last   string `json:"last"`
}

func fetchOKXPrice(ctx context.Context, httpClient *http.Client, baseURL string, symbol string) (*big.Float, error) {
	requestURL := fmt.Sprintf("%s%s?instId=%s", strings.TrimRight(baseURL, "/"), okxTickerPath, url.QueryEscape(symbol))
	body, err := getJSON(ctx, httpClient, requestURL)
	if err != nil {
		return nil, err
	}

	var resp okxTickerResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse OKX JSON response: %w", err)
	}
	if resp.Code != "0" {
		return nil, fmt.Errorf("OKX API error: %s - %s", resp.Code, resp.Msg)
	}
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no OKX ticker data returned for %s", symbol)
	}
	return parsePositiveFloat(resp.Data[0].Last, symbol)
}

func getJSON(ctx context.Context, httpClient *http.Client, requestURL string) ([]byte, error) {
	return getJSONWithHeaders(ctx, httpClient, requestURL, nil)
}

func getJSONWithHeaders(ctx context.Context, httpClient *http.Client, requestURL string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for name, value := range headers {
		req.Header.Set(name, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP status %d: %s", resp.StatusCode, string(body))
	}
	return body, nil
}

func parsePositiveFloat(priceStr string, symbol string) (*big.Float, error) {
	if priceStr == "" {
		return nil, fmt.Errorf("no price data returned for symbol %s", symbol)
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price '%s': %w", priceStr, err)
	}
	if price <= 0 {
		return nil, fmt.Errorf("price must be positive for symbol %s, got %s", symbol, priceStr)
	}
	return big.NewFloat(price), nil
}
