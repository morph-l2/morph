package predeploys

import "github.com/scroll-tech/go-ethereum/common"

// legacy todo : rmove
const (
	DevMorphPortal               = "0x6900000000000000000000000000000000000001"
	DevL1CrossDomainMessenger    = "0x6900000000000000000000000000000000000002"
	DevL1StandardBridge          = "0x6900000000000000000000000000000000000003"
	DevMorphMintableERC20Factory = "0x6900000000000000000000000000000000000004"
	DevAddressManager            = "0x6900000000000000000000000000000000000005"
	DevProxyAdmin                = "0x6900000000000000000000000000000000000006"
	DevWETH9                     = "0x6900000000000000000000000000000000000007"
	DevL1ERC721Bridge            = "0x6900000000000000000000000000000000000008"
	DevSystemConfig              = "0x6900000000000000000000000000000000000009"
	DevRollup                    = "0x6900000000000000000000000000000000000010"
	DevStaking                   = "0x6900000000000000000000000000000000000011"
	DevL1Sequencer               = "0x6900000000000000000000000000000000000012"
)

var (
	DevMorphPortalAddr               = common.HexToAddress(DevMorphPortal)
	DevL1CrossDomainMessengerAddr    = common.HexToAddress(DevL1CrossDomainMessenger)
	DevL1StandardBridgeAddr          = common.HexToAddress(DevL1StandardBridge)
	DevMorphMintableERC20FactoryAddr = common.HexToAddress(DevMorphMintableERC20Factory)
	DevAddressManagerAddr            = common.HexToAddress(DevAddressManager)
	DevProxyAdminAddr                = common.HexToAddress(DevProxyAdmin)
	DevWETH9Addr                     = common.HexToAddress(DevWETH9)
	DevL1ERC721BridgeAddr            = common.HexToAddress(DevL1ERC721Bridge)
	DevSystemConfigAddr              = common.HexToAddress(DevSystemConfig)
	DevRollupAddr                    = common.HexToAddress(DevRollup)
	DevStakingAddr                   = common.HexToAddress(DevStaking)
	DevL1SequencerAddr               = common.HexToAddress(DevL1Sequencer)
	DevPredeploys                    = make(map[string]*common.Address)
)

func init() {
	DevPredeploys["MorphPortal"] = &DevMorphPortalAddr
	DevPredeploys["L1CrossDomainMessenger"] = &DevL1CrossDomainMessengerAddr
	DevPredeploys["L1StandardBridge"] = &DevL1StandardBridgeAddr
	DevPredeploys["MorphMintableERC20Factory"] = &DevMorphMintableERC20FactoryAddr
	DevPredeploys["AddressManager"] = &DevAddressManagerAddr
	DevPredeploys["Admin"] = &DevProxyAdminAddr
	DevPredeploys["WETH9"] = &DevWETH9Addr
	DevPredeploys["L1ERC721Bridge"] = &DevL1ERC721BridgeAddr
	DevPredeploys["SystemConfig"] = &DevSystemConfigAddr
	DevPredeploys["Rollup"] = &DevRollupAddr
	DevPredeploys["Staking"] = &DevStakingAddr
	DevPredeploys["L1Sequencer"] = &DevL1SequencerAddr
}
