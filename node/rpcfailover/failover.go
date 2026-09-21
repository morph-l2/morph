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

const (
	// dialTimeout bounds TCP+TLS setup per attempt, so an endpoint whose host is
	// blackholed cannot consume the caller's entire deadline.
	dialTimeout = 5 * time.Second

	// responseHeaderTimeout bounds the wait for response headers per attempt.
	// This is what makes failover work against an endpoint that accepts the
	// connection and then never answers — the common shape of a hung RPC
	// provider. Because the base transport enforces it per RoundTrip, each
	// attempt gets its own budget without wrapping the caller's context.
	//
	// It does not cover a server that sends headers and then stalls the body;
	// that case remains bounded only by the caller's context.
	responseHeaderTimeout = 10 * time.Second
)

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
			return nil, fmt.Errorf("%s: invalid rpc endpoint %s: %w", name, redactEndpoint(part), err)
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return nil, fmt.Errorf("%s: rpc endpoint %s: failover across multiple endpoints requires http(s), got scheme %q",
				name, redactEndpoint(part), u.Scheme)
		}
		if u.User != nil {
			manageAuth = true
		}
		endpoints = append(endpoints, u)
		redacted = append(redacted, redactEndpoint(part))
	}

	transport := &failoverTransport{
		name:       name,
		base:       newBaseTransport(),
		endpoints:  endpoints,
		redacted:   redacted,
		manageAuth: manageAuth,
		log:        log,
	}
	rpcClient, err := rpc.DialOptions(ctx, parts[0], rpc.WithHTTPClient(&http.Client{Transport: transport}))
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

func newBaseTransport() *http.Transport {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DialContext = (&net.Dialer{Timeout: dialTimeout, KeepAlive: 30 * time.Second}).DialContext
	tr.ResponseHeaderTimeout = responseHeaderTimeout
	return tr
}

// failoverTransport sends each JSON-RPC request to the endpoint it is currently
// stuck to, and walks the remaining endpoints in ring order when that one fails.
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
	// Nothing ever moves it back on its own — a recovered primary is picked up
	// when the current endpoint fails, or on the next process restart.
	cur atomic.Int32

	log tmlog.Logger
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

	start := int(t.cur.Load())
	var lastErr error
	for i := range t.endpoints {
		idx := (start + i) % len(t.endpoints)

		resp, err := t.attempt(req, body, idx)
		if err == nil {
			if idx != start {
				t.cur.Store(int32(idx))
				t.log.Error("switched rpc endpoint", "target", t.name,
					"from", t.redacted[start], "to", t.redacted[idx], "cause", lastErr)
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
	// Only a success is expected to carry a JSON-RPC reply. The statuses that
	// reach here otherwise are the 3xx that http.Client will follow and the 4xx
	// that shouldFailover deliberately passes through; neither promises a JSON
	// body, and judging them here would undo that decision.
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
// us. The remaining 4xx are deliberately excluded: a 400/401/403/404 describes
// the request or the credentials, the next endpoint would almost certainly
// reproduce it, and moving on silently would turn a misconfiguration into a
// mystery. A JSON-RPC error arrives as 200 and is not examined here.
func shouldFailover(status int) bool {
	return status == http.StatusRequestTimeout ||
		status == http.StatusTooManyRequests ||
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
