package tx_summitter

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Float(t *testing.T) {
	// set balance
	// 1.2498745*1e18 -> 12498745*1e11
	balance := big.NewInt(12498745)
	balance.Mul(balance, big.NewInt(1e11))
	// to big.float
	balanceFloat := new(big.Float).SetInt(balance)
	// div 1e18 -> ether unit
	balanceFloat.Quo(balanceFloat, big.NewFloat(1e18))
	// to float64
	balanceFloat64, _ := balanceFloat.Float64()
	require.Equal(t, balanceFloat64, 1.2498745)
}
