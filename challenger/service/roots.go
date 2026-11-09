package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/challenger/metrics"
)

type batchComparer interface {
	Compare(ctx context.Context, batchIndex uint64) (Comparison, error)
}

type rootSource struct {
	rollup   *bindings.Rollup
	l1       *ethclient.Client
	l2       *ethclient.Client
	passer   *bindings.L2ToL1MessagePasser
	logStart uint64
}

func newRootSource(rollup *bindings.Rollup, l1, l2 *ethclient.Client, logStart uint64) (*rootSource, error) {
	passer, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, l2)
	if err != nil {
		return nil, err
	}
	return &rootSource{
		rollup:   rollup,
		l1:       l1,
		l2:       l2,
		passer:   passer,
		logStart: logStart,
	}, nil
}

// Compare reads the batch version currently stored in committedBatches and
// compares its roots with the trusted geth at that batch's last L2 block.
//
// Every L1 value comes from the commit transaction whose CommitBatch event carries
// the current batch hash. revertBatch leaves batchDataStore and committedStateRoots
// in place, so storage can still describe a reverted version.
func (s *rootSource) Compare(ctx context.Context, batchIndex uint64) (Comparison, error) {
	index := new(big.Int).SetUint64(batchIndex)
	raw, err := s.rollup.CommittedBatches(callOpts(ctx), index)
	if err != nil {
		return Comparison{}, &CallError{Err: fmt.Errorf("committedBatches: %w", err)}
	}
	batchHash := common.Hash(raw)
	if batchHash == (common.Hash{}) {
		return Comparison{}, &NotCommittedError{BatchIndex: batchIndex}
	}
	committed, err := s.committedBatch(ctx, index, batchHash)
	if err != nil {
		return Comparison{}, err
	}

	blockNumber := committed.BlockNumber
	obs := ObservedRoots{
		BlockNumber: blockNumber,
		L1State:     committed.StateRoot,
		L1Withdraw:  committed.WithdrawRoot,
	}
	header, err := s.l2.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		if !errors.Is(err, ethereum.NotFound) {
			return Comparison{}, &CallError{Err: fmt.Errorf("l2 header: %w", err)}
		}
	} else if header != nil {
		obs.L2HeaderFound = true
		obs.L2State = header.Root
	}
	if obs.L2HeaderFound {
		root, callErr := s.passer.MessageRoot(&bind.CallOpts{
			Context:     ctx,
			BlockNumber: new(big.Int).SetUint64(blockNumber),
		})
		if callErr != nil {
			obs.L2WithdrawMissing = true
			obs.L2WithdrawErr = callErr.Error()
		} else {
			obs.L2Withdraw = common.BytesToHash(root[:])
		}
	}

	out := Decide(batchIndex, obs)
	out.BatchHash = batchHash
	metrics.Comparisons.WithLabelValues(string(out.Verdict)).Inc()
	return out, nil
}

// committedBatch decodes the commit transaction for this exact batch hash.
// The hash covers every header field read here, so a recommit with the same
// hash carries the same values.
func (s *rootSource) committedBatch(ctx context.Context, index *big.Int, batchHash common.Hash) (commitInput, error) {
	iter, err := s.rollup.FilterCommitBatch(&bind.FilterOpts{
		Start:   s.logStart,
		Context: ctx,
	}, []*big.Int{index}, [][32]byte{batchHash})
	if err != nil {
		return commitInput{}, &CallError{Err: fmt.Errorf("commit log: %w", err)}
	}
	defer iter.Close()

	var txHash common.Hash
	found := false
	for iter.Next() {
		txHash = iter.Event.Raw.TxHash
		found = true
	}
	if err := iter.Error(); err != nil {
		return commitInput{}, &CallError{Err: fmt.Errorf("commit log: %w", err)}
	}
	if !found {
		return commitInput{}, &CallError{Err: fmt.Errorf("commit log not found for batch hash %s", batchHash.Hex())}
	}
	tx, _, err := s.l1.TransactionByHash(ctx, txHash)
	if err != nil {
		return commitInput{}, &CallError{Err: fmt.Errorf("commit transaction: %w", err)}
	}
	input, err := batchInputFromCalldata(tx.Data())
	if err != nil {
		return commitInput{}, &CallError{Err: err}
	}
	return input, nil
}
