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

// L2GatewayRouterMetaData contains all meta data concerning the L2GatewayRouter contract.
var L2GatewayRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDefaultERC20Gateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"SetDefaultERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGateway\",\"type\":\"address\"}],\"name\":\"SetERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldETHGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEthGateway\",\"type\":\"address\"}],\"name\":\"SetETHGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayInContext\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Address\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"setDefaultERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateways\",\"type\":\"address[]\"}],\"name\":\"setERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEthGateway\",\"type\":\"address\"}],\"name\":\"setETHGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611cf8806100e65f395ff3fe608060405260043610610178575f3560e01c80636dc24183116100d1578063a93a4af91161007c578063c7cdea3711610057578063c7cdea37146103c4578063ce8c3e06146103d7578063f2fde38b146103f6575f80fd5b8063a93a4af914610365578063c52a3bbc14610378578063c676ad29146103a5575f80fd5b80638431f5c1116100ac5780638431f5c11461031b5780638c00ce73146103295780638da5cb5b14610348575f80fd5b80636dc24183146102c0578063705b05b8146102d3578063715018a614610307575f80fd5b8063485cc955116101315780635dfd5b9a1161010c5780635dfd5b9a1461026f578063635c86371461028e5780636c07ea43146102ad575f80fd5b8063485cc9551461021e57806354bbd59c1461023d578063575361b61461025c575f80fd5b80633a9a7b20116101615780633a9a7b20146101a45780633d1d31c7146101e057806343c66741146101ff575f80fd5b8063232e87481461017c5780632fcc29fa14610191575b5f80fd5b61018f61018a36600461162f565b610415565b005b61018f61019f36600461169d565b610462565b3480156101af575f80fd5b506068546101c3906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101eb575f80fd5b5061018f6101fa3660046116cf565b61049f565b34801561020a575f80fd5b506101c36102193660046116cf565b610510565b348015610229575f80fd5b5061018f6102383660046116ea565b610545565b348015610248575f80fd5b506101c36102573660046116cf565b610798565b61018f61026a366004611827565b610843565b34801561027a575f80fd5b5061018f6102893660046116cf565b610a21565b348015610299575f80fd5b5061018f6102a8366004611911565b610a92565b61018f6102bb36600461169d565b610c46565b61018f6102ce366004611971565b610c7f565b3480156102de575f80fd5b506101c36102ed3660046116cf565b60676020525f90815260409020546001600160a01b031681565b348015610312575f80fd5b5061018f610e51565b61018f61018a3660046119ce565b348015610334575f80fd5b506065546101c3906001600160a01b031681565b348015610353575f80fd5b506033546001600160a01b03166101c3565b61018f610373366004611a60565b610e64565b348015610383575f80fd5b50610397610392366004611aa3565b610e76565b6040519081526020016101d7565b3480156103b0575f80fd5b506101c36103bf3660046116cf565b611012565b61018f6103d2366004611ae1565b61105c565b3480156103e2575f80fd5b506066546101c3906001600160a01b031681565b348015610401575f80fd5b5061018f6104103660046116cf565b61106b565b60405162461bcd60e51b815260206004820152601660248201527f73686f756c64206e657665722062652063616c6c65640000000000000000000060448201526064015b60405180910390fd5b61049a83835f5b6040519080825280601f01601f191660200182016040528015610493576020820181803683370190505b5084610c7f565b505050565b6104a76110fb565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a905f90a35050565b6001600160a01b038082165f908152606760205260408120549091168061053f57506066546001600160a01b03165b92915050565b5f54610100900460ff161580801561056357505f54600160ff909116105b8061057c5750303b15801561057c57505f5460ff166001145b6105ee5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610459565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561064a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610652611155565b6001600160a01b038216156106c257606680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040515f907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1908290a35b6001600160a01b0383161561073257606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385169081179091556040515f907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a908290a35b801561049a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b5f806107a383610510565b90506001600160a01b0381166107bb57505f92915050565b6040517f54bbd59c0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301528216906354bbd59c90602401602060405180830381865afa158015610818573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083c9190611b01565b9392505050565b6068546001600160a01b03161561089c5760405162461bcd60e51b815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e74657874000000000000000000000000006044820152606401610459565b5f6108a686610510565b90506001600160a01b0381166108fe5760405162461bcd60e51b815260206004820152601460248201527f6e6f206761746577617920617661696c61626c650000000000000000000000006044820152606401610459565b606880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790555f3384604051602001610945929190611b87565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f575361b600000000000000000000000000000000000000000000000000000000825291506001600160a01b0383169063575361b69034906109c2908b908b908b9088908b90600401611ba8565b5f604051808303818588803b1580156109d9575f80fd5b505af11580156109eb573d5f803e3d5ffd5b5050606880547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550505050505050505050565b610a296110fb565b606680546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1905f90a35050565b610a9a6110fb565b8051825114610aeb5760405162461bcd60e51b815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610459565b5f5b825181101561049a575f60675f858481518110610b0c57610b0c611beb565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b03169050828281518110610b5a57610b5a611beb565b602002602001015160675f868581518110610b7757610b77611beb565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550828281518110610bd357610bd3611beb565b60200260200101516001600160a01b0316816001600160a01b0316858481518110610c0057610c00611beb565b60200260200101516001600160a01b03167f0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf60405160405180910390a450600101610aed565b61049a8333845f5b6040519080825280601f01601f191660200182016040528015610c78576020820181803683370190505b5085610843565b6068546001600160a01b031615610cd85760405162461bcd60e51b815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e74657874000000000000000000000000006044820152606401610459565b6065546001600160a01b031680610d315760405162461bcd60e51b815260206004820152601560248201527f657468206761746577617920617661696c61626c6500000000000000000000006044820152606401610459565b606880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790555f3384604051602001610d78929190611b87565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f6dc2418300000000000000000000000000000000000000000000000000000000825291506001600160a01b03831690636dc24183903490610df3908a908a9087908a90600401611c18565b5f604051808303818588803b158015610e0a575f80fd5b505af1158015610e1c573d5f803e3d5ffd5b5050606880547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055505050505050505050565b610e596110fb565b610e625f6111d9565b565b610e708484845f610c4e565b50505050565b6068545f906001600160a01b0316336001600160a01b031614610edb5760405162461bcd60e51b815260206004820152601760248201527f4f6e6c7920696e206465706f73697420636f6e746578740000000000000000006044820152606401610459565b5f336040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192505f918616906370a0823190602401602060405180830381865afa158015610f3f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f639190611c50565b9050610f7a6001600160a01b038616878487611242565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301528291908716906370a0823190602401602060405180830381865afa158015610fda573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ffe9190611c50565b6110089190611c67565b9695505050505050565b60405162461bcd60e51b815260206004820152600b60248201527f756e737570706f7274656400000000000000000000000000000000000000000060448201525f90606401610459565b61106733835f610469565b5050565b6110736110fb565b6001600160a01b0381166110ef5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610459565b6110f8816111d9565b50565b6033546001600160a01b03163314610e625760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610459565b5f54610100900460ff166111d15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610459565b610e626112ca565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610e7090859061134f565b5f54610100900460ff166113465760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610459565b610e62336111d9565b5f6113a3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114359092919063ffffffff16565b905080515f14806113c35750808060200190518101906113c39190611c9f565b61049a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610459565b606061144384845f8561144b565b949350505050565b6060824710156114c35760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610459565b5f80866001600160a01b031685876040516114de9190611cbe565b5f6040518083038185875af1925050503d805f8114611518576040519150601f19603f3d011682016040523d82523d5f602084013e61151d565b606091505b509150915061152e87838387611539565b979650505050505050565b606083156115a75782515f036115a0576001600160a01b0385163b6115a05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610459565b5081611443565b61144383838151156115bc5781518083602001fd5b8060405162461bcd60e51b81526004016104599190611cd9565b6001600160a01b03811681146110f8575f80fd5b5f8083601f8401126115fa575f80fd5b50813567ffffffffffffffff811115611611575f80fd5b602083019150836020828501011115611628575f80fd5b9250929050565b5f805f805f60808688031215611643575f80fd5b853561164e816115d6565b9450602086013561165e816115d6565b935060408601359250606086013567ffffffffffffffff811115611680575f80fd5b61168c888289016115ea565b969995985093965092949392505050565b5f805f606084860312156116af575f80fd5b83356116ba816115d6565b95602085013595506040909401359392505050565b5f602082840312156116df575f80fd5b813561083c816115d6565b5f80604083850312156116fb575f80fd5b8235611706816115d6565b91506020830135611716816115d6565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561179557611795611721565b604052919050565b5f82601f8301126117ac575f80fd5b813567ffffffffffffffff8111156117c6576117c6611721565b6117f760207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161174e565b81815284602083860101111561180b575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a0868803121561183b575f80fd5b8535611846816115d6565b94506020860135611856816115d6565b935060408601359250606086013567ffffffffffffffff811115611878575f80fd5b6118848882890161179d565b95989497509295608001359392505050565b5f82601f8301126118a5575f80fd5b8135602067ffffffffffffffff8211156118c1576118c1611721565b8160051b6118d082820161174e565b92835284810182019282810190878511156118e9575f80fd5b83870192505b8483101561152e578235611902816115d6565b825291830191908301906118ef565b5f8060408385031215611922575f80fd5b823567ffffffffffffffff80821115611939575f80fd5b61194586838701611896565b9350602085013591508082111561195a575f80fd5b5061196785828601611896565b9150509250929050565b5f805f8060808587031215611984575f80fd5b843561198f816115d6565b935060208501359250604085013567ffffffffffffffff8111156119b1575f80fd5b6119bd8782880161179d565b949793965093946060013593505050565b5f805f805f805f60c0888a0312156119e4575f80fd5b87356119ef816115d6565b965060208801356119ff816115d6565b95506040880135611a0f816115d6565b94506060880135611a1f816115d6565b93506080880135925060a088013567ffffffffffffffff811115611a41575f80fd5b611a4d8a828b016115ea565b989b979a50959850939692959293505050565b5f805f8060808587031215611a73575f80fd5b8435611a7e816115d6565b93506020850135611a8e816115d6565b93969395505050506040820135916060013590565b5f805f60608486031215611ab5575f80fd5b8335611ac0816115d6565b92506020840135611ad0816115d6565b929592945050506040919091013590565b5f8060408385031215611af2575f80fd5b50508035926020909101359150565b5f60208284031215611b11575f80fd5b815161083c816115d6565b5f5b83811015611b36578181015183820152602001611b1e565b50505f910152565b5f8151808452611b55816020860160208601611b1c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6001600160a01b0383168152604060208201525f6114436040830184611b3e565b5f6001600160a01b03808816835280871660208401525084604083015260a06060830152611bd960a0830185611b3e565b90508260808301529695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6001600160a01b0385168152836020820152608060408201525f611c3f6080830185611b3e565b905082606083015295945050505050565b5f60208284031215611c60575f80fd5b5051919050565b8181038181111561053f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60208284031215611caf575f80fd5b8151801515811461083c575f80fd5b5f8251611ccf818460208701611b1c565b9190910192915050565b602081525f61083c6020830184611b3e56fea164736f6c6343000818000a",
}

