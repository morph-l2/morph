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

// L2WithdrawLockERC20GatewayMetaData contains all meta data concerning the L2WithdrawLockERC20Gateway contract.
var L2WithdrawLockERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"lock\",\"type\":\"bool\"}],\"name\":\"UpdateWithdrawLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_lock\",\"type\":\"bool\"}],\"name\":\"updateWithdrawLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611c33806100e65f395ff3fe608060405260043610610109575f3560e01c8063a93a4af9116100a1578063cdd0da7c11610071578063f2fde38b11610057578063f2fde38b1461033b578063f887ea401461035a578063fac752eb14610386575f80fd5b8063cdd0da7c146102de578063ebc137d01461031c575f80fd5b8063a93a4af91461024c578063ba27f50b1461025f578063c0c53b8b146102a0578063c676ad29146102bf575f80fd5b8063715018a6116100dc578063715018a6146101cf578063797594b0146101e35780638431f5c11461020f5780638da5cb5b14610222575f80fd5b80633cb747bf1461010d57806354bbd59c14610163578063575361b6146101a75780636c07ea43146101bc575b5f80fd5b348015610118575f80fd5b506099546101399073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561016e575f80fd5b5061013961017d3660046116cc565b73ffffffffffffffffffffffffffffffffffffffff9081165f90815260fa60205260409020541690565b6101ba6101b5366004611733565b6103a5565b005b6101ba6101ca3660046117a9565b6103f0565b3480156101da575f80fd5b506101ba61042e565b3480156101ee575f80fd5b506097546101399073ffffffffffffffffffffffffffffffffffffffff1681565b6101ba61021d3660046117db565b610441565b34801561022d575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610139565b6101ba61025a36600461186d565b61080d565b34801561026a575f80fd5b506101396102793660046116cc565b60fa6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102ab575f80fd5b506101ba6102ba3660046118b0565b61081f565b3480156102ca575f80fd5b506101396102d93660046116cc565b6109f9565b3480156102e9575f80fd5b5061030c6102f83660046116cc565b60fb6020525f908152604090205460ff1681565b604051901515815260200161015a565b348015610327575f80fd5b506101ba6103363660046118f8565b610a43565b348015610346575f80fd5b506101ba6103553660046116cc565b610b37565b348015610365575f80fd5b506098546101399073ffffffffffffffffffffffffffffffffffffffff1681565b348015610391575f80fd5b506101ba6103a0366004611933565b610bd4565b6103e886868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610cfa915050565b505050505050565b6104298333845f5b6040519080825280601f01601f191660200182016040528015610422576020820181803683370190505b5085610cfa565b505050565b610436611145565b61043f5f6111ac565b565b60995473ffffffffffffffffffffffffffffffffffffffff163381146104ae5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104f7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051b919061198c565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146105855760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016104a5565b61058d611222565b34156105db5760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff881661063e5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff8088165f90815260fa60205260409020548982169116146106b65760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016104a5565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690528816906340c10f19906044015f604051808303815f87803b158015610723575f80fd5b505af1158015610735573d5f803e3d5ffd5b505050506107788584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061127b92505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34888888886040516107f294939291906119a7565b60405180910390a461080360018055565b5050505050505050565b6108198484845f6103f8565b50505050565b5f54610100900460ff161580801561083d57505f54600160ff909116105b806108565750303b15801561085657505f5460ff166001145b6108c85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104a5565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610924575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff83166109875760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016104a5565b61099284848461132b565b8015610819575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b60405162461bcd60e51b815260206004820152600d60248201527f756e696d706c656d656e7465640000000000000000000000000000000000000060448201525f906064016104a5565b610a4b611145565b73ffffffffffffffffffffffffffffffffffffffff8216610aae5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff82165f81815260fb602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915591519182527fd8f6792507085b7664354b4599c60b3b600bd3f7e1a758f5e37134d4816b044a910160405180910390a25050565b610b3f611145565b73ffffffffffffffffffffffffffffffffffffffff8116610bc85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104a5565b610bd1816111ac565b50565b610bdc611145565b73ffffffffffffffffffffffffffffffffffffffff8116610c3f5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815260fa6020908152604080832080548787167fffffffffffffffffffffffff00000000000000000000000000000000000000008216811790925560fb90935281842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790559051919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610d02611222565b73ffffffffffffffffffffffffffffffffffffffff85165f90815260fb602052604090205460ff1615610d775760405162461bcd60e51b815260206004820152600d60248201527f7769746864726177206c6f636b0000000000000000000000000000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260fa60205260409020541680610deb5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016104a5565b5f8411610e3a5760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e7400000000000000000000000060448201526064016104a5565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610e755783806020019051810190610e709190611a39565b945090505b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff828116600483015260248201879052881690639dc29fac906044015f604051808303815f87803b158015610ee2575f80fd5b505af1158015610ef4573d5f803e3d5ffd5b505050505f828883898989604051602401610f1496959493929190611b5e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa158015610ff8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101c9190611bb8565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b923492611081929116905f9088908c90600401611bcf565b5f604051808303818588803b158015611098575f80fd5b505af11580156110aa573d5f803e3d5ffd5b50505050508273ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b876040516111299493929190611bcf565b60405180910390a45050505061113e60018055565b5050505050565b60655473ffffffffffffffffffffffffffffffffffffffff16331461043f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104a5565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6002600154036112745760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104a5565b6002600155565b5f81511180156112a157505f8273ffffffffffffffffffffffffffffffffffffffff163b115b15611321576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f906112f8908490600401611c14565b5f604051808303815f87803b15801561130f575f80fd5b505af11580156103e8573d5f803e3d5ffd5b5050565b60018055565b73ffffffffffffffffffffffffffffffffffffffff831661138e5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016104a5565b73ffffffffffffffffffffffffffffffffffffffff81166113f15760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016104a5565b6113f96114a2565b611401611526565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609980548484169216919091179055821615610429576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff1661151e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a5565b61043f6115aa565b5f54610100900460ff166115a25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a5565b61043f611626565b5f54610100900460ff166113255760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a5565b5f54610100900460ff166116a25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a5565b61043f336111ac565b73ffffffffffffffffffffffffffffffffffffffff81168114610bd1575f80fd5b5f602082840312156116dc575f80fd5b81356116e7816116ab565b9392505050565b5f8083601f8401126116fe575f80fd5b50813567ffffffffffffffff811115611715575f80fd5b60208301915083602082850101111561172c575f80fd5b9250929050565b5f805f805f8060a08789031215611748575f80fd5b8635611753816116ab565b95506020870135611763816116ab565b945060408701359350606087013567ffffffffffffffff811115611785575f80fd5b61179189828a016116ee565b979a9699509497949695608090950135949350505050565b5f805f606084860312156117bb575f80fd5b83356117c6816116ab565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a0312156117f1575f80fd5b87356117fc816116ab565b9650602088013561180c816116ab565b9550604088013561181c816116ab565b9450606088013561182c816116ab565b93506080880135925060a088013567ffffffffffffffff81111561184e575f80fd5b61185a8a828b016116ee565b989b979a50959850939692959293505050565b5f805f8060808587031215611880575f80fd5b843561188b816116ab565b9350602085013561189b816116ab565b93969395505050506040820135916060013590565b5f805f606084860312156118c2575f80fd5b83356118cd816116ab565b925060208401356118dd816116ab565b915060408401356118ed816116ab565b809150509250925092565b5f8060408385031215611909575f80fd5b8235611914816116ab565b915060208301358015158114611928575f80fd5b809150509250929050565b5f8060408385031215611944575f80fd5b823561194f816116ab565b91506020830135611928816116ab565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020828403121561199c575f80fd5b81516116e7816116ab565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611a31578181015183820152602001611a19565b50505f910152565b5f8060408385031215611a4a575f80fd5b8251611a55816116ab565b602084015190925067ffffffffffffffff80821115611a72575f80fd5b818501915085601f830112611a85575f80fd5b815181811115611a9757611a9761195f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611add57611add61195f565b81604052828152886020848701011115611af5575f80fd5b611b06836020830160208801611a17565b80955050505050509250929050565b5f8151808452611b2c816020860160208601611a17565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611bac60c0830184611b15565b98975050505050505050565b5f60208284031215611bc8575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f611c036080830185611b15565b905082606083015295945050505050565b602081525f6116e76020830184611b1556fea164736f6c6343000818000a",
}

