package immutables

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

type Config struct {
	L2SequencerAddresses []common.Address `json:"l2SequencerAddresses"`
	L2SequencerTmKeys    []common.Hash    `json:"l2SequencerTmKeys"`
	L2SequencerBlsKeys   [][]byte         `json:"l2SequencerBlsKeys"`

	L2GenesisBlockTimestamp hexutil.Uint64 `json:"l2GenesisBlockTimestamp"`
}
