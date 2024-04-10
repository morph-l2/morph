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

// TypesSequencerHistory is an auto generated low-level Go binding around an user-defined struct.
type TypesSequencerHistory struct {
	SequencerAddresses []common.Address
	Timestamp          *big.Int
}

// TypesSequencerInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesSequencerInfo struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}

// L2SequencerMetaData contains all meta data concerning the L2Sequencer contract.
var L2SequencerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherSequencer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_SUBMITTER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getSequencerHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sequencerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.SequencerHistory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"inSequencersSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"sequencerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"sequencersLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"updateSequencers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040525f6001555f6002555f60035534801561001b575f80fd5b50604051611ebf380380611ebf83398101604081905261003a9161007b565b7353000000000000000000000000000000000000076080526001600160a01b031660a05273530000000000000000000000000000000000000560c0526100a8565b5f6020828403121561008b575f80fd5b81516001600160a01b03811681146100a1575f80fd5b9392505050565b60805160a05160c051611dd16100ee5f395f6101c301525f818161033f01526108f701525f818161016201528181610234015281816108cd015261092e0152611dd15ff3fe608060405234801561000f575f80fd5b506004361061012f575f3560e01c8063ad01732f116100ad578063d1c55fe31161007d578063e597c19e11610063578063e597c19e146102fa578063f6f207ce1461031a578063f81e02a71461033a575f80fd5b8063d1c55fe3146102c7578063d9586467146102f1575f80fd5b8063ad01732f1461025f578063b95cdb7814610272578063be6c5d6814610294578063c9406b1a146102b4575f80fd5b80635942e7c7116101025780637ad9e3ac116100e85780637ad9e3ac1461021c578063927ede2d1461022f5780639d888e8614610256575f80fd5b80635942e7c7146101e55780636d8ce3d2146101fa575f80fd5b8063342b6345146101335780633cb747bf146101605780634a3c980c146101a75780634bbf5252146101be575b5f80fd5b610146610141366004611529565b610361565b604080519283526020830191909152015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610157565b6101b060035481565b604051908152602001610157565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101f86101f336600461175c565b610508565b005b6101b0610208366004611796565b60046020525f908152604090206001015481565b61014661022a3660046117ad565b61084f565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101b060015481565b6101f861026d3660046117cd565b6108b5565b610285610280366004611796565b610ced565b60405161015793929190611872565b6102a76102a23660046117ad565b610dbf565b60405161015791906118af565b6102856102c2366004611796565b610ff0565b6102da6102d5366004611529565b610fff565b604080519215158352602083019190915201610157565b6101b060025481565b61030d6103083660046117ad565b611157565b60405161015791906119b2565b61032d610328366004611796565b61126c565b60405161015791906119c4565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b5f8083801561037157505f600154115b15610485575f5b60045f600180546103899190611a22565b815260208101919091526040015f205481101561041d5760045f600180546103b19190611a22565b81526020019081526020015f205f0181815481106103d1576103d1611a3b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169085160361041557806001805461040b9190611a22565b9250925050610501565b600101610378565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f73657175656e636572206e6f742065786973740000000000000000000000000060448201526064015b60405180910390fd5b5f5b6001545f9081526004602052604090205481101561041d576001545f9081526004602052604090208054829081106104c1576104c1611a3b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036104f9576001549092509050610501565b600101610487565b9250929050565b5f54610100900460ff161580801561052657505f54600160ff909116105b8061053f5750303b15801561053f57505f5460ff166001145b6105cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161047c565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610627575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f825167ffffffffffffffff8111156106425761064261155e565b60405190808252806020026020018201604052801561066b578160200160208202803683370190505b5090505f5b83518110156107875783818151811061068b5761068b611a3b565b60200260200101515f01518282815181106106a8576106a8611a3b565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060058482815181106106f6576106f6611a3b565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117815591810151928201929092556040820151600282019061077c9082611b05565b505050600101610670565b5060408051808201909152818152426020808301919091525f805260048152815180517f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec926107da928492910190611308565b506020820151816001015590505050801561084b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b5f8082801561085f57505f600154115b1561089c5760045f600180546108759190611a22565b815260208101919091526040015f2054600180546108939190611a22565b91509150915091565b50506001545f8181526004602052604090205492909150565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156109d157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610995573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b99190611c21565b73ffffffffffffffffffffffffffffffffffffffff16145b610a5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53657175656e6365723a2066756e6374696f6e2063616e206f6e6c792062652060448201527f63616c6c65642066726f6d20746865206f746865722073657175656e63657200606482015260840161047c565b60018054610a6a91611c3c565b8214610ad2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e76616c69642076657273696f6e0000000000000000000000000000000000604482015260640161047c565b6001829055610ae260065f611390565b60058054610af2916006916113b1565b50610afe60055f611390565b6002805460035543905580515f9067ffffffffffffffff811115610b2457610b2461155e565b604051908082528060200260200182016040528015610b4d578160200160208202803683370190505b5090505f5b8251811015610c6957828181518110610b6d57610b6d611a3b565b60200260200101515f0151828281518110610b8a57610b8a611a3b565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506005838281518110610bd857610bd8611a3b565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091178155918101519282019290925560408201516002820190610c5e9082611b05565b505050600101610b52565b50604080518082018252828152426020808301919091525f86815260048252929092208151805192939192610ca19284920190611308565b50602082015181600101559050507f71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d61827968184604051610ce0929190611c4f565b60405180910390a1505050565b60068181548110610cfc575f80fd5b5f91825260209091206003909102018054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff9093169450909291610d3e90611a68565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a90611a68565b8015610db55780601f10610d8c57610100808354040283529160200191610db5565b820191905f5260205f20905b815481529060010190602001808311610d9857829003601f168201915b5050505050905083565b60608115610ee1576006805480602002602001604051908101604052809291908181526020015f905b82821015610ed6575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610e4790611a68565b80601f0160208091040260200160405190810160405280929190818152602001828054610e7390611a68565b8015610ebe5780601f10610e9557610100808354040283529160200191610ebe565b820191905f5260205f20905b815481529060010190602001808311610ea157829003601f168201915b50505050508152505081526020019060010190610de8565b505050509050919050565b6005805480602002602001604051908101604052809291908181526020015f905b82821015610ed6575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610f6190611a68565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8d90611a68565b8015610fd85780601f10610faf57610100808354040283529160200191610fd8565b820191905f5260205f20905b815481529060010190602001808311610fbb57829003601f168201915b50505050508152505081526020019060010190610f02565b60058181548110610cfc575f80fd5b5f8083801561100f57505f600154115b156110cb575f5b60045f600180546110279190611a22565b815260208101919091526040015f20548110156110b25760045f6001805461104f9190611a22565b81526020019081526020015f205f01818154811061106f5761106f611a3b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036110aa5760018060015461040b9190611a22565b600101611016565b505f600180546110c29190611a22565b91509150610501565b5f5b6001545f90815260046020526040902054811015611148576001545f90815260046020526040902080548290811061110757611107611a3b565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169085160361114057600180549250925050610501565b6001016110cd565b50506001545f91509250929050565b606081801561116757505f600154115b156111f75760045f6001805461117d9190611a22565b81526020019081526020015f205f018054806020026020016040519081016040528092919081815260200182805480156111eb57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116111c0575b50505050509050919050565b6001545f90815260046020908152604091829020805483518184028101840190945280845290918301828280156111eb57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116111c05750505050509050919050565b60408051808201825260608082525f60208084018290528582526004815290849020845181549283028101840186529485018281529394939092849284918401828280156112ee57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116112c3575b505050505081526020016001820154815250509050919050565b828054828255905f5260205f20908101928215611380579160200282015b8281111561138057825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611326565b5061138c92915061145d565b5090565b5080545f8255600302905f5260205f20908101906113ae9190611471565b50565b828054828255905f5260205f20906003028101928215611451575f5260205f209160030282015b8281111561145157825482547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560018084015490830155828260028082019061143f90840182611c70565b505050916003019190600301906113d8565b5061138c929150611471565b5b8082111561138c575f815560010161145e565b8082111561138c5780547fffffffffffffffffffffffff00000000000000000000000000000000000000001681555f600182018190556114b460028301826114bd565b50600301611471565b5080546114c990611a68565b5f825580601f106114d8575050565b601f0160209004905f5260205f20908101906113ae919061145d565b80358015158114611503575f80fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146113ae575f80fd5b5f806040838503121561153a575f80fd5b611543836114f4565b9150602083013561155381611508565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156115ae576115ae61155e565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156115fb576115fb61155e565b604052919050565b5f601f83601f840112611614575f80fd5b8235602067ffffffffffffffff808311156116315761163161155e565b8260051b6116408382016115b4565b9384528681018301938381019089861115611659575f80fd5b84890192505b8583101561174f57823584811115611675575f80fd5b890160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d0381018213156116aa575f80fd5b6116b261158b565b888401356116bf81611508565b81526040848101358a8301529284013592888411156116dc575f80fd5b83850194508e603f8601126116ef575f80fd5b898501359350888411156117055761170561155e565b6117148a848e870116016115b4565b92508383528e81858701011115611729575f80fd5b838186018b8501375f9383018a019390935291820152835250918401919084019061165f565b9998505050505050505050565b5f6020828403121561176c575f80fd5b813567ffffffffffffffff811115611782575f80fd5b61178e84828501611603565b949350505050565b5f602082840312156117a6575f80fd5b5035919050565b5f602082840312156117bd575f80fd5b6117c6826114f4565b9392505050565b5f80604083850312156117de575f80fd5b82359150602083013567ffffffffffffffff8111156117fb575f80fd5b61180785828601611603565b9150509250929050565b5f81518084525f5b8181101561183557602081850181015186830182015201611819565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f6118a66060830184611811565b95945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b83811015611954578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff1684528781015188850152860151606087850181905261194081860183611811565b9689019694505050908601906001016118d6565b509098975050505050505050565b5f815180845260208085019450602084015f5b838110156119a757815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611975565b509495945050505050565b602081525f6117c66020830184611962565b602081525f8251604060208401526119df6060840182611962565b9050602084015160408401528091505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115611a3557611a356119f5565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680611a7c57607f821691505b602082108103611ab3577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f821115611b0057805f5260205f20601f840160051c81016020851015611ade5750805b601f840160051c820191505b81811015611afd575f8155600101611aea565b50505b505050565b815167ffffffffffffffff811115611b1f57611b1f61155e565b611b3381611b2d8454611a68565b84611ab9565b602080601f831160018114611b85575f8415611b4f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611c19565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611bd157888601518255948401946001909101908401611bb2565b5085821015611c0d57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611c31575f80fd5b81516117c681611508565b80820180821115611a3557611a356119f5565b604081525f611c616040830185611962565b90508260208301529392505050565b818103611c7b575050565b611c858254611a68565b67ffffffffffffffff811115611c9d57611c9d61155e565b611cab81611b2d8454611a68565b5f601f821160018114611cfb575f8315611cc55750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611afd565b5f85815260208082208683529082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616925b83811015611d4f5782860154825560019586019590910190602001611d2f565b5085831015611d8b57818501547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122020a9c785b73d3ef4a6e2026110e665d963e3d622265ed74c7f83e400dfff235864736f6c63430008180033",
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

