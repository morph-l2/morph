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

// L2LidoGatewayMetaData contains all meta data concerning the L2LidoGateway contract.
var L2LidoGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorAccountIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorCallerIsNotDepositsDisabler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorCallerIsNotDepositsEnabler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorCallerIsNotWithdrawalsDisabler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorCallerIsNotWithdrawalsEnabler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorDepositsDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorDepositsEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorNonZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorUnsupportedL1Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorUnsupportedL2Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorWithdrawZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorWithdrawalsDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorWithdrawalsEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawAndCallIsNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"disabler\",\"type\":\"address\"}],\"name\":\"DepositsDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enabler\",\"type\":\"address\"}],\"name\":\"DepositsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"disabler\",\"type\":\"address\"}],\"name\":\"WithdrawalsDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enabler\",\"type\":\"address\"}],\"name\":\"WithdrawalsEnabled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSITS_DISABLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSITS_ENABLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWALS_DISABLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWALS_ENABLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositsEnabler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositsDisabler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawalsEnabler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_withdrawalsDisabler\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDepositsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWithdrawalsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b50604051620029a6380380620029a68339810160408190526200003391620001dc565b6001600160a01b03808316608081905290821660a0526200009b5760405162461bcd60e51b815260206004820152601460248201527f7a65726f206c31746f6b656e206164647265737300000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b038116620000f35760405162461bcd60e51b815260206004820152601460248201527f7a65726f206c32546f6b656e2061646472657373000000000000000000000000604482015260640162000092565b620000fd62000105565b505062000212565b5f54610100900460ff16156200016e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840162000092565b5f5460ff90811614620001be575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114620001d7575f80fd5b919050565b5f8060408385031215620001ee575f80fd5b620001f983620001c0565b91506200020960208401620001c0565b90509250929050565b60805160a051612733620002735f395f81816102610152818161062a01528181610a1d01528181611077015261147e01525f818161049b01528181610696015281816109b10152818161100b015281816116a1015261187101526127335ff3fe6080604052600436106101d0575f3560e01c80639010d07c116100fd578063ca15c87311610092578063f27ebced11610062578063f27ebced14610580578063f2fde38b1461059f578063f887ea40146105be578063fadcc54a146105dd575f80fd5b8063ca15c873146104fb578063d547741f1461051a578063e3b523e314610539578063e8bac93b1461056c575f80fd5b8063ad960ce1116100cd578063ad960ce114610476578063c01e1bd61461048a578063c0c53b8b146104bd578063c676ad29146104dc575f80fd5b80639010d07c1461041157806391d1485414610430578063a93a4af91461044f578063ac67e1af14610462575f80fd5b80635ed2c22011610173578063797594b011610143578063797594b01461038f5780638431f5c1146103ae5780638d7601c0146103c15780638da5cb5b146103f4575f80fd5b80635ed2c220146102ec5780636c07ea43146103275780636f18bd221461033a578063715018a61461037b575f80fd5b806356eff267116101ae57806356eff26714610250578063575361b6146102835780635777bf50146102965780635e4c57a4146102d8575f80fd5b80632f2ff15d146101d45780633cb747bf146101f557806354bbd59c14610231575b5f80fd5b3480156101df575f80fd5b506101f36101ee36600461211c565b610610565b005b348015610200575f80fd5b50609954610214906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561023c575f80fd5b5061021461024b36600461214a565b610626565b34801561025b575f80fd5b506102147f000000000000000000000000000000000000000000000000000000000000000081565b6101f36102913660046121b1565b6106bb565b3480156102a1575f80fd5b507fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec5460ff165b6040519015158152602001610228565b3480156102e3575f80fd5b506101f3610706565b3480156102f7575f80fd5b507fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec54610100900460ff166102c8565b6101f3610335366004612227565b610838565b348015610345575f80fd5b5061036d7f63f736f21cb2943826cd50b191eb054ebbea670e4e962d0527611f830cd399d681565b604051908152602001610228565b348015610386575f80fd5b506101f3610876565b34801561039a575f80fd5b50609754610214906001600160a01b031681565b6101f36103bc366004612259565b610889565b3480156103cc575f80fd5b5061036d7f94a954c0bc99227eddbc0715a62a7e1056ed8784cd719c2303b685683908857c81565b3480156103ff575f80fd5b506065546001600160a01b0316610214565b34801561041c575f80fd5b5061021461042b3660046122eb565b610bff565b34801561043b575f80fd5b506102c861044a36600461211c565b610c3e565b6101f361045d36600461230b565b610c74565b34801561046d575f80fd5b506101f3610c86565b348015610481575f80fd5b506101f3610db4565b348015610495575f80fd5b506102147f000000000000000000000000000000000000000000000000000000000000000081565b3480156104c8575f80fd5b506101f36104d736600461234e565b610eca565b3480156104e7575f80fd5b506102146104f636600461214a565b611007565b348015610506575f80fd5b5061036d610515366004612396565b61109c565b348015610525575f80fd5b506101f361053436600461211c565b6110d1565b348015610544575f80fd5b5061036d7f9ab8816a3dc0b3849ec1ac00483f6ec815b07eee2fd766a353311c823ad59d0d81565b348015610577575f80fd5b506101f36110e3565b34801561058b575f80fd5b506101f361059a3660046123ad565b6111fe565b3480156105aa575f80fd5b506101f36105b936600461214a565b611313565b3480156105c9575f80fd5b50609854610214906001600160a01b031681565b3480156105e8575f80fd5b5061036d7f4b43b36766bde12c5e9cbbc37d15f8d1f769f08f54720ab370faeb4ce893753a81565b6106186113a3565b61062282826113fd565b5050565b5f817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614610693576040517f6251ce6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b507f000000000000000000000000000000000000000000000000000000000000000092915050565b6106fe86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250611473915050565b505050505050565b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec5460ff1615610762576040517f4f2c8be200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61078c7f4b43b36766bde12c5e9cbbc37d15f8d1f769f08f54720ab370faeb4ce893753a33610c3e565b6107c2576040517f3d39c5f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560405133907fc36a428b063177e3f28b3b5d340c08f77827847b2ee30114ccf0c40e519c420a905f90a2565b6108718333845f5b6040519080825280601f01601f19166020018201604052801561086a576020820181803683370190505b5085611473565b505050565b61087e6113a3565b6108875f6118eb565b565b6099546001600160a01b03163381146108e95760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610925573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109499190612433565b6097546001600160a01b039081169116146109a65760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016108e0565b6109ae611954565b877f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614610a1a576040517ffe15603f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b877f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614610a86576040517f6251ce6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec5460ff16610ae1576040517fa185a6b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3415610b19576040517f3ddcf11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038881166004830152602482018890528a16906340c10f19906044015f604051808303815f87803b158015610b79575f80fd5b505af1158015610b8b573d5f803e3d5ffd5b50505050876001600160a01b0316896001600160a01b03168b6001600160a01b03167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348a8a8a8a604051610be2949392919061244e565b60405180910390a45050610bf560018055565b5050505050505050565b5f8281527fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ed60205260408120610c3590836119b3565b90505b92915050565b5f8281527fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ed60205260408120610c3590836119be565b610c808484845f610840565b50505050565b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec5460ff16610ce1576040517fa185a6b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d0b7f63f736f21cb2943826cd50b191eb054ebbea670e4e962d0527611f830cd399d633610c3e565b610d41576040517fadd9524c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560405133907f9ca4d309bbfd23c65db3dc38c1712862f5812c7139937e2655de86e803f73bb9905f90a2565b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec54610100900460ff16610e14576040517f77d195b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e3e7f94a954c0bc99227eddbc0715a62a7e1056ed8784cd719c2303b685683908857c33610c3e565b610e74576040517f9e60ca7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec805461ff001916905560405133907f644eeba8ede48fefc32ada09fb240c5f6c0f06507ab1d296d5af41f1521d9fcb905f90a2565b5f54610100900460ff1615808015610ee857505f54600160ff909116105b80610f015750303b158015610f0157505f5460ff166001145b610f735760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108e0565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610fb2575f805461ff0019166101001790555b610fbd8484846119df565b8015610c80575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b5f817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614611074576040517ffe15603f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b507f000000000000000000000000000000000000000000000000000000000000000092915050565b5f8181527fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ed60205260408120610c3890611b22565b6110d96113a3565b6106228282611b2b565b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec54610100900460ff1615611144576040517ff74ad25400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61116e7f9ab8816a3dc0b3849ec1ac00483f6ec815b07eee2fd766a353311c823ad59d0d33610c3e565b6111a4576040517f5c16894300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec805461ff00191661010017905560405133907fb2ed3603bd9051f0182ebfb75f12a21059b4d31b578a2a05c8d0245e9e2d3204905f90a2565b5f54600290610100900460ff1615801561121e57505f5460ff8083169116105b6112905760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108e0565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790556112cc85858585611ba1565b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b61131b6113a3565b6001600160a01b0381166113975760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016108e0565b6113a0816118eb565b50565b6065546001600160a01b031633146108875760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108e0565b5f8281527fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ed602052604090206114339082611d82565b156106225760405133906001600160a01b0383169084907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d905f90a45050565b61147b611954565b847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316146114e7576040517f6251ce6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b038116611528576040517fef6b416200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec54610100900460ff16611588576040517f77d195b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f036115c1576040517f6c18829600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60985433906001600160a01b03168190036115ef57848060200190518101906115ea91906124d3565b955090505b845115611628576040517f998bcee200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f9dc29fac0000000000000000000000000000000000000000000000000000000081526001600160a01b03828116600483015260248201889052891690639dc29fac906044015f604051808303815f87803b158015611688575f80fd5b505af115801561169a573d5f803e3d5ffd5b505050505f7f000000000000000000000000000000000000000000000000000000000000000089838a8a8a6040516024016116da969594939291906125f8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa1580156117b1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117d59190612645565b6099546097546040517fb2267a7b0000000000000000000000000000000000000000000000000000000081529293506001600160a01b039182169263b2267a7b92349261182d929116905f9088908d9060040161265c565b5f604051808303818588803b158015611844575f80fd5b505af1158015611856573d5f803e3d5ffd5b5050505050826001600160a01b03168a6001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48c8c8c876040516118ce949392919061265c565b60405180910390a450505050506118e460018055565b5050505050565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6002600154036119a65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016108e0565b6002600155565b60018055565b5f610c358383611d96565b6001600160a01b0381165f9081526001830160205260408120541515610c35565b6001600160a01b038316611a355760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016108e0565b6001600160a01b038116611a8b5760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016108e0565b611a93611dbc565b611a9b611e40565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561087157609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f610c38825490565b5f8281527fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ed60205260409020611b619082611ec4565b156106225760405133906001600160a01b0383169084907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b905f90a45050565b5f54610100900460ff16611c1d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e0565b7fe20dc80161c3a3e412098d054775959b6cab7cf9e3d46b04fee5a64d0898f0ec80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117815560405133907fc36a428b063177e3f28b3b5d340c08f77827847b2ee30114ccf0c40e519c420a905f90a2805461ff001916610100178155611ca53390565b6001600160a01b03167fb2ed3603bd9051f0182ebfb75f12a21059b4d31b578a2a05c8d0245e9e2d320460405160405180910390a2611d047f4b43b36766bde12c5e9cbbc37d15f8d1f769f08f54720ab370faeb4ce893753a866113fd565b611d2e7f63f736f21cb2943826cd50b191eb054ebbea670e4e962d0527611f830cd399d6856113fd565b611d587f9ab8816a3dc0b3849ec1ac00483f6ec815b07eee2fd766a353311c823ad59d0d846113fd565b6118e47f94a954c0bc99227eddbc0715a62a7e1056ed8784cd719c2303b685683908857c836113fd565b5f610c35836001600160a01b038416611ed8565b5f825f018281548110611dab57611dab612694565b905f5260205f200154905092915050565b5f54610100900460ff16611e385760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e0565b610887611f24565b5f54610100900460ff16611ebc5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e0565b610887611fa0565b5f610c35836001600160a01b038416612025565b5f818152600183016020526040812054611f1d57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610c38565b505f610c38565b5f54610100900460ff166119ad5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e0565b5f54610100900460ff1661201c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e0565b610887336118eb565b5f81815260018301602052604081205480156120ff575f6120476001836126c1565b85549091505f9061205a906001906126c1565b90508181146120b9575f865f01828154811061207857612078612694565b905f5260205f200154905080875f01848154811061209857612098612694565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806120ca576120ca6126f9565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610c38565b5f915050610c38565b6001600160a01b03811681146113a0575f80fd5b5f806040838503121561212d575f80fd5b82359150602083013561213f81612108565b809150509250929050565b5f6020828403121561215a575f80fd5b813561216581612108565b9392505050565b5f8083601f84011261217c575f80fd5b50813567ffffffffffffffff811115612193575f80fd5b6020830191508360208285010111156121aa575f80fd5b9250929050565b5f805f805f8060a087890312156121c6575f80fd5b86356121d181612108565b955060208701356121e181612108565b945060408701359350606087013567ffffffffffffffff811115612203575f80fd5b61220f89828a0161216c565b979a9699509497949695608090950135949350505050565b5f805f60608486031215612239575f80fd5b833561224481612108565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a03121561226f575f80fd5b873561227a81612108565b9650602088013561228a81612108565b9550604088013561229a81612108565b945060608801356122aa81612108565b93506080880135925060a088013567ffffffffffffffff8111156122cc575f80fd5b6122d88a828b0161216c565b989b979a50959850939692959293505050565b5f80604083850312156122fc575f80fd5b50508035926020909101359150565b5f805f806080858703121561231e575f80fd5b843561232981612108565b9350602085013561233981612108565b93969395505050506040820135916060013590565b5f805f60608486031215612360575f80fd5b833561236b81612108565b9250602084013561237b81612108565b9150604084013561238b81612108565b809150509250925092565b5f602082840312156123a6575f80fd5b5035919050565b5f805f80608085870312156123c0575f80fd5b84356123cb81612108565b935060208501356123db81612108565b925060408501356123eb81612108565b915060608501356123fb81612108565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215612443575f80fd5b815161216581612108565b6001600160a01b038516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b838110156124cb5781810151838201526020016124b3565b50505f910152565b5f80604083850312156124e4575f80fd5b82516124ef81612108565b602084015190925067ffffffffffffffff8082111561250c575f80fd5b818501915085601f83011261251f575f80fd5b81518181111561253157612531612406565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561257757612577612406565b8160405282815288602084870101111561258f575f80fd5b6125a08360208301602088016124b1565b80955050505050509250929050565b5f81518084526125c68160208601602086016124b1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261263960c08301846125af565b98975050505050505050565b5f60208284031215612655575f80fd5b5051919050565b6001600160a01b0385168152836020820152608060408201525f61268360808301856125af565b905082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b81810381811115610c38577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c6343000818000a",
}

