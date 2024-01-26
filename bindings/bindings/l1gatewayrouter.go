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

// L1GatewayRouterMetaData contains all meta data concerning the L1GatewayRouter contract.
var L1GatewayRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDefaultERC20Gateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"SetDefaultERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGateway\",\"type\":\"address\"}],\"name\":\"SetERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldETHGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEthGateway\",\"type\":\"address\"}],\"name\":\"SetETHGateway\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayInContext\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Address\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"setDefaultERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateways\",\"type\":\"address[]\"}],\"name\":\"setERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEthGateway\",\"type\":\"address\"}],\"name\":\"setETHGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611d35806100ec6000396000f3fe6080604052600436106101755760003560e01c80638c00ce73116100cb578063c52a3bbc1161007f578063ce8c3e0611610059578063ce8c3e06146103b0578063f219fa66146103d0578063f2fde38b146103e357600080fd5b8063c52a3bbc1461034f578063c676ad291461037d578063ce0b63ce1461039d57600080fd5b80638eaac8a3116100b05780638eaac8a31461031b5780639f8420b314610329578063aac476f81461033c57600080fd5b80638c00ce73146102dd5780638da5cb5b146102fd57600080fd5b8063485cc9551161012d578063705b05b811610107578063705b05b81461027f578063715018a6146102b557806384bd13b0146102ca57600080fd5b8063485cc9551461021f5780635dfd5b9a1461023f578063635c86371461025f57600080fd5b80633a9a7b201161015e5780633a9a7b20146101a25780633d1d31c7146101df57806343c66741146101ff57600080fd5b80630aea8c261461017a57806321425ee01461018f575b600080fd5b61018d610188366004611683565b610403565b005b61018d61019d3660046116f7565b6105ec565b3480156101ae57600080fd5b506068546101c2906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101eb57600080fd5b5061018d6101fa36600461172c565b61062b565b34801561020b57600080fd5b506101c261021a36600461172c565b61069d565b34801561022b57600080fd5b5061018d61023a366004611749565b6106d3565b34801561024b57600080fd5b5061018d61025a36600461172c565b61092e565b34801561026b57600080fd5b5061018d61027a366004611800565b6109a0565b34801561028b57600080fd5b506101c261029a36600461172c565b6067602052600090815260409020546001600160a01b031681565b3480156102c157600080fd5b5061018d610b66565b61018d6102d83660046118ad565b610b7a565b3480156102e957600080fd5b506065546101c2906001600160a01b031681565b34801561030957600080fd5b506033546001600160a01b03166101c2565b61018d6102d8366004611945565b61018d6103373660046119b8565b610bc2565b61018d61034a3660046119da565b610bff565b34801561035b57600080fd5b5061036f61036a366004611a3b565b610d7b565b6040519081526020016101d6565b34801561038957600080fd5b506101c261039836600461172c565b610f1e565b61018d6103ab3660046116f7565b610fcd565b3480156103bc57600080fd5b506066546101c2906001600160a01b031681565b61018d6103de366004611a7c565b610fd9565b3480156103ef57600080fd5b5061018d6103fe36600461172c565b610fec565b6068546001600160a01b0316156104615760405162461bcd60e51b815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e746578740000000000000000000000000060448201526064015b60405180910390fd5b600061046c8661069d565b90506001600160a01b0381166104c45760405162461bcd60e51b815260206004820152601460248201527f6e6f206761746577617920617661696c61626c650000000000000000000000006044820152606401610458565b606880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790556000338460405160200161050c929190611b30565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f0aea8c2600000000000000000000000000000000000000000000000000000000825291506001600160a01b03831690630aea8c26903490610589908b908b908b9088908b90600401611b52565b6000604051808303818588803b1580156105a257600080fd5b505af11580156105b6573d6000803e3d6000fd5b5050606880547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550505050505050505050565b61062683338460005b6040519080825280601f01601f19166020018201604052801561061f576020820181803683370190505b5085610403565b505050565b61063361107c565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a90600090a35050565b6001600160a01b03808216600090815260676020526040812054909116806106cd57506066546001600160a01b03165b92915050565b600054610100900460ff16158080156106f35750600054600160ff909116105b8061070d5750303b15801561070d575060005460ff166001145b61077f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610458565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107dd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6107e56110d6565b6001600160a01b0382161561085657606680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040516000907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1908290a35b6001600160a01b038316156108c757606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385169081179091556040516000907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a908290a35b801561062657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b61093661107c565b606680546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d190600090a35050565b6109a861107c565b80518251146109f95760405162461bcd60e51b815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610458565b60005b825181101561062657600060676000858481518110610a1d57610a1d611b96565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050828281518110610a6d57610a6d611b96565b602002602001015160676000868581518110610a8b57610a8b611b96565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550828281518110610ae957610ae9611b96565b60200260200101516001600160a01b0316816001600160a01b0316858481518110610b1657610b16611b96565b60200260200101516001600160a01b03167f0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf60405160405180910390a45080610b5e81611bf4565b9150506109fc565b610b6e61107c565b610b78600061115b565b565b60405162461bcd60e51b815260206004820152601660248201527f73686f756c64206e657665722062652063616c6c6564000000000000000000006044820152606401610458565b610bfb338360005b6040519080825280601f01601f191660200182016040528015610bf4576020820181803683370190505b5084610bff565b5050565b6068546001600160a01b031615610c585760405162461bcd60e51b815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e74657874000000000000000000000000006044820152606401610458565b6065546001600160a01b031680610cb15760405162461bcd60e51b815260206004820152601560248201527f657468206761746577617920617661696c61626c6500000000000000000000006044820152606401610458565b60003384604051602001610cc6929190611b30565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527faac476f800000000000000000000000000000000000000000000000000000000825291506001600160a01b0383169063aac476f8903490610d41908a908a9087908a90600401611c2c565b6000604051808303818588803b158015610d5a57600080fd5b505af1158015610d6e573d6000803e3d6000fd5b5050505050505050505050565b6068546000906001600160a01b0316336001600160a01b031614610de15760405162461bcd60e51b815260206004820152601760248201527f4f6e6c7920696e206465706f73697420636f6e746578740000000000000000006044820152606401610458565b6000336040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000918616906370a0823190602401602060405180830381865afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d9190611c65565b9050610e846001600160a01b0386168784876111c5565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301528291908716906370a0823190602401602060405180830381865afa158015610ee6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0a9190611c65565b610f149190611c7e565b9695505050505050565b600080610f2a8361069d565b90506001600160a01b038116610f435750600092915050565b6040517fc676ad290000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015282169063c676ad2990602401602060405180830381865afa158015610fa2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc69190611c91565b9392505050565b61062683836000610bca565b610fe684848460006105f5565b50505050565b610ff461107c565b6001600160a01b0381166110705760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610458565b6110798161115b565b50565b6033546001600160a01b03163314610b785760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610458565b600054610100900460ff166111535760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610458565b610b7861124d565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610fe69085906112d3565b600054610100900460ff166112ca5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610458565b610b783361115b565b6000611328826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113bb9092919063ffffffff16565b90508051600014806113495750808060200190518101906113499190611cae565b6106265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610458565b60606113ca84846000856113d2565b949350505050565b60608247101561144a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610458565b600080866001600160a01b031685876040516114669190611cd0565b60006040518083038185875af1925050503d80600081146114a3576040519150601f19603f3d011682016040523d82523d6000602084013e6114a8565b606091505b50915091506114b9878383876114c4565b979650505050505050565b6060831561153357825160000361152c576001600160a01b0385163b61152c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610458565b50816113ca565b6113ca83838151156115485781518083602001fd5b8060405162461bcd60e51b81526004016104589190611cec565b6001600160a01b038116811461107957600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156115ed576115ed611577565b604052919050565b600082601f83011261160657600080fd5b813567ffffffffffffffff81111561162057611620611577565b61165160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016115a6565b81815284602083860101111561166657600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561169b57600080fd5b85356116a681611562565b945060208601356116b681611562565b935060408601359250606086013567ffffffffffffffff8111156116d957600080fd5b6116e5888289016115f5565b95989497509295608001359392505050565b60008060006060848603121561170c57600080fd5b833561171781611562565b95602085013595506040909401359392505050565b60006020828403121561173e57600080fd5b8135610fc681611562565b6000806040838503121561175c57600080fd5b823561176781611562565b9150602083013561177781611562565b809150509250929050565b600082601f83011261179357600080fd5b8135602067ffffffffffffffff8211156117af576117af611577565b8160051b6117be8282016115a6565b92835284810182019282810190878511156117d857600080fd5b83870192505b848310156114b95782356117f181611562565b825291830191908301906117de565b6000806040838503121561181357600080fd5b823567ffffffffffffffff8082111561182b57600080fd5b61183786838701611782565b9350602085013591508082111561184d57600080fd5b5061185a85828601611782565b9150509250929050565b60008083601f84011261187657600080fd5b50813567ffffffffffffffff81111561188e57600080fd5b6020830191508360208285010111156118a657600080fd5b9250929050565b600080600080600080600060c0888a0312156118c857600080fd5b87356118d381611562565b965060208801356118e381611562565b955060408801356118f381611562565b9450606088013561190381611562565b93506080880135925060a088013567ffffffffffffffff81111561192657600080fd5b6119328a828b01611864565b989b979a50959850939692959293505050565b60008060008060006080868803121561195d57600080fd5b853561196881611562565b9450602086013561197881611562565b935060408601359250606086013567ffffffffffffffff81111561199b57600080fd5b6119a788828901611864565b969995985093965092949392505050565b600080604083850312156119cb57600080fd5b50508035926020909101359150565b600080600080608085870312156119f057600080fd5b84356119fb81611562565b935060208501359250604085013567ffffffffffffffff811115611a1e57600080fd5b611a2a878288016115f5565b949793965093946060013593505050565b600080600060608486031215611a5057600080fd5b8335611a5b81611562565b92506020840135611a6b81611562565b929592945050506040919091013590565b60008060008060808587031215611a9257600080fd5b8435611a9d81611562565b93506020850135611aad81611562565b93969395505050506040820135916060013590565b60005b83811015611add578181015183820152602001611ac5565b50506000910152565b60008151808452611afe816020860160208601611ac2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6001600160a01b03831681526040602082015260006113ca6040830184611ae6565b60006001600160a01b03808816835280871660208401525084604083015260a06060830152611b8460a0830185611ae6565b90508260808301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c2557611c25611bc5565b5060010190565b6001600160a01b0385168152836020820152608060408201526000611c546080830185611ae6565b905082606083015295945050505050565b600060208284031215611c7757600080fd5b5051919050565b818103818111156106cd576106cd611bc5565b600060208284031215611ca357600080fd5b8151610fc681611562565b600060208284031215611cc057600080fd5b81518015158114610fc657600080fd5b60008251611ce2818460208701611ac2565b9190910192915050565b602081526000610fc66020830184611ae656fea2646970667358221220a2c4c5ed5cc6833246225ec99b7b69f961066384b2efcfa1a1d55342d7cc4bbe64736f6c63430008100033",
}