// L2GatewayRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use L2GatewayRouterMetaData.ABI instead.
var L2GatewayRouterABI = L2GatewayRouterMetaData.ABI

// L2GatewayRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2GatewayRouterMetaData.Bin instead.
var L2GatewayRouterBin = L2GatewayRouterMetaData.Bin

// DeployL2GatewayRouter deploys a new Ethereum contract, binding an instance of L2GatewayRouter to it.
func DeployL2GatewayRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2GatewayRouter, error) {
	parsed, err := L2GatewayRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2GatewayRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2GatewayRouter{L2GatewayRouterCaller: L2GatewayRouterCaller{contract: contract}, L2GatewayRouterTransactor: L2GatewayRouterTransactor{contract: contract}, L2GatewayRouterFilterer: L2GatewayRouterFilterer{contract: contract}}, nil
}

// L2GatewayRouter is an auto generated Go binding around an Ethereum contract.
type L2GatewayRouter struct {
	L2GatewayRouterCaller     // Read-only binding to the contract
	L2GatewayRouterTransactor // Write-only binding to the contract
	L2GatewayRouterFilterer   // Log filterer for contract events
}

// L2GatewayRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2GatewayRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2GatewayRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2GatewayRouterSession struct {
	Contract     *L2GatewayRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2GatewayRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2GatewayRouterCallerSession struct {
	Contract *L2GatewayRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2GatewayRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2GatewayRouterTransactorSession struct {
	Contract     *L2GatewayRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2GatewayRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2GatewayRouterRaw struct {
	Contract *L2GatewayRouter // Generic contract binding to access the raw methods on
}

// L2GatewayRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2GatewayRouterCallerRaw struct {
	Contract *L2GatewayRouterCaller // Generic read-only contract binding to access the raw methods on
}

// L2GatewayRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactorRaw struct {
	Contract *L2GatewayRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2GatewayRouter creates a new instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouter(address common.Address, backend bind.ContractBackend) (*L2GatewayRouter, error) {
	contract, err := bindL2GatewayRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouter{L2GatewayRouterCaller: L2GatewayRouterCaller{contract: contract}, L2GatewayRouterTransactor: L2GatewayRouterTransactor{contract: contract}, L2GatewayRouterFilterer: L2GatewayRouterFilterer{contract: contract}}, nil
}

// NewL2GatewayRouterCaller creates a new read-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterCaller(address common.Address, caller bind.ContractCaller) (*L2GatewayRouterCaller, error) {
	contract, err := bindL2GatewayRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterCaller{contract: contract}, nil
}

