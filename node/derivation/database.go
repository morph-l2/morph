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
	//ReadLatestBatchBls() types.BatchBls
}

type Writer interface {
	WriteLatestDerivationL1Height(latest uint64)
	//WriteLatestBatchBls(batchBls types.BatchBls)
}
