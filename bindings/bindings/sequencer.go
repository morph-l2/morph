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

// SequencerMetaData contains all meta data concerning the Sequencer contract.
var SequencerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"SequencerSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSetBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_sequencerSet\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCurrentSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerSetVerifyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSequencerSet\",\"type\":\"address[]\"}],\"name\":\"updateSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b507353000000000000000000000000000000000000126080526080516113d26100475f395f8181610291015261068801526113d25ff3fe608060405234801561000f575f80fd5b506004361061018f575f3560e01c806377d7dffb116100dd578063a224cee711610088578063b1bdeab311610063578063b1bdeab314610325578063eae5b1e31461032d578063f2fde38b14610335575f80fd5b8063a224cee7146102f7578063a2e53a941461030a578063a384c12e1461031d575f80fd5b806389609d74116100b857806389609d74146102b35780638da5cb5b146102c65780639b8201a4146102e4575f80fd5b806377d7dffb146102715780637d99e8ac14610279578063807de4431461028c575f80fd5b806338871fac1161013d5780636d46e987116101185780636d46e9871461023c5780636d7f64db1461025f578063715018a614610267575f80fd5b806338871fac14610216578063480265c91461021e57806365fd0f8d14610233575f80fd5b806328d1357a1161016d57806328d1357a146101cc57806329025fcb146101d55780632aae60bd146101de575f80fd5b80630d78725b146101935780630e06ede8146101af57806317f24c2d146101b7575b5f80fd5b61019c60655481565b6040519081526020015b60405180910390f35b606b5461019c565b6101bf610348565b6040516101a691906110b2565b61019c606a5481565b61019c60685481565b6101f16101ec36600461110b565b610499565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a6565b60675461019c565b6102266104ce565b6040516101a69190611122565b61019c60665481565b61024f61024a3660046111b4565b610508565b60405190151581526020016101a6565b6101bf61057e565b61026f6105e9565b005b6101bf6105fc565b6101f161028736600461110b565b610667565b6101f17f000000000000000000000000000000000000000000000000000000000000000081565b6101f16102c136600461110b565b610676565b60335473ffffffffffffffffffffffffffffffffffffffff166101f1565b61026f6102f23660046111d4565b610685565b61026f6103053660046111d4565b61083f565b61024f6103183660046111b4565b610a96565b61019c610bf0565b60695461019c565b6101bf610c18565b61026f6103433660046111b4565b610c83565b6060606a5443106103be57606b8054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575b5050505050905090565b60685443106104305760698054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b60678054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b606781815481106104a8575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b606060665460676068546069606a54606b6040516020016104f49695949392919061128d565b604051602081830303815290604052905090565b5f610578606b80548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610542575b505050505083610d3a565b92915050565b606060678054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b6105f1610dad565b6105fa5f610e2e565b565b6060606b8054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b606b81815481106104a8575f80fd5b606981815481106104a8575f80fd5b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79204c325374616b696e6720636f6e747261637400000000000000000060448201526064015b60405180910390fd5b606a546107374360026112cd565b111561078b5760688054606655606a5490556107544360026112cd565b606a556069805461076791606791610fdc565b50606b805461077891606991610fdc565b50610785606b8383611028565b50610799565b610797606b8383611028565b505b60665460676068546069606a54606b6040516020016107bd9695949392919061128d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206065557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc6282826108244360026112cd565b6040516108339392919061134c565b60405180910390a15050565b5f54610100900460ff161580801561085d57505f54600160ff909116105b806108765750303b15801561087657505f5460ff166001145b610902576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610720565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561095e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b816109c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c69642073657175656e6365722073657400000000000000000000006044820152606401610720565b6109cd610ea4565b6109d960678484611028565b506109e660698484611028565b506109f3606b8484611028565b507f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc6283835f604051610a279392919061134c565b60405180910390a18015610a91575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b5f606a544310610b0d57610578606b80548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b6068544310610b8357610578606980548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b610578606780548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b5f606a544310610c015750606b5490565b6068544310610c11575060695490565b5060675490565b606060698054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b610c8b610dad565b73ffffffffffffffffffffffffffffffffffffffff8116610d2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610720565b610d3781610e2e565b50565b5f805b8351811015610da457838181518110610d5857610d5861136f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d9c576001915050610578565b600101610d3d565b505f9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610720565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610f3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610720565b6105fa5f54610100900460ff16610fd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610720565b6105fa33610e2e565b828054828255905f5260205f20908101928215611018575f5260205f209182015b82811115611018578254825591600101919060010190610ffd565b5061102492915061109e565b5090565b828054828255905f5260205f20908101928215611018579160200282015b828111156110185781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190611046565b5b80821115611024575f815560010161109f565b602080825282518282018190525f9190848201906040850190845b818110156110ff57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016110cd565b50909695505050505050565b5f6020828403121561111b575f80fd5b5035919050565b5f602080835283518060208501525f5b8181101561114e57858101830151858201604001528201611132565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146111af575f80fd5b919050565b5f602082840312156111c4575f80fd5b6111cd8261118c565b9392505050565b5f80602083850312156111e5575f80fd5b823567ffffffffffffffff808211156111fc575f80fd5b818501915085601f83011261120f575f80fd5b81358181111561121d575f80fd5b8660208260051b8501011115611231575f80fd5b60209290920196919550909350505050565b5f8154825f526020805f205f5b8381101561128257815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101611250565b509495945050505050565b8681525f61129e6020830188611243565b8681526112ae6020820187611243565b90508481526112c06020820185611243565b9998505050505050505050565b80820180821115610578577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8183525f60208085019450825f5b858110156112825773ffffffffffffffffffffffffffffffffffffffff6113398361118c565b1687529582019590820190600101611313565b604081525f61135f604083018587611305565b9050826020830152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220189b0c1c9f8476dd3d8c84d14cf667a717234c1ed3571e390aadb4cd9a970d6d64736f6c63430008180033",
}

// SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerMetaData.ABI instead.
var SequencerABI = SequencerMetaData.ABI

// SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerMetaData.Bin instead.
var SequencerBin = SequencerMetaData.Bin

// DeploySequencer deploys a new Ethereum contract, binding an instance of Sequencer to it.
func DeploySequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sequencer, error) {
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// Sequencer is an auto generated Go binding around an Ethereum contract.
type Sequencer struct {
	SequencerCaller     // Read-only binding to the contract
	SequencerTransactor // Write-only binding to the contract
	SequencerFilterer   // Log filterer for contract events
}

// SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSession struct {
	Contract     *Sequencer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerCallerSession struct {
	Contract *SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerTransactorSession struct {
	Contract     *SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRaw struct {
	Contract *Sequencer // Generic contract binding to access the raw methods on
}

// SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerCallerRaw struct {
	Contract *SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerTransactorRaw struct {
	Contract *SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencer creates a new instance of Sequencer, bound to a specific deployed contract.
func NewSequencer(address common.Address, backend bind.ContractBackend) (*Sequencer, error) {
	contract, err := bindSequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// NewSequencerCaller creates a new read-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerCaller(address common.Address, caller bind.ContractCaller) (*SequencerCaller, error) {
	contract, err := bindSequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerCaller{contract: contract}, nil
}

// NewSequencerTransactor creates a new write-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerTransactor, error) {
	contract, err := bindSequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerTransactor{contract: contract}, nil
}

// NewSequencerFilterer creates a new log filterer instance of Sequencer, bound to a specific deployed contract.
func NewSequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerFilterer, error) {
	contract, err := bindSequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerFilterer{contract: contract}, nil
}

// bindSequencer binds a generic wrapper to an already deployed contract.
func bindSequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transact(opts, method, params...)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Sequencer *SequencerCaller) L2STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "L2_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Sequencer *SequencerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Sequencer.Contract.L2STAKINGCONTRACT(&_Sequencer.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Sequencer *SequencerCallerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Sequencer.Contract.L2STAKINGCONTRACT(&_Sequencer.CallOpts)
}

