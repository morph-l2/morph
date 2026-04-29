package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"morph-l2/node/zstd"

	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/rlp"
)

const MaxBlobBytesSize = 4096 * 31

var (
	emptyBlob          = new(kzg4844.Blob)
	emptyBlobCommit, _ = kzg4844.BlobToCommitment(emptyBlob)
	emptyBlobProof, _  = kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
)

// MakeBlobCanonical converts the raw blob data into the canonical blob representation of 4096 BLSFieldElements.
func MakeBlobCanonical(blobBytes []byte) (b *kzg4844.Blob, err error) {
	if len(blobBytes) > MaxBlobBytesSize {
		return nil, fmt.Errorf("data is too large for blob. len=%v", len(blobBytes))
	}
	offset := 0
	b = new(kzg4844.Blob)
	// encode (up to) 31 bytes of remaining input data at a time into the subsequent field element
	for i := 0; i < 4096; i++ {
		offset += copy(b[i*32+1:i*32+32], blobBytes[offset:])
		if offset == len(blobBytes) {
			break
		}
	}
	if offset < len(blobBytes) {
		return nil, fmt.Errorf("failed to fit all data into blob. bytes remaining: %v", len(blobBytes)-offset)
	}
	return
}

func RetrieveBlobBytes(blob *kzg4844.Blob) ([]byte, error) {
	data := make([]byte, MaxBlobBytesSize)
	for i := 0; i < 4096; i++ {
		if blob[i*32] != 0 {
			return nil, fmt.Errorf("invalid blob, found non-zero high order byte %x of field element %d", data[i*32], i)
		}
		copy(data[i*31:i*31+31], blob[i*32+1:i*32+32])
	}
	return data, nil
}

func makeBlobCommitment(bz []byte) (b kzg4844.Blob, c kzg4844.Commitment, err error) {
	blob, err := MakeBlobCanonical(bz)
	if err != nil {
		return
	}
	b = *blob
	c, err = kzg4844.BlobToCommitment(&b)
	if err != nil {
		return
	}
	return
}

func MakeBlobTxSidecar(blobBytes []byte) (*eth.BlobTxSidecar, error) {
	if len(blobBytes) == 0 {
		return &eth.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof},
		}, nil
	}
	if len(blobBytes) > 2*MaxBlobBytesSize {
		return nil, errors.New("only 2 blobs at most is allowed")
	}
	blobCount := len(blobBytes)/(MaxBlobBytesSize+1) + 1
	var (
		err         error
		blobs       = make([]kzg4844.Blob, blobCount)
		commitments = make([]kzg4844.Commitment, blobCount)
	)
	switch blobCount {
	case 1:
		blobs[0], commitments[0], err = makeBlobCommitment(blobBytes)
		if err != nil {
			return nil, err
		}
	case 2:
		blobs[0], commitments[0], err = makeBlobCommitment(blobBytes[:MaxBlobBytesSize])
		if err != nil {
			return nil, err
		}
		blobs[1], commitments[1], err = makeBlobCommitment(blobBytes[MaxBlobBytesSize:])
		if err != nil {
			return nil, err
		}
	}
	return &eth.BlobTxSidecar{
		Blobs:       blobs,
		Commitments: commitments,
	}, nil
}

func CompressBatchBytes(batchBytes []byte) ([]byte, error) {
	if len(batchBytes) == 0 {
		return nil, nil
	}
	compressedBatchBytes, err := zstd.CompressBatchBytes(batchBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to compress batch bytes, err: %w", err)
	}
	return compressedBatchBytes, nil
}

