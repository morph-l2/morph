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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherSequencer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"SequencerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_SUBMITTER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_SEQUENCER\",\"outputs\":[{\"internalType\":\"contractSequencer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"getSequencerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"inSequencersSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSequencerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preVersionHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"sequencerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"previous\",\"type\":\"bool\"}],\"name\":\"sequencersLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.SequencerInfo[]\",\"name\":\"_sequencers\",\"type\":\"tuple[]\"}],\"name\":\"updateSequencers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040525f6001555f6002555f6003555f60045534801561001f575f80fd5b50604051611c04380380611c0483398101604081905261003e9161007f565b7353000000000000000000000000000000000000076080526001600160a01b031660a05273530000000000000000000000000000000000000560c0526100ac565b5f6020828403121561008f575f80fd5b81516001600160a01b03811681146100a5575f80fd5b9392505050565b60805160a05160c051611b0b6100f95f395f81816101dd015261095601525f818161034601526107b301525f818161017c0152818161022c0152818161078901526107ea0152611b0b5ff3fe608060405234801561000f575f80fd5b5060043610610149575f3560e01c8063aeaf9f41116100c7578063d1c55fe31161007d578063dd967ee911610063578063dd967ee91461030e578063e597c19e14610321578063f81e02a714610341575f80fd5b8063d1c55fe3146102db578063d958646714610305575f80fd5b8063be6c5d68116100ad578063be6c5d681461029f578063c9406b1a146102bf578063cfd1eff3146102d2575f80fd5b8063aeaf9f411461026a578063b95cdb781461027d575f80fd5b80635942e7c71161011c578063927ede2d11610102578063927ede2d146102275780639d888e861461024e578063ad01732f14610257575f80fd5b80635942e7c7146101ff5780637ad9e3ac14610214575f80fd5b8063342b63451461014d5780633cb747bf1461017a5780634a3c980c146101c15780634bbf5252146101d8575b5f80fd5b61016061015b366004611291565b610368565b604080519283526020830191909152015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610171565b6101ca60045481565b604051908152602001610171565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b61021261020d3660046114c4565b61049d565b005b6101606102223660046114fe565b61074b565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b6101ca60015481565b61021261026536600461151e565b610771565b61019c610278366004611562565b610b77565b61029061028b366004611562565b610bac565b604051610171939291906115da565b6102b26102ad3660046114fe565b610c7e565b6040516101719190611617565b6102906102cd366004611562565b610eaf565b6101ca60035481565b6102ee6102e9366004611291565b610ebe565b604080519215158352602083019190915201610171565b6101ca60025481565b61019c61031c366004611562565b610fa2565b61033461032f3660046114fe565b610fb1565b60405161017191906116ca565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b5f808315610437575f5b6006548110156103cf576006818154811061038f5761038f611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036103c7576003549092509050610496565b600101610372565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f73657175656e636572206e6f742065786973740000000000000000000000000060448201526064015b60405180910390fd5b5f5b6005548110156103cf576005818154811061045657610456611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169085160361048e576001549092509050610496565b600101610439565b9250929050565b5f54610100900460ff16158080156104bb57505f54600160ff909116105b806104d45750303b1580156104d457505f5460ff166001145b610560576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161042e565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105bc575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f5b82518110156106e35760058382815181106105db576105db611723565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055825160079084908390811061065257610652611723565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911781559181015192820192909255604082015160028201906106d890826117ed565b5050506001016105be565b508015610747575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b5f8082156107625750506006546003549092909150565b50506005546001549092909150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561088d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610851573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108759190611909565b73ffffffffffffffffffffffffffffffffffffffff16145b610919576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53657175656e6365723a2066756e6374696f6e2063616e206f6e6c792062652060448201527f63616c6c65642066726f6d20746865206f746865722073657175656e63657200606482015260840161042e565b6040517f16e2994a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906316e2994a9061098c90600590600401611977565b5f604051808303815f87803b1580156109a3575f80fd5b505af11580156109b5573d5f803e3d5ffd5b5050600154600355506109cb905060085f611091565b6109d660065f6110b2565b600780546109e6916008916110cd565b50600580546109f79160069161117d565b506002546004556001829055610a0e60075f611091565b610a1960055f6110b2565b436002555f5b8151811015610b44576005828281518110610a3c57610a3c611723565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558151600790839083908110610ab357610ab3611723565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091178155918101519282019290925560408201516002820190610b3990826117ed565b505050600101610a1f565b507f71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d618279660058360405161073e929190611989565b60058181548110610b86575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60088181548110610bbb575f80fd5b5f91825260209091206003909102018054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff9093169450909291610bfd90611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2990611750565b8015610c745780601f10610c4b57610100808354040283529160200191610c74565b820191905f5260205f20905b815481529060010190602001808311610c5757829003601f168201915b5050505050905083565b60608115610da0576008805480602002602001604051908101604052809291908181526020015f905b82821015610d95575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610d0690611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3290611750565b8015610d7d5780601f10610d5457610100808354040283529160200191610d7d565b820191905f5260205f20905b815481529060010190602001808311610d6057829003601f168201915b50505050508152505081526020019060010190610ca7565b505050509050919050565b6007805480602002602001604051908101604052809291908181526020015f905b82821015610d95575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610e2090611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4c90611750565b8015610e975780601f10610e6e57610100808354040283529160200191610e97565b820191905f5260205f20905b815481529060010190602001808311610e7a57829003601f168201915b50505050508152505081526020019060010190610dc1565b60078181548110610bbb575f80fd5b5f808315610f33575f5b600654811015610f265760068181548110610ee557610ee5611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603610f1e57505060035460019150610496565b600101610ec8565b50506003545f9150610496565b5f5b600554811015610f935760058181548110610f5257610f52611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603610f8b57600180549250925050610496565b600101610f35565b50506001545f91509250929050565b60068181548110610b86575f80fd5b6060811561102657600680548060200260200160405190810160405280929190818152602001828054801561101a57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610fef575b50505050509050919050565b600580548060200260200160405190810160405280929190818152602001828054801561101a57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610fef5750505050509050919050565b5080545f8255600302905f5260205f20908101906110af91906111c5565b50565b5080545f8255905f5260205f20908101906110af9190611211565b828054828255905f5260205f2090600302810192821561116d575f5260205f209160030282015b8281111561116d57825482547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560018084015490830155828260028082019061115b908401826119aa565b505050916003019190600301906110f4565b506111799291506111c5565b5090565b828054828255905f5260205f209081019282156111b9575f5260205f209182015b828111156111b957825482559160010191906001019061119e565b50611179929150611211565b808211156111795780547fffffffffffffffffffffffff00000000000000000000000000000000000000001681555f600182018190556112086002830182611225565b506003016111c5565b5b80821115611179575f8155600101611212565b50805461123190611750565b5f825580601f10611240575050565b601f0160209004905f5260205f20908101906110af9190611211565b8035801515811461126b575f80fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146110af575f80fd5b5f80604083850312156112a2575f80fd5b6112ab8361125c565b915060208301356112bb81611270565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff81118282101715611316576113166112c6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611363576113636112c6565b604052919050565b5f601f83601f84011261137c575f80fd5b8235602067ffffffffffffffff80831115611399576113996112c6565b8260051b6113a883820161131c565b93845286810183019383810190898611156113c1575f80fd5b84890192505b858310156114b7578235848111156113dd575f80fd5b890160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d038101821315611412575f80fd5b61141a6112f3565b8884013561142781611270565b81526040848101358a830152928401359288841115611444575f80fd5b83850194508e603f860112611457575f80fd5b8985013593508884111561146d5761146d6112c6565b61147c8a848e8701160161131c565b92508383528e81858701011115611491575f80fd5b838186018b8501375f9383018a01939093529182015283525091840191908401906113c7565b9998505050505050505050565b5f602082840312156114d4575f80fd5b813567ffffffffffffffff8111156114ea575f80fd5b6114f68482850161136b565b949350505050565b5f6020828403121561150e575f80fd5b6115178261125c565b9392505050565b5f806040838503121561152f575f80fd5b82359150602083013567ffffffffffffffff81111561154c575f80fd5b6115588582860161136b565b9150509250929050565b5f60208284031215611572575f80fd5b5035919050565b5f81518084525f5b8181101561159d57602081850181015186830182015201611581565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f61160e6060830184611579565b95945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156116bc578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff168452878101518885015286015160608785018190526116a881860183611579565b96890196945050509086019060010161163e565b509098975050505050505050565b602080825282518282018190525f9190848201906040850190845b8181101561171757835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016116e5565b50909695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061176457607f821691505b60208210810361179b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156117e857805f5260205f20601f840160051c810160208510156117c65750805b601f840160051c820191505b818110156117e5575f81556001016117d2565b50505b505050565b815167ffffffffffffffff811115611807576118076112c6565b61181b816118158454611750565b846117a1565b602080601f83116001811461186d575f84156118375750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611901565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156118b95788860151825594840194600190910190840161189a565b50858210156118f557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611919575f80fd5b815161151781611270565b5f815480845260208085019450835f5260205f205f5b8381101561196c57815473ffffffffffffffffffffffffffffffffffffffff168752958201956001918201910161193a565b509495945050505050565b602081525f6115176020830184611924565b604081525f61199b6040830185611924565b90508260208301529392505050565b8181036119b5575050565b6119bf8254611750565b67ffffffffffffffff8111156119d7576119d76112c6565b6119e5816118158454611750565b5f601f821160018114611a35575f83156119ff5750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556117e5565b5f85815260208082208683529082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616925b83811015611a895782860154825560019586019590910190602001611a69565b5085831015611ac557818501547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220409446f504309056dfa4216aa322bf2bf709f1546a20a96dd915225d517f88f964736f6c63430008180033",
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
