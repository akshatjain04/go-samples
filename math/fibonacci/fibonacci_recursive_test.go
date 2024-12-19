package fibonacci

import (
	"math"
	"testing"
)

func TestRecursive(t *testing.T) {

	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{
			name:     "Base Case for Zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "Base Case for One",
			input:    1,
			expected: 1,
		},
		{
			name:     "Small Positive Number",
			input:    5,
			expected: 5,
		},
		{
			name:     "Handling Larger Number",
			input:    20,
			expected: 6765,
		},
		{
			name:     "Maximum Unsigned Int Boundary",
			input:    math.MaxUint32,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := Recursive(tc.input)

			if tc.name == "Maximum Unsigned Int Boundary" {

				t.Log("Testing for no runtime error with maximum uint value")
			} else if result != tc.expected {

				t.Errorf("Failed %s: Expected %d, got %d", tc.name, tc.expected, result)
			} else {

				t.Logf("Passed %s: Expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
