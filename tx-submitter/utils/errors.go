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
