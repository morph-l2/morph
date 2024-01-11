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

// L1SequencerMetaData contains all meta data concerning the L1Sequencer contract.
var L1SequencerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_messenger\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_SEQUENCER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSequencer\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerBLSKeys\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerBLSKeysLength\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_stakingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rollupContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"newestVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rollupContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencerBLSKeys\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAndSendSequencerSet\",\"inputs\":[{\"name\":\"_sequencerBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_sequencerBLSKeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifySignature\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"indexs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerConfirmed\",\"inputs\":[{\"name\":\"sequencers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x610120604052600060025560006003553480156200001c57600080fd5b5060405162001ac638038062001ac68339810160408190526200003f916200014c565b6001600160a01b03811660805273530000000000000000000000000000000000000360a052600160c081905260e0526000610100819052805462ff0000191690556200008a62000091565b506200017e565b6200009b620000f2565b6000805462ff00001916620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620000d53390565b6040516001600160a01b03909116815260200160405180910390a1565b6200010560005462010000900460ff1690565b156200014a5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b6000602082840312156200015f57600080fd5b81516001600160a01b03811681146200017757600080fd5b9392505050565b60805160a05160c05160e051610100516118ee620001d860003960006108cb015260006108a20152600061087901526000818161036f01526105040152600081816101370152818161025701526104d701526118ee6000f3fe6080604052600436106100f75760003560e01c8063927ede2d1161008a578063bfa02ba911610059578063bfa02ba9146102dc578063e4821eb414610309578063ee99205c14610329578063f81e02a71461035d57600080fd5b8063927ede2d1461024557806396091ba1146102795780639d888e8614610299578063b00f376d146102af57600080fd5b806356002467116100c657806356002467146101c35780635c975abb146101e357806373452a921461020c5780638456cb591461023057600080fd5b80633135cb9a146101085780633cb747bf14610128578063485cc9551461018157806354fd4d50146101a157600080fd5b3661010357600080fd5b005b600080fd5b34801561011457600080fd5b506101016101233660046112fa565b610391565b34801561013457600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561018d57600080fd5b5061010161019c366004611408565b610567565b3480156101ad57600080fd5b506101b6610872565b60405161017891906114a9565b3480156101cf57600080fd5b506101b66101de3660046114c3565b610915565b3480156101ef57600080fd5b5060005462010000900460ff166040519015158152602001610178565b34801561021857600080fd5b5061022260035481565b604051908152602001610178565b34801561023c57600080fd5b50610101610a18565b34801561025157600080fd5b506101577f000000000000000000000000000000000000000000000000000000000000000081565b34801561028557600080fd5b506101b66102943660046114c3565b610ab2565b3480156102a557600080fd5b5061022260025481565b3480156102bb57600080fd5b506102226102ca3660046114e5565b60009081526004602052604090205490565b3480156102e857600080fd5b506001546101579073ffffffffffffffffffffffffffffffffffffffff1681565b34801561031557600080fd5b506101016103243660046114fe565b610b6b565b34801561033557600080fd5b50600054610157906301000000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561036957600080fd5b506101577f000000000000000000000000000000000000000000000000000000000000000081565b6000546301000000900473ffffffffffffffffffffffffffffffffffffffff16331461041e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064015b60405180910390fd5b61042782610bfd565b60005462010000900460ff161561049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e7061757365640000000000006044820152606401610415565b6040517f3dbb202b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610530907f000000000000000000000000000000000000000000000000000000000000000090879086906004016115c2565b600060405180830381600087803b15801561054a57600080fd5b505af115801561055e573d6000803e3d6000fd5b50505050505050565b600054610100900460ff16158080156105875750600054600160ff909116105b806105a15750303b1580156105a1575060005460ff166001145b61062d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610415565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561068b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610708576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f696e76616c6964207374616b696e6720636f6e747261637400000000000000006044820152606401610415565b73ffffffffffffffffffffffffffffffffffffffff8216610785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e74726163740000000000000000006044820152606401610415565b600080547fffffffffffffffffff0000000000000000000000000000000000000000ffffff16630100000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001691841691909117905561080a610cbd565b801561086d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b606061089d7f0000000000000000000000000000000000000000000000000000000000000000610d44565b6108c67f0000000000000000000000000000000000000000000000000000000000000000610d44565b6108ef7f0000000000000000000000000000000000000000000000000000000000000000610d44565b60405160200161090193929190611607565b604051602081830303815290604052905090565b600082815260046020526040902054606090801580159061094057508061093d8460016116ac565b11155b15610a00576000848152600460205260409020805484908110610965576109656116bf565b90600052602060002001805461097a906116ee565b80601f01602080910402602001604051908101604052809291908181526020018280546109a6906116ee565b80156109f35780601f106109c8576101008083540402835291602001916109f3565b820191906000526020600020905b8154815290600101906020018083116109d657829003601f168201915b5050505050915050610a12565b50506040805160208101909152600081525b92915050565b6000546301000000900473ffffffffffffffffffffffffffffffffffffffff163314610aa0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e747261637400000000000000000000006044820152606401610415565b610aa8610e02565b610ab0610cbd565b565b60046020528160005260406000208181548110610ace57600080fd5b90600052602060002001600091509150508054610aea906116ee565b80601f0160208091040260200160405190810160405280929190818152602001828054610b16906116ee565b8015610b635780601f10610b3857610100808354040283529160200191610b63565b820191906000526020600020905b815481529060010190602001808311610b4657829003601f168201915b505050505081565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610415565b610bf4610e02565b61086d83610e75565b600354600003610c0f57610c0f610f28565b60005462010000900460ff1615610c82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e7061757365640000000000006044820152606401610415565b60038054906000610c9283611741565b909155505060035460009081526004602090815260409091208251610cb9928401906110d3565b5050565b610cc5610e02565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610d1a3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60606000610d5183610f7f565b600101905060008167ffffffffffffffff811115610d7157610d716111b1565b6040519080825280601f01601f191660200182016040528015610d9b576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610da557509392505050565b60005462010000900460ff1615610ab0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610415565b6002548110158015610e8957506003548111155b610eef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e76616c69642076657273696f6e00000000000000000000000000000000006044820152606401610415565b60015b818111610f22576000818152600460205260408120610f1091611129565b80610f1a81611741565b915050610ef2565b50600255565b610f30611061565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33610d1a565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610fc8577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610ff4576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061101257662386f26fc10000830492506010015b6305f5e100831061102a576305f5e100830492506008015b612710831061103e57612710830492506004015b60648310611050576064830492506002015b600a8310610a125760010192915050565b60005462010000900460ff16610ab0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610415565b828054828255906000526020600020908101928215611119579160200282015b82811115611119578251829061110990826117c7565b50916020019190600101906110f3565b5061112592915061114a565b5090565b5080546000825590600052602060002090810190611147919061114a565b50565b8082111561112557600061115e8282611167565b5060010161114a565b508054611173906116ee565b6000825580601f10611183575050565b601f01602090049060005260206000209081019061114791905b80821115611125576000815560010161119d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611227576112276111b1565b604052919050565b600082601f83011261124057600080fd5b813567ffffffffffffffff81111561125a5761125a6111b1565b61128b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016111e0565b8181528460208386010111156112a057600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156112d7576112d76111b1565b5060051b60200190565b803563ffffffff811681146112f557600080fd5b919050565b60008060006060848603121561130f57600080fd5b833567ffffffffffffffff8082111561132757600080fd5b6113338783880161122f565b945060209150818601358181111561134a57600080fd5b8601601f8101881361135b57600080fd5b803561136e611369826112bd565b6111e0565b81815260059190911b8201840190848101908a83111561138d57600080fd5b8584015b838110156113c5578035868111156113a95760008081fd5b6113b78d898389010161122f565b845250918601918601611391565b508097505050505050506113db604085016112e1565b90509250925092565b803573ffffffffffffffffffffffffffffffffffffffff811681146112f557600080fd5b6000806040838503121561141b57600080fd5b611424836113e4565b9150611432602084016113e4565b90509250929050565b60005b8381101561145657818101518382015260200161143e565b50506000910152565b6000815180845261147781602086016020860161143b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006114bc602083018461145f565b9392505050565b600080604083850312156114d657600080fd5b50508035926020909101359150565b6000602082840312156114f757600080fd5b5035919050565b60008060006060848603121561151357600080fd5b8335925060208085013567ffffffffffffffff8082111561153357600080fd5b818701915087601f83011261154757600080fd5b8135611555611369826112bd565b81815260059190911b8301840190848101908a83111561157457600080fd5b938501935b8285101561159257843582529385019390850190611579565b9650505060408701359250808311156115aa57600080fd5b50506115b88682870161122f565b9150509250925092565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006115f1606083018561145f565b905063ffffffff83166040830152949350505050565b6000845161161981846020890161143b565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611655816001850160208a0161143b565b6001920191820152835161167081600284016020880161143b565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610a1257610a1261167d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168061170257607f821691505b60208210810361173b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117725761177261167d565b5060010190565b601f82111561086d57600081815260208120601f850160051c810160208610156117a05750805b601f850160051c820191505b818110156117bf578281556001016117ac565b505050505050565b815167ffffffffffffffff8111156117e1576117e16111b1565b6117f5816117ef84546116ee565b84611779565b602080601f83116001811461184857600084156118125750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556117bf565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561189557888601518255948401946001909101908401611876565b50858210156118d157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c6343000810000a",
}

