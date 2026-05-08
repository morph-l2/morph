package derivation

import "morph-l2/node/db"

// toDBAnchor converts the in-memory HeadAnchor to the persistent representation.
func (a HeadAnchor) toDBAnchor() *db.DerivationHeadAnchor {
	return &db.DerivationHeadAnchor{
		L2Number: a.L2Number,
		L2Hash:   a.L2Hash,
		L1Number: a.L1Number,
		L1Hash:   a.L1Hash,
	}
}

// headAnchorFromDB inflates a stored anchor back into the in-memory representation.
// Returns nil if the input is nil.
func headAnchorFromDB(a *db.DerivationHeadAnchor) *HeadAnchor {
	if a == nil {
		return nil
	}
	return &HeadAnchor{
		L2Number: a.L2Number,
		L2Hash:   a.L2Hash,
		L1Number: a.L1Number,
		L1Hash:   a.L1Hash,
	}
}

// readSafeHead returns the persisted safe-stage anchor, or nil if unset.
func (d *Derivation) readSafeHead() *HeadAnchor {
	return headAnchorFromDB(d.db.ReadDerivationSafeHead())
}

// readFinalizedHead returns the persisted finalized-stage anchor, or nil if unset.
func (d *Derivation) readFinalizedHead() *HeadAnchor {
	return headAnchorFromDB(d.db.ReadDerivationFinalizedHead())
}

// writeSafeHead persists a new safe-stage anchor.
//
// Per SPEC-005 §3.5 ("Restart and consistency"), this should ideally be written
// atomically with the corresponding L1 anchor window updates so the node never
// observes a half-committed state across a restart. The current implementation
// uses single-key Put because the underlying KV store does not yet expose a
// transactional API; this is acceptable for now because:
//  - safe head ratchets forward inside a single derivation loop iteration;
//  - L1 anchor window writes are append-only and idempotent;
//  - on crash mid-write, the next loop will re-derive from the last persisted
//    LatestDerivationL1Height and re-establish consistency before advancing.
//
// TODO(spec-005): expose a multi-key atomic write helper on db.Store and
// migrate this + WriteDerivationL1Block + WriteLatestDerivationL1Height onto
// it once the rollback executor (P3) lands.
func (d *Derivation) writeSafeHead(anchor HeadAnchor) {
	d.db.WriteDerivationSafeHead(anchor.toDBAnchor())
}

// writeFinalizedHead persists a new finalized-stage anchor.
//
// Per SPEC-005 §3.1, finalized_head is monotonic and never rolls back; callers
// must enforce this invariant before calling.
func (d *Derivation) writeFinalizedHead(anchor HeadAnchor) {
	d.db.WriteDerivationFinalizedHead(anchor.toDBAnchor())
}
