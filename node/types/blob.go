package types

import (
	"encoding/binary"
	"errors"
	"fmt"

	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto/kzg4844"
)

const MaxBlobTxPayloadSize = 4096 * 31

func BlobFromSealedTxPayload(sealedTxPayload []byte) (b *kzg4844.Blob, err error) {
	if len(sealedTxPayload) > MaxBlobTxPayloadSize {
		return nil, fmt.Errorf("data is too large for blob. len=%v", len(sealedTxPayload))
	}
	if len(sealedTxPayload)%31 != 0 {
		return nil, fmt.Errorf("data length has to be a multiple of 31")
	}
	offset := 0
	b = new(kzg4844.Blob)
	// encode (up to) 31 bytes of remaining input data at a time into the subsequent field element
	for i := 0; i < 4096; i++ {
		offset += copy(b[i*32+1:i*32+32], sealedTxPayload[offset:])
		if offset == len(sealedTxPayload) {
			break
		}
	}
	if offset < len(sealedTxPayload) {
		return nil, fmt.Errorf("failed to fit all data into blob. bytes remaining: %v", len(sealedTxPayload)-offset)
	}
	return
}

func DecodeRawTxPayload(b *kzg4844.Blob) ([]byte, error) {
	data := make([]byte, MaxBlobTxPayloadSize)
	for i := 0; i < 4096; i++ {
		if b[i*32] != 0 {
			return nil, fmt.Errorf("invalid blob, found non-zero high order byte %x of field element %d", b[i*32], i)
		}
		copy(data[i*31:i*31+31], b[i*32+1:i*32+32])
	}

	var offset uint32
	var chunkIndex uint16
	var payload []byte
	for {
		if offset >= MaxBlobTxPayloadSize {
			break
		}
		dataLen := binary.LittleEndian.Uint32(data[offset : offset+4])
		remainingLen := MaxBlobTxPayloadSize - offset - 4
		if dataLen > remainingLen {
			return nil, fmt.Errorf("decode error: dataLen is bigger than remainingLen. chunkIndex: %d, dataLen: %d, remaingLen: %d", chunkIndex, dataLen, remainingLen)
		}
		payload = append(payload, data[offset+4:offset+4+dataLen]...)

		ret := (4 + dataLen) / 31
		remainder := (4 + dataLen) % 31
		if remainder > 0 {
			ret += 1
		}
		offset += ret * 31
		chunkIndex++
	}
	return payload, nil
}

func makeBCP(bz []byte) (b kzg4844.Blob, c kzg4844.Commitment, p kzg4844.Proof, err error) {
	blob, err := BlobFromSealedTxPayload(bz)
	if err != nil {
		return
	}
	b = *blob
	c, err = kzg4844.BlobToCommitment(b)
	if err != nil {
		return
	}
	p, err = kzg4844.ComputeBlobProof(b, c)
	if err != nil {
		return
	}
	return
}

func MakeBlobTxSidecarWithTxPayload(sealedTxPayload []byte) (*eth.BlobTxSidecar, error) {
	if len(sealedTxPayload) == 0 {
		return nil, nil
	}
	if len(sealedTxPayload) > 2*MaxBlobTxPayloadSize {
		return nil, errors.New("only 2 blobs at most is allowed")
	}
	blobCount := len(sealedTxPayload)/(MaxBlobTxPayloadSize+1) + 1
	var (
		err         error
		blobs       = make([]kzg4844.Blob, blobCount)
		commitments = make([]kzg4844.Commitment, blobCount)
		proofs      = make([]kzg4844.Proof, blobCount)
	)
	switch blobCount {
	case 1:
		blobs[0], commitments[0], proofs[0], err = makeBCP(sealedTxPayload)
		if err != nil {
			return nil, err
		}
	case 2:
		blobs[0], commitments[0], proofs[0], err = makeBCP(sealedTxPayload[:MaxBlobTxPayloadSize])
		if err != nil {
			return nil, err
		}
		blobs[1], commitments[1], proofs[1], err = makeBCP(sealedTxPayload[MaxBlobTxPayloadSize:])
		if err != nil {
			return nil, err
		}
	}
	return &eth.BlobTxSidecar{
		Blobs:       blobs,
		Commitments: commitments,
		Proofs:      proofs,
	}, nil
}
