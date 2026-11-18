package service

import (
	"fmt"

	"morph-l2/bindings/bindings"
)

func challengeMethodSig() (string, error) {
	parsed, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("failed to parse rollup ABI: %w", err)
	}
	method, ok := parsed.Methods["challengeState"]
	if !ok {
		return "", fmt.Errorf("method challengeState not found in rollup ABI")
	}
	return method.Sig, nil
}

// probeCallData encodes a challengeState call with placeholder arguments.
// It has the selector and length of a real call and is only used to probe the remote signer.
func probeCallData() ([]byte, error) {
	parsed, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to parse rollup ABI: %w", err)
	}
	return parsed.Pack("challengeState", uint64(0), [32]byte{})
}