// BlockHeight0 is a free data retrieval call binding the contract method 0x65fd0f8d.
//
// Solidity: function blockHeight0() view returns(uint256)
func (_Sequencer *SequencerCaller) BlockHeight0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "blockHeight0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockHeight0 is a free data retrieval call binding the contract method 0x65fd0f8d.
//
// Solidity: function blockHeight0() view returns(uint256)
func (_Sequencer *SequencerSession) BlockHeight0() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight0(&_Sequencer.CallOpts)
}

// BlockHeight0 is a free data retrieval call binding the contract method 0x65fd0f8d.
//
// Solidity: function blockHeight0() view returns(uint256)
func (_Sequencer *SequencerCallerSession) BlockHeight0() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight0(&_Sequencer.CallOpts)
}

// BlockHeight1 is a free data retrieval call binding the contract method 0x29025fcb.
//
// Solidity: function blockHeight1() view returns(uint256)
func (_Sequencer *SequencerCaller) BlockHeight1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "blockHeight1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockHeight1 is a free data retrieval call binding the contract method 0x29025fcb.
//
// Solidity: function blockHeight1() view returns(uint256)
func (_Sequencer *SequencerSession) BlockHeight1() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight1(&_Sequencer.CallOpts)
}

// BlockHeight1 is a free data retrieval call binding the contract method 0x29025fcb.
//
// Solidity: function blockHeight1() view returns(uint256)
func (_Sequencer *SequencerCallerSession) BlockHeight1() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight1(&_Sequencer.CallOpts)
}

// BlockHeight2 is a free data retrieval call binding the contract method 0x28d1357a.
//
// Solidity: function blockHeight2() view returns(uint256)
func (_Sequencer *SequencerCaller) BlockHeight2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "blockHeight2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockHeight2 is a free data retrieval call binding the contract method 0x28d1357a.
//
// Solidity: function blockHeight2() view returns(uint256)
func (_Sequencer *SequencerSession) BlockHeight2() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight2(&_Sequencer.CallOpts)
}

// BlockHeight2 is a free data retrieval call binding the contract method 0x28d1357a.
//
// Solidity: function blockHeight2() view returns(uint256)
func (_Sequencer *SequencerCallerSession) BlockHeight2() (*big.Int, error) {
	return _Sequencer.Contract.BlockHeight2(&_Sequencer.CallOpts)
}

// GetCurrentSequencerSet is a free data retrieval call binding the contract method 0x17f24c2d.
//
// Solidity: function getCurrentSequencerSet() view returns(address[])
func (_Sequencer *SequencerCaller) GetCurrentSequencerSet(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getCurrentSequencerSet")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentSequencerSet is a free data retrieval call binding the contract method 0x17f24c2d.
//
// Solidity: function getCurrentSequencerSet() view returns(address[])
func (_Sequencer *SequencerSession) GetCurrentSequencerSet() ([]common.Address, error) {
	return _Sequencer.Contract.GetCurrentSequencerSet(&_Sequencer.CallOpts)
}

// GetCurrentSequencerSet is a free data retrieval call binding the contract method 0x17f24c2d.
//
// Solidity: function getCurrentSequencerSet() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetCurrentSequencerSet() ([]common.Address, error) {
	return _Sequencer.Contract.GetCurrentSequencerSet(&_Sequencer.CallOpts)
}

// GetCurrentSequencerSetSize is a free data retrieval call binding the contract method 0xa384c12e.
//
// Solidity: function getCurrentSequencerSetSize() view returns(uint256)
func (_Sequencer *SequencerCaller) GetCurrentSequencerSetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getCurrentSequencerSetSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSequencerSetSize is a free data retrieval call binding the contract method 0xa384c12e.
//
// Solidity: function getCurrentSequencerSetSize() view returns(uint256)
func (_Sequencer *SequencerSession) GetCurrentSequencerSetSize() (*big.Int, error) {
	return _Sequencer.Contract.GetCurrentSequencerSetSize(&_Sequencer.CallOpts)
}