// L1GatewayRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use L1GatewayRouterMetaData.ABI instead.
var L1GatewayRouterABI = L1GatewayRouterMetaData.ABI

// L1GatewayRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1GatewayRouterMetaData.Bin instead.
var L1GatewayRouterBin = L1GatewayRouterMetaData.Bin

// DeployL1GatewayRouter deploys a new Ethereum contract, binding an instance of L1GatewayRouter to it.
func DeployL1GatewayRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1GatewayRouter, error) {
	parsed, err := L1GatewayRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1GatewayRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1GatewayRouter{L1GatewayRouterCaller: L1GatewayRouterCaller{contract: contract}, L1GatewayRouterTransactor: L1GatewayRouterTransactor{contract: contract}, L1GatewayRouterFilterer: L1GatewayRouterFilterer{contract: contract}}, nil
}

// L1GatewayRouter is an auto generated Go binding around an Ethereum contract.
type L1GatewayRouter struct {
	L1GatewayRouterCaller     // Read-only binding to the contract
	L1GatewayRouterTransactor // Write-only binding to the contract
	L1GatewayRouterFilterer   // Log filterer for contract events
}

// L1GatewayRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1GatewayRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1GatewayRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1GatewayRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1GatewayRouterSession struct {
	Contract     *L1GatewayRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1GatewayRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1GatewayRouterCallerSession struct {
	Contract *L1GatewayRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1GatewayRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1GatewayRouterTransactorSession struct {
	Contract     *L1GatewayRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1GatewayRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1GatewayRouterRaw struct {
	Contract *L1GatewayRouter // Generic contract binding to access the raw methods on
}

// L1GatewayRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1GatewayRouterCallerRaw struct {
	Contract *L1GatewayRouterCaller // Generic read-only contract binding to access the raw methods on
}

// L1GatewayRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1GatewayRouterTransactorRaw struct {
	Contract *L1GatewayRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1GatewayRouter creates a new instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouter(address common.Address, backend bind.ContractBackend) (*L1GatewayRouter, error) {
	contract, err := bindL1GatewayRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouter{L1GatewayRouterCaller: L1GatewayRouterCaller{contract: contract}, L1GatewayRouterTransactor: L1GatewayRouterTransactor{contract: contract}, L1GatewayRouterFilterer: L1GatewayRouterFilterer{contract: contract}}, nil
}

// NewL1GatewayRouterCaller creates a new read-only instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterCaller(address common.Address, caller bind.ContractCaller) (*L1GatewayRouterCaller, error) {
	contract, err := bindL1GatewayRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterCaller{contract: contract}, nil
}

