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

// L1USDCGatewayMetaData contains all meta data concerning the L1USDCGateway contract.
var L1USDCGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1USDC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2USDC\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"burnAllLockedUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circleCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"pauseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"pauseWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBridgedUSDC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateCircleCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b50604051620027433803806200274383398101604081905262000033916200012f565b6200003d62000055565b6001600160a01b039182166080521660a05262000165565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161462000111575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200012a575f80fd5b919050565b5f806040838503121562000141575f80fd5b6200014c8362000113565b91506200015c6020840162000113565b90509250929050565b60805160a05161258b620001b85f395f8181610259015281816103e001528181610f1f015281816110c0015261144001525f8181610382015281816107fd01528181610daf01526113bf015261258b5ff3fe608060405260043610610178575f3560e01c8063797594b0116100d1578063c676ad291161007c578063f219fa6611610057578063f219fa6614610440578063f2fde38b14610453578063f887ea4014610472575f80fd5b8063c676ad29146103c3578063ebd462cb14610402578063f0d7c29c14610421575f80fd5b8063a2604596116100ac578063a26045961461034e578063a6f7366914610371578063c0c53b8b146103a4575f80fd5b8063797594b0146102ff57806384bd13b01461031e5780638da5cb5b14610331575f80fd5b806321846ebb116101315780633cb747bf1161010c5780633cb747bf146102ad578063415855d6146102cc578063715018a6146102eb575f80fd5b806321846ebb1461023457806329e96f9e146102485780632f3ffb9f1461027b575f80fd5b806314298c511161016157806314298c51146101d75780631f878ae6146101ea57806321425ee014610221575f80fd5b806302befd241461017c5780630aea8c26146101c2575b5f80fd5b348015610187575f80fd5b5060fa546101ad9074010000000000000000000000000000000000000000900460ff1681565b60405190151581526020015b60405180910390f35b6101d56101d0366004611f44565b610491565b005b6101d56101e5366004611ff8565b6104a5565b3480156101f5575f80fd5b5060fa54610209906001600160a01b031681565b6040516001600160a01b0390911681526020016101b9565b6101d561022f366004612037565b61072a565b34801561023f575f80fd5b506101d5610763565b348015610253575f80fd5b506102097f000000000000000000000000000000000000000000000000000000000000000081565b348015610286575f80fd5b5060fa546101ad907501000000000000000000000000000000000000000000900460ff1681565b3480156102b8575f80fd5b50609954610209906001600160a01b031681565b3480156102d7575f80fd5b506101d56102e6366004612076565b610858565b3480156102f6575f80fd5b506101d56108aa565b34801561030a575f80fd5b50609754610209906001600160a01b031681565b6101d561032c366004612098565b6108bd565b34801561033c575f80fd5b506065546001600160a01b0316610209565b348015610359575f80fd5b5061036360fb5481565b6040519081526020016101b9565b34801561037c575f80fd5b506102097f000000000000000000000000000000000000000000000000000000000000000081565b3480156103af575f80fd5b506101d56103be36600461212a565b610aad565b3480156103ce575f80fd5b506102096103dd366004612172565b507f000000000000000000000000000000000000000000000000000000000000000090565b34801561040d575f80fd5b506101d561041c366004612076565b610c25565b34801561042c575f80fd5b506101d561043b366004612172565b610c78565b6101d561044e36600461218d565b610cba565b34801561045e575f80fd5b506101d561046d366004612172565b610cc6565b34801561047d575f80fd5b50609854610209906001600160a01b031681565b61049e8585858585610d56565b5050505050565b6099546001600160a01b03163381146105055760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610541573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056591906121d0565b6001600160a01b0316736f297c61b5c92ef107ffd30cd56affe5a273e8416001600160a01b0316146105d95760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e74657874000000000060448201526064016104fc565b6105e161113b565b7f8431f5c10000000000000000000000000000000000000000000000000000000061060f60045f85876121eb565b61061891612212565b7fffffffff0000000000000000000000000000000000000000000000000000000016146106875760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642073656c6563746f720000000000000000000000000000000060448201526064016104fc565b5f808061069785600481896121eb565b8101906106a4919061225a565b5094505093505092506106b8838383611194565b6106cc6001600160a01b03841683836111fd565b816001600160a01b0316836001600160a01b03167fdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a88360405161071191815260200190565b60405180910390a350505061072560018055565b505050565b6107258333845f5b6040519080825280601f01601f19166020018201604052801561075c576020820181803683370190505b5085610d56565b60fa546001600160a01b0316336001600160a01b0316146107c65760405162461bcd60e51b815260206004820152601260248201527f6f6e6c7920636972636c652063616c6c6572000000000000000000000000000060448201526064016104fc565b60fb80545f9091556040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906342966c68906024015f604051808303815f87803b158015610846575f80fd5b505af115801561049e573d5f803e3d5ffd5b6108606112ac565b60fa805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6108b26112ac565b6108bb5f611306565b565b6099546001600160a01b03163381146109185760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064016104fc565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610954573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061097891906121d0565b6097546001600160a01b039081169116146109d55760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016104fc565b6109dd61113b565b6109ec8888888888888861136f565b610a006001600160a01b03891686866111fd565b610a3f8584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061154a92505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a788888888604051610a9294939291906122e5565b60405180910390a4610aa360018055565b5050505050505050565b5f54610100900460ff1615808015610acb57505f54600160ff909116105b80610ae45750303b158015610ae457505f5460ff166001145b610b565760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104fc565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610bb2575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610bbd8484846115df565b8015610c1f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610c2d6112ac565b60fa80549115157501000000000000000000000000000000000000000000027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff909216919091179055565b610c806112ac565b60fa80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610c1f8484845f610732565b610cce6112ac565b6001600160a01b038116610d4a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104fc565b610d5381611306565b50565b610d5e61113b565b5f8311610dad5760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e740000000000000000000000000060448201526064016104fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b031614610e2e5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c79205553444320697320616c6c6f77656400000000000000000000000060448201526064016104fc565b60fa5474010000000000000000000000000000000000000000900460ff1615610e995760405162461bcd60e51b815260206004820152600e60248201527f6465706f7369742070617573656400000000000000000000000000000000000060448201526064016104fc565b5f610ea5868585611722565b8051919650945090915015610efc5760405162461bcd60e51b815260206004820152601360248201527f63616c6c206973206e6f7420616c6c6f7765640000000000000000000000000060448201526064016104fc565b8360fb5f828254610f0d9190612357565b90915550506040515f90610f4f9088907f00000000000000000000000000000000000000000000000000000000000000009085908a908a908a906024016123bd565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8431f5c10000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa158015611008573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061102c919061240a565b6099546097546040517f5f7b15770000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692635f7b1577923492611086929116905f9088908b908b90600401612421565b5f604051808303818588803b15801561109d575f80fd5b505af11580156110af573d5f803e3d5ffd5b5050505050826001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316896001600160a01b03167f1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad886728a8a8a876040516111279493929190612463565b60405180910390a450505061049e60018055565b60026001540361118d5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104fc565b6002600155565b34156111e25760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016104fc565b8060fb5f8282546111f3919061249b565b9091555050505050565b6040516001600160a01b0383166024820152604481018290526107259084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611978565b60018055565b6065546001600160a01b031633146108bb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104fc565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b34156113bd5760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016104fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316876001600160a01b03161461143e5760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206e6f74205553444300000000000000000000000000000060448201526064016104fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316866001600160a01b0316146114bf5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206e6f74205553444300000000000000000000000000000060448201526064016104fc565b60fa547501000000000000000000000000000000000000000000900460ff161561152b5760405162461bcd60e51b815260206004820152600f60248201527f776974686472617720706175736564000000000000000000000000000000000060448201526064016104fc565b8260fb5f82825461153c919061249b565b909155505050505050505050565b5f815111801561156357505f826001600160a01b03163b115b156115db576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f906115ad9084906004016124ae565b5f604051808303815f87803b1580156115c4575f80fd5b505af11580156115d6573d5f803e3d5ffd5b505050505b5050565b6001600160a01b0383166116355760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016104fc565b6001600160a01b03811661168b5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016104fc565b611693611a5e565b61169b611ae2565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561072557609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6098545f908190606090339081906001600160a01b03168190036117f2578580602001905181019061175491906124c0565b6040517fc52a3bbc0000000000000000000000000000000000000000000000000000000081526001600160a01b0380841660048301528b81166024830152604482018b905291985091925083169063c52a3bbc906064016020604051808303815f875af11580156117c7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117eb919061240a565b965061191c565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038a16906370a0823190602401602060405180830381865afa15801561184f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611873919061240a565b905061188a6001600160a01b038a1683308b611b66565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038b16906370a0823190602401602060405180830381865afa1580156118e7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061190b919061240a565b9050611917828261249b565b985050505b5f871161196b5760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e740000000000000000000000000060448201526064016104fc565b9795965093949350505050565b5f6119cc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611bb79092919063ffffffff16565b905080515f14806119ec5750808060200190518101906119ec9190612548565b6107255760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104fc565b5f54610100900460ff16611ada5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104fc565b6108bb611bcd565b5f54610100900460ff16611b5e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104fc565b6108bb611c49565b6040516001600160a01b0380851660248301528316604482015260648101829052610c1f9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611242565b6060611bc584845f85611cce565b949350505050565b5f54610100900460ff166112a65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104fc565b5f54610100900460ff16611cc55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104fc565b6108bb33611306565b606082471015611d465760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016104fc565b5f80866001600160a01b03168587604051611d619190612563565b5f6040518083038185875af1925050503d805f8114611d9b576040519150601f19603f3d011682016040523d82523d5f602084013e611da0565b606091505b5091509150611db187838387611dbc565b979650505050505050565b60608315611e2a5782515f03611e23576001600160a01b0385163b611e235760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104fc565b5081611bc5565b611bc58383815115611e3f5781518083602001fd5b8060405162461bcd60e51b81526004016104fc91906124ae565b6001600160a01b0381168114610d53575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611ec357611ec3611e6d565b604052919050565b5f67ffffffffffffffff821115611ee457611ee4611e6d565b50601f01601f191660200190565b5f82601f830112611f01575f80fd5b8135611f14611f0f82611ecb565b611e9a565b818152846020838601011115611f28575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611f58575f80fd5b8535611f6381611e59565b94506020860135611f7381611e59565b935060408601359250606086013567ffffffffffffffff811115611f95575f80fd5b611fa188828901611ef2565b95989497509295608001359392505050565b5f8083601f840112611fc3575f80fd5b50813567ffffffffffffffff811115611fda575f80fd5b602083019150836020828501011115611ff1575f80fd5b9250929050565b5f8060208385031215612009575f80fd5b823567ffffffffffffffff81111561201f575f80fd5b61202b85828601611fb3565b90969095509350505050565b5f805f60608486031215612049575f80fd5b833561205481611e59565b95602085013595506040909401359392505050565b8015158114610d53575f80fd5b5f60208284031215612086575f80fd5b813561209181612069565b9392505050565b5f805f805f805f60c0888a0312156120ae575f80fd5b87356120b981611e59565b965060208801356120c981611e59565b955060408801356120d981611e59565b945060608801356120e981611e59565b93506080880135925060a088013567ffffffffffffffff81111561210b575f80fd5b6121178a828b01611fb3565b989b979a50959850939692959293505050565b5f805f6060848603121561213c575f80fd5b833561214781611e59565b9250602084013561215781611e59565b9150604084013561216781611e59565b809150509250925092565b5f60208284031215612182575f80fd5b813561209181611e59565b5f805f80608085870312156121a0575f80fd5b84356121ab81611e59565b935060208501356121bb81611e59565b93969395505050506040820135916060013590565b5f602082840312156121e0575f80fd5b815161209181611e59565b5f80858511156121f9575f80fd5b83861115612205575f80fd5b5050820193919092039150565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156122525780818660040360031b1b83161692505b505092915050565b5f805f805f8060c0878903121561226f575f80fd5b863561227a81611e59565b9550602087013561228a81611e59565b9450604087013561229a81611e59565b935060608701356122aa81611e59565b92506080870135915060a087013567ffffffffffffffff8111156122cc575f80fd5b6122d889828a01611ef2565b9150509295509295509295565b6001600160a01b038516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f909201601f191601019392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561236a5761236a61232a565b92915050565b5f5b8381101561238a578181015183820152602001612372565b50505f910152565b5f81518084526123a9816020860160208601612370565b601f01601f19169290920160200192915050565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526123fe60c0830184612392565b98975050505050505050565b5f6020828403121561241a575f80fd5b5051919050565b5f6001600160a01b03808816835286602084015260a0604084015261244960a0840187612392565b606084019590955292909216608090910152509392505050565b6001600160a01b0385168152836020820152608060408201525f61248a6080830185612392565b905082606083015295945050505050565b8181038181111561236a5761236a61232a565b602081525f6120916020830184612392565b5f80604083850312156124d1575f80fd5b82516124dc81611e59565b602084015190925067ffffffffffffffff8111156124f8575f80fd5b8301601f81018513612508575f80fd5b8051612516611f0f82611ecb565b81815286602083850101111561252a575f80fd5b61253b826020830160208601612370565b8093505050509250929050565b5f60208284031215612558575f80fd5b815161209181612069565b5f8251612574818460208701612370565b919091019291505056fea164736f6c6343000818000a",
}

