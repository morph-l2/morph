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

// L1MessageQueueMetaData contains all meta data concerning the L1MessageQueue contract.
var L1MessageQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"DequeueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DropTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"queueIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"UpdateEnforcedTxGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"UpdateGasOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxGasLimit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendEnforcedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"computeTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"dropCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enforcedTxGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_enforcedTxGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageDropped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageSkipped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingQueueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"popCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"updateEnforcedTxGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"updateGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"updateMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611a41806100e65f395ff3fe608060405234801561000f575f80fd5b506004361061019a575f3560e01c806391652461116100e8578063cb23bcb511610093578063e172d3a11161006e578063e172d3a114610397578063f2fde38b146103aa578063f7013ef6146103bd578063fd0ad31e146103d0575f80fd5b8063cb23bcb514610351578063d5ad4a9714610371578063d7704bae14610384575f80fd5b8063ae453cd5116100c3578063ae453cd514610318578063bdc6f0a01461032b578063c94864e11461033e575f80fd5b806391652461146102e95780639b159782146102fc578063a85006ca1461030f575f80fd5b80635d62a8dd11610148578063715018a611610123578063715018a6146102b05780637d82191a146102b85780638da5cb5b146102cb575f80fd5b80635d62a8dd146102745780635e45da231461029457806370cee67f1461029d575f80fd5b80633e83496c116101785780633e83496c1461022c57806355f613ce1461024c5780635ad9945a14610261575f80fd5b806329aa604b1461019e5780633cb747bf146101c45780633e6dada114610209575b5f80fd5b6101b16101ac36600461159e565b6103d8565b6040519081526020015b60405180910390f35b6065546101e49073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bb565b61021c61021736600461159e565b6103f7565b60405190151581526020016101bb565b6067546101e49073ffffffffffffffffffffffffffffffffffffffff1681565b61025f61025a3660046115b5565b61043f565b005b6101b161026f366004611646565b61062a565b6068546101e49073ffffffffffffffffffffffffffffffffffffffff1681565b6101b1606b5481565b61025f6102ab3660046116c2565b61081a565b61025f610898565b61021c6102c636600461159e565b6108ab565b60335473ffffffffffffffffffffffffffffffffffffffff166101e4565b61025f6102f736600461159e565b6108de565b61025f61030a3660046116db565b610af5565b6101b1606a5481565b6101b161032636600461159e565b610c03565b61025f610339366004611731565b610c27565b61025f61034c3660046116c2565b610d84565b6066546101e49073ffffffffffffffffffffffffffffffffffffffff1681565b61025f61037f36600461159e565b610e02565b6101b161039236600461159e565b610e4f565b6101b16103a53660046117d1565b610f0c565b61025f6103b83660046116c2565b610f86565b61025f6103cb36600461189a565b611023565b6069546101b1565b606981815481106103e7575f80fd5b5f91825260209091200154905081565b600881901c5f908152606d6020526040812054600160ff84161b16151580156104395750600882901c5f908152606c6020526040902054600160ff84161b1615155b92915050565b60665473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c15760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c792063616c6c61626c652062792074686520726f6c6c7570000000000060448201526064015b60405180910390fd5b6101008211156105135760405162461bcd60e51b815260206004820152601560248201527f706f7020746f6f206d616e79206d65737361676573000000000000000000000060448201526064016104b8565b82606a54146105645760405162461bcd60e51b815260206004820152601460248201527f737461727420696e646578206d69736d6174636800000000000000000000000060448201526064016104b8565b600883901c5f818152606d6020526040902080546001851b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0193841660ff871681811b9092179092559092919061010081860111156105db57600182015f908152606d6020526040902061010082900385901c90555b505050818301606a5560408051848152602081018490529081018290527fc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e66792679709060600160405180910390a1505050565b5f607e816106d4565b5f8161064157506001919050565b5b81156106575760089190911c90600101610642565b919050565b8060808310600181146106945761067284610633565b60808101835360018301925084816020036008021b83528083019250506106b5565b84841516600181146106a8578483536106ad565b608083535b506001820191505b509392505050565b806094815360609290921b60018301525060150190565b600560405101806106e760018c8361065c565b90506106f56001898361065c565b905061070189826106bd565b905061070f60018b8361065c565b9050600186146001811461077757603887106001811461075c5761073288610633565b8060b701845360018401935088816020036008021b84528084019350508789843791870191610771565b87608001835360018301925087898437918701915b50610788565b6107855f89355f1a8461065c565b91505b506107938c826106bd565b90508181035f8060388310600181146107c6576107af84610633565b60f78101600882021b8517935060010191506107d1565b8360c0019250600191505b5086816008021b821791506001810190508060080292508451831c8284610100031b17915080850394505080845250508181038220925050508092505050979650505050505050565b61082261120d565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e905f90a35050565b6108a061120d565b6108a95f611274565b565b5f606a5482106108bc57505f919050565b600882901c5f908152606d6020526040902054600160ff84161b161515610439565b60655473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109815760405162461bcd60e51b815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016104b8565b606a5481106109d25760405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f742064726f702070656e64696e67206d657373616765000000000060448201526064016104b8565b600881901c5f908152606d6020526040902054600160ff83161b16610a395760405162461bcd60e51b815260206004820152601860248201527f64726f70206e6f6e2d736b6970706564206d657373616765000000000000000060448201526064016104b8565b600881901c5f908152606c6020526040902054600160ff83161b1615610aa15760405162461bcd60e51b815260206004820152601760248201527f6d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016104b8565b600881901c5f908152606c602052604090208054600160ff84161b1790556040518181527f43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf9060200160405180910390a150565b60655473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b985760405162461bcd60e51b815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016104b8565b610bd78383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506112ea92505050565b3373111100000000000000000000000000000000111101610bfc81865f8787876113e9565b5050505050565b5f60698281548110610c1757610c176118f2565b905f5260205f2001549050919050565b60675473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cca5760405162461bcd60e51b815260206004820152602660248201527f4f6e6c792063616c6c61626c652062792074686520456e666f7263656454784760448201527f617465776179000000000000000000000000000000000000000000000000000060648201526084016104b8565b73ffffffffffffffffffffffffffffffffffffffff86163b15610d2f5760405162461bcd60e51b815260206004820152600860248201527f6f6e6c7920454f4100000000000000000000000000000000000000000000000060448201526064016104b8565b610d6e8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506112ea92505050565b610d7c8686868686866113e9565b505050505050565b610d8c61120d565b6067805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd905f90a35050565b610e0a61120d565b606b80549082905560408051828152602081018490527fa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5910160405180910390a15050565b6068545f9073ffffffffffffffffffffffffffffffffffffffff1680610e7757505f92915050565b6040517fd7704bae0000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff82169063d7704bae906024015b602060405180830381865afa158015610ee1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f05919061191f565b9392505050565b6068545f9073ffffffffffffffffffffffffffffffffffffffff1680610f3457505f92915050565b6040517fe172d3a100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063e172d3a190610ec6908690600401611936565b610f8e61120d565b73ffffffffffffffffffffffffffffffffffffffff81166110175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104b8565b61102081611274565b50565b5f54610100900460ff161580801561104157505f54600160ff909116105b8061105a5750303b15801561105a57505f5460ff166001145b6110cc5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104b8565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611128575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61113061149a565b6065805473ffffffffffffffffffffffffffffffffffffffff8089167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560668054888416908316179055606780548784169083161790556068805492861692909116919091179055606b8290558015610d7c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108a95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104b8565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b606b548211156113625760405162461bcd60e51b815260206004820152602560248201527f476173206c696d6974206d757374206e6f7420657863656564206d617847617360448201527f4c696d697400000000000000000000000000000000000000000000000000000060648201526084016104b8565b5f61136c82610f0c565b9050808310156113e45760405162461bcd60e51b815260206004820152603360248201527f496e73756666696369656e7420676173206c696d69742c206d7573742062652060448201527f61626f766520696e7472696e736963206761730000000000000000000000000060648201526084016104b8565b505050565b6069545f6113fc8883888a89898961062a565b606980546001810182555f919091527f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe3999143080181905560405190915073ffffffffffffffffffffffffffffffffffffffff80891691908a16907f69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e90611488908a9087908b908b908b906119a0565b60405180910390a35050505050505050565b5f54610100900460ff166115165760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b8565b6108a95f54610100900460ff166115955760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b8565b6108a933611274565b5f602082840312156115ae575f80fd5b5035919050565b5f805f606084860312156115c7575f80fd5b505081359360208301359350604090920135919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610657575f80fd5b5f8083601f840112611611575f80fd5b50813567ffffffffffffffff811115611628575f80fd5b60208301915083602082850101111561163f575f80fd5b9250929050565b5f805f805f805f60c0888a03121561165c575f80fd5b611665886115de565b96506020880135955060408801359450611681606089016115de565b93506080880135925060a088013567ffffffffffffffff8111156116a3575f80fd5b6116af8a828b01611601565b989b979a50959850939692959293505050565b5f602082840312156116d2575f80fd5b610f05826115de565b5f805f80606085870312156116ee575f80fd5b6116f7856115de565b935060208501359250604085013567ffffffffffffffff811115611719575f80fd5b61172587828801611601565b95989497509550505050565b5f805f805f8060a08789031215611746575f80fd5b61174f876115de565b955061175d602088016115de565b94506040870135935060608701359250608087013567ffffffffffffffff811115611786575f80fd5b61179289828a01611601565b979a9699509497509295939492505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f602082840312156117e1575f80fd5b813567ffffffffffffffff808211156117f8575f80fd5b818401915084601f83011261180b575f80fd5b81358181111561181d5761181d6117a4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611863576118636117a4565b8160405282815287602084870101111561187b575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f805f805f60a086880312156118ae575f80fd5b6118b7866115de565b94506118c5602087016115de565b93506118d3604087016115de565b92506118e1606087016115de565b949793965091946080013592915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020828403121561192f575f80fd5b5051919050565b5f602080835283518060208501525f5b8181101561196257858101830151858201604001528201611946565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b85815267ffffffffffffffff8516602082015283604082015260806060820152816080820152818360a08301375f81830160a090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010194935050505056fea2646970667358221220dcb1394808be3dc565ac19ae301cc3c6582e67f19633747fe7ac09719eb5213064736f6c63430008180033",
}

