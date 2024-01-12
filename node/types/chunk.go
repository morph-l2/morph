package types

import (
	"errors"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
)

const NormalizedRowLimit = 1_000_000

type Chunk struct {
	blockContext  []byte
	txsPayload    []byte
	txHashes      []common.Hash
	accumulatedRc types.RowConsumption
	blockNum      int
}

func NewChunk(blockContext, txsPayload []byte, txHashes []common.Hash, rc types.RowConsumption) *Chunk {
	return &Chunk{
		blockContext:  blockContext,
		txsPayload:    txsPayload,
		txHashes:      txHashes,
		accumulatedRc: rc,
		blockNum:      1,
	}
}

func (ck *Chunk) append(blockContext, txsPayload []byte, txHashes []common.Hash, accRc types.RowConsumption) {
	ck.blockContext = append(ck.blockContext, blockContext...)
	ck.txsPayload = append(ck.txsPayload, txsPayload...)
	ck.txHashes = append(ck.txHashes, txHashes...)
	ck.accumulatedRc = accRc
	ck.blockNum++
}

func (ck *Chunk) accumulateRowUsages(rc types.RowConsumption) (accRc types.RowConsumption, max uint64) {
	if len(ck.accumulatedRc) == 0 {
		return rc, maxRowNumber(rc)
	}
	if len(rc) == 0 {
		copy(accRc, ck.accumulatedRc)
		max = maxRowNumber(ck.accumulatedRc)
		return
	}
	accRowUsagesBefore := make(map[string]uint64)
	for _, element := range ck.accumulatedRc {
		accRowUsagesBefore[element.Name] = element.RowNumber
	}
	addRowUsage := make(map[string]uint64)
	for _, element := range rc {
		addRowUsage[element.Name] = element.RowNumber
	}
	accRowUsagesAfter := make(map[string]uint64)
	for name, rowNumber := range accRowUsagesBefore {
		add, ok := addRowUsage[name]
		if ok {
			accRowUsagesAfter[name] = rowNumber + add
		} else {
			accRowUsagesAfter[name] = rowNumber
		}
	}

	for name, rowNumber := range accRowUsagesAfter {
		accRc = append(accRc, types.SubCircuitRowUsage{Name: name, RowNumber: rowNumber})
		if rowNumber > max {
			max = rowNumber
		}
	}
	return
}

func maxRowNumber(rc types.RowConsumption) (max uint64) {
	for _, subRc := range rc {
		if subRc.RowNumber > max {
			max = subRc.RowNumber
		}
	}
	return
}

func (ck *Chunk) ResetBlockNum(blockNum int) {
	ck.blockNum = blockNum
}

func (ck *Chunk) BlockContext() []byte {
	return ck.blockContext
}

func (ck *Chunk) TxsPayload() []byte {
	return ck.txsPayload
}

func (ck *Chunk) TxHashes() []common.Hash {
	return ck.txHashes
}

func (ck *Chunk) BlockNum() int {
	return ck.blockNum
}

// Encode encodes the chunk into bytes
// Below is the encoding for `Chunk`, total 60*n+1+m bytes.
// Field           Bytes       Type            Index       Comments
// numBlocks       1           uint8           0           The number of blocks in this chunk
// block[0]        60          BlockContext    1           The first block in this chunk
// ......
// block[i]        60          BlockContext    60*i+1      The (i+1)'th block in this chunk
// ......
// block[n-1]      60          BlockContext    60*n-59     The last block in this chunk
// l2Transactions  dynamic     bytes           60*n+1
func (ck *Chunk) Encode() ([]byte, error) {
	if ck == nil || ck.blockNum == 0 {
		return []byte{}, nil
	}
	if ck.blockNum > 255 {
		return nil, errors.New("number of blocks exceeds 1 byte")
	}
	var chunkBytes []byte
	chunkBytes = append(chunkBytes, byte(ck.blockNum))
	chunkBytes = append(chunkBytes, ck.blockContext...)
	chunkBytes = append(chunkBytes, ck.txsPayload...)
	return chunkBytes, nil
}

func (ck *Chunk) Hash() common.Hash {
	var bytes []byte
	for i := 0; i < ck.blockNum; i++ {
		bytes = append(bytes, ck.blockContext[i*60:i*60+58]...)
	}
	for _, txHash := range ck.txHashes {
		bytes = append(bytes, txHash[:]...)
	}
	return crypto.Keccak256Hash(bytes)
}

type Chunks struct {
	data     []*Chunk
	blockNum int

	size int
	hash *common.Hash
}

func NewChunks() *Chunks {
	return &Chunks{
		data: make([]*Chunk, 0),
	}
}

func (cks *Chunks) Append(blockContext, txsPayload []byte, txHashes []common.Hash, rc types.RowConsumption) {
	if cks == nil {
		return
	}
	defer func() {
		cks.size += len(blockContext) + len(txsPayload)
		cks.blockNum++
		cks.hash = nil // clear hash when data is updated
	}()
	if len(cks.data) == 0 {
		cks.data = append(cks.data, NewChunk(blockContext, txsPayload, txHashes, rc))
		cks.size += 1
		return
	}
	lastChunk := cks.data[len(cks.data)-1]
	accRc, max := lastChunk.accumulateRowUsages(rc)
	if max > NormalizedRowLimit { // add a new chunk
		cks.data = append(cks.data, NewChunk(blockContext, txsPayload, txHashes, rc))
		cks.size += 1
		return
	}
	lastChunk.append(blockContext, txsPayload, txHashes, accRc)
	return
}

func (cks *Chunks) Encode() ([][]byte, error) {
	var bytes [][]byte
	for _, ck := range cks.data {
		ckBytes, err := ck.Encode()
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, ckBytes)
	}
	return bytes, nil
}

func (cks *Chunks) DataHash() common.Hash {
	if cks.hash != nil {
		return *cks.hash
	}
	var chunkHashes []byte
	for _, ck := range cks.data {
		hash := ck.Hash()
		chunkHashes = append(chunkHashes, hash[:]...)
	}
	hash := crypto.Keccak256Hash(chunkHashes)
	cks.hash = &hash
	return hash
}

func (cks *Chunks) BlockNum() int { return cks.blockNum }
func (cks *Chunks) ChunkNum() int { return len(cks.data) }
func (cks *Chunks) Size() int     { return cks.size }
func (cks *Chunks) IsChunksAppendedWithAddedRc(rc types.RowConsumption) bool {
	if len(cks.data) == 0 {
		return true
	}
	lastChunk := cks.data[len(cks.data)-1]
	_, max := lastChunk.accumulateRowUsages(rc)
	return max > NormalizedRowLimit
}