// L2LidoGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2LidoGatewayMetaData.ABI instead.
var L2LidoGatewayABI = L2LidoGatewayMetaData.ABI

// L2LidoGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2LidoGatewayMetaData.Bin instead.
var L2LidoGatewayBin = L2LidoGatewayMetaData.Bin

// DeployL2LidoGateway deploys a new Ethereum contract, binding an instance of L2LidoGateway to it.
func DeployL2LidoGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _l1Token common.Address, _l2Token common.Address) (common.Address, *types.Transaction, *L2LidoGateway, error) {
	parsed, err := L2LidoGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2LidoGatewayBin), backend, _l1Token, _l2Token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2LidoGateway{L2LidoGatewayCaller: L2LidoGatewayCaller{contract: contract}, L2LidoGatewayTransactor: L2LidoGatewayTransactor{contract: contract}, L2LidoGatewayFilterer: L2LidoGatewayFilterer{contract: contract}}, nil
}

// L2LidoGateway is an auto generated Go binding around an Ethereum contract.
type L2LidoGateway struct {
	L2LidoGatewayCaller     // Read-only binding to the contract
	L2LidoGatewayTransactor // Write-only binding to the contract
	L2LidoGatewayFilterer   // Log filterer for contract events
}

// L2LidoGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2LidoGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2LidoGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2LidoGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2LidoGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2LidoGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2LidoGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2LidoGatewaySession struct {
	Contract     *L2LidoGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2LidoGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2LidoGatewayCallerSession struct {
	Contract *L2LidoGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2LidoGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2LidoGatewayTransactorSession struct {
	Contract     *L2LidoGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2LidoGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2LidoGatewayRaw struct {
	Contract *L2LidoGateway // Generic contract binding to access the raw methods on
}

// L2LidoGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2LidoGatewayCallerRaw struct {
	Contract *L2LidoGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2LidoGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2LidoGatewayTransactorRaw struct {
	Contract *L2LidoGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2LidoGateway creates a new instance of L2LidoGateway, bound to a specific deployed contract.
func NewL2LidoGateway(address common.Address, backend bind.ContractBackend) (*L2LidoGateway, error) {
	contract, err := bindL2LidoGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2LidoGateway{L2LidoGatewayCaller: L2LidoGatewayCaller{contract: contract}, L2LidoGatewayTransactor: L2LidoGatewayTransactor{contract: contract}, L2LidoGatewayFilterer: L2LidoGatewayFilterer{contract: contract}}, nil
}

// NewL2LidoGatewayCaller creates a new read-only instance of L2LidoGateway, bound to a specific deployed contract.
func NewL2LidoGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2LidoGatewayCaller, error) {
	contract, err := bindL2LidoGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayCaller{contract: contract}, nil
}

// NewL2LidoGatewayTransactor creates a new write-only instance of L2LidoGateway, bound to a specific deployed contract.
func NewL2LidoGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2LidoGatewayTransactor, error) {
	contract, err := bindL2LidoGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayTransactor{contract: contract}, nil
}

// NewL2LidoGatewayFilterer creates a new log filterer instance of L2LidoGateway, bound to a specific deployed contract.
func NewL2LidoGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2LidoGatewayFilterer, error) {
	contract, err := bindL2LidoGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayFilterer{contract: contract}, nil
}

