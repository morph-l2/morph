package main

import (
	"errors"
	"fmt"
)

var UnknownNonceVersion = errors.New("Unknown nonce version")

// checkErr checks if err is not nil, and throws if so.
// Shorthand to ease go's god awful error handling
func checkErr(err error, failReason string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", failReason, err))
	}
}
