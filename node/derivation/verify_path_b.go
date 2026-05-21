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

// SPEC-005 section 4 Path B: blob-independent batch content verification.
//
// In VerifyModePathB the node does not pull blobs from the beacon chain on
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
// the batch via the existing Path A engine API path (which would replace
// the locally divergent blocks via EL forkchoice), then re-run the shared
// verifyBatchRoots. That self-heal is **currently TODO** and not wired
// up here -- it is blocked on the EL number-continuity check (`params.Number
// == latestNumber + 1` in morph-reth `crates/engine-api/src/builder.rs`
// and go-ethereum `eth/catalyst/l2_api.go`) being relaxed in a separate
// spec. Until then a versioned_hash_mismatch falls through to the legacy
// failure path (log + return + retry next poll) under
// path_b_failed_by_kind_total{kind="versioned_hash_mismatch"} and the
// path_b_self_heal_* counters stay at 0.
//
// Mode is selected at startup via --derivation.verify-mode and is not
// switchable at runtime.

// fetchBatchInfoPathB pulls the L1 commitBatch tx, decodes its calldata,
// and populates a BatchInfo using only the calldata + tx blob hashes -- no
// beacon blob fetch. Returned BatchInfo is sufficient for
// verifyBatchContentPathB and verifyBatchRoots.
func (d *Derivation) fetchBatchInfoPathB(ctx context.Context, txHash common.Hash, blockNumber uint64) (*BatchInfo, error) {
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

	bi := new(BatchInfo)
	if err := bi.ParseBatchMetadataOnly(batch); err != nil {
		return nil, fmt.Errorf("parse batch metadata error: %w", err)
	}
	bi.l1BlockNumber = blockNumber
	bi.txHash = txHash
	bi.nonce = tx.Nonce()
	bi.blobHashes = tx.BlobHashes()
	return bi, nil
}

