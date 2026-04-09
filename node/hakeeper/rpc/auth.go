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

// requiresAuth reports whether the request body contains any write JSON-RPC method.
// Handles both single requests ({...}) and batch requests ([...]).
func requiresAuth(body []byte) bool {
	trimmed := bytes.TrimSpace(body)
	if len(trimmed) == 0 {
		return false
	}

	if trimmed[0] == '[' {
		var batch []rpcEnvelope
		if err := json.Unmarshal(trimmed, &batch); err != nil {
			return false
		}
		for _, req := range batch {
			if writeRPCMethods[req.Method] {
				return true
			}
		}
		return false
	}

	var req rpcEnvelope
	if err := json.Unmarshal(trimmed, &req); err != nil {
		return false
	}
	return writeRPCMethods[req.Method]
}
