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

// TypesBatchInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesBatchInfo struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}

// TypesEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesEpochInfo struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}

// SubmitterMetaData contains all meta data concerning the Submitter contract.
var SubmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ACKRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_GOV_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ackRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmedBatchs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"epochUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"getConfirmedBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BatchInfo\",\"name\":\"batchInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EpochInfo\",\"name\":\"epochInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"getTurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"}],\"name\":\"sequencersUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b5073530000000000000000000000000000000000000360805273530000000000000000000000000000000000000460a05260805160a05161180a61007d5f395f818161015201528181610b8901528181610cb10152610efa01525f81816102660152610532015261180a5ff3fe608060405234801561000f575f80fd5b5060043610610149575f3560e01c80637e4fa700116100c7578063a5af40d11161007d578063c6b61e4c11610063578063c6b61e4c1461037d578063e8e39925146103c2578063f2fde38b14610444575f80fd5b8063a5af40d11461031f578063bc0bc6ba14610332575f80fd5b8063843e8a7b116100ad578063843e8a7b146102b45780638da5cb5b146102ee578063965fbb941461030c575f80fd5b80637e4fa700146102a35780638129fc1c146102ac575f80fd5b806322caba241161011c5780636cb23707116101025780636cb23707146102615780636d8ce3d214610288578063715018a61461029b575f80fd5b806322caba24146102375780635c14c3141461024a575f80fd5b8063047d0b6e1461014d57806304b123851461019e57806315123258146101c657806316e2994a14610222575b5f80fd5b6101747f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101b16101ac3660046114ea565b610457565b60408051928352602083019190915201610195565b6101d96101d43660046114ea565b610483565b6040516101959190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b610235610230366004611556565b61051a565b005b610235610245366004611634565b61062f565b61025360665481565b604051908152602001610195565b6101747f000000000000000000000000000000000000000000000000000000000000000081565b6102536102963660046114ea565b61080c565b610235610833565b61025360655481565b610235610846565b6102bc6109d2565b6040805173ffffffffffffffffffffffffffffffffffffffff9094168452602084019290925290820152606001610195565b60335473ffffffffffffffffffffffffffffffffffffffff16610174565b61023561031a3660046114ea565b610c99565b6101b161032d366004611676565b610db0565b6103456103403660046114ea565b6110e6565b60408051825173ffffffffffffffffffffffffffffffffffffffff168152602080840151908201529181015190820152606001610195565b6102bc61038b3660046114ea565b60686020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116919083565b61040d6103d03660046114ea565b60676020525f9081526040902080546001820154600283015460039093015473ffffffffffffffffffffffffffffffffffffffff90921692909184565b6040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152608001610195565b610235610452366004611676565b61116b565b606a8181548110610466575f80fd5b5f9182526020909120600290910201805460019091015490915082565b6104c060405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81525090565b505f908152606760209081526040918290208251608081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281015492820192909252600390910154606082015290565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c79206c322073657175656e63657220636f6e747261637400000000000060448201526064015b60405180910390fd5b6040805180820190915281815242602080830191909152606980546001810182555f919091528251805160029092027f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe399914308019261061f9284929091019061144e565b5060208201518160010155505050565b61063761121f565b60655485146106a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206261746368496e646578000000000000000000000000000060448201526064016105b5565b606654831461070d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617463685374617274426c6f636b00000000000000000060448201526064016105b5565b604080516080808201835273ffffffffffffffffffffffffffffffffffffffff878116808452602080850189815285870189815260608088018a81525f8f8152606786528a9020985189547fffffffffffffffffffffffff000000000000000000000000000000000000000016971696909617885591516001880155516002870155925160039095019490945584518a8152938401529282018690529181018490529081018290527f516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f9060a00160405180910390a160658054905f6107f1836116c3565b9091555061080290508260016116fa565b6066555050505050565b6069818154811061081b575f80fd5b5f918252602090912060016002909202010154905081565b61083b61121f565b6108445f6112a0565b565b5f54610100900460ff161580801561086457505f54600160ff909116105b8061087d5750303b15801561087d57505f5460ff166001145b610909576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105b5565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610965575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61096d611316565b80156109cf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b5f805f8060698054905011610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c69642073657175656e63657200000000000000000000000000000060448201526064016105b5565b606980545f9190610a5690600190611713565b81548110610a6657610a66611726565b905f5260205f2090600202016001015490505f606a80549050118015610abc5750606a8054829190610a9a90600190611713565b81548110610aaa57610aaa611726565b905f5260205f20906002020160010154115b15610af457606a8054610ad190600190611713565b81548110610ae157610ae1611726565b905f5260205f2090600202016001015490505b606980545f9190610b0790600190611713565b81548110610b1757610b17611726565b5f9182526020918290206002909102018054604080518285028101850190915281815292830182828015610b7f57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610b54575b505050505090505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c149190611753565b82519091505f82610c258642611713565b610c2f9190611797565b90505f610c3c83836117aa565b90505f84610c4a8842611713565b610c5491906117aa565b610c5e9042611713565b9050858281518110610c7257610c72611726565b6020026020010151818683610c8791906116fa565b99509950995050505050505050909192565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f6e6c7920676f7620636f6e747261637400000000000000000000000000000060448201526064016105b5565b604080518082019091529081524260208201908152606a80546001810182555f9190915291517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a51600290930292830155517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5290910155565b5f805f60696001606980549050610dc79190611713565b81548110610dd757610dd7611726565b905f5260205f2090600202016001015490505f606a80549050118015610e2d5750606a8054829190610e0b90600190611713565b81548110610e1b57610e1b611726565b905f5260205f20906002020160010154115b15610e6557606a8054610e4290600190611713565b81548110610e5257610e52611726565b905f5260205f2090600202016001015490505b606980545f9190610e7890600190611713565b81548110610e8857610e88611726565b5f9182526020918290206002909102018054604080518285028101850190915281815292830182828015610ef057602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ec5575b505050505090505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f61573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f859190611753565b82519091505f805b82811015610ffc57848181518110610fa757610fa7611726565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603610fea5760019150610ffc565b80610ff4816116c3565b915050610f8d565b81611063576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c6964207375626d697474657200000000000000000000000000000060448201526064016105b5565b5f61106e85836117bd565b61107890886116fa565b90505f61108585876117bd565b9050814211156110ca575f8161109b8442611713565b6110a59190611797565b6110b09060016116fa565b90506110bc82826117bd565b6110c690846116fa565b9250505b816110d587826116fa565b995099505050505050505050915091565b61111d60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b505f908152606860209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600201549181019190915290565b61117361121f565b73ffffffffffffffffffffffffffffffffffffffff8116611216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105b5565b6109cf816112a0565b60335473ffffffffffffffffffffffffffffffffffffffff163314610844576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105b5565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166113ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b5565b6108445f54610100900460ff16611445576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b5565b610844336112a0565b828054828255905f5260205f209081019282156114c6579160200282015b828111156114c657825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019061146c565b506114d29291506114d6565b5090565b5b808211156114d2575f81556001016114d7565b5f602082840312156114fa575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b803573ffffffffffffffffffffffffffffffffffffffff81168114611551575f80fd5b919050565b5f6020808385031215611567575f80fd5b823567ffffffffffffffff8082111561157e575f80fd5b818501915085601f830112611591575f80fd5b8135818111156115a3576115a3611501565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156115e6576115e6611501565b604052918252848201925083810185019188831115611603575f80fd5b938501935b82851015611628576116198561152e565b84529385019392850192611608565b98975050505050505050565b5f805f805f60a08688031215611648575f80fd5b853594506116586020870161152e565b94979496505050506040830135926060810135926080909101359150565b5f60208284031215611686575f80fd5b61168f8261152e565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116f3576116f3611696565b5060010190565b8082018082111561170d5761170d611696565b92915050565b8181038181111561170d5761170d611696565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215611763575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f826117a5576117a561176a565b500490565b5f826117b8576117b861176a565b500690565b808202811582820484141761170d5761170d61169656fea264697066735822122048d8b5394b08e336261ad0fbe1305453dd97be5b73810108ff0390f863b58d5c64736f6c63430008180033",
}

// SubmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use SubmitterMetaData.ABI instead.
var SubmitterABI = SubmitterMetaData.ABI

// SubmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubmitterMetaData.Bin instead.
var SubmitterBin = SubmitterMetaData.Bin

// DeploySubmitter deploys a new Ethereum contract, binding an instance of Submitter to it.
func DeploySubmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Submitter, error) {
	parsed, err := SubmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// Submitter is an auto generated Go binding around an Ethereum contract.
type Submitter struct {
	SubmitterCaller     // Read-only binding to the contract
	SubmitterTransactor // Write-only binding to the contract
	SubmitterFilterer   // Log filterer for contract events
}

// SubmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmitterSession struct {
	Contract     *Submitter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmitterCallerSession struct {
	Contract *SubmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SubmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmitterTransactorSession struct {
	Contract     *SubmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SubmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmitterRaw struct {
	Contract *Submitter // Generic contract binding to access the raw methods on
}

// SubmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmitterCallerRaw struct {
	Contract *SubmitterCaller // Generic read-only contract binding to access the raw methods on
}

// SubmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmitterTransactorRaw struct {
	Contract *SubmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmitter creates a new instance of Submitter, bound to a specific deployed contract.
func NewSubmitter(address common.Address, backend bind.ContractBackend) (*Submitter, error) {
	contract, err := bindSubmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// NewSubmitterCaller creates a new read-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterCaller(address common.Address, caller bind.ContractCaller) (*SubmitterCaller, error) {
	contract, err := bindSubmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterCaller{contract: contract}, nil
}

// NewSubmitterTransactor creates a new write-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmitterTransactor, error) {
	contract, err := bindSubmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterTransactor{contract: contract}, nil
}

// NewSubmitterFilterer creates a new log filterer instance of Submitter, bound to a specific deployed contract.
func NewSubmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmitterFilterer, error) {
	contract, err := bindSubmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmitterFilterer{contract: contract}, nil
}

