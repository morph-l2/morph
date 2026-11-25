package service

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
)

type fakeChainID struct {
	id  *big.Int
	err error
}

func (f fakeChainID) ChainID(context.Context) (*big.Int, error) { return f.id, f.err }

// The L2 endpoint is accepted only when its chain ID is the one the Rollup settles.
func TestVerifyL2ChainID(t *testing.T) {
	rollup := func(id uint64, err error) func(*bind.CallOpts) (uint64, error) {
		return func(*bind.CallOpts) (uint64, error) { return id, err }
	}
	tooBig := new(big.Int).Lsh(big.NewInt(1), 64)
	tooBig.Add(tooBig, big.NewInt(2818))

	tests := []struct {
		name    string
		rollup  func(*bind.CallOpts) (uint64, error)
		l2      fakeChainID
		wantErr string
	}{
		{"same chain", rollup(2818, nil), fakeChainID{id: big.NewInt(2818)}, ""},
		{"other chain", rollup(2818, nil), fakeChainID{id: big.NewInt(2910)}, "chain ID 2910"},
		{"chain ID above uint64", rollup(2818, nil), fakeChainID{id: tooBig}, "but the Rollup settles chain 2818"},
		{"rollup read fails", rollup(0, errors.New("l1 down")), fakeChainID{id: big.NewInt(2818)}, "l1 down"},
		{"l2 read fails", rollup(2818, nil), fakeChainID{err: errors.New("l2 down")}, "l2 down"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyL2ChainID(context.Background(), tt.rollup, tt.l2)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %v, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}
