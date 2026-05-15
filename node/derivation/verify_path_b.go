package derivation

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	tmlog "github.com/tendermint/tendermint/libs/log"

	commonbatch "morph-l2/common/batch"
	commonblob "morph-l2/common/blob"
)

// SPEC-005 section 4 Path B: blob-independent batch content verification.
//
// In VerifyModePathB the node does not pull blobs from the beacon chain.
// Instead it reads the L2 blocks in the batch range from local storage,
// reapplies the sequencer's encoding to rebuild the blob bytes, and compares
// the resulting versioned hashes against the values declared by the L1
// commitBatch tx (carried in BatchInfo.blobHashes).
//
// State / withdrawal root verification (verify.go::verifyBatchRoots) is
// independent of this path and runs after success.
//
// Path A and Path B are mutually exclusive: the mode is fixed at startup by
// `--derivation.verify-mode` and cannot change at runtime. Path A failure
// (e.g. blob unavailable) does NOT auto-fall-back to Path B; the operator
// must restart with the alternate mode.

// fetchBatchInfoPathB pulls the L1 commitBatch tx, decodes its calldata, and
// populates a BatchInfo using only the calldata + tx blob hashes -- no beacon
// blob fetch. Returned BatchInfo is sufficient for verifyBatchContentPathB
// and verifyBatchRoots.
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

// pathBBlockReader is the minimal L2 client surface verifyPathBContent
// needs. Narrowed from types.RetryableClient so unit tests can exercise
// the full Path B encoding pipeline without an authclient stack.
type pathBBlockReader interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*eth.Block, error)
}

// verifyBatchContentPathB rebuilds blob versioned hashes from local L2
// blocks in the [batchInfo.firstBlockNumber, batchInfo.lastBlockNumber]
// range and compares them against batchInfo.blobHashes (taken from the
// L1 commitBatch tx). Returns nil on match.
func (d *Derivation) verifyBatchContentPathB(ctx context.Context, batchInfo *BatchInfo) error {
	return verifyPathBContent(ctx, d.l2Client, d.metrics, d.logger, batchInfo)
}

