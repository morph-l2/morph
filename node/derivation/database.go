package derivation

import (
	"morph-l2/node/sync"
)

type Database interface {
	Reader
	Writer
	sync.Database
}

type Reader interface {
	ReadLatestDerivationL1Height() *uint64
	ReadBlockNumberByIndex(batchIndex uint64) *uint64
}

type Writer interface {
	WriteLatestDerivationL1Height(latest uint64)
	WriteBatchBlockNumber(batchIndex, blockNumber uint64)
}
