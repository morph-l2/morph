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
