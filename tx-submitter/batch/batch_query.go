package batch

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
)

// getLastFinalizeBatchHeaderFromRollupByIndex gets the batch header with the specified index from the rollup contract's FinalizeBatch event
// The finalizeBatch function only receives one parameter: batchHeader bytes, so it can be parsed directly from the transaction
// Query is limited to 10000 block heights, starting from the latest height and querying backwards until data is found
func (bc *BatchCache) getLastFinalizeBatchHeaderFromRollupByIndex(index uint64) (*BatchHeaderBytes, error) {
	// Get the current latest block height
	latestBlock, err := bc.l1Client.BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block number: %w", err)
	}

	const blockRange = uint64(10000) // Query 10000 blocks each time
	var endBlock uint64 = latestBlock
	var startBlock uint64

	// Start from the latest height, query backwards 10000 blocks each time until data is found
	for endBlock > 0 {
		// Calculate the start block for this query
		if endBlock >= blockRange {
			startBlock = endBlock - blockRange + 1
		} else {
			startBlock = 0
		}

		// Set query options
		filterOpts := &bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}

		// Query the FinalizeBatch event with the corresponding index from the rollup contract
		finalizeEventIter, err := bc.rollupContract.FilterFinalizeBatch(filterOpts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			// If query fails, continue querying backwards
			if endBlock < blockRange {
				break // Already queried to block 0, exit loop
			}
			endBlock = startBlock - 1
			continue
		}

		// Iterate through query results
		for finalizeEventIter.Next() {
			event := finalizeEventIter.Event
			// Get transaction hash from event
			txHash := event.Raw.TxHash

			// Get transaction details
			tx, _, err := bc.l1Client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				continue // If getting transaction fails, try next event
			}

			// Parse finalizeBatch transaction data to get batchHeader
			batchHeader, err := parseFinalizeBatchTxData(tx.Data())
			if err != nil {
				continue // If parsing fails, try next event
			}

			// Verify if batch index matches
			batchIndex, err := batchHeader.BatchIndex()
			if err != nil {
				continue
			}
			if batchIndex == index {
				finalizeEventIter.Close()
				return &batchHeader, nil
			}
		}
		finalizeEventIter.Close()

		// Continue querying backwards
		if endBlock < blockRange {
			break // Already queried to block 0, exit loop
		}
		endBlock = startBlock - 1
	}

	return nil, fmt.Errorf("failed to find last finalized batch header for batchIndex %d", index)
}

// parseFinalizeBatchTxData parses the finalizeBatch or importGenesisBatch transaction's input data to get BatchHeaderBytes
// Both finalizeBatch(bytes calldata _batchHeader) and importGenesisBatch(bytes calldata _batchHeader) receive one parameter: batchHeader bytes
// Both methods emit FinalizeBatch event, so we need to support parsing both
func parseFinalizeBatchTxData(txData []byte) (BatchHeaderBytes, error) {
	// Get rollup ABI
	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	// Check if the first 4 bytes of transaction data match the method ID
	if len(txData) < 4 {
		return nil, errors.New("transaction data too short")
	}

	methodID := txData[:4]

	// Try to get finalizeBatch method
	finalizeBatchMethod, ok := rollupAbi.Methods["finalizeBatch"]
	if !ok {
		return nil, errors.New("finalizeBatch method not found in ABI")
	}

	var method abi.Method
	var methodName string

	// Check if method ID matches finalizeBatch
	if bytes.Equal(methodID, finalizeBatchMethod.ID) {
		method = finalizeBatchMethod
		methodName = "finalizeBatch"
	} else {
		// Try importGenesisBatch method
		importGenesisBatchMethod, ok := rollupAbi.Methods["importGenesisBatch"]
		if !ok {
			return nil, errors.New("importGenesisBatch method not found in ABI")
		}
		if bytes.Equal(methodID, importGenesisBatchMethod.ID) {
			method = importGenesisBatchMethod
			methodName = "importGenesisBatch"
		} else {
			return nil, fmt.Errorf("transaction is not a finalizeBatch or importGenesisBatch call, methodID: %x", methodID)
		}
	}

	// Parse parameters (only one parameter: batchHeader bytes)
	args, err := method.Inputs.Unpack(txData[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s transaction parameters: %w", methodName, err)
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no arguments found in %s transaction", methodName)
	}

	// The first parameter is batchHeader bytes
	batchHeaderBytes, ok := args[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("failed to cast batchHeader to []byte in %s transaction", methodName)
	}

	return BatchHeaderBytes(batchHeaderBytes), nil
}

// getCommitBatchDataByIndex gets batchDataInput and batchSignatureInput with the specified index from the rollup contract's CommitBatch event
// Reference the implementation of getLastFinalizeBatchHeaderFromRollupByIndex
// Query is limited to 10000 block heights, starting from the latest height and querying backwards until data is found
func (bc *BatchCache) getCommitBatchDataByIndex(index uint64) (*bindings.IRollupBatchDataInput, *bindings.IRollupBatchSignatureInput, error) {
	// Get the current latest block height
	latestBlock, err := bc.l1Client.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get latest block number: %w", err)
	}

	const blockRange = uint64(10000) // Query 10000 blocks each time
	var endBlock uint64 = latestBlock
	var startBlock uint64

	// Start from the latest height, query backwards 10000 blocks each time until data is found
	for endBlock > 0 {
		// Calculate the start block for this query
		if endBlock >= blockRange {
			startBlock = endBlock - blockRange + 1
		} else {
			startBlock = 0
		}

		// Set query options
		filterOpts := &bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}

		// Query the CommitBatch event with the corresponding index from the rollup contract
		commitEventIter, err := bc.rollupContract.FilterCommitBatch(filterOpts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			// If query fails, continue querying backwards
			if endBlock < blockRange {
				break // Already queried to block 0, exit loop
			}
			endBlock = startBlock - 1
			continue
		}

		// Iterate through query results
		for commitEventIter.Next() {
			event := commitEventIter.Event
			// Get transaction hash from event
			txHash := event.Raw.TxHash

			// Get transaction details
			tx, _, err := bc.l1Client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get transaction by hash: %w", err)
			}

			// Parse commitBatch transaction data to get batchDataInput and batchSignatureInput
			batchDataInput, batchSignatureInput, err := parseCommitBatchTxData(tx.Data())
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse commit batch data: %w", err)
			}

			// Verify if batch index matches (by checking batchIndex in parentBatchHeader)
			if len(batchDataInput.ParentBatchHeader) > 0 {
				parentHeader := BatchHeaderBytes(batchDataInput.ParentBatchHeader)
				parentBatchIndex, err := parentHeader.BatchIndex()
				if err == nil && parentBatchIndex+1 == index {
					commitEventIter.Close()
					return batchDataInput, batchSignatureInput, nil
				}
			}
		}
		commitEventIter.Close()

		// Continue querying backwards
		if endBlock < blockRange {
			break // Already queried to block 0, exit loop
		}
		endBlock = startBlock - 1
	}

	return nil, nil, fmt.Errorf("failed to find commit batch data for index %d", index)
}