// NewL2GatewayRouterTransactor creates a new write-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*L2GatewayRouterTransactor, error) {
	contract, err := bindL2GatewayRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterTransactor{contract: contract}, nil
}

// NewL2GatewayRouterFilterer creates a new log filterer instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*L2GatewayRouterFilterer, error) {
	contract, err := bindL2GatewayRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFilterer{contract: contract}, nil
}

// bindL2GatewayRouter binds a generic wrapper to an already deployed contract.
func bindL2GatewayRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2GatewayRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.L2GatewayRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GatewayInContext(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "gatewayInContext")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GatewayInContext() (common.Address, error) {
	return _L2GatewayRouter.Contract.GatewayInContext(&_L2GatewayRouter.CallOpts)
}

// GatewayInContext is a free data retrieval call binding the contract method 0x3a9a7b20.
//
// Solidity: function gatewayInContext() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GatewayInContext() (common.Address, error) {
	return _L2GatewayRouter.Contract.GatewayInContext(&_L2GatewayRouter.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL1ERC20Address", _l2Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositETH", arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterTransactor) RequestERC20(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "requestERC20", _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RequestERC20(&_L2GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// RequestERC20 is a paid mutator transaction binding the contract method 0xc52a3bbc.
//
// Solidity: function requestERC20(address _sender, address _token, uint256 _amount) returns(uint256)
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) RequestERC20(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RequestERC20(&_L2GatewayRouter.TransactOpts, _sender, _token, _amount)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setDefaultERC20Gateway", _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetETHGateway(opts *bind.TransactOpts, _newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setETHGateway", _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// L2GatewayRouterFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20Iterator struct {
	Event *L2GatewayRouterFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
		it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20 struct {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositERC20Iterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2GatewayRouterFinalizeDepositERC20, error) {
	event := new(L2GatewayRouterFinalizeDepositERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETHIterator struct {
	Event *L2GatewayRouterFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
		it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositETHIterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositETH(log types.Log) (*L2GatewayRouterFinalizeDepositETH, error) {
	event := new(L2GatewayRouterFinalizeDepositETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2GatewayRouter contract.
type L2GatewayRouterInitializedIterator struct {
	Event *L2GatewayRouterInitialized // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterInitialized)
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
		it.Event = new(L2GatewayRouterInitialized)
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
func (it *L2GatewayRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterInitialized represents a Initialized event raised by the L2GatewayRouter contract.
type L2GatewayRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2GatewayRouterInitializedIterator, error) {

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterInitializedIterator{contract: _L2GatewayRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterInitialized)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseInitialized(log types.Log) (*L2GatewayRouterInitialized, error) {
	event := new(L2GatewayRouterInitialized)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferredIterator struct {
	Event *L2GatewayRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterOwnershipTransferred)
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
		it.Event = new(L2GatewayRouterOwnershipTransferred)
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
func (it *L2GatewayRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterOwnershipTransferred represents a OwnershipTransferred event raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2GatewayRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterOwnershipTransferredIterator{contract: _L2GatewayRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterOwnershipTransferred)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseOwnershipTransferred(log types.Log) (*L2GatewayRouterOwnershipTransferred, error) {
	event := new(L2GatewayRouterOwnershipTransferred)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20GatewayIterator struct {
	Event *L2GatewayRouterSetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20Gateway struct {
	OldDefaultERC20Gateway common.Address
	NewDefaultERC20Gateway common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (*L2GatewayRouterSetDefaultERC20GatewayIterator, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetDefaultERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetDefaultERC20Gateway, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (event.Subscription, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetDefaultERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*L2GatewayRouterSetDefaultERC20Gateway, error) {
	event := new(L2GatewayRouterSetDefaultERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20GatewayIterator struct {
	Event *L2GatewayRouterSetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetERC20Gateway)
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
func (it *L2GatewayRouterSetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetERC20Gateway represents a SetERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20Gateway struct {
	Token      common.Address
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (*L2GatewayRouterSetERC20GatewayIterator, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetERC20Gateway, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetERC20Gateway(log types.Log) (*L2GatewayRouterSetERC20Gateway, error) {
	event := new(L2GatewayRouterSetERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGatewayIterator struct {
	Event *L2GatewayRouterSetETHGateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetETHGateway)
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
		it.Event = new(L2GatewayRouterSetETHGateway)
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
func (it *L2GatewayRouterSetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetETHGateway represents a SetETHGateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGateway struct {
	OldETHGateway common.Address
	NewEthGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetETHGateway(opts *bind.FilterOpts, oldETHGateway []common.Address, newEthGateway []common.Address) (*L2GatewayRouterSetETHGatewayIterator, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetETHGatewayIterator{contract: _L2GatewayRouter.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetETHGateway, oldETHGateway []common.Address, newEthGateway []common.Address) (event.Subscription, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetETHGateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetETHGateway(log types.Log) (*L2GatewayRouterSetETHGateway, error) {
	event := new(L2GatewayRouterSetETHGateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20Iterator struct {
	Event *L2GatewayRouterWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawERC20)
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
		it.Event = new(L2GatewayRouterWithdrawERC20)
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
func (it *L2GatewayRouterWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawERC20 represents a WithdrawERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20 struct {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawERC20Iterator{contract: _L2GatewayRouter.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawERC20(log types.Log) (*L2GatewayRouterWithdrawERC20, error) {
	event := new(L2GatewayRouterWithdrawERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETHIterator struct {
	Event *L2GatewayRouterWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawETH)
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
		it.Event = new(L2GatewayRouterWithdrawETH)
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
func (it *L2GatewayRouterWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawETH represents a WithdrawETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawETHIterator{contract: _L2GatewayRouter.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawETH(log types.Log) (*L2GatewayRouterWithdrawETH, error) {
	event := new(L2GatewayRouterWithdrawETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
