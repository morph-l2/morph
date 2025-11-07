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

// IL2TokenRegistryTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type IL2TokenRegistryTokenInfo struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}

// L2TokenRegistryMetaData contains all meta data concerning the L2TokenRegistry contract.
var L2TokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DifferentLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPercent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"PriceRatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"}],\"name\":\"TokenActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"}],\"name\":\"TokenDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"TokenInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScale\",\"type\":\"uint256\"}],\"name\":\"TokenScaleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"tokenIDs\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_tokenIDs\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"}],\"name\":\"batchUpdatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_tokenIDs\",\"type\":\"uint16[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isActives\",\"type\":\"bool[]\"}],\"name\":\"batchUpdateTokenStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIdByAddress\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"internalType\":\"structIL2TokenRegistry.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"}],\"name\":\"isTokenActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"priceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_scale\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_tokenIDs\",\"type\":\"uint16[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_balanceSlots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_scales\",\"type\":\"uint256[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowListEnabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenRegistration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updatePriceRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_balanceSlot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_scale\",\"type\":\"uint256\"}],\"name\":\"updateTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_newScale\",\"type\":\"uint256\"}],\"name\":\"updateTokenScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052609b805460ff1916600117905534801561001c575f80fd5b5061002561002a565b6100e6565b5f54610100900460ff16156100955760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100e4575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6122ff80620000f45f395ff3fe608060405234801561000f575f80fd5b506004361061018f575f3560e01c80639190202e116100dd578063e2f53f2c11610088578063efeadb6d11610063578063efeadb6d1461043d578063f2fde38b14610450578063fce4048914610463575f80fd5b8063e2f53f2c14610404578063e3de72a514610417578063ef0fde0f1461042a575f80fd5b8063c4d66de8116100b8578063c4d66de8146103cb578063dddc98be146103de578063e014d85e146103f1575f80fd5b80639190202e14610392578063a313d007146103a5578063b10b69ee146103b8575f80fd5b8063715018a61161013d5780638c399691116101185780638c399691146103515780638cbab7e4146103645780638da5cb5b14610377575f80fd5b8063715018a61461028e578063724f91ce1461029657806385519c36146102cc575f80fd5b80632848aeaf1161016d5780632848aeaf146102445780632a1ea5a2146102665780632d59c07214610279575f80fd5b806319904c33146101935780631c58e793146101c557806322bd5c1c14610227575b5f80fd5b6101b26101a1366004611be9565b60996020525f908152604090205481565b6040519081526020015b60405180910390f35b6101d86101d3366004611be9565b610476565b6040516101bc91905f60a0820190506001600160a01b0383511682526020830151602083015260408301511515604083015260ff60608401511660608301526080830151608083015292915050565b609b546102349060ff1681565b60405190151581526020016101bc565b610234610252366004611c1f565b609a6020525f908152604090205460ff1681565b6101b2610274366004611be9565b610554565b61028c610287366004611c47565b6105bc565b005b61028c610936565b6102b96102a4366004611c1f565b60986020525f908152604090205461ffff1681565b60405161ffff90911681526020016101bc565b6103176102da366004611be9565b60976020525f908152604090208054600182015460028301546003909301546001600160a01b0390921692909160ff808316926101009004169085565b604080516001600160a01b03909616865260208601949094529115159284019290925260ff9091166060830152608082015260a0016101bc565b61028c61035f366004611e06565b610949565b6102b9610372366004611c1f565b610b38565b6033546040516001600160a01b0390911681526020016101bc565b61028c6103a0366004611eae565b610ba7565b6101b26103b3366004611be9565b610d48565b61028c6103c6366004611f15565b610db3565b61028c6103d9366004611c1f565b610eff565b6101b26103ec366004611f15565b6110bd565b61028c6103ff366004611f3d565b61120a565b61028c610412366004611fdf565b6112e5565b61028c6104253660046120d6565b611407565b61028c610438366004611f15565b611540565b61028c61044b366004612192565b611683565b61028c61045e366004611c1f565b6116f0565b610234610471366004611be9565b61179a565b6040805160a0810182525f8082526020808301829052828401829052606083018290526080830182905261ffff851682526097905291909120546001600160a01b03166104ef576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f90815260976020908152604091829020825160a08101845281546001600160a01b03168152600182015492810192909252600281015460ff80821615159484019490945261010090049092166060820152600390910154608082015290565b61ffff81165f908152609760205260408120546001600160a01b03166105a6576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f9081526099602052604090205490565b6105c46117dd565b6105cc611851565b61ffff85165f908152609760205260409020546001600160a01b031661061e576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841661065e576040517f1eb00b0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384165f9081526098602052604090205461ffff16801580159061069157508561ffff168161ffff1614155b156106c8576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60129050856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610745575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610742918101906121ab565b60015b1561074d5790505b5f60975f8961ffff1661ffff1681526020019081526020015f205f015f9054906101000a90046001600160a01b031690506040518060a00160405280886001600160a01b0316815260200187815260200186151581526020018360ff1681526020018581525060975f8a61ffff1661ffff1681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015f6101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff021916908360ff16021790555060808201518160030155905050866001600160a01b0316816001600160a01b0316146108bd576001600160a01b038181165f9081526098602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155928a1682529020805490911661ffff8a161790555b866001600160a01b03168861ffff167f60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa98888868960405161091a9493929190938452911515602084015260ff166040830152606082015260800190565b60405180910390a350505061092f6001606555565b5050505050565b61093e6117dd565b6109475f6118c4565b565b609b5460ff16801561096a5750335f908152609a602052604090205460ff16155b801561098157506033546001600160a01b03163314155b156109b8576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80518251146109f3576040517f9d89020a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610b33575f6001600160a01b031660975f858481518110610a1d57610a1d6121cb565b60209081029190910181015161ffff1682528101919091526040015f20546001600160a01b031614610b2b57818181518110610a5b57610a5b6121cb565b60200260200101515f0315610b2b57818181518110610a7c57610a7c6121cb565b602002602001015160995f858481518110610a9957610a996121cb565b602002602001015161ffff1661ffff1681526020019081526020015f2081905550828181518110610acc57610acc6121cb565b602002602001015161ffff167fd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467838381518110610b0b57610b0b6121cb565b6020026020010151604051610b2291815260200190565b60405180910390a25b6001016109f5565b505050565b6001600160a01b0381165f9081526098602052604081205461ffff1680158015610b6a57506001600160a01b03831615155b15610ba1576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b610baf6117dd565b828114610be8576040517f9d89020a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561092f575f858583818110610c0557610c056121cb565b9050602002016020810190610c1a9190611be9565b90505f848484818110610c2f57610c2f6121cb565b9050602002016020810190610c449190612192565b61ffff83165f908152609760205260409020549091506001600160a01b0316610c6e575050610d40565b61ffff82165f9081526097602052604090206002015460ff1680151582151514610d3c5761ffff83165f90815260976020526040902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168315801591909117909155610d0d5760405161ffff8416907fba78a86bd882b78fb95715a0d827188ec0a8ad3f500310d33a39b94c9ff677b8905f90a2610d3c565b60405161ffff8416907fa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef905f90a25b5050505b600101610bea565b61ffff81165f908152609760205260408120546001600160a01b0316610d9a576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061ffff165f9081526097602052604090206003015490565b609b5460ff168015610dd45750335f908152609a602052604090205460ff16155b8015610deb57506033546001600160a01b03163314155b15610e22576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f908152609760205260409020546001600160a01b0316610e74576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03610eac576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f8181526099602052604090819020839055517fd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c7046790610ef39084815260200190565b60405180910390a25050565b5f54610100900460ff1615808015610f1d57505f54600160ff909116105b80610f365750303b158015610f3657505f5460ff166001145b610fc7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611023575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61102c826118c4565b609b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156110b9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61ffff82165f908152609760205260408120546001600160a01b031661110f576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff83165f818152609760209081526040808320815160a08101835281546001600160a01b03168152600182015481850152600282015460ff80821615158386015261010090910416606082015260039091015460808201529383526099909152812054908190036111ad576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808260800151856111be91906121f8565b6111c89190612234565b9250825f03611202576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505092915050565b6112126117dd565b61121a611851565b6112268484848461192d565b61ffff84165f81815260976020908152604091829020825160a08101845281546001600160a01b039081168252600183015482850152600283015460ff808216151584880181905261010090920416606080850182905260039095015460808086019190915287518b815296870192909252858701529284018790529351909493881693927fb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e1192908290030190a3506112df6001606555565b50505050565b6112ed6117dd565b8251845114158061130057508151845114155b8061130d57508051845114155b15611344576040517f9d89020a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b84518110156113c7576113bf858281518110611364576113646121cb565b602002602001015185838151811061137e5761137e6121cb565b6020026020010151858481518110611398576113986121cb565b60200260200101518585815181106113b2576113b26121cb565b602002602001015161192d565b600101611346565b507f31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a84846040516113f992919061226c565b60405180910390a150505050565b61140f6117dd565b805182511461144a576040517fd9183d2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610b3357818181518110611467576114676121cb565b6020026020010151609a5f858481518110611484576114846121cb565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548160ff0219169083151502179055508281815181106114d3576114d36121cb565b60200260200101516001600160a01b03167f6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588838381518110611517576115176121cb565b6020026020010151604051611530911515815260200190565b60405180910390a260010161144c565b609b5460ff1680156115615750335f908152609a602052604090205460ff16155b801561157857506033546001600160a01b03163314155b156115af576040517f2af07d2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f908152609760205260409020546001600160a01b0316611601576040517fcbdb7b3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03611639576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff82165f8181526097602052604090819020600301839055517f7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e7490610ef39084815260200190565b61168b6117dd565b609b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb39060200160405180910390a150565b6116f86117dd565b6001600160a01b03811661178e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610fbe565b611797816118c4565b50565b61ffff81165f908152609760205260408120546001600160a01b03166117c157505f919050565b5061ffff165f9081526097602052604090206002015460ff1690565b6033546001600160a01b03163314610947576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610fbe565b6002606554036118bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610fbe565b6002606555565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b03831661196d576040517f1eb00b0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8361ffff165f036119aa576040517f6aa2a93700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff84165f908152609760205260409020546001600160a01b0316156119fd576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0383165f9081526098602052604090205461ffff1615611a50576040517f7d4fffb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60129050836001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611acd575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611aca918101906121ab565b60015b15611ad55790505b6040805160a0810182526001600160a01b0395861680825260208083019687525f83850181815260ff968716606086019081526080860198895261ffff909b1680835260978452868320955186549b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909b169a909a1785559751600185015596516002840180549a51909616610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff911515919091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009a8b16171790945593516003909101559183526098905290208054909216179055565b803561ffff81168114611be4575f80fd5b919050565b5f60208284031215611bf9575f80fd5b611c0282611bd3565b9392505050565b80356001600160a01b0381168114611be4575f80fd5b5f60208284031215611c2f575f80fd5b611c0282611c09565b80358015158114611be4575f80fd5b5f805f805f60a08688031215611c5b575f80fd5b611c6486611bd3565b9450611c7260208701611c09565b935060408601359250611c8760608701611c38565b949793965091946080013592915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611d0c57611d0c611c98565b604052919050565b5f67ffffffffffffffff821115611d2d57611d2d611c98565b5060051b60200190565b5f82601f830112611d46575f80fd5b81356020611d5b611d5683611d14565b611cc5565b8083825260208201915060208460051b870101935086841115611d7c575f80fd5b602086015b84811015611d9f57611d9281611bd3565b8352918301918301611d81565b509695505050505050565b5f82601f830112611db9575f80fd5b81356020611dc9611d5683611d14565b8083825260208201915060208460051b870101935086841115611dea575f80fd5b602086015b84811015611d9f5780358352918301918301611def565b5f8060408385031215611e17575f80fd5b823567ffffffffffffffff80821115611e2e575f80fd5b611e3a86838701611d37565b93506020850135915080821115611e4f575f80fd5b50611e5c85828601611daa565b9150509250929050565b5f8083601f840112611e76575f80fd5b50813567ffffffffffffffff811115611e8d575f80fd5b6020830191508360208260051b8501011115611ea7575f80fd5b9250929050565b5f805f8060408587031215611ec1575f80fd5b843567ffffffffffffffff80821115611ed8575f80fd5b611ee488838901611e66565b90965094506020870135915080821115611efc575f80fd5b50611f0987828801611e66565b95989497509550505050565b5f8060408385031215611f26575f80fd5b611f2f83611bd3565b946020939093013593505050565b5f805f8060808587031215611f50575f80fd5b611f5985611bd3565b9350611f6760208601611c09565b93969395505050506040820135916060013590565b5f82601f830112611f8b575f80fd5b81356020611f9b611d5683611d14565b8083825260208201915060208460051b870101935086841115611fbc575f80fd5b602086015b84811015611d9f57611fd281611c09565b8352918301918301611fc1565b5f805f8060808587031215611ff2575f80fd5b843567ffffffffffffffff80821115612009575f80fd5b61201588838901611d37565b955060209150818701358181111561202b575f80fd5b61203789828a01611f7c565b95505060408701358181111561204b575f80fd5b8701601f8101891361205b575f80fd5b8035612069611d5682611d14565b81815260059190911b8201840190848101908b831115612087575f80fd5b928501925b828410156120a55783358252928501929085019061208c565b965050505060608701359150808211156120bd575f80fd5b506120ca87828801611daa565b91505092959194509250565b5f80604083850312156120e7575f80fd5b823567ffffffffffffffff808211156120fe575f80fd5b61210a86838701611f7c565b9350602091508185013581811115612120575f80fd5b85019050601f81018613612132575f80fd5b8035612140611d5682611d14565b81815260059190911b8201830190838101908883111561215e575f80fd5b928401925b828410156121835761217484611c38565b82529284019290840190612163565b80955050505050509250929050565b5f602082840312156121a2575f80fd5b611c0282611c38565b5f602082840312156121bb575f80fd5b815160ff81168114611c02575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417610ba1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f82612267577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b604080825283519082018190525f906020906060840190828701845b828110156122a857815161ffff1684529284019290840190600101612288565b505050838103828501528451808252858301918301905f5b818110156122e55783516001600160a01b0316835292840192918401916001016122c0565b509097965050505050505056fea164736f6c6343000818000a",
}

