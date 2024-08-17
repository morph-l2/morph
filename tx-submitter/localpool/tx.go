package localpool

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/core/types"
)

func EncodeTx(tx *types.Transaction) (string, error) {
	txbs, err := tx.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("failed to marshal tx: %w", err)
	}
	return string(txbs), nil
}
func ParseTx(tx string) (*types.Transaction, error) {
	var res types.Transaction
	if err := res.UnmarshalJSON([]byte(tx)); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tx: %w", err)
	}
	return &res, nil
}
