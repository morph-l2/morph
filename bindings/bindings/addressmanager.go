// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// AddressManagerMetaData contains all meta data concerning the AddressManager contract.
var AddressManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6105ec8061007a5f395ff3fe608060405234801561000f575f80fd5b5060043610610064575f3560e01c80639b2ea4bd1161004d5780639b2ea4bd146100b4578063bf40fac1146100c7578063f2fde38b146100da575f80fd5b8063715018a6146100685780638da5cb5b14610072575b5f80fd5b6100706100ed565b005b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100706100c23660046104e5565b610100565b61008b6100d5366004610530565b6101d1565b6100706100e836600461056a565b61020a565b6100f56102c6565b6100fe5f610346565b565b6101086102c6565b5f610112836103ba565b5f8181526001602052604090819020805473ffffffffffffffffffffffffffffffffffffffff8681167fffffffffffffffffffffffff00000000000000000000000000000000000000008316179092559151929350169061017490859061058a565b6040805191829003822073ffffffffffffffffffffffffffffffffffffffff808716845284166020840152917f9416a153a346f93d95f94b064ae3f148b6460473c6e82b3f9fc2521b873fcd6c910160405180910390a250505050565b5f60015f6101de846103ba565b815260208101919091526040015f205473ffffffffffffffffffffffffffffffffffffffff1692915050565b6102126102c6565b73ffffffffffffffffffffffffffffffffffffffff81166102ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c381610346565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146100fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b1565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f816040516020016103cc919061058a565b604051602081830303815290604052805190602001209050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112610425575f80fd5b813567ffffffffffffffff80821115610440576104406103e9565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610486576104866103e9565b8160405283815286602085880101111561049e575f80fd5b836020870160208301375f602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146104e0575f80fd5b919050565b5f80604083850312156104f6575f80fd5b823567ffffffffffffffff81111561050c575f80fd5b61051885828601610416565b925050610527602084016104bd565b90509250929050565b5f60208284031215610540575f80fd5b813567ffffffffffffffff811115610556575f80fd5b61056284828501610416565b949350505050565b5f6020828403121561057a575f80fd5b610583826104bd565b9392505050565b5f82515f5b818110156105a9576020818601810151858301520161058f565b505f92019182525091905056fea2646970667358221220d7bde888d6474b9c59b455ad37927da74f75cda5c5f2711d434779ad40b2a93264736f6c63430008180033",
}

// AddressManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressManagerMetaData.ABI instead.
var AddressManagerABI = AddressManagerMetaData.ABI

// AddressManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressManagerMetaData.Bin instead.
var AddressManagerBin = AddressManagerMetaData.Bin

// DeployAddressManager deploys a new Ethereum contract, binding an instance of AddressManager to it.
func DeployAddressManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressManager, error) {
	parsed, err := AddressManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressManager{AddressManagerCaller: AddressManagerCaller{contract: contract}, AddressManagerTransactor: AddressManagerTransactor{contract: contract}, AddressManagerFilterer: AddressManagerFilterer{contract: contract}}, nil
}

// AddressManager is an auto generated Go binding around an Ethereum contract.
type AddressManager struct {
	AddressManagerCaller     // Read-only binding to the contract
	AddressManagerTransactor // Write-only binding to the contract
	AddressManagerFilterer   // Log filterer for contract events
}

// AddressManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressManagerSession struct {
	Contract     *AddressManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressManagerCallerSession struct {
	Contract *AddressManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddressManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressManagerTransactorSession struct {
	Contract     *AddressManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddressManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressManagerRaw struct {
	Contract *AddressManager // Generic contract binding to access the raw methods on
}

// AddressManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressManagerCallerRaw struct {
	Contract *AddressManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AddressManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressManagerTransactorRaw struct {
	Contract *AddressManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressManager creates a new instance of AddressManager, bound to a specific deployed contract.
func NewAddressManager(address common.Address, backend bind.ContractBackend) (*AddressManager, error) {
	contract, err := bindAddressManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressManager{AddressManagerCaller: AddressManagerCaller{contract: contract}, AddressManagerTransactor: AddressManagerTransactor{contract: contract}, AddressManagerFilterer: AddressManagerFilterer{contract: contract}}, nil
}

// NewAddressManagerCaller creates a new read-only instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerCaller(address common.Address, caller bind.ContractCaller) (*AddressManagerCaller, error) {
	contract, err := bindAddressManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressManagerCaller{contract: contract}, nil
}

// NewAddressManagerTransactor creates a new write-only instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressManagerTransactor, error) {
	contract, err := bindAddressManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressManagerTransactor{contract: contract}, nil
}

// NewAddressManagerFilterer creates a new log filterer instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressManagerFilterer, error) {
	contract, err := bindAddressManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressManagerFilterer{contract: contract}, nil
}

