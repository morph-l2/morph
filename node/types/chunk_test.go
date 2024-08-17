package types

import (
	"encoding/binary"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestChunks_AppendNilRows(t *testing.T) {
	blockContext := []byte("123")
	txPayloads := []byte("abc")
	txHashes := []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))}
	chunks := NewChunks()
	chunks.Append(blockContext, txPayloads, txHashes, types.RowConsumption{{Name: "a", RowNumber: 980_000}})
	require.EqualValues(t, 1, len(chunks.data))

	chunks.Append(blockContext, txPayloads, txHashes, nil)
	require.EqualValues(t, 1, len(chunks.data))

	chunks.Append(blockContext, txPayloads, txHashes, types.RowConsumption{{Name: "a", RowNumber: 980_000}})
	require.EqualValues(t, 2, len(chunks.data))
}

func TestChunks_Append(t *testing.T) {
	chunks := NewChunks()
	appended := chunks.isChunksAppendedWithNewBlock(types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.True(t, appended)

	blockContext := []byte("123")
	txPayloads := []byte("abc")
	txHashes := []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))}
	chunks.Append(blockContext, txPayloads, txHashes, types.RowConsumption{{Name: "a", RowNumber: 1}, {Name: "b", RowNumber: 2}})
	require.EqualValues(t, 1, len(chunks.data))
	require.EqualValues(t, 1, chunks.data[0].blockNum)
	require.EqualValues(t, blockContext, chunks.data[0].blockContext)
	require.EqualValues(t, txPayloads, chunks.data[0].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].l1TxHashes)/32)
	require.EqualValues(t, txHashes[0].Bytes(), chunks.data[0].l1TxHashes[0:32])
	require.EqualValues(t, txHashes[1].Bytes(), chunks.data[0].l1TxHashes[32:])
	require.EqualValues(t, 1, chunks.BlockNum())
	require.EqualValues(t, 1, chunks.ChunkNum())

	blockContext1 := []byte("456")
	txPayloads1 := []byte("def")
	chunks.Append(blockContext1, txPayloads1, nil, types.RowConsumption{{Name: "a", RowNumber: 999999}, {Name: "b", RowNumber: 999998}})
	require.EqualValues(t, 1, len(chunks.data))
	require.EqualValues(t, 2, chunks.data[0].blockNum)
	require.EqualValues(t, append(blockContext, blockContext1...), chunks.data[0].blockContext)
	require.EqualValues(t, append(txPayloads, txPayloads1...), chunks.data[0].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].l1TxHashes)/32)
	require.EqualValues(t, 2, chunks.BlockNum())
	require.EqualValues(t, 1, chunks.ChunkNum())

	// the 2nd chunk
	blockContext2 := []byte("789")
	txPayloads2 := []byte("ghi")
	txHashes2 := []common.Hash{common.BigToHash(big.NewInt(3))}
	chunks.Append(blockContext2, txPayloads2, txHashes2, types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.EqualValues(t, 2, len(chunks.data))
	require.EqualValues(t, 2, chunks.data[0].blockNum)
	require.EqualValues(t, 1, chunks.data[1].blockNum)
	require.EqualValues(t, append(blockContext, blockContext1...), chunks.data[0].blockContext)
	require.EqualValues(t, append(txPayloads, txPayloads1...), chunks.data[0].txsPayload)
	require.EqualValues(t, blockContext2, chunks.data[1].blockContext)
	require.EqualValues(t, txPayloads2, chunks.data[1].txsPayload)
	require.EqualValues(t, len(txHashes), len(chunks.data[0].l1TxHashes)/32)
	require.EqualValues(t, len(txHashes2), len(chunks.data[1].l1TxHashes)/32)
	require.EqualValues(t, txHashes2[0].Bytes(), chunks.data[1].l1TxHashes[0:32])
	require.EqualValues(t, 3, chunks.BlockNum())
	require.EqualValues(t, 2, chunks.ChunkNum())

	for i := 0; i < 98; i++ {
		chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{Name: "a", RowNumber: 1}})
	}
	// 99 blocks in 2nd chunk
	require.EqualValues(t, 2, chunks.ChunkNum())
	appended = chunks.isChunksAppendedWithNewBlock(types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.False(t, appended)
	// 100 blocks in 2nd chunk
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.EqualValues(t, 2, chunks.ChunkNum())

	appended = chunks.isChunksAppendedWithNewBlock(types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.True(t, appended)
	// append chunk to 3 chunks totally
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{Name: "a", RowNumber: 1}})
	require.EqualValues(t, 3, chunks.ChunkNum())
}

func TestChunks_ConstructBlobPayload(t *testing.T) {
	chunks := NewChunks()
	txsPayload0 := rand.Bytes(20)
	// 1st chunk has 10 length of tx payload
	chunks.Append(nil, txsPayload0, nil, types.RowConsumption{{Name: "a", RowNumber: 1_000_000}})

	// 2nd chunk has empty txs
	chunks.Append(nil, nil, nil, types.RowConsumption{{Name: "a", RowNumber: 1_000_000}})

	txsPayload1 := rand.Bytes(30)
	// 3rd chunk has 30 length of tx payload
	chunks.Append(nil, txsPayload1, nil, types.RowConsumption{{Name: "a", RowNumber: 500_000}})

	txsPayload2 := rand.Bytes(15)
	// 3rd chunk append 15 length of tx payload
	chunks.Append(nil, txsPayload2, nil, types.RowConsumption{{Name: "a", RowNumber: 500_000}})

	blobBytes := chunks.ConstructBlobPayload()

	skipBytes := 2 + MaxChunks*4
	expectedBytes := make([]byte, skipBytes+65)
	copy(expectedBytes, []byte{0x0, 0x3})
	chunk0Size := make([]byte, 4)
	binary.BigEndian.PutUint32(chunk0Size, 20)
	chunk2Size := make([]byte, 4)
	binary.BigEndian.PutUint32(chunk2Size, 45)
	copy(expectedBytes[2:], chunk0Size)
	copy(expectedBytes[10:], chunk2Size)
	copy(expectedBytes[skipBytes:], append(append(txsPayload0, txsPayload1...), txsPayload2...))
	require.EqualValues(t, expectedBytes, blobBytes)
}

func TestChunk_accumulateRowUsages(t *testing.T) {
	chunk := new(Chunk)
	rc := types.RowConsumption{{Name: "a", RowNumber: 1}}
	accRc, max := chunk.accumulateRowUsages(rc)
	require.True(t, equalRc(rc, accRc))
	require.EqualValues(t, 1, max)

	chunk = NewChunk(nil, nil, nil, types.RowConsumption{{Name: "a", RowNumber: 1}, {Name: "b", RowNumber: 2}})
	rc = types.RowConsumption{{Name: "a", RowNumber: 3}}
	accRc, max = chunk.accumulateRowUsages(rc)
	require.True(t, equalRc(types.RowConsumption{{Name: "a", RowNumber: 4}, {Name: "b", RowNumber: 2}}, accRc))
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
