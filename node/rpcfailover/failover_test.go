package rpcfailover

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

const probeBody = `{"jsonrpc":"2.0","id":1,"method":"eth_blockNumber","params":[]}`

// recorder is a test endpoint that counts requests and remembers what it saw.
type recorder struct {
	server  *httptest.Server
	hits    atomic.Int32
	healthy atomic.Bool // used by newRecoverableEndpoint to bring an endpoint back
	bodies  chan string
	auths   chan string
}

// newEndpoint starts a test endpoint that replies with the given status. A 200
// reply carries a valid eth_blockNumber result so the same helper can be driven
// through ethclient.
func newEndpoint(t *testing.T, status int) *recorder {
	t.Helper()
	r := &recorder{
		bodies: make(chan string, 8),
		auths:  make(chan string, 8),
	}
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.hits.Add(1)
		body, _ := io.ReadAll(req.Body)
		r.bodies <- string(body)
		r.auths <- req.Header.Get("Authorization")

		if status != http.StatusOK {
			w.WriteHeader(status)
			return
		}
		var call struct {
			ID json.RawMessage `json:"id"`
		}
		_ = json.Unmarshal(body, &call)
		if len(call.ID) == 0 {
			call.ID = json.RawMessage("1")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":` + string(call.ID) + `,"result":"0x10"}`))
	}))
	t.Cleanup(r.server.Close)
	return r
}

// url returns the endpoint URL, optionally with basic-auth userinfo spliced in.
func (r *recorder) url(userinfo string) string {
	if userinfo == "" {
		return r.server.URL
	}
	return strings.Replace(r.server.URL, "http://", "http://"+userinfo+"@", 1)
}

func newTestTransport(t *testing.T, rawURLs ...string) *failoverTransport {
	t.Helper()
	endpoints := make([]*url.URL, 0, len(rawURLs))
	redacted := make([]string, 0, len(rawURLs))
	manageAuth := false
	for _, raw := range rawURLs {
		u, err := url.Parse(raw)
		require.NoError(t, err)
		if u.User != nil {
			manageAuth = true
		}
		endpoints = append(endpoints, u)
		redacted = append(redacted, redactEndpoint(raw))
	}
	tr := &failoverTransport{
		name:          "L1",
		base:          newBaseTransport(defaultAttemptTimeouts()),
		endpoints:     endpoints,
		redacted:      redacted,
		manageAuth:    manageAuth,
		failbackAfter: defaultFailbackAfter,
		log:           tmlog.NewNopLogger(),
	}
	tr.lastProbe.Store(time.Now().UnixNano()) // as Dial does
	return tr
}

// post issues a request shaped like the one rpc/http.go builds: a POST whose body
// is wrapped in a NopCloser, which leaves GetBody nil.
func post(t *testing.T, tr http.RoundTripper) (*http.Response, error) {
	t.Helper()
	req, err := http.NewRequest(http.MethodPost, "http://placeholder.invalid", io.NopCloser(strings.NewReader(probeBody)))
	require.NoError(t, err)
	require.Nil(t, req.GetBody, "test request must mimic rpc/http.go, which leaves GetBody nil")
	req.ContentLength = int64(len(probeBody))
	req.Header.Set("Content-Type", "application/json")
	return tr.RoundTrip(req)
}