// L1SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1SequencerMetaData.ABI instead.
var L1SequencerABI = L1SequencerMetaData.ABI

// L1SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1SequencerMetaData.Bin instead.
var L1SequencerBin = L1SequencerMetaData.Bin

// DeployL1Sequencer deploys a new Ethereum contract, binding an instance of L1Sequencer to it.
func DeployL1Sequencer(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address) (common.Address, *types.Transaction, *L1Sequencer, error) {
	parsed, err := L1SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1SequencerBin), backend, _messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1Sequencer{L1SequencerCaller: L1SequencerCaller{contract: contract}, L1SequencerTransactor: L1SequencerTransactor{contract: contract}, L1SequencerFilterer: L1SequencerFilterer{contract: contract}}, nil
}

// L1Sequencer is an auto generated Go binding around an Ethereum contract.
type L1Sequencer struct {
	L1SequencerCaller     // Read-only binding to the contract
	L1SequencerTransactor // Write-only binding to the contract
	L1SequencerFilterer   // Log filterer for contract events
}

// L1SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1SequencerSession struct {
	Contract     *L1Sequencer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1SequencerCallerSession struct {
	Contract *L1SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// L1SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1SequencerTransactorSession struct {
	Contract     *L1SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// L1SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1SequencerRaw struct {
	Contract *L1Sequencer // Generic contract binding to access the raw methods on
}

