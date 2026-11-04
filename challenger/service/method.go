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
