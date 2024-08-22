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

// L2StandardERC20GatewayMetaData contains all meta data concerning the L2StandardERC20Gateway contract.
var L2StandardERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611b51806100e65f395ff3fe6080604052600436106100d9575f3560e01c80638da5cb5b1161007c578063e77772fe11610057578063e77772fe1461020c578063f2fde38b1461022b578063f887ea401461024a578063f8c8765e14610269575f80fd5b80638da5cb5b146101bd578063a93a4af9146101da578063c676ad29146101ed575f80fd5b80636c07ea43116100b75780636c07ea4314610164578063715018a614610177578063797594b01461018b5780638431f5c1146101aa575f80fd5b80633cb747bf146100dd57806354bbd59c14610118578063575361b61461014f575b5f80fd5b3480156100e8575f80fd5b506099546100fc906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b348015610123575f80fd5b506100fc6101323660046114ed565b6001600160a01b039081165f90815260fa60205260409020541690565b61016261015d36600461150f565b610288565b005b6101626101723660046115b1565b6102d3565b348015610182575f80fd5b50610162610311565b348015610196575f80fd5b506097546100fc906001600160a01b031681565b6101626101b83660046116a4565b610324565b3480156101c8575f80fd5b506065546001600160a01b03166100fc565b6101626101e836600461176f565b6107d5565b3480156101f8575f80fd5b506100fc6102073660046114ed565b6107e7565b348015610217575f80fd5b5060fb546100fc906001600160a01b031681565b348015610236575f80fd5b506101626102453660046114ed565b610878565b348015610255575f80fd5b506098546100fc906001600160a01b031681565b348015610274575f80fd5b506101626102833660046117b2565b610908565b6102cb86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610b60915050565b505050505050565b61030c8333845f5b6040519080825280601f01601f191660200182016040528015610305576020820181803683370190505b5085610b60565b505050565b610319610ec7565b6103225f610f21565b565b6099546001600160a01b03163381146103845760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e4919061180b565b6097546001600160a01b039081169116146104415760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161037b565b610449610f8a565b34156104975760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c7565000000000000000000000000000000604482015260640161037b565b6001600160a01b0387166104ed5760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161037b565b60fb546040517f61e98ca10000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0389811660248301525f9216906361e98ca190604401602060405180830381865afa158015610554573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610578919061180b565b9050806001600160a01b0316876001600160a01b0316146105db5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161037b565b505f828060200190518101906105f19190611892565b93509050606080821561061b578480602001905181019061061291906118e4565b92509050610689565b6001600160a01b038981165f90815260fa60205260409020548116908b16146106865760405162461bcd60e51b815260206004820152601660248201527f746f6b656e206d617070696e67206d69736d6174636800000000000000000000604482015260640161037b565b50835b6001600160a01b0389163b6106e7576001600160a01b038981165f90815260fa6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016918c169190911790556106e7828b610fe3565b6040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038881166004830152602482018890528a16906340c10f19906044015f604051808303815f87803b158015610747575f80fd5b505af1158015610759573d5f803e3d5ffd5b5050505061076787826110f7565b876001600160a01b0316896001600160a01b03168b6001600160a01b03167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348a8a866040516107b893929190611983565b60405180910390a45050506107cc60018055565b50505050505050565b6107e18484845f6102db565b50505050565b60fb546040517f61e98ca10000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0383811660248301525f9216906361e98ca190604401602060405180830381865afa15801561084e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610872919061180b565b92915050565b610880610ec7565b6001600160a01b0381166108fc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161037b565b61090581610f21565b50565b5f54610100900460ff161580801561092657505f54600160ff909116105b8061093f5750303b15801561093f57505f5460ff166001145b6109b15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161037b565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a0d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b038416610a635760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f75746572206164647265737300000000000000000000000000604482015260640161037b565b610a6e85858561118d565b6001600160a01b038216610ac45760405162461bcd60e51b815260206004820152601260248201527f7a65726f20746f6b656e20666163746f72790000000000000000000000000000604482015260640161037b565b60fb80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790558015610b59575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610b68610f8a565b5f8311610bb75760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e74000000000000000000000000604482015260640161037b565b60985433906001600160a01b0316819003610be55782806020019051810190610be091906119b3565b935090505b6001600160a01b038087165f90815260fa60205260409020541680610c4c5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e00000000000000604482015260640161037b565b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081526001600160a01b03838116600483015260248201879052881690639dc29fac906044015f604051808303815f87803b158015610cac575f80fd5b505af1158015610cbe573d5f803e3d5ffd5b505050505f818884898989604051602401610cde969594939291906119cf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa158015610db5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd99190611a1c565b6099546097546040517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039182169263b2267a7b923492610e31929116905f9088908c90600401611a33565b5f604051808303818588803b158015610e48575f80fd5b505af1158015610e5a573d5f803e3d5ffd5b5050505050836001600160a01b0316896001600160a01b0316846001600160a01b03167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b87604051610eb29493929190611a33565b60405180910390a450505050610b5960018055565b6065546001600160a01b031633146103225760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161037b565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600260015403610fdc5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161037b565b6002600155565b60fb546040517f7bdbcbbf0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0383811660248301525f921690637bdbcbbf906044016020604051808303815f875af115801561104b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106f919061180b565b90505f805f858060200190518101906110889190611a6b565b925092509250836001600160a01b031663c820f146838584308a6040518663ffffffff1660e01b81526004016110c2959493929190611ae3565b5f604051808303815f87803b1580156110d9575f80fd5b505af11580156110eb573d5f803e3d5ffd5b50505050505050505050565b5f815111801561111057505f826001600160a01b03163b115b15611183576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f9061115a908490600401611b32565b5f604051808303815f87803b158015611171575f80fd5b505af11580156102cb573d5f803e3d5ffd5b5050565b60018055565b6001600160a01b0383166111e35760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161037b565b6001600160a01b0381166112395760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161037b565b6112416112d0565b611249611354565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561030c57609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff1661134c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161037b565b6103226113d8565b5f54610100900460ff166113d05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161037b565b610322611454565b5f54610100900460ff166111875760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161037b565b5f54610100900460ff166114d05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161037b565b61032233610f21565b6001600160a01b0381168114610905575f80fd5b5f602082840312156114fd575f80fd5b8135611508816114d9565b9392505050565b5f805f805f8060a08789031215611524575f80fd5b863561152f816114d9565b9550602087013561153f816114d9565b945060408701359350606087013567ffffffffffffffff80821115611562575f80fd5b818901915089601f830112611575575f80fd5b813581811115611583575f80fd5b8a6020828501011115611594575f80fd5b602083019550809450505050608087013590509295509295509295565b5f805f606084860312156115c3575f80fd5b83356115ce816114d9565b95602085013595506040909401359392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611657576116576115e3565b604052919050565b5f67ffffffffffffffff821115611678576116786115e3565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f805f805f8060c087890312156116b9575f80fd5b86356116c4816114d9565b955060208701356116d4816114d9565b945060408701356116e4816114d9565b935060608701356116f4816114d9565b92506080870135915060a087013567ffffffffffffffff811115611716575f80fd5b8701601f81018913611726575f80fd5b80356117396117348261165f565b611610565b8181528a602083850101111561174d575f80fd5b816020840160208301375f602083830101528093505050509295509295509295565b5f805f8060808587031215611782575f80fd5b843561178d816114d9565b9350602085013561179d816114d9565b93969395505050506040820135916060013590565b5f805f80608085870312156117c5575f80fd5b84356117d0816114d9565b935060208501356117e0816114d9565b925060408501356117f0816114d9565b91506060850135611800816114d9565b939692955090935050565b5f6020828403121561181b575f80fd5b8151611508816114d9565b5f5b83811015611840578181015183820152602001611828565b50505f910152565b5f82601f830112611857575f80fd5b81516118656117348261165f565b818152846020838601011115611879575f80fd5b61188a826020830160208701611826565b949350505050565b5f80604083850312156118a3575f80fd5b825180151581146118b2575f80fd5b602084015190925067ffffffffffffffff8111156118ce575f80fd5b6118da85828601611848565b9150509250929050565b5f80604083850312156118f5575f80fd5b825167ffffffffffffffff8082111561190c575f80fd5b61191886838701611848565b9350602085015191508082111561192d575f80fd5b506118da85828601611848565b5f8151808452611951816020860160208601611826565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6001600160a01b0384168152826020820152606060408201525f6119aa606083018461193a565b95945050505050565b5f80604083850312156119c4575f80fd5b82516118b2816114d9565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611a1060c083018461193a565b98975050505050505050565b5f60208284031215611a2c575f80fd5b5051919050565b6001600160a01b0385168152836020820152608060408201525f611a5a608083018561193a565b905082606083015295945050505050565b5f805f60608486031215611a7d575f80fd5b835167ffffffffffffffff80821115611a94575f80fd5b611aa087838801611848565b94506020860151915080821115611ab5575f80fd5b50611ac286828701611848565b925050604084015160ff81168114611ad8575f80fd5b809150509250925092565b60a081525f611af560a083018861193a565b8281036020840152611b07818861193a565b60ff96909616604084015250506001600160a01b039283166060820152911660809091015292915050565b602081525f611508602083018461193a56fea164736f6c6343000818000a",
}

