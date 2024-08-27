package genesis

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rpc"

	"morph-l2/bindings/hardhat"
	"morph-l2/bindings/predeploys"
	"morph-l2/morph-deployer/eth"
	"morph-l2/morph-deployer/morph-chain-ops/immutables"
	"morph-l2/morph-deployer/morph-chain-ops/state"
	"morph-l2/morph-deployer/rollup"
)

var (
	ErrInvalidDeployConfig     = errors.New("invalid deploy config")
	ErrInvalidImmutablesConfig = errors.New("invalid immutables config")
)

// DeployConfig represents the deployment configuration for Morph
type DeployConfig struct {
	L1StartingBlockTag *MarshalableRPCBlockNumberOrHash `json:"l1StartingBlockTag"`
	L1ChainID          uint64                           `json:"l1ChainID"`
	L2ChainID          uint64                           `json:"l2ChainID"`

	BatchInboxAddress  common.Address `json:"batchInboxAddress"`
	BatchSenderAddress common.Address `json:"batchSenderAddress"`

	// L2 genesis config
	L2GenesisBlockTimestamp     hexutil.Uint64 `json:"l2GenesisBlockTimestamp"`
	L2GenesisBlockNonce         hexutil.Uint64 `json:"l2GenesisBlockNonce"`
	L2GenesisBlockGasLimit      hexutil.Uint64 `json:"l2GenesisBlockGasLimit"`
	L2GenesisBlockDifficulty    *hexutil.Big   `json:"l2GenesisBlockDifficulty"`
	L2GenesisBlockMixHash       common.Hash    `json:"l2GenesisBlockMixHash"`
	L2GenesisBlockNumber        hexutil.Uint64 `json:"l2GenesisBlockNumber"`
	L2GenesisBlockGasUsed       hexutil.Uint64 `json:"l2GenesisBlockGasUsed"`
	L2GenesisBlockParentHash    common.Hash    `json:"l2GenesisBlockParentHash"`
	L2GenesisBlockBaseFeePerGas *hexutil.Big   `json:"l2GenesisBlockBaseFeePerGas"`
	// Seconds after genesis block that Regolith hard fork activates. 0 to activate at genesis. Nil to disable regolith

	MaxTxPerBlock             int `json:"maxTxPerBlock"`
	MaxTxPayloadBytesPerBlock int `json:"maxTxPayloadBytesPerBlock"`

	// System config
	// Owner of the ProxyAdmin predeploy
	ProxyAdminOwner common.Address `json:"proxyAdminOwner"`
	// Owner of the system on L2
	FinalSystemOwner common.Address `json:"finalSystemOwner"`

	// L1 contract address
	// L1Staking proxy address on L1
	L1StakingProxy common.Address `json:"l1StakingProxy"`
	// L1CrossDomainMessenger proxy address on L1
	L1CrossDomainMessengerProxy common.Address `json:"l1CrossDomainMessengerProxy"`
	// Rollup proxy address on L1
	RollupProxy common.Address `json:"RollupProxy"`
	// L1GatewayRouter proxy address on L1
	L1GatewayRouterProxy common.Address `json:"l1GatewayRouterProxy"`
	// L1StandardERC20Gateway proxy address on L1
	L1StandardERC20GatewayProxy common.Address `json:"l1StandardERC20GatewayProxy"`
	// L1CustomERC20GatewayProxy proxy address on L1
	L1CustomERC20GatewayProxy common.Address `json:"l1CustomERC20GatewayProxy"`
	// L1ReverseCustomGatewayProxy proxy address on L1
	L1ReverseCustomGatewayProxy common.Address `json:"l1ReverseCustomGatewayProxy"`
	// L1ETHGateway proxy address on L1
	L1ETHGatewayProxy common.Address `json:"l1ETHGatewayProxy"`
	// L1ERC721Gateway proxy address on L1
	L1ERC721GatewayProxy common.Address `json:"l1ERC721GatewayProxy"`
	// L1ERC1155Gateway proxy address on L1
	L1ERC1155GatewayProxy common.Address `json:"l1ERC1155GatewayProxy"`
	// L1USDCGatewayProxy proxy address on L1
	L1USDCGatewayProxy common.Address `json:"l1USDCGatewayProxy"`
	// L1WETHGatewayProxy proxy address on L1
	L1WETHGatewayProxy common.Address `json:"l1WETHGatewayProxy"`
	// L1WETH address on L1
	L1WETH common.Address `json:"l1WETH"`
	// L1USDC address on L1
	L1USDC common.Address `json:"l1USDC"`
	// L1WithdrawLockERC20GatewayProxy proxy address on L1
	L1WithdrawLockERC20Gateway common.Address `json:"l1WithdrawLockERC20Gateway"`

	// GasPriceOracle config
	// The initial value of the gas overhead
	GasPriceOracleOverhead uint64 `json:"gasPriceOracleOverhead"`
	// The initial value of the gas scalar
	GasPriceOracleScalar uint64 `json:"gasPriceOracleScalar"`
	// The initial value of the gasOracle owner
	GasPriceOracleOwner common.Address `json:"gasPriceOracleOwner"`

	// Fee recipient address
	// L1 recipient of fees accumulated in the L1FeeVault
	L1FeeVaultRecipient common.Address `json:"l1FeeVaultRecipient"`
	// L1 recipient of fees accumulated in the SequencerFeeVault
	SequencerFeeVaultRecipient common.Address `json:"sequencerFeeVaultRecipient"`
	// L2 recipient of fees accumulated in the Bridge
	L2BridgeFeeVaultRecipient common.Address `json:"l2BridgeFeeVaultRecipient"`

	// Gov configs
	GovVotingDuration     uint64 `json:"govVotingDuration"`
	GovBatchBlockInterval uint64 `json:"govBatchBlockInterval"`
	GovBatchMaxBytes      uint64 `json:"govBatchMaxBytes"`
	GovRollupEpoch        uint64 `json:"govRollupEpoch"`
	GovBatchTimeout       uint64 `json:"govBatchTimeout"`
	GovBatchMaxChunks     uint64 `json:"govBatchMaxChunks"`

	// L2Staking configs
	L2StakingSequencerMaxSize      uint64           `json:"l2StakingSequencerMaxSize"`
	L2StakingUnDelegatedLockEpochs uint64           `json:"l2StakingUnDelegatedLockEpochs"`
	L2StakingRewardStartTime       uint64           `json:"l2StakingRewardStartTime"`
	L2StakingAddresses             []common.Address `json:"l2StakingAddresses"`
	L2StakingTmKeys                []common.Hash    `json:"l2StakingTmKeys"`
	L2StakingBlsKeys               []hexutil.Bytes  `json:"l2StakingBlsKeys"`

	// Record configs
	RecordOracleAddress            common.Address `json:"recordOracleAddress"`
	RecordNextBatchSubmissionIndex uint64         `json:"recordNextBatchSubmissionIndex"`

	// MorphToken configs
	MorphTokenOwner              common.Address `json:"morphTokenOwner"`
	MorphTokenName               string         `json:"morphTokenName"`
	MorphTokenSymbol             string         `json:"morphTokenSymbol"`
	MorphTokenInitialSupply      uint64         `json:"morphTokenInitialSupply"`
	MorphTokenDailyInflationRate uint64         `json:"morphTokenDailyInflationRate"`

	FundDevAccounts bool `json:"fundDevAccounts"`
}

