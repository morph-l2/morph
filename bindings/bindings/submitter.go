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

// SubmitterMetaData contains all meta data concerning the Submitter contract.
var SubmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChallengeDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinimumStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RewardPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashRemainingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"SubmitterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"SubmitterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"addSubmitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimSlashRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollupContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPercentage_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"removeSubmitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"stakeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateChallengeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateRewardPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"withdrawalBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"withdrawing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawing\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100d9565b5f54610100900460ff16156100885760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100d7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611d58806100e65f395ff3fe608060405260043610610183575f3560e01c80639f8a13d7116100d1578063c36f41e21161007c578063d13f90b411610057578063d13f90b41461045b578063ec5ffac21461047a578063f2fde38b1461048f575f80fd5b8063c36f41e2146103fe578063c96be4cb1461041d578063cde4cd111461043c575f80fd5b8063ab8c53dc116100ac578063ab8c53dc1461038f578063b2dd5c07146103a4578063bfa02ba9146103d2575f80fd5b80639f8a13d714610332578063a3066aab14610351578063a4f209b014610370575f80fd5b80633bf1944b1161013157806352d472eb1161010c57806352d472eb146102be578063715018a6146102d35780638da5cb5b146102e7575f80fd5b80633bf1944b146102545780633ccfd60b1461027f5780634262336014610293575f80fd5b8063359289911161016157806335928991146101ef5780633a4b66f11461020e5780633baf3a2f14610216575f80fd5b8063072900f9146101875780630d13fd7b146101a8578063233e9903146101d0575b5f80fd5b348015610192575f80fd5b506101a66101a1366004611c12565b6104ae565b005b3480156101b3575f80fd5b506101bd60995481565b6040519081526020015b60405180910390f35b3480156101db575f80fd5b506101a66101ea366004611c32565b610608565b3480156101fa575f80fd5b506101a6610209366004611c32565b6106ae565b6101a6610754565b348015610221575f80fd5b50610244610230366004611c12565b609e6020525f908152604090205460ff1681565b60405190151581526020016101c7565b34801561025f575f80fd5b506101bd61026e366004611c12565b609f6020525f908152604090205481565b34801561028a575f80fd5b506101a66108e9565b34801561029e575f80fd5b506101bd6102ad366004611c12565b609d6020525f908152604090205481565b3480156102c9575f80fd5b506101bd609a5481565b3480156102de575f80fd5b506101a66109ec565b3480156102f2575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c7565b34801561033d575f80fd5b5061024461034c366004611c12565b6109fd565b34801561035c575f80fd5b506101a661036b366004611c12565b610a8e565b34801561037b575f80fd5b506101a661038a366004611c32565b610d21565b34801561039a575f80fd5b506101bd609b5481565b3480156103af575f80fd5b506102446103be366004611c12565b609c6020525f908152604090205460ff1681565b3480156103dd575f80fd5b5060975461030d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610409575f80fd5b506101a6610418366004611c12565b610dd4565b348015610428575f80fd5b506101bd610437366004611c12565b610f0c565b348015610447575f80fd5b506101a6610456366004611c12565b6110ce565b348015610466575f80fd5b506101a6610475366004611c49565b6111fa565b348015610485575f80fd5b506101bd60985481565b34801561049a575f80fd5b506101a66104a9366004611c12565b611642565b6104b66116dc565b73ffffffffffffffffffffffffffffffffffffffff811615801590610500575073ffffffffffffffffffffffffffffffffffffffff81165f908152609c602052604090205460ff16155b8015610531575073ffffffffffffffffffffffffffffffffffffffff81165f908152609e602052604090205460ff16155b801561055f575073ffffffffffffffffffffffffffffffffffffffff81165f908152609d6020526040902054155b6105b05760405162461bcd60e51b815260206004820152601160248201527f696e76616c6964207375626d697474657200000000000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152609c6020526040808220805460ff19166001179055517fb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b9190a250565b6106106116dc565b5f8111801561062157506098548114155b61066d5760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964206d696e696d756d207374616b65000000000000000000000060448201526064016105a7565b60985460408051918252602082018390527fd67ed534faf3e0dcda08ac6043ec96dcf9b6ebec56055fd14bd0cb40a27c0e49910160405180910390a1609855565b6106b66116dc565b5f811180156106c757506099548114155b6107135760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964206368616c6c656e6765206465706f7369740000000000000060448201526064016105a7565b60995460408051918252602082018390527f36f971a40478225aeb80cfbf5e80306e8cb76d3bf7d56fdc5e490945cddb7d55910160405180910390a1609955565b61075c611743565b335f818152609c602052604090205460ff16801561079f575073ffffffffffffffffffffffffffffffffffffffff81165f908152609e602052604090205460ff16155b6107eb5760405162461bcd60e51b815260206004820152600d60248201527f6e6f74207374616b6561626c650000000000000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609d602052604081205461081b903490611cbf565b90505f3411801561082e57506098548110155b61087a5760405162461bcd60e51b815260206004820152601360248201527f62656c6f77206d696e696d756d207374616b650000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff82165f818152609d602090815260409182902084905581513481529081018490527f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90910160405180910390a250506108e76001606555565b565b335f818152609c602052604090205460ff16801561092c575073ffffffffffffffffffffffffffffffffffffffff81165f908152609e602052604090205460ff16155b801561095b575073ffffffffffffffffffffffffffffffffffffffff81165f908152609d602052604090205415155b6109a75760405162461bcd60e51b815260206004820152601060248201527f6e6f7420776974686472617761626c650000000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609c60209081526040808320805460ff19169055609d9091529020546109e99082906117a3565b50565b6109f46116dc565b6108e75f6118b4565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609c602052604081205460ff168015610a56575073ffffffffffffffffffffffffffffffffffffffff82165f908152609e602052604090205460ff16155b8015610a88575060985473ffffffffffffffffffffffffffffffffffffffff83165f908152609d602052604090205410155b92915050565b610a96611743565b3373ffffffffffffffffffffffffffffffffffffffff8216610afa5760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642072656365697665720000000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609e602052604090205460ff16610b6e5760405162461bcd60e51b815260206004820152600f60248201527f6e6f74207769746864726177696e67000000000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff8082165f908152609f60209081526040918290205460975483517f059def6100000000000000000000000000000000000000000000000000000000815293519194169263059def619260048083019391928290030181865afa158015610beb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c0f9190611cd2565b1015610c5d5760405162461bcd60e51b815260206004820152601e60248201527f7769746864726177616c206261746368206e6f742066696e616c697a6564000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609d60209081526040808320805490849055609e8352818420805460ff19169055609f909252822091909155610cae838261192a565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8188e2b4d95f73db30690b4103c71159349bb897df928902c6330ef99e45fef383604051610d0d91815260200190565b60405180910390a350506109e96001606555565b610d296116dc565b5f81118015610d39575060648111155b8015610d475750609a548114155b610d935760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964207265776172642070657263656e746167650000000000000060448201526064016105a7565b609a5460408051918252602082018390527fa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba910160405180910390a1609a55565b610ddc6116dc565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609c602052604090205460ff168015610e35575073ffffffffffffffffffffffffffffffffffffffff81165f908152609e602052604090205460ff16155b610e815760405162461bcd60e51b815260206004820152601160248201527f696e76616c6964207375626d697474657200000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152609c60209081526040808320805460ff19169055609d9091529020548015610ec857610ec882826117a3565b60405173ffffffffffffffffffffffffffffffffffffffff8316907ff84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1905f90a25050565b6097545f9073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f8b5760405162461bcd60e51b815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e747261637400000000000000000000000060448201526064016105a7565b610f93611743565b73ffffffffffffffffffffffffffffffffffffffff8216610fb557505f6110bf565b73ffffffffffffffffffffffffffffffffffffffff82165f908152609d602090815260408083208054609c8452828520805460ff1990811690915591859055609e84528285208054909216909155609f909252822091909155609a5460649061101e9083611ce9565b6110289190611d00565b91506110348282611d38565b609b5f8282546110449190611cbf565b909155505060975461106c9073ffffffffffffffffffffffffffffffffffffffff168361192a565b604080518281526020810184905273ffffffffffffffffffffffffffffffffffffffff8516917f45a371af55b0726877a30f464edc14db5879ab096590bacce682cf6c18223596910160405180910390a2505b6110c96001606555565b919050565b6110d66116dc565b6110de611743565b73ffffffffffffffffffffffffffffffffffffffff81166111415760405162461bcd60e51b815260206004820152601060248201527f696e76616c69642072656365697665720000000000000000000000000000000060448201526064016105a7565b609b54806111915760405162461bcd60e51b815260206004820152601260248201527f6e6f20736c6173682072656d61696e696e67000000000000000000000000000060448201526064016105a7565b5f609b5561119f828261192a565b8173ffffffffffffffffffffffffffffffffffffffff167fa1fefb6c5328a92a416e321ed50997303fe7135fd88c28b0592b21ce42b5cdd9826040516111e791815260200190565b60405180910390a2506109e96001606555565b5f54610100900460ff161580801561121857505f54600160ff909116105b806112315750303b15801561123157505f5460ff166001145b6112a35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105a7565b5f805460ff1916600117905580156112e1575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff86166113445760405162461bcd60e51b815260206004820152600d60248201527f696e76616c6964206f776e65720000000000000000000000000000000000000060448201526064016105a7565b73ffffffffffffffffffffffffffffffffffffffff85161580159061137f57505f8573ffffffffffffffffffffffffffffffffffffffff163b115b6113cb5760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e747261637400000000000000000060448201526064016105a7565b5f841161141a5760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964206d696e696d756d207374616b65000000000000000000000060448201526064016105a7565b5f83116114695760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964206368616c6c656e6765206465706f7369740000000000000060448201526064016105a7565b5f82118015611479575060648211155b6114c55760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964207265776172642070657263656e746167650000000000000060448201526064016105a7565b6114cd6119e6565b6114d5611a6a565b6114de866118b4565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff871617905560988490556099839055609a829055604080515f8152602081018690527fd67ed534faf3e0dcda08ac6043ec96dcf9b6ebec56055fd14bd0cb40a27c0e49910160405180910390a1604080515f8152602081018590527f36f971a40478225aeb80cfbf5e80306e8cb76d3bf7d56fdc5e490945cddb7d55910160405180910390a1604080515f8152602081018490527fa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba910160405180910390a1801561163a575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b61164a6116dc565b73ffffffffffffffffffffffffffffffffffffffff81166116d35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105a7565b6109e9816118b4565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108e75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105a7565b6002606554036117955760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016105a7565b6002606555565b6001606555565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152609e6020908152604091829020805460ff1916600117905560975482517f121dcd50000000000000000000000000000000000000000000000000000000008152925193169263121dcd509260048082019392918290030181865afa158015611829573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061184d9190611cd2565b73ffffffffffffffffffffffffffffffffffffffff83165f818152609f6020526040908190209290925590517fe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44906118a89084815260200190565b60405180910390a25050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b805f03611935575050565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f811461198b576040519150601f19603f3d011682016040523d82523d5f602084013e611990565b606091505b50509050806119e15760405162461bcd60e51b815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016105a7565b505050565b5f54610100900460ff16611a625760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105a7565b6108e7611aee565b5f54610100900460ff16611ae65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105a7565b6108e7611b73565b5f54610100900460ff16611b6a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105a7565b6108e7336118b4565b5f54610100900460ff1661179c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105a7565b803573ffffffffffffffffffffffffffffffffffffffff811681146110c9575f80fd5b5f60208284031215611c22575f80fd5b611c2b82611bef565b9392505050565b5f60208284031215611c42575f80fd5b5035919050565b5f805f805f60a08688031215611c5d575f80fd5b611c6686611bef565b9450611c7460208701611bef565b94979496505050506040830135926060810135926080909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115610a8857610a88611c92565b5f60208284031215611ce2575f80fd5b5051919050565b8082028115828204841417610a8857610a88611c92565b5f82611d33577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b81810381811115610a8857610a88611c9256fea164736f6c6343000818000a",
}

// SubmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use SubmitterMetaData.ABI instead.
var SubmitterABI = SubmitterMetaData.ABI

// SubmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubmitterMetaData.Bin instead.
var SubmitterBin = SubmitterMetaData.Bin

// DeploySubmitter deploys a new Ethereum contract, binding an instance of Submitter to it.
func DeploySubmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Submitter, error) {
	parsed, err := SubmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// Submitter is an auto generated Go binding around an Ethereum contract.
type Submitter struct {
	SubmitterCaller     // Read-only binding to the contract
	SubmitterTransactor // Write-only binding to the contract
	SubmitterFilterer   // Log filterer for contract events
}

// SubmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmitterSession struct {
	Contract     *Submitter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmitterCallerSession struct {
	Contract *SubmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SubmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmitterTransactorSession struct {
	Contract     *SubmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SubmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmitterRaw struct {
	Contract *Submitter // Generic contract binding to access the raw methods on
}

// SubmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmitterCallerRaw struct {
	Contract *SubmitterCaller // Generic read-only contract binding to access the raw methods on
}

// SubmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmitterTransactorRaw struct {
	Contract *SubmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmitter creates a new instance of Submitter, bound to a specific deployed contract.
func NewSubmitter(address common.Address, backend bind.ContractBackend) (*Submitter, error) {
	contract, err := bindSubmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Submitter{SubmitterCaller: SubmitterCaller{contract: contract}, SubmitterTransactor: SubmitterTransactor{contract: contract}, SubmitterFilterer: SubmitterFilterer{contract: contract}}, nil
}

// NewSubmitterCaller creates a new read-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterCaller(address common.Address, caller bind.ContractCaller) (*SubmitterCaller, error) {
	contract, err := bindSubmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterCaller{contract: contract}, nil
}

// NewSubmitterTransactor creates a new write-only instance of Submitter, bound to a specific deployed contract.
func NewSubmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmitterTransactor, error) {
	contract, err := bindSubmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterTransactor{contract: contract}, nil
}

