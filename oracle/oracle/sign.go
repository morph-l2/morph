package oracle

import (
	"fmt"
	"math/big"

	"morph-l2/bindings/predeploys"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
)

var externalSigner *externalsign.ExternalSign

func (o *Oracle) sign(tx *types.Transaction) (*types.Transaction, error) {
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

func (o *Oracle) newRecordTxAndSign(callData []byte) (*types.Transaction, error) {
	from := common.HexToAddress(o.cfg.ExternalSignAddress)
	if !o.cfg.ExternalSign {
		from = crypto.PubkeyToAddress(o.privKey.PublicKey)
	}
	nonce, err := o.l2Client.NonceAt(o.ctx, from, nil)
	if err != nil {
		return nil, err
	}
	// tip and cap
	tip, err := o.l2Client.SuggestGasTipCap(o.ctx)
	if err != nil {
		return nil, err
	}
	head, err := o.l2Client.HeaderByNumber(o.ctx, nil)
	if err != nil {
		return nil, err
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
	gas, err := o.l2Client.EstimateGas(o.ctx, ethereum.CallMsg{
		From:      from,
		To:        &predeploys.RecordAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      callData,
	})
	if err != nil {
		return nil, err
	}
	return o.sign(types.NewTx(&types.DynamicFeeTx{
		ChainID:   o.chainId,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &o.recordAddr,
		Data:      callData}))
}
