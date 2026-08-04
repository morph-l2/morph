// Package devtool holds TEST / TESTNET-ONLY developer helpers for devnet and HA
// testnet bring-up. Nothing in this package may be enabled on a production network.
//
// DO NOT ENABLE IN PRODUCTION.
package devtool

import (
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
// TEST/TESTNET-ONLY. No-op unless --startInSequencerMode /
// MORPH_NODE_START_IN_SEQUENCER_MODE is set.
func ApplyStartInSequencerMode(ctx *cli.Context, logger tmlog.Logger) {
	if !ctx.GlobalBool(startInSequencerModeFlag.Name) {
		return
	}
	upgrade.SetUpgradeBlockHeight(0)
	logger.Info("[TEST-ONLY] start-in-sequencer-mode: pre-set upgrade block height, starting directly in sequencer mode (PBFT phase skipped)",
		"upgradeBlockHeight", upgrade.UpgradeBlockHeight())
}
