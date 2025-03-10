package flags

import "github.com/urfave/cli"

var (
	TxHashFlag = cli.StringFlag{
		Name:     "tx_hash",
		Usage:    "The hash of the transaction to cancel",
		Required: true,
		EnvVar:   prefixEnvVar("TX_HASH"),
	}
)

var CancleTxFlags = []cli.Flag{
	TxHashFlag,
	L1EthRpcFlag,
	PrivateKeyFlag,

	// external sign
	ExternalSign,
	ExternalSignAddress,
	ExternalSignAppid,
	ExternalSignChain,
	ExternalSignUrl,
	ExternalSignRsaPriv,
}
