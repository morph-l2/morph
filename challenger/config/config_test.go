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

func TestLoadConfigLocalKey(t *testing.T) {
	cfg, err := load(
		"--l1-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
		"--private-key", "abc",
		"--http-auth-token", "secret",
	)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.L1RPC != "http://127.0.0.1:8545" {
		t.Fatalf("l1 rpc = %s", cfg.L1RPC)
	}
	if cfg.HTTPAddress() != "0.0.0.0:8080" {
		t.Fatalf("http address = %s", cfg.HTTPAddress())
	}
	if cfg.HTTPAuthToken != "secret" {
		t.Fatalf("auth token = %s", cfg.HTTPAuthToken)
	}
	if cfg.ExternalSign {
		t.Fatal("external sign should be off")
	}
}

func TestLoadConfigRequiresSignerMaterial(t *testing.T) {
	_, err := load(
		"--l1-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
	)
	if err == nil || !strings.Contains(err.Error(), "private key is required") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadConfigExternalSign(t *testing.T) {
	cfg, err := load(
		"--l1-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
		"--external-sign",
		"--external-sign-address", "0x0000000000000000000000000000000000000002",
		"--external-sign-appid", "app",
		"--external-sign-chain", "eth",
		"--external-sign-url", "https://signer.example",
		"--external-sign-rsa-priv", "pem",
	)
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
	_, err := load(
		"--l1-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
		"--external-sign",
		"--external-sign-address", "0x0000000000000000000000000000000000000002",
	)
	if err == nil || !strings.Contains(err.Error(), "external sign is enabled") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadConfigRejectsFeeCapBelowTip(t *testing.T) {
	_, err := load(
		"--l1-eth-rpc", "http://127.0.0.1:8545",
		"--rollup-address", "0x0000000000000000000000000000000000000001",
		"--private-key", "abc",
		"--gas-fee-cap", "1",
		"--gas-tip-cap", "2",
	)
	if err == nil || !strings.Contains(err.Error(), "gas-fee-cap") {
		t.Fatalf("err = %v", err)
	}
}
