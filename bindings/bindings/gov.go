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

// IGovProposalData is an auto generated low-level Go binding around an user-defined struct.
type IGovProposalData struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}

// GovMetaData contains all meta data concerning the Gov contract.
var GovMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchMaxBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rollupEpoch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"isProposalCanBeApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxChunks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIGov.ProposalData\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040525f6065555f6066555f6067555f6068555f6069555f606b55348015610027575f80fd5b5073530000000000000000000000000000000000001260805273530000000000000000000000000000000000000360a05260805160a05161176f6100955f395f8181610286015281816103bb015281816105f201528181610bd301526110a801525f610213015261176f5ff3fe608060405234801561000f575f80fd5b5060043610610163575f3560e01c806385963052116100c7578063b511328d1161007d578063e5aec99511610063578063e5aec9951461035c578063ecded2ae14610365578063f2fde38b14610378575f80fd5b8063b511328d14610315578063bb881e4114610353575f80fd5b80638e21d5fb116100ad5780638e21d5fb14610281578063929a9cbe146102a857806396dea936146102b1575f80fd5b8063859630521461025a5780638da5cb5b14610263575f80fd5b806349c1a5811161011c578063715018a611610102578063715018a6146101fd57806377c7938014610205578063807de4431461020e575f80fd5b806349c1a581146101c757806362b8e1b8146101ea575f80fd5b80632d7aa82b1161014c5780632d7aa82b1461018f578063371fa854146101a25780634063a84e146101be575f80fd5b80630121b93f146101675780630d61b5191461017c575b5f80fd5b61017a6101753660046113c0565b61038b565b005b61017a61018a3660046113c0565b610776565b61017a61019d3660046113d7565b610856565b6101ab606b5481565b6040519081526020015b60405180910390f35b6101ab606a5481565b6101da6101d5366004611437565b610b83565b60405190151581526020016101b5565b61017a6101f83660046114e1565b610ba3565b61017a610e6d565b6101ab60675481565b6102357f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b5565b6101ab60655481565b60335473ffffffffffffffffffffffffffffffffffffffff16610235565b6102357f000000000000000000000000000000000000000000000000000000000000000081565b6101ab60665481565b6102ed6102bf3660046113c0565b606c6020525f9081526040902080546001820154600283015460038401546004909401549293919290919085565b604080519586526020860194909452928401919091526060830152608082015260a0016101b5565b61033e6103233660046113c0565b606d6020525f90815260409020805460019091015460ff1682565b604080519283529015156020830152016101b5565b6101ab60685481565b6101ab60695481565b6101da6103733660046113c0565b610e80565b61017a61038636600461154f565b610ec5565b6040517f6d46e9870000000000000000000000000000000000000000000000000000000081523360048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636d46e98790602401602060405180830381865afa158015610415573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104399190611571565b90508061048d5760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b5f828152606d6020526040902060010154829060ff16156104f05760405162461bcd60e51b815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f766564000000000000006044820152606401610484565b5f818152606d602052604090205442111561054d5760405162461bcd60e51b815260206004820152601460248201527f70726f706f73616c206f7574206f6620646174650000000000000000000000006044820152606401610484565b5f838152606e602052604090206105649033610f62565b156105d75760405162461bcd60e51b815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c0000000000000000000000000000000000000000000000006064820152608401610484565b5f838152606e602052604090206105ee9033610f90565b505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa158015610658573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261069d9190810190611590565b90505f5b8151811015610730576106e68282815181106106bf576106bf61163d565b6020026020010151606e5f8881526020019081526020015f20610f6290919063ffffffff16565b610728576107268282815181106106ff576106ff61163d565b6020026020010151606e5f8881526020019081526020015f20610fb190919063ffffffff16565b505b6001016106a1565b506003815160026107419190611697565b61074b91906116ae565b5f858152606e6020526040902061076190610fd2565b11156107705761077084610fdb565b50505050565b5f818152606d6020526040902060010154819060ff16156107d95760405162461bcd60e51b815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f766564000000000000006044820152606401610484565b5f818152606d60205260409020544211156108365760405162461bcd60e51b815260206004820152601460248201527f70726f706f73616c206f7574206f6620646174650000000000000000000000006044820152606401610484565b5f610840836110a4565b905080156108515761085183610fdb565b505050565b5f54610100900460ff161580801561087457505f54600160ff909116105b8061088d5750303b15801561088d57505f5460ff166001145b6108ff5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610484565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561095b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f87116109aa5760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c000000000000006044820152606401610484565b5f83116109f95760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964206d6178206368756e6b7300000000000000000000000000006044820152606401610484565b5f8211610a485760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f63680000000000000000000000006044820152606401610484565b85151580610a5557508415155b80610a5f57508315155b610aab5760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d730000000000000000000000006044820152606401610484565b606a87905560658690556066859055606784905560688390556069829055604080518781526020810187905290810185905260608101849052608081018390527f360c363478b9fd9bc8883cd807d952eff8bf91de30034c685cc5755bb17f09799060a00160405180910390a18015610b7a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f828152606e60205260408120610b9a9083610f62565b90505b92915050565b6040517f6d46e9870000000000000000000000000000000000000000000000000000000081523360048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636d46e98790602401602060405180830381865afa158015610c2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c519190611571565b905080610ca05760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f73650000000000006044820152606401610484565b81608001515f03610cf35760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f63680000000000000000000000006044820152606401610484565b5f826060015111610d465760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964206d6178206368756e6b7300000000000000000000000000006044820152606401610484565b8151151580610d585750602082015115155b80610d665750604082015115155b610db25760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d730000000000000000000000006044820152606401610484565b606b54610dc09060016116e6565b606b8190555f908152606c6020908152604091829020845181559084015160018201558184015160028201556060840151600382015560808401516004909101558051808201909152606a548190610e1890426116e6565b81525f6020918201819052606b548152606d82526040902082518155910151600190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790555050565b610e756111b4565b610e7e5f61121b565b565b5f818152606d602052604081206001015460ff1615610ea057505f919050565b5f828152606d6020526040902054421115610ebc57505f919050565b610b9d826110a4565b610ecd6111b4565b73ffffffffffffffffffffffffffffffffffffffff8116610f565760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610484565b610f5f8161121b565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610b9a565b5f610b9a8373ffffffffffffffffffffffffffffffffffffffff8416611291565b5f610b9a8373ffffffffffffffffffffffffffffffffffffffff84166112dd565b5f610b9d825490565b5f818152606c6020908152604080832080546065908155600180830154606690815560028401546067908155600385015460689081556004909501546069908155606d885297869020830180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092559154915490549254955484519283529482015291820152606081019290925260808201527f360c363478b9fd9bc8883cd807d952eff8bf91de30034c685cc5755bb17f09799060a00160405180910390a150565b5f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa15801561110e573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111539190810190611590565b90505f805b8251811015611191576111768382815181106106bf576106bf61163d565b15611189576111868260016116e6565b91505b600101611158565b506003825160026111a29190611697565b6111ac91906116ae565b109392505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e7e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610484565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f8181526001830160205260408120546112d657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b9d565b505f610b9d565b5f81815260018301602052604081205480156113b7575f6112ff6001836116f9565b85549091505f90611312906001906116f9565b9050818114611371575f865f0182815481106113305761133061163d565b905f5260205f200154905080875f0184815481106113505761135061163d565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806113825761138261170c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610b9d565b5f915050610b9d565b5f602082840312156113d0575f80fd5b5035919050565b5f805f805f8060c087890312156113ec575f80fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f5f575f80fd5b5f8060408385031215611448575f80fd5b82359150602083013561145a81611416565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156114d9576114d9611465565b604052919050565b5f60a082840312156114f1575f80fd5b60405160a0810181811067ffffffffffffffff8211171561151457611514611465565b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b5f6020828403121561155f575f80fd5b813561156a81611416565b9392505050565b5f60208284031215611581575f80fd5b8151801515811461156a575f80fd5b5f60208083850312156115a1575f80fd5b825167ffffffffffffffff808211156115b8575f80fd5b818501915085601f8301126115cb575f80fd5b8151818111156115dd576115dd611465565b8060051b91506115ee848301611492565b8181529183018401918481019088841115611607575f80fd5b938501935b83851015611631578451925061162183611416565b828252938501939085019061160c565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082028115828204841417610b9d57610b9d61166a565b5f826116e1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b80820180821115610b9d57610b9d61166a565b81810381811115610b9d57610b9d61166a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220023a0e901c20d933beb833e517419017ed7b6eb8d9c02aeeb09ea2cd885116d164736f6c63430008180033",
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

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Gov *GovCaller) L2STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "L2_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Gov *GovSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2STAKINGCONTRACT(&_Gov.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2STAKINGCONTRACT(&_Gov.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCaller) SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.SEQUENCERCONTRACT(&_Gov.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.SEQUENCERCONTRACT(&_Gov.CallOpts)
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

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 _proposalID) view returns(bool)
func (_Gov *GovCaller) IsProposalCanBeApproved(opts *bind.CallOpts, _proposalID *big.Int) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "isProposalCanBeApproved", _proposalID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 _proposalID) view returns(bool)
func (_Gov *GovSession) IsProposalCanBeApproved(_proposalID *big.Int) (bool, error) {
	return _Gov.Contract.IsProposalCanBeApproved(&_Gov.CallOpts, _proposalID)
}

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 _proposalID) view returns(bool)
func (_Gov *GovCallerSession) IsProposalCanBeApproved(_proposalID *big.Int) (bool, error) {
	return _Gov.Contract.IsProposalCanBeApproved(&_Gov.CallOpts, _proposalID)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 _proposalID, address _voter) view returns(bool)
