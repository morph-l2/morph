package derivation

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

const targetCurrentRollupABI = `[
  {"type":"function","name":"commitBatch","stateMutability":"payable","inputs":[{"name":"batchDataInput","type":"tuple","components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}]}],"outputs":[]},
  {"type":"function","name":"commitState","stateMutability":"nonpayable","inputs":[{"name":"batchDataInput","type":"tuple","components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}]}],"outputs":[]},
  {"type":"function","name":"commitBatchWithProof","stateMutability":"nonpayable","inputs":[{"name":"batchDataInput","type":"tuple","components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}]},{"name":"batchHeader","type":"bytes"},{"name":"batchProof","type":"bytes"}],"outputs":[]}
]`

type calldataFixtureCorpus struct {
	SchemaVersion int               `json:"schemaVersion"`
	Fixtures      []calldataFixture `json:"fixtures"`
}

type calldataFixture struct {
	Name     string                  `json:"name"`
	Epoch    string                  `json:"epoch"`
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
	require.Len(t, corpus.Fixtures, 8)
	return corpus
}

func newFixtureDerivation(t *testing.T) *Derivation {
	t.Helper()
	currentABI, err := abi.JSON(strings.NewReader(targetCurrentRollupABI))
	require.NoError(t, err)
	preSubmitterABI, err := types.PreSubmitterRollupMetaData.GetAbi()
	require.NoError(t, err)
	legacyABI, err := types.LegacyRollupMetaData.GetAbi()
	require.NoError(t, err)
	beforeMoveABI, err := types.BeforeMoveBlockCtxABI.GetAbi()
	require.NoError(t, err)
	return &Derivation{
		rollupABI:             &currentABI,
		preSubmitterRollupABI: preSubmitterABI,
		legacyRollupABI:       legacyABI,
		beforeMoveBlockCtxABI: beforeMoveABI,
	}
}

func TestUnPackDataAllHistoricalAndCurrentSelectors(t *testing.T) {
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

func TestUnPackDataRejectsUnknownShortAndMalformedCalldata(t *testing.T) {
	d := newFixtureDerivation(t)
	for length := 0; length < 4; length++ {
		_, err := d.UnPackData(make([]byte, length))
		require.ErrorIs(t, err, types.ErrNotCommitBatchTx)
	}
	_, err := d.UnPackData([]byte{0xde, 0xad, 0xbe, 0xef})
	require.ErrorIs(t, err, types.ErrNotCommitBatchTx)

	currentABI, err := abi.JSON(strings.NewReader(targetCurrentRollupABI))
	require.NoError(t, err)
	_, err = d.UnPackData(currentABI.Methods["commitBatch"].ID)
	require.Error(t, err)
	require.False(t, errors.Is(err, types.ErrNotCommitBatchTx), "known selector with malformed arguments must not be reported as unknown")
}

func TestUnPackDataMixedCutoverWindow(t *testing.T) {
	d := newFixtureDerivation(t)
	fixtures := make(map[string]string)
	for _, fixture := range loadCalldataFixtureCorpus(t).Fixtures {
		fixtures[fixture.Name] = fixture.Data
	}
	for _, name := range []string{
		"pre_submitter_commitBatch",
		"current_commitBatch",
		"pre_submitter_commitState",
		"current_commitState",
		"pre_submitter_commitBatchWithProof",
		"current_commitBatchWithProof",
	} {
		data, err := hexutil.Decode(fixtures[name])
		require.NoError(t, err)
		_, err = d.UnPackData(data)
		require.NoError(t, err, name)
	}
}

func TestSubmissionSelectorClassification(t *testing.T) {
	d := newFixtureDerivation(t)
	for _, fixture := range loadCalldataFixtureCorpus(t).Fixtures {
		data, err := hexutil.Decode(fixture.Data)
		require.NoError(t, err)
		switch fixture.Name {
		case "pre_submitter_commitState", "pre_submitter_commitBatchWithProof", "current_commitState", "current_commitBatchWithProof":
			require.True(t, d.isStoredSourceSubmission(data), fixture.Name)
		case "current_commitBatch":
			require.True(t, d.isCurrentFreshCommitBatch(data), fixture.Name)
		default:
			require.False(t, d.isStoredSourceSubmission(data), fixture.Name)
			require.False(t, d.isCurrentFreshCommitBatch(data), fixture.Name)
		}
	}
}

func TestRollupSelectorFreeze(t *testing.T) {
	beforeMoveABI, err := types.BeforeMoveBlockCtxABI.GetAbi()
	require.NoError(t, err)
	legacyABI, err := types.LegacyRollupMetaData.GetAbi()
	require.NoError(t, err)
	require.Equal(t, "0x4a8d544f", hexutil.Encode(beforeMoveABI.Methods["commitBatch"].ID), "before-move selector")
	require.Equal(t, "0xd63b3549", hexutil.Encode(legacyABI.Methods["commitBatch"].ID), "legacy selector")

	preSubmitterABI, err := types.PreSubmitterRollupMetaData.GetAbi()
	require.NoError(t, err)
	currentABI, err := abi.JSON(strings.NewReader(targetCurrentRollupABI))
	require.NoError(t, err)

	wantOld := map[string]string{
		"commitBatch":          "0x428868b5",
		"commitState":          "0x1e8825be",
		"commitBatchWithProof": "0x4e8f1d67",
	}
	wantNew := map[string]string{
		"commitBatch":          "0x41f756da",
		"commitState":          "0x67caa37a",
		"commitBatchWithProof": "0x1544ba3a",
	}
	for name := range wantOld {
		require.Equal(t, wantOld[name], hexutil.Encode(preSubmitterABI.Methods[name].ID))
		require.Equal(t, wantNew[name], hexutil.Encode(currentABI.Methods[name].ID))
		require.NotEqual(t, preSubmitterABI.Methods[name].ID, currentABI.Methods[name].ID)
	}
}

func TestGeneratedRollupBindingUsesCurrentSelectors(t *testing.T) {
	generatedABI, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	want := map[string]string{
		"commitBatch":          "0x41f756da",
		"commitState":          "0x67caa37a",
		"commitBatchWithProof": "0x1544ba3a",
	}
	for name, selector := range want {
		require.Equal(t, selector, hexutil.Encode(generatedABI.Methods[name].ID), name)
	}
}
