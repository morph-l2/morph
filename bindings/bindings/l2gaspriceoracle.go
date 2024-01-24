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

// L2GasPriceOracleMetaData contains all meta data concerning the L2GasPriceOracle contract.
var L2GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateIntrinsicGasFee\",\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateCrossDomainMessageFee\",\"inputs\":[{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_txGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_txGasContractCreation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_zeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_nonZeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"intrinsicParams\",\"inputs\":[],\"outputs\":[{\"name\":\"txGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txGasContractCreation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"zeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonZeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2BaseFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntrinsicParams\",\"inputs\":[{\"name\":\"_txGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_txGasContractCreation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_zeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_nonZeroGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL2BaseFee\",\"inputs\":[{\"name\":\"_newL2BaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntrinsicParamsUpdated\",\"inputs\":[{\"name\":\"txGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"txGasContractCreation\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"zeroGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonZeroGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2BaseFeeUpdated\",\"inputs\":[{\"name\":\"oldL2BaseFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newL2BaseFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateWhitelist\",\"inputs\":[{\"name\":\"_oldWhitelist\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_newWhitelist\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610e0a806100ec6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063d7704bae11610076578063e172d3a11161005b578063e172d3a1146101e2578063e3176bd5146101f5578063f2fde38b146101fe57600080fd5b8063d7704bae146101ae578063d99bc80e146101cf57600080fd5b8063715018a6116100a7578063715018a61461016b5780638da5cb5b14610173578063accf9a601461019b57600080fd5b80633366ff72146100c357806364431a27146100d8575b600080fd5b6100d66100d1366004610b5c565b610211565b005b6066546101329067ffffffffffffffff808216916801000000000000000081048216917001000000000000000000000000000000008204811691780100000000000000000000000000000000000000000000000090041684565b6040805167ffffffffffffffff958616815293851660208501529184169183019190915290911660608201526080015b60405180910390f35b6100d66103b8565b60335460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610162565b6100d66101a9366004610b5c565b6103cc565b6101c16101bc366004610bb0565b6104ed565b604051908152602001610162565b6100d66101dd366004610bb0565b610503565b6101c16101f0366004610bf8565b610550565b6101c160655481565b6100d661020c366004610cc7565b610652565b600054610100900460ff16158080156102315750600054600160ff909116105b8061024b5750303b15801561024b575060005460ff166001145b6102dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561033a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610342610709565b61034e858585856107a8565b80156103b157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6103c06109a7565b6103ca6000610a28565b565b6103d46109a7565b604080516080808201835267ffffffffffffffff87811680845287821660208086018290528884168688018190529388166060968701819052606680547fffffffffffffffffffffffffffffffff00000000000000000000000000000000168517680100000000000000008502176fffffffffffffffffffffffffffffffff16700100000000000000000000000000000000870277ffffffffffffffffffffffffffffffffffffffffffffffff16177801000000000000000000000000000000000000000000000000830217905587519384529083019190915294810191909152918201929092527f92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854910160405180910390a150505050565b6000606554826104fd9190610d33565b92915050565b61050b6109a7565b606580549082905560408051828152602081018490527f230bc8094d790356a078817d156f95cc1068e9ff6485359f6a986170f567b63b910160405180910390a15050565b606654815160009167ffffffffffffffff8082169270010000000000000000000000000000000083048216927801000000000000000000000000000000000000000000000000900490911690839015610649576000805b875181101561060e578781815181106105c2576105c2610d70565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016156105fc57816105f881610d9f565b9250505b8061060681610d9f565b9150506105a7565b508381885161061d9190610dd7565b6106279190610d33565b6106318483610d33565b61063b9190610dea565b6106459083610dea565b9150505b95945050505050565b61065a6109a7565b73ffffffffffffffffffffffffffffffffffffffff81166106fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102d3565b61070681610a28565b50565b600054610100900460ff166107a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102d3565b6103ca610a9f565b60008467ffffffffffffffff161161081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7478476173206973207a65726f0000000000000000000000000000000000000060448201526064016102d3565b60008267ffffffffffffffff1611610890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7a65726f476173206973207a65726f000000000000000000000000000000000060448201526064016102d3565b60008167ffffffffffffffff1611610904576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6e6f6e5a65726f476173206973207a65726f000000000000000000000000000060448201526064016102d3565b8367ffffffffffffffff168367ffffffffffffffff16116103d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f7478476173436f6e74726163744372656174696f6e206973206c65737320746860448201527f616e20747847617300000000000000000000000000000000000000000000000060648201526084016102d3565b60335473ffffffffffffffffffffffffffffffffffffffff1633146103ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d3565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102d3565b6103ca33610a28565b803567ffffffffffffffff81168114610b5757600080fd5b919050565b60008060008060808587031215610b7257600080fd5b610b7b85610b3f565b9350610b8960208601610b3f565b9250610b9760408601610b3f565b9150610ba560608601610b3f565b905092959194509250565b600060208284031215610bc257600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215610c0a57600080fd5b813567ffffffffffffffff80821115610c2257600080fd5b818401915084601f830112610c3657600080fd5b813581811115610c4857610c48610bc9565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610c8e57610c8e610bc9565b81604052828152876020848701011115610ca757600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610cd957600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610cfd57600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610d6b57610d6b610d04565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610dd057610dd0610d04565b5060010190565b818103818111156104fd576104fd610d04565b808201808211156104fd576104fd610d0456fea164736f6c6343000810000a",
}

