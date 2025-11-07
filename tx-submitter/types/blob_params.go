package types

import (
	"math/big"
)

var (
	DefaultBlobConfig = HoodiChainConfig

	ChainConfigMap = ChainBlobConfigs{
		1:      MainnetChainConfig,
		560048: HoodiChainConfig,
		900:    DevnetChainConfig,
	}
)

func newUint64(val uint64) *uint64 { return &val }

type ChainBlobConfigs map[uint64]*BlobFeeConfig

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
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

	// HoodiChainConfig contains the chain parameters to run a node on the Hoodi test network.
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

	// DevnetChainConfig contains the chain parameters to run a node on the devnet test network.
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

var (
	// DefaultCancunBlobConfig is the default blob configuration for the Cancun fork.
	DefaultCancunBlobConfig = &BlobConfig{
		UpdateFraction: 3338477,
	}
	// DefaultPragueBlobConfig is the default blob configuration for the Prague fork.
	DefaultPragueBlobConfig = &BlobConfig{
		UpdateFraction: 5007716,
	}
	// DefaultOsakaBlobConfig is the default blob configuration for the Osaka fork.
	DefaultOsakaBlobConfig = &BlobConfig{
		UpdateFraction: 5007716,
	}
	// DefaultBPO1BlobConfig is the default blob configuration for the BPO1 fork.
	DefaultBPO1BlobConfig = &BlobConfig{
		UpdateFraction: 8346193,
	}
	// DefaultBPO2BlobConfig is the default blob configuration for the BPO2 fork.
	DefaultBPO2BlobConfig = &BlobConfig{
		UpdateFraction: 11684671,
	}
	// DefaultBPO3BlobConfig is the default blob configuration for the BPO3 fork.
	DefaultBPO3BlobConfig = &BlobConfig{
		UpdateFraction: 20609697,
	}
	// DefaultBPO4BlobConfig is the default blob configuration for the BPO4 fork.
	DefaultBPO4BlobConfig = &BlobConfig{
		UpdateFraction: 13739630,
	}
)