// GetDeployedAddresses will get the deployed addresses of deployed L1 contracts
// required for the L2 genesis creation. Legacy systems use the `Proxy__` prefix
// while modern systems use the `Proxy` suffix. First check for the legacy
// deployments so that this works with upgrading a system.
func (d *DeployConfig) GetDeployedAddresses(hh *hardhat.Hardhat) error {
	if d.L1StakingProxy == (common.Address{}) {
		l1StakingProxyDeployment, err := hh.GetDeployment("Proxy__L1Staking")
		if err != nil {
			return err
		}
		d.L1StakingProxy = l1StakingProxyDeployment.Address
	}

	if d.L1CrossDomainMessengerProxy == (common.Address{}) {
		l1CrossDomainMessengerProxyDeployment, err := hh.GetDeployment("Proxy__L1CrossDomainMessenger")
		if err != nil {
			return err
		}
		d.L1CrossDomainMessengerProxy = l1CrossDomainMessengerProxyDeployment.Address
	}

	if d.RollupProxy == (common.Address{}) {
		RollupProxyDeployment, err := hh.GetDeployment("Proxy__Rollup")
		if err != nil {
			return err
		}
		d.RollupProxy = RollupProxyDeployment.Address
	}

	if d.L1GatewayRouterProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1GatewayRouter")
		if err != nil {
			return err
		}
		d.L1GatewayRouterProxy = deployment.Address
	}

	if d.L1StandardERC20GatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1StandardERC20Gateway")
		if err != nil {
			return err
		}
		d.L1StandardERC20GatewayProxy = deployment.Address
	}

	if d.L1CustomERC20GatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1CustomERC20Gateway")
		if err != nil {
			return err
		}
		d.L1CustomERC20GatewayProxy = deployment.Address
	}

	if d.L1ReverseCustomGatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1ReverseCustomGateway")
		if err != nil {
			return err
		}
		d.L1ReverseCustomGatewayProxy = deployment.Address
	}

	if d.L1ETHGatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1ETHGateway")
		if err != nil {
			return err
		}
		d.L1ETHGatewayProxy = deployment.Address
	}

	if d.L1ERC721GatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1ERC721Gateway")
		if err != nil {
			return err
		}
		d.L1ERC721GatewayProxy = deployment.Address
	}

	if d.L1ERC1155GatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1ERC1155Gateway")
		if err != nil {
			return err
		}
		d.L1ERC1155GatewayProxy = deployment.Address
	}

	if d.L1WETHGatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1WETHGateway")
		if err != nil {
			return err
		}
		d.L1WETHGatewayProxy = deployment.Address
	}

	if d.L1USDCGatewayProxy == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1USDCGateway")
		if err != nil {
			return err
		}
		d.L1USDCGatewayProxy = deployment.Address
	}

	if d.L1WETH == (common.Address{}) {
		deployment, err := hh.GetDeployment("Impl__WETH")
		if err != nil {
			return err
		}
		d.L1WETH = deployment.Address
	}

	if d.L1USDC == (common.Address{}) {
		deployment, err := hh.GetDeployment("Impl__USDC")
		if err != nil {
			return err
		}
		d.L1USDC = deployment.Address
	}

	if d.L1WithdrawLockERC20Gateway == (common.Address{}) {
		deployment, err := hh.GetDeployment("Proxy__L1WithdrawLockERC20Gateway")
		if err != nil {
			return err
		}
		d.L1WithdrawLockERC20Gateway = deployment.Address
	}
	return nil
}

