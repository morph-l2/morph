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

// MultipleVersionRollupVerifierMetaData contains all meta data concerning the MultipleVersionRollupVerifier contract.
var MultipleVersionRollupVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_versions\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorStartBatchIndexFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorStartBatchIndexTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"latestVerifier\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"legacyVerifiers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"legacyVerifiersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_startBatchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_aggrProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_publicInputHash\",\"type\":\"bytes32\"}],\"name\":\"verifyAggregateProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200100c3803806200100c8339810160408190526200003391620002ed565b6200003e33620001a9565b5f5b8251811015620001a0575f6001600160a01b0316828281518110620000695762000069620003ae565b60200260200101516001600160a01b031603620000995760405163a7f9319d60e01b815260040160405180910390fd5b818181518110620000ae57620000ae620003ae565b602002602001015160035f858481518110620000ce57620000ce620003ae565b602002602001015181526020019081526020015f205f0160086101000a8154816001600160a01b0302191690836001600160a01b031602179055507f7a98750a395b9ee50a2644ffda039e31f1d5d06de45510275f972bb20b229b308382815181106200013f576200013f620003ae565b60200260200101515f8484815181106200015d576200015d620003ae565b60200260200101516040516200018f9392919092835260208301919091526001600160a01b0316604082015260600190565b60405180910390a160010162000040565b505050620003c2565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715620002375762000237620001f8565b604052919050565b5f6001600160401b038211156200025a576200025a620001f8565b5060051b60200190565b5f82601f83011262000274575f80fd5b815160206200028d62000287836200023f565b6200020c565b8083825260208201915060208460051b870101935086841115620002af575f80fd5b602086015b84811015620002e25780516001600160a01b0381168114620002d4575f80fd5b8352918301918301620002b4565b509695505050505050565b5f8060408385031215620002ff575f80fd5b82516001600160401b038082111562000316575f80fd5b818501915085601f8301126200032a575f80fd5b815160206200033d62000287836200023f565b82815260059290921b840181019181810190898411156200035c575f80fd5b948201945b838610156200037c5785518252948201949082019062000361565b9188015191965090935050508082111562000395575f80fd5b50620003a48582860162000264565b9150509250929050565b634e487b7160e01b5f52603260045260245ffd5b610c3c80620003d05f395ff3fe608060405234801561000f575f80fd5b50600436106100c4575f3560e01c8063955123061161007d578063c7065b6a11610058578063c7065b6a146101da578063cb23bcb514610227578063f2fde38b14610247575f80fd5b8063955123061461016c578063bd98b2b01461017f578063c4d66de8146101c7575f80fd5b80635027ad2e116100ad5780635027ad2e1461011a578063715018a6146101475780638da5cb5b1461014f575f80fd5b806328aee03f146100c85780632c09a84814610105575b5f80fd5b6100db6100d6366004610a50565b61025a565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610118610113366004610a70565b610397565b005b610139610128366004610af0565b5f9081526002602052604090205490565b6040519081526020016100fc565b61011861042f565b5f5473ffffffffffffffffffffffffffffffffffffffff166100db565b61011861017a366004610b2f565b610442565b61019261018d366004610a50565b610779565b6040805167ffffffffffffffff909316835273ffffffffffffffffffffffffffffffffffffffff9091166020830152016100fc565b6101186101d5366004610b76565b6107d1565b6101926101e8366004610af0565b60036020525f908152604090205467ffffffffffffffff81169068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6001546100db9073ffffffffffffffffffffffffffffffffffffffff1681565b610118610255366004610b76565b6108a5565b5f82815260036020908152604080832081518083019092525467ffffffffffffffff81168083526801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff16928201929092529083101561038c575f84815260026020526040902054805b8015610389575f86815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff830190811061030a5761030a610b96565b5f9182526020918290206040805180820190915291015467ffffffffffffffff81168083526801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff16928201929092529350851015610389577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016102c3565b50505b602001519392505050565b5f6103a2868661025a565b6040517f6b40634100000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff821690636b406341906103fb90879087908790600401610bc3565b5f6040518083038186803b158015610411575f80fd5b505afa158015610423573d5f803e3d5ffd5b50505050505050505050565b61043761095c565b6104405f6109dc565b565b61044a61095c565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663059def616040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104d89190610c18565b8267ffffffffffffffff161161051a576040517f9a4ff10400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8381526003602090815260409182902082518084019093525467ffffffffffffffff8082168085526801000000000000000090920473ffffffffffffffffffffffffffffffffffffffff169284019290925290841610156105a8576040517fb0d8c70400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166105f5576040517fa7f9319d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805167ffffffffffffffff808516911610156106ba57602081015173ffffffffffffffffffffffffffffffffffffffff16156106ac575f84815260026020908152604082208054600181018255908352918190208351920180549184015173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090921667ffffffffffffffff909316929092171790555b67ffffffffffffffff831681525b73ffffffffffffffffffffffffffffffffffffffff82811660208381018281525f8881526003835260409081902086518154935190961668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931667ffffffffffffffff96871617929092179091558051888152938716918401919091528201527f7a98750a395b9ee50a2644ffda039e31f1d5d06de45510275f972bb20b229b309060600160405180910390a150505050565b6002602052815f5260405f208181548110610792575f80fd5b5f9182526020909120015467ffffffffffffffff8116925068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16905082565b6107d961095c565b60015473ffffffffffffffffffffffffffffffffffffffff161561085e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f696e697469616c697a656400000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6108ad61095c565b73ffffffffffffffffffffffffffffffffffffffff8116610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610855565b610959816109dc565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610855565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f8060408385031215610a61575f80fd5b50508035926020909101359150565b5f805f805f60808688031215610a84575f80fd5b8535945060208601359350604086013567ffffffffffffffff80821115610aa9575f80fd5b818801915088601f830112610abc575f80fd5b813581811115610aca575f80fd5b896020828501011115610adb575f80fd5b96999598505060200195606001359392505050565b5f60208284031215610b00575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b2a575f80fd5b919050565b5f805f60608486031215610b41575f80fd5b83359250602084013567ffffffffffffffff81168114610b5f575f80fd5b9150610b6d60408501610b07565b90509250925092565b5f60208284031215610b86575f80fd5b610b8f82610b07565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60408152826040820152828460608301375f606084830101525f60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601168301019050826020830152949350505050565b5f60208284031215610c28575f80fd5b505191905056fea164736f6c6343000818000a",
}

// MultipleVersionRollupVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MultipleVersionRollupVerifierMetaData.ABI instead.
var MultipleVersionRollupVerifierABI = MultipleVersionRollupVerifierMetaData.ABI

// MultipleVersionRollupVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultipleVersionRollupVerifierMetaData.Bin instead.
var MultipleVersionRollupVerifierBin = MultipleVersionRollupVerifierMetaData.Bin

// DeployMultipleVersionRollupVerifier deploys a new Ethereum contract, binding an instance of MultipleVersionRollupVerifier to it.
func DeployMultipleVersionRollupVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _versions []*big.Int, _verifiers []common.Address) (common.Address, *types.Transaction, *MultipleVersionRollupVerifier, error) {
	parsed, err := MultipleVersionRollupVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultipleVersionRollupVerifierBin), backend, _versions, _verifiers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultipleVersionRollupVerifier{MultipleVersionRollupVerifierCaller: MultipleVersionRollupVerifierCaller{contract: contract}, MultipleVersionRollupVerifierTransactor: MultipleVersionRollupVerifierTransactor{contract: contract}, MultipleVersionRollupVerifierFilterer: MultipleVersionRollupVerifierFilterer{contract: contract}}, nil
}

// MultipleVersionRollupVerifier is an auto generated Go binding around an Ethereum contract.
type MultipleVersionRollupVerifier struct {
	MultipleVersionRollupVerifierCaller     // Read-only binding to the contract
	MultipleVersionRollupVerifierTransactor // Write-only binding to the contract
	MultipleVersionRollupVerifierFilterer   // Log filterer for contract events
}

// MultipleVersionRollupVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultipleVersionRollupVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultipleVersionRollupVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultipleVersionRollupVerifierSession struct {
	Contract     *MultipleVersionRollupVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MultipleVersionRollupVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultipleVersionRollupVerifierCallerSession struct {
	Contract *MultipleVersionRollupVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MultipleVersionRollupVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultipleVersionRollupVerifierTransactorSession struct {
	Contract     *MultipleVersionRollupVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MultipleVersionRollupVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierRaw struct {
	Contract *MultipleVersionRollupVerifier // Generic contract binding to access the raw methods on
}

// MultipleVersionRollupVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierCallerRaw struct {
	Contract *MultipleVersionRollupVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MultipleVersionRollupVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultipleVersionRollupVerifierTransactorRaw struct {
	Contract *MultipleVersionRollupVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultipleVersionRollupVerifier creates a new instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifier(address common.Address, backend bind.ContractBackend) (*MultipleVersionRollupVerifier, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifier{MultipleVersionRollupVerifierCaller: MultipleVersionRollupVerifierCaller{contract: contract}, MultipleVersionRollupVerifierTransactor: MultipleVersionRollupVerifierTransactor{contract: contract}, MultipleVersionRollupVerifierFilterer: MultipleVersionRollupVerifierFilterer{contract: contract}}, nil
}

// NewMultipleVersionRollupVerifierCaller creates a new read-only instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierCaller(address common.Address, caller bind.ContractCaller) (*MultipleVersionRollupVerifierCaller, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierCaller{contract: contract}, nil
}

// NewMultipleVersionRollupVerifierTransactor creates a new write-only instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MultipleVersionRollupVerifierTransactor, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierTransactor{contract: contract}, nil
}

// NewMultipleVersionRollupVerifierFilterer creates a new log filterer instance of MultipleVersionRollupVerifier, bound to a specific deployed contract.
func NewMultipleVersionRollupVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MultipleVersionRollupVerifierFilterer, error) {
	contract, err := bindMultipleVersionRollupVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierFilterer{contract: contract}, nil
}

// bindMultipleVersionRollupVerifier binds a generic wrapper to an already deployed contract.
func bindMultipleVersionRollupVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultipleVersionRollupVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.MultipleVersionRollupVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultipleVersionRollupVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.contract.Transact(opts, method, params...)
}

// GetVerifier is a free data retrieval call binding the contract method 0x28aee03f.
//
// Solidity: function getVerifier(uint256 _version, uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) GetVerifier(opts *bind.CallOpts, _version *big.Int, _batchIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "getVerifier", _version, _batchIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x28aee03f.
//
// Solidity: function getVerifier(uint256 _version, uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) GetVerifier(_version *big.Int, _batchIndex *big.Int) (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.GetVerifier(&_MultipleVersionRollupVerifier.CallOpts, _version, _batchIndex)
}

// GetVerifier is a free data retrieval call binding the contract method 0x28aee03f.
//
// Solidity: function getVerifier(uint256 _version, uint256 _batchIndex) view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) GetVerifier(_version *big.Int, _batchIndex *big.Int) (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.GetVerifier(&_MultipleVersionRollupVerifier.CallOpts, _version, _batchIndex)
}

// LatestVerifier is a free data retrieval call binding the contract method 0xc7065b6a.
//
// Solidity: function latestVerifier(uint256 version) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LatestVerifier(opts *bind.CallOpts, version *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "latestVerifier", version)

	outstruct := new(struct {
		StartBatchIndex uint64
		Verifier        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LatestVerifier is a free data retrieval call binding the contract method 0xc7065b6a.
//
// Solidity: function latestVerifier(uint256 version) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LatestVerifier(version *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LatestVerifier(&_MultipleVersionRollupVerifier.CallOpts, version)
}

// LatestVerifier is a free data retrieval call binding the contract method 0xc7065b6a.
//
// Solidity: function latestVerifier(uint256 version) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LatestVerifier(version *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LatestVerifier(&_MultipleVersionRollupVerifier.CallOpts, version)
}

