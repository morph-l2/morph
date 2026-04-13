package types

import (
	"math/big"

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

// Time determination methods (referencing go-ethereum logic).
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
		// If not configured, use default value.
		log.Warn("BlobFeeConfig not set, using default denominator",
			"default", DefaultOsakaBlobConfig)
		return new(big.Int).SetUint64(DefaultOsakaBlobConfig.UpdateFraction)
	}

	cfg := blobFeeConfig
	londonBlock := cfg.LondonBlock // London block number for fork determination.

	// Check in priority order from high to low (BPO5 -> BPO4 -> ... -> Cancun).
	var blobConfig *BlobConfig

	// Check BPO5
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

// isBlockForked returns whether a fork scheduled at block s is active at the
// given head block. Whilst this method is the same as isTimestampForked, they
// are explicitly separate for clearer reading.
func isBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

// isTimestampForked returns whether a fork scheduled at timestamp s is active
// at the given head timestamp. Whilst this method is the same as isBlockForked,
// they are explicitly separate for clearer reading.
func isTimestampForked(s *uint64, head uint64) bool {
	if s == nil {
		return false
	}
	return *s <= head
}