// L2TokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use L2TokenRegistryMetaData.ABI instead.
var L2TokenRegistryABI = L2TokenRegistryMetaData.ABI

// L2TokenRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2TokenRegistryMetaData.Bin instead.
var L2TokenRegistryBin = L2TokenRegistryMetaData.Bin

// DeployL2TokenRegistry deploys a new Ethereum contract, binding an instance of L2TokenRegistry to it.
func DeployL2TokenRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2TokenRegistry, error) {
	parsed, err := L2TokenRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2TokenRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2TokenRegistry{L2TokenRegistryCaller: L2TokenRegistryCaller{contract: contract}, L2TokenRegistryTransactor: L2TokenRegistryTransactor{contract: contract}, L2TokenRegistryFilterer: L2TokenRegistryFilterer{contract: contract}}, nil
}

// L2TokenRegistry is an auto generated Go binding around an Ethereum contract.
type L2TokenRegistry struct {
	L2TokenRegistryCaller     // Read-only binding to the contract
	L2TokenRegistryTransactor // Write-only binding to the contract
	L2TokenRegistryFilterer   // Log filterer for contract events
}

// L2TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2TokenRegistrySession struct {
	Contract     *L2TokenRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2TokenRegistryCallerSession struct {
	Contract *L2TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TokenRegistryTransactorSession struct {
	Contract     *L2TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2TokenRegistryRaw struct {
	Contract *L2TokenRegistry // Generic contract binding to access the raw methods on
}

// L2TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2TokenRegistryCallerRaw struct {
	Contract *L2TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// L2TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TokenRegistryTransactorRaw struct {
	Contract *L2TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2TokenRegistry creates a new instance of L2TokenRegistry, bound to a specific deployed contract.
func NewL2TokenRegistry(address common.Address, backend bind.ContractBackend) (*L2TokenRegistry, error) {
	contract, err := bindL2TokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistry{L2TokenRegistryCaller: L2TokenRegistryCaller{contract: contract}, L2TokenRegistryTransactor: L2TokenRegistryTransactor{contract: contract}, L2TokenRegistryFilterer: L2TokenRegistryFilterer{contract: contract}}, nil
}

// NewL2TokenRegistryCaller creates a new read-only instance of L2TokenRegistry, bound to a specific deployed contract.
func NewL2TokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*L2TokenRegistryCaller, error) {
	contract, err := bindL2TokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryCaller{contract: contract}, nil
}

// NewL2TokenRegistryTransactor creates a new write-only instance of L2TokenRegistry, bound to a specific deployed contract.
func NewL2TokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*L2TokenRegistryTransactor, error) {
	contract, err := bindL2TokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTransactor{contract: contract}, nil
}