// L1SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1SequencerCallerRaw struct {
	Contract *L1SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// L1SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1SequencerTransactorRaw struct {
	Contract *L1SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1Sequencer creates a new instance of L1Sequencer, bound to a specific deployed contract.
func NewL1Sequencer(address common.Address, backend bind.ContractBackend) (*L1Sequencer, error) {
	contract, err := bindL1Sequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1Sequencer{L1SequencerCaller: L1SequencerCaller{contract: contract}, L1SequencerTransactor: L1SequencerTransactor{contract: contract}, L1SequencerFilterer: L1SequencerFilterer{contract: contract}}, nil
}

// NewL1SequencerCaller creates a new read-only instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerCaller(address common.Address, caller bind.ContractCaller) (*L1SequencerCaller, error) {
	contract, err := bindL1Sequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1SequencerCaller{contract: contract}, nil
}

// NewL1SequencerTransactor creates a new write-only instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1SequencerTransactor, error) {
	contract, err := bindL1Sequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1SequencerTransactor{contract: contract}, nil
}

// NewL1SequencerFilterer creates a new log filterer instance of L1Sequencer, bound to a specific deployed contract.
func NewL1SequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1SequencerFilterer, error) {
	contract, err := bindL1Sequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1SequencerFilterer{contract: contract}, nil
}

