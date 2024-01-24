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

// MintManagerMetaData contains all meta data concerning the MintManager contract.
var MintManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_upgrader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_governanceToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINT_CAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINT_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governanceToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractGovernanceToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintPermittedAfter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgrade\",\"inputs\":[{\"name\":\"_newMintManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610aa4380380610aa483398101604081905261002f91610199565b61003833610053565b610041826100a3565b6001600160a01b0316608052506101cc565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6100ab610121565b6001600160a01b0381166101155760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61011e81610053565b50565b6000546001600160a01b0316331461017b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161010c565b565b80516001600160a01b038116811461019457600080fd5b919050565b600080604083850312156101ac57600080fd5b6101b58361017d565b91506101c36020840161017d565b90509250929050565b6080516108a86101fc6000396000818161018101528181610298015281816103a0015261052c01526108a86000f3fe608060405234801561001057600080fd5b50600436106100bd5760003560e01c80638da5cb5b1161007657806398f1312e1161005b57806398f1312e14610161578063f2fde38b14610169578063f96dae0a1461017c57600080fd5b80638da5cb5b14610119578063918f86741461015857600080fd5b806340c10f19116100a757806340c10f19146100f3578063715018a61461010657806383ea6e971461010e57600080fd5b8062f8900c146100c25780630900f010146100de575b600080fd5b6100cb60015481565b6040519081526020015b60405180910390f35b6100f16100ec366004610776565b6101a3565b005b6100f1610101366004610798565b6102f7565b6100f161058c565b6100cb6301e1338081565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d5565b6100cb6103e881565b6100cb601481565b6100f1610177366004610776565b6105a0565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6101ab610657565b73ffffffffffffffffffffffffffffffffffffffff8116610253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4d696e744d616e616765723a206d696e74206d616e616765722063616e6e6f7460448201527f20626520746865207a65726f206164647265737300000000000000000000000060648201526084015b60405180910390fd5b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b90602401600060405180830381600087803b1580156102dc57600080fd5b505af11580156102f0573d6000803e3d6000fd5b5050505050565b6102ff610657565b600154156104cf57426001541115610399576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4d696e744d616e616765723a206d696e74696e67206e6f74207065726d69747460448201527f6564207965740000000000000000000000000000000000000000000000000000606482015260840161024a565b6103e860147f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610409573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042d91906107c2565b610437919061080a565b6104419190610847565b8111156104cf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a206d696e7420616d6f756e74206578636565647360448201527f2063617000000000000000000000000000000000000000000000000000000000606482015260840161024a565b6104dd6301e1338042610882565b6001556040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f000000000000000000000000000000000000000000000000000000000000000016906340c10f1990604401600060405180830381600087803b15801561057057600080fd5b505af1158015610584573d6000803e3d6000fd5b505050505050565b610594610657565b61059e60006106d8565b565b6105a8610657565b73ffffffffffffffffffffffffffffffffffffffff811661064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161024a565b610654816106d8565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161024a565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461077157600080fd5b919050565b60006020828403121561078857600080fd5b6107918261074d565b9392505050565b600080604083850312156107ab57600080fd5b6107b48361074d565b946020939093013593505050565b6000602082840312156107d457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610842576108426107db565b500290565b60008261087d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b80820180821115610895576108956107db565b9291505056fea164736f6c6343000810000a",
}

// MintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MintManagerMetaData.ABI instead.
var MintManagerABI = MintManagerMetaData.ABI

// MintManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintManagerMetaData.Bin instead.
var MintManagerBin = MintManagerMetaData.Bin

// DeployMintManager deploys a new Ethereum contract, binding an instance of MintManager to it.
func DeployMintManager(auth *bind.TransactOpts, backend bind.ContractBackend, _upgrader common.Address, _governanceToken common.Address) (common.Address, *types.Transaction, *MintManager, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintManagerBin), backend, _upgrader, _governanceToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintManager{MintManagerCaller: MintManagerCaller{contract: contract}, MintManagerTransactor: MintManagerTransactor{contract: contract}, MintManagerFilterer: MintManagerFilterer{contract: contract}}, nil
}

// MintManager is an auto generated Go binding around an Ethereum contract.
type MintManager struct {
	MintManagerCaller     // Read-only binding to the contract
	MintManagerTransactor // Write-only binding to the contract
	MintManagerFilterer   // Log filterer for contract events
}

// MintManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintManagerSession struct {
	Contract     *MintManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintManagerCallerSession struct {
	Contract *MintManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MintManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintManagerTransactorSession struct {
	Contract     *MintManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MintManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintManagerRaw struct {
	Contract *MintManager // Generic contract binding to access the raw methods on
}

// MintManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintManagerCallerRaw struct {
	Contract *MintManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MintManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintManagerTransactorRaw struct {
	Contract *MintManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintManager creates a new instance of MintManager, bound to a specific deployed contract.
func NewMintManager(address common.Address, backend bind.ContractBackend) (*MintManager, error) {
	contract, err := bindMintManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintManager{MintManagerCaller: MintManagerCaller{contract: contract}, MintManagerTransactor: MintManagerTransactor{contract: contract}, MintManagerFilterer: MintManagerFilterer{contract: contract}}, nil
}

// NewMintManagerCaller creates a new read-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerCaller(address common.Address, caller bind.ContractCaller) (*MintManagerCaller, error) {
	contract, err := bindMintManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerCaller{contract: contract}, nil
}

// NewMintManagerTransactor creates a new write-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MintManagerTransactor, error) {
	contract, err := bindMintManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerTransactor{contract: contract}, nil
}

// NewMintManagerFilterer creates a new log filterer instance of MintManager, bound to a specific deployed contract.
func NewMintManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MintManagerFilterer, error) {
	contract, err := bindMintManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintManagerFilterer{contract: contract}, nil
}

// bindMintManager binds a generic wrapper to an already deployed contract.
func bindMintManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.MintManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transact(opts, method, params...)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerSession) DENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.DENOMINATOR(&_MintManager.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCallerSession) DENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.DENOMINATOR(&_MintManager.CallOpts)
}

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerCaller) MINTCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "MINT_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerSession) MINTCAP() (*big.Int, error) {
	return _MintManager.Contract.MINTCAP(&_MintManager.CallOpts)
}

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerCallerSession) MINTCAP() (*big.Int, error) {
	return _MintManager.Contract.MINTCAP(&_MintManager.CallOpts)
}

// MINTPERIOD is a free data retrieval call binding the contract method 0x83ea6e97.
//
// Solidity: function MINT_PERIOD() view returns(uint256)
func (_MintManager *MintManagerCaller) MINTPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "MINT_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTPERIOD is a free data retrieval call binding the contract method 0x83ea6e97.
//
// Solidity: function MINT_PERIOD() view returns(uint256)
func (_MintManager *MintManagerSession) MINTPERIOD() (*big.Int, error) {
	return _MintManager.Contract.MINTPERIOD(&_MintManager.CallOpts)
}

// MINTPERIOD is a free data retrieval call binding the contract method 0x83ea6e97.
//
// Solidity: function MINT_PERIOD() view returns(uint256)
func (_MintManager *MintManagerCallerSession) MINTPERIOD() (*big.Int, error) {
	return _MintManager.Contract.MINTPERIOD(&_MintManager.CallOpts)
}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_MintManager *MintManagerCaller) GovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "governanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_MintManager *MintManagerSession) GovernanceToken() (common.Address, error) {
	return _MintManager.Contract.GovernanceToken(&_MintManager.CallOpts)
}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_MintManager *MintManagerCallerSession) GovernanceToken() (common.Address, error) {
	return _MintManager.Contract.GovernanceToken(&_MintManager.CallOpts)
}

// MintPermittedAfter is a free data retrieval call binding the contract method 0x00f8900c.
//
// Solidity: function mintPermittedAfter() view returns(uint256)
func (_MintManager *MintManagerCaller) MintPermittedAfter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "mintPermittedAfter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPermittedAfter is a free data retrieval call binding the contract method 0x00f8900c.
//
// Solidity: function mintPermittedAfter() view returns(uint256)
func (_MintManager *MintManagerSession) MintPermittedAfter() (*big.Int, error) {
	return _MintManager.Contract.MintPermittedAfter(&_MintManager.CallOpts)
}

// MintPermittedAfter is a free data retrieval call binding the contract method 0x00f8900c.
//
// Solidity: function mintPermittedAfter() view returns(uint256)
func (_MintManager *MintManagerCallerSession) MintPermittedAfter() (*big.Int, error) {
	return _MintManager.Contract.MintPermittedAfter(&_MintManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCallerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintManager *MintManagerTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintManager *MintManagerSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Mint(&_MintManager.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintManager *MintManagerTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Mint(&_MintManager.TransactOpts, _account, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newMintManager) returns()
func (_MintManager *MintManagerTransactor) Upgrade(opts *bind.TransactOpts, _newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "upgrade", _newMintManager)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newMintManager) returns()
func (_MintManager *MintManagerSession) Upgrade(_newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.Upgrade(&_MintManager.TransactOpts, _newMintManager)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newMintManager) returns()
func (_MintManager *MintManagerTransactorSession) Upgrade(_newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.Upgrade(&_MintManager.TransactOpts, _newMintManager)
}

// MintManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintManager contract.
type MintManagerOwnershipTransferredIterator struct {
	Event *MintManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerOwnershipTransferred)
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
		it.Event = new(MintManagerOwnershipTransferred)
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
func (it *MintManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MintManager contract.
type MintManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerOwnershipTransferredIterator{contract: _MintManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerOwnershipTransferred)
				if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintManager *MintManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MintManagerOwnershipTransferred, error) {
	event := new(MintManagerOwnershipTransferred)
	if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
