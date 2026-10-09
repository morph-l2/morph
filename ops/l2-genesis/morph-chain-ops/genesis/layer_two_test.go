package genesis

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind/backends"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
)

func testGenesisConfig() *DeployConfig {
	address1 := common.BigToAddress(common.Big1)
	startTime := uint64((time.Now().Unix()/86400 + 1) * 86400)
	return &DeployConfig{
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
		GovRollupEpoch:        600,
		GovBatchTimeout:       100,

		FundDevAccounts: true,
	}
}

func Test_BuildL2DeveloperGenesis(t *testing.T) {
	config := testGenesisConfig()
	curHeader := &types.Header{}
	curHeader.BaseFee = new(big.Int).SetUint64(1)
	l2Genesis, _, err := BuildL2DeveloperGenesis(config, nil, curHeader)
	require.NoError(t, err)

	l2GenesisBlock := l2Genesis.ToBlock(nil)
	genesisBatchHeaderBytes, err := GenesisBatchHeader(l2GenesisBlock.Header())
	require.NoError(t, err)
	t.Logf("generated genesis batch header bytes: %x \n", genesisBatchHeaderBytes)
}

func TestGenesisRetainsConfiguredLegacyPredeployState(t *testing.T) {
	for _, changed := range []bool{false, true} {
		name := "original"
		if changed {
			name = "changed"
		}
		t.Run(name, func(t *testing.T) {
			config := testGenesisConfig()
			if changed {
				config.GovVotingDuration = 1017
				config.GovBatchBlockInterval = 23
				config.GovBatchTimeout = 0
				config.GovRollupEpoch = 701
				config.RecordOracleAddress = common.BigToAddress(big.NewInt(2))
				config.RecordNextBatchSubmissionIndex = 17
				config.L2StakingSequencerMaxSize = 2
				config.L2StakingUnDelegatedLockEpochs = 3
				config.L2StakingRewardStartTime += 86400
				config.L2StakingAddresses[0] = common.BigToAddress(big.NewInt(3))
				config.L2StakingTmKeys[0] = common.BigToHash(big.NewInt(4))
				g2 := bls12381.NewG2()
				config.L2StakingBlsKeys[0] = g2.EncodePoint(g2.One())
			}
			generated, _, err := BuildL2DeveloperGenesis(config, nil, &types.Header{BaseFee: big.NewInt(1)})
			require.NoError(t, err)
			// Run the actual ABI getters against the generated allocation, including
			// the proxy implementations. This checks persisted state, not the input map.
			backend := backends.NewSimulatedBackend(generated.Alloc, 30000000)
			defer backend.Close()
			gov, err := bindings.NewGovCaller(predeploys.GovAddr, backend)
			require.NoError(t, err)
			record, err := bindings.NewRecordCaller(predeploys.RecordAddr, backend)
			require.NoError(t, err)
			staking, err := bindings.NewL2StakingCaller(predeploys.L2StakingAddr, backend)
			require.NoError(t, err)
			sequencer, err := bindings.NewSequencerCaller(predeploys.SequencerAddr, backend)
			require.NoError(t, err)
			for _, check := range []struct {
				name string
				read func(*bind.CallOpts) (*big.Int, error)
				want uint64
			}{
				{"Gov.votingDuration", gov.VotingDuration, config.GovVotingDuration},
				{"Gov.batchBlockInterval", gov.BatchBlockInterval, config.GovBatchBlockInterval},
				{"Gov.batchTimeout", gov.BatchTimeout, config.GovBatchTimeout},
				{"Gov.rollupEpoch", gov.RollupEpoch, config.GovRollupEpoch},
				{"Record.nextBatchSubmissionIndex", record.NextBatchSubmissionIndex, config.RecordNextBatchSubmissionIndex},
				{"L2Staking.sequencerSetMaxSize", staking.SequencerSetMaxSize, config.L2StakingSequencerMaxSize},
				{"L2Staking.undelegateLockEpochs", staking.UndelegateLockEpochs, config.L2StakingUnDelegatedLockEpochs},
				{"L2Staking.rewardStartTime", staking.RewardStartTime, config.L2StakingRewardStartTime},
			} {
				value, err := check.read(nil)
				require.NoError(t, err, check.name)
				require.Equal(t, check.want, value.Uint64(), check.name)
			}
			oracle, err := record.Oracle(nil)
			require.NoError(t, err)
			require.Equal(t, config.RecordOracleAddress, oracle)
			for _, read := range []func(*bind.CallOpts) ([]common.Address, error){
				sequencer.GetSequencerSet0, sequencer.GetSequencerSet1, sequencer.GetSequencerSet2,
			} {
				addresses, err := read(nil)
				require.NoError(t, err)
				require.Equal(t, config.L2StakingAddresses, addresses)
			}
			stakers, err := staking.GetStakesInfo(nil, config.L2StakingAddresses)
			require.NoError(t, err)
			require.Len(t, stakers, len(config.L2StakingAddresses))
			for i, staker := range stakers {
				require.Equal(t, config.L2StakingAddresses[i], staker.Addr)
				require.Equal(t, [32]byte(config.L2StakingTmKeys[i]), staker.TmKey)
				require.Equal(t, []byte(config.L2StakingBlsKeys[i]), staker.BlsKey)
			}
		})
	}
}

