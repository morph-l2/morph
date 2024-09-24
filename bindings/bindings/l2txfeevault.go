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

// L2TxFeeVaultMetaData contains all meta data concerning the L2TxFeeVault contract.
var L2TxFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minWithdrawalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMessenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMessenger\",\"type\":\"address\"}],\"name\":\"UpdateMessenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMinWithdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"UpdateMinWithdrawAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateReceiveAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"UpdateRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateTransferAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"isReceiveAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiveAddr\",\"type\":\"address\"}],\"name\":\"receiveAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferAddr\",\"type\":\"address\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMessenger\",\"type\":\"address\"}],\"name\":\"updateMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateReceiveAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"}],\"name\":\"updateRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateTransferAllowedStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001687380380620016878339810160408190526200003391620000d3565b6200003e8362000068565b600155600380546001600160a01b0319166001600160a01b03929092169190911790555062000111565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000ce575f80fd5b919050565b5f805f60608486031215620000e6575f80fd5b620000f184620000b7565b92506200010160208501620000b7565b9150604084015190509250925092565b611568806200011f5f395ff3fe608060405260043610610140575f3560e01c806384411d65116100bb578063da13f9a211610071578063f2fde38b11610057578063f2fde38b146103c4578063feec756c146103e3578063ff4f354614610402575f80fd5b8063da13f9a214610377578063eff1337c14610396575f80fd5b80639e7adc79116100a15780639e7adc79146102f5578063a03fa7e314610314578063cc198d7414610333575f80fd5b806384411d65146102b55780638da5cb5b146102ca575f80fd5b80633ccfd60b1161011057806366d003ac116100f657806366d003ac14610256578063708125ad14610282578063715018a6146102a1575f80fd5b80633ccfd60b1461021f578063457e1a4914610233575f80fd5b8063151eeb551461014b5780632ccb1b301461018e5780632e1a7d4d146101af5780633cb747bf146101ce575f80fd5b3661014757005b5f80fd5b348015610156575f80fd5b5061017961016536600461130d565b60056020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b348015610199575f80fd5b506101ad6101a836600461132d565b610421565b005b3480156101ba575f80fd5b506101ad6101c9366004611355565b610750565b3480156101d9575f80fd5b506002546101fa9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610185565b34801561022a575f80fd5b506101ad610a55565b34801561023e575f80fd5b5061024860015481565b604051908152602001610185565b348015610261575f80fd5b506003546101fa9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561028d575f80fd5b506101ad61029c3660046113a8565b610ac8565b3480156102ac575f80fd5b506101ad610ce2565b3480156102c0575f80fd5b5061024860045481565b3480156102d5575f80fd5b505f546101fa9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610300575f80fd5b506101ad61030f36600461130d565b610d53565b34801561031f575f80fd5b506101ad61032e36600461130d565b610e2f565b34801561033e575f80fd5b5061017961034d36600461130d565b73ffffffffffffffffffffffffffffffffffffffff165f9081526006602052604090205460ff1690565b348015610382575f80fd5b506101ad6103913660046113a8565b610ebd565b3480156103a1575f80fd5b506101796103b036600461130d565b60066020525f908152604090205460ff1681565b3480156103cf575f80fd5b506101ad6103de36600461130d565b611018565b3480156103ee575f80fd5b506101ad6103fd36600461130d565b6110ea565b34801561040d575f80fd5b506101ad61041c366004611355565b6111c6565b335f9081526005602052604090205460ff168061045457505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6104a55760405162461bcd60e51b815260206004820152601f60248201527f4665655661756c743a2063616c6c6572206973206e6f7420616c6c6f7765640060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661052e5760405162461bcd60e51b815260206004820152603060248201527f4665655661756c743a20726563697069656e7420616464726573732063616e6e60448201527f6f74206265206164647265737328302900000000000000000000000000000000606482015260840161049c565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090205460ff166105c85760405162461bcd60e51b815260206004820152602760248201527f4665655661756c743a20726563697069656e742061646472657373206e6f742060448201527f616c6c6f77656400000000000000000000000000000000000000000000000000606482015260840161049c565b478082111561063f5760405162461bcd60e51b815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f207472616e7366657200000000000000000000000000000000000000000000606482015260840161049c565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517f0a429aba3d89849a2db0153e4534d95c46a1d83c8109d73893f55ebc44010ff49181900360600190a15f8373ffffffffffffffffffffffffffffffffffffffff16836040515f6040518083038185875af1925050503d805f81146106f4576040519150601f19603f3d011682016040523d82523d5f602084013e6106f9565b606091505b505090508061074a5760405162461bcd60e51b815260206004820152601d60248201527f4665655661756c743a20455448207472616e73666572206661696c6564000000604482015260640161049c565b50505050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146107b65760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b60035473ffffffffffffffffffffffffffffffffffffffff166108415760405162461bcd60e51b815260206004820152603060248201527f4665655661756c743a20726563697069656e7420616464726573732063616e6e60448201527f6f74206265206164647265737328302900000000000000000000000000000000606482015260840161049c565b6001548110156108df5760405162461bcd60e51b815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40161049c565b47808211156109565760405162461bcd60e51b815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f20776974686472617700000000000000000000000000000000000000000000606482015260840161049c565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1600254600354604080516020810182525f80825291517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485169463b2267a7b948894610a23949190921692859290600401611498565b5f604051808303818588803b158015610a3a575f80fd5b505af1158015610a4c573d5f803e3d5ffd5b50505050505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610abb5760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b47610ac581610750565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610b2e5760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b5f5b8251811015610cdd575f73ffffffffffffffffffffffffffffffffffffffff16838281518110610b6257610b6261152e565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603610bf35760405162461bcd60e51b815260206004820152602660248201527f4665655661756c743a20616464726573732063616e6e6f74206265206164647260448201527f6573732830290000000000000000000000000000000000000000000000000000606482015260840161049c565b8160065f858481518110610c0957610c0961152e565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507fbd84b6b3e1d029fd61a717f6020a8f35b358486e0971c00d4e64b516503f85e7838281518110610c9357610c9361152e565b602002602001015183604051610ccd92919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a1600101610b30565b505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610d485760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b610d515f611271565b565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610db95760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612905f90a35050565b335f9081526005602052604090205460ff1680610e6257505f5473ffffffffffffffffffffffffffffffffffffffff1633145b610eae5760405162461bcd60e51b815260206004820152601f60248201527f4665655661756c743a2063616c6c6572206973206e6f7420616c6c6f77656400604482015260640161049c565b47610eb98282610421565b5050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610f235760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b5f5b8251811015610cdd578160055f858481518110610f4457610f4461152e565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507fbb5d3e442e800faa1089a4f57bae4f36808d3cf15d051033d78a72147782f24c838281518110610fce57610fce61152e565b60200260200101518360405161100892919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a1600101610f25565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461107e5760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b73ffffffffffffffffffffffffffffffffffffffff81166110e15760405162461bcd60e51b815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f2061646472657373000000604482015260640161049c565b610ac581611271565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146111505760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461122c5760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161049c565b600180549082905560408051828152602081018490527f0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965910160405180910390a15050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611308575f80fd5b919050565b5f6020828403121561131d575f80fd5b611326826112e5565b9392505050565b5f806040838503121561133e575f80fd5b611347836112e5565b946020939093013593505050565b5f60208284031215611365575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b80358015158114611308575f80fd5b5f80604083850312156113b9575f80fd5b823567ffffffffffffffff808211156113d0575f80fd5b818501915085601f8301126113e3575f80fd5b81356020828211156113f7576113f761136c565b8160051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110868211171561143a5761143a61136c565b604052928352818301935084810182019289841115611457575f80fd5b948201945b8386101561147c5761146d866112e5565b8552948201949382019361145c565b965061148b9050878201611399565b9450505050509250929050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b818110156114e75786810183015185820160a0015282016114cb565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a",
}

