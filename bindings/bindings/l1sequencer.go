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

// L1SequencerHistoryRecord is an auto generated low-level Go binding around an user-defined struct.
type L1SequencerHistoryRecord struct {
	StartL2Block  uint64
	SequencerAddr common.Address
}

// L1SequencerMetaData contains all meta data concerning the L1Sequencer contract.
var L1SequencerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSequencer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startL2Block\",\"type\":\"uint64\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"l2Height\",\"type\":\"uint64\"}],\"name\":\"getSequencerAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startL2Block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"internalType\":\"structL1Sequencer.HistoryRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerHistory\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startL2Block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"firstSequencer\",\"type\":\"address\"}],\"name\":\"setFirstSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startL2Block\",\"type\":\"uint64\"}],\"name\":\"updateSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610fc38061001d5f395ff3fe608060405234801561000f575f80fd5b50600436106100c4575f3560e01c8063715018a61161007d578063c4d66de811610058578063c4d66de8146101b6578063f151ce9e146101c9578063f2fde38b146101dc575f80fd5b8063715018a61461017d578063761a90fd146101855780638da5cb5b14610198575f80fd5b80634d96a90a116100ad5780634d96a90a146100f35780636628aea1146101205780636d8ce3d214610135575f80fd5b80630df8955e146100c85780633d5767ce146100dd575b5f80fd5b6100db6100d6366004610df2565b6101ef565b005b6065546040519081526020015b60405180910390f35b6100fb610391565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ea565b610128610434565b6040516100ea9190610e12565b610148610143366004610e80565b6104c1565b6040805167ffffffffffffffff909316835273ffffffffffffffffffffffffffffffffffffffff9091166020830152016100ea565b6100db61050d565b6100db610193366004610eae565b610520565b60335473ffffffffffffffffffffffffffffffffffffffff166100fb565b6100db6101c4366004610df2565b6107a4565b6100fb6101d7366004610edf565b610983565b6100db6101ea366004610df2565b610b4c565b6101f7610be9565b6065541561024c5760405162461bcd60e51b815260206004820152601360248201527f616c726561647920696e697469616c697a65640000000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166102af5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c6964206164647265737300000000000000000000000000000000006044820152606401610243565b6040805180820182525f80825273ffffffffffffffffffffffffffffffffffffffff84811660208085018281526065805460018101825590865295517f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c79096018054915190941668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090911667ffffffffffffffff9096169590951794909417909155925181815290917ffed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8910160405180910390a350565b6065545f906103e25760405162461bcd60e51b815260206004820152601760248201527f6e6f2073657175656e63657220636f6e666967757265640000000000000000006044820152606401610243565b606580546103f290600190610f25565b8154811061040257610402610f3e565b5f9182526020909120015468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16919050565b60606065805480602002602001604051908101604052809291908181526020015f905b828210156104b8575f848152602090819020604080518082019091529084015467ffffffffffffffff8116825268010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681830152825260019092019101610457565b50505050905090565b606581815481106104d0575f80fd5b5f9182526020909120015467ffffffffffffffff8116915068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1682565b610515610be9565b61051e5f610c50565b565b610528610be9565b73ffffffffffffffffffffffffffffffffffffffff821661058b5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c6964206164647265737300000000000000000000000000000000006044820152606401610243565b6065546105da5760405162461bcd60e51b815260206004820152600f60248201527f6e6f7420696e697469616c697a656400000000000000000000000000000000006044820152606401610243565b606580546105ea90600190610f25565b815481106105fa576105fa610f3e565b5f9182526020909120015467ffffffffffffffff908116908216116106875760405162461bcd60e51b815260206004820152602d60248201527f73746172744c32426c6f636b206d75737420626520677265617465722074686160448201527f6e206c617374207265636f7264000000000000000000000000000000000000006064820152608401610243565b606580545f919061069a90600190610f25565b815481106106aa576106aa610f3e565b5f9182526020808320919091015460408051808201825267ffffffffffffffff87811680835273ffffffffffffffffffffffffffffffffffffffff8a8116848801818152606580546001810182559a5294517f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c790990180549551999094167fffffffff00000000000000000000000000000000000000000000000000000000909516949094176801000000000000000098821689021790925592519283529490920490931693509183917ffed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8910160405180910390a3505050565b5f54610100900460ff16158080156107c257505f54600160ff909116105b806107db5750303b1580156107db57505f5460ff166001145b61084d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610243565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108a9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff821661090c5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c6964206f776e6572000000000000000000000000000000000000006044820152606401610243565b610914610cc6565b61091d82610c50565b801561097f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6065545f90806109d55760405162461bcd60e51b815260206004820152601760248201527f6e6f2073657175656e63657220636f6e666967757265640000000000000000006044820152606401610243565b5f806109e2600184610f25565b90505f5b818311610a81575f60026109fa8486610f6b565b610a049190610f7e565b90508667ffffffffffffffff1660658281548110610a2457610a24610f3e565b5f9182526020909120015467ffffffffffffffff1611610a6057809150828103610a4e5750610a81565b610a59816001610f6b565b9350610a7b565b805f03610a6d5750610a81565b610a78600182610f25565b92505b506109e6565b8567ffffffffffffffff1660658281548110610a9f57610a9f610f3e565b5f9182526020909120015467ffffffffffffffff161115610b025760405162461bcd60e51b815260206004820152601660248201527f6e6f2073657175656e63657220617420686569676874000000000000000000006044820152606401610243565b60658181548110610b1557610b15610f3e565b5f9182526020909120015468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff169695505050505050565b610b54610be9565b73ffffffffffffffffffffffffffffffffffffffff8116610bdd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610243565b610be681610c50565b50565b60335473ffffffffffffffffffffffffffffffffffffffff16331461051e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610243565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610d425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610243565b61051e5f54610100900460ff16610dc15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610243565b61051e33610c50565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ded575f80fd5b919050565b5f60208284031215610e02575f80fd5b610e0b82610dca565b9392505050565b602080825282518282018190525f919060409081850190868401855b82811015610e73578151805167ffffffffffffffff16855286015173ffffffffffffffffffffffffffffffffffffffff16868501529284019290850190600101610e2e565b5091979650505050505050565b5f60208284031215610e90575f80fd5b5035919050565b803567ffffffffffffffff81168114610ded575f80fd5b5f8060408385031215610ebf575f80fd5b610ec883610dca565b9150610ed660208401610e97565b90509250929050565b5f60208284031215610eef575f80fd5b610e0b82610e97565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610f3857610f38610ef8565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80820180821115610f3857610f38610ef8565b5f82610fb1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a",
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