// parseCommitBatchTxData parses the commitBatch transaction's input data to get BatchDataInput and BatchSignatureInput
func parseCommitBatchTxData(txData []byte) (*bindings.IRollupBatchDataInput, *bindings.IRollupBatchSignatureInput, error) {
	// Get rollup ABI
	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	// Check if method ID is commitBatch
	commitBatchMethod, ok := rollupAbi.Methods["commitBatch"]
	if !ok {
		return nil, nil, errors.New("commitBatch method not found in ABI")
	}

	// Check if the first 4 bytes of transaction data match the method ID
	if len(txData) < 4 {
		return nil, nil, errors.New("transaction data too short")
	}

	methodID := txData[:4]
	if !bytes.Equal(methodID, commitBatchMethod.ID) {
		// Try commitBatchWithProof
		commitBatchWithProofMethod, ok := rollupAbi.Methods["commitBatchWithProof"]
		if !ok {
			return nil, nil, errors.New("commitBatchWithProof method not found in ABI")
		}
		if bytes.Equal(methodID, commitBatchWithProofMethod.ID) {
			// Use commitBatchWithProof method to parse
			return parseCommitBatchWithProofTxData(txData, rollupAbi)
		}
		return nil, nil, errors.New("transaction is not a commit batch or commitBatchWithProof")
	}

	// Parse parameters
	args, err := commitBatchMethod.Inputs.Unpack(txData[4:])
	if err != nil {
		return nil, nil, err
	}

	// The first parameter is BatchDataInput
	// Note: The struct returned by ABI parsing has JSON tags, need to use matching struct definition
	batchDataInputStruct := args[0].(struct {
		Version           uint8     `json:"version"`
		ParentBatchHeader []uint8   `json:"parentBatchHeader"`
		LastBlockNumber   uint64    `json:"lastBlockNumber"`
		NumL1Messages     uint16    `json:"numL1Messages"`
		PrevStateRoot     [32]uint8 `json:"prevStateRoot"`
		PostStateRoot     [32]uint8 `json:"postStateRoot"`
		WithdrawalRoot    [32]uint8 `json:"withdrawalRoot"`
	})

	// Convert []uint8 to []byte
	parentBatchHeader := make([]byte, len(batchDataInputStruct.ParentBatchHeader))
	for i, v := range batchDataInputStruct.ParentBatchHeader {
		parentBatchHeader[i] = byte(v)
	}

	batchDataInput := &bindings.IRollupBatchDataInput{
		Version:           batchDataInputStruct.Version,
		ParentBatchHeader: parentBatchHeader,
		LastBlockNumber:   batchDataInputStruct.LastBlockNumber,
		NumL1Messages:     batchDataInputStruct.NumL1Messages,
		PrevStateRoot:     batchDataInputStruct.PrevStateRoot,
		PostStateRoot:     batchDataInputStruct.PostStateRoot,
		WithdrawalRoot:    batchDataInputStruct.WithdrawalRoot,
	}

	// The second parameter is BatchSignatureInput
	batchSignatureInputStruct := args[1].(struct {
		SignedSequencersBitmap *big.Int `json:"signedSequencersBitmap"`
		SequencerSets          []uint8  `json:"sequencerSets"`
		Signature              []uint8  `json:"signature"`
	})

	// Convert []uint8 to []byte
	sequencerSets := make([]byte, len(batchSignatureInputStruct.SequencerSets))
	for i, v := range batchSignatureInputStruct.SequencerSets {
		sequencerSets[i] = byte(v)
	}
	signature := make([]byte, len(batchSignatureInputStruct.Signature))
	for i, v := range batchSignatureInputStruct.Signature {
		signature[i] = byte(v)
	}

	batchSignatureInput := &bindings.IRollupBatchSignatureInput{
		SignedSequencersBitmap: batchSignatureInputStruct.SignedSequencersBitmap,
		SequencerSets:          sequencerSets,
		Signature:              signature,
	}

	return batchDataInput, batchSignatureInput, nil
}

