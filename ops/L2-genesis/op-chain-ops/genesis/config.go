package genesis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/morph-l2/morph-deployer/eth"
	"github.com/morph-l2/morph-deployer/rollup"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/bindings/hardhat"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/morph-l2/morph-deployer/op-chain-ops/immutables"
	"github.com/morph-l2/morph-deployer/op-chain-ops/state"
	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/rpc"
)

var (
	ErrInvalidDeployConfig     = errors.New("invalid deploy config")
	ErrInvalidImmutablesConfig = errors.New("invalid immutables config")
)

type Genesis struct {
	// The L1 block that the rollup starts *after* (no derived transactions)
	L1 uint64 `json:"l1"`
	// The L2 block the rollup starts from (no transactions, pre-configured state)
	L2 uint64 `json:"l2"`
	// Timestamp of L2 block
	L2Time uint64 `json:"l2_time"`
}

type Config struct {
	// Genesis anchor point of the rollup
	Genesis Genesis `json:"genesis"`
	// Seconds per L2 block
	BlockTime uint64 `json:"block_time"`
	// Required to verify L1 signatures
	L1ChainID *big.Int `json:"l1_chain_id"`
	// Required to identify the L2 network and create p2p signatures unique for this chain.
	L2ChainID *big.Int `json:"l2_chain_id"`

	// Note: below addresses are part of the block-derivation process,
	// and required to be the same network-wide to stay in consensus.

	// L1 address that batches are sent to.
	BatchInboxAddress common.Address `json:"batch_inbox_address"`
	// L1 Deposit Contract Address
	DepositContractAddress common.Address `json:"deposit_contract_address"`
	// L1 System Config Address
	L1SystemConfigAddress common.Address `json:"l1_system_config_address"`

	L2GenesisStateRoot common.Hash `json:"l2_genesis_state_root"`

	WithdrawRoot common.Hash `json:"withdraw_root"`

	GenesisBatchHeader hexutil.Bytes `json:"genesis_batch_header"`
}