// NewL2TokenRegistryFilterer creates a new log filterer instance of L2TokenRegistry, bound to a specific deployed contract.
func NewL2TokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*L2TokenRegistryFilterer, error) {
	contract, err := bindL2TokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryFilterer{contract: contract}, nil
}

// bindL2TokenRegistry binds a generic wrapper to an already deployed contract.
func bindL2TokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2TokenRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenRegistry *L2TokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenRegistry.Contract.L2TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenRegistry *L2TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.L2TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenRegistry *L2TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.L2TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenRegistry *L2TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenRegistry *L2TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenRegistry *L2TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCaller) AllowList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "allowList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistrySession) AllowList(arg0 common.Address) (bool, error) {
	return _L2TokenRegistry.Contract.AllowList(&_L2TokenRegistry.CallOpts, arg0)
}

// AllowList is a free data retrieval call binding the contract method 0x2848aeaf.
//
// Solidity: function allowList(address ) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) AllowList(arg0 common.Address) (bool, error) {
	return _L2TokenRegistry.Contract.AllowList(&_L2TokenRegistry.CallOpts, arg0)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_L2TokenRegistry *L2TokenRegistrySession) AllowListEnabled() (bool, error) {
	return _L2TokenRegistry.Contract.AllowListEnabled(&_L2TokenRegistry.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) AllowListEnabled() (bool, error) {
	return _L2TokenRegistry.Contract.AllowListEnabled(&_L2TokenRegistry.CallOpts)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_L2TokenRegistry *L2TokenRegistryCaller) CalculateTokenAmount(opts *bind.CallOpts, _tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "calculateTokenAmount", _tokenID, _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_L2TokenRegistry *L2TokenRegistrySession) CalculateTokenAmount(_tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	return _L2TokenRegistry.Contract.CalculateTokenAmount(&_L2TokenRegistry.CallOpts, _tokenID, _ethAmount)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xdddc98be.
//
// Solidity: function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) view returns(uint256 tokenAmount)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) CalculateTokenAmount(_tokenID uint16, _ethAmount *big.Int) (*big.Int, error) {
	return _L2TokenRegistry.Contract.CalculateTokenAmount(&_L2TokenRegistry.CallOpts, _tokenID, _ethAmount)
}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistryCaller) GetTokenIdByAddress(opts *bind.CallOpts, tokenAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "getTokenIdByAddress", tokenAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistrySession) GetTokenIdByAddress(tokenAddress common.Address) (uint16, error) {
	return _L2TokenRegistry.Contract.GetTokenIdByAddress(&_L2TokenRegistry.CallOpts, tokenAddress)
}

