package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/log"
)

// GasCaps holds the calculated gas tip cap and gas fee cap
type GasCaps struct {
	TipCap *big.Int
	FeeCap *big.Int
}

// CalculateGasCaps calculates dynamic gas caps with optional max limits.
// It fetches the suggested tip and base fee from the network, applies configured
// max limits, and ensures the EIP-1559 invariant (tipCap <= feeCap) is maintained.
//
// If no max caps are configured, returns (nil, nil) to indicate caller should use default behavior.
// Use CalculateGasCapsAlways if you always need gas cap values.
func CalculateGasCaps(ctx context.Context, client *L2Client) (*GasCaps, error) {
	maxTipCap := client.GetMaxGasTipCap()
	maxFeeCap := client.GetMaxGasFeeCap()

	// If no caps configured, return nil to indicate default behavior
	if maxTipCap == nil && maxFeeCap == nil {
		return nil, nil
	}

	return doCalculateGasCaps(ctx, client, maxTipCap, maxFeeCap)
}

// CalculateGasCapsAlways calculates dynamic gas caps, always returning values.
// Use this when you need gas cap values regardless of whether max caps are configured.
func CalculateGasCapsAlways(ctx context.Context, client *L2Client) (*GasCaps, error) {
	maxTipCap := client.GetMaxGasTipCap()
	maxFeeCap := client.GetMaxGasFeeCap()
	return doCalculateGasCaps(ctx, client, maxTipCap, maxFeeCap)
}

func doCalculateGasCaps(ctx context.Context, client *L2Client, maxTipCap, maxFeeCap *big.Int) (*GasCaps, error) {

	// Get dynamic gas tip cap from network
	tip, err := client.GetClient().SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested gas tip cap: %w", err)
	}

	// Apply tip cap limit if configured
	if maxTipCap != nil && tip.Cmp(maxTipCap) > 0 {
		log.Debug("Applying gas tip cap limit", "dynamic", tip, "max", maxTipCap)
		tip = new(big.Int).Set(maxTipCap)
	}

	// Get base fee from latest block
	head, err := client.GetClient().HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %w", err)
	}

	// Calculate dynamic gas fee cap: tip + baseFee * 2
	var feeCap *big.Int
	if head.BaseFee != nil {
		feeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		feeCap = new(big.Int).Set(tip)
	}

	// Apply fee cap limit if configured
	if maxFeeCap != nil && feeCap.Cmp(maxFeeCap) > 0 {
		log.Debug("Applying gas fee cap limit", "dynamic", feeCap, "max", maxFeeCap)
		feeCap = new(big.Int).Set(maxFeeCap)
	}

	// Ensure tipCap <= feeCap (EIP-1559 invariant)
	if tip.Cmp(feeCap) > 0 {
		log.Debug("Clamping tip to feeCap for EIP-1559 invariant", "tip", tip, "feeCap", feeCap)
		tip = new(big.Int).Set(feeCap)
	}

	log.Debug("Gas caps calculated", "tipCap", tip, "feeCap", feeCap)

	return &GasCaps{
		TipCap: tip,
		FeeCap: feeCap,
	}, nil
}

