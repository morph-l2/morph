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

// TypesSequencerInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesSequencerInfo struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}

// L2SequencerMetaData contains all meta data concerning the L2Sequencer contract.
var L2SequencerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_otherSequencer\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"L2_SUBMITTER_CONTRACT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_SEQUENCER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSequencer\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentVersionHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerAddresses\",\"inputs\":[{\"name\":\"previous\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerInfos\",\"inputs\":[{\"name\":\"previous\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTypes.SequencerInfo[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inSequencersSet\",\"inputs\":[{\"name\":\"previous\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_sequencers\",\"type\":\"tuple[]\",\"internalType\":\"structTypes.SequencerInfo[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preSequencerAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preSequencerInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preVersionHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencerAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencerIndex\",\"inputs\":[{\"name\":\"previous\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencerInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencersLen\",\"inputs\":[{\"name\":\"previous\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSequencers\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sequencers\",\"type\":\"tuple[]\",\"internalType\":\"structTypes.SequencerInfo[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerUpdated\",\"inputs\":[{\"name\":\"sequencers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610140604052600060015560006002556000600355600060045534801561002557600080fd5b5060405161204c38038061204c83398101604081905261004491610096565b6001608081905260a052600060c05273420000000000000000000000000000000000000760e0526001600160a01b031661010052735300000000000000000000000000000000000005610120526100c6565b6000602082840312156100a857600080fd5b81516001600160a01b03811681146100bf57600080fd5b9392505050565b60805160a05160c05160e0516101005161012051611f14610138600039600081816102010152610a6201526000818161037f01526108bd0152600081816101a0015281816102650152818161089301526108f4015260006105480152600061051f015260006104f60152611f146000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c8063aeaf9f41116100cd578063d1c55fe311610081578063dd967ee911610066578063dd967ee914610347578063e597c19e1461035a578063f81e02a71461037a57600080fd5b8063d1c55fe314610314578063d95864671461033e57600080fd5b8063be6c5d68116100b2578063be6c5d68146102d8578063c9406b1a146102f8578063cfd1eff31461030b57600080fd5b8063aeaf9f41146102a3578063b95cdb78146102b657600080fd5b80635942e7c711610124578063927ede2d11610109578063927ede2d146102605780639d888e8614610287578063ad01732f1461029057600080fd5b80635942e7c7146102385780637ad9e3ac1461024d57600080fd5b80634a3c980c116101555780634a3c980c146101e55780634bbf5252146101fc57806354fd4d501461022357600080fd5b8063342b6345146101715780633cb747bf1461019e575b600080fd5b61018461017f366004611599565b6103a1565b604080519283526020830191909152015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610195565b6101ee60045481565b604051908152602001610195565b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b61022b6104ef565b604051610195919061163e565b61024b610246366004611867565b610592565b005b61018461025b3660046118a4565b610854565b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b6101ee60015481565b61024b61029e3660046118bf565b61087b565b6101c06102b1366004611906565b610c9a565b6102c96102c4366004611906565b610cd1565b6040516101959392919061191f565b6102eb6102e63660046118a4565b610da7565b604051610195919061195d565b6102c9610306366004611906565b610fe0565b6101ee60035481565b610327610322366004611599565b610ff0565b604080519215158352602083019190915201610195565b6101ee60025481565b6101c0610355366004611906565b6110ef565b61036d6103683660046118a4565b6110ff565b6040516101959190611a10565b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b600080831561047d5760005b60065481101561041557600681815481106103ca576103ca611a6a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036104035760035490925090506104e8565b8061040d81611a99565b9150506103ad565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f73657175656e636572206e6f742065786973740000000000000000000000000060448201526064015b60405180910390fd5b60005b600554811015610415576005818154811061049d5761049d611a6a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036104d65760015490925090506104e8565b806104e081611a99565b915050610480565b9250929050565b606061051a7f00000000000000000000000000000000000000000000000000000000000000006111e3565b6105437f00000000000000000000000000000000000000000000000000000000000000006111e3565b61056c7f00000000000000000000000000000000000000000000000000000000000000006111e3565b60405160200161057e93929190611af8565b604051602081830303815290604052905090565b600054610100900460ff16158080156105b25750600054600160ff909116105b806105cc5750303b1580156105cc575060005460ff166001145b610658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610474565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156106b657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60005b82518110156107eb5760058382815181106106d6576106d6611a6a565b6020908102919091018101515182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055825160079084908390811061074e5761074e611a6a565b602090810291909101810151825460018082018555600094855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911781559181015192820192909255604082015160028201906107d59082611c10565b50505080806107e390611a99565b9150506106b9565b50801561085057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600080821561086c5750506006546003549092909150565b50506005546001549092909150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561099957507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561095d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109819190611d2a565b73ffffffffffffffffffffffffffffffffffffffff16145b610a25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53657175656e6365723a2066756e6374696f6e2063616e206f6e6c792062652060448201527f63616c6c65642066726f6d20746865206f746865722073657175656e636572006064820152608401610474565b6040517f16e2994a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906316e2994a90610a9890600590600401611d9d565b600060405180830381600087803b158015610ab257600080fd5b505af1158015610ac6573d6000803e3d6000fd5b505060015460035550610add905060086000611384565b610ae9600660006113a8565b60078054610af9916008916113c6565b5060058054610b0a9160069161147a565b506002546004556001829055610b2260076000611384565b610b2e600560006113a8565b4360025560005b8151811015610c67576005828281518110610b5257610b52611a6a565b6020908102919091018101515182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558151600790839083908110610bca57610bca611a6a565b602090810291909101810151825460018082018555600094855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091178155918101519282019290925560408201516002820190610c519082611c10565b5050508080610c5f90611a99565b915050610b35565b507f71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d6182796600583604051610847929190611db0565b60058181548110610caa57600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60088181548110610ce157600080fd5b600091825260209091206003909102018054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff9093169450909291610d2490611b6e565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5090611b6e565b8015610d9d5780601f10610d7257610100808354040283529160200191610d9d565b820191906000526020600020905b815481529060010190602001808311610d8057829003601f168201915b5050505050905083565b60608115610ecd576008805480602002602001604051908101604052809291908181526020016000905b82821015610ec25760008481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610e3190611b6e565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5d90611b6e565b8015610eaa5780601f10610e7f57610100808354040283529160200191610eaa565b820191906000526020600020905b815481529060010190602001808311610e8d57829003601f168201915b50505050508152505081526020019060010190610dd1565b505050509050919050565b6007805480602002602001604051908101604052809291908181526020016000905b82821015610ec25760008481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610f4f90611b6e565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7b90611b6e565b8015610fc85780601f10610f9d57610100808354040283529160200191610fc8565b820191906000526020600020905b815481529060010190602001808311610fab57829003601f168201915b50505050508152505081526020019060010190610eef565b60078181548110610ce157600080fd5b60008083156110735760005b600654811015611065576006818154811061101957611019611a6a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603611053575050600354600191506104e8565b8061105d81611a99565b915050610ffc565b5050600354600091506104e8565b60005b6005548110156110df576005818154811061109357611093611a6a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036110cd576001805492509250506104e8565b806110d781611a99565b915050611076565b5050600154600091509250929050565b60068181548110610caa57600080fd5b6060811561117657600680548060200260200160405190810160405280929190818152602001828054801561116a57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161113f575b50505050509050919050565b600580548060200260200160405190810160405280929190818152602001828054801561116a5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161113f5750505050509050919050565b606060006111f0836112a1565b600101905060008167ffffffffffffffff81111561121057611210611658565b6040519080825280601f01601f19166020018201604052801561123a576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461124457509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106112ea577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611316576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061133457662386f26fc10000830492506010015b6305f5e100831061134c576305f5e100830492506008015b612710831061136057612710830492506004015b60648310611372576064830492506002015b600a831061137e576001015b92915050565b50805460008255600302906000526020600020908101906113a591906114c6565b50565b50805460008255906000526020600020908101906113a59190611513565b82805482825590600052602060002090600302810192821561146a5760005260206000209160030282015b8281111561146a57825482547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560018084015490830155828260028082019061145890840182611dd2565b505050916003019190600301906113f1565b506114769291506114c6565b5090565b8280548282559060005260206000209081019282156114ba5760005260206000209182015b828111156114ba57825482559160010191906001019061149f565b50611476929150611513565b808211156114765780547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560006001820181905561150a6002830182611528565b506003016114c6565b5b808211156114765760008155600101611514565b50805461153490611b6e565b6000825580601f10611544575050565b601f0160209004906000526020600020908101906113a59190611513565b8035801515811461157257600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146113a557600080fd5b600080604083850312156115ac57600080fd5b6115b583611562565b915060208301356115c581611577565b809150509250929050565b60005b838110156115eb5781810151838201526020016115d3565b50506000910152565b6000815180845261160c8160208601602086016115d0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061165160208301846115f4565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156116aa576116aa611658565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116f7576116f7611658565b604052919050565b6000601f838184011261171157600080fd5b8235602067ffffffffffffffff8083111561172e5761172e611658565b8260051b61173d8382016116b0565b938452868101830193838101908986111561175757600080fd5b84890192505b8583101561185a578235848111156117755760008081fd5b890160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d0381018213156117ac5760008081fd5b6117b4611687565b888401356117c181611577565b81526040848101358a8301529284013592888411156117e05760008081fd5b83850194508e603f8601126117f757600093508384fd5b8985013593508884111561180d5761180d611658565b61181c8a848e870116016116b0565b92508383528e818587010111156118335760008081fd5b838186018b85013760009383018a019390935291820152835250918401919084019061175d565b9998505050505050505050565b60006020828403121561187957600080fd5b813567ffffffffffffffff81111561189057600080fd5b61189c848285016116ff565b949350505050565b6000602082840312156118b657600080fd5b61165182611562565b600080604083850312156118d257600080fd5b82359150602083013567ffffffffffffffff8111156118f057600080fd5b6118fc858286016116ff565b9150509250929050565b60006020828403121561191857600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061195460608301846115f4565b95945050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611a02578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff168452878101518885015286015160608785018190526119ee818601836115f4565b968901969450505090860190600101611984565b509098975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611a5e57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611a2c565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611af1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60008451611b0a8184602089016115d0565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611b46816001850160208a016115d0565b60019201918201528351611b618160028401602088016115d0565b0160020195945050505050565b600181811c90821680611b8257607f821691505b602082108103611bbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115611c0b57600081815260208120601f850160051c81016020861015611be85750805b601f850160051c820191505b81811015611c0757828155600101611bf4565b5050505b505050565b815167ffffffffffffffff811115611c2a57611c2a611658565b611c3e81611c388454611b6e565b84611bc1565b602080601f831160018114611c915760008415611c5b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611c07565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611cde57888601518255948401946001909101908401611cbf565b5085821015611d1a57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600060208284031215611d3c57600080fd5b815161165181611577565b6000815480845260208085019450836000528060002060005b83811015611d9257815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101611d60565b509495945050505050565b6020815260006116516020830184611d47565b604081526000611dc36040830185611d47565b90508260208301529392505050565b818103611ddd575050565b611de78254611b6e565b67ffffffffffffffff811115611dff57611dff611658565b611e0d81611c388454611b6e565b6000601f821160018114611e5f5760008315611e295750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611f00565b6000858152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841690600086815260209020845b83811015611eb75782860154825560019586019590910190602001611e97565b5085831015611ef357818501547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b50505060018360011b0184555b505050505056fea164736f6c6343000810000a",
}