// bindAddressManager binds a generic wrapper to an already deployed contract.
func bindAddressManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressManager *AddressManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressManager.Contract.AddressManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressManager *AddressManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.Contract.AddressManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressManager *AddressManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressManager.Contract.AddressManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressManager *AddressManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressManager *AddressManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressManager *AddressManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressManager.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_AddressManager *AddressManagerCaller) GetAddress(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _AddressManager.contract.Call(opts, &out, "getAddress", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_AddressManager *AddressManagerSession) GetAddress(_name string) (common.Address, error) {
	return _AddressManager.Contract.GetAddress(&_AddressManager.CallOpts, _name)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_AddressManager *AddressManagerCallerSession) GetAddress(_name string) (common.Address, error) {
	return _AddressManager.Contract.GetAddress(&_AddressManager.CallOpts, _name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerSession) Owner() (common.Address, error) {
	return _AddressManager.Contract.Owner(&_AddressManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerCallerSession) Owner() (common.Address, error) {
	return _AddressManager.Contract.Owner(&_AddressManager.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressManager.Contract.RenounceOwnership(&_AddressManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressManager.Contract.RenounceOwnership(&_AddressManager.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_AddressManager *AddressManagerTransactor) SetAddress(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "setAddress", _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_AddressManager *AddressManagerSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.SetAddress(&_AddressManager.TransactOpts, _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_AddressManager *AddressManagerTransactorSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.SetAddress(&_AddressManager.TransactOpts, _name, _address)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.TransferOwnership(&_AddressManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.TransferOwnership(&_AddressManager.TransactOpts, newOwner)
}

// AddressManagerAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AddressManager contract.
type AddressManagerAddressSetIterator struct {
	Event *AddressManagerAddressSet // Event containing the contract specifics and raw log

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
func (it *AddressManagerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressManagerAddressSet)
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
		it.Event = new(AddressManagerAddressSet)
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
func (it *AddressManagerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressManagerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressManagerAddressSet represents a AddressSet event raised by the AddressManager contract.
type AddressManagerAddressSet struct {
	Name       common.Hash
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x9416a153a346f93d95f94b064ae3f148b6460473c6e82b3f9fc2521b873fcd6c.
//
// Solidity: event AddressSet(string indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) FilterAddressSet(opts *bind.FilterOpts, name []string) (*AddressManagerAddressSetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AddressManager.contract.FilterLogs(opts, "AddressSet", nameRule)
	if err != nil {
		return nil, err
	}
	return &AddressManagerAddressSetIterator{contract: _AddressManager.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x9416a153a346f93d95f94b064ae3f148b6460473c6e82b3f9fc2521b873fcd6c.
//
// Solidity: event AddressSet(string indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AddressManagerAddressSet, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AddressManager.contract.WatchLogs(opts, "AddressSet", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressManagerAddressSet)
				if err := _AddressManager.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0x9416a153a346f93d95f94b064ae3f148b6460473c6e82b3f9fc2521b873fcd6c.
//
// Solidity: event AddressSet(string indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) ParseAddressSet(log types.Log) (*AddressManagerAddressSet, error) {
	event := new(AddressManagerAddressSet)
	if err := _AddressManager.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressManager contract.
type AddressManagerOwnershipTransferredIterator struct {
	Event *AddressManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressManagerOwnershipTransferred)
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
		it.Event = new(AddressManagerOwnershipTransferred)
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
func (it *AddressManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AddressManager contract.
type AddressManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressManager *AddressManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressManagerOwnershipTransferredIterator{contract: _AddressManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressManager *AddressManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressManagerOwnershipTransferred)
				if err := _AddressManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressManager *AddressManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AddressManagerOwnershipTransferred, error) {
	event := new(AddressManagerOwnershipTransferred)
	if err := _AddressManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