// DeployConfig represents the deployment configuration for Morph
type DeployConfig struct {
	L1StartingBlockTag *MarshalableRPCBlockNumberOrHash `json:"l1StartingBlockTag"`
	L1ChainID          uint64                           `json:"l1ChainID"`
	L2ChainID          uint64                           `json:"l2ChainID"`

	FinalizationPeriodSeconds uint64         `json:"finalizationPeriodSeconds"`
	BatchInboxAddress         common.Address `json:"batchInboxAddress"`
	BatchSenderAddress        common.Address `json:"batchSenderAddress"`

	L1BlockTime                 uint64         `json:"l1BlockTime"`
	L1GenesisBlockTimestamp     hexutil.Uint64 `json:"l1GenesisBlockTimestamp"`
	L1GenesisBlockNonce         hexutil.Uint64 `json:"l1GenesisBlockNonce"`
	CliqueSignerAddress         common.Address `json:"cliqueSignerAddress"` // proof of stake genesis if left zeroed.
	L1GenesisBlockGasLimit      hexutil.Uint64 `json:"l1GenesisBlockGasLimit"`
	L1GenesisBlockDifficulty    *hexutil.Big   `json:"l1GenesisBlockDifficulty"`
	L1GenesisBlockMixHash       common.Hash    `json:"l1GenesisBlockMixHash"`
	L1GenesisBlockCoinbase      common.Address `json:"l1GenesisBlockCoinbase"`
	L1GenesisBlockNumber        hexutil.Uint64 `json:"l1GenesisBlockNumber"`
	L1GenesisBlockGasUsed       hexutil.Uint64 `json:"l1GenesisBlockGasUsed"`
	L1GenesisBlockParentHash    common.Hash    `json:"l1GenesisBlockParentHash"`
	L1GenesisBlockBaseFeePerGas *hexutil.Big   `json:"l1GenesisBlockBaseFeePerGas"`

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
	L2GenesisRegolithTimeOffset *hexutil.Uint64 `json:"l2GenesisRegolithTimeOffset,omitempty"`

	// Owner of the ProxyAdmin predeploy
	ProxyAdminOwner common.Address `json:"proxyAdminOwner"`
	// Owner of the system on L1
	FinalSystemOwner common.Address `json:"finalSystemOwner"`
	// L1SequencerProxy proxy address on L1
	L1SequencerProxy common.Address `json:"l1SequencerProxy"`
	// L1CrossDomainMessenger proxy address on L1
	L1CrossDomainMessengerProxy common.Address `json:"l1CrossDomainMessengerProxy"`
	// Rollup proxy address on L1
	RollupProxy common.Address `json:"RollupProxy"`
	// The initial value of the gas overhead
	GasPriceOracleOverhead uint64 `json:"gasPriceOracleOverhead"`
	// The initial value of the gas scalar
	GasPriceOracleScalar uint64 `json:"gasPriceOracleScalar"`
	// The initial value of the gasOracle owner
	GasPriceOracleOwner common.Address `json:"gasPriceOracleOwner"`
	// L1 recipient of fees accumulated in the L1FeeVault
	L1FeeVaultRecipient common.Address `json:"l1FeeVaultRecipient"`
	// L1 recipient of fees accumulated in the SequencerFeeVault
	SequencerFeeVaultRecipient common.Address `json:"sequencerFeeVaultRecipient"`

	L2BridgeFeeVaultRecipient common.Address `json:"l2BridgeFeeVaultRecipient"`

	GovProposalInterval   uint64 `json:"govProposalInterval"`
	GovBatchBlockInterval uint64 `json:"govBatchBlockInterval"`
	GovBatchMaxBytes      uint64 `json:"govBatchMaxBytes"`
	GovRollupEpoch        uint64 `json:"govRollupEpoch"`
	GovBatchTimeout       uint64 `json:"govBatchTimeout"`
	GovBatchMaxChunks     uint64 `json:"govBatchMaxChunks"`

	L2SequencerAddresses []common.Address `json:"l2SequencerAddresses"`
	L2SequencerTmKeys    []common.Hash    `json:"l2SequencerTmKeys"`
	L2SequencerBlsKeys   []hexutil.Bytes  `json:"l2SequencerBlsKeys"`

	DeploymentWaitConfirmations int `json:"deploymentWaitConfirmations"`

	EIP1559Elasticity  uint64 `json:"eip1559Elasticity"`
	EIP1559Denominator uint64 `json:"eip1559Denominator"`

	FundDevAccounts bool `json:"fundDevAccounts"`

	MaxTxPerBlock             int `json:"maxTxPerBlock"`
	MaxTxPayloadBytesPerBlock int `json:"maxTxPayloadBytesPerBlock"`
}