// GetSequencerHistory is a free data retrieval call binding the contract method 0xf6f207ce.
//
// Solidity: function getSequencerHistory(uint256 version) view returns((address[],uint256))
func (_L2Sequencer *L2SequencerCaller) GetSequencerHistory(opts *bind.CallOpts, version *big.Int) (TypesSequencerHistory, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "getSequencerHistory", version)

	if err != nil {
		return *new(TypesSequencerHistory), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesSequencerHistory)).(*TypesSequencerHistory)

	return out0, err

}

// GetSequencerHistory is a free data retrieval call binding the contract method 0xf6f207ce.
//
// Solidity: function getSequencerHistory(uint256 version) view returns((address[],uint256))
func (_L2Sequencer *L2SequencerSession) GetSequencerHistory(version *big.Int) (TypesSequencerHistory, error) {
	return _L2Sequencer.Contract.GetSequencerHistory(&_L2Sequencer.CallOpts, version)
}

// GetSequencerHistory is a free data retrieval call binding the contract method 0xf6f207ce.
//
// Solidity: function getSequencerHistory(uint256 version) view returns((address[],uint256))
func (_L2Sequencer *L2SequencerCallerSession) GetSequencerHistory(version *big.Int) (TypesSequencerHistory, error) {
	return _L2Sequencer.Contract.GetSequencerHistory(&_L2Sequencer.CallOpts, version)
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

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_L2Sequencer *L2SequencerCaller) SequencerHistory(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2Sequencer.contract.Call(opts, &out, "sequencerHistory", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_L2Sequencer *L2SequencerSession) SequencerHistory(arg0 *big.Int) (*big.Int, error) {
	return _L2Sequencer.Contract.SequencerHistory(&_L2Sequencer.CallOpts, arg0)
}

// SequencerHistory is a free data retrieval call binding the contract method 0x6d8ce3d2.
//
// Solidity: function sequencerHistory(uint256 ) view returns(uint256 timestamp)
func (_L2Sequencer *L2SequencerCallerSession) SequencerHistory(arg0 *big.Int) (*big.Int, error) {
	return _L2Sequencer.Contract.SequencerHistory(&_L2Sequencer.CallOpts, arg0)
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
