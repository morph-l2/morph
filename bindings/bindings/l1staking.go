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
	_ = abi.ConvertType
)

// L1StakingMetaData contains all meta data concerning the L1Staking contract.
var L1StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"ParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"Withdrawed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_STAKING\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_LOCK_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimSlashRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"add\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signedSequencers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b50604051620027703803806200277083398101604081905262000033916200005d565b6001600160a01b031660805273530000000000000000000000000000000000001260a0526200008c565b5f602082840312156200006e575f80fd5b81516001600160a01b038116811462000085575f80fd5b9392505050565b60805160a05161269f620000d15f395f81816102cb015281816119e50152611b7901525f8181610214015281816103a6015281816119b50152611b49015261269f5ff3fe60806040526004361061018e575f3560e01c80638da5cb5b116100dc578063c0af545b11610087578063d6be695a11610062578063d6be695a14610496578063f2fde38b146104ab578063f6ebdcde146104ca578063fc72b1ed146104df575f80fd5b8063c0af545b14610434578063c7cd469a14610458578063cde4cd1114610477575f80fd5b80639b19251a116100b75780639b19251a146103c8578063a3066aab146103f6578063a574918714610415575f80fd5b80638da5cb5b1461034a5780639168ae7214610367578063927ede2d14610395575f80fd5b8063715018a61161013c57806386489ba91161011757806386489ba9146102ed5780638770d7071461030c5780638b8c24c11461032b575f80fd5b8063715018a61461027b5780637a9262a21461028f578063831cfb58146102ba575f80fd5b80633cb747bf1161016c5780633cb747bf146102065780633ccfd60b146102385780636f1e85331461024c575f80fd5b806302df7ff71461019257806320b651cd146101ba578063348e50c6146101cf575b5f80fd5b34801561019d575f80fd5b506101a760985481565b6040519081526020015b60405180910390f35b6101cd6101c8366004611f0a565b6104f4565b005b3480156101da575f80fd5b506101ee6101e9366004611f5d565b610a51565b6040516001600160a01b0390911681526020016101b1565b348015610211575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000006101ee565b348015610243575f80fd5b506101cd610a79565b348015610257575f80fd5b5061026b610266366004611f74565b610cd7565b60405190151581526020016101b1565b348015610286575f80fd5b506101cd610d2e565b34801561029a575f80fd5b506101a76102a9366004611f74565b60a06020525f908152604090205481565b3480156102c5575f80fd5b506101ee7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102f8575f80fd5b506101cd610307366004611f94565b610d41565b348015610317575f80fd5b506097546101ee906001600160a01b031681565b348015610336575f80fd5b506101a7610345366004612069565b611142565b348015610355575f80fd5b506033546001600160a01b03166101ee565b348015610372575f80fd5b50610386610381366004611f74565b61152c565b6040516101b193929190612104565b3480156103a0575f80fd5b506101ee7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103d3575f80fd5b5061026b6103e2366004611f74565b609d6020525f908152604090205460ff1681565b348015610401575f80fd5b506101cd610410366004611f74565b6115de565b348015610420575f80fd5b506101cd61042f366004611f5d565b611755565b34801561043f575f80fd5b5061026b61044e366004612134565b6001949350505050565b348015610463575f80fd5b506101cd610472366004612208565b6117e7565b348015610482575f80fd5b506101cd610491366004611f74565b6118fc565b3480156104a1575f80fd5b506101a7609b5481565b3480156104b6575f80fd5b506101cd6104c5366004611f74565b611926565b3480156104d5575f80fd5b506101a760995481565b3480156104ea575f80fd5b506101a7609a5481565b6001600160a01b0383165f908152609d6020526040902054839060ff166105625760405162461bcd60e51b815260206004820152601060248201527f6e6f7420696e2077686974656c6973740000000000000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b0384166105b85760405162461bcd60e51b815260206004820152600f60248201527f696e76616c6964206164647265737300000000000000000000000000000000006044820152606401610559565b6001600160a01b038481165f908152609f6020526040902054161561061f5760405162461bcd60e51b815260206004820152601260248201527f616c7265616479207265676973746572656400000000000000000000000000006044820152606401610559565b5f83900361066f5760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642074656e6465726d696e74207075626b6579000000000000006044820152606401610559565b8151610100146106c15760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420626c73207075626b657900000000000000000000000000006044820152606401610559565b60985434146107125760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964207374616b696e672076616c756500000000000000000000006044820152606401610559565b5f5b609e548110156108595783609f5f609e84815481106107355761073561226f565b5f9182526020808320909101546001600160a01b03168352820192909252604001902060010154036107a95760405162461bcd60e51b815260206004820152601860248201527f746d4b657920616c7265616479207265676973746572656400000000000000006044820152606401610559565b8280519060200120609f5f609e84815481106107c7576107c761226f565b5f9182526020808320909101546001600160a01b03168352820192909252604090810190912090516107fc91600201906122ed565b6040518091039020036108515760405162461bcd60e51b815260206004820152601960248201527f626c734b657920616c72656164792072656769737465726564000000000000006044820152606401610559565b600101610714565b50604080516060810182526001600160a01b0386811680835260208084018881528486018881525f938452609f90925294909120835181547fffffffffffffffffffffffff000000000000000000000000000000000000000016931692909217825592516001820155915190919060028201906108d690826123c1565b5050609e80546001810182555f919091527fcfe2a20ff701a1f3e14f63bd70d6c6bc6fba8172ec6d5a505cdab3927c0a9de60180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716179055506040517fb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e1999061096f90869086908690612104565b60405180910390a16001600160a01b038085165f908152609f6020908152604091829020825160608101845281549094168452600181015491840191909152600281018054610a4b9493840191906109c69061229c565b80601f01602080910402602001604051908101604052809291908181526020018280546109f29061229c565b8015610a3d5780601f10610a1457610100808354040283529160200191610a3d565b820191905f5260205f20905b815481529060010190602001808311610a2057829003601f168201915b5050505050815250506119b3565b50505050565b609e8181548110610a60575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f80610a8433611ae5565b9150915081610ad55760405162461bcd60e51b815260206004820152600b60248201527f6f6e6c79207374616b65720000000000000000000000000000000000000000006044820152606401610559565b335f90815260a0602052604090205415610b315760405162461bcd60e51b815260206004820152600b60248201527f7769746864726177696e670000000000000000000000000000000000000000006044820152606401610559565b609954610b3e9043612506565b335f90815260a06020526040902055609e8054610b5d9060019061251f565b81548110610b6d57610b6d61226f565b5f91825260209091200154609e80546001600160a01b039092169183908110610b9857610b9861226f565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550609e805480610bd457610bd4612532565b5f82815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559092019092553380835260a08252604092839020548351918252918101919091527f6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe910160405180910390a16040805160018082528183019092525f916020808301908036833701905050905033815f81518110610ca957610ca961226f565b60200260200101906001600160a01b031690816001600160a01b031681525050610cd281611b47565b505050565b5f805b609e54811015610d2657609e8181548110610cf757610cf761226f565b5f918252602090912001546001600160a01b0390811690841603610d1e5750600192915050565b600101610cda565b505f92915050565b610d36611bb2565b610d3f5f611c0c565b565b5f54610100900460ff1615808015610d5f57505f54600160ff909116105b80610d785750303b158015610d7857505f5460ff166001145b610dea5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610559565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610e46575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b038716610e9c5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642061646d696e000000000000000000000000000000000000006044820152606401610559565b6001600160a01b038616610ef25760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e74726163740000000000000000006044820152606401610559565b5f8411610f675760405162461bcd60e51b815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f30000000000000000000000000000000000000000000000000000000000000006064820152608401610559565b5f8311610fdc5760405162461bcd60e51b815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f30000000000000000000000000000000000000000000000000000000000000006064820152608401610559565b5f821161102b5760405162461bcd60e51b815260206004820152601d60248201527f676173206c696d6974206d7573742067726561746572207468616e20300000006044820152606401610559565b5f8511801561103b575060648511155b6110875760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964207265776172642070657263656e74616765000000000000006044820152606401610559565b609780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038816179055609a85905560988490556099839055609b8290556110d787611c0c565b8015611139575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6097545f906001600160a01b0316331461119e5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610559565b6111a6611c75565b5f805b8351811015611487575f60a05f8684815181106111c8576111c861226f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205411156112bf5760a05f85838151811061120b5761120b61226f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9055609f5f85838151811061124a5761124a61226f565b6020908102919091018101516001600160a01b031682528101919091526040015f90812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906112a96002830182611da4565b50506098546112b89083612506565b9150611422565b5f806112e38684815181106112d6576112d661226f565b6020026020010151611ae5565b91509150811561141f57609e80546112fd9060019061251f565b8154811061130d5761130d61226f565b5f91825260209091200154609e80546001600160a01b0390921691839081106113385761133861226f565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550609e80548061137457611374612532565b600190038181905f5260205f20015f6101000a8154906001600160a01b0302191690559055609f5f8785815181106113ae576113ae61226f565b6020908102919091018101516001600160a01b031682528101919091526040015f90812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155600181018290559061140d6002830182611da4565b505060985461141c9085612506565b93505b50505b609d5f8583815181106114375761143761226f565b6020908102919091018101516001600160a01b031682528101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556001016111a9565b505f6064609a5483611499919061255f565b6114a39190612576565b90506114af818361251f565b609c5f8282546114bf9190612506565b90915550506097546114da906001600160a01b031682611cce565b7f654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a8460405161150991906125ae565b60405180910390a161151a84611b47565b9150506115276001606555565b919050565b609f6020525f90815260409020805460018201546002830180546001600160a01b0390931693919261155d9061229c565b80601f01602080910402602001604051908101604052809291908181526020018280546115899061229c565b80156115d45780601f106115ab576101008083540402835291602001916115d4565b820191905f5260205f20905b8154815290600101906020018083116115b757829003601f168201915b5050505050905083565b6115e6611c75565b335f90815260a060205260409020546116415760405162461bcd60e51b815260206004820152601460248201527f7769746864726177616c206e6f742065786973740000000000000000000000006044820152606401610559565b335f90815260a06020526040902054431161169e5760405162461bcd60e51b815260206004820152601160248201527f7769746864726177616c206c6f636b65640000000000000000000000000000006044820152606401610559565b335f908152609f6020526040812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906116e76002830182611da4565b5050335f81815260a0602090815260408083209290925581519283526001600160a01b038416908301527f89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b910160405180910390a161174881609854611cce565b6117526001606555565b50565b61175d611bb2565b5f81116117ac5760405162461bcd60e51b815260206004820152601d60248201527f676173206c696d6974206d7573742067726561746572207468616e20300000006044820152606401610559565b609b8190556040518181527fa96b260c11da5ffa5f74f6cd6dcb582ef40c552985b8622dd901e63ecee02b3b9060200160405180910390a150565b6117ef611bb2565b5f5b83811015611872576001609d5f8787858181106118105761181061226f565b90506020020160208101906118259190611f74565b6001600160a01b0316815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016117f1565b505f5b818110156118f5575f609d5f8585858181106118935761189361226f565b90506020020160208101906118a89190611f74565b6001600160a01b0316815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101611875565b5050505050565b611904611bb2565b61190c611c75565b61191881609c54611cce565b5f609c556117526001606555565b61192e611bb2565b6001600160a01b0381166119aa5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610559565b61175281611c0c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b2267a7b347f00000000000000000000000000000000000000000000000000000000000000005f636d454d5160e01b86604051602401611a1e91906125fa565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000948516179052609b54905160e088901b9093168352611ab4949392600401612631565b5f604051808303818588803b158015611acb575f80fd5b505af1158015611add573d5f803e3d5ffd5b505050505050565b5f805f5b609e54811015611b3c57836001600160a01b0316609e8281548110611b1057611b1061226f565b5f918252602090912001546001600160a01b031603611b3457600194909350915050565b600101611ae9565b505f93849350915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b2267a7b347f00000000000000000000000000000000000000000000000000000000000000005f630be67fcc60e01b86604051602401611a1e91906125ae565b6033546001600160a01b03163314610d3f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610559565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600260655403611cc75760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610559565b6002606555565b8015611da0575f826001600160a01b031682604051611d10907f3078000000000000000000000000000000000000000000000000000000000000815260020190565b5f6040518083038185875af1925050503d805f8114611d4a576040519150601f19603f3d011682016040523d82523d5f602084013e611d4f565b606091505b5050905080610cd25760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610559565b5050565b508054611db09061229c565b5f825580601f10611dbf575050565b601f0160209004905f5260205f209081019061175291905b80821115611dea575f8155600101611dd7565b5090565b80356001600160a01b0381168114611527575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611e7857611e78611e04565b604052919050565b5f82601f830112611e8f575f80fd5b813567ffffffffffffffff811115611ea957611ea9611e04565b611eda60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611e31565b818152846020838601011115611eee575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f60608486031215611f1c575f80fd5b611f2584611dee565b925060208401359150604084013567ffffffffffffffff811115611f47575f80fd5b611f5386828701611e80565b9150509250925092565b5f60208284031215611f6d575f80fd5b5035919050565b5f60208284031215611f84575f80fd5b611f8d82611dee565b9392505050565b5f805f805f8060c08789031215611fa9575f80fd5b611fb287611dee565b9550611fc060208801611dee565b95989597505050506040840135936060810135936080820135935060a0909101359150565b5f82601f830112611ff4575f80fd5b8135602067ffffffffffffffff82111561201057612010611e04565b8160051b61201f828201611e31565b9283528481018201928281019087851115612038575f80fd5b83870192505b8483101561205e5761204f83611dee565b8252918301919083019061203e565b979650505050505050565b5f60208284031215612079575f80fd5b813567ffffffffffffffff81111561208f575f80fd5b61209b84828501611fe5565b949350505050565b5f81518084525f5b818110156120c7576020818501810151868301820152016120ab565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6001600160a01b0384168152826020820152606060408201525f61212b60608301846120a3565b95945050505050565b5f805f8060808587031215612147575f80fd5b843567ffffffffffffffff8082111561215e575f80fd5b61216a88838901611fe5565b9550602087013591508082111561217f575f80fd5b61218b88838901611fe5565b94506040870135935060608701359150808211156121a7575f80fd5b506121b487828801611e80565b91505092959194509250565b5f8083601f8401126121d0575f80fd5b50813567ffffffffffffffff8111156121e7575f80fd5b6020830191508360208260051b8501011115612201575f80fd5b9250929050565b5f805f806040858703121561221b575f80fd5b843567ffffffffffffffff80821115612232575f80fd5b61223e888389016121c0565b90965094506020870135915080821115612256575f80fd5b50612263878288016121c0565b95989497509550505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c908216806122b057607f821691505b6020821081036122e7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f8083546122fa8161229c565b60018281168015612312576001811461234557612371565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084168752821515830287019450612371565b875f526020805f205f5b858110156123685781548a82015290840190820161234f565b50505082870194505b50929695505050505050565b601f821115610cd257805f5260205f20601f840160051c810160208510156123a25750805b601f840160051c820191505b818110156118f5575f81556001016123ae565b815167ffffffffffffffff8111156123db576123db611e04565b6123ef816123e9845461229c565b8461237d565b602080601f831160018114612441575f841561240b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611add565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561248d5788860151825594840194600190910190840161246e565b50858210156124c957878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115612519576125196124d9565b92915050565b81810381811115612519576125196124d9565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b8082028115828204841417612519576125196124d9565b5f826125a9577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b602080825282518282018190525f9190848201906040850190845b818110156125ee5783516001600160a01b0316835292840192918401916001016125c9565b50909695505050505050565b602081526001600160a01b038251166020820152602082015160408201525f604083015160608084015261209b60808401826120a3565b6001600160a01b0385168152836020820152608060408201525f61265860808301856120a3565b90508260608301529594505050505056fea2646970667358221220ab418431f980d4088a9ea144af58e90ee19e771634adf2d024262a7f307ff3a564736f6c63430008180033",
}

