package db

import (
	"encoding/binary"

	"github.com/morph-l2/go-ethereum/ethdb"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/rlp"

	"morph-l2/node/types"
)

// L1MessageIterator is a wrapper around ethdb.Iterator that
// allows us to iterate over L1 messages in the database. It
// implements an interface similar to ethdb.Iterator.
type L1MessageIterator struct {
	inner     ethdb.Iterator
	keyLength int
}

// IterateL1MessagesFrom creates an L1MessageIterator that iterates over
// all L1 message in the database starting at the provided enqueue index.
func IterateL1MessagesFrom(db ethdb.Iteratee, fromEnqueueIndex uint64) L1MessageIterator {
	start := encodeEnqueueIndex(fromEnqueueIndex)
	it := db.NewIterator(L1MessagePrefix, start)
	keyLength := len(L1MessagePrefix) + 8

	return L1MessageIterator{
		inner:     it,
		keyLength: keyLength,
	}
}

// Next moves the iterator to the next key/value pair.
// It returns false when the iterator is exhausted.
func (it *L1MessageIterator) Next() bool {
	for it.inner.Next() {
		key := it.inner.Key()
		if len(key) == it.keyLength {
			return true
		}
	}
	return false
}

// EnqueueIndex returns the enqueue index of the current L1 message.
func (it *L1MessageIterator) EnqueueIndex() uint64 {
	key := it.inner.Key()
	raw := key[len(L1MessagePrefix) : len(L1MessagePrefix)+8]
	enqueueIndex := binary.BigEndian.Uint64(raw)
	return enqueueIndex
}

// L1Message returns the current L1 message.
func (it *L1MessageIterator) L1Message() types.L1Message {
	data := it.inner.Value()
	var l1Msg types.L1Message
	if err := rlp.DecodeBytes(data, &l1Msg); err != nil {
		log.Crit("Invalid L1 message RLP", "data", data, "err", err)
	}
	return l1Msg
}

// Release releases the associated resources.
func (it *L1MessageIterator) Release() {
	it.inner.Release()
}
