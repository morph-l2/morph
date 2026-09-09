package updater

import (
	"math/big"
	"strings"
	"testing"

	"morph-l2/token-price-oracle/client"
)

func TestCalculatePriceRatioRejectsDecimalsAboveETH(t *testing.T) {
	updater := &PriceUpdater{}
	price := &client.TokenPrice{
		TokenID:       1,
		Symbol:        "TOKEN24",
		TokenPriceUSD: big.NewFloat(1),
		EthPriceUSD:   big.NewFloat(2_000),
	}
	info := &TokenInfo{
		Decimals: 24,
		Scale:    big.NewInt(1_000_000),
		IsActive: true,
	}

	_, err := updater.calculatePriceRatioWithInfo(1, price, info)
	if err == nil {
		t.Fatal("expected decimals above 18 to be rejected")
	}
	if !strings.Contains(err.Error(), "unsupported decimals 24") {
		t.Fatalf("unexpected error: %v", err)
	}
}
