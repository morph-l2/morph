package node

import (
	"errors"
	"fmt"

	"morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// SealBatch seals the accumulated blocks into a batch.
// It ensures proper compression and data preparation based on the batch version.
// It should be called after CalculateBatchSizeWithProposalBlock which ensure the accumulated blocks is correct.
func (e *Executor) SealBatch() ([]byte, []byte, error) {
	// Ensure the batching cache is not empty
	if e.batchingCache.IsEmpty() {
		return nil, nil, errors.New("failed to seal batch: batch cache is empty")
	}

	// Parse the current block from the cache
	block, err := types.WrappedBlockFromBytes(e.batchingCache.currentBlockBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse current block: %w", err)
	}

	// Compress and get data hash based on batch version
	compressedPayload, batchDataHash, err := e.handleBatchSealing(block.Timestamp)
	if err != nil {
		return nil, nil, err
	}

	// Generate sidecar for blob data
	sidecar, err := types.MakeBlobTxSidecar(compressedPayload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create blob sidecar: %w", err)
	}

	// Retrieve sequencer verification hash
	sequencerSetVerifyHash, err := e.sequencerCaller.SequencerSetVerifyHash(nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sequencer set verification hash: %w", err)
	}

	// Create batch header
	batchHeader := e.createBatchHeader(batchDataHash, sidecar, sequencerSetVerifyHash)

	// Cache the sealed header and sidecar
	e.batchingCache.sealedBatchHeader = &batchHeader
	e.batchingCache.sealedSidecar = sidecar

	// Log information about the sealed batch
	e.logSealedBatch(batchHeader)

	batchHash := batchHeader.Hash()
	// Return the batch hash and encoded batch header
	return batchHash[:], batchHeader.Encode(), nil
}

// handleBatchSealing determines which version to use for compression and calculates the data hash.
func (e *Executor) handleBatchSealing(blockTimestamp uint64) ([]byte, common.Hash, error) {
	var (
		compressedPayload []byte
		batchDataHash     common.Hash
		err               error
	)

	// Check if the batch should use the upgraded version
	if e.isBatchUpgraded(blockTimestamp) {
		compressedPayload, err = types.CompressBatchBytes(e.batchingCache.batchData.TxsPayloadV2())
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to compress upgraded payload: %w", err)
		}

		if len(compressedPayload) <= types.MaxBlobBytesSize {
			batchDataHash, err = e.batchingCache.batchData.DataHashV2()
			if err != nil {
				return nil, common.Hash{}, fmt.Errorf("failed to calculate upgraded data hash: %w", err)
			}
			return compressedPayload, batchDataHash, nil
		}
	}

	// Fallback to old version if upgraded is not used
	compressedPayload, err = types.CompressBatchBytes(e.batchingCache.batchData.TxsPayload())
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("failed to compress payload: %w", err)
	}
	batchDataHash = e.batchingCache.batchData.DataHash()

	return compressedPayload, batchDataHash, nil
}

// createBatchHeader creates a BatchHeader from the given parameters.
func (e *Executor) createBatchHeader(dataHash common.Hash, sidecar *eth.BlobTxSidecar, sequencerSetVerifyHash common.Hash) types.BatchHeader {
	blobHashes := []common.Hash{types.EmptyVersionedHash}
	if sidecar != nil && len(sidecar.Blobs) > 0 {
		blobHashes = sidecar.BlobHashes()
	}

	l1MessagePopped := e.batchingCache.totalL1MessagePopped - e.batchingCache.parentBatchHeader.TotalL1MessagePopped

	return types.BatchHeader{
		Version:                0,
		BatchIndex:             e.batchingCache.parentBatchHeader.BatchIndex + 1,
		L1MessagePopped:        l1MessagePopped,
		TotalL1MessagePopped:   e.batchingCache.totalL1MessagePopped,
		DataHash:               dataHash,
		BlobVersionedHash:      blobHashes[0],
		PrevStateRoot:          e.batchingCache.prevStateRoot,
		PostStateRoot:          e.batchingCache.postStateRoot,
		WithdrawalRoot:         e.batchingCache.withdrawRoot,
		SequencerSetVerifyHash: sequencerSetVerifyHash,
		ParentBatchHash:        e.batchingCache.parentBatchHeader.Hash(),
	}
}

// logSealedBatch logs the details of the sealed batch for debugging purposes.
func (e *Executor) logSealedBatch(batchHeader types.BatchHeader) {
	e.logger.Info("Sealed batch header", "batchHash", batchHeader.Hash().Hex())
	e.logger.Info(fmt.Sprintf("===batchIndex: %d \n===L1MessagePopped: %d \n===TotalL1MessagePopped: %d \n===dataHash: %x \n===blockNum: %d \n===ParentBatchHash: %x \n",
		batchHeader.BatchIndex,
		batchHeader.L1MessagePopped,
		batchHeader.TotalL1MessagePopped,
		batchHeader.DataHash,
		e.batchingCache.batchData.BlockNum(),
		batchHeader.ParentBatchHash))

	blockContexts, _ := e.batchingCache.batchData.Encode()
	e.logger.Info(fmt.Sprintf("===blockContexts: %x \n", blockContexts))
}