// L2WithdrawLockERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2WithdrawLockERC20GatewayMetaData.ABI instead.
var L2WithdrawLockERC20GatewayABI = L2WithdrawLockERC20GatewayMetaData.ABI

// L2WithdrawLockERC20GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2WithdrawLockERC20GatewayMetaData.Bin instead.
var L2WithdrawLockERC20GatewayBin = L2WithdrawLockERC20GatewayMetaData.Bin

// DeployL2WithdrawLockERC20Gateway deploys a new Ethereum contract, binding an instance of L2WithdrawLockERC20Gateway to it.
func DeployL2WithdrawLockERC20Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2WithdrawLockERC20Gateway, error) {
	parsed, err := L2WithdrawLockERC20GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2WithdrawLockERC20GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2WithdrawLockERC20Gateway{L2WithdrawLockERC20GatewayCaller: L2WithdrawLockERC20GatewayCaller{contract: contract}, L2WithdrawLockERC20GatewayTransactor: L2WithdrawLockERC20GatewayTransactor{contract: contract}, L2WithdrawLockERC20GatewayFilterer: L2WithdrawLockERC20GatewayFilterer{contract: contract}}, nil
}

// L2WithdrawLockERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L2WithdrawLockERC20Gateway struct {
	L2WithdrawLockERC20GatewayCaller     // Read-only binding to the contract
	L2WithdrawLockERC20GatewayTransactor // Write-only binding to the contract
	L2WithdrawLockERC20GatewayFilterer   // Log filterer for contract events
}

