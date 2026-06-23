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
	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
	"morph-l2/node/types"
)

const (
	// DefaultFetchBlockRange is the number of blocks that we collect in a
	// single eth_getLogs query. 500 (vs sync.DefaultFetchBlockRange=100)
	// trades RPC latency budget for fewer round-trips on the derivation
	// path, where each query is only a CommitBatch event filter and the
	// response stays small even at 500 blocks.
	DefaultFetchBlockRange = uint64(500)

	// DefaultPollInterval is the frequency at which we query for new L1 messages.
	DefaultPollInterval = time.Second * 15

	// DefaultLogProgressInterval is the frequency at which we log progress.
	DefaultLogProgressInterval = time.Second * 10

	VerifyModeLayer1 = "layer1"
	VerifyModeLocal  = "local"

	// DefaultVerifyMode is "local": rebuild + compare locally on the happy
	// path, no beacon blob fetch. Operators who need the legacy "always
	// pull blob" behavior can set --derivation.verify-mode=layer1.
	DefaultVerifyMode = VerifyModeLocal

	// DefaultReorgCheckDepth is the number of recent L1 blocks to check for
	// reorgs in SPEC-005 §4.7.6 detection. 64 covers the post-Merge "finality
	// distance" rule of thumb and provides safety margin if Confirmations is
	// configured below finalized.
	DefaultReorgCheckDepth = uint64(64)

	// DefaultConfirmations: rationale lives on the L1.Confirmations field
	// in DefaultConfig() — fixed-depth (latest-N) paired with the SPEC-005
	// §4.7.6 reorg detector, not a chain tag.
	DefaultConfirmations = rpc.BlockNumber(10)
)

// validateAndDefaultVerifyMode normalises an empty VerifyMode to the default
// and rejects unknown values. Extracted from SetCliContext so the validation
// can be unit-tested without building a cli.Context.
func validateAndDefaultVerifyMode(s string) (string, error) {
	switch s {
	case VerifyModeLayer1, VerifyModeLocal:
		return s, nil
	case "":
		return DefaultVerifyMode, nil
	default:
		return "", fmt.Errorf("invalid derivation.verify-mode %q (must be %q or %q)",
			s, VerifyModeLayer1, VerifyModeLocal)
	}
}

func resolveVerifyMode(current string, verifyModeSet bool, verifyMode string, legacyValidator bool) (string, error) {
	switch {
	case verifyModeSet:
		return validateAndDefaultVerifyMode(verifyMode)
	case legacyValidator:
		return VerifyModeLayer1, nil
	default:
		return validateAndDefaultVerifyMode(current)
	}
}

type Config struct {
	L1                    *types.L1Config `json:"l1"`
	L2                    *types.L2Config `json:"l2"`
	BeaconRpc             string          `json:"beacon_rpc"`
	RollupContractAddress common.Address  `json:"rollup_contract_address"`
	StartHeight           uint64          `json:"start_height"`
	BaseHeight            uint64          `json:"base_height"`
	PollInterval          time.Duration   `json:"poll_interval"`
	LogProgressInterval   time.Duration   `json:"log_progress_interval"`
	FetchBlockRange       uint64          `json:"fetch_block_range"`
	VerifyMode            string          `json:"verify_mode"`
	ReorgCheckDepth       uint64          `json:"reorg_check_depth"`
}

func DefaultConfig() *Config {
	return &Config{
		L1: &types.L1Config{
			// Fixed-depth (latest-N) confirmations rather than the SafeBlockNumber
			// tag: 10 blocks (~2 min on mainnet) keeps lag low, and the always-on
			// L1 reorg detector (SPEC-005 §4.7.6 in reorg.go) rewinds the
			// derivation cursor on hash mismatch so a deeper reorg is recoverable.
			// Operators wanting strict no-reorg-possible reads can still set
			// --derivation.confirmations=-3 (rpc.FinalizedBlockNumber).
			Confirmations: DefaultConfirmations,
		},
		PollInterval:        DefaultPollInterval,
		LogProgressInterval: DefaultLogProgressInterval,
		FetchBlockRange:     DefaultFetchBlockRange,
		VerifyMode:          DefaultVerifyMode,
		ReorgCheckDepth:     DefaultReorgCheckDepth,
		L2:                  new(types.L2Config),
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	c.L1.Addr = ctx.GlobalString(flags.L1NodeAddr.Name)
	// The current setting priority is greater than Env L1Confirmations
	if ctx.GlobalIsSet(flags.DerivationConfirmations.Name) {
		c.L1.Confirmations = rpc.BlockNumber(ctx.GlobalInt64(flags.DerivationConfirmations.Name))
		log.Warn("derivation confirmations reset to ", c.L1.Confirmations)
	}
	if ctx.GlobalIsSet(flags.RollupContractAddress.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.RollupContractAddress.Name))
		c.RollupContractAddress = addr
		if len(c.RollupContractAddress.Bytes()) == 0 {
			return errors.New("invalid DerivationDepositContractAddr")
		}
	} else if ctx.GlobalBool(flags.MainnetFlag.Name) {
		c.RollupContractAddress = types.MainnetRollupContractAddress
	} else if ctx.GlobalBool(flags.HoodiFlag.Name) {
		c.RollupContractAddress = types.HoodiRollupContractAddress
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

	if ctx.GlobalIsSet(flags.DerivationBaseHeight.Name) {
		c.BaseHeight = ctx.GlobalUint64(flags.DerivationBaseHeight.Name)
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

	normalized, err := resolveVerifyMode(
		c.VerifyMode,
		ctx.GlobalIsSet(flags.DerivationVerifyMode.Name),
		ctx.GlobalString(flags.DerivationVerifyMode.Name),
		ctx.GlobalBool(flags.LegacyValidatorMode.Name),
	)
	if err != nil {
		return err
	}
	c.VerifyMode = normalized

	if ctx.GlobalIsSet(flags.DerivationReorgCheckDepth.Name) {
		c.ReorgCheckDepth = ctx.GlobalUint64(flags.DerivationReorgCheckDepth.Name)
	}
	if c.ReorgCheckDepth == 0 {
		c.ReorgCheckDepth = DefaultReorgCheckDepth
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

	return nil
}
