package main

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMissingBatchIsAnError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":null}`))
	}))
	defer server.Close()
	oldFlags, oldArgs := flag.CommandLine, os.Args
	t.Cleanup(func() { flag.CommandLine, os.Args = oldFlags, oldArgs })
	flag.CommandLine = flag.NewFlagSet("batchparse", flag.ContinueOnError)
	os.Args = []string{"batchparse", "-rpc", server.URL, "-batch", "1"}
	if err := run(); err == nil {
		t.Fatal("an unavailable batch must not be reported as successfully parsed")
	}
}

func TestBatchIndexMustBeExplicit(t *testing.T) {
	oldFlags, oldArgs := flag.CommandLine, os.Args
	t.Cleanup(func() { flag.CommandLine, os.Args = oldFlags, oldArgs })
	flag.CommandLine = flag.NewFlagSet("batchparse", flag.ContinueOnError)
	os.Args = []string{"batchparse", "-rpc", "http://127.0.0.1:1"}
	if err := run(); err == nil || !strings.Contains(err.Error(), "-batch") {
		t.Fatalf("expected argument validation before an RPC call, got %v", err)
	}
}
