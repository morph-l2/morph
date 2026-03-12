package types

import (
	"encoding/binary"
	"fmt"
)

func Uint64ToBigEndianBytes(value uint64) []byte {
	valueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(valueBytes, value)
	return valueBytes
}

func Uint16ToBigEndianBytes(value uint16) []byte {
	valueBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(valueBytes, value)
	return valueBytes
}

func HeightFromBlockContextBytes(blockContextBytes []byte) (uint64, error) {
	if len(blockContextBytes) != 60 {
		return 0, fmt.Errorf("wrong block context bytes length, input: %x", blockContextBytes)
	}
	return binary.BigEndian.Uint64(blockContextBytes[:8]), nil
}
