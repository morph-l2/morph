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
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"
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
	Signer  types.Signer
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
	Address string `json:"address"`
	Chain   string `json:"chain"`
	Sha3    string `json:"sha3"`
}
type GenAddrData struct {
	CoinId      string `json:"coinId"`
	ChainCoinId string `json:"chainCoinId"`
	EncryptKey  string `json:"encryptKey"`
	KeyMd5      string `json:"keyMd5"`
	UniqId      string `json:"uniqId"`
	Chain       string `json:"chain"`
}

func init() {
	output := io.Writer(os.Stdout)
	logHandler := log.StreamHandler(output, log.TerminalFormat(false))
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, logHandler))
}

func NewExternalSign(appid string, priv *rsa.PrivateKey, signUrl, addr, chain string, signer types.Signer) *ExternalSign {

	// new resty.client
	client := resty.New()
	return &ExternalSign{
		Appid:   appid,
		Privkey: priv,
		Client:  client,
		signUrl: signUrl,
		Address: addr,
		Chain:   chain,
		Signer:  signer,
	}
}

func (e *ExternalSign) newData(hash string) (*Data, error) {

	return &Data{
		Address: e.Address,
		Chain:   e.Chain,
		Sha3:    hash,
	}, nil
}

func (e *ExternalSign) NewGenAddrData() (*GenAddrData, error) {
	return &GenAddrData{
		Chain: e.Chain,
	}, nil
}

func (e *ExternalSign) craftReqData(data interface{}) (*ReqData, error) {
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

func (e *ExternalSign) RequestSign(tx *types.Transaction) (*types.Transaction, error) {
	hashHex := e.Signer.Hash(tx).Hex()

	data, err := e.newData(hashHex)
	if err != nil {
		return nil, fmt.Errorf("new data error:%s", err)
	}
	reqdata, err := e.craftReqData(*data)
	if err != nil {
		return nil, fmt.Errorf("craft req data error:%s", err)
	}
	signedTx, err := e.requestSign(*reqdata, tx)
	if err != nil {
		return nil, fmt.Errorf("request sign error:%s", err)
	}
	return signedTx, nil
}

func (e *ExternalSign) requestSign(data ReqData, tx *types.Transaction) (*types.Transaction, error) {

	resp, err := e.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(e.signUrl)

	if err != nil {
		return nil, fmt.Errorf("request sign error: %v", err)
	}

	// log resp info
	log.Info("request sign response",
		"status", resp.StatusCode(),
		"body", resp.String(),
		// "result", result,
	)

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response status not ok: %v, resp body:%v", resp.StatusCode(), string(resp.Body()))
	}

	// decode resp
	response := new(Response)
	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal resp err:%w", err)
	}

	if len(response.Result.SignDatas) == 0 {
		return nil, errors.New("respones sha3 empty")
	}

	sig, err := hexutil.Decode(response.Result.SignDatas[0].Sign)
	if err != nil {
		return nil, fmt.Errorf("decide sig failed:%w", err)
	}
	signedTx, err := tx.WithSignature(e.Signer, sig)
	if err != nil {
		return nil, fmt.Errorf("with signature err:%w", err)
	}
	return signedTx, nil
}
