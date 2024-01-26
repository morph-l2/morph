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

// L2CrossDomainMessengerMetaData contains all meta data concerning the L2CrossDomainMessenger contract.
var L2CrossDomainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxFailedExecutionTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxFailedExecutionTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxFailedExecutionTimes\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL1MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageSendTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561001057600080fd5b506001608052600060a081905260c05261002861002d565b6100ec565b600054610100900460ff16156100995760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100ea576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a05160c051611d5f61011b600039600061048901526000610460015260006104370152611d5f6000f3fe6080604052600436106101125760003560e01c8063797594b0116100a5578063bedb86fb11610074578063e70fc93b11610059578063e70fc93b14610331578063ecc704281461036c578063f2fde38b1461038157600080fd5b8063bedb86fb146102f1578063c4d66de81461031157600080fd5b8063797594b0146102665780638da5cb5b146102935780638ef1332e146102be578063b2267a7b146102de57600080fd5b80635c975abb116100e15780635c975abb146101f95780635f7b1577146102115780636e296e4514610224578063715018a61461025157600080fd5b806302345b501461011e5780632a6cccb214610163578063478222c21461018557806354fd4d50146101d757600080fd5b3661011957005b600080fd5b34801561012a57600080fd5b5061014e6101393660046118a7565b60fc6020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561016f57600080fd5b5061018361017e3660046118e9565b6103a1565b005b34801561019157600080fd5b5060cb546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b3480156101e357600080fd5b506101ec610430565b60405161015a9190611979565b34801561020557600080fd5b5060655460ff1661014e565b61018361021f36600461198c565b6104d3565b34801561023057600080fd5b5060c9546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561025d57600080fd5b50610183610526565b34801561027257600080fd5b5060ca546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561029f57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101b2565b3480156102ca57600080fd5b506101836102d9366004611b0c565b61053a565b6101836102ec366004611b7e565b6106b8565b3480156102fd57600080fd5b5061018361030c366004611bdd565b6106d2565b34801561031d57600080fd5b5061018361032c3660046118e9565b6106f3565b34801561033d57600080fd5b5061035e61034c3660046118a7565b60fb6020526000908152604090205481565b60405190815260200161015a565b34801561037857600080fd5b5061035e6108b8565b34801561038d57600080fd5b5061018361039c3660046118e9565b610942565b6103a96109dc565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f591015b60405180910390a15050565b606061045b7f0000000000000000000000000000000000000000000000000000000000000000610a43565b6104847f0000000000000000000000000000000000000000000000000000000000000000610a43565b6104ad7f0000000000000000000000000000000000000000000000000000000000000000610a43565b6040516020016104bf93929190611bff565b604051602081830303815290604052905090565b6104db610b01565b61051e868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250610b54915050565b505050505050565b61052e6109dc565b6105386000610dc9565b565b610542610b01565b60ca5473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff161461060b5760405162461bcd60e51b8152602060048201526024808201527f43616c6c6572206973206e6f74204c3143726f7373446f6d61696e4d6573736560448201527f6e6765720000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600061061a8686868686610e40565b8051602091820120600081815260fc90925260409091205490915060ff16156106ab5760405162461bcd60e51b815260206004820152602960248201527f4d6573736167652077617320616c7265616479207375636365737366756c6c7960448201527f20657865637574656400000000000000000000000000000000000000000000006064820152608401610602565b61051e8686868585610edc565b6106c0610b01565b6106cc84848484610b54565b50505050565b6106da6109dc565b80156106eb576106e8611164565b50565b6106e86111e9565b600054610100900460ff16158080156107135750600054600160ff909116105b8061072d5750303b15801561072d575060005460ff166001145b61079f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610602565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107fd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff821661084a576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610855826000611240565b80156108b457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610424565b5050565b600073530000000000000000000000000000000000000173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610919573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093d9190611c75565b905090565b61094a6109dc565b73ffffffffffffffffffffffffffffffffffffffff81166109d35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610602565b6106e881610dc9565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105385760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610602565b60606000610a5083611374565b600101905060008167ffffffffffffffff811115610a7057610a70611a32565b6040519080825280601f01601f191660200182016040528015610a9a576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610aa457509392505050565b60655460ff16156105385760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610602565b610b5c611457565b823414610bab5760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d6174636800000000000000000000000000006044820152606401610602565b6000735300000000000000000000000000000000000001905060008173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c359190611c75565b90506000610c463388888589610e40565b8051602091820120600081815260fb90925260409091205490915015610cae5760405162461bcd60e51b815260206004820152601260248201527f4475706c696361746564206d65737361676500000000000000000000000000006044820152606401610602565b600081815260fb602052604090819020429055517f600a2e770000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff84169063600a2e77906024016020604051808303816000875af1158015610d2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d509190611c75565b5073ffffffffffffffffffffffffffffffffffffffff87163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8885888a604051610db49493929190611c8e565b60405180910390a35050506106cc6001609755565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60608585858585604051602401610e5b959493929190611cbd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b7fffffffffffffffffffffffffacffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff851601610f875760405162461bcd60e51b815260206004820152602660248201527f466f7262696420746f2063616c6c206c3220746f206c31206d6573736167652060448201527f70617373657200000000000000000000000000000000000000000000000000006064820152608401610602565b610f90846114b7565b60c95473ffffffffffffffffffffffffffffffffffffffff90811690861603610ffb5760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6573736167652073656e646572000000000000000000006044820152606401610602565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff878116919091179091556040516000918616908590611056908690611d0d565b60006040518083038185875af1925050503d8060008114611093576040519150601f19603f3d011682016040523d82523d6000602084013e611098565b606091505b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561113157600082815260fc602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261051e565b60405182907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f90600090a2505050505050565b61116c610b01565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111bf3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6111f161151c565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336111bf565b600054610100900460ff166112bd5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b6112c561156e565b6112cd6115f3565b6112d5611678565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661dead1790915560ca805473ffffffffffffffffffffffffffffffffffffffff858116919093161790558116156108b45760cb805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790555050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106113bd577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106113e9576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061140757662386f26fc10000830492506010015b6305f5e100831061141f576305f5e100830492506008015b612710831061143357612710830492506004015b60648310611445576064830492506002015b600a8310611451576001015b92915050565b6002609754036114a95760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610602565b6002609755565b6001609755565b3073ffffffffffffffffffffffffffffffffffffffff8216036106e85760405162461bcd60e51b815260206004820152601e60248201527f4d657373656e6765723a20466f7262696420746f2063616c6c2073656c6600006044820152606401610602565b60655460ff166105385760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610602565b600054610100900460ff166115eb5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b6105386116fd565b600054610100900460ff166116705760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b610538611783565b600054610100900460ff166116f55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b61053861182a565b600054610100900460ff1661177a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b61053833610dc9565b600054610100900460ff166118005760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600054610100900460ff166114b05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610602565b6000602082840312156118b957600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146118e457600080fd5b919050565b6000602082840312156118fb57600080fd5b611904826118c0565b9392505050565b60005b8381101561192657818101518382015260200161190e565b50506000910152565b6000815180845261194781602086016020860161190b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611904602083018461192f565b60008060008060008060a087890312156119a557600080fd5b6119ae876118c0565b955060208701359450604087013567ffffffffffffffff808211156119d257600080fd5b818901915089601f8301126119e657600080fd5b8135818111156119f557600080fd5b8a6020828501011115611a0757600080fd5b60208301965080955050505060608701359150611a26608088016118c0565b90509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112611a7257600080fd5b813567ffffffffffffffff80821115611a8d57611a8d611a32565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611ad357611ad3611a32565b81604052838152866020858801011115611aec57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a08688031215611b2457600080fd5b611b2d866118c0565b9450611b3b602087016118c0565b93506040860135925060608601359150608086013567ffffffffffffffff811115611b6557600080fd5b611b7188828901611a61565b9150509295509295909350565b60008060008060808587031215611b9457600080fd5b611b9d856118c0565b935060208501359250604085013567ffffffffffffffff811115611bc057600080fd5b611bcc87828801611a61565b949793965093946060013593505050565b600060208284031215611bef57600080fd5b8135801515811461190457600080fd5b60008451611c1181846020890161190b565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611c4d816001850160208a0161190b565b60019201918201528351611c6881600284016020880161190b565b0160020195945050505050565b600060208284031215611c8757600080fd5b5051919050565b848152836020820152826040820152608060608201526000611cb3608083018461192f565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015283606083015260a06080830152611d0260a083018461192f565b979650505050505050565b60008251611d1f81846020870161190b565b919091019291505056fea2646970667358221220c518ccbb54da2d598931b28aeec44db00d71459d71aba7f1a0f3a94f8f3f762764736f6c63430008100033",
}

