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

// TypesBatchInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesBatchInfo struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}

// TypesEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesEpochInfo struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}

// SubmitterMetaData contains all meta data concerning the Submitter contract.
var SubmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_rollup\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ACKRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequencersLen\",\"type\":\"uint256\"}],\"name\":\"EpochUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COUNTERPART\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_GOV_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"name\":\"ackRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatedEpochIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmedBatchs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"epochUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"getConfirmedBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BatchInfo\",\"name\":\"batchInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EpochInfo\",\"name\":\"epochInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"getTurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextEpochStart\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEpochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSubmitterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"}],\"name\":\"sequencersUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateEpochExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610160604052348015610010575f80fd5b50604051611d61380380611d6183398101604081905261002f91610099565b60016080525f60a081905260c05273530000000000000000000000000000000000000760e0526001600160a01b03166101005273530000000000000000000000000000000000000361012052735300000000000000000000000000000000000004610140526100c6565b5f602082840312156100a9575f80fd5b81516001600160a01b03811681146100bf575f80fd5b9392505050565b60805160a05160c05160e051610100516101205161014051611be86101795f395f81816101d20152818161061b01528181610b9801528181610cb801528181610eb6015261121101525f81816102cf0152818161058c01528181610b2801528181610d7001528181610e46015261118c01525f818161018101526106fc01525f818161027a01528181610342015281816106d2015261073301525f610a9201525f610a6901525f610a400152611be85ff3fe608060405234801561000f575f80fd5b5060043610610178575f3560e01c8063843e8a7b116100d2578063bddd8e7311610088578063cc0f858f11610063578063cc0f858f14610440578063e8e3992514610449578063fe4b84df146104cd575f80fd5b8063bddd8e73146103ea578063c58159c4146103f2578063c6b61e4c146103fb575f80fd5b8063965fbb94116100b8578063965fbb9414610364578063a5af40d114610377578063bc0bc6ba1461039f575f80fd5b8063843e8a7b14610303578063927ede2d1461033d575f80fd5b80633cb747bf116101325780636cb237071161010d5780636cb23707146102ca57806373790ab3146102f15780637e4fa700146102fa575f80fd5b80633cb747bf1461027857806354fd4d501461029e5780635c14c314146102b3575f80fd5b8063151232581161016257806315123258146101f457806316e2994a1461025057806322caba2414610265575f80fd5b80628dbdb51461017c578063047d0b6e146101cd575b5f80fd5b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b6102076102023660046117a3565b6104e0565b6040516101c49190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b61026361025e36600461185f565b610574565b005b6102636102733660046118f9565b6106ba565b7f00000000000000000000000000000000000000000000000000000000000000006101a3565b6102a6610a39565b6040516101c4919061195f565b6102bc60025481565b6040519081526020016101c4565b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b6102bc60045481565b6102bc60015481565b61030b610adc565b6040805173ffffffffffffffffffffffffffffffffffffffff90941684526020840192909252908201526060016101c4565b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b6102636103723660046117a3565b610ca0565b61038a610385366004611991565b610dfc565b604080519283526020830191909152016101c4565b6103b26103ad3660046117a3565b6110d6565b60408051825173ffffffffffffffffffffffffffffffffffffffff1681526020808401519082015291810151908201526060016101c4565b61026361115b565b6102bc60055481565b61030b6104093660046117a3565b60076020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116919083565b6102bc60065481565b6104966104573660046117a3565b600360208190525f9182526040909120805460018201546002830154929093015473ffffffffffffffffffffffffffffffffffffffff90911692919084565b6040805173ffffffffffffffffffffffffffffffffffffffff909516855260208501939093529183015260608201526080016101c4565b6102636104db3660046117a3565b6112a8565b61051d60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81525090565b505f908152600360208181526040928390208351608081018552815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281015493820193909352910154606082015290565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610618576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c79206c322073657175656e63657220636f6e747261637400000000000060448201526064015b60405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610682573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a691906119b3565b5f60055590506106b6818361149a565b5050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107d657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561079a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107be91906119ca565b73ffffffffffffffffffffffffffffffffffffffff16145b610862576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f526f6c6c75703a2066756e6374696f6e2063616e206f6e6c792062652063616c60448201527f6c65642066726f6d20746865204c3120726f6c6c757000000000000000000000606482015260840161060f565b60015485146108cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206261746368496e6465780000000000000000000000000000604482015260640161060f565b6002548314610938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617463685374617274426c6f636b000000000000000000604482015260640161060f565b604080516080808201835273ffffffffffffffffffffffffffffffffffffffff878116808452602080850189815285870189815260608088018a81525f8f81526003808752908b902099518a547fffffffffffffffffffffffff000000000000000000000000000000000000000016981697909717895592516001890155905160028801559051959093019490945584518a8152938401529282018690529181018490529081018290527f516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f9060a00160405180910390a160018054905f610a1e83611a12565b90915550610a2f9050826001611a49565b6002555050505050565b6060610a647f0000000000000000000000000000000000000000000000000000000000000000611605565b610a8d7f0000000000000000000000000000000000000000000000000000000000000000611605565b610ab67f0000000000000000000000000000000000000000000000000000000000000000611605565b604051602001610ac893929190611a5c565b604051602081830303815290604052905090565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081525f600482018190529081908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e906024015f60405180830381865afa158015610b6c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b939190810190611ad1565b90505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bff573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c2391906119b3565b825160045460055492935090915b42610c3c8584611a49565b11610c695780610c4b81611a12565b915050828103610c5857505f5b610c628483611a49565b9150610c31565b848181518110610c7b57610c7b611b5b565b6020026020010151828584610c909190611a49565b9750975097505050505050909192565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f6e6c7920676f7620636f6e7472616374000000000000000000000000000000604482015260640161060f565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081525f60048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e906024015f60405180830381865afa158015610dc9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610df09190810190611ad1565b90506106b6828261149a565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081525f60048201819052908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e906024015f60405180830381865afa158015610e8a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610eb19190810190611ad1565b90505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f1d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f4191906119b3565b82519091505f805b82811015610fb857848181518110610f6357610f63611b5b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1603610fa65760019150610fb8565b80610fb081611a12565b915050610f49565b8161101f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c6964207375626d6974746572000000000000000000000000000000604482015260640161060f565b6004546005545b426110318784611a49565b1161105e578061104081611a12565b91505084810361104d57505f5b6110578683611a49565b9150611026565b8083111561109a575f866110728386611b88565b61107c9190611b9b565b9050806110898882611a49565b995099505050505050505050915091565b808310156110b9575f86846110af8489611b88565b6110729190611a49565b6004546110c68782611a49565b9850985050505050505050915091565b61110d60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b505f908152600760209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600201549181019190915290565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081525f60048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e906024015f60405180830381865afa1580156111e5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261120c9190810190611ad1565b90505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015611278573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061129c91906119b3565b90506106b6818361149a565b5f54610100900460ff16158080156112c657505f54600160ff909116105b806112df5750303b1580156112df57505f5460ff166001145b61136b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161060f565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156113c7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f8211611430576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420666972737445706f63685374617274000000000000000000604482015260640161060f565b600482905580156106b6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b80515b42836004546114ac9190611a49565b116115c75760068054905f6114c083611a12565b9190505550604051806060016040528083600554815181106114e4576114e4611b5b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16815260200160045481526020018460045461151d9190611a49565b90526006545f908152600760209081526040808320845181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161781559184015160018301559290920151600290920191909155600580549161159883611a12565b919050555080600554036115ab575f6005555b8260045f8282546115bc9190611a49565b9091555061149d9050565b60408051848152602081018390527fabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579910160405180910390a1505050565b60605f611611836116c1565b60010190505f8167ffffffffffffffff811115611630576116306117ba565b6040519080825280601f01601f19166020018201604052801561165a576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461166457509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611709577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611735576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061175357662386f26fc10000830492506010015b6305f5e100831061176b576305f5e100830492506008015b612710831061177f57612710830492506004015b60648310611791576064830492506002015b600a831061179d576001015b92915050565b5f602082840312156117b3575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611810576118106117ba565b604052919050565b5f67ffffffffffffffff821115611831576118316117ba565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461185c575f80fd5b50565b5f6020808385031215611870575f80fd5b823567ffffffffffffffff811115611886575f80fd5b8301601f81018513611896575f80fd5b80356118a96118a482611818565b6117e7565b81815260059190911b820183019083810190878311156118c7575f80fd5b928401925b828410156118ee5783356118df8161183b565b825292840192908401906118cc565b979650505050505050565b5f805f805f60a0868803121561190d575f80fd5b85359450602086013561191f8161183b565b94979496505050506040830135926060810135926080909101359150565b5f5b8381101561195757818101518382015260200161193f565b50505f910152565b602081525f825180602084015261197d81604085016020870161193d565b601f01601f19169190910160400192915050565b5f602082840312156119a1575f80fd5b81356119ac8161183b565b9392505050565b5f602082840312156119c3575f80fd5b5051919050565b5f602082840312156119da575f80fd5b81516119ac8161183b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a4257611a426119e5565b5060010190565b8082018082111561179d5761179d6119e5565b5f8451611a6d81846020890161193d565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611aa9816001850160208a0161193d565b60019201918201528351611ac481600284016020880161193d565b0160020195945050505050565b5f6020808385031215611ae2575f80fd5b825167ffffffffffffffff811115611af8575f80fd5b8301601f81018513611b08575f80fd5b8051611b166118a482611818565b81815260059190911b82018301908381019087831115611b34575f80fd5b928401925b828410156118ee578351611b4c8161183b565b82529284019290840190611b39565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8181038181111561179d5761179d6119e5565b808202811582820484141761179d5761179d6119e556fea26469706673582212204138a4c492825487e02c32e14dce05b36097c4d651d90f90052984e75116e56d64736f6c63430008180033",
}

// SubmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use SubmitterMetaData.ABI instead.
var SubmitterABI = SubmitterMetaData.ABI

// SubmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubmitterMetaData.Bin instead.
var SubmitterBin = SubmitterMetaData.Bin

// DeploySubmitter deploys a new Ethereum contract, binding an instance of Submitter to it.
func DeploySubmitter(auth *bind.TransactOpts, backend bind.ContractBackend, _rollup common.Address) (common.Address, *types.Transaction, *Submitter, error) {
	parsed, err := SubmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubmitterBin), backend, _rollup)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// Submitter is an auto generated Go binding around an Ethereum contract.
type Submitter struct {
	SubmitterCaller     // Read-only binding to the contract
	SubmitterTransactor // Write-only binding to the contract
	SubmitterFilterer   // Log filterer for contract events
}

// SubmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmitterSession struct {
	Contract     *Submitter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmitterCallerSession struct {
	Contract *SubmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SubmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmitterTransactorSession struct {
	Contract     *SubmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SubmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmitterRaw struct {
	Contract *Submitter // Generic contract binding to access the raw methods on
}

// SubmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmitterCallerRaw struct {
	Contract *SubmitterCaller // Generic read-only contract binding to access the raw methods on
}

// SubmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmitterTransactorRaw struct {
	Contract *SubmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmitter creates a new instance of Submitter, bound to a specific deployed contract.
func NewSubmitter(address common.Address, backend bind.ContractBackend) (*Submitter, error) {
	contract, err := bindSubmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// NewSubmitterCaller creates a new read-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterCaller(address common.Address, caller bind.ContractCaller) (*SubmitterCaller, error) {
	contract, err := bindSubmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterCaller{contract: contract}, nil
}

// NewSubmitterTransactor creates a new write-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmitterTransactor, error) {
	contract, err := bindSubmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterTransactor{contract: contract}, nil
}

