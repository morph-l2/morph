// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/event"
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
	_ = abi.ConvertType
)

// EnforcedTxGatewayMetaData contains all meta data concerning the EnforcedTxGateway contract.
var EnforcedTxGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_queue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611f65806100e65f395ff3fe6080604052600436106100d9575f3560e01c8063715018a61161007c5780638da5cb5b116100575780638da5cb5b1461025c578063bedb86fb14610286578063f2fde38b146102a5578063fb403d7c146102c4575f80fd5b8063715018a6146101f65780637ecebe001461020a57806384b0196e14610235575f80fd5b80633b70c18a116100b75780633b70c18a14610138578063478222c214610189578063485cc955146101b55780635c975abb146101d4575f80fd5b80632a6cccb2146100dd5780633644e515146100fe5780633934ce9d14610125575b5f80fd5b3480156100e8575f80fd5b506100fc6100f73660046119ac565b6102d7565b005b348015610109575f80fd5b506101126103bd565b6040519081526020015b60405180910390f35b6100fc610133366004611a0a565b6103cb565b348015610143575f80fd5b5060fd546101649073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161011c565b348015610194575f80fd5b5060fe546101649073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101c0575f80fd5b506100fc6101cf366004611a6d565b61045e565b3480156101df575f80fd5b5060975460ff16604051901515815260200161011c565b348015610201575f80fd5b506100fc610736565b348015610215575f80fd5b506101126102243660046119ac565b60ff6020525f908152604090205481565b348015610240575f80fd5b50610249610749565b60405161011c9796959493929190611ae1565b348015610267575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610164565b348015610291575f80fd5b506100fc6102a0366004611ba1565b610806565b3480156102b0575f80fd5b506100fc6102bf3660046119ac565b610827565b6100fc6102d2366004611bed565b6108c1565b6102df610ae0565b73ffffffffffffffffffffffffffffffffffffffff81166103475760405162461bcd60e51b815260206004820152601e60248201527f666565207661756c742063616e6e6f742062652061646472657373283029000060448201526064015b60405180910390fd5b60fe805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5905f90a35050565b5f6103c6610b47565b905090565b6103d3610b50565b3332146104485760405162461bcd60e51b815260206004820152603960248201527f4f6e6c7920454f412073656e646572732061726520616c6c6f77656420746f2060448201527f73656e6420656e666f72636564207472616e73616374696f6e00000000000000606482015260840161033e565b61045733868686868633610ba3565b5050505050565b5f54610100900460ff161580801561047c57505f54600160ff909116105b806104955750303b15801561049557505f5460ff166001145b6105075760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161033e565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610563575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff82166105c65760405162461bcd60e51b815260206004820152601e60248201527f666565207661756c742063616e6e6f7420626520616464726573732830290000604482015260640161033e565b6105ce610ea7565b6105d6610f2b565b6105de610faf565b6106526040518060400160405280601181526020017f456e666f726365645478476174657761790000000000000000000000000000008152506040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250611033565b60fd805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fe805492851692909116821790556040515f907f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5908290a38015610731575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b61073e610ae0565b6107475f6110bd565b565b5f6060805f805f606060c9545f801b148015610765575060ca54155b6107b15760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161033e565b6107b9611133565b6107c16111c3565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b61080e610ae0565b801561081f5761081c6111d2565b50565b61081c611257565b61082f610ae0565b73ffffffffffffffffffffffffffffffffffffffff81166108b85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161033e565b61081c816110bd565b6108c9610b50565b824211156109195760405162461bcd60e51b815260206004820152601160248201527f7369676e61747572652065787069726564000000000000000000000000000000604482015260640161033e565b5f60ff5f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f7f302d96da71d942dc3052ca041618b37fc929a10d432f3a337b5be0a8385c9afe8b8b8b8b8b8b604051610990929190611d0c565b60408051918290038220602083019790975273ffffffffffffffffffffffffffffffffffffffff95861690820152939092166060840152608083015260a082015260c081019190915260e0810183905261010081018690526101200160408051601f19818403018152918152815160209283012073ffffffffffffffffffffffffffffffffffffffff8e165f90815260ff9093529082206001850190559150610a38826112ae565b90505f610a4582876112fb565b90508c73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ac25760405162461bcd60e51b815260206004820152601360248201527f496e636f7272656374207369676e617475726500000000000000000000000000604482015260640161033e565b610ad18d8d8d8d8d8d8b610ba3565b50505050505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146107475760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161033e565b5f6103c661131d565b60975460ff16156107475760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161033e565b610bab611390565b60fd546040517f3e4cbbe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff898116600483015260248201879052909116905f908290633e4cbbe690604401602060405180830381865afa158015610c24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c489190611d1b565b905080341015610c9a5760405162461bcd60e51b815260206004820152601a60248201527f496e73756666696369656e742076616c756520666f7220666565000000000000604482015260640161033e565b8015610d4f5760fe546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114610cf7576040519150601f19603f3d011682016040523d82523d5f602084013e610cfc565b606091505b5050905080610d4d5760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f2064656475637420746865206665650000000000000000604482015260640161033e565b505b6040517fbdc6f0a000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063bdc6f0a090610dab908c908c908c908c908c908c90600401611d32565b5f604051808303815f87803b158015610dc2575f80fd5b505af1158015610dd4573d5f803e3d5ffd5b5050503482810391508214610e91575f8473ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610e39576040519150601f19603f3d011682016040523d82523d5f602084013e610e3e565b606091505b5050905080610e8f5760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f20726566756e6420746865206665650000000000000000604482015260640161033e565b505b505050610e9e6001606555565b50505050505050565b5f54610100900460ff16610f235760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b6107476113f0565b5f54610100900460ff16610fa75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b610747611475565b5f54610100900460ff1661102b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b6107476114f1565b5f54610100900460ff166110af5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b6110b98282611597565b5050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b606060cb805461114290611d98565b80601f016020809104026020016040519081016040528092919081815260200182805461116e90611d98565b80156111b95780601f10611190576101008083540402835291602001916111b9565b820191905f5260205f20905b81548152906001019060200180831161119c57829003601f168201915b5050505050905090565b606060cc805461114290611d98565b6111da610b50565b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861122d3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b61125f61163a565b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa3361122d565b5f6112f56112ba610b47565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b92915050565b5f805f611308858561168c565b91509150611315816116ce565b509392505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611347611832565b61134f61188a565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6002606554036113e25760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161033e565b6002606555565b6001606555565b5f54610100900460ff1661146c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b610747336110bd565b5f54610100900460ff166113e95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b5f54610100900460ff1661156d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b5f54610100900460ff166116135760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161033e565b60cb61161f8382611e2d565b5060cc61162c8282611e2d565b50505f60c981905560ca5550565b60975460ff166107475760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161033e565b5f8082516041036116c0576020830151604084015160608501515f1a6116b4878285856118ba565b945094505050506116c7565b505f905060025b9250929050565b5f8160048111156116e1576116e1611f2b565b036116e95750565b60018160048111156116fd576116fd611f2b565b0361174a5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161033e565b600281600481111561175e5761175e611f2b565b036117ab5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161033e565b60038160048111156117bf576117bf611f2b565b0361081c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161033e565b5f8061183c611133565b805190915015611853578051602090910120919050565b60c95480156118625792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b5f806118946111c3565b8051909150156118ab578051602090910120919050565b60ca5480156118625792915050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156118ef57505f9050600361197b565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611940573d5f803e3d5ffd5b5050604051601f19015191505073ffffffffffffffffffffffffffffffffffffffff8116611975575f6001925092505061197b565b91505f90505b94509492505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146119a7575f80fd5b919050565b5f602082840312156119bc575f80fd5b6119c582611984565b9392505050565b5f8083601f8401126119dc575f80fd5b50813567ffffffffffffffff8111156119f3575f80fd5b6020830191508360208285010111156116c7575f80fd5b5f805f805f60808688031215611a1e575f80fd5b611a2786611984565b94506020860135935060408601359250606086013567ffffffffffffffff811115611a50575f80fd5b611a5c888289016119cc565b969995985093965092949392505050565b5f8060408385031215611a7e575f80fd5b611a8783611984565b9150611a9560208401611984565b90509250929050565b5f81518084525f5b81811015611ac257602081850181015186830182015201611aa6565b505f602082860101526020601f19601f83011685010191505092915050565b7fff00000000000000000000000000000000000000000000000000000000000000881681525f602060e06020840152611b1d60e084018a611a9e565b8381036040850152611b2f818a611a9e565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b81811015611b8f57835183529284019291840191600101611b73565b50909c9b505050505050505050505050565b5f60208284031215611bb1575f80fd5b813580151581146119c5575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f805f805f805f6101008a8c031215611c06575f80fd5b611c0f8a611984565b9850611c1d60208b01611984565b975060408a0135965060608a0135955060808a013567ffffffffffffffff80821115611c47575f80fd5b611c538d838e016119cc565b909750955060a08c0135945060c08c0135915080821115611c72575f80fd5b818c0191508c601f830112611c85575f80fd5b813581811115611c9757611c97611bc0565b604051601f8201601f19908116603f01168101908382118183101715611cbf57611cbf611bc0565b816040528281528f6020848701011115611cd7575f80fd5b826020860160208301375f602084830101528096505050505050611cfd60e08b01611984565b90509295985092959850929598565b818382375f9101908152919050565b5f60208284031215611d2b575f80fd5b5051919050565b5f73ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525085604083015284606083015260a060808301528260a0830152828460c08401375f60c0848401015260c0601f19601f8501168301019050979650505050505050565b600181811c90821680611dac57607f821691505b602082108103611de3577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f82111561073157805f5260205f20601f840160051c81016020851015611e0e5750805b601f840160051c820191505b81811015610457575f8155600101611e1a565b815167ffffffffffffffff811115611e4757611e47611bc0565b611e5b81611e558454611d98565b84611de9565b602080601f831160018114611ead575f8415611e775750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611f23565b5f85815260208120601f198616915b82811015611edb57888601518255948401946001909101908401611ebc565b5085821015611f1757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea164736f6c6343000818000a",
}

// EnforcedTxGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use EnforcedTxGatewayMetaData.ABI instead.
var EnforcedTxGatewayABI = EnforcedTxGatewayMetaData.ABI

// EnforcedTxGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnforcedTxGatewayMetaData.Bin instead.
var EnforcedTxGatewayBin = EnforcedTxGatewayMetaData.Bin

// DeployEnforcedTxGateway deploys a new Ethereum contract, binding an instance of EnforcedTxGateway to it.
func DeployEnforcedTxGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnforcedTxGateway, error) {
	parsed, err := EnforcedTxGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnforcedTxGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnforcedTxGateway{EnforcedTxGatewayCaller: EnforcedTxGatewayCaller{contract: contract}, EnforcedTxGatewayTransactor: EnforcedTxGatewayTransactor{contract: contract}, EnforcedTxGatewayFilterer: EnforcedTxGatewayFilterer{contract: contract}}, nil
}

// EnforcedTxGateway is an auto generated Go binding around an Ethereum contract.
type EnforcedTxGateway struct {
	EnforcedTxGatewayCaller     // Read-only binding to the contract
	EnforcedTxGatewayTransactor // Write-only binding to the contract
	EnforcedTxGatewayFilterer   // Log filterer for contract events
}

// EnforcedTxGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnforcedTxGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnforcedTxGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnforcedTxGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnforcedTxGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnforcedTxGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnforcedTxGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnforcedTxGatewaySession struct {
	Contract     *EnforcedTxGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EnforcedTxGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnforcedTxGatewayCallerSession struct {
	Contract *EnforcedTxGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EnforcedTxGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnforcedTxGatewayTransactorSession struct {
	Contract     *EnforcedTxGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EnforcedTxGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnforcedTxGatewayRaw struct {
	Contract *EnforcedTxGateway // Generic contract binding to access the raw methods on
}

// EnforcedTxGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnforcedTxGatewayCallerRaw struct {
	Contract *EnforcedTxGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// EnforcedTxGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnforcedTxGatewayTransactorRaw struct {
	Contract *EnforcedTxGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnforcedTxGateway creates a new instance of EnforcedTxGateway, bound to a specific deployed contract.
func NewEnforcedTxGateway(address common.Address, backend bind.ContractBackend) (*EnforcedTxGateway, error) {
	contract, err := bindEnforcedTxGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGateway{EnforcedTxGatewayCaller: EnforcedTxGatewayCaller{contract: contract}, EnforcedTxGatewayTransactor: EnforcedTxGatewayTransactor{contract: contract}, EnforcedTxGatewayFilterer: EnforcedTxGatewayFilterer{contract: contract}}, nil
}

// NewEnforcedTxGatewayCaller creates a new read-only instance of EnforcedTxGateway, bound to a specific deployed contract.
func NewEnforcedTxGatewayCaller(address common.Address, caller bind.ContractCaller) (*EnforcedTxGatewayCaller, error) {
	contract, err := bindEnforcedTxGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayCaller{contract: contract}, nil
}

// NewEnforcedTxGatewayTransactor creates a new write-only instance of EnforcedTxGateway, bound to a specific deployed contract.
func NewEnforcedTxGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*EnforcedTxGatewayTransactor, error) {
	contract, err := bindEnforcedTxGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayTransactor{contract: contract}, nil
}

// NewEnforcedTxGatewayFilterer creates a new log filterer instance of EnforcedTxGateway, bound to a specific deployed contract.
func NewEnforcedTxGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*EnforcedTxGatewayFilterer, error) {
	contract, err := bindEnforcedTxGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayFilterer{contract: contract}, nil
}

