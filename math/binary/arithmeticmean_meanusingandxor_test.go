package binary

import (
	"binary"
	"fmt"
	"testing"
)

func TestMeanUsingAndXor(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Valid input with different even numbers",
			a:        4,
			b:        8,
			expected: 6,
		},
		{
			name:     "Valid input with different odd numbers",
			a:        5,
			b:        9,
			expected: 7,
		},
		{
			name:     "Valid input with one even and one odd number",
			a:        6,
			b:        9,
			expected: 7,
		},
		{
			name:     "Valid input with identical numbers",
			a:        10,
			b:        10,
			expected: 10,
		},
		{
			name:     "Input with zero and a positive number",
			a:        0,
			b:        8,
			expected: 4,
		},
		{
			name:     "Input with negative numbers",
			a:        -4,
			b:        -8,
			expected: -6,
		},
		{
			name:     "Large integer input",
			a:        2147483647,
			b:        2147483646,
			expected: 2147483646,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binary.MeanUsingAndXor(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("MeanUsingAndXor(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			} else {
				t.Logf("Success: MeanUsingAndXor(%d, %d) = %d", tt.a, tt.b, result)
			}
		})
	}
}