// NewSubmitterFilterer creates a new log filterer instance of Submitter, bound to a specific deployed contract.
func NewSubmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmitterFilterer, error) {
	contract, err := bindSubmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmitterFilterer{contract: contract}, nil
}

// bindSubmitter binds a generic wrapper to an already deployed contract.
func bindSubmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.SubmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.SubmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Submitter *SubmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Submitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Submitter *SubmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Submitter *SubmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Submitter.Contract.contract.Transact(opts, method, params...)
}

// ChallengeDeposit is a free data retrieval call binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() view returns(uint256)
func (_Submitter *SubmitterCaller) ChallengeDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "challengeDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeDeposit is a free data retrieval call binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() view returns(uint256)
func (_Submitter *SubmitterSession) ChallengeDeposit() (*big.Int, error) {
	return _Submitter.Contract.ChallengeDeposit(&_Submitter.CallOpts)
}

// ChallengeDeposit is a free data retrieval call binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() view returns(uint256)
func (_Submitter *SubmitterCallerSession) ChallengeDeposit() (*big.Int, error) {
	return _Submitter.Contract.ChallengeDeposit(&_Submitter.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address submitter) view returns(bool)
func (_Submitter *SubmitterCaller) IsActive(opts *bind.CallOpts, submitter common.Address) (bool, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "isActive", submitter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address submitter) view returns(bool)
func (_Submitter *SubmitterSession) IsActive(submitter common.Address) (bool, error) {
	return _Submitter.Contract.IsActive(&_Submitter.CallOpts, submitter)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address submitter) view returns(bool)
func (_Submitter *SubmitterCallerSession) IsActive(submitter common.Address) (bool, error) {
	return _Submitter.Contract.IsActive(&_Submitter.CallOpts, submitter)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Submitter *SubmitterCaller) MinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "minimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Submitter *SubmitterSession) MinimumStake() (*big.Int, error) {
	return _Submitter.Contract.MinimumStake(&_Submitter.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Submitter *SubmitterCallerSession) MinimumStake() (*big.Int, error) {
	return _Submitter.Contract.MinimumStake(&_Submitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterSession) Owner() (common.Address, error) {
	return _Submitter.Contract.Owner(&_Submitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Submitter *SubmitterCallerSession) Owner() (common.Address, error) {
	return _Submitter.Contract.Owner(&_Submitter.CallOpts)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address submitter) view returns(bool isRegistered)
func (_Submitter *SubmitterCaller) Registered(opts *bind.CallOpts, submitter common.Address) (bool, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "registered", submitter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address submitter) view returns(bool isRegistered)
func (_Submitter *SubmitterSession) Registered(submitter common.Address) (bool, error) {
	return _Submitter.Contract.Registered(&_Submitter.CallOpts, submitter)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address submitter) view returns(bool isRegistered)
func (_Submitter *SubmitterCallerSession) Registered(submitter common.Address) (bool, error) {
	return _Submitter.Contract.Registered(&_Submitter.CallOpts, submitter)
}

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_Submitter *SubmitterCaller) RewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "rewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_Submitter *SubmitterSession) RewardPercentage() (*big.Int, error) {
	return _Submitter.Contract.RewardPercentage(&_Submitter.CallOpts)
}

// RewardPercentage is a free data retrieval call binding the contract method 0x52d472eb.
//
// Solidity: function rewardPercentage() view returns(uint256)
func (_Submitter *SubmitterCallerSession) RewardPercentage() (*big.Int, error) {
	return _Submitter.Contract.RewardPercentage(&_Submitter.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_Submitter *SubmitterCaller) RollupContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "rollupContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_Submitter *SubmitterSession) RollupContract() (common.Address, error) {
	return _Submitter.Contract.RollupContract(&_Submitter.CallOpts)
}

// RollupContract is a free data retrieval call binding the contract method 0xbfa02ba9.
//
// Solidity: function rollupContract() view returns(address)
func (_Submitter *SubmitterCallerSession) RollupContract() (common.Address, error) {
	return _Submitter.Contract.RollupContract(&_Submitter.CallOpts)
}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_Submitter *SubmitterCaller) SlashRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "slashRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_Submitter *SubmitterSession) SlashRemaining() (*big.Int, error) {
	return _Submitter.Contract.SlashRemaining(&_Submitter.CallOpts)
}

