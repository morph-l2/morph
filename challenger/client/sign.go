package client

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	remotesigner "github.com/morph-l2/remote-signer-client/go/signer"
)

// probeTimeout bounds one remote signer health check.
const probeTimeout = 20 * time.Second

// probeGasLimit is the gas limit of the probe transaction. It cannot be
// estimated, because the placeholder batch makes challengeState revert.
const probeGasLimit = 500_000

// remoteSigner is the part of the remote signer client used here.
type remoteSigner interface {
	SignV1(tx *types.Transaction) (*types.Transaction, error)
}

// Signer talks to the remote signer. The request shape was taken from
// token-price-oracle as a reference; this package does not import that service.
type Signer struct {
	remoteClient        remoteSigner
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
	return s.createAndSignTx(ctx, client, to, value, callData, methodSig, maxTipCap, maxFeeCap, 0)
}

// createAndSignTx estimates the gas limit when gasLimit is zero.
func (s *Signer) createAndSignTx(
	ctx context.Context,
	client *ethclient.Client,
	to common.Address,
	value *big.Int,
	callData []byte,
	methodSig string,
	maxTipCap *big.Int,
	maxFeeCap *big.Int,
	gasLimit uint64,
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

	gas := gasLimit
	if gas == 0 {
		gas, err = client.EstimateGas(ctx, ethereum.CallMsg{
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
	}

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

	// L1 signing uses the v1 API, which sends the transaction hash, as
	// tx-submitter and the prover challenge handler do. The signer rejects a
	// v2 request for this transaction because it carries no sha3.
	signed, err := s.remoteClient.SignV1(tx)
	if err != nil {
		return nil, err
	}
	if err := s.verifySender(signed); err != nil {
		return nil, err
	}
	return signed, nil
}

// verifySender checks that the remote signer signed with the configured address.
// The signature is applied to the transaction built here, so only the
// recovered sender can differ from what was requested.
func (s *Signer) verifySender(signed *types.Transaction) error {
	sender, err := types.Sender(types.NewLondonSigner(s.chainID), signed)
	if err != nil {
		return fmt.Errorf("recover remote signature sender: %w", err)
	}
	if sender != s.externalSignAddress {
		return fmt.Errorf("remote signer signed as %s, want %s", sender.Hex(), s.externalSignAddress.Hex())
	}
	return nil
}

// Probe builds a challengeState transaction the same way a real challenge
// does, asks the remote signer to sign it, and checks the recovered sender.
// Only the gas limit is fixed instead of estimated. The signed transaction is
// discarded here and is never broadcast.
func (s *Signer) Probe(
	ctx context.Context,
	client *ethclient.Client,
	to common.Address,
	value *big.Int,
	callData []byte,
	methodSig string,
	maxTipCap *big.Int,
	maxFeeCap *big.Int,
) error {
	ctx, cancel := context.WithTimeout(ctx, probeTimeout)
	defer cancel()

	// The remote signer client takes no context, so the wait is bounded here.
	done := make(chan error, 1)
	go func() {
		_, err := s.createAndSignTx(ctx, client, to, value, callData, methodSig, maxTipCap, maxFeeCap, probeGasLimit)
		done <- err
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("remote signer probe: %w", ctx.Err())
	case err := <-done:
		return err
	}
}
