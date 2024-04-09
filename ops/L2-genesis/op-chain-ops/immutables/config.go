package immutables

import (
	"github.com/scroll-tech/go-ethereum/common"
)

type Config struct {
	L2SequencerAddresses []common.Address `json:"l2SequencerAddresses"`
	L2SequencerTmKeys    []common.Hash    `json:"l2SequencerTmKeys"`
	L2SequencerBlsKeys   [][]byte         `json:"l2SequencerBlsKeys"`
}