// L2TxFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use L2TxFeeVaultMetaData.ABI instead.
var L2TxFeeVaultABI = L2TxFeeVaultMetaData.ABI

// L2TxFeeVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2TxFeeVaultMetaData.Bin instead.
var L2TxFeeVaultBin = L2TxFeeVaultMetaData.Bin

// DeployL2TxFeeVault deploys a new Ethereum contract, binding an instance of L2TxFeeVault to it.
func DeployL2TxFeeVault(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _recipient common.Address, _minWithdrawalAmount *big.Int) (common.Address, *types.Transaction, *L2TxFeeVault, error) {
	parsed, err := L2TxFeeVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2TxFeeVaultBin), backend, _owner, _recipient, _minWithdrawalAmount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2TxFeeVault{L2TxFeeVaultCaller: L2TxFeeVaultCaller{contract: contract}, L2TxFeeVaultTransactor: L2TxFeeVaultTransactor{contract: contract}, L2TxFeeVaultFilterer: L2TxFeeVaultFilterer{contract: contract}}, nil
}

// L2TxFeeVault is an auto generated Go binding around an Ethereum contract.
type L2TxFeeVault struct {
	L2TxFeeVaultCaller     // Read-only binding to the contract
	L2TxFeeVaultTransactor // Write-only binding to the contract
	L2TxFeeVaultFilterer   // Log filterer for contract events
}

