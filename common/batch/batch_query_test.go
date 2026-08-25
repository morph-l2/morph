package batch

import (
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

type historicalBatchSignatureInput struct {
	SignedSequencersBitmap *big.Int
	SequencerSets          []byte
	Signature              []byte
}

func TestParseCommitBatchTxDataSupportsOnlyHistoricalAndCurrentCommitSelectors(t *testing.T) {
	data := batchDataInputStruct{
		Version:           2,
		ParentBatchHeader: []byte{1, 2, 3},
		LastBlockNumber:   12345,
		NumL1Messages:     7,
		PrevStateRoot:     common.HexToHash("0x11"),
		PostStateRoot:     common.HexToHash("0x22"),
		WithdrawalRoot:    common.HexToHash("0x33"),
	}
	signature := historicalBatchSignatureInput{
		SignedSequencersBitmap: big.NewInt(9),
		SequencerSets:          []byte{0xaa, 0xbb},
		Signature:              []byte{0xcc, 0xdd},
	}

	tests := []struct {
		name       string
		rawABI     string
		method     string
		selector   string
		historical bool
	}{
		{name: "historical commitBatch", rawABI: preSubmitterCommitABI, method: "commitBatch", selector: "0x428868b5", historical: true},
		{name: "historical commitBatchWithProof", rawABI: preSubmitterCommitABI, method: "commitBatchWithProof", selector: "0x4e8f1d67", historical: true},
		{name: "current commitBatch", rawABI: submitterCommitABI, method: "commitBatch", selector: "0x41f756da"},
		{name: "current commitBatchWithProof", rawABI: submitterCommitABI, method: "commitBatchWithProof", selector: "0x1544ba3a"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsedABI, err := abi.JSON(strings.NewReader(test.rawABI))
			require.NoError(t, err)
			require.Equal(t, test.selector, hexutil.Encode(parsedABI.Methods[test.method].ID))

			var calldata []byte
			switch {
			case test.historical && test.method == "commitBatch":
				calldata, err = parsedABI.Pack(test.method, data, signature)
			case test.historical:
				calldata, err = parsedABI.Pack(test.method, data, signature, []byte{0x12}, []byte{0x34})
			case test.method == "commitBatch":
				calldata, err = parsedABI.Pack(test.method, data)
			default:
				calldata, err = parsedABI.Pack(test.method, data, []byte{0x12}, []byte{0x34})
			}
			require.NoError(t, err)

			got, err := parseCommitBatchTxData(calldata)
			require.NoError(t, err)
			require.Equal(t, data.Version, got.Version)
			require.Equal(t, data.ParentBatchHeader, got.ParentBatchHeader)
			require.Equal(t, data.LastBlockNumber, got.LastBlockNumber)
			require.Equal(t, data.NumL1Messages, got.NumL1Messages)
			require.Equal(t, data.PrevStateRoot, got.PrevStateRoot)
			require.Equal(t, data.PostStateRoot, got.PostStateRoot)
			require.Equal(t, data.WithdrawalRoot, got.WithdrawalRoot)
		})
	}

	for _, calldata := range [][]byte{
		{0xde, 0xad, 0xbe, 0xef},
		{0x1e, 0x88, 0x25, 0xbe}, // historical commitState
		{0x67, 0xca, 0xa3, 0x7a}, // current commitState
	} {
		_, err := parseCommitBatchTxData(calldata)
		require.Error(t, err)
	}
}

func TestCreateBatchHeaderKeepsWithdrawalRootAndZerosSequencerField(t *testing.T) {
	withdrawalRoot := common.HexToHash("0x1234")
	cache := &BatchCache{
		batchData:            NewBatchData(),
		withdrawRoot:         withdrawalRoot,
		isBatchUpgraded:      func(uint64) bool { return true },
		isBatchV2Upgraded:    func(uint64) bool { return false },
		parentBatchHeader:    nil,
		totalL1MessagePopped: 0,
	}
	header := cache.createBatchHeader(common.HexToHash("0xabcd"), nil, 1)

	gotWithdrawalRoot, err := header.WithdrawalRoot()
	require.NoError(t, err)
	require.Equal(t, withdrawalRoot, gotWithdrawalRoot)
	gotSequencerRoot, err := header.SequencerSetVerifyHash()
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, gotSequencerRoot)
}
