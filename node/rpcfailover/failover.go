// Package rpcfailover dials an Ethereum JSON-RPC endpoint with transparent
// failover across several endpoints.
//
// Nothing here is specific to L1, or to any particular consumer: the transport
// swaps a target URL and judges a response, and the caller supplies a name used
// only to label logs and errors. To get failover for a different set of
// endpoints, pass a different name — do not reimplement this.
//
// # Constraints
//
// These are real limits of the approach, not accidents of the current caller:
//
//   - HTTP(S) only. Over ws:// or IPC the rpc.Client holds a live connection and
//     reconnects through its own reconnectFunc, so moving to another endpoint is
//     not a matter of rewriting a URL. Dial therefore only engages failover when
//     more than one endpoint is configured, and requires http(s) in that case.
//   - JSON-RPC only. Health is partly judged by the body starting like a JSON
//     document. A REST API — the beacon node, say — needs its own scheme; see
//     derivation.FallbackBeaconClient.
//   - Read-only, or writes that are idempotent. A failed attempt is re-sent to
//     the next endpoint, and "no usable response" does not prove "not executed":
//     a proxy can return 502 after the request was processed. Re-sending a
//     pre-signed raw transaction is harmless (same hash), but anything order- or
//     nonce-dependent needs thought before being wired to this. Note that
//     tx-submitter's L2Clients deliberately routes writes to one endpoint only.
//
// # Why the transport layer
//
// Failover lives in an http.RoundTripper rather than in a wrapper around
// ethclient, and that placement is the whole point. rpc/http.go builds a fresh
// http.Request per call, so swapping the target URL underneath the rpc layer is
// invisible to everything above it. Consumers keep taking *ethclient.Client, and
// a contract binding keeps receiving a full bind.ContractBackend — including the
// SubscribeFilterLogs method that no HTTP-only client could honestly implement.
//
// # What counts as a failed endpoint
//
// A transport failure (connection refused, TLS and DNS errors, per-attempt
// timeout), a status of 408, 429 or 5xx, or a reply that is not JSON at all.
//
// That last rule is not theoretical. On a network whose resolver answers for
// names that do not exist, a vanished endpoint does not produce "connection
// refused": the name resolves to a captive-portal address that serves HTTP 200
// and an HTML block page. Judged on status alone that reads as success, so the
// caller keeps using a dead endpoint and every read fails one layer up with
// "invalid character '<' looking for beginning of value" — no failover, and no
// error the caller can act on. The same shape occurs behind load-balancer error
// pages and misrouted DNS.
//
// The check is the first non-whitespace byte of the body, not the Content-Type
// header: a JSON-RPC reply always starts with '{' or '[', whereas the header is
// sometimes mislabelled (Go's own content sniffer reports text/plain for a JSON
// body), and rejecting a working endpoint is worse than missing a broken one.
//
// Two cases are deliberately out of scope:
//
//   - A well-formed JSON-RPC error inside an HTTP 200. That is the endpoint
//     doing its job; the error belongs to the caller. Treating it as a failure
//     would multiply load across every endpoint and return the same error.
//   - An endpoint that answers promptly with a stale view of the chain. Nothing
//     at the HTTP layer can see that.
//
// For the L1 caller both are backstopped by the L1Tracker halt state machine,
// which stops block production once the node's L1 view has been blind for too
// long. A new caller needs its own answer for them.
package rpcfailover

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// defaultFailbackAfter is how long to keep using a fallback endpoint before
// spending one call to see whether the primary is back. "Primary first" is a
// priority, so a recovered primary should be re-adopted rather than waiting for
// the fallback to fail in turn.
//
// Deliberately not a background health check: no goroutine to manage, no probe
// traffic, and the cost of a still-dead primary is one failed attempt per window
// rather than per call.
//
// Separate from attemptTimeouts because it is failover policy, not a bound on a
// single attempt.
const defaultFailbackAfter = 5 * time.Minute

