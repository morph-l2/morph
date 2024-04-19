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

// IRecordBatchSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRecordBatchSubmission struct {
	Index      *big.Int
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}

// IRecordRewardEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type IRecordRewardEpochInfo struct {
	Index               *big.Int
	BlockCount          *big.Int
	Sequencers          []common.Address
	SequencerBlocks     []*big.Int
	SequencerRatios     []*big.Int
	SequencerComissions []*big.Int
}

// IRecordRollupEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type IRecordRollupEpochInfo struct {
	Index     *big.Int
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}

// RecordMetaData contains all meta data concerning the Record contract.
var RecordMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getBatchSubmissions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.BatchSubmission[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getRewardEpochs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerRatios\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerComissions\",\"type\":\"uint256[]\"}],\"internalType\":\"structIRecord.RewardEpochInfo[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getRollupEpochs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.RollupEpochInfo[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchSubmissionIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRewardEpochIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRollupEpochIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.BatchSubmission[]\",\"name\":\"_batchSubmissions\",\"type\":\"tuple[]\"}],\"name\":\"recordFinalizedBatchSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerRatios\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerComissions\",\"type\":\"uint256[]\"}],\"internalType\":\"structIRecord.RewardEpochInfo[]\",\"name\":\"_rewardEpochs\",\"type\":\"tuple[]\"}],\"name\":\"recordRewardEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.RollupEpochInfo[]\",\"name\":\"_rollupEpochs\",\"type\":\"tuple[]\"}],\"name\":\"recordRollupEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardpEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61014060405261271061012052348015610017575f80fd5b5073530000000000000000000000000000000000001060805273530000000000000000000000000000000000001260a05273530000000000000000000000000000000000000360c05273530000000000000000000000000000000000001160e0527353000000000000000000000000000000000000046101005260805160a05160c05160e05161010051610120516126b36100f45f395f6117b601525f61030b01525f818161023201526119c801525f61040701525f81816103c20152610fe301525f81816104ce015281816110d8015261155501526126b35ff3fe608060405234801561000f575f80fd5b5060043610610184575f3560e01c806364b4abe3116100dd5780638da5cb5b11610088578063cb6293e811610063578063cb6293e8146104a9578063d5577141146104c9578063f2fde38b146104f0575f80fd5b80638da5cb5b146103e45780638e21d5fb14610402578063a24231e814610429575f80fd5b80637828a905116100b85780637828a9051461030657806378f908e11461032d578063807de443146103bd575f80fd5b806364b4abe3146102cb5780636ea0396e146102eb578063715018a6146102fe575f80fd5b8063484f8d0f1161013d5780634c69c00f116101185780634c69c00f146102855780634e3ca406146102985780634ecff524146102b8575f80fd5b8063484f8d0f14610254578063485cc9551461025d5780634897504714610272575f80fd5b80632fbf64871161016d5780632fbf6487146101df57806338013f02146101e85780633d9353fe1461022d575f80fd5b80631511e1b1146101885780632b15806b146101a4575b5f80fd5b61019160695481565b6040519081526020015b60405180910390f35b6101ca6101b2366004611fcb565b60686020525f90815260409020805460019091015482565b6040805192835260208301919091520161019b565b610191606b5481565b6065546102089073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6102087f000000000000000000000000000000000000000000000000000000000000000081565b610191606a5481565b61027061026b36600461200a565b610503565b005b61027061028036600461203b565b610720565b6102706102933660046120aa565b610952565b6102ab6102a63660046120ca565b610a04565b60405161019b91906120ea565b6102706102c6366004612163565b610b8c565b6102de6102d93660046120ca565b610de9565b60405161019b91906121c0565b6102706102f9366004612236565b610f7a565b610270611ab5565b6102087f000000000000000000000000000000000000000000000000000000000000000081565b61038061033b366004611fcb565b60666020525f908152604090208054600182015460028301546003840154600490940154929373ffffffffffffffffffffffffffffffffffffffff9092169290919085565b6040805195865273ffffffffffffffffffffffffffffffffffffffff9094166020860152928401919091526060830152608082015260a00161019b565b6102087f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff16610208565b6102087f000000000000000000000000000000000000000000000000000000000000000081565b610474610437366004611fcb565b60676020525f90815260409020805460018201546002830154600390930154919273ffffffffffffffffffffffffffffffffffffffff9091169184565b6040805194855273ffffffffffffffffffffffffffffffffffffffff909316602085015291830152606082015260800161019b565b6104bc6104b73660046120ca565b611ac8565b60405161019b91906122cd565b6102087f000000000000000000000000000000000000000000000000000000000000000081565b6102706104fe3660046120aa565b611d7c565b5f54610100900460ff161580801561052157505f54600160ff909116105b8061053a5750303b15801561053a57505f5460ff166001145b6105b15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561060d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff82166106705760405162461bcd60e51b815260206004820152601660248201527f696e76616c6964206f7261636c6520616464726573730000000000000000000060448201526064016105a8565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556106b983611e19565b801561071b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60655473ffffffffffffffffffffffffffffffffffffffff1633146107875760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f7765640000000000000000000000000060448201526064016105a8565b5f5b8181101561071b5780606a5461079f9190612428565b8383838181106107b1576107b1612441565b9050608002015f0135146108075760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e6465780000000000000000000000000000000000000060448201526064016105a8565b604051806080016040528084848481811061082457610824612441565b9050608002015f0135815260200184848481811061084457610844612441565b905060800201602001602081019061085c91906120aa565b73ffffffffffffffffffffffffffffffffffffffff16815260200184848481811061088957610889612441565b9050608002016040013581526020018484848181106108aa576108aa612441565b9050608002016060013581525060675f8585858181106108cc576108cc612441565b608002919091013582525060208082019290925260409081015f208351815591830151600180840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915590830151600283015560609092015160039091015501610789565b61095a611e8f565b73ffffffffffffffffffffffffffffffffffffffff81166109bd5760405162461bcd60e51b815260206004820152601660248201527f696e76616c6964206f7261636c6520616464726573730000000000000000000060448201526064016105a8565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b606082821015610a565760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e6465780000000000000000000000000000000000000060448201526064016105a8565b610a60838361246e565b610a6b906001612428565b67ffffffffffffffff811115610a8357610a83612481565b604051908082528060200260200182016040528015610af157816020015b610ade60405180608001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b815260200190600190039081610aa15790505b509050825b828111610b85575f81815260676020908152604091829020825160808101845281548152600182015473ffffffffffffffffffffffffffffffffffffffff169281019290925260028101549282019290925260039091015460608201528251839083908110610b6757610b67612441565b60200260200101819052508080610b7d906124ae565b915050610af6565b5092915050565b60655473ffffffffffffffffffffffffffffffffffffffff163314610bf35760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f7765640000000000000000000000000060448201526064016105a8565b5f5b8181101561071b5780606954610c0b9190612428565b838383818110610c1d57610c1d612441565b905060a002015f013514610c735760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e6465780000000000000000000000000000000000000060448201526064016105a8565b6040518060a00160405280848484818110610c9057610c90612441565b905060a002015f01358152602001848484818110610cb057610cb0612441565b905060a002016020016020810190610cc891906120aa565b73ffffffffffffffffffffffffffffffffffffffff168152602001848484818110610cf557610cf5612441565b905060a00201604001358152602001848484818110610d1657610d16612441565b905060a00201606001358152602001848484818110610d3757610d37612441565b905060a002016080013581525060665f858585818110610d5957610d59612441565b60a002919091013582525060208082019290925260409081015f208351815591830151600180840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091559083015160028301556060830151600383015560809092015160049091015501610bf5565b606082821015610e3b5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e6465780000000000000000000000000000000000000060448201526064016105a8565b610e45838361246e565b610e50906001612428565b67ffffffffffffffff811115610e6857610e68612481565b604051908082528060200260200182016040528015610edc57816020015b610ec96040518060a001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81525090565b815260200190600190039081610e865790505b509050825b828111610b85575f81815260666020908152604091829020825160a08101845281548152600182015473ffffffffffffffffffffffffffffffffffffffff16928101929092526002810154928201929092526003820154606082015260049091015460808201528251839083908110610f5c57610f5c612441565b60200260200101819052508080610f72906124ae565b915050610ee1565b60655473ffffffffffffffffffffffffffffffffffffffff163314610fe15760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f7765640000000000000000000000000060448201526064016105a8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663766718086040518163ffffffff1660e01b8152600401602060405180830381865afa15801561104a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106e91906124e5565b606b5460019061107f908490612428565b611089919061246e565b106110d65760405162461bcd60e51b815260206004820152601e60248201527f66757475726520646174612063616e6e6f742062652075706c6f61646564000060448201526064016105a8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a29bfb2c600184849050606b546111259190612428565b61112f919061246e565b6040518263ffffffff1660e01b815260040161114d91815260200190565b5f604051808303815f87803b158015611164575f80fd5b505af1158015611176573d5f803e3d5ffd5b505050505f5b81811015611a97575f83838381811061119757611197612441565b90506020028101906111a991906124fc565b6111b7906040810190612538565b905090505f8484848181106111ce576111ce612441565b90506020028101906111e091906124fc565b606b54903591506111f2908490612428565b81146112405760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642065706f636820696e6465780000000000000000000000000060448201526064016105a8565b8185858581811061125357611253612441565b905060200281019061126591906124fc565b611273906060810190612538565b90501480156112b457508185858581811061129057611290612441565b90506020028101906112a291906124fc565b6112b0906080810190612538565b9050145b80156112f25750818585858181106112ce576112ce612441565b90506020028101906112e091906124fc565b6112ee9060a0810190612538565b9050145b61133e5760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642064617461206c656e6774680000000000000000000000000060448201526064016105a8565b6040518060c0016040528082815260200186868681811061136157611361612441565b905060200281019061137391906124fc565b6020013581526020018367ffffffffffffffff81111561139557611395612481565b6040519080825280602002602001820160405280156113be578160200160208202803683370190505b5081526020018367ffffffffffffffff8111156113dd576113dd612481565b604051908082528060200260200182016040528015611406578160200160208202803683370190505b5081526020018367ffffffffffffffff81111561142557611425612481565b60405190808252806020026020018201604052801561144e578160200160208202803683370190505b5081526020018367ffffffffffffffff81111561146d5761146d612481565b604051908082528060200260200182016040528015611496578160200160208202803683370190505b5090525f82815260686020908152604091829020835181558382015160018201559183015180516114cd9260028501920190611ef6565b50606082015180516114e9916003840191602090910190611f7e565b5060808201518051611505916004840191602090910190611f7e565b5060a08201518051611521916005840191602090910190611f7e565b50506040517fd68d0781000000000000000000000000000000000000000000000000000000008152600481018590525f91507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d68d0781906024016020604051808303815f875af11580156115b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115d491906124e5565b90505f805f8567ffffffffffffffff8111156115f2576115f2612481565b60405190808252806020026020018201604052801561161b578160200160208202803683370190505b5090505f8667ffffffffffffffff81111561163857611638612481565b604051908082528060200260200182016040528015611661578160200160208202803683370190505b5090505f5b878110156118fe5760148b8b8b81811061168257611682612441565b905060200281019061169491906124fc565b6116a29060a0810190612538565b838181106116b2576116b2612441565b9050602002013511156117075760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642073657175656e6365727320636f6d697373696f6e0000000060448201526064016105a8565b8a8a8a81811061171957611719612441565b905060200281019061172b91906124fc565b611739906080810190612538565b8281811061174957611749612441565b905060200201358461175b9190612428565b93508a8a8a81811061176f5761176f612441565b905060200281019061178191906124fc565b61178f906060810190612538565b8281811061179f5761179f612441565b90506020020135856117b19190612428565b94505f7f00000000000000000000000000000000000000000000000000000000000000008c8c8c8181106117e7576117e7612441565b90506020028101906117f991906124fc565b611807906080810190612538565b8481811061181757611817612441565b905060200201358861182991906125a3565b61183391906125ba565b905060648c8c8c81811061184957611849612441565b905060200281019061185b91906124fc565b6118699060a0810190612538565b8481811061187957611879612441565b905060200201358261188b91906125a3565b61189591906125ba565b838b815181106118a7576118a7612441565b602002602001018181525050828a815181106118c5576118c5612441565b6020026020010151816118d8919061246e565b848b815181106118ea576118ea612441565b602090810291909101015250600101611666565b5089898981811061191157611911612441565b905060200281019061192391906124fc565b6020013584146119755760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642073657175656e6365727320626c6f636b730000000000000060448201526064016105a8565b60648311156119c65760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642073657175656e6365727320726174696f730000000000000060448201526064016105a8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cdd0c50e878c8c8c818110611a1557611a15612441565b9050602002810190611a2791906124fc565b611a35906040810190612538565b86866040518663ffffffff1660e01b8152600401611a579594939291906125f2565b5f604051808303815f87803b158015611a6e575f80fd5b505af1158015611a80573d5f803e3d5ffd5b50506001909901985061117c975050505050505050565b5081819050606b5f828254611aac9190612428565b90915550505050565b611abd611e8f565b611ac65f611e19565b565b606082821015611b1a5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e6465780000000000000000000000000000000000000060448201526064016105a8565b611b24838361246e565b611b2f906001612428565b67ffffffffffffffff811115611b4757611b47612481565b604051908082528060200260200182016040528015611baf57816020015b611b9c6040518060c001604052805f81526020015f8152602001606081526020016060815260200160608152602001606081525090565b815260200190600190039081611b655790505b509050825b828111610b85575f81815260686020908152604091829020825160c0810184528154815260018201548184015260028201805485518186028101860187528181529295939493860193830182828015611c4157602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611c16575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015611c9757602002820191905f5260205f20905b815481526020019060010190808311611c83575b5050505050815260200160048201805480602002602001604051908101604052809291908181526020018280548015611ced57602002820191905f5260205f20905b815481526020019060010190808311611cd9575b5050505050815260200160058201805480602002602001604051908101604052809291908181526020018280548015611d4357602002820191905f5260205f20905b815481526020019060010190808311611d2f575b505050505081525050828281518110611d5e57611d5e612441565b60200260200101819052508080611d74906124ae565b915050611bb4565b611d84611e8f565b73ffffffffffffffffffffffffffffffffffffffff8116611e0d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105a8565b611e1681611e19565b50565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60335473ffffffffffffffffffffffffffffffffffffffff163314611ac65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105a8565b828054828255905f5260205f20908101928215611f6e579160200282015b82811115611f6e57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190611f14565b50611f7a929150611fb7565b5090565b828054828255905f5260205f20908101928215611f6e579160200282015b82811115611f6e578251825591602001919060010190611f9c565b5b80821115611f7a575f8155600101611fb8565b5f60208284031215611fdb575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612005575f80fd5b919050565b5f806040838503121561201b575f80fd5b61202483611fe2565b915061203260208401611fe2565b90509250929050565b5f806020838503121561204c575f80fd5b823567ffffffffffffffff80821115612063575f80fd5b818501915085601f830112612076575f80fd5b813581811115612084575f80fd5b8660208260071b8501011115612098575f80fd5b60209290920196919550909350505050565b5f602082840312156120ba575f80fd5b6120c382611fe2565b9392505050565b5f80604083850312156120db575f80fd5b50508035926020909101359150565b602080825282518282018190525f919060409081850190868401855b82811015612156578151805185528681015173ffffffffffffffffffffffffffffffffffffffff168786015285810151868601526060908101519085015260809093019290850190600101612106565b5091979650505050505050565b5f8060208385031215612174575f80fd5b823567ffffffffffffffff8082111561218b575f80fd5b818501915085601f83011261219e575f80fd5b8135818111156121ac575f80fd5b86602060a083028501011115612098575f80fd5b602080825282518282018190525f919060409081850190868401855b82811015612156578151805185528681015173ffffffffffffffffffffffffffffffffffffffff16878601528581015186860152606080820151908601526080908101519085015260a090930192908501906001016121dc565b5f8060208385031215612247575f80fd5b823567ffffffffffffffff8082111561225e575f80fd5b818501915085601f830112612271575f80fd5b81358181111561227f575f80fd5b8660208260051b8501011115612098575f80fd5b5f815180845260208085019450602084015f5b838110156122c2578151875295820195908201906001016122a6565b509495945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156123ed578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051845287810151888501528681015160c0888601819052815190860181905260e08601918a01905f905b8082101561238657825173ffffffffffffffffffffffffffffffffffffffff168452928b0192918b019160019190910190612350565b505050606080830151868303828801526123a08382612293565b92505050608080830151868303828801526123bb8382612293565b9250505060a080830151925085820381870152506123d98183612293565b9689019694505050908601906001016122f4565b509098975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561243b5761243b6123fb565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8181038181111561243b5761243b6123fb565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036124de576124de6123fb565b5060010190565b5f602082840312156124f5575f80fd5b5051919050565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4183360301811261252e575f80fd5b9190910192915050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261256b575f80fd5b83018035915067ffffffffffffffff821115612585575f80fd5b6020019150600581901b360382131561259c575f80fd5b9250929050565b808202811582820484141761243b5761243b6123fb565b5f826125ed577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b858152608060208083018290529082018590525f90869060a08401835b888110156126485773ffffffffffffffffffffffffffffffffffffffff61263585611fe2565b168252928201929082019060010161260f565b50848103604086015261265b8188612293565b9250505082810360608401526126718185612293565b9897505050505050505056fea26469706673582212208d61352fa41bcbd9d68df87d32931ce0b549bff8ab28a4db075aca6daecbbde864736f6c63430008180033",
}