// L1MessageQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MessageQueueMetaData.ABI instead.
var L1MessageQueueABI = L1MessageQueueMetaData.ABI

// L1MessageQueueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1MessageQueueMetaData.Bin instead.
var L1MessageQueueBin = L1MessageQueueMetaData.Bin

// DeployL1MessageQueue deploys a new Ethereum contract, binding an instance of L1MessageQueue to it.
func DeployL1MessageQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1MessageQueue, error) {
	parsed, err := L1MessageQueueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1MessageQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1MessageQueue{L1MessageQueueCaller: L1MessageQueueCaller{contract: contract}, L1MessageQueueTransactor: L1MessageQueueTransactor{contract: contract}, L1MessageQueueFilterer: L1MessageQueueFilterer{contract: contract}}, nil
}

// L1MessageQueue is an auto generated Go binding around an Ethereum contract.
type L1MessageQueue struct {
	L1MessageQueueCaller     // Read-only binding to the contract
	L1MessageQueueTransactor // Write-only binding to the contract
	L1MessageQueueFilterer   // Log filterer for contract events
}

// L1MessageQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1MessageQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1MessageQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1MessageQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1MessageQueueSession struct {
	Contract     *L1MessageQueue   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1MessageQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1MessageQueueCallerSession struct {
	Contract *L1MessageQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L1MessageQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1MessageQueueTransactorSession struct {
	Contract     *L1MessageQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L1MessageQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1MessageQueueRaw struct {
	Contract *L1MessageQueue // Generic contract binding to access the raw methods on
}

// L1MessageQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1MessageQueueCallerRaw struct {
	Contract *L1MessageQueueCaller // Generic read-only contract binding to access the raw methods on
}

// L1MessageQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1MessageQueueTransactorRaw struct {
	Contract *L1MessageQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1MessageQueue creates a new instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueue(address common.Address, backend bind.ContractBackend) (*L1MessageQueue, error) {
	contract, err := bindL1MessageQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueue{L1MessageQueueCaller: L1MessageQueueCaller{contract: contract}, L1MessageQueueTransactor: L1MessageQueueTransactor{contract: contract}, L1MessageQueueFilterer: L1MessageQueueFilterer{contract: contract}}, nil
}

// NewL1MessageQueueCaller creates a new read-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueCaller(address common.Address, caller bind.ContractCaller) (*L1MessageQueueCaller, error) {
	contract, err := bindL1MessageQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueCaller{contract: contract}, nil
}

