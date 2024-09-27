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

// L2ReverseCustomGatewayMetaData contains all meta data concerning the L2ReverseCustomGateway contract.
var L2ReverseCustomGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"counterpart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeDepositERC20\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getL1ERC20Address\",\"inputs\":[{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_counterpart\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenMapping\",\"inputs\":[{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC20AndCall\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"FinalizeDepositERC20\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"inputs\":[{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldL1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newL1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawERC20\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61225280620000e75f395ff3fe6080604052600436106100e4575f3560e01c80638da5cb5b11610087578063c676ad2911610057578063c676ad2914610299578063f2fde38b146102b8578063f887ea40146102d7578063fac752eb14610303575f80fd5b80638da5cb5b146101fc578063a93a4af914610226578063ba27f50b14610239578063c0c53b8b1461027a575f80fd5b80636c07ea43116100c25780636c07ea4314610196578063715018a6146101a9578063797594b0146101bd5780638431f5c1146101e9575f80fd5b80633cb747bf146100e857806354bbd59c1461013d578063575361b614610181575b5f80fd5b3480156100f3575f80fd5b506099546101149073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b348015610148575f80fd5b50610114610157366004611ca3565b73ffffffffffffffffffffffffffffffffffffffff9081165f90815260fa60205260409020541690565b61019461018f366004611d0a565b610322565b005b6101946101a4366004611d80565b61036d565b3480156101b4575f80fd5b506101946103ab565b3480156101c8575f80fd5b506097546101149073ffffffffffffffffffffffffffffffffffffffff1681565b6101946101f7366004611db2565b6103be565b348015610207575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610114565b610194610234366004611e44565b6107aa565b348015610244575f80fd5b50610114610253366004611ca3565b60fa6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610285575f80fd5b50610194610294366004611e87565b6107bc565b3480156102a4575f80fd5b506101146102b3366004611ca3565b6109ca565b3480156102c3575f80fd5b506101946102d2366004611ca3565b610a2e565b3480156102e2575f80fd5b506098546101149073ffffffffffffffffffffffffffffffffffffffff1681565b34801561030e575f80fd5b5061019461031d366004611ecf565b610ae5565b61036586868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610bf2915050565b505050505050565b6103a68333845f5b6040519080825280601f01601f19166020018201604052801561039f576020820181803683370190505b5085610bf2565b505050565b6103b3610f53565b6103bc5f610fd4565b565b60995473ffffffffffffffffffffffffffffffffffffffff16338114610445576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561048e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b29190611f33565b60975473ffffffffffffffffffffffffffffffffffffffff908116911614610536576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161043c565b61053e61104a565b34156105a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c7565000000000000000000000000000000604482015260640161043c565b73ffffffffffffffffffffffffffffffffffffffff8816610623576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161043c565b73ffffffffffffffffffffffffffffffffffffffff8088165f90815260fa60205260409020548982169116146106b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3120746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161043c565b6106d673ffffffffffffffffffffffffffffffffffffffff881686866110bd565b6107158584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061119192505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348888888860405161078f9493929190611f4e565b60405180910390a46107a060018055565b5050505050505050565b6107b68484845f610375565b50505050565b5f54610100900460ff16158080156107da57505f54600160ff909116105b806107f35750303b1580156107f357505f5460ff166001145b61087f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161043c565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108db575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610958576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f75746572206164647265737300000000000000000000000000604482015260640161043c565b610963848484611241565b80156107b6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f756e696d706c656d656e7465640000000000000000000000000000000000000060448201525f9060640161043c565b610a36610f53565b73ffffffffffffffffffffffffffffffffffffffff8116610ad9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161043c565b610ae281610fd4565b50565b610aed610f53565b73ffffffffffffffffffffffffffffffffffffffff8116610b6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161043c565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815260fa602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610bfa61104a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260fa60205260409020541680610c88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e00000000000000604482015260640161043c565b5f8411610cf1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e74000000000000000000000000604482015260640161043c565b33610cfd8786866113ec565b60405191975095509091505f90610d229084908a9085908b908b908b90602401612029565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa158015610e06573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e2a9190612083565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b923492610e8f929116905f9088908c9060040161209a565b5f604051808303818588803b158015610ea6575f80fd5b505af1158015610eb8573d5f803e3d5ffd5b50505050508273ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b87604051610f37949392919061209a565b60405180910390a450505050610f4c60018055565b5050505050565b60655473ffffffffffffffffffffffffffffffffffffffff1633146103bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161043c565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6002600154036110b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161043c565b6002600155565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526103a69084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261169d565b5f81511180156111b757505f8273ffffffffffffffffffffffffffffffffffffffff163b115b15611237576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f9061120e9084906004016120df565b5f604051808303815f87803b158015611225575f80fd5b505af1158015610365573d5f803e3d5ffd5b5050565b60018055565b73ffffffffffffffffffffffffffffffffffffffff83166112be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161043c565b73ffffffffffffffffffffffffffffffffffffffff811661133b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161043c565b6113436117aa565b61134b611848565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156103a6576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6098545f9081906060903390819073ffffffffffffffffffffffffffffffffffffffff168190036114d6578580602001905181019061142b91906120f1565b6040517fc52a3bbc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301528b81166024830152604482018b905291985091925083169063c52a3bbc906064016020604051808303815f875af11580156114ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114cf9190612083565b9650611627565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8a16906370a0823190602401602060405180830381865afa158015611540573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115649190612083565b905061158873ffffffffffffffffffffffffffffffffffffffff8a1683308b6118e6565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa1580156115f2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116169190612083565b905061162282826121cd565b985050505b5f8711611690576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e74000000000000000000000000604482015260640161043c565b9795965093949350505050565b5f6116fe826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166119449092919063ffffffff16565b905080515f148061171e57508080602001905181019061171e919061220b565b6103a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161043c565b5f54610100900460ff16611840576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043c565b6103bc61195a565b5f54610100900460ff166118de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043c565b6103bc6119f0565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526107b69085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161110f565b606061195284845f85611a8f565b949350505050565b5f54610100900460ff1661123b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043c565b5f54610100900460ff16611a86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043c565b6103bc33610fd4565b606082471015611b21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161043c565b5f808673ffffffffffffffffffffffffffffffffffffffff168587604051611b49919061222a565b5f6040518083038185875af1925050503d805f8114611b83576040519150601f19603f3d011682016040523d82523d5f602084013e611b88565b606091505b5091509150611b9987838387611ba4565b979650505050505050565b60608315611c395782515f03611c325773ffffffffffffffffffffffffffffffffffffffff85163b611c32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161043c565b5081611952565b6119528383815115611c4e5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043c91906120df565b73ffffffffffffffffffffffffffffffffffffffff81168114610ae2575f80fd5b5f60208284031215611cb3575f80fd5b8135611cbe81611c82565b9392505050565b5f8083601f840112611cd5575f80fd5b50813567ffffffffffffffff811115611cec575f80fd5b602083019150836020828501011115611d03575f80fd5b9250929050565b5f805f805f8060a08789031215611d1f575f80fd5b8635611d2a81611c82565b95506020870135611d3a81611c82565b945060408701359350606087013567ffffffffffffffff811115611d5c575f80fd5b611d6889828a01611cc5565b979a9699509497949695608090950135949350505050565b5f805f60608486031215611d92575f80fd5b8335611d9d81611c82565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a031215611dc8575f80fd5b8735611dd381611c82565b96506020880135611de381611c82565b95506040880135611df381611c82565b94506060880135611e0381611c82565b93506080880135925060a088013567ffffffffffffffff811115611e25575f80fd5b611e318a828b01611cc5565b989b979a50959850939692959293505050565b5f805f8060808587031215611e57575f80fd5b8435611e6281611c82565b93506020850135611e7281611c82565b93969395505050506040820135916060013590565b5f805f60608486031215611e99575f80fd5b8335611ea481611c82565b92506020840135611eb481611c82565b91506040840135611ec481611c82565b809150509250925092565b5f8060408385031215611ee0575f80fd5b8235611eeb81611c82565b91506020830135611efb81611c82565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215611f43575f80fd5b8151611cbe81611c82565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611fd8578181015183820152602001611fc0565b50505f910152565b5f8151808452611ff7816020860160208601611fbe565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261207760c0830184611fe0565b98975050505050505050565b5f60208284031215612093575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f6120ce6080830185611fe0565b905082606083015295945050505050565b602081525f611cbe6020830184611fe0565b5f8060408385031215612102575f80fd5b825161210d81611c82565b602084015190925067ffffffffffffffff8082111561212a575f80fd5b818501915085601f83011261213d575f80fd5b81518181111561214f5761214f611f06565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561219557612195611f06565b816040528281528860208487010111156121ad575f80fd5b6121be836020830160208801611fbe565b80955050505050509250929050565b81810381811115612205577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b5f6020828403121561221b575f80fd5b81518015158114611cbe575f80fd5b5f825161223b818460208701611fbe565b919091019291505056fea164736f6c6343000818000a",
}

