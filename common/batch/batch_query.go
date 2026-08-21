package batch

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

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
			if startBlock == 0 {
				break // Already queried to block 0, exit loop
			}
			endBlock = startBlock - 1
			continue
		}
		if finalizeEventIter == nil {
			return nil, fmt.Errorf("filter finalized batch %d returned a nil iterator", index)
		}
		defer func() { _ = finalizeEventIter.Close() }()
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
				return &batchHeader, nil
			}
		}

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

// batchDataInputStruct represents the parsed batch data input structure from ABI
type batchDataInputStruct struct {
	Version           uint8     `json:"version"`
	ParentBatchHeader []uint8   `json:"parentBatchHeader"`
	LastBlockNumber   uint64    `json:"lastBlockNumber"`
	NumL1Messages     uint16    `json:"numL1Messages"`
	PrevStateRoot     [32]uint8 `json:"prevStateRoot"`
	PostStateRoot     [32]uint8 `json:"postStateRoot"`
	WithdrawalRoot    [32]uint8 `json:"withdrawalRoot"`
}

// convertBatchDataInput converts the parsed struct to bindings.IRollupBatchDataInput
func convertBatchDataInput(s batchDataInputStruct) *bindings.IRollupBatchDataInput {
	// Convert []uint8 to []byte
	parentBatchHeader := make([]byte, len(s.ParentBatchHeader))
	copy(parentBatchHeader, s.ParentBatchHeader)

	return &bindings.IRollupBatchDataInput{
		Version:           s.Version,
		ParentBatchHeader: parentBatchHeader,
		LastBlockNumber:   s.LastBlockNumber,
		NumL1Messages:     s.NumL1Messages,
		PrevStateRoot:     s.PrevStateRoot,
		PostStateRoot:     s.PostStateRoot,
		WithdrawalRoot:    s.WithdrawalRoot,
	}
}

// parseBatchDataInputFromArgs safely parses BatchDataInput from ABI unpacked arguments
func parseBatchDataInputFromArgs(args []interface{}) (batchDataInputStruct, error) {
	if len(args) < 1 {
		return batchDataInputStruct{}, errors.New("insufficient arguments for batch data input")
	}

	// Use comma-ok assertion for safe type checking
	rawStruct, ok := args[0].(struct {
		Version           uint8     `json:"version"`
		ParentBatchHeader []uint8   `json:"parentBatchHeader"`
		LastBlockNumber   uint64    `json:"lastBlockNumber"`
		NumL1Messages     uint16    `json:"numL1Messages"`
		PrevStateRoot     [32]uint8 `json:"prevStateRoot"`
		PostStateRoot     [32]uint8 `json:"postStateRoot"`
		WithdrawalRoot    [32]uint8 `json:"withdrawalRoot"`
	})
	if !ok {
		return batchDataInputStruct{}, errors.New("failed to cast batch data input to expected struct type")
	}

	return batchDataInputStruct{
		Version:           rawStruct.Version,
		ParentBatchHeader: rawStruct.ParentBatchHeader,
		LastBlockNumber:   rawStruct.LastBlockNumber,
		NumL1Messages:     rawStruct.NumL1Messages,
		PrevStateRoot:     rawStruct.PrevStateRoot,
		PostStateRoot:     rawStruct.PostStateRoot,
		WithdrawalRoot:    rawStruct.WithdrawalRoot,
	}, nil
}

// These two compact ABIs freeze the only historical/current commit methods this
// query supports. The pre-Submitter ABI deliberately includes the complete
// BatchSignatureInput tuple so historical calldata is decoded with its real
// layout even though callers only need BatchDataInput.
const preSubmitterCommitABI = `[
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"},
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"},{"name":"_batchHeader","type":"bytes"},{"name":"_batchProof","type":"bytes"}],"name":"commitBatchWithProof","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

const submitterCommitABI = `[
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"},
  {"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"name":"_batchHeader","type":"bytes"},{"name":"_batchProof","type":"bytes"}],"name":"commitBatchWithProof","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

// parseCommitBatchTxData decodes commitBatch and commitBatchWithProof calldata
// from either side of the Submitter upgrade. No other selector is accepted.
func parseCommitBatchTxData(txData []byte) (*bindings.IRollupBatchDataInput, error) {
	if len(txData) < 4 {
		return nil, errors.New("transaction data too short")
	}

	for _, rawABI := range []string{preSubmitterCommitABI, submitterCommitABI} {
		rollupABI, err := abi.JSON(strings.NewReader(rawABI))
		if err != nil {
			return nil, fmt.Errorf("parse frozen rollup ABI: %w", err)
		}
		method, err := rollupABI.MethodById(txData[:4])
		if err != nil {
			continue
		}
		if method.Name != "commitBatch" && method.Name != "commitBatchWithProof" {
			continue
		}
		args, err := method.Inputs.Unpack(txData[4:])
		if err != nil {
			return nil, fmt.Errorf("unpack %s calldata: %w", method.Name, err)
		}
		batchDataInputRaw, err := parseBatchDataInputFromArgs(args)
		if err != nil {
			return nil, fmt.Errorf("parse batch data input: %w", err)
		}
		return convertBatchDataInput(batchDataInputRaw), nil
	}

	return nil, errors.New("transaction is not commitBatch or commitBatchWithProof")
}
