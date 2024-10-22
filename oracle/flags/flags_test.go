package flags

import (
	"fmt"
	"github.com/morph-l2/go-ethereum/crypto"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

// TestRequiredFlagsSetRequired asserts that all flags deemed required properly
// have the Required field set to true.
func TestRequiredFlagsSetRequired(t *testing.T) {
	for _, flag := range requiredFlags {
		reqFlag, ok := flag.(cli.RequiredFlag)
		require.True(t, ok)
		require.True(t, reqFlag.IsRequired())
	}
}

// TestOptionalFlagsDontSetRequired asserts that all flags deemed optional set
// the Required field to false.
func TestOptionalFlagsDontSetRequired(t *testing.T) {
	for _, flag := range optionalFlags {
		reqFlag, ok := flag.(cli.RequiredFlag)
		require.True(t, ok)
		require.False(t, reqFlag.IsRequired())
	}
}

func TestName(t *testing.T) {

	fmt.Println(len("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	fmt.Println(len("0x0000000000000000000000000000000000000000000000000000000000000000"))
	//ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
	hex := strings.TrimPrefix("0x0000000000000000000000000000000000000000000000000000000000000001", "0x")
	privateKey, err := crypto.HexToECDSA(hex)
	require.NoError(t, err)
	fmt.Println(privateKey)
}