// LegacyVerifiers is a free data retrieval call binding the contract method 0xbd98b2b0.
//
// Solidity: function legacyVerifiers(uint256 version, uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LegacyVerifiers(opts *bind.CallOpts, version *big.Int, arg1 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "legacyVerifiers", version, arg1)

	outstruct := new(struct {
		StartBatchIndex uint64
		Verifier        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LegacyVerifiers is a free data retrieval call binding the contract method 0xbd98b2b0.
//
// Solidity: function legacyVerifiers(uint256 version, uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LegacyVerifiers(version *big.Int, arg1 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiers(&_MultipleVersionRollupVerifier.CallOpts, version, arg1)
}

// LegacyVerifiers is a free data retrieval call binding the contract method 0xbd98b2b0.
//
// Solidity: function legacyVerifiers(uint256 version, uint256 ) view returns(uint64 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LegacyVerifiers(version *big.Int, arg1 *big.Int) (struct {
	StartBatchIndex uint64
	Verifier        common.Address
}, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiers(&_MultipleVersionRollupVerifier.CallOpts, version, arg1)
}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0x5027ad2e.
//
// Solidity: function legacyVerifiersLength(uint256 _version) view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) LegacyVerifiersLength(opts *bind.CallOpts, _version *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "legacyVerifiersLength", _version)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0x5027ad2e.
//
// Solidity: function legacyVerifiersLength(uint256 _version) view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) LegacyVerifiersLength(_version *big.Int) (*big.Int, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiersLength(&_MultipleVersionRollupVerifier.CallOpts, _version)
}

// LegacyVerifiersLength is a free data retrieval call binding the contract method 0x5027ad2e.
//
// Solidity: function legacyVerifiersLength(uint256 _version) view returns(uint256)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) LegacyVerifiersLength(_version *big.Int) (*big.Int, error) {
	return _MultipleVersionRollupVerifier.Contract.LegacyVerifiersLength(&_MultipleVersionRollupVerifier.CallOpts, _version)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Owner() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Owner(&_MultipleVersionRollupVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) Owner() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Owner(&_MultipleVersionRollupVerifier.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Rollup() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Rollup(&_MultipleVersionRollupVerifier.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) Rollup() (common.Address, error) {
	return _MultipleVersionRollupVerifier.Contract.Rollup(&_MultipleVersionRollupVerifier.CallOpts)
}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0x2c09a848.
//
// Solidity: function verifyAggregateProof(uint256 _version, uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCaller) VerifyAggregateProof(opts *bind.CallOpts, _version *big.Int, _batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	var out []interface{}
	err := _MultipleVersionRollupVerifier.contract.Call(opts, &out, "verifyAggregateProof", _version, _batchIndex, _aggrProof, _publicInputHash)

	if err != nil {
		return err
	}

	return err

}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0x2c09a848.
//
// Solidity: function verifyAggregateProof(uint256 _version, uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) VerifyAggregateProof(_version *big.Int, _batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	return _MultipleVersionRollupVerifier.Contract.VerifyAggregateProof(&_MultipleVersionRollupVerifier.CallOpts, _version, _batchIndex, _aggrProof, _publicInputHash)
}

