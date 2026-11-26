package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/log"
)

const (
	resolvedBatchPricePath = "/api/v1/price/resolved/batch"
	// resolvedMaxBatchSize is the largest list the batch endpoint accepts; a longer
	// list is rejected with HTTP 400.
	resolvedMaxBatchSize = 1000
	resolvedCurrency     = "usd"
)

// ResolvedAsset identifies a token the way the resolved price API does: by chain
// name and contract address. A native coin has an empty contract.
type ResolvedAsset struct {
	Chain    string
	Contract string
}

// String renders the asset in its configuration form, "chain/contract".
func (a ResolvedAsset) String() string {
	return a.Chain + "/" + a.Contract
}

// ParseResolvedAsset parses "chain/contract". Only the first slash separates the two,
// so a contract that itself contains slashes survives. "chain" and "chain/" both mean
// the chain's native coin. The chain name is lowercased, because the service echoes it
// lowercased and the echo is compared exactly; the contract is kept as written, because
// the service echoes it unchanged and non-EVM addresses are case-sensitive.
func ParseResolvedAsset(value string) (ResolvedAsset, error) {
	chain, contract, _ := strings.Cut(strings.TrimSpace(value), "/")
	chain = strings.ToLower(strings.TrimSpace(chain))
	if chain == "" {
		return ResolvedAsset{}, fmt.Errorf("invalid resolved price asset %q (expected format: chain/contract)", value)
	}
	return ResolvedAsset{Chain: chain, Contract: strings.TrimSpace(contract)}, nil
}

// ResolvedPriceFeed reads token prices from the resolved price API. The service picks
// the underlying source (on-chain or an exchange) per token, so the feed only supplies
// chain and contract.
type ResolvedPriceFeed struct {
	httpClient  *http.Client
	tokenAssets map[uint16]ResolvedAsset
	ethAsset    ResolvedAsset
	baseURL     string
	log         log.Logger
}

// NewResolvedPriceFeed creates a resolved price API feed. tokenMap maps token IDs to
// "chain/contract"; ethAsset is the asset priced as ETH/USD, in the same form.
func NewResolvedPriceFeed(tokenMap map[uint16]string, baseURL string, ethAsset string) (*ResolvedPriceFeed, error) {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if baseURL == "" {
		return nil, fmt.Errorf("resolved price feed requires --resolved-price-api-base-url")
	}
	eth, err := ParseResolvedAsset(ethAsset)
	if err != nil {
		return nil, fmt.Errorf("invalid --resolved-price-eth-asset: %w", err)
	}
	if len(tokenMap) == 0 {
		return nil, fmt.Errorf("resolved price feed requires token mapping, please configure --token-mapping-resolved")
	}

	tokenAssets := make(map[uint16]ResolvedAsset, len(tokenMap))
	for tokenID, value := range tokenMap {
		asset, err := ParseResolvedAsset(value)
		if err != nil {
			return nil, fmt.Errorf("token %d: %w", tokenID, err)
		}
		tokenAssets[tokenID] = asset
	}

	return &ResolvedPriceFeed{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		tokenAssets: tokenAssets,
		ethAsset:    eth,
		baseURL:     baseURL,
		log:         log.New("component", "resolved_price_feed"),
	}, nil
}

// GetTokenPrice returns token price in USD.
func (f *ResolvedPriceFeed) GetTokenPrice(ctx context.Context, tokenID uint16) (*TokenPrice, error) {
	if _, exists := f.tokenAssets[tokenID]; !exists {
		return nil, fmt.Errorf("token ID %d not mapped to a resolved price asset", tokenID)
	}
	prices, err := f.GetBatchTokenPrices(ctx, []uint16{tokenID})
	if err != nil {
		return nil, err
	}
	price, exists := prices[tokenID]
	if !exists {
		return nil, fmt.Errorf("resolved price API returned no usable price for token %d", tokenID)
	}
	return price, nil
}