// attemptTimeouts bounds one attempt against one endpoint, phase by phase. They
// travel together because they are only meaningful as a set — what is not
// covered by one has to be covered by the next — and because a test that needs
// to shrink one should not have to restate the others.
type attemptTimeouts struct {
	// dial bounds DNS resolution and TCP setup, so an endpoint whose host is
	// blackholed cannot consume the caller's entire deadline. TLS setup is bounded
	// separately by the inherited TLSHandshakeTimeout of 10s.
	dial time.Duration

	// responseHeader bounds the wait for response headers. This is what makes
	// failover work against an endpoint that accepts the connection and then never
	// answers — the common shape of a hung RPC provider. The base transport
	// enforces it per RoundTrip, so each attempt gets its own budget without
	// wrapping the caller's context.
	//
	// It does not cover a server that sends headers and then stalls the body —
	// net/http documents that it "does not include the time to read the response
	// body". idleRead is what bounds that.
	responseHeader time.Duration

	// idleRead bounds how long a read may make no progress at all, and is the only
	// bound on the response body: responseHeader stops at the headers, the rpc
	// layer sets no deadline, and the node's callers pass contexts derived from
	// context.Background(). Without it a server that sends headers and then goes
	// quiet blocks its caller forever, which matters more than it looks: L1Tracker
	// calls the RPC inline from its tick loop and evaluates the halt gate inside
	// that same loop, so a call that never returns leaves the gate stuck open —
	// the node keeps producing against an arbitrarily stale L1 view with nothing
	// to report it.
	//
	// An inactivity bound, not a total one: a legitimately large reply (a wide
	// eth_getLogs can be megabytes) is never cut off while bytes keep arriving.
	//
	// It also applies to a pooled connection waiting for its next response, so it
	// doubles as the idle-connection lifetime. Kept below the inherited
	// IdleConnTimeout of 90s so that relationship stays one-way.
	//
	// The bound is per connection, so it only holds while one request occupies a
	// connection at a time. That is why newBaseTransport disables HTTP/2.
	idleRead time.Duration
}

func defaultAttemptTimeouts() attemptTimeouts {
	return attemptTimeouts{
		dial:           5 * time.Second,
		responseHeader: 10 * time.Second,
		idleRead:       60 * time.Second,
	}
}

// Dial builds a client from a comma-separated list of endpoints, in priority
// order with the primary first.
//
// name identifies which set of endpoints these are ("L1", say). It appears in
// log fields and error messages so that a process dialing more than one set can
// be read, and has no effect on behaviour.
//
// A single endpoint is dialed exactly as a plain ethclient.DialContext would,
// including non-HTTP schemes such as ws:// and IPC paths — failover is opt-in by
// configuring more than one endpoint. Two or more must all be http(s), since the
// failover transport is HTTP-only.
func Dial(ctx context.Context, name, raw string, log tmlog.Logger) (*ethclient.Client, error) {
	parts := splitEndpoints(raw)
	switch len(parts) {
	case 0:
		return nil, fmt.Errorf("%s: no rpc endpoint configured", name)
	case 1:
		return ethclient.DialContext(ctx, parts[0])
	}

	endpoints := make([]*url.URL, 0, len(parts))
	redacted := make([]string, 0, len(parts))
	manageAuth := false
	for _, part := range parts {
		u, err := url.Parse(part)
		if err != nil {
			// url.Error.Error() formats as `parse "<raw url>": ...`, so wrapping it
			// would put the very credentials redactEndpoint exists to hide into the
			// log. Report that parsing failed and nothing more.
			return nil, fmt.Errorf("%s: rpc endpoint %s is not a valid URL", name, redactEndpoint(part))
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return nil, fmt.Errorf("%s: rpc endpoint %s: failover across multiple endpoints requires http(s), got scheme %q",
				name, redactEndpoint(part), u.Scheme)
		}
		// Hostname(), not Host: "http://:8545" parses with a non-empty Host of
		// ":8545" and no host at all.
		if u.Hostname() == "" {
			return nil, fmt.Errorf("%s: rpc endpoint %s has no host", name, redactEndpoint(part))
		}
		if u.User != nil {
			manageAuth = true
		}
		endpoints = append(endpoints, u)
		redacted = append(redacted, redactEndpoint(part))
	}

	transport := &failoverTransport{
		name:          name,
		base:          newBaseTransport(defaultAttemptTimeouts()),
		endpoints:     endpoints,
		redacted:      redacted,
		manageAuth:    manageAuth,
		failbackAfter: defaultFailbackAfter,
		log:           log,
	}
	transport.lastProbe.Store(time.Now().UnixNano())
	// The URL handed to DialOptions is only a bootstrap: it selects the transport
	// by scheme and becomes the default target, which attempt() always replaces.
	// Give it a redacted one, because it is also the URL http.Client reports when
	// wrapping a transport error — and stripPassword only masks userinfo, leaving
	// an API key in the path or query (Infura /v3/, Alchemy /v2/) to travel into
	// every "failed to get L1 header" log line.
	//
	// This also means http.Client.send no longer stamps basic auth from the
	// bootstrap userinfo, which is fine: applyBasicAuth derives it per endpoint.
	rpcClient, err := rpc.DialOptions(ctx, redacted[0], rpc.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		return nil, err
	}
	log.Info("dialed rpc endpoints with failover", "target", name, "endpoints", strings.Join(redacted, ","))
	return ethclient.NewClient(rpcClient), nil
}

