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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerBLSKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"sequencerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sequencerIndex\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sequencerBytes\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_sequencerAddrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sequencerBLSKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"updateAndSendSequencerSet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexs\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040525f6002555f60035534801562000018575f80fd5b5060405162001d0338038062001d038339810160408190526200003b9162000135565b6001600160a01b03811660805273530000000000000000000000000000000000000360a0525f805462ff000019169055620000756200007c565b5062000164565b62000086620000dc565b5f805462ff00001916620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620000bf3390565b6040516001600160a01b03909116815260200160405180910390a1565b620000ee5f5462010000900460ff1690565b15620001335760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640160405180910390fd5b565b5f6020828403121562000146575f80fd5b81516001600160a01b03811681146200015d575f80fd5b9392505050565b60805160a051611b686200019b5f395f818161041a0152610d2201525f818161015d015281816102eb0152610cf30152611b685ff3fe608060405260043610610140575f3560e01c80638a7b00ea116100bb578063bfa02ba911610071578063e73a6ba811610057578063e73a6ba8146103c4578063ee99205c146103d7578063f81e02a714610409575f80fd5b8063bfa02ba914610379578063e4821eb4146103a5575f80fd5b806396091ba1116100a157806396091ba11461030d5780639d888e8614610339578063a1e0ce801461034e575f80fd5b80638a7b00ea146102ae578063927ede2d146102da575f80fd5b80635c975abb116101105780636d46e987116100f65780636d46e9871461025857806373452a92146102775780638456cb591461029a575f80fd5b80635c975abb146102115780635fcd076814610239575f80fd5b80633cb747bf1461014f578063448d6249146101a7578063485cc955146101d35780634df3c5a2146101f2575f80fd5b3661014b575f80fd5b005b5f80fd5b34801561015a575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101b2575f80fd5b506101c66101c13660046111ef565b61043c565b60405161019e9190611206565b3480156101de575f80fd5b506101496101ed366004611287565b6104b2565b3480156101fd575f80fd5b5061017d61020c3660046112b8565b61076d565b34801561021c575f80fd5b505f5462010000900460ff165b604051901515815260200161019e565b348015610244575f80fd5b506101496102533660046113f6565b6107ae565b348015610263575f80fd5b50610229610272366004611458565b6108b9565b348015610282575f80fd5b5061028c60035481565b60405190815260200161019e565b3480156102a5575f80fd5b5061014961094e565b3480156102b9575f80fd5b506102cd6102c83660046111ef565b6109cd565b60405161019e91906114d9565b3480156102e5575f80fd5b5061017d7f000000000000000000000000000000000000000000000000000000000000000081565b348015610318575f80fd5b5061032c6103273660046112b8565b610ab2565b60405161019e9190611559565b348015610344575f80fd5b5061028c60025481565b348015610359575f80fd5b5061028c6103683660046111ef565b5f9081526005602052604090205490565b348015610384575f80fd5b5060015461017d9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103b0575f80fd5b506102296103bf3660046115f5565b610b63565b6101496103d23660046116d8565b610be7565b3480156103e2575f80fd5b505f5461017d906301000000900473ffffffffffffffffffffffffffffffffffffffff1681565b348015610414575f80fd5b5061017d7f000000000000000000000000000000000000000000000000000000000000000081565b5f818152600460209081526040918290208054835181840281018401909452808452606093928301828280156104a657602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161047b575b50505050509050919050565b5f54610100900460ff16158080156104d057505f54600160ff909116105b806104e95750303b1580156104e957505f5460ff166001145b6105605760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105bc575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff831661061f5760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207374616b696e6720636f6e747261637400000000000000006044820152606401610557565b73ffffffffffffffffffffffffffffffffffffffff82166106825760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e74726163740000000000000000006044820152606401610557565b5f80547fffffffffffffffffff0000000000000000000000000000000000000000ffffff16630100000073ffffffffffffffffffffffffffffffffffffffff8681169190910291909117909155600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016918416919091179055610706610d87565b8015610768575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6004602052815f5260405f208181548110610786575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146108155760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610557565b5f805460025482526004602081905260409283902092517f6e1a7a1f000000000000000000000000000000000000000000000000000000008152630100000090920473ffffffffffffffffffffffffffffffffffffffff1692636e1a7a1f92610886928991899189918991016117d4565b5f604051808303815f87803b15801561089d575f80fd5b505af11580156108af573d5f803e3d5ffd5b5050505050505050565b5f805b6002545f90815260046020526040902054811015610946578273ffffffffffffffffffffffffffffffffffffffff1660045f60025481526020019081526020015f20828154811061090f5761090f61189c565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff160361093e5750600192915050565b6001016108bc565b505f92915050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff1633146109bb5760405162461bcd60e51b815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e747261637400000000000000000000006044820152606401610557565b6109c3610e0d565b6109cb610d87565b565b606060055f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610aa7578382905f5260205f20018054610a1c906118c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610a48906118c9565b8015610a935780601f10610a6a57610100808354040283529160200191610a93565b820191905f5260205f20905b815481529060010190602001808311610a7657829003601f168201915b5050505050815260200190600101906109ff565b505050509050919050565b6005602052815f5260405f208181548110610acb575f80fd5b905f5260205f20015f91509150508054610ae4906118c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610b10906118c9565b8015610b5b5780601f10610b3257610100808354040283529160200191610b5b565b820191905f5260205f20905b815481529060010190602001808311610b3e57829003601f168201915b505050505081565b6001545f9073ffffffffffffffffffffffffffffffffffffffff163314610bcc5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610557565b610bd4610e0d565b610bdd84610e65565b5060019392505050565b5f546301000000900473ffffffffffffffffffffffffffffffffffffffff163314610c545760405162461bcd60e51b815260206004820152601560248201527f6f6e6c79207374616b696e6720636f6e747261637400000000000000000000006044820152606401610557565b610c5e8484610f0a565b5f5462010000900460ff1615610cb65760405162461bcd60e51b815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e7061757365640000000000006044820152606401610557565b6040517f5f7b157700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635f7b1577903490610d52907f0000000000000000000000000000000000000000000000000000000000000000905f908b908990899060040161191a565b5f604051808303818588803b158015610d69575f80fd5b505af1158015610d7b573d5f803e3d5ffd5b50505050505050505050565b610d8f610e0d565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16620100001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610de33390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b5f5462010000900460ff16156109cb5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610557565b6002548110158015610e7957506003548111155b610ec55760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642076657273696f6e00000000000000000000000000000000006044820152606401610557565b60015b81811015610f04575f818152600460205260408120610ee691611077565b5f818152600560205260408120610efc91611095565b600101610ec8565b50600255565b6003545f03610f1b57610f1b610fca565b5f5462010000900460ff1615610f735760405162461bcd60e51b815260206004820152601a60248201527f73656e64206d657373616765207768656e20756e7061757365640000000000006044820152606401610557565b60038054905f610f828361196f565b90915550506003545f9081526004602090815260409091208351610fa8928501906110b0565b506003545f908152600560209081526040909120825161076892840190611138565b610fd2611020565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33610de3565b5f5462010000900460ff166109cb5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610557565b5080545f8255905f5260205f20908101906110929190611188565b50565b5080545f8255905f5260205f2090810190611092919061119c565b828054828255905f5260205f20908101928215611128579160200282015b8281111561112857825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906110ce565b50611134929150611188565b5090565b828054828255905f5260205f2090810192821561117c579160200282015b8281111561117c578251829061116c9082611a16565b5091602001919060010190611156565b5061113492915061119c565b5b80821115611134575f8155600101611189565b80821115611134575f6111af82826111b8565b5060010161119c565b5080546111c4906118c9565b5f825580601f106111d3575050565b601f0160209004905f5260205f20908101906110929190611188565b5f602082840312156111ff575f80fd5b5035919050565b602080825282518282018190525f9190848201906040850190845b8181101561125357835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611221565b50909695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611282575f80fd5b919050565b5f8060408385031215611298575f80fd5b6112a18361125f565b91506112af6020840161125f565b90509250929050565b5f80604083850312156112c9575f80fd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561134c5761134c6112d8565b604052919050565b5f67ffffffffffffffff82111561136d5761136d6112d8565b5060051b60200190565b5f82601f830112611386575f80fd5b8135602061139b61139683611354565b611305565b8083825260208201915060208460051b8701019350868411156113bc575f80fd5b602086015b848110156113d857803583529183019183016113c1565b509695505050505050565b803563ffffffff81168114611282575f80fd5b5f805f8060808587031215611409575f80fd5b843567ffffffffffffffff81111561141f575f80fd5b61142b87828801611377565b94505061143a6020860161125f565b9250611448604086016113e3565b9396929550929360600135925050565b5f60208284031215611468575f80fd5b6114718261125f565b9392505050565b5f81518084525f5b8181101561149c57602081850181015186830182015201611480565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f60208083016020845280855180835260408601915060408160051b8701019250602087015f5b8281101561154c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261153a858351611478565b94509285019290850190600101611500565b5092979650505050505050565b602081525f6114716020830184611478565b5f82601f83011261157a575f80fd5b813567ffffffffffffffff811115611594576115946112d8565b6115c560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611305565b8181528460208386010111156115d9575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f60608486031215611607575f80fd5b83359250602084013567ffffffffffffffff80821115611625575f80fd5b61163187838801611377565b93506040860135915080821115611646575f80fd5b506116538682870161156b565b9150509250925092565b5f82601f83011261166c575f80fd5b8135602061167c61139683611354565b82815260059290921b8401810191818101908684111561169a575f80fd5b8286015b848110156113d857803567ffffffffffffffff8111156116bc575f80fd5b6116ca8986838b010161156b565b84525091830191830161169e565b5f805f805f60a086880312156116ec575f80fd5b853567ffffffffffffffff80821115611703575f80fd5b61170f89838a0161156b565b9650602091508188013581811115611725575f80fd5b8801601f81018a13611735575f80fd5b803561174361139682611354565b81815260059190911b8201840190848101908c831115611761575f80fd5b928501925b82841015611786576117778461125f565b82529285019290850190611766565b9850505050604088013591508082111561179e575f80fd5b506117ab8882890161165d565b9350506117ba606087016113e3565b91506117c86080870161125f565b90509295509295909350565b5f60a0820160a0835280885480835260c085019150895f5260209250825f205f5b8281101561182757815473ffffffffffffffffffffffffffffffffffffffff16845292840192600191820191016117f5565b505050838103828501528751808252888301918301905f5b8181101561185b5783518352928401929184019160010161183f565b505073ffffffffffffffffffffffffffffffffffffffff881660408601529250611883915050565b63ffffffff939093166060820152608001529392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c908216806118dd57607f821691505b602082108103611914577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835286602084015260a0604084015261194f60a0840187611478565b63ffffffff95909516606084015292909216608090910152509392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119c4577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5060010190565b601f82111561076857805f5260205f20601f840160051c810160208510156119f05750805b601f840160051c820191505b81811015611a0f575f81556001016119fc565b5050505050565b815167ffffffffffffffff811115611a3057611a306112d8565b611a4481611a3e84546118c9565b846119cb565b602080601f831160018114611a96575f8415611a605750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611b2a565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611ae257888601518255948401946001909101908401611ac3565b5085821015611b1e57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212208a79fed38cc8364435788359a4df7e9f78561793a43933ab839d8038d0d042e764736f6c63430008180033",
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

