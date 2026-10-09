package main

import (
	"testing"

	"morph-l2/bindings/bindings"
)

func TestRollupMethod(t *testing.T) {
	contractABI, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"commitBatch", "commitBatchWithProof"} {
		args := []interface{}{bindings.IRollupBatchDataInput{}}
		if name == "commitBatchWithProof" {
			args = append(args, []byte{}, []byte{})
		}
		data, err := contractABI.Pack(name, args...)
		if err != nil {
			t.Fatal(err)
		}
		got, err := rollupMethod(data)
		if err != nil || got != name {
			t.Fatalf("method = %q, err = %v", got, err)
		}
	}
	for _, data := range [][]byte{nil, {1, 2, 3}, {1, 2, 3, 4}, contractABI.Methods["commitBatch"].ID} {
		if _, err := rollupMethod(data); err == nil {
			t.Fatalf("accepted invalid calldata %x", data)
		}
	}
}

func TestDataGasCost(t *testing.T) {
	if got := dataGasCost([]byte{0, 1, 255}); got != 36 {
		t.Fatalf("gas = %d, want 36", got)
	}
}
