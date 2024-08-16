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

// MorphStandardERC20FactoryMetaData contains all meta data concerning the MorphStandardERC20Factory contract.
var MorphStandardERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"DeployToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"computeL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"deployL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161074738038061074783398101604081905261002e91610105565b610037336100b6565b6001600160a01b0381166100915760405162461bcd60e51b815260206004820152601b60248201527f7a65726f20696d706c656d656e746174696f6e20616464726573730000000000604482015260640160405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055610132565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215610115575f80fd5b81516001600160a01b038116811461012b575f80fd5b9392505050565b6106088061013f5f395ff3fe608060405234801561000f575f80fd5b506004361061006f575f3560e01c80637bdbcbbf1161004d5780637bdbcbbf146100d95780638da5cb5b146100ec578063f2fde38b14610109575f80fd5b80635c60da1b1461007357806361e98ca1146100bc578063715018a6146100cf575b5f80fd5b6001546100939073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100936100ca3660046105b1565b61011c565b6100d7610158565b005b6100936100e73660046105b1565b61016b565b5f5473ffffffffffffffffffffffffffffffffffffffff16610093565b6100d76101173660046105e2565b61020b565b5f8061012884846102c7565b60015490915061014e9073ffffffffffffffffffffffffffffffffffffffff1682610372565b9150505b92915050565b6101606103d4565b6101695f610454565b565b5f6101746103d4565b5f61017f84846102c7565b6001549091505f906101a79073ffffffffffffffffffffffffffffffffffffffff16836104c8565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f07ab516ad4f19b4465f15fa7c2dbc064f18e734a0846d6e0932da244aa3d8a7160405160405180910390a3949350505050565b6102136103d4565b73ffffffffffffffffffffffffffffffffffffffff81166102bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c481610454565b50565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606083901b1660208201525f9083906034016040516020818303038152906040528051906020012060405160200161035492919060609290921b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000168252601482015260340190565b60405160208183030381529060405280519060200120905092915050565b6040513060388201526f5af43d82803e903d91602b57fd5bf3ff602482015260148101839052733d602d80600a3d3981f3363d3d373d3d3d363d738152605881018290526037600c820120607882015260556043909101205f905b9392505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610169576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b2565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b1760205281603760095ff5905073ffffffffffffffffffffffffffffffffffffffff8116610152576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f455243313136373a2063726561746532206661696c656400000000000000000060448201526064016102b2565b803573ffffffffffffffffffffffffffffffffffffffff811681146105ac575f80fd5b919050565b5f80604083850312156105c2575f80fd5b6105cb83610589565b91506105d960208401610589565b90509250929050565b5f602082840312156105f2575f80fd5b6103cd8261058956fea164736f6c6343000818000a",
}

// MorphStandardERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphStandardERC20FactoryMetaData.ABI instead.
var MorphStandardERC20FactoryABI = MorphStandardERC20FactoryMetaData.ABI

// MorphStandardERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MorphStandardERC20FactoryMetaData.Bin instead.
var MorphStandardERC20FactoryBin = MorphStandardERC20FactoryMetaData.Bin

// DeployMorphStandardERC20Factory deploys a new Ethereum contract, binding an instance of MorphStandardERC20Factory to it.
func DeployMorphStandardERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *MorphStandardERC20Factory, error) {
	parsed, err := MorphStandardERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MorphStandardERC20FactoryBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MorphStandardERC20Factory{MorphStandardERC20FactoryCaller: MorphStandardERC20FactoryCaller{contract: contract}, MorphStandardERC20FactoryTransactor: MorphStandardERC20FactoryTransactor{contract: contract}, MorphStandardERC20FactoryFilterer: MorphStandardERC20FactoryFilterer{contract: contract}}, nil
}

// MorphStandardERC20Factory is an auto generated Go binding around an Ethereum contract.
type MorphStandardERC20Factory struct {
	MorphStandardERC20FactoryCaller     // Read-only binding to the contract
	MorphStandardERC20FactoryTransactor // Write-only binding to the contract
	MorphStandardERC20FactoryFilterer   // Log filterer for contract events
}

// MorphStandardERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphStandardERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphStandardERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphStandardERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphStandardERC20FactorySession struct {
	Contract     *MorphStandardERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MorphStandardERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphStandardERC20FactoryCallerSession struct {
	Contract *MorphStandardERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// MorphStandardERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphStandardERC20FactoryTransactorSession struct {
	Contract     *MorphStandardERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// MorphStandardERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphStandardERC20FactoryRaw struct {
	Contract *MorphStandardERC20Factory // Generic contract binding to access the raw methods on
}

// MorphStandardERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphStandardERC20FactoryCallerRaw struct {
	Contract *MorphStandardERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MorphStandardERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphStandardERC20FactoryTransactorRaw struct {
	Contract *MorphStandardERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphStandardERC20Factory creates a new instance of MorphStandardERC20Factory, bound to a specific deployed contract.
func NewMorphStandardERC20Factory(address common.Address, backend bind.ContractBackend) (*MorphStandardERC20Factory, error) {
	contract, err := bindMorphStandardERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20Factory{MorphStandardERC20FactoryCaller: MorphStandardERC20FactoryCaller{contract: contract}, MorphStandardERC20FactoryTransactor: MorphStandardERC20FactoryTransactor{contract: contract}, MorphStandardERC20FactoryFilterer: MorphStandardERC20FactoryFilterer{contract: contract}}, nil
}

// NewMorphStandardERC20FactoryCaller creates a new read-only instance of MorphStandardERC20Factory, bound to a specific deployed contract.
func NewMorphStandardERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*MorphStandardERC20FactoryCaller, error) {
	contract, err := bindMorphStandardERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20FactoryCaller{contract: contract}, nil
}

// NewMorphStandardERC20FactoryTransactor creates a new write-only instance of MorphStandardERC20Factory, bound to a specific deployed contract.
func NewMorphStandardERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphStandardERC20FactoryTransactor, error) {
	contract, err := bindMorphStandardERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20FactoryTransactor{contract: contract}, nil
}

// NewMorphStandardERC20FactoryFilterer creates a new log filterer instance of MorphStandardERC20Factory, bound to a specific deployed contract.
func NewMorphStandardERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphStandardERC20FactoryFilterer, error) {
	contract, err := bindMorphStandardERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20FactoryFilterer{contract: contract}, nil
}

// bindMorphStandardERC20Factory binds a generic wrapper to an already deployed contract.
func bindMorphStandardERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphStandardERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphStandardERC20Factory.Contract.MorphStandardERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.MorphStandardERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.MorphStandardERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphStandardERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// ComputeL2TokenAddress is a free data retrieval call binding the contract method 0x61e98ca1.
//
// Solidity: function computeL2TokenAddress(address _gateway, address _l1Token) view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCaller) ComputeL2TokenAddress(opts *bind.CallOpts, _gateway common.Address, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _MorphStandardERC20Factory.contract.Call(opts, &out, "computeL2TokenAddress", _gateway, _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeL2TokenAddress is a free data retrieval call binding the contract method 0x61e98ca1.
//
// Solidity: function computeL2TokenAddress(address _gateway, address _l1Token) view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) ComputeL2TokenAddress(_gateway common.Address, _l1Token common.Address) (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.ComputeL2TokenAddress(&_MorphStandardERC20Factory.CallOpts, _gateway, _l1Token)
}

// ComputeL2TokenAddress is a free data retrieval call binding the contract method 0x61e98ca1.
//
// Solidity: function computeL2TokenAddress(address _gateway, address _l1Token) view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCallerSession) ComputeL2TokenAddress(_gateway common.Address, _l1Token common.Address) (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.ComputeL2TokenAddress(&_MorphStandardERC20Factory.CallOpts, _gateway, _l1Token)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphStandardERC20Factory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) Implementation() (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.Implementation(&_MorphStandardERC20Factory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCallerSession) Implementation() (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.Implementation(&_MorphStandardERC20Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphStandardERC20Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) Owner() (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.Owner(&_MorphStandardERC20Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryCallerSession) Owner() (common.Address, error) {
	return _MorphStandardERC20Factory.Contract.Owner(&_MorphStandardERC20Factory.CallOpts)
}

// DeployL2Token is a paid mutator transaction binding the contract method 0x7bdbcbbf.
//
// Solidity: function deployL2Token(address _gateway, address _l1Token) returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactor) DeployL2Token(opts *bind.TransactOpts, _gateway common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.contract.Transact(opts, "deployL2Token", _gateway, _l1Token)
}

