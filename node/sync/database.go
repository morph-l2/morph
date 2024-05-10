package sync

import "morph-l2/node/types"

type Database interface {
	Reader
	Writer
}

type Reader interface {
	ReadLatestSyncedL1Height() *uint64
	ReadL1MessagesInRange(start, end uint64) []types.L1Message
	ReadL1MessageByIndex(index uint64) *types.L1Message
}

type Writer interface {
	WriteLatestSyncedL1Height(latest uint64)
	WriteSyncedL1Messages(messages []types.L1Message, latest uint64) error
}
