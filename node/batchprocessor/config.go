package batchprocessor

import (
	"fmt"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/urfave/cli"

	node "morph-l2/node/core"
	"morph-l2/node/flags"
)

const (
	DefaultPollInterval      = 12 * time.Second
	DefaultSafeConfirmations = 10
)

type Config struct {
	L1Addr            string
	RollupAddress     common.Address
	SafeConfirmations uint64
	PollInterval      time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		SafeConfirmations: DefaultSafeConfirmations,
		PollInterval:      DefaultPollInterval,
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	c.L1Addr = ctx.GlobalString(flags.L1NodeAddr.Name)

	if ctx.GlobalBool(flags.MainnetFlag.Name) {
		c.RollupAddress = node.MainnetRollupContractAddress
	} else if ctx.GlobalIsSet(flags.RollupContractAddress.Name) {
		c.RollupAddress = common.HexToAddress(ctx.GlobalString(flags.RollupContractAddress.Name))
	} else {
		return fmt.Errorf("rollup contract address is required: either specify --%s or use --%s for mainnet default",
			flags.RollupContractAddress.Name, flags.MainnetFlag.Name)
	}

	if ctx.GlobalIsSet(flags.BlockTagSafeConfirmations.Name) {
		c.SafeConfirmations = ctx.GlobalUint64(flags.BlockTagSafeConfirmations.Name)
	}
	return nil
}