// NewL1GatewayRouterTransactor creates a new write-only instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*L1GatewayRouterTransactor, error) {
	contract, err := bindL1GatewayRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterTransactor{contract: contract}, nil
}

// NewL1GatewayRouterFilterer creates a new log filterer instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*L1GatewayRouterFilterer, error) {
	contract, err := bindL1GatewayRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterFilterer{contract: contract}, nil
}

// bindL1GatewayRouter binds a generic wrapper to an already deployed contract.
func bindL1GatewayRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1GatewayRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GatewayRouter *L1GatewayRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GatewayRouter.Contract.L1GatewayRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GatewayRouter *L1GatewayRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.L1GatewayRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GatewayRouter *L1GatewayRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.L1GatewayRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GatewayRouter *L1GatewayRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GatewayRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GatewayRouter *L1GatewayRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GatewayRouter *L1GatewayRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.ERC20Gateway(&_L1GatewayRouter.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.ERC20Gateway(&_L1GatewayRouter.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) DefaultERC20Gateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.DefaultERC20Gateway(&_L1GatewayRouter.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.DefaultERC20Gateway(&_L1GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) EthGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.EthGateway(&_L1GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) EthGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.EthGateway(&_L1GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) GatewayInContext(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "gatewayInContext")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) GatewayInContext() (common.Address, error) {
	return _L1GatewayRouter.Contract.GatewayInContext(&_L1GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) GatewayInContext() (common.Address, error) {
	return _L1GatewayRouter.Contract.GatewayInContext(&_L1GatewayRouter.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetERC20Gateway(&_L1GatewayRouter.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetERC20Gateway(&_L1GatewayRouter.CallOpts, _token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "getL2ERC20Address", _l1Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) GetL2ERC20Address(_l1Address common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetL2ERC20Address(&_L1GatewayRouter.CallOpts, _l1Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) GetL2ERC20Address(_l1Address common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetL2ERC20Address(&_L1GatewayRouter.CallOpts, _l1Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) Owner() (common.Address, error) {
	return _L1GatewayRouter.Contract.Owner(&_L1GatewayRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) Owner() (common.Address, error) {
	return _L1GatewayRouter.Contract.Owner(&_L1GatewayRouter.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC20(&_L1GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC20(&_L1GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC200(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC200(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC20AndCall(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositERC20AndCall(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositETH(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositETH", _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETH(&_L1GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETH(&_L1GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositETH0(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositETH0", _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETH0(&_L1GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETH0(&_L1GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) DepositETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "depositETHAndCall", _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETHAndCall(&_L1GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.DepositETHAndCall(&_L1GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeWithdrawERC20(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeWithdrawERC20(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "finalizeWithdrawETH", arg0, arg1, arg2, arg3)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) FinalizeWithdrawETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeWithdrawETH(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) FinalizeWithdrawETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeWithdrawETH(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.Initialize(&_L1GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.Initialize(&_L1GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GatewayRouter *L1GatewayRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.RenounceOwnership(&_L1GatewayRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.RenounceOwnership(&_L1GatewayRouter.TransactOpts)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactor) RequestERC20(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "requestERC20", _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.RequestERC20(&_L1GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.RequestERC20(&_L1GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setDefaultERC20Gateway", _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetDefaultERC20Gateway(&_L1GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetDefaultERC20Gateway(&_L1GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetERC20Gateway(&_L1GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetERC20Gateway(&_L1GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetETHGateway(opts *bind.TransactOpts, _newEthGateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setETHGateway", _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetETHGateway(&_L1GatewayRouter.TransactOpts, _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetETHGateway(&_L1GatewayRouter.TransactOpts, _newEthGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.TransferOwnership(&_L1GatewayRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.TransferOwnership(&_L1GatewayRouter.TransactOpts, newOwner)
}

// L1GatewayRouterDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1GatewayRouter contract.
type L1GatewayRouterDepositERC20Iterator struct {
	Event *L1GatewayRouterDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterDepositERC20)
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
		it.Event = new(L1GatewayRouterDepositERC20)
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
func (it *L1GatewayRouterDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterDepositERC20 represents a DepositERC20 event raised by the L1GatewayRouter contract.
type L1GatewayRouterDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1GatewayRouterDepositERC20Iterator, error) {

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

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterDepositERC20Iterator{contract: _L1GatewayRouter.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterDepositERC20)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseDepositERC20(log types.Log) (*L1GatewayRouterDepositERC20, error) {
	event := new(L1GatewayRouterDepositERC20)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the L1GatewayRouter contract.
type L1GatewayRouterDepositETHIterator struct {
	Event *L1GatewayRouterDepositETH // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterDepositETH)
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
		it.Event = new(L1GatewayRouterDepositETH)
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
func (it *L1GatewayRouterDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterDepositETH represents a DepositETH event raised by the L1GatewayRouter contract.
type L1GatewayRouterDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1GatewayRouterDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterDepositETHIterator{contract: _L1GatewayRouter.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterDepositETH)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseDepositETH(log types.Log) (*L1GatewayRouterDepositETH, error) {
	event := new(L1GatewayRouterDepositETH)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1GatewayRouter contract.
type L1GatewayRouterFinalizeWithdrawERC20Iterator struct {
	Event *L1GatewayRouterFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterFinalizeWithdrawERC20)
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
		it.Event = new(L1GatewayRouterFinalizeWithdrawERC20)
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
func (it *L1GatewayRouterFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1GatewayRouter contract.
type L1GatewayRouterFinalizeWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC20 is a free log retrieval operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1GatewayRouterFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterFinalizeWithdrawERC20Iterator{contract: _L1GatewayRouter.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterFinalizeWithdrawERC20)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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

// ParseFinalizeWithdrawERC20 is a log parse operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1GatewayRouterFinalizeWithdrawERC20, error) {
	event := new(L1GatewayRouterFinalizeWithdrawERC20)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterFinalizeWithdrawETHIterator is returned from FilterFinalizeWithdrawETH and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawETH events raised by the L1GatewayRouter contract.
type L1GatewayRouterFinalizeWithdrawETHIterator struct {
	Event *L1GatewayRouterFinalizeWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterFinalizeWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterFinalizeWithdrawETH)
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
		it.Event = new(L1GatewayRouterFinalizeWithdrawETH)
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
func (it *L1GatewayRouterFinalizeWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterFinalizeWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the L1GatewayRouter contract.
type L1GatewayRouterFinalizeWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawETH is a free log retrieval operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterFinalizeWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1GatewayRouterFinalizeWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterFinalizeWithdrawETHIterator{contract: _L1GatewayRouter.contract, event: "FinalizeWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawETH is a free log subscription operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchFinalizeWithdrawETH(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterFinalizeWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterFinalizeWithdrawETH)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
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

// ParseFinalizeWithdrawETH is a log parse operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseFinalizeWithdrawETH(log types.Log) (*L1GatewayRouterFinalizeWithdrawETH, error) {
	event := new(L1GatewayRouterFinalizeWithdrawETH)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1GatewayRouter contract.
type L1GatewayRouterInitializedIterator struct {
	Event *L1GatewayRouterInitialized // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterInitialized)
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
		it.Event = new(L1GatewayRouterInitialized)
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
func (it *L1GatewayRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterInitialized represents a Initialized event raised by the L1GatewayRouter contract.
type L1GatewayRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1GatewayRouterInitializedIterator, error) {

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterInitializedIterator{contract: _L1GatewayRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterInitialized)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseInitialized(log types.Log) (*L1GatewayRouterInitialized, error) {
	event := new(L1GatewayRouterInitialized)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1GatewayRouter contract.
type L1GatewayRouterOwnershipTransferredIterator struct {
	Event *L1GatewayRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterOwnershipTransferred)
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
		it.Event = new(L1GatewayRouterOwnershipTransferred)
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
func (it *L1GatewayRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterOwnershipTransferred represents a OwnershipTransferred event raised by the L1GatewayRouter contract.
type L1GatewayRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1GatewayRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterOwnershipTransferredIterator{contract: _L1GatewayRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterOwnershipTransferred)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseOwnershipTransferred(log types.Log) (*L1GatewayRouterOwnershipTransferred, error) {
	event := new(L1GatewayRouterOwnershipTransferred)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the L1GatewayRouter contract.
type L1GatewayRouterRefundERC20Iterator struct {
	Event *L1GatewayRouterRefundERC20 // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterRefundERC20)
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
		it.Event = new(L1GatewayRouterRefundERC20)
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
func (it *L1GatewayRouterRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterRefundERC20 represents a RefundERC20 event raised by the L1GatewayRouter contract.
type L1GatewayRouterRefundERC20 struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterRefundERC20(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1GatewayRouterRefundERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterRefundERC20Iterator{contract: _L1GatewayRouter.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterRefundERC20, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterRefundERC20)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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

// ParseRefundERC20 is a log parse operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseRefundERC20(log types.Log) (*L1GatewayRouterRefundERC20, error) {
	event := new(L1GatewayRouterRefundERC20)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterRefundETHIterator is returned from FilterRefundETH and is used to iterate over the raw logs and unpacked data for RefundETH events raised by the L1GatewayRouter contract.
type L1GatewayRouterRefundETHIterator struct {
	Event *L1GatewayRouterRefundETH // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterRefundETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterRefundETH)
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
		it.Event = new(L1GatewayRouterRefundETH)
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
func (it *L1GatewayRouterRefundETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterRefundETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterRefundETH represents a RefundETH event raised by the L1GatewayRouter contract.
type L1GatewayRouterRefundETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundETH is a free log retrieval operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterRefundETH(opts *bind.FilterOpts, recipient []common.Address) (*L1GatewayRouterRefundETHIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterRefundETHIterator{contract: _L1GatewayRouter.contract, event: "RefundETH", logs: logs, sub: sub}, nil
}

// WatchRefundETH is a free log subscription operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchRefundETH(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterRefundETH, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterRefundETH)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "RefundETH", log); err != nil {
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

// ParseRefundETH is a log parse operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseRefundETH(log types.Log) (*L1GatewayRouterRefundETH, error) {
	event := new(L1GatewayRouterRefundETH)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "RefundETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterSetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the L1GatewayRouter contract.
type L1GatewayRouterSetDefaultERC20GatewayIterator struct {
	Event *L1GatewayRouterSetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterSetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterSetDefaultERC20Gateway)
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
		it.Event = new(L1GatewayRouterSetDefaultERC20Gateway)
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
func (it *L1GatewayRouterSetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterSetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterSetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the L1GatewayRouter contract.
type L1GatewayRouterSetDefaultERC20Gateway struct {
	OldDefaultERC20Gateway common.Address
	NewDefaultERC20Gateway common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (*L1GatewayRouterSetDefaultERC20GatewayIterator, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterSetDefaultERC20GatewayIterator{contract: _L1GatewayRouter.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterSetDefaultERC20Gateway, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (event.Subscription, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterSetDefaultERC20Gateway)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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

// ParseSetDefaultERC20Gateway is a log parse operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*L1GatewayRouterSetDefaultERC20Gateway, error) {
	event := new(L1GatewayRouterSetDefaultERC20Gateway)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterSetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the L1GatewayRouter contract.
type L1GatewayRouterSetERC20GatewayIterator struct {
	Event *L1GatewayRouterSetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterSetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterSetERC20Gateway)
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
		it.Event = new(L1GatewayRouterSetERC20Gateway)
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
func (it *L1GatewayRouterSetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterSetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterSetERC20Gateway represents a SetERC20Gateway event raised by the L1GatewayRouter contract.
type L1GatewayRouterSetERC20Gateway struct {
	Token      common.Address
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (*L1GatewayRouterSetERC20GatewayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterSetERC20GatewayIterator{contract: _L1GatewayRouter.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterSetERC20Gateway, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterSetERC20Gateway)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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

// ParseSetERC20Gateway is a log parse operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseSetERC20Gateway(log types.Log) (*L1GatewayRouterSetERC20Gateway, error) {
	event := new(L1GatewayRouterSetERC20Gateway)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterSetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the L1GatewayRouter contract.
type L1GatewayRouterSetETHGatewayIterator struct {
	Event *L1GatewayRouterSetETHGateway // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterSetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterSetETHGateway)
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
		it.Event = new(L1GatewayRouterSetETHGateway)
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
func (it *L1GatewayRouterSetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterSetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterSetETHGateway represents a SetETHGateway event raised by the L1GatewayRouter contract.
type L1GatewayRouterSetETHGateway struct {
	OldETHGateway common.Address
	NewEthGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterSetETHGateway(opts *bind.FilterOpts, oldETHGateway []common.Address, newEthGateway []common.Address) (*L1GatewayRouterSetETHGatewayIterator, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterSetETHGatewayIterator{contract: _L1GatewayRouter.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterSetETHGateway, oldETHGateway []common.Address, newEthGateway []common.Address) (event.Subscription, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterSetETHGateway)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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

// ParseSetETHGateway is a log parse operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseSetETHGateway(log types.Log) (*L1GatewayRouterSetETHGateway, error) {
	event := new(L1GatewayRouterSetETHGateway)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
