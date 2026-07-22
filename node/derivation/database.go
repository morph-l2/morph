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
	// ReadDerivationL1BlockRange returns saved L1 block records in [from, to]
	// inclusive. Used by SPEC-005 §4.7.6 reorg detection.
	ReadDerivationL1BlockRange(from, to uint64) []*db.DerivationL1Block
}

type Writer interface {
	WriteLatestDerivationL1Height(latest uint64)
	// WriteDerivationL1Block records a scanned L1 block's (number, hash) for
	// later reorg detection.
	WriteDerivationL1Block(block *db.DerivationL1Block)
	// DeleteDerivationL1BlocksFrom drops saved L1 block records at height >=
	// height; used after a reorg is detected to clear stale hashes.
	DeleteDerivationL1BlocksFrom(height uint64)
}