// L2WithdrawLockERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2WithdrawLockERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawLockERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2WithdrawLockERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawLockERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2WithdrawLockERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawLockERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2WithdrawLockERC20GatewaySession struct {
	Contract     *L2WithdrawLockERC20Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2WithdrawLockERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2WithdrawLockERC20GatewayCallerSession struct {
	Contract *L2WithdrawLockERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// L2WithdrawLockERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2WithdrawLockERC20GatewayTransactorSession struct {
	Contract     *L2WithdrawLockERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// L2WithdrawLockERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2WithdrawLockERC20GatewayRaw struct {
	Contract *L2WithdrawLockERC20Gateway // Generic contract binding to access the raw methods on
}

// L2WithdrawLockERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2WithdrawLockERC20GatewayCallerRaw struct {
	Contract *L2WithdrawLockERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2WithdrawLockERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2WithdrawLockERC20GatewayTransactorRaw struct {
	Contract *L2WithdrawLockERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2WithdrawLockERC20Gateway creates a new instance of L2WithdrawLockERC20Gateway, bound to a specific deployed contract.
func NewL2WithdrawLockERC20Gateway(address common.Address, backend bind.ContractBackend) (*L2WithdrawLockERC20Gateway, error) {
	contract, err := bindL2WithdrawLockERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20Gateway{L2WithdrawLockERC20GatewayCaller: L2WithdrawLockERC20GatewayCaller{contract: contract}, L2WithdrawLockERC20GatewayTransactor: L2WithdrawLockERC20GatewayTransactor{contract: contract}, L2WithdrawLockERC20GatewayFilterer: L2WithdrawLockERC20GatewayFilterer{contract: contract}}, nil
}

// NewL2WithdrawLockERC20GatewayCaller creates a new read-only instance of L2WithdrawLockERC20Gateway, bound to a specific deployed contract.
func NewL2WithdrawLockERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2WithdrawLockERC20GatewayCaller, error) {
	contract, err := bindL2WithdrawLockERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayCaller{contract: contract}, nil
}

// NewL2WithdrawLockERC20GatewayTransactor creates a new write-only instance of L2WithdrawLockERC20Gateway, bound to a specific deployed contract.
func NewL2WithdrawLockERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2WithdrawLockERC20GatewayTransactor, error) {
	contract, err := bindL2WithdrawLockERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayTransactor{contract: contract}, nil
}

// NewL2WithdrawLockERC20GatewayFilterer creates a new log filterer instance of L2WithdrawLockERC20Gateway, bound to a specific deployed contract.
func NewL2WithdrawLockERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2WithdrawLockERC20GatewayFilterer, error) {
	contract, err := bindL2WithdrawLockERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayFilterer{contract: contract}, nil
}

