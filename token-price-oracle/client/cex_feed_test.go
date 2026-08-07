package client

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParsePriceRejectsNonFiniteValues(t *testing.T) {
	// strconv.ParseFloat accepts these without error, and a bare `<= 0` test lets them
	// through: NaN then panics big.NewFloat, and an Inf survives until big.Float.Int
	// returns nil.
	for _, priceStr := range []string{"NaN", "nan", "Inf", "+Inf", "-Inf", "0", "-1"} {
		if _, err := parsePositiveFloat(priceStr, "BTCUSDT"); err == nil {
			t.Errorf("parsePositiveFloat(%q) succeeded, want error", priceStr)
		}
		if _, err := parseFixedStablecoinPrice(StablecoinPrefix + priceStr); err == nil {
			t.Errorf("parseFixedStablecoinPrice(%q) succeeded, want error", priceStr)
		}
	}

	price, err := parsePositiveFloat("60000.5", "BTCUSDT")
	if err != nil {
		t.Fatal(err)
	}
	if price.Cmp(big.NewFloat(60000.5)) != 0 {
		t.Fatalf("parsePositiveFloat() = %s, want 60000.5", price.String())
	}
}

func TestCEXGetTokenPriceInitializesETHPrice(t *testing.T) {
	fetcher := func(_ context.Context, _ *http.Client, _ string, symbol string) (*big.Float, error) {
		switch symbol {
		case "ETHUSDT":
			return big.NewFloat(3000), nil
		case "BTCUSDT":
			return big.NewFloat(60000), nil
		default:
			t.Fatalf("unexpected symbol: %s", symbol)
			return nil, nil
		}
	}
	feed := newCEXPriceFeed("test", map[uint16]string{1: "BTCUSDT"}, "", "ETHUSDT", fetcher)

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
}

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
