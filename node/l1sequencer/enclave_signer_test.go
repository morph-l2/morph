package l1sequencer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Wire-protocol tests are deliberately omitted: the client is
// vsock-only, and there is no portable way to stand up a vsock
// listener from a test (mdlayher/vsock is Linux-only and CI usually
// lacks vsock support). The wire protocol is exercised end-to-end on
// real Nitro instances; here we only test the pure-Go addr parsing.

func TestParseAddr(t *testing.T) {
	cases := []struct {
		in       string
		wantCID  uint32
		wantPort uint32
		wantErr  bool
	}{
		{"16:5000", 16, 5000, false},
		{"3:8000", 3, 8000, false},
		{"0:1", 0, 1, false},
		{"abc:5000", 0, 0, true},
		{"16:bad", 0, 0, true},
		{"16", 0, 0, true},
		{"", 0, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			cid, port, err := parseAddr(tc.in)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.wantCID, cid)
			require.Equal(t, tc.wantPort, port)
		})
	}
}
