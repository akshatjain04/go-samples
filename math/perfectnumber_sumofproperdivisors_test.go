package math

import (
	"testing"
)

// TestSumOfProperDivisors is a unit test for the SumOfProperDivisors function.
func TestSumOfProperDivisors(t *testing.T) {
	// Define test cases in a table-driven format.
	tests := []struct {
		name     string   // Name of the test case
		input    uint     // Input value for testing
		expected uint     // Expected result
	}{
		{
			name:     "Zero input value",
			input:    0,
			expected: 0,
		},
		{
			name:     "Input value is one",
			input:    1,
			expected: 0,
		},
		{
			name:     "Prime number input",
			input:    7,
			expected: 1,
		},
		{
			name:     "Composite number input",
			input:    10,
			expected: 8,
		},
		{
			name:     "Large number input",
			input:    10000,
			expected: 14211, // TODO: Replace with the actual expected sum of proper divisors for 10000
		},
		{
			name:     "Perfect square input",
			input:    16,
			expected: 15,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := SumOfProperDivisors(tc.input)
			if result != tc.expected {
				t.Errorf("SumOfProperDivisors(%d) = %d; want %d", tc.input, result, tc.expected)
			} else {
				t.Logf("SumOfProperDivisors(%d) = %d, as expected", tc.input, result)
			}
		})
	}
}

// Please note that the actual expected sum of proper divisors for the "Large number input" test case should be calculated
// and provided in place of the TODO comment. The correctness of the test depends on the accuracy of this value.
