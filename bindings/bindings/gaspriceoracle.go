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

// GasPriceOracleMetaData contains all meta data concerning the GasPriceOracle contract.
var GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrAlreadyInCurieFork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrDifferentLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedMaxBlobScalar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedMaxCommitScalar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedMaxOverhead\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedMaxScalar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrSettintSameValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"AllowListAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"BlobScalarUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"CommitScalarUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"L1BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BlobBaseFee\",\"type\":\"uint256\"}],\"name\":\"L1BlobBaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OverheadUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ScalarUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableCurie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1Fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1GasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCurie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BlobBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowListEnabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setBlobScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setCommitScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l1BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL1BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l1BaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l1BlobBaseFee\",\"type\":\"uint256\"}],\"name\":\"setL1BaseFeeAndBlobBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"}],\"name\":\"setOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620014ed380380620014ed83398101604081905262000033916200018c565b6200003e336200005d565b6200004981620000ac565b506004805460ff19166001179055620001bb565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b620000b66200012f565b6001600160a01b038116620001215760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6200012c816200005d565b50565b5f546001600160a01b031633146200018a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000118565b565b5f602082840312156200019d575f80fd5b81516001600160a01b0381168114620001b4575f80fd5b9392505050565b61132480620001c95f395ff3fe608060405234801561000f575f80fd5b5060043610610184575f3560e01c806384189161116100dd578063de26c4a111610088578063efeadb6d11610063578063efeadb6d146102f8578063f2fde38b1461030b578063f45e65d81461031e575f80fd5b8063de26c4a1146102bf578063e3de72a5146102d2578063e88a60ad146102e5575f80fd5b8063a911d77f116100b8578063a911d77f14610282578063babcc5391461028a578063bede39b5146102ac575f80fd5b8063841891611461023f5780638da5cb5b14610248578063944b247f1461026f575f80fd5b806339455d3a1161013d5780636a5e67e5116101185780636a5e67e51461021b5780637046559714610224578063715018a614610237575f80fd5b806339455d3a146101ec57806349948e0e146101ff578063519b4bd314610212575f80fd5b806322bd5c1c1161016d57806322bd5c1c146101c157806323e524ac146101ce5780633577afc5146101d7575f80fd5b80630c18c1621461018857806313dad5be146101a4575b5f80fd5b61019160025481565b6040519081526020015b60405180910390f35b6009546101b19060ff1681565b604051901515815260200161019b565b6004546101b19060ff1681565b61019160075481565b6101ea6101e5366004610f40565b610327565b005b6101ea6101fa366004610f57565b610432565b61019161020d366004610ff3565b610538565b61019160015481565b61019160085481565b6101ea610232366004610f40565b610562565b6101ea610672565b61019160065481565b5f5460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101ea61027d366004610f40565b610685565b6101ea610793565b6101b16102983660046110c3565b60056020525f908152604090205460ff1681565b6101ea6102ba366004610f40565b61088f565b6101916102cd366004610ff3565b610956565b6101ea6102e0366004611188565b610973565b6101ea6102f3366004610f40565b610ad8565b6101ea610306366004611242565b610be6565b6101ea6103193660046110c3565b610c8f565b61019160035481565b336103465f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610383575060045460ff1680156103835750335f9081526005602052604090205460ff165b6103b9576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b621c9c388111156103f6576040517fae85900a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b336104515f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16148061048e575060045460ff16801561048e5750335f9081526005602052604090205460ff165b6104c4576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600182905560068190556040518281527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c449060200160405180910390a16040518181527f9a14bfb5d18c4c3cf14cae19c23d7cf1bcede357ea40ca1f75cd49542c71c2149060200160405180910390a15050565b6009545f9060ff16156105545761054e82610d4b565b92915050565b61054e82610d91565b919050565b336105815f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806105be575060045460ff1680156105be5750335f9081526005602052604090205460ff165b6105f4576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610604633b9aca006103e8611288565b81111561063d576040517f3c89fbd600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60038190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a90602001610427565b61067a610dd4565b6106835f610e54565b565b336106a45f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806106e1575060045460ff1680156106e15750335f9081526005602052604090205460ff165b610717576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610725633b9aca0080611288565b81111561075e576040517f874f603100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60078190556040518181527f2ab3f5a4ebbcbf3c24f62f5454f52f10e1a8c9dcc5acac8f19199ce881a6a10890602001610427565b336107b25f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806107ef575060045460ff1680156107ef5750335f9081526005602052604090205460ff165b610825576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095460ff1615610862576040517f79f9c57500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b336108ae5f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806108eb575060045460ff1680156108eb5750335f9081526005602052604090205460ff165b610921576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c4490602001610427565b6009545f9060ff161561096a57505f919050565b61054e82610ec8565b61097b610dd4565b80518251146109b6576040517f1b9c61c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610ad3578181815181106109d3576109d361129f565b602002602001015160055f8584815181106109f0576109f061129f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550828181518110610a5957610a5961129f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a838381518110610aaa57610aaa61129f565b6020026020010151604051610ac3911515815260200190565b60405180910390a26001016109b8565b505050565b33610af75f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610b34575060045460ff168015610b345750335f9081526005602052604090205460ff165b610b6a576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b78633b9aca0080611288565b811115610bb1576040517ff37ec21500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60088190556040518181527f6b332a036d8c3ead57dcb06c87243bd7a2aed015ddf2d0528c2501dae56331aa90602001610427565b610bee610dd4565b60045460ff16151581151503610c30576040517fd5d1b79c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb390602001610427565b610c97610dd4565b73ffffffffffffffffffffffffffffffffffffffff8116610d3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610d4881610e54565b50565b5f633b9aca006006548351600854610d639190611288565b610d6d9190611288565b600154600754610d7d9190611288565b610d8791906112cc565b61054e91906112df565b5f80610d9c83610ec8565b90505f60015482610dad9190611288565b9050633b9aca0060035482610dc29190611288565b610dcc91906112df565b949350505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610d36565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80515f908190815b81811015610f3157848181518110610eea57610eea61129f565b01602001517fff00000000000000000000000000000000000000000000000000000000000000165f03610f2257600483019250610f29565b6010830192505b600101610ed0565b50506002540160400192915050565b5f60208284031215610f50575f80fd5b5035919050565b5f8060408385031215610f68575f80fd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610feb57610feb610f77565b604052919050565b5f6020808385031215611004575f80fd5b823567ffffffffffffffff8082111561101b575f80fd5b818501915085601f83011261102e575f80fd5b81358181111561104057611040610f77565b611070847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610fa4565b91508082528684828501011115611085575f80fd5b80848401858401375f90820190930192909252509392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461055d575f80fd5b5f602082840312156110d3575f80fd5b6110dc826110a0565b9392505050565b5f67ffffffffffffffff8211156110fc576110fc610f77565b5060051b60200190565b8035801515811461055d575f80fd5b5f82601f830112611124575f80fd5b81356020611139611134836110e3565b610fa4565b8083825260208201915060208460051b87010193508684111561115a575f80fd5b602086015b8481101561117d5761117081611106565b835291830191830161115f565b509695505050505050565b5f8060408385031215611199575f80fd5b823567ffffffffffffffff808211156111b0575f80fd5b818501915085601f8301126111c3575f80fd5b813560206111d3611134836110e3565b82815260059290921b840181019181810190898411156111f1575f80fd5b948201945b8386101561121657611207866110a0565b825294820194908201906111f6565b9650508601359250508082111561122b575f80fd5b5061123885828601611115565b9150509250929050565b5f60208284031215611252575f80fd5b6110dc82611106565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808202811582820484141761054e5761054e61125b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082018082111561054e5761054e61125b565b5f82611312577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a",
}

// GasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use GasPriceOracleMetaData.ABI instead.
var GasPriceOracleABI = GasPriceOracleMetaData.ABI

// GasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasPriceOracleMetaData.Bin instead.
var GasPriceOracleBin = GasPriceOracleMetaData.Bin

// DeployGasPriceOracle deploys a new Ethereum contract, binding an instance of GasPriceOracle to it.
func DeployGasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address) (common.Address, *types.Transaction, *GasPriceOracle, error) {
	parsed, err := GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasPriceOracleBin), backend, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasPriceOracle{GasPriceOracleCaller: GasPriceOracleCaller{contract: contract}, GasPriceOracleTransactor: GasPriceOracleTransactor{contract: contract}, GasPriceOracleFilterer: GasPriceOracleFilterer{contract: contract}}, nil
}

// GasPriceOracle is an auto generated Go binding around an Ethereum contract.
type GasPriceOracle struct {
	GasPriceOracleCaller     // Read-only binding to the contract
	GasPriceOracleTransactor // Write-only binding to the contract
	GasPriceOracleFilterer   // Log filterer for contract events
}

// GasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasPriceOracleSession struct {
	Contract     *GasPriceOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasPriceOracleCallerSession struct {
	Contract *GasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasPriceOracleTransactorSession struct {
	Contract     *GasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasPriceOracleRaw struct {
	Contract *GasPriceOracle // Generic contract binding to access the raw methods on
}

// GasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasPriceOracleCallerRaw struct {
	Contract *GasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// GasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasPriceOracleTransactorRaw struct {
	Contract *GasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasPriceOracle creates a new instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracle(address common.Address, backend bind.ContractBackend) (*GasPriceOracle, error) {
	contract, err := bindGasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracle{GasPriceOracleCaller: GasPriceOracleCaller{contract: contract}, GasPriceOracleTransactor: GasPriceOracleTransactor{contract: contract}, GasPriceOracleFilterer: GasPriceOracleFilterer{contract: contract}}, nil
}

// NewGasPriceOracleCaller creates a new read-only instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*GasPriceOracleCaller, error) {
	contract, err := bindGasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleCaller{contract: contract}, nil
}

// NewGasPriceOracleTransactor creates a new write-only instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*GasPriceOracleTransactor, error) {
	contract, err := bindGasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleTransactor{contract: contract}, nil
}

