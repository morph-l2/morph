package externalsign

import (
	"crypto/rsa"
	"math/big"
	"morph-l2/tx-submitter/utils"
	"testing"

	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func Test_RequestSign(t *testing.T) {
	appid := ""
	rsa, err := rsa.GenerateKey(nil, 2048)
	if err != nil {
		t.Error(err)
	}
	es := NewExternalSign(appid, rsa, "")

	// testdata
	topk, err := crypto.GenerateKey()
	toaddr := crypto.PubkeyToAddress(topk.PublicKey)
	gas := uint64(50000)
	chainid := big.NewInt(4)

	tx := types.NewTx(
		&types.DynamicFeeTx{
			To:        &toaddr,
			Gas:       gas,
			GasFeeCap: big.NewInt(1),
			GasTipCap: big.NewInt(2),
			Value:     big.NewInt(3),
			ChainID:   chainid,
		},
	)
	signedTx, err := es.RequestSign([]types.Transaction{*tx})
	require.NoError(t, err)

	require.Equal(t, tx.Hash(), signedTx.Hash())
	require.Equal(t, tx.Gas(), signedTx.Gas())
	require.Equal(t, tx.GasFeeCap(), signedTx.GasFeeCap())
	require.Equal(t, tx.GasTipCap(), signedTx.GasTipCap())
	require.Equal(t, tx.Value(), signedTx.Value())
	require.Equal(t, tx.Value(), signedTx.Value())
	require.Equal(t, tx.Data(), signedTx.Data())

}

func TestNewWallet(t *testing.T) {

	// business.Raw.AppId = "morph-setup-0D799FE0-401D-4A7C-8C35-32E38F85F37D"
	// business.Raw.Data = "{\"chain\":\"ETH\"}"
	// business.Raw.Noncestr = "93414d51-4f64-4595-a37d-70ff7c5387e6"
	// business.Raw.Timestamp = fmt.Sprintf("%d", time.Now().UnixMilli())

	//test data
	appid := "morph-tx-submitter-399A1722-3F2C-4E39-ABD2-1B65D02C66BA"
	rsaPrivStr := "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"
	signUrl := "http://localhost:8080/v1/sign/gen_address"

	rsaPriv, err := utils.ParseRsaPrivateKey(rsaPrivStr)
	require.NoError(t, err)
	es := NewExternalSign(appid, rsaPriv, signUrl)
	data, err := es.newData(nil)
	data.Chain = "ETH"
	require.NoError(t, err)
	reqData, err := es.craftReqData(*data)
	require.NoError(t, err)
	t.Log("reqData", reqData)
	es.requestSign(*reqData)

}
