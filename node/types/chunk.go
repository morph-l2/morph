package types

import (
	"encoding/binary"
	"errors"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
)

const (
	NormalizedRowLimit = 1_000_000
	MaxBlocksPerChunk  = 100
)

type Chunk struct {
	blockContext  []byte
	txHashes      []byte
	accumulatedRc types.RowConsumption
	blockNum      int
	txsPayload    []byte // the raw txs payload
	sealedPayload []byte
	sealed        bool
}

func NewChunk(blockContext, txsPayload []byte, txHashes []common.Hash, rc types.RowConsumption) *Chunk {
	var txHashBytes []byte
	for _, txHash := range txHashes {
		txHashBytes = append(txHashBytes, txHash.Bytes()...)
	}
	return &Chunk{
		blockContext:  blockContext,
		txsPayload:    txsPayload,
		txHashes:      txHashBytes,
		accumulatedRc: rc,
		blockNum:      1,
	}
}

func (ck *Chunk) append(blockContext, txsPayload []byte, txHashes []common.Hash, accRc types.RowConsumption) {
	ck.blockContext = append(ck.blockContext, blockContext...)
	ck.txsPayload = append(ck.txsPayload, txsPayload...)
	ck.accumulatedRc = accRc
	ck.blockNum++
	for _, txHash := range txHashes {
		ck.txHashes = append(ck.txHashes, txHash.Bytes()...)
	}
}

// Seal build the final txs payload that is ready to be put into the eip4844 blob.
// 1. add first 4bytes in front of the payload to indicates the length of the raw txsPayload
// 2. append zero bytes in the end of the payload to make the whole payload a multiple of 31
func (ck *Chunk) Seal() {
	if ck.sealed {
		return
	}

	finalPayload := make([]byte, 4+len(ck.txsPayload))
	binary.LittleEndian.PutUint32(finalPayload[:4], uint32(len(ck.txsPayload)))
	copy(finalPayload[4:], ck.txsPayload)

	if zeroNum := addedZeroNum(len(finalPayload)); zeroNum > 0 {
		finalPayload = append(finalPayload, make([]byte, zeroNum)...)
	}
	ck.sealedPayload = finalPayload
	ck.sealed = true
}

func (ck *Chunk) Sealed() bool {
	return ck.sealed
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
// l2TxHashes      dynamic     bytes           60*n+1
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
	chunkBytes = append(chunkBytes, ck.txHashes...)
	return chunkBytes, nil
}

func (ck *Chunk) Hash() common.Hash {
	var bytes []byte
	for i := 0; i < ck.blockNum; i++ {
		bytes = append(bytes, ck.blockContext[i*60:i*60+58]...)
	}
	bytes = append(bytes, ck.txHashes...)
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
		cks.size += len(blockContext) + len(txHashes)*common.HashLength
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
	if lastChunk.blockNum+1 > MaxBlocksPerChunk || max > NormalizedRowLimit { // add a new chunk
		lastChunk.Seal()
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

func (cks *Chunks) TxPayload() []byte {
	var bytes []byte
	for _, ck := range cks.data {
		bytes = append(bytes, ck.txsPayload...)
	}
	return bytes
}

func (cks *Chunks) SealTxPayloadForBlob() []byte {
	var bytes []byte
	for _, ck := range cks.data {
		if !ck.Sealed() {
			ck.Seal()
		}
		bytes = append(bytes, ck.sealedPayload...)
	}
	return bytes
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

func (cks *Chunks) CurrentPayloadForBlobSize() int {
	var size int
	for _, ck := range cks.data {
		if ck.Sealed() {
			size += len(ck.sealedPayload)
		} else {
			size += len(ck.txsPayload) + 4
		}
	}
	return size
}

// IsChunksAppendedWithNewBlock check if a new chunk needs to be created with this new block being added.
// If yes, return the number of the zero bytes which are supposed to be added in the last chunk.
func (cks *Chunks) IsChunksAppendedWithNewBlock(blockRc types.RowConsumption) (appended bool, zeroNum int) {
	if len(cks.data) == 0 {
		return true, 0
	}
	lastChunk := cks.data[len(cks.data)-1]
	if lastChunk.blockNum+1 > MaxBlocksPerChunk {
		return true, addedZeroNum(len(lastChunk.txsPayload) + 4) // the extra 4 bytes to store the size of the txsPayload
	}
	_, max := lastChunk.accumulateRowUsages(blockRc)
	if max > NormalizedRowLimit {
		return true, addedZeroNum(len(lastChunk.txsPayload) + 4)
	}
	return false, 0
}

func addedZeroNum(size int) int {
	remainder := size % 31
	if remainder == 0 {
		return 0
	}
	return 31 - remainder
}
