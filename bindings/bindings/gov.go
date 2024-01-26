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

// GovProposalData is an auto generated low-level Go binding around an user-defined struct.
type GovProposalData struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
	MaxChunks          *big.Int
}

// GovMetaData contains all meta data concerning the Gov contract.
var GovMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_SUBMITTER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchMaxBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rollupEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxChunks\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxChunks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalInfos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seqsVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"}],\"internalType\":\"structGov.ProposalData\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencersVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"propID\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600060015560006002556000600355600060045560006005556000600655600060075534801561003357600080fd5b5073530000000000000000000000000000000000000360805273530000000000000000000000000000000000000560a05260805160a0516111f46100ba600039600081816101b0015261085a015260008181610137015281816101ea01528181610381015281816105020152818161075c01528181610c570152610db401526111f46000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806385963052116100b2578063bb881e4111610081578063d441a16811610066578063d441a1681461032e578063e5aec99514610337578063f59993a41461034057600080fd5b8063bb881e41146102e7578063c6e36a32146102f057600080fd5b80638596305214610215578063929a9cbe1461021e57806396dea93614610227578063b511328d1461028c57600080fd5b80634bbf5252116100ee5780634bbf5252146101ab57806362b8e1b8146101d25780636cb23707146101e557806377c793801461020c57600080fd5b80630121b93f146101205780630c3f3517146101355780632d7aa82b146101815780634063a84e14610194575b600080fd5b61013361012e366004610f41565b610349565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61013361018f366004610f5a565b610938565b61019d600b5481565b604051908152602001610178565b6101577f000000000000000000000000000000000000000000000000000000000000000081565b6101336101e0366004610f9d565b610c1f565b6101577f000000000000000000000000000000000000000000000000000000000000000081565b61019d60045481565b61019d60025481565b61019d60035481565b610264610235366004610f41565b600960205260009081526040902080546001820154600283015460038401546004909401549293919290919085565b604080519586526020860194909452928401919091526060830152608082015260a001610178565b6102c561029a366004610f41565b600a60205260009081526040902080546001820154600283015460039093015460ff90921692909184565b6040805194151585526020850193909352918301526060820152608001610178565b61019d60065481565b61031e6102fe366004611034565b600860209081526000928352604080842090915290825290205460ff1681565b6040519015158152602001610178565b61019d60015481565b61019d60055481565b61019d60075481565b6040517fd1c55fe3000000000000000000000000000000000000000000000000000000008152600060048201819052336024830152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d1c55fe3906044016040805180830381865afa1580156103dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104009190611056565b509050806104555760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b6000828152600a602052604090205460ff166104b35760405162461bcd60e51b815260206004820152601160248201527f70726f706f73616c20696e616374697665000000000000000000000000000000604482015260640161044c565b6040517f342b634500000000000000000000000000000000000000000000000000000000815260006004820181905233602483015290819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063342b6345906044016040805180830381865afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c9190611089565b6000868152600860209081526040808320858452909152902054919350915060ff16156106015760405162461bcd60e51b815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c000000000000000000000000000000000000000000000000606482015260840161044c565b6000848152600a602052604090206002015481146106615760405162461bcd60e51b815260206004820152601060248201527f76657273696f6e206d69736d6174636800000000000000000000000000000000604482015260640161044c565b6000848152600a60205260409020600101544211156106c25760405162461bcd60e51b815260206004820152600860248201527f74696d6520656e64000000000000000000000000000000000000000000000000604482015260640161044c565b6000848152600860209081526040808320858452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155878452600a90925282206003018054919290916107259084906110dc565b90915550506040517f7ad9e3ac000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690637ad9e3ac906024016040805180830381865afa1580156107b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107db9190611089565b50905060036107eb8260026110f5565b6107f59190611132565b6000868152600a6020526040902060030154111561093157600085815260096020526040902060030154600554146108cc576005546040517f965fbb9400000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063965fbb9490602401600060405180830381600087803b1580156108b357600080fd5b505af11580156108c7573d6000803e3d6000fd5b505050505b60008581526009602090815260408083208054600290815560018201546003908155908201546004908155908201546005550154600655600a909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050505050565b600054610100900460ff16158080156109585750600054600160ff909116105b806109725750303b158015610972575060005460ff166001145b6109e45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161044c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a4257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60008711610a925760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c00000000000000604482015260640161044c565b60008311610ae25760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f6368000000000000000000000000604482015260640161044c565b60008211610b325760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964206d6178206368756e6b730000000000000000000000000000604482015260640161044c565b85151580610b3f57508415155b80610b4957508315155b610b955760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d73000000000000000000000000604482015260640161044c565b600b879055600286905560038590556004849055600583905560068290558015610c1657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6040517fd1c55fe3000000000000000000000000000000000000000000000000000000008152600060048201819052336024830152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d1c55fe3906044016040805180830381865afa158015610cb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd69190611056565b50905080610d265760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f7365000000000000604482015260640161044c565b8151151580610d385750602082015115155b80610d465750604082015115155b8015610d555750606082015115155b8015610d645750608082015115155b610db05760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d73000000000000000000000000604482015260640161044c565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639d888e866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e41919061116d565b9050600154811115610e535760018190555b60078054906000610e6383611186565b91905055508260096000600754815260200190815260200160002060008201518160000155602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060800160405280600115158152602001600b5442610ed591906110dc565b8152602080820193909352600060409182018190526007548152600a8452819020825181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169015151781559282015160018401558101516002830155606001516003909101555050565b600060208284031215610f5357600080fd5b5035919050565b60008060008060008060c08789031215610f7357600080fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b600060a08284031215610faf57600080fd5b60405160a0810181811067ffffffffffffffff82111715610ff9577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b6000806040838503121561104757600080fd5b50508035926020909101359150565b6000806040838503121561106957600080fd5b8251801515811461107957600080fd5b6020939093015192949293505050565b6000806040838503121561109c57600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156110ef576110ef6110ad565b92915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561112d5761112d6110ad565b500290565b600082611168577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60006020828403121561117f57600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036111b7576111b76110ad565b506001019056fea26469706673582212201cd9c98a8b89f21b8ddf8705023f874f4ca1d478bdf4b8f2569d6f0595122f6d64736f6c63430008100033",
}

