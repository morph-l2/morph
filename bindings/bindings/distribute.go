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

// DistributeMetaData contains all meta data concerning the Distribute contract.
var DistributeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upToEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upToEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECORD_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_EPOCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainsNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"newDelegation\",\"type\":\"bool\"}],\"name\":\"notifyDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainsNumber\",\"type\":\"uint256\"}],\"name\":\"notifyUndelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"queryUnclaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"unclaimedCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"delegatorRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"commissions\",\"type\":\"uint256[]\"}],\"name\":\"updateEpochReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040526201518060e052348015610017575f80fd5b5073530000000000000000000000000000000000001060805273530000000000000000000000000000000000001260c05273530000000000000000000000000000000000000560a05260805160a05160c05160e0516121376100d25f395f61027e01525f8181610166015281816102a3015281816104db015281816109e401528181610af80152610f7d01525f81816101f70152610d4601525f8181610231015281816111870152818161123401526112db01526121375ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c8063996cba6811610093578063d557714111610063578063d55771411461022c578063e16bcc3214610253578063f2fde38b14610266578063fadfa08714610279575f80fd5b8063996cba68146101cc578063ad8e1223146101df578063cd4281d0146101f2578063cdd0c50e14610219575f80fd5b8063807de443116100ce578063807de443146101615780638129fc1c146101a05780638da5cb5b146101a85780639889be51146101b9575f80fd5b806341302560146100ff5780635cf20c7b14610114578063715018a6146101275780637ac3339a1461012f575b5f80fd5b61011261010d366004611d1a565b6102a0565b005b610112610122366004611d1a565b6104d8565b6101126106a7565b61014e61013d366004611d42565b60676020525f908152604090205481565b6040519081526020015b60405180910390f35b6101887f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610158565b6101126106ba565b6033546001600160a01b0316610188565b61014e6101c7366004611d5b565b61082c565b6101126101da366004611d8c565b6109e1565b6101126101ed366004611dc5565b610af5565b6101887f000000000000000000000000000000000000000000000000000000000000000081565b610112610227366004611e56565b610d43565b6101887f000000000000000000000000000000000000000000000000000000000000000081565b610112610261366004611eff565b610f7a565b610112610274366004611d42565b6110c0565b61014e7f000000000000000000000000000000000000000000000000000000000000000081565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161461031d5760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064015b60405180910390fd5b6065545f0361036e5760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e742079657400000000000000000000000000000000000000006044820152606401610314565b8080158061037d575060655482115b156103945760016065546103919190611f95565b90505b6001600160a01b0383165f908152606760205260409020548110156103fb5760405162461bcd60e51b815260206004820152601660248201527f616c6c20636f6d6d697373696f6e20636c61696d6564000000000000000000006044820152606401610314565b6001600160a01b0383165f908152606760205260408120545b82811161045f576001600160a01b0385165f90815260666020908152604080832084845290915290206001015461044b9083611fa8565b91508061045781611fbb565b915050610414565b50801561047057610470848261114d565b61047b826001611fa8565b6001600160a01b0385165f8181526067602090815260409182902093909355805185815292830184905290917fe4760bd616775d8b0ae78f9b8bfa4b453fdde769d10a0559ba473157ce9011d4910160405180910390a250505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146105505760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f7765646044820152606401610314565b6065545f036105a15760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e742079657400000000000000000000000000000000000000006044820152606401610314565b808015806105b0575060655482115b156105c75760016065546105c49190611f95565b90505b5f805b6001600160a01b0385165f9081526068602052604090206105ea906113ab565b811015610690576001600160a01b0385165f90815260686020526040812061061290836113ba565b6001600160a01b0387165f90815260686020526040902090915061063690826113cc565b801561066a57506001600160a01b038087165f90815260686020908152604080832093851683526003909301905220548410155b156106875761067a8187876113ed565b6106849084611fa8565b92505b506001016105ca565b5080156106a1576106a1848261114d565b50505050565b6106af611988565b6106b85f6119e2565b565b5f54610100900460ff16158080156106d857505f54600160ff909116105b806106f15750303b1580156106f157505f5460ff166001145b6107635760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610314565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107bf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6107c7611a4b565b8015610829575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b038082165f9081526068602090815260408083209386168352600390930190529081205481908190805b6065548110156109d7576001600160a01b038088165f9081526066602090815260408083208584528252808320938a168352600690930190522054156108d1576001600160a01b038088165f9081526066602090815260408083208584528252808320938a16835260069093019052205492505b6001600160a01b0387165f90815260666020908152604080832084845290915290206002015415610925576001600160a01b0387165f90815260666020908152604080832084845290915290206002015493505b6001600160a01b0387165f9081526066602090815260408083208484529091529020548490610955908590611ff2565b61095f9190612009565b6109699086611fa8565b6001600160a01b038088165f908152606860209081526040808320938c16835260029093019052205490955060ff1680156109cb57506001600160a01b038087165f908152606860209081526040808320938b16835260049093019052205481145b6109d75760010161085d565b5050505092915050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610a595760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f7765646044820152606401610314565b6065545f03610aaa5760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e742079657400000000000000000000000000000000000000006044820152606401610314565b80801580610ab9575060655482115b15610ad0576001606554610acd9190611f95565b90505b5f610adc8585846113ed565b90508015610aee57610aee848261114d565b5050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610b6d5760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f7765646044820152606401610314565b6001600160a01b0385165f908152606660209081526040808320868452909152902060028101839055600301819055821580610bd057506001600160a01b038085165f908152606860209081526040808320938916835260039093019052205483145b15610cb5576001600160a01b0385165f9081526066602090815260408083208684529091529020610c049060040185611acf565b506001600160a01b038086165f90815260666020908152604080832087845282528083209388168352600690930181528282208290556068905220610c499086611acf565b506001600160a01b038481165f908152606860209081526040808320938916835260028401825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556003840182528083208390556004909301905290812055610aee565b6001600160a01b038085165f9081526068602090815260408083209389168352600290930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155610d139084611f95565b6001600160a01b038086165f908152606860209081526040808320938a1683526004909301905220555050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610dbb5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c79207265636f726420636f6e747261637420616c6c6f776564000000006044820152606401610314565b60658054905f610dca83611fbb565b9190505550866001606554610ddf9190611f95565b14610e2c5760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642065706f636820696e646578000000000000000000000000006044820152606401610314565b8285148015610e3a57508085145b610e865760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642064617461206c656e677468000000000000000000000000006044820152606401610314565b5f5b85811015610f7057848482818110610ea257610ea2612041565b9050602002013560665f898985818110610ebe57610ebe612041565b9050602002016020810190610ed39190611d42565b6001600160a01b0316815260208082019290925260409081015f9081208c8252909252902055828282818110610f0b57610f0b612041565b9050602002013560665f898985818110610f2757610f27612041565b9050602002016020810190610f3c9190611d42565b6001600160a01b0316815260208082019290925260409081015f9081208c8252909252902060019081019190915501610e88565b5050505050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610ff25760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f7765646044820152606401610314565b6001600160a01b0387165f9081526066602090815260408083208884529091529020600281018490556003810183905561102f9060040187611ae3565b506001600160a01b038088165f9081526066602090815260408083208984528252808320938a168352600690930190522084905580156110b7576001600160a01b0386165f90815260686020526040902061108a9088611ae3565b506001600160a01b038087165f908152606860209081526040808320938b16835260039093019052208590555b50505050505050565b6110c8611988565b6001600160a01b0381166111445760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610314565b610829816119e2565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa1580156111ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f2919061206e565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303815f875af115801561127c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a09190612085565b506040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015611322573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611346919061206e565b90505f8311801561135f57508261135d8383611f95565b145b6106a15760405162461bcd60e51b815260206004820152601b60248201527f6d6f72706820746f6b656e207472616e73666572206661696c656400000000006044820152606401610314565b5f6113b4825490565b92915050565b5f6113c58383611af7565b9392505050565b6001600160a01b0381165f90815260018301602052604081205415156113c5565b6001600160a01b0382165f90815260686020526040812061140e90856113cc565b61145a5760405162461bcd60e51b815260206004820152601360248201527f6e6f2072656d61696e696e6720726577617264000000000000000000000000006044820152606401610314565b6001600160a01b038084165f90815260686020908152604080832093881683526003909301905220548210156114d25760405162461bcd60e51b815260206004820152601260248201527f616c6c2072657761726420636c61696d656400000000000000000000000000006044820152606401610314565b6001600160a01b038084165f90815260686020908152604080832093881683526003909301905220545b82811161192a576001600160a01b038086165f9081526066602090815260408083208584528083528184206002810154958a16855260068101845291842054938690529091525461154d9190611ff2565b6115579190612009565b6115619083611fa8565b6001600160a01b0386165f9081526066602052604081209193506115aa91869161158c856001611fa8565b81526020019081526020015f206004016113cc90919063ffffffff16565b611736576001600160a01b0385165f9081526066602052604081206115f5918691906115d7856001611fa8565b81526020019081526020015f20600401611ae390919063ffffffff16565b506001600160a01b038086165f818152606660208181526040808420878552808352818520968b168552600690960182528320549383525290919061163b846001611fa8565b815260208082019290925260409081015f9081206001600160a01b03808a168352600690910184528282209490945592881683526066909152812090611682836001611fa8565b81526020019081526020015f20600201545f03611736576001600160a01b0385165f81815260666020818152604080842086855280835290842060020154948452919052906116d2846001611fa8565b815260208082019290925260409081015f908120600201939093556001600160a01b0388168084526066808452828520868652808552928520600301549185529092529091611722846001611fa8565b815260208101919091526040015f20600301555b6001600160a01b0385165f9081526066602090815260408083208484529091528120600301805491611767836120a0565b90915550506001600160a01b0385165f90815260666020908152604080832084845290915281206003015490036117eb576001600160a01b0385165f908152606660209081526040808320848452909152812081815560018101829055600281018290556003810182905590600482018181816117e48282611cd1565b5050505050505b6001600160a01b038085165f908152606860209081526040808320938916835260029093019052205460ff16801561184a57506001600160a01b038085165f908152606860209081526040808320938916835260049093019052205481145b156118dc576001600160a01b0384165f9081526068602052604090206118709086611acf565b506001600160a01b038481165f908152606860209081526040808320938916835260028401825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055600384018252808320839055600490930190529081205561192a565b6001600160a01b038085165f90815260686020908152604080832093891683526003909301905290812080549161191283611fbb565b9190505550808061192290611fbb565b9150506114fc565b50836001600160a01b0316836001600160a01b03167f7a84a08b02c91f3c62d572853f966fc799bbd121e8ad7833a4494ab8dcfcb4048484604051611979929190918252602082015260400190565b60405180910390a39392505050565b6033546001600160a01b031633146106b85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610314565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611ac75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610314565b6106b8611b1d565b5f6113c5836001600160a01b038416611ba2565b5f6113c5836001600160a01b038416611c85565b5f825f018281548110611b0c57611b0c612041565b905f5260205f200154905092915050565b5f54610100900460ff16611b995760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610314565b6106b8336119e2565b5f8181526001830160205260408120548015611c7c575f611bc4600183611f95565b85549091505f90611bd790600190611f95565b9050818114611c36575f865f018281548110611bf557611bf5612041565b905f5260205f200154905080875f018481548110611c1557611c15612041565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611c4757611c476120d4565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506113b4565b5f9150506113b4565b5f818152600183016020526040812054611cca57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556113b4565b505f6113b4565b5080545f8255905f5260205f209081019061082991905b80821115611cfb575f8155600101611ce8565b5090565b80356001600160a01b0381168114611d15575f80fd5b919050565b5f8060408385031215611d2b575f80fd5b611d3483611cff565b946020939093013593505050565b5f60208284031215611d52575f80fd5b6113c582611cff565b5f8060408385031215611d6c575f80fd5b611d7583611cff565b9150611d8360208401611cff565b90509250929050565b5f805f60608486031215611d9e575f80fd5b611da784611cff565b9250611db560208501611cff565b9150604084013590509250925092565b5f805f805f60a08688031215611dd9575f80fd5b611de286611cff565b9450611df060208701611cff565b94979496505050506040830135926060810135926080909101359150565b5f8083601f840112611e1e575f80fd5b50813567ffffffffffffffff811115611e35575f80fd5b6020830191508360208260051b8501011115611e4f575f80fd5b9250929050565b5f805f805f805f6080888a031215611e6c575f80fd5b87359650602088013567ffffffffffffffff80821115611e8a575f80fd5b611e968b838c01611e0e565b909850965060408a0135915080821115611eae575f80fd5b611eba8b838c01611e0e565b909650945060608a0135915080821115611ed2575f80fd5b50611edf8a828b01611e0e565b989b979a50959850939692959293505050565b8015158114610829575f80fd5b5f805f805f805f60e0888a031215611f15575f80fd5b611f1e88611cff565b9650611f2c60208901611cff565b955060408801359450606088013593506080880135925060a0880135915060c0880135611f5881611ef2565b8091505092959891949750929550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156113b4576113b4611f68565b808201808211156113b4576113b4611f68565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611feb57611feb611f68565b5060010190565b80820281158282048414176113b4576113b4611f68565b5f8261203c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020828403121561207e575f80fd5b5051919050565b5f60208284031215612095575f80fd5b81516113c581611ef2565b5f816120ae576120ae611f68565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212207e88a01a0c47182f77e374f20f19a3c6637d550b898988306d19fec73680c6a864736f6c63430008180033",
}

// DistributeABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributeMetaData.ABI instead.
var DistributeABI = DistributeMetaData.ABI

// DistributeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributeMetaData.Bin instead.
var DistributeBin = DistributeMetaData.Bin

// DeployDistribute deploys a new Ethereum contract, binding an instance of Distribute to it.
func DeployDistribute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Distribute, error) {
	parsed, err := DistributeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Distribute{DistributeCaller: DistributeCaller{contract: contract}, DistributeTransactor: DistributeTransactor{contract: contract}, DistributeFilterer: DistributeFilterer{contract: contract}}, nil
}

// Distribute is an auto generated Go binding around an Ethereum contract.
type Distribute struct {
	DistributeCaller     // Read-only binding to the contract
	DistributeTransactor // Write-only binding to the contract
	DistributeFilterer   // Log filterer for contract events
}

// DistributeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributeSession struct {
	Contract     *Distribute       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributeCallerSession struct {
	Contract *DistributeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DistributeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributeTransactorSession struct {
	Contract     *DistributeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DistributeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributeRaw struct {
	Contract *Distribute // Generic contract binding to access the raw methods on
}

// DistributeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributeCallerRaw struct {
	Contract *DistributeCaller // Generic read-only contract binding to access the raw methods on
}

// DistributeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributeTransactorRaw struct {
	Contract *DistributeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistribute creates a new instance of Distribute, bound to a specific deployed contract.
func NewDistribute(address common.Address, backend bind.ContractBackend) (*Distribute, error) {
	contract, err := bindDistribute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distribute{DistributeCaller: DistributeCaller{contract: contract}, DistributeTransactor: DistributeTransactor{contract: contract}, DistributeFilterer: DistributeFilterer{contract: contract}}, nil
}

// NewDistributeCaller creates a new read-only instance of Distribute, bound to a specific deployed contract.
func NewDistributeCaller(address common.Address, caller bind.ContractCaller) (*DistributeCaller, error) {
	contract, err := bindDistribute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributeCaller{contract: contract}, nil
}

// NewDistributeTransactor creates a new write-only instance of Distribute, bound to a specific deployed contract.
func NewDistributeTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributeTransactor, error) {
	contract, err := bindDistribute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributeTransactor{contract: contract}, nil
}

