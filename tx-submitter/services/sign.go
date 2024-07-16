package services

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/core/types"
)

func (r *Rollup) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if r.cfg.ExternalSign {
		//todo
		return nil, fmt.Errorf("not support external sign yet")
	} else {
		signedTx, err := types.SignTx(tx, types.NewLondonSignerWithEIP4844(r.chainId), r.privKey)
		if err != nil {
			return nil, fmt.Errorf("sign tx error:%v", err)
		}
		return signedTx, nil

	}
}
