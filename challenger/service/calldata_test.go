package service

import (
	"testing"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/bindings/bindings"
)

func TestBatchInputFromCalldata(t *testing.T) {
	parsed, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	var state, withdraw [32]byte
	copy(state[:], common.HexToHash("0x11").Bytes())
	copy(withdraw[:], common.HexToHash("0x22").Bytes())
	input := bindings.IRollupBatchDataInput{
		Version:         1,
		LastBlockNumber: 40,
		PostStateRoot:   state,
		WithdrawalRoot:  withdraw,
	}
	data, err := parsed.Pack("commitBatch", input)
	if err != nil {
		t.Fatal(err)
	}
	got, err := batchInputFromCalldata(data)
	if err != nil {
		t.Fatal(err)
	}
	if got.BlockNumber != 40 || got.StateRoot != common.BytesToHash(state[:]) || got.WithdrawRoot != common.BytesToHash(withdraw[:]) {
		t.Fatalf("input = %+v", got)
	}

	withProof, err := parsed.Pack("commitBatchWithProof", input, []byte{1}, []byte{2})
	if err != nil {
		t.Fatal(err)
	}
	got, err = batchInputFromCalldata(withProof)
	if err != nil {
		t.Fatal(err)
	}
	if got.BlockNumber != 40 || got.WithdrawRoot != common.BytesToHash(withdraw[:]) {
		t.Fatalf("proof input = %+v", got)
	}

	if _, err := batchInputFromCalldata([]byte{0, 1, 2, 3, 4}); err == nil {
		t.Fatal("expected unknown selector to fail")
	}
}
