package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
)

type fakeRemoteSigner struct {
	key  *ecdsa.PrivateKey
	err  error
	wait chan struct{}
	got  *types.Transaction
	sig  string
}

func (f *fakeRemoteSigner) Sign(tx *types.Transaction, methodSig string) (*types.Transaction, error) {
	if f.wait != nil {
		<-f.wait
	}
	f.got, f.sig = tx, methodSig
	if f.err != nil {
		return nil, f.err
	}
	return types.SignTx(tx, types.NewLondonSigner(tx.ChainId()), f.key)
}

func newProbeSigner(t *testing.T, remote *fakeRemoteSigner, address common.Address) *Signer {
	t.Helper()
	return &Signer{remoteClient: remote, externalSignAddress: address, chainID: big.NewInt(1)}
}

func TestSignerProbe(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	from := crypto.PubkeyToAddress(key.PublicKey)
	to := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	data := []byte{1, 2, 3, 4}

	t.Run("signs an unbroadcastable transaction", func(t *testing.T) {
		remote := &fakeRemoteSigner{key: key}
		if err := newProbeSigner(t, remote, from).Probe(context.Background(), to, data, "challengeState(uint64,bytes32)"); err != nil {
			t.Fatalf("Probe: %v", err)
		}
		if remote.sig != "challengeState(uint64,bytes32)" {
			t.Fatalf("method sig = %q", remote.sig)
		}
		tx := remote.got
		if *tx.To() != to || string(tx.Data()) != string(data) {
			t.Fatalf("unexpected probe target: to=%s data=%x", tx.To(), tx.Data())
		}
		if tx.Gas() != 0 || tx.Value().Sign() != 0 {
			t.Fatalf("probe tx must carry no gas and no value: gas=%d value=%s", tx.Gas(), tx.Value())
		}
	})

	t.Run("rejects a different sender", func(t *testing.T) {
		remote := &fakeRemoteSigner{key: key}
		err := newProbeSigner(t, remote, common.HexToAddress("0x01")).Probe(context.Background(), to, data, "m()")
		if err == nil || !strings.Contains(err.Error(), "remote signer signed as") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("returns the signer error", func(t *testing.T) {
		want := errors.New("boom")
		err := newProbeSigner(t, &fakeRemoteSigner{err: want}, from).Probe(context.Background(), to, data, "m()")
		if !errors.Is(err, want) {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("stops waiting when the context ends", func(t *testing.T) {
		remote := &fakeRemoteSigner{key: key, wait: make(chan struct{})}
		defer close(remote.wait)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := newProbeSigner(t, remote, from).Probe(ctx, to, data, "m()")
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("err = %v", err)
		}
	})
}
