package externalsign

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rlp"
)

type ExternalSign struct {
	Appid   string
	Address string
	Privkey *rsa.PrivateKey
	Chain   string

	// http client
	Client *resty.Client
	// url
	signUrl string
}

type BusinessData struct {
	Appid     string `json:"appid"`
	Data      Data   `json:"data"`
	Noncestr  string `json:"noncestr"`
	Timestamp uint64 `json:"timestamp"`
}

type ReqData struct {
	BusinessData
	BizSignature string `json:"bizSignature"`
}

type Data struct {
	Address   string   `json:"address"`
	Chain     string   `json:"chain"`
	SignDatas []string `json:"signDatas"` // raw txs jsons
}

func NewExternalSign(appid string, priv *rsa.PrivateKey, signUrl string) *ExternalSign {

	// new resty.client
	client := resty.New()
	return &ExternalSign{
		Appid:   appid,
		Privkey: priv,
		Client:  client,
		signUrl: signUrl,
	}
}

func (e *ExternalSign) newData(txs []types.Transaction) (*Data, error) {

	// marshal tx
	var signDatas []string
	for _, tx := range txs {
		signData, err := tx.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("create new data failed: %w", err)
		}
		signDatas = append(signDatas, string(signData))
	}

	return &Data{
		Address:   e.Address,
		Chain:     e.Chain,
		SignDatas: signDatas,
	}, nil
}

func (e *ExternalSign) craftReqData(data Data) (*ReqData, error) {
	nonceStr := uuid.NewString()
	databs, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal data failed: %w", err)
	}
	hashed := sha256.Sum256([]byte(databs))
	signature, err := rsa.SignPKCS1v15(nil, e.Privkey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("sign data failed: %w", err)
	}
	hexSig := hex.EncodeToString(signature)

	return &ReqData{
		BusinessData: BusinessData{
			Appid:     e.Appid,
			Data:      data,
			Noncestr:  nonceStr,
			Timestamp: uint64(time.Now().Unix()),
		},
		BizSignature: hexSig,
	}, nil

}

func (e *ExternalSign) RequestSign(txs []types.Transaction) (*types.Transaction, error) {
	data, err := e.newData(txs)
	if err != nil {
		return nil, fmt.Errorf("new data error:%s", err)
	}
	reqdata, err := e.craftReqData(*data)
	if err != nil {
		return nil, fmt.Errorf("craft req data error:%s", err)
	}
	signedTx, err := e.requestSign(*reqdata)
	if err != nil {
		return nil, fmt.Errorf("request sign error:%s", err)
	}
	return signedTx, nil
}

func (e *ExternalSign) requestSign(data ReqData) (*types.Transaction, error) {

	response := new(Response)
	resp, err := e.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		SetResult(response). // or SetResult(AuthSuccess{}).
		Post(e.signUrl)

	if err != nil {
		return nil, fmt.Errorf("request sign error: %v", err)
	}

	// log resp info
	log.Debug("request sign response",
		"status", resp.StatusCode(),
		"body", resp.String(),
	)

	tx := new(types.Transaction)
	err = rlp.DecodeBytes([]byte(response.Result.SignDatas[0].Sign), tx)
	if err != nil {
		return nil, fmt.Errorf("decode signed tx err: %v", err)
	}

	// resp status check
	// if resp.StatusCode() != http.StatusOK {
	// 	return nil, fmt.Errorf("request sign error,code not ok: %v", resp.StatusCode())
	// }

	// respon body check
	// if resp.RawResponse == nil {
	// 	return nil, errors.New("request sign no response")
	// }

	// parse response
	// defer resp.RawResponse.Body.Close()
	// if err := json.NewDecoder(resp.RawResponse.Body).Decode(&response); err != nil {
	// 	return nil, fmt.Errorf("request sign error: %v", err)
	// }

	return tx, nil
}
