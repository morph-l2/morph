package types

import (
	"errors"
	"github.com/morph-l2/go-ethereum/common"
	"math/big"
)

var (
	MaxEpochCount = 50
)

var ErrMemoryDBNotFound = errors.New("not found")

const (
	L1ChangePoint = iota + 1
	L2ChangePoint
)

type SequencerSetUpdateEpoch struct {
	Submitters []common.Address
	StartTime  *big.Int
	EndTime    *big.Int
	EndBlock   *big.Int
}

// TODO
type ChangePoint struct {
	TimeStamp     uint64
	BlockNumber   uint64
	EpochInterval uint64
	Submitters    []common.Address
	ChangeType    uint64
}

type L2Sequencer struct {
	TimeStamp     uint64
	EpochInterval uint64
	Addresses     []common.Address
}

type ActiveStakers struct {
	TimeStamp   uint64
	BlockNumber uint64
	Addresses   []common.Address
}

type ChangeContext struct {
	L2Sequencers        []L2Sequencer
	ActiveStakersByTime []ActiveStakers
	ChangePoints        []ChangePoint
	L1Synced            uint64
	L2synced            uint64
}
