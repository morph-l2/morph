package updater

import (
	"math/big"
	"testing"

	"morph-l2/token-price-oracle/client"
)

func TestCalculatePriceRatioScalesDecimalsAboveETH(t *testing.T) {
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

	// 1e6 * (1 / 0.5) * 10^(18-24) = 2. Must not wrap 18-24 as uint8 250.
	if got.Cmp(big.NewInt(2)) != 0 {
		t.Fatalf("price ratio = %s, want 2", got)
	}
}
