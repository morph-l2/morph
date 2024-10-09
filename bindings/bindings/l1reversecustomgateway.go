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

// L1ReverseCustomGatewayMetaData contains all meta data concerning the L1ReverseCustomGateway contract.
var L1ReverseCustomGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL2Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611c44806100e65f395ff3fe6080604052600436106100e4575f3560e01c80638da5cb5b11610087578063f219fa6611610057578063f219fa661461024b578063f2fde38b1461025e578063f887ea401461027d578063fac752eb1461029c575f80fd5b80638da5cb5b146101a4578063ba27f50b146101c1578063c0c53b8b146101f5578063c676ad2914610214575f80fd5b80633cb747bf116100c25780633cb747bf14610123578063715018a61461015e578063797594b01461017257806384bd13b014610191575f80fd5b80630aea8c26146100e857806314298c51146100fd57806321425ee014610110575b5f80fd5b6100fb6100f6366004611641565b6102bb565b005b6100fb61010b3660046116f5565b6102cf565b6100fb61011e366004611734565b6105b6565b34801561012e575f80fd5b50609954610142906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b348015610169575f80fd5b506100fb6105ef565b34801561017d575f80fd5b50609754610142906001600160a01b031681565b6100fb61019f366004611766565b610602565b3480156101af575f80fd5b506065546001600160a01b0316610142565b3480156101cc575f80fd5b506101426101db3660046117f8565b60fa6020525f90815260409020546001600160a01b031681565b348015610200575f80fd5b506100fb61020f36600461181a565b610854565b34801561021f575f80fd5b5061014261022e3660046117f8565b6001600160a01b039081165f90815260fa60205260409020541690565b6100fb610259366004611862565b610a22565b348015610269575f80fd5b506100fb6102783660046117f8565b610a2e565b348015610288575f80fd5b50609854610142906001600160a01b031681565b3480156102a7575f80fd5b506100fb6102b63660046118a5565b610abe565b6102c88585858585610b97565b5050505050565b6099546001600160a01b031633811461032f5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561036b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038f91906118dc565b6001600160a01b0316736f297c61b5c92ef107ffd30cd56affe5a273e8416001600160a01b0316146104035760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e7465787400000000006044820152606401610326565b61040b610eb1565b7f8431f5c10000000000000000000000000000000000000000000000000000000061043960045f85876118f7565b6104429161191e565b7fffffffff0000000000000000000000000000000000000000000000000000000016146104b15760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642073656c6563746f72000000000000000000000000000000006044820152606401610326565b5f80806104c185600481896118f7565b8101906104ce9190611966565b5094505093505092506104e2838383610f0a565b6040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b158015610542575f80fd5b505af1158015610554573d5f803e3d5ffd5b50505050816001600160a01b0316836001600160a01b03167fdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a88360405161059d91815260200190565b60405180910390a35050506105b160018055565b505050565b6105b18333845f5b6040519080825280601f01601f1916602001820160405280156105e8576020820181803683370190505b5085610b97565b6105f7610f5e565b6106005f610fb8565b565b6099546001600160a01b031633811461065d5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610326565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610699573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106bd91906118dc565b6097546001600160a01b0390811691161461071a5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610326565b610722610eb1565b61073188888888888888611021565b6040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018690528916906340c10f19906044015f604051808303815f87803b158015610791575f80fd5b505af11580156107a3573d5f803e3d5ffd5b505050506107e68584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061113992505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a78888888860405161083994939291906119f1565b60405180910390a461084a60018055565b5050505050505050565b5f54610100900460ff161580801561087257505f54600160ff909116105b8061088b5750303b15801561088b57505f5460ff166001145b6108fd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610326565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610959575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0383166109af5760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f757465722061646472657373000000000000000000000000006044820152606401610326565b6109ba8484846111ce565b8015610a1c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610a1c8484845f6105be565b610a36610f5e565b6001600160a01b038116610ab25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610326565b610abb81610fb8565b50565b610ac6610f5e565b6001600160a01b038116610b1c5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610326565b6001600160a01b038083165f81815260fa602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610b9f610eb1565b6001600160a01b038086165f90815260fa60205260409020541680610c065760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3220746f6b656e000000000000006044820152606401610326565b60985433906001600160a01b0316819003610c345783806020019051810190610c2f9190611a76565b945090505b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081526001600160a01b03828116600483015260248201879052881690639dc29fac906044015f604051808303815f87803b158015610c94575f80fd5b505af1158015610ca6573d5f803e3d5ffd5b505050505f878383898989604051602401610cc696959493929190611b47565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8431f5c10000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa158015610d9d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dc19190611b94565b6099546097546040517f5f7b15770000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692635f7b1577923492610e1b929116905f9088908c908b90600401611bab565b5f604051808303818588803b158015610e32575f80fd5b505af1158015610e44573d5f803e3d5ffd5b5050505050826001600160a01b0316846001600160a01b03168a6001600160a01b03167f1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad886728b8b8b87604051610e9c9493929190611bed565b60405180910390a4505050506102c860018055565b600260015403610f035760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610326565b6002600155565b34156105b15760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c75650000000000000000000000000000006044820152606401610326565b60018055565b6065546001600160a01b031633146106005760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b341561106f5760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c75650000000000000000000000000000006044820152606401610326565b6001600160a01b0386166110c55760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610326565b6001600160a01b038088165f90815260fa60205260409020548782169116146111305760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d617463680000000000000000000000000000006044820152606401610326565b50505050505050565b5f815111801561115257505f826001600160a01b03163b115b156111ca576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f9061119c908490600401611c25565b5f604051808303815f87803b1580156111b3575f80fd5b505af11580156111c5573d5f803e3d5ffd5b505050505b5050565b6001600160a01b0383166112245760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610326565b6001600160a01b03811661127a5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610326565b611282611311565b61128a611395565b609780546001600160a01b038086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105b157609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff1661138d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610326565b610600611419565b5f54610100900460ff166114115760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610326565b610600611495565b5f54610100900460ff16610f585760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610326565b5f54610100900460ff166115115760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610326565b61060033610fb8565b6001600160a01b0381168114610abb575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156115a2576115a261152e565b604052919050565b5f67ffffffffffffffff8211156115c3576115c361152e565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126115fe575f80fd5b813561161161160c826115aa565b61155b565b818152846020838601011115611625575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611655575f80fd5b85356116608161151a565b945060208601356116708161151a565b935060408601359250606086013567ffffffffffffffff811115611692575f80fd5b61169e888289016115ef565b95989497509295608001359392505050565b5f8083601f8401126116c0575f80fd5b50813567ffffffffffffffff8111156116d7575f80fd5b6020830191508360208285010111156116ee575f80fd5b9250929050565b5f8060208385031215611706575f80fd5b823567ffffffffffffffff81111561171c575f80fd5b611728858286016116b0565b90969095509350505050565b5f805f60608486031215611746575f80fd5b83356117518161151a565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a03121561177c575f80fd5b87356117878161151a565b965060208801356117978161151a565b955060408801356117a78161151a565b945060608801356117b78161151a565b93506080880135925060a088013567ffffffffffffffff8111156117d9575f80fd5b6117e58a828b016116b0565b989b979a50959850939692959293505050565b5f60208284031215611808575f80fd5b81356118138161151a565b9392505050565b5f805f6060848603121561182c575f80fd5b83356118378161151a565b925060208401356118478161151a565b915060408401356118578161151a565b809150509250925092565b5f805f8060808587031215611875575f80fd5b84356118808161151a565b935060208501356118908161151a565b93969395505050506040820135916060013590565b5f80604083850312156118b6575f80fd5b82356118c18161151a565b915060208301356118d18161151a565b809150509250929050565b5f602082840312156118ec575f80fd5b81516118138161151a565b5f8085851115611905575f80fd5b83861115611911575f80fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561195e5780818660040360031b1b83161692505b505092915050565b5f805f805f8060c0878903121561197b575f80fd5b86356119868161151a565b955060208701356119968161151a565b945060408701356119a68161151a565b935060608701356119b68161151a565b92506080870135915060a087013567ffffffffffffffff8111156119d8575f80fd5b6119e489828a016115ef565b9150509295509295509295565b6001600160a01b038516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611a6e578181015183820152602001611a56565b50505f910152565b5f8060408385031215611a87575f80fd5b8251611a928161151a565b602084015190925067ffffffffffffffff811115611aae575f80fd5b8301601f81018513611abe575f80fd5b8051611acc61160c826115aa565b818152866020838501011115611ae0575f80fd5b611af1826020830160208601611a54565b8093505050509250929050565b5f8151808452611b15816020860160208601611a54565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611b8860c0830184611afe565b98975050505050505050565b5f60208284031215611ba4575f80fd5b5051919050565b5f6001600160a01b03808816835286602084015260a06040840152611bd360a0840187611afe565b606084019590955292909216608090910152509392505050565b6001600160a01b0385168152836020820152608060408201525f611c146080830185611afe565b905082606083015295945050505050565b602081525f6118136020830184611afe56fea164736f6c6343000818000a",
}

