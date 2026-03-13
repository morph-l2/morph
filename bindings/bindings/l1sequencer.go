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

// L1SequencerMetaData contains all meta data concerning the L1Sequencer contract.
var L1SequencerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getSequencer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sequencer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSequencer\",\"inputs\":[{\"name\":\"newSequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerUpdated\",\"inputs\":[{\"name\":\"oldSequencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSequencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5061081a8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c8063715018a611610058578063715018a6146100f65780638da5cb5b146100fe578063c4d66de81461011c578063f2fde38b1461012f575f80fd5b806343ae20a31461007e5780634d96a90a146100935780635c1bba38146100d6575b5f80fd5b61009161008c3660046107d3565b610142565b005b60655473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6065546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6100916102c7565b60335473ffffffffffffffffffffffffffffffffffffffff166100ad565b61009161012a3660046107d3565b6102da565b61009161013d3660046107d3565b6104ed565b61014a6105a4565b73ffffffffffffffffffffffffffffffffffffffff81166101cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c69642073657175656e63657200000000000000000000000000000060448201526064015b60405180910390fd5b60655473ffffffffffffffffffffffffffffffffffffffff90811690821603610251576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f73616d652073657175656e63657200000000000000000000000000000000000060448201526064016101c3565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fcd58b762453bd126b48db83f2cecd464f5281dd7e5e6824b528c09d0482984d6905f90a35050565b6102cf6105a4565b6102d85f610625565b565b5f54610100900460ff16158080156102f857505f54600160ff909116105b806103115750303b15801561031157505f5460ff166001145b61039d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016101c3565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103f9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8216610476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f696e76616c6964206f776e65720000000000000000000000000000000000000060448201526064016101c3565b61047e61069b565b61048782610625565b80156104e9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6104f56105a4565b73ffffffffffffffffffffffffffffffffffffffff8116610598576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101c3565b6105a181610625565b50565b60335473ffffffffffffffffffffffffffffffffffffffff1633146102d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101c3565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610731576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101c3565b6102d85f54610100900460ff166107ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101c3565b6102d833610625565b5f602082840312156107e3575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610806575f80fd5b939250505056fea164736f6c6343000818000a",
}

// L1SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1SequencerMetaData.ABI instead.
var L1SequencerABI = L1SequencerMetaData.ABI

// L1SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1SequencerMetaData.Bin instead.
var L1SequencerBin = L1SequencerMetaData.Bin

// DeployL1Sequencer deploys a new Ethereum contract, binding an instance of L1Sequencer to it.
func DeployL1Sequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1Sequencer, error) {
	parsed, err := L1SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1SequencerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1Sequencer{L1SequencerCaller: L1SequencerCaller{contract: contract}, L1SequencerTransactor: L1SequencerTransactor{contract: contract}, L1SequencerFilterer: L1SequencerFilterer{contract: contract}}, nil
}

// L1Sequencer is an auto generated Go binding around an Ethereum contract.
type L1Sequencer struct {
	L1SequencerCaller     // Read-only binding to the contract
	L1SequencerTransactor // Write-only binding to the contract
	L1SequencerFilterer   // Log filterer for contract events
}

// L1SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1SequencerSession struct {
	Contract     *L1Sequencer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1SequencerCallerSession struct {
	Contract *L1SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// L1SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1SequencerTransactorSession struct {
	Contract     *L1SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// L1SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1SequencerRaw struct {
	Contract *L1Sequencer // Generic contract binding to access the raw methods on
}

// L1SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1SequencerCallerRaw struct {
	Contract *L1SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// L1SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1SequencerTransactorRaw struct {
	Contract *L1SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1Sequencer creates a new instance of L1Sequencer, bound to a specific deployed contract.
func NewL1Sequencer(address common.Address, backend bind.ContractBackend) (*L1Sequencer, error) {
	contract, err := bindL1Sequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1Sequencer{L1SequencerCaller: L1SequencerCaller{contract: contract}, L1SequencerTransactor: L1SequencerTransactor{contract: contract}, L1SequencerFilterer: L1SequencerFilterer{contract: contract}}, nil
}