// GetTokenIdByAddress is a free data retrieval call binding the contract method 0x8cbab7e4.
//
// Solidity: function getTokenIdByAddress(address tokenAddress) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) GetTokenIdByAddress(tokenAddress common.Address) (uint16, error) {
	return _L2TokenRegistry.Contract.GetTokenIdByAddress(&_L2TokenRegistry.CallOpts, tokenAddress)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_L2TokenRegistry *L2TokenRegistryCaller) GetTokenInfo(opts *bind.CallOpts, _tokenID uint16) (IL2TokenRegistryTokenInfo, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "getTokenInfo", _tokenID)

	if err != nil {
		return *new(IL2TokenRegistryTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IL2TokenRegistryTokenInfo)).(*IL2TokenRegistryTokenInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_L2TokenRegistry *L2TokenRegistrySession) GetTokenInfo(_tokenID uint16) (IL2TokenRegistryTokenInfo, error) {
	return _L2TokenRegistry.Contract.GetTokenInfo(&_L2TokenRegistry.CallOpts, _tokenID)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1c58e793.
//
// Solidity: function getTokenInfo(uint16 _tokenID) view returns((address,bytes32,bool,uint8,uint256))
func (_L2TokenRegistry *L2TokenRegistryCallerSession) GetTokenInfo(_tokenID uint16) (IL2TokenRegistryTokenInfo, error) {
	return _L2TokenRegistry.Contract.GetTokenInfo(&_L2TokenRegistry.CallOpts, _tokenID)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCaller) GetTokenPrice(opts *bind.CallOpts, _tokenID uint16) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "getTokenPrice", _tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistrySession) GetTokenPrice(_tokenID uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.GetTokenPrice(&_L2TokenRegistry.CallOpts, _tokenID)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0x2a1ea5a2.
//
// Solidity: function getTokenPrice(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) GetTokenPrice(_tokenID uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.GetTokenPrice(&_L2TokenRegistry.CallOpts, _tokenID)
}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCaller) GetTokenScale(opts *bind.CallOpts, _tokenID uint16) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "getTokenScale", _tokenID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistrySession) GetTokenScale(_tokenID uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.GetTokenScale(&_L2TokenRegistry.CallOpts, _tokenID)
}

// GetTokenScale is a free data retrieval call binding the contract method 0xa313d007.
//
// Solidity: function getTokenScale(uint16 _tokenID) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) GetTokenScale(_tokenID uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.GetTokenScale(&_L2TokenRegistry.CallOpts, _tokenID)
}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCaller) IsTokenActive(opts *bind.CallOpts, _tokenID uint16) (bool, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "isTokenActive", _tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistrySession) IsTokenActive(_tokenID uint16) (bool, error) {
	return _L2TokenRegistry.Contract.IsTokenActive(&_L2TokenRegistry.CallOpts, _tokenID)
}

