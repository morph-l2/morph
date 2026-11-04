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
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/bindings/bindings"
	"morph-l2/challenger/client"
)

// ChallengeResult is a challenge transaction that was broadcast.
type ChallengeResult struct {
	TxHash      common.Hash
	BlockNumber uint64
}

// Service reads rollup state, assembles challengeState, and submits it.
type Service struct {
	l1        *client.L1Client
	rollup    *bindings.Rollup
	view      RollupView
	txs       *TxManager
	methodSig string
	from      common.Address
	mu        sync.Mutex
}

// New binds the L1 Rollup and prepares the transaction path.
func New(l1 *client.L1Client, rollupAddr common.Address, receiptTimeout time.Duration) (*Service, error) {
	rollup, err := bindings.NewRollup(rollupAddr, l1.GetClient())
	if err != nil {
		return nil, err
	}
	methodSig, err := challengeMethodSig()
	if err != nil {
		return nil, err
	}
	return &Service{
		l1:        l1,
		rollup:    rollup,
		view:      &bindingView{rollup: rollup, l1: l1},
		txs:       NewTxManager(l1, receiptTimeout),
		methodSig: methodSig,
		from:      l1.WalletAddress(),
	}, nil
}

// Challenge preflights the batch and submits challengeState with the required deposit.
func (s *Service) Challenge(ctx context.Context, batchIndex uint64, batchHash common.Hash) (ChallengeResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

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
