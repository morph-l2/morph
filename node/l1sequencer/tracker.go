package l1sequencer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/morph-l2/go-ethereum/ethclient"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// checkInterval is how often the tracker polls L1. Smaller than the halt
// threshold so the gate trips/recovers promptly without burdening the RPC.
const checkInterval = 15 * time.Second

// defaultHaltLag is the threshold at which the node halts: once our reference
// timestamp (the L1 head time on success, or the first-failure time during an
// RPC outage) is more than this behind wall-clock, the sequencer stops producing
// and fullnodes stop syncing. Hardcoded for now; promote to a flag if needed.
const defaultHaltLag = 30 * time.Minute

// verifierRefresher re-syncs the L1 sequencer set on demand. The tracker forces
// one refresh when L1 recovers, before reopening the gate, so block acceptance
// resumes against the freshest sequencer set. Implemented by *SequencerVerifier.
type verifierRefresher interface {
	refresh() error
}

// L1Tracker polls the L1 RPC and reports a single halt signal via IsHalt: when
// our most recent reference timestamp is older than haltLag, the sequencer must
// stop producing and fullnodes must stop syncing, to avoid acting on a stale
// view of L1 sequencer changes. It implements sequencer.L1Tracker.
type L1Tracker struct {
	ctx          context.Context
	cancel       context.CancelFunc
	l1Client     *ethclient.Client
	verifier     verifierRefresher
	lagThreshold time.Duration // warn threshold (log only)
	haltLag      time.Duration // halt threshold
	logger       tmlog.Logger
	stop         chan struct{}

	// State below is only mutated from the single loop goroutine (and tests).
	healthy  atomic.Bool // read concurrently by gate consumers
	lastSeen time.Time   // L1 head time on success, or first-failure time during an outage
	inErr    bool        // in an RPC-failure run; keeps lastSeen anchored at the first failure
}

// NewL1Tracker creates a new L1Tracker. verifier must not be nil.
func NewL1Tracker(
	ctx context.Context,
	l1Client *ethclient.Client,
	verifier verifierRefresher,
	warnLag time.Duration,
	logger tmlog.Logger,
) *L1Tracker {
	ctx, cancel := context.WithCancel(ctx)
	t := &L1Tracker{
		ctx:          ctx,
		cancel:       cancel,
		l1Client:     l1Client,
		verifier:     verifier,
		lagThreshold: warnLag,
		haltLag:      defaultHaltLag,
		logger:       logger.With("module", "l1tracker"),
		stop:         make(chan struct{}),
	}
	t.healthy.Store(true)   // start allowed
	t.lastSeen = time.Now() // grace: tolerate initial RPC failures for haltLag
	return t
}

// IsHalt implements sequencer.L1Tracker.
func (t *L1Tracker) IsHalt() bool { return !t.healthy.Load() }

func (t *L1Tracker) Start() error {
	t.logger.Info("Starting L1Tracker", "warnLag", t.lagThreshold, "haltLag", t.haltLag, "tick", checkInterval)
	go t.loop()
	return nil
}

func (t *L1Tracker) Stop() {
	t.logger.Info("Stopping L1Tracker")
	t.cancel()
	<-t.stop
}

func (t *L1Tracker) loop() {
	defer close(t.stop)
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-ticker.C:
			t.check()
		}
	}
}

// check polls the L1 head, emits the warn-level log, and folds the result into
// the health state.
func (t *L1Tracker) check() {
	header, err := t.l1Client.HeaderByNumber(t.ctx, nil)
	if err != nil {
		t.logger.Error("Failed to get L1 header", "error", err)
		t.update(time.Time{}, false, time.Now())
		return
	}
	headTime := time.Unix(int64(header.Time), 0)
	if lag := time.Since(headTime); lag > t.lagThreshold {
		t.logger.Error("L1 RPC is behind",
			"latestBlock", header.Number.Uint64(),
			"blockTime", headTime.Format(time.RFC3339),
			"lag", lag.Round(time.Second),
		)
	}
	t.update(headTime, true, time.Now())
}

// update folds one poll into the health state. On success it anchors lastSeen at
// the L1 head time; on the first failure of an outage it anchors at now (the
// inErr flag stops later failures from re-anchoring). It then halts when lastSeen
// is older than haltLag, and on recovery refreshes the verifier before reopening.
// now is injected for testability.
func (t *L1Tracker) update(headTime time.Time, ok bool, now time.Time) {
	if ok {
		t.inErr = false
		t.lastSeen = headTime
	} else if !t.inErr {
		t.inErr = true
		t.lastSeen = now
	}

	if now.Sub(t.lastSeen) > t.haltLag {
		if t.healthy.CompareAndSwap(true, false) {
			t.logger.Error("L1 health gate TRIPPED: L1 too stale, halting block production and sync", "haltLag", t.haltLag)
		}
		return
	}

	// L1 is fresh again. On recovery, force a verifier resync before reopening;
	// if it fails, stay halted and retry next tick.
	if !t.healthy.Load() {
		if err := t.verifier.refresh(); err != nil {
			t.logger.Error("verifier refresh on L1 recovery failed; staying halted", "err", err)
			return
		}
		t.healthy.Store(true)
		t.logger.Info("L1 health gate RECOVERED: resuming block production and sync")
	}
}