// splitEndpoints splits a comma-separated endpoint list and trims each entry,
// preserving order. A plain single URL yields a one-element slice, so existing
// single-endpoint configs keep working unchanged. Mirrors
// derivation.Config.BeaconRpcList.
func splitEndpoints(raw string) []string {
	var parts []string
	for _, part := range strings.Split(raw, ",") {
		if part = strings.TrimSpace(part); part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

// redactEndpoint reduces an endpoint to scheme://host for logs and for the
// errors returned to callers. Hosted providers embed credentials everywhere
// except the host — basic-auth userinfo (https://user:pass@host), an API key in
// the path (Infura /v3/ID, Alchemy /v2/KEY), or one in the query (?apikey=) —
// so echoing the raw string would ship those secrets to whatever aggregates the
// node's logs. Host and port are kept: they are not secret, and dropping the
// port would render two endpoints on the same host indistinguishable, which is
// exactly the devnet and local-node case. Mirrors
// derivation.redactBeaconEndpoint.
func redactEndpoint(raw string) string {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "<invalid-endpoint>"
	}
	return parsed.Scheme + "://" + parsed.Host
}

// newBaseTransport builds the transport each attempt runs on.
//
// HTTP/2 is switched off deliberately, and idleRead depends on it. The idle bound
// lives on the TCP connection, and under HTTP/2 one connection carries many
// concurrent streams: traffic on any of them would push the deadline forward, so
// a single stalled response stream would go unbounded again. The node polls L1
// from several goroutines at once, so that is not hypothetical. Multiplexing buys
// nothing here — these are small, infrequent requests — whereas HTTP/1.1 runs one
// request per connection at a time, which is what makes a per-connection deadline
// mean what it says. A non-nil TLSNextProto is the documented way to disable it.
func newBaseTransport(to attemptTimeouts) *http.Transport {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.ForceAttemptHTTP2 = false
	tr.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	// TLSNextProto only stops net/http from handing the connection to its HTTP/2
	// implementation; it does not stop ALPN from negotiating h2 in the first place.
	// If the TLS config still advertises h2 the server may select it, and then the
	// HTTP/1.x code reads frames as a response and reports the connection broken.
	// Pinning NextProtos means h2 can never be selected, whatever else is set here.
	tr.TLSClientConfig = &tls.Config{NextProtos: []string{"http/1.1"}}
	dialer := &net.Dialer{Timeout: to.dial, KeepAlive: 30 * time.Second}
	tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}
		return &idleReadConn{Conn: conn, idle: to.idleRead}, nil
	}
	tr.ResponseHeaderTimeout = to.responseHeader
	return tr
}

// idleReadConn fails a read that makes no progress for idle. The deadline is
// pushed forward before every read, so it measures inactivity rather than total
// time.
type idleReadConn struct {
	net.Conn
	idle time.Duration
}