// L1ReverseCustomGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ReverseCustomGatewayMetaData.ABI instead.
var L1ReverseCustomGatewayABI = L1ReverseCustomGatewayMetaData.ABI

// L1ReverseCustomGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1ReverseCustomGatewayMetaData.Bin instead.
var L1ReverseCustomGatewayBin = L1ReverseCustomGatewayMetaData.Bin

// DeployL1ReverseCustomGateway deploys a new Ethereum contract, binding an instance of L1ReverseCustomGateway to it.
func DeployL1ReverseCustomGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1ReverseCustomGateway, error) {
	parsed, err := L1ReverseCustomGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1ReverseCustomGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1ReverseCustomGateway{L1ReverseCustomGatewayCaller: L1ReverseCustomGatewayCaller{contract: contract}, L1ReverseCustomGatewayTransactor: L1ReverseCustomGatewayTransactor{contract: contract}, L1ReverseCustomGatewayFilterer: L1ReverseCustomGatewayFilterer{contract: contract}}, nil
}

// L1ReverseCustomGateway is an auto generated Go binding around an Ethereum contract.
type L1ReverseCustomGateway struct {
	L1ReverseCustomGatewayCaller     // Read-only binding to the contract
	L1ReverseCustomGatewayTransactor // Write-only binding to the contract
	L1ReverseCustomGatewayFilterer   // Log filterer for contract events
}

