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

// L2WETHGatewayMetaData contains all meta data concerning the L2WETHGateway contract.
var L2WETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L1_WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b5060405162001d9838038062001d9883398101604081905262000033916200012f565b6200003d62000055565b6001600160a01b0391821660a0521660805262000165565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161462000111575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200012a575f80fd5b919050565b5f806040838503121562000141575f80fd5b6200014c8362000113565b91506200015c6020840162000113565b90509250929050565b60805160a051611be1620001b75f395f818160ef0152818161018001528181610327015281816105c60152610a6f01525f81816101cf0152818161022d015281816105450152610ba20152611be15ff3fe6080604052600436106100e7575f3560e01c8063797594b011610087578063c0c53b8b11610057578063c0c53b8b146102eb578063c676ad291461030a578063f2fde38b14610349578063f887ea4014610368575f80fd5b8063797594b0146102895780638431f5c1146102a85780638da5cb5b146102bb578063a93a4af9146102d8575f80fd5b806354bbd59c116100c257806354bbd59c14610210578063575361b61461024f5780636c07ea4314610262578063715018a614610275575f80fd5b806319c4d4c61461016f5780631efd482a146101be5780633cb747bf146101f1575f80fd5b3661016b57337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146101695760405162461bcd60e51b815260206004820152600960248201527f6f6e6c792057455448000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b005b5f80fd5b34801561017a575f80fd5b506101a27f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b3480156101c9575f80fd5b506101a27f000000000000000000000000000000000000000000000000000000000000000081565b3480156101fc575f80fd5b506099546101a2906001600160a01b031681565b34801561021b575f80fd5b506101a261022a366004611690565b507f000000000000000000000000000000000000000000000000000000000000000090565b61016961025d3660046116f7565b610387565b61016961027036600461176d565b6103d2565b348015610280575f80fd5b50610169610410565b348015610294575f80fd5b506097546101a2906001600160a01b031681565b6101696102b636600461179f565b610423565b3480156102c6575f80fd5b506065546001600160a01b03166101a2565b6101696102e6366004611831565b6107a7565b3480156102f6575f80fd5b50610169610305366004611874565b6107b9565b348015610315575f80fd5b506101a2610324366004611690565b507f000000000000000000000000000000000000000000000000000000000000000090565b348015610354575f80fd5b50610169610363366004611690565b610986565b348015610373575f80fd5b506098546101a2906001600160a01b031681565b6103ca86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610a16915050565b505050505050565b61040b8333845f5b6040519080825280601f01601f191660200182016040528015610404576020820181803683370190505b5085610a16565b505050565b610418610ddf565b6104215f610e39565b565b6099546001600160a01b031633811461047e5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610160565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104de91906118e9565b6097546001600160a01b0390811691161461053b5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610160565b610543610ea2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316886001600160a01b0316146105c45760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610160565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316876001600160a01b0316146106455760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610160565b3484146106945760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d6174636800000000000000000000000000006044820152606401610160565b866001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004015f604051808303818588803b1580156106cd575f80fd5b505af11580156106df573d5f803e3d5ffd5b506106fa935050506001600160a01b03891690508686610efb565b6107398584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610fc292505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348888888860405161078c9493929190611904565b60405180910390a461079d60018055565b5050505050505050565b6107b38484845f6103da565b50505050565b5f54610100900460ff16158080156107d757505f54600160ff909116105b806107f05750303b1580156107f057505f5460ff166001145b6108625760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610160565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108be575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0383166109145760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f757465722061646472657373000000000000000000000000006044820152606401610160565b61091f848484611058565b80156107b3575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b61098e610ddf565b6001600160a01b038116610a0a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610160565b610a1381610e39565b50565b610a1e610ea2565b5f8311610a6d5760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610160565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b031614610aee5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c79205745544820697320616c6c6f7765640000000000000000000000006044820152606401610160565b60985433906001600160a01b0316819003610b1c5782806020019051810190610b179190611989565b935090505b610b316001600160a01b03871682308761119b565b6040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b03871690632e1a7d4d906024015f604051808303815f87803b158015610b89575f80fd5b505af1158015610b9b573d5f803e3d5ffd5b50506040517f000000000000000000000000000000000000000000000000000000000000000092505f9150610bde9083908a9086908b908b908b90602401611aae565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b039091169263ecc704289260048083019391928290030181865afa158015610cb5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd99190611afb565b6099549091506001600160a01b031663b2267a7b610cf7348a611b12565b6097546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610d42916001600160a01b0316908c9088908c90600401611b50565b5f604051808303818588803b158015610d59575f80fd5b505af1158015610d6b573d5f803e3d5ffd5b5050505050836001600160a01b0316896001600160a01b0316846001600160a01b03167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b87604051610dc39493929190611b50565b60405180910390a450505050610dd860018055565b5050505050565b6065546001600160a01b031633146104215760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610160565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600260015403610ef45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610160565b6002600155565b6040516001600160a01b03831660248201526044810182905261040b9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526111ec565b5f8151118015610fdb57505f826001600160a01b03163b115b1561104e576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f90611025908490600401611b88565b5f604051808303815f87803b15801561103c575f80fd5b505af11580156103ca573d5f803e3d5ffd5b5050565b60018055565b6001600160a01b0383166110ae5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610160565b6001600160a01b0381166111045760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610160565b61110c6112d2565b611114611356565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561040b57609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6040516001600160a01b03808516602483015283166044820152606481018290526107b39085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610f40565b5f611240826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113da9092919063ffffffff16565b905080515f14806112605750808060200190518101906112609190611b9a565b61040b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610160565b5f54610100900460ff1661134e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610160565b6104216113f0565b5f54610100900460ff166113d25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610160565b61042161146c565b60606113e884845f856114f1565b949350505050565b5f54610100900460ff166110525760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610160565b5f54610100900460ff166114e85760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610160565b61042133610e39565b6060824710156115695760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610160565b5f80866001600160a01b031685876040516115849190611bb9565b5f6040518083038185875af1925050503d805f81146115be576040519150601f19603f3d011682016040523d82523d5f602084013e6115c3565b606091505b50915091506115d4878383876115df565b979650505050505050565b6060831561164d5782515f03611646576001600160a01b0385163b6116465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610160565b50816113e8565b6113e883838151156116625781518083602001fd5b8060405162461bcd60e51b81526004016101609190611b88565b6001600160a01b0381168114610a13575f80fd5b5f602082840312156116a0575f80fd5b81356116ab8161167c565b9392505050565b5f8083601f8401126116c2575f80fd5b50813567ffffffffffffffff8111156116d9575f80fd5b6020830191508360208285010111156116f0575f80fd5b9250929050565b5f805f805f8060a0878903121561170c575f80fd5b86356117178161167c565b955060208701356117278161167c565b945060408701359350606087013567ffffffffffffffff811115611749575f80fd5b61175589828a016116b2565b979a9699509497949695608090950135949350505050565b5f805f6060848603121561177f575f80fd5b833561178a8161167c565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a0312156117b5575f80fd5b87356117c08161167c565b965060208801356117d08161167c565b955060408801356117e08161167c565b945060608801356117f08161167c565b93506080880135925060a088013567ffffffffffffffff811115611812575f80fd5b61181e8a828b016116b2565b989b979a50959850939692959293505050565b5f805f8060808587031215611844575f80fd5b843561184f8161167c565b9350602085013561185f8161167c565b93969395505050506040820135916060013590565b5f805f60608486031215611886575f80fd5b83356118918161167c565b925060208401356118a18161167c565b915060408401356118b18161167c565b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f602082840312156118f9575f80fd5b81516116ab8161167c565b6001600160a01b038516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611981578181015183820152602001611969565b50505f910152565b5f806040838503121561199a575f80fd5b82516119a58161167c565b602084015190925067ffffffffffffffff808211156119c2575f80fd5b818501915085601f8301126119d5575f80fd5b8151818111156119e7576119e76118bc565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611a2d57611a2d6118bc565b81604052828152886020848701011115611a45575f80fd5b611a56836020830160208801611967565b80955050505050509250929050565b5f8151808452611a7c816020860160208601611967565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f6001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611aef60c0830184611a65565b98975050505050505050565b5f60208284031215611b0b575f80fd5b5051919050565b80820180821115611b4a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b6001600160a01b0385168152836020820152608060408201525f611b776080830185611a65565b905082606083015295945050505050565b602081525f6116ab6020830184611a65565b5f60208284031215611baa575f80fd5b815180151581146116ab575f80fd5b5f8251611bca818460208701611967565b919091019291505056fea164736f6c6343000818000a",
}

