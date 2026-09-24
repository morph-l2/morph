package batch

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestExtractInnerTxFullBytesRejectsOversizedDeclaration(t *testing.T) {
	reader := bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff})

	_, err := extractInnerTxFullBytes(0xfb, reader)

	require.EqualError(t, err, "declared tx size 4294967295 exceeds remaining 0 bytes")
}

func TestExtractInnerTxFullBytesAcceptsAvailablePayload(t *testing.T) {
	reader := bytes.NewReader([]byte{3, 1, 2, 3})

	got, err := extractInnerTxFullBytes(0xf8, reader)

	require.NoError(t, err)
	require.Equal(t, []byte{0xf8, 3, 1, 2, 3}, got)
}

func TestMorphTxV2BatchRoundTrip(t *testing.T) {
	to := common.HexToAddress("0x1234")
	tests := []struct {
		name     string
		authList []eth.SetCodeAuthorization
	}{
		{name: "empty authorization list", authList: []eth.SetCodeAuthorization{}},
		{
			name: "non-empty authorization list",
			authList: []eth.SetCodeAuthorization{{
				ChainID: *uint256.NewInt(53077),
				Address: common.HexToAddress("0x5678"),
				Nonce:   7,
				V:       1,
				R:       *uint256.NewInt(2),
				S:       *uint256.NewInt(3),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := eth.NewTx(&eth.MorphTx{
				ChainID:    big.NewInt(53077),
				Nonce:      1,
				GasTipCap:  big.NewInt(1),
				GasFeeCap:  big.NewInt(2),
				Gas:        100000,
				To:         &to,
				Value:      new(big.Int),
				FeeTokenID: 1,
				FeeLimit:   big.NewInt(1000),
				Version:    eth.MorphTxVersion2,
				AuthList:   tt.authList,
				V:          new(big.Int),
				R:          new(big.Int).Lsh(big.NewInt(1), 255),
				S:          new(big.Int).Lsh(big.NewInt(1), 254),
			})
			encoded, err := tx.MarshalBinary()
			require.NoError(t, err)
			require.Equal(t, []byte{eth.MorphTxType, eth.MorphTxVersion2}, encoded[:2])

			decoded, err := DecodeTxsFromBytes(encoded)
			require.NoError(t, err)
			require.Len(t, decoded, 1)
			require.Equal(t, tx.Hash(), decoded[0].Hash())
			reencoded, err := decoded[0].MarshalBinary()
			require.NoError(t, err)
			require.Equal(t, encoded, reencoded)
			require.Equal(t, len(tt.authList), len(decoded[0].AsMorphTx().AuthList))
		})
	}
}

func TestParsingTxsSupportsMorphTxV2(t *testing.T) {
	to := common.HexToAddress("0x1234")
	tx := eth.NewTx(&eth.MorphTx{
		ChainID:    big.NewInt(53077),
		Nonce:      1,
		GasTipCap:  big.NewInt(1),
		GasFeeCap:  big.NewInt(2),
		Gas:        100000,
		To:         &to,
		Value:      new(big.Int),
		FeeTokenID: 1,
		FeeLimit:   big.NewInt(1000),
		Version:    eth.MorphTxVersion2,
		AuthList: []eth.SetCodeAuthorization{{
			ChainID: *uint256.NewInt(53077),
			Address: common.HexToAddress("0x5678"),
			Nonce:   7,
			V:       1,
			R:       *uint256.NewInt(2),
			S:       *uint256.NewInt(3),
		}},
		V: new(big.Int),
		R: new(big.Int).Lsh(big.NewInt(1), 255),
		S: new(big.Int).Lsh(big.NewInt(1), 254),
	})

	want, err := tx.MarshalBinary()
	require.NoError(t, err)

	payload, l1TxHashes, totalL1MessagePopped, l2TxNum, err := ParsingTxs(
		[]*eth.Transaction{tx},
		0,
	)
	require.NoError(t, err)
	require.Equal(t, want, payload)
	require.Empty(t, l1TxHashes)
	require.Zero(t, totalL1MessagePopped)
	require.Equal(t, 1, l2TxNum)
}
