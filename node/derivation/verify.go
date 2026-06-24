package derivation

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// ErrBatchVerifyDivergence is wrapped by verification errors that represent
// a true "verifier reached a verdict of inconsistent" — i.e. the local
// chain disagrees with what L1 committed. Currently produced by:
//   - verifyBatchRoots, when local stateRoot or withdrawalRoot ≠ L1 calldata
//   - verify_local's rebuildBlob, for kinds versioned_hash_mismatch and
//     blob_count_mismatch
//
// Call sites in derivation.go gate `metrics.SetBatchStatus(stateException)`
// on errors.Is(err, ErrBatchVerifyDivergence). Transient or runtime errors
// (RPC down, tx parsing failure, encoding bug, ...) intentionally do NOT
// wrap this sentinel: they reflect "verifier could not run", not "verifier
// determined divergence", and must not light up the divergence alert.
var ErrBatchVerifyDivergence = errors.New("batch verify: divergence verdict")

// verifyBatchRoots verifies the local state root and withdrawal root against the
// values recorded in the L1 commit batch tx calldata.
//
// SPEC-005 section 3.4 invariant: this check is independent of blob data -- both
// batchInfo.root (postStateRoot) and batchInfo.withdrawalRoot are extracted
// from L1 calldata at parse time, so this function runs identically under
// layer1 (beacon blob) and local-rebuild verification modes.
//
// Returns nil on match. On mismatch the error wraps ErrBatchVerifyDivergence
// so callers can distinguish a real divergence verdict from a transient
// failure (e.g. MessageRoot RPC error). Transient failures are returned
// without the sentinel.
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
		return fmt.Errorf("root mismatch: stateRoot(l1=%s, local=%s) withdrawalRoot(l1=%s, local=%s): %w",
			batchInfo.root.Hex(), lastHeader.Root.Hex(),
			batchInfo.withdrawalRoot.Hex(), common.BytesToHash(withdrawalRoot[:]).Hex(),
			ErrBatchVerifyDivergence)
	}
	return nil
}