// NewSubmitterFilterer creates a new log filterer instance of Submitter, bound to a specific deployed contract.
func NewSubmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmitterFilterer, error) {
	contract, err := bindSubmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmitterFilterer{contract: contract}, nil
}

// bindSubmitter binds a generic wrapper to an already deployed contract.
func bindSubmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.SubmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transact(opts, method, params...)
}

// COUNTERPART is a free data retrieval call binding the contract method 0x008dbdb5.
//
// Solidity: function COUNTERPART() view returns(address)
func (_Submitter *SubmitterCaller) COUNTERPART(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "COUNTERPART")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COUNTERPART is a free data retrieval call binding the contract method 0x008dbdb5.
//
// Solidity: function COUNTERPART() view returns(address)
func (_Submitter *SubmitterSession) COUNTERPART() (common.Address, error) {
	return _Submitter.Contract.COUNTERPART(&_Submitter.CallOpts)
}

// COUNTERPART is a free data retrieval call binding the contract method 0x008dbdb5.
//
// Solidity: function COUNTERPART() view returns(address)
func (_Submitter *SubmitterCallerSession) COUNTERPART() (common.Address, error) {
	return _Submitter.Contract.COUNTERPART(&_Submitter.CallOpts)
}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterCaller) L2GOVCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "L2_GOV_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterSession) L2GOVCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2GOVCONTRACT(&_Submitter.CallOpts)
}

