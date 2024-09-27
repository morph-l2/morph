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

// L2GatewayRouterMetaData contains all meta data concerning the L2GatewayRouter contract.
var L2GatewayRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ERC20Gateway\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultERC20Gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethGateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeDepositERC20\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeDepositETH\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gatewayInContext\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getERC20Gateway\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL1ERC20Address\",\"inputs\":[{\"name\":\"_l2Address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ethGateway\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_defaultERC20Gateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestERC20\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultERC20Gateway\",\"inputs\":[{\"name\":\"_newDefaultERC20Gateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setERC20Gateway\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_gateways\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setETHGateway\",\"inputs\":[{\"name\":\"_newEthGateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC20AndCall\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETHAndCall\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"FinalizeDepositERC20\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeDepositETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDefaultERC20Gateway\",\"inputs\":[{\"name\":\"oldDefaultERC20Gateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newDefaultERC20Gateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetERC20Gateway\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldGateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetETHGateway\",\"inputs\":[{\"name\":\"oldETHGateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEthGateway\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawERC20\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61213c80620000e75f395ff3fe608060405260043610610178575f3560e01c80636dc24183116100d1578063a93a4af91161007c578063c7cdea3711610057578063c7cdea3714610405578063ce8c3e0614610418578063f2fde38b14610444575f80fd5b8063a93a4af9146103a6578063c52a3bbc146103b9578063c676ad29146103e6575f80fd5b80638431f5c1116100ac5780638431f5c1146103425780638c00ce73146103505780638da5cb5b1461037c575f80fd5b80636dc24183146102da578063705b05b8146102ed578063715018a61461032e575f80fd5b8063485cc955116101315780635dfd5b9a1161010c5780635dfd5b9a14610289578063635c8637146102a85780636c07ea43146102c7575f80fd5b8063485cc9551461023857806354bbd59c14610257578063575361b614610276575f80fd5b80633a9a7b20116101615780633a9a7b20146101a45780633d1d31c7146101fa57806343c6674114610219575f80fd5b8063232e87481461017c5780632fcc29fa14610191575b5f80fd5b61018f61018a366004611a4c565b610463565b005b61018f61019f366004611aba565b6104ca565b3480156101af575f80fd5b506068546101d09073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610205575f80fd5b5061018f610214366004611aec565b610507565b348015610224575f80fd5b506101d0610233366004611aec565b610585565b348015610243575f80fd5b5061018f610252366004611b07565b6105d4565b348015610262575f80fd5b506101d0610271366004611aec565b610875565b61018f610284366004611c44565b61093a565b348015610294575f80fd5b5061018f6102a3366004611aec565b610b80565b3480156102b3575f80fd5b5061018f6102c2366004611d2e565b610bfe565b61018f6102d5366004611aba565b610e4e565b61018f6102e8366004611d8e565b610e87565b3480156102f8575f80fd5b506101d0610307366004611aec565b60676020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610339575f80fd5b5061018f6110c1565b61018f61018a366004611deb565b34801561035b575f80fd5b506065546101d09073ffffffffffffffffffffffffffffffffffffffff1681565b348015610387575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101d0565b61018f6103b4366004611e7d565b6110d4565b3480156103c4575f80fd5b506103d86103d3366004611ec0565b6110e6565b6040519081526020016101f1565b3480156103f1575f80fd5b506101d0610400366004611aec565b6112dd565b61018f610413366004611efe565b611341565b348015610423575f80fd5b506066546101d09073ffffffffffffffffffffffffffffffffffffffff1681565b34801561044f575f80fd5b5061018f61045e366004611aec565b611350565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f73686f756c64206e657665722062652063616c6c65640000000000000000000060448201526064015b60405180910390fd5b61050283835f5b6040519080825280601f01601f1916602001820160405280156104fb576020820181803683370190505b5084610e87565b505050565b61050f611407565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff8082165f90815260676020526040812054909116806105ce575060665473ffffffffffffffffffffffffffffffffffffffff165b92915050565b5f54610100900460ff16158080156105f257505f54600160ff909116105b8061060b5750303b15801561060b57505f5460ff166001145b610697576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c1565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156106f3575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6106fb611488565b73ffffffffffffffffffffffffffffffffffffffff82161561078557606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040515f907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1908290a35b73ffffffffffffffffffffffffffffffffffffffff83161561080f57606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040515f907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a908290a35b8015610502575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b5f8061088083610585565b905073ffffffffffffffffffffffffffffffffffffffff81166108a557505f92915050565b6040517f54bbd59c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301528216906354bbd59c90602401602060405180830381865afa15801561090f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109339190611f1e565b9392505050565b60685473ffffffffffffffffffffffffffffffffffffffff16156109ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e746578740000000000000000000000000060448201526064016104c1565b5f6109c486610585565b905073ffffffffffffffffffffffffffffffffffffffff8116610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f206761746577617920617661696c61626c6500000000000000000000000060448201526064016104c1565b606880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555f3384604051602001610a97929190611fa4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f575361b6000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff83169063575361b6903490610b21908b908b908b9088908b90600401611fd2565b5f604051808303818588803b158015610b38575f80fd5b505af1158015610b4a573d5f803e3d5ffd5b5050606880547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550505050505050505050565b610b88611407565b6066805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1905f90a35050565b610c06611407565b8051825114610c71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064016104c1565b5f5b8251811015610502575f60675f858481518110610c9257610c92612022565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828281518110610d0757610d07612022565b602002602001015160675f868581518110610d2457610d24612022565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828281518110610db457610db4612022565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16858481518110610dfb57610dfb612022565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf60405160405180910390a450600101610c73565b6105028333845f5b6040519080825280601f01601f191660200182016040528015610e80576020820181803683370190505b508561093a565b60685473ffffffffffffffffffffffffffffffffffffffff1615610f07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e746578740000000000000000000000000060448201526064016104c1565b60655473ffffffffffffffffffffffffffffffffffffffff1680610f87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f657468206761746577617920617661696c61626c65000000000000000000000060448201526064016104c1565b606880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555f3384604051602001610fdb929190611fa4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f6dc24183000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff831690636dc24183903490611063908a908a9087908a9060040161204f565b5f604051808303818588803b15801561107a575f80fd5b505af115801561108c573d5f803e3d5ffd5b5050606880547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055505050505050505050565b6110c9611407565b6110d25f611526565b565b6110e08484845f610e56565b50505050565b6068545f9073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461117f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4f6e6c7920696e206465706f73697420636f6e7465787400000000000000000060448201526064016104c1565b5f336040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192505f918616906370a0823190602401602060405180830381865afa1580156111f0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112149190612094565b905061123873ffffffffffffffffffffffffffffffffffffffff861687848761159c565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301528291908716906370a0823190602401602060405180830381865afa1580156112a5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112c99190612094565b6112d391906120ab565b9695505050505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f756e737570706f7274656400000000000000000000000000000000000000000060448201525f906064016104c1565b61134c33835f6104d1565b5050565b611358611407565b73ffffffffffffffffffffffffffffffffffffffff81166113fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104c1565b61140481611526565b50565b60335473ffffffffffffffffffffffffffffffffffffffff1633146110d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c1565b5f54610100900460ff1661151e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c1565b6110d2611631565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526110e09085906116d0565b5f54610100900460ff166116c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c1565b6110d233611526565b5f611731826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166117dd9092919063ffffffff16565b905080515f148061175157508080602001905181019061175191906120e3565b610502576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104c1565b60606117eb84845f856117f3565b949350505050565b606082471015611885576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016104c1565b5f808673ffffffffffffffffffffffffffffffffffffffff1685876040516118ad9190612102565b5f6040518083038185875af1925050503d805f81146118e7576040519150601f19603f3d011682016040523d82523d5f602084013e6118ec565b606091505b50915091506118fd87838387611908565b979650505050505050565b6060831561199d5782515f036119965773ffffffffffffffffffffffffffffffffffffffff85163b611996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104c1565b50816117eb565b6117eb83838151156119b25781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c1919061211d565b73ffffffffffffffffffffffffffffffffffffffff81168114611404575f80fd5b5f8083601f840112611a17575f80fd5b50813567ffffffffffffffff811115611a2e575f80fd5b602083019150836020828501011115611a45575f80fd5b9250929050565b5f805f805f60808688031215611a60575f80fd5b8535611a6b816119e6565b94506020860135611a7b816119e6565b935060408601359250606086013567ffffffffffffffff811115611a9d575f80fd5b611aa988828901611a07565b969995985093965092949392505050565b5f805f60608486031215611acc575f80fd5b8335611ad7816119e6565b95602085013595506040909401359392505050565b5f60208284031215611afc575f80fd5b8135610933816119e6565b5f8060408385031215611b18575f80fd5b8235611b23816119e6565b91506020830135611b33816119e6565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611bb257611bb2611b3e565b604052919050565b5f82601f830112611bc9575f80fd5b813567ffffffffffffffff811115611be357611be3611b3e565b611c1460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611b6b565b818152846020838601011115611c28575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611c58575f80fd5b8535611c63816119e6565b94506020860135611c73816119e6565b935060408601359250606086013567ffffffffffffffff811115611c95575f80fd5b611ca188828901611bba565b95989497509295608001359392505050565b5f82601f830112611cc2575f80fd5b8135602067ffffffffffffffff821115611cde57611cde611b3e565b8160051b611ced828201611b6b565b9283528481018201928281019087851115611d06575f80fd5b83870192505b848310156118fd578235611d1f816119e6565b82529183019190830190611d0c565b5f8060408385031215611d3f575f80fd5b823567ffffffffffffffff80821115611d56575f80fd5b611d6286838701611cb3565b93506020850135915080821115611d77575f80fd5b50611d8485828601611cb3565b9150509250929050565b5f805f8060808587031215611da1575f80fd5b8435611dac816119e6565b935060208501359250604085013567ffffffffffffffff811115611dce575f80fd5b611dda87828801611bba565b949793965093946060013593505050565b5f805f805f805f60c0888a031215611e01575f80fd5b8735611e0c816119e6565b96506020880135611e1c816119e6565b95506040880135611e2c816119e6565b94506060880135611e3c816119e6565b93506080880135925060a088013567ffffffffffffffff811115611e5e575f80fd5b611e6a8a828b01611a07565b989b979a50959850939692959293505050565b5f805f8060808587031215611e90575f80fd5b8435611e9b816119e6565b93506020850135611eab816119e6565b93969395505050506040820135916060013590565b5f805f60608486031215611ed2575f80fd5b8335611edd816119e6565b92506020840135611eed816119e6565b929592945050506040919091013590565b5f8060408385031215611f0f575f80fd5b50508035926020909101359150565b5f60208284031215611f2e575f80fd5b8151610933816119e6565b5f5b83811015611f53578181015183820152602001611f3b565b50505f910152565b5f8151808452611f72816020860160208601611f39565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6117eb6040830184611f5b565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015260a0606083015261201060a0830185611f5b565b90508260808301529695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f6120836080830185611f5b565b905082606083015295945050505050565b5f602082840312156120a4575f80fd5b5051919050565b818103818111156105ce577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f602082840312156120f3575f80fd5b81518015158114610933575f80fd5b5f8251612113818460208701611f39565b9190910192915050565b602081525f6109336020830184611f5b56fea164736f6c6343000818000a",
}

