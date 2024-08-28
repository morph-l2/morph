package genesis

import (
	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/bindings/predeploys"
)

const (
	// MaxPredeploySlotChecks is the maximum number of storage slots to check
	// when validating the untouched predeploys. This limit is in place
	// to bound execution time of the migration. We can parallelize this
	// in the future.
	MaxPredeploySlotChecks = 1000

	// MaxOVMETHSlotChecks is the maximum number of OVM ETH storage slots to check
	// when validating the OVM ETH migration.
	MaxOVMETHSlotChecks = 5000

	// OVMETHSampleLikelihood is the probability that a storage slot will be checked
	// when validating the OVM ETH migration.
	OVMETHSampleLikelihood = 0.1
)

type StorageCheckMap = map[common.Hash]common.Hash

var (
	L2XDMOwnerSlot      = common.Hash{31: 0x33}
	ProxyAdminOwnerSlot = common.Hash{}

	LegacyETHCheckSlots = map[common.Hash]common.Hash{
		// Bridge
		{31: 0x06}: common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
		// Symbol
		{31: 0x04}: common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000006"),
		// Name
		{31: 0x03}: common.HexToHash("0x457468657200000000000000000000000000000000000000000000000000000a"),
		// Total supply
		{31: 0x02}: {},
	}

	// ExpectedStorageSlots is a map of predeploy addresses to the storage slots and values that are
	// expected to be set in those predeploys after the migration. It does not include any predeploys
	// that were not wiped. It also accounts for the 2 EIP-1967 storage slots in each contract.
	// It does _not_ include L1Block. L1Block is checked separately.
	ExpectedStorageSlots = map[common.Address]StorageCheckMap{
		predeploys.L2CrossDomainMessengerAddr: {
			// Slot 0x00 (0) is a combination of spacer_0_0_20, _initialized, and _initializing
			common.Hash{}: common.HexToHash("0x0000000000000000000000010000000000000000000000000000000000000000"),
			// Slot 0xcc (204) is xDomainMsgSender
			common.Hash{31: 0xcc}: common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000dead"),
			// EIP-1967 storage slots
			AdminSlot:          common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000018"),
			ImplementationSlot: common.HexToHash("0x000000000000000000000000c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d30007"),
		},
		// ProxyAdmin is not a proxy, and only has the _owner slot set.
		predeploys.ProxyAdminAddr: {
			// Slot 0x00 (0) is _owner. Requires custom check, so set to a garbage value
			ProxyAdminOwnerSlot: common.HexToHash("0xbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbad0"),

			// EIP-1967 storage slots
			AdminSlot:          common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000018"),
			ImplementationSlot: common.HexToHash("0x000000000000000000000000c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d30018"),
		},
	}
)

// func eip1967Slots(address common.Address) StorageCheckMap {
// 	codeAddr, err := AddressToCodeNamespace(address)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return StorageCheckMap{
// 		AdminSlot:          predeploys.ProxyAdminAddr.Hash(),
// 		ImplementationSlot: codeAddr.Hash(),
// 	}
// }