// GetDeployedAddresses will get the deployed addresses of deployed L1 contracts
// required for the L2 genesis creation. Legacy systems use the `Proxy__` prefix
// while modern systems use the `Proxy` suffix. First check for the legacy
// deployments so that this works with upgrading a system.
func (d *DeployConfig) GetDeployedAddresses(hh *hardhat.Hardhat) error {
	var err error

	if d.L1SequencerProxy == (common.Address{}) {
		var l1SequencerProxyDeployment *hardhat.Deployment
		l1SequencerProxyDeployment, err = hh.GetDeployment("Proxy__L1Sequencer")
		if errors.Is(err, hardhat.ErrCannotFindDeployment) {
			l1SequencerProxyDeployment, err = hh.GetDeployment("L1SequencerProxy")
			if err != nil {
				return err
			}
		}
		d.L1SequencerProxy = l1SequencerProxyDeployment.Address
	}

	if d.L1CrossDomainMessengerProxy == (common.Address{}) {
		var l1CrossDomainMessengerProxyDeployment *hardhat.Deployment
		l1CrossDomainMessengerProxyDeployment, err = hh.GetDeployment("Proxy__L1CrossDomainMessenger")
		if errors.Is(err, hardhat.ErrCannotFindDeployment) {
			l1CrossDomainMessengerProxyDeployment, err = hh.GetDeployment("L1CrossDomainMessengerProxy")
			if err != nil {
				return err
			}
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
	return nil
}

// InitDeveloperDeployedAddresses will set the dev addresses on the DeployConfig
func (d *DeployConfig) InitDeveloperDeployedAddresses() error {
	d.L1SequencerProxy = predeploys.DevL1SequencerAddr
	d.L1CrossDomainMessengerProxy = predeploys.DevL1CrossDomainMessengerAddr
	return nil
}

func (d *DeployConfig) RegolithTime(genesisTime uint64) *uint64 {
	if d.L2GenesisRegolithTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisRegolithTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
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
	file, err := os.ReadFile(path)
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
func NewL2ImmutableConfig(config *DeployConfig) (immutables.ImmutableConfig, *immutables.Config, error) {
	immutable := make(immutables.ImmutableConfig)

	if config.L1SequencerProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1SequencerProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1CrossDomainMessengerProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("L1CrossDomainMessengerProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.RollupProxy == (common.Address{}) {
		return immutable, nil, fmt.Errorf("RollupProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}

	immutable["GasPriceOracleOwner"] = immutables.ImmutableValues{
		"owner": config.GasPriceOracleOwner,
	}
	immutable["L2Sequencer"] = immutables.ImmutableValues{
		"otherSequencer": config.L1SequencerProxy,
	}
	immutable["L2CrossDomainMessenger"] = immutables.ImmutableValues{
		"otherMessenger": config.L1CrossDomainMessengerProxy,
	}
	immutable["L2ToL1MessagePasser"] = immutables.ImmutableValues{
		"recAddr": config.L1CrossDomainMessengerProxy,
	}
	immutable["L2TxFeeVault"] = immutables.ImmutableValues{
		"owner":               config.FinalSystemOwner,
		"recipient":           config.L1FeeVaultRecipient,
		"minWithdrawalAmount": uint64(1),
	}
	immutable["Submitter"] = immutables.ImmutableValues{
		"rollupAddr": config.RollupProxy,
	}

	blsKeys := make([][]byte, len(config.L2SequencerBlsKeys))
	for i, v := range config.L2SequencerBlsKeys {
		blsKeys[i] = v
	}
	imConfig := &immutables.Config{
		L2SequencerAddresses: config.L2SequencerAddresses,
		L2SequencerTmKeys:    config.L2SequencerTmKeys,
		L2SequencerBlsKeys:   blsKeys,
	}
	return immutable, imConfig, nil
}

// NewL2StorageConfig will create a StorageConfig given an instance of a
// Hardhat and a DeployConfig.
func NewL2StorageConfig(config *DeployConfig, baseFee *big.Int) (state.StorageConfig, error) {
	storage := make(state.StorageConfig)

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
		"feeVault":             config.L2BridgeFeeVaultRecipient,
	}
	storage["L2Sequencer"] = state.StorageValues{
		"_initialized":   1,
		"_initializing":  false,
		"currentVersion": 0,
	}
	storage["Gov"] = state.StorageValues{
		"_initialized":       1,
		"_initializing":      false,
		"proposalInterval":   config.GovProposalInterval,
		"batchBlockInterval": config.GovBatchBlockInterval,
		"batchMaxBytes":      config.GovBatchMaxBytes,
		"batchTimeout":       config.GovBatchTimeout,
		"rollupEpoch":        config.GovRollupEpoch,
		"maxChunks":          config.GovBatchMaxChunks,
	}
	storage["Submitter"] = state.StorageValues{
		"_initialized":   1,
		"_initializing":  false,
		"nextEpochStart": uint64(time.Now().Unix()),
	}
	storage["L2ToL1MessagePasser"] = state.StorageValues{
		"messageRoot": common.HexToHash("0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757"),
	}
	storage["L2TxFeeVault"] = state.StorageValues{
		"owner":             config.FinalSystemOwner,
		"minWithdrawAmount": 0,
		"recipient":         config.SequencerFeeVaultRecipient,
	}
	storage["ProxyAdmin"] = state.StorageValues{
		"_owner": config.ProxyAdminOwner,
	}
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
