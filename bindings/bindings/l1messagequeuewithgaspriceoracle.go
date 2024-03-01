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

// L1MessageQueueWithGasPriceOracleMetaData contains all meta data concerning the L1MessageQueueWithGasPriceOracle contract.
var L1MessageQueueWithGasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_enforcedTxGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"DequeueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DropTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"queueIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"UpdateEnforcedTxGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"UpdateGasOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldL2BaseFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newL2BaseFee\",\"type\":\"uint256\"}],\"name\":\"UpdateL2BaseFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxGasLimit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendEnforcedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"computeTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"dropCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enforcedTxGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageDropped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageSkipped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingQueueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"popCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newL2BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL2BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"updateMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801562000010575f80fd5b50604051620018a2380380620018a2833981016040819052620000339162000185565b6001600160a01b03831615806200005157506001600160a01b038216155b806200006457506001600160a01b038116155b156200008357604051630ecc6fdf60e41b815260040160405180910390fd5b6200008d620000ab565b6001600160a01b0392831660805290821660a0521660c052620001cc565b5f54610100900460ff1615620001175760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161462000167575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b038116811462000180575f80fd5b919050565b5f805f6060848603121562000198575f80fd5b620001a38462000169565b9250620001b36020850162000169565b9150620001c36040850162000169565b90509250925092565b60805160a05160c05161168f620002135f395f818161022d0152610bc901525f8181610313015261042901525f81816101be015281816108520152610a71015261168f5ff3fe608060405234801561000f575f80fd5b506004361061018f575f3560e01c8063a85006ca116100dd578063d99bc80e11610088578063f2fde38b11610063578063f2fde38b14610391578063fd0ad31e146103a4578063fe4b84df146103ac575f80fd5b8063d99bc80e1461035b578063e172d3a11461036e578063e3176bd514610388575f80fd5b8063cb23bcb5116100b8578063cb23bcb51461030e578063d5ad4a9714610335578063d7704bae14610348575f80fd5b8063a85006ca146102df578063ae453cd5146102e8578063bdc6f0a0146102fb575f80fd5b80635e45da231161013d5780638da5cb5b116101185780638da5cb5b1461029b57806391652461146102b95780639b159782146102cc575f80fd5b80635e45da2314610277578063715018a6146102805780637d82191a14610288575f80fd5b80633e83496c1161016d5780633e83496c1461022857806355f613ce1461024f5780635ad9945a14610264575f80fd5b806329aa604b146101935780633cb747bf146101b95780633e6dada114610205575b5f80fd5b6101a66101a1366004611339565b6103bf565b6040519081526020015b60405180910390f35b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b0565b610218610213366004611339565b6103de565b60405190151581526020016101b0565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b61026261025d366004611350565b610426565b005b6101a66102723660046113e1565b610619565b6101a660685481565b610262610809565b610218610296366004611339565b61081c565b60335473ffffffffffffffffffffffffffffffffffffffff166101e0565b6102626102c7366004611339565b61084f565b6102626102da36600461145d565b610a6e565b6101a660675481565b6101a66102f6366004611339565b610b50565b6102626103093660046114b3565b610bc6565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b610262610343366004611339565b610cf7565b6101a6610356366004611339565b610d45565b610262610369366004611339565b610d54565b6101a661037c366004611526565b60100261520801919050565b6101a660655481565b61026261039f366004611565565b610d9a565b6066546101a6565b6102626103ba366004611339565b610e37565b606681815481106103ce575f80fd5b5f91825260209091200154905081565b600881901c5f908152606a6020526040812054600160ff84161b16151580156104205750600882901c5f90815260696020526040902054600160ff84161b1615155b92915050565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16146104b05760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c792063616c6c61626c652062792074686520726f6c6c7570000000000060448201526064015b60405180910390fd5b6101008211156105025760405162461bcd60e51b815260206004820152601560248201527f706f7020746f6f206d616e79206d65737361676573000000000000000000000060448201526064016104a7565b82606754146105535760405162461bcd60e51b815260206004820152601460248201527f737461727420696e646578206d69736d6174636800000000000000000000000060448201526064016104a7565b600883901c5f818152606a6020526040902080546001851b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0193841660ff871681811b9092179092559092919061010081860111156105ca57600182015f908152606a6020526040902061010082900385901c90555b50505081830160675560408051848152602081018490529081018290527fc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e66792679709060600160405180910390a1505050565b5f607e816106c3565b5f8161063057506001919050565b5b81156106465760089190911c90600101610631565b919050565b8060808310600181146106835761066184610622565b60808101835360018301925084816020036008021b83528083019250506106a4565b84841516600181146106975784835361069c565b608083535b506001820191505b509392505050565b806094815360609290921b60018301525060150190565b600560405101806106d660018c8361064b565b90506106e46001898361064b565b90506106f089826106ac565b90506106fe60018b8361064b565b9050600186146001811461076657603887106001811461074b5761072188610622565b8060b701845360018401935088816020036008021b84528084019350508789843791870191610760565b87608001835360018301925087898437918701915b50610777565b6107745f89355f1a8461064b565b91505b506107828c826106ac565b90508181035f8060388310600181146107b55761079e84610622565b60f78101600882021b8517935060010191506107c0565b8360c0019250600191505b5086816008021b821791506001810190508060080292508451831c8284610100031b17915080850394505080845250508181038220925050508092505050979650505050505050565b610811610fab565b61081a5f611012565b565b5f606754821061082d57505f919050565b600882901c5f908152606a6020526040902054600160ff84161b161515610420565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16146108fa5760405162461bcd60e51b815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016104a7565b606754811061094b5760405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f742064726f702070656e64696e67206d657373616765000000000060448201526064016104a7565b600881901c5f908152606a6020526040902054600160ff83161b166109b25760405162461bcd60e51b815260206004820152601860248201527f64726f70206e6f6e2d736b6970706564206d657373616765000000000000000060448201526064016104a7565b600881901c5f90815260696020526040902054600160ff83161b1615610a1a5760405162461bcd60e51b815260206004820152601760248201527f6d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016104a7565b600881901c5f9081526069602052604090208054600160ff84161b1790556040518181527f43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf9060200160405180910390a150565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610b195760405162461bcd60e51b815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016104a7565b610b24838383611088565b3373111100000000000000000000000000000000111101610b4981865f878787611184565b5050505050565b6066545f908210610ba35760405162461bcd60e51b815260206004820152601a60248201527f6d65737361676520696e646578206f7574206f662072616e676500000000000060448201526064016104a7565b60668281548110610bb657610bb6611585565b905f5260205f2001549050919050565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610c715760405162461bcd60e51b815260206004820152602660248201527f4f6e6c792063616c6c61626c652062792074686520456e666f7263656454784760448201527f617465776179000000000000000000000000000000000000000000000000000060648201526084016104a7565b73ffffffffffffffffffffffffffffffffffffffff86163b15610cd65760405162461bcd60e51b815260206004820152600860248201527f6f6e6c7920454f4100000000000000000000000000000000000000000000000060448201526064016104a7565b610ce1838383611088565b610cef868686868686611184565b505050505050565b610cff610fab565b606880549082905560408051828152602081018490527fa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c591015b60405180910390a15050565b5f6065548261042091906115b2565b610d5c610fab565b606580549082905560408051828152602081018490527fc5271ba80b67178cc31f04a3755325121400925878dc608432b6fcaead3663299101610d39565b610da2610fab565b73ffffffffffffffffffffffffffffffffffffffff8116610e2b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104a7565b610e3481611012565b50565b5f54610100900460ff1615808015610e5557505f54600160ff909116105b80610e6e5750303b158015610e6e57505f5460ff166001145b610ee05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104a7565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610f3c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610f44611235565b60688290558015610fa7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610d39565b5050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461081a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104a7565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6068548311156111005760405162461bcd60e51b815260206004820152602560248201527f476173206c696d6974206d757374206e6f7420657863656564206d617847617360448201527f4c696d697400000000000000000000000000000000000000000000000000000060648201526084016104a7565b60108102615208018084101561117e5760405162461bcd60e51b815260206004820152603360248201527f496e73756666696369656e7420676173206c696d69742c206d7573742062652060448201527f61626f766520696e7472696e736963206761730000000000000000000000000060648201526084016104a7565b50505050565b6066545f6111978883888a898989610619565b606680546001810182555f919091527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540181905560405190915073ffffffffffffffffffffffffffffffffffffffff80891691908a16907f69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e90611223908a9087908b908b908b906115ee565b60405180910390a35050505050505050565b5f54610100900460ff166112b15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a7565b61081a5f54610100900460ff166113305760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a7565b61081a33611012565b5f60208284031215611349575f80fd5b5035919050565b5f805f60608486031215611362575f80fd5b505081359360208301359350604090920135919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610646575f80fd5b5f8083601f8401126113ac575f80fd5b50813567ffffffffffffffff8111156113c3575f80fd5b6020830191508360208285010111156113da575f80fd5b9250929050565b5f805f805f805f60c0888a0312156113f7575f80fd5b61140088611379565b9650602088013595506040880135945061141c60608901611379565b93506080880135925060a088013567ffffffffffffffff81111561143e575f80fd5b61144a8a828b0161139c565b989b979a50959850939692959293505050565b5f805f8060608587031215611470575f80fd5b61147985611379565b935060208501359250604085013567ffffffffffffffff81111561149b575f80fd5b6114a78782880161139c565b95989497509550505050565b5f805f805f8060a087890312156114c8575f80fd5b6114d187611379565b95506114df60208801611379565b94506040870135935060608701359250608087013567ffffffffffffffff811115611508575f80fd5b61151489828a0161139c565b979a9699509497509295939492505050565b5f8060208385031215611537575f80fd5b823567ffffffffffffffff81111561154d575f80fd5b6115598582860161139c565b90969095509350505050565b5f60208284031215611575575f80fd5b61157e82611379565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417610420577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b85815267ffffffffffffffff8516602082015283604082015260806060820152816080820152818360a08301375f81830160a090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010194935050505056fea26469706673582212200606d494bda0f0772567e1e207060bed52d4ee571fb7871bbbae25936d5cc19764736f6c63430008180033",
}

