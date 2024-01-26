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

// L1ETHGatewayMetaData contains all meta data concerning the L1ETHGateway contract.
var L1ETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundETH\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611a59806100ec6000396000f3fe6080604052600436106100c75760003560e01c80639f8420b311610074578063ce0b63ce1161004e578063ce0b63ce146101fd578063f2fde38b14610210578063f887ea401461023057600080fd5b80639f8420b3146101b7578063aac476f8146101ca578063c0c53b8b146101dd57600080fd5b8063797594b0116100a5578063797594b01461014c5780638da5cb5b146101795780638eaac8a3146101a457600080fd5b806314298c51146100cc5780633cb747bf146100e1578063715018a614610137575b600080fd5b6100df6100da366004611404565b61025d565b005b3480156100ed57600080fd5b5060995461010e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561014357600080fd5b506100df6105f9565b34801561015857600080fd5b5060975461010e9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561018557600080fd5b5060655473ffffffffffffffffffffffffffffffffffffffff1661010e565b6100df6101b2366004611468565b61060d565b6100df6101c53660046114db565b610910565b6100df6101d83660046114fd565b61094d565b3480156101e957600080fd5b506100df6101f8366004611561565b610997565b6100df61020b3660046115ac565b610b78565b34801561021c57600080fd5b506100df61022b3660046115e1565b610b84565b34801561023c57600080fd5b5060985461010e9073ffffffffffffffffffffffffffffffffffffffff1681565b60995473ffffffffffffffffffffffffffffffffffffffff163381146102ca5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610315573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103399190611605565b73ffffffffffffffffffffffffffffffffffffffff16736f297c61b5c92ef107ffd30cd56affe5a273e84173ffffffffffffffffffffffffffffffffffffffff16146103c75760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e74657874000000000060448201526064016102c1565b6103cf610c21565b7f232e8748000000000000000000000000000000000000000000000000000000006103fe600460008587611622565b6104079161164c565b7fffffffff0000000000000000000000000000000000000000000000000000000016146104765760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642073656c6563746f720000000000000000000000000000000060448201526064016102c1565b6000806104868460048188611622565b8101906104939190611758565b5092505091503481146104e85760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d61746368000000000000000000000000000060448201526064016102c1565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610542576040519150601f19603f3d011682016040523d82523d6000602084013e610547565b606091505b50509050806105985760405162461bcd60e51b815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016102c1565b8273ffffffffffffffffffffffffffffffffffffffff167f289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7836040516105e091815260200190565b60405180910390a25050506105f460018055565b505050565b610601610c80565b61060b6000610ce7565b565b60995473ffffffffffffffffffffffffffffffffffffffff163381146106755760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064016102c1565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e49190611605565b60975473ffffffffffffffffffffffffffffffffffffffff90811691161461074e5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016102c1565b610756610c21565b8334146107a55760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d61746368000000000000000000000000000060448201526064016102c1565b60008573ffffffffffffffffffffffffffffffffffffffff168560405160006040518083038185875af1925050503d80600081146107ff576040519150601f19603f3d011682016040523d82523d6000602084013e610804565b606091505b50509050806108555760405162461bcd60e51b815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016102c1565b6108958685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610d5e92505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e78787876040516108f693929190611807565b60405180910390a35061090860018055565b505050505050565b610949338360005b6040519080825280601f01601f191660200182016040528015610942576020820181803683370190505b5084610e0b565b5050565b610990858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250610e0b915050565b5050505050565b600054610100900460ff16158080156109b75750600054600160ff909116105b806109d15750303b1580156109d1575060005460ff166001145b610a435760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102c1565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610aa157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610b045760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016102c1565b610b0f848484611037565b8015610b7257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6105f483836000610918565b610b8c610c80565b73ffffffffffffffffffffffffffffffffffffffff8116610c155760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102c1565b610c1e81610ce7565b50565b600260015403610c735760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102c1565b6002600155565b60018055565b60655473ffffffffffffffffffffffffffffffffffffffff16331461060b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c1565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008151118015610d86575060008273ffffffffffffffffffffffffffffffffffffffff163b115b15610949576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f90610ddd9084906004016118c9565b600060405180830381600087803b158015610df757600080fd5b505af1158015610908573d6000803e3d6000fd5b610e13610c21565b60008311610e635760405162461bcd60e51b815260206004820152601060248201527f6465706f736974207a65726f206574680000000000000000000000000000000060448201526064016102c1565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610e9e5782806020019051810190610e9991906118dc565b935090505b600081868686604051602401610eb79493929190611969565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f232e87480000000000000000000000000000000000000000000000000000000017905260995460975491517f5f7b157700000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90811692635f7b1577923492610f93929116908a9087908a908a906004016119b2565b6000604051808303818588803b158015610fac57600080fd5b505af1158015610fc0573d6000803e3d6000fd5b50505050508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a6758787604051611024929190611a02565b60405180910390a35050610b7260018055565b73ffffffffffffffffffffffffffffffffffffffff831661109a5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016102c1565b73ffffffffffffffffffffffffffffffffffffffff81166110fd5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016102c1565b6111056111ae565b61110d611233565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105f4576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b600054610100900460ff1661122b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102c1565b61060b6112b8565b600054610100900460ff166112b05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102c1565b61060b611335565b600054610100900460ff16610c7a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102c1565b600054610100900460ff166113b25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102c1565b61060b33610ce7565b60008083601f8401126113cd57600080fd5b50813567ffffffffffffffff8111156113e557600080fd5b6020830191508360208285010111156113fd57600080fd5b9250929050565b6000806020838503121561141757600080fd5b823567ffffffffffffffff81111561142e57600080fd5b61143a858286016113bb565b90969095509350505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c1e57600080fd5b60008060008060006080868803121561148057600080fd5b853561148b81611446565b9450602086013561149b81611446565b935060408601359250606086013567ffffffffffffffff8111156114be57600080fd5b6114ca888289016113bb565b969995985093965092949392505050565b600080604083850312156114ee57600080fd5b50508035926020909101359150565b60008060008060006080868803121561151557600080fd5b853561152081611446565b945060208601359350604086013567ffffffffffffffff81111561154357600080fd5b61154f888289016113bb565b96999598509660600135949350505050565b60008060006060848603121561157657600080fd5b833561158181611446565b9250602084013561159181611446565b915060408401356115a181611446565b809150509250925092565b6000806000606084860312156115c157600080fd5b83356115cc81611446565b95602085013595506040909401359392505050565b6000602082840312156115f357600080fd5b81356115fe81611446565b9392505050565b60006020828403121561161757600080fd5b81516115fe81611446565b6000808585111561163257600080fd5b8386111561163f57600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561168c5780818660040360031b1b83161692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561170a5761170a611694565b604052919050565b600067ffffffffffffffff82111561172c5761172c611694565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000806000806080858703121561176e57600080fd5b843561177981611446565b9350602085013561178981611446565b925060408501359150606085013567ffffffffffffffff8111156117ac57600080fd5b8501601f810187136117bd57600080fd5b80356117d06117cb82611712565b6116c3565b8181528860208385010111156117e557600080fd5b8160208401602083013760006020838301015280935050505092959194509250565b83815260406020820152816040820152818360608301376000818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b60005b8381101561187657818101518382015260200161185e565b50506000910152565b6000815180845261189781602086016020860161185b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115fe602083018461187f565b600080604083850312156118ef57600080fd5b82516118fa81611446565b602084015190925067ffffffffffffffff81111561191757600080fd5b8301601f8101851361192857600080fd5b80516119366117cb82611712565b81815286602083850101111561194b57600080fd5b61195c82602083016020860161185b565b8093505050509250929050565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526119a8608083018461187f565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015260a060408401526119e860a084018761187f565b606084019590955292909216608090910152509392505050565b828152604060208201526000611a1b604083018461187f565b94935050505056fea26469706673582212205e4ed29779fba7ce4d0641f166ca065bbf10b4d5abbfd7ae0e72f437fdde294f64736f6c63430008100033",
}

