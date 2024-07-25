package externalsign

import (
	"crypto/rsa"
	"math/big"
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
	require.NoError(t, err)
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
	signedTx, err := es.RequestSign([]*types.Transaction{tx})
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

	//test data
	appid := ""
	rsaPrivStr := ""
	signUrl := ""

	rsaPriv, err := ParseRsaPrivateKey(rsaPrivStr)
	require.NoError(t, err)
	es := NewExternalSign(appid, rsaPriv, signUrl)
	data, err := es.newData(nil)
	data.Chain = "ETH"
	require.NoError(t, err)
	reqData, err := es.craftReqData(*data)
	require.NoError(t, err)
	pubstr, err := GetPubKeyStr(rsaPriv)
	require.NoError(t, err)
	reqData.Pubkey = pubstr
	require.NoError(t, err)
	t.Log("reqData", reqData)
	es.requestSign(*reqData)

}
