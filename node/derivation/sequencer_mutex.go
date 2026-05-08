package derivation

import "sync"

// SPEC-005 §3.6 / §4 sequencer ↔ derivation mutual exclusion.
//
// Per SPEC-005 §3.7 non-target "do not modify sequencer block production",
// the mutex is enforced **on the morph/node side of the L2Node interface
// (RequestBlockData / DeliverBlock)**. The tendermint consensus layer is
// *not* modified.
//
// This file provides the mutex primitive. Wiring on the sequencer entry
// points (morph/node/sequencer/...) is a separate task tracked in
// tech-design §6.2 task #11.
//
// Granularity (global stop-the-world vs interval lock) is a SPEC-005 §8 #5
// open question. The default scaffold below is a single global RWMutex,
// which gives global exclusion; if interval locking is later chosen, the
// public API stays the same but the internal representation grows a per-
// range structure. Callers should therefore depend only on the methods,
// not on this being a single global lock.

// SequencerMutex coordinates between block production and derivation
// rollback. Any path producing a new unsafe L2 block must acquire a
// production lock; the rollback executor takes an exclusive lock during
// the SetHead → metadata persistence sequence.
type SequencerMutex struct {
	mu sync.RWMutex
}

// NewSequencerMutex returns a fresh mutex. There is one such mutex per
// node process; sharing is established through the constructor wiring.
func NewSequencerMutex() *SequencerMutex {
	return &SequencerMutex{}
}

// AcquireProduction blocks until the rollback executor (if any) has
// released the exclusive lock, then reserves a slot for block production.
// Each call must be paired with a deferred ReleaseProduction.
//
// TODO(spec-005-mutex): once SPEC-005 §8 #5 picks interval locking, this
// signature gains a (from, to) range and the implementation switches to
// a per-range exclusion table.
func (m *SequencerMutex) AcquireProduction() {
	m.mu.RLock()
}

// ReleaseProduction releases a production reservation acquired via
// AcquireProduction. Safe to call from defer.
func (m *SequencerMutex) ReleaseProduction() {
	m.mu.RUnlock()
}

// AcquireRollback blocks until all in-flight production reservations have
// been released, then reserves the exclusive rollback slot. Each call must
// be paired with a deferred ReleaseRollback.
func (m *SequencerMutex) AcquireRollback() {
	m.mu.Lock()
}

// ReleaseRollback releases the exclusive rollback slot. Safe to call from defer.
func (m *SequencerMutex) ReleaseRollback() {
	m.mu.Unlock()
}