// L2GOVCONTRACT is a free data retrieval call binding the contract method 0x047d0b6e.
//
// Solidity: function L2_GOV_CONTRACT() view returns(address)
func (_Submitter *SubmitterCallerSession) L2GOVCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2GOVCONTRACT(&_Submitter.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterCaller) L2SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "L2_SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2SEQUENCERCONTRACT(&_Submitter.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Submitter *SubmitterCallerSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Submitter.Contract.L2SEQUENCERCONTRACT(&_Submitter.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Submitter *SubmitterCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Submitter *SubmitterSession) MESSENGER() (common.Address, error) {
	return _Submitter.Contract.MESSENGER(&_Submitter.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Submitter *SubmitterCallerSession) MESSENGER() (common.Address, error) {
	return _Submitter.Contract.MESSENGER(&_Submitter.CallOpts)
}

// CalculatedEpochIndex is a free data retrieval call binding the contract method 0xcc0f858f.
//
// Solidity: function calculatedEpochIndex() view returns(uint256)
func (_Submitter *SubmitterCaller) CalculatedEpochIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "calculatedEpochIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatedEpochIndex is a free data retrieval call binding the contract method 0xcc0f858f.
//
// Solidity: function calculatedEpochIndex() view returns(uint256)
func (_Submitter *SubmitterSession) CalculatedEpochIndex() (*big.Int, error) {
	return _Submitter.Contract.CalculatedEpochIndex(&_Submitter.CallOpts)
}

// CalculatedEpochIndex is a free data retrieval call binding the contract method 0xcc0f858f.
//
// Solidity: function calculatedEpochIndex() view returns(uint256)
func (_Submitter *SubmitterCallerSession) CalculatedEpochIndex() (*big.Int, error) {
	return _Submitter.Contract.CalculatedEpochIndex(&_Submitter.CallOpts)
}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterCaller) ConfirmedBatchs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "confirmedBatchs", arg0)

	outstruct := new(struct {
		Submitter  common.Address
		StartBlock *big.Int
		EndBlock   *big.Int
		RollupTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RollupTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterSession) ConfirmedBatchs(arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Submitter.Contract.ConfirmedBatchs(&_Submitter.CallOpts, arg0)
}

// ConfirmedBatchs is a free data retrieval call binding the contract method 0xe8e39925.
//
// Solidity: function confirmedBatchs(uint256 ) view returns(address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Submitter *SubmitterCallerSession) ConfirmedBatchs(arg0 *big.Int) (struct {
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Submitter.Contract.ConfirmedBatchs(&_Submitter.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "epochs", arg0)

	outstruct := new(struct {
		Submitter common.Address
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterSession) Epochs(arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Submitter.Contract.Epochs(&_Submitter.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(address submitter, uint256 startTime, uint256 endTime)
func (_Submitter *SubmitterCallerSession) Epochs(arg0 *big.Int) (struct {
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Submitter.Contract.Epochs(&_Submitter.CallOpts, arg0)
}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterCaller) GetConfirmedBatch(opts *bind.CallOpts, batchIndex *big.Int) (TypesBatchInfo, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getConfirmedBatch", batchIndex)

	if err != nil {
		return *new(TypesBatchInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBatchInfo)).(*TypesBatchInfo)

	return out0, err

}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterSession) GetConfirmedBatch(batchIndex *big.Int) (TypesBatchInfo, error) {
	return _Submitter.Contract.GetConfirmedBatch(&_Submitter.CallOpts, batchIndex)
}

// GetConfirmedBatch is a free data retrieval call binding the contract method 0x15123258.
//
// Solidity: function getConfirmedBatch(uint256 batchIndex) view returns((address,uint256,uint256,uint256) batchInfo)
func (_Submitter *SubmitterCallerSession) GetConfirmedBatch(batchIndex *big.Int) (TypesBatchInfo, error) {
	return _Submitter.Contract.GetConfirmedBatch(&_Submitter.CallOpts, batchIndex)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterCaller) GetEpoch(opts *bind.CallOpts, epochIndex *big.Int) (TypesEpochInfo, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getEpoch", epochIndex)

	if err != nil {
		return *new(TypesEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesEpochInfo)).(*TypesEpochInfo)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterSession) GetEpoch(epochIndex *big.Int) (TypesEpochInfo, error) {
	return _Submitter.Contract.GetEpoch(&_Submitter.CallOpts, epochIndex)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 epochIndex) view returns((address,uint256,uint256) epochInfo)
func (_Submitter *SubmitterCallerSession) GetEpoch(epochIndex *big.Int) (TypesEpochInfo, error) {
	return _Submitter.Contract.GetEpoch(&_Submitter.CallOpts, epochIndex)
}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCaller) GetNextSubmitter(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getNextSubmitter")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterSession) GetNextSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetNextSubmitter(&_Submitter.CallOpts)
}

// GetNextSubmitter is a free data retrieval call binding the contract method 0x843e8a7b.
//
// Solidity: function getNextSubmitter() view returns(address, uint256, uint256)
func (_Submitter *SubmitterCallerSession) GetNextSubmitter() (common.Address, *big.Int, *big.Int, error) {
	return _Submitter.Contract.GetNextSubmitter(&_Submitter.CallOpts)
}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterCaller) GetTurn(opts *bind.CallOpts, submitter common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "getTurn", submitter)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterSession) GetTurn(submitter common.Address) (*big.Int, *big.Int, error) {
	return _Submitter.Contract.GetTurn(&_Submitter.CallOpts, submitter)
}

