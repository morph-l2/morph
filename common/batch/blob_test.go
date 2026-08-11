package batch

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractInnerTxFullBytesRejectsOversizedDeclaration(t *testing.T) {
	reader := bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff})

	_, err := extractInnerTxFullBytes(0xfb, reader)

	require.EqualError(t, err, "declared tx size 4294967295 exceeds remaining 0 bytes")
}

func TestExtractInnerTxFullBytesRejectsReaderWithoutLen(t *testing.T) {
	// io.MultiReader wraps the byte reader in a type that does not expose Len(),
	// so the declared size cannot be bounded and the decode must fail closed
	// rather than allocate an attacker-controlled length.
	reader := io.MultiReader(bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 1}))

	_, err := extractInnerTxFullBytes(0xfb, reader)

	require.ErrorContains(t, err, "does not report remaining length")
}

func TestExtractInnerTxFullBytesAcceptsAvailablePayload(t *testing.T) {
	reader := bytes.NewReader([]byte{3, 1, 2, 3})

	got, err := extractInnerTxFullBytes(0xf8, reader)

	require.NoError(t, err)
	require.Equal(t, []byte{0xf8, 3, 1, 2, 3}, got)
}
