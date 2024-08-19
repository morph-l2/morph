package utils

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"morph-l2/bindings/bindings"
	ntype "morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/rpc"
)

// Loop Run the f func periodically.
func Loop(ctx context.Context, period time.Duration, f func()) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for ; ; <-tick.C {
		select {
		case <-ctx.Done():
			return
		default:
			f()
		}
	}
}

func ParseFBatchIndex(calldata []byte) uint64 {
	abi, _ := bindings.RollupMetaData.GetAbi()
	parms, _ := abi.Methods["finalizeBatch"].Inputs.Unpack(calldata[4:])
	batchHeader, _ := ntype.DecodeBatchHeader(parms[0].([]byte))
	return batchHeader.BatchIndex
}

func ParseParentBatchIndex(calldata []byte) uint64 {
	///   * Field                   Bytes       Type        Index   Comments
	///   * version                 1           uint8       0       The batch version
	///   * batchIndex              8           uint64      1       The index of the batch
	///   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch

	abi, _ := bindings.RollupMetaData.GetAbi()
	parms, _ := abi.Methods["commitBatch"].Inputs.UnpackValues(calldata[4:])
	v := reflect.ValueOf(parms[0])
	pbh := v.FieldByName("ParentBatchHeader")
	batchIndex := binary.BigEndian.Uint64(pbh.Bytes()[1:9])
	return batchIndex
}

// ParseL1Mempool parses the L1 mempool and returns the transactions.
func ParseL1Mempool(rpc *rpc.Client, addr common.Address) ([]*types.Transaction, error) {

	var result map[string]map[string]*types.Transaction
	err := rpc.Call(&result, "txpool_contentFrom", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get txpool content: %v", err)
	}

	var txs []*types.Transaction

	// get pending txs
	if pendingTxs, ok := result["pending"]; ok {
		for _, tx := range pendingTxs {
			txs = append(txs, tx)
		}
	}

	// get queued txs
	if pendingTxs, ok := result["queued"]; ok {
		for _, tx := range pendingTxs {
			txs = append(txs, tx)
		}
	}

	return txs, nil

}

func ParseMempoolLatestBatchIndex(id []byte, txs []*types.Transaction) uint64 {

	var res uint64
	for _, tx := range txs {
		if bytes.Equal(tx.Data()[:4], id) {
			pindex := ParseParentBatchIndex(tx.Data())
			if pindex > res {
				res = pindex
			}
		}
	}

	return res + 1

}

func ParseBusinessInfo(tx *types.Transaction, a *abi.ABI) []interface{} {
	// var method string
	// var batchIndex uint64
	// var finalizedIndex uint64
	var res []interface{}
	if len(tx.Data()) > 0 {
		id := tx.Data()[:4]
		if bytes.Equal(id, a.Methods["commitBatch"].ID) {
			method := "commitBatch"
			batchIndex := ParseParentBatchIndex(tx.Data()) + 1
			res = append(res,
				"method", method,
				"batchIndex", batchIndex,
			)
		} else if bytes.Equal(id, a.Methods["finalizeBatch"].ID) {
			method := "finalizeBatch"
			parms, err := a.Methods["finalizeBatch"].Inputs.Unpack(tx.Data()[4:])
			if err != nil {
				log.Error("unpack finalizeBatch error", "err", err)
			}
			batchHeader, _ := ntype.DecodeBatchHeader(parms[0].([]byte))
			res = append(res,
				"method", method,
				"finalizedIndex", batchHeader.BatchIndex,
			)

		}

	} else {
		return []interface{}{}
	}
	return res
}

func ParseMethod(tx *types.Transaction, a *abi.ABI) string {
	if tx.Data() == nil || len(tx.Data()) < 4 {
		return ""
	}
	id := tx.Data()[:4]
	if bytes.Equal(id, a.Methods["commitBatch"].ID) {
		return "commitBatch"
	} else if bytes.Equal(id, a.Methods["finalizeBatch"].ID) {
		return "finalizeBatch"
	} else {
		return ""
	}
}

func ParseNonce(s string) (uint64, uint64, error) {
	re := regexp.MustCompile(`\d+`)

	// Find all matches
	matches := re.FindAllString(s, -1)

	if len(matches) >= 2 {
		// Convert strings to integers
		n1, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0, 0, fmt.Errorf("convert nonce err: %w", err)
		}

		n2, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, 0, fmt.Errorf("convert nonce err: %w", err)
		}

		return uint64(n1), uint64(n2), nil
	} else {
		return 0, 0, fmt.Errorf("expect 2 nonce")
	}
}

func ParseL1MessageCnt(chunks []hexutil.Bytes) uint64 {

	var l1msgcnt uint64

	for _, v := range chunks {
		bknm := v[0]
		chunkbs := v[1:]
		for i := 0; i < int(bknm); i++ {
			l1msgcnt += uint64(binary.BigEndian.Uint16(chunkbs[58:60]))
			chunkbs = chunkbs[60:]
		}

	}

	return l1msgcnt
}
