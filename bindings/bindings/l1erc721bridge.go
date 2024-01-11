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

// L1ERC721BridgeMetaData contains all meta data concerning the L1ERC721Bridge contract.
var L1ERC721BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_otherBridge\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeERC721\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeERC721To\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeBridgeERC721\",\"inputs\":[{\"name\":\"_localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"otherBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ERC721BridgeFinalized\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC721BridgeInitiated\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200163438038062001634833981016040819052620000359162000161565b6001808084846001600160a01b038216620000ac5760405162461bcd60e51b815260206004820152602c60248201527f4552433732314272696467653a206d657373656e6765722063616e6e6f74206260448201526b65206164647265737328302960a01b60648201526084015b60405180910390fd5b6001600160a01b0381166200011c5760405162461bcd60e51b815260206004820152602f60248201527f4552433732314272696467653a206f74686572206272696467652063616e6e6f60448201526e74206265206164647265737328302960881b6064820152608401620000a3565b6001600160a01b039182166080521660a05260c09290925260e0526101005250620001999050565b80516001600160a01b03811681146200015c57600080fd5b919050565b600080604083850312156200017557600080fd5b620001808362000144565b9150620001906020840162000144565b90509250929050565b60805160a05160c05160e051610100516114266200020e6000396000610324015260006102fb015260006102d2015260008181610195015281816101fb015281816103b00152610cc101526000818160da015281816101c401528181610386015281816103e70152610c9201526114266000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80637f46ddb211610076578063927ede2d1161005b578063927ede2d146101bf578063aa557452146101e6578063c89701a2146101f957600080fd5b80637f46ddb2146101905780638129fc1c146101b757600080fd5b806354fd4d50116100a757806354fd4d50146101245780635d93a3fc14610139578063761f44931461017d57600080fd5b80633687011a146100c35780633cb747bf146100d8575b600080fd5b6100d66100d1366004610fd1565b61021f565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012c6102cb565b60405161011b91906110c2565b61016d6101473660046110dc565b603260209081526000938452604080852082529284528284209052825290205460ff1681565b604051901515815260200161011b565b6100d661018b36600461111d565b61036e565b6100fa7f000000000000000000000000000000000000000000000000000000000000000081565b6100d66107ef565b6100fa7f000000000000000000000000000000000000000000000000000000000000000081565b6100d66101f43660046111b5565b610979565b7f00000000000000000000000000000000000000000000000000000000000000006100fa565b333b156102b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c38686333388888888610a35565b505050505050565b60606102f67f0000000000000000000000000000000000000000000000000000000000000000610dac565b61031f7f0000000000000000000000000000000000000000000000000000000000000000610dac565b6103487f0000000000000000000000000000000000000000000000000000000000000000610dac565b60405160200161035a9392919061122c565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561048c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610450573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047491906112a2565b73ffffffffffffffffffffffffffffffffffffffff16145b610518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016102aa565b3073ffffffffffffffffffffffffffffffffffffffff8816036105bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016102aa565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603260209081526040808320938a1683529281528282208683529052205460ff16151560011461068c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c31204272696467650000000000000060648201526084016102aa565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526032602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b15801561074c57600080fd5b505af1158015610760573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107de9493929190611308565b60405180910390a450505050505050565b603154610100900460ff161580801561080f5750603154600160ff909116105b806108295750303b158015610829575060315460ff166001145b6108b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102aa565b603180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561091357603180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b801561097657603180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b73ffffffffffffffffffffffffffffffffffffffff8516610a1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016102aa565b610a2c8787338888888888610a35565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff8716610ad8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016102aa565b600063761f449360e01b888a8989898888604051602401610aff9796959493929190611348565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603286528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610c3f57600080fd5b505af1158015610c53573d6000803e3d6000fd5b50506040517f3dbb202b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169250633dbb202b9150610ced907f000000000000000000000000000000000000000000000000000000000000000090859089906004016113a5565b600060405180830381600087803b158015610d0757600080fd5b505af1158015610d1b573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610d999493929190611308565b60405180910390a4505050505050505050565b60606000610db983610e6a565b600101905060008167ffffffffffffffff811115610dd957610dd96113ea565b6040519080825280601f01601f191660200182016040528015610e03576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610e0d57509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610eb3577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610edf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610efd57662386f26fc10000830492506010015b6305f5e1008310610f15576305f5e100830492506008015b6127108310610f2957612710830492506004015b60648310610f3b576064830492506002015b600a8310610f47576001015b92915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461097657600080fd5b803563ffffffff81168114610f8357600080fd5b919050565b60008083601f840112610f9a57600080fd5b50813567ffffffffffffffff811115610fb257600080fd5b602083019150836020828501011115610fca57600080fd5b9250929050565b60008060008060008060a08789031215610fea57600080fd5b8635610ff581610f4d565b9550602087013561100581610f4d565b94506040870135935061101a60608801610f6f565b9250608087013567ffffffffffffffff81111561103657600080fd5b61104289828a01610f88565b979a9699509497509295939492505050565b60005b8381101561106f578181015183820152602001611057565b50506000910152565b60008151808452611090816020860160208601611054565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006110d56020830184611078565b9392505050565b6000806000606084860312156110f157600080fd5b83356110fc81610f4d565b9250602084013561110c81610f4d565b929592945050506040919091013590565b600080600080600080600060c0888a03121561113857600080fd5b873561114381610f4d565b9650602088013561115381610f4d565b9550604088013561116381610f4d565b9450606088013561117381610f4d565b93506080880135925060a088013567ffffffffffffffff81111561119657600080fd5b6111a28a828b01610f88565b989b979a50959850939692959293505050565b600080600080600080600060c0888a0312156111d057600080fd5b87356111db81610f4d565b965060208801356111eb81610f4d565b955060408801356111fb81610f4d565b94506060880135935061121060808901610f6f565b925060a088013567ffffffffffffffff81111561119657600080fd5b6000845161123e818460208901611054565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161127a816001850160208a01611054565b60019201918201528351611295816002840160208801611054565b0160020195945050505050565b6000602082840312156112b457600080fd5b81516110d581610f4d565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061133e6060830184866112bf565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261139860c0830184866112bf565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006113d46060830185611078565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000810000a",
}

