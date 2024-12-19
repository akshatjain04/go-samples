package binary

import (
	"binary"
	"fmt"
	"testing"
)

func TestLogBase2(t *testing.T) {

	tests := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{
			name:     "Zero input",
			input:    0,
			expected: 0,
		},
		{
			name:     "Maximum uint32",
			input:    4294967295,
			expected: 31,
		},
		{
			name:     "Non-power-of-two",
			input:    20,
			expected: 4,
		},
		{
			name:     "Power of two minus one",
			input:    7,
			expected: 2,
		},
		{
			name:     "Random mid-range value",
			input:    123456,
			expected: 16,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			actual := binary.LogBase2(tc.input)

			if actual != tc.expected {
				t.Errorf("LogBase2(%d) = %d; want %d", tc.input, actual, tc.expected)
			} else {
				t.Logf("LogBase2(%d) = %d; as expected", tc.input, actual)
			}
		})
	}
}

