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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ACKRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_GOV_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ackRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmedBatchs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"epochUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"getConfirmedBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BatchInfo\",\"name\":\"batchInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EpochInfo\",\"name\":\"epochInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"getTurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b5073530000000000000000000000000000000000000360805273530000000000000000000000000000000000000460a05260805160a0516118c76100a05f395f818161011e0152818161095c01528181610c240152610fdc01525f8181610259015281816106a80152818161077c015281816108b601528181610d2801528181610dfc0152610f3601526118c75ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c80637e4fa700116100ad578063a5af40d11161007d578063c6b61e4c11610063578063c6b61e4c14610323578063e8e3992514610368578063f2fde38b146103ea575f80fd5b8063a5af40d1146102c5578063bc0bc6ba146102d8575f80fd5b80637e4fa700146102835780638129fc1c1461028c5780638da5cb5b14610294578063965fbb94146102b2575f80fd5b80635c14c314116100e85780635c14c314146102035780636b511e821461021a5780636cb2370714610254578063715018a61461027b575f80fd5b8063047d0b6e1461011957806304b123851461016a578063151232581461019257806322caba24146101ee575b5f80fd5b6101407f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61017d610178366004611531565b6103fd565b60408051928352602083019190915201610161565b6101a56101a0366004611531565b610429565b6040516101619190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102016101fc366004611569565b6104c0565b005b61020c60665481565b604051908152602001610161565b6102226106a2565b6040805173ffffffffffffffffffffffffffffffffffffffff9094168452602084019290925290820152606001610161565b6101407f000000000000000000000000000000000000000000000000000000000000000081565b610201610a6d565b61020c60655481565b610201610a80565b60335473ffffffffffffffffffffffffffffffffffffffff16610140565b6102016102c0366004611531565b610c0c565b61017d6102d33660046115ad565b610d23565b6102eb6102e6366004611531565b6111c9565b60408051825173ffffffffffffffffffffffffffffffffffffffff168152602080840151908201529181015190820152606001610161565b610222610331366004611531565b60686020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116919083565b6103b3610376366004611531565b60676020525f9081526040902080546001820154600283015460039093015473ffffffffffffffffffffffffffffffffffffffff90921692909184565b6040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152608001610161565b6102016103f83660046115ad565b61124e565b6069818154811061040c575f80fd5b5f9182526020909120600290910201805460019091015490915082565b61046660405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81525090565b505f908152606760209081526040918290208251608081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281015492820192909252600390910154606082015290565b6104c8611302565b6065548514610538576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206261746368496e646578000000000000000000000000000060448201526064015b60405180910390fd5b60665483146105a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617463685374617274426c6f636b000000000000000000604482015260640161052f565b604080516080808201835273ffffffffffffffffffffffffffffffffffffffff878116808452602080850189815285870189815260608088018a81525f8f8152606786528a9020985189547fffffffffffffffffffffffff000000000000000000000000000000000000000016971696909617885591516001880155516002870155925160039095019490945584518a8152938401529282018690529181018490529081018290527f516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f9060a00160405180910390a160658054905f610687836115fc565b909155506106989050826001611633565b6066555050505050565b5f805f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639d888e866040518163ffffffff1660e01b8152600401602060405180830381865afa15801561070f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610733919061164c565b6040517ff6f207ce000000000000000000000000000000000000000000000000000000008152600481018290529091505f9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063f6f207ce906024015f60405180830381865afa1580156107c0573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526108059190810190611708565b602001516069549091501580159061084d57506069805482919061082b906001906117e7565b8154811061083b5761083b6117fa565b905f5260205f20906002020160010154115b156108855760698054610862906001906117e7565b81548110610872576108726117fa565b905f5260205f2090600202016001015490505b6040517ff6f207ce000000000000000000000000000000000000000000000000000000008152600481018390525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063f6f207ce906024015f60405180830381865afa15801561090f573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526109549190810190611708565b5f015190505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e7919061164c565b82519091505f826109f886426117e7565b610a029190611854565b90505f610a0f8383611867565b90505f84610a1d88426117e7565b610a279190611867565b610a3190426117e7565b9050858281518110610a4557610a456117fa565b6020026020010151818683610a5a9190611633565b9a509a509a505050505050505050909192565b610a75611302565b610a7e5f611383565b565b5f54610100900460ff1615808015610a9e57505f54600160ff909116105b80610ab75750303b158015610ab757505f5460ff166001145b610b43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161052f565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610b9f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610ba76113f9565b8015610c09575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f6e6c7920676f7620636f6e7472616374000000000000000000000000000000604482015260640161052f565b604080518082019091529081524260208201908152606980546001810182555f9190915291517f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe399914308600290930292830155517f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe39991430990910155565b5f805f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639d888e866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db3919061164c565b6040517ff6f207ce000000000000000000000000000000000000000000000000000000008152600481018290529091505f9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063f6f207ce906024015f60405180830381865afa158015610e40573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610e859190810190611708565b6020015160695490915015801590610ecd575060698054829190610eab906001906117e7565b81548110610ebb57610ebb6117fa565b905f5260205f20906002020160010154115b15610f055760698054610ee2906001906117e7565b81548110610ef257610ef26117fa565b905f5260205f2090600202016001015490505b6040517ff6f207ce000000000000000000000000000000000000000000000000000000008152600481018390525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063f6f207ce906024015f60405180830381865afa158015610f8f573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610fd49190810190611708565b5f015190505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015611043573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611067919061164c565b82519091505f805b828110156110de57848181518110611089576110896117fa565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff16036110cc57600191506110de565b806110d6816115fc565b91505061106f565b81611145576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c6964207375626d6974746572000000000000000000000000000000604482015260640161052f565b5f611150858361187a565b61115a9088611633565b90505f611167858761187a565b9050814211156111ac575f8161117d84426117e7565b6111879190611854565b611192906001611633565b905061119e828261187a565b6111a89084611633565b9250505b816111b78782611633565b9a509a50505050505050505050915091565b61120060405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b505f908152606860209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600201549181019190915290565b611256611302565b73ffffffffffffffffffffffffffffffffffffffff81166112f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161052f565b610c0981611383565b60335473ffffffffffffffffffffffffffffffffffffffff163314610a7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161052f565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff1661148f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161052f565b610a7e5f54610100900460ff16611528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161052f565b610a7e33611383565b5f60208284031215611541575f80fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c09575f80fd5b5f805f805f60a0868803121561157d575f80fd5b85359450602086013561158f81611548565b94979496505050506040830135926060810135926080909101359150565b5f602082840312156115bd575f80fd5b81356115c881611548565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361162c5761162c6115cf565b5060010190565b80820180821115611646576116466115cf565b92915050565b5f6020828403121561165c575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156116b3576116b3611663565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561170057611700611663565b604052919050565b5f6020808385031215611719575f80fd5b825167ffffffffffffffff80821115611730575f80fd5b9084019060408287031215611743575f80fd5b61174b611690565b825182811115611759575f80fd5b8301601f81018813611769575f80fd5b80518381111561177b5761177b611663565b8060051b935061178c8685016116b9565b818152938201860193868101908a8611156117a5575f80fd5b928701925b858410156117cf57835192506117bf83611548565b82825292870192908701906117aa565b84525050509183015192820192909252949350505050565b81810381811115611646576116466115cf565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f8261186257611862611827565b500490565b5f8261187557611875611827565b500690565b8082028115828204841417611646576116466115cf56fea26469706673582212201e67b76b7b412ae4252a11d3cc1b5f53e564cdb85c28de4b8616cfa098dae49464736f6c63430008180033",
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

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCaller) GetCurrentSubmitter(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getCurrentSubmitter")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterSession) GetCurrentSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetCurrentSubmitter(&_Submitter.CallOpts)
}

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCallerSession) GetCurrentSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetCurrentSubmitter(&_Submitter.CallOpts)
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