// L2WETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2WETHGatewayMetaData.ABI instead.
var L2WETHGatewayABI = L2WETHGatewayMetaData.ABI

// L2WETHGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2WETHGatewayMetaData.Bin instead.
var L2WETHGatewayBin = L2WETHGatewayMetaData.Bin

// DeployL2WETHGateway deploys a new Ethereum contract, binding an instance of L2WETHGateway to it.
func DeployL2WETHGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _l2WETH common.Address, _l1WETH common.Address) (common.Address, *types.Transaction, *L2WETHGateway, error) {
	parsed, err := L2WETHGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2WETHGatewayBin), backend, _l2WETH, _l1WETH)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2WETHGateway{L2WETHGatewayCaller: L2WETHGatewayCaller{contract: contract}, L2WETHGatewayTransactor: L2WETHGatewayTransactor{contract: contract}, L2WETHGatewayFilterer: L2WETHGatewayFilterer{contract: contract}}, nil
}

// L2WETHGateway is an auto generated Go binding around an Ethereum contract.
type L2WETHGateway struct {
	L2WETHGatewayCaller     // Read-only binding to the contract
	L2WETHGatewayTransactor // Write-only binding to the contract
	L2WETHGatewayFilterer   // Log filterer for contract events
}

// L2WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2WETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2WETHGatewaySession struct {
	Contract     *L2WETHGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2WETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2WETHGatewayCallerSession struct {
	Contract *L2WETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2WETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2WETHGatewayTransactorSession struct {
	Contract     *L2WETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2WETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2WETHGatewayRaw struct {
	Contract *L2WETHGateway // Generic contract binding to access the raw methods on
}

// L2WETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2WETHGatewayCallerRaw struct {
	Contract *L2WETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2WETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2WETHGatewayTransactorRaw struct {
	Contract *L2WETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2WETHGateway creates a new instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGateway(address common.Address, backend bind.ContractBackend) (*L2WETHGateway, error) {
	contract, err := bindL2WETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2WETHGateway{L2WETHGatewayCaller: L2WETHGatewayCaller{contract: contract}, L2WETHGatewayTransactor: L2WETHGatewayTransactor{contract: contract}, L2WETHGatewayFilterer: L2WETHGatewayFilterer{contract: contract}}, nil
}

// NewL2WETHGatewayCaller creates a new read-only instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2WETHGatewayCaller, error) {
	contract, err := bindL2WETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayCaller{contract: contract}, nil
}

// NewL2WETHGatewayTransactor creates a new write-only instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2WETHGatewayTransactor, error) {
	contract, err := bindL2WETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayTransactor{contract: contract}, nil
}

// NewL2WETHGatewayFilterer creates a new log filterer instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2WETHGatewayFilterer, error) {
	contract, err := bindL2WETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayFilterer{contract: contract}, nil
}

