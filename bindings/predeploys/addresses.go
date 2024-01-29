package predeploys

import "github.com/scroll-tech/go-ethereum/common"

const (
	L2ToL1MessagePasser    = "0x5300000000000000000000000000000000000001"
	L2GatewayRouter        = "0x5300000000000000000000000000000000000002"
	L2Sequencer            = "0x5300000000000000000000000000000000000003"
	L2Gov                  = "0x5300000000000000000000000000000000000004"
	L2Submitter            = "0x5300000000000000000000000000000000000005"
	L2ETHGateway           = "0x5300000000000000000000000000000000000006"
	L2CrossDomainMessenger = "0x5300000000000000000000000000000000000007"
	L2StandardERC20Gateway = "0x5300000000000000000000000000000000000008"
	L2ERC721Gateway        = "0x5300000000000000000000000000000000000009"
	L2TxFeeVault           = "0x530000000000000000000000000000000000001A"
	ProxyAdmin             = "0x530000000000000000000000000000000000001B"
	GasPriceOracle         = "0x530000000000000000000000000000000000000F"
)

var (
	L2ToL1MessagePasserAddr    = common.HexToAddress(L2ToL1MessagePasser)
	L2CrossDomainMessengerAddr = common.HexToAddress(L2CrossDomainMessenger)
	GasPriceOracleAddr         = common.HexToAddress(GasPriceOracle)
	ProxyAdminAddr             = common.HexToAddress(ProxyAdmin)
	L2SequencerAddr            = common.HexToAddress(L2Sequencer)
	GovAddr                    = common.HexToAddress(L2Gov)
	SubmitterAddr              = common.HexToAddress(L2Submitter)
	L2GatewayRouterAddr        = common.HexToAddress(L2GatewayRouter)
	L2ETHGatewayAddr           = common.HexToAddress(L2ETHGateway)
	L2StandardERC20GatewayAddr = common.HexToAddress(L2StandardERC20Gateway)
	L2ERC721GatewayAddr        = common.HexToAddress(L2ERC721Gateway)
	L2TxFeeVaultAddr           = common.HexToAddress(L2TxFeeVault)

	Predeploys = make(map[string]*common.Address)
)

func init() {
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["L2Sequencer"] = &L2SequencerAddr
	Predeploys["Gov"] = &GovAddr
	Predeploys["Submitter"] = &SubmitterAddr
	Predeploys["L2GatewayRouter"] = &L2GatewayRouterAddr
	Predeploys["L2ETHGateway"] = &L2ETHGatewayAddr
	Predeploys["L2StandardERC20Gateway"] = &L2StandardERC20GatewayAddr
	Predeploys["L2ERC721Gateway"] = &L2ERC721GatewayAddr
	Predeploys["L2TxFeeVault"] = &L2TxFeeVaultAddr
}
