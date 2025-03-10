package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/consensus/misc/eip4844"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
)

func CancleTx(client *ethclient.Client, txHash string) error {
	// check txhash
	if txHash == "" {
		return fmt.Errorf("txHash is empty")
	}
	// get from l1
	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return fmt.Errorf("failed to get tx by hash: %w", err)
	}

	if !pending {
		return fmt.Errorf("tx not find in pending pool")
	}
	if tx == nil {
		return fmt.Errorf("tx is nil")
	}

	switch tx.Type() {
	case types.BlobTxType:
		// craft blob tx
		return fmt.Errorf("blob tx can not be cancled")
	case types.DynamicFeeTxType:
	case types.LegacyTxType:
	default:
		return fmt.Errorf("unsupport tx type: %v", tx.Type())
	}

	// return client.CancleTx(txHash)
}

func newEmptyBlobtx(tx types.BlobTx, client *ethclient.Client) types.BlobTx {

}
func newEmpty1559tx(client *ethclient.Client) types.DynamicFeeTx {

}
func newEmptyLegacyTx(client *ethclient.Client) types.LegacyTx {
}

func GetGasTipAndCap(client *ethclient.Client, chainID int64) (*big.Int, *big.Int, *big.Int, error) {
	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, nil, err
	}
	if head.BaseFee != nil {

	}
	tip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("call SuggestGasTipCap err: %w", err)
	}
	var gasFeeCap *big.Int
	if head.BaseFee != nil {
		gasFeeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		gasFeeCap = new(big.Int).Set(tip)
	}
	// calc blob fee cap
	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas)
	}
	return tip, gasFeeCap, blobFee, nil
}
