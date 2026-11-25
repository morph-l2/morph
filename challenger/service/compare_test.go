package service

import (
	"testing"

	"github.com/morph-l2/go-ethereum/common"
)

func TestDecide(t *testing.T) {
	state := common.HexToHash("0x11")
	withdraw := common.HexToHash("0x22")
	other := common.HexToHash("0x33")
	cases := []struct {
		name    string
		obs     ObservedRoots
		verdict Verdict
		reason  string
	}{
		{
			name: "match",
			obs: ObservedRoots{
				BlockNumber: 9, L1State: state, L2State: state, L2HeaderFound: true,
				L1Withdraw: withdraw, L2Withdraw: withdraw,
			},
			verdict: VerdictMatch,
		},
		{
			name: "state mismatch",
			obs: ObservedRoots{
				BlockNumber: 9, L1State: state, L2State: other, L2HeaderFound: true,
				L1Withdraw: withdraw, L2Withdraw: withdraw,
			},
			verdict: VerdictMismatch,
			reason:  "state root differs",
		},
		{
			name: "withdraw mismatch",
			obs: ObservedRoots{
				BlockNumber: 9, L1State: state, L2State: state, L2HeaderFound: true,
				L1Withdraw: withdraw, L2Withdraw: other,
			},
			verdict: VerdictMismatch,
			reason:  "withdraw root differs",
		},
		{
			name: "missing block",
			obs: ObservedRoots{
				BlockNumber: 9, L1State: state,
			},
			verdict: VerdictInconclusive,
			reason:  "l2 block is not available",
		},
		{
			name: "message root call failed",
			obs: ObservedRoots{
				BlockNumber: 9, L1State: state, L2State: state, L2HeaderFound: true,
				L2WithdrawMissing: true, L2WithdrawErr: "execution reverted",
			},
			verdict: VerdictInconclusive,
			reason:  "messageRoot: execution reverted",
		},
		{
			name:    "empty l1 state root",
			obs:     ObservedRoots{BlockNumber: 9, L2HeaderFound: true, L2State: state},
			verdict: VerdictInconclusive,
			reason:  "l1 state root is empty",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Decide(4, tc.obs)
			if got.Verdict != tc.verdict || got.Reason != tc.reason || got.BlockNumber != tc.obs.BlockNumber {
				t.Fatalf("got %+v", got)
			}
		})
	}
}
