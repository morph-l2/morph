package genesis

import (
	"errors"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/params"
)

// defaultL2GasLimit represents the default gas limit for an L2 block.
const (
	defaultL2GasLimit = 10_000_000
)

// NewL2Genesis will create a new L2 genesis
func NewL2Genesis(config *DeployConfig, block *types.Block) (*core.Genesis, error) {
	if config.L2ChainID == 0 {
		return nil, errors.New("must define L2 ChainID")
	}

	maxTxPerBlock := config.MaxTxPerBlock
	if maxTxPerBlock == 0 {
		maxTxPerBlock = params.ScrollMaxTxPerBlock
	}

	maxTxPayloadBytesPerBlock := config.MaxTxPayloadBytesPerBlock
	if maxTxPayloadBytesPerBlock == 0 {
		maxTxPayloadBytesPerBlock = params.ScrollMaxTxPayloadBytesPerBlock
	}

	morphChainConfig := params.ChainConfig{
		ChainID:                 new(big.Int).SetUint64(config.L2ChainID),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          false,
		EIP150Block:             big.NewInt(0),
		EIP155Block:             big.NewInt(0),
		EIP158Block:             big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           big.NewInt(0),
		MuirGlacierBlock:        nil,
		BerlinBlock:             big.NewInt(0),
		LondonBlock:             big.NewInt(0),
		ArrowGlacierBlock:       nil,
		ArchimedesBlock:         big.NewInt(0),
		ShanghaiBlock:           big.NewInt(0),
		TerminalTotalDifficulty: big.NewInt(0),
		Scroll: params.ScrollConfig{
			UseZktrie:                 true,
			MaxTxPerBlock:             &maxTxPerBlock,
			MaxTxPayloadBytesPerBlock: &maxTxPayloadBytesPerBlock,
			EnableEIP2718:             false,
			EnableEIP1559:             false,
			FeeVaultAddress:           &config.SequencerFeeVaultRecipient,
		},
	}

	gasLimit := config.L2GenesisBlockGasLimit
	if gasLimit == 0 {
		gasLimit = defaultL2GasLimit
	}
	var baseFee *big.Int
	if config.L2GenesisBlockBaseFeePerGas != nil {
		baseFee = config.L2GenesisBlockBaseFeePerGas.ToInt()
	}

	difficulty := config.L2GenesisBlockDifficulty
	if difficulty == nil {
		difficulty = newHexBig(0)
	}

	timestamp := config.L2GenesisBlockTimestamp
	if timestamp == 0 {
		timestamp = hexutil.Uint64(time.Now().Unix())
	}

	return &core.Genesis{
		Config:     &morphChainConfig,
		Nonce:      uint64(config.L2GenesisBlockNonce),
		Timestamp:  uint64(timestamp),
		ExtraData:  []byte{}, // empty extra data
		GasLimit:   uint64(gasLimit),
		Difficulty: difficulty.ToInt(),
		Mixhash:    config.L2GenesisBlockMixHash,
		Coinbase:   common.Address{},
		Number:     uint64(config.L2GenesisBlockNumber),
		GasUsed:    uint64(config.L2GenesisBlockGasUsed),
		ParentHash: config.L2GenesisBlockParentHash,
		BaseFee:    baseFee,
		Alloc:      map[common.Address]core.GenesisAccount{},
	}, nil
}