// GetSequencerAt is a free data retrieval call binding the contract method 0xf151ce9e.
//
// Solidity: function getSequencerAt(uint64 l2Height) view returns(address)
func (_L1Sequencer *L1SequencerCaller) GetSequencerAt(opts *bind.CallOpts, l2Height uint64) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerAt", l2Height)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSequencerAt is a free data retrieval call binding the contract method 0xf151ce9e.
//
// Solidity: function getSequencerAt(uint64 l2Height) view returns(address)
func (_L1Sequencer *L1SequencerSession) GetSequencerAt(l2Height uint64) (common.Address, error) {
	return _L1Sequencer.Contract.GetSequencerAt(&_L1Sequencer.CallOpts, l2Height)
}

// GetSequencerAt is a free data retrieval call binding the contract method 0xf151ce9e.
//
// Solidity: function getSequencerAt(uint64 l2Height) view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerAt(l2Height uint64) (common.Address, error) {
	return _L1Sequencer.Contract.GetSequencerAt(&_L1Sequencer.CallOpts, l2Height)
}

// GetSequencerHistory is a free data retrieval call binding the contract method 0x6628aea1.
//
// Solidity: function getSequencerHistory() view returns((uint64,address)[])
func (_L1Sequencer *L1SequencerCaller) GetSequencerHistory(opts *bind.CallOpts) ([]L1SequencerHistoryRecord, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerHistory")

	if err != nil {
		return *new([]L1SequencerHistoryRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]L1SequencerHistoryRecord)).(*[]L1SequencerHistoryRecord)

	return out0, err

}

// GetSequencerHistory is a free data retrieval call binding the contract method 0x6628aea1.
//
// Solidity: function getSequencerHistory() view returns((uint64,address)[])
func (_L1Sequencer *L1SequencerSession) GetSequencerHistory() ([]L1SequencerHistoryRecord, error) {
	return _L1Sequencer.Contract.GetSequencerHistory(&_L1Sequencer.CallOpts)
}

// GetSequencerHistory is a free data retrieval call binding the contract method 0x6628aea1.
//
// Solidity: function getSequencerHistory() view returns((uint64,address)[])
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerHistory() ([]L1SequencerHistoryRecord, error) {
	return _L1Sequencer.Contract.GetSequencerHistory(&_L1Sequencer.CallOpts)
}