// L1StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StakingMetaData.ABI instead.
var L1StakingABI = L1StakingMetaData.ABI

// L1StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1StakingMetaData.Bin instead.
var L1StakingBin = L1StakingMetaData.Bin

// DeployL1Staking deploys a new Ethereum contract, binding an instance of L1Staking to it.
func DeployL1Staking(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address) (common.Address, *types.Transaction, *L1Staking, error) {
	parsed, err := L1StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1StakingBin), backend, _messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1Staking{L1StakingCaller: L1StakingCaller{contract: contract}, L1StakingTransactor: L1StakingTransactor{contract: contract}, L1StakingFilterer: L1StakingFilterer{contract: contract}}, nil
}

// L1Staking is an auto generated Go binding around an Ethereum contract.
type L1Staking struct {
	L1StakingCaller     // Read-only binding to the contract
	L1StakingTransactor // Write-only binding to the contract
	L1StakingFilterer   // Log filterer for contract events
}

// L1StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StakingSession struct {
	Contract     *L1Staking        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StakingCallerSession struct {
	Contract *L1StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// L1StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StakingTransactorSession struct {
	Contract     *L1StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L1StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StakingRaw struct {
	Contract *L1Staking // Generic contract binding to access the raw methods on
}

// L1StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StakingCallerRaw struct {
	Contract *L1StakingCaller // Generic read-only contract binding to access the raw methods on
}

