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

// L1WETHGatewayMetaData contains all meta data concerning the L1WETHGateway contract.
var L1WETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620024ce380380620024ce833981016040819052620000349162000134565b6200003e62000056565b6001600160a01b0391821660a052166080526200016c565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000115576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200012f57600080fd5b919050565b600080604083850312156200014857600080fd5b620001538362000117565b9150620001636020840162000117565b90509250929050565b60805160a051612305620001c96000396000818160f40152818161029601528181610b0f01528181610e2701526110b9015260008181610244015281816102f601528181610c1b01528181610d54015261113a01526123056000f3fe6080604052600436106100ec5760003560e01c8063885586871161008a578063c676ad2911610059578063c676ad29146102d8578063f219fa6614610318578063f2fde38b1461032b578063f887ea401461034b57600080fd5b806388558687146102325780638da5cb5b14610266578063ad5c464814610284578063c0c53b8b146102b857600080fd5b80633cb747bf116100c65780633cb747bf146101ae578063715018a6146101ea578063797594b0146101ff57806384bd13b01461021f57600080fd5b80630aea8c261461017557806314298c511461018857806321425ee01461019b57600080fd5b3661017057337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161461016e5760405162461bcd60e51b815260206004820152600960248201527f6f6e6c792057455448000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b005b600080fd5b61016e610183366004611c85565b61036b565b61016e610196366004611d42565b61037f565b61016e6101a9366004611d84565b610603565b3480156101ba57600080fd5b506099546101ce906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b3480156101f657600080fd5b5061016e61063d565b34801561020b57600080fd5b506097546101ce906001600160a01b031681565b61016e61022d366004611db9565b610651565b34801561023e57600080fd5b506101ce7f000000000000000000000000000000000000000000000000000000000000000081565b34801561027257600080fd5b506065546001600160a01b03166101ce565b34801561029057600080fd5b506101ce7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102c457600080fd5b5061016e6102d3366004611e51565b610844565b3480156102e457600080fd5b506101ce6102f3366004611e9c565b507f000000000000000000000000000000000000000000000000000000000000000090565b61016e610326366004611ec0565b610a18565b34801561033757600080fd5b5061016e610346366004611e9c565b610a25565b34801561035757600080fd5b506098546101ce906001600160a01b031681565b6103788585858585610ab5565b5050505050565b6099546001600160a01b03163381146103da5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610165565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610418573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043c9190611f06565b6001600160a01b0316736f297c61b5c92ef107ffd30cd56affe5a273e8416001600160a01b0316146104b05760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e7465787400000000006044820152606401610165565b6104b8610dcc565b7f8431f5c1000000000000000000000000000000000000000000000000000000006104e7600460008587611f23565b6104f091611f4d565b7fffffffff00000000000000000000000000000000000000000000000000000000161461055f5760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642073656c6563746f72000000000000000000000000000000006044820152606401610165565b600080806105708560048189611f23565b81019061057d9190611f95565b509450509350509250610591838383610e25565b6105a56001600160a01b0384168383610f44565b816001600160a01b0316836001600160a01b03167fdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8836040516105ea91815260200190565b60405180910390a35050506105fe60018055565b505050565b6105fe83338460005b6040519080825280601f01601f191660200182016040528015610636576020820181803683370190505b5085610ab5565b610645610ff3565b61064f600061104d565b565b6099546001600160a01b03163381146106ac5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610165565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070e9190611f06565b6097546001600160a01b0390811691161461076b5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610165565b610773610dcc565b610782888888888888886110b7565b6107966001600160a01b0389168686610f44565b6107d68584848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061126592505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7888888886040516108299493929190612025565b60405180910390a461083a60018055565b5050505050505050565b600054610100900460ff16158080156108645750600054600160ff909116105b8061087e5750303b15801561087e575060005460ff166001145b6108f05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610165565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561094e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0383166109a45760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f757465722061646472657373000000000000000000000000006044820152606401610165565b6109af848484611301565b8015610a1257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610a12848484600061060c565b610a2d610ff3565b6001600160a01b038116610aa95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610165565b610ab28161104d565b50565b610abd610dcc565b60008311610b0d5760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e74000000000000000000000000006044820152606401610165565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b031614610b8e5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c79205745544820697320616c6c6f7765640000000000000000000000006044820152606401610165565b6000610b9b868585611444565b6040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810183905291965094509091506001600160a01b03871690632e1a7d4d90602401600060405180830381600087803b158015610bfe57600080fd5b505af1158015610c12573d6000803e3d6000fd5b505050506000867f000000000000000000000000000000000000000000000000000000000000000083888888604051602401610c53969594939291906120bb565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8431f5c1000000000000000000000000000000000000000000000000000000001790526099549091506001600160a01b0316635f7b1577610cc93488612138565b6097546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610d16916001600160a01b0316908a9087908a908a90600401612151565b6000604051808303818588803b158015610d2f57600080fd5b505af1158015610d43573d6000803e3d6000fd5b5050505050816001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316886001600160a01b03167f31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25898989604051610db993929190612194565b60405180910390a4505061037860018055565b600260015403610e1e5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610165565b6002600155565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031614610ea65760405162461bcd60e51b815260206004820152600e60248201527f746f6b656e206e6f7420574554480000000000000000000000000000000000006044820152606401610165565b348114610ef55760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d6174636800000000000000000000000000006044820152606401610165565b826001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610f3057600080fd5b505af115801561083a573d6000803e3d6000fd5b6040516001600160a01b0383166024820152604481018290526105fe9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526116a5565b60018055565b6065546001600160a01b0316331461064f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610165565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316876001600160a01b0316146111385760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610165565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316866001600160a01b0316146111b95760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610165565b3483146112085760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d6174636800000000000000000000000000006044820152606401610165565b866001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561124357600080fd5b505af1158015611257573d6000803e3d6000fd5b505050505050505050505050565b6000815111801561128057506000826001600160a01b03163b115b156112fd576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f906112ca9084906004016121c5565b600060405180830381600087803b1580156112e457600080fd5b505af11580156112f8573d6000803e3d6000fd5b505050505b5050565b6001600160a01b0383166113575760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610165565b6001600160a01b0381166113ad5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610165565b6113b561178d565b6113bd611812565b609780546001600160a01b038086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105fe57609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6098546000908190606090339081906001600160a01b0316819003611518578580602001905181019061147791906121d8565b6040517fc52a3bbc0000000000000000000000000000000000000000000000000000000081526001600160a01b0380841660048301528b81166024830152604482018b905291985091925083169063c52a3bbc906064016020604051808303816000875af11580156114ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115119190612265565b9650611648565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038a16906370a0823190602401602060405180830381865afa158015611578573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159c9190612265565b90506115b36001600160a01b038a1683308b611897565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038b16906370a0823190602401602060405180830381865afa158015611613573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116379190612265565b9050611643828261227e565b985050505b600087116116985760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e74000000000000000000000000006044820152606401610165565b9795965093949350505050565b60006116fa826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166118e89092919063ffffffff16565b905080516000148061171b57508080602001905181019061171b9190612291565b6105fe5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610165565b600054610100900460ff1661180a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b61064f6118ff565b600054610100900460ff1661188f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b61064f61197c565b6040516001600160a01b0380851660248301528316604482015260648101829052610a129085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610f89565b60606118f78484600085611a02565b949350505050565b600054610100900460ff16610fed5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b600054610100900460ff166119f95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b61064f3361104d565b606082471015611a7a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610165565b600080866001600160a01b03168587604051611a9691906122b3565b60006040518083038185875af1925050503d8060008114611ad3576040519150601f19603f3d011682016040523d82523d6000602084013e611ad8565b606091505b5091509150611ae987838387611af4565b979650505050505050565b60608315611b63578251600003611b5c576001600160a01b0385163b611b5c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610165565b50816118f7565b6118f78383815115611b785781518083602001fd5b8060405162461bcd60e51b815260040161016591906121c5565b6001600160a01b0381168114610ab257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611bff57611bff611ba7565b604052919050565b600067ffffffffffffffff821115611c2157611c21611ba7565b50601f01601f191660200190565b600082601f830112611c4057600080fd5b8135611c53611c4e82611c07565b611bd6565b818152846020838601011115611c6857600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a08688031215611c9d57600080fd5b8535611ca881611b92565b94506020860135611cb881611b92565b935060408601359250606086013567ffffffffffffffff811115611cdb57600080fd5b611ce788828901611c2f565b95989497509295608001359392505050565b60008083601f840112611d0b57600080fd5b50813567ffffffffffffffff811115611d2357600080fd5b602083019150836020828501011115611d3b57600080fd5b9250929050565b60008060208385031215611d5557600080fd5b823567ffffffffffffffff811115611d6c57600080fd5b611d7885828601611cf9565b90969095509350505050565b600080600060608486031215611d9957600080fd5b8335611da481611b92565b95602085013595506040909401359392505050565b600080600080600080600060c0888a031215611dd457600080fd5b8735611ddf81611b92565b96506020880135611def81611b92565b95506040880135611dff81611b92565b94506060880135611e0f81611b92565b93506080880135925060a088013567ffffffffffffffff811115611e3257600080fd5b611e3e8a828b01611cf9565b989b979a50959850939692959293505050565b600080600060608486031215611e6657600080fd5b8335611e7181611b92565b92506020840135611e8181611b92565b91506040840135611e9181611b92565b809150509250925092565b600060208284031215611eae57600080fd5b8135611eb981611b92565b9392505050565b60008060008060808587031215611ed657600080fd5b8435611ee181611b92565b93506020850135611ef181611b92565b93969395505050506040820135916060013590565b600060208284031215611f1857600080fd5b8151611eb981611b92565b60008085851115611f3357600080fd5b83861115611f4057600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015611f8d5780818660040360031b1b83161692505b505092915050565b60008060008060008060c08789031215611fae57600080fd5b8635611fb981611b92565b95506020870135611fc981611b92565b94506040870135611fd981611b92565b93506060870135611fe981611b92565b92506080870135915060a087013567ffffffffffffffff81111561200c57600080fd5b61201889828a01611c2f565b9150509295509295509295565b6001600160a01b038516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f909201601f191601019392505050565b60005b8381101561208657818101518382015260200161206e565b50506000910152565b600081518084526120a781602086016020860161206b565b601f01601f19169290920160200192915050565b60006001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526120fd60c083018461208f565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561214b5761214b612109565b92915050565b60006001600160a01b03808816835286602084015260a0604084015261217a60a084018761208f565b606084019590955292909216608090910152509392505050565b6001600160a01b03841681528260208201526060604082015260006121bc606083018461208f565b95945050505050565b602081526000611eb9602083018461208f565b600080604083850312156121eb57600080fd5b82516121f681611b92565b602084015190925067ffffffffffffffff81111561221357600080fd5b8301601f8101851361222457600080fd5b8051612232611c4e82611c07565b81815286602083850101111561224757600080fd5b61225882602083016020860161206b565b8093505050509250929050565b60006020828403121561227757600080fd5b5051919050565b8181038181111561214b5761214b612109565b6000602082840312156122a357600080fd5b81518015158114611eb957600080fd5b600082516122c581846020870161206b565b919091019291505056fea26469706673582212209baee73b2a52273271312b15d060f619f6048c24cf7ad06bce26b70eb029b02664736f6c63430008100033",
}

