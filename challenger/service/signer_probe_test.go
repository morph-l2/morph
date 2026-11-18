package service

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/challenger/client"
)

// The loop probes at start and on every interval, whatever the result, and stops with the context.
func TestRunSignerProbe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	calls := make(chan struct{}, 16)
	n := 0
	probe := func(context.Context) error {
		n++
		calls <- struct{}{}
		if n == 1 {
			return errors.New("down")
		}
		return nil
	}

	done := make(chan struct{})
	go func() {
		runSignerProbe(ctx, probe, time.Millisecond)
		close(done)
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-calls:
		case <-time.After(5 * time.Second):
			t.Fatalf("probe %d did not run", i+1)
		}
	}

	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("runSignerProbe did not stop when the context ended")
	}
}

// The probe has the remote signer sign a challengeState transaction built like
// a real challenge, and never broadcasts it.
func TestProbeSignerSignsChallengeTxWithoutBroadcast(t *testing.T) {
	chain := &fakeL1{confirmed: 7, pending: 7}
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, from := newTestKey(t)
	signer := &fakeSigner{key: key}
	signerSrv := httptest.NewServer(signer)
	defer signerSrv.Close()

	svc := probeService(t, remoteL1(t, rpc.URL, signerSrv.URL, from))
	if err := svc.probeSigner(context.Background()); err != nil {
		t.Fatalf("probeSigner: %v", err)
	}

	if sent := chain.sentTxs(); len(sent) != 0 {
		t.Fatalf("probe broadcast %d transactions", len(sent))
	}
	if len(signer.signed) != 1 {
		t.Fatalf("signer was asked %d times, want 1", len(signer.signed))
	}
	if signer.methodSig != "challengeState(uint64,bytes32)" {
		t.Fatalf("method sig = %q", signer.methodSig)
	}
	tx := signer.signed[0]
	if tx.To() == nil || *tx.To() != testRollup {
		t.Fatalf("to = %v", tx.To())
	}
	if !bytes.Equal(tx.Data(), svc.probeData) {
		t.Fatalf("calldata = %x", tx.Data())
	}
	if tx.Value().Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("value = %s, want deposit 100", tx.Value())
	}
	if tx.Nonce() != 7 {
		t.Fatalf("nonce = %d, want 7", tx.Nonce())
	}
	if tx.Gas() == 0 {
		t.Fatal("probe transaction has no gas limit")
	}
	feeCap := new(big.Int).Add(testTip, new(big.Int).Mul(testBaseFee, big.NewInt(2)))
	if tx.GasTipCap().Cmp(testTip) != 0 || tx.GasFeeCap().Cmp(feeCap) != 0 {
		t.Fatalf("tip = %s fee = %s", tx.GasTipCap(), tx.GasFeeCap())
	}
}

// A signature from another key fails the probe.
func TestProbeSignerRejectsWrongSignature(t *testing.T) {
	chain := &fakeL1{}
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	_, from := newTestKey(t)
	otherKey, _ := newTestKey(t)
	signerSrv := httptest.NewServer(&fakeSigner{key: otherKey})
	defer signerSrv.Close()

	svc := probeService(t, remoteL1(t, rpc.URL, signerSrv.URL, from))
	err := svc.probeSigner(context.Background())
	if err == nil || !strings.Contains(err.Error(), "remote signer signed as") {
		t.Fatalf("err = %v", err)
	}
	if sent := chain.sentTxs(); len(sent) != 0 {
		t.Fatalf("probe broadcast %d transactions", len(sent))
	}
}

func probeService(t *testing.T, l1 *client.L1Client) *Service {
	t.Helper()
	svc := mismatchService(t, l1, common.Hash{}, time.Second)
	data, err := probeCallData()
	if err != nil {
		t.Fatal(err)
	}
	svc.rollupAddr = testRollup
	svc.probeData = data
	return svc
}
