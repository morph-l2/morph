package client

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/big"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
)

// Signer handles transaction signing with support for both local and external signing
type Signer struct {
	externalSign        bool
	externalSigner      *externalsign.ExternalSign
	externalSignUrl     string
	externalSignAddress common.Address
	chainID             *big.Int
	signer              types.Signer
}

// NewSigner creates a new Signer instance
func NewSigner(
	externalSign bool,
	externalSignAppid string,
	externalRsaPriv *rsa.PrivateKey,
	externalSignAddress string,
	externalSignChain string,
	externalSignUrl string,
	chainID *big.Int,
) *Signer {
	signer := types.NewLondonSigner(chainID)

	s := &Signer{
		externalSign:        externalSign,
		externalSignUrl:     externalSignUrl,
		externalSignAddress: common.HexToAddress(externalSignAddress),
		chainID:             chainID,
		signer:              signer,
	}

	if externalSign {
		s.externalSigner = externalsign.NewExternalSign(
			externalSignAppid,
			externalRsaPriv,
			externalSignAddress,
			externalSignChain,
			signer,
		)
		log.Info("External signer initialized",
			"address", externalSignAddress,
			"chain", externalSignChain)
	}

	return s
}

// Sign signs a transaction using either external or local signing
func (s *Signer) Sign(tx *types.Transaction) (*types.Transaction, error) {
	if !s.externalSign {
		return nil, fmt.Errorf("local signing not supported in Signer, use bind.TransactOpts")
	}

	signedTx, err := s.externalSigner.RequestSign(s.externalSignUrl, tx)
	if err != nil {
		return nil, fmt.Errorf("external sign request failed: %w", err)
	}
	return signedTx, nil
}

// IsExternalSign returns whether external signing is enabled
func (s *Signer) IsExternalSign() bool {
	return s.externalSign
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
) (*types.Transaction, error) {
	from := s.externalSignAddress

	nonce, err := client.GetClient().NonceAt(ctx, from, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Calculate gas caps (dynamic values with optional max limits)
	caps, err := CalculateGasCapsAlways(ctx, client)
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
		"gasTipCap", caps.TipCap)

	// Sign transaction
	return s.Sign(tx)
}

