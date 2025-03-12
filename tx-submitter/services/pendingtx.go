package services

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
)

const (
	// MethodCommitBatch is the method name for committing a batch
	MethodCommitBatch = "commitBatch"
	// MethodFinalizeBatch is the method name for finalizing a batch
	MethodFinalizeBatch = "finalizeBatch"
)

// Journal defines the interface for transaction journaling
type Journal interface {
	Init() error
	AppendTx(tx *ethtypes.Transaction) error
	ParseAllTxs() ([]*ethtypes.Transaction, error)
	Clean() error
}

// PendingTxs represents the pending transactions pool
type PendingTxs struct {
	mu sync.RWMutex

	txinfos map[common.Hash]*types.TxRecord
	pnonce  uint64 // pending nonce

	failedIndex *uint64
	pindex      uint64 // pending batch index
	pfinalize   uint64

	commitBatchId   []byte
	finalizeBatchId []byte

	journal Journal
}

// NewPendingTxs creates a new PendingTxs instance
func NewPendingTxs(commitBatchMethodId, finalizeBatchMethodId []byte, journal Journal) *PendingTxs {
	pt := &PendingTxs{
		txinfos:         make(map[common.Hash]*types.TxRecord),
		journal:         journal,
		commitBatchId:   commitBatchMethodId,
		finalizeBatchId: finalizeBatchMethodId,
	}
	return pt
}

// Store persists a transaction to the journal
func (pt *PendingTxs) Store(tx *ethtypes.Transaction) error {
	if tx == nil {
		return fmt.Errorf("cannot store nil transaction")
	}
	if err := pt.journal.AppendTx(tx); err != nil {
		return fmt.Errorf("failed to store tx: %w", err)
	}
	return nil
}

// dump persists all transactions to the journal
func (pt *PendingTxs) dump() error {
	if err := pt.journal.Clean(); err != nil {
		return fmt.Errorf("failed to clean journal: %w", err)
	}

	txinfos := pt.getAll()
	for _, info := range txinfos {
		if err := pt.journal.AppendTx(info.Tx); err != nil {
			return fmt.Errorf("failed to store tx: %w", err)
		}
	}
	return nil
}

// Add adds a transaction to the pending pool
func (pt *PendingTxs) Add(tx *ethtypes.Transaction) error {
	if tx == nil {
		return fmt.Errorf("cannot add nil transaction")
	}

	pt.mu.Lock()
	defer pt.mu.Unlock()

	hash := tx.Hash()
	pt.txinfos[hash] = &types.TxRecord{
		Tx:         tx,
		SendTime:   uint64(time.Now().Unix()),
		QueryTimes: 0,
		Confirmed:  false,
	}

	if err := pt.journal.AppendTx(tx); err != nil {
		delete(pt.txinfos, hash)
		return fmt.Errorf("failed to append tx to journal: %w", err)
	}

	return nil
}

// Remove removes a transaction from the pending pool
func (pt *PendingTxs) Remove(txHash common.Hash) error {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	if _, exists := pt.txinfos[txHash]; !exists {
		return nil // tx already removed
	}

	delete(pt.txinfos, txHash)
	return pt.dump()
}

// GetAll returns all pending transactions in nonce order
func (pt *PendingTxs) GetAll() []*types.TxRecord {
	pt.mu.RLock()
	defer pt.mu.RUnlock()
	return pt.getAll()
}

func (pt *PendingTxs) getAll() []*types.TxRecord {
	txs := make([]*types.TxRecord, 0, len(pt.txinfos))
	for _, tx := range pt.txinfos {
		txs = append(txs, tx)
	}

	sort.SliceStable(txs, func(i, j int) bool {
		return txs[i].Tx.Nonce() < txs[j].Tx.Nonce()
	})

	return txs
}

// GetTxRecord returns a transaction record by its hash
func (pt *PendingTxs) GetTxRecord(hash common.Hash) *types.TxRecord {
	pt.mu.RLock()
	defer pt.mu.RUnlock()
	return pt.txinfos[hash]
}

