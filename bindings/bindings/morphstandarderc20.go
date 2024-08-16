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

// MorphStandardERC20MetaData contains all meta data concerning the MorphStandardERC20 contract.
var MorphStandardERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61203080620000e75f395ff3fe608060405234801561000f575f80fd5b506004361061016e575f3560e01c806370a08231116100d25780639dc29fac11610088578063c820f14611610063578063c820f14614610354578063d505accf14610367578063dd62ed3e1461037a575f80fd5b80639dc29fac1461031b578063a457c2d71461032e578063a9059cbb14610341575f80fd5b80637ecebe00116100b85780637ecebe00146102e557806384b0196e146102f857806395d89b4114610313575f80fd5b806370a0823114610290578063797594b0146102c5575f80fd5b8063313ce56711610127578063395093511161010d57806339509351146102555780634000aea01461026857806340c10f191461027b575f80fd5b8063313ce5671461021d5780633644e5151461024d575f80fd5b8063116191b611610157578063116191b6146101b357806318160ddd146101f857806323b872dd1461020a575f80fd5b806306fdde0314610172578063095ea7b314610190575b5f80fd5b61017a6103bf565b60405161018791906119ce565b60405180910390f35b6101a361019e366004611a0f565b61044f565b6040519015158152602001610187565b60cc546101d39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610187565b6035545b604051908152602001610187565b6101a3610218366004611a37565b610468565b60cd5474010000000000000000000000000000000000000000900460ff1660405160ff9091168152602001610187565b6101fc61048b565b6101a3610263366004611a0f565b610499565b6101a3610276366004611a70565b6104e4565b61028e610289366004611a0f565b61054d565b005b6101fc61029e366004611af0565b73ffffffffffffffffffffffffffffffffffffffff165f9081526033602052604090205490565b60cd546101d39073ffffffffffffffffffffffffffffffffffffffff1681565b6101fc6102f3366004611af0565b6105c7565b6103006105f1565b6040516101879796959493929190611b09565b61017a6106ae565b61028e610329366004611a0f565b6106bd565b6101a361033c366004611a0f565b61072e565b6101a361034f366004611a0f565b6107e4565b61028e610362366004611cad565b6107f1565b61028e610375366004611d3d565b610a08565b6101fc610388366004611da2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260346020908152604080832093909416825291909152205490565b6060603680546103ce90611dd3565b80601f01602080910402602001604051908101604052809291908181526020018280546103fa90611dd3565b80156104455780601f1061041c57610100808354040283529160200191610445565b820191905f5260205f20905b81548152906001019060200180831161042857829003601f168201915b5050505050905090565b5f3361045c818585610b90565b60019150505b92915050565b5f33610475858285610d0f565b610480858585610dcb565b506001949350505050565b5f610494610ff1565b905090565b335f81815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061045c90829086906104df908790611e1e565b610b90565b5f6104ef85856107e4565b5073ffffffffffffffffffffffffffffffffffffffff85163b1561048057610480858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610ffa92505050565b60cc5473ffffffffffffffffffffffffffffffffffffffff1633146105b95760405162461bcd60e51b815260206004820152600c60248201527f4f6e6c792047617465776179000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6105c38282611085565b5050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260996020526040812054610462565b5f6060805f805f60606065545f801b14801561060d5750606654155b6106595760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016105b0565b61066161115e565b61066961116d565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b6060603780546103ce90611dd3565b60cc5473ffffffffffffffffffffffffffffffffffffffff1633146107245760405162461bcd60e51b815260206004820152600c60248201527f4f6e6c792047617465776179000000000000000000000000000000000000000060448201526064016105b0565b6105c3828261117c565b335f81815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107d75760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016105b0565b6104808286868403610b90565b5f3361045c818585610dcb565b5f54610100900460ff161580801561080f57505f54600160ff909116105b806108285750303b15801561082857505f5460ff166001145b61089a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105b0565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108f6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6108ff86611309565b61090986866113c7565b60cd805460cc805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925590851660ff88167401000000000000000000000000000000000000000002919091167fffffffffffffffffffffff000000000000000000000000000000000000000000909216919091171790558015610a00575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b83421115610a585760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016105b0565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610a868c61144d565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610aed82611481565b90505f610afc828787876114c8565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b795760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016105b0565b610b848a8a8a610b90565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c185760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff8216610ca15760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152603460209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610dc55781811015610db85760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105b0565b610dc58484848403610b90565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610e545760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff8216610edd5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526033602052604090205481811015610f785760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff8085165f8181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610fe49086815260200190565b60405180910390a3610dc5565b5f6104946114ee565b6040517fa4c0ed36000000000000000000000000000000000000000000000000000000008152839073ffffffffffffffffffffffffffffffffffffffff82169063a4c0ed369061105290339087908790600401611e56565b5f604051808303815f87803b158015611069575f80fd5b505af115801561107b573d5f803e3d5ffd5b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff82166110e85760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105b0565b8060355f8282546110f99190611e1e565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6060606780546103ce90611dd3565b6060606880546103ce90611dd3565b73ffffffffffffffffffffffffffffffffffffffff82166112055760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260336020526040902054818110156112a05760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016105b0565b73ffffffffffffffffffffffffffffffffffffffff83165f8181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d02565b505050565b5f54610100900460ff166113855760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b0565b6113c4816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250611561565b50565b5f54610100900460ff166114435760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b0565b6105c38282611604565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526099602052604090208054600181018255905b50919050565b5f61046261148d610ff1565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f805f6114d787878787611699565b915091506114e481611781565b5095945050505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6115186118e5565b61152061193d565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f54610100900460ff166115dd5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b0565b60676115e98382611ede565b5060686115f68282611ede565b50505f606581905560665550565b5f54610100900460ff166116805760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105b0565b603661168c8382611ede565b5060376113048282611ede565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156116ce57505f90506003611778565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561171f573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611772575f60019250925050611778565b91505f90505b94509492505050565b5f81600481111561179457611794611ff6565b0361179c5750565b60018160048111156117b0576117b0611ff6565b036117fd5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105b0565b600281600481111561181157611811611ff6565b0361185e5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105b0565b600381600481111561187257611872611ff6565b036113c45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016105b0565b5f806118ef61115e565b805190915015611906578051602090910120919050565b60655480156119155792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b5f8061194761116d565b80519091501561195e578051602090910120919050565b60665480156119155792915050565b5f81518084525f5b8181101561199157602081850181015186830182015201611975565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6119e0602083018461196d565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a0a575f80fd5b919050565b5f8060408385031215611a20575f80fd5b611a29836119e7565b946020939093013593505050565b5f805f60608486031215611a49575f80fd5b611a52846119e7565b9250611a60602085016119e7565b9150604084013590509250925092565b5f805f8060608587031215611a83575f80fd5b611a8c856119e7565b935060208501359250604085013567ffffffffffffffff80821115611aaf575f80fd5b818701915087601f830112611ac2575f80fd5b813581811115611ad0575f80fd5b886020828501011115611ae1575f80fd5b95989497505060200194505050565b5f60208284031215611b00575f80fd5b6119e0826119e7565b7fff00000000000000000000000000000000000000000000000000000000000000881681525f602060e06020840152611b4560e084018a61196d565b8381036040850152611b57818a61196d565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b81811015611bb757835183529284019291840191600101611b9b565b50909c9b505050505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611c05575f80fd5b813567ffffffffffffffff80821115611c2057611c20611bc9565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611c6657611c66611bc9565b81604052838152866020858801011115611c7e575f80fd5b836020870160208301375f602085830101528094505050505092915050565b803560ff81168114611a0a575f80fd5b5f805f805f60a08688031215611cc1575f80fd5b853567ffffffffffffffff80821115611cd8575f80fd5b611ce489838a01611bf6565b96506020880135915080821115611cf9575f80fd5b50611d0688828901611bf6565b945050611d1560408701611c9d565b9250611d23606087016119e7565b9150611d31608087016119e7565b90509295509295909350565b5f805f805f805f60e0888a031215611d53575f80fd5b611d5c886119e7565b9650611d6a602089016119e7565b95506040880135945060608801359350611d8660808901611c9d565b925060a0880135915060c0880135905092959891949750929550565b5f8060408385031215611db3575f80fd5b611dbc836119e7565b9150611dca602084016119e7565b90509250929050565b600181811c90821680611de757607f821691505b60208210810361147b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b80820180821115610462577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f611e8a606083018461196d565b95945050505050565b601f82111561130457805f5260205f20601f840160051c81016020851015611eb85750805b601f840160051c820191505b81811015611ed7575f8155600101611ec4565b5050505050565b815167ffffffffffffffff811115611ef857611ef8611bc9565b611f0c81611f068454611dd3565b84611e93565b602080601f831160018114611f5e575f8415611f285750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610a00565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611faa57888601518255948401946001909101908401611f8b565b5085821015611fe657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea164736f6c6343000818000a",
}

// MorphStandardERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphStandardERC20MetaData.ABI instead.
var MorphStandardERC20ABI = MorphStandardERC20MetaData.ABI

// MorphStandardERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MorphStandardERC20MetaData.Bin instead.
var MorphStandardERC20Bin = MorphStandardERC20MetaData.Bin

// DeployMorphStandardERC20 deploys a new Ethereum contract, binding an instance of MorphStandardERC20 to it.
func DeployMorphStandardERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MorphStandardERC20, error) {
	parsed, err := MorphStandardERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MorphStandardERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MorphStandardERC20{MorphStandardERC20Caller: MorphStandardERC20Caller{contract: contract}, MorphStandardERC20Transactor: MorphStandardERC20Transactor{contract: contract}, MorphStandardERC20Filterer: MorphStandardERC20Filterer{contract: contract}}, nil
}

// MorphStandardERC20 is an auto generated Go binding around an Ethereum contract.
type MorphStandardERC20 struct {
	MorphStandardERC20Caller     // Read-only binding to the contract
	MorphStandardERC20Transactor // Write-only binding to the contract
	MorphStandardERC20Filterer   // Log filterer for contract events
}

// MorphStandardERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type MorphStandardERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphStandardERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphStandardERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphStandardERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphStandardERC20Session struct {
	Contract     *MorphStandardERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MorphStandardERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphStandardERC20CallerSession struct {
	Contract *MorphStandardERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MorphStandardERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphStandardERC20TransactorSession struct {
	Contract     *MorphStandardERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MorphStandardERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type MorphStandardERC20Raw struct {
	Contract *MorphStandardERC20 // Generic contract binding to access the raw methods on
}

// MorphStandardERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphStandardERC20CallerRaw struct {
	Contract *MorphStandardERC20Caller // Generic read-only contract binding to access the raw methods on
}

// MorphStandardERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphStandardERC20TransactorRaw struct {
	Contract *MorphStandardERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphStandardERC20 creates a new instance of MorphStandardERC20, bound to a specific deployed contract.
func NewMorphStandardERC20(address common.Address, backend bind.ContractBackend) (*MorphStandardERC20, error) {
	contract, err := bindMorphStandardERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20{MorphStandardERC20Caller: MorphStandardERC20Caller{contract: contract}, MorphStandardERC20Transactor: MorphStandardERC20Transactor{contract: contract}, MorphStandardERC20Filterer: MorphStandardERC20Filterer{contract: contract}}, nil
}

// NewMorphStandardERC20Caller creates a new read-only instance of MorphStandardERC20, bound to a specific deployed contract.
func NewMorphStandardERC20Caller(address common.Address, caller bind.ContractCaller) (*MorphStandardERC20Caller, error) {
	contract, err := bindMorphStandardERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20Caller{contract: contract}, nil
}

// NewMorphStandardERC20Transactor creates a new write-only instance of MorphStandardERC20, bound to a specific deployed contract.
func NewMorphStandardERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MorphStandardERC20Transactor, error) {
	contract, err := bindMorphStandardERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20Transactor{contract: contract}, nil
}