// L2StandardERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2StandardERC20GatewayMetaData.ABI instead.
var L2StandardERC20GatewayABI = L2StandardERC20GatewayMetaData.ABI

// L2StandardERC20GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2StandardERC20GatewayMetaData.Bin instead.
var L2StandardERC20GatewayBin = L2StandardERC20GatewayMetaData.Bin

// DeployL2StandardERC20Gateway deploys a new Ethereum contract, binding an instance of L2StandardERC20Gateway to it.
func DeployL2StandardERC20Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2StandardERC20Gateway, error) {
	parsed, err := L2StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2StandardERC20GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2StandardERC20Gateway{L2StandardERC20GatewayCaller: L2StandardERC20GatewayCaller{contract: contract}, L2StandardERC20GatewayTransactor: L2StandardERC20GatewayTransactor{contract: contract}, L2StandardERC20GatewayFilterer: L2StandardERC20GatewayFilterer{contract: contract}}, nil
}

// L2StandardERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L2StandardERC20Gateway struct {
	L2StandardERC20GatewayCaller     // Read-only binding to the contract
	L2StandardERC20GatewayTransactor // Write-only binding to the contract
	L2StandardERC20GatewayFilterer   // Log filterer for contract events
}

// L2StandardERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2StandardERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2StandardERC20GatewaySession struct {
	Contract     *L2StandardERC20Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2StandardERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2StandardERC20GatewayCallerSession struct {
	Contract *L2StandardERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2StandardERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2StandardERC20GatewayTransactorSession struct {
	Contract     *L2StandardERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2StandardERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2StandardERC20GatewayRaw struct {
	Contract *L2StandardERC20Gateway // Generic contract binding to access the raw methods on
}

// L2StandardERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayCallerRaw struct {
	Contract *L2StandardERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2StandardERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayTransactorRaw struct {
	Contract *L2StandardERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2StandardERC20Gateway creates a new instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20Gateway(address common.Address, backend bind.ContractBackend) (*L2StandardERC20Gateway, error) {
	contract, err := bindL2StandardERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20Gateway{L2StandardERC20GatewayCaller: L2StandardERC20GatewayCaller{contract: contract}, L2StandardERC20GatewayTransactor: L2StandardERC20GatewayTransactor{contract: contract}, L2StandardERC20GatewayFilterer: L2StandardERC20GatewayFilterer{contract: contract}}, nil
}

// NewL2StandardERC20GatewayCaller creates a new read-only instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2StandardERC20GatewayCaller, error) {
	contract, err := bindL2StandardERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayCaller{contract: contract}, nil
}

// NewL2StandardERC20GatewayTransactor creates a new write-only instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2StandardERC20GatewayTransactor, error) {
	contract, err := bindL2StandardERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayTransactor{contract: contract}, nil
}

// NewL2StandardERC20GatewayFilterer creates a new log filterer instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2StandardERC20GatewayFilterer, error) {
	contract, err := bindL2StandardERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayFilterer{contract: contract}, nil
}