func (c *idleReadConn) Read(b []byte) (int, error) {
	if err := c.Conn.SetReadDeadline(time.Now().Add(c.idle)); err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

// failoverTransport sends each JSON-RPC request to the endpoint currently in
// use, walks the remaining endpoints in ring order when that one fails, and
// every failbackAfter lets one call re-try the primary so that a recovered
// primary is re-adopted rather than waiting for the fallback to fail too.
type failoverTransport struct {
	// name labels logs and errors with which set of endpoints these are.
	name string

	base      http.RoundTripper
	endpoints []*url.URL
	redacted  []string // endpoints[i] reduced to scheme://host, for logs

	// manageAuth is true when at least one endpoint carries userinfo, which is
	// the only case where this transport takes ownership of the Authorization
	// header. When no endpoint does, a header supplied via rpc.WithHeader is
	// left untouched.
	manageAuth bool

	// cur is the index of the endpoint that last answered. Sticking to it
	// matters: without it every call would pay the dead primary's timeout again.
	cur atomic.Int32

	// failbackAfter is how long to stay on a fallback before letting one call try
	// the primary again. A field rather than a constant so tests can shrink it.
	failbackAfter time.Duration

	// lastProbe is the unix-nano time the primary was last tried, seeded at
	// construction because the process starts out using it. Compared against
	// failbackAfter to decide when one more probe is due, and advanced only by
	// dueForPrimaryProbe, with CompareAndSwap so exactly one concurrent call pays
	// for it.
	//
	// Switching endpoints deliberately does not touch it. Stamping it on a switch
	// would mean fallbacks flapping faster than failbackAfter keep pushing the
	// next probe out and the primary is never retried; the cost of not stamping is
	// that a switch occurring after the window has already elapsed is followed by
	// one immediate re-probe, which also means a primary that only blipped is
	// picked back up at once instead of after a full window.
	lastProbe atomic.Int64

	log tmlog.Logger
}

// dueForPrimaryProbe reports whether enough time has passed to spend one call
// re-trying the primary. Only the caller that wins the CompareAndSwap probes, so
// a fleet of concurrent requests still costs a single extra attempt per window.
func (t *failoverTransport) dueForPrimaryProbe(now int64) bool {
	last := t.lastProbe.Load()
	if now-last < int64(t.failbackAfter) {
		return false
	}
	return t.lastProbe.CompareAndSwap(last, now)
}

func (t *failoverTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// rpc/http.go sets req.Body but leaves req.GetBody nil, so the body must be
	// buffered here to be replayable against the next endpoint. Request bodies
	// are small (a method name and a few arguments). Response bodies, which can
	// be megabytes, are never buffered — they are handed back untouched.
	var body []byte
	if req.Body != nil {
		var err error
		body, err = io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: read rpc request body: %w", t.name, err)
		}
	}

	// prev is the endpoint in use when this call started; start is where this call
	// begins walking the ring. They differ only when a failback probe is due, in
	// which case this one call starts from the primary instead. If the primary is
	// still down the probe costs one failed attempt and the ring carries on.
	prev := int(t.cur.Load())
	start := prev
	if prev != 0 && t.dueForPrimaryProbe(time.Now().UnixNano()) {
		start = 0
	}

	var lastErr error
	for i := range t.endpoints {
		idx := (start + i) % len(t.endpoints)

		resp, err := t.attempt(req, body, idx)
		if err == nil {
			// Compared against prev, not start, so a successful failback probe moves
			// back to the primary.
			//
			// cur is an approximation on purpose — calls do not coordinate, and no
			// scheme for agreeing on one endpoint is worth its cost here, because a
			// wrong cur only makes the next call spend one failed attempt before
			// moving on. Two consequences worth knowing rather than fixing: a call
			// that began before someone else switched can write its older index over
			// the newer one, and a call whose endpoint equals the one it started from
			// records nothing, so a working primary is not re-adopted here — the
			// failback probe is what does that.
			//
			// cause is nil when this was a failback probe, since nothing failed on the
			// way here. Recovery is therefore logged at Error with a nil cause; the
			// from/to pair is what distinguishes it from a degradation.
			if idx != prev {
				t.cur.Store(int32(idx))
				t.log.Error("switched rpc endpoint", "target", t.name,
					"from", t.redacted[prev], "to", t.redacted[idx], "cause", lastErr)
			}
			return resp, nil
		}
		lastErr = err
		t.log.Debug("rpc endpoint attempt failed", "target", t.name, "endpoint", t.redacted[idx], "err", err)

		// No budget left for another endpoint: the caller either cancelled or ran
		// out of time. Report that rather than an "all endpoints failed" error
		// assembled from attempts that never left the process.
		if ctxErr := req.Context().Err(); ctxErr != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("%s: all %d rpc endpoints failed, last error: %w", t.name, len(t.endpoints), lastErr)
}

// attempt issues the request against endpoints[idx]. A response that is usable
// by the rpc layer is returned as-is; anything else becomes an error so the
// caller can move on.
func (t *failoverTransport) attempt(req *http.Request, body []byte, idx int) (*http.Response, error) {
	target := *t.endpoints[idx] // copy: a RoundTripper must not mutate shared state
	attempt := req.Clone(req.Context())
	attempt.URL = &target
	attempt.Host = "" // let the Host header follow the new URL
	attempt.Body = io.NopCloser(bytes.NewReader(body))
	attempt.ContentLength = int64(len(body))
	// GetBody lets net/http re-send this request itself when the connection it
	// took from the idle pool turns out to be dead. Its shouldRetryRequest only
	// does that for a body it can rewind — the nothingWrittenError branch returns
	// `outgoingLength() == 0 || GetBody != nil`, and a POST with a body fails both
	// unless this is set. Leaving it nil surfaces a routine keep-alive race as an
	// endpoint failure, and the sticky switch then demotes a healthy primary for
	// the rest of the process's life.
	attempt.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}
	if t.manageAuth {
		applyBasicAuth(attempt, &target)
	}

	resp, err := t.base.RoundTrip(attempt)
	if err != nil {
		return nil, err
	}
	if shouldFailover(resp.StatusCode) {
		resp.Body.Close()
		return nil, fmt.Errorf("http %s", resp.Status)
	}
	if err := requireJSONBody(resp); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return resp, nil
}

