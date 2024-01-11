package libraries

import (
	"encoding/binary"
)

type RLPWriter struct{}

func (r *RLPWriter) WriteBytes(input []byte) []byte {
	if len(input) == 1 && input[0] < 128 {
		return input // 单个字节的数据可以直接返回
	}
	return append(r.WriteLength(uint64(len(input)), 128), input...)
}

func (r *RLPWriter) WriteList(input [][]byte) []byte {
	list := r.flatten(input)
	return append(r.WriteLength(uint64(len(list)), 192), list...)
}

func (r *RLPWriter) WriteString(input string) []byte {
	return r.WriteBytes([]byte(input))
}

func (r *RLPWriter) WriteAddress(input [20]byte) []byte {
	return r.WriteBytes(input[:])
}

func (r *RLPWriter) WriteUint(input uint64) []byte {
	return r.WriteBytes(r.toBinary(input))
}

func (r *RLPWriter) WriteBool(input bool) []byte {
	encoded := make([]byte, 1)
	if input {
		encoded[0] = byte(0x01)
	} else {
		encoded[0] = byte(0x80)
	}
	return encoded
}

func (r *RLPWriter) WriteLength(length uint64, offset uint64) []byte {
	if length < 56 {
		return []byte{byte(length) + byte(offset)}
	}

	var encoded []byte
	lenLen := 0
	i := uint64(1)
	for length/i != 0 {
		lenLen++
		i *= 256
	}

	encoded = make([]byte, lenLen+1)
	encoded[0] = byte(lenLen) + byte(offset) + 55
	for i = 1; i <= uint64(lenLen); i++ {
		encoded[i] = byte((length / (256 * (uint64(lenLen) - i + 1))) % 256)
	}

	return encoded
}

func (r *RLPWriter) toBinary(x uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, x)

	i := 0
	for i < 8 && b[i] == 0 {
		i++
	}

	return b[i:]
}

func (r *RLPWriter) flatten(list [][]byte) []byte {
	if len(list) == 0 {
		return []byte{}
	}

	var lenSum uint64
	for _, item := range list {
		lenSum += uint64(len(item))
	}

	flattened := make([]byte, lenSum)
	flattenedPtr := 0

	for _, item := range list {
		copy(flattened[flattenedPtr:], item)
		flattenedPtr += len(item)
	}

	return flattened
}
