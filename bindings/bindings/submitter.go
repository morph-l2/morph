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
	Bin: "0x61016060405234801561001157600080fd5b50604051611e05380380611e058339810160408190526100309161009b565b6001608052600060a081905260c05273530000000000000000000000000000000000000760e0526001600160a01b03166101005273530000000000000000000000000000000000000361012052735300000000000000000000000000000000000004610140526100cb565b6000602082840312156100ad57600080fd5b81516001600160a01b03811681146100c457600080fd5b9392505050565b60805160a05160c05160e051610100516101205161014051611c7e610187600039600081816101dc0152818161062d01528181610bb701528181610cda01528181610ee3015261124e0152600081816102d90152818161059d01528181610b4201528181610d9301528181610e6e01526111c401526000818161018b01526107110152600081816102840152818161034c015281816106e7015261074801526000610aab01526000610a8201526000610a590152611c7e6000f3fe608060405234801561001057600080fd5b50600436106101815760003560e01c8063843e8a7b116100d8578063bddd8e731161008c578063cc0f858f11610066578063cc0f858f1461044b578063e8e3992514610454578063fe4b84df146104d957600080fd5b8063bddd8e73146103f4578063c58159c4146103fc578063c6b61e4c1461040557600080fd5b8063965fbb94116100bd578063965fbb941461036e578063a5af40d114610381578063bc0bc6ba146103a957600080fd5b8063843e8a7b1461030d578063927ede2d1461034757600080fd5b80633cb747bf1161013a5780636cb23707116101145780636cb23707146102d457806373790ab3146102fb5780637e4fa7001461030457600080fd5b80633cb747bf1461028257806354fd4d50146102a85780635c14c314146102bd57600080fd5b8063151232581161016b57806315123258146101fe57806316e2994a1461025a57806322caba241461026f57600080fd5b80628dbdb514610186578063047d0b6e146101d7575b600080fd5b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b61021161020c3660046117f0565b6104ec565b6040516101ce9190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b61026d6102683660046118b2565b610585565b005b61026d61027d366004611951565b6106cf565b7f00000000000000000000000000000000000000000000000000000000000000006101ad565b6102b0610a52565b6040516101ce91906119bd565b6102c660025481565b6040519081526020016101ce565b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6102c660045481565b6102c660015481565b610315610af5565b6040805173ffffffffffffffffffffffffffffffffffffffff90941684526020840192909252908201526060016101ce565b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b61026d61037c3660046117f0565b610cc2565b61039461038f3660046119f0565b610e23565b604080519283526020830191909152016101ce565b6103bc6103b73660046117f0565b611109565b60408051825173ffffffffffffffffffffffffffffffffffffffff1681526020808401519082015291810151908201526060016101ce565b61026d611192565b6102c660055481565b6103156104133660046117f0565b60076020526000908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116919083565b6102c660065481565b6104a26104623660046117f0565b6003602081905260009182526040909120805460018201546002830154929093015473ffffffffffffffffffffffffffffffffffffffff90911692919084565b6040805173ffffffffffffffffffffffffffffffffffffffff909516855260208501939093529183015260608201526080016101ce565b61026d6104e73660046117f0565b6112e7565b61052d6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b506000908152600360208181526040928390208351608081018552815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281015493820193909352910154606082015290565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c79206c322073657175656e63657220636f6e747261637400000000000060448201526064015b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ba9190611a14565b600060055590506106cb81836114e0565b5050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107ed57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d59190611a2d565b73ffffffffffffffffffffffffffffffffffffffff16145b610879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f526f6c6c75703a2066756e6374696f6e2063616e206f6e6c792062652063616c60448201527f6c65642066726f6d20746865204c3120726f6c6c7570000000000000000000006064820152608401610620565b60015485146108e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206261746368496e64657800000000000000000000000000006044820152606401610620565b600254831461094f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617463685374617274426c6f636b0000000000000000006044820152606401610620565b604080516080808201835273ffffffffffffffffffffffffffffffffffffffff878116808452602080850189815285870189815260608088018a815260008f81526003808752908b902099518a547fffffffffffffffffffffffff000000000000000000000000000000000000000016981697909717895592516001890155905160028801559051959093019490945584518a8152938401529282018690529181018490529081018290527f516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f9060a00160405180910390a160018054906000610a3783611a79565b90915550610a489050826001611ab1565b6002555050505050565b6060610a7d7f000000000000000000000000000000000000000000000000000000000000000061164f565b610aa67f000000000000000000000000000000000000000000000000000000000000000061164f565b610acf7f000000000000000000000000000000000000000000000000000000000000000061164f565b604051602001610ae193929190611ac4565b604051602081830303815290604052905090565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081526000600482018190529081908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e90602401600060405180830381865afa158015610b89573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610bb19190810190611b3a565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c449190611a14565b825160045460055492935090915b42610c5d8584611ab1565b11610c8b5780610c6c81611a79565b915050828103610c7a575060005b610c848483611ab1565b9150610c52565b848181518110610c9d57610c9d611bc9565b6020026020010151828584610cb29190611ab1565b9750975097505050505050909192565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f6e6c7920676f7620636f6e74726163740000000000000000000000000000006044820152606401610620565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e90602401600060405180830381865afa158015610def573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e179190810190611b3a565b90506106cb82826114e0565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e90602401600060405180830381865afa158015610eb5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610edd9190810190611b3a565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f709190611a14565b82519091506000805b82811015610fe857848181518110610f9357610f93611bc9565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1603610fd65760019150610fe8565b80610fe081611a79565b915050610f79565b8161104f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c6964207375626d69747465720000000000000000000000000000006044820152606401610620565b6004546005545b426110618784611ab1565b1161108f578061107081611a79565b91505084810361107e575060005b6110888683611ab1565b9150611056565b808311156110cc576000866110a48386611bf8565b6110ae9190611c0b565b9050806110bb8882611ab1565b995099505050505050505050915091565b808310156110ec57600086846110e28489611bf8565b6110a49190611ab1565b6004546110f98782611ab1565b9850985050505050505050915091565b6111436040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b506000908152600760209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600201549181019190915290565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e90602401600060405180830381865afa158015611220573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112489190810190611b3a565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112db9190611a14565b90506106cb81836114e0565b600054610100900460ff16158080156113075750600054600160ff909116105b806113215750303b158015611321575060005460ff166001145b6113ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610620565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561140b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60008211611475576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420666972737445706f636853746172740000000000000000006044820152606401610620565b600482905580156106cb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b80515b42836004546114f29190611ab1565b11611611576006805490600061150783611a79565b91905055506040518060600160405280836005548151811061152b5761152b611bc9565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016004548152602001846004546115649190611ab1565b90526006546000908152600760209081526040808320845181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178155918401516001830155929092015160029092019190915560058054916115e083611a79565b919050555080600554036115f45760006005555b82600460008282546116069190611ab1565b909155506114e39050565b60408051848152602081018390527fabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579910160405180910390a1505050565b6060600061165c8361170d565b600101905060008167ffffffffffffffff81111561167c5761167c611809565b6040519080825280601f01601f1916602001820160405280156116a6576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846116b057509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611756577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611782576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106117a057662386f26fc10000830492506010015b6305f5e10083106117b8576305f5e100830492506008015b61271083106117cc57612710830492506004015b606483106117de576064830492506002015b600a83106117ea576001015b92915050565b60006020828403121561180257600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561186157611861611809565b604052919050565b600067ffffffffffffffff82111561188357611883611809565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff811681146118af57600080fd5b50565b600060208083850312156118c557600080fd5b823567ffffffffffffffff8111156118dc57600080fd5b8301601f810185136118ed57600080fd5b80356119006118fb82611869565b611838565b81815260059190911b8201830190838101908783111561191f57600080fd5b928401925b828410156119465783356119378161188d565b82529284019290840190611924565b979650505050505050565b600080600080600060a0868803121561196957600080fd5b85359450602086013561197b8161188d565b94979496505050506040830135926060810135926080909101359150565b60005b838110156119b457818101518382015260200161199c565b50506000910152565b60208152600082518060208401526119dc816040850160208701611999565b601f01601f19169190910160400192915050565b600060208284031215611a0257600080fd5b8135611a0d8161188d565b9392505050565b600060208284031215611a2657600080fd5b5051919050565b600060208284031215611a3f57600080fd5b8151611a0d8161188d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611aaa57611aaa611a4a565b5060010190565b808201808211156117ea576117ea611a4a565b60008451611ad6818460208901611999565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611b12816001850160208a01611999565b60019201918201528351611b2d816002840160208801611999565b0160020195945050505050565b60006020808385031215611b4d57600080fd5b825167ffffffffffffffff811115611b6457600080fd5b8301601f81018513611b7557600080fd5b8051611b836118fb82611869565b81815260059190911b82018301908381019087831115611ba257600080fd5b928401925b82841015611946578351611bba8161188d565b82529284019290840190611ba7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b818103818111156117ea576117ea611a4a565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c4357611c43611a4a565b50029056fea2646970667358221220862fb82b4e8ed6ff3a93e47af63f90817ada92443f60643571e047623d1c0bf564736f6c63430008100033",
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
