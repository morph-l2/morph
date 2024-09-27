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

// L2ERC1155GatewayMetaData contains all meta data concerning the L2ERC1155Gateway contract.
var L2ERC1155GatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchWithdrawERC1155\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"batchWithdrawERC1155\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"counterpart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeBatchDepositERC1155\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeDepositERC1155\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_counterpart\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenMapping\",\"inputs\":[{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawERC1155\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC1155\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BatchWithdrawERC1155\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeBatchDepositERC1155\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeDepositERC1155\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"inputs\":[{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldL1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newL1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawERC1155\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b62000021565b620000df565b5f54610100900460ff16156200008d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614620000dd575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61262880620000ed5f395ff3fe608060405260043610610123575f3560e01c80638c23d5b2116100a1578063eaa72ad911610071578063f2fde38b11610057578063f2fde38b146103be578063f887ea40146103dd578063fac752eb1461040a575f80fd5b8063eaa72ad91461035b578063f23a6e611461037a575f80fd5b80638c23d5b2146102675780638da5cb5b1461027a578063ba27f50b146102a4578063bc197c81146102e6575f80fd5b80634764cc62116100f657806348de03de116100dc57806348de03de14610213578063715018a614610226578063797594b01461023a575f80fd5b80634764cc62146101d5578063485cc955146101f4575f80fd5b806301ffc9a7146101275780630f2da0801461015b57806321fedfc9146101705780633cb747bf14610183575b5f80fd5b348015610132575f80fd5b50610146610141366004611d8f565b610429565b60405190151581526020015b60405180910390f35b61016e610169366004611df6565b6104c1565b005b61016e61017e366004611e2e565b6104d4565b34801561018e575f80fd5b5061012f546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610152565b3480156101e0575f80fd5b5061016e6101ef366004611e7b565b6104e8565b3480156101ff575f80fd5b5061016e61020e366004611ee5565b610888565b61016e610221366004611f64565b610a29565b348015610231575f80fd5b5061016e610a40565b348015610245575f80fd5b5061012d546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b61016e610275366004611fe9565b610a53565b348015610285575f80fd5b5060fb5473ffffffffffffffffffffffffffffffffffffffff166101b0565b3480156102af575f80fd5b506101b06102be366004612080565b61015e6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f1575f80fd5b5061032a61030036600461221e565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610152565b348015610366575f80fd5b5061016e6103753660046122c5565b610a62565b348015610385575f80fd5b5061032a610394366004612378565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b3480156103c9575f80fd5b5061016e6103d8366004612080565b610e0f565b3480156103e8575f80fd5b5061012e546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b348015610415575f80fd5b5061016e610424366004611ee5565b610ec6565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e00000000000000000000000000000000000000000000000000000000014806104bb57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ce8433858585610fd4565b50505050565b6104e18585858585610fd4565b5050505050565b61012f5473ffffffffffffffffffffffffffffffffffffffff16338114610570576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105dd91906123dc565b61012d5473ffffffffffffffffffffffffffffffffffffffff908116911614610662576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610567565b61066a61132a565b73ffffffffffffffffffffffffffffffffffffffff87166106e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8087165f90815261015e602052604090205488821691161461077a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d617463680000000000000000000000000000006044820152606401610567565b6040517f731133e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260448201849052608060648301525f608483015287169063731133e99060a4015f604051808303815f87803b1580156107fb575f80fd5b505af115801561080d573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff88811682526020820188905291810186905281891693508982169250908a16907f5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef99060600160405180910390a461087f6001609755565b50505050505050565b5f54610100900460ff16158080156108a657505f54600160ff909116105b806108bf5750303b1580156108bf57505f5460ff166001145b61094b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156109a7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6109af6113a4565b6109b76113a4565b6109c2835f8461143a565b8015610a24575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610a38863387878787876115e8565b505050505050565b610a48611a27565b610a515f611aa8565b565b61087f878787878787876115e8565b61012f5473ffffffffffffffffffffffffffffffffffffffff16338114610ae5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610567565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5291906123dc565b61012d5473ffffffffffffffffffffffffffffffffffffffff908116911614610bd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610567565b610bdf61132a565b73ffffffffffffffffffffffffffffffffffffffff8916610c5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8089165f90815261015e60205260409020548a8216911614610cef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d617463680000000000000000000000000000006044820152606401610567565b6040517fb48ab8b600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89169063b48ab8b690610d499089908990899089908990600401612440565b5f604051808303815f87803b158015610d60575f80fd5b505af1158015610d72573d5f803e3d5ffd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167ff07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b77568989898989604051610df295949392919061249f565b60405180910390a4610e046001609755565b505050505050505050565b610e17611a27565b73ffffffffffffffffffffffffffffffffffffffff8116610eba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610567565b610ec381611aa8565b50565b610ece611a27565b73ffffffffffffffffffffffffffffffffffffffff8116610f4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815261015e602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610fdc61132a565b5f8211611045576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815261015e602052604090205416806110d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e000000000000006044820152606401610567565b5f336040517ff5298aca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808316600483015260248201889052604482018790529192509088169063f5298aca906064015f604051808303815f87803b15801561114e575f80fd5b505af1158015611160573d5f803e3d5ffd5b505060405173ffffffffffffffffffffffffffffffffffffffff8086166024830152808b16604483015280851660648301528916608482015260a4810188905260c481018790525f925060e4019050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f730608b30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9081169263b2267a7b92349261128b929116905f9087908b906004016124ed565b5f604051808303818588803b1580156112a2575f80fd5b505af11580156112b4573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8c81168252602082018c90529181018a905281871694508c8216935090871691507f1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a9060600160405180910390a45050506104e16001609755565b600260975403611396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610567565b6002609755565b6001609755565b5f54610100900460ff16610a51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b73ffffffffffffffffffffffffffffffffffffffff83166114b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8116611534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610567565b61153c611b1e565b611544611bbc565b61012d805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925561012f80548484169216919091179055821615610a245761012e805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6115f061132a565b83611657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f20746f6b656e20746f2077697468647261770000000000000000000000006044820152606401610567565b8382146116c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610567565b5f5b82811015611753575f8484838181106116dd576116dd612583565b905060200201351161174b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610567565b6001016116c2565b5073ffffffffffffffffffffffffffffffffffffffff8088165f90815261015e602052604090205416806117e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e000000000000006044820152606401610567565b6040517ff6eb127a000000000000000000000000000000000000000000000000000000008152339073ffffffffffffffffffffffffffffffffffffffff8a169063f6eb127a9061183f9084908b908b908b908b9060040161249f565b5f604051808303815f87803b158015611856575f80fd5b505af1158015611868573d5f803e3d5ffd5b505050505f828a838b8b8b8b8b60405160240161188c9897969594939291906125b0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff92748d30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9081169263b2267a7b923492611968929116905f9087908b906004016124ed565b5f604051808303818588803b15801561197f575f80fd5b505af1158015611991573d5f803e3d5ffd5b50505050508173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f122909708c8c8c8c8c604051611a1295949392919061249f565b60405180910390a450505061087f6001609755565b60fb5473ffffffffffffffffffffffffffffffffffffffff163314610a51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b60fb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611bb4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a51611c5a565b5f54610100900460ff16611c52576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a51611cf0565b5f54610100900460ff1661139d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b5f54610100900460ff16611d86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a5133611aa8565b5f60208284031215611d9f575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611dce575f80fd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610ec3575f80fd5b5f805f8060808587031215611e09575f80fd5b8435611e1481611dd5565b966020860135965060408601359560600135945092505050565b5f805f805f60a08688031215611e42575f80fd5b8535611e4d81611dd5565b94506020860135611e5d81611dd5565b94979496505050506040830135926060810135926080909101359150565b5f805f805f8060c08789031215611e90575f80fd5b8635611e9b81611dd5565b95506020870135611eab81611dd5565b94506040870135611ebb81611dd5565b93506060870135611ecb81611dd5565b9598949750929560808101359460a0909101359350915050565b5f8060408385031215611ef6575f80fd5b8235611f0181611dd5565b91506020830135611f1181611dd5565b809150509250929050565b5f8083601f840112611f2c575f80fd5b50813567ffffffffffffffff811115611f43575f80fd5b6020830191508360208260051b8501011115611f5d575f80fd5b9250929050565b5f805f805f8060808789031215611f79575f80fd5b8635611f8481611dd5565b9550602087013567ffffffffffffffff80821115611fa0575f80fd5b611fac8a838b01611f1c565b90975095506040890135915080821115611fc4575f80fd5b50611fd189828a01611f1c565b979a9699509497949695606090950135949350505050565b5f805f805f805f60a0888a031215611fff575f80fd5b873561200a81611dd5565b9650602088013561201a81611dd5565b9550604088013567ffffffffffffffff80821115612036575f80fd5b6120428b838c01611f1c565b909750955060608a013591508082111561205a575f80fd5b506120678a828b01611f1c565b989b979a50959894979596608090950135949350505050565b5f60208284031215612090575f80fd5b8135611dce81611dd5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561210f5761210f61209b565b604052919050565b5f82601f830112612126575f80fd5b8135602067ffffffffffffffff8211156121425761214261209b565b8160051b6121518282016120c8565b928352848101820192828101908785111561216a575f80fd5b83870192505b8483101561218957823582529183019190830190612170565b979650505050505050565b5f82601f8301126121a3575f80fd5b813567ffffffffffffffff8111156121bd576121bd61209b565b6121ee60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016120c8565b818152846020838601011115612202575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215612232575f80fd5b853561223d81611dd5565b9450602086013561224d81611dd5565b9350604086013567ffffffffffffffff80821115612269575f80fd5b61227589838a01612117565b9450606088013591508082111561228a575f80fd5b61229689838a01612117565b935060808801359150808211156122ab575f80fd5b506122b888828901612194565b9150509295509295909350565b5f805f805f805f8060c0898b0312156122dc575f80fd5b88356122e781611dd5565b975060208901356122f781611dd5565b9650604089013561230781611dd5565b9550606089013561231781611dd5565b9450608089013567ffffffffffffffff80821115612333575f80fd5b61233f8c838d01611f1c565b909650945060a08b0135915080821115612357575f80fd5b506123648b828c01611f1c565b999c989b5096995094979396929594505050565b5f805f805f60a0868803121561238c575f80fd5b853561239781611dd5565b945060208601356123a781611dd5565b93506040860135925060608601359150608086013567ffffffffffffffff8111156123d0575f80fd5b6122b888828901612194565b5f602082840312156123ec575f80fd5b8151611dce81611dd5565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115612427575f80fd5b8260051b80836020870137939093016020019392505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201525f61246f6080830186886123f7565b82810360408401526124828185876123f7565b83810360609094019390935250505f815260200195945050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152606060208201525f6124ce6060830186886123f7565b82810360408401526124e18185876123f7565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b8181101561253c5786810183015185820160a001528201612520565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f73ffffffffffffffffffffffffffffffffffffffff808b168352808a166020840152808916604084015280881660608401525060c060808301526125f960c0830186886123f7565b82810360a084015261260c8185876123f7565b9b9a505050505050505050505056fea164736f6c6343000818000a",
}

