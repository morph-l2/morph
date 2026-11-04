package httpapi

import (
	"context"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/challenger/metrics"
	"morph-l2/challenger/service"
)

// Challenger submits one challengeState transaction.
type Challenger interface {
	Challenge(ctx context.Context, batchIndex uint64, batchHash common.Hash) (service.ChallengeResult, error)
}

// Server exposes the external trigger for a challenge.
type Server struct {
	challenger   Challenger
	token        string
	writeTimeout time.Duration
	httpServer   *http.Server
}

// NewServer builds the challenge HTTP server.
func NewServer(addr, token string, challenger Challenger, writeTimeout time.Duration) *Server {
	s := &Server{
		challenger:   challenger,
		token:        token,
		writeTimeout: writeTimeout,
	}
	s.httpServer = &http.Server{
		Addr:              addr,
		Handler:           s.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      writeTimeout + 30*time.Second,
	}
	return s
}

// Handler returns the HTTP routes.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.handleHealth)
	mux.HandleFunc("POST /challenge", s.handleChallenge)
	return mux
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

type challengeRequest struct {
	BatchIndex *uint64 `json:"batchIndex"`
	BatchHash  string  `json:"batchHash"`
}

type challengeResponse struct {
	TxHash      string `json:"txHash,omitempty"`
	Status      string `json:"status"`
	BlockNumber uint64 `json:"blockNumber,omitempty"`
	BatchIndex  uint64 `json:"batchIndex"`
	Error       string `json:"error,omitempty"`
}

func (s *Server) handleChallenge(w http.ResponseWriter, r *http.Request) {
	if !s.authorized(r) {
		metrics.Challenges.WithLabelValues("unauthorized").Inc()
		writeJSON(w, http.StatusUnauthorized, challengeResponse{Status: "unauthorized", Error: "unauthorized"})
		return
	}

	var req challengeRequest
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		metrics.Challenges.WithLabelValues("bad_request").Inc()
		writeJSON(w, http.StatusBadRequest, challengeResponse{Status: "bad_request", Error: "invalid json body"})
		return
	}
	if req.BatchIndex == nil {
		metrics.Challenges.WithLabelValues("bad_request").Inc()
		writeJSON(w, http.StatusBadRequest, challengeResponse{Status: "bad_request", Error: "batchIndex is required"})
		return
	}
	batchHash, err := parseBatchHash(req.BatchHash)
	if err != nil {
		metrics.Challenges.WithLabelValues("bad_request").Inc()
		writeJSON(w, http.StatusBadRequest, challengeResponse{Status: "bad_request", Error: err.Error(), BatchIndex: *req.BatchIndex})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.writeTimeout)
	defer cancel()

	result, err := s.challenger.Challenge(ctx, *req.BatchIndex, batchHash)
	resp := challengeResponse{
		BatchIndex:  *req.BatchIndex,
		BlockNumber: result.BlockNumber,
	}
	if result.TxHash != (common.Hash{}) {
		resp.TxHash = result.TxHash.Hex()
	}

	var rejected *service.RejectError
	var busy *service.BusyError
	var txErr *service.TxError
	switch {
	case err == nil:
		metrics.Challenges.WithLabelValues("success").Inc()
		resp.Status = "success"
		writeJSON(w, http.StatusOK, resp)
	case errors.As(err, &rejected):
		metrics.Challenges.WithLabelValues("rejected").Inc()
		resp.Status = "rejected"
		resp.Error = rejected.Reason
		log.Info("Challenge rejected", "batchIndex", *req.BatchIndex, "reason", rejected.Reason)
		writeJSON(w, http.StatusConflict, resp)
	case errors.As(err, &busy):
		metrics.Challenges.WithLabelValues("busy").Inc()
		resp.Status = "busy"
		resp.Error = busy.Error()
		writeJSON(w, http.StatusConflict, resp)
	case errors.As(err, &txErr) && txErr.Pending:
		metrics.Challenges.WithLabelValues("pending").Inc()
		resp.Status = "pending"
		resp.Error = txErr.Error()
		if resp.TxHash == "" {
			resp.TxHash = txErr.TxHash
		}
		writeJSON(w, http.StatusAccepted, resp)
	case errors.As(err, &txErr) && txErr.Reverted:
		metrics.Challenges.WithLabelValues("reverted").Inc()
		resp.Status = "reverted"
		resp.Error = txErr.Error()
		if resp.TxHash == "" {
			resp.TxHash = txErr.TxHash
		}
		writeJSON(w, http.StatusBadGateway, resp)
	default:
		metrics.Challenges.WithLabelValues("error").Inc()
		resp.Status = "error"
		resp.Error = err.Error()
		log.Error("Challenge failed", "batchIndex", *req.BatchIndex, "error", err)
		writeJSON(w, http.StatusInternalServerError, resp)
	}
}

func (s *Server) authorized(r *http.Request) bool {
	if s.token == "" {
		return true
	}
	const prefix = "Bearer "
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, prefix) {
		return false
	}
	got := header[len(prefix):]
	if len(got) != len(s.token) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(got), []byte(s.token)) == 1
}

func parseBatchHash(raw string) (common.Hash, error) {
	raw = strings.TrimSpace(raw)
	if !strings.HasPrefix(raw, "0x") && !strings.HasPrefix(raw, "0X") {
		return common.Hash{}, fmt.Errorf("batchHash must be 0x-prefixed")
	}
	encoded := raw[2:]
	if len(encoded) != 64 {
		return common.Hash{}, fmt.Errorf("batchHash must be 32 bytes")
	}
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		return common.Hash{}, fmt.Errorf("batchHash is not hex")
	}
	hash := common.BytesToHash(decoded)
	if hash == (common.Hash{}) {
		return common.Hash{}, fmt.Errorf("batchHash must be non-zero")
	}
	return hash, nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
