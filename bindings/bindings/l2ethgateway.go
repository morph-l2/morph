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

// L2ETHGatewayMetaData contains all meta data concerning the L2ETHGateway contract.
var L2ETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"counterpart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeDepositETH\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_counterpart\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETHAndCall\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"FinalizeDepositETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6116f6806100e65f395ff3fe6080604052600436106100b8575f3560e01c8063797594b011610071578063c7cdea371161004c578063c7cdea37146101d5578063f2fde38b146101e8578063f887ea4014610207575f80fd5b8063797594b0146101605780638da5cb5b1461018c578063c0c53b8b146101b6575f80fd5b80633cb747bf116100a15780633cb747bf146100e45780636dc2418314610139578063715018a61461014c575f80fd5b8063232e8748146100bc5780632fcc29fa146100d1575b5f80fd5b6100cf6100ca3660046111f9565b610233565b005b6100cf6100df366004611290565b61059c565b3480156100ef575f80fd5b506099546101109073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100cf610147366004611383565b6105d9565b348015610157575f80fd5b506100cf6105eb565b34801561016b575f80fd5b506097546101109073ffffffffffffffffffffffffffffffffffffffff1681565b348015610197575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610110565b3480156101c1575f80fd5b506100cf6101d0366004611420565b6105fe565b6100cf6101e3366004611468565b61080c565b3480156101f3575f80fd5b506100cf610202366004611488565b61081b565b348015610212575f80fd5b506098546101109073ffffffffffffffffffffffffffffffffffffffff1681565b60995473ffffffffffffffffffffffffffffffffffffffff163381146102ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610303573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032791906114aa565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146103ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016102b1565b6103b36108d2565b83341461041c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d73672e76616c7565206d69736d61746368000000000000000000000000000060448201526064016102b1565b5f8573ffffffffffffffffffffffffffffffffffffffff16856040515f6040518083038185875af1925050503d805f8114610472576040519150601f19603f3d011682016040523d82523d5f602084013e610477565b606091505b50509050806104e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016102b1565b6105218685858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061094592505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d878787604051610582939291906114c5565b60405180910390a35061059460018055565b505050505050565b6105d483835f5b6040519080825280601f01601f1916602001820160405280156105cd576020820181803683370190505b50846109f1565b505050565b6105e5848484846109f1565b50505050565b6105f3610cc5565b6105fc5f610d46565b565b5f54610100900460ff161580801561061c57505f54600160ff909116105b806106355750303b15801561063557505f5460ff166001145b6106c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102b1565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561071d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff831661079a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016102b1565b6107a5848484610dbc565b80156105e5575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b61081733835f6105a3565b5050565b610823610cc5565b73ffffffffffffffffffffffffffffffffffffffff81166108c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102b1565b6108cf81610d46565b50565b60026001540361093e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102b1565b6002600155565b5f815111801561096b57505f8273ffffffffffffffffffffffffffffffffffffffff163b115b15610817576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f906109c2908490600401611583565b5f604051808303815f87803b1580156109d9575f80fd5b505af1158015610594573d5f803e3d5ffd5b60018055565b6109f96108d2565b5f3411610a62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f7769746864726177207a65726f2065746800000000000000000000000000000060448201526064016102b1565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610a9d5782806020019051810190610a989190611595565b935090505b5f81868686604051602401610ab5949392919061161d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8eaac8a30000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa158015610b99573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bbd9190611665565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b923492610c22929116908b9088908b9060040161167c565b5f604051808303818588803b158015610c39575f80fd5b505af1158015610c4b573d5f803e3d5ffd5b50505050508673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0888885604051610cb1939291906116c1565b60405180910390a35050506105e560018055565b60655473ffffffffffffffffffffffffffffffffffffffff1633146105fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b1565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff8316610e39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016102b1565b73ffffffffffffffffffffffffffffffffffffffff8116610eb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016102b1565b610ebe610f67565b610ec6611005565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105d4576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff16610ffd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102b1565b6105fc6110a3565b5f54610100900460ff1661109b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102b1565b6105fc611139565b5f54610100900460ff166109eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102b1565b5f54610100900460ff166111cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102b1565b6105fc33610d46565b73ffffffffffffffffffffffffffffffffffffffff811681146108cf575f80fd5b5f805f805f6080868803121561120d575f80fd5b8535611218816111d8565b94506020860135611228816111d8565b935060408601359250606086013567ffffffffffffffff8082111561124b575f80fd5b818801915088601f83011261125e575f80fd5b81358181111561126c575f80fd5b89602082850101111561127d575f80fd5b9699959850939650602001949392505050565b5f805f606084860312156112a2575f80fd5b83356112ad816111d8565b95602085013595506040909401359392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611336576113366112c2565b604052919050565b5f67ffffffffffffffff821115611357576113576112c2565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f805f8060808587031215611396575f80fd5b84356113a1816111d8565b935060208501359250604085013567ffffffffffffffff8111156113c3575f80fd5b8501601f810187136113d3575f80fd5b80356113e66113e18261133e565b6112ef565b8181528860208385010111156113fa575f80fd5b816020840160208301375f91810160200191909152949793965093946060013593505050565b5f805f60608486031215611432575f80fd5b833561143d816111d8565b9250602084013561144d816111d8565b9150604084013561145d816111d8565b809150509250925092565b5f8060408385031215611479575f80fd5b50508035926020909101359150565b5f60208284031215611498575f80fd5b81356114a3816111d8565b9392505050565b5f602082840312156114ba575f80fd5b81516114a3816111d8565b83815260406020820152816040820152818360608301375f818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b5f5b8381101561153257818101518382015260200161151a565b50505f910152565b5f8151808452611551816020860160208601611518565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6114a3602083018461153a565b5f80604083850312156115a6575f80fd5b82516115b1816111d8565b602084015190925067ffffffffffffffff8111156115cd575f80fd5b8301601f810185136115dd575f80fd5b80516115eb6113e18261133e565b8181528660208385010111156115ff575f80fd5b611610826020830160208601611518565b8093505050509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261165b608083018461153a565b9695505050505050565b5f60208284031215611675575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f6116b0608083018561153a565b905082606083015295945050505050565b838152606060208201525f6116d9606083018561153a565b905082604083015294935050505056fea164736f6c6343000818000a",
}

