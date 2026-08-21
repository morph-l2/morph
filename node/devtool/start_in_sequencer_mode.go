// Package devtool holds TEST / TESTNET-ONLY developer helpers for devnet and HA
// testnet bring-up. Nothing in this package may be enabled on a production network.
//
// DO NOT ENABLE IN PRODUCTION.
package devtool

import (
	"errors"
	"fmt"

	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/upgrade"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
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

// errProductionNetwork guards the switch against the deployment accident it is
// most exposed to: MORPH_NODE_START_IN_SEQUENCER_MODE surviving a copy-paste into
// a mainnet chart. A node that pre-sets the upgrade height treats itself as
// post-upgrade from block 1, skips PBFT entirely and forks off the network, so
// refuse to boot rather than trust the comments in this file.
var errProductionNetwork = errors.New("start-in-sequencer-mode is TEST-ONLY and must never be enabled on a production network")

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
// Enabling this is not sufficient on its own; see ops/README.md for the
// deployment prerequisites, none of which this function can check.
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
	upgrade.SetUpgradeBlockHeight(0)
	logger.Info("[TEST-ONLY] start-in-sequencer-mode: pre-set upgrade block height, starting directly in sequencer mode (PBFT phase skipped)",
		"upgradeBlockHeight", upgrade.UpgradeBlockHeight())
	return nil
}

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
