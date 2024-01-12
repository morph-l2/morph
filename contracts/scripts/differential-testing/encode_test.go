package main

import (
	"fmt"
	"testing"

	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/morph-l2/contract/scripts/differential-testing/libraries"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

func Test_encodeStaking(t *testing.T) {
	user := common.HexToAddress("0x0000000000000000000000000000000000000080") // user
	tmkey := common.BytesToHash(keccak256.Hash(user.Bytes()))
	blsKey, err := libraries.GenerateRandomBytes(64)
	checkErr(err, "Error generate staking info")
	stakerInfo := struct {
		Addr   common.Address
		TmKey  common.Hash
		BlsKey []byte
	}{
		Addr:   user,
		TmKey:  tmkey,
		BlsKey: blsKey,
	}
	packed, err := stakingInfoInputsArgs.Pack(&stakerInfo)
	checkErr(err, "Error encoding output")
	fmt.Println(hexutil.Encode(packed))
}