// L2GasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L2GasPriceOracleMetaData.ABI instead.
var L2GasPriceOracleABI = L2GasPriceOracleMetaData.ABI

// L2GasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2GasPriceOracleMetaData.Bin instead.
var L2GasPriceOracleBin = L2GasPriceOracleMetaData.Bin

// DeployL2GasPriceOracle deploys a new Ethereum contract, binding an instance of L2GasPriceOracle to it.
func DeployL2GasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2GasPriceOracle, error) {
	parsed, err := L2GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2GasPriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2GasPriceOracle{L2GasPriceOracleCaller: L2GasPriceOracleCaller{contract: contract}, L2GasPriceOracleTransactor: L2GasPriceOracleTransactor{contract: contract}, L2GasPriceOracleFilterer: L2GasPriceOracleFilterer{contract: contract}}, nil
}

// L2GasPriceOracle is an auto generated Go binding around an Ethereum contract.
type L2GasPriceOracle struct {
	L2GasPriceOracleCaller     // Read-only binding to the contract
	L2GasPriceOracleTransactor // Write-only binding to the contract
	L2GasPriceOracleFilterer   // Log filterer for contract events
}

// L2GasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2GasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2GasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2GasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2GasPriceOracleSession struct {
	Contract     *L2GasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2GasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2GasPriceOracleCallerSession struct {
	Contract *L2GasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2GasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2GasPriceOracleTransactorSession struct {
	Contract     *L2GasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2GasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2GasPriceOracleRaw struct {
	Contract *L2GasPriceOracle // Generic contract binding to access the raw methods on
}

// L2GasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2GasPriceOracleCallerRaw struct {
	Contract *L2GasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L2GasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2GasPriceOracleTransactorRaw struct {
	Contract *L2GasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2GasPriceOracle creates a new instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracle(address common.Address, backend bind.ContractBackend) (*L2GasPriceOracle, error) {
	contract, err := bindL2GasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracle{L2GasPriceOracleCaller: L2GasPriceOracleCaller{contract: contract}, L2GasPriceOracleTransactor: L2GasPriceOracleTransactor{contract: contract}, L2GasPriceOracleFilterer: L2GasPriceOracleFilterer{contract: contract}}, nil
}

// NewL2GasPriceOracleCaller creates a new read-only instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L2GasPriceOracleCaller, error) {
	contract, err := bindL2GasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleCaller{contract: contract}, nil
}

// NewL2GasPriceOracleTransactor creates a new write-only instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L2GasPriceOracleTransactor, error) {
	contract, err := bindL2GasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleTransactor{contract: contract}, nil
}

// NewL2GasPriceOracleFilterer creates a new log filterer instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L2GasPriceOracleFilterer, error) {
	contract, err := bindL2GasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleFilterer{contract: contract}, nil
}

