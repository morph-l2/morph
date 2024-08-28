package sync

import (
	"errors"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
	"morph-l2/node/types"
)

const (
	// DefaultFetchBlockRange is the number of blocks that we collect in a single eth_getLogs query.
	DefaultFetchBlockRange = uint64(100)

	// DefaultPollInterval is the frequency at which we query for new L1 messages.
	DefaultPollInterval = time.Second * 15

	// DefaultLogProgressInterval is the frequency at which we log progress.
	DefaultLogProgressInterval = time.Second * 10
)

type Config struct {
	L1                    *types.L1Config `json:"l1"`
	L1MessageQueueAddress *common.Address `json:"l1_message_queue_address"`
	StartHeight           uint64          `json:"start_height"`
	PollInterval          time.Duration   `json:"poll_interval"`
	LogProgressInterval   time.Duration   `json:"log_progress_interval"`
	FetchBlockRange       uint64          `json:"fetch_block_range"`
}

func DefaultConfig() *Config {
	return &Config{
		L1: &types.L1Config{
			Confirmations: rpc.FinalizedBlockNumber,
		},
		PollInterval:        DefaultPollInterval,
		LogProgressInterval: DefaultLogProgressInterval,
		FetchBlockRange:     DefaultFetchBlockRange,
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	c.L1.Addr = ctx.GlobalString(flags.L1NodeAddr.Name)
	if ctx.GlobalIsSet(flags.L1Confirmations.Name) {
		c.L1.Confirmations = rpc.BlockNumber(ctx.GlobalInt64(flags.L1Confirmations.Name))
	}

	if ctx.GlobalIsSet(flags.SyncDepositContractAddr.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.SyncDepositContractAddr.Name))
		c.L1MessageQueueAddress = &addr
		if len(c.L1MessageQueueAddress.Bytes()) == 0 {
			return errors.New("invalid SyncDepositContractAddr")
		}
	}

	if ctx.GlobalIsSet(flags.SyncStartHeight.Name) {
		c.StartHeight = ctx.GlobalUint64(flags.SyncStartHeight.Name)
		if c.StartHeight == 0 {
			return errors.New("invalid SyncStartHeight")
		}
	}

	if ctx.GlobalIsSet(flags.SyncPollInterval.Name) {
		c.PollInterval = ctx.GlobalDuration(flags.SyncPollInterval.Name)
		if c.PollInterval == 0 {
			return errors.New("invalid pollInterval")
		}
	}
	if ctx.GlobalIsSet(flags.SyncLogProgressInterval.Name) {
		c.LogProgressInterval = ctx.GlobalDuration(flags.SyncLogProgressInterval.Name)
		if c.LogProgressInterval == 0 {
			return errors.New("invalid logProgressInterval")
		}
	}
	if ctx.GlobalIsSet(flags.SyncFetchBlockRange.Name) {
		c.FetchBlockRange = ctx.GlobalUint64(flags.SyncFetchBlockRange.Name)
		if c.FetchBlockRange == 0 {
			return errors.New("invalid fetchBlockRange")
		}
	}

	return nil
}