// L1StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StakingTransactorRaw struct {
	Contract *L1StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1Staking creates a new instance of L1Staking, bound to a specific deployed contract.
func NewL1Staking(address common.Address, backend bind.ContractBackend) (*L1Staking, error) {
	contract, err := bindL1Staking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1Staking{L1StakingCaller: L1StakingCaller{contract: contract}, L1StakingTransactor: L1StakingTransactor{contract: contract}, L1StakingFilterer: L1StakingFilterer{contract: contract}}, nil
}

// NewL1StakingCaller creates a new read-only instance of L1Staking, bound to a specific deployed contract.
func NewL1StakingCaller(address common.Address, caller bind.ContractCaller) (*L1StakingCaller, error) {
	contract, err := bindL1Staking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StakingCaller{contract: contract}, nil
}

// NewL1StakingTransactor creates a new write-only instance of L1Staking, bound to a specific deployed contract.
func NewL1StakingTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StakingTransactor, error) {
	contract, err := bindL1Staking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StakingTransactor{contract: contract}, nil
}

// NewL1StakingFilterer creates a new log filterer instance of L1Staking, bound to a specific deployed contract.
func NewL1StakingFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StakingFilterer, error) {
	contract, err := bindL1Staking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StakingFilterer{contract: contract}, nil
}

