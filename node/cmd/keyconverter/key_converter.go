package keyconverter

import (
	"encoding/base64"
	"fmt"

	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/blssignatures"
	"github.com/urfave/cli"
)

func ConvertKey(ctx *cli.Context) error {
	args := ctx.Args()
	if len(args) != 1 {
		return fmt.Errorf("wrong args length, expected: %d, actual: %d", 1, len(args))
	}
	base64EncodedKey := args[0]
	raw, err := base64.StdEncoding.DecodeString(base64EncodedKey)
	if err != nil {
		return fmt.Errorf("failed to decode base64 keys, err: %v", err)
	}
	keyType := ctx.String("type")
	switch keyType {
	case "tm":
		if len(raw) != 32 {
			return fmt.Errorf("wrong public key length, expected: 32, actual: %d", len(raw))
		}
		fmt.Printf("converted tendermint public key: %v \n", hexutil.Encode(raw))
	case "bls":
		blsPublicKey, err := blssignatures.PublicKeyFromBytes(raw, true)
		if err != nil {
			return fmt.Errorf("failed to decode base64 keys, err: %v", err)
		}
		encodedPoint := bls12381.NewG2().EncodePoint(blsPublicKey.Key)
		fmt.Printf("converted bls encoded key: %v \n", hexutil.Encode(encodedPoint))

	default:
		return fmt.Errorf("unknown key type provided. expected: tm/bls")
	}
	return nil
}

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "type",
		Usage: "the key type being converted: tm/bls. tm: tendermint public key for the consensus; bls: bls public key for bls signature",
		Value: "tm",
	},
}
