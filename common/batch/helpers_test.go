package batch

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	ldberrors "github.com/syndtr/goleveldb/leveldb/errors"
	ldbutil "github.com/syndtr/goleveldb/leveldb/util"
)

type testLevelDB struct {
	db *leveldb.DB
}

func openTestKV(t *testing.T) SealedBatchKV {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "ldb")
	_ = os.RemoveAll(dir)
	db, err := leveldb.OpenFile(dir, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	return &testLevelDB{db: db}
}

func (d *testLevelDB) GetBytes(key []byte) ([]byte, error) {
	v, err := d.db.Get(key, nil)
	if err == ldberrors.ErrNotFound {
		return nil, fmt.Errorf("%w", ErrKeyNotFound)
	}
	return v, err
}

func (d *testLevelDB) PutBytes(key, val []byte) error {
	return d.db.Put(key, val, nil)
}

func (d *testLevelDB) Delete(key []byte) error {
	return d.db.Delete(key, nil)
}

func (d *testLevelDB) WriteBatch(puts []KVPair, deletes [][]byte) error {
	b := new(leveldb.Batch)
	for _, kv := range puts {
		b.Put(kv.Key, kv.Value)
	}
	for _, key := range deletes {
		b.Delete(key)
	}
	return d.db.Write(b, nil)
}

func (d *testLevelDB) IteratePrefixKeys(prefix []byte) ([][]byte, error) {
	it := d.db.NewIterator(ldbutil.BytesPrefix(prefix), nil)
	defer it.Release()
	var keys [][]byte
	for it.Next() {
		k := make([]byte, len(it.Key()))
		copy(k, it.Key())
		keys = append(keys, k)
	}
	return keys, it.Error()
}

func testLoop(ctx context.Context, d time.Duration, fn func()) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fn()
		}
	}
}
