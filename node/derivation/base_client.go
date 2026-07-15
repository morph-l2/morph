package derivation

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	tmlog "github.com/tendermint/tendermint/libs/log"
)

const (
	DefaultTimeoutSeconds = 30
)

var _ HTTP = (*BasicHTTPClient)(nil)

type HTTP interface {
	Get(ctx context.Context, path string, headers http.Header) (*http.Response, error)
}

type BasicHTTPClient struct {
	endpoint string
	log      tmlog.Logger
	client   *http.Client
}

func NewBasicHTTPClient(endpoint string, log tmlog.Logger) *BasicHTTPClient {
	// Make sure the endpoint ends in trailing slash
	trimmedEndpoint := strings.TrimSuffix(endpoint, "/") + "/"
	return &BasicHTTPClient{
		endpoint: trimmedEndpoint,
		log:      log,
		client:   &http.Client{Timeout: DefaultTimeoutSeconds * time.Second},
	}
}

func (cl *BasicHTTPClient) Get(ctx context.Context, p string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cl.endpoint+p, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to construct request", err)
	}
	for k, values := range headers {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	return cl.client.Do(req)
}

var _ HTTP = (*FallbackHTTPClient)(nil)

// FallbackHTTPClient wraps several beacon HTTP endpoints and queries them in
// order. When one endpoint fails with a transport error or a non-200 response,
// it transparently falls back to the next. This lets derivation keep making
// progress when the primary beacon node temporarily fails to serve blob
// sidecars (a recurring issue in production), as long as one of the configured
// beacons can serve the request.
//
// The first endpoint returning a 200 response wins. If every endpoint fails,
// the last non-200 response is returned (so the caller produces its usual
// "failed request with status" error), or the last transport error when no
// response was obtained at all. Behaviour with a single endpoint is identical
// to using a BasicHTTPClient directly.
type FallbackHTTPClient struct {
	clients []*BasicHTTPClient
	log     tmlog.Logger
}

func NewFallbackHTTPClient(endpoints []string, log tmlog.Logger) *FallbackHTTPClient {
	clients := make([]*BasicHTTPClient, 0, len(endpoints))
	for _, endpoint := range endpoints {
		clients = append(clients, NewBasicHTTPClient(endpoint, log))
	}
	return &FallbackHTTPClient{clients: clients, log: log}
}

func (cl *FallbackHTTPClient) Get(ctx context.Context, p string, headers http.Header) (*http.Response, error) {
	var (
		lastResp *http.Response
		lastErr  error
	)
	for i, client := range cl.clients {
		// Stop early on cancellation/deadline instead of hammering every
		// fallback endpoint with a request that is already doomed.
		if err := ctx.Err(); err != nil {
			if lastErr == nil {
				lastErr = err
			}
			break
		}
		resp, err := client.Get(ctx, p, headers)
		if err != nil {
			lastErr = err
			if cl.log != nil {
				cl.log.Debug("beacon endpoint request failed, trying fallback",
					"index", i, "endpoint", client.endpoint, "err", err)
			}
			continue
		}
		if resp.StatusCode == http.StatusOK {
			return resp, nil
		}
		// Non-200: keep it as the representative failure and fall through to
		// the next endpoint. Close any previously retained non-200 body so the
		// underlying connection can be reused.
		lastErr = fmt.Errorf("beacon endpoint %s returned status %d", client.endpoint, resp.StatusCode)
		if lastResp != nil {
			_ = lastResp.Body.Close()
		}
		lastResp = resp
		if cl.log != nil {
			cl.log.Debug("beacon endpoint returned non-200, trying fallback",
				"index", i, "endpoint", client.endpoint, "status", resp.StatusCode)
		}
	}
	if lastResp != nil {
		return lastResp, nil
	}
	return nil, lastErr
}