// L1ETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ETHGatewayMetaData.ABI instead.
var L1ETHGatewayABI = L1ETHGatewayMetaData.ABI

// L1ETHGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1ETHGatewayMetaData.Bin instead.
var L1ETHGatewayBin = L1ETHGatewayMetaData.Bin

// DeployL1ETHGateway deploys a new Ethereum contract, binding an instance of L1ETHGateway to it.
func DeployL1ETHGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1ETHGateway, error) {
	parsed, err := L1ETHGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1ETHGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1ETHGateway{L1ETHGatewayCaller: L1ETHGatewayCaller{contract: contract}, L1ETHGatewayTransactor: L1ETHGatewayTransactor{contract: contract}, L1ETHGatewayFilterer: L1ETHGatewayFilterer{contract: contract}}, nil
}

// L1ETHGateway is an auto generated Go binding around an Ethereum contract.
type L1ETHGateway struct {
	L1ETHGatewayCaller     // Read-only binding to the contract
	L1ETHGatewayTransactor // Write-only binding to the contract
	L1ETHGatewayFilterer   // Log filterer for contract events
}

// L1ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ETHGatewaySession struct {
	Contract     *L1ETHGateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ETHGatewayCallerSession struct {
	Contract *L1ETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L1ETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ETHGatewayTransactorSession struct {
	Contract     *L1ETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L1ETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ETHGatewayRaw struct {
	Contract *L1ETHGateway // Generic contract binding to access the raw methods on
}

// L1ETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ETHGatewayCallerRaw struct {
	Contract *L1ETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1ETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ETHGatewayTransactorRaw struct {
	Contract *L1ETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ETHGateway creates a new instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGateway(address common.Address, backend bind.ContractBackend) (*L1ETHGateway, error) {
	contract, err := bindL1ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ETHGateway{L1ETHGatewayCaller: L1ETHGatewayCaller{contract: contract}, L1ETHGatewayTransactor: L1ETHGatewayTransactor{contract: contract}, L1ETHGatewayFilterer: L1ETHGatewayFilterer{contract: contract}}, nil
}

// NewL1ETHGatewayCaller creates a new read-only instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ETHGatewayCaller, error) {
	contract, err := bindL1ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayCaller{contract: contract}, nil
}

// NewL1ETHGatewayTransactor creates a new write-only instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ETHGatewayTransactor, error) {
	contract, err := bindL1ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayTransactor{contract: contract}, nil
}

// NewL1ETHGatewayFilterer creates a new log filterer instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ETHGatewayFilterer, error) {
	contract, err := bindL1ETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayFilterer{contract: contract}, nil
}

