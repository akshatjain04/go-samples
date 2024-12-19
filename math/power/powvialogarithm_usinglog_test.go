package power

import (
	"math"
	"power"
	"testing"
)

func TestUsingLog(t *testing.T) {
	testCases := []struct {
		name           string
		base           float64
		power          float64
		expectedResult float64
		expectError    bool
		delta          float64
	}{
		{
			name:           "Positive base with a fractional power",
			base:           5.0,
			power:          2.5,
			expectedResult: 55.9,
			delta:          0.1,
		},
		{
			name:           "Negative base with an even power",
			base:           -4.0,
			power:          2.0,
			expectedResult: 16.0,
			delta:          0.0,
		},
		{
			name:           "Negative base with an odd power",
			base:           -3.0,
			power:          3.0,
			expectedResult: -27.0,
			delta:          0.0,
		},
		{
			name:           "Base of zero with a positive power",
			base:           0.0,
			power:          5.0,
			expectedResult: 0.0,
			delta:          0.0,
		},
		{
			name:           "Positive base with zero power",
			base:           10.0,
			power:          0.0,
			expectedResult: 1.0,
			delta:          0.0,
		},
		{
			name:           "Base of zero with zero power",
			base:           0.0,
			power:          0.0,
			expectedResult: 1.0,
			delta:          0.0,
		},
		{
			name:           "Negative base with fractional power",
			base:           -2.0,
			power:          0.5,
			expectedResult: math.NaN(),
			delta:          0.0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := power.UsingLog(tc.base, tc.power)

			if tc.expectError {
				if !math.IsNaN(result) {
					t.Errorf("Expected NaN for undefined operation, got %v", result)
				}
			} else if math.Abs(result-tc.expectedResult) > tc.delta {
				t.Errorf("Expected result within %v of %v, got %v", tc.delta, tc.expectedResult, result)
			}

			t.Log("Computed result: ", result)
		})
	}
}

