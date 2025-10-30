// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC20PriceOracleTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type ERC20PriceOracleTokenInfo struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}

// ERC20PriceOracleMetaData contains all meta data concerning the ERC20PriceOracle contract.
var ERC20PriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPercent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"FeeDiscountPercentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"PriceRatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"}],\"name\":\"TokenDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"TokenInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScale\",\"type\":\"uint256\"}],\"name\":\"TokenScaleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"tokenIDs\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_tokenIDs\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"}],\"name\":\"batchUpdatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"deactivateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"feeDiscountPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getFeeDiscountPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIdByAddress\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"internalType\":\"structERC20PriceOracle.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"isTokenActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"priceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_scale\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_tokenIDs\",\"type\":\"uint16[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_balanceSlots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_scales\",\"type\":\"uint256[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowListEnabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenRegistration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"updateFeeDiscountPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updatePriceRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_scale\",\"type\":\"uint256\"}],\"name\":\"updateTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_newScale\",\"type\":\"uint256\"}],\"name\":\"updateTokenScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052606a805460ff1916600117905534801561001c575f80fd5b5061002561002a565b6100e6565b5f54610100900460ff16156100955760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100e4575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6122e880620000f45f395ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c8063a127d680116100f3578063e014d85e11610093578063ef0fde0f1161006e578063ef0fde0f14610490578063efeadb6d146104a3578063f2fde38b146104b6578063fce40489146104c9575f80fd5b8063e014d85e14610457578063e2f53f2c1461046a578063e3de72a51461047d575f80fd5b8063b10b69ee116100ce578063b10b69ee146103ff578063c405454b14610412578063c4d66de814610431578063dddc98be14610444575f80fd5b8063a127d680146103c6578063a313d007146103d9578063a885842e146103ec575f80fd5b80632d59c0721161015e57806385519c361161013957806385519c36146103005780638c399691146103855780638cbab7e4146103985780638da5cb5b146103ab575f80fd5b80632d59c072146102af578063715018a6146102c2578063724f91ce146102ca575f80fd5b806322bd5c1c1161018e57806322bd5c1c1461025d5780632848aeaf1461027a5780632a1ea5a21461029c575f80fd5b80631684d242146101b457806319904c33146101c95780631c58e793146101fb575b5f80fd5b6101c76101c2366004611c81565b6104dc565b005b6101e86101d7366004611c81565b60676020525f908152604090205481565b6040519081526020015b60405180910390f35b61020e610209366004611c81565b61059a565b6040516101f291905f60a0820190506001600160a01b0383511682526020830151602083015260408301511515604083015260ff60608401511660608301526080830151608083015292915050565b606a5461026a9060ff1681565b60405190151581526020016101f2565b61026a610288366004611cb7565b60696020525f908152604090205460ff1681565b6101e86102aa366004611c81565b610678565b6101c76102bd366004611cdf565b6106e0565b6101c7610a48565b6102ed6102d8366004611cb7565b60666020525f908152604090205461ffff1681565b60405161ffff90911681526020016101f2565b61034b61030e366004611c81565b60656020525f908152604090208054600182015460028301546003909301546001600160a01b0390921692909160ff808316926101009004169085565b604080516001600160a01b03909616865260208601949094529115159284019290925260ff9091166060830152608082015260a0016101f2565b6101c7610393366004611e9e565b610a5b565b6102ed6103a6366004611cb7565b610c4a565b6033546040516001600160a01b0390911681526020016101f2565b6101c76103d4366004611efe565b610cb9565b6101e86103e7366004611c81565b610e09565b6101e86103fa366004611c81565b610e74565b6101c761040d366004611efe565b610edc565b6101e8610420366004611c81565b60686020525f908152604090205481565b6101c761043f366004611cb7565b61101c565b6101e8610452366004611efe565b6111da565b6101c7610465366004611f26565b611327565b6101c7610478366004611fc8565b6113f0565b6101c761048b3660046120bf565b611512565b6101c761049e366004611efe565b61164b565b6101c76104b136600461217b565b61178e565b6101c76104c4366004611cb7565b6117fb565b61026a6104d7366004611c81565b6118a5565b6104e46118e8565b61ffff81165f908152606560205260409020546001600160a01b0316610536576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff81165f8181526065602052604080822060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517fa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef9190a250565b6040805160a0810182525f8082526020808301829052828401829052606083018290526080830182905261ffff851682526065905291909120546001600160a01b0316610613576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f90815260656020908152604091829020825160a08101845281546001600160a01b03168152600182015492810192909252600281015460ff80821615159484019490945261010090049092166060820152600390910154608082015290565b61ffff81165f908152606560205260408120546001600160a01b03166106ca576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f9081526067602052604090205490565b6106e86118e8565b61ffff85165f908152606560205260409020546001600160a01b031661073a576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841661077a576040517f1eb00b0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384165f9081526066602052604090205461ffff1680158015906107ad57508561ffff168161ffff1614155b156107e4576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60129050856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610861575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261085e91810190612194565b60015b156108695790505b5f60655f8961ffff1661ffff1681526020019081526020015f205f015f9054906101000a90046001600160a01b031690506040518060a00160405280886001600160a01b0316815260200187815260200186151581526020018360ff1681526020018581525060655f8a61ffff1661ffff1681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015f6101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff021916908360ff16021790555060808201518160030155905050866001600160a01b0316816001600160a01b0316146109d9576001600160a01b038181165f9081526066602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155928a1682529020805490911661ffff8a161790555b866001600160a01b03168861ffff167f60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa988888689604051610a369493929190938452911515602084015260ff166040830152606082015260800190565b60405180910390a35050505050505050565b610a506118e8565b610a595f61195c565b565b606a5460ff168015610a7c5750335f9081526069602052604090205460ff16155b8015610a9357506033546001600160a01b03163314155b15610aca576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051825114610b05576040517f9d89020a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610c45575f6001600160a01b031660655f858481518110610b2f57610b2f6121b4565b60209081029190910181015161ffff1682528101919091526040015f20546001600160a01b031614610c3d57818181518110610b6d57610b6d6121b4565b60200260200101515f0315610c3d57818181518110610b8e57610b8e6121b4565b602002602001015160675f858481518110610bab57610bab6121b4565b602002602001015161ffff1661ffff1681526020019081526020015f2081905550828181518110610bde57610bde6121b4565b602002602001015161ffff167fd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467838381518110610c1d57610c1d6121b4565b6020026020010151604051610c3491815260200190565b60405180910390a25b600101610b07565b505050565b6001600160a01b0381165f9081526066602052604081205461ffff1680158015610c7c57506001600160a01b03831615155b15610cb3576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b606a5460ff168015610cda5750335f9081526069602052604090205460ff16155b8015610cf157506033546001600160a01b03163314155b15610d28576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f908152606560205260409020546001600160a01b0316610d7a576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612710811115610db6576040517fb92e9c7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f8181526068602052604090819020839055517fcfcfd26faabba2194ac421a7561d75b23f925b438bca48676101e3c2579ec42690610dfd9084815260200190565b60405180910390a25050565b61ffff81165f908152606560205260408120546001600160a01b0316610e5b576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f9081526065602052604090206003015490565b61ffff81165f908152606560205260408120546001600160a01b0316610ec6576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f9081526068602052604090205490565b606a5460ff168015610efd5750335f9081526069602052604090205460ff16155b8015610f1457506033546001600160a01b03163314155b15610f4b576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f908152606560205260409020546001600160a01b0316610f9d576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03610fd5576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f8181526067602052604090819020839055517fd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c7046790610dfd9084815260200190565b5f54610100900460ff161580801561103a57505f54600160ff909116105b806110535750303b15801561105357505f5460ff166001145b6110e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611140575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6111498261195c565b606a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156111d6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61ffff82165f908152606560205260408120546001600160a01b031661122c576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff83165f818152606560209081526040808320815160a08101835281546001600160a01b03168152600182015481850152600282015460ff80821615158386015261010090910416606082015260039091015460808201529383526067909152812054908190036112ca576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808260800151856112db91906121e1565b6112e5919061221d565b9250825f0361131f576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505092915050565b61132f6118e8565b61133b848484846119c5565b61ffff84165f81815260656020908152604091829020825160a08101845281546001600160a01b039081168252600183015482850152600283015460ff808216151584880181905261010090920416606080850182905260039095015460808086019190915287518b815296870192909252858701529284018790529351909493881693927fb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e1192908290030190a35050505050565b6113f86118e8565b8251845114158061140b57508151845114155b8061141857508051845114155b1561144f576040517f9d89020a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b84518110156114d2576114ca85828151811061146f5761146f6121b4565b6020026020010151858381518110611489576114896121b4565b60200260200101518584815181106114a3576114a36121b4565b60200260200101518585815181106114bd576114bd6121b4565b60200260200101516119c5565b600101611451565b507f31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a8484604051611504929190612255565b60405180910390a150505050565b61151a6118e8565b8051825114611555576040517fd9183d2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610c4557818181518110611572576115726121b4565b602002602001015160695f85848151811061158f5761158f6121b4565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548160ff0219169083151502179055508281815181106115de576115de6121b4565b60200260200101516001600160a01b03167f6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588838381518110611622576116226121b4565b602002602001015160405161163b911515815260200190565b60405180910390a2600101611557565b606a5460ff16801561166c5750335f9081526069602052604090205460ff16155b801561168357506033546001600160a01b03163314155b156116ba576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f908152606560205260409020546001600160a01b031661170c576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03611744576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f8181526065602052604090819020600301839055517f7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e7490610dfd9084815260200190565b6117966118e8565b606a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb39060200160405180910390a150565b6118036118e8565b6001600160a01b038116611899576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016110db565b6118a28161195c565b50565b61ffff81165f908152606560205260408120546001600160a01b03166118cc57505f919050565b5061ffff165f9081526065602052604090206002015460ff1690565b6033546001600160a01b03163314610a59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016110db565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b038316611a05576040517f1eb00b0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8361ffff165f03611a42576040517f6aa2a93700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff84165f908152606560205260409020546001600160a01b031615611a95576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0383165f9081526066602052604090205461ffff1615611ae8576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60129050836001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611b65575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611b6291810190612194565b60015b15611b6d5790505b6040805160a0810182526001600160a01b0395861680825260208083019687525f83850181815260ff968716606086019081526080860198895261ffff909b1680835260658452868320955186549b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909b169a909a1785559751600185015596516002840180549a51909616610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff911515919091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009a8b16171790945593516003909101559183526066905290208054909216179055565b803561ffff81168114611c7c575f80fd5b919050565b5f60208284031215611c91575f80fd5b611c9a82611c6b565b9392505050565b80356001600160a01b0381168114611c7c575f80fd5b5f60208284031215611cc7575f80fd5b611c9a82611ca1565b80358015158114611c7c575f80fd5b5f805f805f60a08688031215611cf3575f80fd5b611cfc86611c6b565b9450611d0a60208701611ca1565b935060408601359250611d1f60608701611cd0565b949793965091946080013592915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611da457611da4611d30565b604052919050565b5f67ffffffffffffffff821115611dc557611dc5611d30565b5060051b60200190565b5f82601f830112611dde575f80fd5b81356020611df3611dee83611dac565b611d5d565b8083825260208201915060208460051b870101935086841115611e14575f80fd5b602086015b84811015611e3757611e2a81611c6b565b8352918301918301611e19565b509695505050505050565b5f82601f830112611e51575f80fd5b81356020611e61611dee83611dac565b8083825260208201915060208460051b870101935086841115611e82575f80fd5b602086015b84811015611e375780358352918301918301611e87565b5f8060408385031215611eaf575f80fd5b823567ffffffffffffffff80821115611ec6575f80fd5b611ed286838701611dcf565b93506020850135915080821115611ee7575f80fd5b50611ef485828601611e42565b9150509250929050565b5f8060408385031215611f0f575f80fd5b611f1883611c6b565b946020939093013593505050565b5f805f8060808587031215611f39575f80fd5b611f4285611c6b565b9350611f5060208601611ca1565b93969395505050506040820135916060013590565b5f82601f830112611f74575f80fd5b81356020611f84611dee83611dac565b8083825260208201915060208460051b870101935086841115611fa5575f80fd5b602086015b84811015611e3757611fbb81611ca1565b8352918301918301611faa565b5f805f8060808587031215611fdb575f80fd5b843567ffffffffffffffff80821115611ff2575f80fd5b611ffe88838901611dcf565b9550602091508187013581811115612014575f80fd5b61202089828a01611f65565b955050604087013581811115612034575f80fd5b8701601f81018913612044575f80fd5b8035612052611dee82611dac565b81815260059190911b8201840190848101908b831115612070575f80fd5b928501925b8284101561208e57833582529285019290850190612075565b965050505060608701359150808211156120a6575f80fd5b506120b387828801611e42565b91505092959194509250565b5f80604083850312156120d0575f80fd5b823567ffffffffffffffff808211156120e7575f80fd5b6120f386838701611f65565b9350602091508185013581811115612109575f80fd5b85019050601f8101861361211b575f80fd5b8035612129611dee82611dac565b81815260059190911b82018301908381019088831115612147575f80fd5b928401925b8284101561216c5761215d84611cd0565b8252928401929084019061214c565b80955050505050509250929050565b5f6020828403121561218b575f80fd5b611c9a82611cd0565b5f602082840312156121a4575f80fd5b815160ff81168114611c9a575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417610cb3577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f82612250577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b604080825283519082018190525f906020906060840190828701845b8281101561229157815161ffff1684529284019290840190600101612271565b505050838103828501528451808252858301918301905f5b818110156122ce5783516001600160a01b0316835292840192918401916001016122a9565b509097965050505050505056fea164736f6c6343000818000a",
}

