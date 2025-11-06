package types

import (
	"strconv"
	"testing"

	"github.com/morph-l2/go-ethereum/consensus/misc/eip4844"
	"github.com/stretchr/testify/assert"
)

func TestGetBlobFeeDenominator_MainnetChainConfig(t *testing.T) {
	// Use MainnetChainConfig directly since map lookup with *big.Int keys requires exact pointer match
	config := MainnetChainConfig
	assert.NotNil(t, config, "MainnetChainConfig should exist")

	tests := []struct {
		name           string
		blockTime      uint64
		expectedResult uint64
		forkName       string
	}{
		{
			name:           "Before Cancun",
			blockTime:      1710338134,                            // Before Cancun (1710338135)
			expectedResult: DefaultOsakaBlobConfig.UpdateFraction, // Should use Default
			forkName:       "Default",
		},
		{
			name:           "At Cancun",
			blockTime:      1710338135, // Cancun fork time
			expectedResult: DefaultCancunBlobConfig.UpdateFraction,
			forkName:       "Cancun",
		},
		{
			name:           "After Cancun, Before Prague",
			blockTime:      1740000000, // Between Cancun and Prague
			expectedResult: DefaultCancunBlobConfig.UpdateFraction,
			forkName:       "Cancun",
		},
		{
			name:           "At Prague",
			blockTime:      1746612311, // Prague fork time
			expectedResult: DefaultPragueBlobConfig.UpdateFraction,
			forkName:       "Prague",
		},
		{
			name:           "After Prague",
			blockTime:      1746612312, // After Prague
			expectedResult: DefaultPragueBlobConfig.UpdateFraction,
			forkName:       "Prague",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetBlobFeeDenominator(config, tt.blockTime)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedResult, result.Uint64(), "Expected %s fork UpdateFraction", tt.forkName)
		})
	}
}

func TestGetBlobFeeDenominator_HoodiChainConfig(t *testing.T) {
	// Use HoodiChainConfig directly since map lookup with *big.Int keys requires exact pointer match
	config := HoodiChainConfig
	assert.NotNil(t, config, "HoodiChainConfig should exist")

	tests := []struct {
		name           string
		blockTime      uint64
		expectedResult uint64
		forkName       string
	}{
		{
			name:           "At Cancun (0)",
			blockTime:      0, // Cancun fork time is 0
			expectedResult: DefaultCancunBlobConfig.UpdateFraction,
			forkName:       "Cancun",
		},
		{
			name:           "Before Prague",
			blockTime:      1742999831, // Before Prague (1742999832)
			expectedResult: DefaultCancunBlobConfig.UpdateFraction,
			forkName:       "Cancun",
		},
		{
			name:           "At Prague",
			blockTime:      1742999832, // Prague fork time
			expectedResult: DefaultPragueBlobConfig.UpdateFraction,
			forkName:       "Prague",
		},
		{
			name:           "Before Osaka",
			blockTime:      1761677591, // Before Osaka (1761677592)
			expectedResult: DefaultPragueBlobConfig.UpdateFraction,
			forkName:       "Prague",
		},
		{
			name:           "At Osaka",
			blockTime:      1761677592, // Osaka fork time
			expectedResult: DefaultOsakaBlobConfig.UpdateFraction,
			forkName:       "Osaka",
		},
		{
			name:           "Before BPO1",
			blockTime:      1762365719, // Before BPO1 (1762365720)
			expectedResult: DefaultOsakaBlobConfig.UpdateFraction,
			forkName:       "Osaka",
		},
		{
			name:           "At BPO1",
			blockTime:      1762365720, // BPO1 fork time
			expectedResult: DefaultBPO1BlobConfig.UpdateFraction,
			forkName:       "BPO1",
		},
		{
			name:           "Before BPO2",
			blockTime:      1762955543, // Before BPO2 (1762955544)
			expectedResult: DefaultBPO1BlobConfig.UpdateFraction,
			forkName:       "BPO1",
		},
		{
			name:           "At BPO2",
			blockTime:      1762955544, // BPO2 fork time
			expectedResult: DefaultBPO2BlobConfig.UpdateFraction,
			forkName:       "BPO2",
		},
		{
			name:           "After BPO2",
			blockTime:      1762955545, // After BPO2
			expectedResult: DefaultBPO2BlobConfig.UpdateFraction,
			forkName:       "BPO2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetBlobFeeDenominator(config, tt.blockTime)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedResult, result.Uint64(), "Expected %s fork UpdateFraction", tt.forkName)
		})
	}
}

func TestGetBlobFeeDenominator_NilConfig(t *testing.T) {
	t.Run("Nil config uses default", func(t *testing.T) {
		result := GetBlobFeeDenominator(nil, 1000000000)
		assert.NotNil(t, result)
		assert.Equal(t, DefaultOsakaBlobConfig.UpdateFraction, result.Uint64())
	})
}