// L1ERC721BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC721BridgeMetaData.ABI instead.
var L1ERC721BridgeABI = L1ERC721BridgeMetaData.ABI

// L1ERC721BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1ERC721BridgeMetaData.Bin instead.
var L1ERC721BridgeBin = L1ERC721BridgeMetaData.Bin

// DeployL1ERC721Bridge deploys a new Ethereum contract, binding an instance of L1ERC721Bridge to it.
func DeployL1ERC721Bridge(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address, _otherBridge common.Address) (common.Address, *types.Transaction, *L1ERC721Bridge, error) {
	parsed, err := L1ERC721BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1ERC721BridgeBin), backend, _messenger, _otherBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1ERC721Bridge{L1ERC721BridgeCaller: L1ERC721BridgeCaller{contract: contract}, L1ERC721BridgeTransactor: L1ERC721BridgeTransactor{contract: contract}, L1ERC721BridgeFilterer: L1ERC721BridgeFilterer{contract: contract}}, nil
}

// L1ERC721Bridge is an auto generated Go binding around an Ethereum contract.
type L1ERC721Bridge struct {
	L1ERC721BridgeCaller     // Read-only binding to the contract
	L1ERC721BridgeTransactor // Write-only binding to the contract
	L1ERC721BridgeFilterer   // Log filterer for contract events
}

