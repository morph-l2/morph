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

// MultipleVersionRollupVerifierMetaData contains all meta data concerning the MultipleVersionRollupVerifier contract.
var MultipleVersionRollupVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getVerifier\",\"inputs\":[{\"name\":\"_batchIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_rollup\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"startBatchIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"legacyVerifiers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"startBatchIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"legacyVerifiersLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollup\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateVerifier\",\"inputs\":[{\"name\":\"_startBatchIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyAggregateProof\",\"inputs\":[{\"name\":\"_batchIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_aggrProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_publicInputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateVerifier\",\"inputs\":[{\"name\":\"startBatchIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610d59380380610d5983398101604081905261002f91610116565b610038336100c6565b6001600160a01b0381166100925760405162461bcd60e51b815260206004820152601560248201527f7a65726f20766572696669657220616464726573730000000000000000000000604482015260640160405180910390fd5b600280546001600160a01b039092166801000000000000000002600160401b600160e01b0319909216919091179055610146565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561012857600080fd5b81516001600160a01b038116811461013f57600080fd5b9392505050565b610c04806101556000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063b57919e811610081578063cc780aa11161005b578063cc780aa1146101f7578063ef4b52271461020a578063f2fde38b1461021d57600080fd5b8063b57919e8146101b3578063c4d66de8146101c4578063cb23bcb5146101d757600080fd5b80634cf536b3116100b25780634cf536b314610153578063715018a61461018b5780638da5cb5b1461019557600080fd5b80631a2c3cde146100ce5780633561bc271461011b575b600080fd5b6100e16100dc366004610a30565b610230565b6040805167ffffffffffffffff909316835273ffffffffffffffffffffffffffffffffffffffff9091166020830152015b60405180910390f35b61012e610129366004610a30565b61027e565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610112565b6002546100e19067ffffffffffffffff81169068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1682565b610193610379565b005b60005473ffffffffffffffffffffffffffffffffffffffff1661012e565b600154604051908152602001610112565b6101936101d2366004610a72565b61038d565b60035461012e9073ffffffffffffffffffffffffffffffffffffffff1681565b610193610205366004610a94565b610461565b610193610218366004610b16565b6104fc565b61019361022b366004610a72565b610883565b6001818154811061024057600080fd5b60009182526020909120015467ffffffffffffffff8116915068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805180820190915260025467ffffffffffffffff81168083526801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff1660208301526000919083101561036f57600154805b801561036c576001808203815481106102ec576102ec610b58565b60009182526020918290206040805180820190915291015467ffffffffffffffff81168083526801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff1692820192909252935085101561036c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016102d1565b50505b6020015192915050565b61038161093a565b61038b60006109bb565b565b61039561093a565b60035473ffffffffffffffffffffffffffffffffffffffff161561041a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f696e697469616c697a656400000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061046c8561027e565b6040517f6b40634100000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff821690636b406341906104c590879087908790600401610b87565b60006040518083038186803b1580156104dd57600080fd5b505afa1580156104f1573d6000803e3d6000fd5b505050505050505050565b61050461093a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663059def616040518163ffffffff1660e01b8152600401602060405180830381865afa158015610571573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105959190610bde565b8267ffffffffffffffff1611610607576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f737461727420626174636820696e6465782066696e616c697a656400000000006044820152606401610411565b6040805180820190915260025467ffffffffffffffff8082168084526801000000000000000090920473ffffffffffffffffffffffffffffffffffffffff166020840152841610156106b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f737461727420626174636820696e64657820746f6f20736d616c6c00000000006044820152606401610411565b73ffffffffffffffffffffffffffffffffffffffff8216610732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f7a65726f207665726966696572206164647265737300000000000000000000006044820152606401610411565b805167ffffffffffffffff808516911610156107de5760018054808201825560009190915281517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69091018054602084015173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090911667ffffffffffffffff93841617179055831681525b73ffffffffffffffffffffffffffffffffffffffff8216602082810182905282516002805467ffffffffffffffff9283167fffffffff0000000000000000000000000000000000000000000000000000000090911617680100000000000000008502179055604080519187168252918101929092527f1363b06925d4266686ad6ab546259321a7ed3cc0bcc55ada2c6431a754b3b4e2910160405180910390a1505050565b61088b61093a565b73ffffffffffffffffffffffffffffffffffffffff811661092e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610411565b610937816109bb565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461038b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610a4257600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a6d57600080fd5b919050565b600060208284031215610a8457600080fd5b610a8d82610a49565b9392505050565b60008060008060608587031215610aaa57600080fd5b84359350602085013567ffffffffffffffff80821115610ac957600080fd5b818701915087601f830112610add57600080fd5b813581811115610aec57600080fd5b886020828501011115610afe57600080fd5b95986020929092019750949560400135945092505050565b60008060408385031215610b2957600080fd5b823567ffffffffffffffff81168114610b4157600080fd5b9150610b4f60208401610a49565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6040815282604082015282846060830137600060608483010152600060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601168301019050826020830152949350505050565b600060208284031215610bf057600080fd5b505191905056fea164736f6c6343000810000a",
}

// MultipleVersionRollupVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MultipleVersionRollupVerifierMetaData.ABI instead.
var MultipleVersionRollupVerifierABI = MultipleVersionRollupVerifierMetaData.ABI

// MultipleVersionRollupVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultipleVersionRollupVerifierMetaData.Bin instead.
var MultipleVersionRollupVerifierBin = MultipleVersionRollupVerifierMetaData.Bin

// DeployMultipleVersionRollupVerifier deploys a new Ethereum contract, binding an instance of MultipleVersionRollupVerifier to it.
func DeployMultipleVersionRollupVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _verifier common.Address) (common.Address, *types.Transaction, *MultipleVersionRollupVerifier, error) {
	parsed, err := MultipleVersionRollupVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultipleVersionRollupVerifierBin), backend, _verifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultipleVersionRollupVerifier{MultipleVersionRollupVerifierCaller: MultipleVersionRollupVerifierCaller{contract: contract}, MultipleVersionRollupVerifierTransactor: MultipleVersionRollupVerifierTransactor{contract: contract}, MultipleVersionRollupVerifierFilterer: MultipleVersionRollupVerifierFilterer{contract: contract}}, nil
}

// MultipleVersionRollupVerifier is an auto generated Go binding around an Ethereum contract.
type MultipleVersionRollupVerifier struct {
	MultipleVersionRollupVerifierCaller     // Read-only binding to the contract
	MultipleVersionRollupVerifierTransactor // Write-only binding to the contract
	MultipleVersionRollupVerifierFilterer   // Log filterer for contract events
}

// MultipleVersionRollupVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultipleVersionRollupVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultipleVersionRollupVerifierSession struct {
	Contract     *MultipleVersionRollupVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MultipleVersionRollupVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultipleVersionRollupVerifierCallerSession struct {
	Contract *MultipleVersionRollupVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MultipleVersionRollupVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultipleVersionRollupVerifierTransactorSession struct {
	Contract     *MultipleVersionRollupVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MultipleVersionRollupVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierRaw struct {
	Contract *MultipleVersionRollupVerifier // Generic contract binding to access the raw methods on
}

// MultipleVersionRollupVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierCallerRaw struct {
	Contract *MultipleVersionRollupVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MultipleVersionRollupVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierTransactorRaw struct {
	Contract *MultipleVersionRollupVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultipleVersionRollupVerifier creates a new instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifier(address common.Address, backend bind.ContractBackend) (*MultipleVersionRollupVerifier, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifier{MultipleVersionRollupVerifierCaller: MultipleVersionRollupVerifierCaller{contract: contract}, MultipleVersionRollupVerifierTransactor: MultipleVersionRollupVerifierTransactor{contract: contract}, MultipleVersionRollupVerifierFilterer: MultipleVersionRollupVerifierFilterer{contract: contract}}, nil
}

// NewMultipleVersionRollupVerifierCaller creates a new read-only instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierCaller(address common.Address, caller bind.ContractCaller) (*MultipleVersionRollupVerifierCaller, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierCaller{contract: contract}, nil
}

// NewMultipleVersionRollupVerifierTransactor creates a new write-only instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MultipleVersionRollupVerifierTransactor, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierTransactor{contract: contract}, nil
}

// NewMultipleVersionRollupVerifierFilterer creates a new log filterer instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MultipleVersionRollupVerifierFilterer, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierFilterer{contract: contract}, nil
}