// L1USDCGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1USDCGatewayMetaData.ABI instead.
var L1USDCGatewayABI = L1USDCGatewayMetaData.ABI

// L1USDCGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1USDCGatewayMetaData.Bin instead.
var L1USDCGatewayBin = L1USDCGatewayMetaData.Bin

// DeployL1USDCGateway deploys a new Ethereum contract, binding an instance of L1USDCGateway to it.
func DeployL1USDCGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _l1USDC common.Address, _l2USDC common.Address) (common.Address, *types.Transaction, *L1USDCGateway, error) {
	parsed, err := L1USDCGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1USDCGatewayBin), backend, _l1USDC, _l2USDC)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1USDCGateway{L1USDCGatewayCaller: L1USDCGatewayCaller{contract: contract}, L1USDCGatewayTransactor: L1USDCGatewayTransactor{contract: contract}, L1USDCGatewayFilterer: L1USDCGatewayFilterer{contract: contract}}, nil
}

// L1USDCGateway is an auto generated Go binding around an Ethereum contract.
type L1USDCGateway struct {
	L1USDCGatewayCaller     // Read-only binding to the contract
	L1USDCGatewayTransactor // Write-only binding to the contract
	L1USDCGatewayFilterer   // Log filterer for contract events
}

// L1USDCGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1USDCGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1USDCGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1USDCGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1USDCGatewaySession struct {
	Contract     *L1USDCGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1USDCGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1USDCGatewayCallerSession struct {
	Contract *L1USDCGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1USDCGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1USDCGatewayTransactorSession struct {
	Contract     *L1USDCGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1USDCGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1USDCGatewayRaw struct {
	Contract *L1USDCGateway // Generic contract binding to access the raw methods on
}

// L1USDCGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1USDCGatewayCallerRaw struct {
	Contract *L1USDCGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1USDCGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1USDCGatewayTransactorRaw struct {
	Contract *L1USDCGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1USDCGateway creates a new instance of L1USDCGateway, bound to a specific deployed contract.
func NewL1USDCGateway(address common.Address, backend bind.ContractBackend) (*L1USDCGateway, error) {
	contract, err := bindL1USDCGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1USDCGateway{L1USDCGatewayCaller: L1USDCGatewayCaller{contract: contract}, L1USDCGatewayTransactor: L1USDCGatewayTransactor{contract: contract}, L1USDCGatewayFilterer: L1USDCGatewayFilterer{contract: contract}}, nil
}

// NewL1USDCGatewayCaller creates a new read-only instance of L1USDCGateway, bound to a specific deployed contract.
func NewL1USDCGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1USDCGatewayCaller, error) {
	contract, err := bindL1USDCGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayCaller{contract: contract}, nil
}

// NewL1USDCGatewayTransactor creates a new write-only instance of L1USDCGateway, bound to a specific deployed contract.
func NewL1USDCGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1USDCGatewayTransactor, error) {
	contract, err := bindL1USDCGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayTransactor{contract: contract}, nil
}

// NewL1USDCGatewayFilterer creates a new log filterer instance of L1USDCGateway, bound to a specific deployed contract.
func NewL1USDCGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1USDCGatewayFilterer, error) {
	contract, err := bindL1USDCGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayFilterer{contract: contract}, nil
}

