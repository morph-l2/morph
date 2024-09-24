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

// IMorphTokenEpochInflationRate is an auto generated low-level Go binding around an user-defined struct.
type IMorphTokenEpochInflationRate struct {
	Rate                *big.Int
	EffectiveEpochIndex *big.Int
}

// MorphTokenMetaData contains all meta data concerning the MorphToken contract.
var MorphTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InflationMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveEpochIndex\",\"type\":\"uint256\"}],\"name\":\"UpdateEpochInflationRate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECORD_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"epochInflationRates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpochIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMorphToken.EpochInflationRate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"inflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflationMintedEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflationRatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyInflationRate_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upToEpochIndex\",\"type\":\"uint256\"}],\"name\":\"mintInflations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpochIndex\",\"type\":\"uint256\"}],\"name\":\"updateRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5073530000000000000000000000000000000000001560805273530000000000000000000000000000000000001460a05273530000000000000000000000000000000000001260c05260805160a05160c051611bb76100975f395f81816103d20152610a7a01525f81816102330152610d3501525f81816103250152610aff0152611bb75ff3fe608060405234801561000f575f80fd5b506004361061019a575f3560e01c8063715018a6116100e8578063a29bfb2c11610093578063c553f7b31161006e578063c553f7b3146103c5578063cd4281d0146103cd578063dd62ed3e146103f4578063f2fde38b14610439575f80fd5b8063a29bfb2c1461038c578063a457c2d71461039f578063a9059cbb146103b2575f80fd5b80638da5cb5b116100c35780638da5cb5b14610347578063944fa7461461036557806395d89b4114610384575f80fd5b8063715018a614610305578063748231321461030d578063807de44314610320575f80fd5b8063395093511161014857806342966c681161012357806342966c681461028f5780636d0c4a26146102a257806370a08231146102d0575f80fd5b8063395093511461021b5780633d9353fe1461022e578063405abb411461027a575f80fd5b806318160ddd1161017857806318160ddd146101f157806323b872dd146101f9578063313ce5671461020c575f80fd5b806306fdde031461019e578063095ea7b3146101bc5780630b88a984146101df575b5f80fd5b6101a661044c565b6040516101b391906115d5565b60405180910390f35b6101cf6101ca366004611667565b6104dc565b60405190151581526020016101b3565b606c545b6040519081526020016101b3565b6067546101e3565b6101cf61020736600461168f565b6104f5565b604051601281526020016101b3565b6101cf610229366004611667565b610518565b6102557f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b3565b61028d6102883660046116c8565b610563565b005b61028d61029d3660046116e8565b61075e565b6102b56102b03660046116e8565b6107c2565b604080518251815260209283015192810192909252016101b3565b6101e36102de3660046116ff565b73ffffffffffffffffffffffffffffffffffffffff165f9081526068602052604090205490565b61028d610819565b61028d61031b3660046117f3565b61082c565b6102557f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff16610255565b6101e36103733660046116e8565b5f908152606b602052604090205490565b6101a6610a68565b61028d61039a3660046116e8565b610a77565b6101cf6103ad366004611667565b610dbb565b6101cf6103c0366004611667565b610e4b565b606a546101e3565b6102557f000000000000000000000000000000000000000000000000000000000000000081565b6101e3610402366004611873565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260696020908152604080832093909416825291909152205490565b61028d6104473660046116ff565b610e58565b60606065805461045b906118a4565b80601f0160208091040260200160405190810160405280929190818152602001828054610487906118a4565b80156104d25780601f106104a9576101008083540402835291602001916104d2565b820191905f5260205f20905b8154815290600101906020018083116104b557829003601f168201915b5050505050905090565b5f336104e9818585610ef2565b60019150505b92915050565b5f33610502858285611026565b61050d8585856110e2565b506001949350505050565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104e9908290869061055e908790611922565b610ef2565b61056b611297565b606a805483919061057e90600190611935565b8154811061058e5761058e611948565b905f5260205f2090600202015f0154036106155760405162461bcd60e51b815260206004820152602760248201527f6e65772072617465206973207468652073616d6520617320746865206c61746560448201527f737420726174650000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b606a805461062590600190611935565b8154811061063557610635611948565b905f5260205f2090600202016001015481116106b95760405162461bcd60e51b815260206004820152603260248201527f6566666563746976652065706f636873206166746572206d757374206265206760448201527f726561746572207468616e206265666f72650000000000000000000000000000606482015260840161060c565b60408051808201825283815260208101838152606a80546001810182555f91825292517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5160029094029384015590517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a52909201919091559051829184917fbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f49190a35050565b610766611297565b5f81116107b55760405162461bcd60e51b815260206004820152601660248201527f616d6f756e7420746f206275726e206973207a65726f00000000000000000000604482015260640161060c565b6107bf33826112fe565b50565b604080518082019091525f8082526020820152606a82815481106107e8576107e8611948565b905f5260205f2090600202016040518060400160405290815f82015481526020016001820154815250509050919050565b610821611297565b61082a5f611486565b565b5f54610100900460ff161580801561084a57505f54600160ff909116105b806108635750303b15801561086357505f5460ff166001145b6108d55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161060c565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610931575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606561093d87826119c1565b50606661094a86826119c1565b5061095584846114fc565b604080518082019091528281525f60208201818152606a8054600181018255925291517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5160029092029182015590517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a52909101556109d284611486565b6040515f9083907fbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f4908390a38015610a60575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60606066805461045b906118a4565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610afc5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c79207265636f726420636f6e747261637420616c6c6f77656400000000604482015260640161060c565b807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663766718086040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b66573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b8a9190611ad9565b11610bfd5760405162461bcd60e51b815260206004820152602b60248201527f746865207370656369666965642074696d6520686173206e6f7420796574206260448201527f65656e2072656163686564000000000000000000000000000000000000000000606482015260840161060c565b606c54811015610c4f5760405162461bcd60e51b815260206004820152601560248201527f616c6c20696e666c6174696f6e73206d696e7465640000000000000000000000604482015260640161060c565b606c545b818111610da9575f606a5f81548110610c6e57610c6e611948565b5f9182526020822060029091020154606a54909250610c8f90600190611935565b90505b8015610cfc5782606a8281548110610cac57610cac611948565b905f5260205f2090600202016001015411610cea57606a8181548110610cd457610cd4611948565b905f5260205f2090600202015f01549150610cfc565b80610cf481611af0565b915050610c92565b505f662386f26fc1000082606754610d149190611b24565b610d1e9190611b3b565b5f848152606b602052604090208190559050610d5a7f0000000000000000000000000000000000000000000000000000000000000000826114fc565b827f0d82c0920038b8dc7f633e18585f37092ba957b84876fcf833d6841f69eaa32782604051610d8c91815260200190565b60405180910390a250508080610da190611b73565b915050610c53565b50610db5816001611922565b606c5550565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610e3e5760405162461bcd60e51b815260206004820152601e60248201527f64656372656173656420616c6c6f77616e63652062656c6f77207a65726f0000604482015260640161060c565b61050d8286868403610ef2565b5f336104e98185856110e2565b610e60611297565b73ffffffffffffffffffffffffffffffffffffffff8116610ee95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161060c565b6107bf81611486565b73ffffffffffffffffffffffffffffffffffffffff8316610f555760405162461bcd60e51b815260206004820152601d60248201527f617070726f76652066726f6d20746865207a65726f2061646472657373000000604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff8216610fb85760405162461bcd60e51b815260206004820152601b60248201527f617070726f766520746f20746865207a65726f20616464726573730000000000604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526069602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152606960209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146110dc57818110156110cf5760405162461bcd60e51b815260206004820152601660248201527f696e73756666696369656e7420616c6c6f77616e636500000000000000000000604482015260640161060c565b6110dc8484848403610ef2565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166111455760405162461bcd60e51b815260206004820152601e60248201527f7472616e736665722066726f6d20746865207a65726f20616464726573730000604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff82166111a85760405162461bcd60e51b815260206004820152601c60248201527f7472616e7366657220746f20746865207a65726f206164647265737300000000604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff83165f908152606860205260409020548181101561121d5760405162461bcd60e51b815260206004820152601f60248201527f7472616e7366657220616d6f756e7420657863656564732062616c616e636500604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff8085165f8181526068602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906112899086815260200190565b60405180910390a350505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461082a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161060c565b73ffffffffffffffffffffffffffffffffffffffff82166113875760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161060c565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260686020526040902054818110156114225760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161060c565b73ffffffffffffffffffffffffffffffffffffffff83165f8181526068602090815260408083208686039055606780548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611019565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff821661155f5760405162461bcd60e51b815260206004820152601860248201527f6d696e7420746f20746865207a65726f20616464726573730000000000000000604482015260640161060c565b8060675f8282546115709190611922565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152606860209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b5f602080835283518060208501525f5b81811015611601578581018301518582016040015282016115e5565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611662575f80fd5b919050565b5f8060408385031215611678575f80fd5b6116818361163f565b946020939093013593505050565b5f805f606084860312156116a1575f80fd5b6116aa8461163f565b92506116b86020850161163f565b9150604084013590509250925092565b5f80604083850312156116d9575f80fd5b50508035926020909101359150565b5f602082840312156116f8575f80fd5b5035919050565b5f6020828403121561170f575f80fd5b6117188261163f565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f83011261175b575f80fd5b813567ffffffffffffffff808211156117765761177661171f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156117bc576117bc61171f565b816040528381528660208588010111156117d4575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a08688031215611807575f80fd5b853567ffffffffffffffff8082111561181e575f80fd5b61182a89838a0161174c565b9650602088013591508082111561183f575f80fd5b5061184c8882890161174c565b94505061185b6040870161163f565b94979396509394606081013594506080013592915050565b5f8060408385031215611884575f80fd5b61188d8361163f565b915061189b6020840161163f565b90509250929050565b600181811c908216806118b857607f821691505b6020821081036118ef577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156104ef576104ef6118f5565b818103818111156104ef576104ef6118f5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b601f8211156119bc57805f5260205f20601f840160051c8101602085101561199a5750805b601f840160051c820191505b818110156119b9575f81556001016119a6565b50505b505050565b815167ffffffffffffffff8111156119db576119db61171f565b6119ef816119e984546118a4565b84611975565b602080601f831160018114611a41575f8415611a0b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610a60565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611a8d57888601518255948401946001909101908401611a6e565b5085821015611ac957878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b5f60208284031215611ae9575f80fd5b5051919050565b5f81611afe57611afe6118f5565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b80820281158282048414176104ef576104ef6118f5565b5f82611b6e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ba357611ba36118f5565b506001019056fea164736f6c6343000818000a",
}

// MorphTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphTokenMetaData.ABI instead.
var MorphTokenABI = MorphTokenMetaData.ABI

// MorphTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MorphTokenMetaData.Bin instead.
var MorphTokenBin = MorphTokenMetaData.Bin

// DeployMorphToken deploys a new Ethereum contract, binding an instance of MorphToken to it.
func DeployMorphToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MorphToken, error) {
	parsed, err := MorphTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MorphTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MorphToken{MorphTokenCaller: MorphTokenCaller{contract: contract}, MorphTokenTransactor: MorphTokenTransactor{contract: contract}, MorphTokenFilterer: MorphTokenFilterer{contract: contract}}, nil
}

// MorphToken is an auto generated Go binding around an Ethereum contract.
type MorphToken struct {
	MorphTokenCaller     // Read-only binding to the contract
	MorphTokenTransactor // Write-only binding to the contract
	MorphTokenFilterer   // Log filterer for contract events
}

// MorphTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphTokenSession struct {
	Contract     *MorphToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphTokenCallerSession struct {
	Contract *MorphTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MorphTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphTokenTransactorSession struct {
	Contract     *MorphTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MorphTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphTokenRaw struct {
	Contract *MorphToken // Generic contract binding to access the raw methods on
}

// MorphTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphTokenCallerRaw struct {
	Contract *MorphTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MorphTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphTokenTransactorRaw struct {
	Contract *MorphTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphToken creates a new instance of MorphToken, bound to a specific deployed contract.
func NewMorphToken(address common.Address, backend bind.ContractBackend) (*MorphToken, error) {
	contract, err := bindMorphToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphToken{MorphTokenCaller: MorphTokenCaller{contract: contract}, MorphTokenTransactor: MorphTokenTransactor{contract: contract}, MorphTokenFilterer: MorphTokenFilterer{contract: contract}}, nil
}

// NewMorphTokenCaller creates a new read-only instance of MorphToken, bound to a specific deployed contract.
func NewMorphTokenCaller(address common.Address, caller bind.ContractCaller) (*MorphTokenCaller, error) {
	contract, err := bindMorphToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphTokenCaller{contract: contract}, nil
}

// NewMorphTokenTransactor creates a new write-only instance of MorphToken, bound to a specific deployed contract.
func NewMorphTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphTokenTransactor, error) {
	contract, err := bindMorphToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphTokenTransactor{contract: contract}, nil
}

// NewMorphTokenFilterer creates a new log filterer instance of MorphToken, bound to a specific deployed contract.
func NewMorphTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphTokenFilterer, error) {
	contract, err := bindMorphToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphTokenFilterer{contract: contract}, nil
}

// bindMorphToken binds a generic wrapper to an already deployed contract.
func bindMorphToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphToken *MorphTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphToken.Contract.MorphTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphToken *MorphTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphToken.Contract.MorphTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphToken *MorphTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphToken.Contract.MorphTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphToken *MorphTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphToken *MorphTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphToken *MorphTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphToken.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCaller) DISTRIBUTECONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "DISTRIBUTE_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenSession) DISTRIBUTECONTRACT() (common.Address, error) {
	return _MorphToken.Contract.DISTRIBUTECONTRACT(&_MorphToken.CallOpts)
}

