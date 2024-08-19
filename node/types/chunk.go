package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"morph-l2/node/zstd"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
)

const (
	NormalizedRowLimit = 1_000_000
	MaxBlocksPerChunk  = 100
	MaxChunks          = 45
)

type Chunk struct {
	blockContext  []byte
	l1TxHashes    []byte
	accumulatedRc types.RowConsumption
	blockNum      int
	txsPayload    []byte // the raw txs payload
}

func NewChunk(blockContext, txsPayload []byte, l1TxHashes []common.Hash, rc types.RowConsumption) *Chunk {
	var l1TxHashBytes []byte
	for _, txHash := range l1TxHashes {
		l1TxHashBytes = append(l1TxHashBytes, txHash.Bytes()...)
	}
	return &Chunk{
		blockContext:  blockContext,
		txsPayload:    txsPayload,
		l1TxHashes:    l1TxHashBytes,
		accumulatedRc: rc,
		blockNum:      1,
	}
}

func (ck *Chunk) append(blockContext, txsPayload []byte, l1TxHashes []common.Hash, accRc types.RowConsumption) {
	ck.blockContext = append(ck.blockContext, blockContext...)
	ck.txsPayload = append(ck.txsPayload, txsPayload...)
	ck.accumulatedRc = accRc
	ck.blockNum++
	for _, txHash := range l1TxHashes {
		ck.l1TxHashes = append(ck.l1TxHashes, txHash.Bytes()...)
	}
}

