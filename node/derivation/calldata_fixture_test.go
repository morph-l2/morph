package derivation

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

type calldataFixtureCorpus struct {
	SchemaVersion int               `json:"schemaVersion"`
	Fixtures      []calldataFixture `json:"fixtures"`
}

type calldataFixture struct {
	Name     string                  `json:"name"`
	Selector string                  `json:"selector"`
	Data     string                  `json:"data"`
	Expected calldataFixtureExpected `json:"expected"`
}

type calldataFixtureExpected struct {
	Version           uint64 `json:"version"`
	ParentBatchHeader string `json:"parentBatchHeader"`
	BlockContexts     string `json:"blockContexts"`
	LastBlockNumber   uint64 `json:"lastBlockNumber"`
	NumL1Messages     uint16 `json:"numL1Messages"`
	PrevStateRoot     string `json:"prevStateRoot"`
	PostStateRoot     string `json:"postStateRoot"`
	WithdrawRoot      string `json:"withdrawRoot"`
}

func loadCalldataFixtureCorpus(t *testing.T) calldataFixtureCorpus {
	t.Helper()
	raw, err := os.ReadFile("testdata/rollup_calldata_fixtures.json")
	require.NoError(t, err)
	var corpus calldataFixtureCorpus
	require.NoError(t, json.Unmarshal(raw, &corpus))
	require.Equal(t, 1, corpus.SchemaVersion)
	require.Len(t, corpus.Fixtures, 6)
	return corpus
}

func newFixtureDerivation(t *testing.T) *Derivation {
	t.Helper()
	currentABI, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	preSubmitterABI, err := types.PreSubmitterRollupMetaData.GetAbi()
	require.NoError(t, err)
	legacyABI, err := types.LegacyRollupMetaData.GetAbi()
	require.NoError(t, err)
	beforeMoveABI, err := types.BeforeMoveBlockCtxABI.GetAbi()
	require.NoError(t, err)
	return &Derivation{
		rollupABI:             currentABI,
		preSubmitterRollupABI: preSubmitterABI,
		legacyRollupABI:       legacyABI,
		beforeMoveBlockCtxABI: beforeMoveABI,
	}
}

func TestUnPackDataHistoricalAndCurrentSelectors(t *testing.T) {
	d := newFixtureDerivation(t)
	for _, fixture := range loadCalldataFixtureCorpus(t).Fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			data, err := hexutil.Decode(fixture.Data)
			require.NoError(t, err)
			require.Equal(t, fixture.Selector, hexutil.Encode(data[:4]))

			got, err := d.UnPackData(data)
			require.NoError(t, err)
			require.Equal(t, uint(fixture.Expected.Version), got.Version)
			require.Equal(t, common.FromHex(fixture.Expected.ParentBatchHeader), []byte(got.ParentBatchHeader))
			expectedBlockContexts := fixture.Expected.BlockContexts
			if expectedBlockContexts == "" {
				expectedBlockContexts = "0x"
			}
			require.Equal(t, expectedBlockContexts, hexutil.Encode(got.BlockContexts))
			require.Equal(t, fixture.Expected.LastBlockNumber, got.LastBlockNumber)
			require.Equal(t, fixture.Expected.NumL1Messages, got.NumL1Messages)
			require.Equal(t, common.HexToHash(fixture.Expected.PrevStateRoot), got.PrevStateRoot)
			require.Equal(t, common.HexToHash(fixture.Expected.PostStateRoot), got.PostStateRoot)
			require.Equal(t, common.HexToHash(fixture.Expected.WithdrawRoot), got.WithdrawRoot)
		})
	}
}

func TestUnPackDataRejectsCommitStateUnknownShortAndMalformedCalldata(t *testing.T) {
	d := newFixtureDerivation(t)
	for _, input := range [][]byte{
		nil,
		{0x01},
		{0x01, 0x02},
		{0x01, 0x02, 0x03},
		{0xde, 0xad, 0xbe, 0xef},
		{0x1e, 0x88, 0x25, 0xbe}, // pre-upgrade commitState
		{0x67, 0xca, 0xa3, 0x7a}, // current commitState
	} {
		_, err := d.UnPackData(input)
		require.ErrorIs(t, err, types.ErrNotCommitBatchTx)
	}

	_, err := d.UnPackData([]byte{0x41, 0xf7, 0x56, 0xda})
	require.Error(t, err)
	require.False(t, errors.Is(err, types.ErrNotCommitBatchTx), "a known selector with malformed arguments is not an unknown method")
}

func TestRollupSelectorFreeze(t *testing.T) {
	beforeMoveABI, err := types.BeforeMoveBlockCtxABI.GetAbi()
	require.NoError(t, err)
	legacyABI, err := types.LegacyRollupMetaData.GetAbi()
	require.NoError(t, err)
	preSubmitterABI, err := types.PreSubmitterRollupMetaData.GetAbi()
	require.NoError(t, err)
	currentABI, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)

	require.Equal(t, "0x4a8d544f", hexutil.Encode(beforeMoveABI.Methods["commitBatch"].ID))
	require.Equal(t, "0xd63b3549", hexutil.Encode(legacyABI.Methods["commitBatch"].ID))
	require.Equal(t, "0x428868b5", hexutil.Encode(preSubmitterABI.Methods["commitBatch"].ID))
	require.Equal(t, "0x4e8f1d67", hexutil.Encode(preSubmitterABI.Methods["commitBatchWithProof"].ID))
	require.Equal(t, "0x41f756da", hexutil.Encode(currentABI.Methods["commitBatch"].ID))
	require.Equal(t, "0x1544ba3a", hexutil.Encode(currentABI.Methods["commitBatchWithProof"].ID))
	_, oldCommitStatePresent := preSubmitterABI.Methods["commitState"]
	_, currentCommitStatePresent := currentABI.Methods["commitState"]
	require.True(t, oldCommitStatePresent, "the archived full ABI still records historical commitState")
	require.True(t, currentCommitStatePresent, "the full current ABI includes commitState")
}