// bindMultipleVersionRollupVerifier binds a generic wrapper to an already deployed contract.
func bindMultipleVersionRollupVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultipleVersionRollupVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultipleVersionRollupVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.contract.Transact(opts, method, params...)
}

// GetVerifier is a free data retrieval call binding the contract method 0x3561bc27.
//
// Solidity: function getVerifier(uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) GetVerifier(opts *bind.CallOpts, _batchIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "getVerifier", _batchIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x3561bc27.
//
// Solidity: function getVerifier(uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) GetVerifier(_batchIndex *big.Int) (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.GetVerifier(&_MultipleVersionRollupVerifier.CallOpts, _batchIndex)
}

// GetVerifier is a free data retrieval call binding the contract method 0x3561bc27.
//
// Solidity: function getVerifier(uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) GetVerifier(_batchIndex *big.Int) (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.GetVerifier(&_MultipleVersionRollupVerifier.CallOpts, _batchIndex)
}

// LatestVerifier is a free data retrieval call binding the contract method 0x4cf536b3.
//
// Solidity: function latestVerifier() view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LatestVerifier(opts *bind.CallOpts) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "latestVerifier")

	outstruct := new(struct {
		StartBatchIndex uint64
		Verifier        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LatestVerifier is a free data retrieval call binding the contract method 0x4cf536b3.
//
// Solidity: function latestVerifier() view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LatestVerifier() (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LatestVerifier(&_MultipleVersionRollupVerifier.CallOpts)
}

// LatestVerifier is a free data retrieval call binding the contract method 0x4cf536b3.
//
// Solidity: function latestVerifier() view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LatestVerifier() (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LatestVerifier(&_MultipleVersionRollupVerifier.CallOpts)
}

// LegacyVerifiers is a free data retrieval call binding the contract method 0x1a2c3cde.
//
// Solidity: function legacyVerifiers(uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LegacyVerifiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "legacyVerifiers", arg0)

	outstruct := new(struct {
		StartBatchIndex uint64
		Verifier        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LegacyVerifiers is a free data retrieval call binding the contract method 0x1a2c3cde.
//
// Solidity: function legacyVerifiers(uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LegacyVerifiers(arg0 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiers(&_MultipleVersionRollupVerifier.CallOpts, arg0)
}

// LegacyVerifiers is a free data retrieval call binding the contract method 0x1a2c3cde.
//
// Solidity: function legacyVerifiers(uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LegacyVerifiers(arg0 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiers(&_MultipleVersionRollupVerifier.CallOpts, arg0)
}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0xb57919e8.
//
// Solidity: function legacyVerifiersLength() view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LegacyVerifiersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "legacyVerifiersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0xb57919e8.
//
// Solidity: function legacyVerifiersLength() view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LegacyVerifiersLength() (*big.Int, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiersLength(&_MultipleVersionRollupVerifier.CallOpts)
}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0xb57919e8.
//
// Solidity: function legacyVerifiersLength() view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LegacyVerifiersLength() (*big.Int, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiersLength(&_MultipleVersionRollupVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Owner() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Owner(&_MultipleVersionRollupVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) Owner() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Owner(&_MultipleVersionRollupVerifier.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Rollup() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Rollup(&_MultipleVersionRollupVerifier.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) Rollup() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Rollup(&_MultipleVersionRollupVerifier.CallOpts)
}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0xcc780aa1.
//
// Solidity: function verifyAggregateProof(uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) VerifyAggregateProof(opts *bind.CallOpts, _batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "verifyAggregateProof", _batchIndex, _aggrProof, _publicInputHash)

	if err != nil {
		return err
	}

	return err

}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0xcc780aa1.
//
// Solidity: function verifyAggregateProof(uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) VerifyAggregateProof(_batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	return _MultipleVersionRollupVerifier.Contract.VerifyAggregateProof(&_MultipleVersionRollupVerifier.CallOpts, _batchIndex, _aggrProof, _publicInputHash)
}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0xcc780aa1.
//
// Solidity: function verifyAggregateProof(uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) VerifyAggregateProof(_batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	return _MultipleVersionRollupVerifier.Contract.VerifyAggregateProof(&_MultipleVersionRollupVerifier.CallOpts, _batchIndex, _aggrProof, _publicInputHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "initialize", _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Initialize(_rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.Initialize(&_MultipleVersionRollupVerifier.TransactOpts, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) Initialize(_rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.Initialize(&_MultipleVersionRollupVerifier.TransactOpts, _rollup)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.RenounceOwnership(&_MultipleVersionRollupVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.RenounceOwnership(&_MultipleVersionRollupVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.TransferOwnership(&_MultipleVersionRollupVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.TransferOwnership(&_MultipleVersionRollupVerifier.TransactOpts, newOwner)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0xef4b5227.
//
// Solidity: function updateVerifier(uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) UpdateVerifier(opts *bind.TransactOpts, _startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "updateVerifier", _startBatchIndex, _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0xef4b5227.
//
// Solidity: function updateVerifier(uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) UpdateVerifier(_startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.UpdateVerifier(&_MultipleVersionRollupVerifier.TransactOpts, _startBatchIndex, _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0xef4b5227.
//
// Solidity: function updateVerifier(uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) UpdateVerifier(_startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.UpdateVerifier(&_MultipleVersionRollupVerifier.TransactOpts, _startBatchIndex, _verifier)
}

// MultipleVersionRollupVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierOwnershipTransferredIterator struct {
	Event *MultipleVersionRollupVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultipleVersionRollupVerifierOwnershipTransferred)
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
		it.Event = new(MultipleVersionRollupVerifierOwnershipTransferred)
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
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultipleVersionRollupVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MultipleVersionRollupVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultipleVersionRollupVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierOwnershipTransferredIterator{contract: _MultipleVersionRollupVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultipleVersionRollupVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultipleVersionRollupVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultipleVersionRollupVerifierOwnershipTransferred)
				if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*MultipleVersionRollupVerifierOwnershipTransferred, error) {
	event := new(MultipleVersionRollupVerifierOwnershipTransferred)
	if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultipleVersionRollupVerifierUpdateVerifierIterator is returned from FilterUpdateVerifier and is used to iterate over the raw logs and unpacked data for UpdateVerifier events raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierUpdateVerifierIterator struct {
	Event *MultipleVersionRollupVerifierUpdateVerifier // Event containing the contract specifics and raw log

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
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultipleVersionRollupVerifierUpdateVerifier)
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
		it.Event = new(MultipleVersionRollupVerifierUpdateVerifier)
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
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultipleVersionRollupVerifierUpdateVerifier represents a UpdateVerifier event raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierUpdateVerifier struct {
	StartBatchIndex *big.Int
	Verifier        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifier is a free log retrieval operation binding the contract event 0x1363b06925d4266686ad6ab546259321a7ed3cc0bcc55ada2c6431a754b3b4e2.
//
// Solidity: event UpdateVerifier(uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) FilterUpdateVerifier(opts *bind.FilterOpts) (*MultipleVersionRollupVerifierUpdateVerifierIterator, error) {

	logs, sub, err := _MultipleVersionRollupVerifier.contract.FilterLogs(opts, "UpdateVerifier")
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierUpdateVerifierIterator{contract: _MultipleVersionRollupVerifier.contract, event: "UpdateVerifier", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifier is a free log subscription operation binding the contract event 0x1363b06925d4266686ad6ab546259321a7ed3cc0bcc55ada2c6431a754b3b4e2.
//
// Solidity: event UpdateVerifier(uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) WatchUpdateVerifier(opts *bind.WatchOpts, sink chan<- *MultipleVersionRollupVerifierUpdateVerifier) (event.Subscription, error) {

	logs, sub, err := _MultipleVersionRollupVerifier.contract.WatchLogs(opts, "UpdateVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultipleVersionRollupVerifierUpdateVerifier)
				if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
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

// ParseUpdateVerifier is a log parse operation binding the contract event 0x1363b06925d4266686ad6ab546259321a7ed3cc0bcc55ada2c6431a754b3b4e2.
//
// Solidity: event UpdateVerifier(uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) ParseUpdateVerifier(log types.Log) (*MultipleVersionRollupVerifierUpdateVerifier, error) {
	event := new(MultipleVersionRollupVerifierUpdateVerifier)
	if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
