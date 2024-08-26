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

// SequencerMetaData contains all meta data concerning the Sequencer contract.
var SequencerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"SequencerSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSetBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_sequencerSet\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCurrentSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerSetVerifyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSequencerSet\",\"type\":\"address[]\"}],\"name\":\"updateSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b507353000000000000000000000000000000000000156080526080516113f16100475f395f818161029c01526109d801526113f15ff3fe608060405234801561000f575f80fd5b506004361061019a575f3560e01c806377d7dffb116100e85780639b8201a411610093578063b1bdeab31161006e578063b1bdeab314610330578063dc55509014610338578063eae5b1e314610341578063f2fde38b14610349575f80fd5b80639b8201a414610302578063a2e53a9414610315578063a384c12e14610328575f80fd5b806389609d74116100c357806389609d74146102be5780638da5cb5b146102d1578063946d9204146102ef575f80fd5b806377d7dffb1461027c5780637d99e8ac14610284578063807de44314610297575f80fd5b806338871fac116101485780636d46e987116101235780636d46e987146102475780636d7f64db1461026a578063715018a614610272575f80fd5b806338871fac14610221578063480265c91461022957806365fd0f8d1461023e575f80fd5b806328d1357a1161017857806328d1357a146101d757806329025fcb146101e05780632aae60bd146101e9575f80fd5b80630d78725b1461019e5780630e06ede8146101ba57806317f24c2d146101c2575b5f80fd5b6101a760655481565b6040519081526020015b60405180910390f35b606b546101a7565b6101ca61035c565b6040516101b19190611073565b6101a7606a5481565b6101a760685481565b6101fc6101f73660046110cc565b6104ad565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6067546101a7565b6102316104e2565b6040516101b191906110e3565b6101a760665481565b61025a610255366004611175565b61051c565b60405190151581526020016101b1565b6101ca610592565b61027a6105fd565b005b6101ca610610565b6101fc6102923660046110cc565b61067b565b6101fc7f000000000000000000000000000000000000000000000000000000000000000081565b6101fc6102cc3660046110cc565b61068a565b60335473ffffffffffffffffffffffffffffffffffffffff166101fc565b61027a6102fd3660046111dd565b610699565b61027a61031036600461122c565b6109d5565b61025a610323366004611175565b610b8f565b6101a7610ce9565b6069546101a7565b6101a7606c5481565b6101ca610d11565b61027a610357366004611175565b610d7c565b6060606a5443106103d257606b8054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575b5050505050905090565b60685443106104445760698054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575050505050905090565b60678054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575050505050905090565b606781815481106104bc575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b606060665460676068546069606a54606b604051602001610508969594939291906112be565b604051602081830303815290604052905090565b5f61058c606b80548060200260200160405190810160405280929190818152602001828054801561058157602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610556575b505050505083610e33565b92915050565b606060678054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575050505050905090565b610605610ea6565b61060e5f610f27565b565b6060606b8054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575050505050905090565b606b81815481106104bc575f80fd5b606981815481106104bc575f80fd5b5f54610100900460ff16158080156106b757505f54600160ff909116105b806106d05750303b1580156106d057505f5460ff166001145b610761576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107bd575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff841661083a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964206f776e6572206164647265737300000000000000000000006044820152606401610758565b816108a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c69642073657175656e6365722073657400000000000000000000006044820152606401610758565b6108aa84610f27565b6108b660678484610f9d565b506108c360698484610f9d565b506108d0606b8484610f9d565b5042606c55606654606854606a546040516108fa9392606792909160699190606b906020016112be565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905280516020909101206065557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc629061096590859085905f9061135c565b60405180910390a180156109cf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610a74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79204c325374616b696e6720636f6e74726163740000000000000000006044820152606401610758565b606a54610a8243600261137f565b1115610ad65760688054606655606a549055610a9f43600261137f565b606a5560698054610ab291606791611023565b50606b8054610ac391606991611023565b50610ad0606b8383610f9d565b50610ae4565b610ae2606b8383610f9d565b505b42606c55606654606854606a54604051610b0d9392606792909160699190606b906020016112be565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206065557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc628282610b7443600261137f565b604051610b839392919061135c565b60405180910390a15050565b5f606a544310610c065761058c606b80548060200260200160405190810160405280929190818152602001828054801561058157602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161055657505050505083610e33565b6068544310610c7c5761058c606980548060200260200160405190810160405280929190818152602001828054801561058157602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161055657505050505083610e33565b61058c606780548060200260200160405190810160405280929190818152602001828054801561058157602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161055657505050505083610e33565b5f606a544310610cfa5750606b5490565b6068544310610d0a575060695490565b5060675490565b606060698054806020026020016040519081016040528092919081815260200182805480156103c857602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161039d575050505050905090565b610d84610ea6565b73ffffffffffffffffffffffffffffffffffffffff8116610e27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610758565b610e3081610f27565b50565b5f805b8351811015610e9d57838181518110610e5157610e516113b7565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610e9557600191505061058c565b600101610e36565b505f9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461060e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610758565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b828054828255905f5260205f20908101928215611013579160200282015b828111156110135781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190610fbb565b5061101f92915061105f565b5090565b828054828255905f5260205f20908101928215611013575f5260205f209182015b82811115611013578254825591600101919060010190611044565b5b8082111561101f575f8155600101611060565b602080825282518282018190525f9190848201906040850190845b818110156110c057835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161108e565b50909695505050505050565b5f602082840312156110dc575f80fd5b5035919050565b5f602080835283518060208501525f5b8181101561110f578581018301518582016040015282016110f3565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611170575f80fd5b919050565b5f60208284031215611185575f80fd5b61118e8261114d565b9392505050565b5f8083601f8401126111a5575f80fd5b50813567ffffffffffffffff8111156111bc575f80fd5b6020830191508360208260051b85010111156111d6575f80fd5b9250929050565b5f805f604084860312156111ef575f80fd5b6111f88461114d565b9250602084013567ffffffffffffffff811115611213575f80fd5b61121f86828701611195565b9497909650939450505050565b5f806020838503121561123d575f80fd5b823567ffffffffffffffff811115611253575f80fd5b61125f85828601611195565b90969095509350505050565b5f815480845260208085019450835f5260205f205f5b838110156112b357815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101611281565b509495945050505050565b86815260c060208201525f6112d660c083018861126b565b86604084015282810360608401526112ee818761126b565b905084608084015282810360a0840152611308818561126b565b9998505050505050505050565b8183525f60208085019450825f5b858110156112b35773ffffffffffffffffffffffffffffffffffffffff6113498361114d565b1687529582019590820190600101611323565b604081525f61136f604083018587611315565b9050826020830152949350505050565b8082018082111561058c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a",
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
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_Sequencer *SequencerCaller) UpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "updateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_Sequencer *SequencerSession) UpdateTime() (*big.Int, error) {
	return _Sequencer.Contract.UpdateTime(&_Sequencer.CallOpts)
}

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_Sequencer *SequencerCallerSession) UpdateTime() (*big.Int, error) {
	return _Sequencer.Contract.UpdateTime(&_Sequencer.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _sequencerSet) returns()
func (_Sequencer *SequencerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "initialize", _owner, _sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _sequencerSet) returns()
func (_Sequencer *SequencerSession) Initialize(_owner common.Address, _sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _owner, _sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _sequencerSet) returns()
func (_Sequencer *SequencerTransactorSession) Initialize(_owner common.Address, _sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, _owner, _sequencerSet)
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
