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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencersAddr\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sequencersBLS\",\"type\":\"bytes[]\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"sequencerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sequencerBytes\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_sequencerAddrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sequencerBLSKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"}],\"name\":\"updateAndSendSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040525f6002555f60035534801562000018575f80fd5b5060405162001b2438038062001b248339810160408190526200003b9162000135565b6001600160a01b03811660805273530000000000000000000000000000000000000360a0525f805462ff000019169055620000756200007c565b5062000164565b62000086620000dc565b5f805462ff00001916620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620000bf3390565b6040516001600160a01b03909116815260200160405180910390a1565b620000ee5f5462010000900460ff1690565b15620001335760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b5f6020828403121562000146575f80fd5b81516001600160a01b03811681146200015d575f80fd5b9392505050565b60805160a0516119896200019b5f395f818161033b0152610c4101525f8181610177015281816102600152610c1401526119895ff3fe608060405234801561000f575f80fd5b5060043610610149575f3560e01c80638a7b00ea116100c7578063b20d8e681161007d578063e89ff54311610063578063e89ff543146102fd578063ee99205c14610310578063f81e02a714610336575f80fd5b8063b20d8e68146102ca578063bfa02ba9146102dd575f80fd5b806396091ba1116100ad57806396091ba1146102825780639d888e86146102a2578063a1e0ce80146102ab575f80fd5b80638a7b00ea1461023b578063927ede2d1461025b575f80fd5b8063485cc9551161011c5780635c975abb116101025780635c975abb1461020c57806373452a921461021c5780638456cb5914610233575f80fd5b8063485cc955146101e65780634df3c5a2146101f9575f80fd5b806326c9973d1461014d5780633cb747bf146101755780633f4ba83a146101bc578063448d6249146101c6575b5f80fd5b61016061015b366004611188565b61035d565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016c565b6101c4610458565b005b6101d96101d43660046111b0565b6104cf565b60405161016c9190611217565b6101c46101f4366004611230565b610545565b610197610207366004611261565b6107fb565b5f5462010000900460ff16610160565b61022560035481565b60405190815260200161016c565b6101c461083c565b61024e6102493660046111b0565b6108b1565b60405161016c9190611358565b6101977f000000000000000000000000000000000000000000000000000000000000000081565b610295610290366004611261565b610996565b60405161016c919061136a565b61022560025481565b6102256102b93660046111b0565b5f9081526005602052604090205490565b6101c46102d836600461152b565b610a47565b6001546101979073ffffffffffffffffffffffffffffffffffffffff1681565b61016061030b36600461162a565b610cdf565b5f54610197906301000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6101977f000000000000000000000000000000000000000000000000000000000000000081565b5f600254821015801561037257506003548211155b6103c35760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642076657273696f6e000000000000000000000000000000000060448201526064015b60405180910390fd5b5f5b5f8381526004602052604090205481101561044d578373ffffffffffffffffffffffffffffffffffffffff1660045f60025481526020019081526020015f2082815481106104155761041561169a565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1603610445576001915050610452565b6001016103c5565b505f90505b92915050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146104c55760405162461bcd60e51b815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064016103ba565b6104cd610d64565b565b5f8181526004602090815260409182902080548351818402810184019094528084526060939283018282801561053957602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161050e575b50505050509050919050565b5f54610100900460ff161580801561056357505f54600160ff909116105b8061057c5750303b15801561057c57505f5460ff166001145b6105ee5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103ba565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561064a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff83166106ad5760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207374616b696e6720636f6e7472616374000000000000000060448201526064016103ba565b73ffffffffffffffffffffffffffffffffffffffff82166107105760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e747261637400000000000000000060448201526064016103ba565b5f80547fffffffffffffffffff0000000000000000000000000000000000000000ffffff16630100000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016918416919091179055610794610de0565b80156107f6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6004602052815f5260405f208181548110610814575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146108a95760405162461bcd60e51b815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064016103ba565b6104cd610de0565b606060055f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b8282101561098b578382905f5260205f20018054610900906116c7565b80601f016020809104026020016040519081016040528092919081815260200182805461092c906116c7565b80156109775780601f1061094e57610100808354040283529160200191610977565b820191905f5260205f20905b81548152906001019060200180831161095a57829003601f168201915b5050505050815260200190600101906108e3565b505050509050919050565b6005602052815f5260405f2081815481106109af575f80fd5b905f5260205f20015f915091505080546109c8906116c7565b80601f01602080910402602001604051908101604052809291908181526020018280546109f4906116c7565b8015610a3f5780601f10610a1657610100808354040283529160200191610a3f565b820191905f5260205f20905b815481529060010190602001808311610a2257829003601f168201915b505050505081565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff163314610ab45760405162461bcd60e51b815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064016103ba565b600354158015610aeb57505f805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec54155b15610b75575f8052600460209081528351610b2b917f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec9190860190610fe8565b505f8052600560209081528251610b67917f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc9190850190611070565b50610b70610d64565b610c9d565b5f5462010000900460ff1615610bcd5760405162461bcd60e51b815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e70617573656400000000000060448201526064016103ba565b610bd78383610e3c565b6040517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b2267a7b90610c6f907f0000000000000000000000000000000000000000000000000000000000000000905f9089908790600401611718565b5f604051808303815f87803b158015610c86575f80fd5b505af1158015610c98573d5f803e3d5ffd5b505050505b6003547f6cc20213c4a4b36cd7fdbf68276746ce3e071b5be3b94c24d48fb4a8181ee4d48484604051610cd1929190611763565b60405180910390a250505050565b6001545f9073ffffffffffffffffffffffffffffffffffffffff163314610d485760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e747261637400000000000000000000000060448201526064016103ba565b610d50610e93565b610d5985610eeb565b506001949350505050565b610d6c610f91565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610de8610e93565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610db63390565b60038054905f610e4b83611790565b90915550506003545f9081526004602090815260409091208351610e7192850190610fe8565b506003545f90815260056020908152604090912082516107f692840190611070565b5f5462010000900460ff16156104cd5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016103ba565b6002548110158015610eff57506003548111155b610f4b5760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642073657175656e6365722076657273696f6e0000000000000060448201526064016103ba565b6002545b81811015610f8b575f818152600460205260408120610f6d916110c0565b5f818152600560205260408120610f83916110de565b600101610f4f565b50600255565b5f5462010000900460ff166104cd5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016103ba565b828054828255905f5260205f20908101928215611060579160200282015b8281111561106057825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611006565b5061106c9291506110f9565b5090565b828054828255905f5260205f209081019282156110b4579160200282015b828111156110b457825182906110a49082611837565b509160200191906001019061108e565b5061106c92915061110d565b5080545f8255905f5260205f20908101906110db91906110f9565b50565b5080545f8255905f5260205f20908101906110db919061110d565b5b8082111561106c575f81556001016110fa565b8082111561106c575f6111208282611129565b5060010161110d565b508054611135906116c7565b5f825580601f10611144575050565b601f0160209004905f5260205f20908101906110db91906110f9565b803573ffffffffffffffffffffffffffffffffffffffff81168114611183575f80fd5b919050565b5f8060408385031215611199575f80fd5b6111a283611160565b946020939093013593505050565b5f602082840312156111c0575f80fd5b5035919050565b5f815180845260208085019450602084015f5b8381101561120c57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016111da565b509495945050505050565b602081525f61122960208301846111c7565b9392505050565b5f8060408385031215611241575f80fd5b61124a83611160565b915061125860208401611160565b90509250929050565b5f8060408385031215611272575f80fd5b50508035926020909101359150565b5f81518084525f5b818110156112a557602081850181015186830182015201611289565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561134b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018952611339838351611281565b988401989250908301906001016112ff565b5090979650505050505050565b602081525f61122960208301846112e2565b602081525f6112296020830184611281565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113f0576113f061137c565b604052919050565b5f82601f830112611407575f80fd5b813567ffffffffffffffff8111156114215761142161137c565b61145260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016113a9565b818152846020838601011115611466575f80fd5b816020850160208301375f918101602001919091529392505050565b5f67ffffffffffffffff82111561149b5761149b61137c565b5060051b60200190565b5f82601f8301126114b4575f80fd5b813560206114c96114c483611482565b6113a9565b8083825260208201915060208460051b8701019350868411156114ea575f80fd5b602086015b8481101561150d5761150081611160565b83529183019183016114ef565b509695505050505050565b803563ffffffff81168114611183575f80fd5b5f805f806080858703121561153e575f80fd5b843567ffffffffffffffff80821115611555575f80fd5b611561888389016113f8565b9550602091508187013581811115611577575f80fd5b61158389828a016114a5565b955050604087013581811115611597575f80fd5b8701601f810189136115a7575f80fd5b80356115b56114c482611482565b81815260059190911b8201840190848101908b8311156115d3575f80fd5b8584015b83811015611609578035868111156115ed575f80fd5b6115fb8e89838901016113f8565b8452509186019186016115d7565b5080975050505050505061161f60608601611518565b905092959194509250565b5f805f806080858703121561163d575f80fd5b84359350602085013567ffffffffffffffff8082111561165b575f80fd5b611667888389016114a5565b9450604087013591508082111561167c575f80fd5b50611689878288016113f8565b949793965093946060013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c908216806116db57607f821691505b602082108103611712577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f61174c6080830185611281565b905063ffffffff8316606083015295945050505050565b604081525f61177560408301856111c7565b828103602084015261178781856112e2565b95945050505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117e5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5060010190565b601f8211156107f657805f5260205f20601f840160051c810160208510156118115750805b601f840160051c820191505b81811015611830575f815560010161181d565b5050505050565b815167ffffffffffffffff8111156118515761185161137c565b6118658161185f84546116c7565b846117ec565b602080601f8311600181146118b7575f84156118815750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561194b565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611903578886015182559484019460019091019084016118e4565b508582101561193f57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220e965c14eed2367b634eaced99f05d1e2177b686b27da4009c75134693e1c44ac64736f6c63430008180033",
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

