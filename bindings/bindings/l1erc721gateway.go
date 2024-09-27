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

// L1ERC721GatewayMetaData contains all meta data concerning the L1ERC721Gateway contract.
var L1ERC721GatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchDepositERC721\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"batchDepositERC721\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"counterpart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositERC721\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositERC721\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeBatchWithdrawERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeWithdrawERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_counterpart\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onDropMessage\",\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTokenMapping\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BatchDepositERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchRefundERC721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeBatchWithdrawERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeWithdrawERC721\",\"inputs\":[{\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundERC721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"inputs\":[{\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldL2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newL2Token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b62000021565b620000df565b5f54610100900460ff16156200008d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614620000dd575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61292a80620000ed5f395ff3fe608060405260043610610109575f3560e01c8063797594b0116100a1578063d606b4dc11610071578063f2fde38b11610057578063f2fde38b14610340578063f887ea401461035f578063fac752eb1461038b575f80fd5b8063d606b4dc1461030e578063d96c8ecf1461032d575f80fd5b8063797594b0146102585780638da5cb5b146102845780639f0a68b3146102ae578063ba27f50b146102cd575f80fd5b80633cb747bf116100dc5780633cb747bf146101c157806345a4276b14610212578063485cc95514610225578063715018a614610244575f80fd5b80630a7aa1961461010d57806314298c5114610122578063150b7a02146101355780631b997a93146101ae575b5f80fd5b61012061011b366004612126565b6103aa565b005b610120610130366004612169565b6103bc565b348015610140575f80fd5b5061017861014f366004612251565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b6101206101bc366004612374565b610971565b3480156101cc575f80fd5b5060cb546101ed9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a5565b6101206102203660046123dc565b610985565b348015610230575f80fd5b5061012061023f36600461240e565b610991565b34801561024f575f80fd5b50610120610b29565b348015610263575f80fd5b5060c9546101ed9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561028f575f80fd5b5060975473ffffffffffffffffffffffffffffffffffffffff166101ed565b3480156102b9575f80fd5b506101206102c8366004612445565b610b3c565b3480156102d8575f80fd5b506101ed6102e73660046124ce565b60fa6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b348015610319575f80fd5b506101206103283660046124f0565b610f2f565b61012061033b366004612550565b6112af565b34801561034b575f80fd5b5061012061035a3660046124ce565b6112bc565b34801561036a575f80fd5b5060ca546101ed9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610396575f80fd5b506101206103a536600461240e565b611373565b6103b684848484611480565b50505050565b60cb5473ffffffffffffffffffffffffffffffffffffffff16338114610443576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561048c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b091906125a8565b73ffffffffffffffffffffffffffffffffffffffff16736f297c61b5c92ef107ffd30cd56affe5a273e84173ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e746578740000000000604482015260640161043a565b61056061175b565b34156105c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c7565000000000000000000000000000000604482015260640161043a565b7ff8c3cf25000000000000000000000000000000000000000000000000000000006105f660045f85876125c3565b6105ff916125ea565b7fffffffff000000000000000000000000000000000000000000000000000000001603610746575f808061063685600481896125c3565b81019061064391906124f0565b6040517f42842e0e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff808516602483015260448201839052959850929650945050918516916342842e0e91506064015f604051808303815f87803b1580156106c1575f80fd5b505af11580156106d3573d5f803e3d5ffd5b505050508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd141355812088360405161073691815260200190565b60405180910390a3505050610962565b7f982b151f0000000000000000000000000000000000000000000000000000000061077460045f85876125c3565b61077d916125ea565b7fffffffff000000000000000000000000000000000000000000000000000000001603610900575f80806107b485600481896125c3565b8101906107c19190612632565b94505093505092505f5b81518110156108a2578373ffffffffffffffffffffffffffffffffffffffff166342842e0e30858585815181106108045761080461271e565b60209081029190910101516040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff938416600482015292909116602483015260448201526064015f604051808303815f87803b158015610880575f80fd5b505af1158015610892573d5f803e3d5ffd5b5050600190920191506107cb9050565b508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc2283604051610736919061274b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f696e76616c69642073656c6563746f7200000000000000000000000000000000604482015260640161043a565b61096c6001603355565b505050565b61097e85858585856117d5565b5050505050565b61096c83338484611480565b5f54610100900460ff16158080156109af57505f54600160ff909116105b806109c85750303b1580156109c857505f5460ff166001145b610a54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161043a565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610ab0575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610ab8611b5c565b610ac3835f84611bf2565b801561096c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b610b31611d9d565b610b3a5f611e1e565b565b60cb5473ffffffffffffffffffffffffffffffffffffffff16338114610bbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161043a565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c07573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c2b91906125a8565b60c95473ffffffffffffffffffffffffffffffffffffffff908116911614610caf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161043a565b610cb761175b565b73ffffffffffffffffffffffffffffffffffffffff8616610d34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161043a565b73ffffffffffffffffffffffffffffffffffffffff8088165f90815260fa6020526040902054878216911614610dc6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161043a565b5f5b82811015610e9b578773ffffffffffffffffffffffffffffffffffffffff166342842e0e3087878786818110610e0057610e0061271e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815273ffffffffffffffffffffffffffffffffffffffff9586166004820152949093166024850152506020909102013560448201526064015f604051808303815f87803b158015610e79575f80fd5b505af1158015610e8b573d5f803e3d5ffd5b505060019092019150610dc89050565b508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d878787604051610f14939291906127d7565b60405180910390a4610f266001603355565b50505050505050565b60cb5473ffffffffffffffffffffffffffffffffffffffff16338114610fb1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161043a565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ffa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101e91906125a8565b60c95473ffffffffffffffffffffffffffffffffffffffff9081169116146110a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161043a565b6110aa61175b565b73ffffffffffffffffffffffffffffffffffffffff8516611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161043a565b73ffffffffffffffffffffffffffffffffffffffff8087165f90815260fa60205260409020548682169116146111b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d61746368000000000000000000000000000000604482015260640161043a565b6040517f42842e0e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8481166024830152604482018490528716906342842e0e906064015f604051808303815f87803b15801561122c575f80fd5b505af115801561123e573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff878116825260208201879052808916945089811693508a16917facdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9910160405180910390a46112a76001603355565b505050505050565b6103b684338585856117d5565b6112c4611d9d565b73ffffffffffffffffffffffffffffffffffffffff8116611367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161043a565b61137081611e1e565b50565b61137b611d9d565b73ffffffffffffffffffffffffffffffffffffffff81166113f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f74206265203000000000000000604482015260640161043a565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815260fa602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b61148861175b565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260fa60205260409020541680611516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3220746f6b656e00000000000000604482015260640161043a565b5f336040517f42842e0e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808316600483015230602483015260448201879052919250908716906342842e0e906064015f604051808303815f87803b15801561158f575f80fd5b505af11580156115a1573d5f803e3d5ffd5b505060405173ffffffffffffffffffffffffffffffffffffffff808a166024830152808616604483015280851660648301528816608482015260a481018790525f925060c4019050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c3cf250000000000000000000000000000000000000000000000000000000017905260cb5460c95491517f5f7b157700000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90811692635f7b15779234926116c5929116905f9087908b908a9060040161280f565b5f604051808303818588803b1580156116dc575f80fd5b505af11580156116ee573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8b81168252602082018b9052808816955088811694508c1692507ffc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec910160405180910390a45050506103b66001603355565b6002603354036117c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161043a565b6002603355565b6001603355565b6117dd61175b565b81611844576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6e6f20746f6b656e20746f206465706f73697400000000000000000000000000604482015260640161043a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260fa602052604090205416806118d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3220746f6b656e00000000000000604482015260640161043a565b335f5b848110156119a8578773ffffffffffffffffffffffffffffffffffffffff166342842e0e833089898681811061190d5761190d61271e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815273ffffffffffffffffffffffffffffffffffffffff9586166004820152949093166024850152506020909102013560448201526064015f604051808303815f87803b158015611986575f80fd5b505af1158015611998573d5f803e3d5ffd5b5050600190920191506118d59050565b505f8783838989896040516024016119c5969594939291906128c8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f982b151f0000000000000000000000000000000000000000000000000000000017905260cb5460c95491517f5f7b157700000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90811692635f7b1577923492611aa1929116905f9087908b908a9060040161280f565b5f604051808303818588803b158015611ab8575f80fd5b505af1158015611aca573d5f803e3d5ffd5b50505050508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167ff05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd68a8a8a604051611b47939291906127d7565b60405180910390a450505061097e6001603355565b5f54610100900460ff16610b3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043a565b73ffffffffffffffffffffffffffffffffffffffff8316611c6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161043a565b73ffffffffffffffffffffffffffffffffffffffff8116611cec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161043a565b611cf4611e94565b611cfc611f32565b60c9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560cb8054848416921691909117905582161561096c5760ca805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b60975473ffffffffffffffffffffffffffffffffffffffff163314610b3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161043a565b6097805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611f2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043a565b610b3a611fd0565b5f54610100900460ff16611fc8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043a565b610b3a612066565b5f54610100900460ff166117ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043a565b5f54610100900460ff166120fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161043a565b610b3a33611e1e565b73ffffffffffffffffffffffffffffffffffffffff81168114611370575f80fd5b5f805f8060808587031215612139575f80fd5b843561214481612105565b9350602085013561215481612105565b93969395505050506040820135916060013590565b5f806020838503121561217a575f80fd5b823567ffffffffffffffff80821115612191575f80fd5b818501915085601f8301126121a4575f80fd5b8135818111156121b2575f80fd5b8660208285010111156121c3575f80fd5b60209290920196919550909350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612249576122496121d5565b604052919050565b5f805f8060808587031215612264575f80fd5b843561226f81612105565b935060208581013561228081612105565b935060408601359250606086013567ffffffffffffffff808211156122a3575f80fd5b818801915088601f8301126122b6575f80fd5b8135818111156122c8576122c86121d5565b6122f8847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612202565b9150808252898482850101111561230d575f80fd5b80848401858401375f8482840101525080935050505092959194509250565b5f8083601f84011261233c575f80fd5b50813567ffffffffffffffff811115612353575f80fd5b6020830191508360208260051b850101111561236d575f80fd5b9250929050565b5f805f805f60808688031215612388575f80fd5b853561239381612105565b945060208601356123a381612105565b9350604086013567ffffffffffffffff8111156123be575f80fd5b6123ca8882890161232c565b96999598509660600135949350505050565b5f805f606084860312156123ee575f80fd5b83356123f981612105565b95602085013595506040909401359392505050565b5f806040838503121561241f575f80fd5b823561242a81612105565b9150602083013561243a81612105565b809150509250929050565b5f805f805f8060a0878903121561245a575f80fd5b863561246581612105565b9550602087013561247581612105565b9450604087013561248581612105565b9350606087013561249581612105565b9250608087013567ffffffffffffffff8111156124b0575f80fd5b6124bc89828a0161232c565b979a9699509497509295939492505050565b5f602082840312156124de575f80fd5b81356124e981612105565b9392505050565b5f805f805f60a08688031215612504575f80fd5b853561250f81612105565b9450602086013561251f81612105565b9350604086013561252f81612105565b9250606086013561253f81612105565b949793965091946080013592915050565b5f805f8060608587031215612563575f80fd5b843561256e81612105565b9350602085013567ffffffffffffffff811115612589575f80fd5b6125958782880161232c565b9598909750949560400135949350505050565b5f602082840312156125b8575f80fd5b81516124e981612105565b5f80858511156125d1575f80fd5b838611156125dd575f80fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561262a5780818660040360031b1b83161692505b505092915050565b5f805f805f60a08688031215612646575f80fd5b853561265181612105565b945060208681013561266281612105565b9450604087013561267281612105565b9350606087013561268281612105565b9250608087013567ffffffffffffffff8082111561269e575f80fd5b818901915089601f8301126126b1575f80fd5b8135818111156126c3576126c36121d5565b8060051b91506126d4848301612202565b818152918301840191848101908c8411156126ed575f80fd5b938501935b8385101561270b578435825293850193908501906126f2565b8096505050505050509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b602080825282518282018190525f9190848201906040850190845b8181101561278257835183529284019291840191600101612766565b50909695505050505050565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156127be575f80fd5b8260051b80836020870137939093016020019392505050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201525f61280660408301848661278e565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff861681525f602086602084015260a0604084015285518060a08501525f5b8181101561285e5787810183015185820160c001528201612842565b505f60c0828601015260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050508360608301526128be608083018473ffffffffffffffffffffffffffffffffffffffff169052565b9695505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525060a0608083015261291160a08301848661278e565b9897505050505050505056fea164736f6c6343000818000a",
}