// L2TxFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2TxFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2TxFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2TxFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2TxFeeVaultSession struct {
	Contract     *L2TxFeeVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2TxFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2TxFeeVaultCallerSession struct {
	Contract *L2TxFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L2TxFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TxFeeVaultTransactorSession struct {
	Contract     *L2TxFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2TxFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2TxFeeVaultRaw struct {
	Contract *L2TxFeeVault // Generic contract binding to access the raw methods on
}

// L2TxFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2TxFeeVaultCallerRaw struct {
	Contract *L2TxFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// L2TxFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TxFeeVaultTransactorRaw struct {
	Contract *L2TxFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2TxFeeVault creates a new instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVault(address common.Address, backend bind.ContractBackend) (*L2TxFeeVault, error) {
	contract, err := bindL2TxFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVault{L2TxFeeVaultCaller: L2TxFeeVaultCaller{contract: contract}, L2TxFeeVaultTransactor: L2TxFeeVaultTransactor{contract: contract}, L2TxFeeVaultFilterer: L2TxFeeVaultFilterer{contract: contract}}, nil
}

// NewL2TxFeeVaultCaller creates a new read-only instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*L2TxFeeVaultCaller, error) {
	contract, err := bindL2TxFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultCaller{contract: contract}, nil
}

// NewL2TxFeeVaultTransactor creates a new write-only instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*L2TxFeeVaultTransactor, error) {
	contract, err := bindL2TxFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultTransactor{contract: contract}, nil
}

// NewL2TxFeeVaultFilterer creates a new log filterer instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*L2TxFeeVaultFilterer, error) {
	contract, err := bindL2TxFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultFilterer{contract: contract}, nil
}

