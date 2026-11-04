package service

import (
	"testing"
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