// Slash is a paid mutator transaction binding the contract method 0x5fcd0768.
//
// Solidity: function slash(uint256[] sequencerIndex, address challenger, uint32 _minGasLimit, uint256 _gasFee) returns()
func (_L1Sequencer *L1SequencerTransactor) Slash(opts *bind.TransactOpts, sequencerIndex []*big.Int, challenger common.Address, _minGasLimit uint32, _gasFee *big.Int) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "slash", sequencerIndex, challenger, _minGasLimit, _gasFee)
}

// Slash is a paid mutator transaction binding the contract method 0x5fcd0768.
//
// Solidity: function slash(uint256[] sequencerIndex, address challenger, uint32 _minGasLimit, uint256 _gasFee) returns()
func (_L1Sequencer *L1SequencerSession) Slash(sequencerIndex []*big.Int, challenger common.Address, _minGasLimit uint32, _gasFee *big.Int) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Slash(&_L1Sequencer.TransactOpts, sequencerIndex, challenger, _minGasLimit, _gasFee)
}

// Slash is a paid mutator transaction binding the contract method 0x5fcd0768.
//
// Solidity: function slash(uint256[] sequencerIndex, address challenger, uint32 _minGasLimit, uint256 _gasFee) returns()
func (_L1Sequencer *L1SequencerTransactorSession) Slash(sequencerIndex []*big.Int, challenger common.Address, _minGasLimit uint32, _gasFee *big.Int) (*types.Transaction, error) {
	return _L1Sequencer.Contract.Slash(&_L1Sequencer.TransactOpts, sequencerIndex, challenger, _minGasLimit, _gasFee)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xe73a6ba8.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerTransactor) UpdateAndSendSequencerSet(opts *bind.TransactOpts, _sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "updateAndSendSequencerSet", _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit, _refundAddress)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xe73a6ba8.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit, _refundAddress)
}

