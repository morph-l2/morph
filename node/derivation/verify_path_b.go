package derivation

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"

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
	return verifyPathBContent(ctx, d.l2Client, d.metrics, batchInfo)
}

// verifyPathBContent is the testable core of Path B verification. It is
// extracted from the Derivation method above so tests can supply a fake
// pathBBlockReader. Behavior and error messages are unchanged.
func verifyPathBContent(ctx context.Context, reader pathBBlockReader, metrics *Metrics, batchInfo *BatchInfo) error {
	metrics.IncPathBTriggered()

	if batchInfo.firstBlockNumber == 0 || batchInfo.lastBlockNumber < batchInfo.firstBlockNumber {
		metrics.IncPathBFailed()
		return fmt.Errorf("path B: invalid block range [%d, %d]",
			batchInfo.firstBlockNumber, batchInfo.lastBlockNumber)
	}
	if len(batchInfo.blobHashes) == 0 {
		metrics.IncPathBFailed()
		return fmt.Errorf("path B: no blob hashes recorded for batch %d", batchInfo.batchIndex)
	}

	bd := commonbatch.NewBatchData()
	totalL1MessagePopped := batchInfo.parentTotalL1MessagePopped

	for n := batchInfo.firstBlockNumber; n <= batchInfo.lastBlockNumber; n++ {
		block, err := reader.BlockByNumber(ctx, big.NewInt(int64(n)))
		if err != nil {
			metrics.IncPathBFailed()
			return fmt.Errorf("path B: read local block %d failed: %w", n, err)
		}
		if block == nil {
			metrics.IncPathBFailed()
			return fmt.Errorf("path B: local block %d missing", n)
		}

		txsPayload, l1TxHashes, newTotal, l2TxNum, err := commonbatch.ParsingTxs(block.Transactions(), totalL1MessagePopped)
		if err != nil {
			metrics.IncPathBFailed()
			return fmt.Errorf("path B: parsingTxs failed at block %d: %w", n, err)
		}
		l1MsgNum := int(newTotal - totalL1MessagePopped)
		blockCtx := commonbatch.BuildBlockContext(block.Header(), l2TxNum+l1MsgNum, l1MsgNum)
		bd.Append(blockCtx, txsPayload, l1TxHashes)
		totalL1MessagePopped = newTotal
	}

	// Pick V1 or V2 payload format based on batch version. V2 prepends the
	// concatenated block contexts to the tx payload; V1 carries only txs.
	var payload []byte
	if batchInfo.version >= 2 {
		payload = bd.TxsPayloadV2()
	} else {
		payload = bd.TxsPayload()
	}

	compressed, err := commonblob.CompressBatchBytes(payload)
	if err != nil {
		metrics.IncPathBFailed()
		return fmt.Errorf("path B: compress failed: %w", err)
	}

	// maxBlobs is only an upper bound for sidecar capacity; the actual
	// blob count is determined by the size of `compressed`. We pass
	// len(blobHashes) so a payload that would require more blobs than
	// L1 declared is rejected up front rather than producing a sidecar
	// with the wrong blob count and a confusing hash mismatch later.
	sidecar, err := commonblob.MakeBlobTxSidecar(compressed, len(batchInfo.blobHashes))
	if err != nil {
		metrics.IncPathBFailed()
		return fmt.Errorf("path B: build sidecar failed: %w", err)
	}

	rebuilt := sidecar.BlobHashes()
	if len(rebuilt) != len(batchInfo.blobHashes) {
		metrics.IncPathBFailed()
		return fmt.Errorf("path B: blob count mismatch (rebuilt=%d, l1=%d)",
			len(rebuilt), len(batchInfo.blobHashes))
	}
	for i := range rebuilt {
		if rebuilt[i] != batchInfo.blobHashes[i] {
			metrics.IncPathBFailed()
			return fmt.Errorf("path B: versioned hash mismatch at index %d (rebuilt=%s, l1=%s)",
				i, rebuilt[i].Hex(), batchInfo.blobHashes[i].Hex())
		}
	}
	return nil
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
