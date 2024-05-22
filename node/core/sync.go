package node

import (
	"context"
	"fmt"

	"github.com/urfave/cli"

	"morph-l2/node/db"
	"morph-l2/node/sync"
)

func NewSyncer(ctx *cli.Context, home string, config *Config) (*sync.Syncer, error) {
	// configure store
	dbConfig := db.DefaultConfig()
	dbConfig.SetCliContext(ctx)
	store, err := db.NewStore(dbConfig, home)
	if err != nil {
		return nil, err
	}
	// launch syncer
	syncConfig := sync.DefaultConfig()
	if err = syncConfig.SetCliContext(ctx); err != nil {
		return nil, err
	}
	syncer, err := sync.NewSyncer(context.Background(), store, syncConfig, config.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create syncer, error: %v", err)
	}
	return syncer, nil
}
