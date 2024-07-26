package externalsign

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func Test_RequestSign(t *testing.T) {

	appid := ""
	rsaPrivStr := ""
	signUrl := ""

	rsa, err := ParseRsaPrivateKey(rsaPrivStr)
	require.NoError(t, err)
	es := NewExternalSign(appid, rsa, signUrl)

	// testdata
	topk, err := crypto.GenerateKey()
	require.NoError(t, err)
	toaddr := crypto.PubkeyToAddress(topk.PublicKey)
	gas := uint64(50000)
	chainid := big.NewInt(4)

	txdata := &types.DynamicFeeTx{
		To:        &toaddr,
		Gas:       gas,
		GasFeeCap: big.NewInt(1),
		GasTipCap: big.NewInt(2),
		Value:     big.NewInt(3),
		ChainID:   chainid,
	}

	data, err := es.newData([]types.TxData{txdata})
	data.Chain = "ETH"
	// todo: fill it
	data.Address = ""
	require.NoError(t, err)
	reqData, err := es.craftReqData(*data)
	require.NoError(t, err)
	pubstr, err := GetPubKeyStr(rsa)
	require.NoError(t, err)
	reqData.Pubkey = pubstr
	require.NoError(t, err)
	t.Log("reqData", reqData)
	signedTx, err := es.requestSign(*reqData)
	require.NoError(t, err)

	// require.Equal(t, txdata.Hash(), signedTx.Hash())
	require.Equal(t, txdata.Gas, signedTx.Gas())
	require.Equal(t, txdata.GasFeeCap, signedTx.GasFeeCap())
	require.Equal(t, txdata.GasTipCap, signedTx.GasTipCap())
	require.Equal(t, txdata.Value, signedTx.Value())
	require.Equal(t, txdata.Data, signedTx.Data())

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
