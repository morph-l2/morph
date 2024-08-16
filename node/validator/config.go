package validator

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/urfave/cli"

	"morph-l2/node/flags"
)

type Config struct {
	l1RPC           string
	PrivateKey      *ecdsa.PrivateKey
	L1ChainID       *big.Int
	rollupContract  common.Address
	challengeEnable bool
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	l1NodeAddr := ctx.GlobalString(flags.L1NodeAddr.Name)
	l1ChainID := ctx.GlobalUint64(flags.L1ChainID.Name)
	hexPrvKey := ctx.GlobalString(flags.ValidatorPrivateKey.Name)
	hex := strings.TrimPrefix(hexPrvKey, "0x")
	privateKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		return err
	}
	c.challengeEnable = ctx.GlobalIsSet(flags.ValidatorEnable.Name)
	addrHex := ctx.GlobalString(flags.RollupContractAddress.Name)
	rollupContract := common.HexToAddress(addrHex)
	c.l1RPC = l1NodeAddr
	c.L1ChainID = big.NewInt(int64(l1ChainID))
	c.PrivateKey = privateKey
	c.rollupContract = rollupContract
	return nil
}
