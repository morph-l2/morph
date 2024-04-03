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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencersAddr\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sequencersBLS\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"sequencerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sequencerBytes\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_sequencerAddrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sequencerBLSKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"}],\"name\":\"updateAndSendSequencerSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040525f6002555f60035534801562000018575f80fd5b5060405162001bce38038062001bce8339810160408190526200003b9162000135565b6001600160a01b03811660805273530000000000000000000000000000000000000360a0525f805462ff000019169055620000756200007c565b5062000164565b62000086620000dc565b5f805462ff00001916620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620000bf3390565b6040516001600160a01b03909116815260200160405180910390a1565b620000ee5f5462010000900460ff1690565b15620001335760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b5f6020828403121562000146575f80fd5b81516001600160a01b03811681146200015d575f80fd5b9392505050565b60805160a051611a336200019b5f395f81816103370152610ccf01525f818161014f0152818161025c0152610ca20152611a335ff3fe608060405234801561000f575f80fd5b5060043610610149575f3560e01c80638a7b00ea116100c7578063b20d8e681161007d578063e9edab2c11610063578063e9edab2c146102f9578063ee99205c1461030c578063f81e02a714610332575f80fd5b8063b20d8e68146102c6578063bfa02ba9146102d9575f80fd5b806396091ba1116100ad57806396091ba11461027e5780639d888e861461029e578063a1e0ce80146102a7575f80fd5b80638a7b00ea14610237578063927ede2d14610257575f80fd5b80634df3c5a21161011c5780636d46e987116101025780636d46e9871461020557806373452a92146102185780638456cb591461022f575f80fd5b80634df3c5a2146101d65780635c975abb146101e9575f80fd5b80633cb747bf1461014d5780633f4ba83a14610199578063448d6249146101a3578063485cc955146101c3575b5f80fd5b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101a1610359565b005b6101b66101b1366004611219565b6103ef565b6040516101909190611280565b6101a16101d13660046112c1565b610465565b61016f6101e43660046112f2565b610769565b5f5462010000900460ff165b6040519015158152602001610190565b6101f5610213366004611312565b6107aa565b61022160035481565b604051908152602001610190565b6101a161083f565b61024a610245366004611219565b6108ce565b6040516101909190611402565b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b61029161028c3660046112f2565b6109b3565b6040516101909190611414565b61022160025481565b6102216102b5366004611219565b5f9081526005602052604090205490565b6101a16102d43660046115d5565b610a64565b60015461016f9073ffffffffffffffffffffffffffffffffffffffff1681565b6101f56103073660046116d4565b610d31565b5f5461016f906301000000900473ffffffffffffffffffffffffffffffffffffffff1681565b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146103e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064015b60405180910390fd5b6103ed610dcf565b565b5f8181526004602090815260409182902080548351818402810184019094528084526060939283018282801561045957602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161042e575b50505050509050919050565b5f54610100900460ff161580801561048357505f54600160ff909116105b8061049c5750303b15801561049c57505f5460ff166001145b610528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103dc565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610584575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610601576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f696e76616c6964207374616b696e6720636f6e7472616374000000000000000060448201526064016103dc565b73ffffffffffffffffffffffffffffffffffffffff821661067e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e747261637400000000000000000060448201526064016103dc565b5f80547fffffffffffffffffff0000000000000000000000000000000000000000ffffff16630100000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016918416919091179055610702610e4b565b8015610764575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6004602052815f5260405f208181548110610782575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b5f805b6002545f90815260046020526040902054811015610837578273ffffffffffffffffffffffffffffffffffffffff1660045f60025481526020019081526020015f2082815481106108005761080061173c565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff160361082f5750600192915050565b6001016107ad565b505f92915050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146108c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064016103dc565b6103ed610e4b565b606060055f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156109a8578382905f5260205f2001805461091d90611769565b80601f016020809104026020016040519081016040528092919081815260200182805461094990611769565b80156109945780601f1061096b57610100808354040283529160200191610994565b820191905f5260205f20905b81548152906001019060200180831161097757829003601f168201915b505050505081526020019060010190610900565b505050509050919050565b6005602052815f5260405f2081815481106109cc575f80fd5b905f5260205f20015f915091505080546109e590611769565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1190611769565b8015610a5c5780601f10610a3357610100808354040283529160200191610a5c565b820191905f5260205f20905b815481529060010190602001808311610a3f57829003601f168201915b505050505081565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff163314610aeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e7472616374000000000000000000000060448201526064016103dc565b600354158015610b2257505f805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec54155b15610bac575f8052600460209081528351610b62917f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec91908601906110a1565b505f8052600560209081528251610b9e917f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc9190850190611129565b50610ba7610dcf565b610d2b565b5f5462010000900460ff1615610c1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e70617573656400000000000060448201526064016103dc565b610c288383610ea7565b7f3426004122e24543c0ea4a940ddb1b8f7de59b484229ef718fa87e1eeaf62fd28383600354604051610c5d939291906117ba565b60405180910390a16040517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b2267a7b90610cfd907f0000000000000000000000000000000000000000000000000000000000000000905f90899087906004016117ef565b5f604051808303815f87803b158015610d14575f80fd5b505af1158015610d26573d5f803e3d5ffd5b505050505b50505050565b6001545f9073ffffffffffffffffffffffffffffffffffffffff163314610db4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e747261637400000000000000000000000060448201526064016103dc565b610dbc610efe565b610dc584610f70565b5060019392505050565b610dd7611030565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610e53610efe565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e213390565b60038054905f610eb68361183a565b90915550506003545f9081526004602090815260409091208351610edc928501906110a1565b506003545f908152600560209081526040909120825161076492840190611129565b5f5462010000900460ff16156103ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016103dc565b6002548110158015610f8457506003548111155b610fea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642073657175656e6365722076657273696f6e0000000000000060448201526064016103dc565b6002545b8181101561102a575f81815260046020526040812061100c91611179565b5f81815260056020526040812061102291611197565b600101610fee565b50600255565b5f5462010000900460ff166103ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016103dc565b828054828255905f5260205f20908101928215611119579160200282015b8281111561111957825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906110bf565b506111259291506111b2565b5090565b828054828255905f5260205f2090810192821561116d579160200282015b8281111561116d578251829061115d90826118e1565b5091602001919060010190611147565b506111259291506111c6565b5080545f8255905f5260205f209081019061119491906111b2565b50565b5080545f8255905f5260205f209081019061119491906111c6565b5b80821115611125575f81556001016111b3565b80821115611125575f6111d982826111e2565b506001016111c6565b5080546111ee90611769565b5f825580601f106111fd575050565b601f0160209004905f5260205f209081019061119491906111b2565b5f60208284031215611229575f80fd5b5035919050565b5f815180845260208085019450602084015f5b8381101561127557815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611243565b509495945050505050565b602081525f6112926020830184611230565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112bc575f80fd5b919050565b5f80604083850312156112d2575f80fd5b6112db83611299565b91506112e960208401611299565b90509250929050565b5f8060408385031215611303575f80fd5b50508035926020909101359150565b5f60208284031215611322575f80fd5b61129282611299565b5f81518084525f5b8181101561134f57602081850181015186830182015201611333565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f8282518085526020808601955060208260051b840101602086015f5b848110156113f5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030189526113e383835161132b565b988401989250908301906001016113a9565b5090979650505050505050565b602081525f611292602083018461138c565b602081525f611292602083018461132b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561149a5761149a611426565b604052919050565b5f82601f8301126114b1575f80fd5b813567ffffffffffffffff8111156114cb576114cb611426565b6114fc60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611453565b818152846020838601011115611510575f80fd5b816020850160208301375f918101602001919091529392505050565b5f67ffffffffffffffff82111561154557611545611426565b5060051b60200190565b5f82601f83011261155e575f80fd5b8135602061157361156e8361152c565b611453565b8083825260208201915060208460051b870101935086841115611594575f80fd5b602086015b848110156115b7576115aa81611299565b8352918301918301611599565b509695505050505050565b803563ffffffff811681146112bc575f80fd5b5f805f80608085870312156115e8575f80fd5b843567ffffffffffffffff808211156115ff575f80fd5b61160b888389016114a2565b9550602091508187013581811115611621575f80fd5b61162d89828a0161154f565b955050604087013581811115611641575f80fd5b8701601f81018913611651575f80fd5b803561165f61156e8261152c565b81815260059190911b8201840190848101908b83111561167d575f80fd5b8584015b838110156116b357803586811115611697575f80fd5b6116a58e89838901016114a2565b845250918601918601611681565b508097505050505050506116c9606086016115c2565b905092959194509250565b5f805f606084860312156116e6575f80fd5b83359250602084013567ffffffffffffffff80821115611704575f80fd5b6117108783880161154f565b93506040860135915080821115611725575f80fd5b50611732868287016114a2565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061177d57607f821691505b6020821081036117b4577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b606081525f6117cc6060830186611230565b82810360208401526117de818661138c565b915050826040830152949350505050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f611823608083018561132b565b905063ffffffff8316606083015295945050505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361188f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5060010190565b601f82111561076457805f5260205f20601f840160051c810160208510156118bb5750805b601f840160051c820191505b818110156118da575f81556001016118c7565b5050505050565b815167ffffffffffffffff8111156118fb576118fb611426565b61190f816119098454611769565b84611896565b602080601f831160018114611961575f841561192b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556119f5565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156119ad5788860151825594840194600190910190840161198e565b50858210156119e957878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220c98b79071fbe33223aaccca1211cd09541244caf1daa0b744b2d96264d687a5464736f6c63430008180033",
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

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_L1Sequencer *L1SequencerCaller) IsSequencer(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _L1Sequencer.contract.Call(opts, &out, "isSequencer", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_L1Sequencer *L1SequencerSession) IsSequencer(addr common.Address) (bool, error) {
	return _L1Sequencer.Contract.IsSequencer(&_L1Sequencer.CallOpts, addr)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address addr) view returns(bool)
func (_L1Sequencer *L1SequencerCallerSession) IsSequencer(addr common.Address) (bool, error) {
	return _L1Sequencer.Contract.IsSequencer(&_L1Sequencer.CallOpts, addr)
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

// VerifySignature is a paid mutator transaction binding the contract method 0xe9edab2c.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature) returns(bool)
func (_L1Sequencer *L1SequencerTransactor) VerifySignature(opts *bind.TransactOpts, version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "verifySignature", version, sequencers, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe9edab2c.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature) returns(bool)
func (_L1Sequencer *L1SequencerSession) VerifySignature(version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, sequencers, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe9edab2c.
//
// Solidity: function verifySignature(uint256 version, address[] sequencers, bytes signature) returns(bool)
func (_L1Sequencer *L1SequencerTransactorSession) VerifySignature(version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, sequencers, signature)
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
	SequencersAddr []common.Address
	SequencersBLS  [][]byte
	Version        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0x3426004122e24543c0ea4a940ddb1b8f7de59b484229ef718fa87e1eeaf62fd2.
//
// Solidity: event SequencerUpdated(address[] sequencersAddr, bytes[] sequencersBLS, uint256 version)
func (_L1Sequencer *L1SequencerFilterer) FilterSequencerUpdated(opts *bind.FilterOpts) (*L1SequencerSequencerUpdatedIterator, error) {

	logs, sub, err := _L1Sequencer.contract.FilterLogs(opts, "SequencerUpdated")
	if err != nil {
		return nil, err
	}
	return &L1SequencerSequencerUpdatedIterator{contract: _L1Sequencer.contract, event: "SequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0x3426004122e24543c0ea4a940ddb1b8f7de59b484229ef718fa87e1eeaf62fd2.
//
// Solidity: event SequencerUpdated(address[] sequencersAddr, bytes[] sequencersBLS, uint256 version)
func (_L1Sequencer *L1SequencerFilterer) WatchSequencerUpdated(opts *bind.WatchOpts, sink chan<- *L1SequencerSequencerUpdated) (event.Subscription, error) {

	logs, sub, err := _L1Sequencer.contract.WatchLogs(opts, "SequencerUpdated")
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0x3426004122e24543c0ea4a940ddb1b8f7de59b484229ef718fa87e1eeaf62fd2.
//
// Solidity: event SequencerUpdated(address[] sequencersAddr, bytes[] sequencersBLS, uint256 version)
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
