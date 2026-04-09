package rpc

import (
	"fmt"
	"net/http"
	"sync"

	ethrpc "github.com/morph-l2/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
)

// Server is an HTTP JSON-RPC server that exposes the hakeeper management API.
type Server struct {
	log        log.Logger
	listenAddr string
	listenPort int

	rpcServer  *ethrpc.Server
	httpServer *http.Server
	wg         sync.WaitGroup
}

// New creates a new Server. cons must implement ConsensusAdapter (defined in this package).
// token is the auth token for write APIs; pass empty string to disable auth.
func New(log log.Logger, listenAddr string, listenPort int, cons ConsensusAdapter, token string) (*Server, error) {
	rpcSrv := ethrpc.NewServer()

	backend := NewAPIBackend(log, cons)
	if err := rpcSrv.RegisterName(RPCNamespace, backend); err != nil {
		return nil, errors.Wrap(err, "failed to register hakeeper API")
	}

	if token == "" {
		log.Info("hakeeper RPC server has no auth token configured, write APIs are unprotected")
	}

	mux := http.NewServeMux()
	mux.Handle("/", authMiddleware(token, rpcSrv))

	addr := fmt.Sprintf("%s:%d", listenAddr, listenPort)
	httpSrv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &Server{
		log:        log,
		listenAddr: listenAddr,
		listenPort: listenPort,
		rpcServer:  rpcSrv,
		httpServer: httpSrv,
	}, nil
}

// Start begins listening for RPC connections in a background goroutine.
func (s *Server) Start() error {
	s.log.Info("Starting hakeeper RPC server", "addr", s.httpServer.Addr)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("hakeeper RPC server error", "err", err)
		}
	}()
	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop() {
	s.log.Info("Stopping hakeeper RPC server")
	if s.httpServer != nil {
		if err := s.httpServer.Close(); err != nil {
			s.log.Error("hakeeper RPC server shutdown error", "err", err)
		}
	}
	s.wg.Wait()
	if s.rpcServer != nil {
		s.rpcServer.Stop()
	}
	s.log.Info("hakeeper RPC server stopped")
}

// Addr returns the listening address of the server.
func (s *Server) Addr() string {
	return s.httpServer.Addr
}