// bindL2LidoGateway binds a generic wrapper to an already deployed contract.
func bindL2LidoGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2LidoGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2LidoGateway *L2LidoGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2LidoGateway.Contract.L2LidoGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2LidoGateway *L2LidoGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.L2LidoGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2LidoGateway *L2LidoGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.L2LidoGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2LidoGateway *L2LidoGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2LidoGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2LidoGateway *L2LidoGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2LidoGateway *L2LidoGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITSDISABLERROLE is a free data retrieval call binding the contract method 0x6f18bd22.
//
// Solidity: function DEPOSITS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCaller) DEPOSITSDISABLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "DEPOSITS_DISABLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSDISABLERROLE is a free data retrieval call binding the contract method 0x6f18bd22.
//
// Solidity: function DEPOSITS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewaySession) DEPOSITSDISABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.DEPOSITSDISABLERROLE(&_L2LidoGateway.CallOpts)
}

// DEPOSITSDISABLERROLE is a free data retrieval call binding the contract method 0x6f18bd22.
//
// Solidity: function DEPOSITS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCallerSession) DEPOSITSDISABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.DEPOSITSDISABLERROLE(&_L2LidoGateway.CallOpts)
}

// DEPOSITSENABLERROLE is a free data retrieval call binding the contract method 0xfadcc54a.
//
// Solidity: function DEPOSITS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCaller) DEPOSITSENABLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "DEPOSITS_ENABLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSENABLERROLE is a free data retrieval call binding the contract method 0xfadcc54a.
//
// Solidity: function DEPOSITS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewaySession) DEPOSITSENABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.DEPOSITSENABLERROLE(&_L2LidoGateway.CallOpts)
}