// IsTokenActive is a free data retrieval call binding the contract method 0xfce40489.
//
// Solidity: function isTokenActive(uint16 _tokenID) view returns(bool)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) IsTokenActive(_tokenID uint16) (bool, error) {
	return _L2TokenRegistry.Contract.IsTokenActive(&_L2TokenRegistry.CallOpts, _tokenID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TokenRegistry *L2TokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TokenRegistry *L2TokenRegistrySession) Owner() (common.Address, error) {
	return _L2TokenRegistry.Contract.Owner(&_L2TokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) Owner() (common.Address, error) {
	return _L2TokenRegistry.Contract.Owner(&_L2TokenRegistry.CallOpts)
}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCaller) PriceRatio(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "priceRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistrySession) PriceRatio(arg0 uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.PriceRatio(&_L2TokenRegistry.CallOpts, arg0)
}

// PriceRatio is a free data retrieval call binding the contract method 0x19904c33.
//
// Solidity: function priceRatio(uint16 ) view returns(uint256)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) PriceRatio(arg0 uint16) (*big.Int, error) {
	return _L2TokenRegistry.Contract.PriceRatio(&_L2TokenRegistry.CallOpts, arg0)
}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistryCaller) TokenRegistration(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "tokenRegistration", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistrySession) TokenRegistration(arg0 common.Address) (uint16, error) {
	return _L2TokenRegistry.Contract.TokenRegistration(&_L2TokenRegistry.CallOpts, arg0)
}

// TokenRegistration is a free data retrieval call binding the contract method 0x724f91ce.
//
// Solidity: function tokenRegistration(address ) view returns(uint16)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) TokenRegistration(arg0 common.Address) (uint16, error) {
	return _L2TokenRegistry.Contract.TokenRegistration(&_L2TokenRegistry.CallOpts, arg0)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x85519c36.
//
// Solidity: function tokenRegistry(uint16 ) view returns(address tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_L2TokenRegistry *L2TokenRegistryCaller) TokenRegistry(opts *bind.CallOpts, arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	var out []interface{}
	err := _L2TokenRegistry.contract.Call(opts, &out, "tokenRegistry", arg0)

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
func (_L2TokenRegistry *L2TokenRegistrySession) TokenRegistry(arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	return _L2TokenRegistry.Contract.TokenRegistry(&_L2TokenRegistry.CallOpts, arg0)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x85519c36.
//
// Solidity: function tokenRegistry(uint16 ) view returns(address tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_L2TokenRegistry *L2TokenRegistryCallerSession) TokenRegistry(arg0 uint16) (struct {
	TokenAddress common.Address
	BalanceSlot  [32]byte
	IsActive     bool
	Decimals     uint8
	Scale        *big.Int
}, error) {
	return _L2TokenRegistry.Contract.TokenRegistry(&_L2TokenRegistry.CallOpts, arg0)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) BatchUpdatePrices(opts *bind.TransactOpts, _tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "batchUpdatePrices", _tokenIDs, _prices)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) BatchUpdatePrices(_tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.BatchUpdatePrices(&_L2TokenRegistry.TransactOpts, _tokenIDs, _prices)
}

// BatchUpdatePrices is a paid mutator transaction binding the contract method 0x8c399691.
//
// Solidity: function batchUpdatePrices(uint16[] _tokenIDs, uint256[] _prices) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) BatchUpdatePrices(_tokenIDs []uint16, _prices []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.BatchUpdatePrices(&_L2TokenRegistry.TransactOpts, _tokenIDs, _prices)
}

// BatchUpdateTokenStatus is a paid mutator transaction binding the contract method 0x9190202e.
//
// Solidity: function batchUpdateTokenStatus(uint16[] _tokenIDs, bool[] _isActives) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) BatchUpdateTokenStatus(opts *bind.TransactOpts, _tokenIDs []uint16, _isActives []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "batchUpdateTokenStatus", _tokenIDs, _isActives)
}

// BatchUpdateTokenStatus is a paid mutator transaction binding the contract method 0x9190202e.
//
// Solidity: function batchUpdateTokenStatus(uint16[] _tokenIDs, bool[] _isActives) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) BatchUpdateTokenStatus(_tokenIDs []uint16, _isActives []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.BatchUpdateTokenStatus(&_L2TokenRegistry.TransactOpts, _tokenIDs, _isActives)
}

