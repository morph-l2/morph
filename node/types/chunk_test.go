package types

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/core/types"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestChunks_Append(t *testing.T) {
	chunks := NewChunks()
	require.True(t, chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}}))

	blockContext := []byte("123")
	txPayloads := []byte("abc")
	txHashes := []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))}
	chunks.Append(blockContext, txPayloads, txHashes, types.RowConsumption{{"a", 1}, {"b", 2}})
	require.EqualValues(t, 1, len(chunks.data))
	require.EqualValues(t, 1, chunks.data[0].blockNum)
	require.EqualValues(t, blockContext, chunks.data[0].blockContext)
	require.EqualValues(t, txPayloads, chunks.data[0].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].txHashes)/32)
	require.EqualValues(t, txHashes[0].Bytes(), chunks.data[0].txHashes[0:32])
	require.EqualValues(t, txHashes[1].Bytes(), chunks.data[0].txHashes[32:])
	require.EqualValues(t, 1, chunks.BlockNum())
	require.EqualValues(t, 1, chunks.ChunkNum())

	blockContext1 := []byte("456")
	txPayloads1 := []byte("def")
	chunks.Append(blockContext1, txPayloads1, nil, types.RowConsumption{{"a", 999999}, {"b", 999998}})
	require.EqualValues(t, 1, len(chunks.data))
	require.EqualValues(t, 2, chunks.data[0].blockNum)
	require.EqualValues(t, append(blockContext, blockContext1...), chunks.data[0].blockContext)
	require.EqualValues(t, append(txPayloads, txPayloads1...), chunks.data[0].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].txHashes)/32)
	require.EqualValues(t, 2, chunks.BlockNum())
	require.EqualValues(t, 1, chunks.ChunkNum())

	// the 2nd chunk
	blockContext2 := []byte("789")
	txPayloads2 := []byte("ghi")
	txHashes2 := []common.Hash{common.BigToHash(big.NewInt(3))}
	chunks.Append(blockContext2, txPayloads2, txHashes2, types.RowConsumption{{"a", 1}})
	require.EqualValues(t, 2, len(chunks.data))
	require.EqualValues(t, 2, chunks.data[0].blockNum)
	require.EqualValues(t, 1, chunks.data[1].blockNum)
	require.EqualValues(t, append(blockContext, blockContext1...), chunks.data[0].blockContext)
	require.EqualValues(t, append(txPayloads, txPayloads1...), chunks.data[0].txsPayload)
	require.EqualValues(t, blockContext2, chunks.data[1].blockContext)
	require.EqualValues(t, txPayloads2, chunks.data[1].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].txHashes)/32)
	require.EqualValues(t, len(txHashes2), len(chunks.data[1].txHashes)/32)
	require.EqualValues(t, txHashes2[0].Bytes(), chunks.data[1].txHashes[0:32])
	require.EqualValues(t, 3, chunks.BlockNum())
	require.EqualValues(t, 2, chunks.ChunkNum())

	for i := 0; i < 98; i++ {
		chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	}
	// 99 blocks in 2nd chunk
	require.EqualValues(t, 2, chunks.ChunkNum())
	require.False(t, chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}}))
	// 100 blocks in 2nd chunk
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	require.EqualValues(t, 2, chunks.ChunkNum())

	require.True(t, chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}}))
	// append chunk to 3 chunks totally
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	require.EqualValues(t, 3, chunks.ChunkNum())
}

func TestChunk_accumulateRowUsages(t *testing.T) {
	chunk := new(Chunk)
	rc := types.RowConsumption{{"a", 1}}
	accRc, max := chunk.accumulateRowUsages(rc)
	require.True(t, equalRc(rc, accRc))
	require.EqualValues(t, 1, max)

	chunk = NewChunk(nil, nil, nil, types.RowConsumption{{"a", 1}, {"b", 2}})
	rc = types.RowConsumption{{"a", 3}}
	accRc, max = chunk.accumulateRowUsages(rc)
	require.True(t, equalRc(types.RowConsumption{{"a", 4}, {"b", 2}}, accRc))
	require.EqualValues(t, 4, max)
}

func equalRc(arg0, arg1 types.RowConsumption) bool {
	if len(arg0) != len(arg1) {
		return false
	}
	arg0map := make(map[string]uint64)
	for _, rc := range arg0 {
		arg0map[rc.Name] = rc.RowNumber
	}
	arg1map := make(map[string]uint64)
	for _, rc := range arg1 {
		arg1map[rc.Name] = rc.RowNumber
	}

	for arg0K, arg0V := range arg0map {
		arg1V, ok := arg1map[arg0K]
		if !ok || arg0V != arg1V {
			return false
		}
	}
	return true
}
