package rpc

import (
	"context"

	"github.com/tendermint/tendermint/libs/log"
)

// APIBackend implements API, delegating to a ConsensusAdapter.
type APIBackend struct {
	log  log.Logger
	cons ConsensusAdapter
}

// NewAPIBackend creates a new APIBackend.
func NewAPIBackend(log log.Logger, cons ConsensusAdapter) *APIBackend {
	return &APIBackend{log: log, cons: cons}
}

var _ API = (*APIBackend)(nil)

func (api *APIBackend) Leader(ctx context.Context) (bool, error) {
	return api.cons.Leader(), nil
}

func (api *APIBackend) LeaderWithID(ctx context.Context) (*ServerInfo, error) {
	return api.cons.LeaderWithID(), nil
}

func (api *APIBackend) AddServerAsVoter(ctx context.Context, id string, addr string, version uint64) error {
	return api.cons.AddVoter(id, addr, version)
}

func (api *APIBackend) AddServerAsNonvoter(ctx context.Context, id string, addr string, version uint64) error {
	return api.cons.AddNonVoter(id, addr, version)
}

func (api *APIBackend) RemoveServer(ctx context.Context, id string, version uint64) error {
	return api.cons.RemoveServer(id, version)
}

func (api *APIBackend) TransferLeader(ctx context.Context) error {
	return api.cons.TransferLeader()
}

func (api *APIBackend) TransferLeaderToServer(ctx context.Context, id string, addr string) error {
	return api.cons.TransferLeaderTo(id, addr)
}

func (api *APIBackend) ClusterMembership(ctx context.Context) (*ClusterMembership, error) {
	return api.cons.ClusterMembership()
}
