package service

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
)

type fakeView struct {
	challenger    bool
	paused        bool
	revertReq     *big.Int
	inChallenge   bool
	lastFinalized *big.Int
	committed     common.Hash
	exist         bool
	challengerOf  common.Address
	inside        bool
	deposit       *big.Int
	balance       *big.Int
	err           error
}

func (f fakeView) IsChallenger(context.Context, common.Address) (bool, error) {
	return f.challenger, f.err
}
func (f fakeView) Paused(context.Context) (bool, error) { return f.paused, f.err }
func (f fakeView) RevertReqIndex(context.Context) (*big.Int, error) {
	return f.revertReq, f.err
}
func (f fakeView) InChallenge(context.Context) (bool, error) { return f.inChallenge, f.err }
func (f fakeView) LastFinalizedBatchIndex(context.Context) (*big.Int, error) {
	return f.lastFinalized, f.err
}
func (f fakeView) CommittedBatch(context.Context, *big.Int) (common.Hash, error) {
	return f.committed, f.err
}
func (f fakeView) BatchExist(context.Context, *big.Int) (bool, error) { return f.exist, f.err }
func (f fakeView) ChallengerOf(context.Context, *big.Int) (common.Address, error) {
	return f.challengerOf, f.err
}
func (f fakeView) InsideWindow(context.Context, *big.Int) (bool, error) { return f.inside, f.err }
func (f fakeView) ChallengeDeposit(context.Context) (*big.Int, error)   { return f.deposit, f.err }
func (f fakeView) Balance(context.Context, common.Address) (*big.Int, error) {
	return f.balance, f.err
}

func readyView(hash common.Hash) fakeView {
	return fakeView{
		challenger:    true,
		revertReq:     big.NewInt(0),
		lastFinalized: big.NewInt(3),
		committed:     hash,
		exist:         true,
		inside:        true,
		deposit:       big.NewInt(100),
		balance:       big.NewInt(150),
	}
}

func TestCheckAcceptsOpenBatch(t *testing.T) {
	hash := common.HexToHash("0x11")
	deposit, err := Check(context.Background(), readyView(hash), common.HexToAddress("0x1"), 4, hash)
	if err != nil {
		t.Fatal(err)
	}
	if deposit.Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("deposit = %s", deposit)
	}
}

func TestCheckRejectsContractFailures(t *testing.T) {
	hash := common.HexToHash("0x11")
	from := common.HexToAddress("0x1")
	cases := []struct {
		name   string
		mutate func(fakeView) fakeView
		reason string
	}{
		{"not challenger", func(v fakeView) fakeView { v.challenger = false; return v }, "only challenger allowed"},
		{"pending revert", func(v fakeView) fakeView { v.revertReq = big.NewInt(2); return v }, "pending revert request"},
		{"paused", func(v fakeView) fakeView { v.paused = true; return v }, "Pausable: paused"},
		{"in challenge", func(v fakeView) fakeView { v.inChallenge = true; return v }, "already in challenge"},
		{"finalized", func(v fakeView) fakeView { v.lastFinalized = big.NewInt(4); return v }, "batch already finalized"},
		{"hash", func(v fakeView) fakeView { v.committed = common.HexToHash("0x22"); return v }, "incorrect batch hash"},
		{"missing", func(v fakeView) fakeView { v.exist = false; return v }, "batch not exist"},
		{"already", func(v fakeView) fakeView {
			v.challengerOf = common.HexToAddress("0x3")
			return v
		}, "batch already challenged"},
		{"window", func(v fakeView) fakeView { v.inside = false; return v }, "cannot challenge batch outside the challenge window"},
		{"value", func(v fakeView) fakeView { v.balance = big.NewInt(99); return v }, "insufficient value"},
		{"gas", func(v fakeView) fakeView { v.balance = big.NewInt(100); return v }, "insufficient balance for gas"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Check(context.Background(), tc.mutate(readyView(hash)), from, 4, hash)
			var rejected *RejectError
			if !errors.As(err, &rejected) || rejected.Reason != tc.reason {
				t.Fatalf("err = %v", err)
			}
		})
	}
}