// L1MessageQueueWithGasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MessageQueueWithGasPriceOracleMetaData.ABI instead.
var L1MessageQueueWithGasPriceOracleABI = L1MessageQueueWithGasPriceOracleMetaData.ABI

// L1MessageQueueWithGasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1MessageQueueWithGasPriceOracleMetaData.Bin instead.
var L1MessageQueueWithGasPriceOracleBin = L1MessageQueueWithGasPriceOracleMetaData.Bin

// DeployL1MessageQueueWithGasPriceOracle deploys a new Ethereum contract, binding an instance of L1MessageQueueWithGasPriceOracle to it.
func DeployL1MessageQueueWithGasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address, _rollup common.Address, _enforcedTxGateway common.Address) (common.Address, *types.Transaction, *L1MessageQueueWithGasPriceOracle, error) {
	parsed, err := L1MessageQueueWithGasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1MessageQueueWithGasPriceOracleBin), backend, _messenger, _rollup, _enforcedTxGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1MessageQueueWithGasPriceOracle{L1MessageQueueWithGasPriceOracleCaller: L1MessageQueueWithGasPriceOracleCaller{contract: contract}, L1MessageQueueWithGasPriceOracleTransactor: L1MessageQueueWithGasPriceOracleTransactor{contract: contract}, L1MessageQueueWithGasPriceOracleFilterer: L1MessageQueueWithGasPriceOracleFilterer{contract: contract}}, nil
}