// NewDistributeFilterer creates a new log filterer instance of Distribute, bound to a specific deployed contract.
func NewDistributeFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributeFilterer, error) {
	contract, err := bindDistribute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributeFilterer{contract: contract}, nil
}

// bindDistribute binds a generic wrapper to an already deployed contract.
func bindDistribute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribute *DistributeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribute.Contract.DistributeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribute *DistributeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribute.Contract.DistributeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribute *DistributeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribute.Contract.DistributeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribute *DistributeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribute *DistributeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribute *DistributeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribute.Contract.contract.Transact(opts, method, params...)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Distribute *DistributeCaller) L2STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "L2_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Distribute *DistributeSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Distribute.Contract.L2STAKINGCONTRACT(&_Distribute.CallOpts)
}

// L2STAKINGCONTRACT is a free data retrieval call binding the contract method 0x807de443.
//
// Solidity: function L2_STAKING_CONTRACT() view returns(address)
func (_Distribute *DistributeCallerSession) L2STAKINGCONTRACT() (common.Address, error) {
	return _Distribute.Contract.L2STAKINGCONTRACT(&_Distribute.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Distribute *DistributeCaller) MORPHTOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "MORPH_TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Distribute *DistributeSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _Distribute.Contract.MORPHTOKENCONTRACT(&_Distribute.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_Distribute *DistributeCallerSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _Distribute.Contract.MORPHTOKENCONTRACT(&_Distribute.CallOpts)
}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_Distribute *DistributeCaller) RECORDCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "RECORD_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_Distribute *DistributeSession) RECORDCONTRACT() (common.Address, error) {
	return _Distribute.Contract.RECORDCONTRACT(&_Distribute.CallOpts)
}

