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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerBLSKeysLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sequencerBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sequencerBLSKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"updateAndSendSequencerSet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexs\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101206040525f6002555f60035534801562000019575f80fd5b5060405162001a8238038062001a828339810160408190526200003c9162000146565b6001600160a01b03811660805273530000000000000000000000000000000000000360a052600160c0525f60e0819052610100819052805462ff000019169055620000866200008d565b5062000175565b62000097620000ed565b5f805462ff00001916620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620000d03390565b6040516001600160a01b03909116815260200160405180910390a1565b620000ff5f5462010000900460ff1690565b15620001445760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b5f6020828403121562000157575f80fd5b81516001600160a01b03811681146200016e575f80fd5b9392505050565b60805160a05160c05160e051610100516118ba620001c85f395f6106cd01525f6106a401525f61067b01525f81816103490152610acb01525f818161010f015281816102270152610a9c01526118ba5ff3fe6080604052600436106100f2575f3560e01c806396091ba111610087578063bfa02ba911610057578063bfa02ba9146102bb578063e4821eb4146102e7578063ee99205c14610306578063f81e02a714610338575f80fd5b806396091ba11461024957806399a28076146102685780639d888e861461027b578063b00f376d14610290575f80fd5b80635c975abb116100c25780635c975abb146101b857806373452a92146101df5780638456cb5914610202578063927ede2d14610216575f80fd5b80633cb747bf14610101578063485cc9551461015957806354fd4d50146101785780635600246714610199575f80fd5b366100fd575f80fd5b005b5f80fd5b34801561010c575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610164575f80fd5b506100fb610173366004611187565b61036b565b348015610183575f80fd5b5061018c610674565b6040516101509190611223565b3480156101a4575f80fd5b5061018c6101b336600461123c565b610717565b3480156101c3575f80fd5b505f5462010000900460ff166040519015158152602001610150565b3480156101ea575f80fd5b506101f460035481565b604051908152602001610150565b34801561020d575f80fd5b506100fb610813565b348015610221575f80fd5b5061012f7f000000000000000000000000000000000000000000000000000000000000000081565b348015610254575f80fd5b5061018c61026336600461123c565b6108ac565b6100fb610276366004611398565b61095d565b348015610286575f80fd5b506101f460025481565b34801561029b575f80fd5b506101f46102aa36600461148a565b5f9081526004602052604090205490565b3480156102c6575f80fd5b5060015461012f9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f2575f80fd5b506100fb6103013660046114a1565b610b2f565b348015610311575f80fd5b505f5461012f906301000000900473ffffffffffffffffffffffffffffffffffffffff1681565b348015610343575f80fd5b5061012f7f000000000000000000000000000000000000000000000000000000000000000081565b5f54610100900460ff161580801561038957505f54600160ff909116105b806103a25750303b1580156103a257505f5460ff166001145b610433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561048f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff831661050c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f696e76616c6964207374616b696e6720636f6e74726163740000000000000000604482015260640161042a565b73ffffffffffffffffffffffffffffffffffffffff8216610589576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e7472616374000000000000000000604482015260640161042a565b5f80547fffffffffffffffffff0000000000000000000000000000000000000000ffffff16630100000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001691841691909117905561060d610bc1565b801561066f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b606061069f7f0000000000000000000000000000000000000000000000000000000000000000610c47565b6106c87f0000000000000000000000000000000000000000000000000000000000000000610c47565b6106f17f0000000000000000000000000000000000000000000000000000000000000000610c47565b6040516020016107039392919061155e565b604051602081830303815290604052905090565b5f82815260046020526040902054606090801580159061074157508061073e846001611600565b11155b156107fc575f84815260046020526040902080548490811061076557610765611613565b905f5260205f2001805461077890611640565b80601f01602080910402602001604051908101604052809291908181526020018280546107a490611640565b80156107ef5780601f106107c6576101008083540402835291602001916107ef565b820191905f5260205f20905b8154815290600101906020018083116107d257829003601f168201915b505050505091505061080d565b505060408051602081019091525f81525b92915050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff16331461089a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e74726163740000000000000000000000604482015260640161042a565b6108a2610d03565b6108aa610bc1565b565b6004602052815f5260405f2081815481106108c5575f80fd5b905f5260205f20015f915091505080546108de90611640565b80601f016020809104026020016040519081016040528092919081815260200182805461090a90611640565b80156109555780601f1061092c57610100808354040283529160200191610955565b820191905f5260205f20905b81548152906001019060200180831161093857829003601f168201915b505050505081565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146109e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e74726163740000000000000000000000604482015260640161042a565b6109ed83610d75565b5f5462010000900460ff1615610a5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e706175736564000000000000604482015260640161042a565b6040517f5f7b157700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635f7b1577903490610afb907f0000000000000000000000000000000000000000000000000000000000000000905f908a9089908990600401611691565b5f604051808303818588803b158015610b12575f80fd5b505af1158015610b24573d5f803e3d5ffd5b505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bb0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e7472616374000000000000000000000000604482015260640161042a565b610bb8610d03565b61066f83610e31565b610bc9610d03565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c1d3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60605f610c5383610ee3565b60010190505f8167ffffffffffffffff811115610c7257610c7261125c565b6040519080825280601f01601f191660200182016040528015610c9c576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610ca657509392505050565b5f5462010000900460ff16156108aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161042a565b6003545f03610d8657610d86610fc4565b5f5462010000900460ff1615610df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e706175736564000000000000604482015260640161042a565b60038054905f610e07836116e6565b90915550506003545f9081526004602090815260409091208251610e2d9284019061108b565b5050565b6002548110158015610e4557506003548111155b610eab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e76616c69642076657273696f6e0000000000000000000000000000000000604482015260640161042a565b60015b818111610edd575f818152600460205260408120610ecb916110df565b80610ed5816116e6565b915050610eae565b50600255565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610f2b577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610f57576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610f7557662386f26fc10000830492506010015b6305f5e1008310610f8d576305f5e100830492506008015b6127108310610fa157612710830492506004015b60648310610fb3576064830492506002015b600a831061080d5760010192915050565b610fcc61101a565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33610c1d565b5f5462010000900460ff166108aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161042a565b828054828255905f5260205f209081019282156110cf579160200282015b828111156110cf57825182906110bf9082611768565b50916020019190600101906110a9565b506110db9291506110fd565b5090565b5080545f8255905f5260205f20908101906110fa91906110fd565b50565b808211156110db575f6111108282611119565b506001016110fd565b50805461112590611640565b5f825580601f10611134575050565b601f0160209004905f5260205f20908101906110fa91905b808211156110db575f815560010161114c565b803573ffffffffffffffffffffffffffffffffffffffff81168114611182575f80fd5b919050565b5f8060408385031215611198575f80fd5b6111a18361115f565b91506111af6020840161115f565b90509250929050565b5f5b838110156111d25781810151838201526020016111ba565b50505f910152565b5f81518084526111f18160208601602086016111b8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f61123560208301846111da565b9392505050565b5f806040838503121561124d575f80fd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156112d0576112d061125c565b604052919050565b5f82601f8301126112e7575f80fd5b813567ffffffffffffffff8111156113015761130161125c565b61133260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611289565b818152846020838601011115611346575f80fd5b816020850160208301375f918101602001919091529392505050565b5f67ffffffffffffffff82111561137b5761137b61125c565b5060051b60200190565b803563ffffffff81168114611182575f80fd5b5f805f80608085870312156113ab575f80fd5b843567ffffffffffffffff808211156113c2575f80fd5b6113ce888389016112d8565b95506020915081870135818111156113e4575f80fd5b8701601f810189136113f4575f80fd5b803561140761140282611362565b611289565b81815260059190911b8201840190848101908b831115611425575f80fd5b8584015b8381101561145b5780358681111561143f575f80fd5b61144d8e89838901016112d8565b845250918601918601611429565b5080985050505050505061147160408601611385565b915061147f6060860161115f565b905092959194509250565b5f6020828403121561149a575f80fd5b5035919050565b5f805f606084860312156114b3575f80fd5b8335925060208085013567ffffffffffffffff808211156114d2575f80fd5b818701915087601f8301126114e5575f80fd5b81356114f361140282611362565b81815260059190911b8301840190848101908a831115611511575f80fd5b938501935b8285101561152f57843582529385019390850190611516565b965050506040870135925080831115611546575f80fd5b5050611554868287016112d8565b9150509250925092565b5f845161156f8184602089016111b8565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516115ab816001850160208a016111b8565b600192019182015283516115c68160028401602088016111b8565b0160020195945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561080d5761080d6115d3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061165457607f821691505b60208210810361168b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835286602084015260a060408401526116c660a08401876111da565b63ffffffff95909516606084015292909216608090910152509392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611716576117166115d3565b5060010190565b601f82111561066f57805f5260205f20601f840160051c810160208510156117425750805b601f840160051c820191505b81811015611761575f815560010161174e565b5050505050565b815167ffffffffffffffff8111156117825761178261125c565b611796816117908454611640565b8461171d565b602080601f8311600181146117e8575f84156117b25750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561187c565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561183457888601518255948401946001909101908401611815565b508582101561187057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220d19cc4bb7f302af9a3c763125bd20e6a0e301ceecb1172d88c14c2ac955dcaed64736f6c63430008180033",
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

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x99a28076.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateAndSendSequencerSet(opts *bind.TransactOpts, _sequencerBytes []byte, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateAndSendSequencerSet", _sequencerBytes, _sequencerBLSKeys, _gasLimit, _refundAddress)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x99a28076.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerBLSKeys, _gasLimit, _refundAddress)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0x99a28076.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerBLSKeys, _gasLimit, _refundAddress)
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
