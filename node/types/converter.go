package types

import "encoding/binary"

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
