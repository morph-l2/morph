package types

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
	expectedLengthV0 = 249
	expectedLengthV1 = 257

	BatchHeaderVersion0 = 0
	BatchHeaderVersion1 = 1
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