// bindL2StandardERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL2StandardERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Counterpart() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Counterpart(&_L2StandardERC20Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Counterpart(&_L2StandardERC20Gateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL1ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL1ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL2ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL2ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l1Token)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Messenger() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Messenger(&_L2StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Messenger() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Messenger(&_L2StandardERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Owner() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Owner(&_L2StandardERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Owner() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Owner(&_L2StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Router() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Router(&_L2StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Router(&_L2StandardERC20Gateway.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) TokenFactory() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.TokenFactory(&_L2StandardERC20Gateway.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) TokenFactory() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.TokenFactory(&_L2StandardERC20Gateway.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.FinalizeDepositERC20(&_L2StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.FinalizeDepositERC20(&_L2StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.Initialize(&_L2StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.Initialize(&_L2StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _tokenFactory)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.RenounceOwnership(&_L2StandardERC20Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.RenounceOwnership(&_L2StandardERC20Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.TransferOwnership(&_L2StandardERC20Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.TransferOwnership(&_L2StandardERC20Gateway.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20(&_L2StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20(&_L2StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC200(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC200(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20AndCall(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20AndCall(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2StandardERC20GatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayFinalizeDepositERC20Iterator struct {
	Event *L2StandardERC20GatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayFinalizeDepositERC20)
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
		it.Event = new(L2StandardERC20GatewayFinalizeDepositERC20)
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
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayFinalizeDepositERC20 struct {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2StandardERC20GatewayFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayFinalizeDepositERC20Iterator{contract: _L2StandardERC20Gateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayFinalizeDepositERC20)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2StandardERC20GatewayFinalizeDepositERC20, error) {
	event := new(L2StandardERC20GatewayFinalizeDepositERC20)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StandardERC20GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayInitializedIterator struct {
	Event *L2StandardERC20GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayInitialized)
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
		it.Event = new(L2StandardERC20GatewayInitialized)
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
func (it *L2StandardERC20GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayInitialized represents a Initialized event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2StandardERC20GatewayInitializedIterator, error) {

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayInitializedIterator{contract: _L2StandardERC20Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayInitialized)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseInitialized(log types.Log) (*L2StandardERC20GatewayInitialized, error) {
	event := new(L2StandardERC20GatewayInitialized)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StandardERC20GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayOwnershipTransferredIterator struct {
	Event *L2StandardERC20GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayOwnershipTransferred)
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
		it.Event = new(L2StandardERC20GatewayOwnershipTransferred)
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
func (it *L2StandardERC20GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2StandardERC20GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayOwnershipTransferredIterator{contract: _L2StandardERC20Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayOwnershipTransferred)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2StandardERC20GatewayOwnershipTransferred, error) {
	event := new(L2StandardERC20GatewayOwnershipTransferred)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StandardERC20GatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayWithdrawERC20Iterator struct {
	Event *L2StandardERC20GatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayWithdrawERC20)
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
		it.Event = new(L2StandardERC20GatewayWithdrawERC20)
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
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayWithdrawERC20 struct {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2StandardERC20GatewayWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayWithdrawERC20Iterator{contract: _L2StandardERC20Gateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayWithdrawERC20)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2StandardERC20GatewayWithdrawERC20, error) {
	event := new(L2StandardERC20GatewayWithdrawERC20)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