// NewL1SequencerCaller creates a new read-only instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerCaller(address common.Address, caller bind.ContractCaller) (*L1SequencerCaller, error) {
	contract, err := bindL1Sequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1SequencerCaller{contract: contract}, nil
}

// NewL1SequencerTransactor creates a new write-only instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1SequencerTransactor, error) {
	contract, err := bindL1Sequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1SequencerTransactor{contract: contract}, nil
}

// NewL1SequencerFilterer creates a new log filterer instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1SequencerFilterer, error) {
	contract, err := bindL1Sequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1SequencerFilterer{contract: contract}, nil
}

// bindL1Sequencer binds a generic wrapper to an already deployed contract.
func bindL1Sequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1SequencerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Sequencer *L1SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Sequencer.Contract.L1SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Sequencer *L1SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.Contract.L1SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Sequencer *L1SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Sequencer.Contract.L1SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Sequencer *L1SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Sequencer *L1SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Sequencer *L1SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Sequencer.Contract.contract.Transact(opts, method, params...)
}

// GetSequencer is a free data retrieval call binding the contract method 0x4d96a90a.
//
// Solidity: function getSequencer() view returns(address)
func (_L1Sequencer *L1SequencerCaller) GetSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSequencer is a free data retrieval call binding the contract method 0x4d96a90a.
//
// Solidity: function getSequencer() view returns(address)
func (_L1Sequencer *L1SequencerSession) GetSequencer() (common.Address, error) {
	return _L1Sequencer.Contract.GetSequencer(&_L1Sequencer.CallOpts)
}

