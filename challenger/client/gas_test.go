package client

import (
	"context"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/core/types"
)

type stubHeaderClient struct {
	tip    *big.Int
	base   *big.Int
	tipErr error
}

func (s stubHeaderClient) SuggestGasTipCap(context.Context) (*big.Int, error) {
	if s.tipErr != nil {
		return nil, s.tipErr
	}
	return new(big.Int).Set(s.tip), nil
}

func (s stubHeaderClient) HeaderByNumber(context.Context, *big.Int) (*types.Header, error) {
	return &types.Header{BaseFee: s.base}, nil
}

func TestCalculateGasCapsAppliesCeilings(t *testing.T) {
	caps, err := calculateGasCaps(context.Background(), stubHeaderClient{
		tip:  big.NewInt(10),
		base: big.NewInt(3),
	}, big.NewInt(4), big.NewInt(9))
	if err != nil {
		t.Fatal(err)
	}
	if caps.TipCap.Cmp(big.NewInt(4)) != 0 {
		t.Fatalf("tip = %s", caps.TipCap)
	}
	// tip 4 + base 3 * 2 = 10, then the fee ceiling clamps it to 9.
	if caps.FeeCap.Cmp(big.NewInt(9)) != 0 {
		t.Fatalf("fee = %s", caps.FeeCap)
	}
	if caps.TipCap.Cmp(caps.FeeCap) > 0 {
		t.Fatal("tip exceeds fee cap")
	}
}

func TestCalculateGasCapsClampsTipToFee(t *testing.T) {
	caps, err := calculateGasCaps(context.Background(), stubHeaderClient{
		tip:  big.NewInt(20),
		base: big.NewInt(1),
	}, nil, big.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	if caps.FeeCap.Cmp(big.NewInt(5)) != 0 || caps.TipCap.Cmp(big.NewInt(5)) != 0 {
		t.Fatalf("tip=%s fee=%s", caps.TipCap, caps.FeeCap)
	}
}