// L1WETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1WETHGatewayMetaData.ABI instead.
var L1WETHGatewayABI = L1WETHGatewayMetaData.ABI

// L1WETHGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1WETHGatewayMetaData.Bin instead.
var L1WETHGatewayBin = L1WETHGatewayMetaData.Bin

// DeployL1WETHGateway deploys a new Ethereum contract, binding an instance of L1WETHGateway to it.
func DeployL1WETHGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _WETH common.Address, _l2WETH common.Address) (common.Address, *types.Transaction, *L1WETHGateway, error) {
	parsed, err := L1WETHGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1WETHGatewayBin), backend, _WETH, _l2WETH)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1WETHGateway{L1WETHGatewayCaller: L1WETHGatewayCaller{contract: contract}, L1WETHGatewayTransactor: L1WETHGatewayTransactor{contract: contract}, L1WETHGatewayFilterer: L1WETHGatewayFilterer{contract: contract}}, nil
}

// L1WETHGateway is an auto generated Go binding around an Ethereum contract.
type L1WETHGateway struct {
	L1WETHGatewayCaller     // Read-only binding to the contract
	L1WETHGatewayTransactor // Write-only binding to the contract
	L1WETHGatewayFilterer   // Log filterer for contract events
}

// L1WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1WETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1WETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1WETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1WETHGatewaySession struct {
	Contract     *L1WETHGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1WETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1WETHGatewayCallerSession struct {
	Contract *L1WETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1WETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1WETHGatewayTransactorSession struct {
	Contract     *L1WETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1WETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1WETHGatewayRaw struct {
	Contract *L1WETHGateway // Generic contract binding to access the raw methods on
}

// L1WETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1WETHGatewayCallerRaw struct {
	Contract *L1WETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1WETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1WETHGatewayTransactorRaw struct {
	Contract *L1WETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1WETHGateway creates a new instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGateway(address common.Address, backend bind.ContractBackend) (*L1WETHGateway, error) {
	contract, err := bindL1WETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1WETHGateway{L1WETHGatewayCaller: L1WETHGatewayCaller{contract: contract}, L1WETHGatewayTransactor: L1WETHGatewayTransactor{contract: contract}, L1WETHGatewayFilterer: L1WETHGatewayFilterer{contract: contract}}, nil
}

// NewL1WETHGatewayCaller creates a new read-only instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1WETHGatewayCaller, error) {
	contract, err := bindL1WETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayCaller{contract: contract}, nil
}

// NewL1WETHGatewayTransactor creates a new write-only instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1WETHGatewayTransactor, error) {
	contract, err := bindL1WETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayTransactor{contract: contract}, nil
}

// NewL1WETHGatewayFilterer creates a new log filterer instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1WETHGatewayFilterer, error) {
	contract, err := bindL1WETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayFilterer{contract: contract}, nil
}

