package derivation

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
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
	L2                    *types.L2Config `json:"l2"`
	BeaconRpc             string          `json:"beacon_rpc"`
	RollupContractAddress common.Address  `json:"rollup_contract_address"`
	StartHeight           uint64          `json:"start_height"`
	PollInterval          time.Duration   `json:"poll_interval"`
	LogProgressInterval   time.Duration   `json:"log_progress_interval"`
	FetchBlockRange       uint64          `json:"fetch_block_range"`
	MetricsPort           uint64          `json:"metrics_port"`
	MetricsHostname       string          `json:"metrics_hostname"`
	MetricsServerEnable   bool            `json:"metrics_server_enable"`
}

func DefaultConfig() *Config {
	return &Config{
		L1: &types.L1Config{
			Confirmations: rpc.FinalizedBlockNumber,
		},
		PollInterval:        DefaultPollInterval,
		LogProgressInterval: DefaultLogProgressInterval,
		FetchBlockRange:     DefaultFetchBlockRange,
		L2:                  new(types.L2Config),
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	c.L1.Addr = ctx.GlobalString(flags.L1NodeAddr.Name)
	if ctx.GlobalIsSet(flags.L1Confirmations.Name) {
		c.L1.Confirmations = rpc.BlockNumber(ctx.GlobalInt64(flags.L1Confirmations.Name))
	}
	if ctx.GlobalIsSet(flags.RollupContractAddress.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.RollupContractAddress.Name))
		c.RollupContractAddress = addr
		if len(c.RollupContractAddress.Bytes()) == 0 {
			return errors.New("invalid DerivationDepositContractAddr")
		}
	}
	c.BeaconRpc = ctx.GlobalString(flags.L1BeaconAddr.Name)
	if c.BeaconRpc == "" {
		return errors.New("invalid L1BeaconAddr")
	}

	if ctx.GlobalIsSet(flags.DerivationStartHeight.Name) {
		c.StartHeight = ctx.GlobalUint64(flags.DerivationStartHeight.Name)
		if c.StartHeight == 0 {
			return errors.New("invalid DerivationStartHeight")
		}
	}

	if ctx.GlobalIsSet(flags.DerivationPollInterval.Name) {
		c.PollInterval = ctx.GlobalDuration(flags.DerivationPollInterval.Name)
		if c.PollInterval == 0 {
			return errors.New("invalid pollInterval")
		}
	}
	if ctx.GlobalIsSet(flags.DerivationLogProgressInterval.Name) {
		c.LogProgressInterval = ctx.GlobalDuration(flags.DerivationLogProgressInterval.Name)
		if c.LogProgressInterval == 0 {
			return errors.New("invalid logProgressInterval")
		}
	}
	if ctx.GlobalIsSet(flags.DerivationFetchBlockRange.Name) {
		c.FetchBlockRange = ctx.GlobalUint64(flags.DerivationFetchBlockRange.Name)
		if c.FetchBlockRange == 0 {
			return errors.New("invalid fetchBlockRange")
		}
	}

	l2EthAddr := ctx.GlobalString(flags.L2EthAddr.Name)
	l2EngineAddr := ctx.GlobalString(flags.L2EngineAddr.Name)
	fileName := ctx.GlobalString(flags.L2EngineJWTSecret.Name)
	var secret [32]byte
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		return fmt.Errorf("file-name of jwt secret is empty")
	}
	if data, err := os.ReadFile(filepath.Clean(fileName)); err == nil {
		jwtSecret := common.FromHex(strings.TrimSpace(string(data)))
		if len(jwtSecret) != 32 {
			return fmt.Errorf("invalid jwt secret in path %s, not 32 hex-formatted bytes", fileName)
		}
		copy(secret[:], jwtSecret)
	} else {
		if _, err := io.ReadFull(rand.Reader, secret[:]); err != nil {
			return fmt.Errorf("failed to generate jwt secret: %w", err)
		}
		if err := os.WriteFile(fileName, []byte(hexutil.Encode(secret[:])), 0600); err != nil {
			return err
		}
	}
	c.L2.EthAddr = l2EthAddr
	c.L2.EngineAddr = l2EngineAddr
	c.L2.JwtSecret = secret
	c.MetricsServerEnable = ctx.GlobalBool(flags.MetricsServerEnable.Name)
	c.MetricsHostname = ctx.GlobalString(flags.MetricsHostname.Name)
	c.MetricsPort = ctx.GlobalUint64(flags.MetricsPort.Name)

	return nil
}
