// Package devtool holds TEST / TESTNET-ONLY developer helpers for devnet and HA
// testnet bring-up. Nothing in this package may be enabled on a production network.
//
// DO NOT ENABLE IN PRODUCTION.
package devtool

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/tendermint/tendermint/config"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/privval"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/upgrade"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
	"morph-l2/node/types"
)

// startInSequencerModeFlag, when set, makes the node start directly in sequencer
// mode (skipping the pre-upgrade PBFT phase) by pre-setting the consensus upgrade
// block height to 0. Default is false, so production and ordinary devnet runs are
// unaffected.
//
// TEST/TESTNET-ONLY: never enable on a production network.
var startInSequencerModeFlag = cli.BoolFlag{
	Name:   "startInSequencerMode",
	Usage:  "[TEST-ONLY] start directly in sequencer mode (pre-set upgrade block height = 0, skip the PBFT phase); never enable in production",
	EnvVar: "MORPH_NODE_START_IN_SEQUENCER_MODE",
}

// init self-registers the flag into the shared flag list so that cmd/node picks it
// up via `app.Flags = flags.Flags` without editing node/flags/flags.go. It runs
// because cmd/node imports this package to call ApplyStartInSequencerMode.
func init() {
	flags.Flags = append(flags.Flags, startInSequencerModeFlag)
}

// ApplyStartInSequencerMode optionally pre-sets the consensus upgrade block height
// to 0 so that IsUpgraded(1) == true and the node starts the sequencer routines
// directly, never entering the PBFT consensus reactor. Note the tendermint node
// itself still starts; only the pre-upgrade PBFT consensus phase is skipped.
//
// Ordering requirement: this MUST run before the upgrade store is wired (i.e.
// before SetupNode / node startup). At that point upgrade's store is still nil,
// so SetUpgradeBlockHeight only sets the in-memory value and does NOT persist it.
// On a fresh DB the later SetStore finds the key absent and keeps this value.
// Because this mode never enters PBFT, the height is never persisted, so every
// restart simply re-applies it here — making the switch idempotent across restarts
// (verified against upgrade.SetStore / SetUpgradeBlockHeight).
//
// Two guards refuse to start rather than let a misconfiguration through; see
// errProductionNetwork and errSoleGenesisValidator below.
//
// TEST/TESTNET-ONLY. No-op unless --startInSequencerMode /
// MORPH_NODE_START_IN_SEQUENCER_MODE is set.
func ApplyStartInSequencerMode(ctx *cli.Context, logger tmlog.Logger) error {
	if !ctx.GlobalBool(startInSequencerModeFlag.Name) {
		return nil
	}
	if network, production := productionNetwork(ctx); production {
		return fmt.Errorf("%w: --%s (env %s) is set on the %s network",
			errProductionNetwork, startInSequencerModeFlag.Name,
			startInSequencerModeFlag.EnvVar, network)
	}
	if err := checkNotSoleGenesisValidator(ctx, logger); err != nil {
		return err
	}
	upgrade.SetUpgradeBlockHeight(0)
	logger.Info("[TEST-ONLY] start-in-sequencer-mode: pre-set upgrade block height, starting directly in sequencer mode (PBFT phase skipped)",
		"upgradeBlockHeight", upgrade.UpgradeBlockHeight())
	return nil
}

// errProductionNetwork guards the switch against the deployment accident it is
// most exposed to: MORPH_NODE_START_IN_SEQUENCER_MODE surviving a copy-paste into
// a mainnet chart. A mainnet node that pre-sets the upgrade height treats itself
// as post-upgrade from block 1, skips PBFT entirely and forks off the network, so
// refuse to boot rather than trust the comments in this file.
var errProductionNetwork = errors.New("start-in-sequencer-mode is TEST-ONLY and must never be enabled on a production network")

