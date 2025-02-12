package services

import (
	"crypto/sha256"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/params"
)

const (
	// Geth requires a minimum fee bump of 10% for regular tx resubmission, 100% for blob txs
	priceBump     int64 = 10
	blobPriceBump int64 = 100
)

// new = old * (100 + priceBump) / 100
var (
	priceBumpPercent     = big.NewInt(100 + priceBump)
	blobPriceBumpPercent = big.NewInt(100 + blobPriceBump)
	oneHundred           = big.NewInt(100)
)

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

func calcFee(receipt *types.Receipt) float64 {
	if receipt == nil || receipt.EffectiveGasPrice == nil {
		return 0
	}

	calldatafee := new(big.Int).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))
	// blobfee
	blobfee := big.NewInt(0)
	if receipt.Type == types.BlobTxType {
		if receipt.BlobGasPrice != nil {
			blobfee = new(big.Int).Mul(big.NewInt(int64(receipt.BlobGasUsed)), receipt.BlobGasPrice)
		}
	}

	fee := new(big.Int).Add(calldatafee, blobfee)
	return ToEtherFloat(fee)
}

func ToEtherFloat(weiAmt *big.Int) float64 {
	if weiAmt == nil {
		return 0
	}
	etherAmt := new(big.Rat).SetFrac(weiAmt, big.NewInt(params.Ether))
	fEtherAmt, _ := etherAmt.Float64()
	return fEtherAmt
}