// L2ReverseCustomGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ReverseCustomGatewayMetaData.ABI instead.
var L2ReverseCustomGatewayABI = L2ReverseCustomGatewayMetaData.ABI

// L2ReverseCustomGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ReverseCustomGatewayMetaData.Bin instead.
var L2ReverseCustomGatewayBin = L2ReverseCustomGatewayMetaData.Bin

// DeployL2ReverseCustomGateway deploys a new Ethereum contract, binding an instance of L2ReverseCustomGateway to it.
func DeployL2ReverseCustomGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ReverseCustomGateway, error) {
	parsed, err := L2ReverseCustomGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ReverseCustomGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ReverseCustomGateway{L2ReverseCustomGatewayCaller: L2ReverseCustomGatewayCaller{contract: contract}, L2ReverseCustomGatewayTransactor: L2ReverseCustomGatewayTransactor{contract: contract}, L2ReverseCustomGatewayFilterer: L2ReverseCustomGatewayFilterer{contract: contract}}, nil
}

// L2ReverseCustomGateway is an auto generated Go binding around an Ethereum contract.
type L2ReverseCustomGateway struct {
	L2ReverseCustomGatewayCaller     // Read-only binding to the contract
	L2ReverseCustomGatewayTransactor // Write-only binding to the contract
	L2ReverseCustomGatewayFilterer   // Log filterer for contract events
}

