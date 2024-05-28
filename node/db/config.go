package db

import (
	"github.com/urfave/cli"

	"morph-l2/node/flags"
)

type Config struct {
	DBPath          string `json:"db_path"`
	Namespace       string `json:"namespace"`
	DatabaseHandles int    `json:"database_handles"`
	DatabaseCache   int    `json:"database_cache"`
	DatabaseFreezer string `json:"database_freezer"`
}

func DefaultConfig() *Config {
	return &Config{
		Namespace:       "node",
		DatabaseHandles: 256,
		DatabaseCache:   256,
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) {
	if ctx.GlobalIsSet(flags.DBDataDir.Name) {
		c.DBPath = ctx.GlobalString(flags.DBDataDir.Name)
	}
	if ctx.GlobalIsSet(flags.DBNamespace.Name) {
		c.Namespace = ctx.GlobalString(flags.DBNamespace.Name)
	}
	if ctx.GlobalIsSet(flags.DBHandles.Name) {
		c.DatabaseHandles = ctx.GlobalInt(flags.DBHandles.Name)
	}
	if ctx.GlobalIsSet(flags.DBCache.Name) {
		c.DatabaseCache = ctx.GlobalInt(flags.DBCache.Name)
	}
	if ctx.GlobalIsSet(flags.DBFreezer.Name) {
		c.DatabaseFreezer = ctx.GlobalString(flags.DBFreezer.Name)
	}
}