// DEPOSITSENABLERROLE is a free data retrieval call binding the contract method 0xfadcc54a.
//
// Solidity: function DEPOSITS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCallerSession) DEPOSITSENABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.DEPOSITSENABLERROLE(&_L2LidoGateway.CallOpts)
}

// WITHDRAWALSDISABLERROLE is a free data retrieval call binding the contract method 0x8d7601c0.
//
// Solidity: function WITHDRAWALS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCaller) WITHDRAWALSDISABLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "WITHDRAWALS_DISABLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALSDISABLERROLE is a free data retrieval call binding the contract method 0x8d7601c0.
//
// Solidity: function WITHDRAWALS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewaySession) WITHDRAWALSDISABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.WITHDRAWALSDISABLERROLE(&_L2LidoGateway.CallOpts)
}

// WITHDRAWALSDISABLERROLE is a free data retrieval call binding the contract method 0x8d7601c0.
//
// Solidity: function WITHDRAWALS_DISABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCallerSession) WITHDRAWALSDISABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.WITHDRAWALSDISABLERROLE(&_L2LidoGateway.CallOpts)
}

// WITHDRAWALSENABLERROLE is a free data retrieval call binding the contract method 0xe3b523e3.
//
// Solidity: function WITHDRAWALS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCaller) WITHDRAWALSENABLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "WITHDRAWALS_ENABLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALSENABLERROLE is a free data retrieval call binding the contract method 0xe3b523e3.
//
// Solidity: function WITHDRAWALS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewaySession) WITHDRAWALSENABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.WITHDRAWALSENABLERROLE(&_L2LidoGateway.CallOpts)
}

