package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/log"
)

const pythLatestPricePath = "/v2/updates/price/latest"

// pythMaxExponentMagnitude caps |expo| from a Hermes response. Real feeds sit well
// inside this range (typically -12..0).
const pythMaxExponentMagnitude = 32

// pythMaxClockSkew is how far ahead of the local clock a publish time may sit before
// it is rejected. Unlike the Chainlink path, which reads an L1 block timestamp and so
// is always in the past, publish_time is the publisher's near-real-time wall clock at
// second granularity. Rejecting anything ahead of the local clock would turn a small
// amount of host clock drift into a total Pyth outage, so allow a bounded skew while
// still refusing prices dated far into the future.
const pythMaxClockSkew = 30 * time.Second

// PythHermesPriceFeed reads Pyth prices from Hermes as an off-chain data source.
type PythHermesPriceFeed struct {
	httpClient       *http.Client
	mu               sync.RWMutex
	tokenPriceIDs    map[uint16]string
	ethUSDPriceID    string
	maxStaleness     time.Duration
	maxConfidenceBPS uint64
	baseURL          string
	apiKey           string
	log              log.Logger
}

// NewPythHermesPriceFeed creates a Pyth Hermes price feed.
func NewPythHermesPriceFeed(tokenPriceIDs map[uint16]string, baseURL string, apiKey string, ethUSDPriceID string, maxStaleness time.Duration, maxConfidenceBPS uint64) (*PythHermesPriceFeed, error) {
	ethUSDPriceID = normalizePythPriceID(ethUSDPriceID)
	if ethUSDPriceID == "" {
		return nil, fmt.Errorf("pyth price feed requires --pyth-eth-usd-price-id")
	}
	if maxStaleness <= 0 {
		return nil, fmt.Errorf("pyth max staleness must be positive")
	}
	// Hermes requires authentication from 2026-08-18 onwards, on both the current
	// and the upgraded endpoint. Unauthenticated requests still succeed until then,
	// but an updater configured today keeps running past the cutover, at which point
	// every Hermes request starts failing with nothing having changed on our side.
	// Refusing to start is the earlier and louder failure, so the key is required
	// rather than merely recommended.
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("pyth price feed requires --pyth-api-key")
	}

	normalized := make(map[uint16]string, len(tokenPriceIDs))
	for tokenID, priceID := range tokenPriceIDs {
		priceID = normalizePythPriceID(priceID)
		if priceID == "" {
			return nil, fmt.Errorf("invalid pyth price ID for token %d", tokenID)
		}
		normalized[tokenID] = priceID
	}
	if len(normalized) == 0 {
		return nil, fmt.Errorf("pyth price feed requires token mapping, please configure --token-mapping-pyth")
	}

	return &PythHermesPriceFeed{
		httpClient:       &http.Client{Timeout: 10 * time.Second},
		tokenPriceIDs:    normalized,
		ethUSDPriceID:    ethUSDPriceID,
		maxStaleness:     maxStaleness,
		maxConfidenceBPS: maxConfidenceBPS,
		baseURL:          strings.TrimRight(baseURL, "/"),
		apiKey:           apiKey,
		log:              log.New("component", "pyth_price_feed"),
	}, nil
}

// GetTokenPrice returns token price in USD from Pyth Hermes.
func (p *PythHermesPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	p.mu.RLock()
	priceID, exists := p.tokenPriceIDs[tokenID]
	ethUSDPriceID := p.ethUSDPriceID
	p.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("token ID %d not mapped to Pyth price ID", tokenID)
	}

	priceMap, err := p.fetchPrices(ctx, []string{ethUSDPriceID, priceID})
	if err != nil {
		return nil, err
	}

	ethPrice, err := pythPriceToFloat(priceMap[ethUSDPriceID])
	if err != nil {
		return nil, fmt.Errorf("failed to convert ETH/USD Pyth price: %w", err)
	}
	tokenPrice, err := pythPriceToFloat(priceMap[priceID])
	if err != nil {
		return nil, fmt.Errorf("failed to convert token Pyth price for token %d: %w", tokenID, err)
	}

	p.log.Info("Fetched price from Pyth",
		"source", "pyth",
		"token_id", tokenID,
		"price_id", priceID,
		"token_price_usd", tokenPrice.String(),
		"eth_price_usd", ethPrice.String())

	return &TokenPrice{
		TokenID:       tokenID,
		Symbol:        priceID,
		TokenPriceUSD: tokenPrice,
		EthPriceUSD:   ethPrice,
	}, nil
}

// GetBatchTokenPrices returns token prices in USD for multiple tokens. Tokens with no
// configured price ID are omitted rather than failing the batch, so that a mapping
// which covers only part of the requested set still contributes what it can and the
// caller can source the rest elsewhere.
func (p *PythHermesPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	p.mu.RLock()
	priceIDs := make([]string, 0, len(tokenIDs)+1)
	priceIDs = append(priceIDs, p.ethUSDPriceID)
	tokenPriceIDs := make(map[uint16]string, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		priceID, exists := p.tokenPriceIDs[tokenID]
		if !exists {
			continue
		}
		tokenPriceIDs[tokenID] = priceID
		priceIDs = append(priceIDs, priceID)
	}
	p.mu.RUnlock()

	if len(tokenPriceIDs) == 0 {
		return map[uint16]*TokenPrice{}, nil
	}

	priceMap, err := p.fetchPrices(ctx, priceIDs)
	if err != nil {
		return nil, err
	}

	ethPrice, err := pythPriceToFloat(priceMap[p.ethUSDPriceID])
	if err != nil {
		return nil, fmt.Errorf("failed to convert ETH/USD Pyth price: %w", err)
	}

	prices := make(map[uint16]*TokenPrice, len(tokenPriceIDs))
	for _, tokenID := range tokenIDs {
		priceID, exists := tokenPriceIDs[tokenID]
		if !exists {
			continue
		}
		tokenPrice, err := pythPriceToFloat(priceMap[priceID])
		if err != nil {
			return nil, fmt.Errorf("failed to convert token Pyth price for token %d: %w", tokenID, err)
		}
		prices[tokenID] = &TokenPrice{
			TokenID:       tokenID,
			Symbol:        priceID,
			TokenPriceUSD: tokenPrice,
			EthPriceUSD:   new(big.Float).Copy(ethPrice),
		}
	}

	return prices, nil
}

