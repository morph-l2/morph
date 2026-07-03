package zstd

import codeczstd "github.com/morph-l2/morph-da-codec/bindings/codec/zstd"

func CompressBatchBytes(batchBytes []byte) ([]byte, error) {
	if len(batchBytes) == 0 {
		return nil, nil
	}
	return codeczstd.CompressMorphDABatch(batchBytes)
}

// DecompressBatchBytes decompresses the given bytes into batch bytes
func DecompressBatchBytes(compressedBytes []byte) ([]byte, error) {
	return codeczstd.DecompressMorphDABatch(compressedBytes)
}
