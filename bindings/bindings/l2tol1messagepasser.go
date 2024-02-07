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

// L2ToL1MessagePasserMetaData contains all meta data concerning the L2ToL1MessagePasser contract.
var L2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"AppendMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"appendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafNodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5060016080525f60a081905260c05261002661002e565b602155610114565b6020545f90819081805b60208110156100f7578083901c600116600103610094575f816020811061006157610061610100565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506100c1565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160408051601f1981840301815291905280516020909101209150600101610038565b50919392505050565b634e487b7160e01b5f52603260045260245ffd5b60805160a05160c051610ac861013e5f395f61020d01525f6101e401525f6101bb0152610ac85ff3fe608060405234801561000f575f80fd5b506004361061006f575f3560e01c806389c09d381161004d57806389c09d38146100d1578063b58343bb146100d9578063d4b9f4fa146100e2575f80fd5b8063340735f71461007357806354fd4d501461009b578063600a2e77146100b0575b5f80fd5b6100866100813660046106ff565b6100eb565b60405190151581526020015b60405180910390f35b6100a36101b4565b60405161009291906107b5565b6100c36100be366004610805565b610257565b604051908152602001610092565b6100c3610331565b6100c360205481565b6100c360215481565b5f84815b60208110156101a8578085901c600116600103610155578581602081106101185761011861081c565b602002015182604051602001610138929190918252602082015260400190565b6040516020818303038152906040528051906020012091506101a0565b818682602081106101685761016861081c565b6020020151604051602001610187929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b6001016100ef565b50909114949350505050565b60606101df7f0000000000000000000000000000000000000000000000000000000000000000610421565b6102087f0000000000000000000000000000000000000000000000000000000000000000610421565b6102317f0000000000000000000000000000000000000000000000000000000000000000610421565b60405160200161024393929190610849565b604051602081830303815290604052905090565b5f33735300000000000000000000000000000000000007146102d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6f6e6c79206d657373656e676572000000000000000000000000000000000000604482015260640160405180910390fd5b60208054604080519182529181018490527ffaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693910160405180910390a161031e826104dd565b610326610331565b602181905592915050565b6020545f90819081805b6020811015610418578083901c600116600103610397575f81602081106103645761036461081c565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506103c4565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120915060010161033b565b50919392505050565b60605f61042d836105f0565b60010190505f8167ffffffffffffffff81111561044c5761044c6106d2565b6040519080825280601f01601f191660200182016040528015610476576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461048057509392505050565b8060016104ec60206002610a09565b6104f69190610a1b565b60205410610530576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60205f815461053f90610a2e565b918290555090505f5b60208110156105e2578082901c60011660010361057a57825f82602081106105725761057261081c565b015550505050565b5f816020811061058c5761058c61081c565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209250600101610548565b506105eb610a65565b505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610638577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610664576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061068257662386f26fc10000830492506010015b6305f5e100831061069a576305f5e100830492506008015b61271083106106ae57612710830492506004015b606483106106c0576064830492506002015b600a83106106cc576001015b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f806104608587031215610713575f80fd5b84359350602086603f870112610727575f80fd5b604051610400810181811067ffffffffffffffff8211171561074b5761074b6106d2565b60405280610420880189811115610760575f80fd5b602089015b8181101561077c5780358352918401918401610765565b509699919850509435956104400135949350505050565b5f5b838110156107ad578181015183820152602001610795565b50505f910152565b602081525f82518060208401526107d3816040850160208701610793565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b5f60208284031215610815575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f845161085a818460208901610793565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610896816001850160208a01610793565b600192019182015283516108b1816002840160208801610793565b0160020195945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b600181815b8085111561094457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561092a5761092a6108be565b8085161561093757918102915b93841c93908002906108f0565b509250929050565b5f8261095a575060016106cc565b8161096657505f6106cc565b816001811461097c5760028114610986576109a2565b60019150506106cc565b60ff841115610997576109976108be565b50506001821b6106cc565b5060208310610133831016604e8410600b84101617156109c5575081810a6106cc565b6109cf83836108eb565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610a0157610a016108be565b029392505050565b5f610a14838361094c565b9392505050565b818103818111156106cc576106cc6108be565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a5e57610a5e6108be565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea2646970667358221220834eb328736e11f6da8715351fff5b777c9bda9d117b481e336b827449ca455e64736f6c63430008180033",
}

// L2ToL1MessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ToL1MessagePasserMetaData.ABI instead.
var L2ToL1MessagePasserABI = L2ToL1MessagePasserMetaData.ABI

