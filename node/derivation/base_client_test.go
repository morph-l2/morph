package derivation

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

// newBeaconStub returns a server that always answers with the given status and
// body, and a counter recording how many requests it received.
func newBeaconStub(t *testing.T, status int, body string) (string, *int32) {
	t.Helper()
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	return srv.URL, &hits
}

func doGet(t *testing.T, cl HTTP) (*http.Response, error) {
	t.Helper()
	return cl.Get(context.Background(), "eth/v1/beacon/genesis", http.Header{})
}

// The primary endpoint is used when healthy and the fallback is never touched.
func TestFallbackHTTPClient_PrimaryHealthy(t *testing.T) {
	primary, primaryHits := newBeaconStub(t, http.StatusOK, "ok")
	fallback, fallbackHits := newBeaconStub(t, http.StatusOK, "fallback")

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	require.Equal(t, "ok", string(body))
	require.EqualValues(t, 1, atomic.LoadInt32(primaryHits))
	require.EqualValues(t, 0, atomic.LoadInt32(fallbackHits), "fallback must not be queried while primary is healthy")
}

// A non-200 from the primary transparently falls through to a healthy fallback.
func TestFallbackHTTPClient_FallsBackOnNon200(t *testing.T) {
	primary, primaryHits := newBeaconStub(t, http.StatusInternalServerError, "boom")
	fallback, fallbackHits := newBeaconStub(t, http.StatusOK, "recovered")

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	require.Equal(t, "recovered", string(body))
	require.EqualValues(t, 1, atomic.LoadInt32(primaryHits))
	require.EqualValues(t, 1, atomic.LoadInt32(fallbackHits))
}

// A transport error (dead endpoint) also triggers fallback.
func TestFallbackHTTPClient_FallsBackOnTransportError(t *testing.T) {
	fallback, fallbackHits := newBeaconStub(t, http.StatusOK, "recovered")

	// http://127.0.0.1:0 is not listenable, so the first Get fails at the
	// transport layer before any response is produced.
	cl := NewFallbackHTTPClient([]string{"http://127.0.0.1:0", fallback}, nil)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.EqualValues(t, 1, atomic.LoadInt32(fallbackHits))
}

// When every endpoint returns non-200, the last response is surfaced so the
// caller (apiReq) can produce its usual "failed request with status" error.
func TestFallbackHTTPClient_AllNon200ReturnsLastResponse(t *testing.T) {
	primary, _ := newBeaconStub(t, http.StatusServiceUnavailable, "down")
	fallback, _ := newBeaconStub(t, http.StatusNotFound, "missing")

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

// When every endpoint fails at the transport layer, an error is returned.
func TestFallbackHTTPClient_AllTransportErrorsReturnError(t *testing.T) {
	cl := NewFallbackHTTPClient([]string{"http://127.0.0.1:0", "http://127.0.0.1:0"}, nil)
	resp, err := doGet(t, cl)
	require.Error(t, err)
	require.Nil(t, resp)
}
