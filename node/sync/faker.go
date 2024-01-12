package sync

import tmlog "github.com/tendermint/tendermint/libs/log"

func NewFakeSyncer(db Database) *Syncer {
	return &Syncer{
		db:     db,
		isFake: true,
		logger: tmlog.NewNopLogger(),
	}
}
