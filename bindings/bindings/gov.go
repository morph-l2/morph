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

// IGovProposalData is an auto generated low-level Go binding around an user-defined struct.
type IGovProposalData struct {
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
}

// GovMetaData contains all meta data concerning the Gov contract.
var GovMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBatchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchBlockInterval\",\"type\":\"uint256\"}],\"name\":\"BatchBlockIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBatchTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchTimeout\",\"type\":\"uint256\"}],\"name\":\"BatchTimeoutUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalVotingDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalVotingDuration\",\"type\":\"uint256\"}],\"name\":\"VotingDurationUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deleteTo\",\"type\":\"uint256\"}],\"name\":\"cleanUpExpiredProposals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"}],\"internalType\":\"structIGov.ProposalData\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentProposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_votingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestExecutedProposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_votingDuration\",\"type\":\"uint256\"}],\"name\":\"setVotingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undeletedProposalStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b5073530000000000000000000000000000000000001560805273530000000000000000000000000000000000001760a05261004861004d565b610109565b5f54610100900460ff16156100b85760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610107575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a0516118b561013f5f395f81816102c3015281816103c101528181610d5401526111b101525f61025001526118b55ff3fe608060405234801561000f575f80fd5b506004361061016e575f3560e01c8063807de443116100d2578063a83d073711610088578063d5cc89ba11610063578063d5cc89ba1461037a578063d76930751461038d578063f2fde38b14610396575f80fd5b8063a83d073714610320578063acfd4cd814610329578063b511328d1461033c575f80fd5b80638da5cb5b116100b85780638da5cb5b146102a05780638e21d5fb146102be57806396dea936146102e5575f80fd5b8063807de4431461024b5780638596305214610297575f80fd5b80635bcfadb511610127578063715018a61161010d578063715018a61461020857806374c260cf1461021057806377c7938014610242575f80fd5b80635bcfadb5146101ec57806363966190146101ff575f80fd5b8063132002fc11610157578063132002fc1461019a57806349c1a581146101b65780634ec81af1146101d9575f80fd5b80630121b93f146101725780630d61b51914610187575b5f80fd5b6101856101803660046115a5565b6103a9565b005b6101856101953660046115a5565b6106f1565b6101a360675481565b6040519081526020015b60405180910390f35b6101c96101c43660046115dd565b6107a5565b60405190151581526020016101ad565b6101856101e736600461160b565b6107c5565b6101856101fa3660046115a5565b610b01565b6101a360685481565b610185610bd0565b61022361021e3660046115a5565b610be3565b60408051931515845291151560208401521515908201526060016101ad565b6101a360665481565b6102727f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ad565b6101a360655481565b60335473ffffffffffffffffffffffffffffffffffffffff16610272565b6102727f000000000000000000000000000000000000000000000000000000000000000081565b61030b6102f33660046115a5565b606a6020525f90815260409020805460019091015482565b604080519283526020830191909152016101ad565b6101a3606d5481565b6101a3610337366004611643565b610d3b565b61036561034a3660046115a5565b606b6020525f90815260409020805460019091015460ff1682565b604080519283529015156020830152016101ad565b6101856103883660046115a5565b610fb6565b6101a360695481565b6101856103a4366004611659565b6110c1565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa15801561044e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610472919061167b565b9050806104c65760405162461bcd60e51b815260206004820152601660248201527f6f6e6c792073657175656e63657220616c6c6f7765640000000000000000000060448201526064015b60405180910390fd5b6068548211156105185760405162461bcd60e51b815260206004820152601260248201527f696e76616c69642070726f706f73616c4944000000000000000000000000000060448201526064016104bd565b606d5482116105695760405162461bcd60e51b815260206004820152601260248201527f657870697265642070726f706f73616c4944000000000000000000000000000060448201526064016104bd565b6069548210156105bb5760405162461bcd60e51b815260206004820152600f60248201527f70726f706f73616c207072756e6564000000000000000000000000000000000060448201526064016104bd565b5f828152606b60205260409020805460019091015460ff16806105dc575080155b806105e657504281105b156106335760405162461bcd60e51b815260206004820152601060248201527f766f74696e672068617320656e6465640000000000000000000000000000000060448201526064016104bd565b61064a335f858152606c602052604090209061115e565b156106bd5760405162461bcd60e51b815260206004820152602960248201527f73657175656e63657220616c726561647920766f74656420666f72207468697360448201527f2070726f706f73616c000000000000000000000000000000000000000000000060648201526084016104bd565b6106d4335f858152606c602052604090209061118c565b506106de836111ad565b156106ec576106ec836112e4565b505050565b5f806106fc83610be3565b5091509150811561074f5760405162461bcd60e51b815260206004820152601060248201527f766f74696e672068617320656e6465640000000000000000000000000000000060448201526064016104bd565b8061079c5760405162461bcd60e51b815260206004820181905260248201527f70726f706f73616c20686173206e6f74206265656e207061737365642079657460448201526064016104bd565b6106ec836112e4565b5f828152606c602052604081206107bc908361115e565b90505b92915050565b5f54610100900460ff16158080156107e357505f54600160ff909116105b806107fc5750303b1580156107fc57505f5460ff166001145b61086e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104bd565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108ca575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff851661092d5760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964206f776e65722061646472657373000000000000000000000060448201526064016104bd565b5f841161097c5760405162461bcd60e51b815260206004820181905260248201527f696e76616c69642070726f706f73616c20766f74696e67206475726174696f6e60448201526064016104bd565b8215158061098957508115155b6109d55760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104bd565b6109de8561144e565b606784905560658390556066829055604080515f8152602081018690527ffe810823e41a0cf27003f3eac9c17098028ba0aece75bd9783a8da7f75fb3aa3910160405180910390a1604080515f8152602081018590527fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b910160405180910390a1604080515f8152602081018490527fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569910160405180910390a18015610afa575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610b096114c4565b5f81118015610b1a57506067548114155b610b8b5760405162461bcd60e51b8152602060048201526024808201527f696e76616c6964206e65772070726f706f73616c20766f74696e67206475726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016104bd565b606780549082905560408051828152602081018490527ffe810823e41a0cf27003f3eac9c17098028ba0aece75bd9783a8da7f75fb3aa3910160405180910390a15050565b610bd86114c4565b610be15f61144e565b565b5f805f606854841115610c385760405162461bcd60e51b815260206004820152601260248201527f696e76616c69642070726f706f73616c4944000000000000000000000000000060448201526064016104bd565b606d54841015610c8a5760405162461bcd60e51b815260206004820152601060248201527f657870697265642070726f706f73616c0000000000000000000000000000000060448201526064016104bd565b606954841015610cdc5760405162461bcd60e51b815260206004820152600f60248201527f70726f706f73616c207072756e6564000000000000000000000000000000000060448201526064016104bd565b606d548403610cf357506001915081905080610d34565b5f848152606b602052604090206001810154905460ff909116908180610d17575080155b80610d2157504281105b610d2a876111ad565b8394509450945050505b9193909250565b5f8073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610de1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e05919061167b565b905080610e545760405162461bcd60e51b815260206004820152601660248201527f6f6e6c792073657175656e63657220616c6c6f7765640000000000000000000060448201526064016104bd565b8235151580610e665750602083013515155b610eb25760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104bd565b60688054905f610ec1836116c7565b90915550506068545f908152606a602090815260409091208435815590840135600182015550604051806040016040528060675442610f0091906116fe565b81525f60209182018190526068548152606b82526040902082518155910151600190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790553360685460408051863581526020808801359082015273ffffffffffffffffffffffffffffffffffffffff93909316927fd38d9a9a102af286aea1cbcadb8aab8446a90859626389db3710ed4fea4c2c39910160405180910390a35050606854919050565b606d54811061102d5760405162461bcd60e51b815260206004820152603e60248201527f6f6e6c7920616c6c6f7720746f2064656c657465207468652070726f706f736160448201527f6c206265666f72206c6174657374207061737365642070726f706f73616c000060648201526084016104bd565b6069545b8181116110af575f818152606a602090815260408083208381556001908101849055606b83528184208481550180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055606c90915281209081816110988282611577565b5050505080806110a7906116c7565b915050611031565b506110bb8160016116fe565b60695550565b6110c96114c4565b73ffffffffffffffffffffffffffffffffffffffff81166111525760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104bd565b61115b8161144e565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156107bc565b5f6107bc8373ffffffffffffffffffffffffffffffffffffffff841661152b565b5f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa158015611217573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261125c919081019061174e565b90505f805b82518110156112c1576112a683828151811061127f5761127f61182c565b6020026020010151606c5f8881526020019081526020015f2061115e90919063ffffffff16565b156112b9576112b68260016116fe565b91505b600101611261565b506003825160026112d29190611859565b6112dc9190611870565b109392505050565b606d8190555f818152606a60205260409020546065541461135857606580545f838152606a60205260409081902054928390555190917fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b9161134e91848252602082015260400190565b60405180910390a1505b5f818152606a6020526040902060010154606654146113cd57606680545f838152606a60205260409081902060010154928390555190917fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569916113c391848252602082015260400190565b60405180910390a1505b5f818152606b6020908152604091829020600190810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909117905560655460665483519182529181019190915282917f3842d9e8746c3c1b0c5e717a88df9c9ae97f2f8c0e01e42dce160aafdb3840e9910160405180910390a250565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610be15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104bd565b5f81815260018301602052604081205461157057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556107bf565b505f6107bf565b5080545f8255905f5260205f209081019061115b91905b808211156115a1575f815560010161158e565b5090565b5f602082840312156115b5575f80fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461115b575f80fd5b5f80604083850312156115ee575f80fd5b823591506020830135611600816115bc565b809150509250929050565b5f805f806080858703121561161e575f80fd5b8435611629816115bc565b966020860135965060408601359560600135945092505050565b5f60408284031215611653575f80fd5b50919050565b5f60208284031215611669575f80fd5b8135611674816115bc565b9392505050565b5f6020828403121561168b575f80fd5b81518015158114611674575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116f7576116f761169a565b5060010190565b808201808211156107bf576107bf61169a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b8051611749816115bc565b919050565b5f602080838503121561175f575f80fd5b825167ffffffffffffffff80821115611776575f80fd5b818501915085601f830112611789575f80fd5b81518181111561179b5761179b611711565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156117de576117de611711565b6040529182528482019250838101850191888311156117fb575f80fd5b938501935b82851015611820576118118561173e565b84529385019392850192611800565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80820281158282048414176107bf576107bf61169a565b5f826118a3577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a",
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
	parsed, err := GovMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() view returns(uint256)
