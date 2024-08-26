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

// L1StandardERC20GatewayMetaData contains all meta data concerning the L1StandardERC20Gateway contract.
var L1StandardERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2TokenImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2TokenFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b62000021565b620000df565b5f54610100900460ff16156200008d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614620000dd575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61252780620000ed5f395ff3fe6080604052600436106100e4575f3560e01c8063797594b011610087578063eddd5e8211610057578063eddd5e821461021e578063f219fa661461023d578063f2fde38b14610250578063f887ea401461026f575f80fd5b8063797594b0146101b057806384bd13b0146101cf5780638da5cb5b146101e2578063c676ad29146101ff575f80fd5b80631459457a116100c25780631459457a1461014b57806321425ee01461016a5780633cb747bf1461017d578063715018a61461019c575f80fd5b80630aea8c26146100e85780630e28c1f2146100fd57806314298c5114610138575b5f80fd5b6100fb6100f6366004611e16565b61028e565b005b348015610108575f80fd5b5060fa5461011c906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b6100fb610146366004611eca565b6102a2565b348015610156575f80fd5b506100fb610165366004611f09565b610527565b6100fb610178366004611f76565b6107eb565b348015610188575f80fd5b5060995461011c906001600160a01b031681565b3480156101a7575f80fd5b506100fb610824565b3480156101bb575f80fd5b5060975461011c906001600160a01b031681565b6100fb6101dd366004611fa8565b610837565b3480156101ed575f80fd5b506065546001600160a01b031661011c565b34801561020a575f80fd5b5061011c61021936600461203a565b610a27565b348015610229575f80fd5b5060fb5461011c906001600160a01b031681565b6100fb61024b366004612055565b610b4e565b34801561025b575f80fd5b506100fb61026a36600461203a565b610b60565b34801561027a575f80fd5b5060985461011c906001600160a01b031681565b61029b8585858585610bf0565b5050505050565b6099546001600160a01b03163381146103025760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561033e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103629190612098565b6001600160a01b0316736f297c61b5c92ef107ffd30cd56affe5a273e8416001600160a01b0316146103d65760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e74657874000000000060448201526064016102f9565b6103de611031565b7f8431f5c10000000000000000000000000000000000000000000000000000000061040c60045f85876120b3565b610415916120da565b7fffffffff0000000000000000000000000000000000000000000000000000000016146104845760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642073656c6563746f720000000000000000000000000000000060448201526064016102f9565b5f808061049485600481896120b3565b8101906104a19190612122565b5094505093505092506104b583838361108a565b6104c96001600160a01b03841683836110d8565b816001600160a01b0316836001600160a01b03167fdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a88360405161050e91815260200190565b60405180910390a350505061052260018055565b505050565b5f54610100900460ff161580801561054557505f54600160ff909116105b8061055e5750303b15801561055e57505f5460ff166001145b6105d05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102f9565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561062c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0385166106825760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016102f9565b61068d868686611187565b6001600160a01b0383166106e35760405162461bcd60e51b815260206004820152601860248201527f7a65726f20696d706c656d656e746174696f6e2068617368000000000000000060448201526064016102f9565b6001600160a01b0382166107395760405162461bcd60e51b815260206004820152601460248201527f7a65726f20666163746f7279206164647265737300000000000000000000000060448201526064016102f9565b60fa80546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fb80549285169290911691909117905580156107e3575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6105228333845f5b6040519080825280601f01601f19166020018201604052801561081d576020820181803683370190505b5085610bf0565b61082c6112ca565b6108355f611324565b565b6099546001600160a01b03163381146108925760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064016102f9565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f29190612098565b6097546001600160a01b0390811691161461094f5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016102f9565b610957611031565b6109668888888888888861138d565b61097a6001600160a01b03891686866110d8565b6109b98584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061156492505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a788888888604051610a0c94939291906121ad565b60405180910390a4610a1d60018055565b5050505050505050565b6097546040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b1660208201525f9182916001600160a01b039091169060340160405160208183030381529060405280519060200120604051602001610ac392919060609290921b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000168252601482015260340190565b60408051808303601f1901815290829052805160209091012060fa5460fb546001600160a01b0390811660388501526f5af43d82803e903d91602b57fd5bf3ff6024850152166014830152733d602d80600a3d3981f3363d3d373d3d3d363d738252605882018190526037600c830120607883015260556043909201919091209091505b9392505050565b610b5a8484845f6107f3565b50505050565b610b686112ca565b6001600160a01b038116610be45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102f9565b610bed81611324565b50565b610bf8611031565b5f8311610c475760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e740000000000000000000000000060448201526064016102f9565b5f610c538685856115f4565b6001600160a01b03808a165f90815260fc602052604090205492975090955091925016606081610e2257610c8688610a27565b91505f886001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015610cc4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ceb9190810190612242565b90505f896001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015610d29573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d509190810190612242565b90505f8a6001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db39190612287565b9050600188848484604051602001610dcd939291906122d2565b60408051601f1981840301815290829052610deb929160200161230a565b60408051601f1981840301815290829052610e099291602001612337565b6040516020818303038152906040529350505050610e47565b5f85604051602001610e35929190612337565b60405160208183030381529060405290505b5f8883858a8a86604051602401610e6396959493929190612351565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8431f5c10000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa158015610f1c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f40919061239e565b6099546097546040517f5f7b15770000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692635f7b1577923492610f9a929116905f9088908d908d906004016123b5565b5f604051808303818588803b158015610fb1575f80fd5b505af1158015610fc3573d5f803e3d5ffd5b5050505050846001600160a01b0316846001600160a01b03168b6001600160a01b03167f1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad886728c8c8c8760405161101b94939291906123f7565b60405180910390a4505050505061029b60018055565b6002600154036110835760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102f9565b6002600155565b34156105225760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016102f9565b6040516001600160a01b0383166024820152604481018290526105229084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261184a565b60018055565b6001600160a01b0383166111dd5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016102f9565b6001600160a01b0381166112335760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016102f9565b61123b611930565b6112436119b4565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561052257609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6065546001600160a01b031633146108355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102f9565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b34156113db5760405162461bcd60e51b815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016102f9565b6001600160a01b0386166114315760405162461bcd60e51b815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016102f9565b856001600160a01b031661144488610a27565b6001600160a01b03161461149a5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016102f9565b6001600160a01b038088165f90815260fc60205260409020541680611503576001600160a01b038881165f90815260fc6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016918916919091179055610a1d565b866001600160a01b0316816001600160a01b031614610a1d5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016102f9565b5f815111801561157d57505f826001600160a01b03163b115b156115f0576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f906115c790849060040161242f565b5f604051808303815f87803b1580156115de575f80fd5b505af11580156107e3573d5f803e3d5ffd5b5050565b6098545f908190606090339081906001600160a01b03168190036116c457858060200190518101906116269190612441565b6040517fc52a3bbc0000000000000000000000000000000000000000000000000000000081526001600160a01b0380841660048301528b81166024830152604482018b905291985091925083169063c52a3bbc906064016020604051808303815f875af1158015611699573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116bd919061239e565b96506117ee565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038a16906370a0823190602401602060405180830381865afa158015611721573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611745919061239e565b905061175c6001600160a01b038a1683308b611a38565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038b16906370a0823190602401602060405180830381865afa1580156117b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117dd919061239e565b90506117e982826124a2565b985050505b5f871161183d5760405162461bcd60e51b815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e740000000000000000000000000060448201526064016102f9565b9795965093949350505050565b5f61189e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611a899092919063ffffffff16565b905080515f14806118be5750808060200190518101906118be91906124e0565b6105225760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102f9565b5f54610100900460ff166119ac5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102f9565b610835611a9f565b5f54610100900460ff16611a305760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102f9565b610835611b1b565b6040516001600160a01b0380851660248301528316604482015260648101829052610b5a9085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161111d565b6060611a9784845f85611ba0565b949350505050565b5f54610100900460ff166111815760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102f9565b5f54610100900460ff16611b975760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102f9565b61083533611324565b606082471015611c185760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102f9565b5f80866001600160a01b03168587604051611c3391906124ff565b5f6040518083038185875af1925050503d805f8114611c6d576040519150601f19603f3d011682016040523d82523d5f602084013e611c72565b606091505b5091509150611c8387838387611c8e565b979650505050505050565b60608315611cfc5782515f03611cf5576001600160a01b0385163b611cf55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102f9565b5081611a97565b611a978383815115611d115781518083602001fd5b8060405162461bcd60e51b81526004016102f9919061242f565b6001600160a01b0381168114610bed575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d9557611d95611d3f565b604052919050565b5f67ffffffffffffffff821115611db657611db6611d3f565b50601f01601f191660200190565b5f82601f830112611dd3575f80fd5b8135611de6611de182611d9d565b611d6c565b818152846020838601011115611dfa575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611e2a575f80fd5b8535611e3581611d2b565b94506020860135611e4581611d2b565b935060408601359250606086013567ffffffffffffffff811115611e67575f80fd5b611e7388828901611dc4565b95989497509295608001359392505050565b5f8083601f840112611e95575f80fd5b50813567ffffffffffffffff811115611eac575f80fd5b602083019150836020828501011115611ec3575f80fd5b9250929050565b5f8060208385031215611edb575f80fd5b823567ffffffffffffffff811115611ef1575f80fd5b611efd85828601611e85565b90969095509350505050565b5f805f805f60a08688031215611f1d575f80fd5b8535611f2881611d2b565b94506020860135611f3881611d2b565b93506040860135611f4881611d2b565b92506060860135611f5881611d2b565b91506080860135611f6881611d2b565b809150509295509295909350565b5f805f60608486031215611f88575f80fd5b8335611f9381611d2b565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a031215611fbe575f80fd5b8735611fc981611d2b565b96506020880135611fd981611d2b565b95506040880135611fe981611d2b565b94506060880135611ff981611d2b565b93506080880135925060a088013567ffffffffffffffff81111561201b575f80fd5b6120278a828b01611e85565b989b979a50959850939692959293505050565b5f6020828403121561204a575f80fd5b8135610b4781611d2b565b5f805f8060808587031215612068575f80fd5b843561207381611d2b565b9350602085013561208381611d2b565b93969395505050506040820135916060013590565b5f602082840312156120a8575f80fd5b8151610b4781611d2b565b5f80858511156120c1575f80fd5b838611156120cd575f80fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561211a5780818660040360031b1b83161692505b505092915050565b5f805f805f8060c08789031215612137575f80fd5b863561214281611d2b565b9550602087013561215281611d2b565b9450604087013561216281611d2b565b9350606087013561217281611d2b565b92506080870135915060a087013567ffffffffffffffff811115612194575f80fd5b6121a089828a01611dc4565b9150509295509295509295565b6001600160a01b038516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f909201601f191601019392505050565b5f5b8381101561220c5781810151838201526020016121f4565b50505f910152565b5f612221611de184611d9d565b9050828152838383011115612234575f80fd5b610b478360208301846121f2565b5f60208284031215612252575f80fd5b815167ffffffffffffffff811115612268575f80fd5b8201601f81018413612278575f80fd5b611a9784825160208401612214565b5f60208284031215612297575f80fd5b815160ff81168114610b47575f80fd5b5f81518084526122be8160208601602086016121f2565b601f01601f19169290920160200192915050565b606081525f6122e460608301866122a7565b82810360208401526122f681866122a7565b91505060ff83166040830152949350505050565b604081525f61231c60408301856122a7565b828103602084015261232e81856122a7565b95945050505050565b8215158152604060208201525f611a9760408301846122a7565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261239260c08301846122a7565b98975050505050505050565b5f602082840312156123ae575f80fd5b5051919050565b5f6001600160a01b03808816835286602084015260a060408401526123dd60a08401876122a7565b606084019590955292909216608090910152509392505050565b6001600160a01b0385168152836020820152608060408201525f61241e60808301856122a7565b905082606083015295945050505050565b602081525f610b4760208301846122a7565b5f8060408385031215612452575f80fd5b825161245d81611d2b565b602084015190925067ffffffffffffffff811115612479575f80fd5b8301601f81018513612489575f80fd5b61249885825160208401612214565b9150509250929050565b818103818111156124da577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b5f602082840312156124f0575f80fd5b81518015158114610b47575f80fd5b5f82516125108184602087016121f2565b919091019291505056fea164736f6c6343000818000a",
}

