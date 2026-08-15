package batch

import (
	"fmt"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
)

// L2Gov provides the withdrawal-root read used when assembling rollup batches.
// The name is kept to avoid a broad caller rename during this cleanup.
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