// NewL1MessageQueueTransactor creates a new write-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*L1MessageQueueTransactor, error) {
	contract, err := bindL1MessageQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueTransactor{contract: contract}, nil
}

// NewL1MessageQueueFilterer creates a new log filterer instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*L1MessageQueueFilterer, error) {
	contract, err := bindL1MessageQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueFilterer{contract: contract}, nil
}

// bindL1MessageQueue binds a generic wrapper to an already deployed contract.
func bindL1MessageQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1MessageQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueue *L1MessageQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueue.Contract.L1MessageQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueue *L1MessageQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.L1MessageQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueue *L1MessageQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.L1MessageQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueue *L1MessageQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueue *L1MessageQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueue *L1MessageQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.contract.Transact(opts, method, params...)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _calldata []byte) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "calculateIntrinsicGasFee", _calldata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _L1MessageQueue.Contract.CalculateIntrinsicGasFee(&_L1MessageQueue.CallOpts, _calldata)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _L1MessageQueue.Contract.CalculateIntrinsicGasFee(&_L1MessageQueue.CallOpts, _calldata)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) ComputeTransactionHash(opts *bind.CallOpts, _sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "computeTransactionHash", _sender, _queueIndex, _value, _target, _gasLimit, _data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueue *L1MessageQueueSession) ComputeTransactionHash(_sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	return _L1MessageQueue.Contract.ComputeTransactionHash(&_L1MessageQueue.CallOpts, _sender, _queueIndex, _value, _target, _gasLimit, _data)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCallerSession) ComputeTransactionHash(_sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	return _L1MessageQueue.Contract.ComputeTransactionHash(&_L1MessageQueue.CallOpts, _sender, _queueIndex, _value, _target, _gasLimit, _data)
}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) EnforcedTxGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "enforcedTxGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) EnforcedTxGateway() (common.Address, error) {
	return _L1MessageQueue.Contract.EnforcedTxGateway(&_L1MessageQueue.CallOpts)
}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) EnforcedTxGateway() (common.Address, error) {
	return _L1MessageQueue.Contract.EnforcedTxGateway(&_L1MessageQueue.CallOpts)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueue.CallOpts, _gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueue.CallOpts, _gasLimit)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) GasOracle() (common.Address, error) {
	return _L1MessageQueue.Contract.GasOracle(&_L1MessageQueue.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) GasOracle() (common.Address, error) {
	return _L1MessageQueue.Contract.GasOracle(&_L1MessageQueue.CallOpts)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) GetCrossDomainMessage(opts *bind.CallOpts, _queueIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "getCrossDomainMessage", _queueIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.GetCrossDomainMessage(&_L1MessageQueue.CallOpts, _queueIndex)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCallerSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.GetCrossDomainMessage(&_L1MessageQueue.CallOpts, _queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueCaller) IsMessageDropped(opts *bind.CallOpts, _queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "isMessageDropped", _queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueSession) IsMessageDropped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueue.Contract.IsMessageDropped(&_L1MessageQueue.CallOpts, _queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueCallerSession) IsMessageDropped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueue.Contract.IsMessageDropped(&_L1MessageQueue.CallOpts, _queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueCaller) IsMessageSkipped(opts *bind.CallOpts, _queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "isMessageSkipped", _queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueSession) IsMessageSkipped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueue.Contract.IsMessageSkipped(&_L1MessageQueue.CallOpts, _queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 _queueIndex) view returns(bool)
func (_L1MessageQueue *L1MessageQueueCallerSession) IsMessageSkipped(_queueIndex *big.Int) (bool, error) {
	return _L1MessageQueue.Contract.IsMessageSkipped(&_L1MessageQueue.CallOpts, _queueIndex)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) MaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "maxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) MaxGasLimit() (*big.Int, error) {
	return _L1MessageQueue.Contract.MaxGasLimit(&_L1MessageQueue.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) MaxGasLimit() (*big.Int, error) {
	return _L1MessageQueue.Contract.MaxGasLimit(&_L1MessageQueue.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) MessageQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messageQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.MessageQueue(&_L1MessageQueue.CallOpts, arg0)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCallerSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.MessageQueue(&_L1MessageQueue.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) Messenger() (common.Address, error) {
	return _L1MessageQueue.Contract.Messenger(&_L1MessageQueue.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) Messenger() (common.Address, error) {
	return _L1MessageQueue.Contract.Messenger(&_L1MessageQueue.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) NextCrossDomainMessageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "nextCrossDomainMessageIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.NextCrossDomainMessageIndex(&_L1MessageQueue.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.NextCrossDomainMessageIndex(&_L1MessageQueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) Owner() (common.Address, error) {
	return _L1MessageQueue.Contract.Owner(&_L1MessageQueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) Owner() (common.Address, error) {
	return _L1MessageQueue.Contract.Owner(&_L1MessageQueue.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) PendingQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "pendingQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) PendingQueueIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.PendingQueueIndex(&_L1MessageQueue.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) PendingQueueIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.PendingQueueIndex(&_L1MessageQueue.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) Rollup() (common.Address, error) {
	return _L1MessageQueue.Contract.Rollup(&_L1MessageQueue.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) Rollup() (common.Address, error) {
	return _L1MessageQueue.Contract.Rollup(&_L1MessageQueue.CallOpts)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendCrossDomainMessage(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendCrossDomainMessage", _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendCrossDomainMessage(&_L1MessageQueue.TransactOpts, _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendCrossDomainMessage(&_L1MessageQueue.TransactOpts, _target, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendEnforcedTransaction(opts *bind.TransactOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendEnforcedTransaction", _sender, _target, _value, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueSession) AppendEnforcedTransaction(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendEnforcedTransaction(&_L1MessageQueue.TransactOpts, _sender, _target, _value, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) AppendEnforcedTransaction(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendEnforcedTransaction(&_L1MessageQueue.TransactOpts, _sender, _target, _value, _gasLimit, _data)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) DropCrossDomainMessage(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "dropCrossDomainMessage", _index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueue *L1MessageQueueSession) DropCrossDomainMessage(_index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.DropCrossDomainMessage(&_L1MessageQueue.TransactOpts, _index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) DropCrossDomainMessage(_index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.DropCrossDomainMessage(&_L1MessageQueue.TransactOpts, _index)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _messenger, address _rollup, address _enforcedTxGateway, address _gasOracle, uint256 _maxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) Initialize(opts *bind.TransactOpts, _messenger common.Address, _rollup common.Address, _enforcedTxGateway common.Address, _gasOracle common.Address, _maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "initialize", _messenger, _rollup, _enforcedTxGateway, _gasOracle, _maxGasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _messenger, address _rollup, address _enforcedTxGateway, address _gasOracle, uint256 _maxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueSession) Initialize(_messenger common.Address, _rollup common.Address, _enforcedTxGateway common.Address, _gasOracle common.Address, _maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.Initialize(&_L1MessageQueue.TransactOpts, _messenger, _rollup, _enforcedTxGateway, _gasOracle, _maxGasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _messenger, address _rollup, address _enforcedTxGateway, address _gasOracle, uint256 _maxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) Initialize(_messenger common.Address, _rollup common.Address, _enforcedTxGateway common.Address, _gasOracle common.Address, _maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.Initialize(&_L1MessageQueue.TransactOpts, _messenger, _rollup, _enforcedTxGateway, _gasOracle, _maxGasLimit)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) PopCrossDomainMessage(opts *bind.TransactOpts, _startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "popCrossDomainMessage", _startIndex, _count, _skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueue *L1MessageQueueSession) PopCrossDomainMessage(_startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.PopCrossDomainMessage(&_L1MessageQueue.TransactOpts, _startIndex, _count, _skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) PopCrossDomainMessage(_startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.PopCrossDomainMessage(&_L1MessageQueue.TransactOpts, _startIndex, _count, _skippedBitmap)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueue.Contract.RenounceOwnership(&_L1MessageQueue.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueue.Contract.RenounceOwnership(&_L1MessageQueue.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.TransferOwnership(&_L1MessageQueue.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.TransferOwnership(&_L1MessageQueue.TransactOpts, newOwner)
}

// UpdateEnforcedTxGateway is a paid mutator transaction binding the contract method 0xc94864e1.
//
// Solidity: function updateEnforcedTxGateway(address _newGateway) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateEnforcedTxGateway(opts *bind.TransactOpts, _newGateway common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateEnforcedTxGateway", _newGateway)
}

// UpdateEnforcedTxGateway is a paid mutator transaction binding the contract method 0xc94864e1.
//
// Solidity: function updateEnforcedTxGateway(address _newGateway) returns()
func (_L1MessageQueue *L1MessageQueueSession) UpdateEnforcedTxGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateEnforcedTxGateway(&_L1MessageQueue.TransactOpts, _newGateway)
}

// UpdateEnforcedTxGateway is a paid mutator transaction binding the contract method 0xc94864e1.
//
// Solidity: function updateEnforcedTxGateway(address _newGateway) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) UpdateEnforcedTxGateway(_newGateway common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateEnforcedTxGateway(&_L1MessageQueue.TransactOpts, _newGateway)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateGasOracle(opts *bind.TransactOpts, _newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateGasOracle", _newGasOracle)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueSession) UpdateGasOracle(_newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateGasOracle(&_L1MessageQueue.TransactOpts, _newGasOracle)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) UpdateGasOracle(_newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateGasOracle(&_L1MessageQueue.TransactOpts, _newGasOracle)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateMaxGasLimit(opts *bind.TransactOpts, _newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateMaxGasLimit", _newMaxGasLimit)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueSession) UpdateMaxGasLimit(_newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateMaxGasLimit(&_L1MessageQueue.TransactOpts, _newMaxGasLimit)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) UpdateMaxGasLimit(_newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateMaxGasLimit(&_L1MessageQueue.TransactOpts, _newMaxGasLimit)
}

// L1MessageQueueDequeueTransactionIterator is returned from FilterDequeueTransaction and is used to iterate over the raw logs and unpacked data for DequeueTransaction events raised by the L1MessageQueue contract.
type L1MessageQueueDequeueTransactionIterator struct {
	Event *L1MessageQueueDequeueTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueDequeueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueDequeueTransaction)
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
		it.Event = new(L1MessageQueueDequeueTransaction)
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
func (it *L1MessageQueueDequeueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueDequeueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueDequeueTransaction represents a DequeueTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueDequeueTransaction struct {
	StartIndex    *big.Int
	Count         *big.Int
	SkippedBitmap *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDequeueTransaction is a free log retrieval operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterDequeueTransaction(opts *bind.FilterOpts) (*L1MessageQueueDequeueTransactionIterator, error) {

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueDequeueTransactionIterator{contract: _L1MessageQueue.contract, event: "DequeueTransaction", logs: logs, sub: sub}, nil
}

// WatchDequeueTransaction is a free log subscription operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchDequeueTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueDequeueTransaction) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueDequeueTransaction)
				if err := _L1MessageQueue.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseDequeueTransaction(log types.Log) (*L1MessageQueueDequeueTransaction, error) {
	event := new(L1MessageQueueDequeueTransaction)
	if err := _L1MessageQueue.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueDropTransactionIterator is returned from FilterDropTransaction and is used to iterate over the raw logs and unpacked data for DropTransaction events raised by the L1MessageQueue contract.
type L1MessageQueueDropTransactionIterator struct {
	Event *L1MessageQueueDropTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueDropTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueDropTransaction)
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
		it.Event = new(L1MessageQueueDropTransaction)
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
func (it *L1MessageQueueDropTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueDropTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueDropTransaction represents a DropTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueDropTransaction struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDropTransaction is a free log retrieval operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterDropTransaction(opts *bind.FilterOpts) (*L1MessageQueueDropTransactionIterator, error) {

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueDropTransactionIterator{contract: _L1MessageQueue.contract, event: "DropTransaction", logs: logs, sub: sub}, nil
}

// WatchDropTransaction is a free log subscription operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchDropTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueDropTransaction) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueDropTransaction)
				if err := _L1MessageQueue.contract.UnpackLog(event, "DropTransaction", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseDropTransaction(log types.Log) (*L1MessageQueueDropTransaction, error) {
	event := new(L1MessageQueueDropTransaction)
	if err := _L1MessageQueue.contract.UnpackLog(event, "DropTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1MessageQueue contract.
type L1MessageQueueInitializedIterator struct {
	Event *L1MessageQueueInitialized // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueInitialized)
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
		it.Event = new(L1MessageQueueInitialized)
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
func (it *L1MessageQueueInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueInitialized represents a Initialized event raised by the L1MessageQueue contract.
type L1MessageQueueInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1MessageQueueInitializedIterator, error) {

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueInitializedIterator{contract: _L1MessageQueue.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1MessageQueueInitialized) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueInitialized)
				if err := _L1MessageQueue.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseInitialized(log types.Log) (*L1MessageQueueInitialized, error) {
	event := new(L1MessageQueueInitialized)
	if err := _L1MessageQueue.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1MessageQueue contract.
type L1MessageQueueOwnershipTransferredIterator struct {
	Event *L1MessageQueueOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueOwnershipTransferred)
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
		it.Event = new(L1MessageQueueOwnershipTransferred)
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
func (it *L1MessageQueueOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueOwnershipTransferred represents a OwnershipTransferred event raised by the L1MessageQueue contract.
type L1MessageQueueOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1MessageQueueOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueOwnershipTransferredIterator{contract: _L1MessageQueue.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1MessageQueueOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueOwnershipTransferred)
				if err := _L1MessageQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseOwnershipTransferred(log types.Log) (*L1MessageQueueOwnershipTransferred, error) {
	event := new(L1MessageQueueOwnershipTransferred)
	if err := _L1MessageQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueQueueTransactionIterator is returned from FilterQueueTransaction and is used to iterate over the raw logs and unpacked data for QueueTransaction events raised by the L1MessageQueue contract.
type L1MessageQueueQueueTransactionIterator struct {
	Event *L1MessageQueueQueueTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueQueueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueQueueTransaction)
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
		it.Event = new(L1MessageQueueQueueTransaction)
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
func (it *L1MessageQueueQueueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueQueueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueQueueTransaction represents a QueueTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueQueueTransaction struct {
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
func (_L1MessageQueue *L1MessageQueueFilterer) FilterQueueTransaction(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*L1MessageQueueQueueTransactionIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueQueueTransactionIterator{contract: _L1MessageQueue.contract, event: "QueueTransaction", logs: logs, sub: sub}, nil
}

// WatchQueueTransaction is a free log subscription operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchQueueTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueQueueTransaction, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueQueueTransaction)
				if err := _L1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseQueueTransaction(log types.Log) (*L1MessageQueueQueueTransaction, error) {
	event := new(L1MessageQueueQueueTransaction)
	if err := _L1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueUpdateEnforcedTxGatewayIterator is returned from FilterUpdateEnforcedTxGateway and is used to iterate over the raw logs and unpacked data for UpdateEnforcedTxGateway events raised by the L1MessageQueue contract.
type L1MessageQueueUpdateEnforcedTxGatewayIterator struct {
	Event *L1MessageQueueUpdateEnforcedTxGateway // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueUpdateEnforcedTxGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueUpdateEnforcedTxGateway)
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
		it.Event = new(L1MessageQueueUpdateEnforcedTxGateway)
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
func (it *L1MessageQueueUpdateEnforcedTxGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueUpdateEnforcedTxGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueUpdateEnforcedTxGateway represents a UpdateEnforcedTxGateway event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateEnforcedTxGateway struct {
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateEnforcedTxGateway is a free log retrieval operation binding the contract event 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
//
// Solidity: event UpdateEnforcedTxGateway(address indexed _oldGateway, address indexed _newGateway)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterUpdateEnforcedTxGateway(opts *bind.FilterOpts, _oldGateway []common.Address, _newGateway []common.Address) (*L1MessageQueueUpdateEnforcedTxGatewayIterator, error) {

	var _oldGatewayRule []interface{}
	for _, _oldGatewayItem := range _oldGateway {
		_oldGatewayRule = append(_oldGatewayRule, _oldGatewayItem)
	}
	var _newGatewayRule []interface{}
	for _, _newGatewayItem := range _newGateway {
		_newGatewayRule = append(_newGatewayRule, _newGatewayItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "UpdateEnforcedTxGateway", _oldGatewayRule, _newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueUpdateEnforcedTxGatewayIterator{contract: _L1MessageQueue.contract, event: "UpdateEnforcedTxGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateEnforcedTxGateway is a free log subscription operation binding the contract event 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
//
// Solidity: event UpdateEnforcedTxGateway(address indexed _oldGateway, address indexed _newGateway)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchUpdateEnforcedTxGateway(opts *bind.WatchOpts, sink chan<- *L1MessageQueueUpdateEnforcedTxGateway, _oldGateway []common.Address, _newGateway []common.Address) (event.Subscription, error) {

	var _oldGatewayRule []interface{}
	for _, _oldGatewayItem := range _oldGateway {
		_oldGatewayRule = append(_oldGatewayRule, _oldGatewayItem)
	}
	var _newGatewayRule []interface{}
	for _, _newGatewayItem := range _newGateway {
		_newGatewayRule = append(_newGatewayRule, _newGatewayItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "UpdateEnforcedTxGateway", _oldGatewayRule, _newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueUpdateEnforcedTxGateway)
				if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateEnforcedTxGateway", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseUpdateEnforcedTxGateway(log types.Log) (*L1MessageQueueUpdateEnforcedTxGateway, error) {
	event := new(L1MessageQueueUpdateEnforcedTxGateway)
	if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateEnforcedTxGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueUpdateGasOracleIterator is returned from FilterUpdateGasOracle and is used to iterate over the raw logs and unpacked data for UpdateGasOracle events raised by the L1MessageQueue contract.
type L1MessageQueueUpdateGasOracleIterator struct {
	Event *L1MessageQueueUpdateGasOracle // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueUpdateGasOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueUpdateGasOracle)
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
		it.Event = new(L1MessageQueueUpdateGasOracle)
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
func (it *L1MessageQueueUpdateGasOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueUpdateGasOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueUpdateGasOracle represents a UpdateGasOracle event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateGasOracle struct {
	OldGasOracle common.Address
	NewGasOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasOracle is a free log retrieval operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterUpdateGasOracle(opts *bind.FilterOpts, _oldGasOracle []common.Address, _newGasOracle []common.Address) (*L1MessageQueueUpdateGasOracleIterator, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueUpdateGasOracleIterator{contract: _L1MessageQueue.contract, event: "UpdateGasOracle", logs: logs, sub: sub}, nil
}

// WatchUpdateGasOracle is a free log subscription operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchUpdateGasOracle(opts *bind.WatchOpts, sink chan<- *L1MessageQueueUpdateGasOracle, _oldGasOracle []common.Address, _newGasOracle []common.Address) (event.Subscription, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueUpdateGasOracle)
				if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseUpdateGasOracle(log types.Log) (*L1MessageQueueUpdateGasOracle, error) {
	event := new(L1MessageQueueUpdateGasOracle)
	if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueUpdateMaxGasLimitIterator is returned from FilterUpdateMaxGasLimit and is used to iterate over the raw logs and unpacked data for UpdateMaxGasLimit events raised by the L1MessageQueue contract.
type L1MessageQueueUpdateMaxGasLimitIterator struct {
	Event *L1MessageQueueUpdateMaxGasLimit // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueUpdateMaxGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueUpdateMaxGasLimit)
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
		it.Event = new(L1MessageQueueUpdateMaxGasLimit)
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
func (it *L1MessageQueueUpdateMaxGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueUpdateMaxGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueUpdateMaxGasLimit represents a UpdateMaxGasLimit event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateMaxGasLimit struct {
	OldMaxGasLimit *big.Int
	NewMaxGasLimit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxGasLimit is a free log retrieval operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterUpdateMaxGasLimit(opts *bind.FilterOpts) (*L1MessageQueueUpdateMaxGasLimitIterator, error) {

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueUpdateMaxGasLimitIterator{contract: _L1MessageQueue.contract, event: "UpdateMaxGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxGasLimit is a free log subscription operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchUpdateMaxGasLimit(opts *bind.WatchOpts, sink chan<- *L1MessageQueueUpdateMaxGasLimit) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueUpdateMaxGasLimit)
				if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseUpdateMaxGasLimit(log types.Log) (*L1MessageQueueUpdateMaxGasLimit, error) {
	event := new(L1MessageQueueUpdateMaxGasLimit)
	if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
