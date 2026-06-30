package zstd

import (
	"bytes"
	"fmt"

	kzstd "github.com/klauspost/compress/zstd"
	decoderzstd "github.com/morph-l2/morph-da-codec/bindings/decoder/zstd"
)

var magics = []byte{0x28, 0xb5, 0x2f, 0xfd}

func CompressBatchBytes(batchBytes []byte) ([]byte, error) {
	if len(batchBytes) == 0 {
		return nil, nil
	}

	encoder, err := kzstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}
	defer encoder.Close()

	compressed := encoder.EncodeAll(batchBytes, nil)
	if !bytes.HasPrefix(compressed, magics) {
		return nil, fmt.Errorf("compressed batch bytes missing zstd magic")
	}
	return compressed[len(magics):], nil
}

// DecompressBatchBytes decompresses the given bytes into batch bytes
func DecompressBatchBytes(compressedBytes []byte) ([]byte, error) {
	return decoderzstd.DecompressMorphDABatch(compressedBytes)
}
