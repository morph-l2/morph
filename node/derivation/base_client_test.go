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

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil, nil)
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

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil, nil)
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
	cl := NewFallbackHTTPClient([]string{"http://127.0.0.1:0", fallback}, nil, nil)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.EqualValues(t, 1, atomic.LoadInt32(fallbackHits))
}

// When every endpoint fails (non-200 or transport error), the last error is
// returned so the caller (apiReq) reports a failure.
func TestFallbackHTTPClient_AllFailReturnsError(t *testing.T) {
	primary, primaryHits := newBeaconStub(t, http.StatusServiceUnavailable, "down")
	fallback, fallbackHits := newBeaconStub(t, http.StatusNotFound, "missing")

	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil, nil)
	resp, err := doGet(t, cl)
	require.Error(t, err)
	require.Nil(t, resp)
	// Both endpoints are tried before giving up.
	require.EqualValues(t, 1, atomic.LoadInt32(primaryHits))
	require.EqualValues(t, 1, atomic.LoadInt32(fallbackHits))
}

// A per-endpoint failure is exposed via metrics so a flaky beacon is visible on
// dashboards.
func TestFallbackHTTPClient_RecordsFailureMetric(t *testing.T) {
	primary, _ := newBeaconStub(t, http.StatusServiceUnavailable, "down")
	fallback, _ := newBeaconStub(t, http.StatusOK, "recovered")

	m := PrometheusMetrics("morphnode_test_" + t.Name())
	cl := NewFallbackHTTPClient([]string{primary, fallback}, nil, m)
	resp, err := doGet(t, cl)
	require.NoError(t, err)
	resp.Body.Close()
	// Just assert the increment path does not panic with a real Metrics; the
	// counter value is scraped from /metrics in production.
}
