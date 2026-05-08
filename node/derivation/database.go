package derivation

import (
	"morph-l2/node/db"
	"morph-l2/node/sync"
)

type Database interface {
	Reader
	Writer
	sync.Database
}

type Reader interface {
	ReadLatestDerivationL1Height() *uint64
	ReadDerivationL1Block(l1Height uint64) *db.DerivationL1Block
	ReadDerivationL1BlockRange(from, to uint64) []*db.DerivationL1Block
	// SPEC-005: safe / finalized head anchors.
	ReadDerivationSafeHead() *db.DerivationHeadAnchor
	ReadDerivationFinalizedHead() *db.DerivationHeadAnchor
}

type Writer interface {
	WriteLatestDerivationL1Height(latest uint64)
	WriteDerivationL1Block(block *db.DerivationL1Block)
	DeleteDerivationL1BlocksFrom(height uint64)
	// SPEC-005: safe / finalized head anchors.
	WriteDerivationSafeHead(anchor *db.DerivationHeadAnchor)
	WriteDerivationFinalizedHead(anchor *db.DerivationHeadAnchor)
}
