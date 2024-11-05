package utils

import (
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
)

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func MaxOfThreeBig(a, b, c *big.Int) *big.Int {
	max := a

	if b.Cmp(max) > 0 {
		max = b
	}

	if c.Cmp(max) > 0 {
		max = c
	}

	return max
}

// IntersectionOfAddresses
func IntersectionOfAddresses(a, b []common.Address) []common.Address {
	intersection := make([]common.Address, 0)
	aMap := make(map[common.Address]bool)

	// Create a map of addresses in set 'a' for faster lookup
	for _, addr := range a {
		aMap[addr] = true
	}

	// Check each address in 'b' against the map of 'a'
	for _, addr := range b {
		if aMap[addr] {
			intersection = append(intersection, addr)
		}
	}

	return intersection
}