// bindL2GasPriceOracle binds a generic wrapper to an already deployed contract.
func bindL2GasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2GasPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GasPriceOracle *L2GasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GasPriceOracle *L2GasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GasPriceOracle *L2GasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "calculateIntrinsicGasFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) CalculateIntrinsicGasFee(_message []byte) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L2GasPriceOracle.CallOpts, _message)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) CalculateIntrinsicGasFee(_message []byte) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L2GasPriceOracle.CallOpts, _message)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L2GasPriceOracle.CallOpts, _gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L2GasPriceOracle.CallOpts, _gasLimit)
}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) IntrinsicParams(opts *bind.CallOpts) (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "intrinsicParams")

	outstruct := new(struct {
		TxGas                 uint64
		TxGasContractCreation uint64
		ZeroGas               uint64
		NonZeroGas            uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxGas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TxGasContractCreation = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ZeroGas = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.NonZeroGas = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleSession) IntrinsicParams() (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	return _L2GasPriceOracle.Contract.IntrinsicParams(&_L2GasPriceOracle.CallOpts)
}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) IntrinsicParams() (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	return _L2GasPriceOracle.Contract.IntrinsicParams(&_L2GasPriceOracle.CallOpts)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) L2BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "l2BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) L2BaseFee() (*big.Int, error) {
	return _L2GasPriceOracle.Contract.L2BaseFee(&_L2GasPriceOracle.CallOpts)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) L2BaseFee() (*big.Int, error) {
	return _L2GasPriceOracle.Contract.L2BaseFee(&_L2GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleSession) Owner() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Owner(&_L2GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Owner(&_L2GasPriceOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "initialize", _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) Initialize(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.Initialize(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) Initialize(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.Initialize(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.RenounceOwnership(&_L2GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.RenounceOwnership(&_L2GasPriceOracle.TransactOpts)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) SetIntrinsicParams(opts *bind.TransactOpts, _txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "setIntrinsicParams", _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) SetIntrinsicParams(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetIntrinsicParams(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) SetIntrinsicParams(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetIntrinsicParams(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) SetL2BaseFee(opts *bind.TransactOpts, _newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "setL2BaseFee", _newL2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) SetL2BaseFee(_newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetL2BaseFee(&_L2GasPriceOracle.TransactOpts, _newL2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) SetL2BaseFee(_newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetL2BaseFee(&_L2GasPriceOracle.TransactOpts, _newL2BaseFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.TransferOwnership(&_L2GasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.TransferOwnership(&_L2GasPriceOracle.TransactOpts, newOwner)
}

// L2GasPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleInitializedIterator struct {
	Event *L2GasPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleInitialized)
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
		it.Event = new(L2GasPriceOracleInitialized)
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
func (it *L2GasPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleInitialized represents a Initialized event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2GasPriceOracleInitializedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleInitializedIterator{contract: _L2GasPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleInitialized)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseInitialized(log types.Log) (*L2GasPriceOracleInitialized, error) {
	event := new(L2GasPriceOracleInitialized)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleIntrinsicParamsUpdatedIterator is returned from FilterIntrinsicParamsUpdated and is used to iterate over the raw logs and unpacked data for IntrinsicParamsUpdated events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleIntrinsicParamsUpdatedIterator struct {
	Event *L2GasPriceOracleIntrinsicParamsUpdated // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleIntrinsicParamsUpdated)
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
		it.Event = new(L2GasPriceOracleIntrinsicParamsUpdated)
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
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleIntrinsicParamsUpdated represents a IntrinsicParamsUpdated event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleIntrinsicParamsUpdated struct {
	TxGas                 *big.Int
	TxGasContractCreation *big.Int
	ZeroGas               *big.Int
	NonZeroGas            *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterIntrinsicParamsUpdated is a free log retrieval operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterIntrinsicParamsUpdated(opts *bind.FilterOpts) (*L2GasPriceOracleIntrinsicParamsUpdatedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "IntrinsicParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleIntrinsicParamsUpdatedIterator{contract: _L2GasPriceOracle.contract, event: "IntrinsicParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchIntrinsicParamsUpdated is a free log subscription operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchIntrinsicParamsUpdated(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleIntrinsicParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "IntrinsicParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleIntrinsicParamsUpdated)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "IntrinsicParamsUpdated", log); err != nil {
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

// ParseIntrinsicParamsUpdated is a log parse operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseIntrinsicParamsUpdated(log types.Log) (*L2GasPriceOracleIntrinsicParamsUpdated, error) {
	event := new(L2GasPriceOracleIntrinsicParamsUpdated)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "IntrinsicParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleL2BaseFeeUpdatedIterator is returned from FilterL2BaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L2BaseFeeUpdated events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleL2BaseFeeUpdatedIterator struct {
	Event *L2GasPriceOracleL2BaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleL2BaseFeeUpdated)
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
		it.Event = new(L2GasPriceOracleL2BaseFeeUpdated)
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
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleL2BaseFeeUpdated represents a L2BaseFeeUpdated event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleL2BaseFeeUpdated struct {
	OldL2BaseFee *big.Int
	NewL2BaseFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterL2BaseFeeUpdated is a free log retrieval operation binding the contract event 0x230bc8094d790356a078817d156f95cc1068e9ff6485359f6a986170f567b63b.
//
// Solidity: event L2BaseFeeUpdated(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterL2BaseFeeUpdated(opts *bind.FilterOpts) (*L2GasPriceOracleL2BaseFeeUpdatedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "L2BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleL2BaseFeeUpdatedIterator{contract: _L2GasPriceOracle.contract, event: "L2BaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL2BaseFeeUpdated is a free log subscription operation binding the contract event 0x230bc8094d790356a078817d156f95cc1068e9ff6485359f6a986170f567b63b.
//
// Solidity: event L2BaseFeeUpdated(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchL2BaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleL2BaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "L2BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleL2BaseFeeUpdated)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "L2BaseFeeUpdated", log); err != nil {
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

// ParseL2BaseFeeUpdated is a log parse operation binding the contract event 0x230bc8094d790356a078817d156f95cc1068e9ff6485359f6a986170f567b63b.
//
// Solidity: event L2BaseFeeUpdated(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseL2BaseFeeUpdated(log types.Log) (*L2GasPriceOracleL2BaseFeeUpdated, error) {
	event := new(L2GasPriceOracleL2BaseFeeUpdated)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "L2BaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleOwnershipTransferredIterator struct {
	Event *L2GasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleOwnershipTransferred)
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
		it.Event = new(L2GasPriceOracleOwnershipTransferred)
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
func (it *L2GasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2GasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleOwnershipTransferredIterator{contract: _L2GasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleOwnershipTransferred)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*L2GasPriceOracleOwnershipTransferred, error) {
	event := new(L2GasPriceOracleOwnershipTransferred)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleUpdateWhitelistIterator is returned from FilterUpdateWhitelist and is used to iterate over the raw logs and unpacked data for UpdateWhitelist events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleUpdateWhitelistIterator struct {
	Event *L2GasPriceOracleUpdateWhitelist // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleUpdateWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleUpdateWhitelist)
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
		it.Event = new(L2GasPriceOracleUpdateWhitelist)
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
func (it *L2GasPriceOracleUpdateWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleUpdateWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleUpdateWhitelist represents a UpdateWhitelist event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleUpdateWhitelist struct {
	OldWhitelist common.Address
	NewWhitelist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhitelist is a free log retrieval operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterUpdateWhitelist(opts *bind.FilterOpts) (*L2GasPriceOracleUpdateWhitelistIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleUpdateWhitelistIterator{contract: _L2GasPriceOracle.contract, event: "UpdateWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdateWhitelist is a free log subscription operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchUpdateWhitelist(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleUpdateWhitelist) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleUpdateWhitelist)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
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

// ParseUpdateWhitelist is a log parse operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseUpdateWhitelist(log types.Log) (*L2GasPriceOracleUpdateWhitelist, error) {
	event := new(L2GasPriceOracleUpdateWhitelist)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