// RollupConfig converts a DeployConfig to a rollup.Config
func (d *DeployConfig) RollupConfig(l1StartBlock *types.Block, l2GenesisBlockHash common.Hash, l2GenesisBlockNumber uint64, l2GenesisStateRoot common.Hash, withdrawRoot common.Hash, genesisBatchHeader []byte) (*rollup.Config, error) {
	//return nil, nil
	return &rollup.Config{
		Genesis: rollup.Genesis{
			L1: eth.BlockID{
				Hash:   l1StartBlock.Hash(),
				Number: l1StartBlock.NumberU64(),
			},
			L2: eth.BlockID{
				Hash:   l2GenesisBlockHash,
				Number: l2GenesisBlockNumber,
			},
			L2Time: l1StartBlock.Time(),
			SystemConfig: eth.SystemConfig{
				BatcherAddr: d.BatchSenderAddress,
				Overhead:    eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(d.GasPriceOracleOverhead))),
				Scalar:      eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(d.GasPriceOracleScalar))),
				GasLimit:    uint64(d.L2GenesisBlockGasLimit),
			},
		},
		L1ChainID:          new(big.Int).SetUint64(d.L1ChainID),
		L2ChainID:          new(big.Int).SetUint64(d.L2ChainID),
		BatchInboxAddress:  d.BatchInboxAddress,
		L2GenesisStateRoot: l2GenesisStateRoot,
		WithdrawRoot:       withdrawRoot,
		GenesisBatchHeader: genesisBatchHeader,
	}, nil
}

