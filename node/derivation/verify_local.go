package derivation

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"

	commonbatch "morph-l2/common/batch"
	commonblob "morph-l2/common/blob"
)

// SPEC-005 section 4 local verify: blob-independent batch content verification.
//
// In VerifyModeLocal the node does not pull blobs from the beacon chain on
// the happy path. Instead it reads the L2 blocks in the batch range from
// local storage, reapplies the sequencer's encoding to rebuild the blob
// bytes, and compares the resulting versioned hashes against the values
// declared by the L1 commitBatch tx (carried in BatchInfo.blobHashes).
//
// State / withdrawal root verification (verify.go::verifyBatchRoots) is
// independent of this path and runs after success.
//
// On versioned_hash_mismatch the spec (SPEC-005 §4.3) calls for a
// single-batch self-heal: pull the real blob from beacon, decode + derive
// the batch via the layer1 engine API path (which would replace the
// locally divergent blocks via EL forkchoice), then re-run the shared
// verifyBatchRoots. That self-heal is **currently TODO** and not wired
// up here -- it is blocked on the EL number-continuity check (`params.Number
// == latestNumber + 1` in morph-reth `crates/engine-api/src/builder.rs`
// and go-ethereum `eth/catalyst/l2_api.go`) being relaxed in a separate
// spec. Until then a versioned_hash_mismatch falls through to the legacy
// failure path (log + return + retry next poll).
//
// Mode is selected at startup via --derivation.verify-mode and is not
// switchable at runtime.

// fetchBatchInfoOutline pulls the L1 commitBatch tx, decodes its calldata,
// and populates a BatchInfo using only the calldata + tx blob hashes -- no
// beacon blob fetch. Returned BatchInfo is sufficient for
// verifyBatchContentLocal and verifyBatchRoots.
//
// Only the new commitBatch ABI (rollupABI commitBatch / commitBatchWithProof)
// is supported. lastBlockNumber comes from batch.LastBlockNumber and
// firstBlockNumber from parent header's LastBlockNumber + 1. Legacy-ABI
// batches (calldata BlockContexts + V1 blob encoding) are not handled here
// -- they only exist on historical batches that have long since been
// finalized.
func (d *Derivation) fetchBatchInfoOutline(ctx context.Context, txHash common.Hash, blockNumber uint64) (*BatchInfo, error) {
	tx, pending, err := d.l1Client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}
	if pending {
		return nil, errors.New("pending transaction")
	}
	batch, err := d.UnPackData(tx.Data())
	if err != nil {
		return nil, err
	}

	parentHeader := commonbatch.BatchHeaderBytes(batch.ParentBatchHeader)
	parentBatchIndex, err := parentHeader.BatchIndex()
	if err != nil {
		return nil, fmt.Errorf("decode batch header index error:%v", err)
	}
	parentTotalL1Popped, err := parentHeader.TotalL1MessagePopped()
	if err != nil {
		return nil, fmt.Errorf("decode batch header totalL1MessagePopped error:%v", err)
	}

	bi := &BatchInfo{
		batchIndex:                 parentBatchIndex + 1,
		version:                    uint64(batch.Version),
		root:                       batch.PostStateRoot,
		withdrawalRoot:             batch.WithdrawRoot,
		parentTotalL1MessagePopped: parentTotalL1Popped,
		lastBlockNumber:            batch.LastBlockNumber,
		l1BlockNumber:              blockNumber,
		txHash:                     txHash,
		nonce:                      tx.Nonce(),
		blobHashes:                 tx.BlobHashes(),
	}

	parentLast, err := parentHeader.LastBlockNumber()
	if err != nil {
		return nil, fmt.Errorf("decode parent batch header lastBlockNumber error:%v", err)
	}
	bi.firstBlockNumber = parentLast + 1

	return bi, nil
}