// GetSequencerAddrs is a free data retrieval call binding the contract method 0x448d6249.
//
// Solidity: function getSequencerAddrs(uint256 version) view returns(address[])
func (_L1Sequencer *L1SequencerCaller) GetSequencerAddrs(opts *bind.CallOpts, version *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerAddrs", version)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencerAddrs is a free data retrieval call binding the contract method 0x448d6249.
//
// Solidity: function getSequencerAddrs(uint256 version) view returns(address[])
func (_L1Sequencer *L1SequencerSession) GetSequencerAddrs(version *big.Int) ([]common.Address, error) {
	return _L1Sequencer.Contract.GetSequencerAddrs(&_L1Sequencer.CallOpts, version)
}

// GetSequencerAddrs is a free data retrieval call binding the contract method 0x448d6249.
//
// Solidity: function getSequencerAddrs(uint256 version) view returns(address[])
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerAddrs(version *big.Int) ([]common.Address, error) {
	return _L1Sequencer.Contract.GetSequencerAddrs(&_L1Sequencer.CallOpts, version)
}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x8a7b00ea.
//
// Solidity: function getSequencerBLSKeys(uint256 version) view returns(bytes[])
func (_L1Sequencer *L1SequencerCaller) GetSequencerBLSKeys(opts *bind.CallOpts, version *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "getSequencerBLSKeys", version)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x8a7b00ea.
//
// Solidity: function getSequencerBLSKeys(uint256 version) view returns(bytes[])
func (_L1Sequencer *L1SequencerSession) GetSequencerBLSKeys(version *big.Int) ([][]byte, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeys(&_L1Sequencer.CallOpts, version)
}

// GetSequencerBLSKeys is a free data retrieval call binding the contract method 0x8a7b00ea.
//
// Solidity: function getSequencerBLSKeys(uint256 version) view returns(bytes[])
func (_L1Sequencer *L1SequencerCallerSession) GetSequencerBLSKeys(version *big.Int) ([][]byte, error) {
	return _L1Sequencer.Contract.GetSequencerBLSKeys(&_L1Sequencer.CallOpts, version)
}

// IsSequencer is a free data retrieval call binding the contract method 0x26c9973d.
//
// Solidity: function isSequencer(address addr, uint256 version) view returns(bool)
func (_L1Sequencer *L1SequencerCaller) IsSequencer(opts *bind.CallOpts, addr common.Address, version *big.Int) (bool, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "isSequencer", addr, version)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x26c9973d.
//
// Solidity: function isSequencer(address addr, uint256 version) view returns(bool)
func (_L1Sequencer *L1SequencerSession) IsSequencer(addr common.Address, version *big.Int) (bool, error) {
	return _L1Sequencer.Contract.IsSequencer(&_L1Sequencer.CallOpts, addr, version)
}

// IsSequencer is a free data retrieval call binding the contract method 0x26c9973d.
//
// Solidity: function isSequencer(address addr, uint256 version) view returns(bool)
func (_L1Sequencer *L1SequencerCallerSession) IsSequencer(addr common.Address, version *big.Int) (bool, error) {
	return _L1Sequencer.Contract.IsSequencer(&_L1Sequencer.CallOpts, addr, version)
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

// SequencerAddrs is a free data retrieval call binding the contract method 0x4df3c5a2.
//
// Solidity: function sequencerAddrs(uint256 , uint256 ) view returns(address)
func (_L1Sequencer *L1SequencerCaller) SequencerAddrs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "sequencerAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerAddrs is a free data retrieval call binding the contract method 0x4df3c5a2.
//
// Solidity: function sequencerAddrs(uint256 , uint256 ) view returns(address)
func (_L1Sequencer *L1SequencerSession) SequencerAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _L1Sequencer.Contract.SequencerAddrs(&_L1Sequencer.CallOpts, arg0, arg1)
}

// SequencerAddrs is a free data retrieval call binding the contract method 0x4df3c5a2.
//
// Solidity: function sequencerAddrs(uint256 , uint256 ) view returns(address)
func (_L1Sequencer *L1SequencerCallerSession) SequencerAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _L1Sequencer.Contract.SequencerAddrs(&_L1Sequencer.CallOpts, arg0, arg1)
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

// SequencerNum is a free data retrieval call binding the contract method 0xa1e0ce80.
//
// Solidity: function sequencerNum(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerCaller) SequencerNum(opts *bind.CallOpts, version *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "sequencerNum", version)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerNum is a free data retrieval call binding the contract method 0xa1e0ce80.
//
// Solidity: function sequencerNum(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerSession) SequencerNum(version *big.Int) (*big.Int, error) {
	return _L1Sequencer.Contract.SequencerNum(&_L1Sequencer.CallOpts, version)
}