// bindL1Staking binds a generic wrapper to an already deployed contract.
func bindL1Staking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Staking *L1StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Staking.Contract.L1StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Staking *L1StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Staking.Contract.L1StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Staking *L1StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Staking.Contract.L1StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1Staking *L1StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1Staking *L1StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1Staking *L1StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1Staking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTGASLIMIT is a free data retrieval call binding the contract method 0xd6be695a.
//
// Solidity: function DEFAULT_GAS_LIMIT() view returns(uint256)
func (_L1Staking *L1StakingCaller) DEFAULTGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "DEFAULT_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTGASLIMIT is a free data retrieval call binding the contract method 0xd6be695a.
//
// Solidity: function DEFAULT_GAS_LIMIT() view returns(uint256)
func (_L1Staking *L1StakingSession) DEFAULTGASLIMIT() (*big.Int, error) {
	return _L1Staking.Contract.DEFAULTGASLIMIT(&_L1Staking.CallOpts)
}

// DEFAULTGASLIMIT is a free data retrieval call binding the contract method 0xd6be695a.
//
// Solidity: function DEFAULT_GAS_LIMIT() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) DEFAULTGASLIMIT() (*big.Int, error) {
	return _L1Staking.Contract.DEFAULTGASLIMIT(&_L1Staking.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Staking *L1StakingCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Staking *L1StakingSession) MESSENGER() (common.Address, error) {
	return _L1Staking.Contract.MESSENGER(&_L1Staking.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L1Staking *L1StakingCallerSession) MESSENGER() (common.Address, error) {
	return _L1Staking.Contract.MESSENGER(&_L1Staking.CallOpts)
}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L1Staking *L1StakingCaller) OTHERSTAKING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "OTHER_STAKING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L1Staking *L1StakingSession) OTHERSTAKING() (common.Address, error) {
	return _L1Staking.Contract.OTHERSTAKING(&_L1Staking.CallOpts)
}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L1Staking *L1StakingCallerSession) OTHERSTAKING() (common.Address, error) {
	return _L1Staking.Contract.OTHERSTAKING(&_L1Staking.CallOpts)
}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_L1Staking *L1StakingCaller) REWARDPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "REWARD_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_L1Staking *L1StakingSession) REWARDPERCENTAGE() (*big.Int, error) {
	return _L1Staking.Contract.REWARDPERCENTAGE(&_L1Staking.CallOpts)
}

