package services

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/morph-l2/externalsign"
	"github.com/stretchr/testify/require"
)

func TestExternalAddrGenerator(t *testing.T) {
	const want = "0x33D5b507868b7E8AC930CD3BDe9eaDd60c638479"
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, genAddressPath, r.URL.Path)
		var req externalsign.ReqData
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.Equal(t, "app", req.Appid)
		require.JSONEq(t, `{"coinId":"","chainCoinId":"","encryptKey":"","keyMd5":"","uniqId":"","chain":"QANET-L1"}`, req.Data)
		_, _ = w.Write([]byte(`{"status":0,"result":{"address":"` + want + `"}}`))
	}))
	defer proxy.Close()

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	gen, err := NewExternalAddrGenerator(ExternalSignConfig{
		Appid: "app", RsaPriv: priv, Chain: "QANET-L1", Url: proxy.URL + "/v1/sign/tx_sign",
	})
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	gen.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/external_sign/gen_address", nil))
	require.Equal(t, http.StatusOK, rec.Code)
	var resp map[string]string
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	require.Equal(t, want, resp["address"])

	rec = httptest.NewRecorder()
	gen.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/external_sign/gen_address", nil))
	require.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}