func (p *PythHermesPriceFeed) fetchPrices(ctx context.Context, priceIDs []string) (map[string]pythPrice, error) {
	values := url.Values{}
	values.Set("parsed", "true")
	values.Set("encoding", "hex")
	seen := make(map[string]struct{}, len(priceIDs))
	for _, priceID := range priceIDs {
		priceID = normalizePythPriceID(priceID)
		if _, exists := seen[priceID]; exists {
			continue
		}
		seen[priceID] = struct{}{}
		values.Add("ids[]", priceID)
	}

	requestURL := fmt.Sprintf("%s%s?%s", p.baseURL, pythLatestPricePath, values.Encode())
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + p.apiKey,
	}
	body, err := getJSONWithHeaders(ctx, p.httpClient, requestURL, headers)
	if err != nil {
		return nil, err
	}

	var resp pythLatestPriceResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse Pyth Hermes JSON response: %w", err)
	}

	priceMap := make(map[string]pythPrice, len(resp.Parsed))
	now := time.Now()
	for _, parsed := range resp.Parsed {
		priceID := normalizePythPriceID(parsed.ID)
		if err := validatePythPrice(parsed.Price, p.maxStaleness, p.maxConfidenceBPS, now); err != nil {
			return nil, fmt.Errorf("invalid Pyth price for %s: %w", priceID, err)
		}
		priceMap[priceID] = parsed.Price
	}

	for priceID := range seen {
		if _, exists := priceMap[priceID]; !exists {
			return nil, fmt.Errorf("Pyth response missing price ID %s", priceID)
		}
	}

	return priceMap, nil
}

type pythLatestPriceResponse struct {
	Parsed []pythParsedPrice `json:"parsed"`
}

type pythParsedPrice struct {
	ID    string    `json:"id"`
	Price pythPrice `json:"price"`
}

type pythPrice struct {
	Price       string `json:"price"`
	Confidence  string `json:"conf"`
	Exponent    int32  `json:"expo"`
	PublishTime int64  `json:"publish_time"`
}

func validatePythPrice(price pythPrice, maxStaleness time.Duration, maxConfidenceBPS uint64, now time.Time) error {
	priceInt, ok := new(big.Int).SetString(price.Price, 10)
	if !ok {
		return fmt.Errorf("invalid price integer %q", price.Price)
	}
	if priceInt.Sign() <= 0 {
		return fmt.Errorf("price must be positive, got %s", price.Price)
	}

	confInt, ok := new(big.Int).SetString(price.Confidence, 10)
	if !ok {
		return fmt.Errorf("invalid confidence integer %q", price.Confidence)
	}
	if confInt.Sign() < 0 {
		return fmt.Errorf("confidence must be non-negative, got %s", price.Confidence)
	}

	published := time.Unix(price.PublishTime, 0)
	if price.PublishTime <= 0 {
		return fmt.Errorf("publish_time must be positive")
	}
	if published.After(now.Add(pythMaxClockSkew)) {
		return fmt.Errorf("publish_time %s is more than %s in the future", published.UTC().Format(time.RFC3339), pythMaxClockSkew)
	}
	if now.Sub(published) > maxStaleness {
		return fmt.Errorf("price is stale: publish_time=%s maxStaleness=%s", published.UTC().Format(time.RFC3339), maxStaleness)
	}

	if maxConfidenceBPS > 0 {
		confBPS := new(big.Int).Mul(confInt, big.NewInt(10000))
		maxAllowed := new(big.Int).Mul(priceInt, new(big.Int).SetUint64(maxConfidenceBPS))
		if confBPS.Cmp(maxAllowed) > 0 {
			return fmt.Errorf("confidence too wide: conf=%s price=%s max_bps=%d", price.Confidence, price.Price, maxConfidenceBPS)
		}
	}

	return nil
}

func pythPriceToFloat(price pythPrice) (*big.Float, error) {
	priceInt, ok := new(big.Int).SetString(price.Price, 10)
	if !ok {
		return nil, fmt.Errorf("invalid price integer %q", price.Price)
	}

	value := new(big.Float).SetPrec(256).SetInt(priceInt)
	if price.Exponent == 0 {
		return value, nil
	}

	// Bound the magnitude before exponentiating. Exponent is an int32, so the
	// previous MaxInt32 check could never fire, and a hostile or malformed expo
	// of -2147483648 would ask big.Int to materialize 10^2147483648.
	exponent := int64(price.Exponent)
	magnitude := exponent
	if magnitude < 0 {
		magnitude = -magnitude
	}
	if magnitude > pythMaxExponentMagnitude {
		return nil, fmt.Errorf("pyth exponent out of range: %d", exponent)
	}

	scale := new(big.Float).SetPrec(256).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(magnitude), nil))
	if exponent > 0 {
		return value.Mul(value, scale), nil
	}
	return value.Quo(value, scale), nil
}

func normalizePythPriceID(priceID string) string {
	return strings.ToLower(strings.TrimPrefix(strings.TrimSpace(priceID), "0x"))
}
