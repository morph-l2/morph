package batch

import (
	"encoding/binary"
	"errors"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto"
)

type (
	BatchHeaderBytes []byte
)

const (
	expectedLengthV0    = 249
	expectedLengthV1    = 257
	expectedLengthV2Min = 258 // V2 minimum: V1(257) + blobCount(1)

	BatchHeaderVersion0 = 0
	BatchHeaderVersion1 = 1
	BatchHeaderVersion2 = 2
)

var (
	ErrInvalidBatchHeaderLength  = errors.New("invalid BatchHeaderBytes length")
	ErrInvalidBatchHeaderVersion = errors.New("invalid BatchHeaderBytes version")
	ErrEmptyBatchHeaderBytes     = errors.New("empty BatchHeaderBytes")
	ErrNotFoundInBatchHeader     = errors.New("not found in BatchHeaderBytes")
)

func (b BatchHeaderBytes) validate() error {
	version, err := b.Version()
	if err != nil {
		return err
	}
	switch version {
	case BatchHeaderVersion0:
		if len(b) != expectedLengthV0 {
			return ErrInvalidBatchHeaderLength
		}
	case BatchHeaderVersion1:
		if len(b) != expectedLengthV1 {
			return ErrInvalidBatchHeaderLength
		}
	case BatchHeaderVersion2:
		if len(b) < expectedLengthV2Min {
			return ErrInvalidBatchHeaderLength
		}
		blobCount := b[257]
		if blobCount == 0 {
			return ErrInvalidBatchHeaderLength
		}
		expectedLen := expectedLengthV1 + 1 + int(blobCount-1)*32
		if len(b) != expectedLen {
			return ErrInvalidBatchHeaderLength
		}
	default:
		return ErrInvalidBatchHeaderVersion
	}
	return nil
}

func (b BatchHeaderBytes) Bytes() []byte {
	return b[:]
}

func (b BatchHeaderBytes) Hash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(b), nil
}

func (b BatchHeaderBytes) Version() (uint8, error) {
	if len(b) == 0 {
		return 0, ErrEmptyBatchHeaderBytes
	}
	return b[0], nil
}

func (b BatchHeaderBytes) BatchIndex() (uint64, error) {
	if err := b.validate(); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[1:9]), nil
}

func (b BatchHeaderBytes) L1MessagePopped() (uint64, error) {
	if err := b.validate(); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[9:17]), nil
}

func (b BatchHeaderBytes) TotalL1MessagePopped() (uint64, error) {
	if err := b.validate(); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[17:25]), nil
}

func (b BatchHeaderBytes) DataHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[25:57]), nil
}

func (b BatchHeaderBytes) BlobVersionedHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[57:89]), nil
}

func (b BatchHeaderBytes) PrevStateRoot() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[89:121]), nil
}

func (b BatchHeaderBytes) PostStateRoot() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[121:153]), nil
}

func (b BatchHeaderBytes) WithdrawalRoot() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[153:185]), nil
}

func (b BatchHeaderBytes) SequencerSetVerifyHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[185:217]), nil
}

func (b BatchHeaderBytes) ParentBatchHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(b[217:249]), nil
}

func (b BatchHeaderBytes) LastBlockNumber() (uint64, error) {
	if err := b.validate(); err != nil {
		return 0, err
	}
	version, _ := b.Version()
	if version < 1 {
		return 0, errors.New("LastBlockNumber is not available in version 0")
	}
	return binary.BigEndian.Uint64(b[249:257]), nil
}

// structed batch header for version 0
type BatchHeaderV0 struct {
	BatchIndex             uint64
	L1MessagePopped        uint64
	TotalL1MessagePopped   uint64
	DataHash               common.Hash
	BlobVersionedHash      common.Hash
	PrevStateRoot          common.Hash
	PostStateRoot          common.Hash
	WithdrawalRoot         common.Hash
	SequencerSetVerifyHash common.Hash
	ParentBatchHash        common.Hash

	//cache
	EncodedBytes hexutil.Bytes
}