// GetCurrentSequencerSetSize is a free data retrieval call binding the contract method 0xa384c12e.
//
// Solidity: function getCurrentSequencerSetSize() view returns(uint256)
func (_Sequencer *SequencerCallerSession) GetCurrentSequencerSetSize() (*big.Int, error) {
	return _Sequencer.Contract.GetCurrentSequencerSetSize(&_Sequencer.CallOpts)
}

// GetSequencerSet0 is a free data retrieval call binding the contract method 0x6d7f64db.
//
// Solidity: function getSequencerSet0() view returns(address[])
func (_Sequencer *SequencerCaller) GetSequencerSet0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencerSet0 is a free data retrieval call binding the contract method 0x6d7f64db.
//
// Solidity: function getSequencerSet0() view returns(address[])
func (_Sequencer *SequencerSession) GetSequencerSet0() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet0(&_Sequencer.CallOpts)
}

// GetSequencerSet0 is a free data retrieval call binding the contract method 0x6d7f64db.
//
// Solidity: function getSequencerSet0() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetSequencerSet0() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet0(&_Sequencer.CallOpts)
}

// GetSequencerSet0Size is a free data retrieval call binding the contract method 0x38871fac.
//
// Solidity: function getSequencerSet0Size() view returns(uint256)
func (_Sequencer *SequencerCaller) GetSequencerSet0Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet0Size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequencerSet0Size is a free data retrieval call binding the contract method 0x38871fac.
//
// Solidity: function getSequencerSet0Size() view returns(uint256)
func (_Sequencer *SequencerSession) GetSequencerSet0Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet0Size(&_Sequencer.CallOpts)
}

// GetSequencerSet0Size is a free data retrieval call binding the contract method 0x38871fac.
//
// Solidity: function getSequencerSet0Size() view returns(uint256)
func (_Sequencer *SequencerCallerSession) GetSequencerSet0Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet0Size(&_Sequencer.CallOpts)
}

// GetSequencerSet1 is a free data retrieval call binding the contract method 0xeae5b1e3.
//
// Solidity: function getSequencerSet1() view returns(address[])
func (_Sequencer *SequencerCaller) GetSequencerSet1(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet1")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencerSet1 is a free data retrieval call binding the contract method 0xeae5b1e3.
//
// Solidity: function getSequencerSet1() view returns(address[])
func (_Sequencer *SequencerSession) GetSequencerSet1() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet1(&_Sequencer.CallOpts)
}

// GetSequencerSet1 is a free data retrieval call binding the contract method 0xeae5b1e3.
//
// Solidity: function getSequencerSet1() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetSequencerSet1() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet1(&_Sequencer.CallOpts)
}

// GetSequencerSet1Size is a free data retrieval call binding the contract method 0xb1bdeab3.
//
// Solidity: function getSequencerSet1Size() view returns(uint256)
func (_Sequencer *SequencerCaller) GetSequencerSet1Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet1Size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequencerSet1Size is a free data retrieval call binding the contract method 0xb1bdeab3.
//
// Solidity: function getSequencerSet1Size() view returns(uint256)
func (_Sequencer *SequencerSession) GetSequencerSet1Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet1Size(&_Sequencer.CallOpts)
}

// GetSequencerSet1Size is a free data retrieval call binding the contract method 0xb1bdeab3.
//
// Solidity: function getSequencerSet1Size() view returns(uint256)
func (_Sequencer *SequencerCallerSession) GetSequencerSet1Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet1Size(&_Sequencer.CallOpts)
}

// GetSequencerSet2 is a free data retrieval call binding the contract method 0x77d7dffb.
//
// Solidity: function getSequencerSet2() view returns(address[])
func (_Sequencer *SequencerCaller) GetSequencerSet2(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet2")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencerSet2 is a free data retrieval call binding the contract method 0x77d7dffb.
//
// Solidity: function getSequencerSet2() view returns(address[])
func (_Sequencer *SequencerSession) GetSequencerSet2() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet2(&_Sequencer.CallOpts)
}

