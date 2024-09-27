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

// WrappedEtherMetaData contains all meta data concerning the WrappedEther contract.
var WrappedEtherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61016060405234801562000011575f80fd5b506040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b81525080604051806040016040528060018152602001603160f81b8152506040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b815250604051806040016040528060048152602001630ae8aa8960e31b8152508160039081620000a7919062000297565b506004620000b6828262000297565b50620000c89150839050600562000176565b61012052620000d981600662000176565b61014052815160208084019190912060e052815190820120610100524660a0526200016660e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250620003d5565b5f60208351101562000195576200018d83620001ae565b9050620001a8565b81620001a2848262000297565b5060ff90505b92915050565b5f80829050601f81511115620001e4578260405163305a27a960e01b8152600401620001db919062000363565b60405180910390fd5b8051620001f182620003b1565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200022257607f821691505b6020821081036200024157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200029257805f5260205f20601f840160051c810160208510156200026e5750805b601f840160051c820191505b818110156200028f575f81556001016200027a565b50505b505050565b81516001600160401b03811115620002b357620002b3620001f9565b620002cb81620002c484546200020d565b8462000247565b602080601f83116001811462000301575f8415620002e95750858301515b5f19600386901b1c1916600185901b1785556200035b565b5f85815260208120601f198616915b82811015620003315788860151825594840194600190910190840162000310565b50858210156200034f57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f602080835283518060208501525f5b81811015620003915785810183015185820160400152820162000373565b505f604082860101526040601f19601f8301168501019250505092915050565b8051602080830151919081101562000241575f1960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051611a1f620004275f395f61068301525f61065801525f61121d01525f6111f501525f61115001525f61117a01525f6111a40152611a1f5ff3fe60806040526004361061010c575f3560e01c806370a08231116100a1578063a457c2d711610071578063d0e30db011610057578063d0e30db0146102fb578063d505accf14610303578063dd62ed3e14610322575f80fd5b8063a457c2d7146102bd578063a9059cbb146102dc575f80fd5b806370a08231146102225780637ecebe001461026357806384b0196e1461028257806395d89b41146102a9575f80fd5b80632e1a7d4d116100dc5780632e1a7d4d146101b5578063313ce567146101d45780633644e515146101ef5780633950935114610203575f80fd5b806306fdde031461011f578063095ea7b31461014957806318160ddd1461017857806323b872dd14610196575f80fd5b3661011b57610119610373565b005b5f80fd5b34801561012a575f80fd5b506101336103d1565b6040516101409190611732565b60405180910390f35b348015610154575f80fd5b50610168610163366004611773565b610461565b6040519015158152602001610140565b348015610183575f80fd5b506002545b604051908152602001610140565b3480156101a1575f80fd5b506101686101b036600461179b565b61047a565b3480156101c0575f80fd5b506101196101cf3660046117d4565b61049d565b3480156101df575f80fd5b5060405160128152602001610140565b3480156101fa575f80fd5b506101886105c8565b34801561020e575f80fd5b5061016861021d366004611773565b6105d6565b34801561022d575f80fd5b5061018861023c3660046117eb565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b34801561026e575f80fd5b5061018861027d3660046117eb565b610621565b34801561028d575f80fd5b5061029661064b565b6040516101409796959493929190611804565b3480156102b4575f80fd5b506101336106ee565b3480156102c8575f80fd5b506101686102d7366004611773565b6106fd565b3480156102e7575f80fd5b506101686102f6366004611773565b6107cd565b610119610373565b34801561030e575f80fd5b5061011961031d3660046118c4565b6107da565b34801561032d575f80fd5b5061018861033c366004611931565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b3361037e8134610996565b8073ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516103c691815260200190565b60405180910390a250565b6060600380546103e090611962565b80601f016020809104026020016040519081016040528092919081815260200182805461040c90611962565b80156104575780601f1061042e57610100808354040283529160200191610457565b820191905f5260205f20905b81548152906001019060200180831161043a57829003601f168201915b5050505050905090565b5f3361046e818585610a87565b60019150505b92915050565b5f33610487858285610c3a565b610492858585610d10565b506001949350505050565b336104a88183610f7d565b5f8173ffffffffffffffffffffffffffffffffffffffff16836040515f6040518083038185875af1925050503d805f81146104fe576040519150601f19603f3d011682016040523d82523d5f602084013e610503565b606091505b5050905080610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f776974686472617720455448206661696c65640000000000000000000000000060448201526064015b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65846040516105bb91815260200190565b60405180910390a2505050565b5f6105d1611137565b905090565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061046e908290869061061c9087906119ad565b610a87565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260076020526040812054610474565b5f6060808280808361067e7f0000000000000000000000000000000000000000000000000000000000000000600561126d565b6106a97f0000000000000000000000000000000000000000000000000000000000000000600661126d565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b6060600480546103e090611962565b335f81815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161056a565b6104928286868403610a87565b5f3361046e818585610d10565b83421115610844576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e65000000604482015260640161056a565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886108728c611316565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6108d98261134a565b90505f6108e882878787611391565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461097f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e61747572650000604482015260640161056a565b61098a8a8a8a610a87565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8216610a13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161056a565b8060025f828254610a2491906119ad565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8316610b29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff8216610bcc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610d0a5781811015610cfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161056a565b610d0a8484848403610a87565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610db3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff8216610e56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610f0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff8481165f81815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610d0a565b73ffffffffffffffffffffffffffffffffffffffff8216611020576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260208190526040902054818110156110d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161056a565b73ffffffffffffffffffffffffffffffffffffffff83165f818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610c2d565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561119c57507f000000000000000000000000000000000000000000000000000000000000000046145b156111c657507f000000000000000000000000000000000000000000000000000000000000000090565b6105d1604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b606060ff831461128757611280836113b7565b9050610474565b81805461129390611962565b80601f01602080910402602001604051908101604052809291908181526020018280546112bf90611962565b801561130a5780601f106112e15761010080835404028352916020019161130a565b820191905f5260205f20905b8154815290600101906020018083116112ed57829003601f168201915b50505050509050610474565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526007602052604090208054600181018255905b50919050565b5f610474611356611137565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f805f6113a0878787876113f4565b915091506113ad816114dc565b5095945050505050565b60605f6113c383611691565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561142957505f905060036114d3565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561147a573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166114cd575f600192509250506114d3565b91505f90505b94509492505050565b5f8160048111156114ef576114ef6119e5565b036114f75750565b600181600481111561150b5761150b6119e5565b03611572576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161056a565b6002816004811115611586576115866119e5565b036115ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161056a565b6003816004811115611601576116016119e5565b0361168e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161056a565b50565b5f60ff8216601f811115610474576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81518084525f5b818110156116f5576020818501810151868301820152016116d9565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61174460208301846116d1565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461176e575f80fd5b919050565b5f8060408385031215611784575f80fd5b61178d8361174b565b946020939093013593505050565b5f805f606084860312156117ad575f80fd5b6117b68461174b565b92506117c46020850161174b565b9150604084013590509250925092565b5f602082840312156117e4575f80fd5b5035919050565b5f602082840312156117fb575f80fd5b6117448261174b565b7fff00000000000000000000000000000000000000000000000000000000000000881681525f602060e0602084015261184060e084018a6116d1565b8381036040850152611852818a6116d1565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b818110156118b257835183529284019291840191600101611896565b50909c9b505050505050505050505050565b5f805f805f805f60e0888a0312156118da575f80fd5b6118e38861174b565b96506118f16020890161174b565b95506040880135945060608801359350608088013560ff81168114611914575f80fd5b9699959850939692959460a0840135945060c09093013592915050565b5f8060408385031215611942575f80fd5b61194b8361174b565b91506119596020840161174b565b90509250929050565b600181811c9082168061197657607f821691505b602082108103611344577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b80820180821115610474577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea164736f6c6343000818000a",
}

// WrappedEtherABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedEtherMetaData.ABI instead.
var WrappedEtherABI = WrappedEtherMetaData.ABI

// WrappedEtherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedEtherMetaData.Bin instead.
var WrappedEtherBin = WrappedEtherMetaData.Bin

// DeployWrappedEther deploys a new Ethereum contract, binding an instance of WrappedEther to it.
func DeployWrappedEther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WrappedEther, error) {
	parsed, err := WrappedEtherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedEtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedEther{WrappedEtherCaller: WrappedEtherCaller{contract: contract}, WrappedEtherTransactor: WrappedEtherTransactor{contract: contract}, WrappedEtherFilterer: WrappedEtherFilterer{contract: contract}}, nil
}

// WrappedEther is an auto generated Go binding around an Ethereum contract.
type WrappedEther struct {
	WrappedEtherCaller     // Read-only binding to the contract
	WrappedEtherTransactor // Write-only binding to the contract
	WrappedEtherFilterer   // Log filterer for contract events
}

// WrappedEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedEtherSession struct {
	Contract     *WrappedEther     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappedEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedEtherCallerSession struct {
	Contract *WrappedEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WrappedEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedEtherTransactorSession struct {
	Contract     *WrappedEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WrappedEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedEtherRaw struct {
	Contract *WrappedEther // Generic contract binding to access the raw methods on
}

// WrappedEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedEtherCallerRaw struct {
	Contract *WrappedEtherCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedEtherTransactorRaw struct {
	Contract *WrappedEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedEther creates a new instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEther(address common.Address, backend bind.ContractBackend) (*WrappedEther, error) {
	contract, err := bindWrappedEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedEther{WrappedEtherCaller: WrappedEtherCaller{contract: contract}, WrappedEtherTransactor: WrappedEtherTransactor{contract: contract}, WrappedEtherFilterer: WrappedEtherFilterer{contract: contract}}, nil
}

// NewWrappedEtherCaller creates a new read-only instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherCaller(address common.Address, caller bind.ContractCaller) (*WrappedEtherCaller, error) {
	contract, err := bindWrappedEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherCaller{contract: contract}, nil
}

// NewWrappedEtherTransactor creates a new write-only instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedEtherTransactor, error) {
	contract, err := bindWrappedEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherTransactor{contract: contract}, nil
}

// NewWrappedEtherFilterer creates a new log filterer instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedEtherFilterer, error) {
	contract, err := bindWrappedEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherFilterer{contract: contract}, nil
}