func (b BatchHeaderV0) Bytes() BatchHeaderBytes {
	if len(b.EncodedBytes) > 0 {
		return BatchHeaderBytes(b.EncodedBytes)
	}
	batchBytes := make([]byte, expectedLengthV0)
	batchBytes[0] = BatchHeaderVersion0
	binary.BigEndian.PutUint64(batchBytes[1:], b.BatchIndex)
	binary.BigEndian.PutUint64(batchBytes[9:], b.L1MessagePopped)
	binary.BigEndian.PutUint64(batchBytes[17:], b.TotalL1MessagePopped)
	copy(batchBytes[25:], b.DataHash[:])
	copy(batchBytes[57:], b.BlobVersionedHash[:])
	copy(batchBytes[89:], b.PrevStateRoot[:])
	copy(batchBytes[121:], b.PostStateRoot[:])
	copy(batchBytes[153:], b.WithdrawalRoot[:])
	copy(batchBytes[185:], b.SequencerSetVerifyHash[:])
	copy(batchBytes[217:], b.ParentBatchHash[:])
	b.EncodedBytes = batchBytes
	return batchBytes
}

type BatchHeaderV1 struct {
	BatchHeaderV0
	LastBlockNumber uint64

	//cache
	EncodedBytes hexutil.Bytes
}

func (b BatchHeaderV1) Bytes() BatchHeaderBytes {
	if len(b.EncodedBytes) > 0 {
		return BatchHeaderBytes(b.EncodedBytes)
	}
	batchBytes := make([]byte, expectedLengthV1)
	batchBytes[0] = BatchHeaderVersion1
	binary.BigEndian.PutUint64(batchBytes[1:], b.BatchIndex)
	binary.BigEndian.PutUint64(batchBytes[9:], b.L1MessagePopped)
	binary.BigEndian.PutUint64(batchBytes[17:], b.TotalL1MessagePopped)
	copy(batchBytes[25:], b.DataHash[:])
	copy(batchBytes[57:], b.BlobVersionedHash[:])
	copy(batchBytes[89:], b.PrevStateRoot[:])
	copy(batchBytes[121:], b.PostStateRoot[:])
	copy(batchBytes[153:], b.WithdrawalRoot[:])
	copy(batchBytes[185:], b.SequencerSetVerifyHash[:])
	copy(batchBytes[217:], b.ParentBatchHash[:])
	binary.BigEndian.PutUint64(batchBytes[249:], b.LastBlockNumber)

	b.EncodedBytes = batchBytes
	return batchBytes
}

// BatchHeaderV2 extends V1 with a blobCount field and additional blob versioned hashes.
// Format: V1(257B) + blobCount(1B) + blobHash[1..N-1]((N-1)*32B)
// blobHash[0] is retained at the V0/V1 offset 57 for backward compatibility.
type BatchHeaderV2 struct {
	BatchHeaderV1
	BlobCount       uint8
	ExtraBlobHashes []common.Hash // blobHash[1..N-1], does not include blobHash[0]

	//cache
	EncodedBytes hexutil.Bytes
}

func (b BatchHeaderV2) Bytes() BatchHeaderBytes {
	if len(b.EncodedBytes) > 0 {
		return BatchHeaderBytes(b.EncodedBytes)
	}
	// Total size: 257 (V1) + 1 (blobCount) + (N-1)*32 (extra hashes)
	size := expectedLengthV1 + 1 + len(b.ExtraBlobHashes)*32
	batchBytes := make([]byte, size)
	// Copy V1 fields (uses BatchHeaderV1's own Bytes() for correct encoding)
	v1Bytes := b.BatchHeaderV1.Bytes()
	copy(batchBytes, v1Bytes)
	// Override version byte
	batchBytes[0] = BatchHeaderVersion2
	// Write blobCount at offset 257
	batchBytes[257] = b.BlobCount
	// Write extra blob hashes starting at offset 258
	for i, h := range b.ExtraBlobHashes {
		copy(batchBytes[258+i*32:], h[:])
	}
	b.EncodedBytes = batchBytes
	return BatchHeaderBytes(batchBytes)
}

// BlobVersionedHashes returns all blob versioned hashes.
// For V0/V1, returns a single-element slice (the hash at offset 57).
// For V2, returns all N hashes: [blobHash[0], blobHash[1], ..., blobHash[N-1]].
func (b BatchHeaderBytes) BlobVersionedHashes() ([]common.Hash, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	version, _ := b.Version()
	if version < BatchHeaderVersion2 {
		return []common.Hash{common.BytesToHash(b[57:89])}, nil
	}
	blobCount := int(b[257])
	hashes := make([]common.Hash, blobCount)
	hashes[0] = common.BytesToHash(b[57:89])
	for i := 1; i < blobCount; i++ {
		offset := 258 + (i-1)*32
		hashes[i] = common.BytesToHash(b[offset : offset+32])
	}
	return hashes, nil
}