// L2SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2SequencerMetaData.ABI instead.
var L2SequencerABI = L2SequencerMetaData.ABI

// L2SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2SequencerMetaData.Bin instead.
var L2SequencerBin = L2SequencerMetaData.Bin

// DeployL2Sequencer deploys a new Ethereum contract, binding an instance of L2Sequencer to it.
func DeployL2Sequencer(auth *bind.TransactOpts, backend bind.ContractBackend, _otherSequencer common.Address) (common.Address, *types.Transaction, *L2Sequencer, error) {
	parsed, err := L2SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2SequencerBin), backend, _otherSequencer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2Sequencer{L2SequencerCaller: L2SequencerCaller{contract: contract}, L2SequencerTransactor: L2SequencerTransactor{contract: contract}, L2SequencerFilterer: L2SequencerFilterer{contract: contract}}, nil
}

// L2Sequencer is an auto generated Go binding around an Ethereum contract.
type L2Sequencer struct {
	L2SequencerCaller     // Read-only binding to the contract
	L2SequencerTransactor // Write-only binding to the contract
	L2SequencerFilterer   // Log filterer for contract events
}

// L2SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2SequencerSession struct {
	Contract     *L2Sequencer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2SequencerCallerSession struct {
	Contract *L2SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// L2SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2SequencerTransactorSession struct {
	Contract     *L2SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// L2SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2SequencerRaw struct {
	Contract *L2Sequencer // Generic contract binding to access the raw methods on
}

