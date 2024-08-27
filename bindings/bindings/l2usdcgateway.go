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

// L2USDCGatewayMetaData contains all meta data concerning the L2USDCGateway contract.
var L2USDCGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1USDC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2USDC\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"circleCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"pauseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"pauseWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"transferUSDCRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateCircleCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b50604051620022903803806200229083398101604081905262000033916200012f565b6200003d62000055565b6001600160a01b039182166080521660a05262000165565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161462000111575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200012a575f80fd5b919050565b5f806040838503121562000141575f80fd5b6200014c8362000113565b91506200015c6020840162000113565b90509250929050565b60805160a0516120d8620001b85f395f81816102190152818161041e0152818161080d01528181610cd30152610f2c01525f81816102d7015281816103ad01528181610772015261115b01526120d85ff3fe60806040526004361061016d575f3560e01c80638431f5c1116100c6578063c676ad291161007c578063f0d7c29c11610057578063f0d7c29c1461047e578063f2fde38b1461049d578063f887ea40146104bc575f80fd5b8063c676ad2914610401578063c689fc3414610440578063ebd462cb1461045f575f80fd5b8063a6f73669116100ac578063a6f736691461039c578063a93a4af9146103cf578063c0c53b8b146103e2575f80fd5b80638431f5c11461035f5780638da5cb5b14610372575f80fd5b8063415855d6116101265780636c07ea43116101015780636c07ea431461030c578063715018a61461031f578063797594b014610333575f80fd5b8063415855d61461029957806354bbd59c146102ba578063575361b6146102f9575f80fd5b806329e96f9e1161015657806329e96f9e146102085780632f3ffb9f1461023b5780633cb747bf1461026d575f80fd5b806302befd24146101715780631f878ae6146101b7575b5f80fd5b34801561017c575f80fd5b5060fa546101a29074010000000000000000000000000000000000000000900460ff1681565b60405190151581526020015b60405180910390f35b3480156101c2575f80fd5b5060fa546101e39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ae565b348015610213575f80fd5b506101e37f000000000000000000000000000000000000000000000000000000000000000081565b348015610246575f80fd5b5060fa546101a2907501000000000000000000000000000000000000000000900460ff1681565b348015610278575f80fd5b506099546101e39073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102a4575f80fd5b506102b86102b3366004611b66565b6104e8565b005b3480156102c5575f80fd5b506101e36102d4366004611ba9565b507f000000000000000000000000000000000000000000000000000000000000000090565b6102b8610307366004611c09565b61053a565b6102b861031a366004611c7f565b610585565b34801561032a575f80fd5b506102b86105c3565b34801561033e575f80fd5b506097546101e39073ffffffffffffffffffffffffffffffffffffffff1681565b6102b861036d366004611cb1565b6105d6565b34801561037d575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff166101e3565b3480156103a7575f80fd5b506101e37f000000000000000000000000000000000000000000000000000000000000000081565b6102b86103dd366004611d43565b610a88565b3480156103ed575f80fd5b506102b86103fc366004611d86565b610a9a565b34801561040c575f80fd5b506101e361041b366004611ba9565b507f000000000000000000000000000000000000000000000000000000000000000090565b34801561044b575f80fd5b506102b861045a366004611ba9565b610c11565b34801561046a575f80fd5b506102b8610479366004611b66565b610d2d565b348015610489575f80fd5b506102b8610498366004611ba9565b610d80565b3480156104a8575f80fd5b506102b86104b7366004611ba9565b610dcf565b3480156104c7575f80fd5b506098546101e39073ffffffffffffffffffffffffffffffffffffffff1681565b6104f0610e6c565b60fa805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b61057d86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610ed3915050565b505050505050565b6105be8333845f5b6040519080825280601f01601f1916602001820160405280156105b7576020820181803683370190505b5085610ed3565b505050565b6105cb610e6c565b6105d45f6113c1565b565b60995473ffffffffffffffffffffffffffffffffffffffff163381146106435760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106b09190611dfb565b60975473ffffffffffffffffffffffffffffffffffffffff90811691161461071a5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161063a565b610722611437565b34156107705760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c7565000000000000000000000000000000604482015260640161063a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff161461080b5760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206e6f742055534443000000000000000000000000000000604482015260640161063a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16146108a65760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206e6f742055534443000000000000000000000000000000604482015260640161063a565b60fa5474010000000000000000000000000000000000000000900460ff16156109115760405162461bcd60e51b815260206004820152600e60248201527f6465706f73697420706175736564000000000000000000000000000000000000604482015260640161063a565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690528816906340c10f19906044016020604051808303815f875af1158015610983573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a79190611e16565b6109f35760405162461bcd60e51b815260206004820152601060248201527f6d696e742055534443206661696c656400000000000000000000000000000000604482015260640161063a565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba3488888888604051610a6d9493929190611e31565b60405180910390a4610a7e60018055565b5050505050505050565b610a948484845f61058d565b50505050565b5f54610100900460ff1615808015610ab857505f54600160ff909116105b80610ad15750303b158015610ad157505f5460ff166001145b610b435760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161063a565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610b9f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610baa848484611496565b8015610a94575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b60fa5473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c8e5760405162461bcd60e51b815260206004820152601260248201527f6f6e6c7920636972636c652063616c6c65720000000000000000000000000000604482015260640161063a565b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b906024015f604051808303815f87803b158015610d14575f80fd5b505af1158015610d26573d5f803e3d5ffd5b5050505050565b610d35610e6c565b60fa80549115157501000000000000000000000000000000000000000000027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff909216919091179055565b610d88610e6c565b60fa80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610dd7610e6c565b73ffffffffffffffffffffffffffffffffffffffff8116610e605760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161063a565b610e69816113c1565b50565b60655473ffffffffffffffffffffffffffffffffffffffff1633146105d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161063a565b610edb611437565b5f8311610f2a5760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e74000000000000000000000000604482015260640161063a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610fc55760405162461bcd60e51b815260206004820152601460248201527f6f6e6c79205553444320697320616c6c6f776564000000000000000000000000604482015260640161063a565b60fa547501000000000000000000000000000000000000000000900460ff16156110315760405162461bcd60e51b815260206004820152600f60248201527f7769746864726177207061757365640000000000000000000000000000000000604482015260640161063a565b609854339073ffffffffffffffffffffffffffffffffffffffff1681900361106c57828060200190518101906110679190611ec3565b935090505b8251156110bb5760405162461bcd60e51b815260206004820152601360248201527f63616c6c206973206e6f7420616c6c6f77656400000000000000000000000000604482015260640161063a565b6110dd73ffffffffffffffffffffffffffffffffffffffff871682308761160d565b6040517f42966c680000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff8716906342966c68906024015f604051808303815f87803b158015611142575f80fd5b505af1158015611154573d5f803e3d5ffd5b50506040517f000000000000000000000000000000000000000000000000000000000000000092505f91506111979083908a9086908b908b908b90602401611fe8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa15801561127b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061129f9190612042565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b923492611304929116905f9088908c90600401612059565b5f604051808303818588803b15801561131b575f80fd5b505af115801561132d573d5f803e3d5ffd5b50505050508373ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b876040516113ac9493929190612059565b60405180910390a450505050610d2660018055565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6002600154036114895760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161063a565b6002600155565b60018055565b73ffffffffffffffffffffffffffffffffffffffff83166114f95760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161063a565b73ffffffffffffffffffffffffffffffffffffffff811661155c5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161063a565b6115646116a2565b61156c611726565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105be576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610a949085906117aa565b5f54610100900460ff1661171e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161063a565b6105d461189d565b5f54610100900460ff166117a25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161063a565b6105d4611919565b5f61180b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661199e9092919063ffffffff16565b905080515f148061182b57508080602001905181019061182b9190611e16565b6105be5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161063a565b5f54610100900460ff166114905760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161063a565b5f54610100900460ff166119955760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161063a565b6105d4336113c1565b60606119ac84845f856119b4565b949350505050565b606082471015611a2c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161063a565b5f808673ffffffffffffffffffffffffffffffffffffffff168587604051611a54919061209e565b5f6040518083038185875af1925050503d805f8114611a8e576040519150601f19603f3d011682016040523d82523d5f602084013e611a93565b606091505b5091509150611aa487838387611aaf565b979650505050505050565b60608315611b2a5782515f03611b235773ffffffffffffffffffffffffffffffffffffffff85163b611b235760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161063a565b50816119ac565b6119ac8383815115611b3f5781518083602001fd5b8060405162461bcd60e51b815260040161063a91906120b9565b8015158114610e69575f80fd5b5f60208284031215611b76575f80fd5b8135611b8181611b59565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610e69575f80fd5b5f60208284031215611bb9575f80fd5b8135611b8181611b88565b5f8083601f840112611bd4575f80fd5b50813567ffffffffffffffff811115611beb575f80fd5b602083019150836020828501011115611c02575f80fd5b9250929050565b5f805f805f8060a08789031215611c1e575f80fd5b8635611c2981611b88565b95506020870135611c3981611b88565b945060408701359350606087013567ffffffffffffffff811115611c5b575f80fd5b611c6789828a01611bc4565b979a9699509497949695608090950135949350505050565b5f805f60608486031215611c91575f80fd5b8335611c9c81611b88565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a031215611cc7575f80fd5b8735611cd281611b88565b96506020880135611ce281611b88565b95506040880135611cf281611b88565b94506060880135611d0281611b88565b93506080880135925060a088013567ffffffffffffffff811115611d24575f80fd5b611d308a828b01611bc4565b989b979a50959850939692959293505050565b5f805f8060808587031215611d56575f80fd5b8435611d6181611b88565b93506020850135611d7181611b88565b93969395505050506040820135916060013590565b5f805f60608486031215611d98575f80fd5b8335611da381611b88565b92506020840135611db381611b88565b91506040840135611dc381611b88565b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215611e0b575f80fd5b8151611b8181611b88565b5f60208284031215611e26575f80fd5b8151611b8181611b59565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611ebb578181015183820152602001611ea3565b50505f910152565b5f8060408385031215611ed4575f80fd5b8251611edf81611b88565b602084015190925067ffffffffffffffff80821115611efc575f80fd5b818501915085601f830112611f0f575f80fd5b815181811115611f2157611f21611dce565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611f6757611f67611dce565b81604052828152886020848701011115611f7f575f80fd5b611f90836020830160208801611ea1565b80955050505050509250929050565b5f8151808452611fb6816020860160208601611ea1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261203660c0830184611f9f565b98975050505050505050565b5f60208284031215612052575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f61208d6080830185611f9f565b905082606083015295945050505050565b5f82516120af818460208701611ea1565b9190910192915050565b602081525f611b816020830184611f9f56fea164736f6c6343000818000a",
}

