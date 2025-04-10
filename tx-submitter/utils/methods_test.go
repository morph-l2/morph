package utils

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseStringToType(t *testing.T) {
	tests := []struct {
		input    string
		expected any
		hasError bool
	}{
		// int cases
		{"123", int(123), false},
		{"-123", int(-123), false},
		{"notanumber", int(0), true},

		// int8 cases
		{"123", int8(123), false},
		{"-128", int8(-128), false},
		{"notanumber", int8(0), true},

		// int16 cases
		{"123", int16(123), false},
		{"32767", int16(32767), false},
		{"notanumber", int16(0), true},

		// uint cases
		{"123", uint(123), false},
		{"notanumber", uint(0), true},

		// float32 cases
		{"123.45", float32(123.45), false},
		{"-123.45", float32(-123.45), false},
		{"notanumber", float32(0), true},

		// float64 cases
		{"123.45", float64(123.45), false},
		{"-123.45", float64(-123.45), false},
		{"notanumber", float64(0), true},

		// bool cases
		{"true", true, false},
		{"false", false, false},
		{"notabool", false, true},

		// string cases
		{"test", "test", false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tc.input, tc.expected), func(t *testing.T) {
			var result any
			var err error

			switch tc.expected.(type) {
			case int:
				result, err = ParseStringToType[int](tc.input)
			case int8:
				result, err = ParseStringToType[int8](tc.input)
			case int16:
				result, err = ParseStringToType[int16](tc.input)
			case uint:
				result, err = ParseStringToType[uint](tc.input)
			case float32:
				result, err = ParseStringToType[float32](tc.input)
			case float64:
				result, err = ParseStringToType[float64](tc.input)
			case bool:
				result, err = ParseStringToType[bool](tc.input)
			case string:
				result, err = ParseStringToType[string](tc.input)
			}

			if tc.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestWeiToGwei(t *testing.T) {
	tests := []struct {
		name     string
		input    *big.Int
		expected string
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: "0",
		},
		{
			name:     "zero wei",
			input:    big.NewInt(0),
			expected: "0.000000000",
		},
		{
			name:     "1 wei",
			input:    big.NewInt(1),
			expected: "0.000000001",
		},
		{
			name:     "1000 wei",
			input:    big.NewInt(1000),
			expected: "0.000001000",
		},
		{
			name:     "1 million wei",
			input:    big.NewInt(1000000),
			expected: "0.001000000",
		},
		{
			name:     "1 gwei",
			input:    big.NewInt(1e9),
			expected: "1.000000000",
		},
		{
			name:     "1.5 gwei",
			input:    big.NewInt(1.5e9),
			expected: "1.500000000",
		},
		{
			name:     "precise small number",
			input:    big.NewInt(123456789),
			expected: "0.123456789",
		},
		{
			name:     "large number",
			input:    new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000)),
			expected: "1000000000000.000000000",
		},
		{
			name:     "max precision test",
			input:    big.NewInt(1234567891),
			expected: "1.234567891",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WeiToGwei(tt.input)
			if result != tt.expected {
				t.Errorf("WeiToGwei(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
