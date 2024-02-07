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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherSequencer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_SUBMITTER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"inSequencersSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSequencerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"sequencerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"sequencersLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"updateSequencers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040525f6001555f6002555f6003555f600455348015610020575f80fd5b50604051611f33380380611f3383398101604081905261003f91610090565b60016080525f60a081905260c05273530000000000000000000000000000000000000760e0526001600160a01b031661010052735300000000000000000000000000000000000005610120526100bd565b5f602082840312156100a0575f80fd5b81516001600160a01b03811681146100b6575f80fd5b9392505050565b60805160a05160c05160e0516101005161012051611e0c6101275f395f81816101f70152610a2801525f8181610375015261088501525f81816101960152818161025b0152818161085b01526108bc01525f61052501525f6104fc01525f6104d30152611e0c5ff3fe608060405234801561000f575f80fd5b5060043610610163575f3560e01c8063aeaf9f41116100c7578063d1c55fe31161007d578063dd967ee911610063578063dd967ee91461033d578063e597c19e14610350578063f81e02a714610370575f80fd5b8063d1c55fe31461030a578063d958646714610334575f80fd5b8063be6c5d68116100ad578063be6c5d68146102ce578063c9406b1a146102ee578063cfd1eff314610301575f80fd5b8063aeaf9f4114610299578063b95cdb78146102ac575f80fd5b80635942e7c71161011c578063927ede2d11610102578063927ede2d146102565780639d888e861461027d578063ad01732f14610286575f80fd5b80635942e7c71461022e5780637ad9e3ac14610243575f80fd5b80634a3c980c1161014c5780634a3c980c146101db5780634bbf5252146101f257806354fd4d5014610219575f80fd5b8063342b6345146101675780633cb747bf14610194575b5f80fd5b61017a610175366004611501565b610397565b604080519283526020830191909152015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018b565b6101e460045481565b60405190815260200161018b565b6101b67f000000000000000000000000000000000000000000000000000000000000000081565b6102216104cc565b60405161018b91906115a1565b61024161023c3660046117b8565b61056f565b005b61017a6102513660046117f2565b61081d565b6101b67f000000000000000000000000000000000000000000000000000000000000000081565b6101e460015481565b61024161029436600461180b565b610843565b6101b66102a736600461184f565b610c49565b6102bf6102ba36600461184f565b610c7e565b60405161018b93929190611866565b6102e16102dc3660046117f2565b610d50565b60405161018b91906118a3565b6102bf6102fc36600461184f565b610f81565b6101e460035481565b61031d610318366004611501565b610f90565b60408051921515835260208301919091520161018b565b6101e460025481565b6101b661034b36600461184f565b611074565b61036361035e3660046117f2565b611083565b60405161018b9190611956565b6101b67f000000000000000000000000000000000000000000000000000000000000000081565b5f808315610466575f5b6006548110156103fe57600681815481106103be576103be6119af565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036103f65760035490925090506104c5565b6001016103a1565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f73657175656e636572206e6f742065786973740000000000000000000000000060448201526064015b60405180910390fd5b5f5b6005548110156103fe5760058181548110610485576104856119af565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036104bd5760015490925090506104c5565b600101610468565b9250929050565b60606104f77f0000000000000000000000000000000000000000000000000000000000000000611163565b6105207f0000000000000000000000000000000000000000000000000000000000000000611163565b6105497f0000000000000000000000000000000000000000000000000000000000000000611163565b60405160200161055b939291906119dc565b604051602081830303815290604052905090565b5f54610100900460ff161580801561058d57505f54600160ff909116105b806105a65750303b1580156105a657505f5460ff166001145b610632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161045d565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561068e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f5b82518110156107b55760058382815181106106ad576106ad6119af565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558251600790849083908110610724576107246119af565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911781559181015192820192909255604082015160028201906107aa9082611aee565b505050600101610690565b508015610819575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b5f8082156108345750506006546003549092909150565b50506005546001549092909150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561095f57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610923573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109479190611c0a565b73ffffffffffffffffffffffffffffffffffffffff16145b6109eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53657175656e6365723a2066756e6374696f6e2063616e206f6e6c792062652060448201527f63616c6c65642066726f6d20746865206f746865722073657175656e63657200606482015260840161045d565b6040517f16e2994a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906316e2994a90610a5e90600590600401611c78565b5f604051808303815f87803b158015610a75575f80fd5b505af1158015610a87573d5f803e3d5ffd5b505060015460035550610a9d905060085f611301565b610aa860065f611322565b60078054610ab89160089161133d565b5060058054610ac9916006916113ed565b506002546004556001829055610ae060075f611301565b610aeb60055f611322565b436002555f5b8151811015610c16576005828281518110610b0e57610b0e6119af565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558151600790839083908110610b8557610b856119af565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091178155918101519282019290925560408201516002820190610c0b9082611aee565b505050600101610af1565b507f71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d6182796600583604051610810929190611c8a565b60058181548110610c58575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60088181548110610c8d575f80fd5b5f91825260209091206003909102018054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff9093169450909291610ccf90611a51565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfb90611a51565b8015610d465780601f10610d1d57610100808354040283529160200191610d46565b820191905f5260205f20905b815481529060010190602001808311610d2957829003601f168201915b5050505050905083565b60608115610e72576008805480602002602001604051908101604052809291908181526020015f905b82821015610e67575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610dd890611a51565b80601f0160208091040260200160405190810160405280929190818152602001828054610e0490611a51565b8015610e4f5780601f10610e2657610100808354040283529160200191610e4f565b820191905f5260205f20905b815481529060010190602001808311610e3257829003601f168201915b50505050508152505081526020019060010190610d79565b505050509050919050565b6007805480602002602001604051908101604052809291908181526020015f905b82821015610e67575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610ef290611a51565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1e90611a51565b8015610f695780601f10610f4057610100808354040283529160200191610f69565b820191905f5260205f20905b815481529060010190602001808311610f4c57829003601f168201915b50505050508152505081526020019060010190610e93565b60078181548110610c8d575f80fd5b5f808315611005575f5b600654811015610ff85760068181548110610fb757610fb76119af565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603610ff0575050600354600191506104c5565b600101610f9a565b50506003545f91506104c5565b5f5b6005548110156110655760058181548110611024576110246119af565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169085160361105d576001805492509250506104c5565b600101611007565b50506001545f91509250929050565b60068181548110610c58575f80fd5b606081156110f85760068054806020026020016040519081016040528092919081815260200182805480156110ec57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116110c1575b50505050509050919050565b60058054806020026020016040519081016040528092919081815260200182805480156110ec57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116110c15750505050509050919050565b60605f61116f8361121f565b60010190505f8167ffffffffffffffff81111561118e5761118e6115ba565b6040519080825280601f01601f1916602001820160405280156111b8576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846111c257509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611267577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611293576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106112b157662386f26fc10000830492506010015b6305f5e10083106112c9576305f5e100830492506008015b61271083106112dd57612710830492506004015b606483106112ef576064830492506002015b600a83106112fb576001015b92915050565b5080545f8255600302905f5260205f209081019061131f9190611435565b50565b5080545f8255905f5260205f209081019061131f9190611481565b828054828255905f5260205f209060030281019282156113dd575f5260205f209160030282015b828111156113dd57825482547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556001808401549083015582826002808201906113cb90840182611cab565b50505091600301919060030190611364565b506113e9929150611435565b5090565b828054828255905f5260205f20908101928215611429575f5260205f209182015b8281111561142957825482559160010191906001019061140e565b506113e9929150611481565b808211156113e95780547fffffffffffffffffffffffff00000000000000000000000000000000000000001681555f600182018190556114786002830182611495565b50600301611435565b5b808211156113e9575f8155600101611482565b5080546114a190611a51565b5f825580601f106114b0575050565b601f0160209004905f5260205f209081019061131f9190611481565b803580151581146114db575f80fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461131f575f80fd5b5f8060408385031215611512575f80fd5b61151b836114cc565b9150602083013561152b816114e0565b809150509250929050565b5f5b83811015611550578181015183820152602001611538565b50505f910152565b5f815180845261156f816020860160208601611536565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6115b36020830184611558565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561160a5761160a6115ba565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611657576116576115ba565b604052919050565b5f601f83601f840112611670575f80fd5b8235602067ffffffffffffffff8083111561168d5761168d6115ba565b8260051b61169c838201611610565b93845286810183019383810190898611156116b5575f80fd5b84890192505b858310156117ab578235848111156116d1575f80fd5b890160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d038101821315611706575f80fd5b61170e6115e7565b8884013561171b816114e0565b81526040848101358a830152928401359288841115611738575f80fd5b83850194508e603f86011261174b575f80fd5b89850135935088841115611761576117616115ba565b6117708a848e87011601611610565b92508383528e81858701011115611785575f80fd5b838186018b8501375f9383018a01939093529182015283525091840191908401906116bb565b9998505050505050505050565b5f602082840312156117c8575f80fd5b813567ffffffffffffffff8111156117de575f80fd5b6117ea8482850161165f565b949350505050565b5f60208284031215611802575f80fd5b6115b3826114cc565b5f806040838503121561181c575f80fd5b82359150602083013567ffffffffffffffff811115611839575f80fd5b6118458582860161165f565b9150509250929050565b5f6020828403121561185f575f80fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f61189a6060830184611558565b95945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b83811015611948578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff1684528781015188850152860151606087850181905261193481860183611558565b9689019694505050908601906001016118ca565b509098975050505050505050565b602080825282518282018190525f9190848201906040850190845b818110156119a357835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611971565b50909695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f84516119ed818460208901611536565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611a29816001850160208a01611536565b60019201918201528351611a44816002840160208801611536565b0160020195945050505050565b600181811c90821680611a6557607f821691505b602082108103611a9c577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f821115611ae957805f5260205f20601f840160051c81016020851015611ac75750805b601f840160051c820191505b81811015611ae6575f8155600101611ad3565b50505b505050565b815167ffffffffffffffff811115611b0857611b086115ba565b611b1c81611b168454611a51565b84611aa2565b602080601f831160018114611b6e575f8415611b385750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611c02565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611bba57888601518255948401946001909101908401611b9b565b5085821015611bf657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611c1a575f80fd5b81516115b3816114e0565b5f815480845260208085019450835f5260205f205f5b83811015611c6d57815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101611c3b565b509495945050505050565b602081525f6115b36020830184611c25565b604081525f611c9c6040830185611c25565b90508260208301529392505050565b818103611cb6575050565b611cc08254611a51565b67ffffffffffffffff811115611cd857611cd86115ba565b611ce681611b168454611a51565b5f601f821160018114611d36575f8315611d005750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611ae6565b5f85815260208082208683529082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616925b83811015611d8a5782860154825560019586019590910190602001611d6a565b5085831015611dc657818501547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220f4e871adf154faafcfa9e55d2f049b380466de4efb707b07b8e22a6ae61a0f2f64736f6c63430008180033",
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