// bindEnforcedTxGateway binds a generic wrapper to an already deployed contract.
func bindEnforcedTxGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnforcedTxGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnforcedTxGateway *EnforcedTxGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnforcedTxGateway.Contract.EnforcedTxGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnforcedTxGateway *EnforcedTxGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.EnforcedTxGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnforcedTxGateway *EnforcedTxGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.EnforcedTxGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnforcedTxGateway *EnforcedTxGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnforcedTxGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EnforcedTxGateway.Contract.DOMAINSEPARATOR(&_EnforcedTxGateway.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EnforcedTxGateway.Contract.DOMAINSEPARATOR(&_EnforcedTxGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EnforcedTxGateway.Contract.Eip712Domain(&_EnforcedTxGateway.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EnforcedTxGateway.Contract.Eip712Domain(&_EnforcedTxGateway.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) FeeVault() (common.Address, error) {
	return _EnforcedTxGateway.Contract.FeeVault(&_EnforcedTxGateway.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) FeeVault() (common.Address, error) {
	return _EnforcedTxGateway.Contract.FeeVault(&_EnforcedTxGateway.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) MessageQueue() (common.Address, error) {
	return _EnforcedTxGateway.Contract.MessageQueue(&_EnforcedTxGateway.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) MessageQueue() (common.Address, error) {
	return _EnforcedTxGateway.Contract.MessageQueue(&_EnforcedTxGateway.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _EnforcedTxGateway.Contract.Nonces(&_EnforcedTxGateway.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _EnforcedTxGateway.Contract.Nonces(&_EnforcedTxGateway.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) Owner() (common.Address, error) {
	return _EnforcedTxGateway.Contract.Owner(&_EnforcedTxGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) Owner() (common.Address, error) {
	return _EnforcedTxGateway.Contract.Owner(&_EnforcedTxGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EnforcedTxGateway *EnforcedTxGatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EnforcedTxGateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EnforcedTxGateway *EnforcedTxGatewaySession) Paused() (bool, error) {
	return _EnforcedTxGateway.Contract.Paused(&_EnforcedTxGateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EnforcedTxGateway *EnforcedTxGatewayCallerSession) Paused() (bool, error) {
	return _EnforcedTxGateway.Contract.Paused(&_EnforcedTxGateway.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _queue, address _feeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) Initialize(opts *bind.TransactOpts, _queue common.Address, _feeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "initialize", _queue, _feeVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _queue, address _feeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) Initialize(_queue common.Address, _feeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.Initialize(&_EnforcedTxGateway.TransactOpts, _queue, _feeVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _queue, address _feeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) Initialize(_queue common.Address, _feeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.Initialize(&_EnforcedTxGateway.TransactOpts, _queue, _feeVault)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.RenounceOwnership(&_EnforcedTxGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.RenounceOwnership(&_EnforcedTxGateway.TransactOpts)
}

// SendTransaction is a paid mutator transaction binding the contract method 0x3934ce9d.
//
// Solidity: function sendTransaction(address _target, uint256 _value, uint256 _gasLimit, bytes _data) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) SendTransaction(opts *bind.TransactOpts, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "sendTransaction", _target, _value, _gasLimit, _data)
}

// SendTransaction is a paid mutator transaction binding the contract method 0x3934ce9d.
//
// Solidity: function sendTransaction(address _target, uint256 _value, uint256 _gasLimit, bytes _data) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) SendTransaction(_target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SendTransaction(&_EnforcedTxGateway.TransactOpts, _target, _value, _gasLimit, _data)
}

// SendTransaction is a paid mutator transaction binding the contract method 0x3934ce9d.
//
// Solidity: function sendTransaction(address _target, uint256 _value, uint256 _gasLimit, bytes _data) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) SendTransaction(_target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SendTransaction(&_EnforcedTxGateway.TransactOpts, _target, _value, _gasLimit, _data)
}

// SendTransaction0 is a paid mutator transaction binding the contract method 0xfb403d7c.
//
// Solidity: function sendTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data, uint256 _deadline, bytes _signature, address _refundAddress) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) SendTransaction0(opts *bind.TransactOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte, _deadline *big.Int, _signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "sendTransaction0", _sender, _target, _value, _gasLimit, _data, _deadline, _signature, _refundAddress)
}

// SendTransaction0 is a paid mutator transaction binding the contract method 0xfb403d7c.
//
// Solidity: function sendTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data, uint256 _deadline, bytes _signature, address _refundAddress) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) SendTransaction0(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte, _deadline *big.Int, _signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SendTransaction0(&_EnforcedTxGateway.TransactOpts, _sender, _target, _value, _gasLimit, _data, _deadline, _signature, _refundAddress)
}

// SendTransaction0 is a paid mutator transaction binding the contract method 0xfb403d7c.
//
// Solidity: function sendTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data, uint256 _deadline, bytes _signature, address _refundAddress) payable returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) SendTransaction0(_sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte, _deadline *big.Int, _signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SendTransaction0(&_EnforcedTxGateway.TransactOpts, _sender, _target, _value, _gasLimit, _data, _deadline, _signature, _refundAddress)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) SetPause(_status bool) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SetPause(&_EnforcedTxGateway.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.SetPause(&_EnforcedTxGateway.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.TransferOwnership(&_EnforcedTxGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.TransferOwnership(&_EnforcedTxGateway.TransactOpts, newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewaySession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.UpdateFeeVault(&_EnforcedTxGateway.TransactOpts, _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_EnforcedTxGateway *EnforcedTxGatewayTransactorSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _EnforcedTxGateway.Contract.UpdateFeeVault(&_EnforcedTxGateway.TransactOpts, _newFeeVault)
}

// EnforcedTxGatewayEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayEIP712DomainChangedIterator struct {
	Event *EnforcedTxGatewayEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayEIP712DomainChanged)
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
		it.Event = new(EnforcedTxGatewayEIP712DomainChanged)
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
func (it *EnforcedTxGatewayEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayEIP712DomainChanged represents a EIP712DomainChanged event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EnforcedTxGatewayEIP712DomainChangedIterator, error) {

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayEIP712DomainChangedIterator{contract: _EnforcedTxGateway.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayEIP712DomainChanged)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParseEIP712DomainChanged(log types.Log) (*EnforcedTxGatewayEIP712DomainChanged, error) {
	event := new(EnforcedTxGatewayEIP712DomainChanged)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnforcedTxGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayInitializedIterator struct {
	Event *EnforcedTxGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayInitialized)
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
		it.Event = new(EnforcedTxGatewayInitialized)
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
func (it *EnforcedTxGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayInitialized represents a Initialized event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*EnforcedTxGatewayInitializedIterator, error) {

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayInitializedIterator{contract: _EnforcedTxGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayInitialized)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParseInitialized(log types.Log) (*EnforcedTxGatewayInitialized, error) {
	event := new(EnforcedTxGatewayInitialized)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnforcedTxGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayOwnershipTransferredIterator struct {
	Event *EnforcedTxGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayOwnershipTransferred)
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
		it.Event = new(EnforcedTxGatewayOwnershipTransferred)
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
func (it *EnforcedTxGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnforcedTxGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayOwnershipTransferredIterator{contract: _EnforcedTxGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayOwnershipTransferred)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*EnforcedTxGatewayOwnershipTransferred, error) {
	event := new(EnforcedTxGatewayOwnershipTransferred)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnforcedTxGatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayPausedIterator struct {
	Event *EnforcedTxGatewayPaused // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayPaused)
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
		it.Event = new(EnforcedTxGatewayPaused)
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
func (it *EnforcedTxGatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayPaused represents a Paused event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*EnforcedTxGatewayPausedIterator, error) {

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayPausedIterator{contract: _EnforcedTxGateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayPaused) (event.Subscription, error) {

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayPaused)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParsePaused(log types.Log) (*EnforcedTxGatewayPaused, error) {
	event := new(EnforcedTxGatewayPaused)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnforcedTxGatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayUnpausedIterator struct {
	Event *EnforcedTxGatewayUnpaused // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayUnpaused)
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
		it.Event = new(EnforcedTxGatewayUnpaused)
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
func (it *EnforcedTxGatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayUnpaused represents a Unpaused event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EnforcedTxGatewayUnpausedIterator, error) {

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayUnpausedIterator{contract: _EnforcedTxGateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayUnpaused)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParseUnpaused(log types.Log) (*EnforcedTxGatewayUnpaused, error) {
	event := new(EnforcedTxGatewayUnpaused)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnforcedTxGatewayUpdateFeeVaultIterator is returned from FilterUpdateFeeVault and is used to iterate over the raw logs and unpacked data for UpdateFeeVault events raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayUpdateFeeVaultIterator struct {
	Event *EnforcedTxGatewayUpdateFeeVault // Event containing the contract specifics and raw log

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
func (it *EnforcedTxGatewayUpdateFeeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnforcedTxGatewayUpdateFeeVault)
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
		it.Event = new(EnforcedTxGatewayUpdateFeeVault)
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
func (it *EnforcedTxGatewayUpdateFeeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnforcedTxGatewayUpdateFeeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnforcedTxGatewayUpdateFeeVault represents a UpdateFeeVault event raised by the EnforcedTxGateway contract.
type EnforcedTxGatewayUpdateFeeVault struct {
	OldFeeVault common.Address
	NewFeeVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeVault is a free log retrieval operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address indexed _oldFeeVault, address indexed _newFeeVault)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) FilterUpdateFeeVault(opts *bind.FilterOpts, _oldFeeVault []common.Address, _newFeeVault []common.Address) (*EnforcedTxGatewayUpdateFeeVaultIterator, error) {

	var _oldFeeVaultRule []interface{}
	for _, _oldFeeVaultItem := range _oldFeeVault {
		_oldFeeVaultRule = append(_oldFeeVaultRule, _oldFeeVaultItem)
	}
	var _newFeeVaultRule []interface{}
	for _, _newFeeVaultItem := range _newFeeVault {
		_newFeeVaultRule = append(_newFeeVaultRule, _newFeeVaultItem)
	}

	logs, sub, err := _EnforcedTxGateway.contract.FilterLogs(opts, "UpdateFeeVault", _oldFeeVaultRule, _newFeeVaultRule)
	if err != nil {
		return nil, err
	}
	return &EnforcedTxGatewayUpdateFeeVaultIterator{contract: _EnforcedTxGateway.contract, event: "UpdateFeeVault", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeVault is a free log subscription operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address indexed _oldFeeVault, address indexed _newFeeVault)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) WatchUpdateFeeVault(opts *bind.WatchOpts, sink chan<- *EnforcedTxGatewayUpdateFeeVault, _oldFeeVault []common.Address, _newFeeVault []common.Address) (event.Subscription, error) {

	var _oldFeeVaultRule []interface{}
	for _, _oldFeeVaultItem := range _oldFeeVault {
		_oldFeeVaultRule = append(_oldFeeVaultRule, _oldFeeVaultItem)
	}
	var _newFeeVaultRule []interface{}
	for _, _newFeeVaultItem := range _newFeeVault {
		_newFeeVaultRule = append(_newFeeVaultRule, _newFeeVaultItem)
	}

	logs, sub, err := _EnforcedTxGateway.contract.WatchLogs(opts, "UpdateFeeVault", _oldFeeVaultRule, _newFeeVaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnforcedTxGatewayUpdateFeeVault)
				if err := _EnforcedTxGateway.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
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

// ParseUpdateFeeVault is a log parse operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address indexed _oldFeeVault, address indexed _newFeeVault)
func (_EnforcedTxGateway *EnforcedTxGatewayFilterer) ParseUpdateFeeVault(log types.Log) (*EnforcedTxGatewayUpdateFeeVault, error) {
	event := new(EnforcedTxGatewayUpdateFeeVault)
	if err := _EnforcedTxGateway.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
