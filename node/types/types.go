package types

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

var Version1StartedNonce = new(big.Int).SetBytes(hexutil.MustDecode("0x1000000000000000000000000000000000000000000000000000000000000"))

func EncodeNonce(nonce uint64) *big.Int {
	return new(big.Int).Add(Version1StartedNonce, big.NewInt(int64(nonce)))
}

func DecodeNonce(encodedNonce *big.Int) (uint64, error) {
	decoded := new(big.Int).Sub(encodedNonce, Version1StartedNonce)
	if decoded.Sign() < 0 {
		return 0, errors.New(fmt.Sprintf("wrong encoded nonce: %s", encodedNonce.String()))
	}
	return decoded.Uint64(), nil
}

func IsAllZero(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