// ERC20PriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PriceOracleMetaData.ABI instead.
var ERC20PriceOracleABI = ERC20PriceOracleMetaData.ABI

// ERC20PriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20PriceOracleMetaData.Bin instead.
var ERC20PriceOracleBin = ERC20PriceOracleMetaData.Bin

// DeployERC20PriceOracle deploys a new Ethereum contract, binding an instance of ERC20PriceOracle to it.
func DeployERC20PriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20PriceOracle, error) {
	parsed, err := ERC20PriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20PriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20PriceOracle{ERC20PriceOracleCaller: ERC20PriceOracleCaller{contract: contract}, ERC20PriceOracleTransactor: ERC20PriceOracleTransactor{contract: contract}, ERC20PriceOracleFilterer: ERC20PriceOracleFilterer{contract: contract}}, nil
}

// ERC20PriceOracle is an auto generated Go binding around an Ethereum contract.
type ERC20PriceOracle struct {
	ERC20PriceOracleCaller     // Read-only binding to the contract
	ERC20PriceOracleTransactor // Write-only binding to the contract
	ERC20PriceOracleFilterer   // Log filterer for contract events
}

// ERC20PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PriceOracleSession struct {
	Contract     *ERC20PriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PriceOracleCallerSession struct {
	Contract *ERC20PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PriceOracleTransactorSession struct {
	Contract     *ERC20PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PriceOracleRaw struct {
	Contract *ERC20PriceOracle // Generic contract binding to access the raw methods on
}

// ERC20PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PriceOracleCallerRaw struct {
	Contract *ERC20PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PriceOracleTransactorRaw struct {
	Contract *ERC20PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PriceOracle creates a new instance of ERC20PriceOracle, bound to a specific deployed contract.
func NewERC20PriceOracle(address common.Address, backend bind.ContractBackend) (*ERC20PriceOracle, error) {
	contract, err := bindERC20PriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracle{ERC20PriceOracleCaller: ERC20PriceOracleCaller{contract: contract}, ERC20PriceOracleTransactor: ERC20PriceOracleTransactor{contract: contract}, ERC20PriceOracleFilterer: ERC20PriceOracleFilterer{contract: contract}}, nil
}

// NewERC20PriceOracleCaller creates a new read-only instance of ERC20PriceOracle, bound to a specific deployed contract.
func NewERC20PriceOracleCaller(address common.Address, caller bind.ContractCaller) (*ERC20PriceOracleCaller, error) {
	contract, err := bindERC20PriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleCaller{contract: contract}, nil
}

// NewERC20PriceOracleTransactor creates a new write-only instance of ERC20PriceOracle, bound to a specific deployed contract.
func NewERC20PriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PriceOracleTransactor, error) {
	contract, err := bindERC20PriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTransactor{contract: contract}, nil
}

// NewERC20PriceOracleFilterer creates a new log filterer instance of ERC20PriceOracle, bound to a specific deployed contract.
func NewERC20PriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PriceOracleFilterer, error) {
	contract, err := bindERC20PriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleFilterer{contract: contract}, nil
}

// bindERC20PriceOracle binds a generic wrapper to an already deployed contract.
func bindERC20PriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20PriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PriceOracle *ERC20PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PriceOracle.Contract.ERC20PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PriceOracle *ERC20PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.ERC20PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PriceOracle *ERC20PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.ERC20PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PriceOracle *ERC20PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PriceOracle *ERC20PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PriceOracle *ERC20PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) AllowList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "allowList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleSession) AllowList(arg0 common.Address) (bool, error) {
	return _ERC20PriceOracle.Contract.AllowList(&_ERC20PriceOracle.CallOpts, arg0)
}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) AllowList(arg0 common.Address) (bool, error) {
	return _ERC20PriceOracle.Contract.AllowList(&_ERC20PriceOracle.CallOpts, arg0)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleSession) AllowListEnabled() (bool, error) {
	return _ERC20PriceOracle.Contract.AllowListEnabled(&_ERC20PriceOracle.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) AllowListEnabled() (bool, error) {
	return _ERC20PriceOracle.Contract.AllowListEnabled(&_ERC20PriceOracle.CallOpts)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) CalculateTokenAmount(opts *bind.CallOpts, _tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "calculateTokenAmount", _tokenID, _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_ERC20PriceOracle *ERC20PriceOracleSession) CalculateTokenAmount(_tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.CalculateTokenAmount(&_ERC20PriceOracle.CallOpts, _tokenID, _ethAmount)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) CalculateTokenAmount(_tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.CalculateTokenAmount(&_ERC20PriceOracle.CallOpts, _tokenID, _ethAmount)
}

// FeeDiscountPercent is a free data retrieval call binding the contract method 0xc405454b.
//
// Solidity: function feeDiscountPercent(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) FeeDiscountPercent(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "feeDiscountPercent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDiscountPercent is a free data retrieval call binding the contract method 0xc405454b.
//
// Solidity: function feeDiscountPercent(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleSession) FeeDiscountPercent(arg0 uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.FeeDiscountPercent(&_ERC20PriceOracle.CallOpts, arg0)
}

