package batch

import (
	"bytes"
	"fmt"
	"math/big"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto"
)

// L2Gov bundles read-only L2 contracts used when assembling rollup batches.
type L2Gov struct {
	sequencerContract       *bindings.SequencerCaller
	l2MessagePasserContract *bindings.L2ToL1MessagePasserCaller
	govContract             *bindings.GovCaller
}

// NewL2Gov builds an L2Gov using any ContractCaller (e.g. *ethclient.Client or a multi-client backend).
func NewL2Gov(backend bind.ContractCaller) (*L2Gov, error) {
	if backend == nil {
		return nil, fmt.Errorf("nil contract backend")
	}
	sequencerContract, err := bindings.NewSequencerCaller(predeploys.SequencerAddr, backend)
	if err != nil {
		return nil, err
	}
	l2MessagePasserContract, err := bindings.NewL2ToL1MessagePasserCaller(predeploys.L2ToL1MessagePasserAddr, backend)
	if err != nil {
		return nil, err
	}
	govContract, err := bindings.NewGovCaller(predeploys.GovAddr, backend)
	if err != nil {
		return nil, err
	}
	return &L2Gov{
		sequencerContract:       sequencerContract,
		l2MessagePasserContract: l2MessagePasserContract,
		govContract:             govContract,
	}, nil
}

// SequencerSetVerifyHash gets the sequencer set verify hash from the Sequencer contract.
func (c *L2Gov) SequencerSetVerifyHash(opts *bind.CallOpts) ([32]byte, error) {
	return c.sequencerContract.SequencerSetVerifyHash(opts)
}

// GetTreeRoot gets the tree root from the L2ToL1MessagePasser contract.
func (c *L2Gov) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	return c.l2MessagePasserContract.GetTreeRoot(opts)
}

// BatchBlockInterval gets the batch block interval from the Gov contract.
func (c *L2Gov) BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	return c.govContract.BatchBlockInterval(opts)
}

// BatchTimeout gets the batch timeout from the Gov contract.
func (c *L2Gov) BatchTimeout(opts *bind.CallOpts) (*big.Int, error) {
	return c.govContract.BatchTimeout(opts)
}

// GetSequencerSetBytes returns sequencer set bytes after hash consistency check.
func (c *L2Gov) GetSequencerSetBytes(opts *bind.CallOpts) ([]byte, common.Hash, error) {
	hash, err := c.sequencerContract.SequencerSetVerifyHash(opts)
	if err != nil {
		return nil, common.Hash{}, err
	}
	setBytes, err := c.sequencerContract.GetSequencerSetBytes(opts)
	if err != nil {
		return nil, common.Hash{}, err
	}
	if bytes.Equal(hash[:], crypto.Keccak256Hash(setBytes).Bytes()) {
		return setBytes, hash, nil
	}
	return nil, common.Hash{}, fmt.Errorf("sequencer set hash verify failed %v: %v", hexutil.Encode(setBytes), common.BytesToHash(hash[:]).String())
}
