package utils

import (
	"morph-l2/bindings/bindings"
	"os"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func getTestBatchCalldata() []byte {
	bs, _ := os.ReadFile("./testdata/commitbatchcalldata.txt")
	batchBs := common.Hex2Bytes(string(bs))
	return batchBs
}
func getTestBatchData() bindings.IRollupBatchDataInput {
	calldata := getTestBatchCalldata()
	abi, _ := bindings.RollupMetaData.GetAbi()
	params, _ := abi.Methods["commitBatch"].Inputs.UnpackValues(calldata[4:])
	tempStruct := params[0].(struct {
		Version           uint8     "json:\"version\""
		ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
		BlockContexts     []uint8   "json:\"blockContexts\""
		PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot     [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
	})

	return bindings.IRollupBatchDataInput{
		Version:           tempStruct.Version,
		ParentBatchHeader: tempStruct.ParentBatchHeader,
		BlockContexts:     tempStruct.BlockContexts,
		PrevStateRoot:     tempStruct.PrevStateRoot,
		PostStateRoot:     tempStruct.PostStateRoot,
		WithdrawalRoot:    tempStruct.WithdrawalRoot,
	}

}

// test ParseBatchIndex
func TestParseBatchIndex(t *testing.T) {
	calldata := getTestBatchCalldata()
	require.EqualValues(t, 3823, ParseParentBatchIndex(calldata))
}

func TestParseBlockCnt(t *testing.T) {
	batch := getTestBatchData()
	blockCnt := ParseL2BlockCnt(batch.BlockContexts)
	require.EqualValues(t, 461, blockCnt)

}