// NewGasPriceOracleFilterer creates a new log filterer instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*GasPriceOracleFilterer, error) {
	contract, err := bindGasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleFilterer{contract: contract}, nil
}

// bindGasPriceOracle binds a generic wrapper to an already deployed contract.
func bindGasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceOracle *GasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPriceOracle.Contract.GasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceOracle *GasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.GasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceOracle *GasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.GasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceOracle *GasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceOracle *GasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceOracle *GasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_GasPriceOracle *GasPriceOracleSession) AllowListEnabled() (bool, error) {
	return _GasPriceOracle.Contract.AllowListEnabled(&_GasPriceOracle.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCallerSession) AllowListEnabled() (bool, error) {
	return _GasPriceOracle.Contract.AllowListEnabled(&_GasPriceOracle.CallOpts)
}

// BlobScalar is a free data retrieval call binding the contract method 0x6a5e67e5.
//
// Solidity: function blobScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) BlobScalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "blobScalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlobScalar is a free data retrieval call binding the contract method 0x6a5e67e5.
//
// Solidity: function blobScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) BlobScalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.BlobScalar(&_GasPriceOracle.CallOpts)
}

// BlobScalar is a free data retrieval call binding the contract method 0x6a5e67e5.
//
// Solidity: function blobScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) BlobScalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.BlobScalar(&_GasPriceOracle.CallOpts)
}

// CommitScalar is a free data retrieval call binding the contract method 0x23e524ac.
//
// Solidity: function commitScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) CommitScalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "commitScalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitScalar is a free data retrieval call binding the contract method 0x23e524ac.
//
// Solidity: function commitScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) CommitScalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.CommitScalar(&_GasPriceOracle.CallOpts)
}