// GetSequencerSet2 is a free data retrieval call binding the contract method 0x77d7dffb.
//
// Solidity: function getSequencerSet2() view returns(address[])
func (_Sequencer *SequencerCallerSession) GetSequencerSet2() ([]common.Address, error) {
	return _Sequencer.Contract.GetSequencerSet2(&_Sequencer.CallOpts)
}

// GetSequencerSet2Size is a free data retrieval call binding the contract method 0x0e06ede8.
//
// Solidity: function getSequencerSet2Size() view returns(uint256)
func (_Sequencer *SequencerCaller) GetSequencerSet2Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSet2Size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequencerSet2Size is a free data retrieval call binding the contract method 0x0e06ede8.
//
// Solidity: function getSequencerSet2Size() view returns(uint256)
func (_Sequencer *SequencerSession) GetSequencerSet2Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet2Size(&_Sequencer.CallOpts)
}

// GetSequencerSet2Size is a free data retrieval call binding the contract method 0x0e06ede8.
//
// Solidity: function getSequencerSet2Size() view returns(uint256)
func (_Sequencer *SequencerCallerSession) GetSequencerSet2Size() (*big.Int, error) {
	return _Sequencer.Contract.GetSequencerSet2Size(&_Sequencer.CallOpts)
}

// GetSequencerSetBytes is a free data retrieval call binding the contract method 0x480265c9.
//
// Solidity: function getSequencerSetBytes() view returns(bytes)
func (_Sequencer *SequencerCaller) GetSequencerSetBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getSequencerSetBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSequencerSetBytes is a free data retrieval call binding the contract method 0x480265c9.
//
// Solidity: function getSequencerSetBytes() view returns(bytes)
func (_Sequencer *SequencerSession) GetSequencerSetBytes() ([]byte, error) {
	return _Sequencer.Contract.GetSequencerSetBytes(&_Sequencer.CallOpts)
}

// GetSequencerSetBytes is a free data retrieval call binding the contract method 0x480265c9.
//
// Solidity: function getSequencerSetBytes() view returns(bytes)
func (_Sequencer *SequencerCallerSession) GetSequencerSetBytes() ([]byte, error) {
	return _Sequencer.Contract.GetSequencerSetBytes(&_Sequencer.CallOpts)
}

// IsCurrentSequencer is a free data retrieval call binding the contract method 0xa2e53a94.
//
// Solidity: function isCurrentSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerCaller) IsCurrentSequencer(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "isCurrentSequencer", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentSequencer is a free data retrieval call binding the contract method 0xa2e53a94.
//
// Solidity: function isCurrentSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerSession) IsCurrentSequencer(addr common.Address) (bool, error) {
	return _Sequencer.Contract.IsCurrentSequencer(&_Sequencer.CallOpts, addr)
}