func (_Gov *GovCaller) CurrentProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "currentProposalID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() view returns(uint256)
func (_Gov *GovSession) CurrentProposalID() (*big.Int, error) {
	return _Gov.Contract.CurrentProposalID(&_Gov.CallOpts)
}

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() view returns(uint256)
func (_Gov *GovCallerSession) CurrentProposalID() (*big.Int, error) {
	return _Gov.Contract.CurrentProposalID(&_Gov.CallOpts)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 proposalID, address voter) view returns(bool)
func (_Gov *GovCaller) IsVoted(opts *bind.CallOpts, proposalID *big.Int, voter common.Address) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "isVoted", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 proposalID, address voter) view returns(bool)
func (_Gov *GovSession) IsVoted(proposalID *big.Int, voter common.Address) (bool, error) {
	return _Gov.Contract.IsVoted(&_Gov.CallOpts, proposalID, voter)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 proposalID, address voter) view returns(bool)
func (_Gov *GovCallerSession) IsVoted(proposalID *big.Int, voter common.Address) (bool, error) {
	return _Gov.Contract.IsVoted(&_Gov.CallOpts, proposalID, voter)
}

// LatestExecutedProposalID is a free data retrieval call binding the contract method 0xa83d0737.
//
// Solidity: function latestExecutedProposalID() view returns(uint256)
func (_Gov *GovCaller) LatestExecutedProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "latestExecutedProposalID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestExecutedProposalID is a free data retrieval call binding the contract method 0xa83d0737.
//
// Solidity: function latestExecutedProposalID() view returns(uint256)
func (_Gov *GovSession) LatestExecutedProposalID() (*big.Int, error) {
	return _Gov.Contract.LatestExecutedProposalID(&_Gov.CallOpts)
}

