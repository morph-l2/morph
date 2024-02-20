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

// ProxyAdminMetaData contains all meta data concerning the ProxyAdmin contract.
var ProxyAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"contractAddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"implementationName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgrading\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxyType\",\"outputs\":[{\"internalType\":\"enumProxyAdmin.ProxyType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAddressManager\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setImplementationName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumProxyAdmin.ProxyType\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"setProxyType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_upgrading\",\"type\":\"bool\"}],\"name\":\"setUpgrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620019f3380380620019f383398101604081905262000033916200009f565b6200003e3362000050565b620000498162000050565b50620000ce565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215620000b0575f80fd5b81516001600160a01b0381168114620000c7575f80fd5b9392505050565b61191780620000dc5f395ff3fe608060405260043610610109575f3560e01c8063860f7cda116100a157806399a88ec411610071578063b794726211610057578063b794726214610314578063f2fde38b1461034e578063f3b7dead1461036d575f80fd5b806399a88ec4146102d65780639b2ea4bd146102f5575f80fd5b8063860f7cda1461025c5780638d52d4a01461027b5780638da5cb5b1461029a5780639623609d146102c3575f80fd5b80633ab76e9f116100dc5780633ab76e9f146101c25780636bd9f516146101ee578063715018a6146102295780637eff275e1461023d575f80fd5b80630652b57a1461010d57806307c8f7b01461012e578063204e1c7a1461014d578063238181ae14610196575b5f80fd5b348015610118575f80fd5b5061012c6101273660046111ab565b61038c565b005b348015610139575f80fd5b5061012c6101483660046111c6565b6103db565b348015610158575f80fd5b5061016c6101673660046111ab565b61042d565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101a1575f80fd5b506101b56101b03660046111ab565b61064b565b60405161018d9190611250565b3480156101cd575f80fd5b5060035461016c9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101f9575f80fd5b5061021c6102083660046111ab565b60016020525f908152604090205460ff1681565b60405161018d919061128f565b348015610234575f80fd5b5061012c6106e2565b348015610248575f80fd5b5061012c6102573660046112ce565b6106f5565b348015610267575f80fd5b5061012c61027636600461141f565b6108a2565b348015610286575f80fd5b5061012c61029536600461146c565b6108d8565b3480156102a5575f80fd5b505f5473ffffffffffffffffffffffffffffffffffffffff1661016c565b61012c6102d136600461149b565b61094b565b3480156102e1575f80fd5b5061012c6102f03660046112ce565b610b59565b348015610300575f80fd5b5061012c61030f36600461150c565b610de0565b34801561031f575f80fd5b5060035474010000000000000000000000000000000000000000900460ff16604051901515815260200161018d565b348015610359575f80fd5b5061012c6103683660046111ab565b610e71565b348015610378575f80fd5b5061016c6103873660046111ab565b610f28565b610394611096565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6103e3611096565b6003805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001602052604081205460ff168181600281111561046857610468611262565b036104e1578273ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104da9190611550565b9392505050565b60018160028111156104f5576104f5611262565b03610543578273ffffffffffffffffffffffffffffffffffffffff1663aaf10f426040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b6573d5f803e3d5ffd5b600281600281111561055757610557611262565b036105de5760035473ffffffffffffffffffffffffffffffffffffffff8481165f908152600260205260409081902090517fbf40fac1000000000000000000000000000000000000000000000000000000008152919092169163bf40fac1916105c391906004016115b6565b602060405180830381865afa1580156104b6573d5f803e3d5ffd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f50726f787941646d696e3a20756e6b6e6f776e2070726f78792074797065000060448201526064015b60405180910390fd5b50919050565b60026020525f9081526040902080546106639061156b565b80601f016020809104026020016040519081016040528092919081815260200182805461068f9061156b565b80156106da5780601f106106b1576101008083540402835291602001916106da565b820191905f5260205f20905b8154815290600101906020018083116106bd57829003601f168201915b505050505081565b6106ea611096565b6106f35f611116565b565b6106fd611096565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526001602052604081205460ff169081600281111561073857610738611262565b036107bf576040517f8f28397000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690638f283970906024015b5f604051808303815f87803b1580156107a4575f80fd5b505af11580156107b6573d5f803e3d5ffd5b50505050505050565b60018160028111156107d3576107d3611262565b0361082c576040517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301528416906313af40359060240161078d565b600281600281111561084057610840611262565b036105de576003546040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529091169063f2fde38b9060240161078d565b505050565b6108aa611096565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260026020526040902061089d82826116a2565b6108e0611096565b73ffffffffffffffffffffffffffffffffffffffff82165f908152600160208190526040909120805483927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009091169083600281111561094257610942611262565b02179055505050565b610953611096565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526001602052604081205460ff169081600281111561098e5761098e611262565b03610a50576040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851690634f1ef2869034906109e990879087906004016117ba565b5f6040518083038185885af1158015610a04573d5f803e3d5ffd5b50505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610a4a91908101906117f0565b50610b53565b610a5a8484610b59565b5f8473ffffffffffffffffffffffffffffffffffffffff163484604051610a819190611862565b5f6040518083038185875af1925050503d805f8114610abb576040519150601f19603f3d011682016040523d82523d5f602084013e610ac0565b606091505b5050905080610b51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f50726f787941646d696e3a2063616c6c20746f2070726f78792061667465722060448201527f75706772616465206661696c6564000000000000000000000000000000000000606482015260840161063c565b505b50505050565b610b61611096565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526001602052604081205460ff1690816002811115610b9c57610b9c611262565b03610bf5576040517f3659cfe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690633659cfe69060240161078d565b6001816002811115610c0957610c09611262565b03610c88576040517f9b0b0fda0000000000000000000000000000000000000000000000000000000081527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152841690639b0b0fda9060440161078d565b6002816002811115610c9c57610c9c611262565b03610dd85773ffffffffffffffffffffffffffffffffffffffff83165f9081526002602052604081208054610cd09061156b565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfc9061156b565b8015610d475780601f10610d1e57610100808354040283529160200191610d47565b820191905f5260205f20905b815481529060010190602001808311610d2a57829003601f168201915b50506003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815294955073ffffffffffffffffffffffffffffffffffffffff1693639b2ea4bd9350610da59250859150879060040161187d565b5f604051808303815f87803b158015610dbc575f80fd5b505af1158015610dce573d5f803e3d5ffd5b5050505050505050565b61089d6118b4565b610de8611096565b6003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690639b2ea4bd90610e40908590859060040161187d565b5f604051808303815f87803b158015610e57575f80fd5b505af1158015610e69573d5f803e3d5ffd5b505050505050565b610e79611096565b73ffffffffffffffffffffffffffffffffffffffff8116610f1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161063c565b610f2581611116565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001602052604081205460ff1681816002811115610f6357610f63611262565b03610fb1578273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b6573d5f803e3d5ffd5b6001816002811115610fc557610fc5611262565b03611013578273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b6573d5f803e3d5ffd5b600281600281111561102757611027611262565b036105de5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b6573d5f803e3d5ffd5b5f5473ffffffffffffffffffffffffffffffffffffffff1633146106f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161063c565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f25575f80fd5b5f602082840312156111bb575f80fd5b81356104da8161118a565b5f602082840312156111d6575f80fd5b813580151581146104da575f80fd5b5f5b838110156111ff5781810151838201526020016111e7565b50505f910152565b5f815180845261121e8160208601602086016111e5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6104da6020830184611207565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208101600383106112c8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b5f80604083850312156112df575f80fd5b82356112ea8161118a565b915060208301356112fa8161118a565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561137957611379611305565b604052919050565b5f67ffffffffffffffff82111561139a5761139a611305565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f6113d86113d384611381565b611332565b90508281528383830111156113eb575f80fd5b828260208301375f602084830101529392505050565b5f82601f830112611410575f80fd5b6104da838335602085016113c6565b5f8060408385031215611430575f80fd5b823561143b8161118a565b9150602083013567ffffffffffffffff811115611456575f80fd5b61146285828601611401565b9150509250929050565b5f806040838503121561147d575f80fd5b82356114888161118a565b91506020830135600381106112fa575f80fd5b5f805f606084860312156114ad575f80fd5b83356114b88161118a565b925060208401356114c88161118a565b9150604084013567ffffffffffffffff8111156114e3575f80fd5b8401601f810186136114f3575f80fd5b611502868235602084016113c6565b9150509250925092565b5f806040838503121561151d575f80fd5b823567ffffffffffffffff811115611533575f80fd5b61153f85828601611401565b92505060208301356112fa8161118a565b5f60208284031215611560575f80fd5b81516104da8161118a565b600181811c9082168061157f57607f821691505b602082108103610645577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60208083525f84546115c88161156b565b806020870152604060018084165f81146115e9576001811461162357611650565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00851660408a0152604084151560051b8a01019550611650565b895f5260205f205f5b858110156116475781548b820186015290830190880161162c565b8a016040019650505b509398975050505050505050565b601f82111561089d57805f5260205f20601f840160051c810160208510156116835750805b601f840160051c820191505b81811015610b51575f815560010161168f565b815167ffffffffffffffff8111156116bc576116bc611305565b6116d0816116ca845461156b565b8461165e565b602080601f831160018114611722575f84156116ec5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610e69565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561176e5788860151825594840194600190910190840161174f565b50858210156117aa57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6117e86040830184611207565b949350505050565b5f60208284031215611800575f80fd5b815167ffffffffffffffff811115611816575f80fd5b8201601f81018413611826575f80fd5b80516118346113d382611381565b818152856020838501011115611848575f80fd5b6118598260208301602086016111e5565b95945050505050565b5f82516118738184602087016111e5565b9190910192915050565b604081525f61188f6040830185611207565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea26469706673582212207ba73a704064be947692740673a37361648165e4bb185068aa172b4623e0b7d264736f6c63430008180033",
}

// ProxyAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyAdminMetaData.ABI instead.
var ProxyAdminABI = ProxyAdminMetaData.ABI

// ProxyAdminBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyAdminMetaData.Bin instead.
var ProxyAdminBin = ProxyAdminMetaData.Bin

// DeployProxyAdmin deploys a new Ethereum contract, binding an instance of ProxyAdmin to it.
func DeployProxyAdmin(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ProxyAdmin, error) {
	parsed, err := ProxyAdminMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyAdminBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// ProxyAdmin is an auto generated Go binding around an Ethereum contract.
type ProxyAdmin struct {
	ProxyAdminCaller     // Read-only binding to the contract
	ProxyAdminTransactor // Write-only binding to the contract
	ProxyAdminFilterer   // Log filterer for contract events
}

// ProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminSession struct {
	Contract     *ProxyAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminCallerSession struct {
	Contract *ProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminTransactorSession struct {
	Contract     *ProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminRaw struct {
	Contract *ProxyAdmin // Generic contract binding to access the raw methods on
}

// ProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminCallerRaw struct {
	Contract *ProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminTransactorRaw struct {
	Contract *ProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdmin creates a new instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdmin(address common.Address, backend bind.ContractBackend) (*ProxyAdmin, error) {
	contract, err := bindProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// NewProxyAdminCaller creates a new read-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminCaller, error) {
	contract, err := bindProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminCaller{contract: contract}, nil
}

// NewProxyAdminTransactor creates a new write-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminTransactor, error) {
	contract, err := bindProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminTransactor{contract: contract}, nil
}

// NewProxyAdminFilterer creates a new log filterer instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminFilterer, error) {
	contract, err := bindProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminFilterer{contract: contract}, nil
}

// bindProxyAdmin binds a generic wrapper to an already deployed contract.
func bindProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.ProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) AddressManager() (common.Address, error) {
	return _ProxyAdmin.Contract.AddressManager(&_ProxyAdmin.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) AddressManager() (common.Address, error) {
	return _ProxyAdmin.Contract.AddressManager(&_ProxyAdmin.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, _proxy)
}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminCaller) ImplementationName(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "implementationName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminSession) ImplementationName(arg0 common.Address) (string, error) {
	return _ProxyAdmin.Contract.ImplementationName(&_ProxyAdmin.CallOpts, arg0)
}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminCallerSession) ImplementationName(arg0 common.Address) (string, error) {
	return _ProxyAdmin.Contract.ImplementationName(&_ProxyAdmin.CallOpts, arg0)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminCaller) IsUpgrading(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "isUpgrading")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminSession) IsUpgrading() (bool, error) {
	return _ProxyAdmin.Contract.IsUpgrading(&_ProxyAdmin.CallOpts)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminCallerSession) IsUpgrading() (bool, error) {
	return _ProxyAdmin.Contract.IsUpgrading(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminCaller) ProxyType(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "proxyType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminSession) ProxyType(arg0 common.Address) (uint8, error) {
	return _ProxyAdmin.Contract.ProxyType(&_ProxyAdmin.CallOpts, arg0)
}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminCallerSession) ProxyType(arg0 common.Address) (uint8, error) {
	return _ProxyAdmin.Contract.ProxyType(&_ProxyAdmin.CallOpts, arg0)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, _proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "changeProxyAdmin", _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, _proxy, _newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetAddress(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setAddress", _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddress(&_ProxyAdmin.TransactOpts, _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddress(&_ProxyAdmin.TransactOpts, _name, _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetAddressManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setAddressManager", _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminSession) SetAddressManager(_address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddressManager(&_ProxyAdmin.TransactOpts, _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetAddressManager(_address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddressManager(&_ProxyAdmin.TransactOpts, _address)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetImplementationName(opts *bind.TransactOpts, _address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setImplementationName", _address, _name)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminSession) SetImplementationName(_address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetImplementationName(&_ProxyAdmin.TransactOpts, _address, _name)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetImplementationName(_address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetImplementationName(&_ProxyAdmin.TransactOpts, _address, _name)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetProxyType(opts *bind.TransactOpts, _address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setProxyType", _address, _type)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminSession) SetProxyType(_address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetProxyType(&_ProxyAdmin.TransactOpts, _address, _type)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetProxyType(_address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetProxyType(&_ProxyAdmin.TransactOpts, _address, _type)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetUpgrading(opts *bind.TransactOpts, _upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setUpgrading", _upgrading)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminSession) SetUpgrading(_upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetUpgrading(&_ProxyAdmin.TransactOpts, _upgrading)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetUpgrading(_upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetUpgrading(&_ProxyAdmin.TransactOpts, _upgrading)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgrade", _proxy, _implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminSession) Upgrade(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, _proxy, _implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Upgrade(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, _proxy, _implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgradeAndCall", _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, _proxy, _implementation, _data)
}

// ProxyAdminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferredIterator struct {
	Event *ProxyAdminOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyAdminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminOwnershipTransferred)
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
		it.Event = new(ProxyAdminOwnershipTransferred)
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
func (it *ProxyAdminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminOwnershipTransferredIterator{contract: _ProxyAdmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminOwnershipTransferred)
				if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProxyAdmin *ProxyAdminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminOwnershipTransferred, error) {
	event := new(ProxyAdminOwnershipTransferred)
	if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