// SlashRemaining is a free data retrieval call binding the contract method 0xab8c53dc.
//
// Solidity: function slashRemaining() view returns(uint256)
func (_Submitter *SubmitterCallerSession) SlashRemaining() (*big.Int, error) {
	return _Submitter.Contract.SlashRemaining(&_Submitter.CallOpts)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address submitter) view returns(uint256 amount)
func (_Submitter *SubmitterCaller) StakeOf(opts *bind.CallOpts, submitter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "stakeOf", submitter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address submitter) view returns(uint256 amount)
func (_Submitter *SubmitterSession) StakeOf(submitter common.Address) (*big.Int, error) {
	return _Submitter.Contract.StakeOf(&_Submitter.CallOpts, submitter)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address submitter) view returns(uint256 amount)
func (_Submitter *SubmitterCallerSession) StakeOf(submitter common.Address) (*big.Int, error) {
	return _Submitter.Contract.StakeOf(&_Submitter.CallOpts, submitter)
}

// WithdrawalBatchIndex is a free data retrieval call binding the contract method 0x3bf1944b.
//
// Solidity: function withdrawalBatchIndex(address submitter) view returns(uint256 batchIndex)
func (_Submitter *SubmitterCaller) WithdrawalBatchIndex(opts *bind.CallOpts, submitter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "withdrawalBatchIndex", submitter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalBatchIndex is a free data retrieval call binding the contract method 0x3bf1944b.
//
// Solidity: function withdrawalBatchIndex(address submitter) view returns(uint256 batchIndex)
func (_Submitter *SubmitterSession) WithdrawalBatchIndex(submitter common.Address) (*big.Int, error) {
	return _Submitter.Contract.WithdrawalBatchIndex(&_Submitter.CallOpts, submitter)
}

// WithdrawalBatchIndex is a free data retrieval call binding the contract method 0x3bf1944b.
//
// Solidity: function withdrawalBatchIndex(address submitter) view returns(uint256 batchIndex)
func (_Submitter *SubmitterCallerSession) WithdrawalBatchIndex(submitter common.Address) (*big.Int, error) {
	return _Submitter.Contract.WithdrawalBatchIndex(&_Submitter.CallOpts, submitter)
}

// Withdrawing is a free data retrieval call binding the contract method 0x3baf3a2f.
//
// Solidity: function withdrawing(address submitter) view returns(bool isWithdrawing)
func (_Submitter *SubmitterCaller) Withdrawing(opts *bind.CallOpts, submitter common.Address) (bool, error) {
	var out []interface{}
	err := _Submitter.contract.Call(opts, &out, "withdrawing", submitter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdrawing is a free data retrieval call binding the contract method 0x3baf3a2f.
//
// Solidity: function withdrawing(address submitter) view returns(bool isWithdrawing)
func (_Submitter *SubmitterSession) Withdrawing(submitter common.Address) (bool, error) {
	return _Submitter.Contract.Withdrawing(&_Submitter.CallOpts, submitter)
}

// Withdrawing is a free data retrieval call binding the contract method 0x3baf3a2f.
//
// Solidity: function withdrawing(address submitter) view returns(bool isWithdrawing)
func (_Submitter *SubmitterCallerSession) Withdrawing(submitter common.Address) (bool, error) {
	return _Submitter.Contract.Withdrawing(&_Submitter.CallOpts, submitter)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address submitter) returns()
func (_Submitter *SubmitterTransactor) AddSubmitter(opts *bind.TransactOpts, submitter common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "addSubmitter", submitter)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address submitter) returns()
func (_Submitter *SubmitterSession) AddSubmitter(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.AddSubmitter(&_Submitter.TransactOpts, submitter)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address submitter) returns()
func (_Submitter *SubmitterTransactorSession) AddSubmitter(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.AddSubmitter(&_Submitter.TransactOpts, submitter)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_Submitter *SubmitterTransactor) ClaimSlashRemaining(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "claimSlashRemaining", receiver)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_Submitter *SubmitterSession) ClaimSlashRemaining(receiver common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.ClaimSlashRemaining(&_Submitter.TransactOpts, receiver)
}

// ClaimSlashRemaining is a paid mutator transaction binding the contract method 0xcde4cd11.
//
// Solidity: function claimSlashRemaining(address receiver) returns()
func (_Submitter *SubmitterTransactorSession) ClaimSlashRemaining(receiver common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.ClaimSlashRemaining(&_Submitter.TransactOpts, receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_Submitter *SubmitterTransactor) ClaimWithdrawal(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "claimWithdrawal", receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_Submitter *SubmitterSession) ClaimWithdrawal(receiver common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.ClaimWithdrawal(&_Submitter.TransactOpts, receiver)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xa3066aab.
//
// Solidity: function claimWithdrawal(address receiver) returns()
func (_Submitter *SubmitterTransactorSession) ClaimWithdrawal(receiver common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.ClaimWithdrawal(&_Submitter.TransactOpts, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address owner_, address rollupContract_, uint256 minimumStake_, uint256 challengeDeposit_, uint256 rewardPercentage_) returns()
func (_Submitter *SubmitterTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, rollupContract_ common.Address, minimumStake_ *big.Int, challengeDeposit_ *big.Int, rewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "initialize", owner_, rollupContract_, minimumStake_, challengeDeposit_, rewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address owner_, address rollupContract_, uint256 minimumStake_, uint256 challengeDeposit_, uint256 rewardPercentage_) returns()
func (_Submitter *SubmitterSession) Initialize(owner_ common.Address, rollupContract_ common.Address, minimumStake_ *big.Int, challengeDeposit_ *big.Int, rewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts, owner_, rollupContract_, minimumStake_, challengeDeposit_, rewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address owner_, address rollupContract_, uint256 minimumStake_, uint256 challengeDeposit_, uint256 rewardPercentage_) returns()
func (_Submitter *SubmitterTransactorSession) Initialize(owner_ common.Address, rollupContract_ common.Address, minimumStake_ *big.Int, challengeDeposit_ *big.Int, rewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.Initialize(&_Submitter.TransactOpts, owner_, rollupContract_, minimumStake_, challengeDeposit_, rewardPercentage_)
}

// RemoveSubmitter is a paid mutator transaction binding the contract method 0xc36f41e2.
//
// Solidity: function removeSubmitter(address submitter) returns()
func (_Submitter *SubmitterTransactor) RemoveSubmitter(opts *bind.TransactOpts, submitter common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "removeSubmitter", submitter)
}

// RemoveSubmitter is a paid mutator transaction binding the contract method 0xc36f41e2.
//
// Solidity: function removeSubmitter(address submitter) returns()
func (_Submitter *SubmitterSession) RemoveSubmitter(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.RemoveSubmitter(&_Submitter.TransactOpts, submitter)
}

// RemoveSubmitter is a paid mutator transaction binding the contract method 0xc36f41e2.
//
// Solidity: function removeSubmitter(address submitter) returns()
func (_Submitter *SubmitterTransactorSession) RemoveSubmitter(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.RemoveSubmitter(&_Submitter.TransactOpts, submitter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Submitter.Contract.RenounceOwnership(&_Submitter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Submitter *SubmitterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Submitter.Contract.RenounceOwnership(&_Submitter.TransactOpts)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 value) returns()
func (_Submitter *SubmitterTransactor) SetMinimumStake(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "setMinimumStake", value)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 value) returns()
func (_Submitter *SubmitterSession) SetMinimumStake(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.SetMinimumStake(&_Submitter.TransactOpts, value)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 value) returns()
func (_Submitter *SubmitterTransactorSession) SetMinimumStake(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.SetMinimumStake(&_Submitter.TransactOpts, value)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address submitter) returns(uint256 reward)
func (_Submitter *SubmitterTransactor) Slash(opts *bind.TransactOpts, submitter common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "slash", submitter)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address submitter) returns(uint256 reward)
func (_Submitter *SubmitterSession) Slash(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.Slash(&_Submitter.TransactOpts, submitter)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address submitter) returns(uint256 reward)
func (_Submitter *SubmitterTransactorSession) Slash(submitter common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.Slash(&_Submitter.TransactOpts, submitter)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Submitter *SubmitterTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Submitter *SubmitterSession) Stake() (*types.Transaction, error) {
	return _Submitter.Contract.Stake(&_Submitter.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Submitter *SubmitterTransactorSession) Stake() (*types.Transaction, error) {
	return _Submitter.Contract.Stake(&_Submitter.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.TransferOwnership(&_Submitter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Submitter *SubmitterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Submitter.Contract.TransferOwnership(&_Submitter.TransactOpts, newOwner)
}

// UpdateChallengeDeposit is a paid mutator transaction binding the contract method 0x35928991.
//
// Solidity: function updateChallengeDeposit(uint256 value) returns()
func (_Submitter *SubmitterTransactor) UpdateChallengeDeposit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "updateChallengeDeposit", value)
}

// UpdateChallengeDeposit is a paid mutator transaction binding the contract method 0x35928991.
//
// Solidity: function updateChallengeDeposit(uint256 value) returns()
func (_Submitter *SubmitterSession) UpdateChallengeDeposit(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.UpdateChallengeDeposit(&_Submitter.TransactOpts, value)
}

// UpdateChallengeDeposit is a paid mutator transaction binding the contract method 0x35928991.
//
// Solidity: function updateChallengeDeposit(uint256 value) returns()
func (_Submitter *SubmitterTransactorSession) UpdateChallengeDeposit(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.UpdateChallengeDeposit(&_Submitter.TransactOpts, value)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 value) returns()
func (_Submitter *SubmitterTransactor) UpdateRewardPercentage(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "updateRewardPercentage", value)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 value) returns()
func (_Submitter *SubmitterSession) UpdateRewardPercentage(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.UpdateRewardPercentage(&_Submitter.TransactOpts, value)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 value) returns()
func (_Submitter *SubmitterTransactorSession) UpdateRewardPercentage(value *big.Int) (*types.Transaction, error) {
	return _Submitter.Contract.UpdateRewardPercentage(&_Submitter.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Submitter *SubmitterTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Submitter.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Submitter *SubmitterSession) Withdraw() (*types.Transaction, error) {
	return _Submitter.Contract.Withdraw(&_Submitter.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Submitter *SubmitterTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Submitter.Contract.Withdraw(&_Submitter.TransactOpts)
}

// SubmitterChallengeDepositUpdatedIterator is returned from FilterChallengeDepositUpdated and is used to iterate over the raw logs and unpacked data for ChallengeDepositUpdated events raised by the Submitter contract.
type SubmitterChallengeDepositUpdatedIterator struct {
	Event *SubmitterChallengeDepositUpdated // Event containing the contract specifics and raw log

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
func (it *SubmitterChallengeDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterChallengeDepositUpdated)
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
		it.Event = new(SubmitterChallengeDepositUpdated)
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
func (it *SubmitterChallengeDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterChallengeDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterChallengeDepositUpdated represents a ChallengeDepositUpdated event raised by the Submitter contract.
type SubmitterChallengeDepositUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChallengeDepositUpdated is a free log retrieval operation binding the contract event 0x36f971a40478225aeb80cfbf5e80306e8cb76d3bf7d56fdc5e490945cddb7d55.
//
// Solidity: event ChallengeDepositUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) FilterChallengeDepositUpdated(opts *bind.FilterOpts) (*SubmitterChallengeDepositUpdatedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "ChallengeDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &SubmitterChallengeDepositUpdatedIterator{contract: _Submitter.contract, event: "ChallengeDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchChallengeDepositUpdated is a free log subscription operation binding the contract event 0x36f971a40478225aeb80cfbf5e80306e8cb76d3bf7d56fdc5e490945cddb7d55.
//
// Solidity: event ChallengeDepositUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) WatchChallengeDepositUpdated(opts *bind.WatchOpts, sink chan<- *SubmitterChallengeDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "ChallengeDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterChallengeDepositUpdated)
				if err := _Submitter.contract.UnpackLog(event, "ChallengeDepositUpdated", log); err != nil {
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

// ParseChallengeDepositUpdated is a log parse operation binding the contract event 0x36f971a40478225aeb80cfbf5e80306e8cb76d3bf7d56fdc5e490945cddb7d55.
//
// Solidity: event ChallengeDepositUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) ParseChallengeDepositUpdated(log types.Log) (*SubmitterChallengeDepositUpdated, error) {
	event := new(SubmitterChallengeDepositUpdated)
	if err := _Submitter.contract.UnpackLog(event, "ChallengeDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Submitter contract.
type SubmitterInitializedIterator struct {
	Event *SubmitterInitialized // Event containing the contract specifics and raw log

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
func (it *SubmitterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterInitialized)
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
		it.Event = new(SubmitterInitialized)
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
func (it *SubmitterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterInitialized represents a Initialized event raised by the Submitter contract.
type SubmitterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubmitterInitializedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubmitterInitializedIterator{contract: _Submitter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Submitter *SubmitterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubmitterInitialized) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterInitialized)
				if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Submitter *SubmitterFilterer) ParseInitialized(log types.Log) (*SubmitterInitialized, error) {
	event := new(SubmitterInitialized)
	if err := _Submitter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterMinimumStakeUpdatedIterator is returned from FilterMinimumStakeUpdated and is used to iterate over the raw logs and unpacked data for MinimumStakeUpdated events raised by the Submitter contract.
type SubmitterMinimumStakeUpdatedIterator struct {
	Event *SubmitterMinimumStakeUpdated // Event containing the contract specifics and raw log

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
func (it *SubmitterMinimumStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterMinimumStakeUpdated)
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
		it.Event = new(SubmitterMinimumStakeUpdated)
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
func (it *SubmitterMinimumStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterMinimumStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterMinimumStakeUpdated represents a MinimumStakeUpdated event raised by the Submitter contract.
type SubmitterMinimumStakeUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinimumStakeUpdated is a free log retrieval operation binding the contract event 0xd67ed534faf3e0dcda08ac6043ec96dcf9b6ebec56055fd14bd0cb40a27c0e49.
//
// Solidity: event MinimumStakeUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) FilterMinimumStakeUpdated(opts *bind.FilterOpts) (*SubmitterMinimumStakeUpdatedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "MinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &SubmitterMinimumStakeUpdatedIterator{contract: _Submitter.contract, event: "MinimumStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumStakeUpdated is a free log subscription operation binding the contract event 0xd67ed534faf3e0dcda08ac6043ec96dcf9b6ebec56055fd14bd0cb40a27c0e49.
//
// Solidity: event MinimumStakeUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) WatchMinimumStakeUpdated(opts *bind.WatchOpts, sink chan<- *SubmitterMinimumStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "MinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterMinimumStakeUpdated)
				if err := _Submitter.contract.UnpackLog(event, "MinimumStakeUpdated", log); err != nil {
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

// ParseMinimumStakeUpdated is a log parse operation binding the contract event 0xd67ed534faf3e0dcda08ac6043ec96dcf9b6ebec56055fd14bd0cb40a27c0e49.
//
// Solidity: event MinimumStakeUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) ParseMinimumStakeUpdated(log types.Log) (*SubmitterMinimumStakeUpdated, error) {
	event := new(SubmitterMinimumStakeUpdated)
	if err := _Submitter.contract.UnpackLog(event, "MinimumStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Submitter contract.
type SubmitterOwnershipTransferredIterator struct {
	Event *SubmitterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubmitterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterOwnershipTransferred)
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
		it.Event = new(SubmitterOwnershipTransferred)
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
func (it *SubmitterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterOwnershipTransferred represents a OwnershipTransferred event raised by the Submitter contract.
type SubmitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Submitter *SubmitterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubmitterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterOwnershipTransferredIterator{contract: _Submitter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Submitter *SubmitterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubmitterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterOwnershipTransferred)
				if err := _Submitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Submitter *SubmitterFilterer) ParseOwnershipTransferred(log types.Log) (*SubmitterOwnershipTransferred, error) {
	event := new(SubmitterOwnershipTransferred)
	if err := _Submitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterRewardPercentageUpdatedIterator is returned from FilterRewardPercentageUpdated and is used to iterate over the raw logs and unpacked data for RewardPercentageUpdated events raised by the Submitter contract.
type SubmitterRewardPercentageUpdatedIterator struct {
	Event *SubmitterRewardPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *SubmitterRewardPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterRewardPercentageUpdated)
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
		it.Event = new(SubmitterRewardPercentageUpdated)
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
func (it *SubmitterRewardPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterRewardPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterRewardPercentageUpdated represents a RewardPercentageUpdated event raised by the Submitter contract.
type SubmitterRewardPercentageUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardPercentageUpdated is a free log retrieval operation binding the contract event 0xa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba.
//
// Solidity: event RewardPercentageUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) FilterRewardPercentageUpdated(opts *bind.FilterOpts) (*SubmitterRewardPercentageUpdatedIterator, error) {

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "RewardPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &SubmitterRewardPercentageUpdatedIterator{contract: _Submitter.contract, event: "RewardPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardPercentageUpdated is a free log subscription operation binding the contract event 0xa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba.
//
// Solidity: event RewardPercentageUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) WatchRewardPercentageUpdated(opts *bind.WatchOpts, sink chan<- *SubmitterRewardPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "RewardPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterRewardPercentageUpdated)
				if err := _Submitter.contract.UnpackLog(event, "RewardPercentageUpdated", log); err != nil {
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

// ParseRewardPercentageUpdated is a log parse operation binding the contract event 0xa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba.
//
// Solidity: event RewardPercentageUpdated(uint256 oldValue, uint256 newValue)
func (_Submitter *SubmitterFilterer) ParseRewardPercentageUpdated(log types.Log) (*SubmitterRewardPercentageUpdated, error) {
	event := new(SubmitterRewardPercentageUpdated)
	if err := _Submitter.contract.UnpackLog(event, "RewardPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterSlashRemainingClaimedIterator is returned from FilterSlashRemainingClaimed and is used to iterate over the raw logs and unpacked data for SlashRemainingClaimed events raised by the Submitter contract.
type SubmitterSlashRemainingClaimedIterator struct {
	Event *SubmitterSlashRemainingClaimed // Event containing the contract specifics and raw log

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
func (it *SubmitterSlashRemainingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterSlashRemainingClaimed)
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
		it.Event = new(SubmitterSlashRemainingClaimed)
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
func (it *SubmitterSlashRemainingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterSlashRemainingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterSlashRemainingClaimed represents a SlashRemainingClaimed event raised by the Submitter contract.
type SubmitterSlashRemainingClaimed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlashRemainingClaimed is a free log retrieval operation binding the contract event 0xa1fefb6c5328a92a416e321ed50997303fe7135fd88c28b0592b21ce42b5cdd9.
//
// Solidity: event SlashRemainingClaimed(address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) FilterSlashRemainingClaimed(opts *bind.FilterOpts, receiver []common.Address) (*SubmitterSlashRemainingClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "SlashRemainingClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterSlashRemainingClaimedIterator{contract: _Submitter.contract, event: "SlashRemainingClaimed", logs: logs, sub: sub}, nil
}

// WatchSlashRemainingClaimed is a free log subscription operation binding the contract event 0xa1fefb6c5328a92a416e321ed50997303fe7135fd88c28b0592b21ce42b5cdd9.
//
// Solidity: event SlashRemainingClaimed(address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) WatchSlashRemainingClaimed(opts *bind.WatchOpts, sink chan<- *SubmitterSlashRemainingClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "SlashRemainingClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterSlashRemainingClaimed)
				if err := _Submitter.contract.UnpackLog(event, "SlashRemainingClaimed", log); err != nil {
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

// ParseSlashRemainingClaimed is a log parse operation binding the contract event 0xa1fefb6c5328a92a416e321ed50997303fe7135fd88c28b0592b21ce42b5cdd9.
//
// Solidity: event SlashRemainingClaimed(address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) ParseSlashRemainingClaimed(log types.Log) (*SubmitterSlashRemainingClaimed, error) {
	event := new(SubmitterSlashRemainingClaimed)
	if err := _Submitter.contract.UnpackLog(event, "SlashRemainingClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Submitter contract.
type SubmitterSlashedIterator struct {
	Event *SubmitterSlashed // Event containing the contract specifics and raw log

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
func (it *SubmitterSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterSlashed)
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
		it.Event = new(SubmitterSlashed)
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
func (it *SubmitterSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterSlashed represents a Slashed event raised by the Submitter contract.
type SubmitterSlashed struct {
	Submitter common.Address
	Amount    *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x45a371af55b0726877a30f464edc14db5879ab096590bacce682cf6c18223596.
//
// Solidity: event Slashed(address indexed submitter, uint256 amount, uint256 reward)
func (_Submitter *SubmitterFilterer) FilterSlashed(opts *bind.FilterOpts, submitter []common.Address) (*SubmitterSlashedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "Slashed", submitterRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterSlashedIterator{contract: _Submitter.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x45a371af55b0726877a30f464edc14db5879ab096590bacce682cf6c18223596.
//
// Solidity: event Slashed(address indexed submitter, uint256 amount, uint256 reward)
func (_Submitter *SubmitterFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *SubmitterSlashed, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "Slashed", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterSlashed)
				if err := _Submitter.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x45a371af55b0726877a30f464edc14db5879ab096590bacce682cf6c18223596.
//
// Solidity: event Slashed(address indexed submitter, uint256 amount, uint256 reward)
func (_Submitter *SubmitterFilterer) ParseSlashed(log types.Log) (*SubmitterSlashed, error) {
	event := new(SubmitterSlashed)
	if err := _Submitter.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Submitter contract.
type SubmitterStakedIterator struct {
	Event *SubmitterStaked // Event containing the contract specifics and raw log

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
func (it *SubmitterStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterStaked)
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
		it.Event = new(SubmitterStaked)
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
func (it *SubmitterStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterStaked represents a Staked event raised by the Submitter contract.
type SubmitterStaked struct {
	Submitter common.Address
	Amount    *big.Int
	Total     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed submitter, uint256 amount, uint256 total)
func (_Submitter *SubmitterFilterer) FilterStaked(opts *bind.FilterOpts, submitter []common.Address) (*SubmitterStakedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "Staked", submitterRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterStakedIterator{contract: _Submitter.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed submitter, uint256 amount, uint256 total)
func (_Submitter *SubmitterFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *SubmitterStaked, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "Staked", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterStaked)
				if err := _Submitter.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed submitter, uint256 amount, uint256 total)
func (_Submitter *SubmitterFilterer) ParseStaked(log types.Log) (*SubmitterStaked, error) {
	event := new(SubmitterStaked)
	if err := _Submitter.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterSubmitterAddedIterator is returned from FilterSubmitterAdded and is used to iterate over the raw logs and unpacked data for SubmitterAdded events raised by the Submitter contract.
type SubmitterSubmitterAddedIterator struct {
	Event *SubmitterSubmitterAdded // Event containing the contract specifics and raw log

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
func (it *SubmitterSubmitterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterSubmitterAdded)
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
		it.Event = new(SubmitterSubmitterAdded)
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
func (it *SubmitterSubmitterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterSubmitterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterSubmitterAdded represents a SubmitterAdded event raised by the Submitter contract.
type SubmitterSubmitterAdded struct {
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitterAdded is a free log retrieval operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed submitter)
func (_Submitter *SubmitterFilterer) FilterSubmitterAdded(opts *bind.FilterOpts, submitter []common.Address) (*SubmitterSubmitterAddedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "SubmitterAdded", submitterRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterSubmitterAddedIterator{contract: _Submitter.contract, event: "SubmitterAdded", logs: logs, sub: sub}, nil
}

// WatchSubmitterAdded is a free log subscription operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed submitter)
func (_Submitter *SubmitterFilterer) WatchSubmitterAdded(opts *bind.WatchOpts, sink chan<- *SubmitterSubmitterAdded, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "SubmitterAdded", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterSubmitterAdded)
				if err := _Submitter.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
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

// ParseSubmitterAdded is a log parse operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed submitter)
func (_Submitter *SubmitterFilterer) ParseSubmitterAdded(log types.Log) (*SubmitterSubmitterAdded, error) {
	event := new(SubmitterSubmitterAdded)
	if err := _Submitter.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterSubmitterRemovedIterator is returned from FilterSubmitterRemoved and is used to iterate over the raw logs and unpacked data for SubmitterRemoved events raised by the Submitter contract.
type SubmitterSubmitterRemovedIterator struct {
	Event *SubmitterSubmitterRemoved // Event containing the contract specifics and raw log

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
func (it *SubmitterSubmitterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterSubmitterRemoved)
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
		it.Event = new(SubmitterSubmitterRemoved)
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
func (it *SubmitterSubmitterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterSubmitterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterSubmitterRemoved represents a SubmitterRemoved event raised by the Submitter contract.
type SubmitterSubmitterRemoved struct {
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitterRemoved is a free log retrieval operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed submitter)
func (_Submitter *SubmitterFilterer) FilterSubmitterRemoved(opts *bind.FilterOpts, submitter []common.Address) (*SubmitterSubmitterRemovedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "SubmitterRemoved", submitterRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterSubmitterRemovedIterator{contract: _Submitter.contract, event: "SubmitterRemoved", logs: logs, sub: sub}, nil
}

// WatchSubmitterRemoved is a free log subscription operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed submitter)
func (_Submitter *SubmitterFilterer) WatchSubmitterRemoved(opts *bind.WatchOpts, sink chan<- *SubmitterSubmitterRemoved, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "SubmitterRemoved", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterSubmitterRemoved)
				if err := _Submitter.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
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

// ParseSubmitterRemoved is a log parse operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed submitter)
func (_Submitter *SubmitterFilterer) ParseSubmitterRemoved(log types.Log) (*SubmitterSubmitterRemoved, error) {
	event := new(SubmitterSubmitterRemoved)
	if err := _Submitter.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterWithdrawalClaimedIterator is returned from FilterWithdrawalClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalClaimed events raised by the Submitter contract.
type SubmitterWithdrawalClaimedIterator struct {
	Event *SubmitterWithdrawalClaimed // Event containing the contract specifics and raw log

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
func (it *SubmitterWithdrawalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterWithdrawalClaimed)
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
		it.Event = new(SubmitterWithdrawalClaimed)
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
func (it *SubmitterWithdrawalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterWithdrawalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterWithdrawalClaimed represents a WithdrawalClaimed event raised by the Submitter contract.
type SubmitterWithdrawalClaimed struct {
	Submitter common.Address
	Receiver  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalClaimed is a free log retrieval operation binding the contract event 0x8188e2b4d95f73db30690b4103c71159349bb897df928902c6330ef99e45fef3.
//
// Solidity: event WithdrawalClaimed(address indexed submitter, address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) FilterWithdrawalClaimed(opts *bind.FilterOpts, submitter []common.Address, receiver []common.Address) (*SubmitterWithdrawalClaimedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "WithdrawalClaimed", submitterRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterWithdrawalClaimedIterator{contract: _Submitter.contract, event: "WithdrawalClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalClaimed is a free log subscription operation binding the contract event 0x8188e2b4d95f73db30690b4103c71159349bb897df928902c6330ef99e45fef3.
//
// Solidity: event WithdrawalClaimed(address indexed submitter, address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) WatchWithdrawalClaimed(opts *bind.WatchOpts, sink chan<- *SubmitterWithdrawalClaimed, submitter []common.Address, receiver []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "WithdrawalClaimed", submitterRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterWithdrawalClaimed)
				if err := _Submitter.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
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

// ParseWithdrawalClaimed is a log parse operation binding the contract event 0x8188e2b4d95f73db30690b4103c71159349bb897df928902c6330ef99e45fef3.
//
// Solidity: event WithdrawalClaimed(address indexed submitter, address indexed receiver, uint256 amount)
func (_Submitter *SubmitterFilterer) ParseWithdrawalClaimed(log types.Log) (*SubmitterWithdrawalClaimed, error) {
	event := new(SubmitterWithdrawalClaimed)
	if err := _Submitter.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubmitterWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the Submitter contract.
type SubmitterWithdrawalRequestedIterator struct {
	Event *SubmitterWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *SubmitterWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterWithdrawalRequested)
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
		it.Event = new(SubmitterWithdrawalRequested)
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
func (it *SubmitterWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterWithdrawalRequested represents a WithdrawalRequested event raised by the Submitter contract.
type SubmitterWithdrawalRequested struct {
	Submitter common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed submitter, uint256 amount)
func (_Submitter *SubmitterFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, submitter []common.Address) (*SubmitterWithdrawalRequestedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.FilterLogs(opts, "WithdrawalRequested", submitterRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterWithdrawalRequestedIterator{contract: _Submitter.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed submitter, uint256 amount)
func (_Submitter *SubmitterFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *SubmitterWithdrawalRequested, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Submitter.contract.WatchLogs(opts, "WithdrawalRequested", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterWithdrawalRequested)
				if err := _Submitter.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed submitter, uint256 amount)
func (_Submitter *SubmitterFilterer) ParseWithdrawalRequested(log types.Log) (*SubmitterWithdrawalRequested, error) {
	event := new(SubmitterWithdrawalRequested)
	if err := _Submitter.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