// IsCurrentSequencer is a free data retrieval call binding the contract method 0xa2e53a94.
//
// Solidity: function isCurrentSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerCallerSession) IsCurrentSequencer(addr common.Address) (bool, error) {
	return _Sequencer.Contract.IsCurrentSequencer(&_Sequencer.CallOpts, addr)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerCaller) IsSequencer(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "isSequencer", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerSession) IsSequencer(addr common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, addr)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_Sequencer *SequencerCallerSession) IsSequencer(addr common.Address) (bool, error) {
	return _Sequencer.Contract.IsSequencer(&_Sequencer.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sequencer *SequencerCallerSession) Owner() (common.Address, error) {
	return _Sequencer.Contract.Owner(&_Sequencer.CallOpts)
}

// SequencerSet0 is a free data retrieval call binding the contract method 0x2aae60bd.
//
// Solidity: function sequencerSet0(uint256 ) view returns(address)
func (_Sequencer *SequencerCaller) SequencerSet0(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerSet0", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerSet0 is a free data retrieval call binding the contract method 0x2aae60bd.
//
// Solidity: function sequencerSet0(uint256 ) view returns(address)
func (_Sequencer *SequencerSession) SequencerSet0(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet0(&_Sequencer.CallOpts, arg0)
}

// SequencerSet0 is a free data retrieval call binding the contract method 0x2aae60bd.
//
// Solidity: function sequencerSet0(uint256 ) view returns(address)
func (_Sequencer *SequencerCallerSession) SequencerSet0(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet0(&_Sequencer.CallOpts, arg0)
}

// SequencerSet1 is a free data retrieval call binding the contract method 0x89609d74.
//
// Solidity: function sequencerSet1(uint256 ) view returns(address)
func (_Sequencer *SequencerCaller) SequencerSet1(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerSet1", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerSet1 is a free data retrieval call binding the contract method 0x89609d74.
//
// Solidity: function sequencerSet1(uint256 ) view returns(address)
func (_Sequencer *SequencerSession) SequencerSet1(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet1(&_Sequencer.CallOpts, arg0)
}

// SequencerSet1 is a free data retrieval call binding the contract method 0x89609d74.
//
// Solidity: function sequencerSet1(uint256 ) view returns(address)
func (_Sequencer *SequencerCallerSession) SequencerSet1(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet1(&_Sequencer.CallOpts, arg0)
}

// SequencerSet2 is a free data retrieval call binding the contract method 0x7d99e8ac.
//
// Solidity: function sequencerSet2(uint256 ) view returns(address)
func (_Sequencer *SequencerCaller) SequencerSet2(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerSet2", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerSet2 is a free data retrieval call binding the contract method 0x7d99e8ac.
//
// Solidity: function sequencerSet2(uint256 ) view returns(address)
func (_Sequencer *SequencerSession) SequencerSet2(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet2(&_Sequencer.CallOpts, arg0)
}

// SequencerSet2 is a free data retrieval call binding the contract method 0x7d99e8ac.
//
// Solidity: function sequencerSet2(uint256 ) view returns(address)
func (_Sequencer *SequencerCallerSession) SequencerSet2(arg0 *big.Int) (common.Address, error) {
	return _Sequencer.Contract.SequencerSet2(&_Sequencer.CallOpts, arg0)
}

// SequencerSetVerifyHash is a free data retrieval call binding the contract method 0x0d78725b.
//
// Solidity: function sequencerSetVerifyHash() view returns(bytes32)
func (_Sequencer *SequencerCaller) SequencerSetVerifyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "sequencerSetVerifyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerSetVerifyHash is a free data retrieval call binding the contract method 0x0d78725b.
//
// Solidity: function sequencerSetVerifyHash() view returns(bytes32)
func (_Sequencer *SequencerSession) SequencerSetVerifyHash() ([32]byte, error) {
	return _Sequencer.Contract.SequencerSetVerifyHash(&_Sequencer.CallOpts)
}

// SequencerSetVerifyHash is a free data retrieval call binding the contract method 0x0d78725b.
//
// Solidity: function sequencerSetVerifyHash() view returns(bytes32)
func (_Sequencer *SequencerCallerSession) SequencerSetVerifyHash() ([32]byte, error) {
	return _Sequencer.Contract.SequencerSetVerifyHash(&_Sequencer.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _sequencerSet) returns()
func (_Sequencer *SequencerTransactor) Initialize(opts *bind.TransactOpts, _sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "initialize", _sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _sequencerSet) returns()
func (_Sequencer *SequencerSession) Initialize(_sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _sequencerSet) returns()
func (_Sequencer *SequencerTransactorSession) Initialize(_sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _sequencerSet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sequencer *SequencerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sequencer.Contract.RenounceOwnership(&_Sequencer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sequencer *SequencerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.TransferOwnership(&_Sequencer.TransactOpts, newOwner)
}

// UpdateSequencerSet is a paid mutator transaction binding the contract method 0x9b8201a4.
//
// Solidity: function updateSequencerSet(address[] newSequencerSet) returns()
func (_Sequencer *SequencerTransactor) UpdateSequencerSet(opts *bind.TransactOpts, newSequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "updateSequencerSet", newSequencerSet)
}

// UpdateSequencerSet is a paid mutator transaction binding the contract method 0x9b8201a4.
//
// Solidity: function updateSequencerSet(address[] newSequencerSet) returns()
func (_Sequencer *SequencerSession) UpdateSequencerSet(newSequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerSet(&_Sequencer.TransactOpts, newSequencerSet)
}

// UpdateSequencerSet is a paid mutator transaction binding the contract method 0x9b8201a4.
//
// Solidity: function updateSequencerSet(address[] newSequencerSet) returns()
func (_Sequencer *SequencerTransactorSession) UpdateSequencerSet(newSequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.UpdateSequencerSet(&_Sequencer.TransactOpts, newSequencerSet)
}

// SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Sequencer contract.
type SequencerInitializedIterator struct {
	Event *SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInitialized)
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
		it.Event = new(SequencerInitialized)
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
func (it *SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInitialized represents a Initialized event raised by the Sequencer contract.
type SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerInitializedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerInitializedIterator{contract: _Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Sequencer *SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInitialized)
				if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Sequencer *SequencerFilterer) ParseInitialized(log types.Log) (*SequencerInitialized, error) {
	event := new(SequencerInitialized)
	if err := _Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sequencer contract.
type SequencerOwnershipTransferredIterator struct {
	Event *SequencerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SequencerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerOwnershipTransferred)
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
		it.Event = new(SequencerOwnershipTransferred)
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
func (it *SequencerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerOwnershipTransferred represents a OwnershipTransferred event raised by the Sequencer contract.
type SequencerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerOwnershipTransferredIterator{contract: _Sequencer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sequencer *SequencerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SequencerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerOwnershipTransferred)
				if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Sequencer *SequencerFilterer) ParseOwnershipTransferred(log types.Log) (*SequencerOwnershipTransferred, error) {
	event := new(SequencerOwnershipTransferred)
	if err := _Sequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSequencerSetUpdatedIterator is returned from FilterSequencerSetUpdated and is used to iterate over the raw logs and unpacked data for SequencerSetUpdated events raised by the Sequencer contract.
type SequencerSequencerSetUpdatedIterator struct {
	Event *SequencerSequencerSetUpdated // Event containing the contract specifics and raw log

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
func (it *SequencerSequencerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSequencerSetUpdated)
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
		it.Event = new(SequencerSequencerSetUpdated)
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
func (it *SequencerSequencerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSequencerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSequencerSetUpdated represents a SequencerSetUpdated event raised by the Sequencer contract.
type SequencerSequencerSetUpdated struct {
	SequencerSet []common.Address
	BlockHeight  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSequencerSetUpdated is a free log retrieval operation binding the contract event 0x7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc62.
//
// Solidity: event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight)
func (_Sequencer *SequencerFilterer) FilterSequencerSetUpdated(opts *bind.FilterOpts) (*SequencerSequencerSetUpdatedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "SequencerSetUpdated")
	if err != nil {
		return nil, err
	}
	return &SequencerSequencerSetUpdatedIterator{contract: _Sequencer.contract, event: "SequencerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerSetUpdated is a free log subscription operation binding the contract event 0x7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc62.
//
// Solidity: event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight)
func (_Sequencer *SequencerFilterer) WatchSequencerSetUpdated(opts *bind.WatchOpts, sink chan<- *SequencerSequencerSetUpdated) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "SequencerSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSequencerSetUpdated)
				if err := _Sequencer.contract.UnpackLog(event, "SequencerSetUpdated", log); err != nil {
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

// ParseSequencerSetUpdated is a log parse operation binding the contract event 0x7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc62.
//
// Solidity: event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight)
func (_Sequencer *SequencerFilterer) ParseSequencerSetUpdated(log types.Log) (*SequencerSequencerSetUpdated, error) {
	event := new(SequencerSequencerSetUpdated)
	if err := _Sequencer.contract.UnpackLog(event, "SequencerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