func TestDeployConfigInitializerRequirements(t *testing.T) {
	for _, test := range []struct {
		name   string
		change func(*DeployConfig)
		valid  bool
	}{
		{"interval_disabled", func(c *DeployConfig) { c.GovBatchBlockInterval = 0 }, true},
		{"timeout_disabled", func(c *DeployConfig) { c.GovBatchTimeout = 0 }, true},
		{"both_batch_triggers_disabled", func(c *DeployConfig) { c.GovBatchBlockInterval = 0; c.GovBatchTimeout = 0 }, false},
		{"empty_validator_set", func(c *DeployConfig) { c.L2StakingAddresses = nil; c.L2StakingTmKeys = nil; c.L2StakingBlsKeys = nil }, false},
		{"missing_tm_key", func(c *DeployConfig) { c.L2StakingTmKeys = nil }, false},
		{"missing_bls_key", func(c *DeployConfig) { c.L2StakingBlsKeys = nil }, false},
		{"reward_time_not_epoch_aligned", func(c *DeployConfig) { c.L2StakingRewardStartTime++ }, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			config := testGenesisConfig()
			test.change(config)
			err := config.Check()
			if test.valid {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, ErrInvalidDeployConfig)
			}
		})
	}
	for _, input := range []string{`{"l2StakingAddresses":["0x1234"]}`, `{"l2StakingTmKeys":["0x1234"]}`} {
		var config DeployConfig
		require.Error(t, json.Unmarshal([]byte(input), &config), "fixed-width address and Tendermint keys must reject malformed JSON")
	}
}

func TestL2ImmutableConfigRejectsLegacyStakingPlaceholders(t *testing.T) {
	for _, address := range []common.Address{{}, common.HexToAddress("0x000000000000000000000000000000000000dEaD")} {
		config := testGenesisConfig()
		config.L1StakingProxy = address
		_, _, err := NewL2ImmutableConfig(config)
		require.ErrorIs(t, err, ErrInvalidImmutablesConfig)
	}
	_, _, err := NewL2ImmutableConfig(testGenesisConfig())
	require.NoError(t, err, "offline genesis builders accept configured non-placeholder addresses without RPC")
}

func TestRollupConfigMatchesGeneratedL2Block(t *testing.T) {
	for _, explicit := range []bool{false, true} {
		name := "defaults"
		if explicit {
			name = "explicit"
		}
		t.Run(name, func(t *testing.T) {
			config := testGenesisConfig()
			if explicit {
				config.L2GenesisBlockGasLimit = 12345678
				config.L2GenesisBlockTimestamp = 1700000000
			}
			l1Block := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(2), Time: 100})
			generated, err := NewL2Genesis(config, l1Block)
			require.NoError(t, err)
			l2Block := generated.ToBlock(nil)
			if explicit {
				require.Equal(t, uint64(12345678), l2Block.GasLimit())
				require.Equal(t, uint64(1700000000), l2Block.Time())
			} else {
				require.Equal(t, uint64(defaultL2GasLimit), l2Block.GasLimit())
				require.NotZero(t, l2Block.Time())
			}
			// Once generated, block metadata must not be reconstructed from raw
			// configuration values or the current time.
			config.L2GenesisBlockGasLimit = hexutil.Uint64(l2Block.GasLimit() + 100)
			config.L2GenesisBlockTimestamp = hexutil.Uint64(l2Block.Time() + 100)
			config.L2GenesisBlockNumber = 17
			withdrawRoot := common.BigToHash(big.NewInt(19))
			header, err := GenesisBatchHeader(l2Block.Header())
			require.NoError(t, err)
			rollup, err := config.RollupConfig(l1Block, l2Block, withdrawRoot, header)
			require.NoError(t, err)
			require.Equal(t, l1Block.Hash(), rollup.Genesis.L1.Hash)
			require.Equal(t, l1Block.NumberU64(), rollup.Genesis.L1.Number)
			require.Equal(t, l2Block.Hash(), rollup.Genesis.L2.Hash)
			require.Equal(t, l2Block.NumberU64(), rollup.Genesis.L2.Number)
			require.Equal(t, l2Block.Root(), rollup.L2GenesisStateRoot)
			require.Equal(t, l2Block.Time(), rollup.Genesis.L2Time)
			require.NotEqual(t, l1Block.Time(), rollup.Genesis.L2Time)
			require.Equal(t, l2Block.GasLimit(), rollup.Genesis.SystemConfig.GasLimit)
			require.Equal(t, withdrawRoot, rollup.WithdrawRoot)
			require.Equal(t, hexutil.Bytes(header), rollup.GenesisBatchHeader)
		})
	}
}