// bindL1Sequencer binds a generic wrapper to an already deployed contract.
func bindL1Sequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1SequencerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Sequencer *L1SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Sequencer.Contract.L1SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Sequencer *L1SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.Contract.L1SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Sequencer *L1SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Sequencer.Contract.L1SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Sequencer *L1SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Sequencer *L1SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Sequencer *L1SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Sequencer.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Sequencer *L1SequencerCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Sequencer *L1SequencerSession) MESSENGER() (common.Address, error) {
	return _L1Sequencer.Contract.MESSENGER(&_L1Sequencer.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) MESSENGER() (common.Address, error) {
	return _L1Sequencer.Contract.MESSENGER(&_L1Sequencer.CallOpts)
}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L1Sequencer *L1SequencerCaller) OTHERSEQUENCER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "OTHER_SEQUENCER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L1Sequencer *L1SequencerSession) OTHERSEQUENCER() (common.Address, error) {
	return _L1Sequencer.Contract.OTHERSEQUENCER(&_L1Sequencer.CallOpts)
}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) OTHERSEQUENCER() (common.Address, error) {
	return _L1Sequencer.Contract.OTHERSEQUENCER(&_L1Sequencer.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerCaller) CurrentVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "currentVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerSession) CurrentVersion() (*big.Int, error) {
	return _L1Sequencer.Contract.CurrentVersion(&_L1Sequencer.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerCallerSession) CurrentVersion() (*big.Int, error) {
	return _L1Sequencer.Contract.CurrentVersion(&_L1Sequencer.CallOpts)
}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x56002467.
//
// Solidity: function getSequencerBLSKeys(uint256 version, uint256 index) view returns(bytes)
func (_L1Sequencer *L1SequencerCaller) GetSequencerBLSKeys(opts *bind.CallOpts, version *big.Int, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerBLSKeys", version, index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x56002467.
//
// Solidity: function getSequencerBLSKeys(uint256 version, uint256 index) view returns(bytes)
func (_L1Sequencer *L1SequencerSession) GetSequencerBLSKeys(version *big.Int, index *big.Int) ([]byte, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeys(&_L1Sequencer.CallOpts, version, index)
}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x56002467.
//
// Solidity: function getSequencerBLSKeys(uint256 version, uint256 index) view returns(bytes)
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerBLSKeys(version *big.Int, index *big.Int) ([]byte, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeys(&_L1Sequencer.CallOpts, version, index)
}

// GetSequencerBLSKeysLength is a free data retrieval call binding the contract method 0xb00f376d.
//
// Solidity: function getSequencerBLSKeysLength(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerCaller) GetSequencerBLSKeysLength(opts *bind.CallOpts, version *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerBLSKeysLength", version)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequencerBLSKeysLength is a free data retrieval call binding the contract method 0xb00f376d.
//
// Solidity: function getSequencerBLSKeysLength(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerSession) GetSequencerBLSKeysLength(version *big.Int) (*big.Int, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeysLength(&_L1Sequencer.CallOpts, version)
}

// GetSequencerBLSKeysLength is a free data retrieval call binding the contract method 0xb00f376d.
//
// Solidity: function getSequencerBLSKeysLength(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerBLSKeysLength(version *big.Int) (*big.Int, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeysLength(&_L1Sequencer.CallOpts, version)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Sequencer *L1SequencerCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Sequencer *L1SequencerSession) Messenger() (common.Address, error) {
	return _L1Sequencer.Contract.Messenger(&_L1Sequencer.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) Messenger() (common.Address, error) {
	return _L1Sequencer.Contract.Messenger(&_L1Sequencer.CallOpts)
}

// NewestVersion is a free data retrieval call binding the contract method 0x73452a92.
//
// Solidity: function newestVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerCaller) NewestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "newestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewestVersion is a free data retrieval call binding the contract method 0x73452a92.
//
// Solidity: function newestVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerSession) NewestVersion() (*big.Int, error) {
	return _L1Sequencer.Contract.NewestVersion(&_L1Sequencer.CallOpts)
}