// peekLimit is how far requireJSONBody looks for the first non-whitespace byte.
// A JSON-RPC reply has none, so anything beyond a few bytes means the response
// is not one and the window only needs to be generous, not large.
const peekLimit = 8

// requireJSONBody reports an endpoint as failed when its body is not JSON. It
// peeks at the first few bytes and hands them back, so the rpc layer still
// decodes the whole response — full bodies are never buffered, since a single
// eth_getLogs reply can be megabytes.
func requireJSONBody(resp *http.Response) error {
	// Only a success is expected to carry a JSON-RPC reply. The only statuses that
	// reach here otherwise are the 4xx shouldFailover deliberately passes through,
	// which describe the request rather than the endpoint and promise nothing about
	// the body; judging them here would undo that decision.
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil
	}

	// A compressed body cannot be inspected without decompressing it, which is
	// not worth paying for on every call. Go's transport strips Content-Encoding
	// when it decompresses itself, so this only applies when a caller set
	// Accept-Encoding explicitly via rpc.WithHeader.
	if resp.Header.Get("Content-Encoding") != "" {
		return nil
	}

	peeked := make([]byte, peekLimit)
	n, readErr := io.ReadFull(resp.Body, peeked)
	peeked = peeked[:n]
	atEOF := errors.Is(readErr, io.EOF) || errors.Is(readErr, io.ErrUnexpectedEOF)

	// Restore the body before returning, on every path: the caller closes it even
	// when we reject the response.
	resp.Body = replayBody(peeked, resp.Body)

	if readErr != nil && !atEOF {
		return fmt.Errorf("read response body: %w", readErr)
	}

	switch trimmed := bytes.TrimLeft(peeked, " \t\r\n"); {
	case len(trimmed) > 0:
		if trimmed[0] != '{' && trimmed[0] != '[' {
			return fmt.Errorf("endpoint did not return JSON (body starts with %q, content-type %q)",
				trimmed[0], resp.Header.Get("Content-Type"))
		}
	case atEOF:
		// Nothing but optional whitespace in the whole body.
		return errors.New("endpoint returned an empty body")
	default:
		// The peek window held only whitespace and there is more to come. Too
		// unusual to judge, and rejecting a working endpoint is the worse error.
	}
	return nil
}

// replayBody yields the peeked bytes followed by the rest of the original body,
// and closes the original on Close.
func replayBody(peeked []byte, rest io.ReadCloser) io.ReadCloser {
	return struct {
		io.Reader
		io.Closer
	}{io.MultiReader(bytes.NewReader(peeked), rest), rest}
}

// shouldFailover reports whether an HTTP status justifies trying the next
// endpoint. 5xx and 429 are the endpoint's problem, and 408 means it gave up on
// us.
//
// 401 and 403 count too, because endpoints can carry their own credentials (see
// applyBasicAuth): an expired key or an exhausted quota on one provider says
// nothing about the next one, and hosted providers use exactly these statuses
// for it. Refusing to move would blind the node while a correctly configured
// fallback sat idle. A genuine misconfiguration is still visible — every switch
// is logged with its cause, and if all endpoints reject us the aggregate error
// surfaces.
//
// 3xx counts as well. A JSON-RPC endpoint has no business redirecting, and
// following one is not possible here: attempt() rewrites every request's URL to
// its endpoint, so the redirect target would be discarded and the same endpoint
// asked again until http.Client gives up after ten hops — ten wasted round trips
// and no failover. An endpoint configured as http:// behind a server that
// redirects to https:// is the realistic way to hit this.
//
// 400 and 404 are excluded: those describe the request itself, so the next
// endpoint would reproduce them and moving on would only multiply the load. A
// JSON-RPC error arrives as 200 and is not examined here.
func shouldFailover(status int) bool {
	return status == http.StatusRequestTimeout ||
		status == http.StatusTooManyRequests ||
		status == http.StatusUnauthorized ||
		status == http.StatusForbidden ||
		(status >= 300 && status < 400) ||
		status >= 500
}

// applyBasicAuth re-derives the Authorization header for the endpoint actually
// being dialed. http.Client.send fills it in from the *original* URL's userinfo
// before RoundTrip runs, so after swapping the URL the header would otherwise
// carry the primary endpoint's credentials to a different provider.
func applyBasicAuth(req *http.Request, target *url.URL) {
	if target.User == nil {
		req.Header.Del("Authorization")
		return
	}
	password, _ := target.User.Password()
	req.SetBasicAuth(target.User.Username(), password)
}
