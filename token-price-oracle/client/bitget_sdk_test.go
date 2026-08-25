package client

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mirrors TestCEXGetTokenPriceInitializesETHPrice: both CEX implementations must
// satisfy the same precondition, i.e. a standalone GetTokenPrice self-initializes
// the ETH leg rather than failing on an unprimed cache.
func TestBitgetGetTokenPriceInitializesETHPrice(t *testing.T) {
	var ethRequests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		var price string
		switch symbol {
		case "ETHUSDT":
			ethRequests++
			price = "3000"
		case "BTCUSDT":
			price = "60000"
		default:
			t.Errorf("unexpected symbol: %s", symbol)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"code":"00000","msg":"","data":[{"symbol":%q,"lastPr":%q}]}`, symbol, price)
	}))
	defer server.Close()

	feed := NewBitgetSDKPriceFeed(map[uint16]string{1: "BTCUSDT"}, server.URL)

	price, err := feed.GetTokenPrice(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if price.EthPriceUSD.Cmp(big.NewFloat(3000)) != 0 {
		t.Fatalf("ETH price = %s, want 3000", price.EthPriceUSD.String())
	}
	if price.TokenPriceUSD.Cmp(big.NewFloat(60000)) != 0 {
		t.Fatalf("token price = %s, want 60000", price.TokenPriceUSD.String())
	}
	if ethRequests != 1 {
		t.Fatalf("ETH fetches = %d, want 1", ethRequests)
	}

	// The cached ETH leg is reused, so a second call must not re-fetch it.
	if _, err := feed.GetTokenPrice(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
	if ethRequests != 1 {
		t.Fatalf("ETH fetches after second call = %d, want 1", ethRequests)
	}
}

// The batch path primes the ETH leg once and every token in the cycle reuses it,
// keeping a cycle at N+1 requests rather than 2N.
func TestBitgetBatchFetchesETHPriceOncePerCycle(t *testing.T) {
	var ethRequests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		var price string
		switch symbol {
		case "ETHUSDT":
			ethRequests++
			price = "3000"
		case "BTCUSDT":
			price = "60000"
		case "SOLUSDT":
			price = "150"
		default:
			t.Errorf("unexpected symbol: %s", symbol)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"code":"00000","msg":"","data":[{"symbol":%q,"lastPr":%q}]}`, symbol, price)
	}))
	defer server.Close()

	feed := NewBitgetSDKPriceFeed(map[uint16]string{1: "BTCUSDT", 2: "SOLUSDT"}, server.URL)

	prices, err := feed.GetBatchTokenPrices(context.Background(), []uint16{1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(prices) != 2 {
		t.Fatalf("prices returned = %d, want 2", len(prices))
	}
	if ethRequests != 1 {
		t.Fatalf("ETH fetches for a 2-token cycle = %d, want 1", ethRequests)
	}
	for tokenID, price := range prices {
		if price.EthPriceUSD.Cmp(big.NewFloat(3000)) != 0 {
			t.Fatalf("token %d ETH price = %s, want 3000", tokenID, price.EthPriceUSD.String())
		}
	}
}
