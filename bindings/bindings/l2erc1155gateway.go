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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC1155\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61212180620000e75f395ff3fe608060405260043610610123575f3560e01c80638c23d5b2116100a1578063eaa72ad911610071578063f2fde38b11610057578063f2fde38b1461037d578063f887ea401461039c578063fac752eb146103bc575f80fd5b8063eaa72ad91461031a578063f23a6e6114610339575f80fd5b80638c23d5b2146102405780638da5cb5b14610253578063ba27f50b14610270578063bc197c81146102a5575f80fd5b80634764cc62116100f657806348de03de116100dc57806348de03de146101f9578063715018a61461020c578063797594b014610220575f80fd5b80634764cc62146101bb578063485cc955146101da575f80fd5b806301ffc9a7146101275780630f2da0801461015b57806321fedfc9146101705780633cb747bf14610183575b5f80fd5b348015610132575f80fd5b506101466101413660046118c9565b6103db565b60405190151581526020015b60405180910390f35b61016e610169366004611923565b610473565b005b61016e61017e36600461195b565b610486565b34801561018e575f80fd5b5061012f546101a3906001600160a01b031681565b6040516001600160a01b039091168152602001610152565b3480156101c6575f80fd5b5061016e6101d53660046119a8565b61049a565b3480156101e5575f80fd5b5061016e6101f4366004611a12565b610777565b61016e610207366004611a91565b6108fe565b348015610217575f80fd5b5061016e610915565b34801561022b575f80fd5b5061012d546101a3906001600160a01b031681565b61016e61024e366004611b16565b610928565b34801561025e575f80fd5b5060fb546001600160a01b03166101a3565b34801561027b575f80fd5b506101a361028a366004611bad565b61015e6020525f90815260409020546001600160a01b031681565b3480156102b0575f80fd5b506102e96102bf366004611d4b565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610152565b348015610325575f80fd5b5061016e610334366004611df2565b610937565b348015610344575f80fd5b506102e9610353366004611ea5565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b348015610388575f80fd5b5061016e610397366004611bad565b610c07565b3480156103a7575f80fd5b5061012e546101a3906001600160a01b031681565b3480156103c7575f80fd5b5061016e6103d6366004611a12565b610c97565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e000000000000000000000000000000000000000000000000000000000148061046d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104808433858585610d71565b50505050565b6104938585858585610d71565b5050505050565b61012f546001600160a01b03163381146104fb5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610537573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055b9190611f09565b61012d546001600160a01b039081169116146105b95760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016104f2565b6105c1611052565b6001600160a01b0387166106175760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104f2565b6001600160a01b038087165f90815261015e60205260409020548882169116146106835760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016104f2565b6040517f731133e90000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820185905260448201849052608060648301525f608483015287169063731133e99060a4015f604051808303815f87803b1580156106f7575f80fd5b505af1158015610709573d5f803e3d5ffd5b5050604080516001600160a01b0388811682526020820188905291810186905281891693508982169250908a16907f5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef99060600160405180910390a461076e6001609755565b50505050505050565b5f54610100900460ff161580801561079557505f54600160ff909116105b806107ae5750303b1580156107ae57505f5460ff166001145b6108205760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104f2565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561087c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6108846110b2565b61088c6110b2565b610897835f8461112e565b80156108f9575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b61090d86338787878787611274565b505050505050565b61091d6115fd565b6109265f611657565b565b61076e87878787878787611274565b61012f546001600160a01b03163381146109935760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064016104f2565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f39190611f09565b61012d546001600160a01b03908116911614610a515760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016104f2565b610a59611052565b6001600160a01b038916610aaf5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104f2565b6001600160a01b038089165f90815261015e60205260409020548a8216911614610b1b5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016104f2565b6040517fb48ab8b60000000000000000000000000000000000000000000000000000000081526001600160a01b0389169063b48ab8b690610b689089908990899089908990600401611f6d565b5f604051808303815f87803b158015610b7f575f80fd5b505af1158015610b91573d5f803e3d5ffd5b50505050866001600160a01b0316886001600160a01b03168a6001600160a01b03167ff07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b77568989898989604051610bea959493929190611fbf565b60405180910390a4610bfc6001609755565b505050505050505050565b610c0f6115fd565b6001600160a01b038116610c8b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104f2565b610c9481611657565b50565b610c9f6115fd565b6001600160a01b038116610cf55760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104f2565b6001600160a01b038083165f81815261015e602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610d79611052565b5f8211610dc85760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e7400000000000000000000000060448201526064016104f2565b6001600160a01b038086165f90815261015e60205260409020541680610e305760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016104f2565b5f336040517ff5298aca0000000000000000000000000000000000000000000000000000000081526001600160a01b03808316600483015260248201889052604482018790529192509088169063f5298aca906064015f604051808303815f87803b158015610e9d575f80fd5b505af1158015610eaf573d5f803e3d5ffd5b50506040516001600160a01b038086166024830152808b16604483015280851660648301528916608482015260a4810188905260c481018790525f925060e4019050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f730608b30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039081169263b2267a7b923492610fc0929116905f9087908b90600401612000565b5f604051808303818588803b158015610fd7575f80fd5b505af1158015610fe9573d5f803e3d5ffd5b5050604080516001600160a01b038c81168252602082018c90529181018a905281871694508c8216935090871691507f1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a9060600160405180910390a45050506104936001609755565b6002609754036110a45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104f2565b6002609755565b6001609755565b5f54610100900460ff166109265760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f2565b6001600160a01b0383166111845760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016104f2565b6001600160a01b0381166111da5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016104f2565b6111e26116c0565b6111ea611744565b61012d80546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925561012f805484841692169190911790558216156108f95761012e80546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b61127c611052565b836112c95760405162461bcd60e51b815260206004820152601460248201527f6e6f20746f6b656e20746f20776974686472617700000000000000000000000060448201526064016104f2565b8382146113185760405162461bcd60e51b815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064016104f2565b5f5b82811015611391575f84848381811061133557611335612089565b90506020020135116113895760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e7400000000000000000000000060448201526064016104f2565b60010161131a565b506001600160a01b038088165f90815261015e602052604090205416806113fa5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016104f2565b6040517ff6eb127a00000000000000000000000000000000000000000000000000000000815233906001600160a01b038a169063f6eb127a906114499084908b908b908b908b90600401611fbf565b5f604051808303815f87803b158015611460575f80fd5b505af1158015611472573d5f803e3d5ffd5b505050505f828a838b8b8b8b8b6040516024016114969897969594939291906120b6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff92748d30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039081169263b2267a7b923492611565929116905f9087908b90600401612000565b5f604051808303818588803b15801561157c575f80fd5b505af115801561158e573d5f803e3d5ffd5b5050505050816001600160a01b03168a6001600160a01b0316846001600160a01b03167f5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f122909708c8c8c8c8c6040516115e8959493929190611fbf565b60405180910390a450505061076e6001609755565b60fb546001600160a01b031633146109265760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104f2565b60fb80546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff1661173c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f2565b6109266117c8565b5f54610100900460ff166117c05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f2565b610926611844565b5f54610100900460ff166110ab5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f2565b5f54610100900460ff166118c05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f2565b61092633611657565b5f602082840312156118d9575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611908575f80fd5b9392505050565b6001600160a01b0381168114610c94575f80fd5b5f805f8060808587031215611936575f80fd5b84356119418161190f565b966020860135965060408601359560600135945092505050565b5f805f805f60a0868803121561196f575f80fd5b853561197a8161190f565b9450602086013561198a8161190f565b94979496505050506040830135926060810135926080909101359150565b5f805f805f8060c087890312156119bd575f80fd5b86356119c88161190f565b955060208701356119d88161190f565b945060408701356119e88161190f565b935060608701356119f88161190f565b9598949750929560808101359460a0909101359350915050565b5f8060408385031215611a23575f80fd5b8235611a2e8161190f565b91506020830135611a3e8161190f565b809150509250929050565b5f8083601f840112611a59575f80fd5b50813567ffffffffffffffff811115611a70575f80fd5b6020830191508360208260051b8501011115611a8a575f80fd5b9250929050565b5f805f805f8060808789031215611aa6575f80fd5b8635611ab18161190f565b9550602087013567ffffffffffffffff80821115611acd575f80fd5b611ad98a838b01611a49565b90975095506040890135915080821115611af1575f80fd5b50611afe89828a01611a49565b979a9699509497949695606090950135949350505050565b5f805f805f805f60a0888a031215611b2c575f80fd5b8735611b378161190f565b96506020880135611b478161190f565b9550604088013567ffffffffffffffff80821115611b63575f80fd5b611b6f8b838c01611a49565b909750955060608a0135915080821115611b87575f80fd5b50611b948a828b01611a49565b989b979a50959894979596608090950135949350505050565b5f60208284031215611bbd575f80fd5b81356119088161190f565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611c3c57611c3c611bc8565b604052919050565b5f82601f830112611c53575f80fd5b8135602067ffffffffffffffff821115611c6f57611c6f611bc8565b8160051b611c7e828201611bf5565b9283528481018201928281019087851115611c97575f80fd5b83870192505b84831015611cb657823582529183019190830190611c9d565b979650505050505050565b5f82601f830112611cd0575f80fd5b813567ffffffffffffffff811115611cea57611cea611bc8565b611d1b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611bf5565b818152846020838601011115611d2f575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611d5f575f80fd5b8535611d6a8161190f565b94506020860135611d7a8161190f565b9350604086013567ffffffffffffffff80821115611d96575f80fd5b611da289838a01611c44565b94506060880135915080821115611db7575f80fd5b611dc389838a01611c44565b93506080880135915080821115611dd8575f80fd5b50611de588828901611cc1565b9150509295509295909350565b5f805f805f805f8060c0898b031215611e09575f80fd5b8835611e148161190f565b97506020890135611e248161190f565b96506040890135611e348161190f565b95506060890135611e448161190f565b9450608089013567ffffffffffffffff80821115611e60575f80fd5b611e6c8c838d01611a49565b909650945060a08b0135915080821115611e84575f80fd5b50611e918b828c01611a49565b999c989b5096995094979396929594505050565b5f805f805f60a08688031215611eb9575f80fd5b8535611ec48161190f565b94506020860135611ed48161190f565b93506040860135925060608601359150608086013567ffffffffffffffff811115611efd575f80fd5b611de588828901611cc1565b5f60208284031215611f19575f80fd5b81516119088161190f565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611f54575f80fd5b8260051b80836020870137939093016020019392505050565b6001600160a01b0386168152608060208201525f611f8f608083018688611f24565b8281036040840152611fa2818587611f24565b83810360609094019390935250505f815260200195945050505050565b6001600160a01b0386168152606060208201525f611fe1606083018688611f24565b8281036040840152611ff4818587611f24565b98975050505050505050565b6001600160a01b03851681525f60208560208401526080604084015284518060808501525f5b818110156120425786810183015185820160a001528201612026565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6001600160a01b03808b168352808a166020840152808916604084015280881660608401525060c060808301526120f260c083018688611f24565b82810360a0840152612105818587611f24565b9b9a505050505050505050505056fea164736f6c6343000818000a",
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