// L2SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2SequencerCallerRaw struct {
	Contract *L2SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// L2SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2SequencerTransactorRaw struct {
	Contract *L2SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2Sequencer creates a new instance of L2Sequencer, bound to a specific deployed contract.
func NewL2Sequencer(address common.Address, backend bind.ContractBackend) (*L2Sequencer, error) {
	contract, err := bindL2Sequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2Sequencer{L2SequencerCaller: L2SequencerCaller{contract: contract}, L2SequencerTransactor: L2SequencerTransactor{contract: contract}, L2SequencerFilterer: L2SequencerFilterer{contract: contract}}, nil
}

// NewL2SequencerCaller creates a new read-only instance of L2Sequencer, bound to a specific deployed contract.
func NewL2SequencerCaller(address common.Address, caller bind.ContractCaller) (*L2SequencerCaller, error) {
	contract, err := bindL2Sequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2SequencerCaller{contract: contract}, nil
}

// NewL2SequencerTransactor creates a new write-only instance of L2Sequencer, bound to a specific deployed contract.
func NewL2SequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2SequencerTransactor, error) {
	contract, err := bindL2Sequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2SequencerTransactor{contract: contract}, nil
}

// NewL2SequencerFilterer creates a new log filterer instance of L2Sequencer, bound to a specific deployed contract.
func NewL2SequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2SequencerFilterer, error) {
	contract, err := bindL2Sequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2SequencerFilterer{contract: contract}, nil
}

