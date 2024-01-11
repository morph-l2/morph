package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/morph-l2/contract/scripts/differential-testing/libraries"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

var UnknownNonceVersion = errors.New("Unknown nonce version")

// checkOk checks if ok is false, and panics if so.
// Shorthand to ease go's god awful error handling
func checkOk(ok bool) {
	if !ok {
		panic(fmt.Errorf("checkOk failed"))
	}
}

// checkErr checks if err is not nil, and throws if so.
// Shorthand to ease go's god awful error handling
func checkErr(err error, failReason string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", failReason, err))
	}
}

// encodeCrossDomainMessage encodes a versioned cross domain message into a byte array.
func encodeCrossDomainMessage(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) ([]byte, error) {
	_, version := libraries.DecodeVersionedNonce(nonce)

	var encoded []byte
	var err error
	if version.Cmp(big.NewInt(0)) == 0 {
		// Encode cross domain message V0
		encoded, err = libraries.EncodeCrossDomainMessageV0(target, sender, data, nonce)
	} else if version.Cmp(big.NewInt(1)) == 0 {
		// Encode cross domain message V1
		encoded, err = libraries.EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
	} else {
		return nil, UnknownNonceVersion
	}

	return encoded, err
}

// hashWithdrawal hashes a withdrawal transaction.
func hashWithdrawal(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) (common.Hash, error) {
	wd := libraries.Withdrawal{
		Nonce:    nonce,
		Sender:   &sender,
		Target:   &target,
		Value:    value,
		GasLimit: gasLimit,
		Data:     data,
	}
	return wd.Hash()
}

//// hashOutputRootProof hashes an output root proof.
//func hashOutputRootProof(version common.Hash, stateRoot common.Hash, messagePasserStorageRoot common.Hash, latestBlockHash common.Hash) (common.Hash, error) {
//	hash, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
//		Version:                  version,
//		StateRoot:                stateRoot,
//		MessagePasserStorageRoot: messagePasserStorageRoot,
//		LatestBlockhash:          latestBlockHash,
//	})
//	if err != nil {
//		return common.Hash{}, err
//	}
//	return common.Hash(hash), nil
//}

// makeDepositTx creates a deposit transaction type.
func makeDepositTx(
	queueIndex uint64,
	gas uint64,
	to *common.Address,
	value *big.Int,
	data []byte,
	sender common.Address,
) types.L1MessageTx {
	tx := types.L1MessageTx{
		queueIndex,
		gas,
		to,
		value,
		data,
		sender,
	}
	return tx
}
