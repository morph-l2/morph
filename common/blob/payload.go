package blob

import (
	"fmt"

	"morph-l2/node/zstd"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
)

const MaxBlobBytesSize = 4096 * 31

var (
	emptyBlob          = new(kzg4844.Blob)
	emptyBlobCommit, _ = kzg4844.BlobToCommitment(emptyBlob)
	emptyBlobProof, _  = kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
)

// EmptyVersionedHash is the versioned hash of the canonical empty blob (all-zero payload).
var EmptyVersionedHash = common.HexToHash("0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")

// MakeBlobCanonical converts the raw blob data into the canonical blob representation of 4096 BLSFieldElements.
func MakeBlobCanonical(blobBytes []byte) (b *kzg4844.Blob, err error) {
	if len(blobBytes) > MaxBlobBytesSize {
		return nil, fmt.Errorf("data is too large for blob. len=%v", len(blobBytes))
	}
	offset := 0
	b = new(kzg4844.Blob)
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

func MakeBlobTxSidecar(blobBytes []byte, maxBlobs int) (*eth.BlobTxSidecar, error) {
	if len(blobBytes) == 0 {
		return &eth.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof},
		}, nil
	}
	if maxBlobs <= 0 {
		maxBlobs = 1
	}
	if len(blobBytes) > maxBlobs*MaxBlobBytesSize {
		return nil, fmt.Errorf("data size %d exceeds %d blobs capacity (%d bytes)", len(blobBytes), maxBlobs, maxBlobs*MaxBlobBytesSize)
	}
	blobCount := (len(blobBytes) + MaxBlobBytesSize - 1) / MaxBlobBytesSize
	var (
		err         error
		blobs       = make([]kzg4844.Blob, blobCount)
		commitments = make([]kzg4844.Commitment, blobCount)
	)
	for i := 0; i < blobCount; i++ {
		start := i * MaxBlobBytesSize
		end := start + MaxBlobBytesSize
		if end > len(blobBytes) {
			end = len(blobBytes)
		}
		blobs[i], commitments[i], err = makeBlobCommitment(blobBytes[start:end])
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
