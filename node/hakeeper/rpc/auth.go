package rpc

import (
	"bytes"
	"crypto/subtle"
	"encoding/json"
	"io"
	"net/http"
)

// writeRPCMethods is the set of HA JSON-RPC methods that modify cluster state.
// All other methods are read-only and do not require authentication.
var writeRPCMethods = map[string]bool{
	"ha_addServerAsVoter":       true,
	"ha_addServerAsNonvoter":    true,
	"ha_removeServer":           true,
	"ha_transferLeader":         true,
	"ha_transferLeaderToServer": true,
}

// rpcEnvelope captures only the method field from a JSON-RPC request.
type rpcEnvelope struct {
	Method string `json:"method"`
}

// authMiddleware returns an HTTP handler that enforces token auth on write methods.
// If token is empty, the middleware is disabled and all requests pass through.
func authMiddleware(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Read and immediately restore the body so downstream can read it.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(body))

		if requiresAuth(body) {
			got := r.Header.Get("Authorization")
			if subtle.ConstantTimeCompare([]byte(got), []byte(token)) != 1 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":null,"error":{"code":-32001,"message":"unauthorized"}}`))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// requiresAuth reports whether the request body must carry a write-auth token.
//
// It is fail-closed: it returns true (require a token) for anything it cannot
// positively prove to be read-only. Parsing mirrors the downstream
// go-ethereum JSON-RPC server (rpc/json.go readBatch + parseMessage): a
// streaming decoder reads the first JSON value (trailing bytes ignored) and a
// batch is split element-by-element. Using the same parser as the server, plus
// failing closed on any decode error, removes the differential where a
// malformed body classified as "no write method" by the middleware is still
// executed as a write method by the server.
func requiresAuth(body []byte) bool {
	trimmed := bytes.TrimSpace(body)
	if len(trimmed) == 0 {
		return true
	}

	// Read exactly the first JSON value, like the server's codec does. A second
	// value or trailing bytes are ignored by both, so we classify what the
	// server will actually execute.
	dec := json.NewDecoder(bytes.NewReader(trimmed))
	var raw json.RawMessage
	if err := dec.Decode(&raw); err != nil {
		return true
	}

	if isJSONBatch(raw) {
		bd := json.NewDecoder(bytes.NewReader(raw))
		if _, err := bd.Token(); err != nil { // consume '['
			return true
		}
		for bd.More() {
			var req rpcEnvelope
			if err := bd.Decode(&req); err != nil {
				// The server decodes batch elements independently; an element we
				// can't classify might still be executed, so fail closed.
				return true
			}
			if writeRPCMethods[req.Method] {
				return true
			}
		}
		return false
	}

	var req rpcEnvelope
	if err := json.Unmarshal(raw, &req); err != nil {
		return true
	}
	return writeRPCMethods[req.Method]
}

// isJSONBatch reports whether the first non-whitespace byte is '[', matching
// go-ethereum's rpc.isBatch so batch detection stays identical to the server.
func isJSONBatch(raw json.RawMessage) bool {
	for _, c := range raw {
		if c == 0x20 || c == 0x09 || c == 0x0a || c == 0x0d {
			continue
		}
		return c == '['
	}
	return false
}