// GetSequencer is a free data retrieval call binding the contract method 0x4d96a90a.
//
// Solidity: function getSequencer() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) GetSequencer() (common.Address, error) {
	return _L1Sequencer.Contract.GetSequencer(&_L1Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Sequencer *L1SequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Sequencer *L1SequencerSession) Owner() (common.Address, error) {
	return _L1Sequencer.Contract.Owner(&_L1Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) Owner() (common.Address, error) {
	return _L1Sequencer.Contract.Owner(&_L1Sequencer.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_L1Sequencer *L1SequencerCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_L1Sequencer *L1SequencerSession) Sequencer() (common.Address, error) {
	return _L1Sequencer.Contract.Sequencer(&_L1Sequencer.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) Sequencer() (common.Address, error) {
	return _L1Sequencer.Contract.Sequencer(&_L1Sequencer.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1Sequencer *L1SequencerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1Sequencer *L1SequencerSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Initialize(&_L1Sequencer.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L1Sequencer *L1SequencerTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Initialize(&_L1Sequencer.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Sequencer *L1SequencerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Sequencer *L1SequencerSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1Sequencer.Contract.RenounceOwnership(&_L1Sequencer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Sequencer *L1SequencerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1Sequencer.Contract.RenounceOwnership(&_L1Sequencer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Sequencer *L1SequencerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Sequencer *L1SequencerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.TransferOwnership(&_L1Sequencer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Sequencer *L1SequencerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.TransferOwnership(&_L1Sequencer.TransactOpts, newOwner)
}

// UpdateSequencer is a paid mutator transaction binding the contract method 0x43ae20a3.
//
// Solidity: function updateSequencer(address newSequencer) returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateSequencer(opts *bind.TransactOpts, newSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateSequencer", newSequencer)
}

// UpdateSequencer is a paid mutator transaction binding the contract method 0x43ae20a3.
//
// Solidity: function updateSequencer(address newSequencer) returns()
func (_L1Sequencer *L1SequencerSession) UpdateSequencer(newSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateSequencer(&_L1Sequencer.TransactOpts, newSequencer)
}

// UpdateSequencer is a paid mutator transaction binding the contract method 0x43ae20a3.
//
// Solidity: function updateSequencer(address newSequencer) returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateSequencer(newSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateSequencer(&_L1Sequencer.TransactOpts, newSequencer)
}

// L1SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1Sequencer contract.
type L1SequencerInitializedIterator struct {
	Event *L1SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *L1SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerInitialized)
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
		it.Event = new(L1SequencerInitialized)
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
func (it *L1SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerInitialized represents a Initialized event raised by the L1Sequencer contract.
type L1SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Sequencer *L1SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1SequencerInitializedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1SequencerInitializedIterator{contract: _L1Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Sequencer *L1SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerInitialized)
				if err := _L1Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1Sequencer *L1SequencerFilterer) ParseInitialized(log types.Log) (*L1SequencerInitialized, error) {
	event := new(L1SequencerInitialized)
	if err := _L1Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1SequencerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1Sequencer contract.
type L1SequencerOwnershipTransferredIterator struct {
	Event *L1SequencerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1SequencerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerOwnershipTransferred)
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
		it.Event = new(L1SequencerOwnershipTransferred)
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
func (it *L1SequencerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerOwnershipTransferred represents a OwnershipTransferred event raised by the L1Sequencer contract.
type L1SequencerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1Sequencer *L1SequencerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1SequencerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1SequencerOwnershipTransferredIterator{contract: _L1Sequencer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1Sequencer *L1SequencerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1SequencerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerOwnershipTransferred)
				if err := _L1Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1Sequencer *L1SequencerFilterer) ParseOwnershipTransferred(log types.Log) (*L1SequencerOwnershipTransferred, error) {
	event := new(L1SequencerOwnershipTransferred)
	if err := _L1Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1SequencerSequencerUpdatedIterator is returned from FilterSequencerUpdated and is used to iterate over the raw logs and unpacked data for SequencerUpdated events raised by the L1Sequencer contract.
type L1SequencerSequencerUpdatedIterator struct {
	Event *L1SequencerSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *L1SequencerSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerSequencerUpdated)
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
		it.Event = new(L1SequencerSequencerUpdated)
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
func (it *L1SequencerSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerSequencerUpdated represents a SequencerUpdated event raised by the L1Sequencer contract.
type L1SequencerSequencerUpdated struct {
	OldSequencer common.Address
	NewSequencer common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0xcd58b762453bd126b48db83f2cecd464f5281dd7e5e6824b528c09d0482984d6.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer)
func (_L1Sequencer *L1SequencerFilterer) FilterSequencerUpdated(opts *bind.FilterOpts, oldSequencer []common.Address, newSequencer []common.Address) (*L1SequencerSequencerUpdatedIterator, error) {

	var oldSequencerRule []interface{}
	for _, oldSequencerItem := range oldSequencer {
		oldSequencerRule = append(oldSequencerRule, oldSequencerItem)
	}
	var newSequencerRule []interface{}
	for _, newSequencerItem := range newSequencer {
		newSequencerRule = append(newSequencerRule, newSequencerItem)
	}

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "SequencerUpdated", oldSequencerRule, newSequencerRule)
	if err != nil {
		return nil, err
	}
	return &L1SequencerSequencerUpdatedIterator{contract: _L1Sequencer.contract, event: "SequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0xcd58b762453bd126b48db83f2cecd464f5281dd7e5e6824b528c09d0482984d6.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer)
func (_L1Sequencer *L1SequencerFilterer) WatchSequencerUpdated(opts *bind.WatchOpts, sink chan<- *L1SequencerSequencerUpdated, oldSequencer []common.Address, newSequencer []common.Address) (event.Subscription, error) {

	var oldSequencerRule []interface{}
	for _, oldSequencerItem := range oldSequencer {
		oldSequencerRule = append(oldSequencerRule, oldSequencerItem)
	}
	var newSequencerRule []interface{}
	for _, newSequencerItem := range newSequencer {
		newSequencerRule = append(newSequencerRule, newSequencerItem)
	}

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "SequencerUpdated", oldSequencerRule, newSequencerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerSequencerUpdated)
				if err := _L1Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0xcd58b762453bd126b48db83f2cecd464f5281dd7e5e6824b528c09d0482984d6.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer)
func (_L1Sequencer *L1SequencerFilterer) ParseSequencerUpdated(log types.Log) (*L1SequencerSequencerUpdated, error) {
	event := new(L1SequencerSequencerUpdated)
	if err := _L1Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