// LatestExecutedProposalID is a free data retrieval call binding the contract method 0xa83d0737.
//
// Solidity: function latestExecutedProposalID() view returns(uint256)
func (_Gov *GovCallerSession) LatestExecutedProposalID() (*big.Int, error) {
	return _Gov.Contract.LatestExecutedProposalID(&_Gov.CallOpts)
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
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovCaller) ProposalData(opts *bind.CallOpts, proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalData", proposalID)

	outstruct := new(struct {
		BatchBlockInterval *big.Int
		BatchTimeout       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchBlockInterval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchTimeout = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovSession) ProposalData(proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, proposalID)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovCallerSession) ProposalData(proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, proposalID)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 expirationTime, bool executed)
func (_Gov *GovCaller) ProposalInfos(opts *bind.CallOpts, proposalID *big.Int) (struct {
	ExpirationTime *big.Int
	Executed       bool
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInfos", proposalID)

	outstruct := new(struct {
		ExpirationTime *big.Int
		Executed       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpirationTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 expirationTime, bool executed)
func (_Gov *GovSession) ProposalInfos(proposalID *big.Int) (struct {
	ExpirationTime *big.Int
	Executed       bool
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, proposalID)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 expirationTime, bool executed)
func (_Gov *GovCallerSession) ProposalInfos(proposalID *big.Int) (struct {
	ExpirationTime *big.Int
	Executed       bool
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, proposalID)
}

// ProposalStatus is a free data retrieval call binding the contract method 0x74c260cf.
//
// Solidity: function proposalStatus(uint256 proposalID) view returns(bool, bool, bool)
func (_Gov *GovCaller) ProposalStatus(opts *bind.CallOpts, proposalID *big.Int) (bool, bool, bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalStatus", proposalID)

	if err != nil {
		return *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// ProposalStatus is a free data retrieval call binding the contract method 0x74c260cf.
//
// Solidity: function proposalStatus(uint256 proposalID) view returns(bool, bool, bool)
func (_Gov *GovSession) ProposalStatus(proposalID *big.Int) (bool, bool, bool, error) {
	return _Gov.Contract.ProposalStatus(&_Gov.CallOpts, proposalID)
}

// ProposalStatus is a free data retrieval call binding the contract method 0x74c260cf.
//
// Solidity: function proposalStatus(uint256 proposalID) view returns(bool, bool, bool)
func (_Gov *GovCallerSession) ProposalStatus(proposalID *big.Int) (bool, bool, bool, error) {
	return _Gov.Contract.ProposalStatus(&_Gov.CallOpts, proposalID)
}

// UndeletedProposalStart is a free data retrieval call binding the contract method 0xd7693075.
//
// Solidity: function undeletedProposalStart() view returns(uint256)
func (_Gov *GovCaller) UndeletedProposalStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "undeletedProposalStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndeletedProposalStart is a free data retrieval call binding the contract method 0xd7693075.
//
// Solidity: function undeletedProposalStart() view returns(uint256)
func (_Gov *GovSession) UndeletedProposalStart() (*big.Int, error) {
	return _Gov.Contract.UndeletedProposalStart(&_Gov.CallOpts)
}

// UndeletedProposalStart is a free data retrieval call binding the contract method 0xd7693075.
//
// Solidity: function undeletedProposalStart() view returns(uint256)
func (_Gov *GovCallerSession) UndeletedProposalStart() (*big.Int, error) {
	return _Gov.Contract.UndeletedProposalStart(&_Gov.CallOpts)
}

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_Gov *GovCaller) VotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "votingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_Gov *GovSession) VotingDuration() (*big.Int, error) {
	return _Gov.Contract.VotingDuration(&_Gov.CallOpts)
}

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_Gov *GovCallerSession) VotingDuration() (*big.Int, error) {
	return _Gov.Contract.VotingDuration(&_Gov.CallOpts)
}

// CleanUpExpiredProposals is a paid mutator transaction binding the contract method 0xd5cc89ba.
//
// Solidity: function cleanUpExpiredProposals(uint256 deleteTo) returns()
func (_Gov *GovTransactor) CleanUpExpiredProposals(opts *bind.TransactOpts, deleteTo *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "cleanUpExpiredProposals", deleteTo)
}

// CleanUpExpiredProposals is a paid mutator transaction binding the contract method 0xd5cc89ba.
//
// Solidity: function cleanUpExpiredProposals(uint256 deleteTo) returns()
func (_Gov *GovSession) CleanUpExpiredProposals(deleteTo *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.CleanUpExpiredProposals(&_Gov.TransactOpts, deleteTo)
}

// CleanUpExpiredProposals is a paid mutator transaction binding the contract method 0xd5cc89ba.
//
// Solidity: function cleanUpExpiredProposals(uint256 deleteTo) returns()
func (_Gov *GovTransactorSession) CleanUpExpiredProposals(deleteTo *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.CleanUpExpiredProposals(&_Gov.TransactOpts, deleteTo)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xacfd4cd8.
//
// Solidity: function createProposal((uint256,uint256) proposal) returns(uint256)
func (_Gov *GovTransactor) CreateProposal(opts *bind.TransactOpts, proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "createProposal", proposal)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xacfd4cd8.
//
// Solidity: function createProposal((uint256,uint256) proposal) returns(uint256)
func (_Gov *GovSession) CreateProposal(proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.CreateProposal(&_Gov.TransactOpts, proposal)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xacfd4cd8.
//
// Solidity: function createProposal((uint256,uint256) proposal) returns(uint256)
func (_Gov *GovTransactorSession) CreateProposal(proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.CreateProposal(&_Gov.TransactOpts, proposal)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalID) returns()
func (_Gov *GovTransactor) ExecuteProposal(opts *bind.TransactOpts, proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "executeProposal", proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalID) returns()
func (_Gov *GovSession) ExecuteProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.ExecuteProposal(&_Gov.TransactOpts, proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 proposalID) returns()
func (_Gov *GovTransactorSession) ExecuteProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.ExecuteProposal(&_Gov.TransactOpts, proposalID)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _votingDuration, uint256 _batchBlockInterval, uint256 _batchTimeout) returns()
func (_Gov *GovTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _votingDuration *big.Int, _batchBlockInterval *big.Int, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "initialize", _owner, _votingDuration, _batchBlockInterval, _batchTimeout)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _votingDuration, uint256 _batchBlockInterval, uint256 _batchTimeout) returns()
func (_Gov *GovSession) Initialize(_owner common.Address, _votingDuration *big.Int, _batchBlockInterval *big.Int, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _owner, _votingDuration, _batchBlockInterval, _batchTimeout)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _votingDuration, uint256 _batchBlockInterval, uint256 _batchTimeout) returns()
func (_Gov *GovTransactorSession) Initialize(_owner common.Address, _votingDuration *big.Int, _batchBlockInterval *big.Int, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _owner, _votingDuration, _batchBlockInterval, _batchTimeout)
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

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 _votingDuration) returns()
func (_Gov *GovTransactor) SetVotingDuration(opts *bind.TransactOpts, _votingDuration *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "setVotingDuration", _votingDuration)
}

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 _votingDuration) returns()
func (_Gov *GovSession) SetVotingDuration(_votingDuration *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SetVotingDuration(&_Gov.TransactOpts, _votingDuration)
}

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 _votingDuration) returns()
func (_Gov *GovTransactorSession) SetVotingDuration(_votingDuration *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SetVotingDuration(&_Gov.TransactOpts, _votingDuration)
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
// Solidity: function vote(uint256 proposalID) returns()
func (_Gov *GovTransactor) Vote(opts *bind.TransactOpts, proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "vote", proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposalID) returns()
func (_Gov *GovSession) Vote(proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposalID) returns()
func (_Gov *GovTransactorSession) Vote(proposalID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, proposalID)
}