// bindL1USDCGateway binds a generic wrapper to an already deployed contract.
func bindL1USDCGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1USDCGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1USDCGateway *L1USDCGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1USDCGateway.Contract.L1USDCGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1USDCGateway *L1USDCGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.L1USDCGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1USDCGateway *L1USDCGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.L1USDCGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1USDCGateway *L1USDCGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1USDCGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1USDCGateway *L1USDCGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1USDCGateway *L1USDCGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.contract.Transact(opts, method, params...)
}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) CircleCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "circleCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) CircleCaller() (common.Address, error) {
	return _L1USDCGateway.Contract.CircleCaller(&_L1USDCGateway.CallOpts)
}

// CircleCaller is a free data retrieval call binding the contract method 0x1f878ae6.
//
// Solidity: function circleCaller() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) CircleCaller() (common.Address, error) {
	return _L1USDCGateway.Contract.CircleCaller(&_L1USDCGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) Counterpart() (common.Address, error) {
	return _L1USDCGateway.Contract.Counterpart(&_L1USDCGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1USDCGateway.Contract.Counterpart(&_L1USDCGateway.CallOpts)
}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewayCaller) DepositPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "depositPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewaySession) DepositPaused() (bool, error) {
	return _L1USDCGateway.Contract.DepositPaused(&_L1USDCGateway.CallOpts)
}

