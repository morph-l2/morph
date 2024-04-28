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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBatchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchBlockInterval\",\"type\":\"uint256\"}],\"name\":\"BatchBlockIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBatchMaxBytes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchMaxBytes\",\"type\":\"uint256\"}],\"name\":\"BatchMaxBytesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBatchTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchTimeout\",\"type\":\"uint256\"}],\"name\":\"BatchTimeoutUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxChunks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxChunks\",\"type\":\"uint256\"}],\"name\":\"MaxChunksUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalInterval\",\"type\":\"uint256\"}],\"name\":\"ProposalIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"odlRollupEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRollupEpoch\",\"type\":\"uint256\"}],\"name\":\"RollupEpochUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchMaxBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIGov.ProposalData\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentProposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rollupEpoch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"isProposalCanBeApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxChunks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchMaxBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChunks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalInterval\",\"type\":\"uint256\"}],\"name\":\"setProposalInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b5073530000000000000000000000000000000000001260805273530000000000000000000000000000000000000360a05260805160a051611c7e61007d5f395f818161027e015281816103c10152818161062b01528181610dc001526114dd01525f61020b0152611c7e5ff3fe608060405234801561000f575f80fd5b506004361061016e575f3560e01c80638da5cb5b116100d2578063b511328d11610088578063e5aec99511610063578063e5aec9951461037a578063ecded2ae14610383578063f2fde38b14610396575f80fd5b8063b511328d14610320578063bb881e411461035e578063de00d3fd14610367575f80fd5b8063929a9cbe116100b8578063929a9cbe146102a057806396dea936146102a95780639f5039521461030d575f80fd5b80638da5cb5b1461025b5780638e21d5fb14610279575f80fd5b8063639661901161012757806377c793801161010d57806377c79380146101fd578063807de443146102065780638596305214610252575f80fd5b806363966190146101ec578063715018a6146101f5575f80fd5b80632d7aa82b116101575780632d7aa82b1461019a5780634063a84e146101ad57806349c1a581146101c9575f80fd5b80630121b93f146101725780630d61b51914610187575b5f80fd5b6101856101803660046118fe565b6103a9565b005b6101856101953660046118fe565b6107af565b6101856101a8366004611915565b61088f565b6101b6606a5481565b6040519081526020015b60405180910390f35b6101dc6101d7366004611975565b610ccb565b60405190151581526020016101c0565b6101b6606b5481565b610185610ceb565b6101b660675481565b61022d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c0565b6101b660655481565b60335473ffffffffffffffffffffffffffffffffffffffff1661022d565b61022d7f000000000000000000000000000000000000000000000000000000000000000081565b6101b660665481565b6102e56102b73660046118fe565b606c6020525f9081526040902080546001820154600283015460038401546004909401549293919290919085565b604080519586526020860194909452928401919091526060830152608082015260a0016101c0565b61018561031b3660046118fe565b610cfe565b61034961032e3660046118fe565b606d6020525f90815260409020805460019091015460ff1682565b604080519283529015156020830152016101c0565b6101b660685481565b6101856103753660046119a3565b610da8565b6101b660695481565b6101dc6103913660046118fe565b61109e565b6101856103a43660046119b9565b6110e3565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa15801561044e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061047291906119db565b9050806104c65760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b5f828152606d6020526040902060010154829060ff16156105295760405162461bcd60e51b815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104bd565b5f818152606d60205260409020544211156105865760405162461bcd60e51b815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104bd565b61059d335f858152606e6020526040902090611180565b156106105760405162461bcd60e51b815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c00000000000000000000000000000000000000000000000060648201526084016104bd565b610627335f858152606e60205260409020906111ae565b505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa158015610691573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526106d69190810190611a37565b90505f5b81518110156107695761071f8282815181106106f8576106f8611b15565b6020026020010151606e5f8881526020019081526020015f2061118090919063ffffffff16565b6107615761075f82828151811061073857610738611b15565b6020026020010151606e5f8881526020019081526020015f206111cf90919063ffffffff16565b505b6001016106da565b5060038151600261077a9190611b6f565b6107849190611b86565b5f858152606e6020526040902061079a906111f0565b11156107a9576107a9846111f9565b50505050565b5f818152606d6020526040902060010154819060ff16156108125760405162461bcd60e51b815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104bd565b5f818152606d602052604090205442111561086f5760405162461bcd60e51b815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104bd565b5f610879836114d9565b9050801561088a5761088a836111f9565b505050565b5f54610100900460ff16158080156108ad57505f54600160ff909116105b806108c65750303b1580156108c657505f5460ff166001145b6109385760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104bd565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610994575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f87116109e35760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c0000000000000060448201526064016104bd565b5f8311610a325760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104bd565b5f8211610a815760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104bd565b85151580610a8e57508415155b80610a9857508315155b610ae45760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104bd565b610aec6115e9565b606a87905560658690556066859055606784905560688390556069829055604080515f8152602081018990527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a1604080515f8152602081018890527fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b910160405180910390a1604080515f8152602081018790527f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7910160405180910390a1604080515f8152602081018690527fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569910160405180910390a1604080515f8152602081018590527fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a910160405180910390a1604080515f8152602081018490527f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f910160405180910390a18015610cc2575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f828152606e60205260408120610ce29083611180565b90505b92915050565b610cf361166d565b610cfc5f6116d4565b565b610d0661166d565b5f81118015610d1757506065548114155b610d635760405162461bcd60e51b815260206004820152601d60248201527f696e76616c6964206e65772070726f706f73616c20696e74657276616c00000060448201526064016104bd565b606580549082905560408051828152602081018490527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610e4d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e7191906119db565b905080610ec05760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064016104bd565b81608001355f03610f135760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104bd565b5f826060013511610f665760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104bd565b8135151580610f785750602082013515155b80610f865750604082013515155b610fd25760405162461bcd60e51b815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104bd565b606b8054905f610fe183611bbe565b9091555050606b545f908152606c60205260409020829061102d828281358155602082013560018201556040820135600282015560608201356003820155608082013560048201555050565b9050506040518060400160405280606a54426110499190611bf5565b81525f6020918201819052606b548152606d82526040902082518155910151600190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790555050565b5f818152606d602052604081206001015460ff16156110be57505f919050565b5f828152606d60205260409020544211156110da57505f919050565b610ce5826114d9565b6110eb61166d565b73ffffffffffffffffffffffffffffffffffffffff81166111745760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104bd565b61117d816116d4565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610ce2565b5f610ce28373ffffffffffffffffffffffffffffffffffffffff841661174a565b5f610ce28373ffffffffffffffffffffffffffffffffffffffff8416611796565b5f610ce5825490565b5f818152606c60205260409020546065541461126857606580545f838152606c60205260409081902054928390555190917fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b9161125e91848252602082015260400190565b60405180910390a1505b5f818152606c6020526040902060010154606654146112dd57606680545f838152606c60205260409081902060010154928390555190917f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7916112d391848252602082015260400190565b60405180910390a1505b5f818152606c60205260409020600201546067541461135257606780545f838152606c60205260409081902060020154928390555190917fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa69706265699161134891848252602082015260400190565b60405180910390a1505b5f818152606c6020526040902060030154606854146113c757606880545f838152606c60205260409081902060030154928390555190917fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a916113bd91848252602082015260400190565b60405180910390a1505b5f818152606c60205260409020600401546069541461143c57606980545f838152606c60205260409081902060040154928390555190917f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f9161143291848252602082015260400190565b60405180910390a1505b5f818152606d6020908152604091829020600190810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790556065546066546067546068546069548651948552948401929092528285015260608201526080810191909152905182917f146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa919081900360a00190a250565b5f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa158015611543573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526115889190810190611a37565b90505f805b82518110156115c6576115ab8382815181106106f8576106f8611b15565b156115be576115bb826001611bf5565b91505b60010161158d565b506003825160026115d79190611b6f565b6115e19190611b86565b109392505050565b5f54610100900460ff166116655760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bd565b610cfc611879565b60335473ffffffffffffffffffffffffffffffffffffffff163314610cfc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104bd565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f81815260018301602052604081205461178f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610ce5565b505f610ce5565b5f8181526001830160205260408120548015611870575f6117b8600183611c08565b85549091505f906117cb90600190611c08565b905081811461182a575f865f0182815481106117e9576117e9611b15565b905f5260205f200154905080875f01848154811061180957611809611b15565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061183b5761183b611c1b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610ce5565b5f915050610ce5565b5f54610100900460ff166118f55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bd565b610cfc336116d4565b5f6020828403121561190e575f80fd5b5035919050565b5f805f805f8060c0878903121561192a575f80fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b73ffffffffffffffffffffffffffffffffffffffff8116811461117d575f80fd5b5f8060408385031215611986575f80fd5b82359150602083013561199881611954565b809150509250929050565b5f60a082840312156119b3575f80fd5b50919050565b5f602082840312156119c9575f80fd5b81356119d481611954565b9392505050565b5f602082840312156119eb575f80fd5b815180151581146119d4575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b8051611a3281611954565b919050565b5f6020808385031215611a48575f80fd5b825167ffffffffffffffff80821115611a5f575f80fd5b818501915085601f830112611a72575f80fd5b815181811115611a8457611a846119fa565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611ac757611ac76119fa565b604052918252848201925083810185019188831115611ae4575f80fd5b938501935b82851015611b0957611afa85611a27565b84529385019392850192611ae9565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082028115828204841417610ce557610ce5611b42565b5f82611bb9577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bee57611bee611b42565b5060010190565b80820180821115610ce557610ce5611b42565b81810381811115610ce557610ce5611b42565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220360ba45a0cd5cbf99dcc980e12af3ce5eb5560a4b89ccee461aea90411395f1c64736f6c63430008180033",
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

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 proposalID) view returns(bool)
func (_Gov *GovCaller) IsProposalCanBeApproved(opts *bind.CallOpts, proposalID *big.Int) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "isProposalCanBeApproved", proposalID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 proposalID) view returns(bool)
func (_Gov *GovSession) IsProposalCanBeApproved(proposalID *big.Int) (bool, error) {
	return _Gov.Contract.IsProposalCanBeApproved(&_Gov.CallOpts, proposalID)
}