// L2ETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ETHGatewayMetaData.ABI instead.
var L2ETHGatewayABI = L2ETHGatewayMetaData.ABI

// L2ETHGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ETHGatewayMetaData.Bin instead.
var L2ETHGatewayBin = L2ETHGatewayMetaData.Bin

// DeployL2ETHGateway deploys a new Ethereum contract, binding an instance of L2ETHGateway to it.
func DeployL2ETHGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ETHGateway, error) {
	parsed, err := L2ETHGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ETHGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ETHGateway{L2ETHGatewayCaller: L2ETHGatewayCaller{contract: contract}, L2ETHGatewayTransactor: L2ETHGatewayTransactor{contract: contract}, L2ETHGatewayFilterer: L2ETHGatewayFilterer{contract: contract}}, nil
}

// L2ETHGateway is an auto generated Go binding around an Ethereum contract.
type L2ETHGateway struct {
	L2ETHGatewayCaller     // Read-only binding to the contract
	L2ETHGatewayTransactor // Write-only binding to the contract
	L2ETHGatewayFilterer   // Log filterer for contract events
}

// L2ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ETHGatewaySession struct {
	Contract     *L2ETHGateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ETHGatewayCallerSession struct {
	Contract *L2ETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L2ETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ETHGatewayTransactorSession struct {
	Contract     *L2ETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2ETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ETHGatewayRaw struct {
	Contract *L2ETHGateway // Generic contract binding to access the raw methods on
}

// L2ETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ETHGatewayCallerRaw struct {
	Contract *L2ETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2ETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ETHGatewayTransactorRaw struct {
	Contract *L2ETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ETHGateway creates a new instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGateway(address common.Address, backend bind.ContractBackend) (*L2ETHGateway, error) {
	contract, err := bindL2ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ETHGateway{L2ETHGatewayCaller: L2ETHGatewayCaller{contract: contract}, L2ETHGatewayTransactor: L2ETHGatewayTransactor{contract: contract}, L2ETHGatewayFilterer: L2ETHGatewayFilterer{contract: contract}}, nil
}

// NewL2ETHGatewayCaller creates a new read-only instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ETHGatewayCaller, error) {
	contract, err := bindL2ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayCaller{contract: contract}, nil
}

// NewL2ETHGatewayTransactor creates a new write-only instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ETHGatewayTransactor, error) {
	contract, err := bindL2ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayTransactor{contract: contract}, nil
}

// NewL2ETHGatewayFilterer creates a new log filterer instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ETHGatewayFilterer, error) {
	contract, err := bindL2ETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayFilterer{contract: contract}, nil
}

