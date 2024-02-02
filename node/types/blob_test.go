package types

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/scroll-tech/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestBlobFromSealedTxPayload(t *testing.T) {
	sealedTxPayload := rand.Bytes(30)
	_, err := BlobFromSealedTxPayload(sealedTxPayload)
	require.Error(t, err)

	sealedTxPayload = rand.Bytes(MaxBlobTxPayloadSize + 1)
	_, err = BlobFromSealedTxPayload(sealedTxPayload)
	require.Error(t, err)

	sealedTxPayload = rand.Bytes(31)
	_, err = BlobFromSealedTxPayload(sealedTxPayload)
	require.NoError(t, err)

	sealedTxPayload = rand.Bytes(MaxBlobTxPayloadSize)
	_, err = BlobFromSealedTxPayload(sealedTxPayload)
	require.NoError(t, err)

	sealedTxPayload = make([]byte, 0)
	blob, err := BlobFromSealedTxPayload(sealedTxPayload)
	require.NoError(t, err)
	require.EqualValues(t, make([]byte, 4096*32), blob)

	blob, err = BlobFromSealedTxPayload(nil)
	require.NoError(t, err)
	require.EqualValues(t, make([]byte, 4096*32), blob)
}

func TestDecodeRawTxPayload(t *testing.T) {
	// empty blob
	blob := new(kzg4844.Blob)
	rawTxPayload, err := DecodeRawTxPayload(blob)
	require.NoError(t, err)
	require.EqualValues(t, 0, len(rawTxPayload))

	txsPayload10 := rand.Bytes(10)
	txsPayload27 := rand.Bytes(27)
	txsPayload58 := rand.Bytes(58)
	txsPayload59 := rand.Bytes(59)
	len10 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len10, 10)
	len27 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len27, 27)
	len58 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len58, 58)
	len59 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len59, 59)
	sealedTxPayload := bytes.NewBuffer(make([]byte, 31)) // empty chunk
	sealedTxPayload.Write(len10)                         // 10bytes chunk
	sealedTxPayload.Write(txsPayload10)
	sealedTxPayload.Write(make([]byte, 31-14))
	sealedTxPayload.Write(len27) // 27bytes chunk
	sealedTxPayload.Write(txsPayload27)
	sealedTxPayload.Write(make([]byte, 31)) // empty chunk
	sealedTxPayload.Write(len58)            // 58bytes chunk
	sealedTxPayload.Write(txsPayload58)
	sealedTxPayload.Write(len59) // 59bytes chunk
	sealedTxPayload.Write(txsPayload59)
	sealedTxPayload.Write(make([]byte, 30))

	expectedRawTxPayload := bytes.NewBuffer(txsPayload10)
	expectedRawTxPayload.Write(txsPayload27)
	expectedRawTxPayload.Write(txsPayload58)
	expectedRawTxPayload.Write(txsPayload59)

	blob, err = BlobFromSealedTxPayload(sealedTxPayload.Bytes())
	require.NoError(t, err)
	decodedRawTxPayload, err := DecodeRawTxPayload(blob)
	require.NoError(t, err)
	require.EqualValues(t, expectedRawTxPayload.Bytes(), decodedRawTxPayload)
}
