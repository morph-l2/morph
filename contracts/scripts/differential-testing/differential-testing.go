package main

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"os"

	"github.com/iden3/go-iden3-crypto/keccak256"

	"github.com/morph-l2/contract/scripts/differential-testing/libraries"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
)

// ABI types
var (
	// Plain dynamic dynBytes type
	dynBytes, _ = abi.NewType("bytes", "", nil)
	bytesArgs   = abi.Arguments{
		{Type: dynBytes},
	}

	// Plain fixed bytes32 type
	fixedBytes, _  = abi.NewType("bytes32", "", nil)
	fixedBytesArgs = abi.Arguments{
		{Type: fixedBytes},
	}

	// Decoded nonce tuple (nonce, version)
	decodedNonce, _ = abi.NewType("tuple", "DecodedNonce", []abi.ArgumentMarshaling{
		{Name: "nonce", Type: "uint256"},
		{Name: "version", Type: "uint256"},
	})
	decodedNonceArgs = abi.Arguments{
		{Name: "encodedNonce", Type: decodedNonce},
	}

	// Prove withdrawal inputs tuple (bytes32, bytes32, bytes32, bytes32, bytes[])
	proveWithdrawalInputs, _ = abi.NewType("tuple", "ProveWithdrawalInputs", []abi.ArgumentMarshaling{
		{Name: "withdrawalHash", Type: "bytes32"},
		{Name: "withdrawalProof", Type: "bytes32[32]"},
		{Name: "withdrawalRoot", Type: "bytes32"},
	})
	proveWithdrawalInputsArgs = abi.Arguments{
		{Name: "inputs", Type: proveWithdrawalInputs},
	}

	// Prove withdrawal inputs tuple (bytes32, bytes32, bytes32, bytes32, bytes[])
	batchDataInputs, _ = abi.NewType("tuple", "BatchDataInputs", []abi.ArgumentMarshaling{
		{Name: "blockNumber", Type: "uint64"},
		{Name: "transactions", Type: "bytes"},
		{Name: "blockWitness", Type: "bytes"},
		{Name: "preStateRoot", Type: "bytes32"},
		{Name: "postStateRoot", Type: "bytes32"},
		{Name: "withdrawalRoot", Type: "bytes32"},
		{Name: "signature", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "version", Type: "uint256"},
			{Name: "signers", Type: "uint256[]"},
			{Name: "signature", Type: "bytes"},
		}},
	})
	batchDataInputsArgs = abi.Arguments{
		{Name: "inputs", Type: batchDataInputs},
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
	case "generateBatchData":
		blockNum, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		preStateRoot := common.HexToHash(args[2])
		withdrawalRoot := common.HexToHash(args[3])

		bkWitness := make([]byte, 0)
		bkWitness = binary.BigEndian.AppendUint64(bkWitness, blockNum.Uint64())
		//Pack the proof
		type RollupBatchSignature struct {
			Version   *big.Int
			Signers   []*big.Int
			Signature []byte
		}
		BatchData := struct {
			BlockNumber    uint64
			Transactions   []byte
			BlockWitness   []byte
			PreStateRoot   [32]byte
			PostStateRoot  [32]byte
			WithdrawalRoot [32]byte
			Signature      RollupBatchSignature
		}{
			BlockNumber:    blockNum.Uint64(),
			Transactions:   nil,
			BlockWitness:   bkWitness,
			PreStateRoot:   preStateRoot,
			PostStateRoot:  preStateRoot,
			WithdrawalRoot: withdrawalRoot,
			Signature: RollupBatchSignature{
				common.Big0, nil, nil,
			},
		}
		packed, err := batchDataInputsArgs.Pack(&BatchData)
		checkErr(err, "Error encoding output")
		fmt.Print(hexutil.Encode(packed))
	case "decodeVersionedNonce":
		// Parse input arguments
		input, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)

		// Decode versioned nonce
		nonce, version := libraries.DecodeVersionedNonce(input)

		// ABI encode output
		packArgs := struct {
			Nonce   *big.Int
			Version *big.Int
		}{
			nonce,
			version,
		}
		packed, err := decodedNonceArgs.Pack(&packArgs)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "encodeCrossDomainMessage":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Encode cross domain message
		encoded, err := encodeCrossDomainMessage(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error encoding cross domain message")

		// Pack encoded cross domain message
		packed, err := bytesArgs.Pack(&encoded)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashCrossDomainMessage":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Encode cross domain message
		encoded, err := encodeCrossDomainMessage(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error encoding cross domain message")

		// Hash encoded cross domain message
		hash := crypto.Keccak256Hash(encoded)

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashL1MessageTx":
		queueIndex, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		gas, ok := new(big.Int).SetString(args[2], 10)
		checkOk(ok)
		to := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		data := common.FromHex(args[5])
		sender := common.HexToAddress(args[6])
		// Create deposit transaction
		depositTx := makeDepositTx(queueIndex.Uint64(), gas.Uint64(), &to, value, data, sender)
		// RLP encode deposit transaction
		encoded, err := types.NewTx(&depositTx).MarshalBinary()
		checkErr(err, "Error encoding deposit transaction")

		// Hash encoded deposit transaction
		hash := crypto.Keccak256Hash(encoded)

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "encodeL1MessageTx":
		// Parse input arguments
		queueIndex, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		gas, ok := new(big.Int).SetString(args[2], 10)
		checkOk(ok)
		to := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		data := common.FromHex(args[5])
		sender := common.HexToAddress(args[6])
		// Create deposit transaction
		depositTx := makeDepositTx(queueIndex.Uint64(), gas.Uint64(), &to, value, data, sender)

		// RLP encode deposit transaction
		encoded, err := types.NewTx(&depositTx).MarshalBinary()
		checkErr(err, "Failed to RLP encode deposit transaction")
		// Pack rlp encoded deposit transaction
		packed, err := bytesArgs.Pack(&encoded)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashWithdrawal":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Hash withdrawal
		hash, err := hashWithdrawal(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error hashing withdrawal")

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "getProveWithdrawalTransactionInputs":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		wdHash, err := hashWithdrawal(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error hashing withdrawal")
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