// bindL1WETHGateway binds a generic wrapper to an already deployed contract.
func bindL1WETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1WETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1WETHGateway *L1WETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1WETHGateway.Contract.L1WETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1WETHGateway *L1WETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.L1WETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1WETHGateway *L1WETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.L1WETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1WETHGateway *L1WETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1WETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1WETHGateway *L1WETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1WETHGateway *L1WETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) WETH() (common.Address, error) {
	return _L1WETHGateway.Contract.WETH(&_L1WETHGateway.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) WETH() (common.Address, error) {
	return _L1WETHGateway.Contract.WETH(&_L1WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) Counterpart() (common.Address, error) {
	return _L1WETHGateway.Contract.Counterpart(&_L1WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1WETHGateway.Contract.Counterpart(&_L1WETHGateway.CallOpts)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L1WETHGateway.Contract.GetL2ERC20Address(&_L1WETHGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L1WETHGateway.Contract.GetL2ERC20Address(&_L1WETHGateway.CallOpts, arg0)
}

// L2WETH is a free data retrieval call binding the contract method 0x88558687.
//
// Solidity: function l2WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "l2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0x88558687.
//
// Solidity: function l2WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) L2WETH() (common.Address, error) {
	return _L1WETHGateway.Contract.L2WETH(&_L1WETHGateway.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0x88558687.
//
// Solidity: function l2WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) L2WETH() (common.Address, error) {
	return _L1WETHGateway.Contract.L2WETH(&_L1WETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) Messenger() (common.Address, error) {
	return _L1WETHGateway.Contract.Messenger(&_L1WETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) Messenger() (common.Address, error) {
	return _L1WETHGateway.Contract.Messenger(&_L1WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) Owner() (common.Address, error) {
	return _L1WETHGateway.Contract.Owner(&_L1WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) Owner() (common.Address, error) {
	return _L1WETHGateway.Contract.Owner(&_L1WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1WETHGateway *L1WETHGatewaySession) Router() (common.Address, error) {
	return _L1WETHGateway.Contract.Router(&_L1WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCallerSession) Router() (common.Address, error) {
	return _L1WETHGateway.Contract.Router(&_L1WETHGateway.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC20(&_L1WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC20(&_L1WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC200(&_L1WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC200(&_L1WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC20AndCall(&_L1WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.DepositERC20AndCall(&_L1WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.FinalizeWithdrawERC20(&_L1WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.FinalizeWithdrawERC20(&_L1WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1WETHGateway *L1WETHGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.Initialize(&_L1WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.Initialize(&_L1WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.OnDropMessage(&_L1WETHGateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.OnDropMessage(&_L1WETHGateway.TransactOpts, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1WETHGateway *L1WETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1WETHGateway.Contract.RenounceOwnership(&_L1WETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1WETHGateway.Contract.RenounceOwnership(&_L1WETHGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1WETHGateway *L1WETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.TransferOwnership(&_L1WETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.Contract.TransferOwnership(&_L1WETHGateway.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1WETHGateway *L1WETHGatewaySession) Receive() (*types.Transaction, error) {
	return _L1WETHGateway.Contract.Receive(&_L1WETHGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _L1WETHGateway.Contract.Receive(&_L1WETHGateway.TransactOpts)
}

// L1WETHGatewayDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1WETHGateway contract.
type L1WETHGatewayDepositERC20Iterator struct {
	Event *L1WETHGatewayDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1WETHGatewayDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WETHGatewayDepositERC20)
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
		it.Event = new(L1WETHGatewayDepositERC20)
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
func (it *L1WETHGatewayDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WETHGatewayDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WETHGatewayDepositERC20 represents a DepositERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayDepositERC20 struct {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1WETHGatewayDepositERC20Iterator, error) {

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

	logs, sub, err := _L1WETHGateway.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayDepositERC20Iterator{contract: _L1WETHGateway.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1WETHGateway *L1WETHGatewayFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1WETHGatewayDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1WETHGateway.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WETHGatewayDepositERC20)
				if err := _L1WETHGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) ParseDepositERC20(log types.Log) (*L1WETHGatewayDepositERC20, error) {
	event := new(L1WETHGatewayDepositERC20)
	if err := _L1WETHGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1WETHGatewayFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1WETHGateway contract.
type L1WETHGatewayFinalizeWithdrawERC20Iterator struct {
	Event *L1WETHGatewayFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1WETHGatewayFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WETHGatewayFinalizeWithdrawERC20)
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
		it.Event = new(L1WETHGatewayFinalizeWithdrawERC20)
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
func (it *L1WETHGatewayFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WETHGatewayFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WETHGatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayFinalizeWithdrawERC20 struct {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1WETHGatewayFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1WETHGateway.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayFinalizeWithdrawERC20Iterator{contract: _L1WETHGateway.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1WETHGateway *L1WETHGatewayFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1WETHGatewayFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1WETHGateway.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WETHGatewayFinalizeWithdrawERC20)
				if err := _L1WETHGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1WETHGatewayFinalizeWithdrawERC20, error) {
	event := new(L1WETHGatewayFinalizeWithdrawERC20)
	if err := _L1WETHGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1WETHGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1WETHGateway contract.
type L1WETHGatewayInitializedIterator struct {
	Event *L1WETHGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1WETHGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WETHGatewayInitialized)
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
		it.Event = new(L1WETHGatewayInitialized)
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
func (it *L1WETHGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WETHGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WETHGatewayInitialized represents a Initialized event raised by the L1WETHGateway contract.
type L1WETHGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1WETHGateway *L1WETHGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1WETHGatewayInitializedIterator, error) {

	logs, sub, err := _L1WETHGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayInitializedIterator{contract: _L1WETHGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1WETHGateway *L1WETHGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1WETHGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1WETHGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WETHGatewayInitialized)
				if err := _L1WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) ParseInitialized(log types.Log) (*L1WETHGatewayInitialized, error) {
	event := new(L1WETHGatewayInitialized)
	if err := _L1WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1WETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1WETHGateway contract.
type L1WETHGatewayOwnershipTransferredIterator struct {
	Event *L1WETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1WETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WETHGatewayOwnershipTransferred)
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
		it.Event = new(L1WETHGatewayOwnershipTransferred)
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
func (it *L1WETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1WETHGateway contract.
type L1WETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1WETHGateway *L1WETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1WETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1WETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayOwnershipTransferredIterator{contract: _L1WETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1WETHGateway *L1WETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1WETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1WETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WETHGatewayOwnershipTransferred)
				if err := _L1WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1WETHGatewayOwnershipTransferred, error) {
	event := new(L1WETHGatewayOwnershipTransferred)
	if err := _L1WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1WETHGatewayRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the L1WETHGateway contract.
type L1WETHGatewayRefundERC20Iterator struct {
	Event *L1WETHGatewayRefundERC20 // Event containing the contract specifics and raw log

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
func (it *L1WETHGatewayRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1WETHGatewayRefundERC20)
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
		it.Event = new(L1WETHGatewayRefundERC20)
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
func (it *L1WETHGatewayRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1WETHGatewayRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1WETHGatewayRefundERC20 represents a RefundERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayRefundERC20 struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1WETHGateway *L1WETHGatewayFilterer) FilterRefundERC20(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1WETHGatewayRefundERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1WETHGateway.contract.FilterLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayRefundERC20Iterator{contract: _L1WETHGateway.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1WETHGateway *L1WETHGatewayFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *L1WETHGatewayRefundERC20, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1WETHGateway.contract.WatchLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1WETHGatewayRefundERC20)
				if err := _L1WETHGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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
func (_L1WETHGateway *L1WETHGatewayFilterer) ParseRefundERC20(log types.Log) (*L1WETHGatewayRefundERC20, error) {
	event := new(L1WETHGatewayRefundERC20)
	if err := _L1WETHGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
