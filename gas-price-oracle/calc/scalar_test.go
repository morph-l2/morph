package calc

import (
	"testing"
)

func TestScalarCalculator_CalculateScalars(t *testing.T) {
	calc := NewScalarCalculator(50)

	tests := []struct {
		name          string
		rollupGasUsed uint64
		l2TxCount     uint64
		l2DataLen     uint64
		wantCommit    uint64
		wantBlob      uint64
	}{
		{
			name:          "normal case",
			rollupGasUsed: 1000000,
			l2TxCount:     100,
			l2DataLen:     100000,
			wantCommit:    (1000000 + FinalizeBatchGasUsed) * Precision / 100,
			wantBlob:      MaxBlobTxPayloadSize * Precision / 100000,
		},
		{
			name:          "low tx count",
			rollupGasUsed: 500000,
			l2TxCount:     10,
			l2DataLen:     50000,
			wantCommit:    (500000 + FinalizeBatchGasUsed) * Precision / 50, // Using txnPerBatch
			wantBlob:      MaxBlobTxPayloadSize * Precision / 50000,
		},
		{
			name:          "zero data len",
			rollupGasUsed: 1000000,
			l2TxCount:     100,
			l2DataLen:     0,
			wantCommit:    (1000000 + FinalizeBatchGasUsed) * Precision / 100,
			wantBlob:      MaxBlobScalar, // Should return maximum value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commitScalar, blobScalar := calc.CalculateScalars(
				tt.rollupGasUsed,
				tt.l2TxCount,
				tt.l2DataLen,
			)

			if commitScalar != tt.wantCommit {
				t.Errorf("commitScalar = %v, want %v", commitScalar, tt.wantCommit)
			}

			if blobScalar != tt.wantBlob {
				t.Errorf("blobScalar = %v, want %v", blobScalar, tt.wantBlob)
			}
		})
	}
}

func TestExtractTxnNum(t *testing.T) {
	// Construct test data
	// Format: [num_blocks(1)] [block_info(60)]...
	// block_info: [block_num(8)] [other(48)] [num_txs(2)] [num_l1_txs(2)]

	chunk := make([]byte, 61)
	chunk[0] = 1 // 1 block

	// Block 1: block_num=100, num_txs=50, num_l1_txs=5
	// Position: 1 + 56 = 57 (num_txs)
	chunk[57] = 0
	chunk[58] = 50
	// Position: 1 + 58 = 59 (num_l1_txs)
	chunk[59] = 0
	chunk[60] = 5

	chunks := [][]byte{chunk}

	txCount, err := ExtractTxnNum(chunks)
	if err != nil {
		t.Fatalf("ExtractTxnNum failed: %v", err)
	}

	expected := uint64(45) // 50 - 5 = 45
	if txCount != expected {
		t.Errorf("txCount = %v, want %v", txCount, expected)
	}
}

func TestShouldUpdate(t *testing.T) {
	tests := []struct {
		name      string
		latest    uint64
		current   uint64
		threshold uint64
		want      bool
	}{
		{
			name:      "should update - 20% change with 10% threshold",
			latest:    120,
			current:   100,
			threshold: 10,
			want:      true,
		},
		{
			name:      "should not update - 5% change with 10% threshold",
			latest:    105,
			current:   100,
			threshold: 10,
			want:      false,
		},
		{
			name:      "should update - current is zero",
			latest:    100,
			current:   0,
			threshold: 10,
			want:      true,
		},
		{
			name:      "should not update - latest is zero",
			latest:    0,
			current:   100,
			threshold: 10,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldUpdate(tt.latest, tt.current, tt.threshold)
			if got != tt.want {
				t.Errorf("ShouldUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
