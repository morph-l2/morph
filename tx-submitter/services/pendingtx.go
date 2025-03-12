package services

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
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

	pfinalize       uint64
	commitBatchId   []byte
	finalizeBatchId []byte

	// journal
	journal Journal
}

// NewPendingTxs creates a new PendingTxs instance
func NewPendingTxs(commitBatchMethodId, finalizeBatchMethodId []byte, journal Journal) *PendingTxs {
	return &PendingTxs{
		txinfos:         make(map[common.Hash]*types.TxRecord),
		journal:         journal,
		commitBatchId:   commitBatchMethodId,
		finalizeBatchId: finalizeBatchMethodId,
	}
}

func (pt *PendingTxs) Store(tx *ethtypes.Transaction) error {
	err := pt.journal.AppendTx(tx)
	if err != nil {
		return fmt.Errorf("failed to store tx: %v", err)
	}
	return nil
}

func (pt *PendingTxs) dump() error {
	err := pt.journal.Clean()
	if err != nil {
		return fmt.Errorf("failed to dump tx: %v", err)
	}
	txinfos := pt.getAll()
	for _, info := range txinfos {
		err := pt.journal.AppendTx(info.Tx)
		if err != nil {
			return fmt.Errorf("failed to store tx: %v", err)
		}
	}
	return nil
}

// Add adds a transaction to the pending pool
func (pt *PendingTxs) Add(tx *ethtypes.Transaction) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.txinfos[tx.Hash()] = &types.TxRecord{
		Tx:         tx,
		SendTime:   uint64(time.Now().Unix()),
		QueryTimes: 0,
		Confirmed:  false,
	}

	err := pt.journal.AppendTx(tx)
	if err != nil {
		log.Error("failed to append pending txs", "err", err)
	}
}

// Remove removes a transaction from the pending pool
func (pt *PendingTxs) Remove(txHash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	delete(pt.txinfos, txHash)

	err := pt.dump()
	if err != nil {
		log.Error("failed to dump pending txs", "err", err)
	}
}

// GetAll returns all pending transactions
func (pt *PendingTxs) GetAll() []*types.TxRecord {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	return pt.getAll()
}

func (pt *PendingTxs) getAll() []*types.TxRecord {
	// copy txs and return
	txs := make([]*types.TxRecord, 0, len(pt.txinfos))
	for _, tx := range pt.txinfos {
		txs = append(txs, tx)
	}

	// sort by nonce
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

// ClearConfirmedTxs clears all confirmed transactions
func (pt *PendingTxs) ClearConfirmedTxs() {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	// Mark all transactions as unconfirmed
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
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.pnonce = nonce
}

// SetPindex sets the pending index
func (pt *PendingTxs) SetPindex(index uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.pindex = index
}

// SetPFinalize sets the pending finalize index
func (pt *PendingTxs) SetPFinalize(finalize uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.pfinalize = finalize
}

// ExistedIndex checks if a batch index exists
func (pt *PendingTxs) ExistedIndex(index uint64) bool {
	txs := pt.GetAll()
	abi, _ := bindings.RollupMetaData.GetAbi()

	pt.mu.Lock()
	defer pt.mu.Unlock()

	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i].Tx
		if utils.ParseMethod(tx, abi) == "commitBatch" {
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
	return pt.failedIndex != nil
}

// TrySetFailedBatchIndex tries to set the failed batch index
func (pt *PendingTxs) TrySetFailedBatchIndex(index uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	// failed index must be less than pindex
	if index > pt.pindex {
		return
	}

	pt.failedIndex = &index
}

// RemoveRollupRestriction removes the rollup restriction
func (pt *PendingTxs) RemoveRollupRestriction() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.failedIndex = nil
}

// Recover recovers transactions from the journal
func (pt *PendingTxs) Recover(txs []*ethtypes.Transaction, abi *abi.ABI) {
	// restore state from mempool
	if len(txs) > 0 {
		var pbindex, pfindex uint64

		for _, tx := range txs {
			pt.Add(tx)

			method := utils.ParseMethod(tx, abi)
			if method == "commitBatch" {
				index := utils.ParseParentBatchIndex(tx.Data())
				if index > pbindex {
					pbindex = index
				}
			} else if method == "finalizeBatch" {
				findex := utils.ParseFBatchIndex(tx.Data())
				if findex > pfindex {
					pfindex = findex
				}
			}
		}

		pt.SetPindex(pbindex)
		pt.SetPFinalize(pfindex)
		pt.SetNonce(txs[len(txs)-1].Nonce())

		log.Info("Recover from mempool",
			"tx_count", len(txs),
			"max_batch_index", pbindex,
			"max_finalize_index", pfindex,
			"max_nonce", pt.pnonce,
		)
	} else {
		log.Info("journal tx is empty")
	}
}