// L1StandardERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StandardERC20GatewayMetaData.ABI instead.
var L1StandardERC20GatewayABI = L1StandardERC20GatewayMetaData.ABI

// L1StandardERC20GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1StandardERC20GatewayMetaData.Bin instead.
var L1StandardERC20GatewayBin = L1StandardERC20GatewayMetaData.Bin

// DeployL1StandardERC20Gateway deploys a new Ethereum contract, binding an instance of L1StandardERC20Gateway to it.
func DeployL1StandardERC20Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1StandardERC20Gateway, error) {
	parsed, err := L1StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1StandardERC20GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1StandardERC20Gateway{L1StandardERC20GatewayCaller: L1StandardERC20GatewayCaller{contract: contract}, L1StandardERC20GatewayTransactor: L1StandardERC20GatewayTransactor{contract: contract}, L1StandardERC20GatewayFilterer: L1StandardERC20GatewayFilterer{contract: contract}}, nil
}

// L1StandardERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L1StandardERC20Gateway struct {
	L1StandardERC20GatewayCaller     // Read-only binding to the contract
	L1StandardERC20GatewayTransactor // Write-only binding to the contract
	L1StandardERC20GatewayFilterer   // Log filterer for contract events
}

// L1StandardERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StandardERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StandardERC20GatewaySession struct {
	Contract     *L1StandardERC20Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L1StandardERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StandardERC20GatewayCallerSession struct {
	Contract *L1StandardERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L1StandardERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StandardERC20GatewayTransactorSession struct {
	Contract     *L1StandardERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L1StandardERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StandardERC20GatewayRaw struct {
	Contract *L1StandardERC20Gateway // Generic contract binding to access the raw methods on
}

// L1StandardERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayCallerRaw struct {
	Contract *L1StandardERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1StandardERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayTransactorRaw struct {
	Contract *L1StandardERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1StandardERC20Gateway creates a new instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20Gateway(address common.Address, backend bind.ContractBackend) (*L1StandardERC20Gateway, error) {
	contract, err := bindL1StandardERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20Gateway{L1StandardERC20GatewayCaller: L1StandardERC20GatewayCaller{contract: contract}, L1StandardERC20GatewayTransactor: L1StandardERC20GatewayTransactor{contract: contract}, L1StandardERC20GatewayFilterer: L1StandardERC20GatewayFilterer{contract: contract}}, nil
}

// NewL1StandardERC20GatewayCaller creates a new read-only instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1StandardERC20GatewayCaller, error) {
	contract, err := bindL1StandardERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayCaller{contract: contract}, nil
}

// NewL1StandardERC20GatewayTransactor creates a new write-only instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StandardERC20GatewayTransactor, error) {
	contract, err := bindL1StandardERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayTransactor{contract: contract}, nil
}

// NewL1StandardERC20GatewayFilterer creates a new log filterer instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StandardERC20GatewayFilterer, error) {
	contract, err := bindL1StandardERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayFilterer{contract: contract}, nil
}

