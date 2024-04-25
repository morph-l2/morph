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

// L1StakingMetaData contains all meta data concerning the L1Staking contract.
var L1StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"ParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_STAKING\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blsPubkey\",\"type\":\"bytes\"}],\"name\":\"blsKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimSlashRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tmPubkey\",\"type\":\"bytes32\"}],\"name\":\"tmKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"add\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signedSequencers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sequencerSet\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalLockBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"withdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b5060405162002a4338038062002a4383398101604081905262000033916200005d565b6001600160a01b031660805273530000000000000000000000000000000000001260a0526200008c565b5f602082840312156200006e575f80fd5b81516001600160a01b038116811462000085575f80fd5b9392505050565b60805160a051612972620000d15f395f8181610389015281816118fe0152611a7b01525f81816101d601528181610452015281816118c20152611a3f01526129725ff3fe6080604052600436106101af575f3560e01c80638b8c24c1116100e7578063a574918711610087578063c0af545b11610062578063c0af545b14610536578063c7cd469a1461055a578063cde4cd1114610579578063f2fde38b14610598575f80fd5b8063a5749187146104d6578063ab8c53dc146104f5578063bfa02ba91461050a575f80fd5b8063927ede2d116100c2578063927ede2d1461044157806395368d2e146104745780639b19251a14610489578063a3066aab146104b7575f80fd5b80638b8c24c1146103ca5780638da5cb5b146103e95780639168ae7214610413575f80fd5b8063692c565b116101525780637a4e87c31161012d5780637a4e87c3146103135780637a9262a21461034d578063831cfb581461037857806386489ba9146103ab575f80fd5b8063692c565b146102a25780636f1e8533146102e0578063715018a6146102ff575f80fd5b806341de239b1161018d57806341de239b1461023457806343352d61146102575780634d64903a1461027857806352d472eb1461028d575f80fd5b80632a28e5a3146101b35780633cb747bf146101c85780633ccfd60b14610220575b5f80fd5b6101c66101c1366004612213565b6105b7565b005b3480156101d3575f80fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561022b575f80fd5b506101c66109f4565b34801561023f575f80fd5b5061024960995481565b604051908152602001610217565b348015610262575f80fd5b5061026b610b90565b6040516102179190612257565b348015610283575f80fd5b5061024960985481565b348015610298575f80fd5b50610249609a5481565b3480156102ad575f80fd5b506102d06102bc3660046122b0565b60a26020525f908152604090205460ff1681565b6040519015158152602001610217565b3480156102eb575f80fd5b506102d06102fa3660046122ea565b610ba1565b34801561030a575f80fd5b506101c6610bb3565b34801561031e575f80fd5b506102d061032d366004612303565b805160208183018101805160a18252928201919093012091525460ff1681565b348015610358575f80fd5b506102496103673660046122ea565b60a36020525f908152604090205481565b348015610383575f80fd5b506101f67f000000000000000000000000000000000000000000000000000000000000000081565b3480156103b6575f80fd5b506101c66103c536600461233d565b610bc6565b3480156103d5575f80fd5b506102496103e4366004612412565b610ff5565b3480156103f4575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101f6565b34801561041e575f80fd5b5061043261042d3660046122ea565b6113d8565b60405161021793929190612491565b34801561044c575f80fd5b506101f67f000000000000000000000000000000000000000000000000000000000000000081565b34801561047f575f80fd5b50610249609b5481565b348015610494575f80fd5b506102d06104a33660046122ea565b609d6020525f908152604090205460ff1681565b3480156104c2575f80fd5b506101c66104d13660046122ea565b611497565b3480156104e1575f80fd5b506101c66104f03660046122b0565b611613565b348015610500575f80fd5b50610249609c5481565b348015610515575f80fd5b506097546101f69073ffffffffffffffffffffffffffffffffffffffff1681565b348015610541575f80fd5b506102d06105503660046124ce565b6001949350505050565b348015610565575f80fd5b506101c66105743660046125a2565b6116a5565b348015610584575f80fd5b506101c66105933660046122ea565b6117d4565b3480156105a3575f80fd5b506101c66105b23660046122ea565b6117fe565b335f818152609d602052604090205460ff1661061a5760405162461bcd60e51b815260206004820152601060248201527f6e6f7420696e2077686974656c6973740000000000000000000000000000000060448201526064015b60405180910390fd5b335f90815260a0602052604090205473ffffffffffffffffffffffffffffffffffffffff161561068c5760405162461bcd60e51b815260206004820152601260248201527f616c7265616479207265676973746572656400000000000000000000000000006044820152606401610611565b82158015906106a957505f83815260a2602052604090205460ff16155b6106f55760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642074656e6465726d696e74207075626b6579000000000000006044820152606401610611565b8151610100148015610727575060a1826040516107129190612609565b9081526040519081900360200190205460ff16155b6107735760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420626c73207075626b657900000000000000000000000000006044820152606401610611565b60985434146107c45760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964207374616b696e672076616c756500000000000000000000006044820152606401610611565b60405180606001604052806107d63390565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083018790526040928301869052335f90815260a08252839020845181547fffffffffffffffffffffffff0000000000000000000000000000000000000000169316929092178255830151600182015590820151600282019061085590826126b9565b5090505061086b6108633390565b609e90611898565b50600160a18360405161087e9190612609565b9081526040805191829003602090810190922080549315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009485161790555f86815260a29092529020805490911660011790557fb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e1996108fa3390565b848460405161090b93929190612491565b60405180910390a1335f90815260a060209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff1681526001820154928101929092526002810180546109ef948401919061096a90612624565b80601f016020809104026020016040519081016040528092919081815260200182805461099690612624565b80156109e15780601f106109b8576101008083540402835291602001916109e1565b820191905f5260205f20905b8154815290600101906020018083116109c457829003601f168201915b5050505050815250506118c0565b505050565b6109ff609e336119ee565b610a4b5760405162461bcd60e51b815260206004820152600b60248201527f6f6e6c79207374616b65720000000000000000000000000000000000000000006044820152606401610611565b335f90815260a3602052604090205415610aa75760405162461bcd60e51b815260206004820152600b60248201527f7769746864726177696e670000000000000000000000000000000000000000006044820152606401610611565b609954610ab490436127e4565b335f81815260a36020526040902091909155610ad290609e90611a1c565b50335f81815260a360209081526040918290205491519182527f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5910160405180910390a26040805160018082528183019092525f916020808301908036833701905050905033815f81518110610b4a57610b4a6127f7565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050610b8d81611a3d565b50565b6060610b9c609e611b42565b905090565b5f610bad609e836119ee565b92915050565b610bbb611b4e565b610bc45f611bb5565b565b5f54610100900460ff1615808015610be457505f54600160ff909116105b80610bfd5750303b158015610bfd57505f5460ff166001145b610c6f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610611565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610ccb575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8716610d2e5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642061646d696e000000000000000000000000000000000000006044820152606401610611565b73ffffffffffffffffffffffffffffffffffffffff8616610d915760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e74726163740000000000000000006044820152606401610611565b5f8411610e065760405162461bcd60e51b815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f30000000000000000000000000000000000000000000000000000000000000006064820152608401610611565b5f8311610e7b5760405162461bcd60e51b815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f30000000000000000000000000000000000000000000000000000000000000006064820152608401610611565b5f8211610eca5760405162461bcd60e51b815260206004820152601d60248201527f676173206c696d6974206d7573742067726561746572207468616e20300000006044820152606401610611565b5f85118015610eda575060648511155b610f265760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964207265776172642070657263656e74616765000000000000006044820152606401610611565b610f2e611c2b565b610f36611caf565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8816179055609a85905560988490556099839055609b8290558015610fec575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6097545f9073ffffffffffffffffffffffffffffffffffffffff16331461105e5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610611565b611066611d33565b5f805b8351811015611326575f60a35f868481518110611088576110886127f7565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205411156111c05760a35f8583815181106110e5576110e56127f7565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f905560a05f85838151811061113e5761113e6127f7565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040015f90812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906111aa60028301826120ff565b50506098546111b990836127e4565b91506112b4565b6111ed8482815181106111d5576111d56127f7565b6020026020010151609e6119ee90919063ffffffff16565b15611202576098546111ff90836127e4565b91505b61122f848281518110611217576112176127f7565b6020026020010151609e611a1c90919063ffffffff16565b5060a05f858381518110611245576112456127f7565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040015f90812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906112b160028301826120ff565b50505b609d5f8583815181106112c9576112c96127f7565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055600101611069565b505f6064609a54836113389190612824565b611342919061283b565b905061134e8183612873565b609c5f82825461135e91906127e4565b90915550506097546113869073ffffffffffffffffffffffffffffffffffffffff1682611d8c565b7f654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a846040516113b59190612257565b60405180910390a16113c684611a3d565b9150506113d36001606555565b919050565b60a06020525f908152604090208054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff90931693919261141690612624565b80601f016020809104026020016040519081016040528092919081815260200182805461144290612624565b801561148d5780601f106114645761010080835404028352916020019161148d565b820191905f5260205f20905b81548152906001019060200180831161147057829003601f168201915b5050505050905083565b61149f611d33565b335f90815260a360205260409020546114fa5760405162461bcd60e51b815260206004820152601460248201527f7769746864726177616c206e6f742065786973740000000000000000000000006044820152606401610611565b335f90815260a3602052604090205443116115575760405162461bcd60e51b815260206004820152601160248201527f7769746864726177616c206c6f636b65640000000000000000000000000000006044820152606401610611565b335f90815260a06020526040812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906115a060028301826120ff565b5050335f81815260a36020908152604080832092909255905173ffffffffffffffffffffffffffffffffffffffff841681527f89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b910160405180910390a261160981609854611d8c565b610b8d6001606555565b61161b611b4e565b5f811161166a5760405162461bcd60e51b815260206004820152601d60248201527f676173206c696d6974206d7573742067726561746572207468616e20300000006044820152606401610611565b609b8190556040518181527fa96b260c11da5ffa5f74f6cd6dcb582ef40c552985b8622dd901e63ecee02b3b9060200160405180910390a150565b6116ad611b4e565b5f5b8381101561173d576001609d5f8787858181106116ce576116ce6127f7565b90506020020160208101906116e391906122ea565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016116af565b505f5b818110156117cd575f609d5f85858581811061175e5761175e6127f7565b905060200201602081019061177391906122ea565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101611740565b5050505050565b6117dc611b4e565b6117e4611d33565b6117f081609c54611d8c565b5f609c55610b8d6001606555565b611806611b4e565b73ffffffffffffffffffffffffffffffffffffffff811661188f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610611565b610b8d81611bb5565b5f6118b98373ffffffffffffffffffffffffffffffffffffffff8416611e76565b9392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b2267a7b7f00000000000000000000000000000000000000000000000000000000000000005f8460405160240161192f9190612886565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6d454d5100000000000000000000000000000000000000000000000000000000179052609b5490517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526119c594939291906004016128ca565b5f604051808303815f87803b1580156119dc575f80fd5b505af11580156117cd573d5f803e3d5ffd5b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156118b9565b5f6118b98373ffffffffffffffffffffffffffffffffffffffff8416611ec2565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b2267a7b7f00000000000000000000000000000000000000000000000000000000000000005f84604051602401611aac9190612257565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0be67fcc00000000000000000000000000000000000000000000000000000000179052609b5490517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526119c594939291906004016128ca565b60605f6118b983611fa5565b60335473ffffffffffffffffffffffffffffffffffffffff163314610bc45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610611565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611ca75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610611565b610bc4611ffe565b5f54610100900460ff16611d2b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610611565b610bc4612083565b600260655403611d855760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610611565b6002606555565b8015611e6b575f8273ffffffffffffffffffffffffffffffffffffffff1682604051611ddb907f3078000000000000000000000000000000000000000000000000000000000000815260020190565b5f6040518083038185875af1925050503d805f8114611e15576040519150601f19603f3d011682016040523d82523d5f602084013e611e1a565b606091505b50509050806109ef5760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610611565b5050565b6001606555565b5f818152600183016020526040812054611ebb57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610bad565b505f610bad565b5f8181526001830160205260408120548015611f9c575f611ee4600183612873565b85549091505f90611ef790600190612873565b9050818114611f56575f865f018281548110611f1557611f156127f7565b905f5260205f200154905080875f018481548110611f3557611f356127f7565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611f6757611f6761290f565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610bad565b5f915050610bad565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611ff257602002820191905f5260205f20905b815481526020019060010190808311611fde575b50505050509050919050565b5f54610100900460ff1661207a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610611565b610bc433611bb5565b5f54610100900460ff16611e6f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610611565b50805461210b90612624565b5f825580601f1061211a575050565b601f0160209004905f5260205f2090810190610b8d91905b80821115612145575f8155600101612132565b5090565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561219f5761219f612149565b604052919050565b5f82601f8301126121b6575f80fd5b813567ffffffffffffffff8111156121d0576121d0612149565b6121e36020601f19601f84011601612176565b8181528460208386010111156121f7575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215612224575f80fd5b82359150602083013567ffffffffffffffff811115612241575f80fd5b61224d858286016121a7565b9150509250929050565b602080825282518282018190525f9190848201906040850190845b818110156122a457835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612272565b50909695505050505050565b5f602082840312156122c0575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113d3575f80fd5b5f602082840312156122fa575f80fd5b6118b9826122c7565b5f60208284031215612313575f80fd5b813567ffffffffffffffff811115612329575f80fd5b612335848285016121a7565b949350505050565b5f805f805f8060c08789031215612352575f80fd5b61235b876122c7565b9550612369602088016122c7565b95989597505050506040840135936060810135936080820135935060a0909101359150565b5f82601f83011261239d575f80fd5b8135602067ffffffffffffffff8211156123b9576123b9612149565b8160051b6123c8828201612176565b92835284810182019282810190878511156123e1575f80fd5b83870192505b84831015612407576123f8836122c7565b825291830191908301906123e7565b979650505050505050565b5f60208284031215612422575f80fd5b813567ffffffffffffffff811115612438575f80fd5b6123358482850161238e565b5f5b8381101561245e578181015183820152602001612446565b50505f910152565b5f815180845261247d816020860160208601612444565b601f01601f19169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f6124c56060830184612466565b95945050505050565b5f805f80608085870312156124e1575f80fd5b843567ffffffffffffffff808211156124f8575f80fd5b6125048883890161238e565b95506020870135915080821115612519575f80fd5b6125258883890161238e565b9450604087013593506060870135915080821115612541575f80fd5b5061254e878288016121a7565b91505092959194509250565b5f8083601f84011261256a575f80fd5b50813567ffffffffffffffff811115612581575f80fd5b6020830191508360208260051b850101111561259b575f80fd5b9250929050565b5f805f80604085870312156125b5575f80fd5b843567ffffffffffffffff808211156125cc575f80fd5b6125d88883890161255a565b909650945060208701359150808211156125f0575f80fd5b506125fd8782880161255a565b95989497509550505050565b5f825161261a818460208701612444565b9190910192915050565b600181811c9082168061263857607f821691505b60208210810361266f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156109ef57805f5260205f20601f840160051c8101602085101561269a5750805b601f840160051c820191505b818110156117cd575f81556001016126a6565b815167ffffffffffffffff8111156126d3576126d3612149565b6126e7816126e18454612624565b84612675565b602080601f831160018114612739575f84156127035750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556127af565b5f85815260208120601f198616915b8281101561276757888601518255948401946001909101908401612748565b50858210156127a357878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115610bad57610bad6127b7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417610bad57610bad6127b7565b5f8261286e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b81810381811115610bad57610bad6127b7565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201525f60408301516060808401526123356080840182612466565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f6128fe6080830185612466565b905082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea264697066735822122038a925287c90d42f176c71837c5aee3f0e8f2a1cc30d14e11cfcbe821a2d3d0764736f6c63430008180033",
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
	parsed, err := abi.JSON(strings.NewReader(L1StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// BlsKeys is a free data retrieval call binding the contract method 0x7a4e87c3.
//
// Solidity: function blsKeys(bytes blsPubkey) view returns(bool)
func (_L1Staking *L1StakingCaller) BlsKeys(opts *bind.CallOpts, blsPubkey []byte) (bool, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "blsKeys", blsPubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlsKeys is a free data retrieval call binding the contract method 0x7a4e87c3.
//
// Solidity: function blsKeys(bytes blsPubkey) view returns(bool)
func (_L1Staking *L1StakingSession) BlsKeys(blsPubkey []byte) (bool, error) {
	return _L1Staking.Contract.BlsKeys(&_L1Staking.CallOpts, blsPubkey)
}

// BlsKeys is a free data retrieval call binding the contract method 0x7a4e87c3.
//
// Solidity: function blsKeys(bytes blsPubkey) view returns(bool)
func (_L1Staking *L1StakingCallerSession) BlsKeys(blsPubkey []byte) (bool, error) {
	return _L1Staking.Contract.BlsKeys(&_L1Staking.CallOpts, blsPubkey)
}

// DefaultGasLimit is a free data retrieval call binding the contract method 0x95368d2e.
//
// Solidity: function defaultGasLimit() view returns(uint256)
func (_L1Staking *L1StakingCaller) DefaultGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "defaultGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultGasLimit is a free data retrieval call binding the contract method 0x95368d2e.
//
// Solidity: function defaultGasLimit() view returns(uint256)
func (_L1Staking *L1StakingSession) DefaultGasLimit() (*big.Int, error) {
	return _L1Staking.Contract.DefaultGasLimit(&_L1Staking.CallOpts)
}

// DefaultGasLimit is a free data retrieval call binding the contract method 0x95368d2e.
//
// Solidity: function defaultGasLimit() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) DefaultGasLimit() (*big.Int, error) {
	return _L1Staking.Contract.DefaultGasLimit(&_L1Staking.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[])
func (_L1Staking *L1StakingCaller) GetStakers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "getStakers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[])
func (_L1Staking *L1StakingSession) GetStakers() ([]common.Address, error) {
	return _L1Staking.Contract.GetStakers(&_L1Staking.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[])
func (_L1Staking *L1StakingCallerSession) GetStakers() ([]common.Address, error) {
	return _L1Staking.Contract.GetStakers(&_L1Staking.CallOpts)
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

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_L1Staking *L1StakingCaller) RewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "rewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_L1Staking *L1StakingSession) RewardPercentage() (*big.Int, error) {
	return _L1Staking.Contract.RewardPercentage(&_L1Staking.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) RewardPercentage() (*big.Int, error) {
	return _L1Staking.Contract.RewardPercentage(&_L1Staking.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Staking *L1StakingCaller) RollupContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "rollupContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Staking *L1StakingSession) RollupContract() (common.Address, error) {
	return _L1Staking.Contract.RollupContract(&_L1Staking.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_L1Staking *L1StakingCallerSession) RollupContract() (common.Address, error) {
	return _L1Staking.Contract.RollupContract(&_L1Staking.CallOpts)
}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_L1Staking *L1StakingCaller) SlashRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "slashRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_L1Staking *L1StakingSession) SlashRemaining() (*big.Int, error) {
	return _L1Staking.Contract.SlashRemaining(&_L1Staking.CallOpts)
}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) SlashRemaining() (*big.Int, error) {
	return _L1Staking.Contract.SlashRemaining(&_L1Staking.CallOpts)
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

// StakingValue is a free data retrieval call binding the contract method 0x4d64903a.
//
// Solidity: function stakingValue() view returns(uint256)
func (_L1Staking *L1StakingCaller) StakingValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "stakingValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingValue is a free data retrieval call binding the contract method 0x4d64903a.
//
// Solidity: function stakingValue() view returns(uint256)
func (_L1Staking *L1StakingSession) StakingValue() (*big.Int, error) {
	return _L1Staking.Contract.StakingValue(&_L1Staking.CallOpts)
}

// StakingValue is a free data retrieval call binding the contract method 0x4d64903a.
//
// Solidity: function stakingValue() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) StakingValue() (*big.Int, error) {
	return _L1Staking.Contract.StakingValue(&_L1Staking.CallOpts)
}

// TmKeys is a free data retrieval call binding the contract method 0x692c565b.
//
// Solidity: function tmKeys(bytes32 tmPubkey) view returns(bool)
func (_L1Staking *L1StakingCaller) TmKeys(opts *bind.CallOpts, tmPubkey [32]byte) (bool, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "tmKeys", tmPubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TmKeys is a free data retrieval call binding the contract method 0x692c565b.
//
// Solidity: function tmKeys(bytes32 tmPubkey) view returns(bool)
func (_L1Staking *L1StakingSession) TmKeys(tmPubkey [32]byte) (bool, error) {
	return _L1Staking.Contract.TmKeys(&_L1Staking.CallOpts, tmPubkey)
}

// TmKeys is a free data retrieval call binding the contract method 0x692c565b.
//
// Solidity: function tmKeys(bytes32 tmPubkey) view returns(bool)
func (_L1Staking *L1StakingCallerSession) TmKeys(tmPubkey [32]byte) (bool, error) {
	return _L1Staking.Contract.TmKeys(&_L1Staking.CallOpts, tmPubkey)
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

// WithdrawalLockBlocks is a free data retrieval call binding the contract method 0x41de239b.
//
// Solidity: function withdrawalLockBlocks() view returns(uint256)
func (_L1Staking *L1StakingCaller) WithdrawalLockBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "withdrawalLockBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalLockBlocks is a free data retrieval call binding the contract method 0x41de239b.
//
// Solidity: function withdrawalLockBlocks() view returns(uint256)
func (_L1Staking *L1StakingSession) WithdrawalLockBlocks() (*big.Int, error) {
	return _L1Staking.Contract.WithdrawalLockBlocks(&_L1Staking.CallOpts)
}

// WithdrawalLockBlocks is a free data retrieval call binding the contract method 0x41de239b.
//
// Solidity: function withdrawalLockBlocks() view returns(uint256)
func (_L1Staking *L1StakingCallerSession) WithdrawalLockBlocks() (*big.Int, error) {
	return _L1Staking.Contract.WithdrawalLockBlocks(&_L1Staking.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address staker) view returns(uint256)
func (_L1Staking *L1StakingCaller) Withdrawals(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1Staking.contract.Call(opts, &out, "withdrawals", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address staker) view returns(uint256)
func (_L1Staking *L1StakingSession) Withdrawals(staker common.Address) (*big.Int, error) {
	return _L1Staking.Contract.Withdrawals(&_L1Staking.CallOpts, staker)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address staker) view returns(uint256)
func (_L1Staking *L1StakingCallerSession) Withdrawals(staker common.Address) (*big.Int, error) {
	return _L1Staking.Contract.Withdrawals(&_L1Staking.CallOpts, staker)
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

// Register is a paid mutator transaction binding the contract method 0x2a28e5a3.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingTransactor) Register(opts *bind.TransactOpts, tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.contract.Transact(opts, "register", tmKey, blsKey)
}

// Register is a paid mutator transaction binding the contract method 0x2a28e5a3.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingSession) Register(tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.Contract.Register(&_L1Staking.TransactOpts, tmKey, blsKey)
}

// Register is a paid mutator transaction binding the contract method 0x2a28e5a3.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey) payable returns()
func (_L1Staking *L1StakingTransactorSession) Register(tmKey [32]byte, blsKey []byte) (*types.Transaction, error) {
	return _L1Staking.Contract.Register(&_L1Staking.TransactOpts, tmKey, blsKey)
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
// Solidity: event Claimed(address indexed staker, address receiver)
func (_L1Staking *L1StakingFilterer) FilterClaimed(opts *bind.FilterOpts, staker []common.Address) (*L1StakingClaimedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Claimed", stakerRule)
	if err != nil {
		return nil, err
	}
	return &L1StakingClaimedIterator{contract: _L1Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b.
//
// Solidity: event Claimed(address indexed staker, address receiver)
func (_L1Staking *L1StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *L1StakingClaimed, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Claimed", stakerRule)
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
// Solidity: event Claimed(address indexed staker, address receiver)
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

// L1StakingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the L1Staking contract.
type L1StakingWithdrawnIterator struct {
	Event *L1StakingWithdrawn // Event containing the contract specifics and raw log

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
func (it *L1StakingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StakingWithdrawn)
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
		it.Event = new(L1StakingWithdrawn)
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
func (it *L1StakingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StakingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StakingWithdrawn represents a Withdrawn event raised by the L1Staking contract.
type L1StakingWithdrawn struct {
	Addr         common.Address
	UnlockHeight *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed addr, uint256 unlockHeight)
func (_L1Staking *L1StakingFilterer) FilterWithdrawn(opts *bind.FilterOpts, addr []common.Address) (*L1StakingWithdrawnIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _L1Staking.contract.FilterLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return &L1StakingWithdrawnIterator{contract: _L1Staking.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed addr, uint256 unlockHeight)
func (_L1Staking *L1StakingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *L1StakingWithdrawn, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _L1Staking.contract.WatchLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StakingWithdrawn)
				if err := _L1Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed addr, uint256 unlockHeight)
func (_L1Staking *L1StakingFilterer) ParseWithdrawn(log types.Log) (*L1StakingWithdrawn, error) {
	event := new(L1StakingWithdrawn)
	if err := _L1Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
