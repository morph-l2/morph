package services

import (
	"bytes"
	"sort"
	"sync"
	"time"

	"github.com/morph-l2/tx-submitter/utils"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

type TxInfo struct {
	sendTime uint64
	tx       types.Transaction

	queryTimes uint64
}

type PendingTxs struct {
	mu sync.Mutex

	txinfos map[common.Hash]TxInfo
	pnonce  uint64 // pending nonce
	pindex  uint64 // pending batch index

	commitBatchId   []byte
	finalizeBatchId []byte
}

func NewPendingTxs(cid []byte, fid []byte) *PendingTxs {
	return &PendingTxs{
		txinfos:         make(map[common.Hash]TxInfo),
		commitBatchId:   cid,
		finalizeBatchId: fid,
	}
}

func (pt *PendingTxs) Add(tx types.Transaction) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.txinfos[tx.Hash()] = TxInfo{
		sendTime: uint64(time.Now().Unix()),
		tx:       tx,
	}

	// rollup tx: commitBatch
	if len(tx.Data()) > 0 && bytes.Equal(tx.Data()[:4], pt.commitBatchId) {
		pt.pindex = utils.ParseParentBatchIndex(tx.Data()) + 1
	}
	pt.pnonce = tx.Nonce()
}

func (pt *PendingTxs) Remove(txHash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	delete(pt.txinfos, txHash)
}

func (pt *PendingTxs) GetAll() []TxInfo {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	// copy txs and return
	txs := make([]TxInfo, 0, len(pt.txinfos))
	for _, tx := range pt.txinfos {
		txs = append(txs, tx)
	}

	// sort by nonce
	sort.SliceStable(txs, func(i, j int) bool {
		return txs[i].tx.Nonce() < txs[j].tx.Nonce()
	})

	return txs
}

func (pt *PendingTxs) Get(txHash common.Hash) (TxInfo, bool) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	tx, ok := pt.txinfos[txHash]
	return tx, ok
}
