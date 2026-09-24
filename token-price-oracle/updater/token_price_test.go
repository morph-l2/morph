package updater

import (
	"math/big"
	"testing"

	"morph-l2/token-price-oracle/client"
)

func TestCalculatePriceRatioScalesDownDecimalsAbove18(t *testing.T) {
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

	// 1e6 * (1 USD / 0.5 USD per ETH) * 1e(18-24) = 2.
	// A uint8 subtraction would wrap to 250 and scale by 1e250 instead.
	if got.Cmp(big.NewInt(2)) != 0 {
		t.Fatalf("price ratio = %s, want 2", got)
	}
}
