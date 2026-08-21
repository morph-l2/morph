package batch

import (
	"bytes"
	"fmt"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto"
)

// L2Gov provides the historical L2 reads used when assembling rollup batches.
type L2Gov struct {
	sequencerContract       *bindings.SequencerCaller
	l2MessagePasserContract *bindings.L2ToL1MessagePasserCaller
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
	return &L2Gov{
		sequencerContract:       sequencerContract,
		l2MessagePasserContract: l2MessagePasserContract,
	}, nil
}

// GetTreeRoot gets the tree root from the L2ToL1MessagePasser contract.
func (c *L2Gov) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	return c.l2MessagePasserContract.GetTreeRoot(opts)
}

// GetSequencerSetBytes returns the historical sequencer-set encoding and hash
// at opts.BlockNumber. The two contract reads are checked against each other so
// replay never accepts bytes from a different L2 state.
func (c *L2Gov) GetSequencerSetBytes(opts *bind.CallOpts) ([]byte, common.Hash, error) {
	hash, err := c.sequencerContract.SequencerSetVerifyHash(opts)
	if err != nil {
		return nil, common.Hash{}, err
	}
	setBytes, err := c.sequencerContract.GetSequencerSetBytes(opts)
	if err != nil {
		return nil, common.Hash{}, err
	}
	calculated := crypto.Keccak256Hash(setBytes)
	if !bytes.Equal(hash[:], calculated[:]) {
		return nil, common.Hash{}, fmt.Errorf(
			"sequencer set hash verify failed: bytes=%s contract=%s calculated=%s",
			hexutil.Encode(setBytes), common.Hash(hash), calculated,
		)
	}
	return setBytes, common.Hash(hash), nil
}