// L2GatewayRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use L2GatewayRouterMetaData.ABI instead.
var L2GatewayRouterABI = L2GatewayRouterMetaData.ABI

// L2GatewayRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2GatewayRouterMetaData.Bin instead.
var L2GatewayRouterBin = L2GatewayRouterMetaData.Bin

// DeployL2GatewayRouter deploys a new Ethereum contract, binding an instance of L2GatewayRouter to it.
func DeployL2GatewayRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2GatewayRouter, error) {
	parsed, err := L2GatewayRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2GatewayRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2GatewayRouter{L2GatewayRouterCaller: L2GatewayRouterCaller{contract: contract}, L2GatewayRouterTransactor: L2GatewayRouterTransactor{contract: contract}, L2GatewayRouterFilterer: L2GatewayRouterFilterer{contract: contract}}, nil
}

// L2GatewayRouter is an auto generated Go binding around an Ethereum contract.
type L2GatewayRouter struct {
	L2GatewayRouterCaller     // Read-only binding to the contract
	L2GatewayRouterTransactor // Write-only binding to the contract
	L2GatewayRouterFilterer   // Log filterer for contract events
}

// L2GatewayRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2GatewayRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2GatewayRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2GatewayRouterSession struct {
	Contract     *L2GatewayRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2GatewayRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2GatewayRouterCallerSession struct {
	Contract *L2GatewayRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2GatewayRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2GatewayRouterTransactorSession struct {
	Contract     *L2GatewayRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2GatewayRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2GatewayRouterRaw struct {
	Contract *L2GatewayRouter // Generic contract binding to access the raw methods on
}

// L2GatewayRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2GatewayRouterCallerRaw struct {
	Contract *L2GatewayRouterCaller // Generic read-only contract binding to access the raw methods on
}

// L2GatewayRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactorRaw struct {
	Contract *L2GatewayRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2GatewayRouter creates a new instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouter(address common.Address, backend bind.ContractBackend) (*L2GatewayRouter, error) {
	contract, err := bindL2GatewayRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouter{L2GatewayRouterCaller: L2GatewayRouterCaller{contract: contract}, L2GatewayRouterTransactor: L2GatewayRouterTransactor{contract: contract}, L2GatewayRouterFilterer: L2GatewayRouterFilterer{contract: contract}}, nil
}

// NewL2GatewayRouterCaller creates a new read-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterCaller(address common.Address, caller bind.ContractCaller) (*L2GatewayRouterCaller, error) {
	contract, err := bindL2GatewayRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterCaller{contract: contract}, nil
}

