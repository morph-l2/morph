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

// L2ERC721GatewayMetaData contains all meta data concerning the L2ERC721Gateway contract.
var L2ERC721GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC721\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611eab806100ec6000396000f3fe6080604052600436106100f35760003560e01c8063982b151f1161008a578063f2fde38b11610059578063f2fde38b146102c1578063f887ea40146102e1578063f8c3cf2514610301578063fac752eb1461032157600080fd5b8063982b151f14610245578063aa4c115814610265578063ba27f50b14610278578063ee5a8db2146102ae57600080fd5b8063485cc955116100c6578063485cc955146101d2578063715018a6146101f2578063797594b0146102075780638da5cb5b1461022757600080fd5b8063150b7a02146100f85780632a491247146101725780633cb747bf1461018757806346aa3411146101bf575b600080fd5b34801561010457600080fd5b5061013c6101133660046118a2565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b6101856101803660046119a0565b610341565b005b34801561019357600080fd5b5060cb546101a7906001600160a01b031681565b6040516001600160a01b039091168152602001610169565b6101856101cd366004611a21565b610352565b3480156101de57600080fd5b506101856101ed366004611a7d565b610365565b3480156101fe57600080fd5b506101856104ef565b34801561021357600080fd5b5060c9546101a7906001600160a01b031681565b34801561023357600080fd5b506097546001600160a01b03166101a7565b34801561025157600080fd5b50610185610260366004611ab6565b610503565b610185610273366004611b44565b610813565b34801561028457600080fd5b506101a7610293366004611bb1565b60fb602052600090815260409020546001600160a01b031681565b6101856102bc366004611bd5565b610827565b3480156102cd57600080fd5b506101856102dc366004611bb1565b610833565b3480156102ed57600080fd5b5060ca546101a7906001600160a01b031681565b34801561030d57600080fd5b5061018561031c366004611c1b565b6108c3565b34801561032d57600080fd5b5061018561033c366004611a7d565b610b82565b61034d83338484610c5c565b505050565b61035f8433858585610fae565b50505050565b600054610100900460ff16158080156103855750600054600160ff909116105b8061039f5750303b15801561039f575060005460ff166001145b6104165760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561047457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61047c61136d565b610488836000846113ea565b801561034d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b6104f761152d565b6105016000611587565b565b60cb546001600160a01b031633811461055e5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161040d565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561059c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c09190611c7f565b60c9546001600160a01b0390811691161461061d5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161040d565b6106256115f1565b6001600160a01b03871661067b5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161040d565b6001600160a01b03808716600090815260fb60205260409020548882169116146106e75760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161040d565b60005b828110156107a657866001600160a01b03166340c10f198686868581811061071457610714611c9c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526001600160a01b0390941660048501526020029190910135602483015250604401600060405180830381600087803b15801561077b57600080fd5b505af115801561078f573d6000803e3d6000fd5b50505050808061079e90611ccb565b9150506106ea565b50846001600160a01b0316866001600160a01b0316886001600160a01b03167fafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e8787876040516107f893929190611d75565b60405180910390a461080a6001603355565b50505050505050565b6108208585858585610fae565b5050505050565b61035f84848484610c5c565b61083b61152d565b6001600160a01b0381166108b75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161040d565b6108c081611587565b50565b60cb546001600160a01b031633811461091e5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161040d565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561095c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109809190611c7f565b60c9546001600160a01b039081169116146109dd5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161040d565b6109e56115f1565b6001600160a01b038616610a3b5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161040d565b6001600160a01b03808616600090815260fb6020526040902054878216911614610aa75760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161040d565b6040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018490528616906340c10f1990604401600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b5050604080516001600160a01b03878116825260208201879052808916945089811693508a16917fc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41910160405180910390a4610b7a6001603355565b505050505050565b610b8a61152d565b6001600160a01b038116610be05760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161040d565b6001600160a01b03808316600081815260fb602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610c646115f1565b6001600160a01b03808516600090815260fb60205260409020541680610ccc5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e00000000000000604482015260640161040d565b6040517f6352211e00000000000000000000000000000000000000000000000000000000815260048101849052339081906001600160a01b03881690636352211e90602401602060405180830381865afa158015610d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d529190611c7f565b6001600160a01b031614610da85760405162461bcd60e51b815260206004820152600f60248201527f746f6b656e206e6f74206f776e65640000000000000000000000000000000000604482015260640161040d565b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b038716906342966c6890602401600060405180830381600087803b158015610e0357600080fd5b505af1158015610e17573d6000803e3d6000fd5b50506040516001600160a01b038086166024830152808a16604483015280851660648301528816608482015260a481018790526000925060c4019050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd606b4dc0000000000000000000000000000000000000000000000000000000017905260cb5460c95491517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039081169263b2267a7b923492610f219291169060009087908b90600401611da1565b6000604051808303818588803b158015610f3a57600080fd5b505af1158015610f4e573d6000803e3d6000fd5b5050604080516001600160a01b038b81168252602082018b905280881695508c81169450881692507fe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda910160405180910390a450505061035f6001603355565b610fb66115f1565b816110035760405162461bcd60e51b815260206004820152601460248201527f6e6f20746f6b656e20746f207769746864726177000000000000000000000000604482015260640161040d565b6001600160a01b03808616600090815260fb6020526040902054168061106b5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e00000000000000604482015260640161040d565b3360005b848110156111e957816001600160a01b0316886001600160a01b0316636352211e8888858181106110a2576110a2611c9c565b905060200201356040518263ffffffff1660e01b81526004016110c791815260200190565b602060405180830381865afa1580156110e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111089190611c7f565b6001600160a01b03161461115e5760405162461bcd60e51b815260206004820152600f60248201527f746f6b656e206e6f74206f776e65640000000000000000000000000000000000604482015260640161040d565b876001600160a01b03166342966c6887878481811061117f5761117f611c9c565b905060200201356040518263ffffffff1660e01b81526004016111a491815260200190565b600060405180830381600087803b1580156111be57600080fd5b505af11580156111d2573d6000803e3d6000fd5b5050505080806111e190611ccb565b91505061106f565b50600082888389898960405160240161120796959493929190611e2c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9f0a68b30000000000000000000000000000000000000000000000000000000017905260cb5460c95491517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039081169263b2267a7b9234926112d59291169060009087908b90600401611da1565b6000604051808303818588803b1580156112ee57600080fd5b505af1158015611302573d6000803e3d6000fd5b5050505050816001600160a01b0316886001600160a01b0316846001600160a01b03167fbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d8a8a8a60405161135893929190611d75565b60405180910390a45050506108206001603355565b600054610100900460ff166105015760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161040d565b6001600160a01b0383166114405760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161040d565b6001600160a01b0381166114965760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161040d565b61149e611651565b6114a66116d6565b60c980546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560cb8054848416921691909117905582161561034d5760ca80546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6097546001600160a01b031633146105015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161040d565b609780546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6002603354036116435760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161040d565b6002603355565b6001603355565b600054610100900460ff166116ce5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161040d565b61050161175b565b600054610100900460ff166117535760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161040d565b6105016117d8565b600054610100900460ff1661164a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161040d565b600054610100900460ff166118555760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161040d565b61050133611587565b6001600160a01b03811681146108c057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600080608085870312156118b857600080fd5b84356118c38161185e565b935060208501356118d38161185e565b925060408501359150606085013567ffffffffffffffff808211156118f757600080fd5b818701915087601f83011261190b57600080fd5b81358181111561191d5761191d611873565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561196357611963611873565b816040528281528a602084870101111561197c57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806000606084860312156119b557600080fd5b83356119c08161185e565b95602085013595506040909401359392505050565b60008083601f8401126119e757600080fd5b50813567ffffffffffffffff8111156119ff57600080fd5b6020830191508360208260051b8501011115611a1a57600080fd5b9250929050565b60008060008060608587031215611a3757600080fd5b8435611a428161185e565b9350602085013567ffffffffffffffff811115611a5e57600080fd5b611a6a878288016119d5565b9598909750949560400135949350505050565b60008060408385031215611a9057600080fd5b8235611a9b8161185e565b91506020830135611aab8161185e565b809150509250929050565b60008060008060008060a08789031215611acf57600080fd5b8635611ada8161185e565b95506020870135611aea8161185e565b94506040870135611afa8161185e565b93506060870135611b0a8161185e565b9250608087013567ffffffffffffffff811115611b2657600080fd5b611b3289828a016119d5565b979a9699509497509295939492505050565b600080600080600060808688031215611b5c57600080fd5b8535611b678161185e565b94506020860135611b778161185e565b9350604086013567ffffffffffffffff811115611b9357600080fd5b611b9f888289016119d5565b96999598509660600135949350505050565b600060208284031215611bc357600080fd5b8135611bce8161185e565b9392505050565b60008060008060808587031215611beb57600080fd5b8435611bf68161185e565b93506020850135611c068161185e565b93969395505050506040820135916060013590565b600080600080600060a08688031215611c3357600080fd5b8535611c3e8161185e565b94506020860135611c4e8161185e565b93506040860135611c5e8161185e565b92506060860135611c6e8161185e565b949793965091946080013592915050565b600060208284031215611c9157600080fd5b8151611bce8161185e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d23577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611d5c57600080fd5b8260051b80836020870137939093016020019392505050565b6001600160a01b0384168152604060208201526000611d98604083018486611d2a565b95945050505050565b6001600160a01b038516815260006020858184015260806040840152845180608085015260005b81811015611de45786810183015185820160a001528201611dc8565b50600060a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b60006001600160a01b0380891683528088166020840152808716604084015280861660608401525060a06080830152611e6960a083018486611d2a565b9897505050505050505056fea264697066735822122024c60729f578a85f63b698c404a6c8b2252ed14134e56ff4b00168fa5af9c60c64736f6c63430008100033",
}