// L1MessageQueueWithGasPriceOracle is an auto generated Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracle struct {
	L1MessageQueueWithGasPriceOracleCaller     // Read-only binding to the contract
	L1MessageQueueWithGasPriceOracleTransactor // Write-only binding to the contract
	L1MessageQueueWithGasPriceOracleFilterer   // Log filterer for contract events
}

// L1MessageQueueWithGasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueWithGasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueWithGasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1MessageQueueWithGasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueWithGasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1MessageQueueWithGasPriceOracleSession struct {
	Contract     *L1MessageQueueWithGasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L1MessageQueueWithGasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1MessageQueueWithGasPriceOracleCallerSession struct {
	Contract *L1MessageQueueWithGasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// L1MessageQueueWithGasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1MessageQueueWithGasPriceOracleTransactorSession struct {
	Contract     *L1MessageQueueWithGasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// L1MessageQueueWithGasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracleRaw struct {
	Contract *L1MessageQueueWithGasPriceOracle // Generic contract binding to access the raw methods on
}

// L1MessageQueueWithGasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracleCallerRaw struct {
	Contract *L1MessageQueueWithGasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L1MessageQueueWithGasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1MessageQueueWithGasPriceOracleTransactorRaw struct {
	Contract *L1MessageQueueWithGasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1MessageQueueWithGasPriceOracle creates a new instance of L1MessageQueueWithGasPriceOracle, bound to a specific deployed contract.
func NewL1MessageQueueWithGasPriceOracle(address common.Address, backend bind.ContractBackend) (*L1MessageQueueWithGasPriceOracle, error) {
	contract, err := bindL1MessageQueueWithGasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracle{L1MessageQueueWithGasPriceOracleCaller: L1MessageQueueWithGasPriceOracleCaller{contract: contract}, L1MessageQueueWithGasPriceOracleTransactor: L1MessageQueueWithGasPriceOracleTransactor{contract: contract}, L1MessageQueueWithGasPriceOracleFilterer: L1MessageQueueWithGasPriceOracleFilterer{contract: contract}}, nil
}

// NewL1MessageQueueWithGasPriceOracleCaller creates a new read-only instance of L1MessageQueueWithGasPriceOracle, bound to a specific deployed contract.
func NewL1MessageQueueWithGasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L1MessageQueueWithGasPriceOracleCaller, error) {
	contract, err := bindL1MessageQueueWithGasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleCaller{contract: contract}, nil
}

