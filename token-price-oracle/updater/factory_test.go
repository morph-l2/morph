package updater

import "testing"

func TestRedactRPCForLog(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "path API key",
			raw:  "https://rpc.example.com/v3/secret?token=also-secret",
			want: "https://rpc.example.com",
		},
		{
			name: "credentials",
			raw:  "https://user:password@rpc.example.com/path",
			want: "https://rpc.example.com",
		},
		{
			name: "invalid",
			raw:  "not-a-url",
			want: "<invalid_rpc_url>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redactRPCForLog(tt.raw); got != tt.want {
				t.Fatalf("redactRPCForLog() = %q, want %q", got, tt.want)
			}
		})
	}
}
