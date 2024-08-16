package services

import (
	"fmt"

	"github.com/morph-l2/externalsign"

	"github.com/morph-l2/go-ethereum/core/types"
)

var externalSigner *externalsign.ExternalSign

func (r *Rollup) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if r.cfg.ExternalSign {
		if externalSigner == nil {
			externalSigner = externalsign.NewExternalSign(r.cfg.ExternalSignAppid, r.externalRsaPriv, r.cfg.ExternalSignAddress, r.cfg.ExternalSignChain, r.signer)
		}
		signedTx, err := externalSigner.RequestSign(r.cfg.ExternalSignUrl, tx)
		if err != nil {
			return nil, fmt.Errorf("externalsign sign tx error:%v", err)
		}
		return signedTx, nil
	} else {
		signedTx, err := types.SignTx(tx, r.signer, r.privKey)
		if err != nil {
			return nil, fmt.Errorf("sign tx error:%v", err)
		}
		return signedTx, nil

	}
}
