package l1sequencer

import (
	"errors"
	"testing"
	"time"

	tmlog "github.com/tendermint/tendermint/libs/log"
)

type fakeRefresher struct {
	calls int
	err   error
}

func (f *fakeRefresher) refresh() error { f.calls++; return f.err }

func newTrackerForTest() (*L1Tracker, *fakeRefresher) {
	f := &fakeRefresher{}
	t := &L1Tracker{
		verifier:     f,
		lagThreshold: 5 * time.Minute,
		haltLag:      30 * time.Minute,
		logger:       tmlog.NewNopLogger(),
		metrics:      NopMetrics(),
	}
	t.healthy.Store(true)
	return t, f
}

func TestUpdate_LagTripsImmediately(t *testing.T) {
	tr, _ := newTrackerForTest()
	now := time.Unix(1_000_000, 0)
	tr.lastSeen = now
	tr.update(now.Add(-31*time.Minute), true, now)
	if !tr.IsHalt() {
		t.Fatal("an L1 head 31m old should halt immediately")
	}
}

func TestUpdate_LagBoundaryDoesNotTrip(t *testing.T) {
	tr, _ := newTrackerForTest()
	now := time.Unix(1_000_000, 0)
	tr.lastSeen = now
	tr.update(now.Add(-30*time.Minute), true, now) // exactly haltLag
	if tr.IsHalt() {
		t.Fatal("a head exactly haltLag old should NOT halt")
	}
}

func TestUpdate_FirstErrAnchorsAndTripsAfterHaltLag(t *testing.T) {
	tr, _ := newTrackerForTest()
	base := time.Unix(1_000_000, 0)
	tr.lastSeen = base
	tr.update(time.Time{}, false, base) // first failure anchors at base
	if tr.IsHalt() {
		t.Fatal("first failure must not trip")
	}
	tr.update(time.Time{}, false, base.Add(29*time.Minute))
	if tr.IsHalt() {
		t.Fatal("29m of failures must not trip")
	}
	tr.update(time.Time{}, false, base.Add(31*time.Minute))
	if !tr.IsHalt() {
		t.Fatal("31m since the first failure must halt")
	}
}

func TestUpdate_ErrFlagKeepsFirstFailureAnchor(t *testing.T) {
	// Without the inErr flag, every failure would re-anchor lastSeen to `now`
	// and the gate would never trip. Confirm the anchor stays at the FIRST one.
	tr, _ := newTrackerForTest()
	base := time.Unix(1_000_000, 0)
	tr.lastSeen = base
	for i := 0; i <= 31; i++ {
		tr.update(time.Time{}, false, base.Add(time.Duration(i)*time.Minute))
	}
	if !tr.IsHalt() {
		t.Fatal("repeated failures must still trip 30m after the first")
	}
}

func TestUpdate_SuccessClearsErrAndReanchors(t *testing.T) {
	tr, _ := newTrackerForTest()
	base := time.Unix(1_000_000, 0)
	tr.lastSeen = base
	tr.update(time.Time{}, false, base)                                 // err run anchored at base
	tr.update(base.Add(20*time.Minute), true, base.Add(20*time.Minute)) // success clears inErr, re-anchors
	tr.update(time.Time{}, false, base.Add(40*time.Minute))             // new err run anchors at base+40m
	tr.update(time.Time{}, false, base.Add(65*time.Minute))             // only 25m since new anchor
	if tr.IsHalt() {
		t.Fatal("a success should clear the err run and re-anchor the staleness clock")
	}
}

func TestUpdate_RecoversWithRefresh(t *testing.T) {
	tr, f := newTrackerForTest()
	now := time.Unix(1_000_000, 0)
	tr.lastSeen = now
	tr.update(now.Add(-31*time.Minute), true, now) // trip
	if !tr.IsHalt() {
		t.Fatal("should be halted")
	}
	n1 := now.Add(1 * time.Minute)
	tr.update(n1, true, n1) // fresh read -> refresh once, then recover
	if tr.IsHalt() {
		t.Fatal("should recover on the first fresh read")
	}
	if f.calls != 1 {
		t.Fatalf("verifier.refresh should be called once on recovery, got %d", f.calls)
	}
	n2 := now.Add(2 * time.Minute)
	tr.update(n2, true, n2) // already healthy -> no extra refresh
	if f.calls != 1 {
		t.Fatalf("no extra refresh while healthy, got %d", f.calls)
	}
}

func TestUpdate_RecoveryBlockedWhenRefreshFails(t *testing.T) {
	tr, f := newTrackerForTest()
	f.err = errors.New("rpc still flaky")
	now := time.Unix(1_000_000, 0)
	tr.lastSeen = now
	tr.update(now.Add(-31*time.Minute), true, now) // trip
	n1 := now.Add(1 * time.Minute)
	tr.update(n1, true, n1) // fresh but refresh fails -> stay halted
	if !tr.IsHalt() {
		t.Fatal("must stay halted while the recovery refresh fails")
	}
	f.err = nil
	n2 := now.Add(2 * time.Minute)
	tr.update(n2, true, n2) // refresh succeeds -> recover
	if tr.IsHalt() {
		t.Fatal("should recover once the refresh succeeds")
	}
}
