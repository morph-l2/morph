package blocktag

import (
	"fmt"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/urfave/cli"

	node "morph-l2/node/core"
	"morph-l2/node/flags"
)

const (
	// DefaultSafeConfirmations is the default number of L1 blocks to wait before considering a batch as safe
	DefaultSafeConfirmations = 10
	// DefaultPollInterval is the default interval to poll L1 for batch status updates
	DefaultPollInterval = 10 * time.Second
)

// Config holds the configuration for BlockTagService
type Config struct {
	RollupAddress     common.Address
	SafeConfirmations uint64
	PollInterval      time.Duration
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		SafeConfirmations: DefaultSafeConfirmations,
		PollInterval:      DefaultPollInterval,
	}
}

// SetCliContext sets the configuration from CLI context
func (c *Config) SetCliContext(ctx *cli.Context) error {
	// Determine RollupAddress: use explicit flag, or mainnet default, or error
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
