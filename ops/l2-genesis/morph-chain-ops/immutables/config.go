package immutables

import (
	"github.com/scroll-tech/go-ethereum/common"
)

type InitConfig struct {
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
