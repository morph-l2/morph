package rpc

// ServerSuffrage determines whether a Server in a Configuration gets a vote.
type ServerSuffrage int

// These values must match hashicorp/raft's ServerSuffrage exactly: cluster
// membership is read from raft and cast by raw integer value (see
// HAService.ClusterMembership), so any divergence silently mislabels roles.
const (
	// Voter is a server whose vote is counted in elections.
	Voter ServerSuffrage = iota
	// Nonvoter receives log entries but is not considered for elections.
	Nonvoter
	// Staging is a server that acts like a Nonvoter while it catches up, then is
	// promoted to Voter. Present for parity with raft's enum.
	Staging
)

func (s ServerSuffrage) String() string {
	switch s {
	case Voter:
		return "Voter"
	case Nonvoter:
		return "Nonvoter"
	case Staging:
		return "Staging"
	}
	return "ServerSuffrage"
}

// ClusterMembership is a versioned list of servers in the Raft cluster.
type ClusterMembership struct {
	Servers []ServerInfo `json:"servers"`
	// LeaderID is the ID of the current leader, matching one of Servers[i].ID.
	// Empty when no leader is currently known (e.g. during an election).
	LeaderID string `json:"leaderId"`
	Version  uint64 `json:"version"`
}

// ServerInfo describes a single Raft cluster member.
type ServerInfo struct {
	ID       string         `json:"id"`
	Addr     string         `json:"addr"`
	Suffrage ServerSuffrage `json:"suffrage"`
}

// ConsensusAdapter is the interface the RPC backend requires.
// It is implemented directly by HAService in ha_service.go.
type ConsensusAdapter interface {
	Leader() bool
	LeaderWithID() *ServerInfo
	AddVoter(id, addr string, version uint64) error
	AddNonVoter(id, addr string, version uint64) error
	DemoteVoter(id string, version uint64) error
	RemoveServer(id string, version uint64) error
	TransferLeader() error
	TransferLeaderTo(id, addr string) error
	ClusterMembership() (*ClusterMembership, error)
	ServerID() string
	Addr() string
}
