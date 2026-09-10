package utils

import (
	"errors"
	"strings"
)

var ErrExceedFeeLimit = errors.New("exceed fee limit")

// ErrStringMatch returns true if err.Error() is a substring in target.Error() or if both are nil.
// It can accept nil errors without issue.
func ErrStringMatch(err, target error) bool {
	if err == nil && target == nil {
		return true
	} else if err == nil || target == nil {
		return false
	}
	return strings.Contains(err.Error(), target.Error())
}

var rpcErrTargets = []string{
	"timeout",
	"connection refused",
	"connection reset",
	"connection closed",
}

func IsRpcErr(err error) bool {
	if err == nil {
		return false
	}
	for _, target := range rpcErrTargets {
		if strings.Contains(err.Error(), target) {
			return true
		}
	}
	return false
}

// revertErrTargets are the eth_estimateGas failures that mean the node
// simulated the call and the contract rejected it.
var revertErrTargets = []string{
	"execution reverted",
	"always failing transaction",
}

// IsExecutionRevertErr reports whether an eth_estimateGas failure came from the
// contract rejecting the call rather than from a transport or liveness problem.
// A revert does not go away by guessing a gas limit: submitting anyway mines a
// failing transaction, and for a blob tx the blob fee is charged in full even
// though execution reverted.
func IsExecutionRevertErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	for _, target := range revertErrTargets {
		if strings.Contains(msg, target) {
			return true
		}
	}
	return false
}
