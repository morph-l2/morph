package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/bindings/bindings"
	"morph-l2/challenger/client"
)

type bindingView struct {
	rollup *bindings.Rollup
	l1     *client.L1Client
}

func callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func (v *bindingView) IsChallenger(ctx context.Context, addr common.Address) (bool, error) {
	return v.rollup.IsChallenger(callOpts(ctx), addr)
}

func (v *bindingView) Paused(ctx context.Context) (bool, error) {
	return v.rollup.Paused(callOpts(ctx))
}

func (v *bindingView) RevertReqIndex(ctx context.Context) (*big.Int, error) {
	return v.rollup.RevertReqIndex(callOpts(ctx))
}

func (v *bindingView) InChallenge(ctx context.Context) (bool, error) {
	return v.rollup.InChallenge(callOpts(ctx))
}

func (v *bindingView) LastFinalizedBatchIndex(ctx context.Context) (*big.Int, error) {
	return v.rollup.LastFinalizedBatchIndex(callOpts(ctx))
}

func (v *bindingView) CommittedBatch(ctx context.Context, index *big.Int) (common.Hash, error) {
	raw, err := v.rollup.CommittedBatches(callOpts(ctx), index)
	if err != nil {
		return common.Hash{}, err
	}
	return common.Hash(raw), nil
}

func (v *bindingView) BatchExist(ctx context.Context, index *big.Int) (bool, error) {
	return v.rollup.BatchExist(callOpts(ctx), index)
}

func (v *bindingView) ChallengerOf(ctx context.Context, index *big.Int) (common.Address, error) {
	challenge, err := v.rollup.Challenges(callOpts(ctx), index)
	if err != nil {
		return common.Address{}, err
	}
	return challenge.Challenger, nil
}

func (v *bindingView) InsideWindow(ctx context.Context, index *big.Int) (bool, error) {
	return v.rollup.BatchInsideChallengeWindow(callOpts(ctx), index)
}

func (v *bindingView) ChallengeDeposit(ctx context.Context) (*big.Int, error) {
	submitterAddr, err := v.rollup.SubmitterContract(callOpts(ctx))
	if err != nil {
		return nil, err
	}
	if submitterAddr == (common.Address{}) {
		return nil, fmt.Errorf("rollup submitter contract is not set")
	}
	submitter, err := bindings.NewSubmitter(submitterAddr, v.l1.GetClient())
	if err != nil {
		return nil, err
	}
	deposit, err := submitter.ChallengeDeposit(callOpts(ctx))
	if err != nil {
		return nil, err
	}
	if deposit == nil {
		return nil, fmt.Errorf("challenge deposit is nil")
	}
	return deposit, nil
}

func (v *bindingView) Balance(ctx context.Context, addr common.Address) (*big.Int, error) {
	return v.l1.GetBalance(ctx, addr)
}

// VerifyL2ChainID checks that the L2 endpoint serves the chain this Rollup
// settles. Comparing against another chain would report every batch as a
// mismatch and a challenge would then lose its deposit.
func (s *Service) VerifyL2ChainID(ctx context.Context, l2 chainIDReader) error {
	return verifyL2ChainID(ctx, s.rollup.LAYER2CHAINID, l2)
}

type chainIDReader interface {
	ChainID(ctx context.Context) (*big.Int, error)
}

func verifyL2ChainID(ctx context.Context, rollupChainID func(*bind.CallOpts) (uint64, error), l2 chainIDReader) error {
	want, err := rollupChainID(callOpts(ctx))
	if err != nil {
		return fmt.Errorf("read Rollup LAYER_2_CHAIN_ID: %w", err)
	}
	got, err := l2.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("read L2 chain ID: %w", err)
	}
	if !got.IsUint64() || got.Uint64() != want {
		return fmt.Errorf("L2 endpoint has chain ID %s, but the Rollup settles chain %d", got, want)
	}
	return nil
}
