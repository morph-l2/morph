package batch

import (
	"errors"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

type testSignatureInput struct {
	SignedSequencersBitmap *big.Int
	SequencerSets          []byte
	Signature              []byte
}

func testCurrentInput() batchDataInputStruct {
	return batchDataInputStruct{
		Version:           2,
		ParentBatchHeader: []byte{1, 2, 3},
		LastBlockNumber:   123,
		NumL1Messages:     7,
		PrevStateRoot:     [32]byte{1},
		PostStateRoot:     [32]byte{2},
		WithdrawalRoot:    [32]byte{3},
	}
}

func packBatchCall(t *testing.T, contractABI abi.ABI, methodName string, args ...interface{}) []byte {
	t.Helper()
	method, ok := contractABI.Methods[methodName]
	require.True(t, ok)
	encoded, err := method.Inputs.Pack(args...)
	require.NoError(t, err)
	return append(append([]byte(nil), method.ID...), encoded...)
}

func TestDecodeBatchDataInputAllSupportedSelectors(t *testing.T) {
	abis, err := DefaultRollupABIs()
	require.NoError(t, err)

	current := testCurrentInput()
	signature := testSignatureInput{
		SignedSequencersBitmap: big.NewInt(3),
		SequencerSets:          []byte{4, 5},
		Signature:              []byte{6},
	}
	beforeMove := beforeMoveBatchDataInput{
		Version:           current.Version,
		ParentBatchHeader: current.ParentBatchHeader,
		BlockContexts:     []byte{9},
		PrevStateRoot:     current.PrevStateRoot,
		PostStateRoot:     current.PostStateRoot,
		WithdrawalRoot:    current.WithdrawalRoot,
	}
	legacy := legacyBatchDataInput{
		Version:                current.Version,
		ParentBatchHeader:      current.ParentBatchHeader,
		BlockContexts:          []byte{9},
		SkippedL1MessageBitmap: []byte{10},
		PrevStateRoot:          current.PrevStateRoot,
		PostStateRoot:          current.PostStateRoot,
		WithdrawalRoot:         current.WithdrawalRoot,
	}

	fixtures := []struct {
		name        string
		data        []byte
		legacyShape bool
	}{
		{name: "before-move commitBatch", data: packBatchCall(t, abis.BeforeMoveBlockCtx, "commitBatch", beforeMove, signature), legacyShape: true},
		{name: "legacy commitBatch", data: packBatchCall(t, abis.Legacy, "commitBatch", legacy, signature), legacyShape: true},
		{name: "pre-submitter commitBatch", data: packBatchCall(t, abis.PreSubmitter, "commitBatch", current, signature)},
		{name: "pre-submitter commitState", data: packBatchCall(t, abis.PreSubmitter, "commitState", current, signature)},
		{name: "pre-submitter commitBatchWithProof", data: packBatchCall(t, abis.PreSubmitter, "commitBatchWithProof", current, signature, []byte{11}, []byte{12})},
		{name: "current commitBatch", data: packBatchCall(t, abis.Current, "commitBatch", current)},
		{name: "current commitState", data: packBatchCall(t, abis.Current, "commitState", current)},
		{name: "current commitBatchWithProof", data: packBatchCall(t, abis.Current, "commitBatchWithProof", current, []byte{11}, []byte{12})},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got, err := DecodeBatchDataInput(fixture.data, abis)
			require.NoError(t, err)
			require.Equal(t, current.Version, got.Version)
			require.Equal(t, current.ParentBatchHeader, got.ParentBatchHeader)
			require.Equal(t, current.PrevStateRoot, got.PrevStateRoot)
			require.Equal(t, current.PostStateRoot, got.PostStateRoot)
			require.Equal(t, current.WithdrawalRoot, got.WithdrawalRoot)
			if fixture.legacyShape {
				require.Zero(t, got.LastBlockNumber)
				require.Zero(t, got.NumL1Messages)
			} else {
				require.Equal(t, current.LastBlockNumber, got.LastBlockNumber)
				require.Equal(t, current.NumL1Messages, got.NumL1Messages)
			}
		})
	}
}

func TestCurrentBatchSelectorsAreFrozen(t *testing.T) {
	abis, err := DefaultRollupABIs()
	require.NoError(t, err)
	require.Equal(t, []byte{0x41, 0xf7, 0x56, 0xda}, abis.Current.Methods["commitBatch"].ID)
	require.Equal(t, []byte{0x67, 0xca, 0xa3, 0x7a}, abis.Current.Methods["commitState"].ID)
	require.Equal(t, []byte{0x15, 0x44, 0xba, 0x3a}, abis.Current.Methods["commitBatchWithProof"].ID)
	require.Equal(t, []byte{0x42, 0x88, 0x68, 0xb5}, abis.PreSubmitter.Methods["commitBatch"].ID)
	require.Equal(t, []byte{0x1e, 0x88, 0x25, 0xbe}, abis.PreSubmitter.Methods["commitState"].ID)
	require.Equal(t, []byte{0x4e, 0x8f, 0x1d, 0x67}, abis.PreSubmitter.Methods["commitBatchWithProof"].ID)
	_, err = supportedBatchCalls(abis)
	require.NoError(t, err)
}

func TestDecodeBatchDataInputRejectsMalformedAndUnknownCalldata(t *testing.T) {
	abis, err := DefaultRollupABIs()
	require.NoError(t, err)

	_, err = DecodeBatchDataInput([]byte{1, 2, 3}, abis)
	require.ErrorIs(t, err, ErrCalldataTooShort)

	_, err = DecodeBatchDataInput([]byte{0xde, 0xad, 0xbe, 0xef}, abis)
	require.ErrorIs(t, err, ErrUnknownBatchSelector)

	malformed := append([]byte(nil), abis.Current.Methods["commitBatch"].ID...)
	_, err = DecodeBatchDataInput(malformed, abis)
	require.Error(t, err)
	require.False(t, errors.Is(err, ErrUnknownBatchSelector))
}
