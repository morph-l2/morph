package client

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	remotesigner "github.com/morph-l2/remote-signer-client/go/signer"
)

// Signer talks to the remote signer. The request shape was taken from
// token-price-oracle as a reference; this package does not import that service.
type Signer struct {
	remoteClient        *remotesigner.Client
	externalSignAddress common.Address
	chainID             *big.Int
}

// NewSigner creates a remote signer client.
func NewSigner(
	appid string,
	rsaPriv *rsa.PrivateKey,
	address string,
	chain string,
	url string,
	chainID *big.Int,
) (*Signer, error) {
	remoteClient, err := remotesigner.NewClient(appid, rsaPriv, address, chain, url, types.NewLondonSigner(chainID))
	if err != nil {
		return nil, fmt.Errorf("failed to create remote signer client: %w", err)
	}
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid external signer address: %s", address)
	}
	log.Info("External signer initialized", "address", address, "chain", chain)
	return &Signer{
		remoteClient:        remoteClient,
		externalSignAddress: common.HexToAddress(address),
		chainID:             chainID,
	}, nil
}

// CreateAndSignTx builds an EIP-1559 transaction and asks the remote signer to sign it.
// value is the ETH attached to challengeState and must be at least the challenge deposit.
func (s *Signer) CreateAndSignTx(
	ctx context.Context,
	client *ethclient.Client,
	to common.Address,
	value *big.Int,
	callData []byte,
	methodSig string,
	maxTipCap *big.Int,
	maxFeeCap *big.Int,
) (*types.Transaction, error) {
	from := s.externalSignAddress
	if value == nil {
		value = big.NewInt(0)
	}
	value = new(big.Int).Set(value)

	nonce, err := client.NonceAt(ctx, from, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	caps, err := CalculateGasCaps(ctx, client, maxTipCap, maxFeeCap)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate gas caps: %w", err)
	}

	gas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:      from,
		To:        &to,
		Value:     value,
		GasFeeCap: caps.FeeCap,
		GasTipCap: caps.TipCap,
		Data:      callData,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}
	gas = gas * 3 / 2

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.chainID,
		Nonce:     nonce,
		GasTipCap: caps.TipCap,
		GasFeeCap: caps.FeeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
		Data:      callData,
	})

	log.Info("Created transaction for signing",
		"from", from.Hex(),
		"to", to.Hex(),
		"nonce", nonce,
		"value", value,
		"gas", gas,
		"gasFeeCap", caps.FeeCap,
		"gasTipCap", caps.TipCap,
		"methodSig", methodSig)

	return s.remoteClient.Sign(tx, methodSig)
}