// GovABI is the input ABI used to generate the binding from.
// Deprecated: Use GovMetaData.ABI instead.
var GovABI = GovMetaData.ABI

// GovBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovMetaData.Bin instead.
var GovBin = GovMetaData.Bin

// DeployGov deploys a new Ethereum contract, binding an instance of Gov to it.
func DeployGov(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gov, error) {
	parsed, err := GovMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// Gov is an auto generated Go binding around an Ethereum contract.
type Gov struct {
	GovCaller     // Read-only binding to the contract
	GovTransactor // Write-only binding to the contract
	GovFilterer   // Log filterer for contract events
}

// GovCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovSession struct {
	Contract     *Gov              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovCallerSession struct {
	Contract *GovCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovTransactorSession struct {
	Contract     *GovTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovRaw struct {
	Contract *Gov // Generic contract binding to access the raw methods on
}

// GovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovCallerRaw struct {
	Contract *GovCaller // Generic read-only contract binding to access the raw methods on
}

// GovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovTransactorRaw struct {
	Contract *GovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGov creates a new instance of Gov, bound to a specific deployed contract.
func NewGov(address common.Address, backend bind.ContractBackend) (*Gov, error) {
	contract, err := bindGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// NewGovCaller creates a new read-only instance of Gov, bound to a specific deployed contract.
func NewGovCaller(address common.Address, caller bind.ContractCaller) (*GovCaller, error) {
	contract, err := bindGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovCaller{contract: contract}, nil
}

// NewGovTransactor creates a new write-only instance of Gov, bound to a specific deployed contract.
func NewGovTransactor(address common.Address, transactor bind.ContractTransactor) (*GovTransactor, error) {
	contract, err := bindGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovTransactor{contract: contract}, nil
}

// NewGovFilterer creates a new log filterer instance of Gov, bound to a specific deployed contract.
func NewGovFilterer(address common.Address, filterer bind.ContractFilterer) (*GovFilterer, error) {
	contract, err := bindGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovFilterer{contract: contract}, nil
}

// bindGov binds a generic wrapper to an already deployed contract.
func bindGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.GovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transact(opts, method, params...)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCaller) L2SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "L2_SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SEQUENCERCONTRACT(&_Gov.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SEQUENCERCONTRACT(&_Gov.CallOpts)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovCaller) L2SUBMITTERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "L2_SUBMITTER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SUBMITTERCONTRACT(&_Gov.CallOpts)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SUBMITTERCONTRACT(&_Gov.CallOpts)
}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovCaller) BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovSession) BatchBlockInterval() (*big.Int, error) {
	return _Gov.Contract.BatchBlockInterval(&_Gov.CallOpts)
}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovCallerSession) BatchBlockInterval() (*big.Int, error) {
	return _Gov.Contract.BatchBlockInterval(&_Gov.CallOpts)
}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovCaller) BatchMaxBytes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchMaxBytes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovSession) BatchMaxBytes() (*big.Int, error) {
	return _Gov.Contract.BatchMaxBytes(&_Gov.CallOpts)
}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovCallerSession) BatchMaxBytes() (*big.Int, error) {
	return _Gov.Contract.BatchMaxBytes(&_Gov.CallOpts)
}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovCaller) BatchTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovSession) BatchTimeout() (*big.Int, error) {
	return _Gov.Contract.BatchTimeout(&_Gov.CallOpts)
}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovCallerSession) BatchTimeout() (*big.Int, error) {
	return _Gov.Contract.BatchTimeout(&_Gov.CallOpts)
}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovCaller) L2Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "l2Sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovSession) L2Sequencer() (common.Address, error) {
	return _Gov.Contract.L2Sequencer(&_Gov.CallOpts)
}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovCallerSession) L2Sequencer() (common.Address, error) {
	return _Gov.Contract.L2Sequencer(&_Gov.CallOpts)
}

