package types

// DEPRECATED: this file is a duplicate of morph-l2/common/batch's
// batch_header.go and is kept alive only because tx-submitter/utils/utils.go
// still imports BatchHeaderBytes from here. node/types cannot be turned into
// a thin shim re-exporting common/batch because that would close an import
// cycle: common/batch already depends on tx-submitter/db (via BatchCache),
// which depends on tx-submitter/utils, which would then depend back on
// common/batch.
//
// Cleanup path (out of scope for this PR; should be done by the tx-submitter
// owners alongside moving BatchCache out of common/batch):
//  1. Move common/batch/batch_cache.go, batch_storage.go, batch_query.go
//     down to tx-submitter/batch/, so common/batch becomes a true leaf
//     (depends on nothing under tx-submitter/).
//  2. Switch tx-submitter/utils/utils.go to import morph-l2/common/batch.
//  3. Delete this file.

import (
	"encoding/binary"
	"errors"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
)

type (
	BatchHeaderBytes []byte
)

const (
	expectedLengthV0 = 249
	expectedLengthV1 = 257
	// V2 reuses the V1 wire format (257 bytes). The only semantic
	// difference is that the 32-byte field at offset 57 stores
	// keccak256(blobhash(0) || ... || blobhash(N-1)) instead of a
	// single blob versioned hash.
	expectedLengthV2 = 257

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
		if len(b) != expectedLengthV2 {
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

// BlobVersionedHash returns the EIP-4844 blob versioned hash recorded at
// offset [57:89]. This is only meaningful for V0/V1 batches, where the field
// holds the single blob's versioned hash. For V2 batches the same offset
// holds an aggregated hash; callers must use BlobHashesHash instead.
func (b BatchHeaderBytes) BlobVersionedHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	version, _ := b.Version()
	if version >= BatchHeaderVersion2 {
		return common.Hash{}, errors.New("BlobVersionedHash is not available for V2+; use BlobHashesHash")
	}
	return common.BytesToHash(b[57:89]), nil
}

// BlobHashesHash returns the aggregated blob hash recorded at offset [57:89]
// for V2+ batches, defined as keccak256(blobhash(0) || ... || blobhash(N-1)).
// V0/V1 batches do not aggregate and will return an error.
func (b BatchHeaderBytes) BlobHashesHash() (common.Hash, error) {
	if err := b.validate(); err != nil {
		return common.Hash{}, err
	}
	version, _ := b.Version()
	if version < BatchHeaderVersion2 {
		return common.Hash{}, errors.New("BlobHashesHash is only available for V2+; use BlobVersionedHash")
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
