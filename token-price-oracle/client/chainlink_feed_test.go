package client

import (
	"math/big"
	"testing"
	"time"
)

func TestValidateChainlinkRound(t *testing.T) {
	now := time.Unix(1_700_000_000, 0)

	tests := []struct {
		name            string
		answer          *big.Int
		updatedAt       *big.Int
		roundID         *big.Int
		answeredInRound *big.Int
		wantErr         bool
	}{
		{
			name:            "valid",
			answer:          big.NewInt(2000_00000000),
			updatedAt:       big.NewInt(now.Add(-5 * time.Minute).Unix()),
			roundID:         big.NewInt(10),
			answeredInRound: big.NewInt(10),
		},
		{
			name:            "non-positive answer",
			answer:          big.NewInt(0),
			updatedAt:       big.NewInt(now.Add(-5 * time.Minute).Unix()),
			roundID:         big.NewInt(10),
			answeredInRound: big.NewInt(10),
			wantErr:         true,
		},
		{
			name:            "stale",
			answer:          big.NewInt(2000_00000000),
			updatedAt:       big.NewInt(now.Add(-2 * time.Hour).Unix()),
			roundID:         big.NewInt(10),
			answeredInRound: big.NewInt(10),
			wantErr:         true,
		},
		{
			name:            "answered in old round",
			answer:          big.NewInt(2000_00000000),
			updatedAt:       big.NewInt(now.Add(-5 * time.Minute).Unix()),
			roundID:         big.NewInt(10),
			answeredInRound: big.NewInt(9),
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateChainlinkRound(tt.answer, tt.updatedAt, tt.roundID, tt.answeredInRound, time.Hour, now)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateChainlinkRound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChainlinkAnswerToFloat(t *testing.T) {
	price := chainlinkAnswerToFloat(big.NewInt(123456789000), 8)
	got, _ := price.Float64()
	if got != 1234.56789 {
		t.Fatalf("chainlinkAnswerToFloat() = %v, want 1234.56789", got)
	}
}