// L2CrossDomainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2CrossDomainMessengerMetaData.ABI instead.
var L2CrossDomainMessengerABI = L2CrossDomainMessengerMetaData.ABI

// L2CrossDomainMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2CrossDomainMessengerMetaData.Bin instead.
var L2CrossDomainMessengerBin = L2CrossDomainMessengerMetaData.Bin

// DeployL2CrossDomainMessenger deploys a new Ethereum contract, binding an instance of L2CrossDomainMessenger to it.
func DeployL2CrossDomainMessenger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2CrossDomainMessenger, error) {
	parsed, err := L2CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2CrossDomainMessengerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2CrossDomainMessenger{L2CrossDomainMessengerCaller: L2CrossDomainMessengerCaller{contract: contract}, L2CrossDomainMessengerTransactor: L2CrossDomainMessengerTransactor{contract: contract}, L2CrossDomainMessengerFilterer: L2CrossDomainMessengerFilterer{contract: contract}}, nil
}

// L2CrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type L2CrossDomainMessenger struct {
	L2CrossDomainMessengerCaller     // Read-only binding to the contract
	L2CrossDomainMessengerTransactor // Write-only binding to the contract
	L2CrossDomainMessengerFilterer   // Log filterer for contract events
}

// L2CrossDomainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2CrossDomainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2CrossDomainMessengerSession struct {
	Contract     *L2CrossDomainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CrossDomainMessengerCallerSession struct {
	Contract *L2CrossDomainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2CrossDomainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2CrossDomainMessengerTransactorSession struct {
	Contract     *L2CrossDomainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2CrossDomainMessengerRaw struct {
	Contract *L2CrossDomainMessenger // Generic contract binding to access the raw methods on
}

// L2CrossDomainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCallerRaw struct {
	Contract *L2CrossDomainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// L2CrossDomainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactorRaw struct {
	Contract *L2CrossDomainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2CrossDomainMessenger creates a new instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessenger(address common.Address, backend bind.ContractBackend) (*L2CrossDomainMessenger, error) {
	contract, err := bindL2CrossDomainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessenger{L2CrossDomainMessengerCaller: L2CrossDomainMessengerCaller{contract: contract}, L2CrossDomainMessengerTransactor: L2CrossDomainMessengerTransactor{contract: contract}, L2CrossDomainMessengerFilterer: L2CrossDomainMessengerFilterer{contract: contract}}, nil
}

// NewL2CrossDomainMessengerCaller creates a new read-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerCaller(address common.Address, caller bind.ContractCaller) (*L2CrossDomainMessengerCaller, error) {
	contract, err := bindL2CrossDomainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerCaller{contract: contract}, nil
}

// NewL2CrossDomainMessengerTransactor creates a new write-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2CrossDomainMessengerTransactor, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerTransactor{contract: contract}, nil
}

