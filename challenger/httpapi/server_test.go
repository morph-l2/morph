package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/challenger/service"
)

type fakeChallenger struct {
	result     service.ChallengeResult
	err        error
	comparison service.Comparison
	compareErr error
}

func (f fakeChallenger) Challenge(context.Context, uint64, common.Hash) (service.ChallengeResult, error) {
	return f.result, f.err
}

func (f fakeChallenger) CompareBatch(context.Context, uint64) (service.Comparison, error) {
	return f.comparison, f.compareErr
}

func postChallenge(t *testing.T, handler http.Handler, body, token string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/challenge", strings.NewReader(body))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func TestChallengeUnauthorized(t *testing.T) {
	srv := NewServer(":0", "secret", fakeChallenger{}, time.Second)
	rec := postChallenge(t, srv.Handler(), `{"batchIndex":1,"batchHash":"0x11"}`, "")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestChallengeRejectsBadHash(t *testing.T) {
	srv := NewServer(":0", "", fakeChallenger{}, time.Second)
	rec := postChallenge(t, srv.Handler(), `{"batchIndex":4,"batchHash":"0x11"}`, "")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d body %s", rec.Code, rec.Body.String())
	}
}

func TestChallengeSuccess(t *testing.T) {
	hash := common.HexToHash("0x22")
	srv := NewServer(":0", "secret", fakeChallenger{
		result: service.ChallengeResult{TxHash: hash, BlockNumber: 9},
	}, time.Second)
	body := `{"batchIndex":4,"batchHash":"` + hash.Hex() + `"}`
	rec := postChallenge(t, srv.Handler(), body, "secret")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d body %s", rec.Code, rec.Body.String())
	}
	var resp challengeResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Status != "success" || resp.TxHash != hash.Hex() || resp.BlockNumber != 9 {
		t.Fatalf("resp = %+v", resp)
	}
}

func TestChallengeErrorStatus(t *testing.T) {
	hash := common.HexToHash("0x33")
	sent := common.HexToHash("0x44").Hex()
	cases := []struct {
		name   string
		err    error
		code   int
		status string
		txHash string
	}{
		{"rejected", &service.RejectError{Reason: "already in challenge"}, http.StatusConflict, "rejected", ""},
		{"busy", &service.BusyError{Confirmed: 3, Pending: 4}, http.StatusConflict, "busy", ""},
		{"pending", &service.TxError{TxHash: sent, Pending: true, Err: errors.New("timeout")}, http.StatusAccepted, "pending", sent},
		{"reverted", &service.TxError{TxHash: sent, Reverted: true, Err: errors.New("reverted")}, http.StatusBadGateway, "reverted", sent},
		{"rpc", errors.New("dial tcp: refused"), http.StatusInternalServerError, "error", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(":0", "", fakeChallenger{err: tc.err}, time.Second)
			rec := postChallenge(t, srv.Handler(), `{"batchIndex":4,"batchHash":"`+hash.Hex()+`"}`, "")
			if rec.Code != tc.code {
				t.Fatalf("status = %d body %s", rec.Code, rec.Body.String())
			}
			var resp challengeResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatal(err)
			}
			if resp.Status != tc.status || resp.TxHash != tc.txHash {
				t.Fatalf("resp = %+v", resp)
			}
		})
	}
}

func TestGetBatch(t *testing.T) {
	state := common.HexToHash("0x11")
	srv := NewServer(":0", "", fakeChallenger{
		comparison: service.Comparison{
			BatchIndex:  4,
			BlockNumber: 20,
			Verdict:     service.VerdictMatch,
			StateRoot:   service.RootPair{L1: state, L2: state},
		},
	}, time.Second)
	req := httptest.NewRequest(http.MethodGet, "/batch/4", nil)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d body %s", rec.Code, rec.Body.String())
	}
	var got service.Comparison
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Verdict != service.VerdictMatch || got.BlockNumber != 20 {
		t.Fatalf("got %+v", got)
	}
}

func TestGetBatchUnauthorized(t *testing.T) {
	srv := NewServer(":0", "secret", fakeChallenger{}, time.Second)
	req := httptest.NewRequest(http.MethodGet, "/batch/4", nil)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestGetBatchNotCommitted(t *testing.T) {
	srv := NewServer(":0", "", fakeChallenger{compareErr: &service.NotCommittedError{BatchIndex: 4}}, time.Second)
	req := httptest.NewRequest(http.MethodGet, "/batch/4", nil)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d body %s", rec.Code, rec.Body.String())
	}
}

func TestHealth(t *testing.T) {
	srv := NewServer(":0", "", fakeChallenger{}, time.Second)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
}