// BatchUpdateTokenStatus is a paid mutator transaction binding the contract method 0x9190202e.
//
// Solidity: function batchUpdateTokenStatus(uint16[] _tokenIDs, bool[] _isActives) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) BatchUpdateTokenStatus(_tokenIDs []uint16, _isActives []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.BatchUpdateTokenStatus(&_L2TokenRegistry.TransactOpts, _tokenIDs, _isActives)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.Initialize(&_L2TokenRegistry.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.Initialize(&_L2TokenRegistry.TransactOpts, owner_)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) RegisterToken(opts *bind.TransactOpts, _tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "registerToken", _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) RegisterToken(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RegisterToken(&_L2TokenRegistry.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xe014d85e.
//
// Solidity: function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) RegisterToken(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RegisterToken(&_L2TokenRegistry.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _scale)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) RegisterTokens(opts *bind.TransactOpts, _tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "registerTokens", _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) RegisterTokens(_tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RegisterTokens(&_L2TokenRegistry.TransactOpts, _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xe2f53f2c.
//
// Solidity: function registerTokens(uint16[] _tokenIDs, address[] _tokenAddresses, bytes32[] _balanceSlots, uint256[] _scales) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) RegisterTokens(_tokenIDs []uint16, _tokenAddresses []common.Address, _balanceSlots [][32]byte, _scales []*big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RegisterTokens(&_L2TokenRegistry.TransactOpts, _tokenIDs, _tokenAddresses, _balanceSlots, _scales)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TokenRegistry *L2TokenRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RenounceOwnership(&_L2TokenRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.RenounceOwnership(&_L2TokenRegistry.TransactOpts)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "setAllowList", user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.SetAllowList(&_L2TokenRegistry.TransactOpts, user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.SetAllowList(&_L2TokenRegistry.TransactOpts, user, val)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "setAllowListEnabled", _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.SetAllowListEnabled(&_L2TokenRegistry.TransactOpts, _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.SetAllowListEnabled(&_L2TokenRegistry.TransactOpts, _allowListEnabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.TransferOwnership(&_L2TokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.TransferOwnership(&_L2TokenRegistry.TransactOpts, newOwner)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) UpdatePriceRatio(opts *bind.TransactOpts, _tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "updatePriceRatio", _tokenID, _newPrice)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) UpdatePriceRatio(_tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdatePriceRatio(&_L2TokenRegistry.TransactOpts, _tokenID, _newPrice)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xb10b69ee.
//
// Solidity: function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) UpdatePriceRatio(_tokenID uint16, _newPrice *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdatePriceRatio(&_L2TokenRegistry.TransactOpts, _tokenID, _newPrice)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) UpdateTokenInfo(opts *bind.TransactOpts, _tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "updateTokenInfo", _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) UpdateTokenInfo(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdateTokenInfo(&_L2TokenRegistry.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x2d59c072.
//
// Solidity: function updateTokenInfo(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, bool _isActive, uint256 _scale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) UpdateTokenInfo(_tokenID uint16, _tokenAddress common.Address, _balanceSlot [32]byte, _isActive bool, _scale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdateTokenInfo(&_L2TokenRegistry.TransactOpts, _tokenID, _tokenAddress, _balanceSlot, _isActive, _scale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactor) UpdateTokenScale(opts *bind.TransactOpts, _tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.contract.Transact(opts, "updateTokenScale", _tokenID, _newScale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_L2TokenRegistry *L2TokenRegistrySession) UpdateTokenScale(_tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdateTokenScale(&_L2TokenRegistry.TransactOpts, _tokenID, _newScale)
}

// UpdateTokenScale is a paid mutator transaction binding the contract method 0xef0fde0f.
//
// Solidity: function updateTokenScale(uint16 _tokenID, uint256 _newScale) returns()
func (_L2TokenRegistry *L2TokenRegistryTransactorSession) UpdateTokenScale(_tokenID uint16, _newScale *big.Int) (*types.Transaction, error) {
	return _L2TokenRegistry.Contract.UpdateTokenScale(&_L2TokenRegistry.TransactOpts, _tokenID, _newScale)
}