// REWARDPERCENTAGE is a free data retrieval call binding the contract method 0xfc72b1ed.
//
// Solidity: function REWARD_PERCENTAGE() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) REWARDPERCENTAGE() (*big.Int, error) {
	return _L1Staking.Contract.REWARDPERCENTAGE(&_L1Staking.CallOpts)
}

// ROLLUPCONTRACT is a free data retrieval call binding the contract method 0x8770d707.
//
// Solidity: function ROLLUP_CONTRACT() view returns(address)
func (_L1Staking *L1StakingCaller) ROLLUPCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "ROLLUP_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROLLUPCONTRACT is a free data retrieval call binding the contract method 0x8770d707.
//
// Solidity: function ROLLUP_CONTRACT() view returns(address)
func (_L1Staking *L1StakingSession) ROLLUPCONTRACT() (common.Address, error) {
	return _L1Staking.Contract.ROLLUPCONTRACT(&_L1Staking.CallOpts)
}

// ROLLUPCONTRACT is a free data retrieval call binding the contract method 0x8770d707.
//
// Solidity: function ROLLUP_CONTRACT() view returns(address)
func (_L1Staking *L1StakingCallerSession) ROLLUPCONTRACT() (common.Address, error) {
	return _L1Staking.Contract.ROLLUPCONTRACT(&_L1Staking.CallOpts)
}

// STAKINGVALUE is a free data retrieval call binding the contract method 0x02df7ff7.
//
// Solidity: function STAKING_VALUE() view returns(uint256)
func (_L1Staking *L1StakingCaller) STAKINGVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "STAKING_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKINGVALUE is a free data retrieval call binding the contract method 0x02df7ff7.
//
// Solidity: function STAKING_VALUE() view returns(uint256)
func (_L1Staking *L1StakingSession) STAKINGVALUE() (*big.Int, error) {
	return _L1Staking.Contract.STAKINGVALUE(&_L1Staking.CallOpts)
}

// STAKINGVALUE is a free data retrieval call binding the contract method 0x02df7ff7.
//
// Solidity: function STAKING_VALUE() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) STAKINGVALUE() (*big.Int, error) {
	return _L1Staking.Contract.STAKINGVALUE(&_L1Staking.CallOpts)
}

// WITHDRAWALLOCKBLOCKS is a free data retrieval call binding the contract method 0xf6ebdcde.
//
// Solidity: function WITHDRAWAL_LOCK_BLOCKS() view returns(uint256)
func (_L1Staking *L1StakingCaller) WITHDRAWALLOCKBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "WITHDRAWAL_LOCK_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALLOCKBLOCKS is a free data retrieval call binding the contract method 0xf6ebdcde.
//
// Solidity: function WITHDRAWAL_LOCK_BLOCKS() view returns(uint256)
func (_L1Staking *L1StakingSession) WITHDRAWALLOCKBLOCKS() (*big.Int, error) {
	return _L1Staking.Contract.WITHDRAWALLOCKBLOCKS(&_L1Staking.CallOpts)
}

// WITHDRAWALLOCKBLOCKS is a free data retrieval call binding the contract method 0xf6ebdcde.
//
// Solidity: function WITHDRAWAL_LOCK_BLOCKS() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) WITHDRAWALLOCKBLOCKS() (*big.Int, error) {
	return _L1Staking.Contract.WITHDRAWALLOCKBLOCKS(&_L1Staking.CallOpts)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_L1Staking *L1StakingCaller) IsStaker(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "isStaker", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_L1Staking *L1StakingSession) IsStaker(addr common.Address) (bool, error) {
	return _L1Staking.Contract.IsStaker(&_L1Staking.CallOpts, addr)
}