// IsProposalCanBeApproved is a free data retrieval call binding the contract method 0xecded2ae.
//
// Solidity: function isProposalCanBeApproved(uint256 proposalID) view returns(bool)
func (_Gov *GovCallerSession) IsProposalCanBeApproved(proposalID *big.Int) (bool, error) {
	return _Gov.Contract.IsProposalCanBeApproved(&_Gov.CallOpts, proposalID)
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
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovCaller) ProposalData(opts *bind.CallOpts, proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalData", proposalID)

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
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovSession) ProposalData(proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, proposalID)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 proposalID) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovCallerSession) ProposalData(proposalID *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, proposalID)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 endTime, bool approved)
func (_Gov *GovCaller) ProposalInfos(opts *bind.CallOpts, proposalID *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInfos", proposalID)

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
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 endTime, bool approved)
func (_Gov *GovSession) ProposalInfos(proposalID *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, proposalID)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 proposalID) view returns(uint256 endTime, bool approved)
func (_Gov *GovCallerSession) ProposalInfos(proposalID *big.Int) (struct {
	EndTime  *big.Int
	Approved bool
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, proposalID)
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

// CreateProposal is a paid mutator transaction binding the contract method 0xde00d3fd.
//
// Solidity: function createProposal((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactor) CreateProposal(opts *bind.TransactOpts, proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "createProposal", proposal)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xde00d3fd.
//
// Solidity: function createProposal((uint256,uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovSession) CreateProposal(proposal IGovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.CreateProposal(&_Gov.TransactOpts, proposal)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xde00d3fd.
//
// Solidity: function createProposal((uint256,uint256,uint256,uint256,uint256) proposal) returns()
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

// SetProposalInterval is a paid mutator transaction binding the contract method 0x9f503952.
//
// Solidity: function setProposalInterval(uint256 _proposalInterval) returns()
func (_Gov *GovTransactor) SetProposalInterval(opts *bind.TransactOpts, _proposalInterval *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "setProposalInterval", _proposalInterval)
}