// CommitScalar is a free data retrieval call binding the contract method 0x23e524ac.
//
// Solidity: function commitScalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) CommitScalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.CommitScalar(&_GasPriceOracle.CallOpts)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) GetL1Fee(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getL1Fee", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetL1Fee(&_GasPriceOracle.CallOpts, _data)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetL1Fee(&_GasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) GetL1GasUsed(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getL1GasUsed", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetL1GasUsed(&_GasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetL1GasUsed(&_GasPriceOracle.CallOpts, _data)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address addr) view returns(bool allowed)
func (_GasPriceOracle *GasPriceOracleCaller) IsAllowed(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "isAllowed", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address addr) view returns(bool allowed)
func (_GasPriceOracle *GasPriceOracleSession) IsAllowed(addr common.Address) (bool, error) {
	return _GasPriceOracle.Contract.IsAllowed(&_GasPriceOracle.CallOpts, addr)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address addr) view returns(bool allowed)
func (_GasPriceOracle *GasPriceOracleCallerSession) IsAllowed(addr common.Address) (bool, error) {
	return _GasPriceOracle.Contract.IsAllowed(&_GasPriceOracle.CallOpts, addr)
}

// IsCurie is a free data retrieval call binding the contract method 0x13dad5be.
//
// Solidity: function isCurie() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCaller) IsCurie(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "isCurie")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurie is a free data retrieval call binding the contract method 0x13dad5be.
//
// Solidity: function isCurie() view returns(bool)
func (_GasPriceOracle *GasPriceOracleSession) IsCurie() (bool, error) {
	return _GasPriceOracle.Contract.IsCurie(&_GasPriceOracle.CallOpts)
}

// IsCurie is a free data retrieval call binding the contract method 0x13dad5be.
//
// Solidity: function isCurie() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCallerSession) IsCurie() (bool, error) {
	return _GasPriceOracle.Contract.IsCurie(&_GasPriceOracle.CallOpts)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) L1BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "l1BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) L1BaseFee() (*big.Int, error) {
	return _GasPriceOracle.Contract.L1BaseFee(&_GasPriceOracle.CallOpts)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) L1BaseFee() (*big.Int, error) {
	return _GasPriceOracle.Contract.L1BaseFee(&_GasPriceOracle.CallOpts)
}

// L1BlobBaseFee is a free data retrieval call binding the contract method 0x84189161.
//
// Solidity: function l1BlobBaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) L1BlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "l1BlobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1BlobBaseFee is a free data retrieval call binding the contract method 0x84189161.
//
// Solidity: function l1BlobBaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) L1BlobBaseFee() (*big.Int, error) {
	return _GasPriceOracle.Contract.L1BlobBaseFee(&_GasPriceOracle.CallOpts)
}

// L1BlobBaseFee is a free data retrieval call binding the contract method 0x84189161.
//
// Solidity: function l1BlobBaseFee() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) L1BlobBaseFee() (*big.Int, error) {
	return _GasPriceOracle.Contract.L1BlobBaseFee(&_GasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) Overhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "overhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) Overhead() (*big.Int, error) {
	return _GasPriceOracle.Contract.Overhead(&_GasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) Overhead() (*big.Int, error) {
	return _GasPriceOracle.Contract.Overhead(&_GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleSession) Owner() (common.Address, error) {
	return _GasPriceOracle.Contract.Owner(&_GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _GasPriceOracle.Contract.Owner(&_GasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) Scalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "scalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) Scalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.Scalar(&_GasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) Scalar() (*big.Int, error) {
	return _GasPriceOracle.Contract.Scalar(&_GasPriceOracle.CallOpts)
}

// EnableCurie is a paid mutator transaction binding the contract method 0xa911d77f.
//
// Solidity: function enableCurie() returns()
func (_GasPriceOracle *GasPriceOracleTransactor) EnableCurie(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "enableCurie")
}

// EnableCurie is a paid mutator transaction binding the contract method 0xa911d77f.
//
// Solidity: function enableCurie() returns()
func (_GasPriceOracle *GasPriceOracleSession) EnableCurie() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.EnableCurie(&_GasPriceOracle.TransactOpts)
}