// GovBatchBlockIntervalUpdatedIterator is returned from FilterBatchBlockIntervalUpdated and is used to iterate over the raw logs and unpacked data for BatchBlockIntervalUpdated events raised by the Gov contract.
type GovBatchBlockIntervalUpdatedIterator struct {
	Event *GovBatchBlockIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *GovBatchBlockIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovBatchBlockIntervalUpdated)
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
		it.Event = new(GovBatchBlockIntervalUpdated)
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
func (it *GovBatchBlockIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovBatchBlockIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovBatchBlockIntervalUpdated represents a BatchBlockIntervalUpdated event raised by the Gov contract.
type GovBatchBlockIntervalUpdated struct {
	OldBatchBlockInterval *big.Int
	NewBatchBlockInterval *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBatchBlockIntervalUpdated is a free log retrieval operation binding the contract event 0xa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b.
//
// Solidity: event BatchBlockIntervalUpdated(uint256 oldBatchBlockInterval, uint256 newBatchBlockInterval)
func (_Gov *GovFilterer) FilterBatchBlockIntervalUpdated(opts *bind.FilterOpts) (*GovBatchBlockIntervalUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "BatchBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &GovBatchBlockIntervalUpdatedIterator{contract: _Gov.contract, event: "BatchBlockIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchBatchBlockIntervalUpdated is a free log subscription operation binding the contract event 0xa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b.
//
// Solidity: event BatchBlockIntervalUpdated(uint256 oldBatchBlockInterval, uint256 newBatchBlockInterval)
func (_Gov *GovFilterer) WatchBatchBlockIntervalUpdated(opts *bind.WatchOpts, sink chan<- *GovBatchBlockIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "BatchBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovBatchBlockIntervalUpdated)
				if err := _Gov.contract.UnpackLog(event, "BatchBlockIntervalUpdated", log); err != nil {
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

// ParseBatchBlockIntervalUpdated is a log parse operation binding the contract event 0xa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b.
//
// Solidity: event BatchBlockIntervalUpdated(uint256 oldBatchBlockInterval, uint256 newBatchBlockInterval)
func (_Gov *GovFilterer) ParseBatchBlockIntervalUpdated(log types.Log) (*GovBatchBlockIntervalUpdated, error) {
	event := new(GovBatchBlockIntervalUpdated)
	if err := _Gov.contract.UnpackLog(event, "BatchBlockIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovBatchTimeoutUpdatedIterator is returned from FilterBatchTimeoutUpdated and is used to iterate over the raw logs and unpacked data for BatchTimeoutUpdated events raised by the Gov contract.
type GovBatchTimeoutUpdatedIterator struct {
	Event *GovBatchTimeoutUpdated // Event containing the contract specifics and raw log

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
func (it *GovBatchTimeoutUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovBatchTimeoutUpdated)
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
		it.Event = new(GovBatchTimeoutUpdated)
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
func (it *GovBatchTimeoutUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovBatchTimeoutUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovBatchTimeoutUpdated represents a BatchTimeoutUpdated event raised by the Gov contract.
type GovBatchTimeoutUpdated struct {
	OldBatchTimeout *big.Int
	NewBatchTimeout *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBatchTimeoutUpdated is a free log retrieval operation binding the contract event 0xab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569.
//
// Solidity: event BatchTimeoutUpdated(uint256 oldBatchTimeout, uint256 newBatchTimeout)
func (_Gov *GovFilterer) FilterBatchTimeoutUpdated(opts *bind.FilterOpts) (*GovBatchTimeoutUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "BatchTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return &GovBatchTimeoutUpdatedIterator{contract: _Gov.contract, event: "BatchTimeoutUpdated", logs: logs, sub: sub}, nil
}

// WatchBatchTimeoutUpdated is a free log subscription operation binding the contract event 0xab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569.
//
// Solidity: event BatchTimeoutUpdated(uint256 oldBatchTimeout, uint256 newBatchTimeout)
func (_Gov *GovFilterer) WatchBatchTimeoutUpdated(opts *bind.WatchOpts, sink chan<- *GovBatchTimeoutUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "BatchTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovBatchTimeoutUpdated)
				if err := _Gov.contract.UnpackLog(event, "BatchTimeoutUpdated", log); err != nil {
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

// ParseBatchTimeoutUpdated is a log parse operation binding the contract event 0xab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569.
//
// Solidity: event BatchTimeoutUpdated(uint256 oldBatchTimeout, uint256 newBatchTimeout)
func (_Gov *GovFilterer) ParseBatchTimeoutUpdated(log types.Log) (*GovBatchTimeoutUpdated, error) {
	event := new(GovBatchTimeoutUpdated)
	if err := _Gov.contract.UnpackLog(event, "BatchTimeoutUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// GovProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Gov contract.
type GovProposalCreatedIterator struct {
	Event *GovProposalCreated // Event containing the contract specifics and raw log

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
func (it *GovProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovProposalCreated)
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
		it.Event = new(GovProposalCreated)
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
func (it *GovProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovProposalCreated represents a ProposalCreated event raised by the Gov contract.
type GovProposalCreated struct {
	ProposalID         *big.Int
	Creator            common.Address
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xd38d9a9a102af286aea1cbcadb8aab8446a90859626389db3710ed4fea4c2c39.
//
// Solidity: event ProposalCreated(uint256 indexed proposalID, address indexed creator, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalID []*big.Int, creator []common.Address) (*GovProposalCreatedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Gov.contract.FilterLogs(opts, "ProposalCreated", proposalIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &GovProposalCreatedIterator{contract: _Gov.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xd38d9a9a102af286aea1cbcadb8aab8446a90859626389db3710ed4fea4c2c39.
//
// Solidity: event ProposalCreated(uint256 indexed proposalID, address indexed creator, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovProposalCreated, proposalID []*big.Int, creator []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Gov.contract.WatchLogs(opts, "ProposalCreated", proposalIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovProposalCreated)
				if err := _Gov.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xd38d9a9a102af286aea1cbcadb8aab8446a90859626389db3710ed4fea4c2c39.
//
// Solidity: event ProposalCreated(uint256 indexed proposalID, address indexed creator, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) ParseProposalCreated(log types.Log) (*GovProposalCreated, error) {
	event := new(GovProposalCreated)
	if err := _Gov.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
	ProposalID         *big.Int
	BatchBlockInterval *big.Int
	BatchTimeout       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x3842d9e8746c3c1b0c5e717a88df9c9ae97f2f8c0e01e42dce160aafdb3840e9.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalID []*big.Int) (*GovProposalExecutedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Gov.contract.FilterLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &GovProposalExecutedIterator{contract: _Gov.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x3842d9e8746c3c1b0c5e717a88df9c9ae97f2f8c0e01e42dce160aafdb3840e9.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovProposalExecuted, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Gov.contract.WatchLogs(opts, "ProposalExecuted", proposalIDRule)
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x3842d9e8746c3c1b0c5e717a88df9c9ae97f2f8c0e01e42dce160aafdb3840e9.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchTimeout)
func (_Gov *GovFilterer) ParseProposalExecuted(log types.Log) (*GovProposalExecuted, error) {
	event := new(GovProposalExecuted)
	if err := _Gov.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovVotingDurationUpdatedIterator is returned from FilterVotingDurationUpdated and is used to iterate over the raw logs and unpacked data for VotingDurationUpdated events raised by the Gov contract.
type GovVotingDurationUpdatedIterator struct {
	Event *GovVotingDurationUpdated // Event containing the contract specifics and raw log

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
func (it *GovVotingDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovVotingDurationUpdated)
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
		it.Event = new(GovVotingDurationUpdated)
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
func (it *GovVotingDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovVotingDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovVotingDurationUpdated represents a VotingDurationUpdated event raised by the Gov contract.
type GovVotingDurationUpdated struct {
	OldProposalVotingDuration *big.Int
	NewProposalVotingDuration *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterVotingDurationUpdated is a free log retrieval operation binding the contract event 0xfe810823e41a0cf27003f3eac9c17098028ba0aece75bd9783a8da7f75fb3aa3.
//
// Solidity: event VotingDurationUpdated(uint256 oldProposalVotingDuration, uint256 newProposalVotingDuration)
func (_Gov *GovFilterer) FilterVotingDurationUpdated(opts *bind.FilterOpts) (*GovVotingDurationUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "VotingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &GovVotingDurationUpdatedIterator{contract: _Gov.contract, event: "VotingDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchVotingDurationUpdated is a free log subscription operation binding the contract event 0xfe810823e41a0cf27003f3eac9c17098028ba0aece75bd9783a8da7f75fb3aa3.
//
// Solidity: event VotingDurationUpdated(uint256 oldProposalVotingDuration, uint256 newProposalVotingDuration)
func (_Gov *GovFilterer) WatchVotingDurationUpdated(opts *bind.WatchOpts, sink chan<- *GovVotingDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "VotingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovVotingDurationUpdated)
				if err := _Gov.contract.UnpackLog(event, "VotingDurationUpdated", log); err != nil {
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

// ParseVotingDurationUpdated is a log parse operation binding the contract event 0xfe810823e41a0cf27003f3eac9c17098028ba0aece75bd9783a8da7f75fb3aa3.
//
// Solidity: event VotingDurationUpdated(uint256 oldProposalVotingDuration, uint256 newProposalVotingDuration)
func (_Gov *GovFilterer) ParseVotingDurationUpdated(log types.Log) (*GovVotingDurationUpdated, error) {
	event := new(GovVotingDurationUpdated)
	if err := _Gov.contract.UnpackLog(event, "VotingDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
