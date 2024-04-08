package types

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestChunks_Append(t *testing.T) {
	chunks := NewChunks()
	appended, _ := chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}})
	require.True(t, appended)

	blockContext := []byte("123")
	txPayloads := []byte("abc")
	txHashes := []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))}
	chunks.Append(blockContext, txPayloads, txHashes, types.RowConsumption{{"a", 1}, {"b", 2}})
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
	chunks.Append(blockContext1, txPayloads1, nil, types.RowConsumption{{"a", 999999}, {"b", 999998}})
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
	chunks.Append(blockContext2, txPayloads2, txHashes2, types.RowConsumption{{"a", 1}})
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
		chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	}
	// 99 blocks in 2nd chunk
	require.EqualValues(t, 2, chunks.ChunkNum())
	appended, _ = chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}})
	require.False(t, appended)
	// 100 blocks in 2nd chunk
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	require.EqualValues(t, 2, chunks.ChunkNum())

	appended, _ = chunks.IsChunksAppendedWithNewBlock(types.RowConsumption{{"a", 1}})
	require.True(t, appended)
	// append chunk to 3 chunks totally
	chunks.Append([]byte("11"), nil, nil, types.RowConsumption{{"a", 1}})
	require.EqualValues(t, 3, chunks.ChunkNum())
}

func TestChunk_Seal(t *testing.T) {
	chunk := NewChunk(nil, nil, nil, nil)
	chunk.Seal()
	require.EqualValues(t, 31, len(chunk.sealedPayload))
	require.EqualValues(t, make([]byte, 31), chunk.sealedPayload)

	chunk = NewChunk(nil, make([]byte, 0), nil, nil)
	chunk.Seal()
	require.EqualValues(t, 31, len(chunk.sealedPayload))
	require.EqualValues(t, make([]byte, 31), chunk.sealedPayload)

	txPayload := rand.Bytes(10)
	chunk = NewChunk(nil, txPayload, nil, nil)
	require.False(t, chunk.Sealed())
	require.EqualValues(t, 0, len(chunk.sealedPayload))
	chunk.Seal()
	require.True(t, chunk.Sealed())
	require.EqualValues(t, 31, len(chunk.sealedPayload))
	require.EqualValues(t, 10, binary.LittleEndian.Uint32(chunk.sealedPayload[:4]))
	require.EqualValues(t, txPayload, chunk.sealedPayload[4:14])
	require.EqualValues(t, make([]byte, 31-14), chunk.sealedPayload[14:])

	// full one 31bytes
	txPayload = rand.Bytes(27)
	chunk = NewChunk(nil, txPayload, nil, nil)
	require.False(t, chunk.Sealed())
	chunk.Seal()
	require.EqualValues(t, 31, len(chunk.sealedPayload))
	require.EqualValues(t, 27, binary.LittleEndian.Uint32(chunk.sealedPayload[:4]))
	require.EqualValues(t, txPayload, chunk.sealedPayload[4:])

	// full 2 31bytes
	txPayload = rand.Bytes(58)
	chunk = NewChunk(nil, txPayload, nil, nil)
	require.False(t, chunk.Sealed())
	chunk.Seal()
	require.EqualValues(t, 62, len(chunk.sealedPayload))
	require.EqualValues(t, 58, binary.LittleEndian.Uint32(chunk.sealedPayload[:4]))
	require.EqualValues(t, txPayload, chunk.sealedPayload[4:])

	// more 2 31bytes
	txPayload = rand.Bytes(59)
	chunk = NewChunk(nil, txPayload, nil, nil)
	require.False(t, chunk.Sealed())
	chunk.Seal()
	require.EqualValues(t, 93, len(chunk.sealedPayload))
	require.EqualValues(t, 59, binary.LittleEndian.Uint32(chunk.sealedPayload[:4]))
	require.EqualValues(t, txPayload, chunk.sealedPayload[4:63])
	require.EqualValues(t, make([]byte, 93-63), chunk.sealedPayload[63:])
}

func TestChunks_SealTxPayloadForBlob(t *testing.T) {
	chunks := NewChunks()
	// 1st chunk with nil txPayload, takes up 1 31bytes
	chunks.Append(nil, nil, nil, types.RowConsumption{{"a", 1_000_000}})
	require.EqualValues(t, 4, chunks.CurrentPayloadForBlobSize())

	// 2nd chunk with 10bytes txPayload, takes up 1 31bytes
	txsPayload10 := rand.Bytes(10)
	chunks.Append(nil, txsPayload10, nil, types.RowConsumption{{"a", 1_000_000}})
	require.EqualValues(t, 31+4+10, chunks.CurrentPayloadForBlobSize())

	// 3rd chunk with 27bytes txPayload, takes up 1 31bytes
	txsPayload27 := rand.Bytes(27)
	chunks.Append(nil, txsPayload27, nil, types.RowConsumption{{"a", 1_000_000}})
	require.EqualValues(t, 31*2+4+27, chunks.CurrentPayloadForBlobSize())

	// 4th chunk with 58bytes txPayload, takes up 2 31bytes
	txsPayload58 := rand.Bytes(58)
	chunks.Append(nil, txsPayload58, nil, types.RowConsumption{{"a", 1_000_000}})
	require.EqualValues(t, 31*3+4+58, chunks.CurrentPayloadForBlobSize())

	// 5th chunk with 58bytes txPayload, takes up 3 31bytes
	txsPayload59 := rand.Bytes(59)
	chunks.Append(nil, txsPayload59, nil, types.RowConsumption{{"a", 1_000_000}})
	require.EqualValues(t, 31*5+4+59, chunks.CurrentPayloadForBlobSize())

	// chunks totally takes 8 31bytes
	sealedTxPayload := chunks.SealTxPayloadForBlob()
	require.EqualValues(t, 8*31, len(sealedTxPayload))
	len10 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len10, 10)
	len27 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len27, 27)
	len58 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len58, 58)
	len59 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len59, 59)
	expectedSealed := bytes.NewBuffer(make([]byte, 31))
	expectedSealed.Write(len10)
	expectedSealed.Write(txsPayload10)
	expectedSealed.Write(make([]byte, 31-14))
	expectedSealed.Write(len27)
	expectedSealed.Write(txsPayload27)
	expectedSealed.Write(len58)
	expectedSealed.Write(txsPayload58)
	expectedSealed.Write(len59)
	expectedSealed.Write(txsPayload59)
	expectedSealed.Write(make([]byte, 30))
	require.EqualValues(t, expectedSealed.Bytes(), sealedTxPayload)
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
