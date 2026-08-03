package client

import (
	"math/big"
	"testing"
	"time"
)

func TestNewPythHermesPriceFeedRequiresAPIKey(t *testing.T) {
	const ethUSDPriceID = "0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace"
	mapping := map[uint16]string{1: "0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43"}

	for _, apiKey := range []string{"", "   "} {
		if _, err := NewPythHermesPriceFeed(mapping, "https://hermes.pyth.network", apiKey, ethUSDPriceID, time.Hour, 0); err == nil {
			t.Fatalf("NewPythHermesPriceFeed(apiKey=%q) succeeded, want error", apiKey)
		}
	}

	if _, err := NewPythHermesPriceFeed(mapping, "https://hermes.pyth.network", "key", ethUSDPriceID, time.Hour, 0); err != nil {
		t.Fatalf("NewPythHermesPriceFeed with an API key failed: %v", err)
	}
}

func TestValidatePythPrice(t *testing.T) {
	now := time.Unix(1_700_000_000, 0)

	tests := []struct {
		name             string
		price            pythPrice
		maxConfidenceBPS uint64
		wantErr          bool
	}{
		{
			name: "valid",
			price: pythPrice{
				Price:       "175500000000",
				Confidence:  "100000000",
				Exponent:    -8,
				PublishTime: now.Add(-5 * time.Minute).Unix(),
			},
			maxConfidenceBPS: 100,
		},
		{
			name: "stale",
			price: pythPrice{
				Price:       "175500000000",
				Confidence:  "100000000",
				Exponent:    -8,
				PublishTime: now.Add(-2 * time.Hour).Unix(),
			},
			maxConfidenceBPS: 100,
			wantErr:          true,
		},
		{
			name: "too wide confidence",
			price: pythPrice{
				Price:       "100000000",
				Confidence:  "2000000",
				Exponent:    -8,
				PublishTime: now.Add(-5 * time.Minute).Unix(),
			},
			maxConfidenceBPS: 100,
			wantErr:          true,
		},
		{
			// A publish time ahead of the local clock used to be accepted anywhere
			// inside the staleness window, which with the default meant nearly an hour.
			name: "future publish time beyond skew",
			price: pythPrice{
				Price:       "175500000000",
				Confidence:  "100000000",
				Exponent:    -8,
				PublishTime: now.Add(30 * time.Minute).Unix(),
			},
			maxConfidenceBPS: 100,
			wantErr:          true,
		},
		{
			// Host clock drift against the publisher must not reject a fresh price.
			name: "future publish time within skew",
			price: pythPrice{
				Price:       "175500000000",
				Confidence:  "100000000",
				Exponent:    -8,
				PublishTime: now.Add(pythMaxClockSkew / 2).Unix(),
			},
			maxConfidenceBPS: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePythPrice(tt.price, time.Hour, tt.maxConfidenceBPS, now)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validatePythPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPythPriceToFloat(t *testing.T) {
	price, err := pythPriceToFloat(pythPrice{
		Price:    "175500000000",
		Exponent: -8,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := big.NewFloat(1755)
	if price.Cmp(want) != 0 {
		t.Fatalf("pythPriceToFloat() = %s, want %s", price.String(), want.String())
	}
}

func TestPythPriceToFloatRejectsOutOfRangeExponent(t *testing.T) {
	// math.MinInt32 previously reached big.Int.Exp as 10^2147483648.
	for _, exponent := range []int32{-2147483648, 2147483647, pythMaxExponentMagnitude + 1, -(pythMaxExponentMagnitude + 1)} {
		if _, err := pythPriceToFloat(pythPrice{Price: "1", Exponent: exponent}); err == nil {
			t.Fatalf("pythPriceToFloat(expo=%d) succeeded, want error", exponent)
		}
	}

	if _, err := pythPriceToFloat(pythPrice{Price: "1", Exponent: -pythMaxExponentMagnitude}); err != nil {
		t.Fatalf("pythPriceToFloat at the exponent bound failed: %v", err)
	}
}

func TestNormalizePythPriceID(t *testing.T) {
	got := normalizePythPriceID("  0xAbC123  ")
	if got != "abc123" {
		t.Fatalf("normalizePythPriceID() = %q, want abc123", got)
	}
}