func TestFailoverSwitchesOnServerError(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(1), secondary.hits.Load())

	// The buffered request body must be replayed intact onto the second attempt.
	assert.Equal(t, probeBody, <-primary.bodies)
	assert.Equal(t, probeBody, <-secondary.bodies)

	// The response must reach the caller unread, for the rpc layer to decode.
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"result":"0x10"`)
}

func TestFailoverSwitchesOnConnectionRefused(t *testing.T) {
	dead := newEndpoint(t, http.StatusOK)
	deadURL := dead.url("")
	dead.server.Close() // nothing is listening any more

	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, deadURL, secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), secondary.hits.Load())
}

func TestFailoverSticksToNewEndpoint(t *testing.T) {
	primary := newEndpoint(t, http.StatusInternalServerError)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	for i := 0; i < 3; i++ {
		resp, err := post(t, tr)
		require.NoError(t, err)
		resp.Body.Close()
	}

	// The dead primary is probed once, on the call that triggered the switch.
	// Later calls must not pay its cost again.
	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(3), secondary.hits.Load())
}

func TestFailoverKeepsClientErrors(t *testing.T) {
	primary := newEndpoint(t, http.StatusBadRequest)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	// A 400 describes the request, not the endpoint. It is surfaced verbatim and
	// the secondary is never consulted.
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(0), secondary.hits.Load())
}

// newRawEndpoint starts an endpoint that replies with a fixed status, content
// type and body, for responses that are not valid JSON-RPC.
func newRawEndpoint(t *testing.T, status int, contentType, body string) *recorder {
	t.Helper()
	r := &recorder{bodies: make(chan string, 8), auths: make(chan string, 8)}
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.hits.Add(1)
		b, _ := io.ReadAll(req.Body)
		r.bodies <- string(b)
		r.auths <- req.Header.Get("Authorization")
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(r.server.Close)
	return r
}

// TestFailoverOnHTMLBlockPage reproduces the failure seen on the devnet: when an
// endpoint's host stops existing, the corporate DNS resolver answers with a
// captive-portal address that serves HTTP 200 and an HTML block page. Judged on
// status alone that looks like success, so the node kept using a dead endpoint
// and every L1 read failed one layer up with "invalid character '<'".
func TestFailoverOnHTMLBlockPage(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "text/html;charset=utf-8",
		"<!DOCTYPE html>\n<html lang=\"en\"><head><meta charset=\"UTF-8\">")
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(1), secondary.hits.Load(), "an HTML body must be treated as endpoint failure")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"result":"0x10"`)
}

// TestFailoverOnEmptyBody covers a 200 with nothing in it, which is not a
// JSON-RPC reply either.
func TestFailoverOnEmptyBody(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "application/json", "")
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// TestNoFailoverOnJSONRPCError is the other half of the contract: when the
// endpoint answers with a well-formed JSON-RPC error, it is doing its job. The
// error belongs to the caller and must not send us to another endpoint, which
// would multiply load and return the same error anyway.
func TestNoFailoverOnJSONRPCError(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "application/json",
		`{"jsonrpc":"2.0","id":1,"error":{"code":-32000,"message":"execution reverted"}}`)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(0), secondary.hits.Load(), "a JSON-RPC error is an answer, not an endpoint failure")
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), "execution reverted")
}

// TestNoFailoverOnBatchResponse guards the array form: batch replies are valid
// JSON-RPC and must survive the check.
func TestNoFailoverOnBatchResponse(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "application/json",
		`[{"jsonrpc":"2.0","id":1,"result":"0x1"}]`)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, int32(0), secondary.hits.Load())
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, `[{"jsonrpc":"2.0","id":1,"result":"0x1"}]`, string(body),
		"the peeked byte must be handed back to the caller intact")
}