// L1ReverseCustomGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ReverseCustomGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ReverseCustomGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ReverseCustomGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ReverseCustomGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ReverseCustomGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ReverseCustomGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ReverseCustomGatewaySession struct {
	Contract     *L1ReverseCustomGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L1ReverseCustomGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ReverseCustomGatewayCallerSession struct {
	Contract *L1ReverseCustomGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L1ReverseCustomGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ReverseCustomGatewayTransactorSession struct {
	Contract     *L1ReverseCustomGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L1ReverseCustomGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ReverseCustomGatewayRaw struct {
	Contract *L1ReverseCustomGateway // Generic contract binding to access the raw methods on
}

// L1ReverseCustomGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ReverseCustomGatewayCallerRaw struct {
	Contract *L1ReverseCustomGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1ReverseCustomGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ReverseCustomGatewayTransactorRaw struct {
	Contract *L1ReverseCustomGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ReverseCustomGateway creates a new instance of L1ReverseCustomGateway, bound to a specific deployed contract.
func NewL1ReverseCustomGateway(address common.Address, backend bind.ContractBackend) (*L1ReverseCustomGateway, error) {
	contract, err := bindL1ReverseCustomGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGateway{L1ReverseCustomGatewayCaller: L1ReverseCustomGatewayCaller{contract: contract}, L1ReverseCustomGatewayTransactor: L1ReverseCustomGatewayTransactor{contract: contract}, L1ReverseCustomGatewayFilterer: L1ReverseCustomGatewayFilterer{contract: contract}}, nil
}

// NewL1ReverseCustomGatewayCaller creates a new read-only instance of L1ReverseCustomGateway, bound to a specific deployed contract.
func NewL1ReverseCustomGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ReverseCustomGatewayCaller, error) {
	contract, err := bindL1ReverseCustomGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayCaller{contract: contract}, nil
}

// NewL1ReverseCustomGatewayTransactor creates a new write-only instance of L1ReverseCustomGateway, bound to a specific deployed contract.
func NewL1ReverseCustomGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ReverseCustomGatewayTransactor, error) {
	contract, err := bindL1ReverseCustomGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayTransactor{contract: contract}, nil
}

// NewL1ReverseCustomGatewayFilterer creates a new log filterer instance of L1ReverseCustomGateway, bound to a specific deployed contract.
func NewL1ReverseCustomGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ReverseCustomGatewayFilterer, error) {
	contract, err := bindL1ReverseCustomGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayFilterer{contract: contract}, nil
}

