package node

import (
	"context"
	"fmt"

	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"morph-l2/node/db"
	"morph-l2/node/sync"
)

// NewSyncer is a self-contained helper that builds a Syncer from a CLI
// context (store + sync config + L1 client). main.go does NOT use this —
// it shares its own l1Client across all L1-touching components. This
// helper exists for legacy/dev-only paths and dials its own client to
// keep the API simple; if you need to share the client, pass it down
// directly to sync.NewSyncer instead of going through this wrapper.
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
	l1Client, err := ethclient.Dial(syncConfig.L1.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial l1 node, error: %v", err)
	}
	syncer, err := sync.NewSyncer(context.Background(), store, syncConfig, config.Logger, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create syncer, error: %v", err)
	}
	return syncer, nil
}