// NewL2GatewayRouterTransactor creates a new write-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*L2GatewayRouterTransactor, error) {
	contract, err := bindL2GatewayRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterTransactor{contract: contract}, nil
}

// NewL2GatewayRouterFilterer creates a new log filterer instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*L2GatewayRouterFilterer, error) {
	contract, err := bindL2GatewayRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFilterer{contract: contract}, nil
}

// bindL2GatewayRouter binds a generic wrapper to an already deployed contract.
func bindL2GatewayRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2GatewayRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.L2GatewayRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GatewayInContext(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "gatewayInContext")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GatewayInContext() (common.Address, error) {
	return _L2GatewayRouter.Contract.GatewayInContext(&_L2GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GatewayInContext() (common.Address, error) {
	return _L2GatewayRouter.Contract.GatewayInContext(&_L2GatewayRouter.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL1ERC20Address", _l2Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositETH", arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterTransactor) RequestERC20(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "requestERC20", _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RequestERC20(&_L2GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RequestERC20(&_L2GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setDefaultERC20Gateway", _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetETHGateway(opts *bind.TransactOpts, _newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setETHGateway", _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// L2GatewayRouterFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20Iterator struct {
	Event *L2GatewayRouterFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
		it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositERC20Iterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2GatewayRouterFinalizeDepositERC20, error) {
	event := new(L2GatewayRouterFinalizeDepositERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETHIterator struct {
	Event *L2GatewayRouterFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
		it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositETHIterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositETH(log types.Log) (*L2GatewayRouterFinalizeDepositETH, error) {
	event := new(L2GatewayRouterFinalizeDepositETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2GatewayRouter contract.
type L2GatewayRouterInitializedIterator struct {
	Event *L2GatewayRouterInitialized // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterInitialized)
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
		it.Event = new(L2GatewayRouterInitialized)
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
func (it *L2GatewayRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterInitialized represents a Initialized event raised by the L2GatewayRouter contract.
type L2GatewayRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2GatewayRouterInitializedIterator, error) {

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterInitializedIterator{contract: _L2GatewayRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterInitialized)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseInitialized(log types.Log) (*L2GatewayRouterInitialized, error) {
	event := new(L2GatewayRouterInitialized)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferredIterator struct {
	Event *L2GatewayRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterOwnershipTransferred)
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
		it.Event = new(L2GatewayRouterOwnershipTransferred)
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
func (it *L2GatewayRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterOwnershipTransferred represents a OwnershipTransferred event raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2GatewayRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterOwnershipTransferredIterator{contract: _L2GatewayRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterOwnershipTransferred)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseOwnershipTransferred(log types.Log) (*L2GatewayRouterOwnershipTransferred, error) {
	event := new(L2GatewayRouterOwnershipTransferred)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20GatewayIterator struct {
	Event *L2GatewayRouterSetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20Gateway struct {
	OldDefaultERC20Gateway common.Address
	NewDefaultERC20Gateway common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (*L2GatewayRouterSetDefaultERC20GatewayIterator, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetDefaultERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetDefaultERC20Gateway, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (event.Subscription, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetDefaultERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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

// ParseSetDefaultERC20Gateway is a log parse operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*L2GatewayRouterSetDefaultERC20Gateway, error) {
	event := new(L2GatewayRouterSetDefaultERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20GatewayIterator struct {
	Event *L2GatewayRouterSetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetERC20Gateway)
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
func (it *L2GatewayRouterSetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetERC20Gateway represents a SetERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20Gateway struct {
	Token      common.Address
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (*L2GatewayRouterSetERC20GatewayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetERC20Gateway, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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

// ParseSetERC20Gateway is a log parse operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetERC20Gateway(log types.Log) (*L2GatewayRouterSetERC20Gateway, error) {
	event := new(L2GatewayRouterSetERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGatewayIterator struct {
	Event *L2GatewayRouterSetETHGateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetETHGateway)
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
		it.Event = new(L2GatewayRouterSetETHGateway)
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
func (it *L2GatewayRouterSetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetETHGateway represents a SetETHGateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGateway struct {
	OldETHGateway common.Address
	NewEthGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetETHGateway(opts *bind.FilterOpts, oldETHGateway []common.Address, newEthGateway []common.Address) (*L2GatewayRouterSetETHGatewayIterator, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetETHGatewayIterator{contract: _L2GatewayRouter.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetETHGateway, oldETHGateway []common.Address, newEthGateway []common.Address) (event.Subscription, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetETHGateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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

// ParseSetETHGateway is a log parse operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetETHGateway(log types.Log) (*L2GatewayRouterSetETHGateway, error) {
	event := new(L2GatewayRouterSetETHGateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20Iterator struct {
	Event *L2GatewayRouterWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawERC20)
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
		it.Event = new(L2GatewayRouterWithdrawERC20)
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
func (it *L2GatewayRouterWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawERC20 represents a WithdrawERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawERC20Iterator{contract: _L2GatewayRouter.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawERC20(log types.Log) (*L2GatewayRouterWithdrawERC20, error) {
	event := new(L2GatewayRouterWithdrawERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETHIterator struct {
	Event *L2GatewayRouterWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawETH)
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
		it.Event = new(L2GatewayRouterWithdrawETH)
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
func (it *L2GatewayRouterWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawETH represents a WithdrawETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawETHIterator{contract: _L2GatewayRouter.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawETH(log types.Log) (*L2GatewayRouterWithdrawETH, error) {
	event := new(L2GatewayRouterWithdrawETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
