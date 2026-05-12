package derivation

import (
	"bytes"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// verifyBatchRoots verifies the local state root and withdrawal root against the
// values recorded in the L1 commit batch tx calldata.
//
// SPEC-005 §3.4 invariant: this check is independent of blob data — both
// batchInfo.root (postStateRoot) and batchInfo.withdrawalRoot are extracted
// from L1 calldata at parse time, so this function runs identically under
// Path A (online beacon blob) and Path B (local-rebuild) verification modes.
//
// Returns nil on match, error describing the mismatch otherwise.
func (d *Derivation) verifyBatchRoots(batchInfo *BatchInfo, lastHeader *eth.Header) error {
	withdrawalRoot, err := d.L2ToL1MessagePasser.MessageRoot(&bind.CallOpts{
		BlockNumber: lastHeader.Number,
	})
	if err != nil {
		return fmt.Errorf("get withdrawal root failed: %w", err)
	}

	rootMismatch := !bytes.Equal(lastHeader.Root.Bytes(), batchInfo.root.Bytes())
	withdrawalMismatch := !bytes.Equal(withdrawalRoot[:], batchInfo.withdrawalRoot.Bytes())

	if rootMismatch || withdrawalMismatch {
		return fmt.Errorf("root mismatch: stateRoot(l1=%s, local=%s) withdrawalRoot(l1=%s, local=%s)",
			batchInfo.root.Hex(), lastHeader.Root.Hex(),
			batchInfo.withdrawalRoot.Hex(), common.BytesToHash(withdrawalRoot[:]).Hex())
	}
	return nil
}