// RecordABI is the input ABI used to generate the binding from.
// Deprecated: Use RecordMetaData.ABI instead.
var RecordABI = RecordMetaData.ABI

// RecordBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecordMetaData.Bin instead.
var RecordBin = RecordMetaData.Bin

// DeployRecord deploys a new Ethereum contract, binding an instance of Record to it.
func DeployRecord(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Record, error) {
	parsed, err := RecordMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecordBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Record{RecordCaller: RecordCaller{contract: contract}, RecordTransactor: RecordTransactor{contract: contract}, RecordFilterer: RecordFilterer{contract: contract}}, nil
}

// Record is an auto generated Go binding around an Ethereum contract.
type Record struct {
	RecordCaller     // Read-only binding to the contract
	RecordTransactor // Write-only binding to the contract
	RecordFilterer   // Log filterer for contract events
}

// RecordCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecordSession struct {
	Contract     *Record           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecordCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecordCallerSession struct {
	Contract *RecordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RecordTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecordTransactorSession struct {
	Contract     *RecordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecordRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecordRaw struct {
	Contract *Record // Generic contract binding to access the raw methods on
}

// RecordCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecordCallerRaw struct {
	Contract *RecordCaller // Generic read-only contract binding to access the raw methods on
}

// RecordTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecordTransactorRaw struct {
	Contract *RecordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecord creates a new instance of Record, bound to a specific deployed contract.
func NewRecord(address common.Address, backend bind.ContractBackend) (*Record, error) {
	contract, err := bindRecord(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Record{RecordCaller: RecordCaller{contract: contract}, RecordTransactor: RecordTransactor{contract: contract}, RecordFilterer: RecordFilterer{contract: contract}}, nil
}

// NewRecordCaller creates a new read-only instance of Record, bound to a specific deployed contract.
func NewRecordCaller(address common.Address, caller bind.ContractCaller) (*RecordCaller, error) {
	contract, err := bindRecord(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecordCaller{contract: contract}, nil
}

// NewRecordTransactor creates a new write-only instance of Record, bound to a specific deployed contract.
func NewRecordTransactor(address common.Address, transactor bind.ContractTransactor) (*RecordTransactor, error) {
	contract, err := bindRecord(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecordTransactor{contract: contract}, nil
}

// NewRecordFilterer creates a new log filterer instance of Record, bound to a specific deployed contract.
func NewRecordFilterer(address common.Address, filterer bind.ContractFilterer) (*RecordFilterer, error) {
	contract, err := bindRecord(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecordFilterer{contract: contract}, nil
}

// bindRecord binds a generic wrapper to an already deployed contract.
func bindRecord(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecordABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Record *RecordRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Record.Contract.RecordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Record *RecordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Record.Contract.RecordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Record *RecordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Record.Contract.RecordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Record *RecordCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Record.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Record *RecordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Record.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Record *RecordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Record.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_Record *RecordCaller) DISTRIBUTECONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "DISTRIBUTE_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_Record *RecordSession) DISTRIBUTECONTRACT() (common.Address, error) {
	return _Record.Contract.DISTRIBUTECONTRACT(&_Record.CallOpts)
}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_Record *RecordCallerSession) DISTRIBUTECONTRACT() (common.Address, error) {
	return _Record.Contract.DISTRIBUTECONTRACT(&_Record.CallOpts)
}

// GOVCONTRACT is a free data retrieval call binding the contract method 0x7828a905.
//
// Solidity: function GOV_CONTRACT() view returns(address)
func (_Record *RecordCaller) GOVCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "GOV_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVCONTRACT is a free data retrieval call binding the contract method 0x7828a905.
//
// Solidity: function GOV_CONTRACT() view returns(address)
func (_Record *RecordSession) GOVCONTRACT() (common.Address, error) {
	return _Record.Contract.GOVCONTRACT(&_Record.CallOpts)
}

// GOVCONTRACT is a free data retrieval call binding the contract method 0x7828a905.
//
// Solidity: function GOV_CONTRACT() view returns(address)
func (_Record *RecordCallerSession) GOVCONTRACT() (common.Address, error) {
	return _Record.Contract.GOVCONTRACT(&_Record.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Record *RecordCaller) L2STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "L2_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Record *RecordSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Record.Contract.L2STAKINGCONTRACT(&_Record.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Record *RecordCallerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Record.Contract.L2STAKINGCONTRACT(&_Record.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Record *RecordCaller) MORPHTOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "MORPH_TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Record *RecordSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _Record.Contract.MORPHTOKENCONTRACT(&_Record.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Record *RecordCallerSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _Record.Contract.MORPHTOKENCONTRACT(&_Record.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Record *RecordCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Record *RecordSession) ORACLE() (common.Address, error) {
	return _Record.Contract.ORACLE(&_Record.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Record *RecordCallerSession) ORACLE() (common.Address, error) {
	return _Record.Contract.ORACLE(&_Record.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Record *RecordCaller) SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Record *RecordSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _Record.Contract.SEQUENCERCONTRACT(&_Record.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_Record *RecordCallerSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _Record.Contract.SEQUENCERCONTRACT(&_Record.CallOpts)
}

// BatchSubmissions is a free data retrieval call binding the contract method 0x78f908e1.
//
// Solidity: function batchSubmissions(uint256 ) view returns(uint256 index, address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Record *RecordCaller) BatchSubmissions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index      *big.Int
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "batchSubmissions", arg0)

	outstruct := new(struct {
		Index      *big.Int
		Submitter  common.Address
		StartBlock *big.Int
		EndBlock   *big.Int
		RollupTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Submitter = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RollupTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BatchSubmissions is a free data retrieval call binding the contract method 0x78f908e1.
//
// Solidity: function batchSubmissions(uint256 ) view returns(uint256 index, address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Record *RecordSession) BatchSubmissions(arg0 *big.Int) (struct {
	Index      *big.Int
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Record.Contract.BatchSubmissions(&_Record.CallOpts, arg0)
}

// BatchSubmissions is a free data retrieval call binding the contract method 0x78f908e1.
//
// Solidity: function batchSubmissions(uint256 ) view returns(uint256 index, address submitter, uint256 startBlock, uint256 endBlock, uint256 rollupTime)
func (_Record *RecordCallerSession) BatchSubmissions(arg0 *big.Int) (struct {
	Index      *big.Int
	Submitter  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	RollupTime *big.Int
}, error) {
	return _Record.Contract.BatchSubmissions(&_Record.CallOpts, arg0)
}

// GetBatchSubmissions is a free data retrieval call binding the contract method 0x64b4abe3.
//
// Solidity: function getBatchSubmissions(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256)[] res)
func (_Record *RecordCaller) GetBatchSubmissions(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]IRecordBatchSubmission, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "getBatchSubmissions", start, end)

	if err != nil {
		return *new([]IRecordBatchSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRecordBatchSubmission)).(*[]IRecordBatchSubmission)

	return out0, err

}

// GetBatchSubmissions is a free data retrieval call binding the contract method 0x64b4abe3.
//
// Solidity: function getBatchSubmissions(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256)[] res)
func (_Record *RecordSession) GetBatchSubmissions(start *big.Int, end *big.Int) ([]IRecordBatchSubmission, error) {
	return _Record.Contract.GetBatchSubmissions(&_Record.CallOpts, start, end)
}

// GetBatchSubmissions is a free data retrieval call binding the contract method 0x64b4abe3.
//
// Solidity: function getBatchSubmissions(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256)[] res)
func (_Record *RecordCallerSession) GetBatchSubmissions(start *big.Int, end *big.Int) ([]IRecordBatchSubmission, error) {
	return _Record.Contract.GetBatchSubmissions(&_Record.CallOpts, start, end)
}

// GetRewardEpochs is a free data retrieval call binding the contract method 0xcb6293e8.
//
// Solidity: function getRewardEpochs(uint256 start, uint256 end) view returns((uint256,uint256,address[],uint256[],uint256[],uint256[])[] res)
func (_Record *RecordCaller) GetRewardEpochs(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]IRecordRewardEpochInfo, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "getRewardEpochs", start, end)

	if err != nil {
		return *new([]IRecordRewardEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRecordRewardEpochInfo)).(*[]IRecordRewardEpochInfo)

	return out0, err

}

// GetRewardEpochs is a free data retrieval call binding the contract method 0xcb6293e8.
//
// Solidity: function getRewardEpochs(uint256 start, uint256 end) view returns((uint256,uint256,address[],uint256[],uint256[],uint256[])[] res)
func (_Record *RecordSession) GetRewardEpochs(start *big.Int, end *big.Int) ([]IRecordRewardEpochInfo, error) {
	return _Record.Contract.GetRewardEpochs(&_Record.CallOpts, start, end)
}

// GetRewardEpochs is a free data retrieval call binding the contract method 0xcb6293e8.
//
// Solidity: function getRewardEpochs(uint256 start, uint256 end) view returns((uint256,uint256,address[],uint256[],uint256[],uint256[])[] res)
func (_Record *RecordCallerSession) GetRewardEpochs(start *big.Int, end *big.Int) ([]IRecordRewardEpochInfo, error) {
	return _Record.Contract.GetRewardEpochs(&_Record.CallOpts, start, end)
}

// GetRollupEpochs is a free data retrieval call binding the contract method 0x4e3ca406.
//
// Solidity: function getRollupEpochs(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256)[] res)
func (_Record *RecordCaller) GetRollupEpochs(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]IRecordRollupEpochInfo, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "getRollupEpochs", start, end)

	if err != nil {
		return *new([]IRecordRollupEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRecordRollupEpochInfo)).(*[]IRecordRollupEpochInfo)

	return out0, err

}

// GetRollupEpochs is a free data retrieval call binding the contract method 0x4e3ca406.
//
// Solidity: function getRollupEpochs(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256)[] res)
func (_Record *RecordSession) GetRollupEpochs(start *big.Int, end *big.Int) ([]IRecordRollupEpochInfo, error) {
	return _Record.Contract.GetRollupEpochs(&_Record.CallOpts, start, end)
}

// GetRollupEpochs is a free data retrieval call binding the contract method 0x4e3ca406.
//
// Solidity: function getRollupEpochs(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256)[] res)
func (_Record *RecordCallerSession) GetRollupEpochs(start *big.Int, end *big.Int) ([]IRecordRollupEpochInfo, error) {
	return _Record.Contract.GetRollupEpochs(&_Record.CallOpts, start, end)
}

// NextBatchSubmissionIndex is a free data retrieval call binding the contract method 0x1511e1b1.
//
// Solidity: function nextBatchSubmissionIndex() view returns(uint256)
func (_Record *RecordCaller) NextBatchSubmissionIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "nextBatchSubmissionIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBatchSubmissionIndex is a free data retrieval call binding the contract method 0x1511e1b1.
//
// Solidity: function nextBatchSubmissionIndex() view returns(uint256)
func (_Record *RecordSession) NextBatchSubmissionIndex() (*big.Int, error) {
	return _Record.Contract.NextBatchSubmissionIndex(&_Record.CallOpts)
}

// NextBatchSubmissionIndex is a free data retrieval call binding the contract method 0x1511e1b1.
//
// Solidity: function nextBatchSubmissionIndex() view returns(uint256)
func (_Record *RecordCallerSession) NextBatchSubmissionIndex() (*big.Int, error) {
	return _Record.Contract.NextBatchSubmissionIndex(&_Record.CallOpts)
}

// NextRewardEpochIndex is a free data retrieval call binding the contract method 0x2fbf6487.
//
// Solidity: function nextRewardEpochIndex() view returns(uint256)
func (_Record *RecordCaller) NextRewardEpochIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "nextRewardEpochIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRewardEpochIndex is a free data retrieval call binding the contract method 0x2fbf6487.
//
// Solidity: function nextRewardEpochIndex() view returns(uint256)
func (_Record *RecordSession) NextRewardEpochIndex() (*big.Int, error) {
	return _Record.Contract.NextRewardEpochIndex(&_Record.CallOpts)
}

// NextRewardEpochIndex is a free data retrieval call binding the contract method 0x2fbf6487.
//
// Solidity: function nextRewardEpochIndex() view returns(uint256)
func (_Record *RecordCallerSession) NextRewardEpochIndex() (*big.Int, error) {
	return _Record.Contract.NextRewardEpochIndex(&_Record.CallOpts)
}

// NextRollupEpochIndex is a free data retrieval call binding the contract method 0x484f8d0f.
//
// Solidity: function nextRollupEpochIndex() view returns(uint256)
func (_Record *RecordCaller) NextRollupEpochIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "nextRollupEpochIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRollupEpochIndex is a free data retrieval call binding the contract method 0x484f8d0f.
//
// Solidity: function nextRollupEpochIndex() view returns(uint256)
func (_Record *RecordSession) NextRollupEpochIndex() (*big.Int, error) {
	return _Record.Contract.NextRollupEpochIndex(&_Record.CallOpts)
}

// NextRollupEpochIndex is a free data retrieval call binding the contract method 0x484f8d0f.
//
// Solidity: function nextRollupEpochIndex() view returns(uint256)
func (_Record *RecordCallerSession) NextRollupEpochIndex() (*big.Int, error) {
	return _Record.Contract.NextRollupEpochIndex(&_Record.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Record *RecordCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Record *RecordSession) Owner() (common.Address, error) {
	return _Record.Contract.Owner(&_Record.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Record *RecordCallerSession) Owner() (common.Address, error) {
	return _Record.Contract.Owner(&_Record.CallOpts)
}

// RewardpEpochs is a free data retrieval call binding the contract method 0x2b15806b.
//
// Solidity: function rewardpEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordCaller) RewardpEpochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "rewardpEpochs", arg0)

	outstruct := new(struct {
		Index      *big.Int
		BlockCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardpEpochs is a free data retrieval call binding the contract method 0x2b15806b.
//
// Solidity: function rewardpEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordSession) RewardpEpochs(arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	return _Record.Contract.RewardpEpochs(&_Record.CallOpts, arg0)
}

// RewardpEpochs is a free data retrieval call binding the contract method 0x2b15806b.
//
// Solidity: function rewardpEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordCallerSession) RewardpEpochs(arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	return _Record.Contract.RewardpEpochs(&_Record.CallOpts, arg0)
}

// RollupEpochs is a free data retrieval call binding the contract method 0xa24231e8.
//
// Solidity: function rollupEpochs(uint256 ) view returns(uint256 index, address submitter, uint256 startTime, uint256 endTime)
func (_Record *RecordCaller) RollupEpochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index     *big.Int
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "rollupEpochs", arg0)

	outstruct := new(struct {
		Index     *big.Int
		Submitter common.Address
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Submitter = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RollupEpochs is a free data retrieval call binding the contract method 0xa24231e8.
//
// Solidity: function rollupEpochs(uint256 ) view returns(uint256 index, address submitter, uint256 startTime, uint256 endTime)
func (_Record *RecordSession) RollupEpochs(arg0 *big.Int) (struct {
	Index     *big.Int
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Record.Contract.RollupEpochs(&_Record.CallOpts, arg0)
}

// RollupEpochs is a free data retrieval call binding the contract method 0xa24231e8.
//
// Solidity: function rollupEpochs(uint256 ) view returns(uint256 index, address submitter, uint256 startTime, uint256 endTime)
func (_Record *RecordCallerSession) RollupEpochs(arg0 *big.Int) (struct {
	Index     *big.Int
	Submitter common.Address
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _Record.Contract.RollupEpochs(&_Record.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _oracle) returns()
func (_Record *RecordTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "initialize", _admin, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _oracle) returns()
func (_Record *RecordSession) Initialize(_admin common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Record.Contract.Initialize(&_Record.TransactOpts, _admin, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _oracle) returns()
func (_Record *RecordTransactorSession) Initialize(_admin common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Record.Contract.Initialize(&_Record.TransactOpts, _admin, _oracle)
}

// RecordFinalizedBatchSubmissions is a paid mutator transaction binding the contract method 0x4ecff524.
//
// Solidity: function recordFinalizedBatchSubmissions((uint256,address,uint256,uint256,uint256)[] _batchSubmissions) returns()
func (_Record *RecordTransactor) RecordFinalizedBatchSubmissions(opts *bind.TransactOpts, _batchSubmissions []IRecordBatchSubmission) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "recordFinalizedBatchSubmissions", _batchSubmissions)
}

// RecordFinalizedBatchSubmissions is a paid mutator transaction binding the contract method 0x4ecff524.
//
// Solidity: function recordFinalizedBatchSubmissions((uint256,address,uint256,uint256,uint256)[] _batchSubmissions) returns()
func (_Record *RecordSession) RecordFinalizedBatchSubmissions(_batchSubmissions []IRecordBatchSubmission) (*types.Transaction, error) {
	return _Record.Contract.RecordFinalizedBatchSubmissions(&_Record.TransactOpts, _batchSubmissions)
}

// RecordFinalizedBatchSubmissions is a paid mutator transaction binding the contract method 0x4ecff524.
//
// Solidity: function recordFinalizedBatchSubmissions((uint256,address,uint256,uint256,uint256)[] _batchSubmissions) returns()
func (_Record *RecordTransactorSession) RecordFinalizedBatchSubmissions(_batchSubmissions []IRecordBatchSubmission) (*types.Transaction, error) {
	return _Record.Contract.RecordFinalizedBatchSubmissions(&_Record.TransactOpts, _batchSubmissions)
}

// RecordRewardEpochs is a paid mutator transaction binding the contract method 0x6ea0396e.
//
// Solidity: function recordRewardEpochs((uint256,uint256,address[],uint256[],uint256[],uint256[])[] _rewardEpochs) returns()
func (_Record *RecordTransactor) RecordRewardEpochs(opts *bind.TransactOpts, _rewardEpochs []IRecordRewardEpochInfo) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "recordRewardEpochs", _rewardEpochs)
}

// RecordRewardEpochs is a paid mutator transaction binding the contract method 0x6ea0396e.
//
// Solidity: function recordRewardEpochs((uint256,uint256,address[],uint256[],uint256[],uint256[])[] _rewardEpochs) returns()
func (_Record *RecordSession) RecordRewardEpochs(_rewardEpochs []IRecordRewardEpochInfo) (*types.Transaction, error) {
	return _Record.Contract.RecordRewardEpochs(&_Record.TransactOpts, _rewardEpochs)
}

// RecordRewardEpochs is a paid mutator transaction binding the contract method 0x6ea0396e.
//
// Solidity: function recordRewardEpochs((uint256,uint256,address[],uint256[],uint256[],uint256[])[] _rewardEpochs) returns()
func (_Record *RecordTransactorSession) RecordRewardEpochs(_rewardEpochs []IRecordRewardEpochInfo) (*types.Transaction, error) {
	return _Record.Contract.RecordRewardEpochs(&_Record.TransactOpts, _rewardEpochs)
}

// RecordRollupEpochs is a paid mutator transaction binding the contract method 0x48975047.
//
// Solidity: function recordRollupEpochs((uint256,address,uint256,uint256)[] _rollupEpochs) returns()
func (_Record *RecordTransactor) RecordRollupEpochs(opts *bind.TransactOpts, _rollupEpochs []IRecordRollupEpochInfo) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "recordRollupEpochs", _rollupEpochs)
}

// RecordRollupEpochs is a paid mutator transaction binding the contract method 0x48975047.
//
// Solidity: function recordRollupEpochs((uint256,address,uint256,uint256)[] _rollupEpochs) returns()
func (_Record *RecordSession) RecordRollupEpochs(_rollupEpochs []IRecordRollupEpochInfo) (*types.Transaction, error) {
	return _Record.Contract.RecordRollupEpochs(&_Record.TransactOpts, _rollupEpochs)
}

// RecordRollupEpochs is a paid mutator transaction binding the contract method 0x48975047.
//
// Solidity: function recordRollupEpochs((uint256,address,uint256,uint256)[] _rollupEpochs) returns()
func (_Record *RecordTransactorSession) RecordRollupEpochs(_rollupEpochs []IRecordRollupEpochInfo) (*types.Transaction, error) {
	return _Record.Contract.RecordRollupEpochs(&_Record.TransactOpts, _rollupEpochs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Record *RecordTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Record *RecordSession) RenounceOwnership() (*types.Transaction, error) {
	return _Record.Contract.RenounceOwnership(&_Record.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Record *RecordTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Record.Contract.RenounceOwnership(&_Record.TransactOpts)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _oracle) returns()
func (_Record *RecordTransactor) SetOracleAddress(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "setOracleAddress", _oracle)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _oracle) returns()
func (_Record *RecordSession) SetOracleAddress(_oracle common.Address) (*types.Transaction, error) {
	return _Record.Contract.SetOracleAddress(&_Record.TransactOpts, _oracle)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _oracle) returns()
func (_Record *RecordTransactorSession) SetOracleAddress(_oracle common.Address) (*types.Transaction, error) {
	return _Record.Contract.SetOracleAddress(&_Record.TransactOpts, _oracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Record *RecordTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Record *RecordSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Record.Contract.TransferOwnership(&_Record.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Record *RecordTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Record.Contract.TransferOwnership(&_Record.TransactOpts, newOwner)
}

// RecordInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Record contract.
type RecordInitializedIterator struct {
	Event *RecordInitialized // Event containing the contract specifics and raw log

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
func (it *RecordInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordInitialized)
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
		it.Event = new(RecordInitialized)
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
func (it *RecordInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordInitialized represents a Initialized event raised by the Record contract.
type RecordInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Record *RecordFilterer) FilterInitialized(opts *bind.FilterOpts) (*RecordInitializedIterator, error) {

	logs, sub, err := _Record.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RecordInitializedIterator{contract: _Record.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Record *RecordFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RecordInitialized) (event.Subscription, error) {

	logs, sub, err := _Record.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordInitialized)
				if err := _Record.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Record *RecordFilterer) ParseInitialized(log types.Log) (*RecordInitialized, error) {
	event := new(RecordInitialized)
	if err := _Record.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Record contract.
type RecordOwnershipTransferredIterator struct {
	Event *RecordOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RecordOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordOwnershipTransferred)
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
		it.Event = new(RecordOwnershipTransferred)
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
func (it *RecordOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordOwnershipTransferred represents a OwnershipTransferred event raised by the Record contract.
type RecordOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Record *RecordFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RecordOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Record.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RecordOwnershipTransferredIterator{contract: _Record.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Record *RecordFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecordOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Record.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordOwnershipTransferred)
				if err := _Record.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Record *RecordFilterer) ParseOwnershipTransferred(log types.Log) (*RecordOwnershipTransferred, error) {
	event := new(RecordOwnershipTransferred)
	if err := _Record.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
