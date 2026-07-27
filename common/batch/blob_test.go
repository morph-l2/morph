package batch

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractInnerTxFullBytesRejectsOversizedDeclaration(t *testing.T) {
	reader := bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff})

	_, err := extractInnerTxFullBytes(0xfb, reader)

	require.EqualError(t, err, "declared tx size 4294967295 exceeds remaining 0 bytes")
}

func TestExtractInnerTxFullBytesAcceptsAvailablePayload(t *testing.T) {
	reader := bytes.NewReader([]byte{3, 1, 2, 3})

	got, err := extractInnerTxFullBytes(0xf8, reader)

	require.NoError(t, err)
	require.Equal(t, []byte{0xf8, 3, 1, 2, 3}, got)
}
