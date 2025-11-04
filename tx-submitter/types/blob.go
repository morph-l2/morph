package types

import (
	"crypto/sha256"

	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
)

// BlobHashes computes the blob hashes of the given blobs.
func BlobHashes(blobs []kzg4844.Blob, commitments []kzg4844.Commitment) []common.Hash {
	hasher := sha256.New()
	h := make([]common.Hash, len(commitments))
	for i := range blobs {
		h[i] = kzg4844.CalcBlobHashV1(hasher, &commitments[i])
	}
	return h
}

func MakeBlobProof(blobs []kzg4844.Blob, commitment []kzg4844.Commitment) (p []kzg4844.Proof, err error) {
	p = make([]kzg4844.Proof, 0, len(blobs))
	for i, _ := range blobs {
		p[i], err = kzg4844.ComputeBlobProof(&blobs[i], commitment[i])
		if err != nil {
			return nil, err
		}
	}
	return
}

func MakeCellProof(blobs []kzg4844.Blob) ([]kzg4844.Proof, error) {
	proofs := make([]kzg4844.Proof, 0, len(blobs)*kzg4844.CellProofsPerBlob)
	for _, blob := range blobs {
		cellProofs, err := kzg4844.ComputeCellProofs(&blob)
		if err != nil {
			return nil, err
		}
		proofs = append(proofs, cellProofs...)
	}
	return proofs, nil
}

func DetermineBlobVersion(head *ethtypes.Header, chainID uint64) byte {
	if head == nil {
		return ethtypes.BlobSidecarVersion0
	}
	blobConfig, exist := ChainConfigMap[chainID]
	if !exist {
		blobConfig = DefaultBlobConfig
	}
	if blobConfig.OsakaTime != nil && blobConfig.IsOsaka(head.Number, head.Time) {
		return ethtypes.BlobSidecarVersion1
	}
	return ethtypes.BlobSidecarVersion0
}

// BlobSidecarVersionToV1 converts the BlobSidecar to version 1, attaching the cell proofs.
func BlobSidecarVersionToV1(sc *ethtypes.BlobTxSidecar) error {
	if sc.Version == ethtypes.BlobSidecarVersion1 {
		return nil
	}
	if sc.Version == ethtypes.BlobSidecarVersion0 {
		proofs := make([]kzg4844.Proof, 0, len(sc.Blobs)*kzg4844.CellProofsPerBlob)
		for _, blob := range sc.Blobs {
			cellProofs, err := kzg4844.ComputeCellProofs(&blob)
			if err != nil {
				return err
			}
			proofs = append(proofs, cellProofs...)
		}
		sc.Version = ethtypes.BlobSidecarVersion1
		sc.Proofs = proofs
	}
	return nil
}
