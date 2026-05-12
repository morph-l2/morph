package types

import (
	"fmt"

	"morph-l2/common/batch"
	"morph-l2/tx-submitter/iface"
)

// L2Caller reads L2 gov / sequencer state for batch assembly (see batch.L2Gov).
type L2Caller = batch.L2Gov

// NewL2Caller builds an L2Caller backed by the given L2 RPC clients.
func NewL2Caller(l2Clients []iface.L2Client) (*batch.L2Gov, error) {
	if len(l2Clients) == 0 {
		return nil, fmt.Errorf("no l2clients provided")
	}
	for _, l2Client := range l2Clients {
		if l2Client == nil {
			return nil, fmt.Errorf("nil l2client")
		}
	}
	return batch.NewL2Gov(&iface.L2Clients{Clients: l2Clients})
}