// bindSubmitter binds a generic wrapper to an already deployed contract.
func bindSubmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.SubmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transact(opts, method, params...)
}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterCaller) L2GOVCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "L2_GOV_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterSession) L2GOVCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2GOVCONTRACT(&_Submitter.CallOpts)
}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterCallerSession) L2GOVCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2GOVCONTRACT(&_Submitter.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterCaller) L2SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "L2_SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2SEQUENCERCONTRACT(&_Submitter.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterCallerSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2SEQUENCERCONTRACT(&_Submitter.CallOpts)
}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterCaller) ConfirmedBatchs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "confirmedBatchs", arg0)

	outstruct := new(struct {
		Submitter  common.Address
		StartBlock *big.Int
		EndBlock   *big.Int
		RollupTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RollupTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterSession) ConfirmedBatchs(arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Submitter.Contract.ConfirmedBatchs(&_Submitter.CallOpts, arg0)
}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterCallerSession) ConfirmedBatchs(arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Submitter.Contract.ConfirmedBatchs(&_Submitter.CallOpts, arg0)
}

// EpochHistory is a free data retrieval call binding the contract method 0x04b12385.
//
// Solidity: function epochHistory(uint256 ) view returns(uint256 epoch, uint256 timestamp)
func (_Submitter *SubmitterCaller) EpochHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "epochHistory", arg0)

	outstruct := new(struct {
		Epoch     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochHistory is a free data retrieval call binding the contract method 0x04b12385.
//
// Solidity: function epochHistory(uint256 ) view returns(uint256 epoch, uint256 timestamp)
func (_Submitter *SubmitterSession) EpochHistory(arg0 *big.Int) (struct {
	Epoch     *big.Int
	Timestamp *big.Int
}, error) {
	return _Submitter.Contract.EpochHistory(&_Submitter.CallOpts, arg0)
}

// EpochHistory is a free data retrieval call binding the contract method 0x04b12385.
//
// Solidity: function epochHistory(uint256 ) view returns(uint256 epoch, uint256 timestamp)
func (_Submitter *SubmitterCallerSession) EpochHistory(arg0 *big.Int) (struct {
	Epoch     *big.Int
	Timestamp *big.Int
}, error) {
	return _Submitter.Contract.EpochHistory(&_Submitter.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "epochs", arg0)

	outstruct := new(struct {
		Submitter common.Address
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterSession) Epochs(arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Submitter.Contract.Epochs(&_Submitter.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterCallerSession) Epochs(arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Submitter.Contract.Epochs(&_Submitter.CallOpts, arg0)
}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterCaller) GetConfirmedBatch(opts *bind.CallOpts, batchIndex *big.Int) (TypesBatchInfo, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getConfirmedBatch", batchIndex)

	if err != nil {
		return *new(TypesBatchInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBatchInfo)).(*TypesBatchInfo)

	return out0, err

}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterSession) GetConfirmedBatch(batchIndex *big.Int) (TypesBatchInfo, error) {
	return _Submitter.Contract.GetConfirmedBatch(&_Submitter.CallOpts, batchIndex)
}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterCallerSession) GetConfirmedBatch(batchIndex *big.Int) (TypesBatchInfo, error) {
	return _Submitter.Contract.GetConfirmedBatch(&_Submitter.CallOpts, batchIndex)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterCaller) GetEpoch(opts *bind.CallOpts, epochIndex *big.Int) (TypesEpochInfo, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getEpoch", epochIndex)

	if err != nil {
		return *new(TypesEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesEpochInfo)).(*TypesEpochInfo)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterSession) GetEpoch(epochIndex *big.Int) (TypesEpochInfo, error) {
	return _Submitter.Contract.GetEpoch(&_Submitter.CallOpts, epochIndex)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterCallerSession) GetEpoch(epochIndex *big.Int) (TypesEpochInfo, error) {
	return _Submitter.Contract.GetEpoch(&_Submitter.CallOpts, epochIndex)
}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCaller) GetNextSubmitter(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getNextSubmitter")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterSession) GetNextSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetNextSubmitter(&_Submitter.CallOpts)
}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCallerSession) GetNextSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetNextSubmitter(&_Submitter.CallOpts)
}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterCaller) GetTurn(opts *bind.CallOpts, submitter common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getTurn", submitter)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterSession) GetTurn(submitter common.Address) (*big.Int, *big.Int, error) {
	return _Submitter.Contract.GetTurn(&_Submitter.CallOpts, submitter)
}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterCallerSession) GetTurn(submitter common.Address) (*big.Int, *big.Int, error) {
	return _Submitter.Contract.GetTurn(&_Submitter.CallOpts, submitter)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterCaller) NextBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterSession) NextBatchIndex() (*big.Int, error) {
	return _Submitter.Contract.NextBatchIndex(&_Submitter.CallOpts)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextBatchIndex() (*big.Int, error) {
	return _Submitter.Contract.NextBatchIndex(&_Submitter.CallOpts)
}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterCaller) NextBatchStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextBatchStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterSession) NextBatchStartBlock() (*big.Int, error) {
	return _Submitter.Contract.NextBatchStartBlock(&_Submitter.CallOpts)
}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextBatchStartBlock() (*big.Int, error) {
	return _Submitter.Contract.NextBatchStartBlock(&_Submitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterSession) Owner() (common.Address, error) {
	return _Submitter.Contract.Owner(&_Submitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterCallerSession) Owner() (common.Address, error) {
	return _Submitter.Contract.Owner(&_Submitter.CallOpts)
}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_Submitter *SubmitterCaller) SequencerHistory(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "sequencerHistory", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_Submitter *SubmitterSession) SequencerHistory(arg0 *big.Int) (*big.Int, error) {
	return _Submitter.Contract.SequencerHistory(&_Submitter.CallOpts, arg0)
}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_Submitter *SubmitterCallerSession) SequencerHistory(arg0 *big.Int) (*big.Int, error) {
	return _Submitter.Contract.SequencerHistory(&_Submitter.CallOpts, arg0)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterTransactor) AckRollup(opts *bind.TransactOpts, batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "ackRollup", batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterSession) AckRollup(batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.AckRollup(&_Submitter.TransactOpts, batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterTransactorSession) AckRollup(batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.AckRollup(&_Submitter.TransactOpts, batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterTransactor) EpochUpdated(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "epochUpdated", epoch)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterSession) EpochUpdated(epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.EpochUpdated(&_Submitter.TransactOpts, epoch)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterTransactorSession) EpochUpdated(epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.EpochUpdated(&_Submitter.TransactOpts, epoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Submitter *SubmitterTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Submitter *SubmitterSession) Initialize() (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Submitter *SubmitterTransactorSession) Initialize() (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Submitter.Contract.RenounceOwnership(&_Submitter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Submitter.Contract.RenounceOwnership(&_Submitter.TransactOpts)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterTransactor) SequencersUpdated(opts *bind.TransactOpts, sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "sequencersUpdated", sequencers)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterSession) SequencersUpdated(sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.SequencersUpdated(&_Submitter.TransactOpts, sequencers)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterTransactorSession) SequencersUpdated(sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.SequencersUpdated(&_Submitter.TransactOpts, sequencers)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.TransferOwnership(&_Submitter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.TransferOwnership(&_Submitter.TransactOpts, newOwner)
}

// SubmitterACKRollupIterator is returned from FilterACKRollup and is used to iterate over the raw logs and unpacked data for ACKRollup events raised by the Submitter contract.
type SubmitterACKRollupIterator struct {
	Event *SubmitterACKRollup // Event containing the contract specifics and raw log

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
func (it *SubmitterACKRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterACKRollup)
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
		it.Event = new(SubmitterACKRollup)
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
func (it *SubmitterACKRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterACKRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterACKRollup represents a ACKRollup event raised by the Submitter contract.
type SubmitterACKRollup struct {
	BatchIndex      *big.Int
	Submitter       common.Address
	BatchStartBlock *big.Int
	BatchEndBlock   *big.Int
	RollupTime      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterACKRollup is a free log retrieval operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) FilterACKRollup(opts *bind.FilterOpts) (*SubmitterACKRollupIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "ACKRollup")
	if err != nil {
		return nil, err
	}
	return &SubmitterACKRollupIterator{contract: _Submitter.contract, event: "ACKRollup", logs: logs, sub: sub}, nil
}

// WatchACKRollup is a free log subscription operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) WatchACKRollup(opts *bind.WatchOpts, sink chan<- *SubmitterACKRollup) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "ACKRollup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterACKRollup)
				if err := _Submitter.contract.UnpackLog(event, "ACKRollup", log); err != nil {
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

// ParseACKRollup is a log parse operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) ParseACKRollup(log types.Log) (*SubmitterACKRollup, error) {
	event := new(SubmitterACKRollup)
	if err := _Submitter.contract.UnpackLog(event, "ACKRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Submitter contract.
type SubmitterInitializedIterator struct {
	Event *SubmitterInitialized // Event containing the contract specifics and raw log

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
func (it *SubmitterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterInitialized)
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
		it.Event = new(SubmitterInitialized)
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
func (it *SubmitterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterInitialized represents a Initialized event raised by the Submitter contract.
type SubmitterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubmitterInitializedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubmitterInitializedIterator{contract: _Submitter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubmitterInitialized) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterInitialized)
				if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Submitter *SubmitterFilterer) ParseInitialized(log types.Log) (*SubmitterInitialized, error) {
	event := new(SubmitterInitialized)
	if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Submitter contract.
type SubmitterOwnershipTransferredIterator struct {
	Event *SubmitterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubmitterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterOwnershipTransferred)
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
		it.Event = new(SubmitterOwnershipTransferred)
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
func (it *SubmitterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterOwnershipTransferred represents a OwnershipTransferred event raised by the Submitter contract.
type SubmitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Submitter *SubmitterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubmitterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterOwnershipTransferredIterator{contract: _Submitter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Submitter *SubmitterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubmitterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterOwnershipTransferred)
				if err := _Submitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Submitter *SubmitterFilterer) ParseOwnershipTransferred(log types.Log) (*SubmitterOwnershipTransferred, error) {
	event := new(SubmitterOwnershipTransferred)
	if err := _Submitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
