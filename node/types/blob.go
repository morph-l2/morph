package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"morph-l2/node/zstd"

	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/rlp"
)

const MaxBlobBytesSize = 4096 * 31

// MakeBlobCanonical converts the raw blob data into the canonical blob representation of 4096 BLSFieldElements.
func MakeBlobCanonical(blobBytes []byte) (b *kzg4844.Blob, err error) {
	if len(blobBytes) > MaxBlobBytesSize {
		return nil, fmt.Errorf("data is too large for blob. len=%v", len(blobBytes))
	}
	offset := 0
	b = new(kzg4844.Blob)
	// encode (up to) 31 bytes of remaining input data at a time into the subsequent field element
	for i := 0; i < 4096; i++ {
		offset += copy(b[i*32+1:i*32+32], blobBytes[offset:])
		if offset == len(blobBytes) {
			break
		}
	}
	if offset < len(blobBytes) {
		return nil, fmt.Errorf("failed to fit all data into blob. bytes remaining: %v", len(blobBytes)-offset)
	}
	return
}

func RetrieveBlobBytes(blob *kzg4844.Blob) ([]byte, error) {
	data := make([]byte, MaxBlobBytesSize)
	for i := 0; i < 4096; i++ {
		if blob[i*32] != 0 {
			return nil, fmt.Errorf("invalid blob, found non-zero high order byte %x of field element %d", data[i*32], i)
		}
		copy(data[i*31:i*31+31], blob[i*32+1:i*32+32])
	}
	return data, nil
}

func DecodeRawTxPayload(b *kzg4844.Blob) ([]byte, error) {
	data, err := RetrieveBlobBytes(b)
	if err != nil {
		return nil, err
	}

	var offset uint32
	var chunkIndex uint16
	var payload []byte
	for {
		if offset >= MaxBlobBytesSize {
			break
		}
		dataLen := binary.LittleEndian.Uint32(data[offset : offset+4])
		remainingLen := MaxBlobBytesSize - offset - 4
		if dataLen > remainingLen {
			return nil, fmt.Errorf("decode error: dataLen is bigger than remainingLen. chunkIndex: %d, dataLen: %d, remaingLen: %d", chunkIndex, dataLen, remainingLen)
		}
		payload = append(payload, data[offset+4:offset+4+dataLen]...)

		ret := (4 + dataLen) / 31
		remainder := (4 + dataLen) % 31
		if remainder > 0 {
			ret += 1
		}
		offset += ret * 31
		chunkIndex++
	}
	return payload, nil
}

func makeBCP(bz []byte) (b kzg4844.Blob, c kzg4844.Commitment, p kzg4844.Proof, err error) {
	blob, err := MakeBlobCanonical(bz)
	if err != nil {
		return
	}
	b = *blob
	c, err = kzg4844.BlobToCommitment(&b)
	if err != nil {
		return
	}
	p, err = kzg4844.ComputeBlobProof(&b, c)
	if err != nil {
		return
	}
	return
}

func MakeBlobTxSidecar(blobBytes []byte) (*eth.BlobTxSidecar, error) {
	if len(blobBytes) == 0 {
		return nil, nil
	}
	if len(blobBytes) > 2*MaxBlobBytesSize {
		return nil, errors.New("only 2 blobs at most is allowed")
	}
	blobCount := len(blobBytes)/(MaxBlobBytesSize+1) + 1
	var (
		err         error
		blobs       = make([]kzg4844.Blob, blobCount)
		commitments = make([]kzg4844.Commitment, blobCount)
		proofs      = make([]kzg4844.Proof, blobCount)
	)
	switch blobCount {
	case 1:
		blobs[0], commitments[0], proofs[0], err = makeBCP(blobBytes)
		if err != nil {
			return nil, err
		}
	case 2:
		blobs[0], commitments[0], proofs[0], err = makeBCP(blobBytes[:MaxBlobBytesSize])
		if err != nil {
			return nil, err
		}
		blobs[1], commitments[1], proofs[1], err = makeBCP(blobBytes[MaxBlobBytesSize:])
		if err != nil {
			return nil, err
		}
	}
	return &eth.BlobTxSidecar{
		Blobs:       blobs,
		Commitments: commitments,
		Proofs:      proofs,
	}, nil
}

func EncodeBatchBytesToBlob(batchBytes []byte) (*eth.BlobTxSidecar, error) {
	compressedBatchBytes, err := zstd.CompressBatchBytes(batchBytes)
	if err != nil {
		return nil, err
	}
	return MakeBlobTxSidecar(compressedBatchBytes)
}

