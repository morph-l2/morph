package externalsign

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

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

func GetPubKeyStr(priv *rsa.PrivateKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key:%w", err)
	}
	return base64.StdEncoding.EncodeToString(pubKeyBytes), nil

}