// L2USDCGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2USDCGatewayMetaData.ABI instead.
var L2USDCGatewayABI = L2USDCGatewayMetaData.ABI

// L2USDCGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2USDCGatewayMetaData.Bin instead.
var L2USDCGatewayBin = L2USDCGatewayMetaData.Bin

// DeployL2USDCGateway deploys a new Ethereum contract, binding an instance of L2USDCGateway to it.
func DeployL2USDCGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _l1USDC common.Address, _l2USDC common.Address) (common.Address, *types.Transaction, *L2USDCGateway, error) {
	parsed, err := L2USDCGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2USDCGatewayBin), backend, _l1USDC, _l2USDC)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2USDCGateway{L2USDCGatewayCaller: L2USDCGatewayCaller{contract: contract}, L2USDCGatewayTransactor: L2USDCGatewayTransactor{contract: contract}, L2USDCGatewayFilterer: L2USDCGatewayFilterer{contract: contract}}, nil
}

// L2USDCGateway is an auto generated Go binding around an Ethereum contract.
type L2USDCGateway struct {
	L2USDCGatewayCaller     // Read-only binding to the contract
	L2USDCGatewayTransactor // Write-only binding to the contract
	L2USDCGatewayFilterer   // Log filterer for contract events
}

// L2USDCGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2USDCGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2USDCGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2USDCGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2USDCGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2USDCGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2USDCGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2USDCGatewaySession struct {
	Contract     *L2USDCGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2USDCGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2USDCGatewayCallerSession struct {
	Contract *L2USDCGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2USDCGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2USDCGatewayTransactorSession struct {
	Contract     *L2USDCGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2USDCGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2USDCGatewayRaw struct {
	Contract *L2USDCGateway // Generic contract binding to access the raw methods on
}

// L2USDCGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2USDCGatewayCallerRaw struct {
	Contract *L2USDCGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2USDCGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2USDCGatewayTransactorRaw struct {
	Contract *L2USDCGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2USDCGateway creates a new instance of L2USDCGateway, bound to a specific deployed contract.
func NewL2USDCGateway(address common.Address, backend bind.ContractBackend) (*L2USDCGateway, error) {
	contract, err := bindL2USDCGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2USDCGateway{L2USDCGatewayCaller: L2USDCGatewayCaller{contract: contract}, L2USDCGatewayTransactor: L2USDCGatewayTransactor{contract: contract}, L2USDCGatewayFilterer: L2USDCGatewayFilterer{contract: contract}}, nil
}

// NewL2USDCGatewayCaller creates a new read-only instance of L2USDCGateway, bound to a specific deployed contract.
func NewL2USDCGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2USDCGatewayCaller, error) {
	contract, err := bindL2USDCGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayCaller{contract: contract}, nil
}

// NewL2USDCGatewayTransactor creates a new write-only instance of L2USDCGateway, bound to a specific deployed contract.
func NewL2USDCGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2USDCGatewayTransactor, error) {
	contract, err := bindL2USDCGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayTransactor{contract: contract}, nil
}

// NewL2USDCGatewayFilterer creates a new log filterer instance of L2USDCGateway, bound to a specific deployed contract.
func NewL2USDCGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2USDCGatewayFilterer, error) {
	contract, err := bindL2USDCGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayFilterer{contract: contract}, nil
}

