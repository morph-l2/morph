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

// MorphMintableERC20FactoryMetaData contains all meta data concerning the MorphMintableERC20Factory contract.
var MorphMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createMorphMintableERC20\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createStandardL2Token\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MorphMintableERC20Created\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StandardL2TokenCreated\",\"inputs\":[{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161223438038061223483398101604081905261003091610050565b6001608081905260a052600060c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e0516121756100bf6000396000818160ce015261019e0152600061030d015260006102e2015260006102b701526121756000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c80631fa79434146200005757806354fd4d501462000098578063896f93d114620000b1578063ee9a31a214620000c8575b600080fd5b6200006e6200006836600462000610565b620000f0565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b620000a2620002af565b6040516200008f919062000719565b6200006e620000c236600462000610565b6200035a565b6200006e7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff84166200019a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f4d6f7270684d696e7461626c654552433230466163746f72793a206d7573742060448201527f70726f766964652072656d6f746520746f6b656e206164647265737300000000606482015260840160405180910390fd5b60007f0000000000000000000000000000000000000000000000000000000000000000858585604051620001ce9062000520565b620001dd949392919062000735565b604051809103906000f080158015620001fa573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a360405133815273ffffffffffffffffffffffffffffffffffffffff80871691908316907f10990cc0cc2bb901dd0cf179e544af30985f0fa4307de1da2be8b63a409235af9060200160405180910390a3949350505050565b6060620002dc7f000000000000000000000000000000000000000000000000000000000000000062000371565b620003077f000000000000000000000000000000000000000000000000000000000000000062000371565b620003327f000000000000000000000000000000000000000000000000000000000000000062000371565b60405160200162000346939291906200078f565b604051602081830303815290604052905090565b600062000369848484620000f0565b949350505050565b60606000620003808362000436565b600101905060008167ffffffffffffffff811115620003a357620003a36200052e565b6040519080825280601f01601f191660200182016040528015620003ce576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084620003d857509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831062000480577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310620004ad576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310620004cc57662386f26fc10000830492506010015b6305f5e1008310620004e5576305f5e100830492506008015b6127108310620004fa57612710830492506004015b606483106200050d576064830492506002015b600a83106200051a576001015b92915050565b61195d806200080c83390190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200056f57600080fd5b813567ffffffffffffffff808211156200058d576200058d6200052e565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620005d657620005d66200052e565b81604052838152866020858801011115620005f057600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200062657600080fd5b833573ffffffffffffffffffffffffffffffffffffffff811681146200064b57600080fd5b9250602084013567ffffffffffffffff808211156200066957600080fd5b62000677878388016200055d565b935060408601359150808211156200068e57600080fd5b506200069d868287016200055d565b9150509250925092565b60005b83811015620006c4578181015183820152602001620006aa565b50506000910152565b60008151808452620006e7816020860160208601620006a7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006200072e6020830184620006cd565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152620007706080830185620006cd565b8281036060840152620007848185620006cd565b979650505050505050565b60008451620007a3818460208901620006a7565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551620007e1816001850160208a01620006a7565b60019201918201528351620007fe816002840160208801620006a7565b016002019594505050505056fe6101206040523480156200001257600080fd5b506040516200195d3803806200195d833981016040819052620000359162000165565b6001600080848460036200004a838262000284565b50600462000059828262000284565b50505060809290925260a05260c05250506001600160a01b0390811660e052166101005262000350565b80516001600160a01b03811681146200009b57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c857600080fd5b81516001600160401b0380821115620000e557620000e5620000a0565b604051601f8301601f19908116603f01168101908282118183101715620001105762000110620000a0565b816040528381526020925086838588010111156200012d57600080fd5b600091505b8382101562000151578582018301518183018401529082019062000132565b600093810190920192909252949350505050565b600080600080608085870312156200017c57600080fd5b620001878562000083565b9350620001976020860162000083565b60408601519093506001600160401b0380821115620001b557600080fd5b620001c388838901620000b6565b93506060870151915080821115620001da57600080fd5b50620001e987828801620000b6565b91505092959194509250565b600181811c908216806200020a57607f821691505b6020821081036200022b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027f57600081815260208120601f850160051c810160208610156200025a5750805b601f850160051c820191505b818110156200027b5782815560010162000266565b5050505b505050565b81516001600160401b03811115620002a057620002a0620000a0565b620002b881620002b18454620001f5565b8462000231565b602080601f831160018114620002f05760008415620002d75750858301515b600019600386901b1c1916600185901b1785556200027b565b600085815260208120601f198616915b82811015620003215788860151825594840194600190910190840162000300565b5085821015620003405787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516115ac620003b1600039600081816102f50152818161038a015281816105d101526107ab0152600081816101a9015261031b0152600061073a01526000610711015260006106e801526115ac6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e1461033f578063e78cea92146102f3578063ee9a31a21461038557600080fd5b8063ae1f6aaf146102f3578063c01e1bd614610319578063d6c0b2c41461031957600080fd5b80639dc29fac116100bd5780639dc29fac146102ba578063a457c2d7146102cd578063a9059cbb146102e057600080fd5b806370a082311461027c57806395d89b41146102b257600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a3660046112d2565b6103ac565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f861049d565b60405161019b919061133f565b61018f6102133660046113b9565b61052f565b6002545b60405190815260200161019b565b61018f6102383660046113e3565b610549565b6040516012815260200161019b565b61018f61025a3660046113b9565b61056d565b61027261026d3660046113b9565b6105b9565b005b6101f86106e1565b61021c61028a36600461141f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610784565b6102726102c83660046113b9565b610793565b61018f6102db3660046113b9565b6108aa565b61018f6102ee3660046113b9565b61097b565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61034d36600461143a565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061046557507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061049457507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ac9061146d565b80601f01602080910402602001604051908101604052809291908181526020018280546104d89061146d565b80156105255780601f106104fa57610100808354040283529160200191610525565b820191906000526020600020905b81548152906001019060200180831161050857829003601f168201915b5050505050905090565b60003361053d818585610989565b60019150505b92915050565b600033610557858285610b3d565b610562858585610c14565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061053d90829086906105b49087906114c0565b610989565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4d6f7270684d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e00000000000000000000000000000060648201526084015b60405180910390fd5b61068d8282610e83565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106d591815260200190565b60405180910390a25050565b606061070c7f0000000000000000000000000000000000000000000000000000000000000000610f76565b6107357f0000000000000000000000000000000000000000000000000000000000000000610f76565b61075e7f0000000000000000000000000000000000000000000000000000000000000000610f76565b604051602001610770939291906114fa565b604051602081830303815290604052905090565b6060600480546104ac9061146d565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610858576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4d6f7270684d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e000000000000000000000000000000606482015260840161067a565b6108628282611034565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106d591815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561096e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161067a565b6105628286868403610989565b60003361053d818585610c14565b73ffffffffffffffffffffffffffffffffffffffff8316610a2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff8216610ace576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c0e5781811015610c01576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161067a565b610c0e8484848403610989565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff8216610d5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610c0e565b73ffffffffffffffffffffffffffffffffffffffff8216610f00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161067a565b8060026000828254610f1291906114c0565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b60606000610f83836111f0565b600101905060008167ffffffffffffffff811115610fa357610fa3611570565b6040519080825280601f01601f191660200182016040528015610fcd576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610fd757509392505050565b73ffffffffffffffffffffffffffffffffffffffff82166110d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561118d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161067a565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610b30565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611239577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611265576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061128357662386f26fc10000830492506010015b6305f5e100831061129b576305f5e100830492506008015b61271083106112af57612710830492506004015b606483106112c1576064830492506002015b600a83106105435760010192915050565b6000602082840312156112e457600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461131457600080fd5b9392505050565b60005b8381101561133657818101518382015260200161131e565b50506000910152565b602081526000825180602084015261135e81604085016020870161131b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113b457600080fd5b919050565b600080604083850312156113cc57600080fd5b6113d583611390565b946020939093013593505050565b6000806000606084860312156113f857600080fd5b61140184611390565b925061140f60208501611390565b9150604084013590509250925092565b60006020828403121561143157600080fd5b61131482611390565b6000806040838503121561144d57600080fd5b61145683611390565b915061146460208401611390565b90509250929050565b600181811c9082168061148157607f821691505b6020821081036114ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115610543577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000845161150c81846020890161131b565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611548816001850160208a0161131b565b6001920191820152835161156381600284016020880161131b565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000810000aa164736f6c6343000810000a",
}

// MorphMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphMintableERC20FactoryMetaData.ABI instead.
var MorphMintableERC20FactoryABI = MorphMintableERC20FactoryMetaData.ABI

// MorphMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MorphMintableERC20FactoryMetaData.Bin instead.
var MorphMintableERC20FactoryBin = MorphMintableERC20FactoryMetaData.Bin

// DeployMorphMintableERC20Factory deploys a new Ethereum contract, binding an instance of MorphMintableERC20Factory to it.
func DeployMorphMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *MorphMintableERC20Factory, error) {
	parsed, err := MorphMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MorphMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MorphMintableERC20Factory{MorphMintableERC20FactoryCaller: MorphMintableERC20FactoryCaller{contract: contract}, MorphMintableERC20FactoryTransactor: MorphMintableERC20FactoryTransactor{contract: contract}, MorphMintableERC20FactoryFilterer: MorphMintableERC20FactoryFilterer{contract: contract}}, nil
}

// MorphMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type MorphMintableERC20Factory struct {
	MorphMintableERC20FactoryCaller     // Read-only binding to the contract
	MorphMintableERC20FactoryTransactor // Write-only binding to the contract
	MorphMintableERC20FactoryFilterer   // Log filterer for contract events
}

// MorphMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphMintableERC20FactorySession struct {
	Contract     *MorphMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MorphMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphMintableERC20FactoryCallerSession struct {
	Contract *MorphMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// MorphMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphMintableERC20FactoryTransactorSession struct {
	Contract     *MorphMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// MorphMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphMintableERC20FactoryRaw struct {
	Contract *MorphMintableERC20Factory // Generic contract binding to access the raw methods on
}

// MorphMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphMintableERC20FactoryCallerRaw struct {
	Contract *MorphMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MorphMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphMintableERC20FactoryTransactorRaw struct {
	Contract *MorphMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphMintableERC20Factory creates a new instance of MorphMintableERC20Factory, bound to a specific deployed contract.
func NewMorphMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*MorphMintableERC20Factory, error) {
	contract, err := bindMorphMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20Factory{MorphMintableERC20FactoryCaller: MorphMintableERC20FactoryCaller{contract: contract}, MorphMintableERC20FactoryTransactor: MorphMintableERC20FactoryTransactor{contract: contract}, MorphMintableERC20FactoryFilterer: MorphMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewMorphMintableERC20FactoryCaller creates a new read-only instance of MorphMintableERC20Factory, bound to a specific deployed contract.
func NewMorphMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*MorphMintableERC20FactoryCaller, error) {
	contract, err := bindMorphMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20FactoryCaller{contract: contract}, nil
}

// NewMorphMintableERC20FactoryTransactor creates a new write-only instance of MorphMintableERC20Factory, bound to a specific deployed contract.
func NewMorphMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphMintableERC20FactoryTransactor, error) {
	contract, err := bindMorphMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewMorphMintableERC20FactoryFilterer creates a new log filterer instance of MorphMintableERC20Factory, bound to a specific deployed contract.
func NewMorphMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphMintableERC20FactoryFilterer, error) {
	contract, err := bindMorphMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindMorphMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindMorphMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MorphMintableERC20FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphMintableERC20Factory.Contract.MorphMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.MorphMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.MorphMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _MorphMintableERC20Factory.Contract.BRIDGE(&_MorphMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _MorphMintableERC20Factory.Contract.BRIDGE(&_MorphMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MorphMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MorphMintableERC20Factory *MorphMintableERC20FactorySession) Version() (string, error) {
	return _MorphMintableERC20Factory.Contract.Version(&_MorphMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryCallerSession) Version() (string, error) {
	return _MorphMintableERC20Factory.Contract.Version(&_MorphMintableERC20Factory.CallOpts)
}

// CreateMorphMintableERC20 is a paid mutator transaction binding the contract method 0x1fa79434.
//
// Solidity: function createMorphMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactor) CreateMorphMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.contract.Transact(opts, "createMorphMintableERC20", _remoteToken, _name, _symbol)
}

// CreateMorphMintableERC20 is a paid mutator transaction binding the contract method 0x1fa79434.
//
// Solidity: function createMorphMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactorySession) CreateMorphMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.CreateMorphMintableERC20(&_MorphMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateMorphMintableERC20 is a paid mutator transaction binding the contract method 0x1fa79434.
//
// Solidity: function createMorphMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactorSession) CreateMorphMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.CreateMorphMintableERC20(&_MorphMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactor) CreateStandardL2Token(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.contract.Transact(opts, "createStandardL2Token", _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactorySession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.CreateStandardL2Token(&_MorphMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryTransactorSession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MorphMintableERC20Factory.Contract.CreateStandardL2Token(&_MorphMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// MorphMintableERC20FactoryMorphMintableERC20CreatedIterator is returned from FilterMorphMintableERC20Created and is used to iterate over the raw logs and unpacked data for MorphMintableERC20Created events raised by the MorphMintableERC20Factory contract.
type MorphMintableERC20FactoryMorphMintableERC20CreatedIterator struct {
	Event *MorphMintableERC20FactoryMorphMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *MorphMintableERC20FactoryMorphMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphMintableERC20FactoryMorphMintableERC20Created)
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
		it.Event = new(MorphMintableERC20FactoryMorphMintableERC20Created)
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
func (it *MorphMintableERC20FactoryMorphMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphMintableERC20FactoryMorphMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphMintableERC20FactoryMorphMintableERC20Created represents a MorphMintableERC20Created event raised by the MorphMintableERC20Factory contract.
type MorphMintableERC20FactoryMorphMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMorphMintableERC20Created is a free log retrieval operation binding the contract event 0x10990cc0cc2bb901dd0cf179e544af30985f0fa4307de1da2be8b63a409235af.
//
// Solidity: event MorphMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) FilterMorphMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*MorphMintableERC20FactoryMorphMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MorphMintableERC20Factory.contract.FilterLogs(opts, "MorphMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20FactoryMorphMintableERC20CreatedIterator{contract: _MorphMintableERC20Factory.contract, event: "MorphMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchMorphMintableERC20Created is a free log subscription operation binding the contract event 0x10990cc0cc2bb901dd0cf179e544af30985f0fa4307de1da2be8b63a409235af.
//
// Solidity: event MorphMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) WatchMorphMintableERC20Created(opts *bind.WatchOpts, sink chan<- *MorphMintableERC20FactoryMorphMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MorphMintableERC20Factory.contract.WatchLogs(opts, "MorphMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphMintableERC20FactoryMorphMintableERC20Created)
				if err := _MorphMintableERC20Factory.contract.UnpackLog(event, "MorphMintableERC20Created", log); err != nil {
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

// ParseMorphMintableERC20Created is a log parse operation binding the contract event 0x10990cc0cc2bb901dd0cf179e544af30985f0fa4307de1da2be8b63a409235af.
//
// Solidity: event MorphMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) ParseMorphMintableERC20Created(log types.Log) (*MorphMintableERC20FactoryMorphMintableERC20Created, error) {
	event := new(MorphMintableERC20FactoryMorphMintableERC20Created)
	if err := _MorphMintableERC20Factory.contract.UnpackLog(event, "MorphMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphMintableERC20FactoryStandardL2TokenCreatedIterator is returned from FilterStandardL2TokenCreated and is used to iterate over the raw logs and unpacked data for StandardL2TokenCreated events raised by the MorphMintableERC20Factory contract.
type MorphMintableERC20FactoryStandardL2TokenCreatedIterator struct {
	Event *MorphMintableERC20FactoryStandardL2TokenCreated // Event containing the contract specifics and raw log

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
func (it *MorphMintableERC20FactoryStandardL2TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphMintableERC20FactoryStandardL2TokenCreated)
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
		it.Event = new(MorphMintableERC20FactoryStandardL2TokenCreated)
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
func (it *MorphMintableERC20FactoryStandardL2TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphMintableERC20FactoryStandardL2TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphMintableERC20FactoryStandardL2TokenCreated represents a StandardL2TokenCreated event raised by the MorphMintableERC20Factory contract.
type MorphMintableERC20FactoryStandardL2TokenCreated struct {
	RemoteToken common.Address
	LocalToken  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStandardL2TokenCreated is a free log retrieval operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) FilterStandardL2TokenCreated(opts *bind.FilterOpts, remoteToken []common.Address, localToken []common.Address) (*MorphMintableERC20FactoryStandardL2TokenCreatedIterator, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _MorphMintableERC20Factory.contract.FilterLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &MorphMintableERC20FactoryStandardL2TokenCreatedIterator{contract: _MorphMintableERC20Factory.contract, event: "StandardL2TokenCreated", logs: logs, sub: sub}, nil
}

// WatchStandardL2TokenCreated is a free log subscription operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) WatchStandardL2TokenCreated(opts *bind.WatchOpts, sink chan<- *MorphMintableERC20FactoryStandardL2TokenCreated, remoteToken []common.Address, localToken []common.Address) (event.Subscription, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _MorphMintableERC20Factory.contract.WatchLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphMintableERC20FactoryStandardL2TokenCreated)
				if err := _MorphMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
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

// ParseStandardL2TokenCreated is a log parse operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MorphMintableERC20Factory *MorphMintableERC20FactoryFilterer) ParseStandardL2TokenCreated(log types.Log) (*MorphMintableERC20FactoryStandardL2TokenCreated, error) {
	event := new(MorphMintableERC20FactoryStandardL2TokenCreated)
	if err := _MorphMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
