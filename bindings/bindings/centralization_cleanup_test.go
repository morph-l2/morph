package bindings

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestCurrentRollupBindingShape(t *testing.T) {
	current, err := RollupMetaData.GetAbi()
	if err != nil || current == nil {
		t.Fatalf("parse generated Rollup ABI: %v", err)
	}
	if strings.Contains(RollupMetaData.ABI, "BatchSignatureInput") {
		t.Fatal("generated Rollup binding still contains BatchSignatureInput")
	}

	expected := map[string]struct {
		selector string
		inputs   int
	}{
		"commitBatch":          {selector: "41f756da", inputs: 1},
		"commitState":          {selector: "67caa37a", inputs: 1},
		"commitBatchWithProof": {selector: "1544ba3a", inputs: 3},
		"initialize4":          {selector: "f8fa010f", inputs: 1},
	}
	for name, want := range expected {
		method, ok := current.Methods[name]
		if !ok {
			t.Fatalf("generated Rollup ABI is missing %s", name)
		}
		if got := hex.EncodeToString(method.ID); got != want.selector {
			t.Fatalf("%s selector = %s, want %s", name, got, want.selector)
		}
		if got := len(method.Inputs); got != want.inputs {
			t.Fatalf("%s input count = %d, want %d", name, got, want.inputs)
		}
	}

	batchDataStore, ok := current.Methods["batchDataStore"]
	if !ok {
		t.Fatal("generated Rollup ABI is missing batchDataStore")
	}
	if len(batchDataStore.Outputs) != 4 || batchDataStore.Outputs[3].Type.String() != "address" {
		t.Fatalf("batchDataStore fourth output must be address: %v", batchDataStore.Outputs)
	}
	if _, ok := current.Methods["pendingBatchCount"]; ok {
		t.Fatal("generated Rollup binding must not contain pendingBatchCount")
	}
}

func TestCurrentSubmitterBindingShape(t *testing.T) {
	current, err := SubmitterMetaData.GetAbi()
	if err != nil || current == nil {
		t.Fatalf("parse generated Submitter ABI: %v", err)
	}
	for _, name := range []string{
		"addSubmitter",
		"challengeDeposit",
		"claimSlashRemaining",
		"claimWithdrawal",
		"initialize",
		"isActive",
		"removeSubmitter",
		"slash",
		"stake",
		"withdraw",
		"withdrawalBatchIndex",
	} {
		if _, ok := current.Methods[name]; !ok {
			t.Fatalf("generated Submitter ABI is missing %s", name)
		}
	}

	method := current.Methods["withdrawalBatchIndex"]
	if got := hex.EncodeToString(method.ID); got != "3bf1944b" {
		t.Fatalf("withdrawalBatchIndex selector = %s, want 3bf1944b", got)
	}
	if len(method.Inputs) != 1 || method.Inputs[0].Type.String() != "address" {
		t.Fatalf("withdrawalBatchIndex input must be address: %v", method.Inputs)
	}
	if len(method.Outputs) != 1 || method.Outputs[0].Type.String() != "uint256" {
		t.Fatalf("withdrawalBatchIndex output must be uint256: %v", method.Outputs)
	}
}

func TestCurrentStorageLayoutsAreGenerated(t *testing.T) {
	wantRollupSlots := map[string]uint{
		"submitterContract":        151,
		"batchDataStore":           162,
		"batchBlobVersionedHashes": 173,
	}
	for _, entry := range RollupStorageLayout.Storage {
		if entry.Label == "pendingBatchCount" {
			t.Fatal("Rollup storage layout must not contain pendingBatchCount")
		}
		if want, ok := wantRollupSlots[entry.Label]; ok {
			if entry.Slot != want {
				t.Fatalf("Rollup %s slot = %d, want %d", entry.Label, entry.Slot, want)
			}
			delete(wantRollupSlots, entry.Label)
		}
	}
	if len(wantRollupSlots) != 0 {
		t.Fatalf("Rollup storage layout is missing entries: %v", wantRollupSlots)
	}

	wantSubmitterSlots := map[string]uint{
		"withdrawing":          158,
		"withdrawalBatchIndex": 159,
	}
	foundFinalGap := false
	for _, entry := range SubmitterStorageLayout.Storage {
		if want, ok := wantSubmitterSlots[entry.Label]; ok {
			if entry.Slot != want {
				t.Fatalf("Submitter %s slot = %d, want %d", entry.Label, entry.Slot, want)
			}
			delete(wantSubmitterSlots, entry.Label)
		}
		if entry.Label == "__gap" && entry.Slot == 160 {
			layoutType := SubmitterStorageLayout.Types[entry.Type]
			if layoutType.Label != "uint256[49]" {
				t.Fatalf("Submitter final gap type = %s, want uint256[49]", layoutType.Label)
			}
			foundFinalGap = true
		}
	}
	if len(wantSubmitterSlots) != 0 {
		t.Fatalf("Submitter storage layout is missing entries: %v", wantSubmitterSlots)
	}
	if !foundFinalGap {
		t.Fatal("Submitter storage layout is missing the final gap at slot 160")
	}
}