// TestNoFailoverOnLeadingWhitespace allows a body that starts with whitespace.
func TestNoFailoverOnLeadingWhitespace(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "application/json",
		"\r\n  {\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x2\"}")
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, int32(0), secondary.hits.Load())
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"result":"0x2"`)
}

// TestNoPeekOnCompressedBody: a compressed body cannot be inspected cheaply, so
// it is passed through rather than guessed at. Go's transport strips
// Content-Encoding when it does the decompression itself, so this only applies
// when a caller set Accept-Encoding explicitly via rpc.WithHeader.
func TestNoPeekOnCompressedBody(t *testing.T) {
	primary := &recorder{bodies: make(chan string, 8), auths: make(chan string, 8)}
	primary.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		primary.hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		// Deliberately not valid JSON and not valid gzip: the point is that the
		// body is never looked at when it announces an encoding.
		_, _ = w.Write([]byte("\x1f\x8b\x08 not-json-but-compressed"))
	}))
	t.Cleanup(primary.server.Close)

	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	// Go's transport only leaves Content-Encoding on the response when it did not
	// add Accept-Encoding itself — otherwise it decompresses and strips the
	// header. Setting it explicitly reproduces the one condition under which the
	// guard is reachable, which is a caller using rpc.WithHeader.
	req, err := http.NewRequest(http.MethodPost, "http://placeholder.invalid", io.NopCloser(strings.NewReader(probeBody)))
	require.NoError(t, err)
	req.Header.Set("Accept-Encoding", "gzip")
	req.ContentLength = int64(len(probeBody))

	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, "gzip", resp.Header.Get("Content-Encoding"),
		"precondition: the transport must not have stripped Content-Encoding")
	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(0), secondary.hits.Load(), "a compressed body must be passed through unexamined")
}

// TestFailoverOnHTMLBlockPageWithoutContentType pins that the decision is made
// on the body, not the header: the same block page with no Content-Type at all
// must still be rejected.
func TestFailoverOnHTMLBlockPageWithoutContentType(t *testing.T) {
	primary := newRawEndpoint(t, http.StatusOK, "", "<!DOCTYPE html><html></html>")
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// TestClientErrorBodyNotJudged pins the scoping decision: the JSON check applies
// to 2xx only, so a 4xx with an empty or HTML body is still returned verbatim
// rather than being turned into a failover.
func TestClientErrorBodyNotJudged(t *testing.T) {
	// 404 rather than 403: 403 now fails over on its own, which would hide what
	// this test is about.
	primary := newRawEndpoint(t, http.StatusNotFound, "text/html", "<html>not found</html>")
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, int32(0), secondary.hits.Load())
}

// breakableConn fails writes once the test flips its flag, while reads keep
// delegating to the real connection. That combination is what produces net/http's
// nothingWrittenError: the write fails before a single byte reaches the wire, and
// because the read side never sees a close, readLoop does not evict the
// connection from the idle pool first (transport.go removeIdleConn).
type breakableConn struct {
	net.Conn
	broken *atomic.Bool
}

func (c *breakableConn) Write(b []byte) (int, error) {
	if c.broken.Load() {
		return 0, errors.New("simulated dead pooled connection")
	}
	return c.Conn.Write(b)
}

// TestStalePooledConnectionRetriesSameEndpoint pins that a connection dying in
// the pool is not mistaken for the endpoint being down.
//
// net/http can re-send such a request on a fresh connection by itself, but only
// if it can rewind the body: shouldRetryRequest's nothingWrittenError branch
// returns `req.outgoingLength() == 0 || req.GetBody != nil`, and a POST with a
// body and no GetBody fails both. Without GetBody the error surfaces here
// instead, and the sticky failover then demotes a perfectly healthy primary for
// the rest of the process's life.
func TestStalePooledConnectionRetriesSameEndpoint(t *testing.T) {
	primary := newEndpoint(t, http.StatusOK)
	secondary := newEndpoint(t, http.StatusOK)

	var broken atomic.Bool
	var dials atomic.Int32
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	base := http.DefaultTransport.(*http.Transport).Clone()
	base.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}
		// Only the first connection — the one that will be reused — is breakable.
		// A redial must succeed, which is what the retry depends on.
		if dials.Add(1) == 1 {
			return &breakableConn{Conn: conn, broken: &broken}, nil
		}
		return conn, nil
	}

	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	tr.base = base

	// First call establishes the connection. Draining the body is what returns it
	// to the idle pool.
	resp, err := post(t, tr)
	require.NoError(t, err)
	_, err = io.Copy(io.Discard, resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	require.Equal(t, int32(1), primary.hits.Load())

	// The pooled connection dies without the client noticing.
	broken.Store(true)

	resp, err = post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, int32(0), secondary.hits.Load(),
		"a connection that died in the pool is not an endpoint failure and must not trigger failover")
	assert.Equal(t, int32(2), primary.hits.Load(),
		"the request should have been re-sent to the same endpoint on a fresh connection")
	assert.Equal(t, int32(0), tr.cur.Load(), "the sticky endpoint must still be the primary")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"result":"0x10"`)
}