func (ck *Chunk) accumulateRowUsages(rc types.RowConsumption) (accRc types.RowConsumption, max uint64) {
	if len(ck.accumulatedRc) == 0 {
		return rc, maxRowNumber(rc)
	}
	if len(rc) == 0 {
		accRc = make(types.RowConsumption, len(ck.accumulatedRc))
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

func (ck *Chunk) BlockContext() []byte {
	return ck.blockContext
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
	return chunkBytes, nil
}

func (ck *Chunk) Decode(chunkBytes []byte) error {
	reader := bytes.NewReader(chunkBytes)
	var blockNum uint8
	if err := binary.Read(reader, binary.BigEndian, &blockNum); err != nil {
		return err
	}

	bcs := make([]byte, 0)
	for i := 0; i < int(blockNum); i++ {
		bc := make([]byte, 60)
		if err := binary.Read(reader, binary.BigEndian, &bc); err != nil {
			return err
		}
		bcs = append(bcs, bc...)
	}
	ck.blockContext = bcs
	ck.blockNum = int(blockNum)
	return nil
}

func (ck *Chunk) Hash() common.Hash {
	var bz []byte
	for i := 0; i < ck.blockNum; i++ {
		bz = append(bz, ck.blockContext[i*60:i*60+58]...)
	}
	bz = append(bz, ck.l1TxHashes...)
	return crypto.Keccak256Hash(bz)
}

type Chunks struct {
	data           []*Chunk
	blockNum       int
	sizeInCalldata int
	blobPayload    []byte

	hash *common.Hash
}

func NewChunks() *Chunks {
	return &Chunks{
		data:        make([]*Chunk, 0),
		blobPayload: make([]byte, 0),
	}
}

func (cks *Chunks) Append(blockContext, txsPayload []byte, l1TxHashes []common.Hash, rc types.RowConsumption) {
	if cks == nil {
		return
	}
	defer func() {
		cks.sizeInCalldata += len(blockContext)
		cks.blockNum++
		cks.hash = nil // clear hash when data is updated
	}()

	if len(cks.data) == 0 {
		cks.blobPayload = cks.appendBlobBytes(txsPayload, true)
		cks.data = append(cks.data, NewChunk(blockContext, txsPayload, l1TxHashes, rc))
		cks.sizeInCalldata += 1
		return
	}
	lastChunk := cks.data[len(cks.data)-1]
	accRc, maxRowUsages := lastChunk.accumulateRowUsages(rc)
	if lastChunk.blockNum+1 > MaxBlocksPerChunk || maxRowUsages > NormalizedRowLimit { // add a new chunk
		cks.blobPayload = cks.appendBlobBytes(txsPayload, true)
		cks.data = append(cks.data, NewChunk(blockContext, txsPayload, l1TxHashes, rc))
		cks.sizeInCalldata += 1
		return
	}

	cks.blobPayload = cks.appendBlobBytes(txsPayload, false)
	lastChunk.append(blockContext, txsPayload, l1TxHashes, accRc)
}

func (cks *Chunks) Encode() ([][]byte, error) {
	var bytesArray [][]byte
	for _, ck := range cks.data {
		ckBytes, err := ck.Encode()
		if err != nil {
			return nil, err
		}
		bytesArray = append(bytesArray, ckBytes)
	}
	return bytesArray, nil
}

func (cks *Chunks) appendBlobBytes(txsPayload []byte, appendChunk bool) []byte {
	blobBytes := make([]byte, len(cks.blobPayload))
	copy(blobBytes, cks.blobPayload)

	if len(blobBytes) == 0 && !appendChunk {
		panic("Incorrect state. Chunk has not been appended while blobPayload is empty")
	}
	if len(cks.data) == MaxChunks && appendChunk {
		panic(fmt.Errorf("can not append chunk up to more than %d", MaxChunks))
	}

	if len(blobBytes) == 0 {
		// metadata consists of num_chunks (2 bytes) and chunki_size (4 bytes per chunk)
		metadataLength := 2 + MaxChunks*4

		// the raw (un-padded) blob payload
		blobBytes = make([]byte, metadataLength)
	}
	if appendChunk {
		// update chunk num
		preChunkNum := binary.BigEndian.Uint16(blobBytes[:2])
		binary.BigEndian.PutUint16(blobBytes[0:], preChunkNum+1)

		if len(txsPayload) > 0 {
			// update new chunk size
			newChunkIndex := len(cks.data)
			binary.BigEndian.PutUint32(blobBytes[2+newChunkIndex*4:], uint32(len(txsPayload)))

			// append blob
			blobBytes = append(blobBytes, txsPayload...)
		}
	} else if len(txsPayload) > 0 { // update chunk size and append payload
		var chunkIndex int
		if len(cks.data) == 0 {
			chunkIndex = 0
		} else {
			chunkIndex = len(cks.data) - 1
		}
		preChunkSize := binary.BigEndian.Uint32(blobBytes[2+chunkIndex*4:])
		binary.BigEndian.PutUint32(blobBytes[2+chunkIndex*4:], uint32(len(txsPayload))+preChunkSize)

		// append blob
		blobBytes = append(blobBytes, txsPayload...)
	}
	return blobBytes
}

func (cks *Chunks) ConstructBlobPayload() []byte {
	if len(cks.blobPayload) > 0 {
		return cks.blobPayload
	}

	// metadata consists of num_chunks (2 bytes) and chunki_size (4 bytes per chunk)
	metadataLength := 2 + MaxChunks*4

	// the raw (un-padded) blob payload
	blobBytes := make([]byte, metadataLength)

	// the number of chunks that contain at least one L2 transaction
	for i, chunk := range cks.data {
		chunkSize := len(chunk.txsPayload)
		if chunkSize != 0 {
			blobBytes = append(blobBytes, chunk.txsPayload...)
		}
		// blob metadata: chunki_size
		binary.BigEndian.PutUint32(blobBytes[2+4*i:], uint32(chunkSize))
	}
	// blob metadata: num_chunks
	binary.BigEndian.PutUint16(blobBytes[0:], uint16(len(cks.data)))
	return blobBytes
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

func (cks *Chunks) BlockNum() int       { return cks.blockNum }
func (cks *Chunks) ChunkNum() int       { return len(cks.data) }
func (cks *Chunks) SizeInCalldata() int { return cks.sizeInCalldata }

// isChunksAppendedWithNewBlock check if a new chunk needs to be created with this new block being added.
func (cks *Chunks) isChunksAppendedWithNewBlock(blockRc types.RowConsumption) (appended bool) {
	if len(cks.data) == 0 {
		return true
	}
	lastChunk := cks.data[len(cks.data)-1]
	if lastChunk.blockNum+1 > MaxBlocksPerChunk {
		return true
	}
	_, maxRowUsages := lastChunk.accumulateRowUsages(blockRc)
	return maxRowUsages > NormalizedRowLimit
}

func (cks *Chunks) EstimateCompressedSizeWithNewPayload(txPayload []byte, blockRc types.RowConsumption) (appendChunk, sizeOverflow bool, err error) {
	appendChunk = cks.isChunksAppendedWithNewBlock(blockRc)
	if appendChunk && len(cks.data) == MaxChunks {
		// return sizeOverflow as true to notify the upper stream to seal the batch, to prevent the batch from involving over MaxChunks
		return appendChunk, true, err
	}
	blobBytes := cks.appendBlobBytes(txPayload, appendChunk)
	compressed, err := zstd.CompressBatchBytes(blobBytes)
	if err != nil {
		return false, false, err
	}
	if len(compressed) > MaxBlobBytesSize {
		sizeOverflow = true
	}
	return
}
