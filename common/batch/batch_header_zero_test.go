package batch

import (
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNewBatchHeaderSequencerSetVerifyHashIsZeroForAllVersions(t *testing.T) {
	tests := []struct {
		name     string
		upgraded bool
		v2       bool
		wantLen  int
	}{
		{name: "v0", wantLen: expectedLengthV0},
		{name: "v1", upgraded: true, wantLen: expectedLengthV1},
		{name: "v2", upgraded: true, v2: true, wantLen: expectedLengthV2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &BatchCache{
				isBatchUpgraded:   func(uint64) bool { return test.upgraded },
				isBatchV2Upgraded: func(uint64) bool { return test.v2 },
			}
			header := cache.createBatchHeader(common.Hash{1}, nil, 1)
			verifyHash, err := header.SequencerSetVerifyHash()
			require.NoError(t, err)
			require.Equal(t, common.Hash{}, verifyHash)
			require.Equal(t, BatchHeaderBytes(make([]byte, 32)), header[185:217])
			require.Len(t, header, test.wantLen)
		})
	}
}
