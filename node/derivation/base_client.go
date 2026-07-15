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

// FallbackHTTPClient wraps several beacon HTTP endpoints and tries them in the
// configured order. It returns the first endpoint that answers with a 200; on
// a transport error or a non-200 it records the failing endpoint in metrics and
// moves on to the next. This keeps derivation making progress when a beacon
// temporarily fails to serve blob sidecars, and surfaces the flaky node via the
// beacon_request_failure_total metric. With a single endpoint it behaves like a
// BasicHTTPClient.
type FallbackHTTPClient struct {
	clients []*BasicHTTPClient
	metrics *Metrics
}

func NewFallbackHTTPClient(endpoints []string, log tmlog.Logger, metrics *Metrics) *FallbackHTTPClient {
	clients := make([]*BasicHTTPClient, 0, len(endpoints))
	for _, endpoint := range endpoints {
		clients = append(clients, NewBasicHTTPClient(endpoint, log))
	}
	return &FallbackHTTPClient{clients: clients, metrics: metrics}
}

func (cl *FallbackHTTPClient) Get(ctx context.Context, p string, headers http.Header) (*http.Response, error) {
	var lastErr error
	for _, client := range cl.clients {
		resp, err := client.Get(ctx, p, headers)
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("beacon endpoint %s returned status %d", client.endpoint, resp.StatusCode)
			_ = resp.Body.Close()
		}
		if cl.metrics != nil {
			cl.metrics.IncBeaconRequestFailure(client.endpoint)
		}
	}
	return nil, lastErr
}