// L1ERC721BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC721BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC721BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC721BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC721BridgeSession struct {
	Contract     *L1ERC721Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC721BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC721BridgeCallerSession struct {
	Contract *L1ERC721BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L1ERC721BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC721BridgeTransactorSession struct {
	Contract     *L1ERC721BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L1ERC721BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC721BridgeRaw struct {
	Contract *L1ERC721Bridge // Generic contract binding to access the raw methods on
}

// L1ERC721BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC721BridgeCallerRaw struct {
	Contract *L1ERC721BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC721BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC721BridgeTransactorRaw struct {
	Contract *L1ERC721BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC721Bridge creates a new instance of L1ERC721Bridge, bound to a specific deployed contract.
func NewL1ERC721Bridge(address common.Address, backend bind.ContractBackend) (*L1ERC721Bridge, error) {
	contract, err := bindL1ERC721Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC721Bridge{L1ERC721BridgeCaller: L1ERC721BridgeCaller{contract: contract}, L1ERC721BridgeTransactor: L1ERC721BridgeTransactor{contract: contract}, L1ERC721BridgeFilterer: L1ERC721BridgeFilterer{contract: contract}}, nil
}

// NewL1ERC721BridgeCaller creates a new read-only instance of L1ERC721Bridge, bound to a specific deployed contract.
func NewL1ERC721BridgeCaller(address common.Address, caller bind.ContractCaller) (*L1ERC721BridgeCaller, error) {
	contract, err := bindL1ERC721Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeCaller{contract: contract}, nil
}

// NewL1ERC721BridgeTransactor creates a new write-only instance of L1ERC721Bridge, bound to a specific deployed contract.
func NewL1ERC721BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC721BridgeTransactor, error) {
	contract, err := bindL1ERC721Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeTransactor{contract: contract}, nil
}

// NewL1ERC721BridgeFilterer creates a new log filterer instance of L1ERC721Bridge, bound to a specific deployed contract.
func NewL1ERC721BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC721BridgeFilterer, error) {
	contract, err := bindL1ERC721Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeFilterer{contract: contract}, nil
}

// bindL1ERC721Bridge binds a generic wrapper to an already deployed contract.
func bindL1ERC721Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC721BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC721Bridge *L1ERC721BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC721Bridge.Contract.L1ERC721BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC721Bridge *L1ERC721BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.L1ERC721BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC721Bridge *L1ERC721BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.L1ERC721BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC721Bridge *L1ERC721BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC721Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC721Bridge *L1ERC721BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC721Bridge *L1ERC721BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeSession) MESSENGER() (common.Address, error) {
	return _L1ERC721Bridge.Contract.MESSENGER(&_L1ERC721Bridge.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) MESSENGER() (common.Address, error) {
	return _L1ERC721Bridge.Contract.MESSENGER(&_L1ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeSession) OTHERBRIDGE() (common.Address, error) {
	return _L1ERC721Bridge.Contract.OTHERBRIDGE(&_L1ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _L1ERC721Bridge.Contract.OTHERBRIDGE(&_L1ERC721Bridge.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x5d93a3fc.
//
// Solidity: function deposits(address , address , uint256 ) view returns(bool)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "deposits", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x5d93a3fc.
//
// Solidity: function deposits(address , address , uint256 ) view returns(bool)
func (_L1ERC721Bridge *L1ERC721BridgeSession) Deposits(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _L1ERC721Bridge.Contract.Deposits(&_L1ERC721Bridge.CallOpts, arg0, arg1, arg2)
}

// Deposits is a free data retrieval call binding the contract method 0x5d93a3fc.
//
// Solidity: function deposits(address , address , uint256 ) view returns(bool)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) Deposits(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _L1ERC721Bridge.Contract.Deposits(&_L1ERC721Bridge.CallOpts, arg0, arg1, arg2)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeSession) Messenger() (common.Address, error) {
	return _L1ERC721Bridge.Contract.Messenger(&_L1ERC721Bridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) Messenger() (common.Address, error) {
	return _L1ERC721Bridge.Contract.Messenger(&_L1ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeSession) OtherBridge() (common.Address, error) {
	return _L1ERC721Bridge.Contract.OtherBridge(&_L1ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) OtherBridge() (common.Address, error) {
	return _L1ERC721Bridge.Contract.OtherBridge(&_L1ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1ERC721Bridge *L1ERC721BridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1ERC721Bridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1ERC721Bridge *L1ERC721BridgeSession) Version() (string, error) {
	return _L1ERC721Bridge.Contract.Version(&_L1ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1ERC721Bridge *L1ERC721BridgeCallerSession) Version() (string, error) {
	return _L1ERC721Bridge.Contract.Version(&_L1ERC721Bridge.CallOpts)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactor) BridgeERC721(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.contract.Transact(opts, "bridgeERC721", _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeSession) BridgeERC721(_localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.BridgeERC721(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactorSession) BridgeERC721(_localToken common.Address, _remoteToken common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.BridgeERC721(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactor) BridgeERC721To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.contract.Transact(opts, "bridgeERC721To", _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeSession) BridgeERC721To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.BridgeERC721To(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactorSession) BridgeERC721To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.BridgeERC721To(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactor) FinalizeBridgeERC721(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.contract.Transact(opts, "finalizeBridgeERC721", _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeSession) FinalizeBridgeERC721(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.FinalizeBridgeERC721(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactorSession) FinalizeBridgeERC721(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.FinalizeBridgeERC721(&_L1ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Bridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L1ERC721Bridge *L1ERC721BridgeSession) Initialize() (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.Initialize(&_L1ERC721Bridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L1ERC721Bridge *L1ERC721BridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _L1ERC721Bridge.Contract.Initialize(&_L1ERC721Bridge.TransactOpts)
}

// L1ERC721BridgeERC721BridgeFinalizedIterator is returned from FilterERC721BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC721BridgeFinalized events raised by the L1ERC721Bridge contract.
type L1ERC721BridgeERC721BridgeFinalizedIterator struct {
	Event *L1ERC721BridgeERC721BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L1ERC721BridgeERC721BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721BridgeERC721BridgeFinalized)
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
		it.Event = new(L1ERC721BridgeERC721BridgeFinalized)
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
func (it *L1ERC721BridgeERC721BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721BridgeERC721BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721BridgeERC721BridgeFinalized represents a ERC721BridgeFinalized event raised by the L1ERC721Bridge contract.
type L1ERC721BridgeERC721BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeFinalized is a free log retrieval operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) FilterERC721BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1ERC721BridgeERC721BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeERC721BridgeFinalizedIterator{contract: _L1ERC721Bridge.contract, event: "ERC721BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeFinalized is a free log subscription operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) WatchERC721BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L1ERC721BridgeERC721BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721BridgeERC721BridgeFinalized)
				if err := _L1ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
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

// ParseERC721BridgeFinalized is a log parse operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) ParseERC721BridgeFinalized(log types.Log) (*L1ERC721BridgeERC721BridgeFinalized, error) {
	event := new(L1ERC721BridgeERC721BridgeFinalized)
	if err := _L1ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721BridgeERC721BridgeInitiatedIterator is returned from FilterERC721BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC721BridgeInitiated events raised by the L1ERC721Bridge contract.
type L1ERC721BridgeERC721BridgeInitiatedIterator struct {
	Event *L1ERC721BridgeERC721BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L1ERC721BridgeERC721BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721BridgeERC721BridgeInitiated)
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
		it.Event = new(L1ERC721BridgeERC721BridgeInitiated)
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
func (it *L1ERC721BridgeERC721BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721BridgeERC721BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721BridgeERC721BridgeInitiated represents a ERC721BridgeInitiated event raised by the L1ERC721Bridge contract.
type L1ERC721BridgeERC721BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeInitiated is a free log retrieval operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) FilterERC721BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1ERC721BridgeERC721BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeERC721BridgeInitiatedIterator{contract: _L1ERC721Bridge.contract, event: "ERC721BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeInitiated is a free log subscription operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) WatchERC721BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L1ERC721BridgeERC721BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721BridgeERC721BridgeInitiated)
				if err := _L1ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
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

// ParseERC721BridgeInitiated is a log parse operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) ParseERC721BridgeInitiated(log types.Log) (*L1ERC721BridgeERC721BridgeInitiated, error) {
	event := new(L1ERC721BridgeERC721BridgeInitiated)
	if err := _L1ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC721BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1ERC721Bridge contract.
type L1ERC721BridgeInitializedIterator struct {
	Event *L1ERC721BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L1ERC721BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC721BridgeInitialized)
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
		it.Event = new(L1ERC721BridgeInitialized)
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
func (it *L1ERC721BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC721BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC721BridgeInitialized represents a Initialized event raised by the L1ERC721Bridge contract.
type L1ERC721BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1ERC721BridgeInitializedIterator, error) {

	logs, sub, err := _L1ERC721Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1ERC721BridgeInitializedIterator{contract: _L1ERC721Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1ERC721BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L1ERC721Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC721BridgeInitialized)
				if err := _L1ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1ERC721Bridge *L1ERC721BridgeFilterer) ParseInitialized(log types.Log) (*L1ERC721BridgeInitialized, error) {
	event := new(L1ERC721BridgeInitialized)
	if err := _L1ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
