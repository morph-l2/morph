package externalsign

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestRequestSign(t *testing.T) {

	appid := "morph-setup-0D799FE0-401D-4A7C-8C35-32E38F85F37D"
	rsaPrivStr := ""
	signUrl := "http://localhost:8080/v1/sign/tx_sign"
	addr := "0x83da35ef5eae423f43023edb20b3a43de443cbca"
	chain := "QANET-L2"
	chainid := big.NewInt(53077)
	signer := types.LatestSignerForChainID(chainid)

	rsa, err := ParseRsaPrivateKey(rsaPrivStr)
	require.NoError(t, err)
	es := NewExternalSign(appid, rsa, signUrl, addr, chain, signer)

	// testdata
	topk, err := crypto.GenerateKey()
	require.NoError(t, err)
	toaddr := crypto.PubkeyToAddress(topk.PublicKey)
	gas := uint64(50000)

	txdata := &types.DynamicFeeTx{
		To:        &toaddr,
		Gas:       gas,
		GasFeeCap: big.NewInt(1),
		GasTipCap: big.NewInt(2),
		Value:     big.NewInt(3),
		ChainID:   chainid,
	}

	tx := types.NewTx(txdata)
	hashHex := signer.Hash(tx).Hex()
	data, err := es.newData(hashHex)
	require.NoError(t, err)
	reqData, err := es.craftReqData(*data)
	require.NoError(t, err)
	pubstr, err := GetPubKeyStr(rsa)
	require.NoError(t, err)
	reqData.Pubkey = pubstr
	require.NoError(t, err)
	t.Log("reqData", reqData)
	signedTx, err := es.requestSign(*reqData, tx)
	require.NoError(t, err)

	from, err := signer.Sender(signedTx)
	require.NoError(t, err)
	require.Equal(t, hexutil.Encode(from.Bytes()), addr)
	require.Equal(t, txdata.Gas, signedTx.Gas())
	require.Equal(t, txdata.GasFeeCap, signedTx.GasFeeCap())
	require.Equal(t, txdata.GasTipCap, signedTx.GasTipCap())
	require.Equal(t, txdata.Value, signedTx.Value())
	require.Equal(t, txdata.Data, signedTx.Data())

}

func TestNewWallet(t *testing.T) {

	//test data
	appid := "morph-setup-0D799FE0-401D-4A7C-8C35-32E38F85F37D"
	rsaPrivStr := ""
	signUrl := "http://localhost:8080/v1/sign/gen_address"
	addr := "0x83da35ef5eae423f43023edb20b3a43de443cbca"
	chain := "QANET-L2"
	chainid := big.NewInt(53077)
	signer := types.LatestSignerForChainID(chainid)

	rsaPriv, err := ParseRsaPrivateKey(rsaPrivStr)
	require.NoError(t, err)
	es := NewExternalSign(appid, rsaPriv, signUrl, addr, chain, signer)
	data, err := es.newData("")
	require.NoError(t, err)
	reqData, err := es.craftReqData(*data)
	require.NoError(t, err)
	pubstr, err := GetPubKeyStr(rsaPriv)
	require.NoError(t, err)
	reqData.Pubkey = pubstr
	require.NoError(t, err)
	t.Log("reqData", reqData)
	es.requestSign(*reqData, nil)

}