// GetTurn is a free data retrieval call binding the contract method 0xa5af40d1.
//
// Solidity: function getTurn(address submitter) view returns(uint256, uint256)
func (_Submitter *SubmitterCallerSession) GetTurn(submitter common.Address) (*big.Int, *big.Int, error) {
	return _Submitter.Contract.GetTurn(&_Submitter.CallOpts, submitter)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Submitter *SubmitterCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Submitter *SubmitterSession) Messenger() (common.Address, error) {
	return _Submitter.Contract.Messenger(&_Submitter.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Submitter *SubmitterCallerSession) Messenger() (common.Address, error) {
	return _Submitter.Contract.Messenger(&_Submitter.CallOpts)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterCaller) NextBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterSession) NextBatchIndex() (*big.Int, error) {
	return _Submitter.Contract.NextBatchIndex(&_Submitter.CallOpts)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextBatchIndex() (*big.Int, error) {
	return _Submitter.Contract.NextBatchIndex(&_Submitter.CallOpts)
}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterCaller) NextBatchStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextBatchStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterSession) NextBatchStartBlock() (*big.Int, error) {
	return _Submitter.Contract.NextBatchStartBlock(&_Submitter.CallOpts)
}

// NextBatchStartBlock is a free data retrieval call binding the contract method 0x5c14c314.
//
// Solidity: function nextBatchStartBlock() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextBatchStartBlock() (*big.Int, error) {
	return _Submitter.Contract.NextBatchStartBlock(&_Submitter.CallOpts)
}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint256)
func (_Submitter *SubmitterCaller) NextEpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextEpochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint256)
func (_Submitter *SubmitterSession) NextEpochStart() (*big.Int, error) {
	return _Submitter.Contract.NextEpochStart(&_Submitter.CallOpts)
}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextEpochStart() (*big.Int, error) {
	return _Submitter.Contract.NextEpochStart(&_Submitter.CallOpts)
}

// NextSubmitterIndex is a free data retrieval call binding the contract method 0xc58159c4.
//
// Solidity: function nextSubmitterIndex() view returns(uint256)
func (_Submitter *SubmitterCaller) NextSubmitterIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "nextSubmitterIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSubmitterIndex is a free data retrieval call binding the contract method 0xc58159c4.
//
// Solidity: function nextSubmitterIndex() view returns(uint256)
func (_Submitter *SubmitterSession) NextSubmitterIndex() (*big.Int, error) {
	return _Submitter.Contract.NextSubmitterIndex(&_Submitter.CallOpts)
}