// TestAttemptSuppliesGetBody is the narrow contract the retry above depends on.
func TestAttemptSuppliesGetBody(t *testing.T) {
	only := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, only.url(""), only.url(""))

	var captured *http.Request
	tr.base = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`)),
		}, nil
	})

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotNil(t, captured)
	require.NotNil(t, captured.GetBody, "net/http needs GetBody to rewind a POST body")
	rewound, err := captured.GetBody()
	require.NoError(t, err)
	replayed, err := io.ReadAll(rewound)
	require.NoError(t, err)
	assert.Equal(t, probeBody, string(replayed), "GetBody must yield the full original body")
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// --- #4: credentials are per endpoint, so 401/403 must move on ---

// TestFailoverOnUnauthorized covers an expired or revoked key on the primary.
// Endpoints can carry their own credentials, so 401 on one says nothing about
// the next; refusing to move would blind the caller while a working fallback
// sat idle.
func TestFailoverOnUnauthorized(t *testing.T) {
	primary := newEndpoint(t, http.StatusUnauthorized)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url("alice:expired"), secondary.url("bob:valid"))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// TestFailoverOnForbidden covers an exhausted quota, which hosted providers
// report as 403.
func TestFailoverOnForbidden(t *testing.T) {
	primary := newEndpoint(t, http.StatusForbidden)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// TestNoFailoverOnBadRequestOrNotFound keeps the other half of the rule: these
// describe the request, so another endpoint would only reproduce them.
func TestNoFailoverOnBadRequestOrNotFound(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusNotFound} {
		primary := newEndpoint(t, status)
		secondary := newEndpoint(t, http.StatusOK)
		tr := newTestTransport(t, primary.url(""), secondary.url(""))

		resp, err := post(t, tr)
		require.NoError(t, err)
		resp.Body.Close()
		assert.Equal(t, status, resp.StatusCode)
		assert.Equal(t, int32(0), secondary.hits.Load(), "status %d must not fail over", status)
	}
}

// --- #5: a malformed endpoint must not leak its credentials into the error ---

func TestDialErrorDoesNotLeakCredentials(t *testing.T) {
	const secret = "s3cr3t-api-key"
	// A control character makes url.Parse fail; its own error text embeds the raw
	// URL, which is exactly what must not reach the caller.
	bad := "http://user:" + secret + "@rpc.example.com/v3/" + secret + "/\x7f"

	_, err := Dial(context.Background(), "L1", bad+",http://ok.invalid", tmlog.NewNopLogger())
	require.Error(t, err)
	assert.NotContains(t, err.Error(), secret, "error must not echo credentials from the raw URL")
	assert.Contains(t, err.Error(), "is not a valid URL")
}

func TestDialRejectsEndpointWithoutHost(t *testing.T) {
	// "http://:8545" parses with a non-empty Host of ":8545" and no host at all,
	// so checking Host alone would let it through.
	for _, bad := range []string{"http://", "http://:8545"} {
		_, err := Dial(context.Background(), "L1", bad+",http://ok.invalid", tmlog.NewNopLogger())
		require.Error(t, err, "endpoint %q must be rejected", bad)
		assert.Contains(t, err.Error(), "has no host")
	}
}

// TestRuntimeErrorDoesNotLeakCredentials covers the error path that outlives
// startup: when every endpoint fails, http.Client wraps the transport error in a
// url.Error carrying the request URL, and its stripPassword only masks userinfo —
// an API key in the path would ride along into every caller's error log.
func TestRuntimeErrorDoesNotLeakCredentials(t *testing.T) {
	const secret = "s3cr3t-api-key"
	a := newEndpoint(t, http.StatusOK)
	b := newEndpoint(t, http.StatusOK)
	// The key in the path is how hosted providers embed it (Infura /v3/, Alchemy /v2/).
	urlA := a.url("") + "/v3/" + secret
	urlB := b.url("") + "/v3/" + secret
	a.server.Close() // both refuse connections, so the call fails at runtime
	b.server.Close()

	client, err := Dial(context.Background(), "L1", urlA+","+urlB, tmlog.NewNopLogger())
	require.NoError(t, err)
	defer client.Close()

	_, err = client.BlockNumber(context.Background())
	require.Error(t, err)
	assert.NotContains(t, err.Error(), secret,
		"the URL http.Client reports on a transport error must not carry the API key")
}

// TestFailoverOnRedirect: attempt() rewrites every request's URL to its endpoint,
// so a redirect target would be discarded and the same endpoint asked again until
// http.Client gives up after ten hops. Treating 3xx as an endpoint failure is
// what keeps that from happening.
func TestFailoverOnRedirect(t *testing.T) {
	redirecting := &recorder{bodies: make(chan string, 8), auths: make(chan string, 8)}
	redirecting.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		redirecting.hits.Add(1)
		w.Header().Set("Location", "https://elsewhere.invalid/rpc")
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(redirecting.server.Close)

	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, redirecting.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), redirecting.hits.Load(), "the redirect must be asked once, not looped over")
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// --- #6: a recovered primary is re-adopted ---

func TestFailsBackToPrimaryAfterWindow(t *testing.T) {
	primary := newRecoverableEndpoint(t)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	tr.failbackAfter = time.Millisecond

	primary.healthy.Store(false)
	resp, err := post(t, tr)
	require.NoError(t, err)
	resp.Body.Close()
	require.Equal(t, int32(1), tr.cur.Load(), "should have moved to the secondary")

	primary.healthy.Store(true)
	time.Sleep(2 * time.Millisecond) // let the failback window elapse

	resp, err = post(t, tr)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, int32(0), tr.cur.Load(), "a healthy primary must be re-adopted")
	assert.Equal(t, int32(1), secondary.hits.Load(), "the secondary should not have been needed again")
}

// TestNormalCallDoesNotTouchProbeWindow pins that the failback clock is only
// advanced when the primary is actually tried: if every successful call restamped
// it, a busy node would never reach the window and the primary would never be
// retried.
func TestNormalCallDoesNotTouchProbeWindow(t *testing.T) {
	only := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, only.url(""), only.url(""))
	tr.lastProbe.Store(12345)

	for i := 0; i < 3; i++ {
		resp, err := post(t, tr)
		require.NoError(t, err)
		resp.Body.Close()
	}

	assert.Equal(t, int64(12345), tr.lastProbe.Load(),
		"a call served by the endpoint already in use must not move the failback clock")
}

// TestSwitchDoesNotPostponeProbe: only dueForPrimaryProbe advances the failback
// clock. If switching endpoints stamped it too, fallbacks flapping faster than
// failbackAfter would keep pushing the next probe out and the primary would never
// be retried.
func TestSwitchDoesNotPostponeProbe(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	// Long window, so no probe is due during this call and the clock can only
	// change if the switch itself stamps it.
	tr.failbackAfter = time.Hour
	stamp := time.Now().Add(-time.Minute).UnixNano()
	tr.lastProbe.Store(stamp)

	resp, err := post(t, tr)
	require.NoError(t, err)
	resp.Body.Close()

	require.Equal(t, int32(1), tr.cur.Load(), "should have switched to the secondary")
	assert.Equal(t, stamp, tr.lastProbe.Load(), "a switch must leave the failback clock alone")
}

// TestSwitchAfterWindowElapsedReprobesOnce is the accepted cost of the above: a
// switch that happens once the window has already passed is followed by one
// immediate re-probe. That also means a primary which only blipped is picked back
// up at once rather than after a full window.
func TestSwitchAfterWindowElapsedReprobesOnce(t *testing.T) {
	primary := newRecoverableEndpoint(t)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	tr.failbackAfter = time.Millisecond

	primary.healthy.Store(false)
	resp, err := post(t, tr)
	require.NoError(t, err)
	resp.Body.Close()
	require.Equal(t, int32(1), tr.cur.Load())
	require.Equal(t, int32(1), primary.hits.Load())

	time.Sleep(2 * time.Millisecond) // window elapses while still on the fallback

	// Primary came back between the two calls, so the re-probe adopts it again.
	primary.healthy.Store(true)
	resp, err = post(t, tr)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, int32(2), primary.hits.Load(), "the primary is re-probed once the window has passed")
	assert.Equal(t, int32(0), tr.cur.Load(), "and re-adopted now that it answers")
}

func TestDoesNotProbePrimaryEveryCallWhileItIsDown(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	// Long window: the first call switches away, later calls must not re-probe.
	tr.failbackAfter = time.Hour

	for i := 0; i < 4; i++ {
		resp, err := post(t, tr)
		require.NoError(t, err)
		resp.Body.Close()
	}

	assert.Equal(t, int32(1), primary.hits.Load(),
		"the dead primary is probed once, not on every call")
	assert.Equal(t, int32(4), secondary.hits.Load())
}

// newRecoverableEndpoint serves 503 or a valid reply depending on a flag, so a
// test can bring an endpoint back up.
func newRecoverableEndpoint(t *testing.T) *recorder {
	t.Helper()
	r := &recorder{bodies: make(chan string, 8), auths: make(chan string, 8)}
	r.healthy.Store(true)
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.hits.Add(1)
		if !r.healthy.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x10"}`))
	}))
	t.Cleanup(r.server.Close)
	return r
}