// bindL2Sequencer binds a generic wrapper to an already deployed contract.
func bindL2Sequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2SequencerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Sequencer *L2SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Sequencer.Contract.L2SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Sequencer *L2SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Sequencer.Contract.L2SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Sequencer *L2SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Sequencer.Contract.L2SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Sequencer *L2SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Sequencer *L2SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Sequencer *L2SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Sequencer.Contract.contract.Transact(opts, method, params...)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_L2Sequencer *L2SequencerCaller) L2SUBMITTERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "L2_SUBMITTER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_L2Sequencer *L2SequencerSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _L2Sequencer.Contract.L2SUBMITTERCONTRACT(&_L2Sequencer.CallOpts)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _L2Sequencer.Contract.L2SUBMITTERCONTRACT(&_L2Sequencer.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Sequencer *L2SequencerCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Sequencer *L2SequencerSession) MESSENGER() (common.Address, error) {
	return _L2Sequencer.Contract.MESSENGER(&_L2Sequencer.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) MESSENGER() (common.Address, error) {
	return _L2Sequencer.Contract.MESSENGER(&_L2Sequencer.CallOpts)
}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L2Sequencer *L2SequencerCaller) OTHERSEQUENCER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "OTHER_SEQUENCER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L2Sequencer *L2SequencerSession) OTHERSEQUENCER() (common.Address, error) {
	return _L2Sequencer.Contract.OTHERSEQUENCER(&_L2Sequencer.CallOpts)
}

// OTHERSEQUENCER is a free data retrieval call binding the contract method 0xf81e02a7.
//
// Solidity: function OTHER_SEQUENCER() view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) OTHERSEQUENCER() (common.Address, error) {
	return _L2Sequencer.Contract.OTHERSEQUENCER(&_L2Sequencer.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerCaller) CurrentVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "currentVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerSession) CurrentVersion() (*big.Int, error) {
	return _L2Sequencer.Contract.CurrentVersion(&_L2Sequencer.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerCallerSession) CurrentVersion() (*big.Int, error) {
	return _L2Sequencer.Contract.CurrentVersion(&_L2Sequencer.CallOpts)
}

