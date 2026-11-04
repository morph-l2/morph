package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
)

// GasCaps holds the EIP-1559 tip and fee caps applied to a challenge transaction.
type GasCaps struct {
	TipCap *big.Int
	FeeCap *big.Int
}

type headerClient interface {
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// CalculateGasCaps fetches the suggested tip and base fee, then applies optional ceilings.
// The result keeps tipCap <= feeCap.
func CalculateGasCaps(ctx context.Context, client *ethclient.Client, maxTipCap, maxFeeCap *big.Int) (*GasCaps, error) {
	return calculateGasCaps(ctx, client, maxTipCap, maxFeeCap)
}

func calculateGasCaps(ctx context.Context, client headerClient, maxTipCap, maxFeeCap *big.Int) (*GasCaps, error) {
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested gas tip cap: %w", err)
	}
	if maxTipCap != nil && tip.Cmp(maxTipCap) > 0 {
		log.Debug("Applying gas tip cap limit", "dynamic", tip, "max", maxTipCap)
		tip = new(big.Int).Set(maxTipCap)
	}

	head, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %w", err)
	}

	var feeCap *big.Int
	if head.BaseFee != nil {
		feeCap = new(big.Int).Add(tip, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))
	} else {
		feeCap = new(big.Int).Set(tip)
	}
	if maxFeeCap != nil && feeCap.Cmp(maxFeeCap) > 0 {
		log.Debug("Applying gas fee cap limit", "dynamic", feeCap, "max", maxFeeCap)
		feeCap = new(big.Int).Set(maxFeeCap)
	}
	if tip.Cmp(feeCap) > 0 {
		log.Debug("Clamping tip to feeCap for EIP-1559 invariant", "tip", tip, "feeCap", feeCap)
		tip = new(big.Int).Set(feeCap)
	}

	return &GasCaps{TipCap: tip, FeeCap: feeCap}, nil
}