// L2ERC1155GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ERC1155GatewayMetaData.ABI instead.
var L2ERC1155GatewayABI = L2ERC1155GatewayMetaData.ABI

// L2ERC1155GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ERC1155GatewayMetaData.Bin instead.
var L2ERC1155GatewayBin = L2ERC1155GatewayMetaData.Bin

// DeployL2ERC1155Gateway deploys a new Ethereum contract, binding an instance of L2ERC1155Gateway to it.
func DeployL2ERC1155Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ERC1155Gateway, error) {
	parsed, err := L2ERC1155GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ERC1155GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ERC1155Gateway{L2ERC1155GatewayCaller: L2ERC1155GatewayCaller{contract: contract}, L2ERC1155GatewayTransactor: L2ERC1155GatewayTransactor{contract: contract}, L2ERC1155GatewayFilterer: L2ERC1155GatewayFilterer{contract: contract}}, nil
}

// L2ERC1155Gateway is an auto generated Go binding around an Ethereum contract.
type L2ERC1155Gateway struct {
	L2ERC1155GatewayCaller     // Read-only binding to the contract
	L2ERC1155GatewayTransactor // Write-only binding to the contract
	L2ERC1155GatewayFilterer   // Log filterer for contract events
}

// L2ERC1155GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC1155GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC1155GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC1155GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC1155GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ERC1155GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC1155GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ERC1155GatewaySession struct {
	Contract     *L2ERC1155Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ERC1155GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ERC1155GatewayCallerSession struct {
	Contract *L2ERC1155GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2ERC1155GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ERC1155GatewayTransactorSession struct {
	Contract     *L2ERC1155GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2ERC1155GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ERC1155GatewayRaw struct {
	Contract *L2ERC1155Gateway // Generic contract binding to access the raw methods on
}

// L2ERC1155GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ERC1155GatewayCallerRaw struct {
	Contract *L2ERC1155GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2ERC1155GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ERC1155GatewayTransactorRaw struct {
	Contract *L2ERC1155GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ERC1155Gateway creates a new instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155Gateway(address common.Address, backend bind.ContractBackend) (*L2ERC1155Gateway, error) {
	contract, err := bindL2ERC1155Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155Gateway{L2ERC1155GatewayCaller: L2ERC1155GatewayCaller{contract: contract}, L2ERC1155GatewayTransactor: L2ERC1155GatewayTransactor{contract: contract}, L2ERC1155GatewayFilterer: L2ERC1155GatewayFilterer{contract: contract}}, nil
}

// NewL2ERC1155GatewayCaller creates a new read-only instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ERC1155GatewayCaller, error) {
	contract, err := bindL2ERC1155Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayCaller{contract: contract}, nil
}

// NewL2ERC1155GatewayTransactor creates a new write-only instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC1155GatewayTransactor, error) {
	contract, err := bindL2ERC1155Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayTransactor{contract: contract}, nil
}

// NewL2ERC1155GatewayFilterer creates a new log filterer instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ERC1155GatewayFilterer, error) {
	contract, err := bindL2ERC1155Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayFilterer{contract: contract}, nil
}