// bindL1StandardERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL1StandardERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Counterpart() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Counterpart(&_L1StandardERC20Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Counterpart(&_L1StandardERC20Gateway.CallOpts)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.GetL2ERC20Address(&_L1StandardERC20Gateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.GetL2ERC20Address(&_L1StandardERC20Gateway.CallOpts, _l1Token)
}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) L2TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "l2TokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) L2TokenFactory() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenFactory(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) L2TokenFactory() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenFactory(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) L2TokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "l2TokenImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) L2TokenImplementation() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenImplementation(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) L2TokenImplementation() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenImplementation(&_L1StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Messenger() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Messenger(&_L1StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Messenger() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Messenger(&_L1StandardERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Owner() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Owner(&_L1StandardERC20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Owner() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Owner(&_L1StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Router() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Router(&_L1StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Router(&_L1StandardERC20Gateway.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20(&_L1StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20(&_L1StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC200(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC200(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20AndCall(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20AndCall(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.FinalizeWithdrawERC20(&_L1StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.FinalizeWithdrawERC20(&_L1StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.Initialize(&_L1StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.Initialize(&_L1StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.OnDropMessage(&_L1StandardERC20Gateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.OnDropMessage(&_L1StandardERC20Gateway.TransactOpts, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.RenounceOwnership(&_L1StandardERC20Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.RenounceOwnership(&_L1StandardERC20Gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.TransferOwnership(&_L1StandardERC20Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.TransferOwnership(&_L1StandardERC20Gateway.TransactOpts, newOwner)
}

// L1StandardERC20GatewayDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayDepositERC20Iterator struct {
	Event *L1StandardERC20GatewayDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayDepositERC20)
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
		it.Event = new(L1StandardERC20GatewayDepositERC20)
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
func (it *L1StandardERC20GatewayDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayDepositERC20 represents a DepositERC20 event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayDepositERC20 struct {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardERC20GatewayDepositERC20Iterator, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayDepositERC20Iterator{contract: _L1StandardERC20Gateway.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x1a6c38816de45937fd5cd974f9694fe10e64163ba12a92abf0f4b6b23ad88672.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayDepositERC20)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseDepositERC20(log types.Log) (*L1StandardERC20GatewayDepositERC20, error) {
	event := new(L1StandardERC20GatewayDepositERC20)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardERC20GatewayFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayFinalizeWithdrawERC20Iterator struct {
	Event *L1StandardERC20GatewayFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayFinalizeWithdrawERC20)
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
		it.Event = new(L1StandardERC20GatewayFinalizeWithdrawERC20)
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
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayFinalizeWithdrawERC20 struct {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardERC20GatewayFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayFinalizeWithdrawERC20Iterator{contract: _L1StandardERC20Gateway.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayFinalizeWithdrawERC20)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1StandardERC20GatewayFinalizeWithdrawERC20, error) {
	event := new(L1StandardERC20GatewayFinalizeWithdrawERC20)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardERC20GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayInitializedIterator struct {
	Event *L1StandardERC20GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayInitialized)
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
		it.Event = new(L1StandardERC20GatewayInitialized)
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
func (it *L1StandardERC20GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayInitialized represents a Initialized event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1StandardERC20GatewayInitializedIterator, error) {

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayInitializedIterator{contract: _L1StandardERC20Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayInitialized)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseInitialized(log types.Log) (*L1StandardERC20GatewayInitialized, error) {
	event := new(L1StandardERC20GatewayInitialized)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardERC20GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayOwnershipTransferredIterator struct {
	Event *L1StandardERC20GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayOwnershipTransferred)
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
		it.Event = new(L1StandardERC20GatewayOwnershipTransferred)
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
func (it *L1StandardERC20GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1StandardERC20GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayOwnershipTransferredIterator{contract: _L1StandardERC20Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayOwnershipTransferred)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1StandardERC20GatewayOwnershipTransferred, error) {
	event := new(L1StandardERC20GatewayOwnershipTransferred)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardERC20GatewayRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayRefundERC20Iterator struct {
	Event *L1StandardERC20GatewayRefundERC20 // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayRefundERC20)
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
		it.Event = new(L1StandardERC20GatewayRefundERC20)
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
func (it *L1StandardERC20GatewayRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayRefundERC20 represents a RefundERC20 event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayRefundERC20 struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterRefundERC20(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1StandardERC20GatewayRefundERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayRefundERC20Iterator{contract: _L1StandardERC20Gateway.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
//
// Solidity: event RefundERC20(address indexed token, address indexed recipient, uint256 amount)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayRefundERC20, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "RefundERC20", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayRefundERC20)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseRefundERC20(log types.Log) (*L1StandardERC20GatewayRefundERC20, error) {
	event := new(L1StandardERC20GatewayRefundERC20)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
