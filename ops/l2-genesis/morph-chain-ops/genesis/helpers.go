package genesis

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
)

var (
	// codeNamespace represents the namespace of implementations of predeploys
	codeNamespace = common.HexToAddress("0xc0D3C0d3C0d3C0D3c0d3C0d3c0D3C0d3c0d30000")
	// l2PredeployNamespace_53 represents the namespace of L2 predeploys
	l2PredeployNamespace_53 = common.HexToAddress("0x5300000000000000000000000000000000000000")
	// bigL2PredeployNamespace_53 represents the predeploy namespace as a big.Int
	bigL2PredeployNamespace_53 = new(big.Int).SetBytes(l2PredeployNamespace_53.Bytes())
	// bigCodeNamespace represents the predeploy namespace as a big.Int
	bigCodeNameSpace = new(big.Int).SetBytes(codeNamespace.Bytes())
	// implementationSlot represents the EIP 1967 implementation storage slot
	ImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
	// implementationSlot represents the EIP 1967 admin storage slot
	AdminSlot = common.HexToHash("0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103")
)

// DevAccounts represent the standard hardhat development accounts.
// These are funded if the deploy config has funding development
// accounts enabled.
var DevAccounts = []common.Address{
	common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"),
	common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
	common.HexToAddress("0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec"),
	common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"),
	common.HexToAddress("0x2546BcD3c84621e976D8185a91A922aE77ECEc30"),
	common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
	common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
	common.HexToAddress("0x71bE63f3384f5fb98995898A86B02Fb2426c5788"),
	common.HexToAddress("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"),
	common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"),
	common.HexToAddress("0x976EA74026E726554dB657fA54763abd0C3a0aa9"),
	common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"),
	common.HexToAddress("0xBcd4042DE499D14e55001CcbB24a551F3b954096"),
	common.HexToAddress("0xFABB0ac9d68B0B445fB7357272Ff202C5651694a"),
	common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"),
	common.HexToAddress("0xbDA5747bFD65F08deb54cb465eB87D40e51B197E"),
	common.HexToAddress("0xcd3B766CCDd6AE721141F452C550Ca635964ce71"),
	common.HexToAddress("0xdD2FD4581271e230360230F9337D5c0430Bf44C0"),
	common.HexToAddress("0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097"),
	common.HexToAddress("0xde3829a23df1479438622a08a116e8eb3f620bb5"),
	common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
	// Test account used by geth tests
	common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7"),
	common.HexToAddress("0xca062B0Fd91172d89BcD4Bb084ac4e21972CC467"),
}

// The devBalance is the amount of wei that a dev account is funded with.
var devBalance = hexutil.MustDecodeBig("0x20000000000000000000000000000000000000000000000000000000000000")

// The lockedBalance is the amount of wei that the account who holds the total locked ETH in layer2 is funded with.
var lockedBalance = hexutil.MustDecodeBig("0xa56fa5b99019a5c8000000") // 200M

// AddressToCodeNamespace takes a predeploy address and computes
// the implementation address that the implementation should be deployed at
func AddressToCodeNamespace(addr common.Address) (common.Address, error) {
	if !IsL2DevPredeploy(addr) {
		return common.Address{}, fmt.Errorf("cannot handle non predeploy: %s", addr)
	}
	bigAddress := new(big.Int).SetBytes(addr[16:])
	num := new(big.Int).Or(bigCodeNameSpace, bigAddress)
	return common.BigToAddress(num), nil
}

func IsL2DevPredeploy(addr common.Address) bool {
	return bytes.Equal(addr[0:2], []byte{0x53, 0x00})
}

func newHexBig(in uint64) *hexutil.Big {
	b := new(big.Int).SetUint64(in)
	hb := hexutil.Big(*b)
	return &hb
}