// NewDeployConfig reads a config file given a path on the filesystem.
func NewDeployConfig(path string) (*DeployConfig, error) {
	file, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("deploy config at %s not found: %w", path, err)
	}

	var config DeployConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("cannot unmarshal deploy config: %w", err)
	}
	fmt.Printf("owner: %s\n", config.GasPriceOracleOwner.String())
	return &config, nil
}

// NewL2ImmutableConfig will create an ImmutableConfig given an instance of a
// DeployConfig and a block.
func NewL2ImmutableConfig(config *DeployConfig) (immutables.ImmutableConfig, *immutables.InitConfig, error) {
	immutable := make(immutables.ImmutableConfig)

	if config.L1StakingProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1StakingProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1CrossDomainMessengerProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1CrossDomainMessengerProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.RollupProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("RollupProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1GatewayRouterProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1GatewayRouterProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1StandardERC20GatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1StandardERC20GatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1CustomERC20GatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1CustomERC20GatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ReverseCustomGatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1ReverseCustomGatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ETHGatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1ETHGatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ERC721GatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1ERC721GatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ERC1155GatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1ERC1155GatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1WETHGatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1WETHGatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1USDCGatewayProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1USDCGatewayProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1WithdrawLockERC20Gateway == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1WithdrawLockERC20Gateway cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1WETH == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1WETH cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1USDC == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1USDC cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	immutable["L2Staking"] = immutables.ImmutableValues{
		"OTHER_STAKING": config.L1StakingProxy,
	}
	immutable["L2WETHGateway"] = immutables.ImmutableValues{
		"l1WETH": config.L1WETH,
	}
	immutable["L2USDCGateway"] = immutables.ImmutableValues{
		"l1USDC": config.L1USDC,
	}
	blsKeys := make([][]byte, len(config.L2StakingBlsKeys))
	for i, v := range config.L2StakingBlsKeys {
		blsKeys[i] = v
	}

	imConfig := &immutables.InitConfig{
		// L2USDC
		USDCTokenName:     "Bridged USDC",
		USDCTokenSymbol:   "USDC.e",
		USDCTokenCurrency: "USD",
		USDCTokenDecimals: 18,
		USDCMasterMinter:  config.FinalSystemOwner,
		USDCPauser:        config.FinalSystemOwner,
		USDCBlackLister:   config.FinalSystemOwner,
		USDCOwner:         config.FinalSystemOwner,
		// MorphToken
		MorphTokenOwner:              config.MorphTokenOwner,
		MorphTokenName:               config.MorphTokenName,
		MorphTokenSymbol:             config.MorphTokenSymbol,
		MorphTokenInitialSupply:      config.MorphTokenInitialSupply,
		MorphTokenDailyInflationRate: config.MorphTokenDailyInflationRate,
		// L2Staking
		L2StakingOwner:                config.FinalSystemOwner,
		L2StakingSequencersMaxSize:    config.L2StakingSequencerMaxSize,
		L2StakingRewardStartTime:      config.L2StakingRewardStartTime,
		L2StakingUnDelegateLockEpochs: config.L2StakingUnDelegatedLockEpochs,
		L2StakingAddresses:            config.L2StakingAddresses,
		L2StakingTmKeys:               config.L2StakingTmKeys,
		L2StakingBlsKeys:              blsKeys,
	}
	return immutable, imConfig, nil
}

