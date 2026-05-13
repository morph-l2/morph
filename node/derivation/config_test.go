package derivation

import (
	"strings"
	"testing"
)

// SPEC-005 section 4.2 + 5.1 verify-mode dispatch tests. The mode is bound at
// startup; the validation switch in SetCliContext rejects unknown values
// fail-fast so a typo never reaches the main loop.

func TestVerifyMode_DefaultIsPathA(t *testing.T) {
	if got := DefaultConfig().VerifyMode; got != VerifyModePathA {
		t.Fatalf("DefaultConfig().VerifyMode = %q, want %q", got, VerifyModePathA)
	}

	got, err := validateAndDefaultVerifyMode("")
	if err != nil {
		t.Fatalf("empty verify-mode rejected: %v", err)
	}
	if got != VerifyModePathA {
		t.Fatalf("empty verify-mode normalised to %q, want %q", got, VerifyModePathA)
	}
}

func TestVerifyMode_AcceptsPathB(t *testing.T) {
	got, err := validateAndDefaultVerifyMode(VerifyModePathB)
	if err != nil {
		t.Fatalf("pathB rejected: %v", err)
	}
	if got != VerifyModePathB {
		t.Fatalf("pathB normalised to %q, want %q", got, VerifyModePathB)
	}
}

func TestVerifyMode_RejectsUnknown(t *testing.T) {
	if _, err := validateAndDefaultVerifyMode("pathC"); err == nil {
		t.Fatal("expected error on unknown verify-mode, got nil")
	} else if !strings.Contains(err.Error(), "pathC") {
		t.Fatalf("error should mention the offending value; got: %v", err)
	}

	if _, err := validateAndDefaultVerifyMode("PATHA"); err == nil {
		t.Fatal("verify-mode is case-sensitive; uppercase should be rejected")
	}
}
