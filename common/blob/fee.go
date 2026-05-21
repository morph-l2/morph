package blob

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/log"
)

// BlobFeeConfig is used to configure blob fee calculation parameters.
type BlobFeeConfig struct {
	ChainID *big.Int

	// LondonBlock is the London block number (used for time determination).
	LondonBlock *big.Int

	// Timestamps for each EIP upgrade (Unix timestamp, seconds).
	CancunTime *uint64 // EIP-4844 (Cancun)
	PragueTime *uint64 // Prague
	OsakaTime  *uint64 // Osaka
	BPO1Time   *uint64 // BPO1
	BPO2Time   *uint64 // BPO2
	BPO3Time   *uint64 // BPO3
	BPO4Time   *uint64 // BPO4
	BPO5Time   *uint64 // BPO5

	// BlobConfig corresponding to each EIP.
	Cancun *BlobConfig
	Prague *BlobConfig
	Osaka  *BlobConfig
	BPO1   *BlobConfig
	BPO2   *BlobConfig
	BPO3   *BlobConfig
	BPO4   *BlobConfig
	BPO5   *BlobConfig

	Default *BlobConfig
}

// BlobConfig contains the parameters required for blob fee calculation.
type BlobConfig struct {
	// UpdateFraction is the denominator in the fakeExponential function.
	UpdateFraction uint64
}

// IsCancun returns whether time is either equal to the Cancun fork time or greater.
func (c *BlobFeeConfig) IsCancun(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.CancunTime, time)
}

// IsPrague returns whether time is either equal to the Prague fork time or greater.
func (c *BlobFeeConfig) IsPrague(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.PragueTime, time)
}

// IsOsaka returns whether time is either equal to the Osaka fork time or greater.
func (c *BlobFeeConfig) IsOsaka(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.OsakaTime, time)
}

// IsBPO1 returns whether time is either equal to the BPO1 fork time or greater.
func (c *BlobFeeConfig) IsBPO1(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.BPO1Time, time)
}

// IsBPO2 returns whether time is either equal to the BPO2 fork time or greater.
func (c *BlobFeeConfig) IsBPO2(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.BPO2Time, time)
}

// IsBPO3 returns whether time is either equal to the BPO3 fork time or greater.
func (c *BlobFeeConfig) IsBPO3(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.BPO3Time, time)
}

// IsBPO4 returns whether time is either equal to the BPO4 fork time or greater.
func (c *BlobFeeConfig) IsBPO4(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.BPO4Time, time)
}

// IsBPO5 returns whether time is either equal to the BPO5 fork time or greater.
func (c *BlobFeeConfig) IsBPO5(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.BPO5Time, time)
}

// IsLondon returns whether num is either equal to the London fork block or greater.
func (c *BlobFeeConfig) IsLondon(num *big.Int) bool {
	return isBlockForked(c.LondonBlock, num)
}

// GetBlobFeeDenominator returns the corresponding UpdateFraction based on the time.
func GetBlobFeeDenominator(blobFeeConfig *BlobFeeConfig, blockTime uint64) *big.Int {
	if blobFeeConfig == nil {
		log.Warn("BlobFeeConfig not set, using default denominator",
			"default", DefaultOsakaBlobConfig)
		return new(big.Int).SetUint64(DefaultOsakaBlobConfig.UpdateFraction)
	}

	cfg := blobFeeConfig
	londonBlock := cfg.LondonBlock

	var blobConfig *BlobConfig

	if cfg.BPO5Time != nil && cfg.IsBPO5(londonBlock, blockTime) && cfg.BPO5 != nil {
		blobConfig = cfg.BPO5
	} else if cfg.BPO4Time != nil && cfg.IsBPO4(londonBlock, blockTime) && cfg.BPO4 != nil {
		blobConfig = cfg.BPO4
	} else if cfg.BPO3Time != nil && cfg.IsBPO3(londonBlock, blockTime) && cfg.BPO3 != nil {
		blobConfig = cfg.BPO3
	} else if cfg.BPO2Time != nil && cfg.IsBPO2(londonBlock, blockTime) && cfg.BPO2 != nil {
		blobConfig = cfg.BPO2
	} else if cfg.BPO1Time != nil && cfg.IsBPO1(londonBlock, blockTime) && cfg.BPO1 != nil {
		blobConfig = cfg.BPO1
	} else if cfg.OsakaTime != nil && cfg.IsOsaka(londonBlock, blockTime) && cfg.Osaka != nil {
		blobConfig = cfg.Osaka
	} else if cfg.PragueTime != nil && cfg.IsPrague(londonBlock, blockTime) && cfg.Prague != nil {
		blobConfig = cfg.Prague
	} else if cfg.CancunTime != nil && cfg.IsCancun(londonBlock, blockTime) && cfg.Cancun != nil {
		blobConfig = cfg.Cancun
	} else if cfg.Default != nil {
		blobConfig = cfg.Default
	}

	if blobConfig == nil {
		log.Warn("No blob config found for current time, using default",
			"blockTime", blockTime,
			"londonBlock", londonBlock,
			"default", DefaultOsakaBlobConfig)
		return new(big.Int).SetUint64(DefaultOsakaBlobConfig.UpdateFraction)
	}

	return new(big.Int).SetUint64(blobConfig.UpdateFraction)
}

func isBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func isTimestampForked(s *uint64, head uint64) bool {
	if s == nil {
		return false
	}
	return *s <= head
}

func newUint64(val uint64) *uint64 { return &val }

var (
	DefaultCancunBlobConfig = &BlobConfig{
		UpdateFraction: 3338477,
	}
	DefaultPragueBlobConfig = &BlobConfig{
		UpdateFraction: 5007716,
	}
	DefaultOsakaBlobConfig = &BlobConfig{
		UpdateFraction: 5007716,
	}
	DefaultBPO1BlobConfig = &BlobConfig{
		UpdateFraction: 8346193,
	}
	DefaultBPO2BlobConfig = &BlobConfig{
		UpdateFraction: 11684671,
	}
	DefaultBPO3BlobConfig = &BlobConfig{
		UpdateFraction: 20609697,
	}
	DefaultBPO4BlobConfig = &BlobConfig{
		UpdateFraction: 13739630,
	}
)

var (
	MainnetChainConfig = &BlobFeeConfig{
		ChainID:     big.NewInt(1),
		LondonBlock: big.NewInt(12_965_000),
		CancunTime:  newUint64(1710338135),
		PragueTime:  newUint64(1746612311),
		OsakaTime:   newUint64(1764798551),
		BPO1Time:    newUint64(1765290071),
		BPO2Time:    newUint64(1767747671),
		Cancun:      DefaultCancunBlobConfig,
		Prague:      DefaultPragueBlobConfig,
		Osaka:       DefaultOsakaBlobConfig,
		BPO1:        DefaultBPO1BlobConfig,
		BPO2:        DefaultBPO2BlobConfig,
		Default:     DefaultOsakaBlobConfig,
	}

	HoodiChainConfig = &BlobFeeConfig{
		ChainID:     big.NewInt(560048),
		LondonBlock: big.NewInt(0),
		CancunTime:  newUint64(0),
		PragueTime:  newUint64(1742999832),
		OsakaTime:   newUint64(1761677592),
		BPO1Time:    newUint64(1762365720),
		BPO2Time:    newUint64(1762955544),
		Cancun:      DefaultCancunBlobConfig,
		Prague:      DefaultPragueBlobConfig,
		Osaka:       DefaultOsakaBlobConfig,
		BPO1:        DefaultBPO1BlobConfig,
		BPO2:        DefaultBPO2BlobConfig,
		Default:     DefaultOsakaBlobConfig,
	}

	DevnetChainConfig = &BlobFeeConfig{
		ChainID:     big.NewInt(900),
		LondonBlock: big.NewInt(0),
		CancunTime:  newUint64(0),
		PragueTime:  newUint64(1742999832),
		OsakaTime:   newUint64(1761677592),
		BPO1Time:    newUint64(1762365720),
		BPO2Time:    newUint64(1762955544),
		Cancun:      DefaultCancunBlobConfig,
		Prague:      DefaultPragueBlobConfig,
		Osaka:       DefaultOsakaBlobConfig,
		BPO1:        DefaultBPO1BlobConfig,
		BPO2:        DefaultBPO2BlobConfig,
		Default:     DefaultOsakaBlobConfig,
	}
)

// ChainBlobConfigs maps chain ID to blob fee configuration.
type ChainBlobConfigs map[uint64]*BlobFeeConfig

var (
	DefaultBlobConfig = HoodiChainConfig

	ChainConfigMap = ChainBlobConfigs{
		1:      MainnetChainConfig,
		560048: HoodiChainConfig,
		900:    DevnetChainConfig,
	}
)

// BlobHashes computes the blob hashes of the given blobs.
func BlobHashes(blobs []kzg4844.Blob, commitments []kzg4844.Commitment) []common.Hash {
	hasher := sha256.New()
	h := make([]common.Hash, len(commitments))
	for i := range commitments {
		h[i] = kzg4844.CalcBlobHashV1(hasher, &commitments[i])
	}
	return h
}

// MakeBlobProof builds KZG proofs for blob transactions (sidecar v0).
func MakeBlobProof(blobs []kzg4844.Blob, commitment []kzg4844.Commitment) ([]kzg4844.Proof, error) {
	if len(blobs) != len(commitment) {
		return nil, fmt.Errorf("blob/commitment length mismatch: %d != %d", len(blobs), len(commitment))
	}
	proofs := make([]kzg4844.Proof, len(blobs))
	for i := range blobs {
		proof, err := kzg4844.ComputeBlobProof(&blobs[i], commitment[i])
		if err != nil {
			return nil, err
		}
		proofs[i] = proof
	}
	return proofs, nil
}

// MakeCellProof builds cell proofs for blob sidecar v1.
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

// DetermineBlobVersion selects blob sidecar version from header time and chain config.
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
		proofs, err := MakeCellProof(sc.Blobs)
		if err != nil {
			return err
		}
		sc.Version = ethtypes.BlobSidecarVersion1
		sc.Proofs = proofs
	}
	return nil
}
