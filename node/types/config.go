package types

import (
	"github.com/morph-l2/go-ethereum/rpc"
)

var DefaultHomeDir = ".morphnode"

type L1Config struct {
	Addr          string          `json:"addr"`
	Confirmations rpc.BlockNumber `json:"confirmations"`
}

type L2Config struct {
	EthAddr    string   `json:"eth"`
	EngineAddr string   `json:"engine"`
	JwtSecret  [32]byte `json:"jwt_secret"`
}