// verifyBatchContentLocal rebuilds blob versioned hashes from local L2
// blocks in the [batchInfo.firstBlockNumber, batchInfo.lastBlockNumber]
// range and compares them against batchInfo.blobHashes (taken from the L1
// commitBatch tx). Returns nil on match.
//
// Failure paths intentionally inline metric inc + structured log + error
// construction at each kind site rather than route through a shared
// helper. One error-wrapping invariant the call site (derivation.go)
// relies on:
//
//   - kind=versioned_hash_mismatch and kind=blob_count_mismatch wrap
//     ErrBatchVerifyDivergence so the call site flips BatchStatus to
//     stateException ONLY on a real "verifier reached unequal verdict";
//     transient / runtime errors must NOT light up the divergence alert.
//     versioned_hash_mismatch will additionally be the self-heal trigger
//     once the EL change lands (see file-level comment).
//
// All other kinds are plain errors. When you add a new kind, decide
// deliberately whether it represents "verifier could not run" (no
// sentinel) vs "verifier produced a divergence verdict" (wrap
// ErrBatchVerifyDivergence) and update the SentinelContract test.
func (d *Derivation) rebuildBlob(ctx context.Context, batchInfo *BatchInfo) ([]common.Hash, error) {
	d.metrics.IncLocalVerifyTriggered()

	// Standard log fields used by every failure-path Error log. Per-site
	// kvs are appended at the call site.
	logBase := []interface{}{
		"batchIndex", batchInfo.batchIndex,
		"version", batchInfo.version,
		"firstBlock", batchInfo.firstBlockNumber,
		"lastBlock", batchInfo.lastBlockNumber,
		"parentTotalL1Popped", batchInfo.parentTotalL1MessagePopped,
		"expectedBlobs", len(batchInfo.blobHashes),
	}

	if batchInfo.firstBlockNumber == 0 || batchInfo.lastBlockNumber < batchInfo.firstBlockNumber {
		d.logger.Error("local verify verification failed: invalid block range",
			append([]interface{}{"kind", "invalid_block_range"}, logBase...)...)
		return nil, fmt.Errorf("local verify [invalid_block_range]: invalid block range [%d, %d]",
			batchInfo.firstBlockNumber, batchInfo.lastBlockNumber)
	}
	if len(batchInfo.blobHashes) == 0 {
		d.logger.Error("local verify verification failed: no blob hashes recorded",
			append([]interface{}{"kind", "empty_blob_hashes"}, logBase...)...)
		return nil, fmt.Errorf("local verify [empty_blob_hashes]: no blob hashes recorded for batch %d", batchInfo.batchIndex)
	}

	bd := commonbatch.NewBatchData()
	totalL1MessagePopped := batchInfo.parentTotalL1MessagePopped

	for n := batchInfo.firstBlockNumber; n <= batchInfo.lastBlockNumber; n++ {
		block, err := d.l2Client.BlockByNumber(ctx, big.NewInt(int64(n)))
		if err != nil {
			d.logger.Error("local verify verification failed: read local block",
				append([]interface{}{"kind", "local_block_read_error", "blockNumber", n, "cause", err}, logBase...)...)
			return nil, fmt.Errorf("local verify [local_block_read_error]: read local block %d failed: %w", n, err)
		}
		if block == nil {
			d.logger.Error("local verify verification failed: local block missing",
				append([]interface{}{"kind", "local_block_missing", "blockNumber", n}, logBase...)...)
			return nil, fmt.Errorf("local verify [local_block_missing]: local block %d missing", n)
		}

		txsPayload, l1TxHashes, newTotal, l2TxNum, err := commonbatch.ParsingTxs(block.Transactions(), totalL1MessagePopped)
		if err != nil {
			d.logger.Error("local verify verification failed: parse local block txs",
				append([]interface{}{"kind", "parsing_txs_error", "blockNumber", n, "cause", err}, logBase...)...)
			return nil, fmt.Errorf("local verify [parsing_txs_error]: parsingTxs failed at block %d: %w", n, err)
		}
		l1MsgNum := int(newTotal - totalL1MessagePopped)
		blockCtx := commonbatch.BuildBlockContext(block.Header(), l2TxNum+l1MsgNum, l1MsgNum)
		bd.Append(blockCtx, txsPayload, l1TxHashes)
		totalL1MessagePopped = newTotal
	}

	// New-ABI only: blob payload is V2-encoded (blockContexts || txs at the
	// blob head). Legacy-ABI batches are out of scope for local verify.
	payload := bd.TxsPayloadV2()
	const chosenEncoding = "V2"

	compressed, err := commonblob.CompressBatchBytes(payload)
	if err != nil {
		d.logger.Error("local verify verification failed: compress",
			append([]interface{}{
				"kind", "compress_error",
				"encoding", chosenEncoding, "payloadLen", len(payload), "cause", err,
			}, logBase...)...)
		return nil, fmt.Errorf("local verify [compress_error]: compress failed: %w", err)
	}

	// maxBlobs is only an upper bound for sidecar capacity; the actual
	// blob count is determined by the size of `compressed`. We pass
	// len(blobHashes) so a payload that would require more blobs than L1
	// declared is rejected up front rather than producing a sidecar with
	// the wrong blob count and a confusing hash mismatch later.
	sidecar, err := commonblob.MakeBlobTxSidecar(compressed, len(batchInfo.blobHashes))
	if err != nil {
		d.logger.Error("local verify verification failed: build sidecar",
			append([]interface{}{
				"kind", "sidecar_build_error",
				"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed), "cause", err,
			}, logBase...)...)
		return nil, fmt.Errorf("local verify [sidecar_build_error]: build sidecar failed: %w", err)
	}

	rebuilt := sidecar.BlobHashes()
	if len(rebuilt) != len(batchInfo.blobHashes) {
		d.logger.Error("local verify verification failed: blob count mismatch",
			append([]interface{}{
				"kind", "blob_count_mismatch",
				"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed),
				"rebuiltBlobs", len(rebuilt),
				"rebuiltHashes", hashesHexCSV(rebuilt),
				"expectedHashes", hashesHexCSV(batchInfo.blobHashes),
			}, logBase...)...)
		return nil, fmt.Errorf("local verify [blob_count_mismatch]: blob count mismatch (rebuilt=%d, l1=%d): %w",
			len(rebuilt), len(batchInfo.blobHashes), ErrBatchVerifyDivergence)
	}
	return rebuilt, nil
}

// hashesHexCSV renders a small slice of hashes as a comma-separated hex
// list, suitable for a one-line log field. Used in divergence diagnostics
// where the per-index hex helps an operator spot which blob diverged.
func hashesHexCSV(hs []common.Hash) string {
	parts := make([]string, len(hs))
	for i, h := range hs {
		parts[i] = h.Hex()
	}
	return strings.Join(parts, ",")
}

// fetchLocalLastHeader returns the local L2 header at
// batchInfo.lastBlockNumber. Used by local verify after content verification
// succeeds, to feed verifyBatchRoots.
func (d *Derivation) fetchLocalLastHeader(ctx context.Context, batchInfo *BatchInfo) (*eth.Header, error) {
	header, err := d.l2Client.HeaderByNumber(ctx, big.NewInt(int64(batchInfo.lastBlockNumber)))
	if err != nil {
		return nil, fmt.Errorf("local verify: read local header at %d failed: %w", batchInfo.lastBlockNumber, err)
	}
	if header == nil {
		return nil, fmt.Errorf("local verify: local header at %d missing", batchInfo.lastBlockNumber)
	}
	return header, nil
}
