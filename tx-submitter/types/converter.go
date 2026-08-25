package types

import "morph-l2/common/batch"

func Uint64ToBigEndianBytes(value uint64) []byte {
	return batch.Uint64ToBigEndianBytes(value)
}

func Uint16ToBigEndianBytes(value uint16) []byte {
	return batch.Uint16ToBigEndianBytes(value)
}

func HeightFromBlockContextBytes(blockContextBytes []byte) (uint64, error) {
	return batch.HeightFromBlockContextBytes(blockContextBytes)
}