// MaxChunks is a free data retrieval call binding the contract method 0xbb881e41.
//
// Solidity: function maxChunks() view returns(uint256)
func (_Gov *GovCaller) MaxChunks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "maxChunks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxChunks is a free data retrieval call binding the contract method 0xbb881e41.
//
// Solidity: function maxChunks() view returns(uint256)
func (_Gov *GovSession) MaxChunks() (*big.Int, error) {
	return _Gov.Contract.MaxChunks(&_Gov.CallOpts)
}

// MaxChunks is a free data retrieval call binding the contract method 0xbb881e41.
//
// Solidity: function maxChunks() view returns(uint256)
func (_Gov *GovCallerSession) MaxChunks() (*big.Int, error) {
	return _Gov.Contract.MaxChunks(&_Gov.CallOpts)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch, uint256 maxChunks)
func (_Gov *GovCaller) ProposalData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
	MaxChunks          *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalData", arg0)

	outstruct := new(struct {
		BatchBlockInterval *big.Int
		BatchMaxBytes      *big.Int
		BatchTimeout       *big.Int
		RollupEpoch        *big.Int
		MaxChunks          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchBlockInterval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchMaxBytes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BatchTimeout = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RollupEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxChunks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch, uint256 maxChunks)
func (_Gov *GovSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
	MaxChunks          *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch, uint256 maxChunks)
func (_Gov *GovCallerSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
	MaxChunks          *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovCaller) ProposalInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInfos", arg0)

	outstruct := new(struct {
		Active      bool
		EndTime     *big.Int
		SeqsVersion *big.Int
		Votes       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SeqsVersion = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovSession) ProposalInfos(arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, arg0)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovCallerSession) ProposalInfos(arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, arg0)
}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovCaller) ProposalInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovSession) ProposalInterval() (*big.Int, error) {
	return _Gov.Contract.ProposalInterval(&_Gov.CallOpts)
}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovCallerSession) ProposalInterval() (*big.Int, error) {
	return _Gov.Contract.ProposalInterval(&_Gov.CallOpts)
}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovCaller) ProposalNumbers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalNumbers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovSession) ProposalNumbers() (*big.Int, error) {
	return _Gov.Contract.ProposalNumbers(&_Gov.CallOpts)
}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovCallerSession) ProposalNumbers() (*big.Int, error) {
	return _Gov.Contract.ProposalNumbers(&_Gov.CallOpts)
}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovCaller) RollupEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "rollupEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovSession) RollupEpoch() (*big.Int, error) {
	return _Gov.Contract.RollupEpoch(&_Gov.CallOpts)
}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovCallerSession) RollupEpoch() (*big.Int, error) {
	return _Gov.Contract.RollupEpoch(&_Gov.CallOpts)
}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovCaller) SequencersVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "sequencersVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovSession) SequencersVersion() (*big.Int, error) {
	return _Gov.Contract.SequencersVersion(&_Gov.CallOpts)
}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovCallerSession) SequencersVersion() (*big.Int, error) {
	return _Gov.Contract.SequencersVersion(&_Gov.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovSession) Votes(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Gov.Contract.Votes(&_Gov.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovCallerSession) Votes(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Gov.Contract.Votes(&_Gov.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch, uint256 _maxChunks) returns()
func (_Gov *GovTransactor) Initialize(opts *bind.TransactOpts, _proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int, _maxChunks *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "initialize", _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch, _maxChunks)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch, uint256 _maxChunks) returns()
func (_Gov *GovSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int, _maxChunks *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch, _maxChunks)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch, uint256 _maxChunks) returns()
func (_Gov *GovTransactorSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int, _maxChunks *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch, _maxChunks)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactor) Propose(opts *bind.TransactOpts, proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "propose", proposal)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovSession) Propose(proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactorSession) Propose(proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovTransactor) Vote(opts *bind.TransactOpts, propID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "vote", propID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovSession) Vote(propID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, propID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovTransactorSession) Vote(propID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, propID)
}

// GovInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gov contract.
type GovInitializedIterator struct {
	Event *GovInitialized // Event containing the contract specifics and raw log

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
func (it *GovInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovInitialized)
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
		it.Event = new(GovInitialized)
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
func (it *GovInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovInitialized represents a Initialized event raised by the Gov contract.
type GovInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gov *GovFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovInitializedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovInitializedIterator{contract: _Gov.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gov *GovFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovInitialized) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovInitialized)
				if err := _Gov.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Gov *GovFilterer) ParseInitialized(log types.Log) (*GovInitialized, error) {
	event := new(GovInitialized)
	if err := _Gov.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