// VerifyAggregateProof is a free data retrieval call binding the contract method 0x2c09a848.
//
// Solidity: function verifyAggregateProof(uint256 _version, uint256 _batchIndex, bytes _aggrProof, bytes32 _publicInputHash) view returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierCallerSession) VerifyAggregateProof(_version *big.Int, _batchIndex *big.Int, _aggrProof []byte, _publicInputHash [32]byte) error {
	return _MultipleVersionRollupVerifier.Contract.VerifyAggregateProof(&_MultipleVersionRollupVerifier.CallOpts, _version, _batchIndex, _aggrProof, _publicInputHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "initialize", _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) Initialize(_rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.Initialize(&_MultipleVersionRollupVerifier.TransactOpts, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rollup) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) Initialize(_rollup common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.Initialize(&_MultipleVersionRollupVerifier.TransactOpts, _rollup)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.RenounceOwnership(&_MultipleVersionRollupVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.RenounceOwnership(&_MultipleVersionRollupVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.TransferOwnership(&_MultipleVersionRollupVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.TransferOwnership(&_MultipleVersionRollupVerifier.TransactOpts, newOwner)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x95512306.
//
// Solidity: function updateVerifier(uint256 _version, uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactor) UpdateVerifier(opts *bind.TransactOpts, _version *big.Int, _startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.contract.Transact(opts, "updateVerifier", _version, _startBatchIndex, _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x95512306.
//
// Solidity: function updateVerifier(uint256 _version, uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierSession) UpdateVerifier(_version *big.Int, _startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.UpdateVerifier(&_MultipleVersionRollupVerifier.TransactOpts, _version, _startBatchIndex, _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x95512306.
//
// Solidity: function updateVerifier(uint256 _version, uint64 _startBatchIndex, address _verifier) returns()
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierTransactorSession) UpdateVerifier(_version *big.Int, _startBatchIndex uint64, _verifier common.Address) (*types.Transaction, error) {
	return _MultipleVersionRollupVerifier.Contract.UpdateVerifier(&_MultipleVersionRollupVerifier.TransactOpts, _version, _startBatchIndex, _verifier)
}

// MultipleVersionRollupVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierOwnershipTransferredIterator struct {
	Event *MultipleVersionRollupVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultipleVersionRollupVerifierOwnershipTransferred)
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
		it.Event = new(MultipleVersionRollupVerifierOwnershipTransferred)
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
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultipleVersionRollupVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultipleVersionRollupVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MultipleVersionRollupVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultipleVersionRollupVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierOwnershipTransferredIterator{contract: _MultipleVersionRollupVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultipleVersionRollupVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultipleVersionRollupVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultipleVersionRollupVerifierOwnershipTransferred)
				if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*MultipleVersionRollupVerifierOwnershipTransferred, error) {
	event := new(MultipleVersionRollupVerifierOwnershipTransferred)
	if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultipleVersionRollupVerifierUpdateVerifierIterator is returned from FilterUpdateVerifier and is used to iterate over the raw logs and unpacked data for UpdateVerifier events raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierUpdateVerifierIterator struct {
	Event *MultipleVersionRollupVerifierUpdateVerifier // Event containing the contract specifics and raw log

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
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultipleVersionRollupVerifierUpdateVerifier)
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
		it.Event = new(MultipleVersionRollupVerifierUpdateVerifier)
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
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultipleVersionRollupVerifierUpdateVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultipleVersionRollupVerifierUpdateVerifier represents a UpdateVerifier event raised by the MultipleVersionRollupVerifier contract.
type MultipleVersionRollupVerifierUpdateVerifier struct {
	Version         *big.Int
	StartBatchIndex *big.Int
	Verifier        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifier is a free log retrieval operation binding the contract event 0x7a98750a395b9ee50a2644ffda039e31f1d5d06de45510275f972bb20b229b30.
//
// Solidity: event UpdateVerifier(uint256 version, uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) FilterUpdateVerifier(opts *bind.FilterOpts) (*MultipleVersionRollupVerifierUpdateVerifierIterator, error) {

	logs, sub, err := _MultipleVersionRollupVerifier.contract.FilterLogs(opts, "UpdateVerifier")
	if err != nil {
		return nil, err
	}
	return &MultipleVersionRollupVerifierUpdateVerifierIterator{contract: _MultipleVersionRollupVerifier.contract, event: "UpdateVerifier", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifier is a free log subscription operation binding the contract event 0x7a98750a395b9ee50a2644ffda039e31f1d5d06de45510275f972bb20b229b30.
//
// Solidity: event UpdateVerifier(uint256 version, uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) WatchUpdateVerifier(opts *bind.WatchOpts, sink chan<- *MultipleVersionRollupVerifierUpdateVerifier) (event.Subscription, error) {

	logs, sub, err := _MultipleVersionRollupVerifier.contract.WatchLogs(opts, "UpdateVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultipleVersionRollupVerifierUpdateVerifier)
				if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
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

// ParseUpdateVerifier is a log parse operation binding the contract event 0x7a98750a395b9ee50a2644ffda039e31f1d5d06de45510275f972bb20b229b30.
//
// Solidity: event UpdateVerifier(uint256 version, uint256 startBatchIndex, address verifier)
func (_MultipleVersionRollupVerifier *MultipleVersionRollupVerifierFilterer) ParseUpdateVerifier(log types.Log) (*MultipleVersionRollupVerifierUpdateVerifier, error) {
	event := new(MultipleVersionRollupVerifierUpdateVerifier)
	if err := _MultipleVersionRollupVerifier.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
