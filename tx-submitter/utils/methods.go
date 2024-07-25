package utils

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto"
)

func ParsePkAndWallet(pkStr, walletStr string) (*ecdsa.PrivateKey, *common.Address, error) {
	hex := strings.TrimPrefix(pkStr, "0x")
	priv, err := crypto.HexToECDSA(hex)
	if err != nil {
		return nil, nil, err
	}
	wallet := common.HexToAddress(walletStr)

	return priv, &wallet, nil
}
func ParseRsaPrivateKey(pkStr string) (*rsa.PrivateKey, error) {
	// decode private
	var derBytes []byte
	blockPub, rest := pem.Decode([]byte(pkStr))
	switch string(rest) {
	case pkStr:
		// private key only
		publicBytes, err := base64.StdEncoding.DecodeString(pkStr)
		if err != nil {
			return nil, err
		}
		derBytes = publicBytes
	default:
		// contains the private key that begins with -----BEGIN
		derBytes = blockPub.Bytes
	}
	priKey, err := x509.ParsePKCS1PrivateKey(derBytes)
	if err != nil {
		privateKeyF, err := x509.ParsePKCS8PrivateKey(derBytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		priKey, ok = privateKeyF.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("%s", "failed to parse to rsa private key")
		}
	}

	return priKey, nil
}

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
