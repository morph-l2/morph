package service

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	remotesigner "github.com/morph-l2/remote-signer-client/go/signer"

	"morph-l2/bindings/bindings"
	"morph-l2/challenger/client"
	"morph-l2/challenger/config"
)

var (
	testChainID = big.NewInt(900)
	testRollup  = common.HexToAddress("0xa513e6e4b8f2a923d98304ec87f64353c4d5c853")
	testTip     = big.NewInt(1_000_000_000)
	testBaseFee = big.NewInt(10_000_000_000)
)

const testGas = 100_000

// fakeL1 answers the JSON-RPC calls made while building, sending, and
// confirming one transaction. Nothing is executed; sent transactions are kept.
type fakeL1 struct {
	mu        sync.Mutex
	confirmed uint64
	pending   uint64
	status    *uint64
	sent      []*types.Transaction
}

type rpcRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func (f *fakeL1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req rpcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := f.call(req)
	resp := map[string]any{"jsonrpc": "2.0", "id": req.ID}
	if err != nil {
		resp["error"] = map[string]any{"code": -32000, "message": err.Error()}
	} else {
		resp["result"] = result
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (f *fakeL1) call(req rpcRequest) (any, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	switch req.Method {
	case "eth_chainId":
		return (*hexutil.Big)(testChainID), nil
	case "eth_getTransactionCount":
		var tag string
		if err := json.Unmarshal(req.Params[1], &tag); err != nil {
			return nil, err
		}
		if tag == "pending" {
			return hexutil.Uint64(f.pending), nil
		}
		return hexutil.Uint64(f.confirmed), nil
	case "eth_maxPriorityFeePerGas":
		return (*hexutil.Big)(testTip), nil
	case "eth_getBlockByNumber":
		return &types.Header{Number: big.NewInt(100), Difficulty: big.NewInt(0), GasLimit: 30_000_000, BaseFee: testBaseFee}, nil
	case "eth_getCode":
		return hexutil.Bytes{0x60, 0x00}, nil
	case "eth_estimateGas":
		return hexutil.Uint64(testGas), nil
	case "eth_sendRawTransaction":
		var raw hexutil.Bytes
		if err := json.Unmarshal(req.Params[0], &raw); err != nil {
			return nil, err
		}
		tx := new(types.Transaction)
		if err := tx.UnmarshalBinary(raw); err != nil {
			return nil, err
		}
		f.sent = append(f.sent, tx)
		f.pending++
		return tx.Hash(), nil
	case "eth_getTransactionReceipt":
		if f.status == nil || len(f.sent) == 0 {
			return nil, nil
		}
		tx := f.sent[len(f.sent)-1]
		return &types.Receipt{
			Status:      *f.status,
			TxHash:      tx.Hash(),
			BlockNumber: big.NewInt(101),
			GasUsed:     testGas,
			Logs:        []*types.Log{},
		}, nil
	}
	return nil, errors.New("unexpected method " + req.Method)
}

func (f *fakeL1) mined(status uint64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.status = &status
}

func (f *fakeL1) sentTxs() []*types.Transaction {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]*types.Transaction(nil), f.sent...)
}

// fakeSigner plays the remote signer: it signs the raw transaction it is sent
// and records the method signature of the request.
type fakeSigner struct {
	key       *ecdsa.PrivateKey
	mu        sync.Mutex
	methodSig string
}

