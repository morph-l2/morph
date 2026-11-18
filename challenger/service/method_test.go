package service

import (
	"bytes"
	"testing"

	"morph-l2/bindings/bindings"
)

func TestChallengeMethodSig(t *testing.T) {
	sig, err := challengeMethodSig()
	if err != nil {
		t.Fatal(err)
	}
	if sig != "challengeState(uint64,bytes32)" {
		t.Fatalf("sig = %s", sig)
	}
}

func TestProbeCallData(t *testing.T) {
	data, err := probeCallData()
	if err != nil {
		t.Fatalf("probeCallData: %v", err)
	}
	parsed, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	// Selector plus two 32-byte words, the same shape as a real challengeState call.
	if len(data) != 4+64 {
		t.Fatalf("len = %d, want 68", len(data))
	}
	if !bytes.Equal(data[:4], parsed.Methods["challengeState"].ID) {
		t.Fatalf("selector = %x, want %x", data[:4], parsed.Methods["challengeState"].ID)
	}
}
