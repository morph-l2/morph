package config

import (
	"strings"
	"testing"

	"github.com/urfave/cli"

	"morph-l2/challenger/flags"
)

func load(args ...string) (*Config, error) {
	var cfg *Config
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Action = func(ctx *cli.Context) error {
		loaded, err := LoadConfig(ctx)
		cfg = loaded
		return err
	}
	err := app.Run(append([]string{"challenger"}, args...))
	return cfg, err
}

// base returns the flags every configuration needs, followed by extra.
func base(extra ...string) []string {
	return append([]string{
		"--l1-eth-rpc", "http://127.0.0.1:9545",
		"--l2-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
	}, extra...)
}

func without(args []string, flag string) []string {
	out := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		if args[i] == flag {
			i++
			continue
		}
		out = append(out, args[i])
	}
	return out
}

func TestLoadConfigLocalKey(t *testing.T) {
	cfg, err := load(base("--private-key", "abc")...)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.L1RPC != "http://127.0.0.1:9545" || cfg.L2RPC != "http://127.0.0.1:8545" {
		t.Fatalf("rpcs = %s %s", cfg.L1RPC, cfg.L2RPC)
	}
	if cfg.HTTPAddress() != "0.0.0.0:8080" {
		t.Fatalf("http address = %s", cfg.HTTPAddress())
	}
	if cfg.HTTPAuthEnable || cfg.HTTPAuthToken != "" {
		t.Fatalf("auth should be off by default: enable=%t token=%q", cfg.HTTPAuthEnable, cfg.HTTPAuthToken)
	}
	if cfg.ExternalSign {
		t.Fatal("external sign should be off")
	}
}

func TestLoadConfigAuthEnabled(t *testing.T) {
	cfg, err := load(base("--private-key", "abc", "--http-auth-enable", "--http-auth-token", "secret")...)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.HTTPAuthEnable || cfg.HTTPAuthToken != "secret" {
		t.Fatalf("enable=%t token=%q", cfg.HTTPAuthEnable, cfg.HTTPAuthToken)
	}
}

func TestLoadConfigAuthEnabledRequiresToken(t *testing.T) {
	for _, extra := range [][]string{
		{"--http-auth-enable"},
		{"--http-auth-enable", "--http-auth-token", "   "},
	} {
		_, err := load(base(append([]string{"--private-key", "abc"}, extra...)...)...)
		if err == nil || !strings.Contains(err.Error(), "http auth token is required") {
			t.Fatalf("%v: err = %v", extra, err)
		}
	}
}

func TestLoadConfigAuthDisabledIgnoresToken(t *testing.T) {
	cfg, err := load(base("--private-key", "abc", "--http-auth-token", "secret")...)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.HTTPAuthToken != "" {
		t.Fatalf("token = %q, want ignored while auth is off", cfg.HTTPAuthToken)
	}
}

func TestLoadConfigRequiresRollupAddress(t *testing.T) {
	_, err := load(without(base("--private-key", "abc"), "--rollup-address")...)
	if err == nil || !strings.Contains(err.Error(), "rollup-address") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadConfigRequiresSignerMaterial(t *testing.T) {
	_, err := load(base()...)
	if err == nil || !strings.Contains(err.Error(), "private key is required") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadConfigExternalSign(t *testing.T) {
	cfg, err := load(base(
		"--external-sign",
		"--external-sign-address", "0x0000000000000000000000000000000000000002",
		"--external-sign-appid", "app",
		"--external-sign-chain", "eth",
		"--external-sign-url", "https://signer.example",
		"--external-sign-rsa-priv", "pem",
	)...)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.ExternalSign {
		t.Fatal("external sign should be on")
	}
	if cfg.ExternalSignAddress != "0x0000000000000000000000000000000000000002" {
		t.Fatalf("signer address = %s", cfg.ExternalSignAddress)
	}
}

func TestLoadConfigExternalSignMissingFields(t *testing.T) {
	_, err := load(base(
		"--external-sign",
		"--external-sign-address", "0x0000000000000000000000000000000000000002",
	)...)
	if err == nil || !strings.Contains(err.Error(), "external sign is enabled") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadConfigRejectsFeeCapBelowTip(t *testing.T) {
	_, err := load(base(
		"--private-key", "abc",
		"--gas-fee-cap", "1",
		"--gas-tip-cap", "2",
	)...)
	if err == nil || !strings.Contains(err.Error(), "gas-fee-cap") {
		t.Fatalf("err = %v", err)
	}
}
