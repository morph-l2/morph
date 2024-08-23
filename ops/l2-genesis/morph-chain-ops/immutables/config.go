package immutables

import (
	"github.com/morph-l2/go-ethereum/common"
)

type InitConfig struct {
	// USDC
	USDCTokenName     string         `json:"USDCTokenName"`
	USDCTokenSymbol   string         `json:"USDCTokenSymbol"`
	USDCTokenCurrency string         `json:"USDCTokenCurrency"`
	USDCTokenDecimals uint8          `json:"USDCTokenDecimals"`
	USDCMasterMinter  common.Address `json:"USDCTMasterMinter"`
	USDCPauser        common.Address `json:"USDCPauser"`
	USDCBlackLister   common.Address `json:"USDCBlackLister"`
	USDCOwner         common.Address `json:"USDCOwner"`

	// MorphToken
	MorphTokenOwner              common.Address `json:"morphTokenOwner"`
	MorphTokenName               string         `json:"morphTokenName"`
	MorphTokenSymbol             string         `json:"morphTokenSymbol"`
	MorphTokenInitialSupply      uint64         `json:"morphTokenInitialSupply"`
	MorphTokenDailyInflationRate uint64         `json:"morphTokenDailyInflationRate"`

	// L2Staking
	L2StakingOwner                common.Address   `json:"l2StakingOwner"`
	L2StakingSequencersMaxSize    uint64           `json:"l2StakingSequencersMaxSize"`
	L2StakingUnDelegateLockEpochs uint64           `json:"l2StakingUnDelegateLockEpochs"`
	L2StakingRewardStartTime      uint64           `json:"l2StakingRewardStartTime"`
	L2StakingAddresses            []common.Address `json:"l2StakingAddresses"`
	L2StakingTmKeys               []common.Hash    `json:"l2StakingTmKeys"`
	L2StakingBlsKeys              [][]byte         `json:"l2StakingBlsKeys"`
}
