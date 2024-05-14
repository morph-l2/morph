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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upToEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upToEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECORD_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetEpochIndex\",\"type\":\"uint256\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainsNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"newDelegation\",\"type\":\"bool\"}],\"name\":\"notifyDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainsNumber\",\"type\":\"uint256\"}],\"name\":\"notifyUndelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"queryUnclaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"unclaimedCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"delegatorRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"commissions\",\"type\":\"uint256[]\"}],\"name\":\"updateEpochReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5073530000000000000000000000000000000000001360805273530000000000000000000000000000000000001560a05273530000000000000000000000000000000000001260c05260805160a05160c0516120dc6100c15f395f81816101ec0152610d1401525f818161015b01528181610271015281816104a9015281816109b201528181610ac60152610f4b01525f8181610226015281816111550152818161120201526112a901526120dc5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c80639889be5111610093578063cdd0c50e11610063578063cdd0c50e1461020e578063d557714114610221578063e16bcc3214610248578063f2fde38b1461025b575f80fd5b80639889be51146101ae578063996cba68146101c1578063ad8e1223146101d4578063cd4281d0146101e7575f80fd5b80637ac3339a116100ce5780637ac3339a14610124578063807de443146101565780638129fc1c146101955780638da5cb5b1461019d575f80fd5b806341302560146100f45780635cf20c7b14610109578063715018a61461011c575b5f80fd5b610107610102366004611ce8565b61026e565b005b610107610117366004611ce8565b6104a6565b610107610675565b610143610132366004611d10565b60676020525f908152604090205481565b6040519081526020015b60405180910390f35b61017d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161014d565b610107610688565b6033546001600160a01b031661017d565b6101436101bc366004611d29565b6107fa565b6101076101cf366004611d5a565b6109af565b6101076101e2366004611d93565b610ac3565b61017d7f000000000000000000000000000000000000000000000000000000000000000081565b61010761021c366004611e24565b610d11565b61017d7f000000000000000000000000000000000000000000000000000000000000000081565b610107610256366004611ecd565b610f48565b610107610269366004611d10565b61108e565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146102eb5760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064015b60405180910390fd5b6065545f0361033c5760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e7420796574000000000000000000000000000000000000000060448201526064016102e2565b8080158061034b575060655482115b1561036257600160655461035f9190611f63565b90505b6001600160a01b0383165f908152606760205260409020548110156103c95760405162461bcd60e51b815260206004820152601660248201527f616c6c20636f6d6d697373696f6e20636c61696d65640000000000000000000060448201526064016102e2565b6001600160a01b0383165f908152606760205260408120545b82811161042d576001600160a01b0385165f9081526066602090815260408083208484529091529020600101546104199083611f76565b91508061042581611f89565b9150506103e2565b50801561043e5761043e848261111b565b610449826001611f76565b6001600160a01b0385165f8181526067602090815260409182902093909355805185815292830184905290917fe4760bd616775d8b0ae78f9b8bfa4b453fdde769d10a0559ba473157ce9011d4910160405180910390a250505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161461051e5760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064016102e2565b6065545f0361056f5760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e7420796574000000000000000000000000000000000000000060448201526064016102e2565b8080158061057e575060655482115b156105955760016065546105929190611f63565b90505b5f805b6001600160a01b0385165f9081526068602052604090206105b890611379565b81101561065e576001600160a01b0385165f9081526068602052604081206105e09083611388565b6001600160a01b0387165f908152606860205260409020909150610604908261139a565b801561063857506001600160a01b038087165f90815260686020908152604080832093851683526003909301905220548410155b15610655576106488187866113bb565b6106529084611f76565b92505b50600101610598565b50801561066f5761066f848261111b565b50505050565b61067d611956565b6106865f6119b0565b565b5f54610100900460ff16158080156106a657505f54600160ff909116105b806106bf5750303b1580156106bf57505f5460ff166001145b6107315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102e2565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561078d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610795611a19565b80156107f7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b038082165f9081526068602090815260408083209386168352600390930190529081205481908190805b6065548110156109a5576001600160a01b038088165f9081526066602090815260408083208584528252808320938a1683526006909301905220541561089f576001600160a01b038088165f9081526066602090815260408083208584528252808320938a16835260069093019052205492505b6001600160a01b0387165f908152606660209081526040808320848452909152902060020154156108f3576001600160a01b0387165f90815260666020908152604080832084845290915290206002015493505b6001600160a01b0387165f9081526066602090815260408083208484529091529020548490610923908590611fc0565b61092d9190611fd7565b6109379086611f76565b6001600160a01b038088165f908152606860209081526040808320938c16835260029093019052205490955060ff16801561099957506001600160a01b038087165f908152606860209081526040808320938b16835260049093019052205481145b6109a55760010161082b565b5050505092915050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610a275760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064016102e2565b6065545f03610a785760405162461bcd60e51b815260206004820152600c60248201527f6e6f74206d696e7420796574000000000000000000000000000000000000000060448201526064016102e2565b80801580610a87575060655482115b15610a9e576001606554610a9b9190611f63565b90505b5f610aaa8585846113bb565b90508015610abc57610abc848261111b565b5050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610b3b5760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064016102e2565b6001600160a01b0385165f908152606660209081526040808320868452909152902060028101839055600301819055821580610b9e57506001600160a01b038085165f908152606860209081526040808320938916835260039093019052205483145b15610c83576001600160a01b0385165f9081526066602090815260408083208684529091529020610bd29060040185611a9d565b506001600160a01b038086165f90815260666020908152604080832087845282528083209388168352600690930181528282208290556068905220610c179086611a9d565b506001600160a01b038481165f908152606860209081526040808320938916835260028401825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556003840182528083208390556004909301905290812055610abc565b6001600160a01b038085165f9081526068602090815260408083209389168352600290930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155610ce19084611f63565b6001600160a01b038086165f908152606860209081526040808320938a1683526004909301905220555050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610d895760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c79207265636f726420636f6e747261637420616c6c6f7765640000000060448201526064016102e2565b60658054905f610d9883611f89565b9190505550866001606554610dad9190611f63565b14610dfa5760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642065706f636820696e6465780000000000000000000000000060448201526064016102e2565b8285148015610e0857508085145b610e545760405162461bcd60e51b815260206004820152601360248201527f696e76616c69642064617461206c656e6774680000000000000000000000000060448201526064016102e2565b5f5b85811015610f3e57848482818110610e7057610e7061200f565b9050602002013560665f898985818110610e8c57610e8c61200f565b9050602002016020810190610ea19190611d10565b6001600160a01b0316815260208082019290925260409081015f9081208c8252909252902055828282818110610ed957610ed961200f565b9050602002013560665f898985818110610ef557610ef561200f565b9050602002016020810190610f0a9190611d10565b6001600160a01b0316815260208082019290925260409081015f9081208c8252909252902060019081019190915501610e56565b5050505050505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610fc05760405162461bcd60e51b815260206004820181905260248201527f6f6e6c79206c32207374616b696e6720636f6e747261637420616c6c6f77656460448201526064016102e2565b6001600160a01b0387165f90815260666020908152604080832088845290915290206002810184905560038101839055610ffd9060040187611ab1565b506001600160a01b038088165f9081526066602090815260408083208984528252808320938a16835260069093019052208490558015611085576001600160a01b0386165f9081526068602052604090206110589088611ab1565b506001600160a01b038087165f908152606860209081526040808320938b16835260039093019052208590555b50505050505050565b611096611956565b6001600160a01b0381166111125760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102e2565b6107f7816119b0565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa15801561119c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c0919061203c565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303815f875af115801561124a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061126e9190612053565b506040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa1580156112f0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611314919061203c565b90505f8311801561132d57508261132b8383611f63565b145b61066f5760405162461bcd60e51b815260206004820152601b60248201527f6d6f72706820746f6b656e207472616e73666572206661696c6564000000000060448201526064016102e2565b5f611382825490565b92915050565b5f6113938383611ac5565b9392505050565b6001600160a01b0381165f9081526001830160205260408120541515611393565b6001600160a01b0382165f9081526068602052604081206113dc908561139a565b6114285760405162461bcd60e51b815260206004820152601360248201527f6e6f2072656d61696e696e67207265776172640000000000000000000000000060448201526064016102e2565b6001600160a01b038084165f90815260686020908152604080832093881683526003909301905220548210156114a05760405162461bcd60e51b815260206004820152601260248201527f616c6c2072657761726420636c61696d6564000000000000000000000000000060448201526064016102e2565b6001600160a01b038084165f90815260686020908152604080832093881683526003909301905220545b8281116118f8576001600160a01b038086165f9081526066602090815260408083208584528083528184206002810154958a16855260068101845291842054938690529091525461151b9190611fc0565b6115259190611fd7565b61152f9083611f76565b6001600160a01b0386165f90815260666020526040812091935061157891869161155a856001611f76565b81526020019081526020015f2060040161139a90919063ffffffff16565b611704576001600160a01b0385165f9081526066602052604081206115c3918691906115a5856001611f76565b81526020019081526020015f20600401611ab190919063ffffffff16565b506001600160a01b038086165f818152606660208181526040808420878552808352818520968b1685526006909601825283205493835252909190611609846001611f76565b815260208082019290925260409081015f9081206001600160a01b03808a168352600690910184528282209490945592881683526066909152812090611650836001611f76565b81526020019081526020015f20600201545f03611704576001600160a01b0385165f81815260666020818152604080842086855280835290842060020154948452919052906116a0846001611f76565b815260208082019290925260409081015f908120600201939093556001600160a01b03881680845260668084528285208686528085529285206003015491855290925290916116f0846001611f76565b815260208101919091526040015f20600301555b6001600160a01b0385165f90815260666020908152604080832084845290915281206003018054916117358361206e565b90915550506001600160a01b0385165f90815260666020908152604080832084845290915281206003015490036117b9576001600160a01b0385165f908152606660209081526040808320848452909152812081815560018101829055600281018290556003810182905590600482018181816117b28282611c9f565b5050505050505b6001600160a01b038085165f908152606860209081526040808320938916835260029093019052205460ff16801561181857506001600160a01b038085165f908152606860209081526040808320938916835260049093019052205481145b156118aa576001600160a01b0384165f90815260686020526040902061183e9086611a9d565b506001600160a01b038481165f908152606860209081526040808320938916835260028401825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560038401825280832083905560049093019052908120556118f8565b6001600160a01b038085165f9081526068602090815260408083209389168352600390930190529081208054916118e083611f89565b919050555080806118f090611f89565b9150506114ca565b50836001600160a01b0316836001600160a01b03167f7a84a08b02c91f3c62d572853f966fc799bbd121e8ad7833a4494ab8dcfcb4048484604051611947929190918252602082015260400190565b60405180910390a39392505050565b6033546001600160a01b031633146106865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e2565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611a955760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102e2565b610686611aeb565b5f611393836001600160a01b038416611b70565b5f611393836001600160a01b038416611c53565b5f825f018281548110611ada57611ada61200f565b905f5260205f200154905092915050565b5f54610100900460ff16611b675760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102e2565b610686336119b0565b5f8181526001830160205260408120548015611c4a575f611b92600183611f63565b85549091505f90611ba590600190611f63565b9050818114611c04575f865f018281548110611bc357611bc361200f565b905f5260205f200154905080875f018481548110611be357611be361200f565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611c1557611c156120a2565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611382565b5f915050611382565b5f818152600183016020526040812054611c9857508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611382565b505f611382565b5080545f8255905f5260205f20908101906107f791905b80821115611cc9575f8155600101611cb6565b5090565b80356001600160a01b0381168114611ce3575f80fd5b919050565b5f8060408385031215611cf9575f80fd5b611d0283611ccd565b946020939093013593505050565b5f60208284031215611d20575f80fd5b61139382611ccd565b5f8060408385031215611d3a575f80fd5b611d4383611ccd565b9150611d5160208401611ccd565b90509250929050565b5f805f60608486031215611d6c575f80fd5b611d7584611ccd565b9250611d8360208501611ccd565b9150604084013590509250925092565b5f805f805f60a08688031215611da7575f80fd5b611db086611ccd565b9450611dbe60208701611ccd565b94979496505050506040830135926060810135926080909101359150565b5f8083601f840112611dec575f80fd5b50813567ffffffffffffffff811115611e03575f80fd5b6020830191508360208260051b8501011115611e1d575f80fd5b9250929050565b5f805f805f805f6080888a031215611e3a575f80fd5b87359650602088013567ffffffffffffffff80821115611e58575f80fd5b611e648b838c01611ddc565b909850965060408a0135915080821115611e7c575f80fd5b611e888b838c01611ddc565b909650945060608a0135915080821115611ea0575f80fd5b50611ead8a828b01611ddc565b989b979a50959850939692959293505050565b80151581146107f7575f80fd5b5f805f805f805f60e0888a031215611ee3575f80fd5b611eec88611ccd565b9650611efa60208901611ccd565b955060408801359450606088013593506080880135925060a0880135915060c0880135611f2681611ec0565b8091505092959891949750929550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561138257611382611f36565b8082018082111561138257611382611f36565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611fb957611fb9611f36565b5060010190565b808202811582820484141761138257611382611f36565b5f8261200a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020828403121561204c575f80fd5b5051919050565b5f60208284031215612063575f80fd5b815161139381611ec0565b5f8161207c5761207c611f36565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c6343000818000a",
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