// parseCommitBatchWithProofTxData parses the commitBatchWithProof transaction's input data
// commitBatchWithProof has 4 parameters: batchDataInput, batchSignatureInput, _batchHeader, _batchProof
func parseCommitBatchWithProofTxData(txData []byte, rollupAbi *abi.ABI) (*bindings.IRollupBatchDataInput, *bindings.IRollupBatchSignatureInput, error) {
	commitBatchWithProofMethod, ok := rollupAbi.Methods["commitBatchWithProof"]
	if !ok {
		return nil, nil, errors.New("commitBatchWithProof method not found in ABI")
	}

	// Parse parameters
	args, err := commitBatchWithProofMethod.Inputs.Unpack(txData[4:])
	if err != nil {
		return nil, nil, err
	}

	// The first parameter is BatchDataInput
	// Note: The struct returned by ABI parsing has JSON tags, need to use matching struct definition
	batchDataInputStruct := args[0].(struct {
		Version           uint8     `json:"version"`
		ParentBatchHeader []uint8   `json:"parentBatchHeader"`
		LastBlockNumber   uint64    `json:"lastBlockNumber"`
		NumL1Messages     uint16    `json:"numL1Messages"`
		PrevStateRoot     [32]uint8 `json:"prevStateRoot"`
		PostStateRoot     [32]uint8 `json:"postStateRoot"`
		WithdrawalRoot    [32]uint8 `json:"withdrawalRoot"`
	})

	// Convert []uint8 to []byte
	parentBatchHeader := make([]byte, len(batchDataInputStruct.ParentBatchHeader))
	for i, v := range batchDataInputStruct.ParentBatchHeader {
		parentBatchHeader[i] = byte(v)
	}

	batchDataInput := &bindings.IRollupBatchDataInput{
		Version:           batchDataInputStruct.Version,
		ParentBatchHeader: parentBatchHeader,
		LastBlockNumber:   batchDataInputStruct.LastBlockNumber,
		NumL1Messages:     batchDataInputStruct.NumL1Messages,
		PrevStateRoot:     batchDataInputStruct.PrevStateRoot,
		PostStateRoot:     batchDataInputStruct.PostStateRoot,
		WithdrawalRoot:    batchDataInputStruct.WithdrawalRoot,
	}

	// The second parameter is BatchSignatureInput
	batchSignatureInputStruct := args[1].(struct {
		SignedSequencersBitmap *big.Int `json:"signedSequencersBitmap"`
		SequencerSets          []uint8  `json:"sequencerSets"`
		Signature              []uint8  `json:"signature"`
	})

	// Convert []uint8 to []byte
	sequencerSets := make([]byte, len(batchSignatureInputStruct.SequencerSets))
	for i, v := range batchSignatureInputStruct.SequencerSets {
		sequencerSets[i] = byte(v)
	}
	signature := make([]byte, len(batchSignatureInputStruct.Signature))
	for i, v := range batchSignatureInputStruct.Signature {
		signature[i] = byte(v)
	}

	batchSignatureInput := &bindings.IRollupBatchSignatureInput{
		SignedSequencersBitmap: batchSignatureInputStruct.SignedSequencersBitmap,
		SequencerSets:          sequencerSets,
		Signature:              signature,
	}

	// The third parameter is _batchHeader (bytes)
	// The fourth parameter is _batchProof (bytes)
	// These parameters don't need to be returned, but can be used for verification

	return batchDataInput, batchSignatureInput, nil
}
