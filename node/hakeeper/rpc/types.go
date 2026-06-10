package rpc

// ServerSuffrage determines whether a Server in a Configuration gets a vote.
type ServerSuffrage int

const (
	// Nonvoter receives log entries but is not considered for elections.
	// Zero value — safer default (no voting rights).
	Nonvoter ServerSuffrage = iota
	// Voter is a server whose vote is counted in elections.
	Voter
)

func (s ServerSuffrage) String() string {
	switch s {
	case Voter:
		return "Voter"
	case Nonvoter:
		return "Nonvoter"
	}
	return "ServerSuffrage"
}

// ClusterMembership is a versioned list of servers in the Raft cluster.
type ClusterMembership struct {
	Servers []ServerInfo `json:"servers"`
	Version uint64       `json:"version"`
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