// bindL2ETHGateway binds a generic wrapper to an already deployed contract.
func bindL2ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ETHGateway *L2ETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ETHGateway.Contract.L2ETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ETHGateway *L2ETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.L2ETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ETHGateway *L2ETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.L2ETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ETHGateway *L2ETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ETHGateway *L2ETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ETHGateway *L2ETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ETHGateway *L2ETHGatewaySession) Counterpart() (common.Address, error) {
	return _L2ETHGateway.Contract.Counterpart(&_L2ETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2ETHGateway.Contract.Counterpart(&_L2ETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ETHGateway *L2ETHGatewaySession) Messenger() (common.Address, error) {
	return _L2ETHGateway.Contract.Messenger(&_L2ETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2ETHGateway.Contract.Messenger(&_L2ETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ETHGateway *L2ETHGatewaySession) Owner() (common.Address, error) {
	return _L2ETHGateway.Contract.Owner(&_L2ETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCallerSession) Owner() (common.Address, error) {
	return _L2ETHGateway.Contract.Owner(&_L2ETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ETHGateway *L2ETHGatewaySession) Router() (common.Address, error) {
	return _L2ETHGateway.Contract.Router(&_L2ETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCallerSession) Router() (common.Address, error) {
	return _L2ETHGateway.Contract.Router(&_L2ETHGateway.CallOpts)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) FinalizeDepositETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "finalizeDepositETH", _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ETHGateway *L2ETHGatewaySession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.FinalizeDepositETH(&_L2ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.FinalizeDepositETH(&_L2ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ETHGateway *L2ETHGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.Initialize(&_L2ETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.Initialize(&_L2ETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ETHGateway *L2ETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ETHGateway.Contract.RenounceOwnership(&_L2ETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ETHGateway.Contract.RenounceOwnership(&_L2ETHGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ETHGateway *L2ETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.TransferOwnership(&_L2ETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.TransferOwnership(&_L2ETHGateway.TransactOpts, newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewaySession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETH(&_L2ETHGateway.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETH(&_L2ETHGateway.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewaySession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETH0(&_L2ETHGateway.TransactOpts, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETH0(&_L2ETHGateway.TransactOpts, _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewaySession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETHAndCall(&_L2ETHGateway.TransactOpts, _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactorSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.Contract.WithdrawETHAndCall(&_L2ETHGateway.TransactOpts, _to, _amount, _data, _gasLimit)
}

// L2ETHGatewayFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the L2ETHGateway contract.
type L2ETHGatewayFinalizeDepositETHIterator struct {
	Event *L2ETHGatewayFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *L2ETHGatewayFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ETHGatewayFinalizeDepositETH)
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
		it.Event = new(L2ETHGatewayFinalizeDepositETH)
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
func (it *L2ETHGatewayFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ETHGatewayFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ETHGatewayFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2ETHGateway contract.
type L2ETHGatewayFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ETHGateway *L2ETHGatewayFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2ETHGatewayFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ETHGateway.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayFinalizeDepositETHIterator{contract: _L2ETHGateway.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ETHGateway *L2ETHGatewayFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *L2ETHGatewayFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ETHGateway.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ETHGatewayFinalizeDepositETH)
				if err := _L2ETHGateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ETHGateway *L2ETHGatewayFilterer) ParseFinalizeDepositETH(log types.Log) (*L2ETHGatewayFinalizeDepositETH, error) {
	event := new(L2ETHGatewayFinalizeDepositETH)
	if err := _L2ETHGateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ETHGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ETHGateway contract.
type L2ETHGatewayInitializedIterator struct {
	Event *L2ETHGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2ETHGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ETHGatewayInitialized)
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
		it.Event = new(L2ETHGatewayInitialized)
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
func (it *L2ETHGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ETHGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ETHGatewayInitialized represents a Initialized event raised by the L2ETHGateway contract.
type L2ETHGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ETHGateway *L2ETHGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ETHGatewayInitializedIterator, error) {

	logs, sub, err := _L2ETHGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayInitializedIterator{contract: _L2ETHGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ETHGateway *L2ETHGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ETHGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ETHGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ETHGatewayInitialized)
				if err := _L2ETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ETHGateway *L2ETHGatewayFilterer) ParseInitialized(log types.Log) (*L2ETHGatewayInitialized, error) {
	event := new(L2ETHGatewayInitialized)
	if err := _L2ETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2ETHGateway contract.
type L2ETHGatewayOwnershipTransferredIterator struct {
	Event *L2ETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2ETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ETHGatewayOwnershipTransferred)
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
		it.Event = new(L2ETHGatewayOwnershipTransferred)
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
func (it *L2ETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ETHGateway contract.
type L2ETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ETHGateway *L2ETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2ETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayOwnershipTransferredIterator{contract: _L2ETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ETHGateway *L2ETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2ETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ETHGatewayOwnershipTransferred)
				if err := _L2ETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2ETHGateway *L2ETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2ETHGatewayOwnershipTransferred, error) {
	event := new(L2ETHGatewayOwnershipTransferred)
	if err := _L2ETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ETHGatewayWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the L2ETHGateway contract.
type L2ETHGatewayWithdrawETHIterator struct {
	Event *L2ETHGatewayWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L2ETHGatewayWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ETHGatewayWithdrawETH)
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
		it.Event = new(L2ETHGatewayWithdrawETH)
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
func (it *L2ETHGatewayWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ETHGatewayWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ETHGatewayWithdrawETH represents a WithdrawETH event raised by the L2ETHGateway contract.
type L2ETHGatewayWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2ETHGateway *L2ETHGatewayFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2ETHGatewayWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ETHGateway.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayWithdrawETHIterator{contract: _L2ETHGateway.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2ETHGateway *L2ETHGatewayFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *L2ETHGatewayWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ETHGateway.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ETHGatewayWithdrawETH)
				if err := _L2ETHGateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x22b1de295ba82e3c7a822438f4741347553ea2d59af4e3b98febc5af9d77d0d0.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data, uint256 nonce)
func (_L2ETHGateway *L2ETHGatewayFilterer) ParseWithdrawETH(log types.Log) (*L2ETHGatewayWithdrawETH, error) {
	event := new(L2ETHGatewayWithdrawETH)
	if err := _L2ETHGateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