// GetSequencerHistoryLength is a free data retrieval call binding the contract method 0x3d5767ce.
//
// Solidity: function getSequencerHistoryLength() view returns(uint256)
func (_L1Sequencer *L1SequencerCaller) GetSequencerHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequencerHistoryLength is a free data retrieval call binding the contract method 0x3d5767ce.
//
// Solidity: function getSequencerHistoryLength() view returns(uint256)
func (_L1Sequencer *L1SequencerSession) GetSequencerHistoryLength() (*big.Int, error) {
	return _L1Sequencer.Contract.GetSequencerHistoryLength(&_L1Sequencer.CallOpts)
}

// GetSequencerHistoryLength is a free data retrieval call binding the contract method 0x3d5767ce.
//
// Solidity: function getSequencerHistoryLength() view returns(uint256)
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerHistoryLength() (*big.Int, error) {
	return _L1Sequencer.Contract.GetSequencerHistoryLength(&_L1Sequencer.CallOpts)
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

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint64 startL2Block, address sequencerAddr)
func (_L1Sequencer *L1SequencerCaller) SequencerHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartL2Block  uint64
	SequencerAddr common.Address
}, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "sequencerHistory", arg0)

	outstruct := new(struct {
		StartL2Block  uint64
		SequencerAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartL2Block = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SequencerAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint64 startL2Block, address sequencerAddr)
func (_L1Sequencer *L1SequencerSession) SequencerHistory(arg0 *big.Int) (struct {
	StartL2Block  uint64
	SequencerAddr common.Address
}, error) {
	return _L1Sequencer.Contract.SequencerHistory(&_L1Sequencer.CallOpts, arg0)
}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint64 startL2Block, address sequencerAddr)
func (_L1Sequencer *L1SequencerCallerSession) SequencerHistory(arg0 *big.Int) (struct {
	StartL2Block  uint64
	SequencerAddr common.Address
}, error) {
	return _L1Sequencer.Contract.SequencerHistory(&_L1Sequencer.CallOpts, arg0)
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

// SetFirstSequencer is a paid mutator transaction binding the contract method 0x0df8955e.
//
// Solidity: function setFirstSequencer(address firstSequencer) returns()
func (_L1Sequencer *L1SequencerTransactor) SetFirstSequencer(opts *bind.TransactOpts, firstSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "setFirstSequencer", firstSequencer)
}

// SetFirstSequencer is a paid mutator transaction binding the contract method 0x0df8955e.
//
// Solidity: function setFirstSequencer(address firstSequencer) returns()
func (_L1Sequencer *L1SequencerSession) SetFirstSequencer(firstSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.SetFirstSequencer(&_L1Sequencer.TransactOpts, firstSequencer)
}

// SetFirstSequencer is a paid mutator transaction binding the contract method 0x0df8955e.
//
// Solidity: function setFirstSequencer(address firstSequencer) returns()
func (_L1Sequencer *L1SequencerTransactorSession) SetFirstSequencer(firstSequencer common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.SetFirstSequencer(&_L1Sequencer.TransactOpts, firstSequencer)
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

// UpdateSequencer is a paid mutator transaction binding the contract method 0x761a90fd.
//
// Solidity: function updateSequencer(address newSequencer, uint64 startL2Block) returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateSequencer(opts *bind.TransactOpts, newSequencer common.Address, startL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateSequencer", newSequencer, startL2Block)
}

// UpdateSequencer is a paid mutator transaction binding the contract method 0x761a90fd.
//
// Solidity: function updateSequencer(address newSequencer, uint64 startL2Block) returns()
func (_L1Sequencer *L1SequencerSession) UpdateSequencer(newSequencer common.Address, startL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateSequencer(&_L1Sequencer.TransactOpts, newSequencer, startL2Block)
}

// UpdateSequencer is a paid mutator transaction binding the contract method 0x761a90fd.
//
// Solidity: function updateSequencer(address newSequencer, uint64 startL2Block) returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateSequencer(newSequencer common.Address, startL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateSequencer(&_L1Sequencer.TransactOpts, newSequencer, startL2Block)
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
	StartL2Block uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0xfed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer, uint64 startL2Block)
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

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0xfed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer, uint64 startL2Block)
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0xfed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8.
//
// Solidity: event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer, uint64 startL2Block)
func (_L1Sequencer *L1SequencerFilterer) ParseSequencerUpdated(log types.Log) (*L1SequencerSequencerUpdated, error) {
	event := new(L1SequencerSequencerUpdated)
	if err := _L1Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
