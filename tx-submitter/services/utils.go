package services

import (
	"crypto/sha256"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto/kzg4844"
	"github.com/scroll-tech/go-ethereum/params"
)

const (
	// Geth requires a minimum fee bump of 10% for regular tx resubmission, 100% for blob txs
	priceBump     int64 = 10
	blobPriceBump int64 = 100
)

// new = old * (100 + priceBump) / 100
var priceBumpPercent = big.NewInt(100 + priceBump)
var blobPriceBumpPercent = big.NewInt(100 + blobPriceBump)
var oneHundred = big.NewInt(100)
var two = big.NewInt(2)

var blobCommitmentVersionKZG uint8 = 0x01

// kZGToVersionedHash implements kzg_to_versioned_hash from EIP-4844
func kZGToVersionedHash(kzg kzg4844.Commitment) common.Hash {
	h := sha256.Sum256(kzg[:])
	h[0] = blobCommitmentVersionKZG

	return h
}

// calcThresholdValue returns x * priceBumpPercent / 100
func calcThresholdValue(x *big.Int, isBlobTx bool) *big.Int {
	var percent *big.Int
	if isBlobTx {
		percent = blobPriceBumpPercent
	} else {
		percent = priceBumpPercent
	}
	threshold := new(big.Int).Mul(percent, x)
	threshold = threshold.Div(threshold, oneHundred)
	return threshold
}

// for rollup
func RoughEstimateGas(msgcnt uint64) uint64 {
	base := uint64(400_000)
	gasPerMsg := uint64(4200)

	return (base + msgcnt*gasPerMsg) * 12 / 10
}
func calcFee(tx types.Transaction, receipt *types.Receipt) float64 {
	calldatafee := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(receipt.GasUsed)))
	// blobfee
	blobfee := big.NewInt(0)
	if tx.BlobTxSidecar() != nil {
		if receipt.BlobGasPrice == nil {
			return 0
		}
		blobfee = new(big.Int).Mul(big.NewInt(int64(receipt.BlobGasUsed)), receipt.BlobGasPrice)
	}

	fee := new(big.Int).Add(calldatafee, blobfee)
	feeEther := new(big.Rat).SetFrac(fee, big.NewInt(params.Ether))
	fEtherFee, _ := feeEther.Float64()
	return fEtherFee
}
