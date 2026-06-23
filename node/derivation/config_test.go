package derivation

import (
	"strings"
	"testing"
)

// SPEC-005 section 4.2 + 5.1 verify-mode dispatch tests. The mode is bound at
// startup; the validation switch in SetCliContext rejects unknown values
// fail-fast so a typo never reaches the main loop.

func TestVerifyMode_DefaultIsLocal(t *testing.T) {
	if got := DefaultConfig().VerifyMode; got != VerifyModeLocal {
		t.Fatalf("DefaultConfig().VerifyMode = %q, want %q", got, VerifyModeLocal)
	}

	got, err := validateAndDefaultVerifyMode("")
	if err != nil {
		t.Fatalf("empty verify-mode rejected: %v", err)
	}
	if got != VerifyModeLocal {
		t.Fatalf("empty verify-mode normalised to %q, want %q", got, VerifyModeLocal)
	}
}

func TestVerifyMode_AcceptsExplicitModes(t *testing.T) {
	for _, mode := range []string{VerifyModeLayer1, VerifyModeLocal} {
		got, err := validateAndDefaultVerifyMode(mode)
		if err != nil {
			t.Fatalf("%s rejected: %v", mode, err)
		}
		if got != mode {
			t.Fatalf("%s normalised to %q, want %q", mode, got, mode)
		}
	}
}

func TestVerifyMode_LegacyValidatorAlias(t *testing.T) {
	got, err := resolveVerifyMode(DefaultVerifyMode, false, "", true)
	if err != nil {
		t.Fatalf("legacy validator alias rejected: %v", err)
	}
	if got != VerifyModeLayer1 {
		t.Fatalf("legacy validator alias resolved to %q, want %q", got, VerifyModeLayer1)
	}
}

func TestVerifyMode_ExplicitModeOverridesLegacyValidatorAlias(t *testing.T) {
	got, err := resolveVerifyMode(DefaultVerifyMode, true, VerifyModeLocal, true)
	if err != nil {
		t.Fatalf("explicit verify-mode rejected: %v", err)
	}
	if got != VerifyModeLocal {
		t.Fatalf("explicit verify-mode resolved to %q, want %q", got, VerifyModeLocal)
	}
}

func TestVerifyMode_RejectsUnknown(t *testing.T) {
	// "hybrid" was the old default; ensure post-removal it's rejected so
	// stale operator configs fail loud rather than silently falling back to
	// local.
	for _, bad := range []string{"pathC", "hybrid"} {
		err := validateAndDefaultVerifyModeErr(t, bad)
		if !strings.Contains(err.Error(), bad) {
			t.Fatalf("error should mention the offending value %q; got: %v", bad, err)
		}
		// Error message should enumerate the valid modes so a typo's fix
		// is obvious from the log line alone.
		for _, mode := range []string{VerifyModeLayer1, VerifyModeLocal} {
			if !strings.Contains(err.Error(), mode) {
				t.Fatalf("error should list %q as a valid mode; got: %v", mode, err)
			}
		}
	}

	if _, err := validateAndDefaultVerifyMode("PATHA"); err == nil {
		t.Fatal("verify-mode is case-sensitive; uppercase should be rejected")
	}
}

func validateAndDefaultVerifyModeErr(t *testing.T, s string) error {
	t.Helper()
	_, err := validateAndDefaultVerifyMode(s)
	if err == nil {
		t.Fatalf("expected error on verify-mode %q, got nil", s)
	}
	return err
}