// bindL2ERC1155Gateway binds a generic wrapper to an already deployed contract.
func bindL2ERC1155Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ERC1155GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC1155Gateway *L2ERC1155GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC1155Gateway.Contract.L2ERC1155GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC1155Gateway *L2ERC1155GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.L2ERC1155GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC1155Gateway *L2ERC1155GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.L2ERC1155GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC1155Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) Counterpart() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Counterpart(&_L2ERC1155Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Counterpart(&_L2ERC1155Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) Messenger() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Messenger(&_L2ERC1155Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) Messenger() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Messenger(&_L2ERC1155Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) Owner() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Owner(&_L2ERC1155Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) Owner() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Owner(&_L2ERC1155Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) Router() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Router(&_L2ERC1155Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) Router() (common.Address, error) {
	return _L2ERC1155Gateway.Contract.Router(&_L2ERC1155Gateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2ERC1155Gateway.Contract.SupportsInterface(&_L2ERC1155Gateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2ERC1155Gateway.Contract.SupportsInterface(&_L2ERC1155Gateway.CallOpts, interfaceId)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ERC1155Gateway.Contract.TokenMapping(&_L2ERC1155Gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2ERC1155Gateway.Contract.TokenMapping(&_L2ERC1155Gateway.CallOpts, arg0)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) BatchWithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "batchWithdrawERC1155", _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) BatchWithdrawERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.BatchWithdrawERC1155(&_L2ERC1155Gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) BatchWithdrawERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.BatchWithdrawERC1155(&_L2ERC1155Gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) BatchWithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "batchWithdrawERC11550", _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) BatchWithdrawERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.BatchWithdrawERC11550(&_L2ERC1155Gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) BatchWithdrawERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.BatchWithdrawERC11550(&_L2ERC1155Gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) FinalizeBatchDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "finalizeBatchDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) FinalizeBatchDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.FinalizeBatchDepositERC1155(&_L2ERC1155Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) FinalizeBatchDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.FinalizeBatchDepositERC1155(&_L2ERC1155Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) FinalizeDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "finalizeDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) FinalizeDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.FinalizeDepositERC1155(&_L2ERC1155Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) FinalizeDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.FinalizeDepositERC1155(&_L2ERC1155Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.Initialize(&_L2ERC1155Gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.Initialize(&_L2ERC1155Gateway.TransactOpts, _counterpart, _messenger)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.OnERC1155BatchReceived(&_L2ERC1155Gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.OnERC1155BatchReceived(&_L2ERC1155Gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.OnERC1155Received(&_L2ERC1155Gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.OnERC1155Received(&_L2ERC1155Gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.RenounceOwnership(&_L2ERC1155Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.RenounceOwnership(&_L2ERC1155Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.TransferOwnership(&_L2ERC1155Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.TransferOwnership(&_L2ERC1155Gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.UpdateTokenMapping(&_L2ERC1155Gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.UpdateTokenMapping(&_L2ERC1155Gateway.TransactOpts, _l2Token, _l1Token)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) WithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "withdrawERC1155", _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) WithdrawERC1155(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.WithdrawERC1155(&_L2ERC1155Gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) WithdrawERC1155(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.WithdrawERC1155(&_L2ERC1155Gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) WithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "withdrawERC11550", _token, _to, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewaySession) WithdrawERC11550(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.WithdrawERC11550(&_L2ERC1155Gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactorSession) WithdrawERC11550(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.Contract.WithdrawERC11550(&_L2ERC1155Gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// L2ERC1155GatewayBatchWithdrawERC1155Iterator is returned from FilterBatchWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC1155 events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayBatchWithdrawERC1155Iterator struct {
	Event *L2ERC1155GatewayBatchWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayBatchWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayBatchWithdrawERC1155)
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
		it.Event = new(L2ERC1155GatewayBatchWithdrawERC1155)
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
func (it *L2ERC1155GatewayBatchWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayBatchWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayBatchWithdrawERC1155 represents a BatchWithdrawERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayBatchWithdrawERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchWithdrawERC1155 is a free log retrieval operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterBatchWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC1155GatewayBatchWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayBatchWithdrawERC1155Iterator{contract: _L2ERC1155Gateway.contract, event: "BatchWithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchBatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayBatchWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayBatchWithdrawERC1155)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
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

// ParseBatchWithdrawERC1155 is a log parse operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseBatchWithdrawERC1155(log types.Log) (*L2ERC1155GatewayBatchWithdrawERC1155, error) {
	event := new(L2ERC1155GatewayBatchWithdrawERC1155)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator is returned from FilterFinalizeBatchDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC1155 events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator struct {
	Event *L2ERC1155GatewayFinalizeBatchDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayFinalizeBatchDepositERC1155)
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
		it.Event = new(L2ERC1155GatewayFinalizeBatchDepositERC1155)
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
func (it *L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayFinalizeBatchDepositERC1155 represents a FinalizeBatchDepositERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeBatchDepositERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchDepositERC1155 is a free log retrieval operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterFinalizeBatchDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayFinalizeBatchDepositERC1155Iterator{contract: _L2ERC1155Gateway.contract, event: "FinalizeBatchDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC1155 is a free log subscription operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchFinalizeBatchDepositERC1155(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayFinalizeBatchDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayFinalizeBatchDepositERC1155)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
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

// ParseFinalizeBatchDepositERC1155 is a log parse operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseFinalizeBatchDepositERC1155(log types.Log) (*L2ERC1155GatewayFinalizeBatchDepositERC1155, error) {
	event := new(L2ERC1155GatewayFinalizeBatchDepositERC1155)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayFinalizeDepositERC1155Iterator is returned from FilterFinalizeDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC1155 events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeDepositERC1155Iterator struct {
	Event *L2ERC1155GatewayFinalizeDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayFinalizeDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayFinalizeDepositERC1155)
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
		it.Event = new(L2ERC1155GatewayFinalizeDepositERC1155)
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
func (it *L2ERC1155GatewayFinalizeDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayFinalizeDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayFinalizeDepositERC1155 represents a FinalizeDepositERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeDepositERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC1155 is a free log retrieval operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterFinalizeDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC1155GatewayFinalizeDepositERC1155Iterator, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayFinalizeDepositERC1155Iterator{contract: _L2ERC1155Gateway.contract, event: "FinalizeDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC1155 is a free log subscription operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchFinalizeDepositERC1155(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayFinalizeDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayFinalizeDepositERC1155)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
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

// ParseFinalizeDepositERC1155 is a log parse operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseFinalizeDepositERC1155(log types.Log) (*L2ERC1155GatewayFinalizeDepositERC1155, error) {
	event := new(L2ERC1155GatewayFinalizeDepositERC1155)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayInitializedIterator struct {
	Event *L2ERC1155GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayInitialized)
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
		it.Event = new(L2ERC1155GatewayInitialized)
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
func (it *L2ERC1155GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayInitialized represents a Initialized event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ERC1155GatewayInitializedIterator, error) {

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayInitializedIterator{contract: _L2ERC1155Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayInitialized)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseInitialized(log types.Log) (*L2ERC1155GatewayInitialized, error) {
	event := new(L2ERC1155GatewayInitialized)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayOwnershipTransferredIterator struct {
	Event *L2ERC1155GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayOwnershipTransferred)
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
		it.Event = new(L2ERC1155GatewayOwnershipTransferred)
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
func (it *L2ERC1155GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2ERC1155GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayOwnershipTransferredIterator{contract: _L2ERC1155Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayOwnershipTransferred)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2ERC1155GatewayOwnershipTransferred, error) {
	event := new(L2ERC1155GatewayOwnershipTransferred)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayUpdateTokenMappingIterator struct {
	Event *L2ERC1155GatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayUpdateTokenMapping)
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
		it.Event = new(L2ERC1155GatewayUpdateTokenMapping)
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
func (it *L2ERC1155GatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2ERC1155GatewayUpdateTokenMappingIterator, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayUpdateTokenMappingIterator{contract: _L2ERC1155Gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayUpdateTokenMapping)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2ERC1155GatewayUpdateTokenMapping, error) {
	event := new(L2ERC1155GatewayUpdateTokenMapping)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC1155GatewayWithdrawERC1155Iterator is returned from FilterWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawERC1155 events raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayWithdrawERC1155Iterator struct {
	Event *L2ERC1155GatewayWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *L2ERC1155GatewayWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC1155GatewayWithdrawERC1155)
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
		it.Event = new(L2ERC1155GatewayWithdrawERC1155)
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
func (it *L2ERC1155GatewayWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC1155GatewayWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC1155GatewayWithdrawERC1155 represents a WithdrawERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayWithdrawERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC1155 is a free log retrieval operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) FilterWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ERC1155GatewayWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.FilterLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayWithdrawERC1155Iterator{contract: _L2ERC1155Gateway.contract, event: "WithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) WatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *L2ERC1155GatewayWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ERC1155Gateway.contract.WatchLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC1155GatewayWithdrawERC1155)
				if err := _L2ERC1155Gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
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

// ParseWithdrawERC1155 is a log parse operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2ERC1155Gateway *L2ERC1155GatewayFilterer) ParseWithdrawERC1155(log types.Log) (*L2ERC1155GatewayWithdrawERC1155, error) {
	event := new(L2ERC1155GatewayWithdrawERC1155)
	if err := _L2ERC1155Gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