// bindL1ReverseCustomGateway binds a generic wrapper to an already deployed contract.
func bindL1ReverseCustomGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1ReverseCustomGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ReverseCustomGateway.Contract.L1ReverseCustomGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.L1ReverseCustomGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.L1ReverseCustomGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ReverseCustomGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) Counterpart() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Counterpart(&_L1ReverseCustomGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Counterpart(&_L1ReverseCustomGateway.CallOpts)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.GetL2ERC20Address(&_L1ReverseCustomGateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.GetL2ERC20Address(&_L1ReverseCustomGateway.CallOpts, _l1Token)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) Messenger() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Messenger(&_L1ReverseCustomGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) Messenger() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Messenger(&_L1ReverseCustomGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) Owner() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Owner(&_L1ReverseCustomGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) Owner() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Owner(&_L1ReverseCustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) Router() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Router(&_L1ReverseCustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) Router() (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.Router(&_L1ReverseCustomGateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ReverseCustomGateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.TokenMapping(&_L1ReverseCustomGateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1ReverseCustomGateway.Contract.TokenMapping(&_L1ReverseCustomGateway.CallOpts, arg0)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC20(&_L1ReverseCustomGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC20(&_L1ReverseCustomGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC200(&_L1ReverseCustomGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC200(&_L1ReverseCustomGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC20AndCall(&_L1ReverseCustomGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.DepositERC20AndCall(&_L1ReverseCustomGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.FinalizeWithdrawERC20(&_L1ReverseCustomGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.FinalizeWithdrawERC20(&_L1ReverseCustomGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.Initialize(&_L1ReverseCustomGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.Initialize(&_L1ReverseCustomGateway.TransactOpts, _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.OnDropMessage(&_L1ReverseCustomGateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.OnDropMessage(&_L1ReverseCustomGateway.TransactOpts, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.RenounceOwnership(&_L1ReverseCustomGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.RenounceOwnership(&_L1ReverseCustomGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.TransferOwnership(&_L1ReverseCustomGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.TransferOwnership(&_L1ReverseCustomGateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewaySession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.UpdateTokenMapping(&_L1ReverseCustomGateway.TransactOpts, _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayTransactorSession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ReverseCustomGateway.Contract.UpdateTokenMapping(&_L1ReverseCustomGateway.TransactOpts, _l1Token, _l2Token)
}

// L1ReverseCustomGatewayDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayDepositERC20Iterator struct {
	Event *L1ReverseCustomGatewayDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayDepositERC20)
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
		it.Event = new(L1ReverseCustomGatewayDepositERC20)
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
func (it *L1ReverseCustomGatewayDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayDepositERC20 represents a DepositERC20 event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0x1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad88672.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1ReverseCustomGatewayDepositERC20Iterator, error) {

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

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayDepositERC20Iterator{contract: _L1ReverseCustomGateway.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad88672.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayDepositERC20)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0x1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad88672.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseDepositERC20(log types.Log) (*L1ReverseCustomGatewayDepositERC20, error) {
	event := new(L1ReverseCustomGatewayDepositERC20)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator struct {
	Event *L1ReverseCustomGatewayFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayFinalizeWithdrawERC20)
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
		it.Event = new(L1ReverseCustomGatewayFinalizeWithdrawERC20)
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
func (it *L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayFinalizeWithdrawERC20 struct {
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
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayFinalizeWithdrawERC20Iterator{contract: _L1ReverseCustomGateway.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayFinalizeWithdrawERC20)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1ReverseCustomGatewayFinalizeWithdrawERC20, error) {
	event := new(L1ReverseCustomGatewayFinalizeWithdrawERC20)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ReverseCustomGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayInitializedIterator struct {
	Event *L1ReverseCustomGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayInitialized)
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
		it.Event = new(L1ReverseCustomGatewayInitialized)
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
func (it *L1ReverseCustomGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayInitialized represents a Initialized event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1ReverseCustomGatewayInitializedIterator, error) {

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayInitializedIterator{contract: _L1ReverseCustomGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayInitialized)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseInitialized(log types.Log) (*L1ReverseCustomGatewayInitialized, error) {
	event := new(L1ReverseCustomGatewayInitialized)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ReverseCustomGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayOwnershipTransferredIterator struct {
	Event *L1ReverseCustomGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayOwnershipTransferred)
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
		it.Event = new(L1ReverseCustomGatewayOwnershipTransferred)
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
func (it *L1ReverseCustomGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1ReverseCustomGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayOwnershipTransferredIterator{contract: _L1ReverseCustomGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayOwnershipTransferred)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1ReverseCustomGatewayOwnershipTransferred, error) {
	event := new(L1ReverseCustomGatewayOwnershipTransferred)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ReverseCustomGatewayRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayRefundERC20Iterator struct {
	Event *L1ReverseCustomGatewayRefundERC20 // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayRefundERC20)
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
		it.Event = new(L1ReverseCustomGatewayRefundERC20)
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
func (it *L1ReverseCustomGatewayRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayRefundERC20 represents a RefundERC20 event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayRefundERC20 struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterRefundERC20(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1ReverseCustomGatewayRefundERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayRefundERC20Iterator{contract: _L1ReverseCustomGateway.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayRefundERC20, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayRefundERC20)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseRefundERC20(log types.Log) (*L1ReverseCustomGatewayRefundERC20, error) {
	event := new(L1ReverseCustomGatewayRefundERC20)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ReverseCustomGatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayUpdateTokenMappingIterator struct {
	Event *L1ReverseCustomGatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L1ReverseCustomGatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ReverseCustomGatewayUpdateTokenMapping)
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
		it.Event = new(L1ReverseCustomGatewayUpdateTokenMapping)
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
func (it *L1ReverseCustomGatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ReverseCustomGatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ReverseCustomGatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1ReverseCustomGateway contract.
type L1ReverseCustomGatewayUpdateTokenMapping struct {
	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (*L1ReverseCustomGatewayUpdateTokenMappingIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var oldL2TokenRule []interface{}
	for _, oldL2TokenItem := range oldL2Token {
		oldL2TokenRule = append(oldL2TokenRule, oldL2TokenItem)
	}
	var newL2TokenRule []interface{}
	for _, newL2TokenItem := range newL2Token {
		newL2TokenRule = append(newL2TokenRule, newL2TokenItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.FilterLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ReverseCustomGatewayUpdateTokenMappingIterator{contract: _L1ReverseCustomGateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L1ReverseCustomGatewayUpdateTokenMapping, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var oldL2TokenRule []interface{}
	for _, oldL2TokenItem := range oldL2Token {
		oldL2TokenRule = append(oldL2TokenRule, oldL2TokenItem)
	}
	var newL2TokenRule []interface{}
	for _, newL2TokenItem := range newL2Token {
		newL2TokenRule = append(newL2TokenRule, newL2TokenItem)
	}

	logs, sub, err := _L1ReverseCustomGateway.contract.WatchLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ReverseCustomGatewayUpdateTokenMapping)
				if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1ReverseCustomGateway *L1ReverseCustomGatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L1ReverseCustomGatewayUpdateTokenMapping, error) {
	event := new(L1ReverseCustomGatewayUpdateTokenMapping)
	if err := _L1ReverseCustomGateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