// NextSubmitterIndex is a free data retrieval call binding the contract method 0xc58159c4.
//
// Solidity: function nextSubmitterIndex() view returns(uint256)
func (_Submitter *SubmitterCallerSession) NextSubmitterIndex() (*big.Int, error) {
	return _Submitter.Contract.NextSubmitterIndex(&_Submitter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Submitter *SubmitterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Submitter *SubmitterSession) Version() (string, error) {
	return _Submitter.Contract.Version(&_Submitter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Submitter *SubmitterCallerSession) Version() (string, error) {
	return _Submitter.Contract.Version(&_Submitter.CallOpts)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterTransactor) AckRollup(opts *bind.TransactOpts, batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "ackRollup", batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterSession) AckRollup(batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.AckRollup(&_Submitter.TransactOpts, batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// AckRollup is a paid mutator transaction binding the contract method 0x22caba24.
//
// Solidity: function ackRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime) returns()
func (_Submitter *SubmitterTransactorSession) AckRollup(batchIndex *big.Int, submitter common.Address, batchStartBlock *big.Int, batchEndBlock *big.Int, rollupTime *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.AckRollup(&_Submitter.TransactOpts, batchIndex, submitter, batchStartBlock, batchEndBlock, rollupTime)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterTransactor) EpochUpdated(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "epochUpdated", epoch)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterSession) EpochUpdated(epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.EpochUpdated(&_Submitter.TransactOpts, epoch)
}

// EpochUpdated is a paid mutator transaction binding the contract method 0x965fbb94.
//
// Solidity: function epochUpdated(uint256 epoch) returns()
func (_Submitter *SubmitterTransactorSession) EpochUpdated(epoch *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.EpochUpdated(&_Submitter.TransactOpts, epoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _nextEpochStart) returns()
func (_Submitter *SubmitterTransactor) Initialize(opts *bind.TransactOpts, _nextEpochStart *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "initialize", _nextEpochStart)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _nextEpochStart) returns()
func (_Submitter *SubmitterSession) Initialize(_nextEpochStart *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts, _nextEpochStart)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _nextEpochStart) returns()
func (_Submitter *SubmitterTransactorSession) Initialize(_nextEpochStart *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts, _nextEpochStart)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterTransactor) SequencersUpdated(opts *bind.TransactOpts, sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "sequencersUpdated", sequencers)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterSession) SequencersUpdated(sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.SequencersUpdated(&_Submitter.TransactOpts, sequencers)
}

// SequencersUpdated is a paid mutator transaction binding the contract method 0x16e2994a.
//
// Solidity: function sequencersUpdated(address[] sequencers) returns()
func (_Submitter *SubmitterTransactorSession) SequencersUpdated(sequencers []common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.SequencersUpdated(&_Submitter.TransactOpts, sequencers)
}

// UpdateEpochExternal is a paid mutator transaction binding the contract method 0xbddd8e73.
//
// Solidity: function updateEpochExternal() returns()
func (_Submitter *SubmitterTransactor) UpdateEpochExternal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "updateEpochExternal")
}

// UpdateEpochExternal is a paid mutator transaction binding the contract method 0xbddd8e73.
//
// Solidity: function updateEpochExternal() returns()
func (_Submitter *SubmitterSession) UpdateEpochExternal() (*types.Transaction, error) {
	return _Submitter.Contract.UpdateEpochExternal(&_Submitter.TransactOpts)
}

// UpdateEpochExternal is a paid mutator transaction binding the contract method 0xbddd8e73.
//
// Solidity: function updateEpochExternal() returns()
func (_Submitter *SubmitterTransactorSession) UpdateEpochExternal() (*types.Transaction, error) {
	return _Submitter.Contract.UpdateEpochExternal(&_Submitter.TransactOpts)
}