// SetProposalInterval is a paid mutator transaction binding the contract method 0x9f503952.
//
// Solidity: function setProposalInterval(uint256 _proposalInterval) returns()
func (_Gov *GovSession) SetProposalInterval(_proposalInterval *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SetProposalInterval(&_Gov.TransactOpts, _proposalInterval)
}

// SetProposalInterval is a paid mutator transaction binding the contract method 0x9f503952.
//
// Solidity: function setProposalInterval(uint256 _proposalInterval) returns()
func (_Gov *GovTransactorSession) SetProposalInterval(_proposalInterval *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SetProposalInterval(&_Gov.TransactOpts, _proposalInterval)
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

// GovBatchMaxBytesUpdatedIterator is returned from FilterBatchMaxBytesUpdated and is used to iterate over the raw logs and unpacked data for BatchMaxBytesUpdated events raised by the Gov contract.
type GovBatchMaxBytesUpdatedIterator struct {
	Event *GovBatchMaxBytesUpdated // Event containing the contract specifics and raw log

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
func (it *GovBatchMaxBytesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovBatchMaxBytesUpdated)
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
		it.Event = new(GovBatchMaxBytesUpdated)
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
func (it *GovBatchMaxBytesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovBatchMaxBytesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovBatchMaxBytesUpdated represents a BatchMaxBytesUpdated event raised by the Gov contract.
type GovBatchMaxBytesUpdated struct {
	OldBatchMaxBytes *big.Int
	NewBatchMaxBytes *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBatchMaxBytesUpdated is a free log retrieval operation binding the contract event 0x11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7.
//
// Solidity: event BatchMaxBytesUpdated(uint256 oldBatchMaxBytes, uint256 newBatchMaxBytes)
func (_Gov *GovFilterer) FilterBatchMaxBytesUpdated(opts *bind.FilterOpts) (*GovBatchMaxBytesUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "BatchMaxBytesUpdated")
	if err != nil {
		return nil, err
	}
	return &GovBatchMaxBytesUpdatedIterator{contract: _Gov.contract, event: "BatchMaxBytesUpdated", logs: logs, sub: sub}, nil
}

// WatchBatchMaxBytesUpdated is a free log subscription operation binding the contract event 0x11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7.
//
// Solidity: event BatchMaxBytesUpdated(uint256 oldBatchMaxBytes, uint256 newBatchMaxBytes)
func (_Gov *GovFilterer) WatchBatchMaxBytesUpdated(opts *bind.WatchOpts, sink chan<- *GovBatchMaxBytesUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "BatchMaxBytesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovBatchMaxBytesUpdated)
				if err := _Gov.contract.UnpackLog(event, "BatchMaxBytesUpdated", log); err != nil {
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

// ParseBatchMaxBytesUpdated is a log parse operation binding the contract event 0x11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7.
//
// Solidity: event BatchMaxBytesUpdated(uint256 oldBatchMaxBytes, uint256 newBatchMaxBytes)
func (_Gov *GovFilterer) ParseBatchMaxBytesUpdated(log types.Log) (*GovBatchMaxBytesUpdated, error) {
	event := new(GovBatchMaxBytesUpdated)
	if err := _Gov.contract.UnpackLog(event, "BatchMaxBytesUpdated", log); err != nil {
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

// GovMaxChunksUpdatedIterator is returned from FilterMaxChunksUpdated and is used to iterate over the raw logs and unpacked data for MaxChunksUpdated events raised by the Gov contract.
type GovMaxChunksUpdatedIterator struct {
	Event *GovMaxChunksUpdated // Event containing the contract specifics and raw log

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
func (it *GovMaxChunksUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovMaxChunksUpdated)
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
		it.Event = new(GovMaxChunksUpdated)
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
func (it *GovMaxChunksUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovMaxChunksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovMaxChunksUpdated represents a MaxChunksUpdated event raised by the Gov contract.
type GovMaxChunksUpdated struct {
	OldMaxChunks *big.Int
	NewMaxChunks *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxChunksUpdated is a free log retrieval operation binding the contract event 0xd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a.
//
// Solidity: event MaxChunksUpdated(uint256 oldMaxChunks, uint256 newMaxChunks)
func (_Gov *GovFilterer) FilterMaxChunksUpdated(opts *bind.FilterOpts) (*GovMaxChunksUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "MaxChunksUpdated")
	if err != nil {
		return nil, err
	}
	return &GovMaxChunksUpdatedIterator{contract: _Gov.contract, event: "MaxChunksUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxChunksUpdated is a free log subscription operation binding the contract event 0xd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a.
//
// Solidity: event MaxChunksUpdated(uint256 oldMaxChunks, uint256 newMaxChunks)
func (_Gov *GovFilterer) WatchMaxChunksUpdated(opts *bind.WatchOpts, sink chan<- *GovMaxChunksUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "MaxChunksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovMaxChunksUpdated)
				if err := _Gov.contract.UnpackLog(event, "MaxChunksUpdated", log); err != nil {
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

// ParseMaxChunksUpdated is a log parse operation binding the contract event 0xd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a.
//
// Solidity: event MaxChunksUpdated(uint256 oldMaxChunks, uint256 newMaxChunks)
func (_Gov *GovFilterer) ParseMaxChunksUpdated(log types.Log) (*GovMaxChunksUpdated, error) {
	event := new(GovMaxChunksUpdated)
	if err := _Gov.contract.UnpackLog(event, "MaxChunksUpdated", log); err != nil {
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
	ProposalID         *big.Int
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	MaxChunks          *big.Int
	RollupEpoch        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
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

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 maxChunks, uint256 rollupEpoch)
func (_Gov *GovFilterer) ParseProposalExecuted(log types.Log) (*GovProposalExecuted, error) {
	event := new(GovProposalExecuted)
	if err := _Gov.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovProposalIntervalUpdatedIterator is returned from FilterProposalIntervalUpdated and is used to iterate over the raw logs and unpacked data for ProposalIntervalUpdated events raised by the Gov contract.
type GovProposalIntervalUpdatedIterator struct {
	Event *GovProposalIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *GovProposalIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovProposalIntervalUpdated)
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
		it.Event = new(GovProposalIntervalUpdated)
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
func (it *GovProposalIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovProposalIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovProposalIntervalUpdated represents a ProposalIntervalUpdated event raised by the Gov contract.
type GovProposalIntervalUpdated struct {
	OldProposalInterval *big.Int
	NewProposalInterval *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterProposalIntervalUpdated is a free log retrieval operation binding the contract event 0x9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b.
//
// Solidity: event ProposalIntervalUpdated(uint256 oldProposalInterval, uint256 newProposalInterval)
func (_Gov *GovFilterer) FilterProposalIntervalUpdated(opts *bind.FilterOpts) (*GovProposalIntervalUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "ProposalIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &GovProposalIntervalUpdatedIterator{contract: _Gov.contract, event: "ProposalIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchProposalIntervalUpdated is a free log subscription operation binding the contract event 0x9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b.
//
// Solidity: event ProposalIntervalUpdated(uint256 oldProposalInterval, uint256 newProposalInterval)
func (_Gov *GovFilterer) WatchProposalIntervalUpdated(opts *bind.WatchOpts, sink chan<- *GovProposalIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "ProposalIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovProposalIntervalUpdated)
				if err := _Gov.contract.UnpackLog(event, "ProposalIntervalUpdated", log); err != nil {
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

// ParseProposalIntervalUpdated is a log parse operation binding the contract event 0x9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b.
//
// Solidity: event ProposalIntervalUpdated(uint256 oldProposalInterval, uint256 newProposalInterval)
func (_Gov *GovFilterer) ParseProposalIntervalUpdated(log types.Log) (*GovProposalIntervalUpdated, error) {
	event := new(GovProposalIntervalUpdated)
	if err := _Gov.contract.UnpackLog(event, "ProposalIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovRollupEpochUpdatedIterator is returned from FilterRollupEpochUpdated and is used to iterate over the raw logs and unpacked data for RollupEpochUpdated events raised by the Gov contract.
type GovRollupEpochUpdatedIterator struct {
	Event *GovRollupEpochUpdated // Event containing the contract specifics and raw log

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
func (it *GovRollupEpochUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovRollupEpochUpdated)
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
		it.Event = new(GovRollupEpochUpdated)
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
func (it *GovRollupEpochUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovRollupEpochUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovRollupEpochUpdated represents a RollupEpochUpdated event raised by the Gov contract.
type GovRollupEpochUpdated struct {
	OdlRollupEpoch *big.Int
	NewRollupEpoch *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRollupEpochUpdated is a free log retrieval operation binding the contract event 0x9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f.
//
// Solidity: event RollupEpochUpdated(uint256 odlRollupEpoch, uint256 newRollupEpoch)
func (_Gov *GovFilterer) FilterRollupEpochUpdated(opts *bind.FilterOpts) (*GovRollupEpochUpdatedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "RollupEpochUpdated")
	if err != nil {
		return nil, err
	}
	return &GovRollupEpochUpdatedIterator{contract: _Gov.contract, event: "RollupEpochUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupEpochUpdated is a free log subscription operation binding the contract event 0x9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f.
//
// Solidity: event RollupEpochUpdated(uint256 odlRollupEpoch, uint256 newRollupEpoch)
func (_Gov *GovFilterer) WatchRollupEpochUpdated(opts *bind.WatchOpts, sink chan<- *GovRollupEpochUpdated) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "RollupEpochUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovRollupEpochUpdated)
				if err := _Gov.contract.UnpackLog(event, "RollupEpochUpdated", log); err != nil {
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

// ParseRollupEpochUpdated is a log parse operation binding the contract event 0x9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f.
//
// Solidity: event RollupEpochUpdated(uint256 odlRollupEpoch, uint256 newRollupEpoch)
func (_Gov *GovFilterer) ParseRollupEpochUpdated(log types.Log) (*GovRollupEpochUpdated, error) {
	event := new(GovRollupEpochUpdated)
	if err := _Gov.contract.UnpackLog(event, "RollupEpochUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
