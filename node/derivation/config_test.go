package derivation

import (
	"flag"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
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

func TestMetricsPort_AcceptsDefaultInLayer1(t *testing.T) {
	cfg := DefaultConfig()
	if err := cfg.SetCliContext(newVerifyModeTestContext(t, map[string]string{
		flags.DerivationVerifyMode.Name: VerifyModeLayer1,
	})); err != nil {
		t.Fatalf("layer1 with default metrics-port rejected: %v", err)
	}
	if cfg.MetricsPort != 26660 {
		t.Fatalf("default metrics-port = %d, want 26660", cfg.MetricsPort)
	}
}

func TestMetricsPort_RejectsOutOfRange(t *testing.T) {
	cfg := DefaultConfig()
	err := cfg.SetCliContext(newVerifyModeTestContext(t, map[string]string{
		flags.DerivationVerifyMode.Name: VerifyModeLayer1,
		flags.MetricsPort.Name:          "70000",
	}))
	if err == nil {
		t.Fatal("out-of-range metrics-port accepted, want error")
	}
	if !strings.Contains(err.Error(), flags.MetricsPort.Name) {
		t.Fatalf("error should mention %q; got: %v", flags.MetricsPort.Name, err)
	}
}

func TestBeaconRpcList(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want []string
	}{
		{"single", "http://a", []string{"http://a"}},
		{"multiple", "http://a,http://b,http://c", []string{"http://a", "http://b", "http://c"}},
		{"trims spaces", " http://a , http://b ", []string{"http://a", "http://b"}},
		{"drops empties", "http://a,,http://b,", []string{"http://a", "http://b"}},
		{"empty", "", nil},
		{"only separators", " , , ", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := (&Config{BeaconRpc: tc.in}).BeaconRpcList()
			require.Equal(t, tc.want, got)
		})
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

func newVerifyModeTestContext(t *testing.T, values map[string]string) *cli.Context {
	t.Helper()

	flagSet := flag.NewFlagSet(t.Name(), flag.ContinueOnError)
	for _, f := range []cli.Flag{
		flags.DerivationVerifyMode,
		flags.MetricsPort,
		flags.L1BeaconAddr,
		flags.L2EngineJWTSecret,
	} {
		f.Apply(flagSet)
	}

	defaults := map[string]string{
		flags.L1BeaconAddr.Name:      "http://beacon.example",
		flags.L2EngineJWTSecret.Name: filepath.Join(t.TempDir(), "jwt.hex"),
	}
	for name, value := range defaults {
		if err := flagSet.Set(name, value); err != nil {
			t.Fatalf("set default flag %s: %v", name, err)
		}
	}
	for name, value := range values {
		if err := flagSet.Set(name, value); err != nil {
			t.Fatalf("set flag %s: %v", name, err)
		}
	}

	return cli.NewContext(cli.NewApp(), flagSet, nil)
}