func TestGetBlobFeeDenominator_AllChainConfigs(t *testing.T) {
	// Test that all chain configs in ChainConfigMap work correctly
	for chainID, config := range ChainConfigMap {
		t.Run("ChainID_"+strconv.FormatUint(chainID, 10), func(t *testing.T) {
			assert.NotNil(t, config, "Config should not be nil for chainID %s", chainID)
			assert.NotNil(t, config.LondonBlock, "LondonBlock should not be nil for chainID %s", chainID)

			// Test with a very large timestamp to ensure it works
			result := GetBlobFeeDenominator(config, 9999999999)
			assert.NotNil(t, result)
			assert.Greater(t, result.Uint64(), uint64(0), "UpdateFraction should be greater than 0")
		})
	}
}

func TestGetBlobFeeDenominator_ChainConfigMap(t *testing.T) {
	// Test GetBlobFeeDenominator using ChainConfigMap
	t.Run("Mainnet from ChainConfigMap", func(t *testing.T) {
		// Find Mainnet config by ChainID
		var mainnetConfig *BlobFeeConfig
		for chainID, config := range ChainConfigMap {
			if chainID == 1 {
				mainnetConfig = config
				break
			}
		}
		assert.NotNil(t, mainnetConfig, "Mainnet config should be found in ChainConfigMap")

		// Test various timestamps
		result := GetBlobFeeDenominator(mainnetConfig, 1710338135) // At Cancun
		assert.Equal(t, DefaultCancunBlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(mainnetConfig, 1746612311) // At Prague
		assert.Equal(t, DefaultPragueBlobConfig.UpdateFraction, result.Uint64())
	})

	t.Run("Hoodi from ChainConfigMap", func(t *testing.T) {
		// Find Hoodi config by ChainID
		var hoodiConfig *BlobFeeConfig
		for chainID, config := range ChainConfigMap {
			if chainID == 560048 {
				hoodiConfig = config
				break
			}
		}
		assert.NotNil(t, hoodiConfig, "Hoodi config should be found in ChainConfigMap")

		// Test various timestamps
		result := GetBlobFeeDenominator(hoodiConfig, 0) // At Cancun (0)
		assert.Equal(t, DefaultCancunBlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(hoodiConfig, 1761677592) // At Osaka
		assert.Equal(t, DefaultOsakaBlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(hoodiConfig, 1762365720) // At BPO1
		assert.Equal(t, DefaultBPO1BlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(hoodiConfig, 1762955544) // At BPO2
		assert.Equal(t, DefaultBPO2BlobConfig.UpdateFraction, result.Uint64())
	})

	t.Run("Devnet from ChainConfigMap", func(t *testing.T) {
		// Find Devnet config by ChainID
		var devnetConfig *BlobFeeConfig
		for chainID, config := range ChainConfigMap {
			if chainID == 900 {
				devnetConfig = config
				break
			}
		}
		assert.NotNil(t, devnetConfig, "Devnet config should be found in ChainConfigMap")

		// Test various timestamps
		result := GetBlobFeeDenominator(devnetConfig, 0) // At Cancun (0)
		assert.Equal(t, DefaultCancunBlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(devnetConfig, 1761677592) // At Osaka
		assert.Equal(t, DefaultOsakaBlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(devnetConfig, 1762365720) // At BPO1
		assert.Equal(t, DefaultBPO1BlobConfig.UpdateFraction, result.Uint64())

		result = GetBlobFeeDenominator(devnetConfig, 1762955544) // At BPO2
		assert.Equal(t, DefaultBPO2BlobConfig.UpdateFraction, result.Uint64())
	})
}

func TestGetBlobFeeDenominator_ForkPriority(t *testing.T) {
	// Test that higher priority forks are checked first
	config := HoodiChainConfig
	assert.NotNil(t, config)

	// At BPO2 time, should return BPO2, not BPO1, Osaka, Prague, or Cancun
	bpo2Time := uint64(1762955544)
	result := GetBlobFeeDenominator(config, bpo2Time)
	assert.Equal(t, DefaultBPO2BlobConfig.UpdateFraction, result.Uint64())

	// At BPO1 time, should return BPO1, not Osaka, Prague, or Cancun
	bpo1Time := uint64(1762365720)
	result = GetBlobFeeDenominator(config, bpo1Time)
	assert.Equal(t, DefaultBPO1BlobConfig.UpdateFraction, result.Uint64())

	// At Osaka time, should return Osaka, not Prague or Cancun
	osakaTime := uint64(1761677592)
	result = GetBlobFeeDenominator(config, osakaTime)
	assert.Equal(t, DefaultOsakaBlobConfig.UpdateFraction, result.Uint64())
}

func TestCal(t *testing.T) {
	// Test that higher priority forks are checked first
	config := HoodiChainConfig
	assert.NotNil(t, config)

	timeNow := uint64(1762395060)
	result := GetBlobFeeDenominator(config, timeNow)
	t.Log(result)
	fee := eip4844.CalcBlobFee(170172092, result.Uint64())
	t.Log(fee.Uint64())
}
