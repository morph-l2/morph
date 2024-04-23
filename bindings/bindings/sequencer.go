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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"SequencerSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_SET_VERIFY_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet0Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet1Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSet2Size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencerSetBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_sequencerSet\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCurrentSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerSet2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSequencerSet\",\"type\":\"address[]\"}],\"name\":\"updateSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b507353000000000000000000000000000000000000126080526080516110b06100475f395f8181610262015261062c01526110b05ff3fe608060405234801561000f575f80fd5b506004361061016e575f3560e01c80636d7f64db116100d25780639b8201a411610088578063a384c12e11610063578063a384c12e146102d2578063b1bdeab3146102da578063eae5b1e3146102e2575f80fd5b80639b8201a414610297578063a224cee7146102ac578063a2e53a94146102bf575f80fd5b80637d99e8ac116100b85780637d99e8ac1461024a578063807de4431461025d57806389609d7414610284575f80fd5b80636d7f64db1461023a57806377d7dffb14610242575f80fd5b806338871fac11610127578063646f7da01161010d578063646f7da01461020557806365fd0f8d1461020e5780636d46e98714610217575f80fd5b806338871fac146101e8578063480265c9146101f0575f80fd5b806328d1357a1161015757806328d1357a1461019e57806329025fcb146101a75780632aae60bd146101b0575f80fd5b80630e06ede81461017257806317f24c2d14610189575b5f80fd5b6007545b6040519081526020015b60405180910390f35b6101916102ea565b6040516101809190610d8f565b61017660065481565b61017660045481565b6101c36101be366004610da8565b61043b565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610180565b600354610176565b6101f8610470565b6040516101809190610dbf565b61017660015481565b61017660025481565b61022a610225366004610e51565b6104aa565b6040519015158152602001610180565b610191610520565b61019161058b565b6101c3610258366004610da8565b6105f6565b6101c37f000000000000000000000000000000000000000000000000000000000000000081565b6101c3610292366004610da8565b610605565b6102aa6102a5366004610e97565b610614565b005b6102aa6102ba366004610e97565b6107dc565b61022a6102cd366004610e51565b610a07565b610176610b61565b600554610176565b610191610b89565b6060600654431061036057600780548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575b5050505050905090565b60045443106103d257600580548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575050505050905090565b600380548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575050505050905090565b6003818154811061044a575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b606060025460036004546005600654600760405160200161049696959493929190610fb4565b604051602081830303815290604052905090565b5f61051a600780548060200260200160405190810160405280929190818152602001828054801561050f57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116104e4575b505050505083610bf4565b92915050565b6060600380548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575050505050905090565b6060600780548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575050505050905090565b6007818154811061044a575f80fd5b6005818154811061044a575f80fd5b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79204c325374616b696e6720636f6e747261637400000000000000000060448201526064015b60405180910390fd5b6006546106c6436002610ff4565b1115610724576004805460029081556006549091556106e6904390610ff4565b600655600580546106f991600391610c67565b506007805461070a91600591610c67565b50805161071e906007906020840190610cb3565b50610739565b8051610737906007906020840190610cb3565b505b60025460036004546005600654600760405160200161075d96959493929190610fb4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206001557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc62816107c3436002610ff4565b6040516107d192919061102c565b60405180910390a150565b5f54610100900460ff16158080156107fa57505f54600160ff909116105b806108135750303b15801561081357505f5460ff166001145b61089f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106af565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108fb575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f825111610965576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c69642073657175656e63657220736574000000000000000000000060448201526064016106af565b8151610978906003906020850190610cb3565b50815161098c906005906020850190610cb3565b5081516109a0906007906020850190610cb3565b508015610a03575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b5f6006544310610a7e5761051a600780548060200260200160405190810160405280929190818152602001828054801561050f57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116104e457505050505083610bf4565b6004544310610af45761051a600580548060200260200160405190810160405280929190818152602001828054801561050f57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116104e457505050505083610bf4565b61051a600380548060200260200160405190810160405280929190818152602001828054801561050f57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116104e457505050505083610bf4565b5f6006544310610b72575060075490565b6004544310610b82575060055490565b5060035490565b6060600580548060200260200160405190810160405280929190818152602001828054801561035657602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161032b575050505050905090565b5f805b8351811015610c5e57838181518110610c1257610c1261104d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c5657600191505061051a565b600101610bf7565b505f9392505050565b828054828255905f5260205f20908101928215610ca3575f5260205f209182015b82811115610ca3578254825591600101919060010190610c88565b50610caf929150610d2b565b5090565b828054828255905f5260205f20908101928215610ca3579160200282015b82811115610ca357825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610cd1565b5b80821115610caf575f8155600101610d2c565b5f815180845260208085019450602084015f5b83811015610d8457815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610d52565b509495945050505050565b602081525f610da16020830184610d3f565b9392505050565b5f60208284031215610db8575f80fd5b5035919050565b5f602080835283518060208501525f5b81811015610deb57858101830151858201604001528201610dcf565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e4c575f80fd5b919050565b5f60208284031215610e61575f80fd5b610da182610e29565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020808385031215610ea8575f80fd5b823567ffffffffffffffff80821115610ebf575f80fd5b818501915085601f830112610ed2575f80fd5b813581811115610ee457610ee4610e6a565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715610f2757610f27610e6a565b604052918252848201925083810185019188831115610f44575f80fd5b938501935b82851015610f6957610f5a85610e29565b84529385019392850192610f49565b98975050505050505050565b5f8154825f526020805f205f5b83811015610d8457815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101610f82565b8681525f610fc56020830188610f75565b868152610fd56020820187610f75565b9050848152610fe76020820185610f75565b9998505050505050505050565b8082018082111561051a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b604081525f61103e6040830185610d3f565b90508260208301529392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212202370b2f86d87cdb54ee30361af545502930a9958a56df338a8b77ea61df4f66364736f6c63430008180033",
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