// EnableCurie is a paid mutator transaction binding the contract method 0xa911d77f.
//
// Solidity: function enableCurie() returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) EnableCurie() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.EnableCurie(&_GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RenounceOwnership(&_GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RenounceOwnership(&_GasPriceOracle.TransactOpts)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setAllowList", user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetAllowList(&_GasPriceOracle.TransactOpts, user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetAllowList(&_GasPriceOracle.TransactOpts, user, val)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setAllowListEnabled", _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetAllowListEnabled(&_GasPriceOracle.TransactOpts, _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetAllowListEnabled(&_GasPriceOracle.TransactOpts, _allowListEnabled)
}

// SetBlobScalar is a paid mutator transaction binding the contract method 0xe88a60ad.
//
// Solidity: function setBlobScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetBlobScalar(opts *bind.TransactOpts, _scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setBlobScalar", _scalar)
}

// SetBlobScalar is a paid mutator transaction binding the contract method 0xe88a60ad.
//
// Solidity: function setBlobScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetBlobScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetBlobScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// SetBlobScalar is a paid mutator transaction binding the contract method 0xe88a60ad.
//
// Solidity: function setBlobScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetBlobScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetBlobScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// SetCommitScalar is a paid mutator transaction binding the contract method 0x944b247f.
//
// Solidity: function setCommitScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetCommitScalar(opts *bind.TransactOpts, _scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setCommitScalar", _scalar)
}

// SetCommitScalar is a paid mutator transaction binding the contract method 0x944b247f.
//
// Solidity: function setCommitScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetCommitScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetCommitScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// SetCommitScalar is a paid mutator transaction binding the contract method 0x944b247f.
//
// Solidity: function setCommitScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetCommitScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetCommitScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetL1BaseFee(opts *bind.TransactOpts, _l1BaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setL1BaseFee", _l1BaseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetL1BaseFee(_l1BaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetL1BaseFee(&_GasPriceOracle.TransactOpts, _l1BaseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetL1BaseFee(_l1BaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetL1BaseFee(&_GasPriceOracle.TransactOpts, _l1BaseFee)
}

// SetL1BaseFeeAndBlobBaseFee is a paid mutator transaction binding the contract method 0x39455d3a.
//
// Solidity: function setL1BaseFeeAndBlobBaseFee(uint256 _l1BaseFee, uint256 _l1BlobBaseFee) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetL1BaseFeeAndBlobBaseFee(opts *bind.TransactOpts, _l1BaseFee *big.Int, _l1BlobBaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setL1BaseFeeAndBlobBaseFee", _l1BaseFee, _l1BlobBaseFee)
}

// SetL1BaseFeeAndBlobBaseFee is a paid mutator transaction binding the contract method 0x39455d3a.
//
// Solidity: function setL1BaseFeeAndBlobBaseFee(uint256 _l1BaseFee, uint256 _l1BlobBaseFee) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetL1BaseFeeAndBlobBaseFee(_l1BaseFee *big.Int, _l1BlobBaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetL1BaseFeeAndBlobBaseFee(&_GasPriceOracle.TransactOpts, _l1BaseFee, _l1BlobBaseFee)
}

// SetL1BaseFeeAndBlobBaseFee is a paid mutator transaction binding the contract method 0x39455d3a.
//
// Solidity: function setL1BaseFeeAndBlobBaseFee(uint256 _l1BaseFee, uint256 _l1BlobBaseFee) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetL1BaseFeeAndBlobBaseFee(_l1BaseFee *big.Int, _l1BlobBaseFee *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetL1BaseFeeAndBlobBaseFee(&_GasPriceOracle.TransactOpts, _l1BaseFee, _l1BlobBaseFee)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetOverhead(opts *bind.TransactOpts, _overhead *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setOverhead", _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetOverhead(&_GasPriceOracle.TransactOpts, _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetOverhead(&_GasPriceOracle.TransactOpts, _overhead)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetScalar(opts *bind.TransactOpts, _scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setScalar", _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetScalar(&_GasPriceOracle.TransactOpts, _scalar)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.TransferOwnership(&_GasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.TransferOwnership(&_GasPriceOracle.TransactOpts, newOwner)
}

