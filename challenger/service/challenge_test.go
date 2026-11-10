package service

import (
	"context"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
)

type staticCompare struct {
	cmp Comparison
	err error
}

func (s *staticCompare) Compare(context.Context, uint64) (Comparison, error) {
	return s.cmp, s.err
}

func rejectReason(t *testing.T, err error) string {
	t.Helper()
	rejected, ok := err.(*RejectError)
	if !ok {
		t.Fatalf("err = %v, want RejectError", err)
	}
	return rejected.Reason
}

func TestChallengeRefusesUnlessRootsMismatch(t *testing.T) {
	hash := common.HexToHash("0x11")
	match := &Service{compare: &staticCompare{cmp: Comparison{BatchHash: hash, Verdict: VerdictMatch}}}
	_, err := match.Challenge(context.Background(), 4, hash)
	if reason := rejectReason(t, err); reason != "batch roots match" {
		t.Fatalf("match reason = %s", reason)
	}

	notChallenger := readyView(hash)
	notChallenger.challenger = false
	mismatch := &Service{
		compare: &staticCompare{cmp: Comparison{BatchHash: hash, Verdict: VerdictMismatch}},
		view:    notChallenger,
		from:    common.HexToAddress("0x1"),
	}
	_, err = mismatch.Challenge(context.Background(), 4, hash)
	if reason := rejectReason(t, err); reason != "only challenger allowed" {
		t.Fatalf("mismatch reason = %s", reason)
	}
}

func TestChallengeRefusesAnotherVersion(t *testing.T) {
	current := common.HexToHash("0x22")
	requested := common.HexToHash("0x11")
	svc := &Service{compare: &staticCompare{cmp: Comparison{BatchHash: current, Verdict: VerdictMismatch}}}
	_, err := svc.Challenge(context.Background(), 4, requested)
	if reason := rejectReason(t, err); reason != "incorrect batch hash" {
		t.Fatalf("reason = %s", reason)
	}
}

func TestChallengeRefusesRevertedBatch(t *testing.T) {
	svc := &Service{compare: &staticCompare{err: &NotCommittedError{BatchIndex: 4}}}
	_, err := svc.Challenge(context.Background(), 4, common.HexToHash("0x11"))
	if reason := rejectReason(t, err); reason != "batch is not committed" {
		t.Fatalf("reason = %s", reason)
	}
}

func TestChallengeRefusesFinalizedBatch(t *testing.T) {
	svc := &Service{compare: &staticCompare{err: &RejectError{Reason: "batch already finalized"}}}
	_, err := svc.Challenge(context.Background(), 4, common.HexToHash("0x11"))
	if reason := rejectReason(t, err); reason != "batch already finalized" {
		t.Fatalf("reason = %s", reason)
	}
}