// WITHDRAWALSENABLERROLE is a free data retrieval call binding the contract method 0xe3b523e3.
//
// Solidity: function WITHDRAWALS_ENABLER_ROLE() view returns(bytes32)
func (_L2LidoGateway *L2LidoGatewayCallerSession) WITHDRAWALSENABLERROLE() ([32]byte, error) {
	return _L2LidoGateway.Contract.WITHDRAWALSENABLERROLE(&_L2LidoGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) Counterpart() (common.Address, error) {
	return _L2LidoGateway.Contract.Counterpart(&_L2LidoGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2LidoGateway.Contract.Counterpart(&_L2LidoGateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2LidoGateway.Contract.GetL1ERC20Address(&_L2LidoGateway.CallOpts, _l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2LidoGateway.Contract.GetL1ERC20Address(&_L2LidoGateway.CallOpts, _l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2LidoGateway.Contract.GetL2ERC20Address(&_L2LidoGateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2LidoGateway.Contract.GetL2ERC20Address(&_L2LidoGateway.CallOpts, _l1Token)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 _role, uint256 _index) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) GetRoleMember(opts *bind.CallOpts, _role [32]byte, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "getRoleMember", _role, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 _role, uint256 _index) view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) GetRoleMember(_role [32]byte, _index *big.Int) (common.Address, error) {
	return _L2LidoGateway.Contract.GetRoleMember(&_L2LidoGateway.CallOpts, _role, _index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 _role, uint256 _index) view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) GetRoleMember(_role [32]byte, _index *big.Int) (common.Address, error) {
	return _L2LidoGateway.Contract.GetRoleMember(&_L2LidoGateway.CallOpts, _role, _index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 _role) view returns(uint256)
func (_L2LidoGateway *L2LidoGatewayCaller) GetRoleMemberCount(opts *bind.CallOpts, _role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "getRoleMemberCount", _role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 _role) view returns(uint256)
func (_L2LidoGateway *L2LidoGatewaySession) GetRoleMemberCount(_role [32]byte) (*big.Int, error) {
	return _L2LidoGateway.Contract.GetRoleMemberCount(&_L2LidoGateway.CallOpts, _role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 _role) view returns(uint256)
func (_L2LidoGateway *L2LidoGatewayCallerSession) GetRoleMemberCount(_role [32]byte) (*big.Int, error) {
	return _L2LidoGateway.Contract.GetRoleMemberCount(&_L2LidoGateway.CallOpts, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCaller) HasRole(opts *bind.CallOpts, _role [32]byte, _account common.Address) (bool, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "hasRole", _role, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_L2LidoGateway *L2LidoGatewaySession) HasRole(_role [32]byte, _account common.Address) (bool, error) {
	return _L2LidoGateway.Contract.HasRole(&_L2LidoGateway.CallOpts, _role, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCallerSession) HasRole(_role [32]byte, _account common.Address) (bool, error) {
	return _L2LidoGateway.Contract.HasRole(&_L2LidoGateway.CallOpts, _role, _account)
}

// IsDepositsEnabled is a free data retrieval call binding the contract method 0x5777bf50.
//
// Solidity: function isDepositsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCaller) IsDepositsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "isDepositsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositsEnabled is a free data retrieval call binding the contract method 0x5777bf50.
//
// Solidity: function isDepositsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewaySession) IsDepositsEnabled() (bool, error) {
	return _L2LidoGateway.Contract.IsDepositsEnabled(&_L2LidoGateway.CallOpts)
}

// IsDepositsEnabled is a free data retrieval call binding the contract method 0x5777bf50.
//
// Solidity: function isDepositsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCallerSession) IsDepositsEnabled() (bool, error) {
	return _L2LidoGateway.Contract.IsDepositsEnabled(&_L2LidoGateway.CallOpts)
}

// IsWithdrawalsEnabled is a free data retrieval call binding the contract method 0x5ed2c220.
//
// Solidity: function isWithdrawalsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCaller) IsWithdrawalsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "isWithdrawalsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalsEnabled is a free data retrieval call binding the contract method 0x5ed2c220.
//
// Solidity: function isWithdrawalsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewaySession) IsWithdrawalsEnabled() (bool, error) {
	return _L2LidoGateway.Contract.IsWithdrawalsEnabled(&_L2LidoGateway.CallOpts)
}

// IsWithdrawalsEnabled is a free data retrieval call binding the contract method 0x5ed2c220.
//
// Solidity: function isWithdrawalsEnabled() view returns(bool)
func (_L2LidoGateway *L2LidoGatewayCallerSession) IsWithdrawalsEnabled() (bool, error) {
	return _L2LidoGateway.Contract.IsWithdrawalsEnabled(&_L2LidoGateway.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) L1Token() (common.Address, error) {
	return _L2LidoGateway.Contract.L1Token(&_L2LidoGateway.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) L1Token() (common.Address, error) {
	return _L2LidoGateway.Contract.L1Token(&_L2LidoGateway.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) L2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "l2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) L2Token() (common.Address, error) {
	return _L2LidoGateway.Contract.L2Token(&_L2LidoGateway.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) L2Token() (common.Address, error) {
	return _L2LidoGateway.Contract.L2Token(&_L2LidoGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) Messenger() (common.Address, error) {
	return _L2LidoGateway.Contract.Messenger(&_L2LidoGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2LidoGateway.Contract.Messenger(&_L2LidoGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) Owner() (common.Address, error) {
	return _L2LidoGateway.Contract.Owner(&_L2LidoGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) Owner() (common.Address, error) {
	return _L2LidoGateway.Contract.Owner(&_L2LidoGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2LidoGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2LidoGateway *L2LidoGatewaySession) Router() (common.Address, error) {
	return _L2LidoGateway.Contract.Router(&_L2LidoGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2LidoGateway *L2LidoGatewayCallerSession) Router() (common.Address, error) {
	return _L2LidoGateway.Contract.Router(&_L2LidoGateway.CallOpts)
}

// DisableDeposits is a paid mutator transaction binding the contract method 0xac67e1af.
//
// Solidity: function disableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) DisableDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "disableDeposits")
}

// DisableDeposits is a paid mutator transaction binding the contract method 0xac67e1af.
//
// Solidity: function disableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewaySession) DisableDeposits() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.DisableDeposits(&_L2LidoGateway.TransactOpts)
}

// DisableDeposits is a paid mutator transaction binding the contract method 0xac67e1af.
//
// Solidity: function disableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) DisableDeposits() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.DisableDeposits(&_L2LidoGateway.TransactOpts)
}

// DisableWithdrawals is a paid mutator transaction binding the contract method 0xad960ce1.
//
// Solidity: function disableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) DisableWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "disableWithdrawals")
}

// DisableWithdrawals is a paid mutator transaction binding the contract method 0xad960ce1.
//
// Solidity: function disableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewaySession) DisableWithdrawals() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.DisableWithdrawals(&_L2LidoGateway.TransactOpts)
}

// DisableWithdrawals is a paid mutator transaction binding the contract method 0xad960ce1.
//
// Solidity: function disableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) DisableWithdrawals() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.DisableWithdrawals(&_L2LidoGateway.TransactOpts)
}

// EnableDeposits is a paid mutator transaction binding the contract method 0x5e4c57a4.
//
// Solidity: function enableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) EnableDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "enableDeposits")
}

// EnableDeposits is a paid mutator transaction binding the contract method 0x5e4c57a4.
//
// Solidity: function enableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewaySession) EnableDeposits() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.EnableDeposits(&_L2LidoGateway.TransactOpts)
}

// EnableDeposits is a paid mutator transaction binding the contract method 0x5e4c57a4.
//
// Solidity: function enableDeposits() returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) EnableDeposits() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.EnableDeposits(&_L2LidoGateway.TransactOpts)
}

// EnableWithdrawals is a paid mutator transaction binding the contract method 0xe8bac93b.
//
// Solidity: function enableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) EnableWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "enableWithdrawals")
}

// EnableWithdrawals is a paid mutator transaction binding the contract method 0xe8bac93b.
//
// Solidity: function enableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewaySession) EnableWithdrawals() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.EnableWithdrawals(&_L2LidoGateway.TransactOpts)
}

// EnableWithdrawals is a paid mutator transaction binding the contract method 0xe8bac93b.
//
// Solidity: function enableWithdrawals() returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) EnableWithdrawals() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.EnableWithdrawals(&_L2LidoGateway.TransactOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2LidoGateway *L2LidoGatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.FinalizeDepositERC20(&_L2LidoGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.FinalizeDepositERC20(&_L2LidoGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) GrantRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "grantRole", _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewaySession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.GrantRole(&_L2LidoGateway.TransactOpts, _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.GrantRole(&_L2LidoGateway.TransactOpts, _role, _account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2LidoGateway *L2LidoGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.Initialize(&_L2LidoGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.Initialize(&_L2LidoGateway.TransactOpts, _counterpart, _router, _messenger)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf27ebced.
//
// Solidity: function initializeV2(address _depositsEnabler, address _depositsDisabler, address _withdrawalsEnabler, address _withdrawalsDisabler) returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) InitializeV2(opts *bind.TransactOpts, _depositsEnabler common.Address, _depositsDisabler common.Address, _withdrawalsEnabler common.Address, _withdrawalsDisabler common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "initializeV2", _depositsEnabler, _depositsDisabler, _withdrawalsEnabler, _withdrawalsDisabler)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf27ebced.
//
// Solidity: function initializeV2(address _depositsEnabler, address _depositsDisabler, address _withdrawalsEnabler, address _withdrawalsDisabler) returns()
func (_L2LidoGateway *L2LidoGatewaySession) InitializeV2(_depositsEnabler common.Address, _depositsDisabler common.Address, _withdrawalsEnabler common.Address, _withdrawalsDisabler common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.InitializeV2(&_L2LidoGateway.TransactOpts, _depositsEnabler, _depositsDisabler, _withdrawalsEnabler, _withdrawalsDisabler)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf27ebced.
//
// Solidity: function initializeV2(address _depositsEnabler, address _depositsDisabler, address _withdrawalsEnabler, address _withdrawalsDisabler) returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) InitializeV2(_depositsEnabler common.Address, _depositsDisabler common.Address, _withdrawalsEnabler common.Address, _withdrawalsDisabler common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.InitializeV2(&_L2LidoGateway.TransactOpts, _depositsEnabler, _depositsDisabler, _withdrawalsEnabler, _withdrawalsDisabler)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2LidoGateway *L2LidoGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.RenounceOwnership(&_L2LidoGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2LidoGateway.Contract.RenounceOwnership(&_L2LidoGateway.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) RevokeRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "revokeRole", _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewaySession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.RevokeRole(&_L2LidoGateway.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.RevokeRole(&_L2LidoGateway.TransactOpts, _role, _account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2LidoGateway *L2LidoGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.TransferOwnership(&_L2LidoGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.TransferOwnership(&_L2LidoGateway.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC20(&_L2LidoGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC20(&_L2LidoGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC200(&_L2LidoGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC200(&_L2LidoGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC20AndCall(&_L2LidoGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2LidoGateway *L2LidoGatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2LidoGateway.Contract.WithdrawERC20AndCall(&_L2LidoGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2LidoGatewayDepositsDisabledIterator is returned from FilterDepositsDisabled and is used to iterate over the raw logs and unpacked data for DepositsDisabled events raised by the L2LidoGateway contract.
type L2LidoGatewayDepositsDisabledIterator struct {
	Event *L2LidoGatewayDepositsDisabled // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayDepositsDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayDepositsDisabled)
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
		it.Event = new(L2LidoGatewayDepositsDisabled)
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
func (it *L2LidoGatewayDepositsDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayDepositsDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayDepositsDisabled represents a DepositsDisabled event raised by the L2LidoGateway contract.
type L2LidoGatewayDepositsDisabled struct {
	Disabler common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositsDisabled is a free log retrieval operation binding the contract event 0x9ca4d309bbfd23c65db3dc38c1712862f5812c7139937e2655de86e803f73bb9.
//
// Solidity: event DepositsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterDepositsDisabled(opts *bind.FilterOpts, disabler []common.Address) (*L2LidoGatewayDepositsDisabledIterator, error) {

	var disablerRule []interface{}
	for _, disablerItem := range disabler {
		disablerRule = append(disablerRule, disablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "DepositsDisabled", disablerRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayDepositsDisabledIterator{contract: _L2LidoGateway.contract, event: "DepositsDisabled", logs: logs, sub: sub}, nil
}

// WatchDepositsDisabled is a free log subscription operation binding the contract event 0x9ca4d309bbfd23c65db3dc38c1712862f5812c7139937e2655de86e803f73bb9.
//
// Solidity: event DepositsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchDepositsDisabled(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayDepositsDisabled, disabler []common.Address) (event.Subscription, error) {

	var disablerRule []interface{}
	for _, disablerItem := range disabler {
		disablerRule = append(disablerRule, disablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "DepositsDisabled", disablerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayDepositsDisabled)
				if err := _L2LidoGateway.contract.UnpackLog(event, "DepositsDisabled", log); err != nil {
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

// ParseDepositsDisabled is a log parse operation binding the contract event 0x9ca4d309bbfd23c65db3dc38c1712862f5812c7139937e2655de86e803f73bb9.
//
// Solidity: event DepositsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseDepositsDisabled(log types.Log) (*L2LidoGatewayDepositsDisabled, error) {
	event := new(L2LidoGatewayDepositsDisabled)
	if err := _L2LidoGateway.contract.UnpackLog(event, "DepositsDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayDepositsEnabledIterator is returned from FilterDepositsEnabled and is used to iterate over the raw logs and unpacked data for DepositsEnabled events raised by the L2LidoGateway contract.
type L2LidoGatewayDepositsEnabledIterator struct {
	Event *L2LidoGatewayDepositsEnabled // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayDepositsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayDepositsEnabled)
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
		it.Event = new(L2LidoGatewayDepositsEnabled)
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
func (it *L2LidoGatewayDepositsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayDepositsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayDepositsEnabled represents a DepositsEnabled event raised by the L2LidoGateway contract.
type L2LidoGatewayDepositsEnabled struct {
	Enabler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositsEnabled is a free log retrieval operation binding the contract event 0xc36a428b063177e3f28b3b5d340c08f77827847b2ee30114ccf0c40e519c420a.
//
// Solidity: event DepositsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterDepositsEnabled(opts *bind.FilterOpts, enabler []common.Address) (*L2LidoGatewayDepositsEnabledIterator, error) {

	var enablerRule []interface{}
	for _, enablerItem := range enabler {
		enablerRule = append(enablerRule, enablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "DepositsEnabled", enablerRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayDepositsEnabledIterator{contract: _L2LidoGateway.contract, event: "DepositsEnabled", logs: logs, sub: sub}, nil
}

// WatchDepositsEnabled is a free log subscription operation binding the contract event 0xc36a428b063177e3f28b3b5d340c08f77827847b2ee30114ccf0c40e519c420a.
//
// Solidity: event DepositsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchDepositsEnabled(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayDepositsEnabled, enabler []common.Address) (event.Subscription, error) {

	var enablerRule []interface{}
	for _, enablerItem := range enabler {
		enablerRule = append(enablerRule, enablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "DepositsEnabled", enablerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayDepositsEnabled)
				if err := _L2LidoGateway.contract.UnpackLog(event, "DepositsEnabled", log); err != nil {
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

// ParseDepositsEnabled is a log parse operation binding the contract event 0xc36a428b063177e3f28b3b5d340c08f77827847b2ee30114ccf0c40e519c420a.
//
// Solidity: event DepositsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseDepositsEnabled(log types.Log) (*L2LidoGatewayDepositsEnabled, error) {
	event := new(L2LidoGatewayDepositsEnabled)
	if err := _L2LidoGateway.contract.UnpackLog(event, "DepositsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2LidoGateway contract.
type L2LidoGatewayFinalizeDepositERC20Iterator struct {
	Event *L2LidoGatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayFinalizeDepositERC20)
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
		it.Event = new(L2LidoGatewayFinalizeDepositERC20)
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
func (it *L2LidoGatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2LidoGateway contract.
type L2LidoGatewayFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2LidoGatewayFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayFinalizeDepositERC20Iterator{contract: _L2LidoGateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayFinalizeDepositERC20)
				if err := _L2LidoGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2LidoGatewayFinalizeDepositERC20, error) {
	event := new(L2LidoGatewayFinalizeDepositERC20)
	if err := _L2LidoGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2LidoGateway contract.
type L2LidoGatewayInitializedIterator struct {
	Event *L2LidoGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayInitialized)
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
		it.Event = new(L2LidoGatewayInitialized)
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
func (it *L2LidoGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayInitialized represents a Initialized event raised by the L2LidoGateway contract.
type L2LidoGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2LidoGatewayInitializedIterator, error) {

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayInitializedIterator{contract: _L2LidoGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayInitialized)
				if err := _L2LidoGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseInitialized(log types.Log) (*L2LidoGatewayInitialized, error) {
	event := new(L2LidoGatewayInitialized)
	if err := _L2LidoGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2LidoGateway contract.
type L2LidoGatewayOwnershipTransferredIterator struct {
	Event *L2LidoGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayOwnershipTransferred)
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
		it.Event = new(L2LidoGatewayOwnershipTransferred)
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
func (it *L2LidoGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2LidoGateway contract.
type L2LidoGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2LidoGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayOwnershipTransferredIterator{contract: _L2LidoGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayOwnershipTransferred)
				if err := _L2LidoGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2LidoGatewayOwnershipTransferred, error) {
	event := new(L2LidoGatewayOwnershipTransferred)
	if err := _L2LidoGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2LidoGateway contract.
type L2LidoGatewayRoleGrantedIterator struct {
	Event *L2LidoGatewayRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayRoleGranted)
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
		it.Event = new(L2LidoGatewayRoleGranted)
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
func (it *L2LidoGatewayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayRoleGranted represents a RoleGranted event raised by the L2LidoGateway contract.
type L2LidoGatewayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2LidoGatewayRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayRoleGrantedIterator{contract: _L2LidoGateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayRoleGranted)
				if err := _L2LidoGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseRoleGranted(log types.Log) (*L2LidoGatewayRoleGranted, error) {
	event := new(L2LidoGatewayRoleGranted)
	if err := _L2LidoGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2LidoGateway contract.
type L2LidoGatewayRoleRevokedIterator struct {
	Event *L2LidoGatewayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayRoleRevoked)
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
		it.Event = new(L2LidoGatewayRoleRevoked)
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
func (it *L2LidoGatewayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayRoleRevoked represents a RoleRevoked event raised by the L2LidoGateway contract.
type L2LidoGatewayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2LidoGatewayRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayRoleRevokedIterator{contract: _L2LidoGateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayRoleRevoked)
				if err := _L2LidoGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseRoleRevoked(log types.Log) (*L2LidoGatewayRoleRevoked, error) {
	event := new(L2LidoGatewayRoleRevoked)
	if err := _L2LidoGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawERC20Iterator struct {
	Event *L2LidoGatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayWithdrawERC20)
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
		it.Event = new(L2LidoGatewayWithdrawERC20)
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
func (it *L2LidoGatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2LidoGatewayWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayWithdrawERC20Iterator{contract: _L2LidoGateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayWithdrawERC20)
				if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2LidoGatewayWithdrawERC20, error) {
	event := new(L2LidoGatewayWithdrawERC20)
	if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayWithdrawalsDisabledIterator is returned from FilterWithdrawalsDisabled and is used to iterate over the raw logs and unpacked data for WithdrawalsDisabled events raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawalsDisabledIterator struct {
	Event *L2LidoGatewayWithdrawalsDisabled // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayWithdrawalsDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayWithdrawalsDisabled)
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
		it.Event = new(L2LidoGatewayWithdrawalsDisabled)
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
func (it *L2LidoGatewayWithdrawalsDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayWithdrawalsDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayWithdrawalsDisabled represents a WithdrawalsDisabled event raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawalsDisabled struct {
	Disabler common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsDisabled is a free log retrieval operation binding the contract event 0x644eeba8ede48fefc32ada09fb240c5f6c0f06507ab1d296d5af41f1521d9fcb.
//
// Solidity: event WithdrawalsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterWithdrawalsDisabled(opts *bind.FilterOpts, disabler []common.Address) (*L2LidoGatewayWithdrawalsDisabledIterator, error) {

	var disablerRule []interface{}
	for _, disablerItem := range disabler {
		disablerRule = append(disablerRule, disablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "WithdrawalsDisabled", disablerRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayWithdrawalsDisabledIterator{contract: _L2LidoGateway.contract, event: "WithdrawalsDisabled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsDisabled is a free log subscription operation binding the contract event 0x644eeba8ede48fefc32ada09fb240c5f6c0f06507ab1d296d5af41f1521d9fcb.
//
// Solidity: event WithdrawalsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchWithdrawalsDisabled(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayWithdrawalsDisabled, disabler []common.Address) (event.Subscription, error) {

	var disablerRule []interface{}
	for _, disablerItem := range disabler {
		disablerRule = append(disablerRule, disablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "WithdrawalsDisabled", disablerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayWithdrawalsDisabled)
				if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawalsDisabled", log); err != nil {
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

// ParseWithdrawalsDisabled is a log parse operation binding the contract event 0x644eeba8ede48fefc32ada09fb240c5f6c0f06507ab1d296d5af41f1521d9fcb.
//
// Solidity: event WithdrawalsDisabled(address indexed disabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseWithdrawalsDisabled(log types.Log) (*L2LidoGatewayWithdrawalsDisabled, error) {
	event := new(L2LidoGatewayWithdrawalsDisabled)
	if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawalsDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2LidoGatewayWithdrawalsEnabledIterator is returned from FilterWithdrawalsEnabled and is used to iterate over the raw logs and unpacked data for WithdrawalsEnabled events raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawalsEnabledIterator struct {
	Event *L2LidoGatewayWithdrawalsEnabled // Event containing the contract specifics and raw log

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
func (it *L2LidoGatewayWithdrawalsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2LidoGatewayWithdrawalsEnabled)
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
		it.Event = new(L2LidoGatewayWithdrawalsEnabled)
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
func (it *L2LidoGatewayWithdrawalsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2LidoGatewayWithdrawalsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2LidoGatewayWithdrawalsEnabled represents a WithdrawalsEnabled event raised by the L2LidoGateway contract.
type L2LidoGatewayWithdrawalsEnabled struct {
	Enabler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsEnabled is a free log retrieval operation binding the contract event 0xb2ed3603bd9051f0182ebfb75f12a21059b4d31b578a2a05c8d0245e9e2d3204.
//
// Solidity: event WithdrawalsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) FilterWithdrawalsEnabled(opts *bind.FilterOpts, enabler []common.Address) (*L2LidoGatewayWithdrawalsEnabledIterator, error) {

	var enablerRule []interface{}
	for _, enablerItem := range enabler {
		enablerRule = append(enablerRule, enablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.FilterLogs(opts, "WithdrawalsEnabled", enablerRule)
	if err != nil {
		return nil, err
	}
	return &L2LidoGatewayWithdrawalsEnabledIterator{contract: _L2LidoGateway.contract, event: "WithdrawalsEnabled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsEnabled is a free log subscription operation binding the contract event 0xb2ed3603bd9051f0182ebfb75f12a21059b4d31b578a2a05c8d0245e9e2d3204.
//
// Solidity: event WithdrawalsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) WatchWithdrawalsEnabled(opts *bind.WatchOpts, sink chan<- *L2LidoGatewayWithdrawalsEnabled, enabler []common.Address) (event.Subscription, error) {

	var enablerRule []interface{}
	for _, enablerItem := range enabler {
		enablerRule = append(enablerRule, enablerItem)
	}

	logs, sub, err := _L2LidoGateway.contract.WatchLogs(opts, "WithdrawalsEnabled", enablerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2LidoGatewayWithdrawalsEnabled)
				if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawalsEnabled", log); err != nil {
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

// ParseWithdrawalsEnabled is a log parse operation binding the contract event 0xb2ed3603bd9051f0182ebfb75f12a21059b4d31b578a2a05c8d0245e9e2d3204.
//
// Solidity: event WithdrawalsEnabled(address indexed enabler)
func (_L2LidoGateway *L2LidoGatewayFilterer) ParseWithdrawalsEnabled(log types.Log) (*L2LidoGatewayWithdrawalsEnabled, error) {
	event := new(L2LidoGatewayWithdrawalsEnabled)
	if err := _L2LidoGateway.contract.UnpackLog(event, "WithdrawalsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
