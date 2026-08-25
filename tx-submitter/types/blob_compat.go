package types

import (
	"math/big"

	"morph-l2/common/blob"

	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
)

type BlobFeeConfig = blob.BlobFeeConfig
type BlobConfig = blob.BlobConfig
type ChainBlobConfigs = blob.ChainBlobConfigs

var (
	ChainConfigMap          = blob.ChainConfigMap
	DefaultBlobConfig       = blob.DefaultBlobConfig
	MainnetChainConfig      = blob.MainnetChainConfig
	HoodiChainConfig        = blob.HoodiChainConfig
	DevnetChainConfig       = blob.DevnetChainConfig
	DefaultCancunBlobConfig = blob.DefaultCancunBlobConfig
	DefaultPragueBlobConfig = blob.DefaultPragueBlobConfig
	DefaultOsakaBlobConfig  = blob.DefaultOsakaBlobConfig
	DefaultBPO1BlobConfig   = blob.DefaultBPO1BlobConfig
	DefaultBPO2BlobConfig   = blob.DefaultBPO2BlobConfig
	DefaultBPO3BlobConfig   = blob.DefaultBPO3BlobConfig
	DefaultBPO4BlobConfig   = blob.DefaultBPO4BlobConfig
)

func GetBlobFeeDenominator(blobFeeConfig *BlobFeeConfig, blockTime uint64) *big.Int {
	return blob.GetBlobFeeDenominator(blobFeeConfig, blockTime)
}

func BlobHashes(blobs []kzg4844.Blob, commitments []kzg4844.Commitment) []common.Hash {
	return blob.BlobHashes(blobs, commitments)
}

func MakeBlobProof(blobs []kzg4844.Blob, commitment []kzg4844.Commitment) ([]kzg4844.Proof, error) {
	return blob.MakeBlobProof(blobs, commitment)
}

func MakeCellProof(blobs []kzg4844.Blob) ([]kzg4844.Proof, error) {
	return blob.MakeCellProof(blobs)
}

func DetermineBlobVersion(head *ethtypes.Header, chainID uint64) byte {
	return blob.DetermineBlobVersion(head, chainID)
}

func BlobSidecarVersionToV1(sc *ethtypes.BlobTxSidecar) error {
	return blob.BlobSidecarVersionToV1(sc)
}
