package rpc

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// okHandler is a stub downstream handler that always returns 200.
var okHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":true}`))
})

func TestAuthMiddleware_ReadMethod_NoToken_Passes(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `{"jsonrpc":"2.0","method":"ha_leader","params":[],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestAuthMiddleware_WriteMethod_ValidToken_Passes(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `{"jsonrpc":"2.0","method":"ha_removeServer","params":["node-2",1],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "secret")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestAuthMiddleware_WriteMethod_NoToken_Returns401(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `{"jsonrpc":"2.0","method":"ha_removeServer","params":["node-2",1],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestAuthMiddleware_WriteMethod_WrongToken_Returns401(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `{"jsonrpc":"2.0","method":"ha_addServerAsVoter","params":["id","addr",0],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "wrong-token")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestAuthMiddleware_EmptyToken_AllMethodsPass(t *testing.T) {
	h := authMiddleware("", okHandler)
	body := `{"jsonrpc":"2.0","method":"ha_removeServer","params":["node-2",1],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 (auth disabled), got %d", rr.Code)
	}
}

func TestAuthMiddleware_BatchRequest_WithWriteMethod_NoToken_Returns401(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `[{"jsonrpc":"2.0","method":"ha_leader","params":[],"id":1},{"jsonrpc":"2.0","method":"ha_removeServer","params":["node-2",1],"id":2}]`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for batch with write method, got %d", rr.Code)
	}
}

func TestAuthMiddleware_BatchRequest_OnlyReadMethods_Passes(t *testing.T) {
	h := authMiddleware("secret", okHandler)
	body := `[{"jsonrpc":"2.0","method":"ha_leader","params":[],"id":1},{"jsonrpc":"2.0","method":"ha_clusterMembership","params":[],"id":2}]`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 for batch with only read methods, got %d", rr.Code)
	}
}

func TestAuthMiddleware_BodyReadable(t *testing.T) {
	var captured string
	downstream := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		captured = string(b)
		w.WriteHeader(http.StatusOK)
	})
	h := authMiddleware("secret", downstream)
	body := `{"jsonrpc":"2.0","method":"ha_leader","params":[],"id":1}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	req.Header.Set("Authorization", "secret")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if captured != body {
		t.Fatalf("body not restored: got %q", captured)
	}
}