// DISTRIBUTECONTRACT is a free data retrieval call binding the contract method 0x3d9353fe.
//
// Solidity: function DISTRIBUTE_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCallerSession) DISTRIBUTECONTRACT() (common.Address, error) {
	return _MorphToken.Contract.DISTRIBUTECONTRACT(&_MorphToken.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCaller) L2STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "L2_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _MorphToken.Contract.L2STAKINGCONTRACT(&_MorphToken.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCallerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _MorphToken.Contract.L2STAKINGCONTRACT(&_MorphToken.CallOpts)
}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCaller) RECORDCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "RECORD_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenSession) RECORDCONTRACT() (common.Address, error) {
	return _MorphToken.Contract.RECORDCONTRACT(&_MorphToken.CallOpts)
}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_MorphToken *MorphTokenCallerSession) RECORDCONTRACT() (common.Address, error) {
	return _MorphToken.Contract.RECORDCONTRACT(&_MorphToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphToken *MorphTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphToken *MorphTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphToken.Contract.Allowance(&_MorphToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphToken.Contract.Allowance(&_MorphToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphToken *MorphTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphToken *MorphTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphToken.Contract.BalanceOf(&_MorphToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphToken.Contract.BalanceOf(&_MorphToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MorphToken *MorphTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MorphToken *MorphTokenSession) Decimals() (uint8, error) {
	return _MorphToken.Contract.Decimals(&_MorphToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MorphToken *MorphTokenCallerSession) Decimals() (uint8, error) {
	return _MorphToken.Contract.Decimals(&_MorphToken.CallOpts)
}

// EpochInflationRates is a free data retrieval call binding the contract method 0x6d0c4a26.
//
// Solidity: function epochInflationRates(uint256 index) view returns((uint256,uint256))
func (_MorphToken *MorphTokenCaller) EpochInflationRates(opts *bind.CallOpts, index *big.Int) (IMorphTokenEpochInflationRate, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "epochInflationRates", index)

	if err != nil {
		return *new(IMorphTokenEpochInflationRate), err
	}

	out0 := *abi.ConvertType(out[0], new(IMorphTokenEpochInflationRate)).(*IMorphTokenEpochInflationRate)

	return out0, err

}

// EpochInflationRates is a free data retrieval call binding the contract method 0x6d0c4a26.
//
// Solidity: function epochInflationRates(uint256 index) view returns((uint256,uint256))
func (_MorphToken *MorphTokenSession) EpochInflationRates(index *big.Int) (IMorphTokenEpochInflationRate, error) {
	return _MorphToken.Contract.EpochInflationRates(&_MorphToken.CallOpts, index)
}

// EpochInflationRates is a free data retrieval call binding the contract method 0x6d0c4a26.
//
// Solidity: function epochInflationRates(uint256 index) view returns((uint256,uint256))
func (_MorphToken *MorphTokenCallerSession) EpochInflationRates(index *big.Int) (IMorphTokenEpochInflationRate, error) {
	return _MorphToken.Contract.EpochInflationRates(&_MorphToken.CallOpts, index)
}

// Inflation is a free data retrieval call binding the contract method 0x944fa746.
//
// Solidity: function inflation(uint256 epochIndex) view returns(uint256)
func (_MorphToken *MorphTokenCaller) Inflation(opts *bind.CallOpts, epochIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "inflation", epochIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Inflation is a free data retrieval call binding the contract method 0x944fa746.
//
// Solidity: function inflation(uint256 epochIndex) view returns(uint256)
func (_MorphToken *MorphTokenSession) Inflation(epochIndex *big.Int) (*big.Int, error) {
	return _MorphToken.Contract.Inflation(&_MorphToken.CallOpts, epochIndex)
}

// Inflation is a free data retrieval call binding the contract method 0x944fa746.
//
// Solidity: function inflation(uint256 epochIndex) view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) Inflation(epochIndex *big.Int) (*big.Int, error) {
	return _MorphToken.Contract.Inflation(&_MorphToken.CallOpts, epochIndex)
}

// InflationMintedEpochs is a free data retrieval call binding the contract method 0x0b88a984.
//
// Solidity: function inflationMintedEpochs() view returns(uint256)
func (_MorphToken *MorphTokenCaller) InflationMintedEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "inflationMintedEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationMintedEpochs is a free data retrieval call binding the contract method 0x0b88a984.
//
// Solidity: function inflationMintedEpochs() view returns(uint256)
func (_MorphToken *MorphTokenSession) InflationMintedEpochs() (*big.Int, error) {
	return _MorphToken.Contract.InflationMintedEpochs(&_MorphToken.CallOpts)
}

// InflationMintedEpochs is a free data retrieval call binding the contract method 0x0b88a984.
//
// Solidity: function inflationMintedEpochs() view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) InflationMintedEpochs() (*big.Int, error) {
	return _MorphToken.Contract.InflationMintedEpochs(&_MorphToken.CallOpts)
}

// InflationRatesCount is a free data retrieval call binding the contract method 0xc553f7b3.
//
// Solidity: function inflationRatesCount() view returns(uint256)
func (_MorphToken *MorphTokenCaller) InflationRatesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "inflationRatesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationRatesCount is a free data retrieval call binding the contract method 0xc553f7b3.
//
// Solidity: function inflationRatesCount() view returns(uint256)
func (_MorphToken *MorphTokenSession) InflationRatesCount() (*big.Int, error) {
	return _MorphToken.Contract.InflationRatesCount(&_MorphToken.CallOpts)
}

// InflationRatesCount is a free data retrieval call binding the contract method 0xc553f7b3.
//
// Solidity: function inflationRatesCount() view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) InflationRatesCount() (*big.Int, error) {
	return _MorphToken.Contract.InflationRatesCount(&_MorphToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphToken *MorphTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphToken *MorphTokenSession) Name() (string, error) {
	return _MorphToken.Contract.Name(&_MorphToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphToken *MorphTokenCallerSession) Name() (string, error) {
	return _MorphToken.Contract.Name(&_MorphToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphToken *MorphTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphToken *MorphTokenSession) Owner() (common.Address, error) {
	return _MorphToken.Contract.Owner(&_MorphToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MorphToken *MorphTokenCallerSession) Owner() (common.Address, error) {
	return _MorphToken.Contract.Owner(&_MorphToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphToken *MorphTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphToken *MorphTokenSession) Symbol() (string, error) {
	return _MorphToken.Contract.Symbol(&_MorphToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphToken *MorphTokenCallerSession) Symbol() (string, error) {
	return _MorphToken.Contract.Symbol(&_MorphToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphToken *MorphTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphToken *MorphTokenSession) TotalSupply() (*big.Int, error) {
	return _MorphToken.Contract.TotalSupply(&_MorphToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphToken *MorphTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MorphToken.Contract.TotalSupply(&_MorphToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Approve(&_MorphToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Approve(&_MorphToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MorphToken *MorphTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MorphToken *MorphTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Burn(&_MorphToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MorphToken *MorphTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Burn(&_MorphToken.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphToken *MorphTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphToken *MorphTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.DecreaseAllowance(&_MorphToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphToken *MorphTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.DecreaseAllowance(&_MorphToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphToken *MorphTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphToken *MorphTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.IncreaseAllowance(&_MorphToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphToken *MorphTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.IncreaseAllowance(&_MorphToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name_, string symbol_, address _owner, uint256 initialSupply_, uint256 dailyInflationRate_) returns()
func (_MorphToken *MorphTokenTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, _owner common.Address, initialSupply_ *big.Int, dailyInflationRate_ *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "initialize", name_, symbol_, _owner, initialSupply_, dailyInflationRate_)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name_, string symbol_, address _owner, uint256 initialSupply_, uint256 dailyInflationRate_) returns()
func (_MorphToken *MorphTokenSession) Initialize(name_ string, symbol_ string, _owner common.Address, initialSupply_ *big.Int, dailyInflationRate_ *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Initialize(&_MorphToken.TransactOpts, name_, symbol_, _owner, initialSupply_, dailyInflationRate_)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name_, string symbol_, address _owner, uint256 initialSupply_, uint256 dailyInflationRate_) returns()
func (_MorphToken *MorphTokenTransactorSession) Initialize(name_ string, symbol_ string, _owner common.Address, initialSupply_ *big.Int, dailyInflationRate_ *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Initialize(&_MorphToken.TransactOpts, name_, symbol_, _owner, initialSupply_, dailyInflationRate_)
}

// MintInflations is a paid mutator transaction binding the contract method 0xa29bfb2c.
//
// Solidity: function mintInflations(uint256 upToEpochIndex) returns()
func (_MorphToken *MorphTokenTransactor) MintInflations(opts *bind.TransactOpts, upToEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "mintInflations", upToEpochIndex)
}

// MintInflations is a paid mutator transaction binding the contract method 0xa29bfb2c.
//
// Solidity: function mintInflations(uint256 upToEpochIndex) returns()
func (_MorphToken *MorphTokenSession) MintInflations(upToEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.MintInflations(&_MorphToken.TransactOpts, upToEpochIndex)
}

// MintInflations is a paid mutator transaction binding the contract method 0xa29bfb2c.
//
// Solidity: function mintInflations(uint256 upToEpochIndex) returns()
func (_MorphToken *MorphTokenTransactorSession) MintInflations(upToEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.MintInflations(&_MorphToken.TransactOpts, upToEpochIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphToken *MorphTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphToken *MorphTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphToken.Contract.RenounceOwnership(&_MorphToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MorphToken *MorphTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MorphToken.Contract.RenounceOwnership(&_MorphToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Transfer(&_MorphToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.Transfer(&_MorphToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.TransferFrom(&_MorphToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphToken *MorphTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.TransferFrom(&_MorphToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphToken *MorphTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphToken *MorphTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphToken.Contract.TransferOwnership(&_MorphToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MorphToken *MorphTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MorphToken.Contract.TransferOwnership(&_MorphToken.TransactOpts, newOwner)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x405abb41.
//
// Solidity: function updateRate(uint256 newRate, uint256 effectiveEpochIndex) returns()
func (_MorphToken *MorphTokenTransactor) UpdateRate(opts *bind.TransactOpts, newRate *big.Int, effectiveEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.contract.Transact(opts, "updateRate", newRate, effectiveEpochIndex)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x405abb41.
//
// Solidity: function updateRate(uint256 newRate, uint256 effectiveEpochIndex) returns()
func (_MorphToken *MorphTokenSession) UpdateRate(newRate *big.Int, effectiveEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.UpdateRate(&_MorphToken.TransactOpts, newRate, effectiveEpochIndex)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x405abb41.
//
// Solidity: function updateRate(uint256 newRate, uint256 effectiveEpochIndex) returns()
func (_MorphToken *MorphTokenTransactorSession) UpdateRate(newRate *big.Int, effectiveEpochIndex *big.Int) (*types.Transaction, error) {
	return _MorphToken.Contract.UpdateRate(&_MorphToken.TransactOpts, newRate, effectiveEpochIndex)
}

// MorphTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MorphToken contract.
type MorphTokenApprovalIterator struct {
	Event *MorphTokenApproval // Event containing the contract specifics and raw log

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
func (it *MorphTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenApproval)
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
		it.Event = new(MorphTokenApproval)
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
func (it *MorphTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenApproval represents a Approval event raised by the MorphToken contract.
type MorphTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphToken *MorphTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MorphTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MorphTokenApprovalIterator{contract: _MorphToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphToken *MorphTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MorphTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenApproval)
				if err := _MorphToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphToken *MorphTokenFilterer) ParseApproval(log types.Log) (*MorphTokenApproval, error) {
	event := new(MorphTokenApproval)
	if err := _MorphToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphTokenInflationMintedIterator is returned from FilterInflationMinted and is used to iterate over the raw logs and unpacked data for InflationMinted events raised by the MorphToken contract.
type MorphTokenInflationMintedIterator struct {
	Event *MorphTokenInflationMinted // Event containing the contract specifics and raw log

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
func (it *MorphTokenInflationMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenInflationMinted)
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
		it.Event = new(MorphTokenInflationMinted)
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
func (it *MorphTokenInflationMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenInflationMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenInflationMinted represents a InflationMinted event raised by the MorphToken contract.
type MorphTokenInflationMinted struct {
	EpochIndex *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInflationMinted is a free log retrieval operation binding the contract event 0x0d82c0920038b8dc7f633e18585f37092ba957b84876fcf833d6841f69eaa327.
//
// Solidity: event InflationMinted(uint256 indexed epochIndex, uint256 amount)
func (_MorphToken *MorphTokenFilterer) FilterInflationMinted(opts *bind.FilterOpts, epochIndex []*big.Int) (*MorphTokenInflationMintedIterator, error) {

	var epochIndexRule []interface{}
	for _, epochIndexItem := range epochIndex {
		epochIndexRule = append(epochIndexRule, epochIndexItem)
	}

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "InflationMinted", epochIndexRule)
	if err != nil {
		return nil, err
	}
	return &MorphTokenInflationMintedIterator{contract: _MorphToken.contract, event: "InflationMinted", logs: logs, sub: sub}, nil
}

// WatchInflationMinted is a free log subscription operation binding the contract event 0x0d82c0920038b8dc7f633e18585f37092ba957b84876fcf833d6841f69eaa327.
//
// Solidity: event InflationMinted(uint256 indexed epochIndex, uint256 amount)
func (_MorphToken *MorphTokenFilterer) WatchInflationMinted(opts *bind.WatchOpts, sink chan<- *MorphTokenInflationMinted, epochIndex []*big.Int) (event.Subscription, error) {

	var epochIndexRule []interface{}
	for _, epochIndexItem := range epochIndex {
		epochIndexRule = append(epochIndexRule, epochIndexItem)
	}

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "InflationMinted", epochIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenInflationMinted)
				if err := _MorphToken.contract.UnpackLog(event, "InflationMinted", log); err != nil {
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

// ParseInflationMinted is a log parse operation binding the contract event 0x0d82c0920038b8dc7f633e18585f37092ba957b84876fcf833d6841f69eaa327.
//
// Solidity: event InflationMinted(uint256 indexed epochIndex, uint256 amount)
func (_MorphToken *MorphTokenFilterer) ParseInflationMinted(log types.Log) (*MorphTokenInflationMinted, error) {
	event := new(MorphTokenInflationMinted)
	if err := _MorphToken.contract.UnpackLog(event, "InflationMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MorphToken contract.
type MorphTokenInitializedIterator struct {
	Event *MorphTokenInitialized // Event containing the contract specifics and raw log

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
func (it *MorphTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenInitialized)
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
		it.Event = new(MorphTokenInitialized)
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
func (it *MorphTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenInitialized represents a Initialized event raised by the MorphToken contract.
type MorphTokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MorphToken *MorphTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*MorphTokenInitializedIterator, error) {

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MorphTokenInitializedIterator{contract: _MorphToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MorphToken *MorphTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MorphTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenInitialized)
				if err := _MorphToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MorphToken *MorphTokenFilterer) ParseInitialized(log types.Log) (*MorphTokenInitialized, error) {
	event := new(MorphTokenInitialized)
	if err := _MorphToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MorphToken contract.
type MorphTokenOwnershipTransferredIterator struct {
	Event *MorphTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MorphTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenOwnershipTransferred)
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
		it.Event = new(MorphTokenOwnershipTransferred)
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
func (it *MorphTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MorphToken contract.
type MorphTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphToken *MorphTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MorphTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MorphTokenOwnershipTransferredIterator{contract: _MorphToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MorphToken *MorphTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MorphTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenOwnershipTransferred)
				if err := _MorphToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MorphToken *MorphTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MorphTokenOwnershipTransferred, error) {
	event := new(MorphTokenOwnershipTransferred)
	if err := _MorphToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MorphToken contract.
type MorphTokenTransferIterator struct {
	Event *MorphTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MorphTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenTransfer)
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
		it.Event = new(MorphTokenTransfer)
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
func (it *MorphTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenTransfer represents a Transfer event raised by the MorphToken contract.
type MorphTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphToken *MorphTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MorphTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MorphTokenTransferIterator{contract: _MorphToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphToken *MorphTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MorphTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenTransfer)
				if err := _MorphToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphToken *MorphTokenFilterer) ParseTransfer(log types.Log) (*MorphTokenTransfer, error) {
	event := new(MorphTokenTransfer)
	if err := _MorphToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphTokenUpdateEpochInflationRateIterator is returned from FilterUpdateEpochInflationRate and is used to iterate over the raw logs and unpacked data for UpdateEpochInflationRate events raised by the MorphToken contract.
type MorphTokenUpdateEpochInflationRateIterator struct {
	Event *MorphTokenUpdateEpochInflationRate // Event containing the contract specifics and raw log

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
func (it *MorphTokenUpdateEpochInflationRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphTokenUpdateEpochInflationRate)
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
		it.Event = new(MorphTokenUpdateEpochInflationRate)
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
func (it *MorphTokenUpdateEpochInflationRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphTokenUpdateEpochInflationRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphTokenUpdateEpochInflationRate represents a UpdateEpochInflationRate event raised by the MorphToken contract.
type MorphTokenUpdateEpochInflationRate struct {
	Rate                *big.Int
	EffectiveEpochIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateEpochInflationRate is a free log retrieval operation binding the contract event 0xbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f4.
//
// Solidity: event UpdateEpochInflationRate(uint256 indexed rate, uint256 indexed effectiveEpochIndex)
func (_MorphToken *MorphTokenFilterer) FilterUpdateEpochInflationRate(opts *bind.FilterOpts, rate []*big.Int, effectiveEpochIndex []*big.Int) (*MorphTokenUpdateEpochInflationRateIterator, error) {

	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}
	var effectiveEpochIndexRule []interface{}
	for _, effectiveEpochIndexItem := range effectiveEpochIndex {
		effectiveEpochIndexRule = append(effectiveEpochIndexRule, effectiveEpochIndexItem)
	}

	logs, sub, err := _MorphToken.contract.FilterLogs(opts, "UpdateEpochInflationRate", rateRule, effectiveEpochIndexRule)
	if err != nil {
		return nil, err
	}
	return &MorphTokenUpdateEpochInflationRateIterator{contract: _MorphToken.contract, event: "UpdateEpochInflationRate", logs: logs, sub: sub}, nil
}

// WatchUpdateEpochInflationRate is a free log subscription operation binding the contract event 0xbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f4.
//
// Solidity: event UpdateEpochInflationRate(uint256 indexed rate, uint256 indexed effectiveEpochIndex)
func (_MorphToken *MorphTokenFilterer) WatchUpdateEpochInflationRate(opts *bind.WatchOpts, sink chan<- *MorphTokenUpdateEpochInflationRate, rate []*big.Int, effectiveEpochIndex []*big.Int) (event.Subscription, error) {

	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}
	var effectiveEpochIndexRule []interface{}
	for _, effectiveEpochIndexItem := range effectiveEpochIndex {
		effectiveEpochIndexRule = append(effectiveEpochIndexRule, effectiveEpochIndexItem)
	}

	logs, sub, err := _MorphToken.contract.WatchLogs(opts, "UpdateEpochInflationRate", rateRule, effectiveEpochIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphTokenUpdateEpochInflationRate)
				if err := _MorphToken.contract.UnpackLog(event, "UpdateEpochInflationRate", log); err != nil {
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

// ParseUpdateEpochInflationRate is a log parse operation binding the contract event 0xbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f4.
//
// Solidity: event UpdateEpochInflationRate(uint256 indexed rate, uint256 indexed effectiveEpochIndex)
func (_MorphToken *MorphTokenFilterer) ParseUpdateEpochInflationRate(log types.Log) (*MorphTokenUpdateEpochInflationRate, error) {
	event := new(MorphTokenUpdateEpochInflationRate)
	if err := _MorphToken.contract.UnpackLog(event, "UpdateEpochInflationRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
