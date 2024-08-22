package main

import (
	"fmt"
	"testing"

	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"morph-l2/contract/scripts/differential-testing/libraries"
)

func Test_encodeStaking(t *testing.T) {
	user := common.HexToAddress("0x0000000000000000000000000000000000000080") // user
	tmkey := common.BytesToHash(keccak256.Hash(user.Bytes()))
	blsKey, err := libraries.GenerateRandomBytes(64)
	require.NoError(t, err)
	stakerInfo := struct {
		Addr   common.Address
		TmKey  common.Hash
		BlsKey []byte
	}{
		Addr:   user,
		TmKey:  tmkey,
		BlsKey: blsKey,
	}
	packed, err := stakerInfoInputsArgs.Pack(&stakerInfo)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(packed))
}
