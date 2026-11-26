package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const testResolvedUSDT = "0xdac17f958d2ee523a2206206994597c13d831ec7"

// newResolvedTestServer serves the batch endpoint from a table keyed by
// "chain/contract"; each value is the JSON tail of that item (price and flags).
func newResolvedTestServer(t *testing.T, items map[string]string, requests *int) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != resolvedBatchPricePath {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var req resolvedBatchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("bad request body: %v", err)
		}
		if req.Currency != "usd" {
			t.Errorf("currency = %q, want usd", req.Currency)
		}
		if requests != nil {
			*requests++
		}

		parts := make([]string, 0, len(req.List))
		for _, item := range req.List {
			tail, ok := items[item.Chain+"/"+item.Contract]
			if !ok {
				// What the real service returns for a token it cannot price.
				tail = `"price":0,"price_raw":0,"priceFrom":"bk","isSuspicious":false`
			}
			parts = append(parts, fmt.Sprintf(`{"chain":%q,"contract":%q,%s}`, item.Chain, item.Contract, tail))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":0,"data":[%s]}`, strings.Join(parts, ","))
	}))
}

func TestParseResolvedAsset(t *testing.T) {
	tests := []struct {
		in      string
		want    ResolvedAsset
		wantErr bool
	}{
		{in: "eth/" + testResolvedUSDT, want: ResolvedAsset{Chain: "eth", Contract: testResolvedUSDT}},
		{in: "bnb/", want: ResolvedAsset{Chain: "bnb"}},
		{in: "bnb", want: ResolvedAsset{Chain: "bnb"}},
		{in: "ETH/0xdAC17F958D2ee523a2206206994597C13D831ec7", want: ResolvedAsset{Chain: "eth", Contract: "0xdAC17F958D2ee523a2206206994597C13D831ec7"}},
		{in: " cosmos / ibc/ABC ", want: ResolvedAsset{Chain: "cosmos", Contract: "ibc/ABC"}},
		{in: "", wantErr: true},
		{in: "/0xabc", wantErr: true},
	}
	for _, tt := range tests {
		got, err := ParseResolvedAsset(tt.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseResolvedAsset(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseResolvedAsset(%q) = %+v, want %+v", tt.in, got, tt.want)
		}
	}
}

func TestNewResolvedPriceFeedValidatesConfig(t *testing.T) {
	mapping := map[uint16]string{1: "eth/" + testResolvedUSDT}
	if _, err := NewResolvedPriceFeed(mapping, "", "eth/"); err == nil {
		t.Error("empty base URL accepted")
	}
	if _, err := NewResolvedPriceFeed(mapping, "http://example.com", ""); err == nil {
		t.Error("empty ETH asset accepted")
	}
	if _, err := NewResolvedPriceFeed(nil, "http://example.com", "eth/"); err == nil {
		t.Error("empty mapping accepted")
	}
	if _, err := NewResolvedPriceFeed(map[uint16]string{1: "/0xabc"}, "http://example.com", "eth/"); err == nil {
		t.Error("mapping without chain accepted")
	}
}

func TestResolvedGetBatchTokenPrices(t *testing.T) {
	requests := 0
	server := newResolvedTestServer(t, map[string]string{
		"eth/":                    `"price":2000.5,"price_raw":2000.5,"priceFrom":"bk","isSuspicious":false`,
		"eth/" + testResolvedUSDT: `"price":1.0001,"price_raw":1.0001,"priceFrom":"binance","isSuspicious":false`,
		"eth/0xsuspicious":        `"price":0,"price_raw":3000,"priceFrom":"bk","isSuspicious":true`,
		"bnb/":                    `"price":1e-7,"price_raw":1e-7,"priceFrom":"bk","isSuspicious":false`,
	}, &requests)
	defer server.Close()

	feed, err := NewResolvedPriceFeed(map[uint16]string{
		1: "eth/" + testResolvedUSDT,
		2: "eth/0xsuspicious",
		3: "eth/0xunknown",
		4: "bnb",
	}, server.URL+"/", "eth/")
	if err != nil {
		t.Fatal(err)
	}

	// Token 9 has no mapping; it must be omitted rather than fail the batch.
	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2, 3, 4, 9})
	if err != nil {
		t.Fatal(err)
	}
	if requests != 1 {
		t.Errorf("requests = %d, want 1 (ETH and tokens share one batch call)", requests)
	}
	if len(prices) != 2 {
		t.Fatalf("resolved %d tokens, want 2 (suspicious, unpriced and unmapped tokens omitted): %v", len(prices), prices)
	}

	usdt := prices[1]
	if usdt == nil || usdt.TokenPriceUSD.Text('f', 4) != "1.0001" || usdt.EthPriceUSD.Text('f', 1) != "2000.5" {
		t.Errorf("token 1 = %+v, want token 1.0001 / eth 2000.5", usdt)
	}
	if usdt != nil && usdt.Symbol != "eth/"+testResolvedUSDT {
		t.Errorf("token 1 symbol = %q", usdt.Symbol)
	}
	if bnb := prices[4]; bnb == nil || bnb.TokenPriceUSD.Cmp(big.NewFloat(0)) <= 0 || bnb.TokenPriceUSD.Text('e', 0) != "1e-07" {
		t.Errorf("token 4 = %+v, want 1e-7 parsed from exponent notation", bnb)
	}
}