// NewestVersion is a free data retrieval call binding the contract method 0x73452a92.
//
// Solidity: function newestVersion() view returns(uint256)
func (_L1Sequencer *L1SequencerCallerSession) NewestVersion() (*big.Int, error) {
	return _L1Sequencer.Contract.NewestVersion(&_L1Sequencer.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1Sequencer *L1SequencerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1Sequencer *L1SequencerSession) Paused() (bool, error) {
	return _L1Sequencer.Contract.Paused(&_L1Sequencer.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1Sequencer *L1SequencerCallerSession) Paused() (bool, error) {
	return _L1Sequencer.Contract.Paused(&_L1Sequencer.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Sequencer *L1SequencerCaller) RollupContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "rollupContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Sequencer *L1SequencerSession) RollupContract() (common.Address, error) {
	return _L1Sequencer.Contract.RollupContract(&_L1Sequencer.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) RollupContract() (common.Address, error) {
	return _L1Sequencer.Contract.RollupContract(&_L1Sequencer.CallOpts)
}

// SequencerBLSKeys is a free data retrieval call binding the contract method 0x96091ba1.
//
// Solidity: function sequencerBLSKeys(uint256 , uint256 ) view returns(bytes)
func (_L1Sequencer *L1SequencerCaller) SequencerBLSKeys(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "sequencerBLSKeys", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SequencerBLSKeys is a free data retrieval call binding the contract method 0x96091ba1.
//
// Solidity: function sequencerBLSKeys(uint256 , uint256 ) view returns(bytes)
func (_L1Sequencer *L1SequencerSession) SequencerBLSKeys(arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	return _L1Sequencer.Contract.SequencerBLSKeys(&_L1Sequencer.CallOpts, arg0, arg1)
}

// SequencerBLSKeys is a free data retrieval call binding the contract method 0x96091ba1.
//
// Solidity: function sequencerBLSKeys(uint256 , uint256 ) view returns(bytes)
func (_L1Sequencer *L1SequencerCallerSession) SequencerBLSKeys(arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	return _L1Sequencer.Contract.SequencerBLSKeys(&_L1Sequencer.CallOpts, arg0, arg1)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_L1Sequencer *L1SequencerCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_L1Sequencer *L1SequencerSession) StakingContract() (common.Address, error) {
	return _L1Sequencer.Contract.StakingContract(&_L1Sequencer.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) StakingContract() (common.Address, error) {
	return _L1Sequencer.Contract.StakingContract(&_L1Sequencer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1Sequencer *L1SequencerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1Sequencer *L1SequencerSession) Version() (string, error) {
	return _L1Sequencer.Contract.Version(&_L1Sequencer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1Sequencer *L1SequencerCallerSession) Version() (string, error) {
	return _L1Sequencer.Contract.Version(&_L1Sequencer.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _stakingContract, address _rollupContract) returns()
func (_L1Sequencer *L1SequencerTransactor) Initialize(opts *bind.TransactOpts, _stakingContract common.Address, _rollupContract common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "initialize", _stakingContract, _rollupContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _stakingContract, address _rollupContract) returns()
func (_L1Sequencer *L1SequencerSession) Initialize(_stakingContract common.Address, _rollupContract common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Initialize(&_L1Sequencer.TransactOpts, _stakingContract, _rollupContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _stakingContract, address _rollupContract) returns()
func (_L1Sequencer *L1SequencerTransactorSession) Initialize(_stakingContract common.Address, _rollupContract common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Initialize(&_L1Sequencer.TransactOpts, _stakingContract, _rollupContract)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1Sequencer *L1SequencerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1Sequencer *L1SequencerSession) Pause() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Pause(&_L1Sequencer.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1Sequencer *L1SequencerTransactorSession) Pause() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Pause(&_L1Sequencer.TransactOpts)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x3135cb9a.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _minGasLimit) returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateAndSendSequencerSet(opts *bind.TransactOpts, _sequencerBytes []byte, _sequencerBLSKeys [][]byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateAndSendSequencerSet", _sequencerBytes, _sequencerBLSKeys, _minGasLimit)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x3135cb9a.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _minGasLimit) returns()
func (_L1Sequencer *L1SequencerSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerBLSKeys [][]byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerBLSKeys, _minGasLimit)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x3135cb9a.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _minGasLimit) returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerBLSKeys [][]byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerBLSKeys, _minGasLimit)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns()
func (_L1Sequencer *L1SequencerTransactor) VerifySignature(opts *bind.TransactOpts, version *big.Int, indexs []*big.Int, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "verifySignature", version, indexs, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns()
func (_L1Sequencer *L1SequencerSession) VerifySignature(version *big.Int, indexs []*big.Int, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, indexs, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns()
func (_L1Sequencer *L1SequencerTransactorSession) VerifySignature(version *big.Int, indexs []*big.Int, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, indexs, signature)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1Sequencer *L1SequencerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1Sequencer *L1SequencerSession) Receive() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Receive(&_L1Sequencer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1Sequencer *L1SequencerTransactorSession) Receive() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Receive(&_L1Sequencer.TransactOpts)
}

// L1SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1Sequencer contract.
type L1SequencerInitializedIterator struct {
	Event *L1SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *L1SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerInitialized)
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
		it.Event = new(L1SequencerInitialized)
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
func (it *L1SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerInitialized represents a Initialized event raised by the L1Sequencer contract.
type L1SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Sequencer *L1SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1SequencerInitializedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1SequencerInitializedIterator{contract: _L1Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Sequencer *L1SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerInitialized)
				if err := _L1Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1Sequencer *L1SequencerFilterer) ParseInitialized(log types.Log) (*L1SequencerInitialized, error) {
	event := new(L1SequencerInitialized)
	if err := _L1Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1SequencerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1Sequencer contract.
type L1SequencerPausedIterator struct {
	Event *L1SequencerPaused // Event containing the contract specifics and raw log

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
func (it *L1SequencerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerPaused)
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
		it.Event = new(L1SequencerPaused)
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
func (it *L1SequencerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerPaused represents a Paused event raised by the L1Sequencer contract.
type L1SequencerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1Sequencer *L1SequencerFilterer) FilterPaused(opts *bind.FilterOpts) (*L1SequencerPausedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1SequencerPausedIterator{contract: _L1Sequencer.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1Sequencer *L1SequencerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1SequencerPaused) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerPaused)
				if err := _L1Sequencer.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1Sequencer *L1SequencerFilterer) ParsePaused(log types.Log) (*L1SequencerPaused, error) {
	event := new(L1SequencerPaused)
	if err := _L1Sequencer.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1SequencerSequencerConfirmedIterator is returned from FilterSequencerConfirmed and is used to iterate over the raw logs and unpacked data for SequencerConfirmed events raised by the L1Sequencer contract.
type L1SequencerSequencerConfirmedIterator struct {
	Event *L1SequencerSequencerConfirmed // Event containing the contract specifics and raw log

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
func (it *L1SequencerSequencerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerSequencerConfirmed)
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
		it.Event = new(L1SequencerSequencerConfirmed)
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
func (it *L1SequencerSequencerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerSequencerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerSequencerConfirmed represents a SequencerConfirmed event raised by the L1Sequencer contract.
type L1SequencerSequencerConfirmed struct {
	Sequencers []common.Address
	Version    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequencerConfirmed is a free log retrieval operation binding the contract event 0x50a525bf6ba5f5d52f74ae26ad49967b0c29d7f16b6bb634250b74dd792e8b05.
//
// Solidity: event SequencerConfirmed(address[] sequencers, uint256 version)
func (_L1Sequencer *L1SequencerFilterer) FilterSequencerConfirmed(opts *bind.FilterOpts) (*L1SequencerSequencerConfirmedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "SequencerConfirmed")
	if err != nil {
		return nil, err
	}
	return &L1SequencerSequencerConfirmedIterator{contract: _L1Sequencer.contract, event: "SequencerConfirmed", logs: logs, sub: sub}, nil
}

// WatchSequencerConfirmed is a free log subscription operation binding the contract event 0x50a525bf6ba5f5d52f74ae26ad49967b0c29d7f16b6bb634250b74dd792e8b05.
//
// Solidity: event SequencerConfirmed(address[] sequencers, uint256 version)
func (_L1Sequencer *L1SequencerFilterer) WatchSequencerConfirmed(opts *bind.WatchOpts, sink chan<- *L1SequencerSequencerConfirmed) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "SequencerConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerSequencerConfirmed)
				if err := _L1Sequencer.contract.UnpackLog(event, "SequencerConfirmed", log); err != nil {
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

// ParseSequencerConfirmed is a log parse operation binding the contract event 0x50a525bf6ba5f5d52f74ae26ad49967b0c29d7f16b6bb634250b74dd792e8b05.
//
// Solidity: event SequencerConfirmed(address[] sequencers, uint256 version)
func (_L1Sequencer *L1SequencerFilterer) ParseSequencerConfirmed(log types.Log) (*L1SequencerSequencerConfirmed, error) {
	event := new(L1SequencerSequencerConfirmed)
	if err := _L1Sequencer.contract.UnpackLog(event, "SequencerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1SequencerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1Sequencer contract.
type L1SequencerUnpausedIterator struct {
	Event *L1SequencerUnpaused // Event containing the contract specifics and raw log

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
func (it *L1SequencerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerUnpaused)
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
		it.Event = new(L1SequencerUnpaused)
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
func (it *L1SequencerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerUnpaused represents a Unpaused event raised by the L1Sequencer contract.
type L1SequencerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1Sequencer *L1SequencerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1SequencerUnpausedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1SequencerUnpausedIterator{contract: _L1Sequencer.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1Sequencer *L1SequencerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1SequencerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerUnpaused)
				if err := _L1Sequencer.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1Sequencer *L1SequencerFilterer) ParseUnpaused(log types.Log) (*L1SequencerUnpaused, error) {
	event := new(L1SequencerUnpaused)
	if err := _L1Sequencer.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