// IncQueryTimes increments the query times for a transaction
func (pt *PendingTxs) IncQueryTimes(txHash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	if tx, ok := pt.txinfos[txHash]; ok {
		tx.QueryTimes++
	}
}

// MarkConfirmed marks a transaction as confirmed
func (pt *PendingTxs) MarkConfirmed(hash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	if record, exists := pt.txinfos[hash]; exists {
		record.Confirmed = true
	}
}

// ClearConfirmedTxs marks all confirmed transactions as unconfirmed
func (pt *PendingTxs) ClearConfirmedTxs() {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	for _, record := range pt.txinfos {
		if record.Confirmed {
			log.Info("Marking transaction as unconfirmed due to reorg",
				"hash", record.Tx.Hash().String(),
				"nonce", record.Tx.Nonce())
			record.Confirmed = false
		}
	}
}

// SetNonce sets the pending nonce
func (pt *PendingTxs) SetNonce(nonce uint64) {
	atomic.StoreUint64(&pt.pnonce, nonce)
}

// GetNonce gets the current pending nonce
func (pt *PendingTxs) GetNonce() uint64 {
	return atomic.LoadUint64(&pt.pnonce)
}

// SetPindex sets the pending index
func (pt *PendingTxs) SetPindex(index uint64) {
	atomic.StoreUint64(&pt.pindex, index)
}

// GetPindex gets the current pending index
func (pt *PendingTxs) GetPindex() uint64 {
	return atomic.LoadUint64(&pt.pindex)
}

// SetPFinalize sets the pending finalize index
func (pt *PendingTxs) SetPFinalize(finalize uint64) {
	atomic.StoreUint64(&pt.pfinalize, finalize)
}

// GetPFinalize gets the current pending finalize index
func (pt *PendingTxs) GetPFinalize() uint64 {
	return atomic.LoadUint64(&pt.pfinalize)
}

// ExistedIndex checks if a batch index exists
func (pt *PendingTxs) ExistedIndex(index uint64) bool {
	txs := pt.GetAll() // already has RLock
	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		log.Error("Failed to get ABI", "err", err)
		return false
	}

	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i].Tx
		if utils.ParseMethod(tx, abi) == MethodCommitBatch {
			pindex := utils.ParseParentBatchIndex(tx.Data()) + 1
			if index == pindex {
				return true
			}
		}
	}
	return false
}

// HaveFailed checks if there are any failed transactions
func (pt *PendingTxs) HaveFailed() bool {
	pt.mu.RLock()
	defer pt.mu.RUnlock()
	return pt.failedIndex != nil
}

// TrySetFailedBatchIndex tries to set the failed batch index
func (pt *PendingTxs) TrySetFailedBatchIndex(index uint64) {
	currentPindex := pt.GetPindex()
	if index > currentPindex {
		return
	}

	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.failedIndex = &index
}

// RemoveRollupRestriction removes the rollup restriction
func (pt *PendingTxs) RemoveRollupRestriction() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.failedIndex = nil
}

// Recover recovers transactions from the journal
func (pt *PendingTxs) Recover(txs []*ethtypes.Transaction, abi *abi.ABI) error {
	if len(txs) == 0 {
		log.Info("Journal tx is empty")
		return nil
	}

	var pbindex, pfindex uint64
	for _, tx := range txs {
		if err := pt.Add(tx); err != nil {
			return fmt.Errorf("failed to add tx during recovery: %w", err)
		}

		method := utils.ParseMethod(tx, abi)
		switch method {
		case MethodCommitBatch:
			index := utils.ParseParentBatchIndex(tx.Data())
			if index > pbindex {
				pbindex = index
			}
		case MethodFinalizeBatch:
			findex := utils.ParseFBatchIndex(tx.Data())
			if findex > pfindex {
				pfindex = findex
			}
		}
	}

	pt.SetPindex(pbindex)
	pt.SetPFinalize(pfindex)
	pt.SetNonce(txs[len(txs)-1].Nonce())

	log.Info("Recovered from mempool",
		"tx_count", len(txs),
		"max_batch_index", pbindex,
		"max_finalize_index", pfindex,
		"max_nonce", pt.GetNonce(),
	)

	return nil
}
