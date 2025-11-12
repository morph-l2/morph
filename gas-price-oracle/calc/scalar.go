package calc

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

const (
	// Precision factor
	Precision = 1e9

	// MaxCommitScalar maximum commit scalar value
	MaxCommitScalar = 1e16 // 10^(9+7)

	// MaxBlobScalar maximum blob scalar value
	MaxBlobScalar = 1e11 // 10^(9+2)

	// FinalizeBatchGasUsed gas consumed by finalize batch
	FinalizeBatchGasUsed = 156400
)

// ScalarCalculator calculates scalar values
type ScalarCalculator struct {
	txnPerBatch uint64
}

// NewScalarCalculator creates a new scalar calculator
func NewScalarCalculator(txnPerBatch uint64) *ScalarCalculator {
	return &ScalarCalculator{
		txnPerBatch: txnPerBatch,
	}
}

// CalculateScalars calculates commit scalar and blob scalar
// rollupGasUsed: gas consumed by rollup transaction
// l2TxCount: number of L2 transactions
// l2DataLen: L2 data length (extracted from blob)
func (c *ScalarCalculator) CalculateScalars(rollupGasUsed uint64, l2TxCount uint64, l2DataLen uint64) (commitScalar, blobScalar uint64) {
	// Calculate commit scalar
	// commit_scalar = (rollup_gas_used + finalize_gas) * PRECISION / max(l2_tx_count, txn_per_batch)
	totalGas := rollupGasUsed + FinalizeBatchGasUsed
	txCount := max(l2TxCount, c.txnPerBatch)
	commitScalar = totalGas * Precision / txCount

	// Limit to maximum value
	if commitScalar > MaxCommitScalar {
		commitScalar = MaxCommitScalar
	}

	// Calculate blob scalar
	// blob_scalar = MAX_BLOB_SIZE * PRECISION / l2_data_len
	if l2DataLen > 0 {
		blobScalar = MaxBlobTxPayloadSize * Precision / l2DataLen
	} else {
		blobScalar = MaxBlobScalar
	}

	// Limit to maximum value
	if blobScalar > MaxBlobScalar {
		blobScalar = MaxBlobScalar
	}

	return commitScalar, blobScalar
}

// ExtractTxnNum extracts transaction count from chunk data
// Batch chunk format:
// |   1 byte   | 60 bytes | ... | 60 bytes |
// | num blocks |  block 1 | ... |  block n |
//
// Each block (60 bytes):
// | 8 bytes | 48 bytes | 2 bytes | 2 bytes |
// | block_num | ... | num_txs | num_l1_txs |
func ExtractTxnNum(chunks [][]byte) (uint64, error) {
	if len(chunks) == 0 {
		return 0, fmt.Errorf("empty chunks")
	}

	var totalTxn uint64
	var totalL1Txn uint64

	for _, chunk := range chunks {
		if len(chunk) < 1 {
			continue
		}

		// Read number of blocks
		numBlocks := uint64(chunk[0])

		// Parse each block information
		for i := uint64(0); i < numBlocks; i++ {
			offset := 1 + i*60
			if offset+60 > uint64(len(chunk)) {
				break
			}

			// Read transaction count (offset + 56, 2 bytes)
			numTxs := binary.BigEndian.Uint16(chunk[offset+56 : offset+58])
			totalTxn += uint64(numTxs)

			// Read L1 message transaction count (offset + 58, 2 bytes)
			numL1Txs := binary.BigEndian.Uint16(chunk[offset+58 : offset+60])
			totalL1Txn += uint64(numL1Txs)
		}
	}

	// L2 transaction count = total transactions - L1 message transactions
	if totalTxn < totalL1Txn {
		return 0, fmt.Errorf("total txn %d < l1 txn %d", totalTxn, totalL1Txn)
	}

	return totalTxn - totalL1Txn, nil
}

// ShouldUpdate determines whether an update is needed
// latest: latest calculated value
// current: current on-chain value
// threshold: threshold percentage (e.g. 10 means 10%)
func ShouldUpdate(latest, current, threshold uint64) bool {
	// If latest is zero, don't update (potentially invalid data)
	if latest == 0 {
		return false
	}

	// If current is zero, update if latest is non-zero
	if current == 0 {
		return true
	}

	actualChange := absDiff(latest, current)
	expectedChange := current * threshold / 100

	return actualChange > expectedChange
}

// ShouldUpdateBigInt determines whether an update is needed (big.Int version)
func ShouldUpdateBigInt(latest, current *big.Int, threshold uint64) bool {
	if current.Sign() == 0 {
		return latest.Sign() > 0
	}

	actualChange := new(big.Int).Abs(new(big.Int).Sub(latest, current))
	expectedChange := new(big.Int).Div(
		new(big.Int).Mul(current, big.NewInt(int64(threshold))),
		big.NewInt(100),
	)

	return actualChange.Cmp(expectedChange) > 0
}

// absDiff calculates absolute difference
func absDiff(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}

// max returns the larger of two uint64 values
func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
