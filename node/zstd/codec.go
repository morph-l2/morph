package zstd

/*
#include <stdint.h>
char* compress_scroll_batch_bytes(uint8_t* src, uint64_t src_size, uint8_t* output_buf, uint64_t *output_buf_size);
*/
import "C"

import (
	"bytes"
	"fmt"
	"unsafe"

	"github.com/klauspost/compress/zstd"
)

var magics = []byte{0x28, 0xb5, 0x2f, 0xfd}

func CompressBatchBytes(bytes []byte) ([]byte, error) {
	srcSize := C.uint64_t(len(bytes))
	outbufSize := C.uint64_t(len(bytes) + 128) // Allocate output buffer with extra 128 bytes
	outbuf := make([]byte, outbufSize)

	if err := C.compress_scroll_batch_bytes((*C.uchar)(unsafe.Pointer(&bytes[0])), srcSize,
		(*C.uchar)(unsafe.Pointer(&outbuf[0])), &outbufSize); err != nil {
		return nil, fmt.Errorf("failed to compress batch bytes: %s", C.GoString(err))
	}

	return outbuf[:int(outbufSize)], nil
}

// DecompressBatchBytes decompresses the given bytes into batch bytes
func DecompressBatchBytes(compressedBytes []byte) ([]byte, error) {
	// add magics
	data := make([]byte, len(compressedBytes)+len(magics))
	copy(data, magics)
	copy(data[len(magics):], compressedBytes)

	// decompress data in stream and in batches of bytes, because we don't know actual length of compressed data
	var res []byte
	readBatchSize := 131072
	batchOfBytes := make([]byte, readBatchSize)

	r := bytes.NewReader(data)
	zr, err := zstd.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	for {
		i, _ := zr.Read(batchOfBytes) //
		res = append(res, batchOfBytes[:i]...)
		if i < readBatchSize {
			break
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("failed to decompress blob bytes")
	}
	return res, nil
}