// --- a body that stops arriving must be bounded, not block forever ---

// newStallingEndpoint starts an endpoint that writes its headers and prefix,
// flushes them onto the wire, and then stops writing without closing the
// connection — the shape neither ResponseHeaderTimeout nor a TCP error catches.
// The returned channel must be closed by the test before its server is torn
// down, since httptest.Server.Close waits for outstanding handlers.
func newStallingEndpoint(t *testing.T, prefix string) (*recorder, chan struct{}) {
	t.Helper()
	release := make(chan struct{})
	r := &recorder{bodies: make(chan string, 8), auths: make(chan string, 8)}
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if prefix != "" {
			_, _ = io.WriteString(w, prefix)
		}
		w.(http.Flusher).Flush()
		<-release
	}))
	t.Cleanup(r.server.Close)
	return r, release
}

// TestFailoverWhenBodyNeverArrives: the stall lands inside requireJSONBody's
// peek, which is still inside the attempt, so it is a normal endpoint failure and
// the next endpoint serves the call.
func TestFailoverWhenBodyNeverArrives(t *testing.T) {
	primary, release := newStallingEndpoint(t, "")
	defer close(release)
	secondary := newEndpoint(t, http.StatusOK)

	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	to := defaultAttemptTimeouts()
	to.idleRead = 100 * time.Millisecond
	tr.base = newBaseTransport(to)

	started := time.Now()
	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(1), secondary.hits.Load(), "a body that never arrives is an endpoint failure")
	assert.Less(t, time.Since(started), 5*time.Second, "must not have waited indefinitely on the stalled endpoint")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"result":"0x10"`)
}

// TestStalledBodyMidStreamFailsBounded is the residual case: enough bytes arrive
// to satisfy the peek, so the response is handed up and the stall happens past
// the failover decision point. No switch is possible there — but the read must
// fail rather than hang, which is what lets L1Tracker's loop iterate again and
// its halt gate keep working.
func TestStalledBodyMidStreamFailsBounded(t *testing.T) {
	primary, release := newStallingEndpoint(t, `{"jsonrpc":"2.0",`)
	defer close(release)
	secondary := newEndpoint(t, http.StatusOK)

	tr := newTestTransport(t, primary.url(""), secondary.url(""))
	to := defaultAttemptTimeouts()
	to.idleRead = 100 * time.Millisecond
	tr.base = newBaseTransport(to)

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, int32(0), secondary.hits.Load(), "the peek succeeded, so this call does not fail over")

	started := time.Now()
	_, err = io.ReadAll(resp.Body)
	require.Error(t, err, "a stalled body must fail rather than block forever")
	assert.Less(t, time.Since(started), 5*time.Second)
}

// TestHTTP2IsDisabled pins the assumption idleRead rests on. Under HTTP/2 one
// connection carries many concurrent streams, so traffic on any of them would
// push a per-connection read deadline forward and a single stalled stream would
// be unbounded again. The server here offers h2 over ALPN; the transport must
// still come back on HTTP/1.1.
func TestHTTP2IsDisabled(t *testing.T) {
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x10"}`))
	}))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	t.Cleanup(srv.Close)

	// Control: without this the assertion below would also pass against a server
	// that never offered h2 in the first place.
	ctrl, err := srv.Client().Post(srv.URL, "application/json", strings.NewReader(probeBody))
	require.NoError(t, err)
	require.NoError(t, ctrl.Body.Close())
	require.Equal(t, "HTTP/2.0", ctrl.Proto, "precondition: the test server must offer HTTP/2")

	tr := newTestTransport(t, srv.URL, srv.URL)
	base := newBaseTransport(defaultAttemptTimeouts())
	// Trust the test server's certificate, but only that: replacing the whole TLS
	// config would re-advertise h2 over ALPN, and the server would then select a
	// protocol the HTTP/1.x code cannot read.
	base.TLSClientConfig.RootCAs = srv.Client().Transport.(*http.Transport).TLSClientConfig.RootCAs
	tr.base = base

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, "HTTP/1.1", resp.Proto,
		"HTTP/2 must stay disabled: idleRead is a per-connection bound and needs one request per connection")
}