func (d *DeployConfig) Check() error {
	if d.FinalSystemOwner == (common.Address{}) {
		return fmt.Errorf("FinalSystemOwner cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.ProxyAdminOwner == (common.Address{}) {
		return fmt.Errorf("ProxyAdminOwner cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.L1FeeVaultRecipient == (common.Address{}) {
		return fmt.Errorf("L1FeeVaultRecipient cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.GovVotingDuration <= 0 {
		return fmt.Errorf("GovVotingDuration must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GovBatchBlockInterval <= 0 {
		return fmt.Errorf("GovBatchBlockInterval must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GovBatchMaxBytes <= 0 {
		return fmt.Errorf("GovBatchMaxBytes must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GovBatchTimeout <= 0 {
		return fmt.Errorf("GovBatchTimeout must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GovRollupEpoch <= 0 {
		return fmt.Errorf("GovRollupEpoch must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GovBatchMaxChunks <= 0 {
		return fmt.Errorf("GovBatchMaxChunks must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GasPriceOracleOwner == (common.Address{}) {
		return fmt.Errorf("GasPriceOracleOwner cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.GasPriceOracleOverhead <= 0 {
		return fmt.Errorf("GasPriceOracleOverhead must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.GasPriceOracleScalar <= 0 {
		return fmt.Errorf("GasPriceOracleScalar must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.RecordOracleAddress == (common.Address{}) {
		return fmt.Errorf("RecordOracleAddress cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.RecordNextBatchSubmissionIndex <= 0 {
		return fmt.Errorf("RecordNextBatchSubmissionIndex cannot be address(0): %w", ErrInvalidDeployConfig)
	}
	if d.L2StakingSequencerMaxSize <= 0 {
		return fmt.Errorf("L2StakingSequencerMaxSize must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.L2StakingRewardStartTime <= 0 {
		return fmt.Errorf("L2StakingRewardStartTime must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.L2StakingUnDelegatedLockEpochs <= 0 {
		return fmt.Errorf("L2StakingUnDelegatedLockEpochs must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.MorphTokenOwner == (common.Address{}) {
		return fmt.Errorf("MorphTokenOwner canot be nil: %w", ErrInvalidDeployConfig)
	}
	if d.MorphTokenName == "" {
		return fmt.Errorf("MorphTokenName canot be nil: %w", ErrInvalidDeployConfig)
	}
	if d.MorphTokenSymbol == "" {
		return fmt.Errorf("MorphTokenSymbol canot be nil: %w", ErrInvalidDeployConfig)
	}
	if d.MorphTokenDailyInflationRate <= 0 {
		return fmt.Errorf("MorphTokenDailyInflationRate must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	if d.MorphTokenInitialSupply <= 0 {
		return fmt.Errorf("MorphTokenInitialSupply must be greater than 0: %w", ErrInvalidDeployConfig)
	}
	return nil
}

// NewL2StorageConfig will create a StorageConfig given an instance of a
// Hardhat and a DeployConfig.
func NewL2StorageConfig(config *DeployConfig, baseFee *big.Int) (state.StorageConfig, error) {
	l2BridgeFeeVaultRecipient := config.L2BridgeFeeVaultRecipient
	if l2BridgeFeeVaultRecipient == types.EmptyAddress {
		l2BridgeFeeVaultRecipient = predeploys.L2TxFeeVaultAddr
	}
	storage := make(state.StorageConfig)
	err := config.Check()
	if err != nil {
		return nil, err
	}
	storage["GasPriceOracle"] = state.StorageValues{
		"_owner":           config.GasPriceOracleOwner,
		"overhead":         config.GasPriceOracleOverhead,
		"scalar":           config.GasPriceOracleScalar,
		"l1BaseFee":        baseFee,
		"allowListEnabled": true,
	}
	storage["L2CrossDomainMessenger"] = state.StorageValues{
		"_status":              1,
		"_initialized":         1,
		"_initializing":        false,
		"_owner":               config.FinalSystemOwner,
		"_paused":              false,
		"xDomainMessageSender": "0x000000000000000000000000000000000000dEaD",
		"counterpart":          config.L1CrossDomainMessengerProxy,
		"feeVault":             l2BridgeFeeVaultRecipient,
	}
	storage["MorphToken"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
	}
	storage["Record"] = state.StorageValues{
		"_initialized":             1,
		"_initializing":            false,
		"_owner":                   config.FinalSystemOwner,
		"oracle":                   config.RecordOracleAddress,
		"nextBatchSubmissionIndex": config.RecordNextBatchSubmissionIndex,
	}
	storage["Distribute"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
	}
	storage["L2Staking"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
	}
	storage["Sequencer"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
	}
	storage["Gov"] = state.StorageValues{
		"_initialized":       1,
		"_initializing":      false,
		"_owner":             config.FinalSystemOwner,
		"votingDuration":     config.GovVotingDuration,
		"batchBlockInterval": config.GovBatchBlockInterval,
		"batchMaxBytes":      config.GovBatchMaxBytes,
		"batchTimeout":       config.GovBatchTimeout,
		"maxChunks":          config.GovBatchMaxChunks,
		"rollupEpoch":        config.GovRollupEpoch,
	}
	storage["L2ToL1MessagePasser"] = state.StorageValues{
		"messageRoot": common.HexToHash("0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757"),
	}
	storage["L2TxFeeVault"] = state.StorageValues{
		"owner":             config.FinalSystemOwner,
		"minWithdrawAmount": 0,
		"recipient":         config.L1FeeVaultRecipient,
	}
	storage["ProxyAdmin"] = state.StorageValues{
		"_owner": config.ProxyAdminOwner,
	}
	storage["L2GatewayRouter"] = state.StorageValues{
		"_initialized":        1,
		"_initializing":       false,
		"_owner":              config.ProxyAdminOwner,
		"ethGateway":          predeploys.L2ETHGatewayAddr,
		"defaultERC20Gateway": predeploys.L2StandardERC20GatewayAddr,
	}
	storage["L2StandardERC20Gateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"tokenFactory":  predeploys.MorphStandardERC20FactoryAddr,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1StandardERC20GatewayProxy,
	}
	storage["L2CustomERC20Gateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1CustomERC20GatewayProxy,
	}
	storage["L2ReverseCustomGateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1ReverseCustomGatewayProxy,
	}
	storage["L2ETHGateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1ETHGatewayProxy,
	}
	storage["L2WETHGateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"counterpart":   config.L1WETHGatewayProxy,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
	}
	storage["L2USDCGateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"counterpart":   config.L1USDCGatewayProxy,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
	}
	storage["L2WithdrawLockERC20Gateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"router":        predeploys.L2GatewayRouterAddr,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1WithdrawLockERC20Gateway,
	}
	storage["L2ERC721Gateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1ERC721GatewayProxy,
		"router":        common.BigToAddress(common.Big0),
	}
	storage["L2ERC1155Gateway"] = state.StorageValues{
		"_status":       1, // ReentrancyGuard
		"_initialized":  1,
		"_initializing": false,
		"_owner":        config.FinalSystemOwner,
		"messenger":     predeploys.L2CrossDomainMessengerAddr,
		"counterpart":   config.L1ERC1155GatewayProxy,
		"router":        common.BigToAddress(common.Big0),
	}
	storage["MorphStandardERC20Factory"] = state.StorageValues{
		"_owner":         predeploys.L2StandardERC20GatewayAddr,
		"implementation": predeploys.MorphStandardERC20Addr,
	}
	storage["L2USDC"] = state.StorageValues{}
	return storage, nil
}

type MarshalableRPCBlockNumberOrHash rpc.BlockNumberOrHash

func (m *MarshalableRPCBlockNumberOrHash) MarshalJSON() ([]byte, error) {
	r := rpc.BlockNumberOrHash(*m)
	if hash, ok := r.Hash(); ok {
		return json.Marshal(hash)
	}
	if num, ok := r.Number(); ok {
		// never errors
		text, _ := num.MarshalText()
		return json.Marshal(string(text))
	}
	return json.Marshal(nil)
}

func (m *MarshalableRPCBlockNumberOrHash) UnmarshalJSON(b []byte) error {
	var r rpc.BlockNumberOrHash
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	asMarshalable := MarshalableRPCBlockNumberOrHash(r)
	*m = asMarshalable
	return nil
}

// Number wraps the rpc.BlockNumberOrHash Number method.
func (m *MarshalableRPCBlockNumberOrHash) Number() (rpc.BlockNumber, bool) {
	return (*rpc.BlockNumberOrHash)(m).Number()
}

// Hash wraps the rpc.BlockNumberOrHash Hash method.
func (m *MarshalableRPCBlockNumberOrHash) Hash() (common.Hash, bool) {
	return (*rpc.BlockNumberOrHash)(m).Hash()
}

// String wraps the rpc.BlockNumberOrHash String method.
func (m *MarshalableRPCBlockNumberOrHash) String() string {
	return (*rpc.BlockNumberOrHash)(m).String()
}