// bindL2USDCGateway binds a generic wrapper to an already deployed contract.
func bindL2USDCGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2USDCGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2USDCGateway *L2USDCGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2USDCGateway.Contract.L2USDCGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2USDCGateway *L2USDCGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.L2USDCGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2USDCGateway *L2USDCGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.L2USDCGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2USDCGateway *L2USDCGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2USDCGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2USDCGateway *L2USDCGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2USDCGateway *L2USDCGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.contract.Transact(opts, method, params...)
}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) CircleCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "circleCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) CircleCaller() (common.Address, error) {
	return _L2USDCGateway.Contract.CircleCaller(&_L2USDCGateway.CallOpts)
}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) CircleCaller() (common.Address, error) {
	return _L2USDCGateway.Contract.CircleCaller(&_L2USDCGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) Counterpart() (common.Address, error) {
	return _L2USDCGateway.Contract.Counterpart(&_L2USDCGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2USDCGateway.Contract.Counterpart(&_L2USDCGateway.CallOpts)
}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewayCaller) DepositPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "depositPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewaySession) DepositPaused() (bool, error) {
	return _L2USDCGateway.Contract.DepositPaused(&_L2USDCGateway.CallOpts)
}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewayCallerSession) DepositPaused() (bool, error) {
	return _L2USDCGateway.Contract.DepositPaused(&_L2USDCGateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "getL1ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2USDCGateway.Contract.GetL1ERC20Address(&_L2USDCGateway.CallOpts, arg0)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2USDCGateway.Contract.GetL1ERC20Address(&_L2USDCGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2USDCGateway.Contract.GetL2ERC20Address(&_L2USDCGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2USDCGateway.Contract.GetL2ERC20Address(&_L2USDCGateway.CallOpts, arg0)
}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) L1USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "l1USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) L1USDC() (common.Address, error) {
	return _L2USDCGateway.Contract.L1USDC(&_L2USDCGateway.CallOpts)
}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) L1USDC() (common.Address, error) {
	return _L2USDCGateway.Contract.L1USDC(&_L2USDCGateway.CallOpts)
}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) L2USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "l2USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) L2USDC() (common.Address, error) {
	return _L2USDCGateway.Contract.L2USDC(&_L2USDCGateway.CallOpts)
}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) L2USDC() (common.Address, error) {
	return _L2USDCGateway.Contract.L2USDC(&_L2USDCGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) Messenger() (common.Address, error) {
	return _L2USDCGateway.Contract.Messenger(&_L2USDCGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2USDCGateway.Contract.Messenger(&_L2USDCGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) Owner() (common.Address, error) {
	return _L2USDCGateway.Contract.Owner(&_L2USDCGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) Owner() (common.Address, error) {
	return _L2USDCGateway.Contract.Owner(&_L2USDCGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2USDCGateway *L2USDCGatewaySession) Router() (common.Address, error) {
	return _L2USDCGateway.Contract.Router(&_L2USDCGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2USDCGateway *L2USDCGatewayCallerSession) Router() (common.Address, error) {
	return _L2USDCGateway.Contract.Router(&_L2USDCGateway.CallOpts)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewayCaller) WithdrawPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2USDCGateway.contract.Call(opts, &out, "withdrawPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewaySession) WithdrawPaused() (bool, error) {
	return _L2USDCGateway.Contract.WithdrawPaused(&_L2USDCGateway.CallOpts)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L2USDCGateway *L2USDCGatewayCallerSession) WithdrawPaused() (bool, error) {
	return _L2USDCGateway.Contract.WithdrawPaused(&_L2USDCGateway.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2USDCGateway *L2USDCGatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.FinalizeDepositERC20(&_L2USDCGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.FinalizeDepositERC20(&_L2USDCGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2USDCGateway *L2USDCGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.Initialize(&_L2USDCGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.Initialize(&_L2USDCGateway.TransactOpts, _counterpart, _router, _messenger)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) PauseDeposit(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "pauseDeposit", _paused)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewaySession) PauseDeposit(_paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.PauseDeposit(&_L2USDCGateway.TransactOpts, _paused)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) PauseDeposit(_paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.PauseDeposit(&_L2USDCGateway.TransactOpts, _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) PauseWithdraw(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "pauseWithdraw", _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewaySession) PauseWithdraw(_paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.PauseWithdraw(&_L2USDCGateway.TransactOpts, _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) PauseWithdraw(_paused bool) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.PauseWithdraw(&_L2USDCGateway.TransactOpts, _paused)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2USDCGateway *L2USDCGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2USDCGateway.Contract.RenounceOwnership(&_L2USDCGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2USDCGateway.Contract.RenounceOwnership(&_L2USDCGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2USDCGateway *L2USDCGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.TransferOwnership(&_L2USDCGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.TransferOwnership(&_L2USDCGateway.TransactOpts, newOwner)
}

// TransferUSDCRoles is a paid mutator transaction binding the contract method 0xc689fc34.
//
// Solidity: function transferUSDCRoles(address _owner) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) TransferUSDCRoles(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "transferUSDCRoles", _owner)
}

// TransferUSDCRoles is a paid mutator transaction binding the contract method 0xc689fc34.
//
// Solidity: function transferUSDCRoles(address _owner) returns()
func (_L2USDCGateway *L2USDCGatewaySession) TransferUSDCRoles(_owner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.TransferUSDCRoles(&_L2USDCGateway.TransactOpts, _owner)
}

// TransferUSDCRoles is a paid mutator transaction binding the contract method 0xc689fc34.
//
// Solidity: function transferUSDCRoles(address _owner) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) TransferUSDCRoles(_owner common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.TransferUSDCRoles(&_L2USDCGateway.TransactOpts, _owner)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) UpdateCircleCaller(opts *bind.TransactOpts, _caller common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "updateCircleCaller", _caller)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L2USDCGateway *L2USDCGatewaySession) UpdateCircleCaller(_caller common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.UpdateCircleCaller(&_L2USDCGateway.TransactOpts, _caller)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) UpdateCircleCaller(_caller common.Address) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.UpdateCircleCaller(&_L2USDCGateway.TransactOpts, _caller)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC20(&_L2USDCGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC20(&_L2USDCGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC200(&_L2USDCGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC200(&_L2USDCGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC20AndCall(&_L2USDCGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2USDCGateway *L2USDCGatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2USDCGateway.Contract.WithdrawERC20AndCall(&_L2USDCGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2USDCGatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2USDCGateway contract.
type L2USDCGatewayFinalizeDepositERC20Iterator struct {
	Event *L2USDCGatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2USDCGatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2USDCGatewayFinalizeDepositERC20)
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
		it.Event = new(L2USDCGatewayFinalizeDepositERC20)
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
func (it *L2USDCGatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2USDCGatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2USDCGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2USDCGateway contract.
type L2USDCGatewayFinalizeDepositERC20 struct {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2USDCGatewayFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2USDCGateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayFinalizeDepositERC20Iterator{contract: _L2USDCGateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2USDCGateway *L2USDCGatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2USDCGatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2USDCGateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2USDCGatewayFinalizeDepositERC20)
				if err := _L2USDCGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2USDCGatewayFinalizeDepositERC20, error) {
	event := new(L2USDCGatewayFinalizeDepositERC20)
	if err := _L2USDCGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2USDCGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2USDCGateway contract.
type L2USDCGatewayInitializedIterator struct {
	Event *L2USDCGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2USDCGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2USDCGatewayInitialized)
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
		it.Event = new(L2USDCGatewayInitialized)
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
func (it *L2USDCGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2USDCGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2USDCGatewayInitialized represents a Initialized event raised by the L2USDCGateway contract.
type L2USDCGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2USDCGateway *L2USDCGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2USDCGatewayInitializedIterator, error) {

	logs, sub, err := _L2USDCGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayInitializedIterator{contract: _L2USDCGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2USDCGateway *L2USDCGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2USDCGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2USDCGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2USDCGatewayInitialized)
				if err := _L2USDCGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) ParseInitialized(log types.Log) (*L2USDCGatewayInitialized, error) {
	event := new(L2USDCGatewayInitialized)
	if err := _L2USDCGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2USDCGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2USDCGateway contract.
type L2USDCGatewayOwnershipTransferredIterator struct {
	Event *L2USDCGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2USDCGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2USDCGatewayOwnershipTransferred)
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
		it.Event = new(L2USDCGatewayOwnershipTransferred)
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
func (it *L2USDCGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2USDCGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2USDCGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2USDCGateway contract.
type L2USDCGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2USDCGateway *L2USDCGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2USDCGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2USDCGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayOwnershipTransferredIterator{contract: _L2USDCGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2USDCGateway *L2USDCGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2USDCGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2USDCGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2USDCGatewayOwnershipTransferred)
				if err := _L2USDCGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2USDCGatewayOwnershipTransferred, error) {
	event := new(L2USDCGatewayOwnershipTransferred)
	if err := _L2USDCGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2USDCGatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2USDCGateway contract.
type L2USDCGatewayWithdrawERC20Iterator struct {
	Event *L2USDCGatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2USDCGatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2USDCGatewayWithdrawERC20)
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
		it.Event = new(L2USDCGatewayWithdrawERC20)
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
func (it *L2USDCGatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2USDCGatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2USDCGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2USDCGateway contract.
type L2USDCGatewayWithdrawERC20 struct {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2USDCGatewayWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2USDCGateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2USDCGatewayWithdrawERC20Iterator{contract: _L2USDCGateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2USDCGateway *L2USDCGatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2USDCGatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2USDCGateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2USDCGatewayWithdrawERC20)
				if err := _L2USDCGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2USDCGateway *L2USDCGatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2USDCGatewayWithdrawERC20, error) {
	event := new(L2USDCGatewayWithdrawERC20)
	if err := _L2USDCGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