func (s *fakeSigner) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req remotesigner.ReqData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data remotesigner.DataV2
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	raw, err := hex.DecodeString(data.RawData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx := new(types.Transaction)
	if err := tx.UnmarshalBinary(raw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sig, err := crypto.Sign(types.NewLondonSigner(testChainID).Hash(tx).Bytes(), s.key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.mu.Lock()
	s.methodSig = data.MethodSig
	s.mu.Unlock()
	_ = json.NewEncoder(w).Encode(remotesigner.Response{
		Result: remotesigner.Result{SignDatas: []*remotesigner.SignData{{Sign: hexutil.Encode(sig)}}},
	})
}

func newTestKey(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	return key, crypto.PubkeyToAddress(key.PublicKey)
}

func localL1(t *testing.T, url string, key *ecdsa.PrivateKey) *client.L1Client {
	t.Helper()
	l1, err := client.NewL1Client(&config.Config{
		L1RPC:      url,
		PrivateKey: hexutil.Encode(crypto.FromECDSA(key)),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(l1.Close)
	return l1
}

func remoteL1(t *testing.T, url, signerURL string, from common.Address) *client.L1Client {
	t.Helper()
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	l1, err := client.NewL1Client(&config.Config{
		L1RPC:               url,
		ExternalSign:        true,
		ExternalSignAddress: from.Hex(),
		ExternalSignAppid:   "challenger-test",
		ExternalSignChain:   "ETH",
		ExternalSignUrl:     signerURL,
		ExternalSignRsaPriv: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(rsaKey)),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(l1.Close)
	return l1
}

func mismatchService(t *testing.T, l1 *client.L1Client, hash common.Hash, receiptTimeout time.Duration) *Service {
	t.Helper()
	rollup, err := bindings.NewRollup(testRollup, l1.GetClient())
	if err != nil {
		t.Fatal(err)
	}
	methodSig, err := challengeMethodSig()
	if err != nil {
		t.Fatal(err)
	}
	return &Service{
		l1:        l1,
		rollup:    rollup,
		view:      readyView(hash),
		compare:   &staticCompare{cmp: Comparison{BatchHash: hash, Verdict: VerdictMismatch}},
		txs:       NewTxManager(l1, receiptTimeout),
		methodSig: methodSig,
		from:      l1.WalletAddress(),
	}
}

// checkChallengeTx asserts the broadcast transaction is challengeState(4, hash)
// from the expected sender, carrying the deposit, with capped fees and padded gas.
func checkChallengeTx(t *testing.T, tx *types.Transaction, from common.Address, hash common.Hash) {
	t.Helper()
	sender, err := types.Sender(types.NewLondonSigner(testChainID), tx)
	if err != nil {
		t.Fatal(err)
	}
	if sender != from {
		t.Fatalf("sender = %s, want %s", sender, from)
	}
	if tx.To() == nil || *tx.To() != testRollup {
		t.Fatalf("to = %v", tx.To())
	}
	if tx.Value().Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("value = %s, want deposit 100", tx.Value())
	}
	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	want, err := abi.Pack("challengeState", uint64(4), [32]byte(hash))
	if err != nil {
		t.Fatal(err)
	}
	if hexutil.Encode(tx.Data()) != hexutil.Encode(want) {
		t.Fatalf("calldata = %x", tx.Data())
	}
	if tx.Gas() != testGas*3/2 {
		t.Fatalf("gas = %d", tx.Gas())
	}
	feeCap := new(big.Int).Add(testTip, new(big.Int).Mul(testBaseFee, big.NewInt(2)))
	if tx.GasTipCap().Cmp(testTip) != 0 || tx.GasFeeCap().Cmp(feeCap) != 0 {
		t.Fatalf("tip = %s fee = %s", tx.GasTipCap(), tx.GasFeeCap())
	}
}

func TestChallengeSendsWhenRootsMismatch(t *testing.T) {
	chain := &fakeL1{}
	chain.mined(types.ReceiptStatusSuccessful)
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, from := newTestKey(t)
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, localL1(t, rpc.URL, key), hash, 5*time.Second)
	result, err := svc.Challenge(context.Background(), 4, hash)
	if err != nil {
		t.Fatal(err)
	}
	sent := chain.sentTxs()
	if len(sent) != 1 {
		t.Fatalf("sent %d transactions", len(sent))
	}
	checkChallengeTx(t, sent[0], from, hash)
	if result.TxHash != sent[0].Hash() || result.BlockNumber != 101 {
		t.Fatalf("result = %+v", result)
	}
}

func TestChallengeSendsWithRemoteSigner(t *testing.T) {
	chain := &fakeL1{}
	chain.mined(types.ReceiptStatusSuccessful)
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, from := newTestKey(t)
	signer := &fakeSigner{key: key}
	signerSrv := httptest.NewServer(signer)
	defer signerSrv.Close()
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, remoteL1(t, rpc.URL, signerSrv.URL, from), hash, 5*time.Second)
	result, err := svc.Challenge(context.Background(), 4, hash)
	if err != nil {
		t.Fatal(err)
	}
	sent := chain.sentTxs()
	if len(sent) != 1 {
		t.Fatalf("sent %d transactions", len(sent))
	}
	checkChallengeTx(t, sent[0], from, hash)
	if result.TxHash != sent[0].Hash() {
		t.Fatalf("result = %+v", result)
	}
	if signer.methodSig != svc.methodSig {
		t.Fatalf("signer methodSig = %q, want %q", signer.methodSig, svc.methodSig)
	}
}

func TestChallengeRefusesWrongRemoteSignature(t *testing.T) {
	chain := &fakeL1{}
	chain.mined(types.ReceiptStatusSuccessful)
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	_, from := newTestKey(t)
	otherKey, other := newTestKey(t)
	signerSrv := httptest.NewServer(&fakeSigner{key: otherKey})
	defer signerSrv.Close()
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, remoteL1(t, rpc.URL, signerSrv.URL, from), hash, 5*time.Second)
	_, err := svc.Challenge(context.Background(), 4, hash)
	if err == nil || !strings.Contains(err.Error(), other.Hex()) {
		t.Fatalf("err = %v, want wrong sender %s", err, other.Hex())
	}
	if n := len(chain.sentTxs()); n != 0 {
		t.Fatalf("broadcast %d transactions signed by the wrong key", n)
	}
}

func TestChallengeSkipsWhenSignerBusy(t *testing.T) {
	chain := &fakeL1{confirmed: 3, pending: 4}
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, _ := newTestKey(t)
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, localL1(t, rpc.URL, key), hash, 5*time.Second)
	_, err := svc.Challenge(context.Background(), 4, hash)
	var busy *BusyError
	if !errors.As(err, &busy) || busy.Confirmed != 3 || busy.Pending != 4 {
		t.Fatalf("err = %v, want BusyError", err)
	}
	if n := len(chain.sentTxs()); n != 0 {
		t.Fatalf("sent %d transactions while busy", n)
	}
}