// L1ERC721GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC721GatewayMetaData.ABI instead.
var L1ERC721GatewayABI = L1ERC721GatewayMetaData.ABI

// L1ERC721GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1ERC721GatewayMetaData.Bin instead.
var L1ERC721GatewayBin = L1ERC721GatewayMetaData.Bin

// DeployL1ERC721Gateway deploys a new Ethereum contract, binding an instance of L1ERC721Gateway to it.
func DeployL1ERC721Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1ERC721Gateway, error) {
	parsed, err := L1ERC721GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1ERC721GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1ERC721Gateway{L1ERC721GatewayCaller: L1ERC721GatewayCaller{contract: contract}, L1ERC721GatewayTransactor: L1ERC721GatewayTransactor{contract: contract}, L1ERC721GatewayFilterer: L1ERC721GatewayFilterer{contract: contract}}, nil
}

// L1ERC721Gateway is an auto generated Go binding around an Ethereum contract.
type L1ERC721Gateway struct {
	L1ERC721GatewayCaller     // Read-only binding to the contract
	L1ERC721GatewayTransactor // Write-only binding to the contract
	L1ERC721GatewayFilterer   // Log filterer for contract events
}

// L1ERC721GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC721GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC721GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC721GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC721GatewaySession struct {
	Contract     *L1ERC721Gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC721GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC721GatewayCallerSession struct {
	Contract *L1ERC721GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1ERC721GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC721GatewayTransactorSession struct {
	Contract     *L1ERC721GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1ERC721GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC721GatewayRaw struct {
	Contract *L1ERC721Gateway // Generic contract binding to access the raw methods on
}

// L1ERC721GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC721GatewayCallerRaw struct {
	Contract *L1ERC721GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC721GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC721GatewayTransactorRaw struct {
	Contract *L1ERC721GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC721Gateway creates a new instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721Gateway(address common.Address, backend bind.ContractBackend) (*L1ERC721Gateway, error) {
	contract, err := bindL1ERC721Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC721Gateway{L1ERC721GatewayCaller: L1ERC721GatewayCaller{contract: contract}, L1ERC721GatewayTransactor: L1ERC721GatewayTransactor{contract: contract}, L1ERC721GatewayFilterer: L1ERC721GatewayFilterer{contract: contract}}, nil
}

// NewL1ERC721GatewayCaller creates a new read-only instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ERC721GatewayCaller, error) {
	contract, err := bindL1ERC721Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayCaller{contract: contract}, nil
}

// NewL1ERC721GatewayTransactor creates a new write-only instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC721GatewayTransactor, error) {
	contract, err := bindL1ERC721Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayTransactor{contract: contract}, nil
}

// NewL1ERC721GatewayFilterer creates a new log filterer instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC721GatewayFilterer, error) {
	contract, err := bindL1ERC721Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayFilterer{contract: contract}, nil
}

