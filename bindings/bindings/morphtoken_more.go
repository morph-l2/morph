// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const MorphTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_string_storage\"},{\"astId\":1006,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1009,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1010,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_dailyInflationRates\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_array(t_struct(DailyInflationRate)1017_storage)dyn_storage\"},{\"astId\":1011,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_inflations\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_mapping(t_uint256,t_uint256)\"},{\"astId\":1012,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"_inflationMintedDays\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_uint256\"},{\"astId\":1013,\"contract\":\"contracts/L2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_array(t_uint256)1014_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(DailyInflationRate)1017_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IMorphToken.DailyInflationRate[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[38]\",\"numberOfBytes\":\"1216\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(DailyInflationRate)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IMorphToken.DailyInflationRate\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var MorphTokenStorageLayout = new(solc.StorageLayout)

var MorphTokenDeployedBin = "0x608060405234801561000f575f80fd5b506004361061018f575f3560e01c8063715018a6116100dd578063a457c2d711610088578063cd4281d011610063578063cd4281d0146103af578063dd62ed3e146103d6578063f2fde38b1461041b575f80fd5b8063a457c2d714610381578063a9059cbb14610394578063c553f7b3146103a7575f80fd5b8063944fa746116100b8578063944fa7461461034757806395d89b4114610366578063a29bfb2c1461036e575f80fd5b8063715018a6146102fa578063807de443146103025780638da5cb5b14610329575f80fd5b8063395093511161013d5780635ea94e6f116101185780635ea94e6f146102aa5780636fe0e395146102b257806370a08231146102c5575f80fd5b806339509351146102365780633d9353fe14610249578063405abb4114610295575f80fd5b806318160ddd1161016d57806318160ddd1461020257806323b872dd14610214578063313ce56714610227575f80fd5b806306fdde0314610193578063095ea7b3146101b15780630cb92c13146101d4575b5f80fd5b61019b61042e565b6040516101a89190611603565b60405180910390f35b6101c46101bf366004611695565b6104be565b60405190151581526020016101a8565b6101e76101e23660046116bd565b6104d7565b604080518251815260209283015192810192909252016101a8565b6067545b6040519081526020016101a8565b6101c46102223660046116d4565b61052e565b604051601281526020016101a8565b6101c4610244366004611695565b610551565b6102707f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a8565b6102a86102a336600461170d565b61059c565b005b606c54610206565b6102a86102c0366004611801565b61071e565b6102066102d3366004611870565b73ffffffffffffffffffffffffffffffffffffffff165f9081526068602052604090205490565b6102a8610971565b6102707f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff16610270565b6102066103553660046116bd565b5f908152606b602052604090205490565b61019b610984565b6102a861037c3660046116bd565b610993565b6101c461038f366004611695565b610d2e565b6101c46103a2366004611695565b610dd8565b606a54610206565b6102707f000000000000000000000000000000000000000000000000000000000000000081565b6102066103e4366004611890565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260696020908152604080832093909416825291909152205490565b6102a8610429366004611870565b610de5565b60606065805461043d906118c1565b80601f0160208091040260200160405190810160405280929190818152602001828054610469906118c1565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610e9c565b60019150505b92915050565b604080518082019091525f8082526020820152606a82815481106104fd576104fd611912565b905f5260205f2090600202016040518060400160405290815f82015481526020016001820154815250509050919050565b5f3361053b858285611003565b6105468585856110d9565b506001949350505050565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104cb908290869061059790879061196c565b610e9c565b6105a46112dc565b5f8111610638576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f6566666563746976652064617973206166746572206d7573742062652067726560448201527f61746572207468616e203000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b606a80545f91839161064c9060019061197f565b8154811061065c5761065c611912565b905f5260205f20906002020160010154610676919061196c565b60408051808201825285815260208101838152606a80546001810182555f91825292517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5160029094029384015590517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a52909201919091559051919250829185917fceb8faf32dab6cac28ca92d86325e14aa7016715513ad8b353b7de90fb7f02cf91a3505050565b5f54610100900460ff161580801561073c57505f54600160ff909116105b806107555750303b15801561075557505f5460ff166001145b6107e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161062f565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561083d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61084561135d565b606561085186826119d6565b50606661085e85826119d6565b5061086933846113fb565b6040805180820182528381525f60208201818152606a805460018101825590835292517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a51600290940293840155517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5290920191909155905183907fceb8faf32dab6cac28ca92d86325e14aa7016715513ad8b353b7de90fb7f02cf908390a3801561096a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6109796112dc565b6109825f6114ee565b565b60606066805461043d906118c1565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6f6e6c79207265636f726420636f6e747261637420616c6c6f77656400000000604482015260640161062f565b5f7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633faa50096040518163ffffffff1660e01b8152600401602060405180830381865afa158015610abd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ae19190611af2565b610aeb904261197f565b610af59190611b09565b610b0090600161196c565b9050818111610b91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f746865207370656369666965642074696d6520686173206e6f7420796574206260448201527f65656e2072656163686564000000000000000000000000000000000000000000606482015260840161062f565b606c548111610bfc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f616c6c20696e666c6174696f6e73206d696e7465640000000000000000000000604482015260640161062f565b606c545b828111610d29575f606a5f81548110610c1b57610c1b611912565b5f9182526020822060029091020154606a54909250610c3c9060019061197f565b90505b8015610ca55782606a8281548110610c5957610c59611912565b905f5260205f2090600202016001015411610c9357606a8181548110610c8157610c81611912565b905f5260205f2090600202015f015491505b80610c9d81611b41565b915050610c3f565b50662386f26fc1000081606754610cbc9190611b75565b610cc69190611b09565b5f838152606b60205260409020819055610d01907f0000000000000000000000000000000000000000000000000000000000000000906113fb565b606c8054905f610d1083611b8c565b9190505550508080610d2190611b8c565b915050610c00565b505050565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610dcb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f64656372656173656420616c6c6f77616e63652062656c6f77207a65726f0000604482015260640161062f565b6105468286868403610e9c565b5f336104cb8185856110d9565b610ded6112dc565b73ffffffffffffffffffffffffffffffffffffffff8116610e90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161062f565b610e99816114ee565b50565b73ffffffffffffffffffffffffffffffffffffffff8316610f19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f617070726f76652066726f6d20746865207a65726f2061646472657373000000604482015260640161062f565b73ffffffffffffffffffffffffffffffffffffffff8216610f96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f617070726f766520746f20746865207a65726f20616464726573730000000000604482015260640161062f565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526069602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152606960209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146110d357818110156110c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f696e73756666696369656e7420616c6c6f77616e636500000000000000000000604482015260640161062f565b6110d38484848403610e9c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611156576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7472616e736665722066726f6d20746865207a65726f20616464726573730000604482015260640161062f565b73ffffffffffffffffffffffffffffffffffffffff82166111d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7472616e7366657220746f20746865207a65726f206164647265737300000000604482015260640161062f565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526068602052604090205481811015611262576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f7472616e7366657220616d6f756e7420657863656564732062616c616e636500604482015260640161062f565b73ffffffffffffffffffffffffffffffffffffffff8085165f8181526068602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906112ce9086815260200190565b60405180910390a350505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610982576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161062f565b5f54610100900460ff166113f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161062f565b610982611564565b73ffffffffffffffffffffffffffffffffffffffff8216611478576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6d696e7420746f20746865207a65726f20616464726573730000000000000000604482015260640161062f565b8060675f828254611489919061196c565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152606860209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166115fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161062f565b610982336114ee565b5f602080835283518060208501525f5b8181101561162f57858101830151858201604001528201611613565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611690575f80fd5b919050565b5f80604083850312156116a6575f80fd5b6116af8361166d565b946020939093013593505050565b5f602082840312156116cd575f80fd5b5035919050565b5f805f606084860312156116e6575f80fd5b6116ef8461166d565b92506116fd6020850161166d565b9150604084013590509250925092565b5f806040838503121561171e575f80fd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611769575f80fd5b813567ffffffffffffffff808211156117845761178461172d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156117ca576117ca61172d565b816040528381528660208588010111156117e2575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f8060808587031215611814575f80fd5b843567ffffffffffffffff8082111561182b575f80fd5b6118378883890161175a565b9550602087013591508082111561184c575f80fd5b506118598782880161175a565b949794965050505060408301359260600135919050565b5f60208284031215611880575f80fd5b6118898261166d565b9392505050565b5f80604083850312156118a1575f80fd5b6118aa8361166d565b91506118b86020840161166d565b90509250929050565b600181811c908216806118d557607f821691505b60208210810361190c577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156104d1576104d161193f565b818103818111156104d1576104d161193f565b601f821115610d2957805f5260205f20601f840160051c810160208510156119b75750805b601f840160051c820191505b8181101561096a575f81556001016119c3565b815167ffffffffffffffff8111156119f0576119f061172d565b611a04816119fe84546118c1565b84611992565b602080601f831160018114611a56575f8415611a205750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611aea565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611aa257888601518255948401946001909101908401611a83565b5085821015611ade57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611b02575f80fd5b5051919050565b5f82611b3c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f81611b4f57611b4f61193f565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b80820281158282048414176104d1576104d161193f565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bbc57611bbc61193f565b506001019056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(MorphTokenStorageLayoutJSON), MorphTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["MorphToken"] = MorphTokenStorageLayout
	deployedBytecodes["MorphToken"] = MorphTokenDeployedBin
}