// NewL1MessageQueueWithGasPriceOracleTransactor creates a new write-only instance of L1MessageQueueWithGasPriceOracle, bound to a specific deployed contract.
func NewL1MessageQueueWithGasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L1MessageQueueWithGasPriceOracleTransactor, error) {
	contract, err := bindL1MessageQueueWithGasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleTransactor{contract: contract}, nil
}

// NewL1MessageQueueWithGasPriceOracleFilterer creates a new log filterer instance of L1MessageQueueWithGasPriceOracle, bound to a specific deployed contract.
func NewL1MessageQueueWithGasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L1MessageQueueWithGasPriceOracleFilterer, error) {
	contract, err := bindL1MessageQueueWithGasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleFilterer{contract: contract}, nil
}

// bindL1MessageQueueWithGasPriceOracle binds a generic wrapper to an already deployed contract.
func bindL1MessageQueueWithGasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1MessageQueueWithGasPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueueWithGasPriceOracle.Contract.L1MessageQueueWithGasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.L1MessageQueueWithGasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.L1MessageQueueWithGasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueueWithGasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) pure returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _calldata []byte) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "calculateIntrinsicGasFee", _calldata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) pure returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L1MessageQueueWithGasPriceOracle.CallOpts, _calldata)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) pure returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L1MessageQueueWithGasPriceOracle.CallOpts, _calldata)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) ComputeTransactionHash(opts *bind.CallOpts, _sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "computeTransactionHash", _sender, _queueIndex, _value, _target, _gasLimit, _data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) ComputeTransactionHash(_sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.ComputeTransactionHash(&_L1MessageQueueWithGasPriceOracle.CallOpts, _sender, _queueIndex, _value, _target, _gasLimit, _data)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) ComputeTransactionHash(_sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.ComputeTransactionHash(&_L1MessageQueueWithGasPriceOracle.CallOpts, _sender, _queueIndex, _value, _target, _gasLimit, _data)
}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) EnforcedTxGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "enforcedTxGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) EnforcedTxGateway() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.EnforcedTxGateway(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) EnforcedTxGateway() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.EnforcedTxGateway(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueueWithGasPriceOracle.CallOpts, _gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueueWithGasPriceOracle.CallOpts, _gasLimit)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) GetCrossDomainMessage(opts *bind.CallOpts, _queueIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "getCrossDomainMessage", _queueIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.GetCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.GetCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) IsMessageDropped(opts *bind.CallOpts, _queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "isMessageDropped", _queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) IsMessageDropped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.IsMessageDropped(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) IsMessageDropped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.IsMessageDropped(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) IsMessageSkipped(opts *bind.CallOpts, _queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "isMessageSkipped", _queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) IsMessageSkipped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.IsMessageSkipped(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) IsMessageSkipped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.IsMessageSkipped(&_L1MessageQueueWithGasPriceOracle.CallOpts, _queueIndex)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) L2BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "l2BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) L2BaseFee() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.L2BaseFee(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) L2BaseFee() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.L2BaseFee(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) MaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "maxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) MaxGasLimit() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.MaxGasLimit(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) MaxGasLimit() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.MaxGasLimit(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) MessageQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "messageQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.MessageQueue(&_L1MessageQueueWithGasPriceOracle.CallOpts, arg0)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.MessageQueue(&_L1MessageQueueWithGasPriceOracle.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) Messenger() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Messenger(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) Messenger() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Messenger(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) NextCrossDomainMessageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "nextCrossDomainMessageIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.NextCrossDomainMessageIndex(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.NextCrossDomainMessageIndex(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) Owner() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Owner(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Owner(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) PendingQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "pendingQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) PendingQueueIndex() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.PendingQueueIndex(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) PendingQueueIndex() (*big.Int, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.PendingQueueIndex(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueueWithGasPriceOracle.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) Rollup() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Rollup(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleCallerSession) Rollup() (common.Address, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Rollup(&_L1MessageQueueWithGasPriceOracle.CallOpts)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) AppendCrossDomainMessage(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "appendCrossDomainMessage", _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.AppendCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.AppendCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _target, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) AppendEnforcedTransaction(opts *bind.TransactOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "appendEnforcedTransaction", _sender, _target, _value, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) AppendEnforcedTransaction(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.AppendEnforcedTransaction(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _sender, _target, _value, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) AppendEnforcedTransaction(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.AppendEnforcedTransaction(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _sender, _target, _value, _gasLimit, _data)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) DropCrossDomainMessage(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "dropCrossDomainMessage", _index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) DropCrossDomainMessage(_index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.DropCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) DropCrossDomainMessage(_index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.DropCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _index)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "initialize", _maxGasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) Initialize(_maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Initialize(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _maxGasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _maxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) Initialize(_maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.Initialize(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _maxGasLimit)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) PopCrossDomainMessage(opts *bind.TransactOpts, _startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "popCrossDomainMessage", _startIndex, _count, _skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) PopCrossDomainMessage(_startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.PopCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _startIndex, _count, _skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) PopCrossDomainMessage(_startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.PopCrossDomainMessage(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _startIndex, _count, _skippedBitmap)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.RenounceOwnership(&_L1MessageQueueWithGasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.RenounceOwnership(&_L1MessageQueueWithGasPriceOracle.TransactOpts)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) SetL2BaseFee(opts *bind.TransactOpts, _newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "setL2BaseFee", _newL2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) SetL2BaseFee(_newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.SetL2BaseFee(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _newL2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _newL2BaseFee) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) SetL2BaseFee(_newL2BaseFee *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.SetL2BaseFee(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _newL2BaseFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.TransferOwnership(&_L1MessageQueueWithGasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.TransferOwnership(&_L1MessageQueueWithGasPriceOracle.TransactOpts, newOwner)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactor) UpdateMaxGasLimit(opts *bind.TransactOpts, _newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.contract.Transact(opts, "updateMaxGasLimit", _newMaxGasLimit)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleSession) UpdateMaxGasLimit(_newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.UpdateMaxGasLimit(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _newMaxGasLimit)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleTransactorSession) UpdateMaxGasLimit(_newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueueWithGasPriceOracle.Contract.UpdateMaxGasLimit(&_L1MessageQueueWithGasPriceOracle.TransactOpts, _newMaxGasLimit)
}

// L1MessageQueueWithGasPriceOracleDequeueTransactionIterator is returned from FilterDequeueTransaction and is used to iterate over the raw logs and unpacked data for DequeueTransaction events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleDequeueTransactionIterator struct {
	Event *L1MessageQueueWithGasPriceOracleDequeueTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleDequeueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleDequeueTransaction)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleDequeueTransaction)
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
func (it *L1MessageQueueWithGasPriceOracleDequeueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleDequeueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleDequeueTransaction represents a DequeueTransaction event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleDequeueTransaction struct {
	StartIndex    *big.Int
	Count         *big.Int
	SkippedBitmap *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDequeueTransaction is a free log retrieval operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterDequeueTransaction(opts *bind.FilterOpts) (*L1MessageQueueWithGasPriceOracleDequeueTransactionIterator, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleDequeueTransactionIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "DequeueTransaction", logs: logs, sub: sub}, nil
}

// WatchDequeueTransaction is a free log subscription operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchDequeueTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleDequeueTransaction) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleDequeueTransaction)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
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

// ParseDequeueTransaction is a log parse operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseDequeueTransaction(log types.Log) (*L1MessageQueueWithGasPriceOracleDequeueTransaction, error) {
	event := new(L1MessageQueueWithGasPriceOracleDequeueTransaction)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleDropTransactionIterator is returned from FilterDropTransaction and is used to iterate over the raw logs and unpacked data for DropTransaction events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleDropTransactionIterator struct {
	Event *L1MessageQueueWithGasPriceOracleDropTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleDropTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleDropTransaction)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleDropTransaction)
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
func (it *L1MessageQueueWithGasPriceOracleDropTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleDropTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleDropTransaction represents a DropTransaction event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleDropTransaction struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDropTransaction is a free log retrieval operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterDropTransaction(opts *bind.FilterOpts) (*L1MessageQueueWithGasPriceOracleDropTransactionIterator, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleDropTransactionIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "DropTransaction", logs: logs, sub: sub}, nil
}

// WatchDropTransaction is a free log subscription operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchDropTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleDropTransaction) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleDropTransaction)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "DropTransaction", log); err != nil {
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

// ParseDropTransaction is a log parse operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseDropTransaction(log types.Log) (*L1MessageQueueWithGasPriceOracleDropTransaction, error) {
	event := new(L1MessageQueueWithGasPriceOracleDropTransaction)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "DropTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleInitializedIterator struct {
	Event *L1MessageQueueWithGasPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleInitialized)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleInitialized)
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
func (it *L1MessageQueueWithGasPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleInitialized represents a Initialized event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1MessageQueueWithGasPriceOracleInitializedIterator, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleInitializedIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleInitialized)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseInitialized(log types.Log) (*L1MessageQueueWithGasPriceOracleInitialized, error) {
	event := new(L1MessageQueueWithGasPriceOracleInitialized)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator struct {
	Event *L1MessageQueueWithGasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleOwnershipTransferred)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleOwnershipTransferred)
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
func (it *L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleOwnershipTransferredIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleOwnershipTransferred)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*L1MessageQueueWithGasPriceOracleOwnershipTransferred, error) {
	event := new(L1MessageQueueWithGasPriceOracleOwnershipTransferred)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleQueueTransactionIterator is returned from FilterQueueTransaction and is used to iterate over the raw logs and unpacked data for QueueTransaction events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleQueueTransactionIterator struct {
	Event *L1MessageQueueWithGasPriceOracleQueueTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleQueueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleQueueTransaction)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleQueueTransaction)
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
func (it *L1MessageQueueWithGasPriceOracleQueueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleQueueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleQueueTransaction represents a QueueTransaction event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleQueueTransaction struct {
	Sender     common.Address
	Target     common.Address
	Value      *big.Int
	QueueIndex uint64
	GasLimit   *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueTransaction is a free log retrieval operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterQueueTransaction(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*L1MessageQueueWithGasPriceOracleQueueTransactionIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleQueueTransactionIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "QueueTransaction", logs: logs, sub: sub}, nil
}

// WatchQueueTransaction is a free log subscription operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchQueueTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleQueueTransaction, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleQueueTransaction)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
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

// ParseQueueTransaction is a log parse operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseQueueTransaction(log types.Log) (*L1MessageQueueWithGasPriceOracleQueueTransaction, error) {
	event := new(L1MessageQueueWithGasPriceOracleQueueTransaction)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator is returned from FilterUpdateEnforcedTxGateway and is used to iterate over the raw logs and unpacked data for UpdateEnforcedTxGateway events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator struct {
	Event *L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway)
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
func (it *L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway represents a UpdateEnforcedTxGateway event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway struct {
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateEnforcedTxGateway is a free log retrieval operation binding the contract event 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
//
// Solidity: event UpdateEnforcedTxGateway(address indexed _oldGateway, address indexed _newGateway)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterUpdateEnforcedTxGateway(opts *bind.FilterOpts, _oldGateway []common.Address, _newGateway []common.Address) (*L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator, error) {

	var _oldGatewayRule []interface{}
	for _, _oldGatewayItem := range _oldGateway {
		_oldGatewayRule = append(_oldGatewayRule, _oldGatewayItem)
	}
	var _newGatewayRule []interface{}
	for _, _newGatewayItem := range _newGateway {
		_newGatewayRule = append(_newGatewayRule, _newGatewayItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "UpdateEnforcedTxGateway", _oldGatewayRule, _newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGatewayIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "UpdateEnforcedTxGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateEnforcedTxGateway is a free log subscription operation binding the contract event 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
//
// Solidity: event UpdateEnforcedTxGateway(address indexed _oldGateway, address indexed _newGateway)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchUpdateEnforcedTxGateway(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway, _oldGateway []common.Address, _newGateway []common.Address) (event.Subscription, error) {

	var _oldGatewayRule []interface{}
	for _, _oldGatewayItem := range _oldGateway {
		_oldGatewayRule = append(_oldGatewayRule, _oldGatewayItem)
	}
	var _newGatewayRule []interface{}
	for _, _newGatewayItem := range _newGateway {
		_newGatewayRule = append(_newGatewayRule, _newGatewayItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "UpdateEnforcedTxGateway", _oldGatewayRule, _newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateEnforcedTxGateway", log); err != nil {
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

// ParseUpdateEnforcedTxGateway is a log parse operation binding the contract event 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
//
// Solidity: event UpdateEnforcedTxGateway(address indexed _oldGateway, address indexed _newGateway)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseUpdateEnforcedTxGateway(log types.Log) (*L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway, error) {
	event := new(L1MessageQueueWithGasPriceOracleUpdateEnforcedTxGateway)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateEnforcedTxGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator is returned from FilterUpdateGasOracle and is used to iterate over the raw logs and unpacked data for UpdateGasOracle events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator struct {
	Event *L1MessageQueueWithGasPriceOracleUpdateGasOracle // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleUpdateGasOracle)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleUpdateGasOracle)
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
func (it *L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleUpdateGasOracle represents a UpdateGasOracle event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateGasOracle struct {
	OldGasOracle common.Address
	NewGasOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasOracle is a free log retrieval operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterUpdateGasOracle(opts *bind.FilterOpts, _oldGasOracle []common.Address, _newGasOracle []common.Address) (*L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleUpdateGasOracleIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "UpdateGasOracle", logs: logs, sub: sub}, nil
}

// WatchUpdateGasOracle is a free log subscription operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchUpdateGasOracle(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleUpdateGasOracle, _oldGasOracle []common.Address, _newGasOracle []common.Address) (event.Subscription, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleUpdateGasOracle)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
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

// ParseUpdateGasOracle is a log parse operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseUpdateGasOracle(log types.Log) (*L1MessageQueueWithGasPriceOracleUpdateGasOracle, error) {
	event := new(L1MessageQueueWithGasPriceOracleUpdateGasOracle)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator is returned from FilterUpdateL2BaseFee and is used to iterate over the raw logs and unpacked data for UpdateL2BaseFee events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator struct {
	Event *L1MessageQueueWithGasPriceOracleUpdateL2BaseFee // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleUpdateL2BaseFee)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleUpdateL2BaseFee)
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
func (it *L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleUpdateL2BaseFee represents a UpdateL2BaseFee event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateL2BaseFee struct {
	OldL2BaseFee *big.Int
	NewL2BaseFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateL2BaseFee is a free log retrieval operation binding the contract event 0xc5271ba80b67178cc31f04a3755325121400925878dc608432b6fcaead366329.
//
// Solidity: event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterUpdateL2BaseFee(opts *bind.FilterOpts) (*L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "UpdateL2BaseFee")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleUpdateL2BaseFeeIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "UpdateL2BaseFee", logs: logs, sub: sub}, nil
}

// WatchUpdateL2BaseFee is a free log subscription operation binding the contract event 0xc5271ba80b67178cc31f04a3755325121400925878dc608432b6fcaead366329.
//
// Solidity: event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchUpdateL2BaseFee(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleUpdateL2BaseFee) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "UpdateL2BaseFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleUpdateL2BaseFee)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateL2BaseFee", log); err != nil {
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

// ParseUpdateL2BaseFee is a log parse operation binding the contract event 0xc5271ba80b67178cc31f04a3755325121400925878dc608432b6fcaead366329.
//
// Solidity: event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseUpdateL2BaseFee(log types.Log) (*L1MessageQueueWithGasPriceOracleUpdateL2BaseFee, error) {
	event := new(L1MessageQueueWithGasPriceOracleUpdateL2BaseFee)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateL2BaseFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator is returned from FilterUpdateMaxGasLimit and is used to iterate over the raw logs and unpacked data for UpdateMaxGasLimit events raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator struct {
	Event *L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit)
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
		it.Event = new(L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit)
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
func (it *L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit represents a UpdateMaxGasLimit event raised by the L1MessageQueueWithGasPriceOracle contract.
type L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit struct {
	OldMaxGasLimit *big.Int
	NewMaxGasLimit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxGasLimit is a free log retrieval operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) FilterUpdateMaxGasLimit(opts *bind.FilterOpts) (*L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.FilterLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueWithGasPriceOracleUpdateMaxGasLimitIterator{contract: _L1MessageQueueWithGasPriceOracle.contract, event: "UpdateMaxGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxGasLimit is a free log subscription operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) WatchUpdateMaxGasLimit(opts *bind.WatchOpts, sink chan<- *L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueueWithGasPriceOracle.contract.WatchLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit)
				if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
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

// ParseUpdateMaxGasLimit is a log parse operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_L1MessageQueueWithGasPriceOracle *L1MessageQueueWithGasPriceOracleFilterer) ParseUpdateMaxGasLimit(log types.Log) (*L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit, error) {
	event := new(L1MessageQueueWithGasPriceOracleUpdateMaxGasLimit)
	if err := _L1MessageQueueWithGasPriceOracle.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
