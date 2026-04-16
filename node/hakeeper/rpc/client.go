package rpc

import (
	"context"

	ethrpc "github.com/morph-l2/go-ethereum/rpc"
)

// RPCNamespace is the JSON-RPC namespace for the HA management API.
var RPCNamespace = "ha"

// APIClient provides an RPC client for calling hakeeper API methods.
type APIClient struct {
	c *ethrpc.Client
}

var _ API = (*APIClient)(nil)

// NewAPIClient creates a new APIClient wrapping a go-ethereum rpc.Client.
func NewAPIClient(c *ethrpc.Client) *APIClient {
	return &APIClient{c: c}
}

// DialAPIClient dials a hakeeper RPC server at the given address and returns
// an APIClient. token is sent as the Authorization header on every request;
// pass empty string if the server has no auth configured.
// The caller is responsible for calling Close() when done.
func DialAPIClient(ctx context.Context, addr string, token string) (*APIClient, error) {
	c, err := ethrpc.DialContext(ctx, "http://"+addr)
	if err != nil {
		return nil, err
	}
	if token != "" {
		c.SetHeader("Authorization", token)
	}
	return NewAPIClient(c), nil
}

func prefixRPC(method string) string {
	return RPCNamespace + "_" + method
}

// Close closes the underlying RPC client.
func (c *APIClient) Close() {
	c.c.Close()
}

func (c *APIClient) Leader(ctx context.Context) (bool, error) {
	var leader bool
	err := c.c.CallContext(ctx, &leader, prefixRPC("leader"))
	return leader, err
}

func (c *APIClient) LeaderWithID(ctx context.Context) (*ServerInfo, error) {
	var info *ServerInfo
	err := c.c.CallContext(ctx, &info, prefixRPC("leaderWithID"))
	return info, err
}

func (c *APIClient) AddServerAsVoter(ctx context.Context, id string, addr string, version uint64) error {
	return c.c.CallContext(ctx, nil, prefixRPC("addServerAsVoter"), id, addr, version)
}

func (c *APIClient) AddServerAsNonvoter(ctx context.Context, id string, addr string, version uint64) error {
	return c.c.CallContext(ctx, nil, prefixRPC("addServerAsNonvoter"), id, addr, version)
}

func (c *APIClient) RemoveServer(ctx context.Context, id string, version uint64) error {
	return c.c.CallContext(ctx, nil, prefixRPC("removeServer"), id, version)
}

func (c *APIClient) TransferLeader(ctx context.Context) error {
	return c.c.CallContext(ctx, nil, prefixRPC("transferLeader"))
}

func (c *APIClient) TransferLeaderToServer(ctx context.Context, id string, addr string) error {
	return c.c.CallContext(ctx, nil, prefixRPC("transferLeaderToServer"), id, addr)
}

func (c *APIClient) ClusterMembership(ctx context.Context) (*ClusterMembership, error) {
	var membership ClusterMembership
	err := c.c.CallContext(ctx, &membership, prefixRPC("clusterMembership"))
	return &membership, err
}