// bindL2TxFeeVault binds a generic wrapper to an already deployed contract.
func bindL2TxFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2TxFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TxFeeVault.Contract.L2TxFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.L2TxFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.L2TxFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TxFeeVault *L2TxFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TxFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TxFeeVault *L2TxFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TxFeeVault *L2TxFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.contract.Transact(opts, method, params...)
}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xcc198d74.
//
// Solidity: function isReceiveAllowed(address _to) view returns(bool)
func (_L2TxFeeVault *L2TxFeeVaultCaller) IsReceiveAllowed(opts *bind.CallOpts, _to common.Address) (bool, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "isReceiveAllowed", _to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xcc198d74.
//
// Solidity: function isReceiveAllowed(address _to) view returns(bool)
func (_L2TxFeeVault *L2TxFeeVaultSession) IsReceiveAllowed(_to common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.IsReceiveAllowed(&_L2TxFeeVault.CallOpts, _to)
}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xcc198d74.
//
// Solidity: function isReceiveAllowed(address _to) view returns(bool)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) IsReceiveAllowed(_to common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.IsReceiveAllowed(&_L2TxFeeVault.CallOpts, _to)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Messenger() (common.Address, error) {
	return _L2TxFeeVault.Contract.Messenger(&_L2TxFeeVault.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Messenger() (common.Address, error) {
	return _L2TxFeeVault.Contract.Messenger(&_L2TxFeeVault.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCaller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultSession) MinWithdrawAmount() (*big.Int, error) {
	return _L2TxFeeVault.Contract.MinWithdrawAmount(&_L2TxFeeVault.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _L2TxFeeVault.Contract.MinWithdrawAmount(&_L2TxFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Owner() (common.Address, error) {
	return _L2TxFeeVault.Contract.Owner(&_L2TxFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Owner() (common.Address, error) {
	return _L2TxFeeVault.Contract.Owner(&_L2TxFeeVault.CallOpts)
}

// ReceiveAllowed is a free data retrieval call binding the contract method 0xeff1337c.
//
// Solidity: function receiveAllowed(address receiveAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultCaller) ReceiveAllowed(opts *bind.CallOpts, receiveAddr common.Address) (bool, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "receiveAllowed", receiveAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReceiveAllowed is a free data retrieval call binding the contract method 0xeff1337c.
//
// Solidity: function receiveAllowed(address receiveAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultSession) ReceiveAllowed(receiveAddr common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.ReceiveAllowed(&_L2TxFeeVault.CallOpts, receiveAddr)
}

// ReceiveAllowed is a free data retrieval call binding the contract method 0xeff1337c.
//
// Solidity: function receiveAllowed(address receiveAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) ReceiveAllowed(receiveAddr common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.ReceiveAllowed(&_L2TxFeeVault.CallOpts, receiveAddr)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Recipient() (common.Address, error) {
	return _L2TxFeeVault.Contract.Recipient(&_L2TxFeeVault.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Recipient() (common.Address, error) {
	return _L2TxFeeVault.Contract.Recipient(&_L2TxFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultSession) TotalProcessed() (*big.Int, error) {
	return _L2TxFeeVault.Contract.TotalProcessed(&_L2TxFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _L2TxFeeVault.Contract.TotalProcessed(&_L2TxFeeVault.CallOpts)
}

// TransferAllowed is a free data retrieval call binding the contract method 0x151eeb55.
//
// Solidity: function transferAllowed(address transferAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultCaller) TransferAllowed(opts *bind.CallOpts, transferAddr common.Address) (bool, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "transferAllowed", transferAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferAllowed is a free data retrieval call binding the contract method 0x151eeb55.
//
// Solidity: function transferAllowed(address transferAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultSession) TransferAllowed(transferAddr common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.TransferAllowed(&_L2TxFeeVault.CallOpts, transferAddr)
}

// TransferAllowed is a free data retrieval call binding the contract method 0x151eeb55.
//
// Solidity: function transferAllowed(address transferAddr) view returns(bool allowed)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) TransferAllowed(transferAddr common.Address) (bool, error) {
	return _L2TxFeeVault.Contract.TransferAllowed(&_L2TxFeeVault.CallOpts, transferAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.RenounceOwnership(&_L2TxFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.RenounceOwnership(&_L2TxFeeVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferOwnership(&_L2TxFeeVault.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferOwnership(&_L2TxFeeVault.TransactOpts, _newOwner)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address _to, uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) TransferTo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "transferTo", _to, _value)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address _to, uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) TransferTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferTo(&_L2TxFeeVault.TransactOpts, _to, _value)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address _to, uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) TransferTo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferTo(&_L2TxFeeVault.TransactOpts, _to, _value)
}

// TransferTo0 is a paid mutator transaction binding the contract method 0xa03fa7e3.
//
// Solidity: function transferTo(address _to) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) TransferTo0(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "transferTo0", _to)
}

// TransferTo0 is a paid mutator transaction binding the contract method 0xa03fa7e3.
//
// Solidity: function transferTo(address _to) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) TransferTo0(_to common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferTo0(&_L2TxFeeVault.TransactOpts, _to)
}

// TransferTo0 is a paid mutator transaction binding the contract method 0xa03fa7e3.
//
// Solidity: function transferTo(address _to) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) TransferTo0(_to common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferTo0(&_L2TxFeeVault.TransactOpts, _to)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateMessenger(opts *bind.TransactOpts, _newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateMessenger", _newMessenger)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateMessenger(_newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMessenger(&_L2TxFeeVault.TransactOpts, _newMessenger)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateMessenger(_newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMessenger(&_L2TxFeeVault.TransactOpts, _newMessenger)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateMinWithdrawAmount(opts *bind.TransactOpts, _newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateMinWithdrawAmount", _newMinWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateMinWithdrawAmount(_newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMinWithdrawAmount(&_L2TxFeeVault.TransactOpts, _newMinWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateMinWithdrawAmount(_newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMinWithdrawAmount(&_L2TxFeeVault.TransactOpts, _newMinWithdrawAmount)
}

// UpdateReceiveAllowed is a paid mutator transaction binding the contract method 0x708125ad.
//
// Solidity: function updateReceiveAllowed(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateReceiveAllowed(opts *bind.TransactOpts, _accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateReceiveAllowed", _accounts, _status)
}

// UpdateReceiveAllowed is a paid mutator transaction binding the contract method 0x708125ad.
//
// Solidity: function updateReceiveAllowed(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateReceiveAllowed(_accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateReceiveAllowed(&_L2TxFeeVault.TransactOpts, _accounts, _status)
}

// UpdateReceiveAllowed is a paid mutator transaction binding the contract method 0x708125ad.
//
// Solidity: function updateReceiveAllowed(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateReceiveAllowed(_accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateReceiveAllowed(&_L2TxFeeVault.TransactOpts, _accounts, _status)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateRecipient(opts *bind.TransactOpts, _newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateRecipient", _newRecipient)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateRecipient(&_L2TxFeeVault.TransactOpts, _newRecipient)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateRecipient(&_L2TxFeeVault.TransactOpts, _newRecipient)
}

// UpdateTransferAllowedStatus is a paid mutator transaction binding the contract method 0xda13f9a2.
//
// Solidity: function updateTransferAllowedStatus(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateTransferAllowedStatus(opts *bind.TransactOpts, _accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateTransferAllowedStatus", _accounts, _status)
}

// UpdateTransferAllowedStatus is a paid mutator transaction binding the contract method 0xda13f9a2.
//
// Solidity: function updateTransferAllowedStatus(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateTransferAllowedStatus(_accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateTransferAllowedStatus(&_L2TxFeeVault.TransactOpts, _accounts, _status)
}

