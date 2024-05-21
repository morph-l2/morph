package genesis

import (
	"math/big"
	"testing"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	node "morph-l2/node/core"
)

func Test_BuildL2DeveloperGenesis(t *testing.T) {
	address1 := common.BigToAddress(common.Big1)

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
		L1ETHGatewayProxy:           address1,
		L1ERC721GatewayProxy:        address1,
		L1ERC1155GatewayProxy:       address1,
		L1WETHGatewayProxy:          address1,
		L1WETH:                      address1,

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
		L2StakingRewardStartTime:       uint64(time.Now().Unix()),

		GovVotingDuration:     1000,
		GovBatchBlockInterval: 20,
		GovBatchMaxBytes:      124928,
		GovRollupEpoch:        600,
		GovBatchTimeout:       100,
		GovBatchMaxChunks:     15,

		// MorphToken
		MorphTokenName:               "Morph Token",
		MorphTokenSymbol:             "Morph",
		MorphTokenInitialSupply:      1,
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