// L2TokenRegistryAllowListEnabledUpdatedIterator is returned from FilterAllowListEnabledUpdated and is used to iterate over the raw logs and unpacked data for AllowListEnabledUpdated events raised by the L2TokenRegistry contract.
type L2TokenRegistryAllowListEnabledUpdatedIterator struct {
	Event *L2TokenRegistryAllowListEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryAllowListEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryAllowListEnabledUpdated)
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
		it.Event = new(L2TokenRegistryAllowListEnabledUpdated)
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
func (it *L2TokenRegistryAllowListEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryAllowListEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryAllowListEnabledUpdated represents a AllowListEnabledUpdated event raised by the L2TokenRegistry contract.
type L2TokenRegistryAllowListEnabledUpdated struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowListEnabledUpdated is a free log retrieval operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterAllowListEnabledUpdated(opts *bind.FilterOpts) (*L2TokenRegistryAllowListEnabledUpdatedIterator, error) {

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryAllowListEnabledUpdatedIterator{contract: _L2TokenRegistry.contract, event: "AllowListEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListEnabledUpdated is a free log subscription operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchAllowListEnabledUpdated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryAllowListEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryAllowListEnabledUpdated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseAllowListEnabledUpdated(log types.Log) (*L2TokenRegistryAllowListEnabledUpdated, error) {
	event := new(L2TokenRegistryAllowListEnabledUpdated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryAllowListSetIterator is returned from FilterAllowListSet and is used to iterate over the raw logs and unpacked data for AllowListSet events raised by the L2TokenRegistry contract.
type L2TokenRegistryAllowListSetIterator struct {
	Event *L2TokenRegistryAllowListSet // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryAllowListSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryAllowListSet)
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
		it.Event = new(L2TokenRegistryAllowListSet)
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
func (it *L2TokenRegistryAllowListSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryAllowListSet represents a AllowListSet event raised by the L2TokenRegistry contract.
type L2TokenRegistryAllowListSet struct {
	User common.Address
	Val  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowListSet is a free log retrieval operation binding the contract event 0x6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588.
//
// Solidity: event AllowListSet(address indexed user, bool val)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterAllowListSet(opts *bind.FilterOpts, user []common.Address) (*L2TokenRegistryAllowListSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "AllowListSet", userRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryAllowListSetIterator{contract: _L2TokenRegistry.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

// WatchAllowListSet is a free log subscription operation binding the contract event 0x6dad0aed33f4b7f07095619b668698e17943fd9f4c83e7cfcc7f6dd880a11588.
//
// Solidity: event AllowListSet(address indexed user, bool val)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryAllowListSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "AllowListSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryAllowListSet)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseAllowListSet(log types.Log) (*L2TokenRegistryAllowListSet, error) {
	event := new(L2TokenRegistryAllowListSet)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2TokenRegistry contract.
type L2TokenRegistryInitializedIterator struct {
	Event *L2TokenRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryInitialized)
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
		it.Event = new(L2TokenRegistryInitialized)
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
func (it *L2TokenRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryInitialized represents a Initialized event raised by the L2TokenRegistry contract.
type L2TokenRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2TokenRegistryInitializedIterator, error) {

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryInitializedIterator{contract: _L2TokenRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryInitialized)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseInitialized(log types.Log) (*L2TokenRegistryInitialized, error) {
	event := new(L2TokenRegistryInitialized)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2TokenRegistry contract.
type L2TokenRegistryOwnershipTransferredIterator struct {
	Event *L2TokenRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryOwnershipTransferred)
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
		it.Event = new(L2TokenRegistryOwnershipTransferred)
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
func (it *L2TokenRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the L2TokenRegistry contract.
type L2TokenRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2TokenRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryOwnershipTransferredIterator{contract: _L2TokenRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryOwnershipTransferred)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*L2TokenRegistryOwnershipTransferred, error) {
	event := new(L2TokenRegistryOwnershipTransferred)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryPriceRatioUpdatedIterator is returned from FilterPriceRatioUpdated and is used to iterate over the raw logs and unpacked data for PriceRatioUpdated events raised by the L2TokenRegistry contract.
type L2TokenRegistryPriceRatioUpdatedIterator struct {
	Event *L2TokenRegistryPriceRatioUpdated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryPriceRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryPriceRatioUpdated)
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
		it.Event = new(L2TokenRegistryPriceRatioUpdated)
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
func (it *L2TokenRegistryPriceRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryPriceRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryPriceRatioUpdated represents a PriceRatioUpdated event raised by the L2TokenRegistry contract.
type L2TokenRegistryPriceRatioUpdated struct {
	TokenID  uint16
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceRatioUpdated is a free log retrieval operation binding the contract event 0xd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467.
//
// Solidity: event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterPriceRatioUpdated(opts *bind.FilterOpts, tokenID []uint16) (*L2TokenRegistryPriceRatioUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "PriceRatioUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryPriceRatioUpdatedIterator{contract: _L2TokenRegistry.contract, event: "PriceRatioUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceRatioUpdated is a free log subscription operation binding the contract event 0xd73999ac164146908368455e72209122b67c149b37aab024e2707394a2c70467.
//
// Solidity: event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchPriceRatioUpdated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryPriceRatioUpdated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "PriceRatioUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryPriceRatioUpdated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "PriceRatioUpdated", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParsePriceRatioUpdated(log types.Log) (*L2TokenRegistryPriceRatioUpdated, error) {
	event := new(L2TokenRegistryPriceRatioUpdated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "PriceRatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokenActivatedIterator is returned from FilterTokenActivated and is used to iterate over the raw logs and unpacked data for TokenActivated events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenActivatedIterator struct {
	Event *L2TokenRegistryTokenActivated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokenActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokenActivated)
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
		it.Event = new(L2TokenRegistryTokenActivated)
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
func (it *L2TokenRegistryTokenActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokenActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokenActivated represents a TokenActivated event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenActivated struct {
	TokenID uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenActivated is a free log retrieval operation binding the contract event 0xba78a86bd882b78fb95715a0d827188ec0a8ad3f500310d33a39b94c9ff677b8.
//
// Solidity: event TokenActivated(uint16 indexed tokenID)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokenActivated(opts *bind.FilterOpts, tokenID []uint16) (*L2TokenRegistryTokenActivatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokenActivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokenActivatedIterator{contract: _L2TokenRegistry.contract, event: "TokenActivated", logs: logs, sub: sub}, nil
}

// WatchTokenActivated is a free log subscription operation binding the contract event 0xba78a86bd882b78fb95715a0d827188ec0a8ad3f500310d33a39b94c9ff677b8.
//
// Solidity: event TokenActivated(uint16 indexed tokenID)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokenActivated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokenActivated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokenActivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokenActivated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenActivated", log); err != nil {
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

// ParseTokenActivated is a log parse operation binding the contract event 0xba78a86bd882b78fb95715a0d827188ec0a8ad3f500310d33a39b94c9ff677b8.
//
// Solidity: event TokenActivated(uint16 indexed tokenID)
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokenActivated(log types.Log) (*L2TokenRegistryTokenActivated, error) {
	event := new(L2TokenRegistryTokenActivated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokenDeactivatedIterator is returned from FilterTokenDeactivated and is used to iterate over the raw logs and unpacked data for TokenDeactivated events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenDeactivatedIterator struct {
	Event *L2TokenRegistryTokenDeactivated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokenDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokenDeactivated)
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
		it.Event = new(L2TokenRegistryTokenDeactivated)
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
func (it *L2TokenRegistryTokenDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokenDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokenDeactivated represents a TokenDeactivated event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenDeactivated struct {
	TokenID uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeactivated is a free log retrieval operation binding the contract event 0xa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef.
//
// Solidity: event TokenDeactivated(uint16 indexed tokenID)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokenDeactivated(opts *bind.FilterOpts, tokenID []uint16) (*L2TokenRegistryTokenDeactivatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokenDeactivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokenDeactivatedIterator{contract: _L2TokenRegistry.contract, event: "TokenDeactivated", logs: logs, sub: sub}, nil
}

// WatchTokenDeactivated is a free log subscription operation binding the contract event 0xa625871090c2595895650b8e9222d1a3267cedf9de819bf446400962ce1357ef.
//
// Solidity: event TokenDeactivated(uint16 indexed tokenID)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokenDeactivated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokenDeactivated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokenDeactivated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokenDeactivated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenDeactivated", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokenDeactivated(log types.Log) (*L2TokenRegistryTokenDeactivated, error) {
	event := new(L2TokenRegistryTokenDeactivated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokenInfoUpdatedIterator is returned from FilterTokenInfoUpdated and is used to iterate over the raw logs and unpacked data for TokenInfoUpdated events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenInfoUpdatedIterator struct {
	Event *L2TokenRegistryTokenInfoUpdated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokenInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokenInfoUpdated)
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
		it.Event = new(L2TokenRegistryTokenInfoUpdated)
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
func (it *L2TokenRegistryTokenInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokenInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokenInfoUpdated represents a TokenInfoUpdated event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenInfoUpdated struct {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokenInfoUpdated(opts *bind.FilterOpts, tokenID []uint16, tokenAddress []common.Address) (*L2TokenRegistryTokenInfoUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokenInfoUpdated", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokenInfoUpdatedIterator{contract: _L2TokenRegistry.contract, event: "TokenInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenInfoUpdated is a free log subscription operation binding the contract event 0x60281b1abf645864e8443ca11a3c3b51a6a9203a376da58db7919f7cfebc4aa9.
//
// Solidity: event TokenInfoUpdated(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokenInfoUpdated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokenInfoUpdated, tokenID []uint16, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokenInfoUpdated", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokenInfoUpdated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokenInfoUpdated(log types.Log) (*L2TokenRegistryTokenInfoUpdated, error) {
	event := new(L2TokenRegistryTokenInfoUpdated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenRegisteredIterator struct {
	Event *L2TokenRegistryTokenRegistered // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokenRegistered)
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
		it.Event = new(L2TokenRegistryTokenRegistered)
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
func (it *L2TokenRegistryTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokenRegistered represents a TokenRegistered event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenRegistered struct {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokenRegistered(opts *bind.FilterOpts, tokenID []uint16, tokenAddress []common.Address) (*L2TokenRegistryTokenRegisteredIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokenRegistered", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokenRegisteredIterator{contract: _L2TokenRegistry.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xb9d0acb419ab21384716fbeaa0bcbc172f6347c9bf4fc0614c4e79fc47b36e11.
//
// Solidity: event TokenRegistered(uint16 indexed tokenID, address indexed tokenAddress, bytes32 balanceSlot, bool isActive, uint8 decimals, uint256 scale)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokenRegistered, tokenID []uint16, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokenRegistered", tokenIDRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokenRegistered)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokenRegistered(log types.Log) (*L2TokenRegistryTokenRegistered, error) {
	event := new(L2TokenRegistryTokenRegistered)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokenScaleUpdatedIterator is returned from FilterTokenScaleUpdated and is used to iterate over the raw logs and unpacked data for TokenScaleUpdated events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenScaleUpdatedIterator struct {
	Event *L2TokenRegistryTokenScaleUpdated // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokenScaleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokenScaleUpdated)
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
		it.Event = new(L2TokenRegistryTokenScaleUpdated)
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
func (it *L2TokenRegistryTokenScaleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokenScaleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokenScaleUpdated represents a TokenScaleUpdated event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokenScaleUpdated struct {
	TokenID  uint16
	NewScale *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenScaleUpdated is a free log retrieval operation binding the contract event 0x7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e74.
//
// Solidity: event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokenScaleUpdated(opts *bind.FilterOpts, tokenID []uint16) (*L2TokenRegistryTokenScaleUpdatedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokenScaleUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokenScaleUpdatedIterator{contract: _L2TokenRegistry.contract, event: "TokenScaleUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenScaleUpdated is a free log subscription operation binding the contract event 0x7b614d0c690ae942aec30d9378eb72c3678dd8cb74a55343c87baf8dfe078e74.
//
// Solidity: event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokenScaleUpdated(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokenScaleUpdated, tokenID []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokenScaleUpdated", tokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokenScaleUpdated)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenScaleUpdated", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokenScaleUpdated(log types.Log) (*L2TokenRegistryTokenScaleUpdated, error) {
	event := new(L2TokenRegistryTokenScaleUpdated)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokenScaleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenRegistryTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the L2TokenRegistry contract.
type L2TokenRegistryTokensRegisteredIterator struct {
	Event *L2TokenRegistryTokensRegistered // Event containing the contract specifics and raw log

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
func (it *L2TokenRegistryTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenRegistryTokensRegistered)
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
		it.Event = new(L2TokenRegistryTokensRegistered)
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
func (it *L2TokenRegistryTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenRegistryTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenRegistryTokensRegistered represents a TokensRegistered event raised by the L2TokenRegistry contract.
type L2TokenRegistryTokensRegistered struct {
	TokenIDs       []uint16
	TokenAddresses []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0x31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a.
//
// Solidity: event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses)
func (_L2TokenRegistry *L2TokenRegistryFilterer) FilterTokensRegistered(opts *bind.FilterOpts) (*L2TokenRegistryTokensRegisteredIterator, error) {

	logs, sub, err := _L2TokenRegistry.contract.FilterLogs(opts, "TokensRegistered")
	if err != nil {
		return nil, err
	}
	return &L2TokenRegistryTokensRegisteredIterator{contract: _L2TokenRegistry.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0x31d3859b7231c34728c90804bf84d54252e90f91806a23ede786587a3937922a.
//
// Solidity: event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses)
func (_L2TokenRegistry *L2TokenRegistryFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *L2TokenRegistryTokensRegistered) (event.Subscription, error) {

	logs, sub, err := _L2TokenRegistry.contract.WatchLogs(opts, "TokensRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenRegistryTokensRegistered)
				if err := _L2TokenRegistry.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
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
func (_L2TokenRegistry *L2TokenRegistryFilterer) ParseTokensRegistered(log types.Log) (*L2TokenRegistryTokensRegistered, error) {
	event := new(L2TokenRegistryTokensRegistered)
	if err := _L2TokenRegistry.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
