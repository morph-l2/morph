package utils

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"time"

	"morph-l2/bindings/bindings"
	ntype "morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
)

var ErrNotOwnedCommit = errors.New("calldata is not an owned current commitBatch")

// ParseOwnedCommit accepts only the current, post-cutover commitBatch selector.
// Historical commitBatch, commitState, commitBatchWithProof, and unknown
// calldata must never be imported into the official writer's pending state.
func ParseOwnedCommit(data []byte, current *abi.ABI) (methodName string, input *bindings.IRollupBatchDataInput, err error) {
	if current == nil || len(data) < 4 {
		return "", nil, ErrNotOwnedCommit
	}
	method, ok := current.Methods["commitBatch"]
	if !ok || !bytes.Equal(data[:4], method.ID) {
		return "", nil, ErrNotOwnedCommit
	}
	values, err := method.Inputs.Unpack(data[4:])
	if err != nil || len(values) != 1 {
		if err == nil {
			err = errors.New("unexpected commitBatch argument count")
		}
		return "", nil, fmt.Errorf("decode owned commitBatch: %w", err)
	}
	defer func() {
		if recovered := recover(); recovered != nil {
			methodName = ""
			input = nil
			err = fmt.Errorf("decode owned commitBatch tuple: %v", recovered)
		}
	}()
	converted := abi.ConvertType(values[0], new(bindings.IRollupBatchDataInput))
	input, ok = converted.(*bindings.IRollupBatchDataInput)
	if !ok || input == nil {
		return "", nil, errors.New("decode owned commitBatch: invalid batch tuple")
	}
	return method.Name, input, nil
}

func ParseOwnedCommitBatchIndex(data []byte, current *abi.ABI) (uint64, error) {
	_, input, err := ParseOwnedCommit(data, current)
	if err != nil {
		return 0, err
	}
	if len(input.ParentBatchHeader) < 9 {
		return 0, errors.New("owned commitBatch parent header is too short")
	}
	parentIndex := binary.BigEndian.Uint64(input.ParentBatchHeader[1:9])
	if parentIndex == ^uint64(0) {
		return 0, errors.New("owned commitBatch parent index overflows")
	}
	return parentIndex + 1, nil
}

// ParseParentBatchIndex is retained for non-authoritative diagnostics. It is
// strict: only the current commitBatch selector is accepted.
func ParseParentBatchIndex(data []byte) uint64 {
	current, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return 0
	}
	batchIndex, err := ParseOwnedCommitBatchIndex(data, current)
	if err != nil || batchIndex == 0 {
		return 0
	}
	return batchIndex - 1
}

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

func ParseBusinessInfo(tx *types.Transaction, a *abi.ABI) []interface{} {
	// var method string
	// var batchIndex uint64
	// var finalizedIndex uint64
	var res []interface{}
	if len(tx.Data()) >= 4 {
		id := tx.Data()[:4]
		if bytes.Equal(id, a.Methods["commitBatch"].ID) {
			method := "commitBatch"
			batchIndex, err := ParseOwnedCommitBatchIndex(tx.Data(), a)
			if err != nil {
				return []interface{}{}
			}
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
			batchIndex, _ := ntype.BatchHeaderBytes(parms[0].([]byte)).BatchIndex()
			res = append(res,
				"method", method,
				"finalizedIndex", batchIndex,
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
