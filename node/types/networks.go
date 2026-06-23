package types

import "github.com/morph-l2/go-ethereum/common"

// Default L1 contract addresses and derivation start heights per network.
// Defined in this leaf package so sync/, derivation/ and cmd/ can all consume
// them without creating an import cycle through node/core.

var (
	// L1 Mainnet Contract Addresses
	MainnetRollupContractAddress      = common.HexToAddress("0x759894ced0e6af42c26668076ffa84d02e3cef60")
	MainnetSyncDepositContractAddress = common.HexToAddress("0x3931ade842f5bb8763164bdd81e5361dce6cc1ef")
	MainnetL1SequencerContractAddress = common.HexToAddress("")

	// L1 Hoodi Contract Addresses
	HoodiRollupContractAddress      = common.HexToAddress("0x57e0e6dde89dc52c01fe785774271504b1e04664")
	HoodiSyncDepositContractAddress = common.HexToAddress("0xd7f39d837f4790b215ba67e0ab63665912648dbe")
	HoodiL1SequencerContractAddress = common.HexToAddress("")
)

// Network-specific default block timestamps (Unix milliseconds) at which consensus switches to
// centralized sequencer mode. They are used when --mainnet / --hoodi is selected and the generic
// sequencerUpgradeTime flag is unset. A value <= 0 disables timestamp-triggered upgrade.
var (
	MainnetCentralizedSequencerUpgradeTime int64 = 0 // ms
	HoodiCentralizedSequencerUpgradeTime   int64 = 0 // ms
)