// CurrentVersionHeight is a free data retrieval call binding the contract method 0xd9586467.
//
// Solidity: function currentVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerCaller) CurrentVersionHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "currentVersionHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentVersionHeight is a free data retrieval call binding the contract method 0xd9586467.
//
// Solidity: function currentVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerSession) CurrentVersionHeight() (*big.Int, error) {
	return _L2Sequencer.Contract.CurrentVersionHeight(&_L2Sequencer.CallOpts)
}

// CurrentVersionHeight is a free data retrieval call binding the contract method 0xd9586467.
//
// Solidity: function currentVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerCallerSession) CurrentVersionHeight() (*big.Int, error) {
	return _L2Sequencer.Contract.CurrentVersionHeight(&_L2Sequencer.CallOpts)
}

// GetSequencerAddresses is a free data retrieval call binding the contract method 0xe597c19e.
//
// Solidity: function getSequencerAddresses(bool previous) view returns(address[])
func (_L2Sequencer *L2SequencerCaller) GetSequencerAddresses(opts *bind.CallOpts, previous bool) ([]common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "getSequencerAddresses", previous)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencerAddresses is a free data retrieval call binding the contract method 0xe597c19e.
//
// Solidity: function getSequencerAddresses(bool previous) view returns(address[])
func (_L2Sequencer *L2SequencerSession) GetSequencerAddresses(previous bool) ([]common.Address, error) {
	return _L2Sequencer.Contract.GetSequencerAddresses(&_L2Sequencer.CallOpts, previous)
}

// GetSequencerAddresses is a free data retrieval call binding the contract method 0xe597c19e.
//
// Solidity: function getSequencerAddresses(bool previous) view returns(address[])
func (_L2Sequencer *L2SequencerCallerSession) GetSequencerAddresses(previous bool) ([]common.Address, error) {
	return _L2Sequencer.Contract.GetSequencerAddresses(&_L2Sequencer.CallOpts, previous)
}

// GetSequencerInfos is a free data retrieval call binding the contract method 0xbe6c5d68.
//
// Solidity: function getSequencerInfos(bool previous) view returns((address,bytes32,bytes)[])
func (_L2Sequencer *L2SequencerCaller) GetSequencerInfos(opts *bind.CallOpts, previous bool) ([]TypesSequencerInfo, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "getSequencerInfos", previous)

	if err != nil {
		return *new([]TypesSequencerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TypesSequencerInfo)).(*[]TypesSequencerInfo)

	return out0, err

}

// GetSequencerInfos is a free data retrieval call binding the contract method 0xbe6c5d68.
//
// Solidity: function getSequencerInfos(bool previous) view returns((address,bytes32,bytes)[])
func (_L2Sequencer *L2SequencerSession) GetSequencerInfos(previous bool) ([]TypesSequencerInfo, error) {
	return _L2Sequencer.Contract.GetSequencerInfos(&_L2Sequencer.CallOpts, previous)
}

// GetSequencerInfos is a free data retrieval call binding the contract method 0xbe6c5d68.
//
// Solidity: function getSequencerInfos(bool previous) view returns((address,bytes32,bytes)[])
func (_L2Sequencer *L2SequencerCallerSession) GetSequencerInfos(previous bool) ([]TypesSequencerInfo, error) {
	return _L2Sequencer.Contract.GetSequencerInfos(&_L2Sequencer.CallOpts, previous)
}

// InSequencersSet is a free data retrieval call binding the contract method 0xd1c55fe3.
//
// Solidity: function inSequencersSet(bool previous, address checkAddr) view returns(bool, uint256)
func (_L2Sequencer *L2SequencerCaller) InSequencersSet(opts *bind.CallOpts, previous bool, checkAddr common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "inSequencersSet", previous, checkAddr)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// InSequencersSet is a free data retrieval call binding the contract method 0xd1c55fe3.
//
// Solidity: function inSequencersSet(bool previous, address checkAddr) view returns(bool, uint256)
func (_L2Sequencer *L2SequencerSession) InSequencersSet(previous bool, checkAddr common.Address) (bool, *big.Int, error) {
	return _L2Sequencer.Contract.InSequencersSet(&_L2Sequencer.CallOpts, previous, checkAddr)
}

