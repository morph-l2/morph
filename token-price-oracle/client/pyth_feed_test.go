package client

import (
	"math/big"
	"testing"
	"time"
)

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

func TestNormalizePythPriceID(t *testing.T) {
	got := normalizePythPriceID("  0xAbC123  ")
	if got != "abc123" {
		t.Fatalf("normalizePythPriceID() = %q, want abc123", got)
	}
}