// SequencerNum is a free data retrieval call binding the contract method 0xa1e0ce80.
//
// Solidity: function sequencerNum(uint256 version) view returns(uint256)
func (_L1Sequencer *L1SequencerCallerSession) SequencerNum(version *big.Int) (*big.Int, error) {
	return _L1Sequencer.Contract.SequencerNum(&_L1Sequencer.CallOpts, version)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1Sequencer *L1SequencerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1Sequencer *L1SequencerSession) Unpause() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Unpause(&_L1Sequencer.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1Sequencer *L1SequencerTransactorSession) Unpause() (*types.Transaction, error) {
	return _L1Sequencer.Contract.Unpause(&_L1Sequencer.TransactOpts)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xb20d8e68.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit) returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateAndSendSequencerSet(opts *bind.TransactOpts, _sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateAndSendSequencerSet", _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xb20d8e68.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit) returns()
func (_L1Sequencer *L1SequencerSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xb20d8e68.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit) returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe89ff543.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature, bytes32 batchHash) returns(bool)
func (_L1Sequencer *L1SequencerTransactor) VerifySignature(opts *bind.TransactOpts, version *big.Int, sequencers []common.Address, signature []byte, batchHash [32]byte) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "verifySignature", version, sequencers, signature, batchHash)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe89ff543.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature, bytes32 batchHash) returns(bool)
func (_L1Sequencer *L1SequencerSession) VerifySignature(version *big.Int, sequencers []common.Address, signature []byte, batchHash [32]byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, sequencers, signature, batchHash)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe89ff543.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature, bytes32 batchHash) returns(bool)
func (_L1Sequencer *L1SequencerTransactorSession) VerifySignature(version *big.Int, sequencers []common.Address, signature []byte, batchHash [32]byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, sequencers, signature, batchHash)
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