// DepositPaused is a free data retrieval call binding the contract method 0x02befd24.
//
// Solidity: function depositPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewayCallerSession) DepositPaused() (bool, error) {
	return _L1USDCGateway.Contract.DepositPaused(&_L1USDCGateway.CallOpts)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L1USDCGateway.Contract.GetL2ERC20Address(&_L1USDCGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L1USDCGateway.Contract.GetL2ERC20Address(&_L1USDCGateway.CallOpts, arg0)
}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) L1USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "l1USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) L1USDC() (common.Address, error) {
	return _L1USDCGateway.Contract.L1USDC(&_L1USDCGateway.CallOpts)
}

// L1USDC is a free data retrieval call binding the contract method 0xa6f73669.
//
// Solidity: function l1USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) L1USDC() (common.Address, error) {
	return _L1USDCGateway.Contract.L1USDC(&_L1USDCGateway.CallOpts)
}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) L2USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "l2USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) L2USDC() (common.Address, error) {
	return _L1USDCGateway.Contract.L2USDC(&_L1USDCGateway.CallOpts)
}

// L2USDC is a free data retrieval call binding the contract method 0x29e96f9e.
//
// Solidity: function l2USDC() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) L2USDC() (common.Address, error) {
	return _L1USDCGateway.Contract.L2USDC(&_L1USDCGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) Messenger() (common.Address, error) {
	return _L1USDCGateway.Contract.Messenger(&_L1USDCGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) Messenger() (common.Address, error) {
	return _L1USDCGateway.Contract.Messenger(&_L1USDCGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) Owner() (common.Address, error) {
	return _L1USDCGateway.Contract.Owner(&_L1USDCGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) Owner() (common.Address, error) {
	return _L1USDCGateway.Contract.Owner(&_L1USDCGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1USDCGateway *L1USDCGatewaySession) Router() (common.Address, error) {
	return _L1USDCGateway.Contract.Router(&_L1USDCGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1USDCGateway *L1USDCGatewayCallerSession) Router() (common.Address, error) {
	return _L1USDCGateway.Contract.Router(&_L1USDCGateway.CallOpts)
}

// TotalBridgedUSDC is a free data retrieval call binding the contract method 0xa2604596.
//
// Solidity: function totalBridgedUSDC() view returns(uint256)
func (_L1USDCGateway *L1USDCGatewayCaller) TotalBridgedUSDC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "totalBridgedUSDC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgedUSDC is a free data retrieval call binding the contract method 0xa2604596.
//
// Solidity: function totalBridgedUSDC() view returns(uint256)
func (_L1USDCGateway *L1USDCGatewaySession) TotalBridgedUSDC() (*big.Int, error) {
	return _L1USDCGateway.Contract.TotalBridgedUSDC(&_L1USDCGateway.CallOpts)
}

// TotalBridgedUSDC is a free data retrieval call binding the contract method 0xa2604596.
//
// Solidity: function totalBridgedUSDC() view returns(uint256)
func (_L1USDCGateway *L1USDCGatewayCallerSession) TotalBridgedUSDC() (*big.Int, error) {
	return _L1USDCGateway.Contract.TotalBridgedUSDC(&_L1USDCGateway.CallOpts)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewayCaller) WithdrawPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1USDCGateway.contract.Call(opts, &out, "withdrawPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewaySession) WithdrawPaused() (bool, error) {
	return _L1USDCGateway.Contract.WithdrawPaused(&_L1USDCGateway.CallOpts)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_L1USDCGateway *L1USDCGatewayCallerSession) WithdrawPaused() (bool, error) {
	return _L1USDCGateway.Contract.WithdrawPaused(&_L1USDCGateway.CallOpts)
}

// BurnAllLockedUSDC is a paid mutator transaction binding the contract method 0x21846ebb.
//
// Solidity: function burnAllLockedUSDC() returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) BurnAllLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "burnAllLockedUSDC")
}

// BurnAllLockedUSDC is a paid mutator transaction binding the contract method 0x21846ebb.
//
// Solidity: function burnAllLockedUSDC() returns()
func (_L1USDCGateway *L1USDCGatewaySession) BurnAllLockedUSDC() (*types.Transaction, error) {
	return _L1USDCGateway.Contract.BurnAllLockedUSDC(&_L1USDCGateway.TransactOpts)
}

// BurnAllLockedUSDC is a paid mutator transaction binding the contract method 0x21846ebb.
//
// Solidity: function burnAllLockedUSDC() returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) BurnAllLockedUSDC() (*types.Transaction, error) {
	return _L1USDCGateway.Contract.BurnAllLockedUSDC(&_L1USDCGateway.TransactOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewaySession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC20(&_L1USDCGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC20(&_L1USDCGateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewaySession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC200(&_L1USDCGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC200(&_L1USDCGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewaySession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC20AndCall(&_L1USDCGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.DepositERC20AndCall(&_L1USDCGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1USDCGateway *L1USDCGatewaySession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.FinalizeWithdrawERC20(&_L1USDCGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.FinalizeWithdrawERC20(&_L1USDCGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1USDCGateway *L1USDCGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.Initialize(&_L1USDCGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.Initialize(&_L1USDCGateway.TransactOpts, _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1USDCGateway *L1USDCGatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.OnDropMessage(&_L1USDCGateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.OnDropMessage(&_L1USDCGateway.TransactOpts, _message)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) PauseDeposit(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "pauseDeposit", _paused)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewaySession) PauseDeposit(_paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.PauseDeposit(&_L1USDCGateway.TransactOpts, _paused)
}

// PauseDeposit is a paid mutator transaction binding the contract method 0x415855d6.
//
// Solidity: function pauseDeposit(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) PauseDeposit(_paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.PauseDeposit(&_L1USDCGateway.TransactOpts, _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) PauseWithdraw(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "pauseWithdraw", _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewaySession) PauseWithdraw(_paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.PauseWithdraw(&_L1USDCGateway.TransactOpts, _paused)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0xebd462cb.
//
// Solidity: function pauseWithdraw(bool _paused) returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) PauseWithdraw(_paused bool) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.PauseWithdraw(&_L1USDCGateway.TransactOpts, _paused)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCGateway *L1USDCGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1USDCGateway.Contract.RenounceOwnership(&_L1USDCGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1USDCGateway.Contract.RenounceOwnership(&_L1USDCGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCGateway *L1USDCGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.TransferOwnership(&_L1USDCGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.TransferOwnership(&_L1USDCGateway.TransactOpts, newOwner)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L1USDCGateway *L1USDCGatewayTransactor) UpdateCircleCaller(opts *bind.TransactOpts, _caller common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.contract.Transact(opts, "updateCircleCaller", _caller)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L1USDCGateway *L1USDCGatewaySession) UpdateCircleCaller(_caller common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.UpdateCircleCaller(&_L1USDCGateway.TransactOpts, _caller)
}

// UpdateCircleCaller is a paid mutator transaction binding the contract method 0xf0d7c29c.
//
// Solidity: function updateCircleCaller(address _caller) returns()
func (_L1USDCGateway *L1USDCGatewayTransactorSession) UpdateCircleCaller(_caller common.Address) (*types.Transaction, error) {
	return _L1USDCGateway.Contract.UpdateCircleCaller(&_L1USDCGateway.TransactOpts, _caller)
}

// L1USDCGatewayDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1USDCGateway contract.
type L1USDCGatewayDepositERC20Iterator struct {
	Event *L1USDCGatewayDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1USDCGatewayDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCGatewayDepositERC20)
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
		it.Event = new(L1USDCGatewayDepositERC20)
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
func (it *L1USDCGatewayDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCGatewayDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCGatewayDepositERC20 represents a DepositERC20 event raised by the L1USDCGateway contract.
type L1USDCGatewayDepositERC20 struct {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1USDCGatewayDepositERC20Iterator, error) {

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

	logs, sub, err := _L1USDCGateway.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayDepositERC20Iterator{contract: _L1USDCGateway.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad88672.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L1USDCGateway *L1USDCGatewayFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1USDCGatewayDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1USDCGateway.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCGatewayDepositERC20)
				if err := _L1USDCGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) ParseDepositERC20(log types.Log) (*L1USDCGatewayDepositERC20, error) {
	event := new(L1USDCGatewayDepositERC20)
	if err := _L1USDCGateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCGatewayFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1USDCGateway contract.
type L1USDCGatewayFinalizeWithdrawERC20Iterator struct {
	Event *L1USDCGatewayFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1USDCGatewayFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCGatewayFinalizeWithdrawERC20)
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
		it.Event = new(L1USDCGatewayFinalizeWithdrawERC20)
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
func (it *L1USDCGatewayFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCGatewayFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCGatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1USDCGateway contract.
type L1USDCGatewayFinalizeWithdrawERC20 struct {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1USDCGatewayFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1USDCGateway.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayFinalizeWithdrawERC20Iterator{contract: _L1USDCGateway.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1USDCGateway *L1USDCGatewayFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1USDCGatewayFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1USDCGateway.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCGatewayFinalizeWithdrawERC20)
				if err := _L1USDCGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1USDCGatewayFinalizeWithdrawERC20, error) {
	event := new(L1USDCGatewayFinalizeWithdrawERC20)
	if err := _L1USDCGateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1USDCGateway contract.
type L1USDCGatewayInitializedIterator struct {
	Event *L1USDCGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1USDCGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCGatewayInitialized)
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
		it.Event = new(L1USDCGatewayInitialized)
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
func (it *L1USDCGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCGatewayInitialized represents a Initialized event raised by the L1USDCGateway contract.
type L1USDCGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1USDCGateway *L1USDCGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1USDCGatewayInitializedIterator, error) {

	logs, sub, err := _L1USDCGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayInitializedIterator{contract: _L1USDCGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1USDCGateway *L1USDCGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1USDCGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1USDCGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCGatewayInitialized)
				if err := _L1USDCGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) ParseInitialized(log types.Log) (*L1USDCGatewayInitialized, error) {
	event := new(L1USDCGatewayInitialized)
	if err := _L1USDCGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1USDCGateway contract.
type L1USDCGatewayOwnershipTransferredIterator struct {
	Event *L1USDCGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1USDCGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCGatewayOwnershipTransferred)
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
		it.Event = new(L1USDCGatewayOwnershipTransferred)
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
func (it *L1USDCGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1USDCGateway contract.
type L1USDCGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1USDCGateway *L1USDCGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1USDCGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayOwnershipTransferredIterator{contract: _L1USDCGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1USDCGateway *L1USDCGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1USDCGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCGatewayOwnershipTransferred)
				if err := _L1USDCGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1USDCGatewayOwnershipTransferred, error) {
	event := new(L1USDCGatewayOwnershipTransferred)
	if err := _L1USDCGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCGatewayRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the L1USDCGateway contract.
type L1USDCGatewayRefundERC20Iterator struct {
	Event *L1USDCGatewayRefundERC20 // Event containing the contract specifics and raw log

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
func (it *L1USDCGatewayRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCGatewayRefundERC20)
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
		it.Event = new(L1USDCGatewayRefundERC20)
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
func (it *L1USDCGatewayRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCGatewayRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCGatewayRefundERC20 represents a RefundERC20 event raised by the L1USDCGateway contract.
type L1USDCGatewayRefundERC20 struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1USDCGateway *L1USDCGatewayFilterer) FilterRefundERC20(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1USDCGatewayRefundERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1USDCGateway.contract.FilterLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCGatewayRefundERC20Iterator{contract: _L1USDCGateway.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1USDCGateway *L1USDCGatewayFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *L1USDCGatewayRefundERC20, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1USDCGateway.contract.WatchLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCGatewayRefundERC20)
				if err := _L1USDCGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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
func (_L1USDCGateway *L1USDCGatewayFilterer) ParseRefundERC20(log types.Log) (*L1USDCGatewayRefundERC20, error) {
	event := new(L1USDCGatewayRefundERC20)
	if err := _L1USDCGateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