func TestFailoverOnTooManyRequests(t *testing.T) {
	primary := newEndpoint(t, http.StatusTooManyRequests)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), secondary.hits.Load())
}

func TestFailoverAllEndpointsFail(t *testing.T) {
	primary := newEndpoint(t, http.StatusBadGateway)
	secondary := newEndpoint(t, http.StatusServiceUnavailable)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	resp, err := post(t, tr)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "all 2 rpc endpoints failed")
	assert.Equal(t, int32(1), primary.hits.Load())
	assert.Equal(t, int32(1), secondary.hits.Load())
}

func TestFailoverAppliesPerEndpointBasicAuth(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url("alice:primary-secret"), secondary.url("bob:secondary-secret"))

	// http.Client.send would have stamped the primary's credentials on the
	// request before RoundTrip; reproduce that so the test covers the leak.
	req, err := http.NewRequest(http.MethodPost, primary.url("alice:primary-secret"), io.NopCloser(strings.NewReader(probeBody)))
	require.NoError(t, err)
	req.SetBasicAuth("alice", "primary-secret")
	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	primaryAuth, secondaryAuth := <-primary.auths, <-secondary.auths
	assert.NotEmpty(t, primaryAuth)
	assert.NotEqual(t, primaryAuth, secondaryAuth, "primary credentials must not be forwarded to the secondary endpoint")

	user, pass, ok := (&http.Request{Header: http.Header{"Authorization": {secondaryAuth}}}).BasicAuth()
	require.True(t, ok)
	assert.Equal(t, "bob", user)
	assert.Equal(t, "secondary-secret", pass)
}