func TestResolvedGetBatchTokenPricesNoMappedTokens(t *testing.T) {
	requests := 0
	server := newResolvedTestServer(t, nil, &requests)
	defer server.Close()
	feed, err := NewResolvedPriceFeed(map[uint16]string{1: "eth/" + testResolvedUSDT}, server.URL, "eth/")
	if err != nil {
		t.Fatal(err)
	}

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{7, 8})
	if err != nil || len(prices) != 0 {
		t.Fatalf("got %v, %v; want empty map and nil error", prices, err)
	}
	if requests != 0 {
		t.Errorf("requests = %d, want 0 when no requested token is mapped", requests)
	}
}

func TestResolvedUnusableETHPriceFailsBatch(t *testing.T) {
	for name, ethTail := range map[string]string{
		"suspicious": `"price":0,"price_raw":2000,"priceFrom":"bk","isSuspicious":true`,
		"zero":       `"price":0,"price_raw":0,"priceFrom":"bk","isSuspicious":false`,
		"missing":    `"priceFrom":"bk","isSuspicious":false`,
	} {
		t.Run(name, func(t *testing.T) {
			server := newResolvedTestServer(t, map[string]string{
				"eth/":                    ethTail,
				"eth/" + testResolvedUSDT: `"price":1,"price_raw":1,"priceFrom":"bk","isSuspicious":false`,
			}, nil)
			defer server.Close()
			feed, err := NewResolvedPriceFeed(map[uint16]string{1: "eth/" + testResolvedUSDT}, server.URL, "eth/")
			if err != nil {
				t.Fatal(err)
			}
			if _, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1}); err == nil {
				t.Fatal("batch succeeded without a usable ETH price")
			}
		})
	}
}