// L2ERC721GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ERC721GatewayMetaData.ABI instead.
var L2ERC721GatewayABI = L2ERC721GatewayMetaData.ABI

// L2ERC721GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ERC721GatewayMetaData.Bin instead.
var L2ERC721GatewayBin = L2ERC721GatewayMetaData.Bin

// DeployL2ERC721Gateway deploys a new Ethereum contract, binding an instance of L2ERC721Gateway to it.
func DeployL2ERC721Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ERC721Gateway, error) {
	parsed, err := L2ERC721GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ERC721GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ERC721Gateway{L2ERC721GatewayCaller: L2ERC721GatewayCaller{contract: contract}, L2ERC721GatewayTransactor: L2ERC721GatewayTransactor{contract: contract}, L2ERC721GatewayFilterer: L2ERC721GatewayFilterer{contract: contract}}, nil
}

// L2ERC721Gateway is an auto generated Go binding around an Ethereum contract.
type L2ERC721Gateway struct {
	L2ERC721GatewayCaller     // Read-only binding to the contract
	L2ERC721GatewayTransactor // Write-only binding to the contract
	L2ERC721GatewayFilterer   // Log filterer for contract events
}

// L2ERC721GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC721GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC721GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ERC721GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ERC721GatewaySession struct {
	Contract     *L2ERC721Gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ERC721GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ERC721GatewayCallerSession struct {
	Contract *L2ERC721GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2ERC721GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ERC721GatewayTransactorSession struct {
	Contract     *L2ERC721GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2ERC721GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ERC721GatewayRaw struct {
	Contract *L2ERC721Gateway // Generic contract binding to access the raw methods on
}

// L2ERC721GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ERC721GatewayCallerRaw struct {
	Contract *L2ERC721GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2ERC721GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ERC721GatewayTransactorRaw struct {
	Contract *L2ERC721GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ERC721Gateway creates a new instance of L2ERC721Gateway, bound to a specific deployed contract.
func NewL2ERC721Gateway(address common.Address, backend bind.ContractBackend) (*L2ERC721Gateway, error) {
	contract, err := bindL2ERC721Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ERC721Gateway{L2ERC721GatewayCaller: L2ERC721GatewayCaller{contract: contract}, L2ERC721GatewayTransactor: L2ERC721GatewayTransactor{contract: contract}, L2ERC721GatewayFilterer: L2ERC721GatewayFilterer{contract: contract}}, nil
}

// NewL2ERC721GatewayCaller creates a new read-only instance of L2ERC721Gateway, bound to a specific deployed contract.
func NewL2ERC721GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ERC721GatewayCaller, error) {
	contract, err := bindL2ERC721Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayCaller{contract: contract}, nil
}

// NewL2ERC721GatewayTransactor creates a new write-only instance of L2ERC721Gateway, bound to a specific deployed contract.
func NewL2ERC721GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC721GatewayTransactor, error) {
	contract, err := bindL2ERC721Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayTransactor{contract: contract}, nil
}

// NewL2ERC721GatewayFilterer creates a new log filterer instance of L2ERC721Gateway, bound to a specific deployed contract.
func NewL2ERC721GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ERC721GatewayFilterer, error) {
	contract, err := bindL2ERC721Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayFilterer{contract: contract}, nil
}