// SubmitterACKRollupIterator is returned from FilterACKRollup and is used to iterate over the raw logs and unpacked data for ACKRollup events raised by the Submitter contract.
type SubmitterACKRollupIterator struct {
	Event *SubmitterACKRollup // Event containing the contract specifics and raw log

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
func (it *SubmitterACKRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterACKRollup)
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
		it.Event = new(SubmitterACKRollup)
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
func (it *SubmitterACKRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterACKRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterACKRollup represents a ACKRollup event raised by the Submitter contract.
type SubmitterACKRollup struct {
	BatchIndex      *big.Int
	Submitter       common.Address
	BatchStartBlock *big.Int
	BatchEndBlock   *big.Int
	RollupTime      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterACKRollup is a free log retrieval operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) FilterACKRollup(opts *bind.FilterOpts) (*SubmitterACKRollupIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "ACKRollup")
	if err != nil {
		return nil, err
	}
	return &SubmitterACKRollupIterator{contract: _Submitter.contract, event: "ACKRollup", logs: logs, sub: sub}, nil
}

// WatchACKRollup is a free log subscription operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) WatchACKRollup(opts *bind.WatchOpts, sink chan<- *SubmitterACKRollup) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "ACKRollup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterACKRollup)
				if err := _Submitter.contract.UnpackLog(event, "ACKRollup", log); err != nil {
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

// ParseACKRollup is a log parse operation binding the contract event 0x516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f.
//
// Solidity: event ACKRollup(uint256 batchIndex, address submitter, uint256 batchStartBlock, uint256 batchEndBlock, uint256 rollupTime)
func (_Submitter *SubmitterFilterer) ParseACKRollup(log types.Log) (*SubmitterACKRollup, error) {
	event := new(SubmitterACKRollup)
	if err := _Submitter.contract.UnpackLog(event, "ACKRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterEpochUpdatedIterator is returned from FilterEpochUpdated and is used to iterate over the raw logs and unpacked data for EpochUpdated events raised by the Submitter contract.
type SubmitterEpochUpdatedIterator struct {
	Event *SubmitterEpochUpdated // Event containing the contract specifics and raw log

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
func (it *SubmitterEpochUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterEpochUpdated)
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
		it.Event = new(SubmitterEpochUpdated)
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
func (it *SubmitterEpochUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterEpochUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterEpochUpdated represents a EpochUpdated event raised by the Submitter contract.
type SubmitterEpochUpdated struct {
	Interval      *big.Int
	SequencersLen *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEpochUpdated is a free log retrieval operation binding the contract event 0xabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579.
//
// Solidity: event EpochUpdated(uint256 interval, uint256 sequencersLen)
func (_Submitter *SubmitterFilterer) FilterEpochUpdated(opts *bind.FilterOpts) (*SubmitterEpochUpdatedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "EpochUpdated")
	if err != nil {
		return nil, err
	}
	return &SubmitterEpochUpdatedIterator{contract: _Submitter.contract, event: "EpochUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochUpdated is a free log subscription operation binding the contract event 0xabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579.
//
// Solidity: event EpochUpdated(uint256 interval, uint256 sequencersLen)
func (_Submitter *SubmitterFilterer) WatchEpochUpdated(opts *bind.WatchOpts, sink chan<- *SubmitterEpochUpdated) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "EpochUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterEpochUpdated)
				if err := _Submitter.contract.UnpackLog(event, "EpochUpdated", log); err != nil {
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

// ParseEpochUpdated is a log parse operation binding the contract event 0xabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579.
//
// Solidity: event EpochUpdated(uint256 interval, uint256 sequencersLen)
func (_Submitter *SubmitterFilterer) ParseEpochUpdated(log types.Log) (*SubmitterEpochUpdated, error) {
	event := new(SubmitterEpochUpdated)
	if err := _Submitter.contract.UnpackLog(event, "EpochUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Submitter contract.
type SubmitterInitializedIterator struct {
	Event *SubmitterInitialized // Event containing the contract specifics and raw log

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
func (it *SubmitterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterInitialized)
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
		it.Event = new(SubmitterInitialized)
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
func (it *SubmitterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterInitialized represents a Initialized event raised by the Submitter contract.
type SubmitterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubmitterInitializedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubmitterInitializedIterator{contract: _Submitter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubmitterInitialized) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterInitialized)
				if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Submitter *SubmitterFilterer) ParseInitialized(log types.Log) (*SubmitterInitialized, error) {
	event := new(SubmitterInitialized)
	if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
