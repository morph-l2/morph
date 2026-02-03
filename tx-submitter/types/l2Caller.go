package types

import (
	"math/big"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
)

type L2Caller struct {
	l2Clients               *iface.L2Clients
	sequencerContract       *bindings.SequencerCaller
	l2MessagePasserContract *bindings.L2ToL1MessagePasserCaller
	govContract             *bindings.GovCaller
}

func NewL2Caller(l2Clients []iface.L2Client) (*L2Caller, error) {
	clients := &iface.L2Clients{Clients: l2Clients}

	// Initialize Sequencer contract
	sequencerContract, err := bindings.NewSequencerCaller(predeploys.SequencerAddr, clients)
	if err != nil {
		return nil, err
	}

	// Initialize L2ToL1MessagePasser contract
	l2MessagePasserContract, err := bindings.NewL2ToL1MessagePasserCaller(predeploys.L2ToL1MessagePasserAddr, clients)
	if err != nil {
		return nil, err
	}

	// Initialize Gov contract
	govContract, err := bindings.NewGovCaller(predeploys.GovAddr, clients)
	if err != nil {
		return nil, err
	}

	return &L2Caller{
		l2Clients:               clients,
		sequencerContract:       sequencerContract,
		l2MessagePasserContract: l2MessagePasserContract,
		govContract:             govContract,
	}, nil
}

// SequencerSetVerifyHash gets the sequencer set verify hash from the Sequencer contract
func (c *L2Caller) SequencerSetVerifyHash(opts *bind.CallOpts) ([32]byte, error) {
	return c.sequencerContract.SequencerSetVerifyHash(opts)
}

// GetTreeRoot gets the tree root from the L2ToL1MessagePasser contract
func (c *L2Caller) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	return c.l2MessagePasserContract.GetTreeRoot(opts)
}

// BatchBlockInterval gets the batch block interval from the Gov contract
func (c *L2Caller) BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	return c.govContract.BatchBlockInterval(opts)
}

// BatchTimeout gets the batch timeout from the Gov contract
func (c *L2Caller) BatchTimeout(opts *bind.CallOpts) (*big.Int, error) {
	return c.govContract.BatchTimeout(opts)
}