// FeeDiscountPercent is a free data retrieval call binding the contract method 0xc405454b.
//
// Solidity: function feeDiscountPercent(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) FeeDiscountPercent(arg0 uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.FeeDiscountPercent(&_ERC20PriceOracle.CallOpts, arg0)
}

// GetFeeDiscountPercent is a free data retrieval call binding the contract method 0xa885842e.
//
// Solidity: function getFeeDiscountPercent(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) GetFeeDiscountPercent(opts *bind.CallOpts, _tokenID uint16) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "getFeeDiscountPercent", _tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeDiscountPercent is a free data retrieval call binding the contract method 0xa885842e.
//
// Solidity: function getFeeDiscountPercent(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleSession) GetFeeDiscountPercent(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetFeeDiscountPercent(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetFeeDiscountPercent is a free data retrieval call binding the contract method 0xa885842e.
//
// Solidity: function getFeeDiscountPercent(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) GetFeeDiscountPercent(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetFeeDiscountPercent(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) GetTokenIdByAddress(opts *bind.CallOpts, tokenAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "getTokenIdByAddress", tokenAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleSession) GetTokenIdByAddress(tokenAddress common.Address) (uint16, error) {
	return _ERC20PriceOracle.Contract.GetTokenIdByAddress(&_ERC20PriceOracle.CallOpts, tokenAddress)
}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) GetTokenIdByAddress(tokenAddress common.Address) (uint16, error) {
	return _ERC20PriceOracle.Contract.GetTokenIdByAddress(&_ERC20PriceOracle.CallOpts, tokenAddress)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_ERC20PriceOracle *ERC20PriceOracleCaller) GetTokenInfo(opts *bind.CallOpts, _tokenID uint16) (ERC20PriceOracleTokenInfo, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "getTokenInfo", _tokenID)

	if err != nil {
		return *new(ERC20PriceOracleTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20PriceOracleTokenInfo)).(*ERC20PriceOracleTokenInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_ERC20PriceOracle *ERC20PriceOracleSession) GetTokenInfo(_tokenID uint16) (ERC20PriceOracleTokenInfo, error) {
	return _ERC20PriceOracle.Contract.GetTokenInfo(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) GetTokenInfo(_tokenID uint16) (ERC20PriceOracleTokenInfo, error) {
	return _ERC20PriceOracle.Contract.GetTokenInfo(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) GetTokenPrice(opts *bind.CallOpts, _tokenID uint16) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "getTokenPrice", _tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleSession) GetTokenPrice(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetTokenPrice(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) GetTokenPrice(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetTokenPrice(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) GetTokenScale(opts *bind.CallOpts, _tokenID uint16) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "getTokenScale", _tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleSession) GetTokenScale(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetTokenScale(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) GetTokenScale(_tokenID uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.GetTokenScale(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) IsTokenActive(opts *bind.CallOpts, _tokenID uint16) (bool, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "isTokenActive", _tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleSession) IsTokenActive(_tokenID uint16) (bool, error) {
	return _ERC20PriceOracle.Contract.IsTokenActive(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) IsTokenActive(_tokenID uint16) (bool, error) {
	return _ERC20PriceOracle.Contract.IsTokenActive(&_ERC20PriceOracle.CallOpts, _tokenID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20PriceOracle *ERC20PriceOracleSession) Owner() (common.Address, error) {
	return _ERC20PriceOracle.Contract.Owner(&_ERC20PriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) Owner() (common.Address, error) {
	return _ERC20PriceOracle.Contract.Owner(&_ERC20PriceOracle.CallOpts)
}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) PriceRatio(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "priceRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleSession) PriceRatio(arg0 uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.PriceRatio(&_ERC20PriceOracle.CallOpts, arg0)
}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) PriceRatio(arg0 uint16) (*big.Int, error) {
	return _ERC20PriceOracle.Contract.PriceRatio(&_ERC20PriceOracle.CallOpts, arg0)
}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) TokenRegistration(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "tokenRegistration", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleSession) TokenRegistration(arg0 common.Address) (uint16, error) {
	return _ERC20PriceOracle.Contract.TokenRegistration(&_ERC20PriceOracle.CallOpts, arg0)
}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) TokenRegistration(arg0 common.Address) (uint16, error) {
	return _ERC20PriceOracle.Contract.TokenRegistration(&_ERC20PriceOracle.CallOpts, arg0)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x85519c36.
//
// Solidity: function tokenRegistry(uint16 ) view returns(address tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleCaller) TokenRegistry(opts *bind.CallOpts, arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	var out []interface{}
	err := _ERC20PriceOracle.contract.Call(opts, &out, "tokenRegistry", arg0)

	outstruct := new(struct {
		TokenAddress common.Address
		BalanceSlot  [32]byte
		IsActive     bool
		Decimals     uint8
		Scale        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BalanceSlot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Scale = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenRegistry is a free data retrieval call binding the contract method 0x85519c36.
//
// Solidity: function tokenRegistry(uint16 ) view returns(address tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleSession) TokenRegistry(arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	return _ERC20PriceOracle.Contract.TokenRegistry(&_ERC20PriceOracle.CallOpts, arg0)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x85519c36.
//
// Solidity: function tokenRegistry(uint16 ) view returns(address tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleCallerSession) TokenRegistry(arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	return _ERC20PriceOracle.Contract.TokenRegistry(&_ERC20PriceOracle.CallOpts, arg0)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) BatchUpdatePrices(opts *bind.TransactOpts, _tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "batchUpdatePrices", _tokenIDs, _prices)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) BatchUpdatePrices(_tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.BatchUpdatePrices(&_ERC20PriceOracle.TransactOpts, _tokenIDs, _prices)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) BatchUpdatePrices(_tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.BatchUpdatePrices(&_ERC20PriceOracle.TransactOpts, _tokenIDs, _prices)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x1684d242.
//
// Solidity: function deactivateToken(uint16 _tokenID) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) DeactivateToken(opts *bind.TransactOpts, _tokenID uint16) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "deactivateToken", _tokenID)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x1684d242.
//
// Solidity: function deactivateToken(uint16 _tokenID) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) DeactivateToken(_tokenID uint16) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.DeactivateToken(&_ERC20PriceOracle.TransactOpts, _tokenID)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x1684d242.
//
// Solidity: function deactivateToken(uint16 _tokenID) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) DeactivateToken(_tokenID uint16) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.DeactivateToken(&_ERC20PriceOracle.TransactOpts, _tokenID)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.Initialize(&_ERC20PriceOracle.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.Initialize(&_ERC20PriceOracle.TransactOpts, owner_)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) RegisterToken(opts *bind.TransactOpts, _tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "registerToken", _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) RegisterToken(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RegisterToken(&_ERC20PriceOracle.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) RegisterToken(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RegisterToken(&_ERC20PriceOracle.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) RegisterTokens(opts *bind.TransactOpts, _tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "registerTokens", _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) RegisterTokens(_tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RegisterTokens(&_ERC20PriceOracle.TransactOpts, _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) RegisterTokens(_tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RegisterTokens(&_ERC20PriceOracle.TransactOpts, _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RenounceOwnership(&_ERC20PriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.RenounceOwnership(&_ERC20PriceOracle.TransactOpts)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "setAllowList", user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.SetAllowList(&_ERC20PriceOracle.TransactOpts, user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.SetAllowList(&_ERC20PriceOracle.TransactOpts, user, val)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "setAllowListEnabled", _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.SetAllowListEnabled(&_ERC20PriceOracle.TransactOpts, _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.SetAllowListEnabled(&_ERC20PriceOracle.TransactOpts, _allowListEnabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.TransferOwnership(&_ERC20PriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.TransferOwnership(&_ERC20PriceOracle.TransactOpts, newOwner)
}

// UpdateFeeDiscountPercent is a paid mutator transaction binding the contract method 0xa127d680.
//
// Solidity: function updateFeeDiscountPercent(uint16 _tokenID, uint256 _newPercent) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) UpdateFeeDiscountPercent(opts *bind.TransactOpts, _tokenID uint16, _newPercent *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "updateFeeDiscountPercent", _tokenID, _newPercent)
}

// UpdateFeeDiscountPercent is a paid mutator transaction binding the contract method 0xa127d680.
//
// Solidity: function updateFeeDiscountPercent(uint16 _tokenID, uint256 _newPercent) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) UpdateFeeDiscountPercent(_tokenID uint16, _newPercent *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateFeeDiscountPercent(&_ERC20PriceOracle.TransactOpts, _tokenID, _newPercent)
}

// UpdateFeeDiscountPercent is a paid mutator transaction binding the contract method 0xa127d680.
//
// Solidity: function updateFeeDiscountPercent(uint16 _tokenID, uint256 _newPercent) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) UpdateFeeDiscountPercent(_tokenID uint16, _newPercent *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateFeeDiscountPercent(&_ERC20PriceOracle.TransactOpts, _tokenID, _newPercent)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) UpdatePriceRatio(opts *bind.TransactOpts, _tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "updatePriceRatio", _tokenID, _newPrice)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) UpdatePriceRatio(_tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdatePriceRatio(&_ERC20PriceOracle.TransactOpts, _tokenID, _newPrice)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) UpdatePriceRatio(_tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdatePriceRatio(&_ERC20PriceOracle.TransactOpts, _tokenID, _newPrice)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) UpdateTokenInfo(opts *bind.TransactOpts, _tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "updateTokenInfo", _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) UpdateTokenInfo(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateTokenInfo(&_ERC20PriceOracle.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) UpdateTokenInfo(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateTokenInfo(&_ERC20PriceOracle.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactor) UpdateTokenScale(opts *bind.TransactOpts, _tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.contract.Transact(opts, "updateTokenScale", _tokenID, _newScale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleSession) UpdateTokenScale(_tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateTokenScale(&_ERC20PriceOracle.TransactOpts, _tokenID, _newScale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_ERC20PriceOracle *ERC20PriceOracleTransactorSession) UpdateTokenScale(_tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _ERC20PriceOracle.Contract.UpdateTokenScale(&_ERC20PriceOracle.TransactOpts, _tokenID, _newScale)
}

// ERC20PriceOracleAllowListEnabledUpdatedIterator is returned from FilterAllowListEnabledUpdated and is used to iterate over the raw logs and unpacked data for AllowListEnabledUpdated events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleAllowListEnabledUpdatedIterator struct {
	Event *ERC20PriceOracleAllowListEnabledUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleAllowListEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleAllowListEnabledUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleAllowListEnabledUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleAllowListEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleAllowListEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleAllowListEnabledUpdated represents a AllowListEnabledUpdated event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleAllowListEnabledUpdated struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowListEnabledUpdated is a free log retrieval operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterAllowListEnabledUpdated(opts *bind.FilterOpts) (*ERC20PriceOracleAllowListEnabledUpdatedIterator, error) {

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleAllowListEnabledUpdatedIterator{contract: _ERC20PriceOracle.contract, event: "AllowListEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListEnabledUpdated is a free log subscription operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchAllowListEnabledUpdated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleAllowListEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleAllowListEnabledUpdated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowListEnabledUpdated is a log parse operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseAllowListEnabledUpdated(log types.Log) (*ERC20PriceOracleAllowListEnabledUpdated, error) {
	event := new(ERC20PriceOracleAllowListEnabledUpdated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleAllowListSetIterator is returned from FilterAllowListSet and is used to iterate over the raw logs and unpacked data for AllowListSet events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleAllowListSetIterator struct {
	Event *ERC20PriceOracleAllowListSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleAllowListSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleAllowListSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleAllowListSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleAllowListSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleAllowListSet represents a AllowListSet event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleAllowListSet struct {
	User common.Address
	Val  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowListSet is a free log retrieval operation binding the contract event 0x6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588.
//
// Solidity: event AllowListSet(address indexed user, bool val)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterAllowListSet(opts *bind.FilterOpts, user []common.Address) (*ERC20PriceOracleAllowListSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "AllowListSet", userRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleAllowListSetIterator{contract: _ERC20PriceOracle.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

// WatchAllowListSet is a free log subscription operation binding the contract event 0x6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588.
//
// Solidity: event AllowListSet(address indexed user, bool val)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleAllowListSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "AllowListSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleAllowListSet)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "AllowListSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowListSet is a log parse operation binding the contract event 0x6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588.
//
// Solidity: event AllowListSet(address indexed user, bool val)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseAllowListSet(log types.Log) (*ERC20PriceOracleAllowListSet, error) {
	event := new(ERC20PriceOracleAllowListSet)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleFeeDiscountPercentUpdatedIterator is returned from FilterFeeDiscountPercentUpdated and is used to iterate over the raw logs and unpacked data for FeeDiscountPercentUpdated events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleFeeDiscountPercentUpdatedIterator struct {
	Event *ERC20PriceOracleFeeDiscountPercentUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleFeeDiscountPercentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleFeeDiscountPercentUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleFeeDiscountPercentUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleFeeDiscountPercentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleFeeDiscountPercentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleFeeDiscountPercentUpdated represents a FeeDiscountPercentUpdated event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleFeeDiscountPercentUpdated struct {
	TokenID    uint16
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeDiscountPercentUpdated is a free log retrieval operation binding the contract event 0xcfcfd26faabba2194ac421a7561d75b23f925b438bca48676101e3c2579ec426.
//
// Solidity: event FeeDiscountPercentUpdated(uint16 indexed tokenID, uint256 newPercent)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterFeeDiscountPercentUpdated(opts *bind.FilterOpts, tokenID []uint16) (*ERC20PriceOracleFeeDiscountPercentUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "FeeDiscountPercentUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleFeeDiscountPercentUpdatedIterator{contract: _ERC20PriceOracle.contract, event: "FeeDiscountPercentUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeDiscountPercentUpdated is a free log subscription operation binding the contract event 0xcfcfd26faabba2194ac421a7561d75b23f925b438bca48676101e3c2579ec426.
//
// Solidity: event FeeDiscountPercentUpdated(uint16 indexed tokenID, uint256 newPercent)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchFeeDiscountPercentUpdated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleFeeDiscountPercentUpdated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "FeeDiscountPercentUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleFeeDiscountPercentUpdated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "FeeDiscountPercentUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeDiscountPercentUpdated is a log parse operation binding the contract event 0xcfcfd26faabba2194ac421a7561d75b23f925b438bca48676101e3c2579ec426.
//
// Solidity: event FeeDiscountPercentUpdated(uint16 indexed tokenID, uint256 newPercent)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseFeeDiscountPercentUpdated(log types.Log) (*ERC20PriceOracleFeeDiscountPercentUpdated, error) {
	event := new(ERC20PriceOracleFeeDiscountPercentUpdated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "FeeDiscountPercentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleInitializedIterator struct {
	Event *ERC20PriceOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleInitialized represents a Initialized event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20PriceOracleInitializedIterator, error) {

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleInitializedIterator{contract: _ERC20PriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleInitialized)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseInitialized(log types.Log) (*ERC20PriceOracleInitialized, error) {
	event := new(ERC20PriceOracleInitialized)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleOwnershipTransferredIterator struct {
	Event *ERC20PriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20PriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleOwnershipTransferredIterator{contract: _ERC20PriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleOwnershipTransferred)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20PriceOracleOwnershipTransferred, error) {
	event := new(ERC20PriceOracleOwnershipTransferred)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOraclePriceRatioUpdatedIterator is returned from FilterPriceRatioUpdated and is used to iterate over the raw logs and unpacked data for PriceRatioUpdated events raised by the ERC20PriceOracle contract.
type ERC20PriceOraclePriceRatioUpdatedIterator struct {
	Event *ERC20PriceOraclePriceRatioUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOraclePriceRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOraclePriceRatioUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOraclePriceRatioUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOraclePriceRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOraclePriceRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOraclePriceRatioUpdated represents a PriceRatioUpdated event raised by the ERC20PriceOracle contract.
type ERC20PriceOraclePriceRatioUpdated struct {
	TokenID  uint16
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceRatioUpdated is a free log retrieval operation binding the contract event 0xd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467.
//
// Solidity: event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterPriceRatioUpdated(opts *bind.FilterOpts, tokenID []uint16) (*ERC20PriceOraclePriceRatioUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "PriceRatioUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOraclePriceRatioUpdatedIterator{contract: _ERC20PriceOracle.contract, event: "PriceRatioUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceRatioUpdated is a free log subscription operation binding the contract event 0xd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467.
//
// Solidity: event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchPriceRatioUpdated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOraclePriceRatioUpdated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "PriceRatioUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOraclePriceRatioUpdated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "PriceRatioUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceRatioUpdated is a log parse operation binding the contract event 0xd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467.
//
// Solidity: event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParsePriceRatioUpdated(log types.Log) (*ERC20PriceOraclePriceRatioUpdated, error) {
	event := new(ERC20PriceOraclePriceRatioUpdated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "PriceRatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleTokenDeactivatedIterator is returned from FilterTokenDeactivated and is used to iterate over the raw logs and unpacked data for TokenDeactivated events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenDeactivatedIterator struct {
	Event *ERC20PriceOracleTokenDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleTokenDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleTokenDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleTokenDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleTokenDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleTokenDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleTokenDeactivated represents a TokenDeactivated event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenDeactivated struct {
	TokenID uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeactivated is a free log retrieval operation binding the contract event 0xa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef.
//
// Solidity: event TokenDeactivated(uint16 indexed tokenID)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterTokenDeactivated(opts *bind.FilterOpts, tokenID []uint16) (*ERC20PriceOracleTokenDeactivatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "TokenDeactivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTokenDeactivatedIterator{contract: _ERC20PriceOracle.contract, event: "TokenDeactivated", logs: logs, sub: sub}, nil
}

// WatchTokenDeactivated is a free log subscription operation binding the contract event 0xa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef.
//
// Solidity: event TokenDeactivated(uint16 indexed tokenID)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchTokenDeactivated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleTokenDeactivated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "TokenDeactivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleTokenDeactivated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenDeactivated is a log parse operation binding the contract event 0xa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef.
//
// Solidity: event TokenDeactivated(uint16 indexed tokenID)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseTokenDeactivated(log types.Log) (*ERC20PriceOracleTokenDeactivated, error) {
	event := new(ERC20PriceOracleTokenDeactivated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleTokenInfoUpdatedIterator is returned from FilterTokenInfoUpdated and is used to iterate over the raw logs and unpacked data for TokenInfoUpdated events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenInfoUpdatedIterator struct {
	Event *ERC20PriceOracleTokenInfoUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleTokenInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleTokenInfoUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleTokenInfoUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleTokenInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleTokenInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleTokenInfoUpdated represents a TokenInfoUpdated event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenInfoUpdated struct {
	TokenID      uint16
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenInfoUpdated is a free log retrieval operation binding the contract event 0x60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa9.
//
// Solidity: event TokenInfoUpdated(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterTokenInfoUpdated(opts *bind.FilterOpts, tokenID []uint16, tokenAddress []common.Address) (*ERC20PriceOracleTokenInfoUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "TokenInfoUpdated", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTokenInfoUpdatedIterator{contract: _ERC20PriceOracle.contract, event: "TokenInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenInfoUpdated is a free log subscription operation binding the contract event 0x60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa9.
//
// Solidity: event TokenInfoUpdated(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchTokenInfoUpdated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleTokenInfoUpdated, tokenID []uint16, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "TokenInfoUpdated", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleTokenInfoUpdated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenInfoUpdated is a log parse operation binding the contract event 0x60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa9.
//
// Solidity: event TokenInfoUpdated(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseTokenInfoUpdated(log types.Log) (*ERC20PriceOracleTokenInfoUpdated, error) {
	event := new(ERC20PriceOracleTokenInfoUpdated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenRegisteredIterator struct {
	Event *ERC20PriceOracleTokenRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleTokenRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleTokenRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleTokenRegistered represents a TokenRegistered event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenRegistered struct {
	TokenID      uint16
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e11.
//
// Solidity: event TokenRegistered(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterTokenRegistered(opts *bind.FilterOpts, tokenID []uint16, tokenAddress []common.Address) (*ERC20PriceOracleTokenRegisteredIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "TokenRegistered", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTokenRegisteredIterator{contract: _ERC20PriceOracle.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e11.
//
// Solidity: event TokenRegistered(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleTokenRegistered, tokenID []uint16, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "TokenRegistered", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleTokenRegistered)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenRegistered is a log parse operation binding the contract event 0xb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e11.
//
// Solidity: event TokenRegistered(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseTokenRegistered(log types.Log) (*ERC20PriceOracleTokenRegistered, error) {
	event := new(ERC20PriceOracleTokenRegistered)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleTokenScaleUpdatedIterator is returned from FilterTokenScaleUpdated and is used to iterate over the raw logs and unpacked data for TokenScaleUpdated events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenScaleUpdatedIterator struct {
	Event *ERC20PriceOracleTokenScaleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleTokenScaleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleTokenScaleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleTokenScaleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleTokenScaleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleTokenScaleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleTokenScaleUpdated represents a TokenScaleUpdated event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokenScaleUpdated struct {
	TokenID  uint16
	NewScale *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenScaleUpdated is a free log retrieval operation binding the contract event 0x7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e74.
//
// Solidity: event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterTokenScaleUpdated(opts *bind.FilterOpts, tokenID []uint16) (*ERC20PriceOracleTokenScaleUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "TokenScaleUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTokenScaleUpdatedIterator{contract: _ERC20PriceOracle.contract, event: "TokenScaleUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenScaleUpdated is a free log subscription operation binding the contract event 0x7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e74.
//
// Solidity: event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchTokenScaleUpdated(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleTokenScaleUpdated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "TokenScaleUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleTokenScaleUpdated)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenScaleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenScaleUpdated is a log parse operation binding the contract event 0x7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e74.
//
// Solidity: event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseTokenScaleUpdated(log types.Log) (*ERC20PriceOracleTokenScaleUpdated, error) {
	event := new(ERC20PriceOracleTokenScaleUpdated)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokenScaleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PriceOracleTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokensRegisteredIterator struct {
	Event *ERC20PriceOracleTokensRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PriceOracleTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PriceOracleTokensRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PriceOracleTokensRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PriceOracleTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PriceOracleTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PriceOracleTokensRegistered represents a TokensRegistered event raised by the ERC20PriceOracle contract.
type ERC20PriceOracleTokensRegistered struct {
	TokenIDs       []uint16
	TokenAddresses []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0x31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a.
//
// Solidity: event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) FilterTokensRegistered(opts *bind.FilterOpts) (*ERC20PriceOracleTokensRegisteredIterator, error) {

	logs, sub, err := _ERC20PriceOracle.contract.FilterLogs(opts, "TokensRegistered")
	if err != nil {
		return nil, err
	}
	return &ERC20PriceOracleTokensRegisteredIterator{contract: _ERC20PriceOracle.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0x31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a.
//
// Solidity: event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *ERC20PriceOracleTokensRegistered) (event.Subscription, error) {

	logs, sub, err := _ERC20PriceOracle.contract.WatchLogs(opts, "TokensRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PriceOracleTokensRegistered)
				if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensRegistered is a log parse operation binding the contract event 0x31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a.
//
// Solidity: event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses)
func (_ERC20PriceOracle *ERC20PriceOracleFilterer) ParseTokensRegistered(log types.Log) (*ERC20PriceOracleTokensRegistered, error) {
	event := new(ERC20PriceOracleTokensRegistered)
	if err := _ERC20PriceOracle.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