// verifyBatchContentPathB rebuilds blob versioned hashes from local L2
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
func (d *Derivation) verifyBatchContentPathB(ctx context.Context, batchInfo *BatchInfo) error {
	d.metrics.IncPathBTriggered()

	// Standard log fields used by every failure-path Error log. Per-site
	// kvs are appended at the call site.
	logBase := []interface{}{
		"batchIndex", batchInfo.batchIndex,
		"version", batchInfo.version,
		"hasCalldataBlockContexts", batchInfo.hasCalldataBlockContexts,
		"firstBlock", batchInfo.firstBlockNumber,
		"lastBlock", batchInfo.lastBlockNumber,
		"parentTotalL1Popped", batchInfo.parentTotalL1MessagePopped,
		"expectedBlobs", len(batchInfo.blobHashes),
	}

	if batchInfo.firstBlockNumber == 0 || batchInfo.lastBlockNumber < batchInfo.firstBlockNumber {
		d.metrics.IncPathBFailedKind("invalid_block_range")
		d.logger.Error("path B verification failed: invalid block range",
			append([]interface{}{"kind", "invalid_block_range"}, logBase...)...)
		return fmt.Errorf("path B [invalid_block_range]: invalid block range [%d, %d]",
			batchInfo.firstBlockNumber, batchInfo.lastBlockNumber)
	}
	if len(batchInfo.blobHashes) == 0 {
		d.metrics.IncPathBFailedKind("empty_blob_hashes")
		d.logger.Error("path B verification failed: no blob hashes recorded",
			append([]interface{}{"kind", "empty_blob_hashes"}, logBase...)...)
		return fmt.Errorf("path B [empty_blob_hashes]: no blob hashes recorded for batch %d", batchInfo.batchIndex)
	}

	bd := commonbatch.NewBatchData()
	totalL1MessagePopped := batchInfo.parentTotalL1MessagePopped

	for n := batchInfo.firstBlockNumber; n <= batchInfo.lastBlockNumber; n++ {
		block, err := d.l2Client.BlockByNumber(ctx, big.NewInt(int64(n)))
		if err != nil {
			d.metrics.IncPathBFailedKind("local_block_read_error")
			d.logger.Error("path B verification failed: read local block",
				append([]interface{}{"kind", "local_block_read_error", "blockNumber", n, "cause", err}, logBase...)...)
			return fmt.Errorf("path B [local_block_read_error]: read local block %d failed: %w", n, err)
		}
		if block == nil {
			d.metrics.IncPathBFailedKind("local_block_missing")
			d.logger.Error("path B verification failed: local block missing",
				append([]interface{}{"kind", "local_block_missing", "blockNumber", n}, logBase...)...)
			return fmt.Errorf("path B [local_block_missing]: local block %d missing", n)
		}

		txsPayload, l1TxHashes, newTotal, l2TxNum, err := commonbatch.ParsingTxs(block.Transactions(), totalL1MessagePopped)
		if err != nil {
			d.metrics.IncPathBFailedKind("parsing_txs_error")
			d.logger.Error("path B verification failed: parse local block txs",
				append([]interface{}{"kind", "parsing_txs_error", "blockNumber", n, "cause", err}, logBase...)...)
			return fmt.Errorf("path B [parsing_txs_error]: parsingTxs failed at block %d: %w", n, err)
		}
		l1MsgNum := int(newTotal - totalL1MessagePopped)
		blockCtx := commonbatch.BuildBlockContext(block.Header(), l2TxNum+l1MsgNum, l1MsgNum)
		bd.Append(blockCtx, txsPayload, l1TxHashes)
		totalL1MessagePopped = newTotal
	}

	// Pick V1 or V2 blob payload format. The discriminator is the L1
	// commitBatch ABI variant — NOT the BatchHeader version byte:
	//
	//   - Legacy ABI (BlockContexts in calldata) -> blob = TxsPayload (V1)
	//   - New ABI (LastBlockNumber + NumL1Messages, no BlockContexts in
	//     calldata) -> blob = TxsPayloadV2 (V2; blockContexts || txs at
	//     blob head)
	//
	// Sequencer's createBatchHeader sets version byte from
	// (isBatchUpgraded, isBatchV2Upgraded) while handleBatchSealing chooses
	// encoding from (isBatchUpgraded, V2-fits-in-cap); during the V1->V2
	// transition window a single batch can have version=1 + V2 encoding.
	// Path A already keys off `batch.BlockContexts != nil`
	// (batch_info.go::ParseBatch); Path B mirrors that here via the
	// `hasCalldataBlockContexts` flag set in ParseBatchMetadataOnly.
	var (
		payload        []byte
		chosenEncoding string
	)
	if batchInfo.hasCalldataBlockContexts {
		payload = bd.TxsPayload()
		chosenEncoding = "V1"
	} else {
		payload = bd.TxsPayloadV2()
		chosenEncoding = "V2"
	}

	compressed, err := commonblob.CompressBatchBytes(payload)
	if err != nil {
		d.metrics.IncPathBFailedKind("compress_error")
		d.logger.Error("path B verification failed: compress",
			append([]interface{}{
				"kind", "compress_error",
				"encoding", chosenEncoding, "payloadLen", len(payload), "cause", err,
			}, logBase...)...)
		return fmt.Errorf("path B [compress_error]: compress failed: %w", err)
	}

	// maxBlobs is only an upper bound for sidecar capacity; the actual
	// blob count is determined by the size of `compressed`. We pass
	// len(blobHashes) so a payload that would require more blobs than L1
	// declared is rejected up front rather than producing a sidecar with
	// the wrong blob count and a confusing hash mismatch later.
	sidecar, err := commonblob.MakeBlobTxSidecar(compressed, len(batchInfo.blobHashes))
	if err != nil {
		d.metrics.IncPathBFailedKind("sidecar_build_error")
		d.logger.Error("path B verification failed: build sidecar",
			append([]interface{}{
				"kind", "sidecar_build_error",
				"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed), "cause", err,
			}, logBase...)...)
		return fmt.Errorf("path B [sidecar_build_error]: build sidecar failed: %w", err)
	}

	rebuilt := sidecar.BlobHashes()
	if len(rebuilt) != len(batchInfo.blobHashes) {
		d.metrics.IncPathBFailedKind("blob_count_mismatch")
		d.logger.Error("path B verification failed: blob count mismatch",
			append([]interface{}{
				"kind", "blob_count_mismatch",
				"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed),
				"rebuiltBlobs", len(rebuilt),
				"rebuiltHashes", hashesHexCSV(rebuilt),
				"expectedHashes", hashesHexCSV(batchInfo.blobHashes),
			}, logBase...)...)
		return fmt.Errorf("path B [blob_count_mismatch]: blob count mismatch (rebuilt=%d, l1=%d): %w",
			len(rebuilt), len(batchInfo.blobHashes), ErrBatchVerifyDivergence)
	}
	for i := range rebuilt {
		if rebuilt[i] != batchInfo.blobHashes[i] {
			d.metrics.IncPathBFailedKind("versioned_hash_mismatch")
			d.logger.Error("path B verification failed: versioned hash mismatch",
				append([]interface{}{
					"kind", "versioned_hash_mismatch",
					"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed),
					"rebuiltBlobs", len(rebuilt),
					"mismatchIndex", i,
					"rebuiltHashes", hashesHexCSV(rebuilt),
					"expectedHashes", hashesHexCSV(batchInfo.blobHashes),
				}, logBase...)...)
			return fmt.Errorf("path B [versioned_hash_mismatch]: versioned hash mismatch at index %d (rebuilt=%s, l1=%s): %w",
				i, rebuilt[i].Hex(), batchInfo.blobHashes[i].Hex(), ErrBatchVerifyDivergence)
		}
	}
	return nil
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
// batchInfo.lastBlockNumber. Used by Path B after content verification
// succeeds, to feed verifyBatchRoots.
func (d *Derivation) fetchLocalLastHeader(ctx context.Context, batchInfo *BatchInfo) (*eth.Header, error) {
	header, err := d.l2Client.HeaderByNumber(ctx, big.NewInt(int64(batchInfo.lastBlockNumber)))
	if err != nil {
		return nil, fmt.Errorf("path B: read local header at %d failed: %w", batchInfo.lastBlockNumber, err)
	}
	if header == nil {
		return nil, fmt.Errorf("path B: local header at %d missing", batchInfo.lastBlockNumber)
	}
	return header, nil
}
