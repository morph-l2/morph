package main

import (
	"fmt"
	"os"

	"github.com/iden3/go-iden3-crypto/keccak256"

	"github.com/morph-l2/contract/scripts/differential-testing/libraries"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

// ABI types
var (
	// Prove withdrawal inputs tuple (bytes32, bytes32, bytes32, bytes32, bytes[])
	proveWithdrawalInputs, _ = abi.NewType("tuple", "ProveWithdrawalInputs", []abi.ArgumentMarshaling{
		{Name: "withdrawalHash", Type: "bytes32"},
		{Name: "withdrawalProof", Type: "bytes32[32]"},
		{Name: "withdrawalRoot", Type: "bytes32"},
	})
	proveWithdrawalInputsArgs = abi.Arguments{
		{Name: "inputs", Type: proveWithdrawalInputs},
	}

	// Prove withdrawal inputs tuple (address, bytes32, bytes)
	stakingInfoInputs, _ = abi.NewType("tuple", "StakingInfo", []abi.ArgumentMarshaling{
		{Name: "addr", Type: "address"},
		{Name: "tmKey", Type: "bytes32"},
		{Name: "blsKey", Type: "bytes"},
	})
	stakingInfoInputsArgs = abi.Arguments{
		{Name: "inputs", Type: stakingInfoInputs},
	}
)

func main() {
	args := os.Args[1:]

	// This command requires arguments
	if len(args) == 0 {
		panic("Error: No arguments provided")
	}

	switch args[0] {
	case "getTest":
		testNum := uint64(12345)
		uint64Type, _ := abi.NewType("uint64", "", nil)
		uint64Arg := abi.Arguments{
			{Type: uint64Type},
		}
		data, err := uint64Arg.Pack(&testNum)
		checkErr(err, "Error encoding output")
		fmt.Print(hexutil.Encode(data))
	case "generateStakingInfo":
		user := common.HexToAddress(args[1]) // user
		tmKey := common.BytesToHash(keccak256.Hash(user.Bytes()))
		blsKey, err := libraries.GenerateRandomBytes(256)
		checkErr(err, "Error generate staking info")

		stakerInfo := struct {
			Addr   common.Address
			TmKey  [32]byte
			BlsKey []byte
		}{
			Addr:   user,
			TmKey:  tmKey,
			BlsKey: blsKey,
		}
		packed, err := stakingInfoInputsArgs.Pack(&stakerInfo)
		checkErr(err, "Error encoding output")
		fmt.Print(hexutil.Encode(packed))
	case "getProveWithdrawalTransactionInputs":
		// Parse input arguments
		wdHash := common.HexToHash(args[1])
		smt := libraries.NewSMT(32)
		smt.Add(wdHash)
		wdProof := smt.GetProofTreeByIndex(0)
		wdRoot := smt.GetRoot()
		// Pack the proof
		Proof := struct {
			WithdrawalHash  common.Hash
			WithdrawalProof []common.Hash
			WithdrawalRoot  common.Hash
		}{
			WithdrawalHash:  wdHash,
			WithdrawalProof: wdProof,
			WithdrawalRoot:  wdRoot,
		}
		packed, err := proveWithdrawalInputsArgs.Pack(&Proof)
		checkErr(err, "Error encoding output")
		// Print the output
		fmt.Print(hexutil.Encode(packed[:]))
	default:
		panic(fmt.Errorf("Unknown command: %s", args[0]))
	}
}
