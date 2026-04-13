package batch

import (
	"encoding/binary"
	"fmt"

	"morph-l2/node/zstd"
	"morph-l2/tx-submitter/types"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
)

var (
	EmptyVersionedHash = common.HexToHash("0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")
)

type BatchData struct {
	blockContexts []byte
	l1TxHashes    []byte
	l1TxNum       uint16
	blockNum      uint16
	txsPayload    []byte

	hash *common.Hash
}

func NewBatchData() *BatchData {
	return &BatchData{
		blockContexts: make([]byte, 0),
		l1TxHashes:    make([]byte, 0),
		txsPayload:    make([]byte, 0),
	}
}

func (cks *BatchData) Append(blockContext, txsPayload []byte, l1TxHashes []common.Hash) {
	if cks == nil {
		return
	}
	cks.blockContexts = append(cks.blockContexts, blockContext...)
	cks.txsPayload = append(cks.txsPayload, txsPayload...)
	cks.blockNum++
	for _, txHash := range l1TxHashes {
		cks.l1TxHashes = append(cks.l1TxHashes, txHash.Bytes()...)
	}
	cks.l1TxNum += uint16(len(l1TxHashes))
}

// Encode encodes the data into bytes
// Below is the encoding, total 60*n+1+m bytes.
// Field           Bytes       Type            Index       Comments
// numBlocks       2           uint16          0           The number of blocks in this chunk
// block[0]        60          BlockContext    1           The first block in this chunk
// ......
// block[i]        60          BlockContext    60*i+1      The (i+1)'th block in this chunk
// ......
// block[n-1]      60          BlockContext    60*n-59     The last block in this chunk
func (cks *BatchData) Encode() ([]byte, error) {
	if cks == nil || cks.blockNum == 0 {
		return []byte{}, nil
	}

	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, cks.blockNum)
	data = append(data, cks.blockContexts...)
	return data, nil
}

func (cks *BatchData) IsEmpty() bool {
	return cks == nil || len(cks.blockContexts) == 0
}

func (cks *BatchData) DataHash() common.Hash {
	if cks.hash != nil {
		return *cks.hash
	}

	var bz []byte
	for i := 0; i < int(cks.blockNum); i++ {
		bz = append(bz, cks.blockContexts[i*60:i*60+58]...)
	}
	bz = append(bz, cks.l1TxHashes...)
	return crypto.Keccak256Hash(bz)
}

// DataHashV2 computes the Keccak-256 hash of the batch data, incorporating
// the last block height, L1 transaction count, and L1 transaction hashes.
func (cks *BatchData) DataHashV2() (common.Hash, error) {
	// Validate blockContexts length
	if len(cks.blockContexts) < 60 {
		return common.Hash{}, fmt.Errorf("blockContexts too short, length: %d", len(cks.blockContexts))
	}

	// Extract the last 60 bytes
	lastBlockContext := cks.blockContexts[len(cks.blockContexts)-60:]

	// Parse block height
	height, err := types.HeightFromBlockContextBytes(lastBlockContext)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse blockContext: context length=%d, lastBlockContext=%x, err=%w",
			len(cks.blockContexts), lastBlockContext, err)
	}

	// Compute the hash
	return cks.calculateHash(height), nil
}

func (cks *BatchData) calculateHash(height uint64) common.Hash {
	// Preallocate memory for efficiency
	hashData := make([]byte, 8+2+len(cks.l1TxHashes)) // 8 bytes for height, 2 bytes for l1TxNum
	copy(hashData[:8], types.Uint64ToBigEndianBytes(height))
	copy(hashData[8:10], types.Uint16ToBigEndianBytes(cks.l1TxNum))
	copy(hashData[10:], cks.l1TxHashes)

	return crypto.Keccak256Hash(hashData)
}

func (cks *BatchData) TxsPayload() []byte {
	return cks.txsPayload
}

// TxsPayloadV2 returns the bytes combining the block contexts with the tx payload
func (cks *BatchData) TxsPayloadV2() []byte {
	return append(cks.blockContexts, cks.txsPayload...)
}

func (cks *BatchData) BlockNum() uint16 { return cks.blockNum }

func (cks *BatchData) EstimateCompressedSizeWithNewPayload(txPayload []byte) (bool, error) {
	blobBytes := append(cks.txsPayload, txPayload...)
	if len(blobBytes) <= MaxBlobBytesSize {
		return false, nil
	}
	compressed, err := zstd.CompressBatchBytes(blobBytes)
	if err != nil {
		return false, err
	}
	return len(compressed) > MaxBlobBytesSize, nil
}

func (cks *BatchData) combinePayloads(newBlockContext, newTxPayload []byte) []byte {
	totalLength := len(cks.blockContexts) + len(newBlockContext) + len(cks.txsPayload) + len(newTxPayload)
	combined := make([]byte, totalLength)
	copy(combined, cks.blockContexts)
	copy(combined[len(cks.blockContexts):], newBlockContext)
	copy(combined[len(cks.blockContexts)+len(newBlockContext):], cks.txsPayload)
	copy(combined[len(cks.blockContexts)+len(newBlockContext)+len(cks.txsPayload):], newTxPayload)
	return combined
}

// WillExceedCompressedSizeLimit checks if the size of the combined block contexts
// and transaction payloads (after compression) exceeds the maximum allowed size.
func (cks *BatchData) WillExceedCompressedSizeLimit(newBlockContext, newTxPayload []byte) (bool, error) {
	// Combine the existing and new block contexts and transaction payloads
	combinedBytes := cks.combinePayloads(newBlockContext, newTxPayload)
	if len(combinedBytes) <= MaxBlobBytesSize {
		return false, nil
	}
	compressed, err := zstd.CompressBatchBytes(combinedBytes)
	if err != nil {
		return false, fmt.Errorf("compression failed: %w", err)
	}
	return len(compressed) > MaxBlobBytesSize, nil
}