func DecodeTxsFromBytes(txsBytes []byte) (eth.Transactions, error) {
	reader := bytes.NewReader(txsBytes)
	txs := make(eth.Transactions, 0)
	for {
		var (
			typeByte byte
			err      error
		)
		if err = binary.Read(reader, binary.BigEndian, &typeByte); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if typeByte == 0 {
			break
		}

		switch typeByte {
		case eth.AccessListTxType, eth.DynamicFeeTxType, eth.SetCodeTxType:
			tx, err := decodeTypedTx(typeByte, reader)
			if err != nil {
				return nil, err
			}
			txs = append(txs, tx)

		case eth.MorphTxType:
			tx, err := decodeMorphTx(reader)
			if err != nil {
				return nil, err
			}
			txs = append(txs, tx)

		default:
			if typeByte <= 0xf7 {
				return nil, fmt.Errorf("not supported tx type: %d", typeByte)
			}
			fullTxBytes, err := extractInnerTxFullBytes(typeByte, reader)
			if err != nil {
				return nil, err
			}
			var inner eth.LegacyTx
			if err = rlp.DecodeBytes(fullTxBytes, &inner); err != nil {
				return nil, err
			}
			txs = append(txs, eth.NewTx(&inner))
		}
	}
	return txs, nil
}

// decodeTypedTx decodes a standard EIP-2718 typed tx (AccessList, DynamicFee, SetCode)
// from the reader. The type byte has already been consumed; the next byte is the RLP prefix.
func decodeTypedTx(typeByte byte, reader io.Reader) (*eth.Transaction, error) {
	var rlpPrefix byte
	if err := binary.Read(reader, binary.BigEndian, &rlpPrefix); err != nil {
		return nil, err
	}
	rlpBytes, err := extractInnerTxFullBytes(rlpPrefix, reader)
	if err != nil {
		return nil, err
	}
	txBinary := make([]byte, 0, 1+len(rlpBytes))
	txBinary = append(txBinary, typeByte)
	txBinary = append(txBinary, rlpBytes...)

	var tx eth.Transaction
	if err := tx.UnmarshalBinary(txBinary); err != nil {
		return nil, err
	}
	return &tx, nil
}

// decodeMorphTx decodes a MorphTx from the reader. The type byte (0x7f) has already
// been consumed. MorphTx has two wire formats:
//   - V0: type(0x7f) || RLP(fields)              — next byte is RLP prefix (>= 0xC0)
//   - V1: type(0x7f) || version(0x01) || RLP(fields) — next byte is version, then RLP prefix
func decodeMorphTx(reader io.Reader) (*eth.Transaction, error) {
	var nextByte byte
	if err := binary.Read(reader, binary.BigEndian, &nextByte); err != nil {
		return nil, err
	}

	var versionPrefix []byte
	rlpFirstByte := nextByte
	if nextByte != 0 && nextByte < 0xC0 {
		// V1+: nextByte is the version byte, read the actual RLP prefix
		versionPrefix = []byte{nextByte}
		if err := binary.Read(reader, binary.BigEndian, &rlpFirstByte); err != nil {
			return nil, err
		}
	}

	rlpBytes, err := extractInnerTxFullBytes(rlpFirstByte, reader)
	if err != nil {
		return nil, err
	}

	txBinary := make([]byte, 0, 1+len(versionPrefix)+len(rlpBytes))
	txBinary = append(txBinary, eth.MorphTxType)
	txBinary = append(txBinary, versionPrefix...)
	txBinary = append(txBinary, rlpBytes...)

	var tx eth.Transaction
	if err := tx.UnmarshalBinary(txBinary); err != nil {
		return nil, err
	}
	return &tx, nil
}

func extractInnerTxFullBytes(firstByte byte, reader io.Reader) ([]byte, error) {
	sizeByteLen := firstByte - 0xf7
	if sizeByteLen > 4 {
		return nil, fmt.Errorf("invalid RLP size byte length: %d (firstByte=0x%x)", sizeByteLen, firstByte)
	}

	sizeByte := make([]byte, sizeByteLen)
	if err := binary.Read(reader, binary.BigEndian, sizeByte); err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint32(append(make([]byte, 4-len(sizeByte)), sizeByte...))

	txRaw := make([]byte, size)
	if err := binary.Read(reader, binary.BigEndian, txRaw); err != nil {
		return nil, err
	}
	fullTxBytes := make([]byte, 1+uint32(sizeByteLen)+size)
	copy(fullTxBytes[:1], []byte{firstByte})
	copy(fullTxBytes[1:1+sizeByteLen], sizeByte)
	copy(fullTxBytes[1+sizeByteLen:], txRaw)

	return fullTxBytes, nil
}
