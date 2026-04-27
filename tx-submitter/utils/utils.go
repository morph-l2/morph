package utils

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
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
	if len(calldata) < 4 {
		return 0
	}

	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return 0
	}

	method, exists := abi.Methods["finalizeBatch"]
	if !exists {
		return 0
	}

	parms, err := method.Inputs.Unpack(calldata[4:])
	if err != nil || len(parms) == 0 {
		return 0
	}

	batchBytes, ok := parms[0].([]byte)
	if !ok || len(batchBytes) < 9 {
		return 0
	}

	// 1-9 is batch index
	return binary.BigEndian.Uint64(batchBytes[1:9])
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

// SetFBatchIndex sets the batch index in the calldata while preserving all other data
func SetFBatchIndex(calldata []byte, batchIndex uint64) error {
	if len(calldata) < 4 {
		return fmt.Errorf("calldata too short")
	}

	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ABI: %w", err)
	}

	method, exists := abi.Methods["finalizeBatch"]
	if !exists {
		return fmt.Errorf("finalizeBatch method not found in ABI")
	}

	parms, err := method.Inputs.Unpack(calldata[4:])
	if err != nil || len(parms) == 0 {
		return fmt.Errorf("failed to unpack parameters: %w", err)
	}

	batchBytes, ok := parms[0].([]byte)
	if !ok || len(batchBytes) < 9 {
		return fmt.Errorf("invalid batch bytes")
	}

	// Modify only the batch index (bytes 1-9) while keeping other data unchanged
	binary.BigEndian.PutUint64(batchBytes[1:9], batchIndex)

	// Re-encode the parameters
	encodedParams, err := method.Inputs.Pack(batchBytes)
	if err != nil {
		return fmt.Errorf("failed to pack parameters: %w", err)
	}

	// Update only the parameter portion, keeping the method ID unchanged
	copy(calldata[4:], encodedParams)
	return nil
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

func ParseL1MessageCnt(blockContexts hexutil.Bytes) uint64 {

	var l1msgcnt uint64
	blockNum := binary.BigEndian.Uint16(blockContexts[:2])
	remainingBz := blockContexts[2:]
	for i := 0; i < int(blockNum); i++ {
		l1msgcnt += uint64(binary.BigEndian.Uint16(remainingBz[58:60]))
		remainingBz = remainingBz[60:]
	}

	return l1msgcnt
}

// FormatTime formats a timestamp into RFC3339 format string.
// Returns "N/A" for nil or non-positive timestamps.
func FormatTime(timestamp *big.Int) string {
	if timestamp == nil || timestamp.Int64() <= 0 {
		return "N/A"
	}
	return time.Unix(timestamp.Int64(), 0).Format(time.RFC3339)
}