// GetBatchTokenPrices returns token prices in USD for multiple tokens. The ETH leg and
// every token are read in the same request, so all ratios in a cycle share one sample.
// Tokens that are unmapped, flagged suspicious or unpriced are omitted so the caller
// can source them from another feed; an unusable ETH price fails the whole batch,
// because no ratio can be computed without it.
func (f *ResolvedPriceFeed) GetBatchTokenPrices(ctx context.Context, tokenIDs []uint16) (map[uint16]*TokenPrice, error) {
	assets := []ResolvedAsset{f.ethAsset}
	seen := map[ResolvedAsset]struct{}{f.ethAsset: {}}
	mapped := make([]uint16, 0, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		asset, exists := f.tokenAssets[tokenID]
		if !exists {
			continue
		}
		mapped = append(mapped, tokenID)
		if _, dup := seen[asset]; !dup {
			seen[asset] = struct{}{}
			assets = append(assets, asset)
		}
	}
	if len(mapped) == 0 {
		return map[uint16]*TokenPrice{}, nil
	}

	quotes, err := f.fetchQuotes(ctx, assets)
	if err != nil {
		return nil, err
	}

	ethQuote := quotes[f.ethAsset]
	ethPrice, err := ethQuote.usablePrice()
	if err != nil {
		return nil, fmt.Errorf("unusable ETH price from resolved price API (%s): %w", f.ethAsset, err)
	}

	prices := make(map[uint16]*TokenPrice, len(mapped))
	for _, tokenID := range mapped {
		asset := f.tokenAssets[tokenID]
		quote := quotes[asset]
		tokenPrice, err := quote.usablePrice()
		if err != nil {
			f.log.Warn("Unusable price from resolved price API, skipping",
				"token_id", tokenID,
				"asset", asset.String(),
				"price_from", quote.PriceFrom,
				"error", err)
			continue
		}

		f.log.Info("Fetched price from resolved price API",
			"source", "resolved",
			"token_id", tokenID,
			"asset", asset.String(),
			"price_from", quote.PriceFrom,
			"token_price_usd", tokenPrice.String(),
			"eth_price_usd", ethPrice.String(),
			"eth_price_from", ethQuote.PriceFrom)

		prices[tokenID] = &TokenPrice{
			TokenID:       tokenID,
			Symbol:        asset.String(),
			TokenPriceUSD: tokenPrice,
			EthPriceUSD:   new(big.Float).Copy(ethPrice),
		}
	}
	return prices, nil
}

type resolvedBatchRequest struct {
	Currency string                 `json:"currency"`
	List     []resolvedBatchReqItem `json:"list"`
}

type resolvedBatchReqItem struct {
	Chain    string `json:"chain"`
	Contract string `json:"contract"`
}

type resolvedBatchResponse struct {
	Status int             `json:"status"`
	Error  string          `json:"error"`
	Data   []resolvedQuote `json:"data"`
}

type resolvedQuote struct {
	Chain        string      `json:"chain"`
	Contract     string      `json:"contract"`
	Price        json.Number `json:"price"`
	PriceFrom    string      `json:"priceFrom"`
	IsSuspicious bool        `json:"isSuspicious"`
}

// usablePrice returns the quote's price, or an error when it must not be used. The
// service reports a token it cannot price as price 0 with a success status, and masks
// a suspicious price to 0 as well, so a non-positive price always means "no price".
func (q resolvedQuote) usablePrice() (*big.Float, error) {
	if q.IsSuspicious {
		return nil, fmt.Errorf("price is flagged suspicious")
	}
	price, ok := new(big.Float).SetPrec(256).SetString(q.Price.String())
	if !ok {
		return nil, fmt.Errorf("invalid price %q", q.Price.String())
	}
	if price.IsInf() || price.Sign() <= 0 {
		return nil, fmt.Errorf("price must be a positive finite number, got %s", q.Price.String())
	}
	return price, nil
}

func (f *ResolvedPriceFeed) fetchQuotes(ctx context.Context, assets []ResolvedAsset) (map[ResolvedAsset]resolvedQuote, error) {
	quotes := make(map[ResolvedAsset]resolvedQuote, len(assets))
	for start := 0; start < len(assets); start += resolvedMaxBatchSize {
		end := min(start+resolvedMaxBatchSize, len(assets))
		if err := f.fetchQuoteChunk(ctx, assets[start:end], quotes); err != nil {
			return nil, err
		}
	}
	return quotes, nil
}

func (f *ResolvedPriceFeed) fetchQuoteChunk(ctx context.Context, assets []ResolvedAsset, quotes map[ResolvedAsset]resolvedQuote) error {
	reqBody := resolvedBatchRequest{Currency: resolvedCurrency, List: make([]resolvedBatchReqItem, len(assets))}
	for i, asset := range assets {
		reqBody.List[i] = resolvedBatchReqItem(asset)
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to encode resolved price request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, f.baseURL+resolvedBatchPricePath, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := doJSONRequest(f.httpClient, req)
	if err != nil {
		return err
	}

	var resp resolvedBatchResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("failed to parse resolved price JSON response: %w", err)
	}
	if resp.Status != 0 {
		return fmt.Errorf("resolved price API error: status=%d error=%s", resp.Status, resp.Error)
	}
	// The response carries no key other than position plus the echoed chain and
	// contract. Check both, so that a reordered or truncated response cannot attach
	// one token's price to another.
	if len(resp.Data) != len(assets) {
		return fmt.Errorf("resolved price API returned %d items for %d requested", len(resp.Data), len(assets))
	}
	for i, asset := range assets {
		quote := resp.Data[i]
		if quote.Chain != asset.Chain || quote.Contract != asset.Contract {
			return fmt.Errorf("resolved price API item %d is for %s/%s, requested %s", i, quote.Chain, quote.Contract, asset)
		}
		quotes[asset] = quote
	}
	return nil
}