// UpdateAndSendSequencerSet is a paid mutator transaction binding the contract method 0xe73a6ba8.
//
// Solidity: function updateAndSendSequencerSet(bytes _sequencerBytes, address[] _sequencerAddrs, bytes[] _sequencerBLSKeys, uint32 _gasLimit, address _refundAddress) payable returns()
func (_L1Sequencer *L1SequencerTransactorSession) UpdateAndSendSequencerSet(_sequencerBytes []byte, _sequencerAddrs []common.Address, _sequencerBLSKeys [][]byte, _gasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1Sequencer.Contract.UpdateAndSendSequencerSet(&_L1Sequencer.TransactOpts, _sequencerBytes, _sequencerAddrs, _sequencerBLSKeys, _gasLimit, _refundAddress)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns(bool)
func (_L1Sequencer *L1SequencerTransactor) VerifySignature(opts *bind.TransactOpts, version *big.Int, indexs []*big.Int, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.contract.Transact(opts, "verifySignature", version, indexs, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns(bool)
func (_L1Sequencer *L1SequencerSession) VerifySignature(version *big.Int, indexs []*big.Int, signature []byte) (*types.Transaction, error) {
	return _L1Sequencer.Contract.VerifySignature(&_L1Sequencer.TransactOpts, version, indexs, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xe4821eb4.
//
// Solidity: function verifySignature(uint256 version, uint256[] indexs, bytes signature) returns(bool)
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
