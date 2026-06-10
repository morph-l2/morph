package rpc

import "context"

// API defines the interface for the hakeeper management RPC API.
type API interface {
	// Leader returns true if the server is the leader.
	Leader(ctx context.Context) (bool, error)
	// LeaderWithID returns the current leader's server info.
	LeaderWithID(ctx context.Context) (*ServerInfo, error)
	// AddServerAsVoter adds a server as a voter to the cluster.
	AddServerAsVoter(ctx context.Context, id string, addr string, version uint64) error
	// AddServerAsNonvoter adds a server as a non-voter to the cluster.
	AddServerAsNonvoter(ctx context.Context, id string, addr string, version uint64) error
	// RemoveServer removes a server from the cluster.
	RemoveServer(ctx context.Context, id string, version uint64) error
	// TransferLeader transfers leadership to another server.
	TransferLeader(ctx context.Context) error
	// TransferLeaderToServer transfers leadership to a specific server.
	TransferLeaderToServer(ctx context.Context, id string, addr string) error
	// ClusterMembership returns the current cluster membership configuration.
	ClusterMembership(ctx context.Context) (*ClusterMembership, error)
}