// Deprecated: DecodeTxsFromBlob is recommended
func DecodeLegacyTxsFromBlob(b *kzg4844.Blob) (eth.Transactions, error) {
	data, err := RetrieveBlobBytes(b)
	if err != nil {
		return nil, err
	}

	// metadata || tx_payload
	// metadata consists of num_chunks (2 bytes) and chunki_size (4 bytes per chunk)
	dataReader := bytes.NewReader(data[2:])
	var txPayloadSize uint32
	for i := 0; i < 15; i++ {
		var size uint32
		if err := binary.Read(dataReader, binary.BigEndian, size); err != nil {
			return nil, err
		}
		txPayloadSize += size
	}
	txPayload := data[62 : 62+txPayloadSize]

	var byteOccupied int
	var sizeBytes []byte
	b3 := byte(txPayloadSize >> 16)
	b2 := byte(txPayloadSize >> 8)
	b1 := byte(txPayloadSize)
	if b3 > 0 {
		byteOccupied = 3
		sizeBytes = []byte{b3, b2, b1}
	} else if b2 > 0 {
		byteOccupied = 2
		sizeBytes = []byte{b2, b1}
	} else {
		byteOccupied = 1
		sizeBytes = []byte{b1}
	}

	fistByte := byte(247 + byteOccupied)
	simulatedRLP := append(append([]byte{fistByte}, sizeBytes...), txPayload...)
	decoded := make([]*eth.Transaction, 0)
	if err := rlp.DecodeBytes(simulatedRLP, &decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeTxsFromBlob(blob *kzg4844.Blob) (eth.Transactions, error) {
	data, err := RetrieveBlobBytes(blob)
	if err != nil {
		return nil, err
	}
	batchBytes, err := zstd.DecompressBatchBytes(data)
	if err != nil {
		return nil, err
	}
	data = batchBytes
	nonEmptyChunkNum := binary.BigEndian.Uint16(data[:2])
	if nonEmptyChunkNum == 0 {
		return nil, nil
	}
	// skip metadata: 2bytes(chunkNum) + maxChunks*4bytes(size per chunk)
	skipBytes := 2 + MaxChunks*4
	reader := bytes.NewReader(data[skipBytes:])
	txs := make(eth.Transactions, 0)
	for {
		var (
			firstByte   byte
			fullTxBytes []byte
			innerTx     eth.TxData
			err         error
		)
		if err = binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
			// if the blob byte array is completely consumed, then break the loop
			if err == io.EOF {
				break
			}
			return nil, err
		}
		// zero byte is found after valid tx bytes, break the loop
		if firstByte == 0 {
			break
		}

		switch firstByte {
		case eth.AccessListTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.AccessListTx)
		case eth.DynamicFeeTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.DynamicFeeTx)
		default:
			if firstByte <= 0xf7 { // legacy tx first byte must be greater than 0xf7(247)
				return nil, fmt.Errorf("not supported tx type: %d", firstByte)
			}
			innerTx = new(eth.LegacyTx)
		}

		// we support the tx types of LegacyTxType/AccessListTxType/DynamicFeeTxType
		//if firstByte == eth.AccessListTxType || firstByte == eth.DynamicFeeTxType {
		//	// the firstByte here is used to indicate tx type, so skip it
		//	if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
		//		return nil, err
		//	}
		//} else if firstByte <= 0xf7 { // legacy tx first byte must be greater than 0xf7(247)
		//	return nil, fmt.Errorf("not supported tx type: %d", firstByte)
		//}
		fullTxBytes, err = extractInnerTxFullBytes(firstByte, reader)
		if err != nil {
			return nil, err
		}
		if err = rlp.DecodeBytes(fullTxBytes, innerTx); err != nil {
			return nil, err
		}
		txs = append(txs, eth.NewTx(innerTx))
	}
	return txs, nil
}

func extractInnerTxFullBytes(firstByte byte, reader io.Reader) ([]byte, error) {
	//the occupied byte length for storing the size of the following rlp encoded bytes
	sizeByteLen := firstByte - 0xf7

	// the size of the following rlp encoded bytes
	sizeByte := make([]byte, sizeByteLen)
	if err := binary.Read(reader, binary.BigEndian, sizeByte); err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint32(append(make([]byte, 4-len(sizeByte)), sizeByte...))

	txRaw := make([]byte, size)
	if err := binary.Read(reader, binary.BigEndian, txRaw); err != nil {
		return nil, err
	}
	fullTxBytes := make([]byte, 1+uint32(sizeByteLen)+size)
	copy(fullTxBytes[:1], []byte{firstByte})
	copy(fullTxBytes[1:1+sizeByteLen], sizeByte)
	copy(fullTxBytes[1+sizeByteLen:], txRaw)

	return fullTxBytes, nil
}