func TestResolvedGetTokenPrice(t *testing.T) {
	server := newResolvedTestServer(t, map[string]string{
		"eth/":                    `"price":2000,"price_raw":2000,"priceFrom":"bk","isSuspicious":false`,
		"eth/" + testResolvedUSDT: `"price":1,"price_raw":1,"priceFrom":"bk","isSuspicious":false`,
	}, nil)
	defer server.Close()
	feed, err := NewResolvedPriceFeed(map[uint16]string{1: "eth/" + testResolvedUSDT, 2: "eth/0xunknown"}, server.URL, "eth/")
	if err != nil {
		t.Fatal(err)
	}

	price, err := feed.GetTokenPrice(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if price.TokenPriceUSD.Cmp(big.NewFloat(1)) != 0 || price.EthPriceUSD.Cmp(big.NewFloat(2000)) != 0 {
		t.Errorf("price = %+v", price)
	}
	if _, err := feed.GetTokenPrice(context.Background(), 2); err == nil {
		t.Error("unpriced token returned a price")
	}
	if _, err := feed.GetTokenPrice(context.Background(), 3); err == nil {
		t.Error("unmapped token returned a price")
	}
}

func TestResolvedRejectsBadResponses(t *testing.T) {
	tests := []struct {
		name   string
		status int
		body   string
	}{
		{name: "api error status", status: 200, body: `{"status":1,"error":"chain参数不能为空"}`},
		{name: "service unavailable", status: 503, body: `{"status":1,"error":"resolved price service unavailable"}`},
		{name: "truncated list", status: 200, body: `{"status":0,"data":[{"chain":"eth","contract":"","price":2000,"priceFrom":"bk","isSuspicious":false}]}`},
		{
			name:   "reordered list",
			status: 200,
			body: `{"status":0,"data":[` +
				`{"chain":"eth","contract":"` + testResolvedUSDT + `","price":1,"priceFrom":"bk","isSuspicious":false},` +
				`{"chain":"eth","contract":"","price":2000,"priceFrom":"bk","isSuspicious":false}]}`,
		},
		{name: "not json", status: 200, body: `<html>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				fmt.Fprint(w, tt.body)
			}))
			defer server.Close()
			feed, err := NewResolvedPriceFeed(map[uint16]string{1: "eth/" + testResolvedUSDT}, server.URL, "eth/")
			if err != nil {
				t.Fatal(err)
			}
			if _, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1}); err == nil {
				t.Fatal("bad response accepted")
			}
		})
	}
}

// TestResolvedLive queries a real resolved price service. It is skipped unless
// TOKEN_PRICE_ORACLE_RESOLVED_PRICE_API_BASE_URL is set.
func TestResolvedLive(t *testing.T) {
	baseURL := os.Getenv("TOKEN_PRICE_ORACLE_RESOLVED_PRICE_API_BASE_URL")
	if baseURL == "" {
		t.Skip("TOKEN_PRICE_ORACLE_RESOLVED_PRICE_API_BASE_URL not set")
	}
	feed, err := NewResolvedPriceFeed(map[uint16]string{
		1: "eth/" + testResolvedUSDT,
		2: "bnb",
		3: "eth/0x000000000000000000000000000000000000dead", // unpriced: must be omitted
	}, baseURL, "eth/")
	if err != nil {
		t.Fatal(err)
	}

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, tokenID := range []uint16{1, 2} {
		price, ok := prices[tokenID]
		if !ok {
			t.Errorf("token %d not resolved", tokenID)
			continue
		}
		t.Logf("token %d (%s): token=%s eth=%s", tokenID, price.Symbol,
			price.TokenPriceUSD.Text('g', 12), price.EthPriceUSD.Text('g', 12))
	}
	if _, ok := prices[3]; ok {
		t.Errorf("unpriced token resolved: %+v", prices[3])
	}
}

func TestResolvedSplitsLargeBatches(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req resolvedBatchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("bad request body: %v", err)
		}
		if len(req.List) > resolvedMaxBatchSize {
			t.Errorf("batch of %d exceeds the %d limit", len(req.List), resolvedMaxBatchSize)
		}
		requests++
		parts := make([]string, 0, len(req.List))
		for _, item := range req.List {
			parts = append(parts, fmt.Sprintf(`{"chain":%q,"contract":%q,"price":2,"priceFrom":"bk","isSuspicious":false}`, item.Chain, item.Contract))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":0,"data":[%s]}`, strings.Join(parts, ","))
	}))
	defer server.Close()

	const tokenCount = resolvedMaxBatchSize + 5
	mapping := make(map[uint16]string, tokenCount)
	tokenIDs := make([]uint16, 0, tokenCount)
	for i := 1; i <= tokenCount; i++ {
		mapping[uint16(i)] = fmt.Sprintf("eth/0x%040x", i)
		tokenIDs = append(tokenIDs, uint16(i))
	}
	feed, err := NewResolvedPriceFeed(mapping, server.URL, "eth/")
	if err != nil {
		t.Fatal(err)
	}

	prices, err := feed.GetBatchTokenPrices(context.Background(), tokenIDs)
	if err != nil {
		t.Fatal(err)
	}
	if len(prices) != tokenCount || requests != 2 {
		t.Fatalf("resolved %d tokens in %d requests, want %d tokens in 2 requests", len(prices), requests, tokenCount)
	}
}
