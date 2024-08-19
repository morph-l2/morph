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

// TransparentUpgradeableProxyMetaData contains all meta data concerning the TransparentUpgradeableProxy contract.
var TransparentUpgradeableProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040516200108138038062001081833981016040819052620000269162000415565b82816200003582825f6200004c565b50620000439050826200007d565b50505062000540565b6200005783620000ee565b5f82511180620000645750805b1562000078576200007683836200012f565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620000be5f805160206200103a833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1620000eb816200015e565b50565b620000f981620001fb565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606200015783836040518060600160405280602781526020016200105a6027913962000292565b9392505050565b6001600160a01b038116620001c95760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f805160206200103a8339815191525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6200026a5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401620001c0565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc620001da565b60605f80856001600160a01b031685604051620002b09190620004ef565b5f60405180830381855af49150503d805f8114620002ea576040519150601f19603f3d011682016040523d82523d5f602084013e620002ef565b606091505b50909250905062000303868383876200030d565b9695505050505050565b60608315620003805782515f0362000378576001600160a01b0385163b620003785760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401620001c0565b50816200038c565b6200038c838362000394565b949350505050565b815115620003a55781518083602001fd5b8060405162461bcd60e51b8152600401620001c091906200050c565b80516001600160a01b0381168114620003d8575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200040d578181015183820152602001620003f3565b50505f910152565b5f805f6060848603121562000428575f80fd5b6200043384620003c1565b92506200044360208501620003c1565b60408501519092506001600160401b038082111562000460575f80fd5b818601915086601f83011262000474575f80fd5b815181811115620004895762000489620003dd565b604051601f8201601f19908116603f01168101908382118183101715620004b457620004b4620003dd565b81604052828152896020848701011115620004cd575f80fd5b620004e0836020830160208801620003f1565b80955050505050509250925092565b5f825162000502818460208701620003f1565b9190910192915050565b602081525f82518060208401526200052c816040850160208701620003f1565b601f01601f19169190910160400192915050565b610aec806200054e5f395ff3fe60806040523661001357610011610017565b005b6100115b61001f6102a0565b73ffffffffffffffffffffffffffffffffffffffff1633036102965760607fffffffff000000000000000000000000000000000000000000000000000000005f35167fc9a6301a000000000000000000000000000000000000000000000000000000008101610097576100906102df565b915061028e565b7fb0e10d7a000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016100e857610090610332565b7f70d7c690000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160161013957610090610376565b7f07ae5bc0000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160161018a576100906103a6565b7fa39f25e5000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016101db576100906103f2565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b815160208301f35b61029e610405565b565b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b60606102e9610415565b5f6102f736600481846108bc565b810190610304919061090b565b905061031f8160405180602001604052805f8152505f61041f565b505060408051602081019091525f815290565b60605f8061034336600481846108bc565b8101906103509190610951565b915091506103608282600161041f565b60405180602001604052805f8152509250505090565b6060610380610415565b5f61038e36600481846108bc565b81019061039b919061090b565b905061031f8161044a565b60606103b0610415565b5f6103b96102a0565b6040805173ffffffffffffffffffffffffffffffffffffffff831660208201529192500160405160208183030381529060405291505090565b60606103fc610415565b5f6103b96104ae565b61029e6104106104ae565b6104bc565b341561029e575f80fd5b610428836104da565b5f825111806104345750805b15610445576104438383610526565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104736102a0565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a16104ab81610552565b50565b5f6104b761065e565b905090565b365f80375f80365f845af43d5f803e8080156104d6573d5ff35b3d5ffd5b6104e381610685565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b606061054b8383604051806060016040528060278152602001610ab960279139610750565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81166105f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610285565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6102c3565b73ffffffffffffffffffffffffffffffffffffffff81163b610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610285565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610618565b60605f808573ffffffffffffffffffffffffffffffffffffffff16856040516107799190610a4d565b5f60405180830381855af49150503d805f81146107b1576040519150601f19603f3d011682016040523d82523d5f602084013e6107b6565b606091505b50915091506107c7868383876107d1565b9695505050505050565b606083156108665782515f0361085f5773ffffffffffffffffffffffffffffffffffffffff85163b61085f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610285565b5081610870565b6108708383610878565b949350505050565b8151156108885781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102859190610a68565b5f80858511156108ca575f80fd5b838611156108d6575f80fd5b5050820193919092039150565b803573ffffffffffffffffffffffffffffffffffffffff81168114610906575f80fd5b919050565b5f6020828403121561091b575f80fd5b61054b826108e3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610962575f80fd5b61096b836108e3565b9150602083013567ffffffffffffffff80821115610987575f80fd5b818501915085601f83011261099a575f80fd5b8135818111156109ac576109ac610924565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156109f2576109f2610924565b81604052828152886020848701011115610a0a575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f5b83811015610a45578181015183820152602001610a2d565b50505f910152565b5f8251610a5e818460208701610a2b565b9190910192915050565b602081525f8251806020840152610a86816040850160208701610a2b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c6343000818000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// TransparentUpgradeableProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use TransparentUpgradeableProxyMetaData.ABI instead.
var TransparentUpgradeableProxyABI = TransparentUpgradeableProxyMetaData.ABI

// TransparentUpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransparentUpgradeableProxyMetaData.Bin instead.
var TransparentUpgradeableProxyBin = TransparentUpgradeableProxyMetaData.Bin

// DeployTransparentUpgradeableProxy deploys a new Ethereum contract, binding an instance of TransparentUpgradeableProxy to it.
func DeployTransparentUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin_ common.Address, _data []byte) (common.Address, *types.Transaction, *TransparentUpgradeableProxy, error) {
	parsed, err := TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransparentUpgradeableProxyBin), backend, _logic, admin_, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// TransparentUpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type TransparentUpgradeableProxy struct {
	TransparentUpgradeableProxyCaller     // Read-only binding to the contract
	TransparentUpgradeableProxyTransactor // Write-only binding to the contract
	TransparentUpgradeableProxyFilterer   // Log filterer for contract events
}

// TransparentUpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransparentUpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentUpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransparentUpgradeableProxySession struct {
	Contract     *TransparentUpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransparentUpgradeableProxyCallerSession struct {
	Contract *TransparentUpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TransparentUpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransparentUpgradeableProxyTransactorSession struct {
	Contract     *TransparentUpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TransparentUpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransparentUpgradeableProxyRaw struct {
	Contract *TransparentUpgradeableProxy // Generic contract binding to access the raw methods on
}

// TransparentUpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyCallerRaw struct {
	Contract *TransparentUpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TransparentUpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransparentUpgradeableProxyTransactorRaw struct {
	Contract *TransparentUpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransparentUpgradeableProxy creates a new instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*TransparentUpgradeableProxy, error) {
	contract, err := bindTransparentUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxy{TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// NewTransparentUpgradeableProxyCaller creates a new read-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*TransparentUpgradeableProxyCaller, error) {
	contract, err := bindTransparentUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyCaller{contract: contract}, nil
}

// NewTransparentUpgradeableProxyTransactor creates a new write-only instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TransparentUpgradeableProxyTransactor, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyTransactor{contract: contract}, nil
}

// NewTransparentUpgradeableProxyFilterer creates a new log filterer instance of TransparentUpgradeableProxy, bound to a specific deployed contract.
func NewTransparentUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TransparentUpgradeableProxyFilterer, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyFilterer{contract: contract}, nil
}

// bindTransparentUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindTransparentUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Receive(&_TransparentUpgradeableProxy.TransactOpts)
}

// TransparentUpgradeableProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyAdminChangedIterator struct {
	Event *TransparentUpgradeableProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyAdminChanged)
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
		it.Event = new(TransparentUpgradeableProxyAdminChanged)
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
func (it *TransparentUpgradeableProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyAdminChanged represents a AdminChanged event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxyAdminChangedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyAdminChangedIterator{contract: _TransparentUpgradeableProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyAdminChanged)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseAdminChanged(log types.Log) (*TransparentUpgradeableProxyAdminChanged, error) {
	event := new(TransparentUpgradeableProxyAdminChanged)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyBeaconUpgradedIterator struct {
	Event *TransparentUpgradeableProxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyBeaconUpgraded)
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
		it.Event = new(TransparentUpgradeableProxyBeaconUpgraded)
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
func (it *TransparentUpgradeableProxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyBeaconUpgraded represents a BeaconUpgraded event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TransparentUpgradeableProxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyBeaconUpgradedIterator{contract: _TransparentUpgradeableProxy.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyBeaconUpgraded)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseBeaconUpgraded(log types.Log) (*TransparentUpgradeableProxyBeaconUpgraded, error) {
	event := new(TransparentUpgradeableProxyBeaconUpgraded)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentUpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgradedIterator struct {
	Event *TransparentUpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentUpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyUpgraded)
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
		it.Event = new(TransparentUpgradeableProxyUpgraded)
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
func (it *TransparentUpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentUpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentUpgradeableProxyUpgraded represents a Upgraded event raised by the TransparentUpgradeableProxy contract.
type TransparentUpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyUpgradedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentUpgradeableProxyUpgraded)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*TransparentUpgradeableProxyUpgraded, error) {
	event := new(TransparentUpgradeableProxyUpgraded)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