// bindL2WithdrawLockERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL2WithdrawLockERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2WithdrawLockERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WithdrawLockERC20Gateway.Contract.L2WithdrawLockERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.L2WithdrawLockERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.L2WithdrawLockERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WithdrawLockERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) Counterpart() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Counterpart(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Counterpart(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.GetL1ERC20Address(&_L2WithdrawLockERC20Gateway.CallOpts, _l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.GetL1ERC20Address(&_L2WithdrawLockERC20Gateway.CallOpts, _l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.GetL2ERC20Address(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.GetL2ERC20Address(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) Messenger() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Messenger(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) Messenger() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Messenger(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) Owner() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Owner(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) Owner() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Owner(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) Router() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Router(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Router(&_L2WithdrawLockERC20Gateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.TokenMapping(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2WithdrawLockERC20Gateway.Contract.TokenMapping(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// WithdrawLock is a free data retrieval call binding the contract method 0xcdd0da7c.
//
// Solidity: function withdrawLock(address ) view returns(bool)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCaller) WithdrawLock(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L2WithdrawLockERC20Gateway.contract.Call(opts, &out, "withdrawLock", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawLock is a free data retrieval call binding the contract method 0xcdd0da7c.
//
// Solidity: function withdrawLock(address ) view returns(bool)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) WithdrawLock(arg0 common.Address) (bool, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawLock(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// WithdrawLock is a free data retrieval call binding the contract method 0xcdd0da7c.
//
// Solidity: function withdrawLock(address ) view returns(bool)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayCallerSession) WithdrawLock(arg0 common.Address) (bool, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawLock(&_L2WithdrawLockERC20Gateway.CallOpts, arg0)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.FinalizeDepositERC20(&_L2WithdrawLockERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.FinalizeDepositERC20(&_L2WithdrawLockERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Initialize(&_L2WithdrawLockERC20Gateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.Initialize(&_L2WithdrawLockERC20Gateway.TransactOpts, _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.RenounceOwnership(&_L2WithdrawLockERC20Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.RenounceOwnership(&_L2WithdrawLockERC20Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.TransferOwnership(&_L2WithdrawLockERC20Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.TransferOwnership(&_L2WithdrawLockERC20Gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.UpdateTokenMapping(&_L2WithdrawLockERC20Gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.UpdateTokenMapping(&_L2WithdrawLockERC20Gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateWithdrawLock is a paid mutator transaction binding the contract method 0xebc137d0.
//
// Solidity: function updateWithdrawLock(address _l2Token, bool _lock) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) UpdateWithdrawLock(opts *bind.TransactOpts, _l2Token common.Address, _lock bool) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "updateWithdrawLock", _l2Token, _lock)
}

// UpdateWithdrawLock is a paid mutator transaction binding the contract method 0xebc137d0.
//
// Solidity: function updateWithdrawLock(address _l2Token, bool _lock) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) UpdateWithdrawLock(_l2Token common.Address, _lock bool) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.UpdateWithdrawLock(&_L2WithdrawLockERC20Gateway.TransactOpts, _l2Token, _lock)
}