// bindL1ETHGateway binds a generic wrapper to an already deployed contract.
func bindL1ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ETHGateway *L1ETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ETHGateway.Contract.L1ETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ETHGateway *L1ETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.L1ETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ETHGateway *L1ETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.L1ETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ETHGateway *L1ETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ETHGateway *L1ETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ETHGateway *L1ETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ETHGateway *L1ETHGatewaySession) Counterpart() (common.Address, error) {
	return _L1ETHGateway.Contract.Counterpart(&_L1ETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1ETHGateway.Contract.Counterpart(&_L1ETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ETHGateway *L1ETHGatewaySession) Messenger() (common.Address, error) {
	return _L1ETHGateway.Contract.Messenger(&_L1ETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCallerSession) Messenger() (common.Address, error) {
	return _L1ETHGateway.Contract.Messenger(&_L1ETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ETHGateway *L1ETHGatewaySession) Owner() (common.Address, error) {
	return _L1ETHGateway.Contract.Owner(&_L1ETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCallerSession) Owner() (common.Address, error) {
	return _L1ETHGateway.Contract.Owner(&_L1ETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ETHGateway *L1ETHGatewaySession) Router() (common.Address, error) {
	return _L1ETHGateway.Contract.Router(&_L1ETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCallerSession) Router() (common.Address, error) {
	return _L1ETHGateway.Contract.Router(&_L1ETHGateway.CallOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETH(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETH", _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewaySession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETH(&_L1ETHGateway.TransactOpts, _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETH(&_L1ETHGateway.TransactOpts, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETH0(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETH0", _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewaySession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETH0(&_L1ETHGateway.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETH0(&_L1ETHGateway.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETHAndCall", _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewaySession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETHAndCall(&_L1ETHGateway.TransactOpts, _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.DepositETHAndCall(&_L1ETHGateway.TransactOpts, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "finalizeWithdrawETH", _from, _to, _amount, _data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ETHGateway *L1ETHGatewaySession) FinalizeWithdrawETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.FinalizeWithdrawETH(&_L1ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) FinalizeWithdrawETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.FinalizeWithdrawETH(&_L1ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ETHGateway *L1ETHGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.Initialize(&_L1ETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.Initialize(&_L1ETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ETHGateway *L1ETHGatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.OnDropMessage(&_L1ETHGateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.OnDropMessage(&_L1ETHGateway.TransactOpts, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ETHGateway *L1ETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ETHGateway.Contract.RenounceOwnership(&_L1ETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ETHGateway.Contract.RenounceOwnership(&_L1ETHGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ETHGateway *L1ETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.TransferOwnership(&_L1ETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ETHGateway *L1ETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.Contract.TransferOwnership(&_L1ETHGateway.TransactOpts, newOwner)
}

// L1ETHGatewayDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the L1ETHGateway contract.
type L1ETHGatewayDepositETHIterator struct {
	Event *L1ETHGatewayDepositETH // Event containing the contract specifics and raw log

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
func (it *L1ETHGatewayDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ETHGatewayDepositETH)
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
		it.Event = new(L1ETHGatewayDepositETH)
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
func (it *L1ETHGatewayDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ETHGatewayDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ETHGatewayDepositETH represents a DepositETH event raised by the L1ETHGateway contract.
type L1ETHGatewayDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1ETHGatewayDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ETHGateway.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayDepositETHIterator{contract: _L1ETHGateway.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *L1ETHGatewayDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ETHGateway.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ETHGatewayDepositETH)
				if err := _L1ETHGateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) ParseDepositETH(log types.Log) (*L1ETHGatewayDepositETH, error) {
	event := new(L1ETHGatewayDepositETH)
	if err := _L1ETHGateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ETHGatewayFinalizeWithdrawETHIterator is returned from FilterFinalizeWithdrawETH and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawETH events raised by the L1ETHGateway contract.
type L1ETHGatewayFinalizeWithdrawETHIterator struct {
	Event *L1ETHGatewayFinalizeWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L1ETHGatewayFinalizeWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ETHGatewayFinalizeWithdrawETH)
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
		it.Event = new(L1ETHGatewayFinalizeWithdrawETH)
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
func (it *L1ETHGatewayFinalizeWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ETHGatewayFinalizeWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ETHGatewayFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the L1ETHGateway contract.
type L1ETHGatewayFinalizeWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawETH is a free log retrieval operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) FilterFinalizeWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1ETHGatewayFinalizeWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ETHGateway.contract.FilterLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayFinalizeWithdrawETHIterator{contract: _L1ETHGateway.contract, event: "FinalizeWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawETH is a free log subscription operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) WatchFinalizeWithdrawETH(opts *bind.WatchOpts, sink chan<- *L1ETHGatewayFinalizeWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ETHGateway.contract.WatchLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ETHGatewayFinalizeWithdrawETH)
				if err := _L1ETHGateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
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

// ParseFinalizeWithdrawETH is a log parse operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L1ETHGateway *L1ETHGatewayFilterer) ParseFinalizeWithdrawETH(log types.Log) (*L1ETHGatewayFinalizeWithdrawETH, error) {
	event := new(L1ETHGatewayFinalizeWithdrawETH)
	if err := _L1ETHGateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ETHGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1ETHGateway contract.
type L1ETHGatewayInitializedIterator struct {
	Event *L1ETHGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1ETHGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ETHGatewayInitialized)
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
		it.Event = new(L1ETHGatewayInitialized)
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
func (it *L1ETHGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ETHGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ETHGatewayInitialized represents a Initialized event raised by the L1ETHGateway contract.
type L1ETHGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ETHGateway *L1ETHGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1ETHGatewayInitializedIterator, error) {

	logs, sub, err := _L1ETHGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayInitializedIterator{contract: _L1ETHGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ETHGateway *L1ETHGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1ETHGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1ETHGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ETHGatewayInitialized)
				if err := _L1ETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1ETHGateway *L1ETHGatewayFilterer) ParseInitialized(log types.Log) (*L1ETHGatewayInitialized, error) {
	event := new(L1ETHGatewayInitialized)
	if err := _L1ETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1ETHGateway contract.
type L1ETHGatewayOwnershipTransferredIterator struct {
	Event *L1ETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1ETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ETHGatewayOwnershipTransferred)
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
		it.Event = new(L1ETHGatewayOwnershipTransferred)
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
func (it *L1ETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ETHGateway contract.
type L1ETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ETHGateway *L1ETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1ETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayOwnershipTransferredIterator{contract: _L1ETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ETHGateway *L1ETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1ETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ETHGatewayOwnershipTransferred)
				if err := _L1ETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1ETHGateway *L1ETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1ETHGatewayOwnershipTransferred, error) {
	event := new(L1ETHGatewayOwnershipTransferred)
	if err := _L1ETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ETHGatewayRefundETHIterator is returned from FilterRefundETH and is used to iterate over the raw logs and unpacked data for RefundETH events raised by the L1ETHGateway contract.
type L1ETHGatewayRefundETHIterator struct {
	Event *L1ETHGatewayRefundETH // Event containing the contract specifics and raw log

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
func (it *L1ETHGatewayRefundETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ETHGatewayRefundETH)
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
		it.Event = new(L1ETHGatewayRefundETH)
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
func (it *L1ETHGatewayRefundETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ETHGatewayRefundETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ETHGatewayRefundETH represents a RefundETH event raised by the L1ETHGateway contract.
type L1ETHGatewayRefundETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundETH is a free log retrieval operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1ETHGateway *L1ETHGatewayFilterer) FilterRefundETH(opts *bind.FilterOpts, recipient []common.Address) (*L1ETHGatewayRefundETHIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ETHGateway.contract.FilterLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayRefundETHIterator{contract: _L1ETHGateway.contract, event: "RefundETH", logs: logs, sub: sub}, nil
}

// WatchRefundETH is a free log subscription operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1ETHGateway *L1ETHGatewayFilterer) WatchRefundETH(opts *bind.WatchOpts, sink chan<- *L1ETHGatewayRefundETH, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1ETHGateway.contract.WatchLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ETHGatewayRefundETH)
				if err := _L1ETHGateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
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

// ParseRefundETH is a log parse operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_L1ETHGateway *L1ETHGatewayFilterer) ParseRefundETH(log types.Log) (*L1ETHGatewayRefundETH, error) {
	event := new(L1ETHGatewayRefundETH)
	if err := _L1ETHGateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