// DeployL2Token is a paid mutator transaction binding the contract method 0x7bdbcbbf.
//
// Solidity: function deployL2Token(address _gateway, address _l1Token) returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) DeployL2Token(_gateway common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.DeployL2Token(&_MorphStandardERC20Factory.TransactOpts, _gateway, _l1Token)
}

// DeployL2Token is a paid mutator transaction binding the contract method 0x7bdbcbbf.
//
// Solidity: function deployL2Token(address _gateway, address _l1Token) returns(address)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactorSession) DeployL2Token(_gateway common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.DeployL2Token(&_MorphStandardERC20Factory.TransactOpts, _gateway, _l1Token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.RenounceOwnership(&_MorphStandardERC20Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.RenounceOwnership(&_MorphStandardERC20Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.TransferOwnership(&_MorphStandardERC20Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20Factory.Contract.TransferOwnership(&_MorphStandardERC20Factory.TransactOpts, newOwner)
}

// MorphStandardERC20FactoryDeployTokenIterator is returned from FilterDeployToken and is used to iterate over the raw logs and unpacked data for DeployToken events raised by the MorphStandardERC20Factory contract.
type MorphStandardERC20FactoryDeployTokenIterator struct {
	Event *MorphStandardERC20FactoryDeployToken // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20FactoryDeployTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20FactoryDeployToken)
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
		it.Event = new(MorphStandardERC20FactoryDeployToken)
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
func (it *MorphStandardERC20FactoryDeployTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20FactoryDeployTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20FactoryDeployToken represents a DeployToken event raised by the MorphStandardERC20Factory contract.
type MorphStandardERC20FactoryDeployToken struct {
	L1Token common.Address
	L2Token common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeployToken is a free log retrieval operation binding the contract event 0x07ab516ad4f19b4465f15fa7c2dbc064f18e734a0846d6e0932da244aa3d8a71.
//
// Solidity: event DeployToken(address indexed _l1Token, address indexed _l2Token)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) FilterDeployToken(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address) (*MorphStandardERC20FactoryDeployTokenIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}

	logs, sub, err := _MorphStandardERC20Factory.contract.FilterLogs(opts, "DeployToken", _l1TokenRule, _l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20FactoryDeployTokenIterator{contract: _MorphStandardERC20Factory.contract, event: "DeployToken", logs: logs, sub: sub}, nil
}

// WatchDeployToken is a free log subscription operation binding the contract event 0x07ab516ad4f19b4465f15fa7c2dbc064f18e734a0846d6e0932da244aa3d8a71.
//
// Solidity: event DeployToken(address indexed _l1Token, address indexed _l2Token)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) WatchDeployToken(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20FactoryDeployToken, _l1Token []common.Address, _l2Token []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}

	logs, sub, err := _MorphStandardERC20Factory.contract.WatchLogs(opts, "DeployToken", _l1TokenRule, _l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20FactoryDeployToken)
				if err := _MorphStandardERC20Factory.contract.UnpackLog(event, "DeployToken", log); err != nil {
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

// ParseDeployToken is a log parse operation binding the contract event 0x07ab516ad4f19b4465f15fa7c2dbc064f18e734a0846d6e0932da244aa3d8a71.
//
// Solidity: event DeployToken(address indexed _l1Token, address indexed _l2Token)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) ParseDeployToken(log types.Log) (*MorphStandardERC20FactoryDeployToken, error) {
	event := new(MorphStandardERC20FactoryDeployToken)
	if err := _MorphStandardERC20Factory.contract.UnpackLog(event, "DeployToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphStandardERC20FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MorphStandardERC20Factory contract.
type MorphStandardERC20FactoryOwnershipTransferredIterator struct {
	Event *MorphStandardERC20FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20FactoryOwnershipTransferred)
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
		it.Event = new(MorphStandardERC20FactoryOwnershipTransferred)
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
func (it *MorphStandardERC20FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the MorphStandardERC20Factory contract.
type MorphStandardERC20FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphStandardERC20FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphStandardERC20Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20FactoryOwnershipTransferredIterator{contract: _MorphStandardERC20Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphStandardERC20Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20FactoryOwnershipTransferred)
				if err := _MorphStandardERC20Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MorphStandardERC20Factory *MorphStandardERC20FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MorphStandardERC20FactoryOwnershipTransferred, error) {
	event := new(MorphStandardERC20FactoryOwnershipTransferred)
	if err := _MorphStandardERC20Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
