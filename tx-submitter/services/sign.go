package services

import (
	"encoding/hex"
	"fmt"

	"github.com/morph-l2/externalsign"

	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
)

var externalSigner *externalsign.ExternalSign

func (r *Rollup) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if r.cfg.ExternalSign {
		if externalSigner == nil {
			externalSigner = externalsign.NewExternalSign(r.cfg.ExternalSignAppid, r.externalRsaPriv, r.cfg.ExternalSignAddress, r.cfg.ExternalSignChain, r.signer)
		}
		logExternalSignRequestSize(tx)
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

// logExternalSignRequestSize reports the unsigned-tx encoding size that dominates
// the JSON body sent to the external signer (txData is hex of MarshalBinary).
func logExternalSignRequestSize(tx *types.Transaction) {
	txBs, err := tx.MarshalBinary()
	if err != nil {
		log.Warn("failed to marshal tx for external sign size log", "err", err)
		return
	}

	marshalBytes := len(txBs)
	txDataHexBytes := hex.EncodedLen(marshalBytes)
	ctx := []interface{}{
		"tx_type", tx.Type(),
		"marshal_bytes", marshalBytes,
		"tx_data_hex_bytes", txDataHexBytes,
	}
	if sc := tx.BlobTxSidecar(); sc != nil {
		ctx = append(ctx,
			"blobs", len(sc.Blobs),
			"sidecar_version", sc.Version,
		)
	}
	log.Info("external sign request size", ctx...)
}