// NewMorphStandardERC20Filterer creates a new log filterer instance of MorphStandardERC20, bound to a specific deployed contract.
func NewMorphStandardERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MorphStandardERC20Filterer, error) {
	contract, err := bindMorphStandardERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20Filterer{contract: contract}, nil
}

// bindMorphStandardERC20 binds a generic wrapper to an already deployed contract.
func bindMorphStandardERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphStandardERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphStandardERC20 *MorphStandardERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphStandardERC20.Contract.MorphStandardERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphStandardERC20 *MorphStandardERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.MorphStandardERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphStandardERC20 *MorphStandardERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.MorphStandardERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphStandardERC20 *MorphStandardERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphStandardERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphStandardERC20 *MorphStandardERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphStandardERC20 *MorphStandardERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphStandardERC20 *MorphStandardERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphStandardERC20 *MorphStandardERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphStandardERC20.Contract.DOMAINSEPARATOR(&_MorphStandardERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MorphStandardERC20.Contract.DOMAINSEPARATOR(&_MorphStandardERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.Allowance(&_MorphStandardERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.Allowance(&_MorphStandardERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.BalanceOf(&_MorphStandardERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.BalanceOf(&_MorphStandardERC20.CallOpts, account)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20Session) Counterpart() (common.Address, error) {
	return _MorphStandardERC20.Contract.Counterpart(&_MorphStandardERC20.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Counterpart() (common.Address, error) {
	return _MorphStandardERC20.Contract.Counterpart(&_MorphStandardERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphStandardERC20 *MorphStandardERC20Session) Decimals() (uint8, error) {
	return _MorphStandardERC20.Contract.Decimals(&_MorphStandardERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Decimals() (uint8, error) {
	return _MorphStandardERC20.Contract.Decimals(&_MorphStandardERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "eip712Domain")

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
func (_MorphStandardERC20 *MorphStandardERC20Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MorphStandardERC20.Contract.Eip712Domain(&_MorphStandardERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MorphStandardERC20.Contract.Eip712Domain(&_MorphStandardERC20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20Session) Gateway() (common.Address, error) {
	return _MorphStandardERC20.Contract.Gateway(&_MorphStandardERC20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Gateway() (common.Address, error) {
	return _MorphStandardERC20.Contract.Gateway(&_MorphStandardERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20Session) Name() (string, error) {
	return _MorphStandardERC20.Contract.Name(&_MorphStandardERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Name() (string, error) {
	return _MorphStandardERC20.Contract.Name(&_MorphStandardERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.Nonces(&_MorphStandardERC20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MorphStandardERC20.Contract.Nonces(&_MorphStandardERC20.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20Session) Symbol() (string, error) {
	return _MorphStandardERC20.Contract.Symbol(&_MorphStandardERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) Symbol() (string, error) {
	return _MorphStandardERC20.Contract.Symbol(&_MorphStandardERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MorphStandardERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20Session) TotalSupply() (*big.Int, error) {
	return _MorphStandardERC20.Contract.TotalSupply(&_MorphStandardERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MorphStandardERC20 *MorphStandardERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MorphStandardERC20.Contract.TotalSupply(&_MorphStandardERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Approve(&_MorphStandardERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Approve(&_MorphStandardERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Burn(&_MorphStandardERC20.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Burn(&_MorphStandardERC20.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.DecreaseAllowance(&_MorphStandardERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.DecreaseAllowance(&_MorphStandardERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.IncreaseAllowance(&_MorphStandardERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.IncreaseAllowance(&_MorphStandardERC20.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc820f146.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address _gateway, address _counterpart) returns()
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _gateway common.Address, _counterpart common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "initialize", _name, _symbol, _decimals, _gateway, _counterpart)
}

// Initialize is a paid mutator transaction binding the contract method 0xc820f146.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address _gateway, address _counterpart) returns()
func (_MorphStandardERC20 *MorphStandardERC20Session) Initialize(_name string, _symbol string, _decimals uint8, _gateway common.Address, _counterpart common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Initialize(&_MorphStandardERC20.TransactOpts, _name, _symbol, _decimals, _gateway, _counterpart)
}

// Initialize is a paid mutator transaction binding the contract method 0xc820f146.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address _gateway, address _counterpart) returns()
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Initialize(_name string, _symbol string, _decimals uint8, _gateway common.Address, _counterpart common.Address) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Initialize(&_MorphStandardERC20.TransactOpts, _name, _symbol, _decimals, _gateway, _counterpart)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Mint(&_MorphStandardERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Mint(&_MorphStandardERC20.TransactOpts, _to, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphStandardERC20 *MorphStandardERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Permit(&_MorphStandardERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Permit(&_MorphStandardERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Transfer(&_MorphStandardERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.Transfer(&_MorphStandardERC20.TransactOpts, to, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) TransferAndCall(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "transferAndCall", receiver, amount, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_MorphStandardERC20 *MorphStandardERC20Session) TransferAndCall(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.TransferAndCall(&_MorphStandardERC20.TransactOpts, receiver, amount, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) TransferAndCall(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.TransferAndCall(&_MorphStandardERC20.TransactOpts, receiver, amount, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.TransferFrom(&_MorphStandardERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MorphStandardERC20 *MorphStandardERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MorphStandardERC20.Contract.TransferFrom(&_MorphStandardERC20.TransactOpts, from, to, amount)
}

// MorphStandardERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MorphStandardERC20 contract.
type MorphStandardERC20ApprovalIterator struct {
	Event *MorphStandardERC20Approval // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20Approval)
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
		it.Event = new(MorphStandardERC20Approval)
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
func (it *MorphStandardERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20Approval represents a Approval event raised by the MorphStandardERC20 contract.
type MorphStandardERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MorphStandardERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphStandardERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20ApprovalIterator{contract: _MorphStandardERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MorphStandardERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20Approval)
				if err := _MorphStandardERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MorphStandardERC20 *MorphStandardERC20Filterer) ParseApproval(log types.Log) (*MorphStandardERC20Approval, error) {
	event := new(MorphStandardERC20Approval)
	if err := _MorphStandardERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphStandardERC20EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MorphStandardERC20 contract.
type MorphStandardERC20EIP712DomainChangedIterator struct {
	Event *MorphStandardERC20EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20EIP712DomainChanged)
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
		it.Event = new(MorphStandardERC20EIP712DomainChanged)
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
func (it *MorphStandardERC20EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20EIP712DomainChanged represents a EIP712DomainChanged event raised by the MorphStandardERC20 contract.
type MorphStandardERC20EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MorphStandardERC20 *MorphStandardERC20Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MorphStandardERC20EIP712DomainChangedIterator, error) {

	logs, sub, err := _MorphStandardERC20.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20EIP712DomainChangedIterator{contract: _MorphStandardERC20.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MorphStandardERC20 *MorphStandardERC20Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MorphStandardERC20.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20EIP712DomainChanged)
				if err := _MorphStandardERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_MorphStandardERC20 *MorphStandardERC20Filterer) ParseEIP712DomainChanged(log types.Log) (*MorphStandardERC20EIP712DomainChanged, error) {
	event := new(MorphStandardERC20EIP712DomainChanged)
	if err := _MorphStandardERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphStandardERC20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MorphStandardERC20 contract.
type MorphStandardERC20InitializedIterator struct {
	Event *MorphStandardERC20Initialized // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20Initialized)
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
		it.Event = new(MorphStandardERC20Initialized)
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
func (it *MorphStandardERC20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20Initialized represents a Initialized event raised by the MorphStandardERC20 contract.
type MorphStandardERC20Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) FilterInitialized(opts *bind.FilterOpts) (*MorphStandardERC20InitializedIterator, error) {

	logs, sub, err := _MorphStandardERC20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20InitializedIterator{contract: _MorphStandardERC20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20Initialized) (event.Subscription, error) {

	logs, sub, err := _MorphStandardERC20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20Initialized)
				if err := _MorphStandardERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MorphStandardERC20 *MorphStandardERC20Filterer) ParseInitialized(log types.Log) (*MorphStandardERC20Initialized, error) {
	event := new(MorphStandardERC20Initialized)
	if err := _MorphStandardERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphStandardERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MorphStandardERC20 contract.
type MorphStandardERC20TransferIterator struct {
	Event *MorphStandardERC20Transfer // Event containing the contract specifics and raw log

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
func (it *MorphStandardERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphStandardERC20Transfer)
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
		it.Event = new(MorphStandardERC20Transfer)
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
func (it *MorphStandardERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphStandardERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphStandardERC20Transfer represents a Transfer event raised by the MorphStandardERC20 contract.
type MorphStandardERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MorphStandardERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphStandardERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MorphStandardERC20TransferIterator{contract: _MorphStandardERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MorphStandardERC20 *MorphStandardERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MorphStandardERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MorphStandardERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphStandardERC20Transfer)
				if err := _MorphStandardERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MorphStandardERC20 *MorphStandardERC20Filterer) ParseTransfer(log types.Log) (*MorphStandardERC20Transfer, error) {
	event := new(MorphStandardERC20Transfer)
	if err := _MorphStandardERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