func TestFailoverDropsInheritedAuthWhenEndpointHasNone(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url("alice:primary-secret"), secondary.url(""))

	req, err := http.NewRequest(http.MethodPost, primary.url("alice:primary-secret"), io.NopCloser(strings.NewReader(probeBody)))
	require.NoError(t, err)
	req.SetBasicAuth("alice", "primary-secret")
	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	<-primary.auths
	assert.Empty(t, <-secondary.auths, "an endpoint without userinfo must not receive another endpoint's credentials")
}

// TestDialFailoverThroughEthclient checks the whole stack: a dead primary, and
// ethclient decoding a real JSON-RPC response served by the secondary.
func TestDialFailoverThroughEthclient(t *testing.T) {
	dead := newEndpoint(t, http.StatusOK)
	deadURL := dead.url("")
	dead.server.Close()

	secondary := newEndpoint(t, http.StatusOK)

	client, err := Dial(context.Background(), "L1", deadURL+" , "+secondary.url(""), tmlog.NewNopLogger())
	require.NoError(t, err)
	defer client.Close()

	height, err := client.BlockNumber(context.Background())
	require.NoError(t, err)
	assert.Equal(t, uint64(0x10), height)
	assert.Equal(t, int32(1), secondary.hits.Load())
}

// TestDialSingleEndpoint pins the compatibility promise: one endpoint behaves
// exactly like the previous ethclient.Dial call.
func TestDialSingleEndpoint(t *testing.T) {
	only := newEndpoint(t, http.StatusOK)

	client, err := Dial(context.Background(), "L1", only.url(""), tmlog.NewNopLogger())
	require.NoError(t, err)
	defer client.Close()

	height, err := client.BlockNumber(context.Background())
	require.NoError(t, err)
	assert.Equal(t, uint64(0x10), height)
}

