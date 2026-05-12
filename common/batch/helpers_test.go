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
