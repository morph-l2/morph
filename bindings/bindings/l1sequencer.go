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

// L1SequencerSequencerRecord is an auto generated low-level Go binding around an user-defined struct.
type L1SequencerSequencerRecord struct {
	StartL2Block  uint64
	SequencerAddr common.Address
}

// L1SequencerMetaData contains all meta data concerning the L1Sequencer contract.
var L1SequencerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerAt\",\"inputs\":[{\"name\":\"l2Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerHistory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structL1Sequencer.SequencerRecord[]\",\"components\":[{\"name\":\"startL2Block\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequencerAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerHistoryLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeHistory\",\"inputs\":[{\"name\":\"firstSequencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"upgradeL2Block\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sequencerHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"startL2Block\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequencerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSequencer\",\"inputs\":[{\"name\":\"newSequencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startL2Block\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerUpdated\",\"inputs\":[{\"name\":\"oldSequencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSequencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startL2Block\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b506111968061001d5f395ff3fe608060405234801561000f575f80fd5b50600436106100cf575f3560e01c8063761a90fd1161007d578063f151ce9e11610058578063f151ce9e146101ee578063f198e27f14610201578063f2fde38b14610214575f80fd5b8063761a90fd146101aa5780638da5cb5b146101bd578063c4d66de8146101db575f80fd5b80636628aea1116100ad5780636628aea1146101435780636d8ce3d214610158578063715018a6146101a0575f80fd5b80633d5767ce146100d35780633ef5e8cc146100e95780634d96a90a14610116575b5f80fd5b6065546040519081526020015b60405180910390f35b6066546100fd9067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100e0565b61011e610227565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e0565b61014b6102e9565b6040516100e09190610f9d565b61016b61016636600461100b565b610376565b6040805167ffffffffffffffff909316835273ffffffffffffffffffffffffffffffffffffffff9091166020830152016100e0565b6101a86103c2565b005b6101a86101b8366004611061565b6103d5565b60335473ffffffffffffffffffffffffffffffffffffffff1661011e565b6101a86101e9366004611092565b6106a7565b61011e6101fc3660046110b2565b6108ba565b6101a861020f366004611061565b610ab7565b6101a8610222366004611092565b610cb7565b6065545f90610297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6e6f2073657175656e63657220636f6e6669677572656400000000000000000060448201526064015b60405180910390fd5b606580546102a7906001906110f8565b815481106102b7576102b7611111565b5f9182526020909120015468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16919050565b60606065805480602002602001604051908101604052809291908181526020015f905b8282101561036d575f848152602090819020604080518082019091529084015467ffffffffffffffff8116825268010000000000000000900473ffffffffffffffffffffffffffffffffffffffff168183015282526001909201910161030c565b50505050905090565b60658181548110610385575f80fd5b5f9182526020909120015467ffffffffffffffff8116915068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6103ca610d6e565b6103d35f610def565b565b6103dd610d6e565b73ffffffffffffffffffffffffffffffffffffffff821661045a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e76616c696420616464726573730000000000000000000000000000000000604482015260640161028e565b6065546104c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6e6f7420696e697469616c697a65640000000000000000000000000000000000604482015260640161028e565b606580546104d3906001906110f8565b815481106104e3576104e3611111565b5f9182526020909120015467ffffffffffffffff9081169082161161058a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f73746172744c32426c6f636b206d75737420626520677265617465722074686160448201527f6e206c617374207265636f726400000000000000000000000000000000000000606482015260840161028e565b606580545f919061059d906001906110f8565b815481106105ad576105ad611111565b5f9182526020808320919091015460408051808201825267ffffffffffffffff87811680835273ffffffffffffffffffffffffffffffffffffffff8a8116848801818152606580546001810182559a5294517f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c790990180549551999094167fffffffff00000000000000000000000000000000000000000000000000000000909516949094176801000000000000000098821689021790925592519283529490920490931693509183917ffed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8910160405180910390a3505050565b5f54610100900460ff16158080156106c557505f54600160ff909116105b806106de5750303b1580156106de57505f5460ff166001145b61076a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161028e565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107c6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8216610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f696e76616c6964206f776e657200000000000000000000000000000000000000604482015260640161028e565b61084b610e65565b61085482610def565b80156108b6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6065545f9080610926576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6e6f2073657175656e63657220636f6e66696775726564000000000000000000604482015260640161028e565b5f806109336001846110f8565b90505f5b8183116109d2575f600261094b848661113e565b6109559190611151565b90508667ffffffffffffffff166065828154811061097557610975611111565b5f9182526020909120015467ffffffffffffffff16116109b15780915082810361099f57506109d2565b6109aa81600161113e565b93506109cc565b805f036109be57506109d2565b6109c96001826110f8565b92505b50610937565b8567ffffffffffffffff16606582815481106109f0576109f0611111565b5f9182526020909120015467ffffffffffffffff161115610a6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6e6f2073657175656e6365722061742068656967687400000000000000000000604482015260640161028e565b60658181548110610a8057610a80611111565b5f9182526020909120015468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff169695505050505050565b610abf610d6e565b60655415610b29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f616c726561647920696e697469616c697a656400000000000000000000000000604482015260640161028e565b73ffffffffffffffffffffffffffffffffffffffff8216610ba6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e76616c696420616464726573730000000000000000000000000000000000604482015260640161028e565b60408051808201825267ffffffffffffffff83811680835273ffffffffffffffffffffffffffffffffffffffff8681166020808601828152606580546001810182555f91825297517f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c79098018054925190951668010000000000000000027fffffffff00000000000000000000000000000000000000000000000000000000909216979096169690961795909517909155606680547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001683179055935190815290917ffed767db50732333bba543b785430d53a3a836d71064a68ae91809e50eca7bb8910160405180910390a35050565b610cbf610d6e565b73ffffffffffffffffffffffffffffffffffffffff8116610d62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161028e565b610d6b81610def565b50565b60335473ffffffffffffffffffffffffffffffffffffffff1633146103d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161028e565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610efb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161028e565b6103d35f54610100900460ff16610f94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161028e565b6103d333610def565b602080825282518282018190525f919060409081850190868401855b82811015610ffe578151805167ffffffffffffffff16855286015173ffffffffffffffffffffffffffffffffffffffff16868501529284019290850190600101610fb9565b5091979650505050505050565b5f6020828403121561101b575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611045575f80fd5b919050565b803567ffffffffffffffff81168114611045575f80fd5b5f8060408385031215611072575f80fd5b61107b83611022565b91506110896020840161104a565b90509250929050565b5f602082840312156110a2575f80fd5b6110ab82611022565b9392505050565b5f602082840312156110c2575f80fd5b6110ab8261104a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561110b5761110b6110cb565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082018082111561110b5761110b6110cb565b5f82611184577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a",
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

