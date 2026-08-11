package bindings

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi"
)

const (
	oldCommitBatchSelector          = "428868b5"
	oldCommitStateSelector          = "1e8825be"
	oldCommitBatchWithProofSelector = "4e8f1d67"
	newCommitBatchSelector          = "41f756da"
	newCommitStateSelector          = "67caa37a"
	newCommitBatchWithProofSelector = "1544ba3a"
)

type calldataCorpus struct {
	SchemaVersion int `json:"schemaVersion"`
	Fixtures      []struct {
		Name     string `json:"name"`
		Epoch    string `json:"epoch"`
		Selector string `json:"selector"`
		Data     string `json:"data"`
	} `json:"fixtures"`
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test source path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}

func readABI(t *testing.T, path string) abi.ABI {
	t.Helper()
	raw, err := os.Open(path)
	if err != nil {
		t.Fatalf("open ABI %s: %v", path, err)
	}
	defer raw.Close()
	parsed, err := abi.JSON(raw)
	if err != nil {
		t.Fatalf("parse ABI %s: %v", path, err)
	}
	return parsed
}

func methodSelector(t *testing.T, parsed abi.ABI, name string) string {
	t.Helper()
	method, ok := parsed.Methods[name]
	if !ok {
		t.Fatalf("ABI is missing method %s", name)
	}
	return hex.EncodeToString(method.ID)
}

func TestGeneratedRollupBindingShape(t *testing.T) {
	current, err := RollupMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse generated Rollup ABI: %v", err)
	}
	if current == nil {
		t.Fatal("generated Rollup ABI is nil")
	}
	if strings.Contains(RollupMetaData.ABI, "BatchSignatureInput") {
		t.Fatal("generated Rollup binding still contains BatchSignatureInput")
	}

	expected := map[string]struct {
		selector string
		inputs   int
	}{
		"commitBatch":          {newCommitBatchSelector, 1},
		"commitState":          {newCommitStateSelector, 1},
		"commitBatchWithProof": {newCommitBatchWithProofSelector, 3},
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
		t.Fatalf("batchDataStore output shape = %v, want fourth output address", batchDataStore.Outputs)
	}
	if _, ok := current.Methods["pendingBatchCount"]; !ok {
		t.Fatal("generated Rollup ABI is missing pendingBatchCount")
	}
}

func TestGeneratedSubmitterBindingShape(t *testing.T) {
	current, err := SubmitterMetaData.GetAbi()
	if err != nil {
		t.Fatalf("parse generated Submitter ABI: %v", err)
	}
	if current == nil {
		t.Fatal("generated Submitter ABI is nil")
	}
	for _, name := range []string{
		"addSubmitter",
		"claimCredit",
		"claimWithdrawal",
		"initialize",
		"isActive",
		"removeSubmitter",
		"slash",
		"stake",
		"withdraw",
	} {
		if _, ok := current.Methods[name]; !ok {
			t.Fatalf("generated Submitter ABI is missing %s", name)
		}
	}
	active := current.Methods["isActive"]
	if len(active.Inputs) != 1 || len(active.Outputs) != 1 || active.Outputs[0].Type.String() != "bool" {
		t.Fatalf("isActive ABI shape changed: inputs=%v outputs=%v", active.Inputs, active.Outputs)
	}
}

func TestPreSubmitterAndCurrentSelectorsRemainDistinct(t *testing.T) {
	root := repositoryRoot(t)
	preSubmitter := readABI(t, filepath.Join(
		root,
		"contracts",
		"abi",
		"archive",
		"pre-submitter",
		"Rollup.json",
	))
	current, err := RollupMetaData.GetAbi()
	if err != nil || current == nil {
		t.Fatalf("parse generated Rollup ABI: %v", err)
	}

	tests := []struct {
		name          string
		oldSelector   string
		newSelector   string
		oldInputCount int
		newInputCount int
	}{
		{"commitBatch", oldCommitBatchSelector, newCommitBatchSelector, 2, 1},
		{"commitState", oldCommitStateSelector, newCommitStateSelector, 2, 1},
		{"commitBatchWithProof", oldCommitBatchWithProofSelector, newCommitBatchWithProofSelector, 4, 3},
	}
	for _, test := range tests {
		oldMethod := preSubmitter.Methods[test.name]
		newMethod := current.Methods[test.name]
		if got := methodSelector(t, preSubmitter, test.name); got != test.oldSelector {
			t.Fatalf("pre-submitter %s selector = %s, want %s", test.name, got, test.oldSelector)
		}
		if got := methodSelector(t, *current, test.name); got != test.newSelector {
			t.Fatalf("current %s selector = %s, want %s", test.name, got, test.newSelector)
		}
		if bytes.Equal(oldMethod.ID, newMethod.ID) {
			t.Fatalf("pre-submitter and current %s selectors collide", test.name)
		}
		if len(oldMethod.Inputs) != test.oldInputCount || len(newMethod.Inputs) != test.newInputCount {
			t.Fatalf(
				"%s input counts = old %d/current %d, want old %d/current %d",
				test.name,
				len(oldMethod.Inputs),
				len(newMethod.Inputs),
				test.oldInputCount,
				test.newInputCount,
			)
		}
	}
}

func TestSupportedRollupSelectorsHaveNoCollisions(t *testing.T) {
	root := repositoryRoot(t)
	raw, err := os.ReadFile(filepath.Join(
		root,
		"node",
		"derivation",
		"testdata",
		"rollup_calldata_fixtures.json",
	))
	if err != nil {
		t.Fatalf("read shared calldata corpus: %v", err)
	}
	var corpus calldataCorpus
	if err := json.Unmarshal(raw, &corpus); err != nil {
		t.Fatalf("parse shared calldata corpus: %v", err)
	}
	if corpus.SchemaVersion != 1 || len(corpus.Fixtures) != 8 {
		t.Fatalf("shared calldata corpus shape = version %d/%d fixtures, want version 1/8 fixtures", corpus.SchemaVersion, len(corpus.Fixtures))
	}

	seen := make(map[string]string, len(corpus.Fixtures))
	for _, fixture := range corpus.Fixtures {
		selector := strings.TrimPrefix(strings.ToLower(fixture.Selector), "0x")
		data := strings.TrimPrefix(strings.ToLower(fixture.Data), "0x")
		if len(selector) != 8 || len(data) < 8 || data[:8] != selector {
			t.Fatalf("fixture %s has inconsistent selector/data", fixture.Name)
		}
		if previous, ok := seen[selector]; ok {
			t.Fatalf("selector collision %s between %s and %s", selector, previous, fixture.Name)
		}
		seen[selector] = fixture.Name
	}

	for _, selector := range []string{
		oldCommitBatchSelector,
		oldCommitStateSelector,
		oldCommitBatchWithProofSelector,
		newCommitBatchSelector,
		newCommitStateSelector,
		newCommitBatchWithProofSelector,
	} {
		if _, ok := seen[selector]; !ok {
			t.Fatalf("shared calldata corpus is missing protected selector %s", selector)
		}
	}
}
