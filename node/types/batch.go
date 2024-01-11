package types

import (
	"encoding/binary"
	"fmt"
	"github.com/scroll-tech/go-ethereum/common/hexutil"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/crypto"
)

type BatchHeader struct {
	// Encoded in BatchHeaderV0Codec
	Version                uint8
	BatchIndex             uint64
	L1MessagePopped        uint64
	TotalL1MessagePopped   uint64
	DataHash               common.Hash
	ParentBatchHash        common.Hash
	SkippedL1MessageBitmap hexutil.Bytes

	//cache
	Bytes hexutil.Bytes
}

// Encode encodes the BatchHeader into RollupV2 BatchHeaderV0Codec Encoding.
func (b *BatchHeader) Encode() []byte {
	if len(b.Bytes) > 0 {
		return b.Bytes
	}
	batchBytes := make([]byte, 89+len(b.SkippedL1MessageBitmap))
	batchBytes[0] = b.Version
	binary.BigEndian.PutUint64(batchBytes[1:], b.BatchIndex)
	binary.BigEndian.PutUint64(batchBytes[9:], b.L1MessagePopped)
	binary.BigEndian.PutUint64(batchBytes[17:], b.TotalL1MessagePopped)
	copy(batchBytes[25:], b.DataHash[:])
	copy(batchBytes[57:], b.ParentBatchHash[:])
	copy(batchBytes[89:], b.SkippedL1MessageBitmap[:])
	b.Bytes = batchBytes
	return batchBytes
}

// Hash calculates the hash of the batch header.
func (b *BatchHeader) Hash() common.Hash {
	if len(b.Bytes) == 0 {
		b.Encode()
	}
	return crypto.Keccak256Hash(b.Bytes)
}

// DecodeBatchHeader attempts to decode the given byte slice into a BatchHeader.
func DecodeBatchHeader(data []byte) (BatchHeader, error) {
	if len(data) < 89 {
		return BatchHeader{}, fmt.Errorf("insufficient data for BatchHeader")
	}
	b := BatchHeader{
		Version: data[0],

		BatchIndex:             binary.BigEndian.Uint64(data[1:9]),
		L1MessagePopped:        binary.BigEndian.Uint64(data[9:17]),
		TotalL1MessagePopped:   binary.BigEndian.Uint64(data[17:25]),
		DataHash:               common.BytesToHash(data[25:57]),
		ParentBatchHash:        common.BytesToHash(data[57:89]),
		SkippedL1MessageBitmap: data[89:],

		Bytes: data,
	}
	return b, nil
}