// verifyPathBContent is the testable core of Path B verification. It is
// extracted from the Derivation method above so tests can supply a fake
// pathBBlockReader. Behavior and error messages are unchanged; on every
// failure path it emits a single structured Error log carrying the
// fields an operator needs to diagnose the mismatch (kind, batchIndex,
// version, block range, parent total L1 messages popped, chosen
// encoding when reached, payload / compressed lengths, rebuilt vs
// expected blob hashes) and increments the per-kind PathBFailed metric.
func verifyPathBContent(ctx context.Context, reader pathBBlockReader, metrics *Metrics, logger tmlog.Logger, batchInfo *BatchInfo) error {
	metrics.IncPathBTriggered()

	if batchInfo.firstBlockNumber == 0 || batchInfo.lastBlockNumber < batchInfo.firstBlockNumber {
		return pathBFail(logger, metrics, batchInfo, "invalid_block_range", nil,
			fmt.Sprintf("invalid block range [%d, %d]", batchInfo.firstBlockNumber, batchInfo.lastBlockNumber))
	}
	if len(batchInfo.blobHashes) == 0 {
		return pathBFail(logger, metrics, batchInfo, "empty_blob_hashes", nil,
			fmt.Sprintf("no blob hashes recorded for batch %d", batchInfo.batchIndex))
	}

	bd := commonbatch.NewBatchData()
	totalL1MessagePopped := batchInfo.parentTotalL1MessagePopped

	for n := batchInfo.firstBlockNumber; n <= batchInfo.lastBlockNumber; n++ {
		block, err := reader.BlockByNumber(ctx, big.NewInt(int64(n)))
		if err != nil {
			return pathBFail(logger, metrics, batchInfo, "local_block_read_error", err,
				fmt.Sprintf("read local block %d failed", n), "blockNumber", n)
		}
		if block == nil {
			return pathBFail(logger, metrics, batchInfo, "local_block_missing", nil,
				fmt.Sprintf("local block %d missing", n), "blockNumber", n)
		}

		txsPayload, l1TxHashes, newTotal, l2TxNum, err := commonbatch.ParsingTxs(block.Transactions(), totalL1MessagePopped)
		if err != nil {
			return pathBFail(logger, metrics, batchInfo, "parsing_txs_error", err,
				fmt.Sprintf("parsingTxs failed at block %d", n), "blockNumber", n)
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
	// (isBatchUpgraded, isBatchV2Upgraded) while handleBatchSealing
	// chooses encoding from (isBatchUpgraded, V2-fits-in-cap); during
	// the V1->V2 transition window a single batch can have version=1 +
	// V2 encoding. Path A already keys off `batch.BlockContexts != nil`
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
		return pathBFail(logger, metrics, batchInfo, "compress_error", err,
			"compress failed",
			"encoding", chosenEncoding, "payloadLen", len(payload))
	}

	// maxBlobs is only an upper bound for sidecar capacity; the actual
	// blob count is determined by the size of `compressed`. We pass
	// len(blobHashes) so a payload that would require more blobs than
	// L1 declared is rejected up front rather than producing a sidecar
	// with the wrong blob count and a confusing hash mismatch later.
	sidecar, err := commonblob.MakeBlobTxSidecar(compressed, len(batchInfo.blobHashes))
	if err != nil {
		return pathBFail(logger, metrics, batchInfo, "sidecar_build_error", err,
			"build sidecar failed",
			"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed))
	}

	rebuilt := sidecar.BlobHashes()
	if len(rebuilt) != len(batchInfo.blobHashes) {
		return pathBFail(logger, metrics, batchInfo, "blob_count_mismatch", nil,
			fmt.Sprintf("blob count mismatch (rebuilt=%d, l1=%d)", len(rebuilt), len(batchInfo.blobHashes)),
			"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed),
			"rebuiltBlobs", len(rebuilt),
			"rebuiltHashes", hashesHexCSV(rebuilt),
			"expectedHashes", hashesHexCSV(batchInfo.blobHashes))
	}
	for i := range rebuilt {
		if rebuilt[i] != batchInfo.blobHashes[i] {
			return pathBFail(logger, metrics, batchInfo, "versioned_hash_mismatch", nil,
				fmt.Sprintf("versioned hash mismatch at index %d (rebuilt=%s, l1=%s)",
					i, rebuilt[i].Hex(), batchInfo.blobHashes[i].Hex()),
				"encoding", chosenEncoding, "payloadLen", len(payload), "compressedLen", len(compressed),
				"rebuiltBlobs", len(rebuilt),
				"mismatchIndex", i,
				"rebuiltHashes", hashesHexCSV(rebuilt),
				"expectedHashes", hashesHexCSV(batchInfo.blobHashes))
		}
	}
	return nil
}

// pathBFail is the single failure exit for verifyPathBContent. It emits one
// structured Error log carrying the always-present diagnostic fields plus any
// per-site context kvs the caller supplies, increments the per-kind
// PathBFailed metric, and returns a wrapped error so the upstream
// derivationBlock log retains the same surface as before. cause may be nil
// for sanity-check failures (no underlying error to wrap).
func pathBFail(logger tmlog.Logger, metrics *Metrics, batchInfo *BatchInfo, kind string, cause error, msg string, kvs ...interface{}) error {
	metrics.IncPathBFailedKind(kind)

	args := []interface{}{
		"kind", kind,
		"batchIndex", batchInfo.batchIndex,
		"version", batchInfo.version,
		"hasCalldataBlockContexts", batchInfo.hasCalldataBlockContexts,
		"firstBlock", batchInfo.firstBlockNumber,
		"lastBlock", batchInfo.lastBlockNumber,
		"parentTotalL1Popped", batchInfo.parentTotalL1MessagePopped,
		"expectedBlobs", len(batchInfo.blobHashes),
	}
	args = append(args, kvs...)
	if cause != nil {
		args = append(args, "cause", cause)
	}
	logger.Error("path B verification failed: "+msg, args...)

	if cause != nil {
		return fmt.Errorf("path B [%s]: %s: %w", kind, msg, cause)
	}
	return fmt.Errorf("path B [%s]: %s", kind, msg)
}

// hashesHexCSV renders a small slice of hashes as a comma-separated hex list,
// suitable for a one-line log field. Used in failure diagnostics where the
// per-index hex helps an operator spot which blob diverged.
func hashesHexCSV(hs []common.Hash) string {
	parts := make([]string, len(hs))
	for i, h := range hs {
		parts[i] = h.Hex()
	}
	return strings.Join(parts, ",")
}

// fetchLocalLastHeader returns the local L2 header at batchInfo.lastBlockNumber.
// Used by Path B after content verification succeeds, to feed verifyBatchRoots.
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
