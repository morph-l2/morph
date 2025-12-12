package client

import (
	"context"
	"math/big"
	"testing"
	"time"
)

// TestBitgetSDK_FetchETHPrice tests fetching ETH price using official Bitget SDK
func TestBitgetSDK_FetchETHPrice(t *testing.T) {
	feed := NewBitgetSDKPriceFeed(map[uint16]string{
		1: "ETHUSDT",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Test fetching ETH price directly
	price, err := feed.fetchPrice(ctx, "ETHUSDT")
	if err != nil {
		t.Fatalf("Failed to fetch ETH price: %v", err)
	}

	if price == nil {
		t.Fatal("Price is nil")
	}

	if price.Cmp(big.NewFloat(0)) <= 0 {
		t.Errorf("Expected positive price, got %v", price)
	}

	t.Logf("ETH Price: %v USDT", price)
}
