package service

import (
	"bytes"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/bindings/bindings"
)

// commitInput is the batch data carried by the commit transaction.
// The contract deletes committedStateRoots and batchDataStore when the next batch
// finalizes, so this is the record that remains for an older batch.
type commitInput struct {
	BlockNumber  uint64
	StateRoot    common.Hash
	WithdrawRoot common.Hash
}

// batchInputFromCalldata reads BatchDataInput from a commit transaction.
// commitBatch, commitState, and commitBatchWithProof all carry that struct as the first argument.
func batchInputFromCalldata(data []byte) (commitInput, error) {
	if len(data) < 4 {
		return commitInput{}, fmt.Errorf("commit calldata is too short")
	}
	parsed, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return commitInput{}, fmt.Errorf("parse rollup abi: %w", err)
	}
	method, ok := methodBySelector(parsed, data[:4])
	if !ok {
		return commitInput{}, fmt.Errorf("commit calldata has no batch data")
	}
	switch method.Name {
	case "commitBatch", "commitState", "commitBatchWithProof":
	default:
		return commitInput{}, fmt.Errorf("commit calldata has no batch data")
	}
	args, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return commitInput{}, fmt.Errorf("unpack %s: %w", method.Name, err)
	}
	if len(args) == 0 {
		return commitInput{}, fmt.Errorf("%s has no batch data", method.Name)
	}
	input, ok := abi.ConvertType(args[0], new(bindings.IRollupBatchDataInput)).(*bindings.IRollupBatchDataInput)
	if !ok || input == nil {
		return commitInput{}, fmt.Errorf("decode %s batch data", method.Name)
	}
	return commitInput{
		BlockNumber:  input.LastBlockNumber,
		StateRoot:    common.BytesToHash(input.PostStateRoot[:]),
		WithdrawRoot: common.BytesToHash(input.WithdrawalRoot[:]),
	}, nil
}

func methodBySelector(parsed *abi.ABI, selector []byte) (*abi.Method, bool) {
	for _, method := range parsed.Methods {
		if bytes.Equal(method.ID, selector) {
			found := method
			return &found, true
		}
	}
	return nil, false
}
