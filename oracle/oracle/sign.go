package oracle

import (
	"fmt"
	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum/core/types"
)

var externalSigner *externalsign.ExternalSign

func (o *Oracle) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if o.cfg.ExternalSign {
		if externalSigner == nil {
			externalSigner = externalsign.NewExternalSign(o.cfg.ExternalSignAppid, o.externalRsaPriv, o.cfg.ExternalSignAddress, o.cfg.ExternalSignChain, o.signer)
		}
		signedTx, err := externalSigner.RequestSign(o.cfg.ExternalSignUrl, tx)
		if err != nil {
			return nil, fmt.Errorf("externalsign sign tx error:%v", err)
		}
		return signedTx, nil
	} else {
		signedTx, err := types.SignTx(tx, o.signer, o.privKey)
		if err != nil {
			return nil, fmt.Errorf("sign tx error:%v", err)
		}
		return signedTx, nil

	}
}
