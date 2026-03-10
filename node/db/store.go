package db

import (
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/morph-l2/go-ethereum/core/rawdb"
	"github.com/morph-l2/go-ethereum/ethdb"
	"github.com/morph-l2/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"

	"morph-l2/node/types"
)

type Store struct {
	db ethdb.Database
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
		db: db,
	}, nil
}

func (s *Store) ReadLatestDerivationL1Height() *uint64 {
	data, err := s.db.Get(derivationL1HeightKey)
	if err != nil && !isNotFoundErr(err) {
		panic(fmt.Sprintf("Failed to read batch index from database,err:%v", err))
	}
	if len(data) == 0 {
		return nil
	}

	number := new(big.Int).SetBytes(data)
	if !number.IsUint64() {
		panic(fmt.Sprintf("unexpected derivation l1 height in database, number: %d", number))
	}

	value := number.Uint64()
	return &value
}

func (s *Store) ReadLatestSyncedL1Height() *uint64 {
	data, err := s.db.Get(syncedL1HeightKey)
	if err != nil && !isNotFoundErr(err) {
		panic(fmt.Sprintf("failed to read synced L1 block number from database, err: %v", err))
	}
	if len(data) == 0 {
		return nil
	}

	number := new(big.Int).SetBytes(data)
	if !number.IsUint64() {
		panic(fmt.Sprintf("unexpected synced L1 block number in database, number: %d", number))
	}

	value := number.Uint64()
	return &value
}

func (s *Store) ReadL1MessagesInRange(start, end uint64) []types.L1Message {
	if start > end {
		return nil
	}
	//expectedCount := end - start + 1
	//messages := make([]types.L1Message, 0, expectedCount)
	var messages []types.L1Message
	it := IterateL1MessagesFrom(s.db, start)
	defer it.Release()

	for it.Next() {
		if it.EnqueueIndex() > end {
			break
		}
		messages = append(messages, it.L1Message())
	}

	return messages
}

func (s *Store) ReadL1MessageByIndex(index uint64) *types.L1Message {
	data, err := s.db.Get(L1MessageKey(index))
	if err != nil && !isNotFoundErr(err) {
		panic(fmt.Sprintf("failed to read L1 message from database, err: %v", err))
	}
	if len(data) == 0 {
		return nil
	}
	var l1Msg types.L1Message
	if err := rlp.DecodeBytes(data, &l1Msg); err != nil {
		panic(fmt.Sprintf("invalid L1 message RLP, err: %v", err))

	}
	return &l1Msg
}

func (s *Store) WriteLatestDerivationL1Height(latest uint64) {
	if err := s.db.Put(derivationL1HeightKey, new(big.Int).SetUint64(latest).Bytes()); err != nil {
		panic(fmt.Sprintf("failed to update derivation synced L1 height, err: %v", err))
	}
}

func (s *Store) WriteLatestSyncedL1Height(latest uint64) {
	if err := s.db.Put(syncedL1HeightKey, new(big.Int).SetUint64(latest).Bytes()); err != nil {
		panic(fmt.Sprintf("failed to update synced L1 height, err: %v", err))
	}
}

func (s *Store) WriteSyncedL1Messages(messages []types.L1Message, latestSynced uint64) error {
	if len(messages) == 0 {
		return nil
	}
	batch := s.db.NewBatch()
	for _, msg := range messages {
		bytes, err := rlp.EncodeToBytes(msg)
		if err != nil {
			panic(fmt.Sprintf("failed to RLP encode L1 message, err: %v", err))
		}
		enqueueIndex := msg.QueueIndex
		if err := batch.Put(L1MessageKey(enqueueIndex), bytes); err != nil {
			panic(fmt.Sprintf("failed to store L1 message, err: %v", err))
		}
	}
	if err := batch.Put(syncedL1HeightKey, new(big.Int).SetUint64(latestSynced).Bytes()); err != nil {
		panic(fmt.Sprintf("failed to update synced L1 height, err: %v", err))
	}
	return batch.Write()
}

// DerivationL1Block stores L1 block info for reorg detection.
type DerivationL1Block struct {
	Number     uint64
	Hash       [32]byte
	BatchIndex uint64 // 0 means no batch in this block
	L2EndBlock uint64 // last L2 block number in the batch (0 if no batch)
}

func (s *Store) WriteDerivationL1Block(block *DerivationL1Block) {
	data, err := rlp.EncodeToBytes(block)
	if err != nil {
		panic(fmt.Sprintf("failed to RLP encode DerivationL1Block, err: %v", err))
	}
	if err := s.db.Put(DerivationL1BlockKey(block.Number), data); err != nil {
		panic(fmt.Sprintf("failed to write DerivationL1Block, err: %v", err))
	}
}

func (s *Store) ReadDerivationL1Block(l1Height uint64) *DerivationL1Block {
	data, err := s.db.Get(DerivationL1BlockKey(l1Height))
	if err != nil && !isNotFoundErr(err) {
		panic(fmt.Sprintf("failed to read DerivationL1Block, err: %v", err))
	}
	if len(data) == 0 {
		return nil
	}
	var block DerivationL1Block
	if err := rlp.DecodeBytes(data, &block); err != nil {
		panic(fmt.Sprintf("invalid DerivationL1Block RLP, err: %v", err))
	}
	return &block
}

func (s *Store) ReadDerivationL1BlockRange(from, to uint64) []*DerivationL1Block {
	var blocks []*DerivationL1Block
	for h := from; h <= to; h++ {
		b := s.ReadDerivationL1Block(h)
		if b != nil {
			blocks = append(blocks, b)
		}
	}
	return blocks
}

func (s *Store) DeleteDerivationL1BlocksFrom(height uint64) {
	batch := s.db.NewBatch()
	for h := height; ; h++ {
		key := DerivationL1BlockKey(h)
		has, err := s.db.Has(key)
		if err != nil || !has {
			break
		}
		if err := batch.Delete(key); err != nil {
			panic(fmt.Sprintf("failed to delete DerivationL1Block at %d, err: %v", h, err))
		}
	}
	if err := batch.Write(); err != nil {
		panic(fmt.Sprintf("failed to write batch delete for DerivationL1Blocks, err: %v", err))
	}
}

func isNotFoundErr(err error) bool {
	return err.Error() == leveldb.ErrNotFound.Error() || err.Error() == types.ErrMemoryDBNotFound.Error()
}