// InSequencersSet is a free data retrieval call binding the contract method 0xd1c55fe3.
//
// Solidity: function inSequencersSet(bool previous, address checkAddr) view returns(bool, uint256)
func (_L2Sequencer *L2SequencerCallerSession) InSequencersSet(previous bool, checkAddr common.Address) (bool, *big.Int, error) {
	return _L2Sequencer.Contract.InSequencersSet(&_L2Sequencer.CallOpts, previous, checkAddr)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Sequencer *L2SequencerCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Sequencer *L2SequencerSession) Messenger() (common.Address, error) {
	return _L2Sequencer.Contract.Messenger(&_L2Sequencer.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) Messenger() (common.Address, error) {
	return _L2Sequencer.Contract.Messenger(&_L2Sequencer.CallOpts)
}

// PreSequencerAddresses is a free data retrieval call binding the contract method 0xdd967ee9.
//
// Solidity: function preSequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerCaller) PreSequencerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "preSequencerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreSequencerAddresses is a free data retrieval call binding the contract method 0xdd967ee9.
//
// Solidity: function preSequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerSession) PreSequencerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Sequencer.Contract.PreSequencerAddresses(&_L2Sequencer.CallOpts, arg0)
}

// PreSequencerAddresses is a free data retrieval call binding the contract method 0xdd967ee9.
//
// Solidity: function preSequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) PreSequencerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Sequencer.Contract.PreSequencerAddresses(&_L2Sequencer.CallOpts, arg0)
}

// PreSequencerInfos is a free data retrieval call binding the contract method 0xb95cdb78.
//
// Solidity: function preSequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerCaller) PreSequencerInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "preSequencerInfos", arg0)

	outstruct := new(struct {
		Addr   common.Address
		TmKey  [32]byte
		BlsKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TmKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// PreSequencerInfos is a free data retrieval call binding the contract method 0xb95cdb78.
//
// Solidity: function preSequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerSession) PreSequencerInfos(arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Sequencer.Contract.PreSequencerInfos(&_L2Sequencer.CallOpts, arg0)
}

// PreSequencerInfos is a free data retrieval call binding the contract method 0xb95cdb78.
//
// Solidity: function preSequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerCallerSession) PreSequencerInfos(arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Sequencer.Contract.PreSequencerInfos(&_L2Sequencer.CallOpts, arg0)
}

// PreVersion is a free data retrieval call binding the contract method 0xcfd1eff3.
//
// Solidity: function preVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerCaller) PreVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "preVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreVersion is a free data retrieval call binding the contract method 0xcfd1eff3.
//
// Solidity: function preVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerSession) PreVersion() (*big.Int, error) {
	return _L2Sequencer.Contract.PreVersion(&_L2Sequencer.CallOpts)
}

// PreVersion is a free data retrieval call binding the contract method 0xcfd1eff3.
//
// Solidity: function preVersion() view returns(uint256)
func (_L2Sequencer *L2SequencerCallerSession) PreVersion() (*big.Int, error) {
	return _L2Sequencer.Contract.PreVersion(&_L2Sequencer.CallOpts)
}

// PreVersionHeight is a free data retrieval call binding the contract method 0x4a3c980c.
//
// Solidity: function preVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerCaller) PreVersionHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "preVersionHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreVersionHeight is a free data retrieval call binding the contract method 0x4a3c980c.
//
// Solidity: function preVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerSession) PreVersionHeight() (*big.Int, error) {
	return _L2Sequencer.Contract.PreVersionHeight(&_L2Sequencer.CallOpts)
}

// PreVersionHeight is a free data retrieval call binding the contract method 0x4a3c980c.
//
// Solidity: function preVersionHeight() view returns(uint256)
func (_L2Sequencer *L2SequencerCallerSession) PreVersionHeight() (*big.Int, error) {
	return _L2Sequencer.Contract.PreVersionHeight(&_L2Sequencer.CallOpts)
}

// SequencerAddresses is a free data retrieval call binding the contract method 0xaeaf9f41.
//
// Solidity: function sequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerCaller) SequencerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "sequencerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerAddresses is a free data retrieval call binding the contract method 0xaeaf9f41.
//
// Solidity: function sequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerSession) SequencerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Sequencer.Contract.SequencerAddresses(&_L2Sequencer.CallOpts, arg0)
}

