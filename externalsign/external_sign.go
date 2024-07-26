package externalsign

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
	Data      string `json:"data"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
}

type ReqData struct {
	BusinessData
	BizSignature string `json:"bizSignature"`
	Pubkey       string `json:"publicKey"`
}

type Data struct {
	Address   string         `json:"address"`
	Chain     string         `json:"chain"`
	SignDatas []types.TxData `json:"signDatas"` // raw txs jsons
}

func init() {
	output := io.Writer(os.Stdout)
	logHandler := log.StreamHandler(output, log.TerminalFormat(false))
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, logHandler))
}

func NewExternalSign(appid string, priv *rsa.PrivateKey, signUrl, addr, chain string) *ExternalSign {

	// new resty.client
	client := resty.New()
	return &ExternalSign{
		Appid:   appid,
		Privkey: priv,
		Client:  client,
		signUrl: signUrl,
		Address: addr,
		Chain:   chain,
	}
}

func (e *ExternalSign) newData(txDatas []types.TxData) (*Data, error) {

	return &Data{
		Address:   e.Address,
		Chain:     e.Chain,
		SignDatas: txDatas,
	}, nil
}

func (e *ExternalSign) craftReqData(data Data) (*ReqData, error) {
	nonceStr := uuid.NewString()
	dataBs, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal data failed: %w", err)
	}
	businessData := BusinessData{
		Appid:     e.Appid,
		Data:      string(dataBs),
		Noncestr:  nonceStr,
		Timestamp: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
	businessDataBs, err := json.Marshal(businessData)
	if err != nil {
		return nil, fmt.Errorf("marshal data failed: %w", err)
	}
	hashed := sha256.Sum256([]byte(businessDataBs))
	signature, err := rsa.SignPKCS1v15(nil, e.Privkey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("sign data failed: %w", err)
	}
	hexSig := hex.EncodeToString(signature)

	return &ReqData{
		BusinessData: businessData,
		BizSignature: hexSig,
	}, nil

}

func (e *ExternalSign) RequestSign(txdatas []types.TxData) (*types.Transaction, error) {
	data, err := e.newData(txdatas)
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
	log.Info("request sign response",
		"status", resp.StatusCode(),
		"body", resp.String(),
		"result", resp.Result(),
	)

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response status not ok: %v, resp body:%v", resp.StatusCode(), string(resp.Body()))
	}

	if len(response.Result.SignDatas) == 0 {
		return nil, errors.New("signDatas empty")
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes([]byte(response.Result.SignDatas[0].Sign), tx)
	if err != nil {
		return nil, fmt.Errorf("decode signed tx err: %v", err)
	}
	return tx, nil
}
