package types

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto"
)

type BatchHeaderWithBlobHashes struct {
	BatchHeader
	blobHashes []common.Hash
}

func NewBatchHeaderWithBlobHashes(batchHeader BatchHeader, blobHashes []common.Hash) *BatchHeaderWithBlobHashes {
	return &BatchHeaderWithBlobHashes{
		BatchHeader: batchHeader,
		blobHashes:  blobHashes,
	}
}

func (b *BatchHeaderWithBlobHashes) MarshalBinary() []byte {
	var ret []byte
	ret = append(ret, byte(len(b.blobHashes)))
	for _, blobHash := range b.blobHashes {
		ret = append(ret, blobHash.Bytes()...)
	}
	ret = append(ret, b.BatchHeader.Encode()...)
	return ret
}

func (b *BatchHeaderWithBlobHashes) UnmarshalBinary(input []byte) (err error) {
	if len(input) < 90 {
		return errors.New("insufficient data for BatchHeaderWithBlobHashes")
	}
	var blobHashes []common.Hash
	blobHashCount := input[0]
	remaining := input[1:]
	for i := 0; i < int(blobHashCount); i++ {
		var hash common.Hash
		copy(hash[:], remaining[:32])
		blobHashes = append(blobHashes, hash)
		remaining = remaining[32:]
	}
	b.BatchHeader, err = DecodeBatchHeader(remaining)
	if err != nil {
		return nil
	}
	b.blobHashes = blobHashes
	return
}

func (b *BatchHeaderWithBlobHashes) BatchHash() common.Hash {
	if len(b.blobHashes) == 0 {
		return b.BatchHeader.Hash()
	}
	bytes := b.BatchHeader.Hash().Bytes()
	for _, bh := range b.blobHashes {
		bytes = append(bytes, bh.Bytes()...)
	}
	return crypto.Keccak256Hash(bytes)
}

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
	EncodedBytes hexutil.Bytes
}

// Encode encodes the BatchHeader into RollupV2 BatchHeaderV0Codec Encoding.
func (b *BatchHeader) Encode() []byte {
	if len(b.EncodedBytes) > 0 {
		return b.EncodedBytes
	}
	batchBytes := make([]byte, 89+len(b.SkippedL1MessageBitmap))
	batchBytes[0] = b.Version
	binary.BigEndian.PutUint64(batchBytes[1:], b.BatchIndex)
	binary.BigEndian.PutUint64(batchBytes[9:], b.L1MessagePopped)
	binary.BigEndian.PutUint64(batchBytes[17:], b.TotalL1MessagePopped)
	copy(batchBytes[25:], b.DataHash[:])
	copy(batchBytes[57:], b.ParentBatchHash[:])
	copy(batchBytes[89:], b.SkippedL1MessageBitmap[:])
	b.EncodedBytes = batchBytes
	return batchBytes
}

// Hash calculates the hash of the batch header.
func (b *BatchHeader) Hash() common.Hash {
	if len(b.EncodedBytes) == 0 {
		b.Encode()
	}

	return crypto.Keccak256Hash(b.EncodedBytes)
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

		EncodedBytes: data,
	}
	return b, nil
}
