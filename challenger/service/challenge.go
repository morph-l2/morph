package service

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/bindings/bindings"
	"morph-l2/challenger/client"
)

// ChallengeResult is a challenge transaction that was broadcast.
type ChallengeResult struct {
	TxHash      common.Hash
	BlockNumber uint64
}

// Service reads rollup state, compares batch roots, and submits challengeState.
type Service struct {
	l1        *client.L1Client
	rollup    *bindings.Rollup
	view      RollupView
	compare   batchComparer
	txs       *TxManager
	methodSig string
	from      common.Address
	mu        sync.Mutex
}

// New binds the L1 Rollup and the trusted L2 geth, and prepares the transaction path.
func New(l1 *client.L1Client, l2 *ethclient.Client, rollupAddr common.Address, receiptTimeout time.Duration, logStart uint64) (*Service, error) {
	rollup, err := bindings.NewRollup(rollupAddr, l1.GetClient())
	if err != nil {
		return nil, err
	}
	methodSig, err := challengeMethodSig()
	if err != nil {
		return nil, err
	}
	roots, err := newRootSource(rollup, l1.GetClient(), l2, logStart)
	if err != nil {
		return nil, err
	}
	return &Service{
		l1:        l1,
		rollup:    rollup,
		view:      &bindingView{rollup: rollup, l1: l1},
		compare:   roots,
		txs:       NewTxManager(l1, receiptTimeout),
		methodSig: methodSig,
		from:      l1.WalletAddress(),
	}, nil
}

// CompareBatch reads L1 and the trusted L2 geth and compares the two roots.
func (s *Service) CompareBatch(ctx context.Context, batchIndex uint64) (Comparison, error) {
	return s.compare.Compare(ctx, batchIndex)
}

// Challenge compares the batch roots, then preflights and submits challengeState.
// A batch is signed only when at least one root mismatches.
func (s *Service) Challenge(ctx context.Context, batchIndex uint64, batchHash common.Hash) (ChallengeResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	comparison, err := s.CompareBatch(ctx, batchIndex)
	if err != nil {
		var missing *NotCommittedError
		if errors.As(err, &missing) {
			return ChallengeResult{}, &RejectError{Reason: missing.Error()}
		}
		return ChallengeResult{}, err
	}
	// The verdict holds only for the version it read. A different hash is
	// another version of this index, or a stale request.
	if comparison.BatchHash != batchHash {
		return ChallengeResult{}, &RejectError{Reason: "incorrect batch hash"}
	}
	if comparison.Verdict != VerdictMismatch {
		return ChallengeResult{}, &RejectError{Reason: refuseReason(comparison)}
	}

	deposit, err := Check(ctx, s.view, s.from, batchIndex, batchHash)
	if err != nil {
		return ChallengeResult{}, err
	}

	log.Info("Submitting challenge",
		"batchIndex", batchIndex,
		"batchHash", batchHash.Hex(),
		"from", s.from.Hex(),
		"value", deposit)

	receipt, err := s.txs.SendTransaction(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Value = new(big.Int).Set(deposit)
		return s.rollup.ChallengeState(opts, batchIndex, batchHash)
	}, s.methodSig)
	return resultFrom(receipt, err)
}

func resultFrom(receipt *types.Receipt, err error) (ChallengeResult, error) {
	var result ChallengeResult
	if receipt != nil {
		result.TxHash = receipt.TxHash
		if receipt.BlockNumber != nil {
			result.BlockNumber = receipt.BlockNumber.Uint64()
		}
	}
	var txErr *TxError
	if errors.As(err, &txErr) && result.TxHash == (common.Hash{}) {
		result.TxHash = common.HexToHash(txErr.TxHash)
	}
	return result, err
}
