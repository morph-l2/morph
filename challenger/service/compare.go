package service

import "github.com/morph-l2/go-ethereum/common"

// Verdict is the result of comparing one batch's roots.
type Verdict string

const (
	// VerdictMatch means both the state root and the withdraw root are equal.
	VerdictMatch Verdict = "match"
	// VerdictMismatch means at least one of the two roots differs.
	VerdictMismatch Verdict = "mismatch"
	// VerdictInconclusive means the trusted geth could not supply a root.
	VerdictInconclusive Verdict = "inconclusive"
)

// RootPair is one root read from L1 and from the trusted L2 geth.
type RootPair struct {
	L1 common.Hash `json:"l1"`
	L2 common.Hash `json:"l2"`
}

// Comparison is one batch version checked at its last L2 block.
// BatchHash is the committedBatches value the roots were read for.
type Comparison struct {
	BatchIndex   uint64      `json:"batchIndex"`
	BatchHash    common.Hash `json:"batchHash"`
	BlockNumber  uint64      `json:"blockNumber"`
	StateRoot    RootPair    `json:"stateRoot"`
	WithdrawRoot RootPair    `json:"withdrawRoot"`
	Verdict      Verdict     `json:"verdict"`
	Reason       string      `json:"reason,omitempty"`
}

// ObservedRoots is the chain data used to decide a verdict.
type ObservedRoots struct {
	BlockNumber       uint64
	L1State           common.Hash
	L1Withdraw        common.Hash
	L2HeaderFound     bool
	L2State           common.Hash
	L2Withdraw        common.Hash
	L2WithdrawMissing bool
	L2WithdrawErr     string
}

// Decide compares the two root pairs. A zero L1 state root is not a mismatch:
// it means the L1 side could not be read, not that the batch is wrong.
func Decide(index uint64, obs ObservedRoots) Comparison {
	out := Comparison{
		BatchIndex:  index,
		BlockNumber: obs.BlockNumber,
		StateRoot:   RootPair{L1: obs.L1State, L2: obs.L2State},
		WithdrawRoot: RootPair{
			L1: obs.L1Withdraw,
			L2: obs.L2Withdraw,
		},
	}
	if obs.L1State == (common.Hash{}) {
		out.Verdict = VerdictInconclusive
		out.Reason = "l1 state root is empty"
		return out
	}
	if !obs.L2HeaderFound {
		out.Verdict = VerdictInconclusive
		out.Reason = "l2 block is not available"
		return out
	}
	if obs.L2WithdrawMissing {
		out.Verdict = VerdictInconclusive
		out.Reason = "messageRoot: " + obs.L2WithdrawErr
		return out
	}
	stateOK := obs.L1State == obs.L2State
	withdrawOK := obs.L1Withdraw == obs.L2Withdraw
	if stateOK && withdrawOK {
		out.Verdict = VerdictMatch
		return out
	}
	out.Verdict = VerdictMismatch
	switch {
	case !stateOK && !withdrawOK:
		out.Reason = "state root and withdraw root differ"
	case !stateOK:
		out.Reason = "state root differs"
	default:
		out.Reason = "withdraw root differs"
	}
	return out
}

func refuseReason(cmp Comparison) string {
	if cmp.Verdict == VerdictMatch {
		return "batch roots match"
	}
	if cmp.Reason == "" {
		return "batch comparison inconclusive"
	}
	return "batch comparison inconclusive: " + cmp.Reason
}
