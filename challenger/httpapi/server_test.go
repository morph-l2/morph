package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/challenger/service"
)

type fakeChallenger struct {
	result service.ChallengeResult
	err    error
}

func (f fakeChallenger) Challenge(context.Context, uint64, common.Hash) (service.ChallengeResult, error) {
	return f.result, f.err
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

func TestChallengeConflict(t *testing.T) {
	srv := NewServer(":0", "", fakeChallenger{err: &service.RejectError{Reason: "already in challenge"}}, time.Second)
	hash := common.HexToHash("0x33")
	rec := postChallenge(t, srv.Handler(), `{"batchIndex":4,"batchHash":"`+hash.Hex()+`"}`, "")
	if rec.Code != http.StatusConflict {
		t.Fatalf("status = %d", rec.Code)
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