// bindL2WETHGateway binds a generic wrapper to an already deployed contract.
func bindL2WETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2WETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WETHGateway *L2WETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WETHGateway.Contract.L2WETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WETHGateway *L2WETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.L2WETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WETHGateway *L2WETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.L2WETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WETHGateway *L2WETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WETHGateway *L2WETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WETHGateway *L2WETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.contract.Transact(opts, method, params...)
}

// L1WETH is a free data retrieval call binding the contract method 0x1efd482a.
//
// Solidity: function L1_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) L1WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "L1_WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1WETH is a free data retrieval call binding the contract method 0x1efd482a.
//
// Solidity: function L1_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) L1WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L1WETH(&_L2WETHGateway.CallOpts)
}

// L1WETH is a free data retrieval call binding the contract method 0x1efd482a.
//
// Solidity: function L1_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) L1WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L1WETH(&_L2WETHGateway.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0x19c4d4c6.
//
// Solidity: function L2_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "L2_WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0x19c4d4c6.
//
// Solidity: function L2_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) L2WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L2WETH(&_L2WETHGateway.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0x19c4d4c6.
//
// Solidity: function L2_WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) L2WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L2WETH(&_L2WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Counterpart() (common.Address, error) {
	return _L2WETHGateway.Contract.Counterpart(&_L2WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2WETHGateway.Contract.Counterpart(&_L2WETHGateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "getL1ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL1ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL1ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL2ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL2ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Messenger() (common.Address, error) {
	return _L2WETHGateway.Contract.Messenger(&_L2WETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2WETHGateway.Contract.Messenger(&_L2WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Owner() (common.Address, error) {
	return _L2WETHGateway.Contract.Owner(&_L2WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Owner() (common.Address, error) {
	return _L2WETHGateway.Contract.Owner(&_L2WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Router() (common.Address, error) {
	return _L2WETHGateway.Contract.Router(&_L2WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Router() (common.Address, error) {
	return _L2WETHGateway.Contract.Router(&_L2WETHGateway.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.FinalizeDepositERC20(&_L2WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.FinalizeDepositERC20(&_L2WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Initialize(&_L2WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Initialize(&_L2WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.RenounceOwnership(&_L2WETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.RenounceOwnership(&_L2WETHGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.TransferOwnership(&_L2WETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.TransferOwnership(&_L2WETHGateway.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20(&_L2WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20(&_L2WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC200(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC200(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20AndCall(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20AndCall(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) Receive() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Receive(&_L2WETHGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Receive(&_L2WETHGateway.TransactOpts)
}

// L2WETHGatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2WETHGateway contract.
type L2WETHGatewayFinalizeDepositERC20Iterator struct {
	Event *L2WETHGatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayFinalizeDepositERC20)
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
		it.Event = new(L2WETHGatewayFinalizeDepositERC20)
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
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2WETHGateway contract.
type L2WETHGatewayFinalizeDepositERC20 struct {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WETHGatewayFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayFinalizeDepositERC20Iterator{contract: _L2WETHGateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayFinalizeDepositERC20)
				if err := _L2WETHGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2WETHGatewayFinalizeDepositERC20, error) {
	event := new(L2WETHGatewayFinalizeDepositERC20)
	if err := _L2WETHGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2WETHGateway contract.
type L2WETHGatewayInitializedIterator struct {
	Event *L2WETHGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayInitialized)
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
		it.Event = new(L2WETHGatewayInitialized)
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
func (it *L2WETHGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayInitialized represents a Initialized event raised by the L2WETHGateway contract.
type L2WETHGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2WETHGatewayInitializedIterator, error) {

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayInitializedIterator{contract: _L2WETHGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayInitialized)
				if err := _L2WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseInitialized(log types.Log) (*L2WETHGatewayInitialized, error) {
	event := new(L2WETHGatewayInitialized)
	if err := _L2WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2WETHGateway contract.
type L2WETHGatewayOwnershipTransferredIterator struct {
	Event *L2WETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayOwnershipTransferred)
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
		it.Event = new(L2WETHGatewayOwnershipTransferred)
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
func (it *L2WETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2WETHGateway contract.
type L2WETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2WETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayOwnershipTransferredIterator{contract: _L2WETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayOwnershipTransferred)
				if err := _L2WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2WETHGatewayOwnershipTransferred, error) {
	event := new(L2WETHGatewayOwnershipTransferred)
	if err := _L2WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2WETHGateway contract.
type L2WETHGatewayWithdrawERC20Iterator struct {
	Event *L2WETHGatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayWithdrawERC20)
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
		it.Event = new(L2WETHGatewayWithdrawERC20)
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
func (it *L2WETHGatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2WETHGateway contract.
type L2WETHGatewayWithdrawERC20 struct {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WETHGatewayWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayWithdrawERC20Iterator{contract: _L2WETHGateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c4.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data, uint256 nonce)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayWithdrawERC20)
				if err := _L2WETHGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2WETHGatewayWithdrawERC20, error) {
	event := new(L2WETHGatewayWithdrawERC20)
	if err := _L2WETHGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
