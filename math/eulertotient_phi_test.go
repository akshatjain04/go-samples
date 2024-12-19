package math_test

import (
	"math"
	"testing"
)

// TestPhi tests the Phi function for various scenarios.
func TestPhi(t *testing.T) {
	// Define the test cases using a table-driven approach.
	testCases := []struct {
		name     string
		input    int64
		expected int64
	}{
		{
			name:     "Basic functionality with a prime number",
			input:    13,
			expected: 12,
		},
		{
			name:     "Test with a non-prime number",
			input:    10,
			expected: 4,
		},
		{
			name:     "Test with the number one",
			input:    1,
			expected: 1,
		},
		{
			name:     "Test with a perfect square",
			input:    9,
			expected: 6,
		},
		// TODO: Define the expected behavior for negative numbers in the Phi function.
		// The following test case assumes Phi returns 0 for negative inputs.
		{
			name:     "Test with a negative number",
			input:    -5,
			expected: 0, // Assuming the expected behavior is to return 0 for negative inputs.
		},
	}

	// Iterate over each test case.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act: Invoke the Phi function with the test input.
			result := math.Phi(tc.input)

			// Assert: Check if the result matches the expected value.
			if result != tc.expected {
				t.Errorf("Phi(%d) = %d, want %d", tc.input, result, tc.expected)
			} else {
				t.Logf("Phi(%d) = %d, as expected", tc.input, result)
			}
		})
	}
}
