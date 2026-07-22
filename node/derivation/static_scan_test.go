package derivation

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// SPEC-005 section 5.1 static-assertion tests. These guard against regressions where
// someone accidentally re-introduces validator/blocktag references or pulls
// the wrong common package after a refactor.

// walkNodeRepoSourceFiles walks up from this test file to the morph repo
// root (parent of node/) and yields every .go source file under node/
// (excluding test files and vendored code).
func walkNodeRepoSourceFiles(t *testing.T) (string, []string) {
	t.Helper()

	wd, err := os.Getwd() // .../morph/node/derivation
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	nodeRoot := filepath.Dir(wd) // .../morph/node

	var files []string
	err = filepath.WalkDir(nodeRoot, func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			// Skip vendored / test-fixtures dirs if any; nothing matches today
			// but cheap to keep the door closed.
			name := d.Name()
			if name == "node_modules" || name == "vendor" || name == "ops-morph" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatalf("walk node tree: %v", err)
	}
	return nodeRoot, files
}

func TestNoValidatorReferences(t *testing.T) {
	_, files := walkNodeRepoSourceFiles(t)

	// Symbols that the SPEC-005 validator-role removal must keep out of node/.
	// We are specifically guarding against accidental re-introduction; the
	// patterns are narrow on purpose so legitimate uses (e.g., Tendermint
	// consensus validator pubkeys) don't false-positive.
	banned := []string{
		"node/validator",             // import path
		"validator.NewValidator",     // factory call
		"validator.NewConfig",        // config call
		"flags.ValidatorEnable",      // role flag
		"validator.challengeEnable",  // legacy flag string
		"validator.privateKey",       // legacy flag string
		"VALIDATOR_PRIVATE_KEY",      // legacy envvar
		"VALIDATOR_CHALLENGE_ENABLE", // legacy envvar
		// We deliberately do NOT ban "ChallengeEnable" / "ChallengeState"
		// in source -- they appear in the Rollup contract ABI string in
		// node/types/batch.go and are immutable on-chain identifiers we
		// must keep in sync with. The node-side challenge bypass that
		// SPEC-005 removes is keyed by validator.* flags above, which
		// uniquely identify the deleted code paths.
	}

	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		body := string(b)
		for _, p := range banned {
			if strings.Contains(body, p) {
				t.Errorf("validator residue: %q found in %s", p, f)
			}
		}
	}
}

func TestNoBlocktagReferences(t *testing.T) {
	_, files := walkNodeRepoSourceFiles(t)

	banned := []string{
		"node/blocktag",               // import path
		"BlockTagService",             // service type
		"NewBlockTagService",          // factory
		"BlockTagSafeConfirmations",   // flag symbol
		"BLOCKTAG_SAFE_CONFIRMATIONS", // envvar
		"blocktag.safeConfirmations",  // flag name string
		"blocktag.DefaultConfig",      // config factory
	}

	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		body := string(b)
		for _, p := range banned {
			if strings.Contains(body, p) {
				t.Errorf("blocktag residue: %q found in %s", p, f)
			}
		}
	}
}

// TestLocalVerifyUsesCommonBlobPackage guards SPEC-005 section 3.4: local verify must use
// `common/blob` helpers (the same set tx-submitter calls), not the duplicate
// implementations under `common/batch/blob.go`. Codec drift between the two
// would cause permanent versioned hash mismatches.
func TestLocalVerifyUsesCommonBlobPackage(t *testing.T) {
	body, err := os.ReadFile("verify_local.go")
	if err != nil {
		t.Fatalf("read verify_local.go: %v", err)
	}
	src := string(body)

	if !strings.Contains(src, `"morph-l2/common/blob"`) {
		t.Fatalf("verify_local.go must import morph-l2/common/blob")
	}
	// Sanity check the actual call sites -- import is necessary but not
	// sufficient; mismatched calls (e.g., commonbatch.CompressBatchBytes)
	// would still drift codecs.
	required := []string{"commonblob.CompressBatchBytes", "commonblob.MakeBlobTxSidecar"}
	for _, sym := range required {
		if !strings.Contains(src, sym) {
			t.Errorf("verify_local.go missing required call %q", sym)
		}
	}
}