// RECORDCONTRACT is a free data retrieval call binding the contract method 0xcd4281d0.
//
// Solidity: function RECORD_CONTRACT() view returns(address)
func (_Distribute *DistributeCallerSession) RECORDCONTRACT() (common.Address, error) {
	return _Distribute.Contract.RECORDCONTRACT(&_Distribute.CallOpts)
}

// REWARDEPOCH is a free data retrieval call binding the contract method 0xfadfa087.
//
// Solidity: function REWARD_EPOCH() view returns(uint256)
func (_Distribute *DistributeCaller) REWARDEPOCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "REWARD_EPOCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDEPOCH is a free data retrieval call binding the contract method 0xfadfa087.
//
// Solidity: function REWARD_EPOCH() view returns(uint256)
func (_Distribute *DistributeSession) REWARDEPOCH() (*big.Int, error) {
	return _Distribute.Contract.REWARDEPOCH(&_Distribute.CallOpts)
}

// REWARDEPOCH is a free data retrieval call binding the contract method 0xfadfa087.
//
// Solidity: function REWARD_EPOCH() view returns(uint256)
func (_Distribute *DistributeCallerSession) REWARDEPOCH() (*big.Int, error) {
	return _Distribute.Contract.REWARDEPOCH(&_Distribute.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribute *DistributeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribute *DistributeSession) Owner() (common.Address, error) {
	return _Distribute.Contract.Owner(&_Distribute.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribute *DistributeCallerSession) Owner() (common.Address, error) {
	return _Distribute.Contract.Owner(&_Distribute.CallOpts)
}

// QueryUnclaimed is a free data retrieval call binding the contract method 0x9889be51.
//
// Solidity: function queryUnclaimed(address delegatee, address delegator) view returns(uint256 reward)
func (_Distribute *DistributeCaller) QueryUnclaimed(opts *bind.CallOpts, delegatee common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "queryUnclaimed", delegatee, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryUnclaimed is a free data retrieval call binding the contract method 0x9889be51.
//
// Solidity: function queryUnclaimed(address delegatee, address delegator) view returns(uint256 reward)
func (_Distribute *DistributeSession) QueryUnclaimed(delegatee common.Address, delegator common.Address) (*big.Int, error) {
	return _Distribute.Contract.QueryUnclaimed(&_Distribute.CallOpts, delegatee, delegator)
}

// QueryUnclaimed is a free data retrieval call binding the contract method 0x9889be51.
//
// Solidity: function queryUnclaimed(address delegatee, address delegator) view returns(uint256 reward)
func (_Distribute *DistributeCallerSession) QueryUnclaimed(delegatee common.Address, delegator common.Address) (*big.Int, error) {
	return _Distribute.Contract.QueryUnclaimed(&_Distribute.CallOpts, delegatee, delegator)
}

// UnclaimedCommission is a free data retrieval call binding the contract method 0x7ac3339a.
//
// Solidity: function unclaimedCommission(address delegatee) view returns(uint256 epochIndex)
func (_Distribute *DistributeCaller) UnclaimedCommission(opts *bind.CallOpts, delegatee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Distribute.contract.Call(opts, &out, "unclaimedCommission", delegatee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedCommission is a free data retrieval call binding the contract method 0x7ac3339a.
//
// Solidity: function unclaimedCommission(address delegatee) view returns(uint256 epochIndex)
func (_Distribute *DistributeSession) UnclaimedCommission(delegatee common.Address) (*big.Int, error) {
	return _Distribute.Contract.UnclaimedCommission(&_Distribute.CallOpts, delegatee)
}

// UnclaimedCommission is a free data retrieval call binding the contract method 0x7ac3339a.
//
// Solidity: function unclaimedCommission(address delegatee) view returns(uint256 epochIndex)
func (_Distribute *DistributeCallerSession) UnclaimedCommission(delegatee common.Address) (*big.Int, error) {
	return _Distribute.Contract.UnclaimedCommission(&_Distribute.CallOpts, delegatee)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address delegatee, address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactor) Claim(opts *bind.TransactOpts, delegatee common.Address, delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "claim", delegatee, delegator, targetEpochIndex)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address delegatee, address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeSession) Claim(delegatee common.Address, delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.Claim(&_Distribute.TransactOpts, delegatee, delegator, targetEpochIndex)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address delegatee, address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactorSession) Claim(delegatee common.Address, delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.Claim(&_Distribute.TransactOpts, delegatee, delegator, targetEpochIndex)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5cf20c7b.
//
// Solidity: function claimAll(address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactor) ClaimAll(opts *bind.TransactOpts, delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "claimAll", delegator, targetEpochIndex)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5cf20c7b.
//
// Solidity: function claimAll(address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeSession) ClaimAll(delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.ClaimAll(&_Distribute.TransactOpts, delegator, targetEpochIndex)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5cf20c7b.
//
// Solidity: function claimAll(address delegator, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactorSession) ClaimAll(delegator common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.ClaimAll(&_Distribute.TransactOpts, delegator, targetEpochIndex)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x41302560.
//
// Solidity: function claimCommission(address delegatee, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactor) ClaimCommission(opts *bind.TransactOpts, delegatee common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "claimCommission", delegatee, targetEpochIndex)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x41302560.
//
// Solidity: function claimCommission(address delegatee, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeSession) ClaimCommission(delegatee common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.ClaimCommission(&_Distribute.TransactOpts, delegatee, targetEpochIndex)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x41302560.
//
// Solidity: function claimCommission(address delegatee, uint256 targetEpochIndex) returns()
func (_Distribute *DistributeTransactorSession) ClaimCommission(delegatee common.Address, targetEpochIndex *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.ClaimCommission(&_Distribute.TransactOpts, delegatee, targetEpochIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Distribute *DistributeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Distribute *DistributeSession) Initialize() (*types.Transaction, error) {
	return _Distribute.Contract.Initialize(&_Distribute.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Distribute *DistributeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Distribute.Contract.Initialize(&_Distribute.TransactOpts)
}

// NotifyDelegation is a paid mutator transaction binding the contract method 0xe16bcc32.
//
// Solidity: function notifyDelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 amount, uint256 totalAmount, uint256 remainsNumber, bool newDelegation) returns()
func (_Distribute *DistributeTransactor) NotifyDelegation(opts *bind.TransactOpts, delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, amount *big.Int, totalAmount *big.Int, remainsNumber *big.Int, newDelegation bool) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "notifyDelegation", delegatee, delegator, effectiveEpoch, amount, totalAmount, remainsNumber, newDelegation)
}

// NotifyDelegation is a paid mutator transaction binding the contract method 0xe16bcc32.
//
// Solidity: function notifyDelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 amount, uint256 totalAmount, uint256 remainsNumber, bool newDelegation) returns()
func (_Distribute *DistributeSession) NotifyDelegation(delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, amount *big.Int, totalAmount *big.Int, remainsNumber *big.Int, newDelegation bool) (*types.Transaction, error) {
	return _Distribute.Contract.NotifyDelegation(&_Distribute.TransactOpts, delegatee, delegator, effectiveEpoch, amount, totalAmount, remainsNumber, newDelegation)
}

// NotifyDelegation is a paid mutator transaction binding the contract method 0xe16bcc32.
//
// Solidity: function notifyDelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 amount, uint256 totalAmount, uint256 remainsNumber, bool newDelegation) returns()
func (_Distribute *DistributeTransactorSession) NotifyDelegation(delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, amount *big.Int, totalAmount *big.Int, remainsNumber *big.Int, newDelegation bool) (*types.Transaction, error) {
	return _Distribute.Contract.NotifyDelegation(&_Distribute.TransactOpts, delegatee, delegator, effectiveEpoch, amount, totalAmount, remainsNumber, newDelegation)
}

// NotifyUndelegation is a paid mutator transaction binding the contract method 0xad8e1223.
//
// Solidity: function notifyUndelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 totalAmount, uint256 remainsNumber) returns()
func (_Distribute *DistributeTransactor) NotifyUndelegation(opts *bind.TransactOpts, delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, totalAmount *big.Int, remainsNumber *big.Int) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "notifyUndelegation", delegatee, delegator, effectiveEpoch, totalAmount, remainsNumber)
}

// NotifyUndelegation is a paid mutator transaction binding the contract method 0xad8e1223.
//
// Solidity: function notifyUndelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 totalAmount, uint256 remainsNumber) returns()
func (_Distribute *DistributeSession) NotifyUndelegation(delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, totalAmount *big.Int, remainsNumber *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.NotifyUndelegation(&_Distribute.TransactOpts, delegatee, delegator, effectiveEpoch, totalAmount, remainsNumber)
}

// NotifyUndelegation is a paid mutator transaction binding the contract method 0xad8e1223.
//
// Solidity: function notifyUndelegation(address delegatee, address delegator, uint256 effectiveEpoch, uint256 totalAmount, uint256 remainsNumber) returns()
func (_Distribute *DistributeTransactorSession) NotifyUndelegation(delegatee common.Address, delegator common.Address, effectiveEpoch *big.Int, totalAmount *big.Int, remainsNumber *big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.NotifyUndelegation(&_Distribute.TransactOpts, delegatee, delegator, effectiveEpoch, totalAmount, remainsNumber)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribute *DistributeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribute *DistributeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distribute.Contract.RenounceOwnership(&_Distribute.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribute *DistributeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distribute.Contract.RenounceOwnership(&_Distribute.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribute *DistributeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribute *DistributeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distribute.Contract.TransferOwnership(&_Distribute.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribute *DistributeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distribute.Contract.TransferOwnership(&_Distribute.TransactOpts, newOwner)
}

// UpdateEpochReward is a paid mutator transaction binding the contract method 0xcdd0c50e.
//
// Solidity: function updateEpochReward(uint256 epochIndex, address[] sequencers, uint256[] delegatorRewards, uint256[] commissions) returns()
func (_Distribute *DistributeTransactor) UpdateEpochReward(opts *bind.TransactOpts, epochIndex *big.Int, sequencers []common.Address, delegatorRewards []*big.Int, commissions []*big.Int) (*types.Transaction, error) {
	return _Distribute.contract.Transact(opts, "updateEpochReward", epochIndex, sequencers, delegatorRewards, commissions)
}

// UpdateEpochReward is a paid mutator transaction binding the contract method 0xcdd0c50e.
//
// Solidity: function updateEpochReward(uint256 epochIndex, address[] sequencers, uint256[] delegatorRewards, uint256[] commissions) returns()
func (_Distribute *DistributeSession) UpdateEpochReward(epochIndex *big.Int, sequencers []common.Address, delegatorRewards []*big.Int, commissions []*big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.UpdateEpochReward(&_Distribute.TransactOpts, epochIndex, sequencers, delegatorRewards, commissions)
}

// UpdateEpochReward is a paid mutator transaction binding the contract method 0xcdd0c50e.
//
// Solidity: function updateEpochReward(uint256 epochIndex, address[] sequencers, uint256[] delegatorRewards, uint256[] commissions) returns()
func (_Distribute *DistributeTransactorSession) UpdateEpochReward(epochIndex *big.Int, sequencers []common.Address, delegatorRewards []*big.Int, commissions []*big.Int) (*types.Transaction, error) {
	return _Distribute.Contract.UpdateEpochReward(&_Distribute.TransactOpts, epochIndex, sequencers, delegatorRewards, commissions)
}

// DistributeCommissionClaimedIterator is returned from FilterCommissionClaimed and is used to iterate over the raw logs and unpacked data for CommissionClaimed events raised by the Distribute contract.
type DistributeCommissionClaimedIterator struct {
	Event *DistributeCommissionClaimed // Event containing the contract specifics and raw log

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
func (it *DistributeCommissionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributeCommissionClaimed)
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
		it.Event = new(DistributeCommissionClaimed)
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
func (it *DistributeCommissionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributeCommissionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributeCommissionClaimed represents a CommissionClaimed event raised by the Distribute contract.
type DistributeCommissionClaimed struct {
	Delegatee common.Address
	UpToEpoch *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionClaimed is a free log retrieval operation binding the contract event 0xe4760bd616775d8b0ae78f9b8bfa4b453fdde769d10a0559ba473157ce9011d4.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) FilterCommissionClaimed(opts *bind.FilterOpts, delegatee []common.Address) (*DistributeCommissionClaimedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Distribute.contract.FilterLogs(opts, "CommissionClaimed", delegateeRule)
	if err != nil {
		return nil, err
	}
	return &DistributeCommissionClaimedIterator{contract: _Distribute.contract, event: "CommissionClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionClaimed is a free log subscription operation binding the contract event 0xe4760bd616775d8b0ae78f9b8bfa4b453fdde769d10a0559ba473157ce9011d4.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) WatchCommissionClaimed(opts *bind.WatchOpts, sink chan<- *DistributeCommissionClaimed, delegatee []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Distribute.contract.WatchLogs(opts, "CommissionClaimed", delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributeCommissionClaimed)
				if err := _Distribute.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
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

// ParseCommissionClaimed is a log parse operation binding the contract event 0xe4760bd616775d8b0ae78f9b8bfa4b453fdde769d10a0559ba473157ce9011d4.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) ParseCommissionClaimed(log types.Log) (*DistributeCommissionClaimed, error) {
	event := new(DistributeCommissionClaimed)
	if err := _Distribute.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Distribute contract.
type DistributeInitializedIterator struct {
	Event *DistributeInitialized // Event containing the contract specifics and raw log

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
func (it *DistributeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributeInitialized)
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
		it.Event = new(DistributeInitialized)
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
func (it *DistributeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributeInitialized represents a Initialized event raised by the Distribute contract.
type DistributeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Distribute *DistributeFilterer) FilterInitialized(opts *bind.FilterOpts) (*DistributeInitializedIterator, error) {

	logs, sub, err := _Distribute.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DistributeInitializedIterator{contract: _Distribute.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Distribute *DistributeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DistributeInitialized) (event.Subscription, error) {

	logs, sub, err := _Distribute.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributeInitialized)
				if err := _Distribute.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Distribute *DistributeFilterer) ParseInitialized(log types.Log) (*DistributeInitialized, error) {
	event := new(DistributeInitialized)
	if err := _Distribute.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Distribute contract.
type DistributeOwnershipTransferredIterator struct {
	Event *DistributeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistributeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributeOwnershipTransferred)
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
		it.Event = new(DistributeOwnershipTransferred)
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
func (it *DistributeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributeOwnershipTransferred represents a OwnershipTransferred event raised by the Distribute contract.
type DistributeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distribute *DistributeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistributeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distribute.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistributeOwnershipTransferredIterator{contract: _Distribute.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distribute *DistributeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistributeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distribute.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributeOwnershipTransferred)
				if err := _Distribute.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Distribute *DistributeFilterer) ParseOwnershipTransferred(log types.Log) (*DistributeOwnershipTransferred, error) {
	event := new(DistributeOwnershipTransferred)
	if err := _Distribute.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributeRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Distribute contract.
type DistributeRewardClaimedIterator struct {
	Event *DistributeRewardClaimed // Event containing the contract specifics and raw log

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
func (it *DistributeRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributeRewardClaimed)
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
		it.Event = new(DistributeRewardClaimed)
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
func (it *DistributeRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributeRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributeRewardClaimed represents a RewardClaimed event raised by the Distribute contract.
type DistributeRewardClaimed struct {
	Delegator common.Address
	Delegatee common.Address
	UpToEpoch *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x7a84a08b02c91f3c62d572853f966fc799bbd121e8ad7833a4494ab8dcfcb404.
//
// Solidity: event RewardClaimed(address indexed delegator, address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) FilterRewardClaimed(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*DistributeRewardClaimedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Distribute.contract.FilterLogs(opts, "RewardClaimed", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &DistributeRewardClaimedIterator{contract: _Distribute.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x7a84a08b02c91f3c62d572853f966fc799bbd121e8ad7833a4494ab8dcfcb404.
//
// Solidity: event RewardClaimed(address indexed delegator, address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *DistributeRewardClaimed, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Distribute.contract.WatchLogs(opts, "RewardClaimed", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributeRewardClaimed)
				if err := _Distribute.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x7a84a08b02c91f3c62d572853f966fc799bbd121e8ad7833a4494ab8dcfcb404.
//
// Solidity: event RewardClaimed(address indexed delegator, address indexed delegatee, uint256 upToEpoch, uint256 amount)
func (_Distribute *DistributeFilterer) ParseRewardClaimed(log types.Log) (*DistributeRewardClaimed, error) {
	event := new(DistributeRewardClaimed)
	if err := _Distribute.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