// UpdateTransferAllowedStatus is a paid mutator transaction binding the contract method 0xda13f9a2.
//
// Solidity: function updateTransferAllowedStatus(address[] _accounts, bool _status) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateTransferAllowedStatus(_accounts []common.Address, _status bool) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateTransferAllowedStatus(&_L2TxFeeVault.TransactOpts, _accounts, _status)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw(&_L2TxFeeVault.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw(&_L2TxFeeVault.TransactOpts, _value)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Withdraw0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "withdraw0")
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Withdraw0() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw0(&_L2TxFeeVault.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Withdraw0() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw0(&_L2TxFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Receive() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Receive(&_L2TxFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Receive(&_L2TxFeeVault.TransactOpts)
}

// L2TxFeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2TxFeeVault contract.
type L2TxFeeVaultOwnershipTransferredIterator struct {
	Event *L2TxFeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultOwnershipTransferred)
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
		it.Event = new(L2TxFeeVaultOwnershipTransferred)
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
func (it *L2TxFeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the L2TxFeeVault contract.
type L2TxFeeVaultOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _oldOwner []common.Address, _newOwner []common.Address) (*L2TxFeeVaultOwnershipTransferredIterator, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultOwnershipTransferredIterator{contract: _L2TxFeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultOwnershipTransferred, _oldOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultOwnershipTransferred)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*L2TxFeeVaultOwnershipTransferred, error) {
	event := new(L2TxFeeVaultOwnershipTransferred)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L2TxFeeVault contract.
type L2TxFeeVaultTransferIterator struct {
	Event *L2TxFeeVaultTransfer // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultTransfer)
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
		it.Event = new(L2TxFeeVaultTransfer)
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
func (it *L2TxFeeVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultTransfer represents a Transfer event raised by the L2TxFeeVault contract.
type L2TxFeeVaultTransfer struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x0a429aba3d89849a2db0153e4534d95c46a1d83c8109d73893f55ebc44010ff4.
//
// Solidity: event Transfer(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterTransfer(opts *bind.FilterOpts) (*L2TxFeeVaultTransferIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultTransferIterator{contract: _L2TxFeeVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x0a429aba3d89849a2db0153e4534d95c46a1d83c8109d73893f55ebc44010ff4.
//
// Solidity: event Transfer(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultTransfer) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultTransfer)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x0a429aba3d89849a2db0153e4534d95c46a1d83c8109d73893f55ebc44010ff4.
//
// Solidity: event Transfer(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseTransfer(log types.Log) (*L2TxFeeVaultTransfer, error) {
	event := new(L2TxFeeVaultTransfer)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateMessengerIterator is returned from FilterUpdateMessenger and is used to iterate over the raw logs and unpacked data for UpdateMessenger events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMessengerIterator struct {
	Event *L2TxFeeVaultUpdateMessenger // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultUpdateMessengerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateMessenger)
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
		it.Event = new(L2TxFeeVaultUpdateMessenger)
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
func (it *L2TxFeeVaultUpdateMessengerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateMessengerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateMessenger represents a UpdateMessenger event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMessenger struct {
	OldMessenger common.Address
	NewMessenger common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateMessenger is a free log retrieval operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateMessenger(opts *bind.FilterOpts, oldMessenger []common.Address, newMessenger []common.Address) (*L2TxFeeVaultUpdateMessengerIterator, error) {

	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateMessenger", oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateMessengerIterator{contract: _L2TxFeeVault.contract, event: "UpdateMessenger", logs: logs, sub: sub}, nil
}

// WatchUpdateMessenger is a free log subscription operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateMessenger(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateMessenger, oldMessenger []common.Address, newMessenger []common.Address) (event.Subscription, error) {

	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateMessenger", oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateMessenger)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMessenger", log); err != nil {
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

// ParseUpdateMessenger is a log parse operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateMessenger(log types.Log) (*L2TxFeeVaultUpdateMessenger, error) {
	event := new(L2TxFeeVaultUpdateMessenger)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMessenger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateMinWithdrawAmountIterator is returned from FilterUpdateMinWithdrawAmount and is used to iterate over the raw logs and unpacked data for UpdateMinWithdrawAmount events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMinWithdrawAmountIterator struct {
	Event *L2TxFeeVaultUpdateMinWithdrawAmount // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateMinWithdrawAmount)
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
		it.Event = new(L2TxFeeVaultUpdateMinWithdrawAmount)
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
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateMinWithdrawAmount represents a UpdateMinWithdrawAmount event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMinWithdrawAmount struct {
	OldMinWithdrawAmount *big.Int
	NewMinWithdrawAmount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinWithdrawAmount is a free log retrieval operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateMinWithdrawAmount(opts *bind.FilterOpts) (*L2TxFeeVaultUpdateMinWithdrawAmountIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateMinWithdrawAmountIterator{contract: _L2TxFeeVault.contract, event: "UpdateMinWithdrawAmount", logs: logs, sub: sub}, nil
}

// WatchUpdateMinWithdrawAmount is a free log subscription operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateMinWithdrawAmount(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateMinWithdrawAmount) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateMinWithdrawAmount)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMinWithdrawAmount", log); err != nil {
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

// ParseUpdateMinWithdrawAmount is a log parse operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateMinWithdrawAmount(log types.Log) (*L2TxFeeVaultUpdateMinWithdrawAmount, error) {
	event := new(L2TxFeeVaultUpdateMinWithdrawAmount)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMinWithdrawAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateReceiveAllowedIterator is returned from FilterUpdateReceiveAllowed and is used to iterate over the raw logs and unpacked data for UpdateReceiveAllowed events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateReceiveAllowedIterator struct {
	Event *L2TxFeeVaultUpdateReceiveAllowed // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultUpdateReceiveAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateReceiveAllowed)
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
		it.Event = new(L2TxFeeVaultUpdateReceiveAllowed)
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
func (it *L2TxFeeVaultUpdateReceiveAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateReceiveAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateReceiveAllowed represents a UpdateReceiveAllowed event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateReceiveAllowed struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateReceiveAllowed is a free log retrieval operation binding the contract event 0xbd84b6b3e1d029fd61a717f6020a8f35b358486e0971c00d4e64b516503f85e7.
//
// Solidity: event UpdateReceiveAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateReceiveAllowed(opts *bind.FilterOpts) (*L2TxFeeVaultUpdateReceiveAllowedIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateReceiveAllowed")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateReceiveAllowedIterator{contract: _L2TxFeeVault.contract, event: "UpdateReceiveAllowed", logs: logs, sub: sub}, nil
}

// WatchUpdateReceiveAllowed is a free log subscription operation binding the contract event 0xbd84b6b3e1d029fd61a717f6020a8f35b358486e0971c00d4e64b516503f85e7.
//
// Solidity: event UpdateReceiveAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateReceiveAllowed(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateReceiveAllowed) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateReceiveAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateReceiveAllowed)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateReceiveAllowed", log); err != nil {
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

// ParseUpdateReceiveAllowed is a log parse operation binding the contract event 0xbd84b6b3e1d029fd61a717f6020a8f35b358486e0971c00d4e64b516503f85e7.
//
// Solidity: event UpdateReceiveAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateReceiveAllowed(log types.Log) (*L2TxFeeVaultUpdateReceiveAllowed, error) {
	event := new(L2TxFeeVaultUpdateReceiveAllowed)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateReceiveAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateRecipientIterator is returned from FilterUpdateRecipient and is used to iterate over the raw logs and unpacked data for UpdateRecipient events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateRecipientIterator struct {
	Event *L2TxFeeVaultUpdateRecipient // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultUpdateRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateRecipient)
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
		it.Event = new(L2TxFeeVaultUpdateRecipient)
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
func (it *L2TxFeeVaultUpdateRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateRecipient represents a UpdateRecipient event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateRecipient struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateRecipient is a free log retrieval operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateRecipient(opts *bind.FilterOpts, oldRecipient []common.Address, newRecipient []common.Address) (*L2TxFeeVaultUpdateRecipientIterator, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateRecipient", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateRecipientIterator{contract: _L2TxFeeVault.contract, event: "UpdateRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateRecipient is a free log subscription operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateRecipient(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateRecipient, oldRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateRecipient", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateRecipient)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateRecipient", log); err != nil {
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

// ParseUpdateRecipient is a log parse operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateRecipient(log types.Log) (*L2TxFeeVaultUpdateRecipient, error) {
	event := new(L2TxFeeVaultUpdateRecipient)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateTransferAllowedIterator is returned from FilterUpdateTransferAllowed and is used to iterate over the raw logs and unpacked data for UpdateTransferAllowed events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateTransferAllowedIterator struct {
	Event *L2TxFeeVaultUpdateTransferAllowed // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultUpdateTransferAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateTransferAllowed)
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
		it.Event = new(L2TxFeeVaultUpdateTransferAllowed)
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
func (it *L2TxFeeVaultUpdateTransferAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateTransferAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateTransferAllowed represents a UpdateTransferAllowed event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateTransferAllowed struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTransferAllowed is a free log retrieval operation binding the contract event 0xbb5d3e442e800faa1089a4f57bae4f36808d3cf15d051033d78a72147782f24c.
//
// Solidity: event UpdateTransferAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateTransferAllowed(opts *bind.FilterOpts) (*L2TxFeeVaultUpdateTransferAllowedIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateTransferAllowed")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateTransferAllowedIterator{contract: _L2TxFeeVault.contract, event: "UpdateTransferAllowed", logs: logs, sub: sub}, nil
}

// WatchUpdateTransferAllowed is a free log subscription operation binding the contract event 0xbb5d3e442e800faa1089a4f57bae4f36808d3cf15d051033d78a72147782f24c.
//
// Solidity: event UpdateTransferAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateTransferAllowed(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateTransferAllowed) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateTransferAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateTransferAllowed)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateTransferAllowed", log); err != nil {
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

// ParseUpdateTransferAllowed is a log parse operation binding the contract event 0xbb5d3e442e800faa1089a4f57bae4f36808d3cf15d051033d78a72147782f24c.
//
// Solidity: event UpdateTransferAllowed(address account, bool status)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateTransferAllowed(log types.Log) (*L2TxFeeVaultUpdateTransferAllowed, error) {
	event := new(L2TxFeeVaultUpdateTransferAllowed)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateTransferAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the L2TxFeeVault contract.
type L2TxFeeVaultWithdrawalIterator struct {
	Event *L2TxFeeVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *L2TxFeeVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultWithdrawal)
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
		it.Event = new(L2TxFeeVaultWithdrawal)
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
func (it *L2TxFeeVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultWithdrawal represents a Withdrawal event raised by the L2TxFeeVault contract.
type L2TxFeeVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*L2TxFeeVaultWithdrawalIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultWithdrawalIterator{contract: _L2TxFeeVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultWithdrawal)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseWithdrawal(log types.Log) (*L2TxFeeVaultWithdrawal, error) {
	event := new(L2TxFeeVaultWithdrawal)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