// IsStaker is a free data retrieval call binding the contract method 0x6f1e8533.
//
// Solidity: function isStaker(address addr) view returns(bool)
func (_L1Staking *L1StakingCallerSession) IsStaker(addr common.Address) (bool, error) {
	return _L1Staking.Contract.IsStaker(&_L1Staking.CallOpts, addr)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Staking *L1StakingCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Staking *L1StakingSession) Messenger() (common.Address, error) {
	return _L1Staking.Contract.Messenger(&_L1Staking.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1Staking *L1StakingCallerSession) Messenger() (common.Address, error) {
	return _L1Staking.Contract.Messenger(&_L1Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Staking *L1StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Staking *L1StakingSession) Owner() (common.Address, error) {
	return _L1Staking.Contract.Owner(&_L1Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1Staking *L1StakingCallerSession) Owner() (common.Address, error) {
	return _L1Staking.Contract.Owner(&_L1Staking.CallOpts)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_L1Staking *L1StakingCaller) StakerList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "stakerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_L1Staking *L1StakingSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _L1Staking.Contract.StakerList(&_L1Staking.CallOpts, arg0)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_L1Staking *L1StakingCallerSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _L1Staking.Contract.StakerList(&_L1Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Addr   common.Address
		TmKey  [32]byte
		BlsKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TmKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingSession) Stakers(arg0 common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L1Staking.Contract.Stakers(&_L1Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingCallerSession) Stakers(arg0 common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L1Staking.Contract.Stakers(&_L1Staking.CallOpts, arg0)
}

// VerifySignature is a free data retrieval call binding the contract method 0xc0af545b.
//
// Solidity: function verifySignature(address[] signedSequencers, address[] sequencerSet, bytes32 msgHash, bytes signature) pure returns(bool)
func (_L1Staking *L1StakingCaller) VerifySignature(opts *bind.CallOpts, signedSequencers []common.Address, sequencerSet []common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "verifySignature", signedSequencers, sequencerSet, msgHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0xc0af545b.
//
// Solidity: function verifySignature(address[] signedSequencers, address[] sequencerSet, bytes32 msgHash, bytes signature) pure returns(bool)
func (_L1Staking *L1StakingSession) VerifySignature(signedSequencers []common.Address, sequencerSet []common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	return _L1Staking.Contract.VerifySignature(&_L1Staking.CallOpts, signedSequencers, sequencerSet, msgHash, signature)
}

// VerifySignature is a free data retrieval call binding the contract method 0xc0af545b.
//
// Solidity: function verifySignature(address[] signedSequencers, address[] sequencerSet, bytes32 msgHash, bytes signature) pure returns(bool)
func (_L1Staking *L1StakingCallerSession) VerifySignature(signedSequencers []common.Address, sequencerSet []common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	return _L1Staking.Contract.VerifySignature(&_L1Staking.CallOpts, signedSequencers, sequencerSet, msgHash, signature)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_L1Staking *L1StakingCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_L1Staking *L1StakingSession) Whitelist(arg0 common.Address) (bool, error) {
	return _L1Staking.Contract.Whitelist(&_L1Staking.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_L1Staking *L1StakingCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _L1Staking.Contract.Whitelist(&_L1Staking.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_L1Staking *L1StakingCaller) Withdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "withdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_L1Staking *L1StakingSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _L1Staking.Contract.Withdrawals(&_L1Staking.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_L1Staking *L1StakingCallerSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _L1Staking.Contract.Withdrawals(&_L1Staking.CallOpts, arg0)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_L1Staking *L1StakingTransactor) ClaimSlashRemaining(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "claimSlashRemaining", receiver)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_L1Staking *L1StakingSession) ClaimSlashRemaining(receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.ClaimSlashRemaining(&_L1Staking.TransactOpts, receiver)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_L1Staking *L1StakingTransactorSession) ClaimSlashRemaining(receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.ClaimSlashRemaining(&_L1Staking.TransactOpts, receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_L1Staking *L1StakingTransactor) ClaimWithdrawal(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "claimWithdrawal", receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_L1Staking *L1StakingSession) ClaimWithdrawal(receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.ClaimWithdrawal(&_L1Staking.TransactOpts, receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_L1Staking *L1StakingTransactorSession) ClaimWithdrawal(receiver common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.ClaimWithdrawal(&_L1Staking.TransactOpts, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _admin, address _rollupContract, uint256 _rewardPercentage, uint256 _stakingValue, uint256 _lockBlocks, uint256 _gasLimit) returns()
func (_L1Staking *L1StakingTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _rollupContract common.Address, _rewardPercentage *big.Int, _stakingValue *big.Int, _lockBlocks *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "initialize", _admin, _rollupContract, _rewardPercentage, _stakingValue, _lockBlocks, _gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _admin, address _rollupContract, uint256 _rewardPercentage, uint256 _stakingValue, uint256 _lockBlocks, uint256 _gasLimit) returns()
func (_L1Staking *L1StakingSession) Initialize(_admin common.Address, _rollupContract common.Address, _rewardPercentage *big.Int, _stakingValue *big.Int, _lockBlocks *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.Contract.Initialize(&_L1Staking.TransactOpts, _admin, _rollupContract, _rewardPercentage, _stakingValue, _lockBlocks, _gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _admin, address _rollupContract, uint256 _rewardPercentage, uint256 _stakingValue, uint256 _lockBlocks, uint256 _gasLimit) returns()
func (_L1Staking *L1StakingTransactorSession) Initialize(_admin common.Address, _rollupContract common.Address, _rewardPercentage *big.Int, _stakingValue *big.Int, _lockBlocks *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.Contract.Initialize(&_L1Staking.TransactOpts, _admin, _rollupContract, _rewardPercentage, _stakingValue, _lockBlocks, _gasLimit)
}

// Register is a paid mutator transaction binding the contract method 0x20b651cd.
//
// Solidity: function register(address addr, bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingTransactor) Register(opts *bind.TransactOpts, addr common.Address, tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "register", addr, tmKey, blsKey)
}

// Register is a paid mutator transaction binding the contract method 0x20b651cd.
//
// Solidity: function register(address addr, bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingSession) Register(addr common.Address, tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.Contract.Register(&_L1Staking.TransactOpts, addr, tmKey, blsKey)
}

// Register is a paid mutator transaction binding the contract method 0x20b651cd.
//
// Solidity: function register(address addr, bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingTransactorSession) Register(addr common.Address, tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.Contract.Register(&_L1Staking.TransactOpts, addr, tmKey, blsKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Staking *L1StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Staking *L1StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1Staking.Contract.RenounceOwnership(&_L1Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1Staking *L1StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1Staking.Contract.RenounceOwnership(&_L1Staking.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] sequencers) returns(uint256)
func (_L1Staking *L1StakingTransactor) Slash(opts *bind.TransactOpts, sequencers []common.Address) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "slash", sequencers)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] sequencers) returns(uint256)
func (_L1Staking *L1StakingSession) Slash(sequencers []common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.Slash(&_L1Staking.TransactOpts, sequencers)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] sequencers) returns(uint256)
func (_L1Staking *L1StakingTransactorSession) Slash(sequencers []common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.Slash(&_L1Staking.TransactOpts, sequencers)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Staking *L1StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Staking *L1StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.TransferOwnership(&_L1Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1Staking *L1StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.TransferOwnership(&_L1Staking.TransactOpts, newOwner)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xa5749187.
//
// Solidity: function updateParams(uint256 gasLimit) returns()
func (_L1Staking *L1StakingTransactor) UpdateParams(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "updateParams", gasLimit)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xa5749187.
//
// Solidity: function updateParams(uint256 gasLimit) returns()
func (_L1Staking *L1StakingSession) UpdateParams(gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.Contract.UpdateParams(&_L1Staking.TransactOpts, gasLimit)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xa5749187.
//
// Solidity: function updateParams(uint256 gasLimit) returns()
func (_L1Staking *L1StakingTransactorSession) UpdateParams(gasLimit *big.Int) (*types.Transaction, error) {
	return _L1Staking.Contract.UpdateParams(&_L1Staking.TransactOpts, gasLimit)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xc7cd469a.
//
// Solidity: function updateWhitelist(address[] add, address[] remove) returns()
func (_L1Staking *L1StakingTransactor) UpdateWhitelist(opts *bind.TransactOpts, add []common.Address, remove []common.Address) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "updateWhitelist", add, remove)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xc7cd469a.
//
// Solidity: function updateWhitelist(address[] add, address[] remove) returns()
func (_L1Staking *L1StakingSession) UpdateWhitelist(add []common.Address, remove []common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.UpdateWhitelist(&_L1Staking.TransactOpts, add, remove)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xc7cd469a.
//
// Solidity: function updateWhitelist(address[] add, address[] remove) returns()
func (_L1Staking *L1StakingTransactorSession) UpdateWhitelist(add []common.Address, remove []common.Address) (*types.Transaction, error) {
	return _L1Staking.Contract.UpdateWhitelist(&_L1Staking.TransactOpts, add, remove)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L1Staking *L1StakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L1Staking *L1StakingSession) Withdraw() (*types.Transaction, error) {
	return _L1Staking.Contract.Withdraw(&_L1Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L1Staking *L1StakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _L1Staking.Contract.Withdraw(&_L1Staking.TransactOpts)
}

// L1StakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the L1Staking contract.
type L1StakingClaimedIterator struct {
	Event *L1StakingClaimed // Event containing the contract specifics and raw log

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
func (it *L1StakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingClaimed)
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
		it.Event = new(L1StakingClaimed)
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
func (it *L1StakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingClaimed represents a Claimed event raised by the L1Staking contract.
type L1StakingClaimed struct {
	Staker   common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b.
//
// Solidity: event Claimed(address staker, address receiver)
func (_L1Staking *L1StakingFilterer) FilterClaimed(opts *bind.FilterOpts) (*L1StakingClaimedIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &L1StakingClaimedIterator{contract: _L1Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b.
//
// Solidity: event Claimed(address staker, address receiver)
func (_L1Staking *L1StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *L1StakingClaimed) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingClaimed)
				if err := _L1Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b.
//
// Solidity: event Claimed(address staker, address receiver)
func (_L1Staking *L1StakingFilterer) ParseClaimed(log types.Log) (*L1StakingClaimed, error) {
	event := new(L1StakingClaimed)
	if err := _L1Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1Staking contract.
type L1StakingInitializedIterator struct {
	Event *L1StakingInitialized // Event containing the contract specifics and raw log

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
func (it *L1StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingInitialized)
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
		it.Event = new(L1StakingInitialized)
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
func (it *L1StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingInitialized represents a Initialized event raised by the L1Staking contract.
type L1StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Staking *L1StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1StakingInitializedIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1StakingInitializedIterator{contract: _L1Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1Staking *L1StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingInitialized)
				if err := _L1Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1Staking *L1StakingFilterer) ParseInitialized(log types.Log) (*L1StakingInitialized, error) {
	event := new(L1StakingInitialized)
	if err := _L1Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1Staking contract.
type L1StakingOwnershipTransferredIterator struct {
	Event *L1StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingOwnershipTransferred)
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
		it.Event = new(L1StakingOwnershipTransferred)
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
func (it *L1StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingOwnershipTransferred represents a OwnershipTransferred event raised by the L1Staking contract.
type L1StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1Staking *L1StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1StakingOwnershipTransferredIterator{contract: _L1Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1Staking *L1StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingOwnershipTransferred)
				if err := _L1Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1Staking *L1StakingFilterer) ParseOwnershipTransferred(log types.Log) (*L1StakingOwnershipTransferred, error) {
	event := new(L1StakingOwnershipTransferred)
	if err := _L1Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingParamsUpdatedIterator is returned from FilterParamsUpdated and is used to iterate over the raw logs and unpacked data for ParamsUpdated events raised by the L1Staking contract.
type L1StakingParamsUpdatedIterator struct {
	Event *L1StakingParamsUpdated // Event containing the contract specifics and raw log

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
func (it *L1StakingParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingParamsUpdated)
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
		it.Event = new(L1StakingParamsUpdated)
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
func (it *L1StakingParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingParamsUpdated represents a ParamsUpdated event raised by the L1Staking contract.
type L1StakingParamsUpdated struct {
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterParamsUpdated is a free log retrieval operation binding the contract event 0xa96b260c11da5ffa5f74f6cd6dcb582ef40c552985b8622dd901e63ecee02b3b.
//
// Solidity: event ParamsUpdated(uint256 gasLimit)
func (_L1Staking *L1StakingFilterer) FilterParamsUpdated(opts *bind.FilterOpts) (*L1StakingParamsUpdatedIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &L1StakingParamsUpdatedIterator{contract: _L1Staking.contract, event: "ParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchParamsUpdated is a free log subscription operation binding the contract event 0xa96b260c11da5ffa5f74f6cd6dcb582ef40c552985b8622dd901e63ecee02b3b.
//
// Solidity: event ParamsUpdated(uint256 gasLimit)
func (_L1Staking *L1StakingFilterer) WatchParamsUpdated(opts *bind.WatchOpts, sink chan<- *L1StakingParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingParamsUpdated)
				if err := _L1Staking.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
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

// ParseParamsUpdated is a log parse operation binding the contract event 0xa96b260c11da5ffa5f74f6cd6dcb582ef40c552985b8622dd901e63ecee02b3b.
//
// Solidity: event ParamsUpdated(uint256 gasLimit)
func (_L1Staking *L1StakingFilterer) ParseParamsUpdated(log types.Log) (*L1StakingParamsUpdated, error) {
	event := new(L1StakingParamsUpdated)
	if err := _L1Staking.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the L1Staking contract.
type L1StakingRegisteredIterator struct {
	Event *L1StakingRegistered // Event containing the contract specifics and raw log

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
func (it *L1StakingRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingRegistered)
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
		it.Event = new(L1StakingRegistered)
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
func (it *L1StakingRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingRegistered represents a Registered event raised by the L1Staking contract.
type L1StakingRegistered struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e199.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingFilterer) FilterRegistered(opts *bind.FilterOpts) (*L1StakingRegisteredIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &L1StakingRegisteredIterator{contract: _L1Staking.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e199.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *L1StakingRegistered) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingRegistered)
				if err := _L1Staking.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e199.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey)
func (_L1Staking *L1StakingFilterer) ParseRegistered(log types.Log) (*L1StakingRegistered, error) {
	event := new(L1StakingRegistered)
	if err := _L1Staking.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the L1Staking contract.
type L1StakingSlashedIterator struct {
	Event *L1StakingSlashed // Event containing the contract specifics and raw log

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
func (it *L1StakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingSlashed)
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
		it.Event = new(L1StakingSlashed)
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
func (it *L1StakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingSlashed represents a Slashed event raised by the L1Staking contract.
type L1StakingSlashed struct {
	Stakers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a.
//
// Solidity: event Slashed(address[] stakers)
func (_L1Staking *L1StakingFilterer) FilterSlashed(opts *bind.FilterOpts) (*L1StakingSlashedIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return &L1StakingSlashedIterator{contract: _L1Staking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a.
//
// Solidity: event Slashed(address[] stakers)
func (_L1Staking *L1StakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *L1StakingSlashed) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingSlashed)
				if err := _L1Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a.
//
// Solidity: event Slashed(address[] stakers)
func (_L1Staking *L1StakingFilterer) ParseSlashed(log types.Log) (*L1StakingSlashed, error) {
	event := new(L1StakingSlashed)
	if err := _L1Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StakingWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the L1Staking contract.
type L1StakingWithdrawedIterator struct {
	Event *L1StakingWithdrawed // Event containing the contract specifics and raw log

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
func (it *L1StakingWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingWithdrawed)
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
		it.Event = new(L1StakingWithdrawed)
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
func (it *L1StakingWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingWithdrawed represents a Withdrawed event raised by the L1Staking contract.
type L1StakingWithdrawed struct {
	Addr       common.Address
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 unlockTime)
func (_L1Staking *L1StakingFilterer) FilterWithdrawed(opts *bind.FilterOpts) (*L1StakingWithdrawedIterator, error) {

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return &L1StakingWithdrawedIterator{contract: _L1Staking.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 unlockTime)
func (_L1Staking *L1StakingFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *L1StakingWithdrawed) (event.Subscription, error) {

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingWithdrawed)
				if err := _L1Staking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
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

// ParseWithdrawed is a log parse operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 unlockTime)
func (_L1Staking *L1StakingFilterer) ParseWithdrawed(log types.Log) (*L1StakingWithdrawed, error) {
	event := new(L1StakingWithdrawed)
	if err := _L1Staking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