// errSoleGenesisValidator guards the one topology that silently cannot work.
//
// Pre-setting the upgrade height closes both entry points into sequencer mode for
// a node that holds the only genesis validator key:
//
//   - tendermint node.go:903 computes blockSync as
//     `config.BlockSyncMode && !onlyValidatorIsUs(state, pubKey)`, so this node
//     gets blockSync=false, blocksync/reactor.go:187 does nothing, poolRoutine
//     never runs, and the caught-up hand-over at blocksync/reactor.go:608 that
//     would call StartSequencerRoutines() never happens;
//   - IsUpgraded(1) is true, so consensus/reactor.go:81-85 returns early, PBFT
//     never runs, and the upgrade callback at node.go:1700 never fires either.
//
// StateV2 is therefore never started and the node produces nothing — while
// looking completely healthy: containers up, RPC serving, no error logged, block
// height pinned at 0. That is expensive to diagnose, so fail fast instead.
//
// The fix is the one the HA cluster already relies on: do not hold the genesis
// validator key. setup_nodes.py copies priv_validator_key.json for node0 only and
// deletes it elsewhere, so every other node boots with a key tendermint generated
// for it (privval.LoadOrGenFilePV) and onlyValidatorIsUs is false. This is safe
// because PBFT never runs here, so the genesis validator set is never used to
// vote; block production is gated on the L1 sequencer contract plus the sequencer
// signing key, not on tendermint validator identity.
//
// Note this does not cover the second requirement, which cannot be checked at
// startup because peers connect asynchronously: BlockPool.IsCaughtUp() returns
// false while the pool has no peers (blocksync/pool.go:186), so at least one
// other node running tendermint is still needed. A node with
// MORPH_NODE_DERIVATION_VERIFY_MODE=layer1 does not start tendermint and does not
// count as a peer.
var errSoleGenesisValidator = errors.New("start-in-sequencer-mode cannot work on a node holding the only genesis validator key: tendermint disables block sync for it, so the blocksync hand-over never starts the sequencer routines, and the pre-set upgrade height keeps PBFT from starting either. Use a priv_validator_key.json that is not in the genesis validator set (see ops/docker/docker-compose-devnet-skip-tendermint.yml), or drop --startInSequencerMode")

// productionNetwork reports the selected network and whether it is one where this
// switch must never be honoured. It mirrors sequencerUpgradeNetwork in
// cmd/node/main.go, which resolves the same two flags one call earlier.
func productionNetwork(ctx *cli.Context) (string, bool) {
	if ctx.GlobalBool(flags.MainnetFlag.Name) {
		return "mainnet", true
	}
	if ctx.GlobalBool(flags.HoodiFlag.Name) {
		return "hoodi", true
	}
	return "dev", false
}

// checkNotSoleGenesisValidator returns errSoleGenesisValidator when this node's
// private validator key is the single validator in the genesis document.
//
// It deliberately fails open. Anything that stops it from reaching a definite
// answer — no key file (the expected, working case), no genesis file, unreadable
// or malformed either — is reported as "cannot tell" and lets startup continue,
// because a guard for a test-only switch must not become a new way to refuse to
// boot. Only a positive match blocks. For the same reason the files are read and
// unmarshalled here rather than via privval.LoadFilePV, which calls os.Exit on a
// malformed key.
func checkNotSoleGenesisValidator(ctx *cli.Context, logger tmlog.Logger) error {
	home, err := homeDir(ctx)
	if err != nil {
		logger.Debug("[TEST-ONLY] start-in-sequencer-mode: cannot resolve home dir, skipping sole-validator check", "err", err)
		return nil
	}
	// Paths come from tendermint's defaults rooted at home. A config.toml that
	// relocates either file makes this read miss and fail open, per the contract
	// above.
	tmCfg := cfg.DefaultConfig().SetRoot(home)

	keyPath := tmCfg.PrivValidatorKeyFile()
	keyJSON, err := os.ReadFile(keyPath)
	if err != nil {
		// The common case: setup deleted the key so tendermint will generate a
		// non-genesis one. Exactly the configuration this switch needs.
		logger.Debug("[TEST-ONLY] start-in-sequencer-mode: no private validator key on disk, sole-validator check not applicable", "path", keyPath)
		return nil
	}
	var pvKey privval.FilePVKey
	if err := tmjson.Unmarshal(keyJSON, &pvKey); err != nil {
		logger.Debug("[TEST-ONLY] start-in-sequencer-mode: cannot parse private validator key, skipping sole-validator check", "path", keyPath, "err", err)
		return nil
	}
	if pvKey.PrivKey == nil {
		return nil
	}

	genDoc, err := tmtypes.GenesisDocFromFile(tmCfg.GenesisFile())
	if err != nil {
		logger.Debug("[TEST-ONLY] start-in-sequencer-mode: cannot read genesis, skipping sole-validator check", "path", tmCfg.GenesisFile(), "err", err)
		return nil
	}
	if len(genDoc.Validators) != 1 {
		return nil
	}
	// Derive the address from the public key rather than trusting the file's
	// "address" field, which is what privval does when it loads a key.
	if !bytes.Equal(genDoc.Validators[0].Address, pvKey.PrivKey.PubKey().Address()) {
		return nil
	}
	return fmt.Errorf("%w (key %s, genesis %s)", errSoleGenesisValidator, keyPath, tmCfg.GenesisFile())
}

// homeDir mirrors homeDir in cmd/node/main.go. It is duplicated rather than
// exported from there because this package must not be imported by main's own
// helpers, and the switch has to resolve the config dir before SetupNode runs.
func homeDir(ctx *cli.Context) (string, error) {
	home := ctx.GlobalString(flags.Home.Name)
	if home == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		home = filepath.Join(userHome, types.DefaultHomeDir)
	}
	return home, nil
}