// bindL1ERC721Gateway binds a generic wrapper to an already deployed contract.
func bindL1ERC721Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1ERC721GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC721Gateway *L1ERC721GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC721Gateway.Contract.L1ERC721GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC721Gateway *L1ERC721GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.L1ERC721GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC721Gateway *L1ERC721GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.L1ERC721GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC721Gateway *L1ERC721GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC721Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC721Gateway *L1ERC721GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC721Gateway *L1ERC721GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewaySession) Counterpart() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Counterpart(&_L1ERC721Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Counterpart(&_L1ERC721Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewaySession) Messenger() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Messenger(&_L1ERC721Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCallerSession) Messenger() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Messenger(&_L1ERC721Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewaySession) Owner() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Owner(&_L1ERC721Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCallerSession) Owner() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Owner(&_L1ERC721Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewaySession) Router() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Router(&_L1ERC721Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCallerSession) Router() (common.Address, error) {
	return _L1ERC721Gateway.Contract.Router(&_L1ERC721Gateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1ERC721Gateway.Contract.TokenMapping(&_L1ERC721Gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1ERC721Gateway.Contract.TokenMapping(&_L1ERC721Gateway.CallOpts, arg0)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) BatchDepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "batchDepositERC721", _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.BatchDepositERC721(&_L1ERC721Gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.BatchDepositERC721(&_L1ERC721Gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) BatchDepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "batchDepositERC7210", _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.BatchDepositERC7210(&_L1ERC721Gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.BatchDepositERC7210(&_L1ERC721Gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) DepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "depositERC721", _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.DepositERC721(&_L1ERC721Gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.DepositERC721(&_L1ERC721Gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) DepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "depositERC7210", _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.DepositERC7210(&_L1ERC721Gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.DepositERC7210(&_L1ERC721Gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) FinalizeBatchWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "finalizeBatchWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.FinalizeBatchWithdrawERC721(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.FinalizeBatchWithdrawERC721(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) FinalizeWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "finalizeWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.FinalizeWithdrawERC721(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.FinalizeWithdrawERC721(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.Initialize(&_L1ERC721Gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.Initialize(&_L1ERC721Gateway.TransactOpts, _counterpart, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.OnDropMessage(&_L1ERC721Gateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.OnDropMessage(&_L1ERC721Gateway.TransactOpts, _message)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1ERC721Gateway *L1ERC721GatewaySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.OnERC721Received(&_L1ERC721Gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.OnERC721Received(&_L1ERC721Gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.RenounceOwnership(&_L1ERC721Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.RenounceOwnership(&_L1ERC721Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.TransferOwnership(&_L1ERC721Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.TransferOwnership(&_L1ERC721Gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ERC721Gateway *L1ERC721GatewaySession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.UpdateTokenMapping(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactorSession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.Contract.UpdateTokenMapping(&_L1ERC721Gateway.TransactOpts, _l1Token, _l2Token)
}

// L1ERC721GatewayBatchDepositERC721Iterator is returned from FilterBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for BatchDepositERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchDepositERC721Iterator struct {
	Event *L1ERC721GatewayBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayBatchDepositERC721)
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
		it.Event = new(L1ERC721GatewayBatchDepositERC721)
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
func (it *L1ERC721GatewayBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayBatchDepositERC721 represents a BatchDepositERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchDepositERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchDepositERC721 is a free log retrieval operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterBatchDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1ERC721GatewayBatchDepositERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayBatchDepositERC721Iterator{contract: _L1ERC721Gateway.contract, event: "BatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchBatchDepositERC721 is a free log subscription operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayBatchDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayBatchDepositERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
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

// ParseBatchDepositERC721 is a log parse operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseBatchDepositERC721(log types.Log) (*L1ERC721GatewayBatchDepositERC721, error) {
	event := new(L1ERC721GatewayBatchDepositERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayBatchRefundERC721Iterator is returned from FilterBatchRefundERC721 and is used to iterate over the raw logs and unpacked data for BatchRefundERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchRefundERC721Iterator struct {
	Event *L1ERC721GatewayBatchRefundERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayBatchRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayBatchRefundERC721)
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
		it.Event = new(L1ERC721GatewayBatchRefundERC721)
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
func (it *L1ERC721GatewayBatchRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayBatchRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayBatchRefundERC721 represents a BatchRefundERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchRefundERC721 is a free log retrieval operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterBatchRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1ERC721GatewayBatchRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayBatchRefundERC721Iterator{contract: _L1ERC721Gateway.contract, event: "BatchRefundERC721", logs: logs, sub: sub}, nil
}

// WatchBatchRefundERC721 is a free log subscription operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchBatchRefundERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayBatchRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayBatchRefundERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
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

// ParseBatchRefundERC721 is a log parse operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseBatchRefundERC721(log types.Log) (*L1ERC721GatewayBatchRefundERC721, error) {
	event := new(L1ERC721GatewayBatchRefundERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayDepositERC721Iterator is returned from FilterDepositERC721 and is used to iterate over the raw logs and unpacked data for DepositERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayDepositERC721Iterator struct {
	Event *L1ERC721GatewayDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayDepositERC721)
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
		it.Event = new(L1ERC721GatewayDepositERC721)
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
func (it *L1ERC721GatewayDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayDepositERC721 represents a DepositERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayDepositERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC721 is a free log retrieval operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1ERC721GatewayDepositERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayDepositERC721Iterator{contract: _L1ERC721Gateway.contract, event: "DepositERC721", logs: logs, sub: sub}, nil
}

// WatchDepositERC721 is a free log subscription operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayDepositERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
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

// ParseDepositERC721 is a log parse operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseDepositERC721(log types.Log) (*L1ERC721GatewayDepositERC721, error) {
	event := new(L1ERC721GatewayDepositERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator is returned from FilterFinalizeBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchWithdrawERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator struct {
	Event *L1ERC721GatewayFinalizeBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayFinalizeBatchWithdrawERC721)
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
		it.Event = new(L1ERC721GatewayFinalizeBatchWithdrawERC721)
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
func (it *L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayFinalizeBatchWithdrawERC721 represents a FinalizeBatchWithdrawERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeBatchWithdrawERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchWithdrawERC721 is a free log retrieval operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterFinalizeBatchWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayFinalizeBatchWithdrawERC721Iterator{contract: _L1ERC721Gateway.contract, event: "FinalizeBatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchWithdrawERC721 is a free log subscription operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchFinalizeBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayFinalizeBatchWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayFinalizeBatchWithdrawERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
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

// ParseFinalizeBatchWithdrawERC721 is a log parse operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseFinalizeBatchWithdrawERC721(log types.Log) (*L1ERC721GatewayFinalizeBatchWithdrawERC721, error) {
	event := new(L1ERC721GatewayFinalizeBatchWithdrawERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayFinalizeWithdrawERC721Iterator is returned from FilterFinalizeWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeWithdrawERC721Iterator struct {
	Event *L1ERC721GatewayFinalizeWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayFinalizeWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayFinalizeWithdrawERC721)
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
		it.Event = new(L1ERC721GatewayFinalizeWithdrawERC721)
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
func (it *L1ERC721GatewayFinalizeWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayFinalizeWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayFinalizeWithdrawERC721 represents a FinalizeWithdrawERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeWithdrawERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC721 is a free log retrieval operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterFinalizeWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1ERC721GatewayFinalizeWithdrawERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayFinalizeWithdrawERC721Iterator{contract: _L1ERC721Gateway.contract, event: "FinalizeWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC721 is a free log subscription operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchFinalizeWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayFinalizeWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayFinalizeWithdrawERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
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

// ParseFinalizeWithdrawERC721 is a log parse operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseFinalizeWithdrawERC721(log types.Log) (*L1ERC721GatewayFinalizeWithdrawERC721, error) {
	event := new(L1ERC721GatewayFinalizeWithdrawERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayInitializedIterator struct {
	Event *L1ERC721GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayInitialized)
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
		it.Event = new(L1ERC721GatewayInitialized)
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
func (it *L1ERC721GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayInitialized represents a Initialized event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1ERC721GatewayInitializedIterator, error) {

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayInitializedIterator{contract: _L1ERC721Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayInitialized)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseInitialized(log types.Log) (*L1ERC721GatewayInitialized, error) {
	event := new(L1ERC721GatewayInitialized)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayOwnershipTransferredIterator struct {
	Event *L1ERC721GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayOwnershipTransferred)
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
		it.Event = new(L1ERC721GatewayOwnershipTransferred)
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
func (it *L1ERC721GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1ERC721GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayOwnershipTransferredIterator{contract: _L1ERC721Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayOwnershipTransferred)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1ERC721GatewayOwnershipTransferred, error) {
	event := new(L1ERC721GatewayOwnershipTransferred)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayRefundERC721Iterator is returned from FilterRefundERC721 and is used to iterate over the raw logs and unpacked data for RefundERC721 events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayRefundERC721Iterator struct {
	Event *L1ERC721GatewayRefundERC721 // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayRefundERC721)
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
		it.Event = new(L1ERC721GatewayRefundERC721)
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
func (it *L1ERC721GatewayRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayRefundERC721 represents a RefundERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC721 is a free log retrieval operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1ERC721GatewayRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayRefundERC721Iterator{contract: _L1ERC721Gateway.contract, event: "RefundERC721", logs: logs, sub: sub}, nil
}

// WatchRefundERC721 is a free log subscription operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchRefundERC721(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayRefundERC721)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
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

// ParseRefundERC721 is a log parse operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseRefundERC721(log types.Log) (*L1ERC721GatewayRefundERC721, error) {
	event := new(L1ERC721GatewayRefundERC721)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721GatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L1ERC721Gateway contract.
type L1ERC721GatewayUpdateTokenMappingIterator struct {
	Event *L1ERC721GatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L1ERC721GatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721GatewayUpdateTokenMapping)
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
		it.Event = new(L1ERC721GatewayUpdateTokenMapping)
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
func (it *L1ERC721GatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721GatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayUpdateTokenMapping struct {
	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (*L1ERC721GatewayUpdateTokenMappingIterator, error) {

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

	logs, sub, err := _L1ERC721Gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayUpdateTokenMappingIterator{contract: _L1ERC721Gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L1ERC721GatewayUpdateTokenMapping, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC721Gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721GatewayUpdateTokenMapping)
				if err := _L1ERC721Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
func (_L1ERC721Gateway *L1ERC721GatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L1ERC721GatewayUpdateTokenMapping, error) {
	event := new(L1ERC721GatewayUpdateTokenMapping)
	if err := _L1ERC721Gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