// UpdateWithdrawLock is a paid mutator transaction binding the contract method 0xebc137d0.
//
// Solidity: function updateWithdrawLock(address _l2Token, bool _lock) returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) UpdateWithdrawLock(_l2Token common.Address, _lock bool) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.UpdateWithdrawLock(&_L2WithdrawLockERC20Gateway.TransactOpts, _l2Token, _lock)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC20(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC20(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC200(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC200(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC20AndCall(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WithdrawLockERC20Gateway.Contract.WithdrawERC20AndCall(&_L2WithdrawLockERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator struct {
	Event *L2WithdrawLockERC20GatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayFinalizeDepositERC20)
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
		it.Event = new(L2WithdrawLockERC20GatewayFinalizeDepositERC20)
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
func (it *L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayFinalizeDepositERC20Iterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayFinalizeDepositERC20)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2WithdrawLockERC20GatewayFinalizeDepositERC20, error) {
	event := new(L2WithdrawLockERC20GatewayFinalizeDepositERC20)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawLockERC20GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayInitializedIterator struct {
	Event *L2WithdrawLockERC20GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayInitialized)
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
		it.Event = new(L2WithdrawLockERC20GatewayInitialized)
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
func (it *L2WithdrawLockERC20GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayInitialized represents a Initialized event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2WithdrawLockERC20GatewayInitializedIterator, error) {

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayInitializedIterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayInitialized)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseInitialized(log types.Log) (*L2WithdrawLockERC20GatewayInitialized, error) {
	event := new(L2WithdrawLockERC20GatewayInitialized)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawLockERC20GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayOwnershipTransferredIterator struct {
	Event *L2WithdrawLockERC20GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayOwnershipTransferred)
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
		it.Event = new(L2WithdrawLockERC20GatewayOwnershipTransferred)
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
func (it *L2WithdrawLockERC20GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2WithdrawLockERC20GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayOwnershipTransferredIterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayOwnershipTransferred)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2WithdrawLockERC20GatewayOwnershipTransferred, error) {
	event := new(L2WithdrawLockERC20GatewayOwnershipTransferred)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawLockERC20GatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayUpdateTokenMappingIterator struct {
	Event *L2WithdrawLockERC20GatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayUpdateTokenMapping)
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
		it.Event = new(L2WithdrawLockERC20GatewayUpdateTokenMapping)
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
func (it *L2WithdrawLockERC20GatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2WithdrawLockERC20GatewayUpdateTokenMappingIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayUpdateTokenMappingIterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayUpdateTokenMapping)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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

// ParseUpdateTokenMapping is a log parse operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2WithdrawLockERC20GatewayUpdateTokenMapping, error) {
	event := new(L2WithdrawLockERC20GatewayUpdateTokenMapping)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator is returned from FilterUpdateWithdrawLock and is used to iterate over the raw logs and unpacked data for UpdateWithdrawLock events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator struct {
	Event *L2WithdrawLockERC20GatewayUpdateWithdrawLock // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayUpdateWithdrawLock)
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
		it.Event = new(L2WithdrawLockERC20GatewayUpdateWithdrawLock)
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
func (it *L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayUpdateWithdrawLock represents a UpdateWithdrawLock event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayUpdateWithdrawLock struct {
	L2Token common.Address
	Lock    bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawLock is a free log retrieval operation binding the contract event 0xd8f6792507085b7664354b4599c60b3b600bd3f7e1a758f5e37134d4816b044a.
//
// Solidity: event UpdateWithdrawLock(address indexed l2Token, bool lock)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterUpdateWithdrawLock(opts *bind.FilterOpts, l2Token []common.Address) (*L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "UpdateWithdrawLock", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayUpdateWithdrawLockIterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "UpdateWithdrawLock", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawLock is a free log subscription operation binding the contract event 0xd8f6792507085b7664354b4599c60b3b600bd3f7e1a758f5e37134d4816b044a.
//
// Solidity: event UpdateWithdrawLock(address indexed l2Token, bool lock)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchUpdateWithdrawLock(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayUpdateWithdrawLock, l2Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "UpdateWithdrawLock", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayUpdateWithdrawLock)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "UpdateWithdrawLock", log); err != nil {
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

// ParseUpdateWithdrawLock is a log parse operation binding the contract event 0xd8f6792507085b7664354b4599c60b3b600bd3f7e1a758f5e37134d4816b044a.
//
// Solidity: event UpdateWithdrawLock(address indexed l2Token, bool lock)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseUpdateWithdrawLock(log types.Log) (*L2WithdrawLockERC20GatewayUpdateWithdrawLock, error) {
	event := new(L2WithdrawLockERC20GatewayUpdateWithdrawLock)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "UpdateWithdrawLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WithdrawLockERC20GatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayWithdrawERC20Iterator struct {
	Event *L2WithdrawLockERC20GatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2WithdrawLockERC20GatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WithdrawLockERC20GatewayWithdrawERC20)
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
		it.Event = new(L2WithdrawLockERC20GatewayWithdrawERC20)
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
func (it *L2WithdrawLockERC20GatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WithdrawLockERC20GatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WithdrawLockERC20GatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2WithdrawLockERC20Gateway contract.
type L2WithdrawLockERC20GatewayWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WithdrawLockERC20GatewayWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawLockERC20GatewayWithdrawERC20Iterator{contract: _L2WithdrawLockERC20Gateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2WithdrawLockERC20GatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WithdrawLockERC20Gateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WithdrawLockERC20GatewayWithdrawERC20)
				if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2WithdrawLockERC20Gateway *L2WithdrawLockERC20GatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2WithdrawLockERC20GatewayWithdrawERC20, error) {
	event := new(L2WithdrawLockERC20GatewayWithdrawERC20)
	if err := _L2WithdrawLockERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