// GasPriceOracleAllowListAddressSetIterator is returned from FilterAllowListAddressSet and is used to iterate over the raw logs and unpacked data for AllowListAddressSet events raised by the GasPriceOracle contract.
type GasPriceOracleAllowListAddressSetIterator struct {
	Event *GasPriceOracleAllowListAddressSet // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleAllowListAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleAllowListAddressSet)
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
		it.Event = new(GasPriceOracleAllowListAddressSet)
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
func (it *GasPriceOracleAllowListAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleAllowListAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleAllowListAddressSet represents a AllowListAddressSet event raised by the GasPriceOracle contract.
type GasPriceOracleAllowListAddressSet struct {
	User common.Address
	Val  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowListAddressSet is a free log retrieval operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterAllowListAddressSet(opts *bind.FilterOpts, user []common.Address) (*GasPriceOracleAllowListAddressSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleAllowListAddressSetIterator{contract: _GasPriceOracle.contract, event: "AllowListAddressSet", logs: logs, sub: sub}, nil
}

// WatchAllowListAddressSet is a free log subscription operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchAllowListAddressSet(opts *bind.WatchOpts, sink chan<- *GasPriceOracleAllowListAddressSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleAllowListAddressSet)
				if err := _GasPriceOracle.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
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

// ParseAllowListAddressSet is a log parse operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseAllowListAddressSet(log types.Log) (*GasPriceOracleAllowListAddressSet, error) {
	event := new(GasPriceOracleAllowListAddressSet)
	if err := _GasPriceOracle.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleAllowListEnabledUpdatedIterator is returned from FilterAllowListEnabledUpdated and is used to iterate over the raw logs and unpacked data for AllowListEnabledUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleAllowListEnabledUpdatedIterator struct {
	Event *GasPriceOracleAllowListEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleAllowListEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleAllowListEnabledUpdated)
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
		it.Event = new(GasPriceOracleAllowListEnabledUpdated)
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
func (it *GasPriceOracleAllowListEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleAllowListEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleAllowListEnabledUpdated represents a AllowListEnabledUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleAllowListEnabledUpdated struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowListEnabledUpdated is a free log retrieval operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterAllowListEnabledUpdated(opts *bind.FilterOpts) (*GasPriceOracleAllowListEnabledUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleAllowListEnabledUpdatedIterator{contract: _GasPriceOracle.contract, event: "AllowListEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListEnabledUpdated is a free log subscription operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchAllowListEnabledUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleAllowListEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleAllowListEnabledUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
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

// ParseAllowListEnabledUpdated is a log parse operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseAllowListEnabledUpdated(log types.Log) (*GasPriceOracleAllowListEnabledUpdated, error) {
	event := new(GasPriceOracleAllowListEnabledUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleBlobScalarUpdatedIterator is returned from FilterBlobScalarUpdated and is used to iterate over the raw logs and unpacked data for BlobScalarUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleBlobScalarUpdatedIterator struct {
	Event *GasPriceOracleBlobScalarUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleBlobScalarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleBlobScalarUpdated)
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
		it.Event = new(GasPriceOracleBlobScalarUpdated)
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
func (it *GasPriceOracleBlobScalarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleBlobScalarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleBlobScalarUpdated represents a BlobScalarUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleBlobScalarUpdated struct {
	Scalar *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlobScalarUpdated is a free log retrieval operation binding the contract event 0x6b332a036d8c3ead57dcb06c87243bd7a2aed015ddf2d0528c2501dae56331aa.
//
// Solidity: event BlobScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterBlobScalarUpdated(opts *bind.FilterOpts) (*GasPriceOracleBlobScalarUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "BlobScalarUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleBlobScalarUpdatedIterator{contract: _GasPriceOracle.contract, event: "BlobScalarUpdated", logs: logs, sub: sub}, nil
}

// WatchBlobScalarUpdated is a free log subscription operation binding the contract event 0x6b332a036d8c3ead57dcb06c87243bd7a2aed015ddf2d0528c2501dae56331aa.
//
// Solidity: event BlobScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchBlobScalarUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleBlobScalarUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "BlobScalarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleBlobScalarUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "BlobScalarUpdated", log); err != nil {
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

// ParseBlobScalarUpdated is a log parse operation binding the contract event 0x6b332a036d8c3ead57dcb06c87243bd7a2aed015ddf2d0528c2501dae56331aa.
//
// Solidity: event BlobScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseBlobScalarUpdated(log types.Log) (*GasPriceOracleBlobScalarUpdated, error) {
	event := new(GasPriceOracleBlobScalarUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "BlobScalarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleCommitScalarUpdatedIterator is returned from FilterCommitScalarUpdated and is used to iterate over the raw logs and unpacked data for CommitScalarUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleCommitScalarUpdatedIterator struct {
	Event *GasPriceOracleCommitScalarUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleCommitScalarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleCommitScalarUpdated)
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
		it.Event = new(GasPriceOracleCommitScalarUpdated)
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
func (it *GasPriceOracleCommitScalarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleCommitScalarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleCommitScalarUpdated represents a CommitScalarUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleCommitScalarUpdated struct {
	Scalar *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitScalarUpdated is a free log retrieval operation binding the contract event 0x2ab3f5a4ebbcbf3c24f62f5454f52f10e1a8c9dcc5acac8f19199ce881a6a108.
//
// Solidity: event CommitScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterCommitScalarUpdated(opts *bind.FilterOpts) (*GasPriceOracleCommitScalarUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "CommitScalarUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleCommitScalarUpdatedIterator{contract: _GasPriceOracle.contract, event: "CommitScalarUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitScalarUpdated is a free log subscription operation binding the contract event 0x2ab3f5a4ebbcbf3c24f62f5454f52f10e1a8c9dcc5acac8f19199ce881a6a108.
//
// Solidity: event CommitScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchCommitScalarUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleCommitScalarUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "CommitScalarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleCommitScalarUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "CommitScalarUpdated", log); err != nil {
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

// ParseCommitScalarUpdated is a log parse operation binding the contract event 0x2ab3f5a4ebbcbf3c24f62f5454f52f10e1a8c9dcc5acac8f19199ce881a6a108.
//
// Solidity: event CommitScalarUpdated(uint256 scalar)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseCommitScalarUpdated(log types.Log) (*GasPriceOracleCommitScalarUpdated, error) {
	event := new(GasPriceOracleCommitScalarUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "CommitScalarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleL1BaseFeeUpdatedIterator is returned from FilterL1BaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L1BaseFeeUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleL1BaseFeeUpdatedIterator struct {
	Event *GasPriceOracleL1BaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleL1BaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleL1BaseFeeUpdated)
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
		it.Event = new(GasPriceOracleL1BaseFeeUpdated)
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
func (it *GasPriceOracleL1BaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleL1BaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleL1BaseFeeUpdated represents a L1BaseFeeUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleL1BaseFeeUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterL1BaseFeeUpdated is a free log retrieval operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterL1BaseFeeUpdated(opts *bind.FilterOpts) (*GasPriceOracleL1BaseFeeUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleL1BaseFeeUpdatedIterator{contract: _GasPriceOracle.contract, event: "L1BaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL1BaseFeeUpdated is a free log subscription operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchL1BaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleL1BaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleL1BaseFeeUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
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

// ParseL1BaseFeeUpdated is a log parse operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseL1BaseFeeUpdated(log types.Log) (*GasPriceOracleL1BaseFeeUpdated, error) {
	event := new(GasPriceOracleL1BaseFeeUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleL1BlobBaseFeeUpdatedIterator is returned from FilterL1BlobBaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L1BlobBaseFeeUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleL1BlobBaseFeeUpdatedIterator struct {
	Event *GasPriceOracleL1BlobBaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleL1BlobBaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleL1BlobBaseFeeUpdated)
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
		it.Event = new(GasPriceOracleL1BlobBaseFeeUpdated)
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
func (it *GasPriceOracleL1BlobBaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleL1BlobBaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleL1BlobBaseFeeUpdated represents a L1BlobBaseFeeUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleL1BlobBaseFeeUpdated struct {
	L1BlobBaseFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL1BlobBaseFeeUpdated is a free log retrieval operation binding the contract event 0x9a14bfb5d18c4c3cf14cae19c23d7cf1bcede357ea40ca1f75cd49542c71c214.
//
// Solidity: event L1BlobBaseFeeUpdated(uint256 l1BlobBaseFee)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterL1BlobBaseFeeUpdated(opts *bind.FilterOpts) (*GasPriceOracleL1BlobBaseFeeUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "L1BlobBaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleL1BlobBaseFeeUpdatedIterator{contract: _GasPriceOracle.contract, event: "L1BlobBaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL1BlobBaseFeeUpdated is a free log subscription operation binding the contract event 0x9a14bfb5d18c4c3cf14cae19c23d7cf1bcede357ea40ca1f75cd49542c71c214.
//
// Solidity: event L1BlobBaseFeeUpdated(uint256 l1BlobBaseFee)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchL1BlobBaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleL1BlobBaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "L1BlobBaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleL1BlobBaseFeeUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "L1BlobBaseFeeUpdated", log); err != nil {
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

// ParseL1BlobBaseFeeUpdated is a log parse operation binding the contract event 0x9a14bfb5d18c4c3cf14cae19c23d7cf1bcede357ea40ca1f75cd49542c71c214.
//
// Solidity: event L1BlobBaseFeeUpdated(uint256 l1BlobBaseFee)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseL1BlobBaseFeeUpdated(log types.Log) (*GasPriceOracleL1BlobBaseFeeUpdated, error) {
	event := new(GasPriceOracleL1BlobBaseFeeUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "L1BlobBaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleOverheadUpdatedIterator is returned from FilterOverheadUpdated and is used to iterate over the raw logs and unpacked data for OverheadUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleOverheadUpdatedIterator struct {
	Event *GasPriceOracleOverheadUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleOverheadUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleOverheadUpdated)
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
		it.Event = new(GasPriceOracleOverheadUpdated)
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
func (it *GasPriceOracleOverheadUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleOverheadUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleOverheadUpdated represents a OverheadUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleOverheadUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOverheadUpdated is a free log retrieval operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterOverheadUpdated(opts *bind.FilterOpts) (*GasPriceOracleOverheadUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleOverheadUpdatedIterator{contract: _GasPriceOracle.contract, event: "OverheadUpdated", logs: logs, sub: sub}, nil
}

// WatchOverheadUpdated is a free log subscription operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchOverheadUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleOverheadUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleOverheadUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
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

// ParseOverheadUpdated is a log parse operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseOverheadUpdated(log types.Log) (*GasPriceOracleOverheadUpdated, error) {
	event := new(GasPriceOracleOverheadUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasPriceOracle contract.
type GasPriceOracleOwnershipTransferredIterator struct {
	Event *GasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleOwnershipTransferred)
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
		it.Event = new(GasPriceOracleOwnershipTransferred)
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
func (it *GasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the GasPriceOracle contract.
type GasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleOwnershipTransferredIterator{contract: _GasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleOwnershipTransferred)
				if err := _GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*GasPriceOracleOwnershipTransferred, error) {
	event := new(GasPriceOracleOwnershipTransferred)
	if err := _GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleScalarUpdatedIterator is returned from FilterScalarUpdated and is used to iterate over the raw logs and unpacked data for ScalarUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleScalarUpdatedIterator struct {
	Event *GasPriceOracleScalarUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleScalarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleScalarUpdated)
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
		it.Event = new(GasPriceOracleScalarUpdated)
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
func (it *GasPriceOracleScalarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleScalarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleScalarUpdated represents a ScalarUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleScalarUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterScalarUpdated is a free log retrieval operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterScalarUpdated(opts *bind.FilterOpts) (*GasPriceOracleScalarUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleScalarUpdatedIterator{contract: _GasPriceOracle.contract, event: "ScalarUpdated", logs: logs, sub: sub}, nil
}

// WatchScalarUpdated is a free log subscription operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchScalarUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleScalarUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleScalarUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
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

// ParseScalarUpdated is a log parse operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 arg0)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseScalarUpdated(log types.Log) (*GasPriceOracleScalarUpdated, error) {
	event := new(GasPriceOracleScalarUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
