package db

import (
	"github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/store"
)

type TmDB struct {
	BlockStore *store.BlockStore
	StateStore state.Store
}

func NewTmDB(blockStore *store.BlockStore, stateStore state.Store) *TmDB {
	return &TmDB{
		BlockStore: blockStore,
		StateStore: stateStore,
	}
}
