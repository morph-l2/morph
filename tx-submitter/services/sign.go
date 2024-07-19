package services

import (
	"fmt"

	"morphl2/externalsign"

	"github.com/scroll-tech/go-ethereum/core/types"
)

var externalSigner *externalsign.ExternalSign

func (r *Rollup) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if r.cfg.ExternalSign {
		if externalSigner == nil {
			externalsign.NewExternalSign(r.cfg.ExternalSignAppid, r.privKey, r.cfg.ExternalSignUrl)
		}

		signedTx, err := externalsigner.RequestSign([]types.Transaction{*tx})
		if err != nil {
			return nil, fmt.Errorf("externalsign sign tx error:%v", err)
		}
		return signedTx, nil
	} else {
		signedTx, err := types.SignTx(tx, types.NewLondonSignerWithEIP4844(r.chainId), r.privKey)
		if err != nil {
			return nil, fmt.Errorf("sign tx error:%v", err)
		}
		return signedTx, nil

	}
}
