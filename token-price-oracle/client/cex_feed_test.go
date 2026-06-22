package client

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchBinancePrice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != binanceTickerPath {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "BTCUSDT" {
			t.Fatalf("unexpected symbol: %s", r.URL.Query().Get("symbol"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"symbol":"BTCUSDT","price":"64385.12"}`))
	}))
	defer server.Close()

	price, err := fetchBinancePrice(context.Background(), server.Client(), server.URL, "BTCUSDT")
	if err != nil {
		t.Fatal(err)
	}
	if price.Cmp(big.NewFloat(64385.12)) != 0 {
		t.Fatalf("price = %s, want 64385.12", price.String())
	}
}

func TestFetchOKXPrice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != okxTickerPath {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("instId") != "BTC-USDT" {
			t.Fatalf("unexpected instId: %s", r.URL.Query().Get("instId"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"code":"0","msg":"","data":[{"instId":"BTC-USDT","last":"64386.45"}]}`))
	}))
	defer server.Close()

	price, err := fetchOKXPrice(context.Background(), server.Client(), server.URL, "BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	if price.Cmp(big.NewFloat(64386.45)) != 0 {
		t.Fatalf("price = %s, want 64386.45", price.String())
	}
}

func TestParseFixedStablecoinPrice(t *testing.T) {
	price, err := parseFixedStablecoinPrice("$1.0")
	if err != nil {
		t.Fatal(err)
	}
	if price.Cmp(big.NewFloat(1.0)) != 0 {
		t.Fatalf("price = %s, want 1", price.String())
	}
}