func (_Gov *GovCaller) IsVoted(opts *bind.CallOpts, _proposalID *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "isVoted", _proposalID, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 _proposalID, address _voter) view returns(bool)
func (_Gov *GovSession) IsVoted(_proposalID *big.Int, _voter common.Address) (bool, error) {
	return _Gov.Contract.IsVoted(&_Gov.CallOpts, _proposalID, _voter)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 _proposalID, address _voter) view returns(bool)
func (_Gov *GovCallerSession) IsVoted(_proposalID *big.Int, _voter common.Address) (bool, error) {
	return _Gov.Contract.IsVoted(&_Gov.CallOpts, _proposalID, _voter)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gov *GovCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gov *GovSession) Owner() (common.Address, error) {
	return _Gov.Contract.Owner(&_Gov.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gov *GovCallerSession) Owner() (common.Address, error) {
	return _Gov.Contract.Owner(&_Gov.CallOpts)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovCaller) ProposalData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalData", arg0)

	outstruct := new(struct {
		BatchBlockInterval *big.Int
		BatchMaxBytes      *big.Int
		BatchTimeout       *big.Int
		MaxChunks          *big.Int
		RollupEpoch        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchBlockInterval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchMaxBytes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BatchTimeout = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxChunks = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RollupEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovCallerSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_Gov *GovCaller) ProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_Gov *GovSession) ProposalID() (*big.Int, error) {
	return _Gov.Contract.ProposalID(&_Gov.CallOpts)
}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_Gov *GovCallerSession) ProposalID() (*big.Int, error) {
	return _Gov.Contract.ProposalID(&_Gov.CallOpts)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(uint256 endTime, bool approved)
func (_Gov *GovCaller) ProposalInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInfos", arg0)

	outstruct := new(struct {
		EndTime  *big.Int
		Approved bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Approved = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(uint256 endTime, bool approved)
func (_Gov *GovSession) ProposalInfos(arg0 *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, arg0)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(uint256 endTime, bool approved)
func (_Gov *GovCallerSession) ProposalInfos(arg0 *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
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

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalID) returns()
func (_Gov *GovTransactor) ExecuteProposal(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "executeProposal", _proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalID) returns()
func (_Gov *GovSession) ExecuteProposal(_proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.ExecuteProposal(&_Gov.TransactOpts, _proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalID) returns()
func (_Gov *GovTransactorSession) ExecuteProposal(_proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.ExecuteProposal(&_Gov.TransactOpts, _proposalID)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _maxChunks, uint256 _rollupEpoch) returns()
func (_Gov *GovTransactor) Initialize(opts *bind.TransactOpts, _proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _maxChunks *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "initialize", _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _maxChunks, _rollupEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _maxChunks, uint256 _rollupEpoch) returns()
func (_Gov *GovSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _maxChunks *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _maxChunks, _rollupEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x2d7aa82b.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _maxChunks, uint256 _rollupEpoch) returns()
func (_Gov *GovTransactorSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _maxChunks *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _maxChunks, _rollupEpoch)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactor) Propose(opts *bind.TransactOpts, proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "propose", proposal)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovSession) Propose(proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// Propose is a paid mutator transaction binding the contract method 0x62b8e1b8.
//
// Solidity: function propose((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactorSession) Propose(proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gov *GovTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gov *GovSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gov.Contract.RenounceOwnership(&_Gov.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gov *GovTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gov.Contract.RenounceOwnership(&_Gov.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gov *GovTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gov *GovSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gov.Contract.TransferOwnership(&_Gov.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gov *GovTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gov.Contract.TransferOwnership(&_Gov.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalID) returns()
func (_Gov *GovTransactor) Vote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "vote", _proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalID) returns()
func (_Gov *GovSession) Vote(_proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, _proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalID) returns()
func (_Gov *GovTransactorSession) Vote(_proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, _proposalID)
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

// GovOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gov contract.
type GovOwnershipTransferredIterator struct {
	Event *GovOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovOwnershipTransferred)
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
		it.Event = new(GovOwnershipTransferred)
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
func (it *GovOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovOwnershipTransferred represents a OwnershipTransferred event raised by the Gov contract.
type GovOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gov *GovFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gov.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovOwnershipTransferredIterator{contract: _Gov.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gov *GovFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gov.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovOwnershipTransferred)
				if err := _Gov.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gov *GovFilterer) ParseOwnershipTransferred(log types.Log) (*GovOwnershipTransferred, error) {
	event := new(GovOwnershipTransferred)
	if err := _Gov.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Gov contract.
type GovProposalExecutedIterator struct {
	Event *GovProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovProposalExecuted)
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
		it.Event = new(GovProposalExecuted)
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
func (it *GovProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovProposalExecuted represents a ProposalExecuted event raised by the Gov contract.
type GovProposalExecuted struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x360c363478b9fd9bc8883cd807d952eff8bf91de30034c685cc5755bb17f0979.
//
// Solidity: event ProposalExecuted(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GovProposalExecutedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GovProposalExecutedIterator{contract: _Gov.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x360c363478b9fd9bc8883cd807d952eff8bf91de30034c685cc5755bb17f0979.
//
// Solidity: event ProposalExecuted(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovProposalExecuted)
				if err := _Gov.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x360c363478b9fd9bc8883cd807d952eff8bf91de30034c685cc5755bb17f0979.
//
// Solidity: event ProposalExecuted(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovFilterer) ParseProposalExecuted(log types.Log) (*GovProposalExecuted, error) {
	event := new(GovProposalExecuted)
	if err := _Gov.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
