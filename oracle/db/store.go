package db

import (
	"fmt"

	"path/filepath"
	"sync"

	"github.com/morph-l2/go-ethereum/core/rawdb"
	"github.com/morph-l2/go-ethereum/ethdb"
	"github.com/morph-l2/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"

	"morph-l2/oracle/types"
)

type Store struct {
	db             ethdb.Database
	ChainPointSync sync.Mutex
}

func NewMemoryStore() *Store {
	return &Store{
		db: rawdb.NewMemoryDatabase(),
	}
}

func NewStore(config *Config, home string) (*Store, error) {
	var (
		db      ethdb.Database
		err     error
		dbPath  = config.DBPath
		freezer = config.DatabaseFreezer
	)

	if dbPath == "" {
		if home == "" {
			return nil, fmt.Errorf("either Home or DB path has to be provided")
		}
		dbPath = filepath.Join(home, "node-data")
	}

	if config.DatabaseFreezer == "" {
		freezer = filepath.Join(dbPath, "ancient")
	}
	db, err = rawdb.NewLevelDBDatabaseWithFreezer(dbPath, config.DatabaseCache, config.DatabaseHandles, freezer, config.Namespace, false)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:             db,
		ChainPointSync: sync.Mutex{},
	}, nil
}

func (s *Store) WriteLatestChangeContext(changePoints types.ChangeContext) error {
	data, err := rlp.EncodeToBytes(changePoints)
	if err != nil {
		return err
	}
	if err := s.db.Put(changePointsKey, data); err != nil {
		panic(fmt.Sprintf("failed to update change points failed, err: %v", err))
	}
	return nil
}

func (s *Store) ReadLatestChangePoints() types.ChangeContext {
	data, err := s.db.Get(changePointsKey)
	if err != nil && !isNotFoundErr(err) {
		panic(fmt.Sprintf("failed to read change points, err: %v", err))
	}
	if err != nil {
		panic(fmt.Sprintf("failed to sync change points, err: %v", err))
	}
	var changeCtx types.ChangeContext
	if err := rlp.DecodeBytes(data, &changeCtx); err != nil {
		panic(fmt.Sprintf("decode data to changepoint error:%v", err))
	}
	return changeCtx
}

func isNotFoundErr(err error) bool {
	return err.Error() == leveldb.ErrNotFound.Error() || err.Error() == types.ErrMemoryDBNotFound.Error()
}
