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
	Index                *big.Int
	BlockCount           *big.Int
	Sequencers           []common.Address
	SequencerBlocks      []*big.Int
	SequencerRatios      []*big.Int
	SequencerCommissions []*big.Int
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getBatchSubmissions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.BatchSubmission[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getRewardEpochs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerRatios\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerCommissions\",\"type\":\"uint256[]\"}],\"internalType\":\"structIRecord.RewardEpochInfo[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getRollupEpochs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.RollupEpochInfo[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRewardEpochBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchSubmissionIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRewardEpochIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRollupEpochIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rollupTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.BatchSubmission[]\",\"name\":\"_batchSubmissions\",\"type\":\"tuple[]\"}],\"name\":\"recordFinalizedBatchSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerRatios\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sequencerCommissions\",\"type\":\"uint256[]\"}],\"internalType\":\"structIRecord.RewardEpochInfo[]\",\"name\":\"_rewardEpochs\",\"type\":\"tuple[]\"}],\"name\":\"recordRewardEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRecord.RollupEpochInfo[]\",\"name\":\"_rollupEpochs\",\"type\":\"tuple[]\"}],\"name\":\"recordRollupEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_latestBlock\",\"type\":\"uint256\"}],\"name\":\"setLatestRewardEpochBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61014060405261271061012052348015610017575f80fd5b5073530000000000000000000000000000000000001060805273530000000000000000000000000000000000001260a05273530000000000000000000000000000000000000360c05273530000000000000000000000000000000000001160e0527353000000000000000000000000000000000000046101005260805160a05160c05160e05161010051610120516128d26100fb5f395f81816119b60152611b7701525f61030201525f818161022b0152611be701525f6103fe01525f81816103b9015261112101525f818161050001528181611216015261175601526128d25ff3fe608060405234801561000f575f80fd5b506004361061019a575f3560e01c806364b4abe3116100e85780638da5cb5b11610093578063a795f4091161006e578063a795f409146104a0578063cb6293e8146104db578063d5577141146104fb578063f2fde38b14610522575f80fd5b80638da5cb5b146103db5780638e21d5fb146103f9578063a24231e814610420575f80fd5b80637828a905116100c35780637828a905146102fd57806378f908e114610324578063807de443146103b4575f80fd5b806364b4abe3146102c25780636ea0396e146102e2578063715018a6146102f5575f80fd5b8063484f8d0f116101485780634c69c00f116101235780634c69c00f1461027c5780634e3ca4061461028f5780634ecff524146102af575f80fd5b8063484f8d0f1461024d578063485cc955146102565780634897504714610269575f80fd5b80632fbf6487116101785780632fbf6487146101d857806338013f02146101e15780633d9353fe14610226575f80fd5b80630776c0f71461019e57806310c9873f146101ba5780631511e1b1146101cf575b5f80fd5b6101a7606c5481565b6040519081526020015b60405180910390f35b6101cd6101c83660046121ea565b610535565b005b6101a760695481565b6101a7606b5481565b6065546102019073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6102017f000000000000000000000000000000000000000000000000000000000000000081565b6101a7606a5481565b6101cd610264366004612229565b6105f5565b6101cd61027736600461225a565b61080d565b6101cd61028a3660046122c9565b610a3f565b6102a261029d3660046122e9565b610af1565b6040516101b19190612309565b6101cd6102bd366004612382565b610c79565b6102d56102d03660046122e9565b610ed6565b6040516101b191906123df565b6101cd6102f0366004612455565b611067565b6101cd611cd4565b6102017f000000000000000000000000000000000000000000000000000000000000000081565b6103776103323660046121ea565b60666020525f908152604090208054600182015460028301546003840154600490940154929373ffffffffffffffffffffffffffffffffffffffff9092169290919085565b6040805195865273ffffffffffffffffffffffffffffffffffffffff9094166020860152928401919091526060830152608082015260a0016101b1565b6102017f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff16610201565b6102017f000000000000000000000000000000000000000000000000000000000000000081565b61046b61042e3660046121ea565b60676020525f90815260409020805460018201546002830154600390930154919273ffffffffffffffffffffffffffffffffffffffff9091169184565b6040805194855273ffffffffffffffffffffffffffffffffffffffff90931660208501529183015260608201526080016101b1565b6104c66104ae3660046121ea565b60686020525f90815260409020805460019091015482565b604080519283526020830191909152016101b1565b6104ee6104e93660046122e9565b611ce7565b6040516101b191906124ec565b6102017f000000000000000000000000000000000000000000000000000000000000000081565b6101cd6105303660046122c9565b611f9b565b60655473ffffffffffffffffffffffffffffffffffffffff1633146105a15760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f7765640000000000000000000000000060448201526064015b60405180910390fd5b5f81116105f05760405162461bcd60e51b815260206004820152601460248201527f696e76616c6964206c617465737420626c6f636b0000000000000000000000006044820152606401610598565b606c55565b5f54610100900460ff161580801561061357505f54600160ff909116105b8061062c5750303b15801561062c57505f5460ff166001145b61069e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610598565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156106fa575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff821661075d5760405162461bcd60e51b815260206004820152601660248201527f696e76616c6964206f7261636c652061646472657373000000000000000000006044820152606401610598565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556107a683612038565b8015610808575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60655473ffffffffffffffffffffffffffffffffffffffff1633146108745760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f776564000000000000000000000000006044820152606401610598565b5f5b818110156108085780606a5461088c9190612647565b83838381811061089e5761089e612660565b9050608002015f0135146108f45760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e646578000000000000000000000000000000000000006044820152606401610598565b604051806080016040528084848481811061091157610911612660565b9050608002015f0135815260200184848481811061093157610931612660565b905060800201602001602081019061094991906122c9565b73ffffffffffffffffffffffffffffffffffffffff16815260200184848481811061097657610976612660565b90506080020160400135815260200184848481811061099757610997612660565b9050608002016060013581525060675f8585858181106109b9576109b9612660565b608002919091013582525060208082019290925260409081015f208351815591830151600180840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915590830151600283015560609092015160039091015501610876565b610a476120ae565b73ffffffffffffffffffffffffffffffffffffffff8116610aaa5760405162461bcd60e51b815260206004820152601660248201527f696e76616c6964206f7261636c652061646472657373000000000000000000006044820152606401610598565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b606082821015610b435760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e646578000000000000000000000000000000000000006044820152606401610598565b610b4d838361268d565b610b58906001612647565b67ffffffffffffffff811115610b7057610b706126a0565b604051908082528060200260200182016040528015610bde57816020015b610bcb60405180608001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b815260200190600190039081610b8e5790505b509050825b828111610c72575f81815260676020908152604091829020825160808101845281548152600182015473ffffffffffffffffffffffffffffffffffffffff169281019290925260028101549282019290925260039091015460608201528251839083908110610c5457610c54612660565b60200260200101819052508080610c6a906126cd565b915050610be3565b5092915050565b60655473ffffffffffffffffffffffffffffffffffffffff163314610ce05760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f776564000000000000000000000000006044820152606401610598565b5f5b818110156108085780606954610cf89190612647565b838383818110610d0a57610d0a612660565b905060a002015f013514610d605760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e646578000000000000000000000000000000000000006044820152606401610598565b6040518060a00160405280848484818110610d7d57610d7d612660565b905060a002015f01358152602001848484818110610d9d57610d9d612660565b905060a002016020016020810190610db591906122c9565b73ffffffffffffffffffffffffffffffffffffffff168152602001848484818110610de257610de2612660565b905060a00201604001358152602001848484818110610e0357610e03612660565b905060a00201606001358152602001848484818110610e2457610e24612660565b905060a002016080013581525060665f858585818110610e4657610e46612660565b60a002919091013582525060208082019290925260409081015f208351815591830151600180840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091559083015160028301556060830151600383015560809092015160049091015501610ce2565b606082821015610f285760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e646578000000000000000000000000000000000000006044820152606401610598565b610f32838361268d565b610f3d906001612647565b67ffffffffffffffff811115610f5557610f556126a0565b604051908082528060200260200182016040528015610fc957816020015b610fb66040518060a001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81525090565b815260200190600190039081610f735790505b509050825b828111610c72575f81815260666020908152604091829020825160a08101845281548152600182015473ffffffffffffffffffffffffffffffffffffffff1692810192909252600281015492820192909252600382015460608201526004909101546080820152825183908390811061104957611049612660565b6020026020010181905250808061105f906126cd565b915050610fce565b60655473ffffffffffffffffffffffffffffffffffffffff1633146110ce5760405162461bcd60e51b815260206004820152601360248201527f6f6e6c79206f7261636c6520616c6c6f776564000000000000000000000000006044820152606401610598565b5f606c541161111f5760405162461bcd60e51b815260206004820152601960248201527f737461727420626c6f636b2073686f756c6420626520736574000000000000006044820152606401610598565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663766718086040518163ffffffff1660e01b8152600401602060405180830381865afa158015611188573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111ac9190612704565b606b546001906111bd908490612647565b6111c7919061268d565b106112145760405162461bcd60e51b815260206004820152601e60248201527f66757475726520646174612063616e6e6f742062652075706c6f6164656400006044820152606401610598565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a29bfb2c600184849050606b546112639190612647565b61126d919061268d565b6040518263ffffffff1660e01b815260040161128b91815260200190565b5f604051808303815f87803b1580156112a2575f80fd5b505af11580156112b4573d5f803e3d5ffd5b505050505f5b81811015611cb6575f8383838181106112d5576112d5612660565b90506020028101906112e7919061271b565b6112f5906040810190612757565b905090505f84848481811061130c5761130c612660565b905060200281019061131e919061271b565b606b5490359150611330908490612647565b811461137e5760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642065706f636820696e646578000000000000000000000000006044820152606401610598565b8185858581811061139157611391612660565b90506020028101906113a3919061271b565b6113b1906060810190612757565b90501480156113f25750818585858181106113ce576113ce612660565b90506020028101906113e0919061271b565b6113ee906080810190612757565b9050145b801561143057508185858581811061140c5761140c612660565b905060200281019061141e919061271b565b61142c9060a0810190612757565b9050145b61147c5760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642064617461206c656e677468000000000000000000000000006044820152606401610598565b84848481811061148e5761148e612660565b90506020028101906114a0919061271b565b60200135606c5f8282546114b49190612647565b925050819055506040518060c001604052808281526020018686868181106114de576114de612660565b90506020028101906114f0919061271b565b60200135815260200186868681811061150b5761150b612660565b905060200281019061151d919061271b565b61152b906040810190612757565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050509082525060200186868681811061157357611573612660565b9050602002810190611585919061271b565b611593906060810190612757565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152505050908252506020018686868181106115db576115db612660565b90506020028101906115ed919061271b565b6115fb906080810190612757565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050509082525060200186868681811061164357611643612660565b9050602002810190611655919061271b565b6116639060a0810190612757565b808060200260200160405190810160405280939291908181526020018383602002808284375f92018290525093909452505083815260686020908152604091829020845181558482015160018201559184015180519293506116ce9260028501929190910190612115565b50606082015180516116ea91600384019160209091019061219d565b506080820151805161170691600484019160209091019061219d565b5060a0820151805161172291600584019160209091019061219d565b50506040517f944fa746000000000000000000000000000000000000000000000000000000008152600481018390525f91507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063944fa74690602401602060405180830381865afa1580156117b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117d49190612704565b90505f805f8567ffffffffffffffff8111156117f2576117f26126a0565b60405190808252806020026020018201604052801561181b578160200160208202803683370190505b5090505f8667ffffffffffffffff811115611838576118386126a0565b604051908082528060200260200182016040528015611861578160200160208202803683370190505b5090505f5b87811015611afe5760148b8b8b81811061188257611882612660565b9050602002810190611894919061271b565b6118a29060a0810190612757565b838181106118b2576118b2612660565b9050602002013511156119075760405162461bcd60e51b815260206004820152601d60248201527f696e76616c69642073657175656e6365727320636f6d6d697373696f6e0000006044820152606401610598565b8a8a8a81811061191957611919612660565b905060200281019061192b919061271b565b611939906080810190612757565b8281811061194957611949612660565b905060200201358461195b9190612647565b93508a8a8a81811061196f5761196f612660565b9050602002810190611981919061271b565b61198f906060810190612757565b8281811061199f5761199f612660565b90506020020135856119b19190612647565b94505f7f00000000000000000000000000000000000000000000000000000000000000008c8c8c8181106119e7576119e7612660565b90506020028101906119f9919061271b565b611a07906080810190612757565b84818110611a1757611a17612660565b9050602002013588611a2991906127c2565b611a3391906127d9565b905060648c8c8c818110611a4957611a49612660565b9050602002810190611a5b919061271b565b611a699060a0810190612757565b84818110611a7957611a79612660565b9050602002013582611a8b91906127c2565b611a9591906127d9565b838381518110611aa757611aa7612660565b602002602001018181525050828281518110611ac557611ac5612660565b602002602001015181611ad8919061268d565b848381518110611aea57611aea612660565b602090810291909101015250600101611866565b50898989818110611b1157611b11612660565b9050602002810190611b23919061271b565b602001358414611b755760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642073657175656e6365727320626c6f636b73000000000000006044820152606401610598565b7f0000000000000000000000000000000000000000000000000000000000000000831115611be55760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642073657175656e6365727320726174696f73000000000000006044820152606401610598565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cdd0c50e878c8c8c818110611c3457611c34612660565b9050602002810190611c46919061271b565b611c54906040810190612757565b86866040518663ffffffff1660e01b8152600401611c76959493929190612811565b5f604051808303815f87803b158015611c8d575f80fd5b505af1158015611c9f573d5f803e3d5ffd5b5050600190990198506112ba975050505050505050565b5081819050606b5f828254611ccb9190612647565b90915550505050565b611cdc6120ae565b611ce55f612038565b565b606082821015611d395760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420696e646578000000000000000000000000000000000000006044820152606401610598565b611d43838361268d565b611d4e906001612647565b67ffffffffffffffff811115611d6657611d666126a0565b604051908082528060200260200182016040528015611dce57816020015b611dbb6040518060c001604052805f81526020015f8152602001606081526020016060815260200160608152602001606081525090565b815260200190600190039081611d845790505b509050825b828111610c72575f81815260686020908152604091829020825160c0810184528154815260018201548184015260028201805485518186028101860187528181529295939493860193830182828015611e6057602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611e35575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015611eb657602002820191905f5260205f20905b815481526020019060010190808311611ea2575b5050505050815260200160048201805480602002602001604051908101604052809291908181526020018280548015611f0c57602002820191905f5260205f20905b815481526020019060010190808311611ef8575b5050505050815260200160058201805480602002602001604051908101604052809291908181526020018280548015611f6257602002820191905f5260205f20905b815481526020019060010190808311611f4e575b505050505081525050828281518110611f7d57611f7d612660565b60200260200101819052508080611f93906126cd565b915050611dd3565b611fa36120ae565b73ffffffffffffffffffffffffffffffffffffffff811661202c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610598565b61203581612038565b50565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60335473ffffffffffffffffffffffffffffffffffffffff163314611ce55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610598565b828054828255905f5260205f2090810192821561218d579160200282015b8281111561218d57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190612133565b506121999291506121d6565b5090565b828054828255905f5260205f2090810192821561218d579160200282015b8281111561218d5782518255916020019190600101906121bb565b5b80821115612199575f81556001016121d7565b5f602082840312156121fa575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612224575f80fd5b919050565b5f806040838503121561223a575f80fd5b61224383612201565b915061225160208401612201565b90509250929050565b5f806020838503121561226b575f80fd5b823567ffffffffffffffff80821115612282575f80fd5b818501915085601f830112612295575f80fd5b8135818111156122a3575f80fd5b8660208260071b85010111156122b7575f80fd5b60209290920196919550909350505050565b5f602082840312156122d9575f80fd5b6122e282612201565b9392505050565b5f80604083850312156122fa575f80fd5b50508035926020909101359150565b602080825282518282018190525f919060409081850190868401855b82811015612375578151805185528681015173ffffffffffffffffffffffffffffffffffffffff168786015285810151868601526060908101519085015260809093019290850190600101612325565b5091979650505050505050565b5f8060208385031215612393575f80fd5b823567ffffffffffffffff808211156123aa575f80fd5b818501915085601f8301126123bd575f80fd5b8135818111156123cb575f80fd5b86602060a0830285010111156122b7575f80fd5b602080825282518282018190525f919060409081850190868401855b82811015612375578151805185528681015173ffffffffffffffffffffffffffffffffffffffff16878601528581015186860152606080820151908601526080908101519085015260a090930192908501906001016123fb565b5f8060208385031215612466575f80fd5b823567ffffffffffffffff8082111561247d575f80fd5b818501915085601f830112612490575f80fd5b81358181111561249e575f80fd5b8660208260051b85010111156122b7575f80fd5b5f815180845260208085019450602084015f5b838110156124e1578151875295820195908201906001016124c5565b509495945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b8381101561260c578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051845287810151888501528681015160c0888601819052815190860181905260e08601918a01905f905b808210156125a557825173ffffffffffffffffffffffffffffffffffffffff168452928b0192918b01916001919091019061256f565b505050606080830151868303828801526125bf83826124b2565b92505050608080830151868303828801526125da83826124b2565b9250505060a080830151925085820381870152506125f881836124b2565b968901969450505090860190600101612513565b509098975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561265a5761265a61261a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8181038181111561265a5761265a61261a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036126fd576126fd61261a565b5060010190565b5f60208284031215612714575f80fd5b5051919050565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4183360301811261274d575f80fd5b9190910192915050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261278a575f80fd5b83018035915067ffffffffffffffff8211156127a4575f80fd5b6020019150600581901b36038213156127bb575f80fd5b9250929050565b808202811582820484141761265a5761265a61261a565b5f8261280c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b858152608060208083018290529082018590525f90869060a08401835b888110156128675773ffffffffffffffffffffffffffffffffffffffff61285485612201565b168252928201929082019060010161282e565b50848103604086015261287a81886124b2565b92505050828103606084015261289081856124b2565b9897505050505050505056fea2646970667358221220433e7d1f6191556caa0aaf80e790753f1f71ed52f81825082dd75fececbb9c1564736f6c63430008180033",
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

// LatestRewardEpochBlock is a free data retrieval call binding the contract method 0x0776c0f7.
//
// Solidity: function latestRewardEpochBlock() view returns(uint256)
func (_Record *RecordCaller) LatestRewardEpochBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "latestRewardEpochBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRewardEpochBlock is a free data retrieval call binding the contract method 0x0776c0f7.
//
// Solidity: function latestRewardEpochBlock() view returns(uint256)
func (_Record *RecordSession) LatestRewardEpochBlock() (*big.Int, error) {
	return _Record.Contract.LatestRewardEpochBlock(&_Record.CallOpts)
}

// LatestRewardEpochBlock is a free data retrieval call binding the contract method 0x0776c0f7.
//
// Solidity: function latestRewardEpochBlock() view returns(uint256)
func (_Record *RecordCallerSession) LatestRewardEpochBlock() (*big.Int, error) {
	return _Record.Contract.LatestRewardEpochBlock(&_Record.CallOpts)
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

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordCaller) RewardEpochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "rewardEpochs", arg0)

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

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordSession) RewardEpochs(arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	return _Record.Contract.RewardEpochs(&_Record.CallOpts, arg0)
}

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 ) view returns(uint256 index, uint256 blockCount)
func (_Record *RecordCallerSession) RewardEpochs(arg0 *big.Int) (struct {
	Index      *big.Int
	BlockCount *big.Int
}, error) {
	return _Record.Contract.RewardEpochs(&_Record.CallOpts, arg0)
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

// SetLatestRewardEpochBlock is a paid mutator transaction binding the contract method 0x10c9873f.
//
// Solidity: function setLatestRewardEpochBlock(uint256 _latestBlock) returns()
func (_Record *RecordTransactor) SetLatestRewardEpochBlock(opts *bind.TransactOpts, _latestBlock *big.Int) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "setLatestRewardEpochBlock", _latestBlock)
}

// SetLatestRewardEpochBlock is a paid mutator transaction binding the contract method 0x10c9873f.
//
// Solidity: function setLatestRewardEpochBlock(uint256 _latestBlock) returns()
func (_Record *RecordSession) SetLatestRewardEpochBlock(_latestBlock *big.Int) (*types.Transaction, error) {
	return _Record.Contract.SetLatestRewardEpochBlock(&_Record.TransactOpts, _latestBlock)
}

// SetLatestRewardEpochBlock is a paid mutator transaction binding the contract method 0x10c9873f.
//
// Solidity: function setLatestRewardEpochBlock(uint256 _latestBlock) returns()
func (_Record *RecordTransactorSession) SetLatestRewardEpochBlock(_latestBlock *big.Int) (*types.Transaction, error) {
	return _Record.Contract.SetLatestRewardEpochBlock(&_Record.TransactOpts, _latestBlock)
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