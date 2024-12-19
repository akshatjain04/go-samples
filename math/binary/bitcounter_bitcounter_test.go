package binary

import (
	"fmt"
	"os"
	"testing"
)

func TestBitCounter(t *testing.T) {

	testCases := []struct {
		name     string
		input    uint
		expected int
	}{
		{
			name:     "Alternating bits",
			input:    5,
			expected: 2,
		},
		{
			name:     "Power of two",
			input:    32,
			expected: 1,
		},
		{
			name:     "Maximum unsigned integer",
			input:    ^uint(0),
			expected: 32,
		},
		{
			name:     "Single bit set",
			input:    2,
			expected: 1,
		},
		{
			name:     "All bits set",
			input:    31,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BitCounter(tc.input)
			if result != tc.expected {
				t.Errorf("BitCounter(%d) = %d, want %d", tc.input, result, tc.expected)
			} else {
				t.Logf("BitCounter(%d) = %d, as expected", tc.input, result)
			}
		})
	}
}