func TestDialRejectsNonHTTPWithMultipleEndpoints(t *testing.T) {
	_, err := Dial(context.Background(), "L1", "http://a.invalid,ws://b.invalid", tmlog.NewNopLogger())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requires http(s)")
	// The rejected endpoint must not be echoed with its credentials.
	assert.Contains(t, err.Error(), "ws://b.invalid")
}

func TestDialAllowsNonHTTPSingleEndpoint(t *testing.T) {
	// ws:// with a single endpoint keeps the old code path. The dial fails because
	// nothing is listening, not because the scheme was rejected.
	_, err := Dial(context.Background(), "L1", "ws://127.0.0.1:1/", tmlog.NewNopLogger())
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "requires http(s)")
}

func TestDialNoEndpoint(t *testing.T) {
	_, err := Dial(context.Background(), "L1", "  ,  ", tmlog.NewNopLogger())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no rpc endpoint configured")
}

func TestSplitEndpoints(t *testing.T) {
	assert.Equal(t, []string{"http://a"}, splitEndpoints("http://a"))
	assert.Equal(t, []string{"http://a", "http://b"}, splitEndpoints(" http://a , http://b "))
	assert.Equal(t, []string{"http://a", "http://b"}, splitEndpoints("http://a,,http://b"))
	assert.Nil(t, splitEndpoints(""))
}

func TestRedactEndpoint(t *testing.T) {
	// Credentials hide in userinfo, path and query; all three must be dropped.
	assert.Equal(t, "https://rpc.example.com",
		redactEndpoint("https://alice:secret@rpc.example.com/v3/abcdef?apikey=xyz"))
	assert.Equal(t, "<invalid-endpoint>", redactEndpoint("not a url"))

	// The port must survive, otherwise two endpoints on one host — the devnet and
	// local-node shape — become indistinguishable in the failover log line.
	assert.Equal(t, "http://127.0.0.1:8545", redactEndpoint("http://127.0.0.1:8545"))
	assert.NotEqual(t,
		redactEndpoint("http://127.0.0.1:8545"),
		redactEndpoint("http://127.0.0.1:8546"))
}

// TestRedactedEndpointsAreDistinct guards the log line itself: two test servers
// differ only by port, so a redaction that dropped it would make every switch
// message read "from=http://127.0.0.1 to=http://127.0.0.1".
func TestRedactedEndpointsAreDistinct(t *testing.T) {
	primary := newEndpoint(t, http.StatusServiceUnavailable)
	secondary := newEndpoint(t, http.StatusOK)
	tr := newTestTransport(t, primary.url(""), secondary.url(""))

	require.Len(t, tr.redacted, 2)
	assert.NotEqual(t, tr.redacted[0], tr.redacted[1])
}

func TestShouldFailover(t *testing.T) {
	// 401 and 403 are in this list because endpoints can carry their own
	// credentials: an expired key or exhausted quota on one provider says nothing
	// about the next. 3xx is here because attempt() rewrites the URL, so a redirect
	// could never be followed — only looped over.
	for _, status := range []int{408, 429, 401, 403, 301, 302, 307, 308, 500, 502, 503, 504} {
		assert.True(t, shouldFailover(status), "status %d should trigger failover", status)
	}
	// These describe the request, so another endpoint would only reproduce them.
	for _, status := range []int{200, 201, 400, 404} {
		assert.False(t, shouldFailover(status), "status %d should not trigger failover", status)
	}
}
