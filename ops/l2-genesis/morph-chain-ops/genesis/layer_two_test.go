package genesis

import (
	"math/big"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	node "morph-l2/node/core"
)

func Test_BuildL2DeveloperGenesis(t *testing.T) {
	address1 := common.BigToAddress(common.Big1)
	startTime := uint64((time.Now().Unix()/86400 + 1) * 86400)
	config := &DeployConfig{
		L1ChainID: 900,
		L2ChainID: 53077,

		BatchInboxAddress:  address1,
		BatchSenderAddress: address1,

		L1StakingProxy:              address1,
		ProxyAdminOwner:             address1,
		FinalSystemOwner:            address1,
		L1CrossDomainMessengerProxy: address1,
		RollupProxy:                 address1,
		L1GatewayRouterProxy:        address1,
		L1StandardERC20GatewayProxy: address1,
		L1CustomERC20GatewayProxy:   address1,
		L1ReverseCustomGatewayProxy: address1,
		L1ETHGatewayProxy:           address1,
		L1ERC721GatewayProxy:        address1,
		L1ERC1155GatewayProxy:       address1,
		L1WETHGatewayProxy:          address1,
		L1WETH:                      address1,
		L1USDC:                      address1,
		L1USDCGatewayProxy:          address1,
		L1WithdrawLockERC20Gateway:  address1,

		GasPriceOracleOverhead: 1,
		GasPriceOracleScalar:   1,
		GasPriceOracleOwner:    address1,

		L1FeeVaultRecipient:        address1,
		SequencerFeeVaultRecipient: address1,
		L2BridgeFeeVaultRecipient:  address1,

		RecordOracleAddress:            address1,
		RecordNextBatchSubmissionIndex: 1,

		// L2Staking configs
		L2StakingSequencerMaxSize:      1,
		L2StakingUnDelegatedLockEpochs: 1,
		L2StakingRewardStartTime:       startTime,
		L2StakingAddresses:             []common.Address{common.HexToAddress("0x783698dCDEBdc96785c5c60ED96113612bA09c2b")},
		L2StakingTmKeys:                []common.Hash{common.HexToHash("0x5280d0eee2a64d3ad29480d15ffd1b048ce5908f180b5ccd65cc3dcf00941abb")},
		L2StakingBlsKeys:               []hexutil.Bytes{hexutil.MustDecode("0x00000000000000000000000000000000095ad465c2895ee825c7d4f1b60a18734db57d4108369e47c6e3a94ee15846f825c06dad5d98f503bd31ece1d9f94b11000000000000000000000000000000000c5d6ba04bc9b9674dd2acbfc5caed3976c1b8be2ec90a03d78dffe924648b4fba82225aff43c744310c6a60185b75ac000000000000000000000000000000000fce6be001c871a11b9db1c6c15f0a6999de5646941a74486206dc784f0b3ffe11799212f3f44ef754b4a0f1ecf85639000000000000000000000000000000000b2f06634e5ea719682c30911c94dfb560f0b7656b5c34a871ea035e3fe7b041885420f8fe1e251f1cce5cdb7514869e")},

		GovVotingDuration:     1000,
		GovBatchBlockInterval: 20,
		GovBatchMaxBytes:      124928,
		GovRollupEpoch:        600,
		GovBatchTimeout:       100,
		GovBatchMaxChunks:     15,

		// MorphToken
		MorphTokenOwner:              address1,
		MorphTokenName:               "Morph Token",
		MorphTokenSymbol:             "Morph",
		MorphTokenInitialSupply:      1000000000,
		MorphTokenDailyInflationRate: 1,

		FundDevAccounts: true,
		MaxTxPerBlock:   1000,
	}
	curHeader := &types.Header{}
	curHeader.BaseFee = new(big.Int).SetUint64(1)
	l2Genesis, _, err := BuildL2DeveloperGenesis(config, nil, curHeader)
	require.NoError(t, err)

	l2GenesisBlock := l2Genesis.ToBlock(nil)
	genesisBatchHeader, err := node.GenesisBatchHeader(l2GenesisBlock.Header())
	require.NoError(t, err)
	t.Logf("generated genesis batch header bytes: %x \n", genesisBatchHeader.Encode())
}