// L2ToL1MessagePasserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ToL1MessagePasserMetaData.Bin instead.
var L2ToL1MessagePasserBin = L2ToL1MessagePasserMetaData.Bin

// DeployL2ToL1MessagePasser deploys a new Ethereum contract, binding an instance of L2ToL1MessagePasser to it.
func DeployL2ToL1MessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ToL1MessagePasser, error) {
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ToL1MessagePasserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// L2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type L2ToL1MessagePasser struct {
	L2ToL1MessagePasserCaller     // Read-only binding to the contract
	L2ToL1MessagePasserTransactor // Write-only binding to the contract
	L2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// L2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ToL1MessagePasserSession struct {
	Contract     *L2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ToL1MessagePasserCallerSession struct {
	Contract *L2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// L2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ToL1MessagePasserTransactorSession struct {
	Contract     *L2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ToL1MessagePasserRaw struct {
	Contract *L2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// L2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCallerRaw struct {
	Contract *L2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// L2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactorRaw struct {
	Contract *L2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ToL1MessagePasser creates a new instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasser(address common.Address, backend bind.ContractBackend) (*L2ToL1MessagePasser, error) {
	contract, err := bindL2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewL2ToL1MessagePasserCaller creates a new read-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserCaller(address common.Address, caller bind.ContractCaller) (*L2ToL1MessagePasserCaller, error) {
	contract, err := bindL2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewL2ToL1MessagePasserTransactor creates a new write-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ToL1MessagePasserTransactor, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewL2ToL1MessagePasserFilterer creates a new log filterer instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ToL1MessagePasserFilterer, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindL2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindL2ToL1MessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "getTreeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) GetTreeRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.GetTreeRoot(&_L2ToL1MessagePasser.CallOpts)
}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) GetTreeRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.GetTreeRoot(&_L2ToL1MessagePasser.CallOpts)
}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) LeafNodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "leafNodesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) LeafNodesCount() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.LeafNodesCount(&_L2ToL1MessagePasser.CallOpts)
}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) LeafNodesCount() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.LeafNodesCount(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.VerifyMerkleProof(&_L2ToL1MessagePasser.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.VerifyMerkleProof(&_L2ToL1MessagePasser.CallOpts, leafHash, smtProof, index, root)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// AppendMessage is a paid mutator transaction binding the contract method 0x600a2e77.
//
// Solidity: function appendMessage(bytes32 _messageHash) returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) AppendMessage(opts *bind.TransactOpts, _messageHash [32]byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "appendMessage", _messageHash)
}

// AppendMessage is a paid mutator transaction binding the contract method 0x600a2e77.
//
// Solidity: function appendMessage(bytes32 _messageHash) returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) AppendMessage(_messageHash [32]byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.AppendMessage(&_L2ToL1MessagePasser.TransactOpts, _messageHash)
}

// AppendMessage is a paid mutator transaction binding the contract method 0x600a2e77.
//
// Solidity: function appendMessage(bytes32 _messageHash) returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) AppendMessage(_messageHash [32]byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.AppendMessage(&_L2ToL1MessagePasser.TransactOpts, _messageHash)
}

// L2ToL1MessagePasserAppendMessageIterator is returned from FilterAppendMessage and is used to iterate over the raw logs and unpacked data for AppendMessage events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserAppendMessageIterator struct {
	Event *L2ToL1MessagePasserAppendMessage // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserAppendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserAppendMessage)
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
		it.Event = new(L2ToL1MessagePasserAppendMessage)
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
func (it *L2ToL1MessagePasserAppendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserAppendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserAppendMessage represents a AppendMessage event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserAppendMessage struct {
	Index       *big.Int
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAppendMessage is a free log retrieval operation binding the contract event 0xfaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterAppendMessage(opts *bind.FilterOpts) (*L2ToL1MessagePasserAppendMessageIterator, error) {

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "AppendMessage")
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserAppendMessageIterator{contract: _L2ToL1MessagePasser.contract, event: "AppendMessage", logs: logs, sub: sub}, nil
}

// WatchAppendMessage is a free log subscription operation binding the contract event 0xfaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchAppendMessage(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserAppendMessage) (event.Subscription, error) {

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "AppendMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserAppendMessage)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "AppendMessage", log); err != nil {
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

// ParseAppendMessage is a log parse operation binding the contract event 0xfaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseAppendMessage(log types.Log) (*L2ToL1MessagePasserAppendMessage, error) {
	event := new(L2ToL1MessagePasserAppendMessage)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "AppendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