// L2ReverseCustomGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ReverseCustomGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ReverseCustomGatewaySession struct {
	Contract     *L2ReverseCustomGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2ReverseCustomGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ReverseCustomGatewayCallerSession struct {
	Contract *L2ReverseCustomGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2ReverseCustomGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ReverseCustomGatewayTransactorSession struct {
	Contract     *L2ReverseCustomGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2ReverseCustomGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ReverseCustomGatewayRaw struct {
	Contract *L2ReverseCustomGateway // Generic contract binding to access the raw methods on
}

// L2ReverseCustomGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayCallerRaw struct {
	Contract *L2ReverseCustomGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2ReverseCustomGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayTransactorRaw struct {
	Contract *L2ReverseCustomGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ReverseCustomGateway creates a new instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGateway(address common.Address, backend bind.ContractBackend) (*L2ReverseCustomGateway, error) {
	contract, err := bindL2ReverseCustomGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGateway{L2ReverseCustomGatewayCaller: L2ReverseCustomGatewayCaller{contract: contract}, L2ReverseCustomGatewayTransactor: L2ReverseCustomGatewayTransactor{contract: contract}, L2ReverseCustomGatewayFilterer: L2ReverseCustomGatewayFilterer{contract: contract}}, nil
}

// NewL2ReverseCustomGatewayCaller creates a new read-only instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ReverseCustomGatewayCaller, error) {
	contract, err := bindL2ReverseCustomGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayCaller{contract: contract}, nil
}

// NewL2ReverseCustomGatewayTransactor creates a new write-only instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ReverseCustomGatewayTransactor, error) {
	contract, err := bindL2ReverseCustomGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayTransactor{contract: contract}, nil
}

// NewL2ReverseCustomGatewayFilterer creates a new log filterer instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ReverseCustomGatewayFilterer, error) {
	contract, err := bindL2ReverseCustomGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayFilterer{contract: contract}, nil
}

