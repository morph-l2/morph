package main

import (
	"encoding/json"
	"flag"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
)

func TestSenderKeysPersistBeforeFunding(t *testing.T) {
	path := filepath.Join(t.TempDir(), "senders.json")
	created, err := loadSenders(path, 2)
	if err != nil {
		t.Fatal(err)
	}
	reloaded, err := loadSenders(path, 2)
	if err != nil {
		t.Fatal(err)
	}
	for i := range created {
		if crypto.PubkeyToAddress(created[i].PublicKey) != crypto.PubkeyToAddress(reloaded[i].PublicKey) {
			t.Fatal("sender changed on restart")
		}
	}
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("sender file must use mode 0600, err=%v", err)
	}
	if _, err := loadSenders(path, 3); err == nil {
		t.Fatal("accepted mismatched sender count")
	}
}

func TestInvalidSenderFilesArePreserved(t *testing.T) {
	for _, data := range []string{"incomplete", `["bad key"]`} {
		path := filepath.Join(t.TempDir(), "senders.json")
		if err := os.WriteFile(path, []byte(data), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := loadSenders(path, 1); err == nil {
			t.Fatal("accepted invalid file")
		}
		after, err := os.ReadFile(path)
		if err != nil || string(after) != data {
			t.Fatal("invalid file was overwritten")
		}
	}
}

func TestSenderFilesMustBePrivateRegularFiles(t *testing.T) {
	path := filepath.Join(t.TempDir(), "senders.json")
	if _, err := loadSenders(path, 1); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadSenders(path, 1); err == nil {
		t.Fatal("accepted publicly readable sender keys")
	}
	if err := os.Chmod(path, 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(t.TempDir(), "senders-link.json")
	if err := os.Symlink(path, link); err != nil {
		t.Fatal(err)
	}
	if _, err := loadSenders(link, 1); err == nil {
		t.Fatal("accepted a symbolic link")
	}
}

func TestDifferentGenesisRejectedBeforeFunding(t *testing.T) {
	newRPC := func(extra byte) *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var request struct {
				ID     json.RawMessage `json:"id"`
				Method string          `json:"method"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Error(err)
				return
			}
			var result interface{}
			switch request.Method {
			case "eth_chainId":
				result = "0x1"
			case "eth_getBlockByNumber":
				result = &types.Header{Number: big.NewInt(0), Difficulty: big.NewInt(0), Extra: []byte{extra}}
			default:
				t.Errorf("unexpected RPC call before genesis validation: %s", request.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": request.ID, "result": result})
		}))
	}
	first, second := newRPC(1), newRPC(2)
	defer first.Close()
	defer second.Close()
	fundingKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("FUNDING_PRIVATE_KEY", common.Bytes2Hex(crypto.FromECDSA(fundingKey)))
	oldFlags, oldArgs := flag.CommandLine, os.Args
	t.Cleanup(func() { flag.CommandLine, os.Args = oldFlags, oldArgs })
	flag.CommandLine = flag.NewFlagSet("multisend", flag.ContinueOnError)
	accounts := filepath.Join(t.TempDir(), "senders.json")
	os.Args = []string{"multisend", "-rpc", first.URL + "," + second.URL, "-chain-id", "1", "-accounts", accounts, "-timeout", "1s"}
	if err := run(); err == nil || !strings.Contains(err.Error(), "different genesis") {
		t.Fatalf("expected genesis mismatch, got %v", err)
	}
	if _, err := os.Stat(accounts); !os.IsNotExist(err) {
		t.Fatal("sender accounts created before validating all RPC identities")
	}
}
