package rollup

import (
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"

	"morph-l2/morph-deployer/eth"
)

type Genesis struct {
	// The L1 block that the rollup starts *after* (no derived transactions)
	L1 eth.BlockID `json:"l1"`
	// The L2 block the rollup starts from (no transactions, pre-configured state)
	L2 eth.BlockID `json:"l2"`
	// Timestamp of L2 block
	L2Time uint64 `json:"l2_time"`
	// Initial system configuration values.
	// The L2 genesis block may not include transactions, and thus cannot encode the config values,
	// unlike later L2 blocks.
	SystemConfig eth.SystemConfig `json:"system_config"`
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