func TestChallengeReportsRevert(t *testing.T) {
	chain := &fakeL1{}
	chain.mined(types.ReceiptStatusFailed)
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, _ := newTestKey(t)
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, localL1(t, rpc.URL, key), hash, 5*time.Second)
	result, err := svc.Challenge(context.Background(), 4, hash)
	var txErr *TxError
	if !errors.As(err, &txErr) || !txErr.Reverted {
		t.Fatalf("err = %v, want reverted TxError", err)
	}
	if sent := chain.sentTxs(); len(sent) != 1 || result.TxHash != sent[0].Hash() {
		t.Fatalf("result = %+v", result)
	}
}

func TestChallengeReportsPendingAfterTimeout(t *testing.T) {
	chain := &fakeL1{}
	rpc := httptest.NewServer(chain)
	defer rpc.Close()
	key, _ := newTestKey(t)
	hash := common.HexToHash("0x11")

	svc := mismatchService(t, localL1(t, rpc.URL, key), hash, 10*time.Millisecond)
	result, err := svc.Challenge(context.Background(), 4, hash)
	var txErr *TxError
	if !errors.As(err, &txErr) || !txErr.Pending {
		t.Fatalf("err = %v, want pending TxError", err)
	}
	if sent := chain.sentTxs(); len(sent) != 1 || result.TxHash != sent[0].Hash() {
		t.Fatalf("result = %+v", result)
	}

	_, err = svc.Challenge(context.Background(), 4, hash)
	var busy *BusyError
	if !errors.As(err, &busy) {
		t.Fatalf("second challenge err = %v, want BusyError", err)
	}
	if n := len(chain.sentTxs()); n != 1 {
		t.Fatalf("sent %d transactions, want 1", n)
	}
}
