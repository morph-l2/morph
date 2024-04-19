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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"SequencerSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_SET_VERIFY_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCurrentSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSequencerSet\",\"type\":\"address[]\"}],\"name\":\"updateSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50735300000000000000000000000000000000000012608052608051610fec6100475f395f818161024201526105d20152610fec5ff3fe608060405234801561000f575f80fd5b5060043610610163575f3560e01c806377d7dffb116100c7578063a224cee71161007d578063a384c12e11610063578063a384c12e146102b2578063b1bdeab3146102ba578063eae5b1e3146102c2575f80fd5b8063a224cee71461028c578063a2e53a941461029f575f80fd5b8063807de443116100ad578063807de4431461023d57806389609d74146102645780639b8201a414610277575f80fd5b806377d7dffb146102225780637d99e8ac1461022a575f80fd5b806338871fac1161011c57806365fd0f8d1161010257806365fd0f8d146101ee5780636d46e987146101f75780636d7f64db1461021a575f80fd5b806338871fac146101dd578063646f7da0146101e5575f80fd5b806328d1357a1161014c57806328d1357a1461019357806329025fcb1461019c5780632aae60bd146101a5575f80fd5b80630e06ede81461016757806317f24c2d1461017e575b5f80fd5b6007545b6040519081526020015b60405180910390f35b6101866102ca565b6040516101759190610d35565b61016b60065481565b61016b60045481565b6101b86101b3366004610d4e565b61041b565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610175565b60035461016b565b61016b60015481565b61016b60025481565b61020a610205366004610d8d565b610450565b6040519015158152602001610175565b6101866104c6565b610186610531565b6101b8610238366004610d4e565b61059c565b6101b87f000000000000000000000000000000000000000000000000000000000000000081565b6101b8610272366004610d4e565b6105ab565b61028a610285366004610dd3565b6105ba565b005b61028a61029a366004610dd3565b610782565b61020a6102ad366004610d8d565b6109ad565b61016b610b07565b60055461016b565b610186610b2f565b6060600654431061034057600780548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575b5050505050905090565b60045443106103b257600580548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575050505050905090565b600380548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575050505050905090565b6003818154811061042a575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b5f6104c060078054806020026020016040519081016040528092919081815260200182805480156104b557602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161048a575b505050505083610b9a565b92915050565b6060600380548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575050505050905090565b6060600780548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575050505050905090565b6007818154811061042a575f80fd5b6005818154811061042a575f80fd5b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461065e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79204c325374616b696e6720636f6e747261637400000000000000000060448201526064015b60405180910390fd5b60065461066c436002610eb1565b11156106ca5760048054600290815560065490915561068c904390610eb1565b6006556005805461069f91600391610c0d565b50600780546106b091600591610c0d565b5080516106c4906007906020840190610c59565b506106df565b80516106dd906007906020840190610c59565b505b60025460036004546005600654600760405160200161070396959493929190610f28565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206001557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc6281610769436002610eb1565b604051610777929190610f68565b60405180910390a150565b5f54610100900460ff16158080156107a057505f54600160ff909116105b806107b95750303b1580156107b957505f5460ff166001145b610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610655565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108a1575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f82511161090b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c69642073657175656e6365722073657400000000000000000000006044820152606401610655565b815161091e906003906020850190610c59565b508151610932906005906020850190610c59565b508151610946906007906020850190610c59565b5080156109a9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b5f6006544310610a24576104c060078054806020026020016040519081016040528092919081815260200182805480156104b557602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161048a57505050505083610b9a565b6004544310610a9a576104c060058054806020026020016040519081016040528092919081815260200182805480156104b557602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161048a57505050505083610b9a565b6104c060038054806020026020016040519081016040528092919081815260200182805480156104b557602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161048a57505050505083610b9a565b5f6006544310610b18575060075490565b6004544310610b28575060055490565b5060035490565b6060600580548060200260200160405190810160405280929190818152602001828054801561033657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161030b575050505050905090565b5f805b8351811015610c0457838181518110610bb857610bb8610f89565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bfc5760019150506104c0565b600101610b9d565b505f9392505050565b828054828255905f5260205f20908101928215610c49575f5260205f209182015b82811115610c49578254825591600101919060010190610c2e565b50610c55929150610cd1565b5090565b828054828255905f5260205f20908101928215610c49579160200282015b82811115610c4957825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610c77565b5b80821115610c55575f8155600101610cd2565b5f815180845260208085019450602084015f5b83811015610d2a57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610cf8565b509495945050505050565b602081525f610d476020830184610ce5565b9392505050565b5f60208284031215610d5e575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d88575f80fd5b919050565b5f60208284031215610d9d575f80fd5b610d4782610d65565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020808385031215610de4575f80fd5b823567ffffffffffffffff80821115610dfb575f80fd5b818501915085601f830112610e0e575f80fd5b813581811115610e2057610e20610da6565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715610e6357610e63610da6565b604052918252848201925083810185019188831115610e80575f80fd5b938501935b82851015610ea557610e9685610d65565b84529385019392850192610e85565b98975050505050505050565b808201808211156104c0577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8154825f526020805f205f5b83811015610d2a57815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101610ef6565b8681525f610f396020830188610ee9565b868152610f496020820187610ee9565b9050848152610f5b6020820185610ee9565b9998505050505050505050565b604081525f610f7a6040830185610ce5565b90508260208301529392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220ddefb8932ae8ddcd924fc230b4a18528ab526d64691badba9147ba834691bde664736f6c63430008180033",
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

// SEQUENCERSETVERIFYHASH is a free data retrieval call binding the contract method 0x646f7da0.
//
// Solidity: function SEQUENCER_SET_VERIFY_HASH() view returns(bytes32)
func (_Sequencer *SequencerCaller) SEQUENCERSETVERIFYHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "SEQUENCER_SET_VERIFY_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERSETVERIFYHASH is a free data retrieval call binding the contract method 0x646f7da0.
//
// Solidity: function SEQUENCER_SET_VERIFY_HASH() view returns(bytes32)
func (_Sequencer *SequencerSession) SEQUENCERSETVERIFYHASH() ([32]byte, error) {
	return _Sequencer.Contract.SEQUENCERSETVERIFYHASH(&_Sequencer.CallOpts)
}

// SEQUENCERSETVERIFYHASH is a free data retrieval call binding the contract method 0x646f7da0.
//
// Solidity: function SEQUENCER_SET_VERIFY_HASH() view returns(bytes32)
func (_Sequencer *SequencerCallerSession) SEQUENCERSETVERIFYHASH() ([32]byte, error) {
	return _Sequencer.Contract.SEQUENCERSETVERIFYHASH(&_Sequencer.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] sequencerSet) returns()
func (_Sequencer *SequencerTransactor) Initialize(opts *bind.TransactOpts, sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "initialize", sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] sequencerSet) returns()
func (_Sequencer *SequencerSession) Initialize(sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, sequencerSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] sequencerSet) returns()
func (_Sequencer *SequencerTransactorSession) Initialize(sequencerSet []common.Address) (*types.Transaction, error) {
	return _Sequencer.Contract.Initialize(&_Sequencer.TransactOpts, sequencerSet)
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
