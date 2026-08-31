package constants

import "testing"

func TestIsRollupMethod(t *testing.T) {
	t.Parallel()
	tests := []struct {
		method string
		want   bool
	}{
		{MethodCommitBatch, true},
		{MethodCommitState, true},
		{MethodFinalizeBatch, true},
		{"transfer", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			t.Parallel()
			if got := IsRollupMethod(tt.method); got != tt.want {
				t.Fatalf("IsRollupMethod(%q) = %v, want %v", tt.method, got, tt.want)
			}
		})
	}
}
