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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appendMessage\",\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTreeRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"leafNodesCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMerkleProof\",\"inputs\":[{\"name\":\"leafHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"smtProof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AppendMessage\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MerkleTreeFull\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561001057600080fd5b506001608052600060a081905260c052610028610030565b60215561014a565b602054600090819081805b6020811015610104578083901c60011660010361009857600081602081106100655761006561010d565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506100c5565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806100fc90610123565b91505061003b565b50919392505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161014357634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051610aa5610179600039600061021c015260006101f3015260006101ca0152610aa56000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806389c09d381161005057806389c09d38146100d5578063b58343bb146100dd578063d4b9f4fa146100e657600080fd5b8063340735f71461007757806354fd4d501461009f578063600a2e77146100b4575b600080fd5b61008a6100853660046106f1565b6100ef565b60405190151581526020015b60405180910390f35b6100a76101c3565b60405161009691906107ad565b6100c76100c23660046107fe565b610266565b604051908152602001610096565b6100c7610341565b6100c760205481565b6100c760215481565b600084815b60208110156101b7578085901c60011660010361015a5785816020811061011d5761011d610817565b60200201518260405160200161013d929190918252602082015260400190565b6040516020818303038152906040528051906020012091506101a5565b8186826020811061016d5761016d610817565b602002015160405160200161018c929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806101af81610875565b9150506100f4565b50909114949350505050565b60606101ee7f000000000000000000000000000000000000000000000000000000000000000061041e565b6102177f000000000000000000000000000000000000000000000000000000000000000061041e565b6102407f000000000000000000000000000000000000000000000000000000000000000061041e565b604051602001610252939291906108ad565b604051602081830303815290604052905090565b600033735300000000000000000000000000000000000007146102e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6f6e6c79206d657373656e676572000000000000000000000000000000000000604482015260640160405180910390fd5b60208054604080519182529181018490527ffaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693910160405180910390a161032e826104dc565b610336610341565b602181905592915050565b602054600090819081805b6020811015610415578083901c6001166001036103a9576000816020811061037657610376610817565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506103d6565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604051602081830303815290604052805190602001209150808061040d90610875565b91505061034c565b50919392505050565b6060600061042b836105df565b600101905060008167ffffffffffffffff81111561044b5761044b6106c2565b6040519080825280601f01601f191660200182016040528015610475576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461047f57509392505050565b8060016104eb60206002610a43565b6104f59190610a56565b6020541061052f576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060206000815461054090610875565b9182905550905060005b60208110156105d1578082901c60011660010361057d57826000826020811061057557610575610817565b015550505050565b6000816020811061059057610590610817565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925080806105c990610875565b91505061054a565b506105da610a69565b505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610628577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610654576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061067257662386f26fc10000830492506010015b6305f5e100831061068a576305f5e100830492506008015b612710831061069e57612710830492506004015b606483106106b0576064830492506002015b600a83106106bc576001015b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600080610460858703121561070857600080fd5b84359350602086603f87011261071d57600080fd5b604051610400810181811067ffffffffffffffff82111715610741576107416106c2565b6040528061042088018981111561075757600080fd5b8389015b81811015610772578035835291840191840161075b565b509699919850509435956104400135949350505050565b60005b838110156107a457818101518382015260200161078c565b50506000910152565b60208152600082518060208401526107cc816040850160208701610789565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561081057600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108a6576108a6610846565b5060010190565b600084516108bf818460208901610789565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516108fb816001850160208a01610789565b60019201918201528351610916816002840160208801610789565b0160020195945050505050565b600181815b8085111561097c57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561096257610962610846565b8085161561096f57918102915b93841c9390800290610928565b509250929050565b600082610993575060016106bc565b816109a0575060006106bc565b81600181146109b657600281146109c0576109dc565b60019150506106bc565b60ff8411156109d1576109d1610846565b50506001821b6106bc565b5060208310610133831016604e8410600b84101617156109ff575081810a6106bc565b610a098383610923565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610a3b57610a3b610846565b029392505050565b6000610a4f8383610984565b9392505050565b818103818111156106bc576106bc610846565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c6343000810000a",
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
