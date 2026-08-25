package derivation

import "testing"

func TestEnforceBatchRootVerification(t *testing.T) {
	tests := []struct {
		name       string
		verifyMode string
		want       bool
	}{
		{
			name:       "layer1 validator enforces",
			verifyMode: VerifyModeLayer1,
			want:       true,
		},
		{
			name:       "local node observes",
			verifyMode: VerifyModeLocal,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enforceBatchRootVerification(tt.verifyMode); got != tt.want {
				t.Fatalf("enforceBatchRootVerification(%q) = %t, want %t", tt.verifyMode, got, tt.want)
			}
		})
	}
}
