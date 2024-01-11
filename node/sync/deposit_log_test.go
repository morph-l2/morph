package sync

import (
	"math/big"
	"testing"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestParseRelayMessage(t *testing.T) {
	abi, err := bindings.L2CrossDomainMessengerMetaData.GetAbi()
	require.NoError(t, err)

	nonce := big.NewInt(1000)
	sender := common.HexToAddress("0x123")
	target := common.HexToAddress("0x456")
	value := big.NewInt(2e18)
	gasLimit := big.NewInt(1000000)
	message := []byte("abcd")
	packedBytes, err := abi.Pack("relayMessage", nonce, sender, target, value, gasLimit, message)
	require.NoError(t, err)

	unpacked, err := unpackRelayMessage(packedBytes)
	require.NoError(t, err)

	require.EqualValues(t, nonce, unpacked.nonce)
	require.EqualValues(t, sender, unpacked.sender)
	require.EqualValues(t, target, unpacked.target)
	require.EqualValues(t, value.Int64(), unpacked.value.Int64())
	require.EqualValues(t, gasLimit.Int64(), unpacked.minGasLimit.Int64())
	require.EqualValues(t, message, unpacked.message)
}