// L1SequencerSequencerUpdatedIterator is returned from FilterSequencerUpdated and is used to iterate over the raw logs and unpacked data for SequencerUpdated events raised by the L1Sequencer contract.
type L1SequencerSequencerUpdatedIterator struct {
	Event *L1SequencerSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *L1SequencerSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1SequencerSequencerUpdated)
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
		it.Event = new(L1SequencerSequencerUpdated)
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
func (it *L1SequencerSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1SequencerSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1SequencerSequencerUpdated represents a SequencerUpdated event raised by the L1Sequencer contract.
type L1SequencerSequencerUpdated struct {
	Version        *big.Int
	SequencersAddr []common.Address
	SequencersBLS  [][]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0x6cc20213c4a4b36cd7fdbf68276746ce3e071b5be3b94c24d48fb4a8181ee4d4.
//
// Solidity: event SequencerUpdated(uint256 indexed version, address[] sequencersAddr, bytes[] sequencersBLS)
func (_L1Sequencer *L1SequencerFilterer) FilterSequencerUpdated(opts *bind.FilterOpts, version []*big.Int) (*L1SequencerSequencerUpdatedIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "SequencerUpdated", versionRule)
	if err != nil {
		return nil, err
	}
	return &L1SequencerSequencerUpdatedIterator{contract: _L1Sequencer.contract, event: "SequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0x6cc20213c4a4b36cd7fdbf68276746ce3e071b5be3b94c24d48fb4a8181ee4d4.
//
// Solidity: event SequencerUpdated(uint256 indexed version, address[] sequencersAddr, bytes[] sequencersBLS)
func (_L1Sequencer *L1SequencerFilterer) WatchSequencerUpdated(opts *bind.WatchOpts, sink chan<- *L1SequencerSequencerUpdated, version []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "SequencerUpdated", versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1SequencerSequencerUpdated)
				if err := _L1Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0x6cc20213c4a4b36cd7fdbf68276746ce3e071b5be3b94c24d48fb4a8181ee4d4.
//
// Solidity: event SequencerUpdated(uint256 indexed version, address[] sequencersAddr, bytes[] sequencersBLS)
func (_L1Sequencer *L1SequencerFilterer) ParseSequencerUpdated(log types.Log) (*L1SequencerSequencerUpdated, error) {
	event := new(L1SequencerSequencerUpdated)
	if err := _L1Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
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
