package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
)

// RollupView is the chain state read before a challenge deposit is sent.
type RollupView interface {
	IsChallenger(ctx context.Context, addr common.Address) (bool, error)
	Paused(ctx context.Context) (bool, error)
	RevertReqIndex(ctx context.Context) (*big.Int, error)
	InChallenge(ctx context.Context) (bool, error)
	LastFinalizedBatchIndex(ctx context.Context) (*big.Int, error)
	CommittedBatch(ctx context.Context, index *big.Int) (common.Hash, error)
	BatchExist(ctx context.Context, index *big.Int) (bool, error)
	ChallengerOf(ctx context.Context, index *big.Int) (common.Address, error)
	InsideWindow(ctx context.Context, index *big.Int) (bool, error)
	ChallengeDeposit(ctx context.Context) (*big.Int, error)
	Balance(ctx context.Context, addr common.Address) (*big.Int, error)
}

// Check decides whether challengeState can be submitted for this batch.
// On success it returns the deposit that must be attached to the transaction.
// Checks follow the require order in Rollup.challengeState.
func Check(ctx context.Context, view RollupView, from common.Address, batchIndex uint64, batchHash common.Hash) (*big.Int, error) {
	isChallenger, err := view.IsChallenger(ctx, from)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("isChallenger: %w", err)}
	}
	if !isChallenger {
		return nil, &RejectError{Reason: "only challenger allowed"}
	}

	revertReq, err := view.RevertReqIndex(ctx)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("revertReqIndex: %w", err)}
	}
	if revertReq != nil && revertReq.Sign() > 0 {
		return nil, &RejectError{Reason: "pending revert request"}
	}

	paused, err := view.Paused(ctx)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("paused: %w", err)}
	}
	if paused {
		return nil, &RejectError{Reason: "Pausable: paused"}
	}

	inChallenge, err := view.InChallenge(ctx)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("inChallenge: %w", err)}
	}
	if inChallenge {
		return nil, &RejectError{Reason: "already in challenge"}
	}

	lastFinalized, err := view.LastFinalizedBatchIndex(ctx)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("lastFinalizedBatchIndex: %w", err)}
	}
	if lastFinalized == nil {
		return nil, &CallError{Err: fmt.Errorf("lastFinalizedBatchIndex is nil")}
	}
	index := new(big.Int).SetUint64(batchIndex)
	if lastFinalized.Cmp(index) >= 0 {
		return nil, &RejectError{Reason: "batch already finalized"}
	}

	committed, err := view.CommittedBatch(ctx, index)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("committedBatches: %w", err)}
	}
	if committed != batchHash {
		return nil, &RejectError{Reason: "incorrect batch hash"}
	}

	exists, err := view.BatchExist(ctx, index)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("batchExist: %w", err)}
	}
	if !exists {
		return nil, &RejectError{Reason: "batch not exist"}
	}

	challenger, err := view.ChallengerOf(ctx, index)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("challenges: %w", err)}
	}
	if challenger != (common.Address{}) {
		return nil, &RejectError{Reason: "batch already challenged"}
	}

	inside, err := view.InsideWindow(ctx, index)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("batchInsideChallengeWindow: %w", err)}
	}
	if !inside {
		return nil, &RejectError{Reason: "cannot challenge batch outside the challenge window"}
	}

	deposit, err := view.ChallengeDeposit(ctx)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("challengeDeposit: %w", err)}
	}
	if deposit == nil {
		return nil, &CallError{Err: fmt.Errorf("challenge deposit is nil")}
	}

	balance, err := view.Balance(ctx, from)
	if err != nil {
		return nil, &CallError{Err: fmt.Errorf("balance: %w", err)}
	}
	if balance == nil || balance.Cmp(deposit) < 0 {
		return nil, &RejectError{Reason: "insufficient value"}
	}
	if balance.Cmp(deposit) == 0 {
		return nil, &RejectError{Reason: "insufficient balance for gas"}
	}

	return new(big.Int).Set(deposit), nil
}