// bindL2ReverseCustomGateway binds a generic wrapper to an already deployed contract.
func bindL2ReverseCustomGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ReverseCustomGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ReverseCustomGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Counterpart() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Counterpart(&_L2ReverseCustomGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Counterpart(&_L2ReverseCustomGateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.GetL1ERC20Address(&_L2ReverseCustomGateway.CallOpts, _l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.GetL1ERC20Address(&_L2ReverseCustomGateway.CallOpts, _l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.GetL2ERC20Address(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.GetL2ERC20Address(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Messenger() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Messenger(&_L2ReverseCustomGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Messenger(&_L2ReverseCustomGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Owner() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Owner(&_L2ReverseCustomGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) Owner() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Owner(&_L2ReverseCustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Router() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Router(&_L2ReverseCustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) Router() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Router(&_L2ReverseCustomGateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.TokenMapping(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.TokenMapping(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.FinalizeDepositERC20(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.FinalizeDepositERC20(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.Initialize(&_L2ReverseCustomGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.Initialize(&_L2ReverseCustomGateway.TransactOpts, _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.RenounceOwnership(&_L2ReverseCustomGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.RenounceOwnership(&_L2ReverseCustomGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.TransferOwnership(&_L2ReverseCustomGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.TransferOwnership(&_L2ReverseCustomGateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.UpdateTokenMapping(&_L2ReverseCustomGateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.UpdateTokenMapping(&_L2ReverseCustomGateway.TransactOpts, _l2Token, _l1Token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC20(&_L2ReverseCustomGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC20(&_L2ReverseCustomGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC200(&_L2ReverseCustomGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC200(&_L2ReverseCustomGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC20AndCall(&_L2ReverseCustomGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.WithdrawERC20AndCall(&_L2ReverseCustomGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2ReverseCustomGatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayFinalizeDepositERC20Iterator struct {
	Event *L2ReverseCustomGatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayFinalizeDepositERC20)
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
		it.Event = new(L2ReverseCustomGatewayFinalizeDepositERC20)
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
func (it *L2ReverseCustomGatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayFinalizeDepositERC20 struct {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ReverseCustomGatewayFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayFinalizeDepositERC20Iterator{contract: _L2ReverseCustomGateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayFinalizeDepositERC20)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2ReverseCustomGatewayFinalizeDepositERC20, error) {
	event := new(L2ReverseCustomGatewayFinalizeDepositERC20)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayInitializedIterator struct {
	Event *L2ReverseCustomGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayInitialized)
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
		it.Event = new(L2ReverseCustomGatewayInitialized)
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
func (it *L2ReverseCustomGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayInitialized represents a Initialized event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ReverseCustomGatewayInitializedIterator, error) {

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayInitializedIterator{contract: _L2ReverseCustomGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayInitialized)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseInitialized(log types.Log) (*L2ReverseCustomGatewayInitialized, error) {
	event := new(L2ReverseCustomGatewayInitialized)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayOwnershipTransferredIterator struct {
	Event *L2ReverseCustomGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayOwnershipTransferred)
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
		it.Event = new(L2ReverseCustomGatewayOwnershipTransferred)
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
func (it *L2ReverseCustomGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2ReverseCustomGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayOwnershipTransferredIterator{contract: _L2ReverseCustomGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayOwnershipTransferred)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2ReverseCustomGatewayOwnershipTransferred, error) {
	event := new(L2ReverseCustomGatewayOwnershipTransferred)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayUpdateTokenMappingIterator struct {
	Event *L2ReverseCustomGatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayUpdateTokenMapping)
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
		it.Event = new(L2ReverseCustomGatewayUpdateTokenMapping)
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
func (it *L2ReverseCustomGatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2ReverseCustomGatewayUpdateTokenMappingIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayUpdateTokenMappingIterator{contract: _L2ReverseCustomGateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayUpdateTokenMapping)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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

// ParseUpdateTokenMapping is a log parse operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2ReverseCustomGatewayUpdateTokenMapping, error) {
	event := new(L2ReverseCustomGatewayUpdateTokenMapping)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayWithdrawERC20Iterator struct {
	Event *L2ReverseCustomGatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayWithdrawERC20)
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
		it.Event = new(L2ReverseCustomGatewayWithdrawERC20)
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
func (it *L2ReverseCustomGatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayWithdrawERC20 struct {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ReverseCustomGatewayWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayWithdrawERC20Iterator{contract: _L2ReverseCustomGateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayWithdrawERC20)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2ReverseCustomGatewayWithdrawERC20, error) {
	event := new(L2ReverseCustomGatewayWithdrawERC20)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