// ActiveHeight is a free data retrieval call binding the contract method 0x3ef5e8cc.
//
// Solidity: function activeHeight() view returns(uint64)
func (_L1Sequencer *L1SequencerCaller) ActiveHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "activeHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ActiveHeight is a free data retrieval call binding the contract method 0x3ef5e8cc.
//
// Solidity: function activeHeight() view returns(uint64)
func (_L1Sequencer *L1SequencerSession) ActiveHeight() (uint64, error) {
	return _L1Sequencer.Contract.ActiveHeight(&_L1Sequencer.CallOpts)
}

// ActiveHeight is a free data retrieval call binding the contract method 0x3ef5e8cc.
//
// Solidity: function activeHeight() view returns(uint64)
func (_L1Sequencer *L1SequencerCallerSession) ActiveHeight() (uint64, error) {
	return _L1Sequencer.Contract.ActiveHeight(&_L1Sequencer.CallOpts)
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
func (_L1Sequencer *L1SequencerCaller) GetSequencerHistory(opts *bind.CallOpts) ([]L1SequencerSequencerRecord, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerHistory")

	if err != nil {
		return *new([]L1SequencerSequencerRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]L1SequencerSequencerRecord)).(*[]L1SequencerSequencerRecord)

	return out0, err

}

// GetSequencerHistory is a free data retrieval call binding the contract method 0x6628aea1.
//
// Solidity: function getSequencerHistory() view returns((uint64,address)[])
func (_L1Sequencer *L1SequencerSession) GetSequencerHistory() ([]L1SequencerSequencerRecord, error) {
	return _L1Sequencer.Contract.GetSequencerHistory(&_L1Sequencer.CallOpts)
}

// GetSequencerHistory is a free data retrieval call binding the contract method 0x6628aea1.
//
// Solidity: function getSequencerHistory() view returns((uint64,address)[])
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerHistory() ([]L1SequencerSequencerRecord, error) {
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

// InitializeHistory is a paid mutator transaction binding the contract method 0xf198e27f.
//
// Solidity: function initializeHistory(address firstSequencer, uint64 upgradeL2Block) returns()
func (_L1Sequencer *L1SequencerTransactor) InitializeHistory(opts *bind.TransactOpts, firstSequencer common.Address, upgradeL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "initializeHistory", firstSequencer, upgradeL2Block)
}

// InitializeHistory is a paid mutator transaction binding the contract method 0xf198e27f.
//
// Solidity: function initializeHistory(address firstSequencer, uint64 upgradeL2Block) returns()
func (_L1Sequencer *L1SequencerSession) InitializeHistory(firstSequencer common.Address, upgradeL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.Contract.InitializeHistory(&_L1Sequencer.TransactOpts, firstSequencer, upgradeL2Block)
}

// InitializeHistory is a paid mutator transaction binding the contract method 0xf198e27f.
//
// Solidity: function initializeHistory(address firstSequencer, uint64 upgradeL2Block) returns()
func (_L1Sequencer *L1SequencerTransactorSession) InitializeHistory(firstSequencer common.Address, upgradeL2Block uint64) (*types.Transaction, error) {
	return _L1Sequencer.Contract.InitializeHistory(&_L1Sequencer.TransactOpts, firstSequencer, upgradeL2Block)
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
