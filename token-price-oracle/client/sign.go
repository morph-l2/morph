package client

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
	remotesigner "github.com/morph-l2/remote-signer-client/go/signer"
)

// Signer handles transaction signing with support for both local and external signing
type Signer struct {
	remoteClient        *remotesigner.Client
	externalSignAddress common.Address
	chainID             *big.Int
}

// NewSigner creates a new Signer instance
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

	log.Info("External signer initialized",
		"address", address,
		"chain", chain)

	return &Signer{
		remoteClient:        remoteClient,
		externalSignAddress: common.HexToAddress(address),
		chainID:             chainID,
	}, nil
}

// GetFromAddress returns the signer's address
func (s *Signer) GetFromAddress() common.Address {
	return s.externalSignAddress
}

// CreateAndSignTx creates a new transaction and signs it
func (s *Signer) CreateAndSignTx(
	ctx context.Context,
	client *L2Client,
	to common.Address,
	callData []byte,
	methodSig string,
) (*types.Transaction, error) {
	from := s.externalSignAddress

	nonce, err := client.GetClient().NonceAt(ctx, from, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Calculate gas caps (dynamic values with optional max limits)
	caps, err := CalculateGasCaps(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate gas caps: %w", err)
	}

	// Estimate gas
	gas, err := client.GetClient().EstimateGas(ctx, ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasFeeCap: caps.FeeCap,
		GasTipCap: caps.TipCap,
		Data:      callData,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Add 50% buffer to gas estimate
	gas = gas * 3 / 2

	// Create transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.chainID,
		Nonce:     nonce,
		GasTipCap: caps.TipCap,
		GasFeeCap: caps.FeeCap,
		Gas:       gas,
		To:        &to,
		Data:      callData,
	})

	log.Info("Created transaction for signing",
		"from", from.Hex(),
		"to", to.Hex(),
		"nonce", nonce,
		"gas", gas,
		"gasFeeCap", caps.FeeCap,
		"gasTipCap", caps.TipCap,
		"methodSig", methodSig)

	// Sign transaction
	return s.remoteClient.Sign(tx, methodSig)
}