// bindWrappedEther binds a generic wrapper to an already deployed contract.
func bindWrappedEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedEtherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedEther *WrappedEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedEther.Contract.WrappedEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedEther *WrappedEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.Contract.WrappedEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedEther *WrappedEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedEther.Contract.WrappedEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedEther *WrappedEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedEther *WrappedEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedEther *WrappedEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedEther.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WrappedEther.Contract.DOMAINSEPARATOR(&_WrappedEther.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WrappedEther.Contract.DOMAINSEPARATOR(&_WrappedEther.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Allowance(&_WrappedEther.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Allowance(&_WrappedEther.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.BalanceOf(&_WrappedEther.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.BalanceOf(&_WrappedEther.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherSession) Decimals() (uint8, error) {
	return _WrappedEther.Contract.Decimals(&_WrappedEther.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherCallerSession) Decimals() (uint8, error) {
	return _WrappedEther.Contract.Decimals(&_WrappedEther.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WrappedEther *WrappedEtherCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "eip712Domain")

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
func (_WrappedEther *WrappedEtherSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WrappedEther.Contract.Eip712Domain(&_WrappedEther.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WrappedEther *WrappedEtherCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WrappedEther.Contract.Eip712Domain(&_WrappedEther.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherSession) Name() (string, error) {
	return _WrappedEther.Contract.Name(&_WrappedEther.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherCallerSession) Name() (string, error) {
	return _WrappedEther.Contract.Name(&_WrappedEther.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Nonces(&_WrappedEther.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Nonces(&_WrappedEther.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherSession) Symbol() (string, error) {
	return _WrappedEther.Contract.Symbol(&_WrappedEther.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherCallerSession) Symbol() (string, error) {
	return _WrappedEther.Contract.Symbol(&_WrappedEther.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherSession) TotalSupply() (*big.Int, error) {
	return _WrappedEther.Contract.TotalSupply(&_WrappedEther.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedEther.Contract.TotalSupply(&_WrappedEther.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Approve(&_WrappedEther.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Approve(&_WrappedEther.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.DecreaseAllowance(&_WrappedEther.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.DecreaseAllowance(&_WrappedEther.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherSession) Deposit() (*types.Transaction, error) {
	return _WrappedEther.Contract.Deposit(&_WrappedEther.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedEther.Contract.Deposit(&_WrappedEther.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.IncreaseAllowance(&_WrappedEther.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.IncreaseAllowance(&_WrappedEther.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.Contract.Permit(&_WrappedEther.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.Contract.Permit(&_WrappedEther.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Transfer(&_WrappedEther.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Transfer(&_WrappedEther.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.TransferFrom(&_WrappedEther.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.TransferFrom(&_WrappedEther.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Withdraw(&_WrappedEther.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Withdraw(&_WrappedEther.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherSession) Receive() (*types.Transaction, error) {
	return _WrappedEther.Contract.Receive(&_WrappedEther.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedEther.Contract.Receive(&_WrappedEther.TransactOpts)
}

// WrappedEtherApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedEther contract.
type WrappedEtherApprovalIterator struct {
	Event *WrappedEtherApproval // Event containing the contract specifics and raw log

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
func (it *WrappedEtherApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherApproval)
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
		it.Event = new(WrappedEtherApproval)
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
func (it *WrappedEtherApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherApproval represents a Approval event raised by the WrappedEther contract.
type WrappedEtherApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedEtherApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherApprovalIterator{contract: _WrappedEther.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedEtherApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherApproval)
				if err := _WrappedEther.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WrappedEther *WrappedEtherFilterer) ParseApproval(log types.Log) (*WrappedEtherApproval, error) {
	event := new(WrappedEtherApproval)
	if err := _WrappedEther.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WrappedEther contract.
type WrappedEtherDepositIterator struct {
	Event *WrappedEtherDeposit // Event containing the contract specifics and raw log

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
func (it *WrappedEtherDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherDeposit)
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
		it.Event = new(WrappedEtherDeposit)
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
func (it *WrappedEtherDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherDeposit represents a Deposit event raised by the WrappedEther contract.
type WrappedEtherDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WrappedEtherDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherDepositIterator{contract: _WrappedEther.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WrappedEtherDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherDeposit)
				if err := _WrappedEther.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) ParseDeposit(log types.Log) (*WrappedEtherDeposit, error) {
	event := new(WrappedEtherDeposit)
	if err := _WrappedEther.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WrappedEther contract.
type WrappedEtherEIP712DomainChangedIterator struct {
	Event *WrappedEtherEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WrappedEtherEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherEIP712DomainChanged)
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
		it.Event = new(WrappedEtherEIP712DomainChanged)
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
func (it *WrappedEtherEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherEIP712DomainChanged represents a EIP712DomainChanged event raised by the WrappedEther contract.
type WrappedEtherEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WrappedEther *WrappedEtherFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WrappedEtherEIP712DomainChangedIterator, error) {

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WrappedEtherEIP712DomainChangedIterator{contract: _WrappedEther.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WrappedEther *WrappedEtherFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WrappedEtherEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherEIP712DomainChanged)
				if err := _WrappedEther.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_WrappedEther *WrappedEtherFilterer) ParseEIP712DomainChanged(log types.Log) (*WrappedEtherEIP712DomainChanged, error) {
	event := new(WrappedEtherEIP712DomainChanged)
	if err := _WrappedEther.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedEther contract.
type WrappedEtherTransferIterator struct {
	Event *WrappedEtherTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedEtherTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherTransfer)
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
		it.Event = new(WrappedEtherTransfer)
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
func (it *WrappedEtherTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherTransfer represents a Transfer event raised by the WrappedEther contract.
type WrappedEtherTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedEtherTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherTransferIterator{contract: _WrappedEther.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedEtherTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherTransfer)
				if err := _WrappedEther.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WrappedEther *WrappedEtherFilterer) ParseTransfer(log types.Log) (*WrappedEtherTransfer, error) {
	event := new(WrappedEtherTransfer)
	if err := _WrappedEther.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WrappedEther contract.
type WrappedEtherWithdrawalIterator struct {
	Event *WrappedEtherWithdrawal // Event containing the contract specifics and raw log

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
func (it *WrappedEtherWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherWithdrawal)
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
		it.Event = new(WrappedEtherWithdrawal)
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
func (it *WrappedEtherWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherWithdrawal represents a Withdrawal event raised by the WrappedEther contract.
type WrappedEtherWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WrappedEtherWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherWithdrawalIterator{contract: _WrappedEther.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WrappedEtherWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherWithdrawal)
				if err := _WrappedEther.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) ParseWithdrawal(log types.Log) (*WrappedEtherWithdrawal, error) {
	event := new(WrappedEtherWithdrawal)
	if err := _WrappedEther.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
