package utils

import (
	"crypto/ecdsa"
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