// SequencerAddresses is a free data retrieval call binding the contract method 0xaeaf9f41.
//
// Solidity: function sequencerAddresses(uint256 ) view returns(address)
func (_L2Sequencer *L2SequencerCallerSession) SequencerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Sequencer.Contract.SequencerAddresses(&_L2Sequencer.CallOpts, arg0)
}

// SequencerIndex is a free data retrieval call binding the contract method 0x342b6345.
//
// Solidity: function sequencerIndex(bool previous, address checkAddr) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerCaller) SequencerIndex(opts *bind.CallOpts, previous bool, checkAddr common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "sequencerIndex", previous, checkAddr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SequencerIndex is a free data retrieval call binding the contract method 0x342b6345.
//
// Solidity: function sequencerIndex(bool previous, address checkAddr) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerSession) SequencerIndex(previous bool, checkAddr common.Address) (*big.Int, *big.Int, error) {
	return _L2Sequencer.Contract.SequencerIndex(&_L2Sequencer.CallOpts, previous, checkAddr)
}

// SequencerIndex is a free data retrieval call binding the contract method 0x342b6345.
//
// Solidity: function sequencerIndex(bool previous, address checkAddr) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerCallerSession) SequencerIndex(previous bool, checkAddr common.Address) (*big.Int, *big.Int, error) {
	return _L2Sequencer.Contract.SequencerIndex(&_L2Sequencer.CallOpts, previous, checkAddr)
}

// SequencerInfos is a free data retrieval call binding the contract method 0xc9406b1a.
//
// Solidity: function sequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerCaller) SequencerInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "sequencerInfos", arg0)

	outstruct := new(struct {
		Addr   common.Address
		TmKey  [32]byte
		BlsKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TmKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// SequencerInfos is a free data retrieval call binding the contract method 0xc9406b1a.
//
// Solidity: function sequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerSession) SequencerInfos(arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Sequencer.Contract.SequencerInfos(&_L2Sequencer.CallOpts, arg0)
}

// SequencerInfos is a free data retrieval call binding the contract method 0xc9406b1a.
//
// Solidity: function sequencerInfos(uint256 ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Sequencer *L2SequencerCallerSession) SequencerInfos(arg0 *big.Int) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Sequencer.Contract.SequencerInfos(&_L2Sequencer.CallOpts, arg0)
}

// SequencersLen is a free data retrieval call binding the contract method 0x7ad9e3ac.
//
// Solidity: function sequencersLen(bool previous) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerCaller) SequencersLen(opts *bind.CallOpts, previous bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "sequencersLen", previous)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SequencersLen is a free data retrieval call binding the contract method 0x7ad9e3ac.
//
// Solidity: function sequencersLen(bool previous) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerSession) SequencersLen(previous bool) (*big.Int, *big.Int, error) {
	return _L2Sequencer.Contract.SequencersLen(&_L2Sequencer.CallOpts, previous)
}

// SequencersLen is a free data retrieval call binding the contract method 0x7ad9e3ac.
//
// Solidity: function sequencersLen(bool previous) view returns(uint256, uint256)
func (_L2Sequencer *L2SequencerCallerSession) SequencersLen(previous bool) (*big.Int, *big.Int, error) {
	return _L2Sequencer.Contract.SequencersLen(&_L2Sequencer.CallOpts, previous)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2Sequencer *L2SequencerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2Sequencer *L2SequencerSession) Version() (string, error) {
	return _L2Sequencer.Contract.Version(&_L2Sequencer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2Sequencer *L2SequencerCallerSession) Version() (string, error) {
	return _L2Sequencer.Contract.Version(&_L2Sequencer.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x5942e7c7.
//
// Solidity: function initialize((address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerTransactor) Initialize(opts *bind.TransactOpts, _sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.contract.Transact(opts, "initialize", _sequencers)
}

// Initialize is a paid mutator transaction binding the contract method 0x5942e7c7.
//
// Solidity: function initialize((address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerSession) Initialize(_sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.Contract.Initialize(&_L2Sequencer.TransactOpts, _sequencers)
}

// Initialize is a paid mutator transaction binding the contract method 0x5942e7c7.
//
// Solidity: function initialize((address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerTransactorSession) Initialize(_sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.Contract.Initialize(&_L2Sequencer.TransactOpts, _sequencers)
}

