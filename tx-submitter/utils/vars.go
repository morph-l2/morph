package utils

import (
	"crypto/sha256"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
)

var (
	OneHash = common.Hash(sha256.Sum256([]byte("1")))
	OneBig  = big.NewInt(1)
	OneUint = uint64(1)
	Addr    = common.HexToAddress("0x76F3768A7D61Bd7F2bF3B44024D3da59EF237e65")
)
