package batch

import (
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
)

const canonicalBootstrapConfirmations uint64 = 6

type confirmedL1Snapshot struct {
	number uint64
	hash   common.Hash
}

// withConfirmedL1Snapshot pins every contract getter and event query performed
// by fn to one confirmed L1 block number, then verifies the block hash again.
// The surrounding initMu serialization guarantees there is at most one active
// bootstrap snapshot per cache.
func (bc *BatchCache) withConfirmedL1Snapshot(fn func() error) error {
	if bc.l1Snapshot != nil {
		return fn()
	}

	latest, err := bc.l1Client.BlockNumber(bc.ctx)
	if err != nil {
		return fmt.Errorf("read latest L1 block for canonical bootstrap: %w", err)
	}
	snapshotNumber := uint64(0)
	if latest > canonicalBootstrapConfirmations {
		snapshotNumber = latest - canonicalBootstrapConfirmations
	}
	header, err := bc.l1Client.HeaderByNumber(bc.ctx, new(big.Int).SetUint64(snapshotNumber))
	if err != nil {
		return fmt.Errorf("read confirmed L1 snapshot %d: %w", snapshotNumber, err)
	}
	if header == nil {
		return fmt.Errorf("read confirmed L1 snapshot %d: nil header", snapshotNumber)
	}
	if header.Number == nil || header.Number.Uint64() != snapshotNumber {
		return fmt.Errorf("confirmed L1 snapshot number mismatch: requested=%d returned=%v", snapshotNumber, header.Number)
	}

	bc.l1Snapshot = &confirmedL1Snapshot{number: snapshotNumber, hash: header.Hash()}
	opErr := fn()
	verifyErr := bc.verifyConfirmedL1Snapshot()
	bc.l1Snapshot = nil
	if opErr != nil {
		if verifyErr != nil {
			return fmt.Errorf("canonical bootstrap failed (%v) and snapshot verification failed: %w", opErr, verifyErr)
		}
		return opErr
	}
	return verifyErr
}

func (bc *BatchCache) verifyConfirmedL1Snapshot() error {
	if bc.l1Snapshot == nil {
		return fmt.Errorf("canonical bootstrap has no confirmed L1 snapshot")
	}
	header, err := bc.l1Client.HeaderByNumber(bc.ctx, new(big.Int).SetUint64(bc.l1Snapshot.number))
	if err != nil {
		return fmt.Errorf("re-read confirmed L1 snapshot %d: %w", bc.l1Snapshot.number, err)
	}
	if header == nil {
		return fmt.Errorf("re-read confirmed L1 snapshot %d: nil header", bc.l1Snapshot.number)
	}
	if header.Number == nil || header.Number.Uint64() != bc.l1Snapshot.number {
		return fmt.Errorf("re-read confirmed L1 snapshot number mismatch: requested=%d returned=%v", bc.l1Snapshot.number, header.Number)
	}
	if header.Hash() != bc.l1Snapshot.hash {
		return fmt.Errorf(
			"confirmed L1 snapshot changed during bootstrap: number=%d expected=%s actual=%s",
			bc.l1Snapshot.number,
			bc.l1Snapshot.hash,
			header.Hash(),
		)
	}
	return nil
}

func (bc *BatchCache) snapshotCallOpts() *bind.CallOpts {
	opts := &bind.CallOpts{Context: bc.ctx}
	if bc.l1Snapshot != nil {
		opts.BlockNumber = new(big.Int).SetUint64(bc.l1Snapshot.number)
	}
	return opts
}

func (bc *BatchCache) snapshotFilterOpts(start, end uint64) (*bind.FilterOpts, error) {
	if bc.l1Snapshot == nil {
		return nil, fmt.Errorf("canonical event query requires a confirmed L1 snapshot")
	}
	if end > bc.l1Snapshot.number {
		return nil, fmt.Errorf("canonical event query end %d exceeds snapshot %d", end, bc.l1Snapshot.number)
	}
	endCopy := end
	return &bind.FilterOpts{Start: start, End: &endCopy, Context: bc.ctx}, nil
}
