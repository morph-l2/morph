package types

import (
	"encoding/binary"
	"fmt"

	"morph-l2/node/zstd"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto"
)

var EmptyVersionedHash = common.HexToHash("0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")

type BatchHeader struct {
	// Encoded in BatchHeaderV0Codec
	Version                uint8
	BatchIndex             uint64
	L1MessagePopped        uint64
	TotalL1MessagePopped   uint64
	DataHash               common.Hash
	BlobVersionedHash      common.Hash
	PrevStateRoot          common.Hash
	PostStateRoot          common.Hash
	WithdrawalRoot         common.Hash
	SequencerSetVerifyHash common.Hash
	ParentBatchHash        common.Hash
	SkippedL1MessageBitmap hexutil.Bytes

	//cache
	EncodedBytes hexutil.Bytes
}

// Encode encodes the BatchHeader into RollupV2 BatchHeaderV0Codec Encoding.
func (b *BatchHeader) Encode() []byte {
	if len(b.EncodedBytes) > 0 {
		return b.EncodedBytes
	}
	batchBytes := make([]byte, 249+len(b.SkippedL1MessageBitmap))
	batchBytes[0] = b.Version
	binary.BigEndian.PutUint64(batchBytes[1:], b.BatchIndex)
	binary.BigEndian.PutUint64(batchBytes[9:], b.L1MessagePopped)
	binary.BigEndian.PutUint64(batchBytes[17:], b.TotalL1MessagePopped)
	copy(batchBytes[25:], b.DataHash[:])
	copy(batchBytes[57:], b.BlobVersionedHash[:])
	copy(batchBytes[89:], b.PrevStateRoot[:])
	copy(batchBytes[121:], b.PostStateRoot[:])
	copy(batchBytes[153:], b.WithdrawalRoot[:])
	copy(batchBytes[185:], b.SequencerSetVerifyHash[:])
	copy(batchBytes[217:], b.ParentBatchHash[:])
	copy(batchBytes[249:], b.SkippedL1MessageBitmap[:])
	b.EncodedBytes = batchBytes
	return batchBytes
}

// Hash calculates the hash of the batch header.
func (b *BatchHeader) Hash() common.Hash {
	if len(b.EncodedBytes) == 0 {
		b.Encode()
	}

	return crypto.Keccak256Hash(b.EncodedBytes)
}

// DecodeBatchHeader attempts to decode the given byte slice into a BatchHeader.
func DecodeBatchHeader(data []byte) (BatchHeader, error) {
	if len(data) < 249 {
		return BatchHeader{}, fmt.Errorf("insufficient data for BatchHeader")
	}
	b := BatchHeader{
		Version: data[0],

		BatchIndex:             binary.BigEndian.Uint64(data[1:9]),
		L1MessagePopped:        binary.BigEndian.Uint64(data[9:17]),
		TotalL1MessagePopped:   binary.BigEndian.Uint64(data[17:25]),
		DataHash:               common.BytesToHash(data[25:57]),
		BlobVersionedHash:      common.BytesToHash(data[57:89]),
		PrevStateRoot:          common.BytesToHash(data[89:121]),
		PostStateRoot:          common.BytesToHash(data[121:153]),
		WithdrawalRoot:         common.BytesToHash(data[153:185]),
		SequencerSetVerifyHash: common.BytesToHash(data[185:217]),
		ParentBatchHash:        common.BytesToHash(data[217:249]),
		SkippedL1MessageBitmap: data[249:],

		EncodedBytes: data,
	}
	return b, nil
}

type BatchData struct {
	blockContexts []byte
	l1TxHashes    []byte
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

func (cks *BatchData) TxsPayload() []byte {
	return cks.txsPayload
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