// bindL2ERC721Gateway binds a generic wrapper to an already deployed contract.
func bindL2ERC721Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC721GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Gateway *L2ERC721GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Gateway.Contract.L2ERC721GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Gateway *L2ERC721GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.L2ERC721GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Gateway *L2ERC721GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.L2ERC721GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Gateway *L2ERC721GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Gateway *L2ERC721GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Gateway *L2ERC721GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewaySession) Counterpart() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Counterpart(&_L2ERC721Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Counterpart(&_L2ERC721Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewaySession) Messenger() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Messenger(&_L2ERC721Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCallerSession) Messenger() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Messenger(&_L2ERC721Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewaySession) Owner() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Owner(&_L2ERC721Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCallerSession) Owner() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Owner(&_L2ERC721Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewaySession) Router() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Router(&_L2ERC721Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCallerSession) Router() (common.Address, error) {
	return _L2ERC721Gateway.Contract.Router(&_L2ERC721Gateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC721Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ERC721Gateway.Contract.TokenMapping(&_L2ERC721Gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC721Gateway *L2ERC721GatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ERC721Gateway.Contract.TokenMapping(&_L2ERC721Gateway.CallOpts, arg0)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) BatchWithdrawERC721(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "batchWithdrawERC721", _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) BatchWithdrawERC721(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.BatchWithdrawERC721(&_L2ERC721Gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) BatchWithdrawERC721(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.BatchWithdrawERC721(&_L2ERC721Gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) BatchWithdrawERC7210(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "batchWithdrawERC7210", _token, _to, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) BatchWithdrawERC7210(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.BatchWithdrawERC7210(&_L2ERC721Gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) BatchWithdrawERC7210(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.BatchWithdrawERC7210(&_L2ERC721Gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) FinalizeBatchDepositERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "finalizeBatchDepositERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) FinalizeBatchDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.FinalizeBatchDepositERC721(&_L2ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) FinalizeBatchDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.FinalizeBatchDepositERC721(&_L2ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) FinalizeDepositERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "finalizeDepositERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) FinalizeDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.FinalizeDepositERC721(&_L2ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) FinalizeDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.FinalizeDepositERC721(&_L2ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.Initialize(&_L2ERC721Gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.Initialize(&_L2ERC721Gateway.TransactOpts, _counterpart, _messenger)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2ERC721Gateway *L2ERC721GatewaySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.OnERC721Received(&_L2ERC721Gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.OnERC721Received(&_L2ERC721Gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.RenounceOwnership(&_L2ERC721Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.RenounceOwnership(&_L2ERC721Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.TransferOwnership(&_L2ERC721Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.TransferOwnership(&_L2ERC721Gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.UpdateTokenMapping(&_L2ERC721Gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.UpdateTokenMapping(&_L2ERC721Gateway.TransactOpts, _l2Token, _l1Token)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) WithdrawERC721(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "withdrawERC721", _token, _tokenId, _gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) WithdrawERC721(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.WithdrawERC721(&_L2ERC721Gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) WithdrawERC721(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.WithdrawERC721(&_L2ERC721Gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactor) WithdrawERC7210(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.contract.Transact(opts, "withdrawERC7210", _token, _to, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewaySession) WithdrawERC7210(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.WithdrawERC7210(&_L2ERC721Gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2ERC721Gateway *L2ERC721GatewayTransactorSession) WithdrawERC7210(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC721Gateway.Contract.WithdrawERC7210(&_L2ERC721Gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// L2ERC721GatewayBatchWithdrawERC721Iterator is returned from FilterBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC721 events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayBatchWithdrawERC721Iterator struct {
	Event *L2ERC721GatewayBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayBatchWithdrawERC721)
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
		it.Event = new(L2ERC721GatewayBatchWithdrawERC721)
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
func (it *L2ERC721GatewayBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayBatchWithdrawERC721 represents a BatchWithdrawERC721 event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayBatchWithdrawERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchWithdrawERC721 is a free log retrieval operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterBatchWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC721GatewayBatchWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayBatchWithdrawERC721Iterator{contract: _L2ERC721Gateway.contract, event: "BatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC721 is a free log subscription operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayBatchWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayBatchWithdrawERC721)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
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

// ParseBatchWithdrawERC721 is a log parse operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseBatchWithdrawERC721(log types.Log) (*L2ERC721GatewayBatchWithdrawERC721, error) {
	event := new(L2ERC721GatewayBatchWithdrawERC721)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayFinalizeBatchDepositERC721Iterator is returned from FilterFinalizeBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC721 events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayFinalizeBatchDepositERC721Iterator struct {
	Event *L2ERC721GatewayFinalizeBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayFinalizeBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayFinalizeBatchDepositERC721)
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
		it.Event = new(L2ERC721GatewayFinalizeBatchDepositERC721)
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
func (it *L2ERC721GatewayFinalizeBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayFinalizeBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayFinalizeBatchDepositERC721 represents a FinalizeBatchDepositERC721 event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayFinalizeBatchDepositERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchDepositERC721 is a free log retrieval operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterFinalizeBatchDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC721GatewayFinalizeBatchDepositERC721Iterator, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayFinalizeBatchDepositERC721Iterator{contract: _L2ERC721Gateway.contract, event: "FinalizeBatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC721 is a free log subscription operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchFinalizeBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayFinalizeBatchDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayFinalizeBatchDepositERC721)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
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

// ParseFinalizeBatchDepositERC721 is a log parse operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseFinalizeBatchDepositERC721(log types.Log) (*L2ERC721GatewayFinalizeBatchDepositERC721, error) {
	event := new(L2ERC721GatewayFinalizeBatchDepositERC721)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayFinalizeDepositERC721Iterator is returned from FilterFinalizeDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC721 events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayFinalizeDepositERC721Iterator struct {
	Event *L2ERC721GatewayFinalizeDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayFinalizeDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayFinalizeDepositERC721)
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
		it.Event = new(L2ERC721GatewayFinalizeDepositERC721)
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
func (it *L2ERC721GatewayFinalizeDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayFinalizeDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayFinalizeDepositERC721 represents a FinalizeDepositERC721 event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayFinalizeDepositERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC721 is a free log retrieval operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterFinalizeDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC721GatewayFinalizeDepositERC721Iterator, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayFinalizeDepositERC721Iterator{contract: _L2ERC721Gateway.contract, event: "FinalizeDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC721 is a free log subscription operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchFinalizeDepositERC721(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayFinalizeDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayFinalizeDepositERC721)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
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

// ParseFinalizeDepositERC721 is a log parse operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseFinalizeDepositERC721(log types.Log) (*L2ERC721GatewayFinalizeDepositERC721, error) {
	event := new(L2ERC721GatewayFinalizeDepositERC721)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayInitializedIterator struct {
	Event *L2ERC721GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayInitialized)
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
		it.Event = new(L2ERC721GatewayInitialized)
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
func (it *L2ERC721GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayInitialized represents a Initialized event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ERC721GatewayInitializedIterator, error) {

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayInitializedIterator{contract: _L2ERC721Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayInitialized)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseInitialized(log types.Log) (*L2ERC721GatewayInitialized, error) {
	event := new(L2ERC721GatewayInitialized)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayOwnershipTransferredIterator struct {
	Event *L2ERC721GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayOwnershipTransferred)
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
		it.Event = new(L2ERC721GatewayOwnershipTransferred)
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
func (it *L2ERC721GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2ERC721GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayOwnershipTransferredIterator{contract: _L2ERC721Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayOwnershipTransferred)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2ERC721GatewayOwnershipTransferred, error) {
	event := new(L2ERC721GatewayOwnershipTransferred)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayUpdateTokenMappingIterator struct {
	Event *L2ERC721GatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayUpdateTokenMapping)
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
		it.Event = new(L2ERC721GatewayUpdateTokenMapping)
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
func (it *L2ERC721GatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2ERC721GatewayUpdateTokenMappingIterator, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayUpdateTokenMappingIterator{contract: _L2ERC721Gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayUpdateTokenMapping)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2ERC721GatewayUpdateTokenMapping, error) {
	event := new(L2ERC721GatewayUpdateTokenMapping)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721GatewayWithdrawERC721Iterator is returned from FilterWithdrawERC721 and is used to iterate over the raw logs and unpacked data for WithdrawERC721 events raised by the L2ERC721Gateway contract.
type L2ERC721GatewayWithdrawERC721Iterator struct {
	Event *L2ERC721GatewayWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L2ERC721GatewayWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721GatewayWithdrawERC721)
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
		it.Event = new(L2ERC721GatewayWithdrawERC721)
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
func (it *L2ERC721GatewayWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721GatewayWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721GatewayWithdrawERC721 represents a WithdrawERC721 event raised by the L2ERC721Gateway contract.
type L2ERC721GatewayWithdrawERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC721 is a free log retrieval operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) FilterWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC721GatewayWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.FilterLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721GatewayWithdrawERC721Iterator{contract: _L2ERC721Gateway.contract, event: "WithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC721 is a free log subscription operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) WatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L2ERC721GatewayWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC721Gateway.contract.WatchLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721GatewayWithdrawERC721)
				if err := _L2ERC721Gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
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

// ParseWithdrawERC721 is a log parse operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2ERC721Gateway *L2ERC721GatewayFilterer) ParseWithdrawERC721(log types.Log) (*L2ERC721GatewayWithdrawERC721, error) {
	event := new(L2ERC721GatewayWithdrawERC721)
	if err := _L2ERC721Gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
