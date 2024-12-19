package math_test

import (
	"math"
	"testing"
)

// TestCos will test the Cos function in the math package.
func TestCos(t *testing.T) {
	// Define a struct for test cases
	type testCase struct {
		name     string
		input    float64
		expected float64
		delta    float64
	}

	// Initialize test cases
	testCases := []testCase{
		{
			name:     "Validate Cosine of Zero",
			input:    0,
			expected: 1,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Pi/2",
			input:    math.Pi / 2,
			expected: 0,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Pi",
			input:    math.Pi,
			expected: -1,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Negative Input",
			input:    -math.Pi,
			expected: -1,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Large Input",
			input:    1000 * math.Pi,
			expected: 1, // cos(1000 * Pi) is equivalent to cos(0)
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Very Small Input",
			input:    1e-10,
			expected: 1,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Half Pi Multiple",
			input:    3 * math.Pi / 2,
			expected: 0,
			delta:    0.000001,
		},
		{
			name:     "Validate Cosine of Two Pi Multiple",
			input:    2 * math.Pi,
			expected: 1,
			delta:    0.000001,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Executing test case:", tc.name)

			// Act
			result := math.Cos(tc.input)

			// Assert
			if math.Abs(result-tc.expected) > tc.delta {
				t.Errorf("Cos(%v) = %v; want %v (within delta %v)", tc.input, result, tc.expected, tc.delta)
			} else {
				t.Logf("Success: Cos(%v) = %v; matches expected value %v (within delta %v)", tc.input, result, tc.expected, tc.delta)
			}
		})
	}
}