// UpdateSequencers is a paid mutator transaction binding the contract method 0xad01732f.
//
// Solidity: function updateSequencers(uint256 version, (address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerTransactor) UpdateSequencers(opts *bind.TransactOpts, version *big.Int, _sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.contract.Transact(opts, "updateSequencers", version, _sequencers)
}

// UpdateSequencers is a paid mutator transaction binding the contract method 0xad01732f.
//
// Solidity: function updateSequencers(uint256 version, (address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerSession) UpdateSequencers(version *big.Int, _sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.Contract.UpdateSequencers(&_L2Sequencer.TransactOpts, version, _sequencers)
}

// UpdateSequencers is a paid mutator transaction binding the contract method 0xad01732f.
//
// Solidity: function updateSequencers(uint256 version, (address,bytes32,bytes)[] _sequencers) returns()
func (_L2Sequencer *L2SequencerTransactorSession) UpdateSequencers(version *big.Int, _sequencers []TypesSequencerInfo) (*types.Transaction, error) {
	return _L2Sequencer.Contract.UpdateSequencers(&_L2Sequencer.TransactOpts, version, _sequencers)
}

// L2SequencerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2Sequencer contract.
type L2SequencerInitializedIterator struct {
	Event *L2SequencerInitialized // Event containing the contract specifics and raw log

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
func (it *L2SequencerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2SequencerInitialized)
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
		it.Event = new(L2SequencerInitialized)
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
func (it *L2SequencerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2SequencerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2SequencerInitialized represents a Initialized event raised by the L2Sequencer contract.
type L2SequencerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2Sequencer *L2SequencerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2SequencerInitializedIterator, error) {

	logs, sub, err := _L2Sequencer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2SequencerInitializedIterator{contract: _L2Sequencer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2Sequencer *L2SequencerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2SequencerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2Sequencer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2SequencerInitialized)
				if err := _L2Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2Sequencer *L2SequencerFilterer) ParseInitialized(log types.Log) (*L2SequencerInitialized, error) {
	event := new(L2SequencerInitialized)
	if err := _L2Sequencer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2SequencerSequencerUpdatedIterator is returned from FilterSequencerUpdated and is used to iterate over the raw logs and unpacked data for SequencerUpdated events raised by the L2Sequencer contract.
type L2SequencerSequencerUpdatedIterator struct {
	Event *L2SequencerSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *L2SequencerSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2SequencerSequencerUpdated)
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
		it.Event = new(L2SequencerSequencerUpdated)
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
func (it *L2SequencerSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2SequencerSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2SequencerSequencerUpdated represents a SequencerUpdated event raised by the L2Sequencer contract.
type L2SequencerSequencerUpdated struct {
	Sequencers []common.Address
	Version    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0x71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d6182796.
//
// Solidity: event SequencerUpdated(address[] sequencers, uint256 version)
func (_L2Sequencer *L2SequencerFilterer) FilterSequencerUpdated(opts *bind.FilterOpts) (*L2SequencerSequencerUpdatedIterator, error) {

	logs, sub, err := _L2Sequencer.contract.FilterLogs(opts, "SequencerUpdated")
	if err != nil {
		return nil, err
	}
	return &L2SequencerSequencerUpdatedIterator{contract: _L2Sequencer.contract, event: "SequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0x71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d6182796.
//
// Solidity: event SequencerUpdated(address[] sequencers, uint256 version)
func (_L2Sequencer *L2SequencerFilterer) WatchSequencerUpdated(opts *bind.WatchOpts, sink chan<- *L2SequencerSequencerUpdated) (event.Subscription, error) {

	logs, sub, err := _L2Sequencer.contract.WatchLogs(opts, "SequencerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2SequencerSequencerUpdated)
				if err := _L2Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0x71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d6182796.
//
// Solidity: event SequencerUpdated(address[] sequencers, uint256 version)
func (_L2Sequencer *L2SequencerFilterer) ParseSequencerUpdated(log types.Log) (*L2SequencerSequencerUpdated, error) {
	event := new(L2SequencerSequencerUpdated)
	if err := _L2Sequencer.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
