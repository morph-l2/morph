package updater

import (
	"math/big"
	"testing"

	"morph-l2/token-price-oracle/client"
)

func TestCalculatePriceRatioDoesNotWrapDecimalsAbove18(t *testing.T) {
	updater := &PriceUpdater{}
	price := &client.TokenPrice{
		TokenID:       1,
		Symbol:        "TOKEN24",
		TokenPriceUSD: big.NewFloat(1),
		EthPriceUSD:   big.NewFloat(0.5),
	}
	info := &TokenInfo{
		Decimals: 24,
		Scale:    big.NewInt(1_000_000),
		IsActive: true,
	}

	got, err := updater.calculatePriceRatioWithInfo(1, price, info)
	if err != nil {
		t.Fatal(err)
	}

	// 18-tokenDecimals as uint8 wraps to 250, which would scale by 1e250.
	if got.Cmp(new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil)) >= 0 {
		t.Fatalf("price ratio %s looks like a wrapped 1e250 exponent", got)
	}
}
