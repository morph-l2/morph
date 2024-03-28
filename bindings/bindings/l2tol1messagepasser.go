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
	_ = abi.ConvertType
)

// L2ToL1MessagePasserMetaData contains all meta data concerning the L2ToL1MessagePasser contract.
var L2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"AppendMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"appendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafNodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610018610020565b602155610106565b6020545f90819081805b60208110156100e9578083901c600116600103610086575f8160208110610053576100536100f2565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506100b3565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160408051601f198184030181529190528051602090910120915060010161002a565b50919392505050565b634e487b7160e01b5f52603260045260245ffd5b61079a806101135f395ff3fe608060405234801561000f575f80fd5b5060043610610064575f3560e01c806389c09d381161004d57806389c09d38146100b1578063b58343bb146100b9578063d4b9f4fa146100c2575f80fd5b8063340735f714610068578063600a2e7714610090575b5f80fd5b61007b6100763660046104b2565b6100cb565b60405190151581526020015b60405180910390f35b6100a361009e366004610546565b610194565b604051908152602001610087565b6100a3610282565b6100a360205481565b6100a360215481565b5f84815b6020811015610188578085901c600116600103610135578581602081106100f8576100f861055d565b602002015182604051602001610118929190918252602082015260400190565b604051602081830303815290604052805190602001209150610180565b818682602081106101485761014861055d565b6020020151604051602001610167929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b6001016100cf565b50909114949350505050565b5f3373530000000000000000000000000000000000000714610216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6f6e6c79206d657373656e676572000000000000000000000000000000000000604482015260640160405180910390fd5b61021f82610372565b610227610282565b6021556020547f509758f52fb5e05d2e0d4379024275cbab7c27923c22777fcdb7e12a4d9499639061025b906001906105b7565b602154604080519283526020830186905282015260600160405180910390a1505060215490565b6020545f90819081805b6020811015610369578083901c6001166001036102e8575f81602081106102b5576102b561055d565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610315565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120915060010161028c565b50919392505050565b806001610381602060026106ee565b61038b91906105b7565b602054106103c5576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60205f81546103d490610700565b918290555090505f5b6020811015610477578082901c60011660010361040f57825f82602081106104075761040761055d565b015550505050565b5f81602081106104215761042161055d565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012092506001016103dd565b50610480610737565b505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f8061046085870312156104c6575f80fd5b84359350602086603f8701126104da575f80fd5b604051610400810181811067ffffffffffffffff821117156104fe576104fe610485565b60405280610420880189811115610513575f80fd5b602089015b8181101561052f5780358352918401918401610518565b509699919850509435956104400135949350505050565b5f60208284031215610556575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156105ca576105ca61058a565b92915050565b600181815b8085111561062957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561060f5761060f61058a565b8085161561061c57918102915b93841c93908002906105d5565b509250929050565b5f8261063f575060016105ca565b8161064b57505f6105ca565b8160018114610661576002811461066b57610687565b60019150506105ca565b60ff84111561067c5761067c61058a565b50506001821b6105ca565b5060208310610133831016604e8410600b84101617156106aa575081810a6105ca565b6106b483836105d0565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156106e6576106e661058a565b029392505050565b5f6106f98383610631565b9392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107305761073061058a565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea2646970667358221220379f9a67cc35fe8f0c311c533610e293a1f475f33e130716ce5b303cda19a4d364736f6c63430008180033",
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
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	RootHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAppendMessage is a free log retrieval operation binding the contract event 0x509758f52fb5e05d2e0d4379024275cbab7c27923c22777fcdb7e12a4d949963.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterAppendMessage(opts *bind.FilterOpts) (*L2ToL1MessagePasserAppendMessageIterator, error) {

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "AppendMessage")
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserAppendMessageIterator{contract: _L2ToL1MessagePasser.contract, event: "AppendMessage", logs: logs, sub: sub}, nil
}

// WatchAppendMessage is a free log subscription operation binding the contract event 0x509758f52fb5e05d2e0d4379024275cbab7c27923c22777fcdb7e12a4d949963.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash, bytes32 rootHash)
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

// ParseAppendMessage is a log parse operation binding the contract event 0x509758f52fb5e05d2e0d4379024275cbab7c27923c22777fcdb7e12a4d949963.
//
// Solidity: event AppendMessage(uint256 index, bytes32 messageHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseAppendMessage(log types.Log) (*L2ToL1MessagePasserAppendMessage, error) {
	event := new(L2ToL1MessagePasserAppendMessage)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "AppendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
