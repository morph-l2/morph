package batch

import (
	"fmt"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
)

// L2Gov bundles read-only L2 contracts used when assembling rollup batches.
// The name is retained for API compatibility, but batching no longer reads the
// L2 Sequencer or Gov contracts. Only the withdrawal tree root remains an L2
// contract dependency.
type L2Gov struct {
	l2MessagePasserContract *bindings.L2ToL1MessagePasserCaller
}

// NewL2Gov builds an L2Gov using any ContractCaller (e.g. *ethclient.Client or a multi-client backend).
func NewL2Gov(backend bind.ContractCaller) (*L2Gov, error) {
	if backend == nil {
		return nil, fmt.Errorf("nil contract backend")
	}
	l2MessagePasserContract, err := bindings.NewL2ToL1MessagePasserCaller(predeploys.L2ToL1MessagePasserAddr, backend)
	if err != nil {
		return nil, err
	}
	return &L2Gov{
		l2MessagePasserContract: l2MessagePasserContract,
	}, nil
}

// GetTreeRoot gets the tree root from the L2ToL1MessagePasser contract.
func (c *L2Gov) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	return c.l2MessagePasserContract.GetTreeRoot(opts)
}
