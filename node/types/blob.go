package types

import (
	"errors"

	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto/kzg4844"
)

func makeBCP(bz []byte) (b kzg4844.Blob, c kzg4844.Commitment, p kzg4844.Proof, err error) {
	err = (&b).FromData(bz)
	if err != nil {
		return
	}
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

func MakeBlobTxSidecar(blobBytes []byte) (*eth.BlobTxSidecar, error) {
	if len(blobBytes) == 0 {
		return nil, nil
	}
	if len(blobBytes) > 2*kzg4844.MaxBlobDataSize {
		return nil, errors.New("only 2 blobs at most is allowed")
	}
	blobCount := len(blobBytes)/(kzg4844.MaxBlobDataSize+1) + 1
	var (
		err         error
		blobs       = make([]kzg4844.Blob, blobCount)
		commitments = make([]kzg4844.Commitment, blobCount)
		proofs      = make([]kzg4844.Proof, blobCount)
	)
	switch blobCount {
	case 1:
		blobs[0], commitments[0], proofs[0], err = makeBCP(blobBytes)
		if err != nil {
			return nil, err
		}
	case 2:
		blobs[0], commitments[0], proofs[0], err = makeBCP(blobBytes[:kzg4844.MaxBlobDataSize])
		if err != nil {
			return nil, err
		}
		blobs[1], commitments[1], proofs[1], err = makeBCP(blobBytes[kzg4844.MaxBlobDataSize:])
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
