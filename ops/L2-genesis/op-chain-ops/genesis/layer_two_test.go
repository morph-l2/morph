package genesis

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func Test_BuildL2DeveloperGenesis(t *testing.T) {
	address1 := common.BigToAddress(common.Big1)

	config := &DeployConfig{
		L1ChainID: 900,
		L2ChainID: 53077,

		FinalizationPeriodSeconds: 1,

		ProxyAdminOwner:             address1,
		FinalSystemOwner:            address1,
		L1SequencerProxy:            address1,
		L1CrossDomainMessengerProxy: address1,
		RollupProxy:                 address1,
		GasPriceOracleOverhead:      1,
		GasPriceOracleScalar:        1,
		GasPriceOracleOwner:         address1,
		L1FeeVaultRecipient:         address1,
		SequencerFeeVaultRecipient:  address1,

		GovProposalInterval:   1000,
		GovBatchBlockInterval: 20,
		GovBatchMaxBytes:      124928,
		GovRollupEpoch:        600,
		GovBatchTimeout:       100,
		GovBatchMaxChunks:     15,

		FundDevAccounts: true,
		MaxTxPerBlock:   1000,
	}
	curHeader := &types.Header{}
	curHeader.BaseFee = new(big.Int).SetUint64(1)
	_, _, err := BuildL2DeveloperGenesis(config, nil, curHeader)
	require.NoError(t, err)
}
