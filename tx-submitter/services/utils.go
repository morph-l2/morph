package services

import (
	"crypto/sha256"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/params"
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

// calcFee calculates the total fee for a transaction
func calcFee(tx *ethtypes.Transaction, receipt *ethtypes.Receipt) *big.Float {
	gasUsed := new(big.Float).SetUint64(receipt.GasUsed)
	effectiveGasPrice := new(big.Float).SetInt(receipt.EffectiveGasPrice)
	txFee := new(big.Float).Mul(gasUsed, effectiveGasPrice)
	txFeeEth := new(big.Float).Quo(txFee, new(big.Float).SetInt(big.NewInt(params.Ether)))

	// Add blob fee if it's a blob transaction
	if tx.Type() == ethtypes.BlobTxType {
		blobGasUsed := new(big.Float).SetUint64(tx.BlobGas())
		blobGasPrice := new(big.Float).SetInt(tx.BlobGasFeeCap())
		blobFee := new(big.Float).Mul(blobGasUsed, blobGasPrice)
		blobFeeEth := new(big.Float).Quo(blobFee, new(big.Float).SetInt(big.NewInt(params.Ether)))
		txFeeEth.Add(txFeeEth, blobFeeEth)
	}

	return txFeeEth
}

func ToEtherFloat(weiAmt *big.Int) float64 {
	if weiAmt == nil {
		return 0
	}
	etherAmt := new(big.Rat).SetFrac(weiAmt, big.NewInt(params.Ether))
	fEtherAmt, _ := etherAmt.Float64()
	return fEtherAmt

}