// NewL2CrossDomainMessengerFilterer creates a new log filterer instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2CrossDomainMessengerFilterer, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFilterer{contract: contract}, nil
}

// bindL2CrossDomainMessenger binds a generic wrapper to an already deployed contract.
func bindL2CrossDomainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2CrossDomainMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Counterpart() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.Counterpart(&_L2CrossDomainMessenger.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Counterpart() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.Counterpart(&_L2CrossDomainMessenger.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) FeeVault() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.FeeVault(&_L2CrossDomainMessenger.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) FeeVault() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.FeeVault(&_L2CrossDomainMessenger.CallOpts)
}

// IsL1MessageExecuted is a free data retrieval call binding the contract method 0x02345b50.
//
// Solidity: function isL1MessageExecuted(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) IsL1MessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "isL1MessageExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL1MessageExecuted is a free data retrieval call binding the contract method 0x02345b50.
//
// Solidity: function isL1MessageExecuted(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) IsL1MessageExecuted(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.IsL1MessageExecuted(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// IsL1MessageExecuted is a free data retrieval call binding the contract method 0x02345b50.
//
// Solidity: function isL1MessageExecuted(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) IsL1MessageExecuted(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.IsL1MessageExecuted(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MessageSendTimestamp(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "messageSendTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MessageSendTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageSendTimestamp(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MessageSendTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageSendTimestamp(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Owner() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.Owner(&_L2CrossDomainMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Owner() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.Owner(&_L2CrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Paused() (bool, error) {
	return _L2CrossDomainMessenger.Contract.Paused(&_L2CrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Paused() (bool, error) {
	return _L2CrossDomainMessenger.Contract.Paused(&_L2CrossDomainMessenger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Version() (string, error) {
	return _L2CrossDomainMessenger.Contract.Version(&_L2CrossDomainMessenger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Version() (string, error) {
	return _L2CrossDomainMessenger.Contract.Version(&_L2CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _counterpart) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "initialize", _counterpart)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _counterpart) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Initialize(_counterpart common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Initialize(&_L2CrossDomainMessenger.TransactOpts, _counterpart)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _counterpart) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) Initialize(_counterpart common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Initialize(&_L2CrossDomainMessenger.TransactOpts, _counterpart)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "relayMessage", _from, _to, _value, _nonce, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RelayMessage(_from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _from, _to, _value, _nonce, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) RelayMessage(_from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _from, _to, _value, _nonce, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RenounceOwnership(&_L2CrossDomainMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RenounceOwnership(&_L2CrossDomainMessenger.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address ) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "sendMessage", _to, _value, _message, _gasLimit, arg4)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address ) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SendMessage(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _to, _value, _message, _gasLimit, arg4)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address ) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) SendMessage(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _to, _value, _message, _gasLimit, arg4)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) SendMessage0(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "sendMessage0", _to, _value, _message, _gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SendMessage0(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage0(&_L2CrossDomainMessenger.TransactOpts, _to, _value, _message, _gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) SendMessage0(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage0(&_L2CrossDomainMessenger.TransactOpts, _to, _value, _message, _gasLimit)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SetPause(_status bool) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SetPause(&_L2CrossDomainMessenger.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SetPause(&_L2CrossDomainMessenger.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.TransferOwnership(&_L2CrossDomainMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.TransferOwnership(&_L2CrossDomainMessenger.TransactOpts, newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.UpdateFeeVault(&_L2CrossDomainMessenger.TransactOpts, _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.UpdateFeeVault(&_L2CrossDomainMessenger.TransactOpts, _newFeeVault)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Receive() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Receive(&_L2CrossDomainMessenger.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) Receive() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Receive(&_L2CrossDomainMessenger.TransactOpts)
}

// L2CrossDomainMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
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
		it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
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
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*L2CrossDomainMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFailedRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerFailedRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*L2CrossDomainMessengerFailedRelayedMessage, error) {
	event := new(L2CrossDomainMessengerFailedRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerInitializedIterator struct {
	Event *L2CrossDomainMessengerInitialized // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerInitialized)
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
		it.Event = new(L2CrossDomainMessengerInitialized)
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
func (it *L2CrossDomainMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerInitialized represents a Initialized event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2CrossDomainMessengerInitializedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerInitializedIterator{contract: _L2CrossDomainMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerInitialized)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseInitialized(log types.Log) (*L2CrossDomainMessengerInitialized, error) {
	event := new(L2CrossDomainMessengerInitialized)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerOwnershipTransferredIterator struct {
	Event *L2CrossDomainMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerOwnershipTransferred)
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
		it.Event = new(L2CrossDomainMessengerOwnershipTransferred)
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
func (it *L2CrossDomainMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2CrossDomainMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerOwnershipTransferredIterator{contract: _L2CrossDomainMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerOwnershipTransferred)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*L2CrossDomainMessengerOwnershipTransferred, error) {
	event := new(L2CrossDomainMessengerOwnershipTransferred)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerPausedIterator struct {
	Event *L2CrossDomainMessengerPaused // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerPaused)
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
		it.Event = new(L2CrossDomainMessengerPaused)
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
func (it *L2CrossDomainMessengerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerPaused represents a Paused event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterPaused(opts *bind.FilterOpts) (*L2CrossDomainMessengerPausedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerPausedIterator{contract: _L2CrossDomainMessenger.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerPaused) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerPaused)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParsePaused(log types.Log) (*L2CrossDomainMessengerPaused, error) {
	event := new(L2CrossDomainMessengerPaused)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerRelayedMessage)
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
		it.Event = new(L2CrossDomainMessengerRelayedMessage)
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
func (it *L2CrossDomainMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*L2CrossDomainMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseRelayedMessage(log types.Log) (*L2CrossDomainMessengerRelayedMessage, error) {
	event := new(L2CrossDomainMessengerRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessageIterator struct {
	Event *L2CrossDomainMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerSentMessage)
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
		it.Event = new(L2CrossDomainMessengerSentMessage)
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
func (it *L2CrossDomainMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerSentMessage represents a SentMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*L2CrossDomainMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerSentMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerSentMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseSentMessage(log types.Log) (*L2CrossDomainMessengerSentMessage, error) {
	event := new(L2CrossDomainMessengerSentMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUnpausedIterator struct {
	Event *L2CrossDomainMessengerUnpaused // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerUnpaused)
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
		it.Event = new(L2CrossDomainMessengerUnpaused)
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
func (it *L2CrossDomainMessengerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerUnpaused represents a Unpaused event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2CrossDomainMessengerUnpausedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerUnpausedIterator{contract: _L2CrossDomainMessenger.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerUnpaused)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseUnpaused(log types.Log) (*L2CrossDomainMessengerUnpaused, error) {
	event := new(L2CrossDomainMessengerUnpaused)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerUpdateFeeVaultIterator is returned from FilterUpdateFeeVault and is used to iterate over the raw logs and unpacked data for UpdateFeeVault events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUpdateFeeVaultIterator struct {
	Event *L2CrossDomainMessengerUpdateFeeVault // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerUpdateFeeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerUpdateFeeVault)
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
		it.Event = new(L2CrossDomainMessengerUpdateFeeVault)
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
func (it *L2CrossDomainMessengerUpdateFeeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerUpdateFeeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerUpdateFeeVault represents a UpdateFeeVault event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUpdateFeeVault struct {
	OldFeeVault common.Address
	NewFeeVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeVault is a free log retrieval operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterUpdateFeeVault(opts *bind.FilterOpts) (*L2CrossDomainMessengerUpdateFeeVaultIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerUpdateFeeVaultIterator{contract: _L2CrossDomainMessenger.contract, event: "UpdateFeeVault", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeVault is a free log subscription operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchUpdateFeeVault(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerUpdateFeeVault) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerUpdateFeeVault)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
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

// ParseUpdateFeeVault is a log parse operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseUpdateFeeVault(log types.Log) (*L2CrossDomainMessengerUpdateFeeVault, error) {
	event := new(L2CrossDomainMessengerUpdateFeeVault)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator is returned from FilterUpdateMaxFailedExecutionTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxFailedExecutionTimes events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator struct {
	Event *L2CrossDomainMessengerUpdateMaxFailedExecutionTimes // Event containing the contract specifics and raw log

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
func (it *L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerUpdateMaxFailedExecutionTimes)
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
		it.Event = new(L2CrossDomainMessengerUpdateMaxFailedExecutionTimes)
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
func (it *L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerUpdateMaxFailedExecutionTimes represents a UpdateMaxFailedExecutionTimes event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUpdateMaxFailedExecutionTimes struct {
	OldMaxFailedExecutionTimes *big.Int
	NewMaxFailedExecutionTimes *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxFailedExecutionTimes is a free log retrieval operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterUpdateMaxFailedExecutionTimes(opts *bind.FilterOpts) (*L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "UpdateMaxFailedExecutionTimes")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerUpdateMaxFailedExecutionTimesIterator{contract: _L2CrossDomainMessenger.contract, event: "UpdateMaxFailedExecutionTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxFailedExecutionTimes is a free log subscription operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchUpdateMaxFailedExecutionTimes(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerUpdateMaxFailedExecutionTimes) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "UpdateMaxFailedExecutionTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerUpdateMaxFailedExecutionTimes)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "UpdateMaxFailedExecutionTimes", log); err != nil {
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

// ParseUpdateMaxFailedExecutionTimes is a log parse operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseUpdateMaxFailedExecutionTimes(log types.Log) (*L2CrossDomainMessengerUpdateMaxFailedExecutionTimes, error) {
	event := new(L2CrossDomainMessengerUpdateMaxFailedExecutionTimes)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "UpdateMaxFailedExecutionTimes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
