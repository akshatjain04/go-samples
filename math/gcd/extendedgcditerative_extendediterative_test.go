package gcd

import (
	"testing"
	"gcd"
)

func TestExtendedIterative(t *testing.T) {

	testCases := []struct {
		name           string
		a, b           int64
		expectedGCD    int64
		expectedCoeffX int64
		expectedCoeffY int64
	}{
		{
			name:           "GCD where a < b",
			a:              15,
			b:              35,
			expectedGCD:    5,
			expectedCoeffX: -2,
			expectedCoeffY: 1,
		},
		{
			name:           "GCD where a > b",
			a:              1071,
			b:              462,
			expectedGCD:    21,
			expectedCoeffX: -2,
			expectedCoeffY: 5,
		},
		{
			name:           "GCD of two identical integers",
			a:              56,
			b:              56,
			expectedGCD:    56,
			expectedCoeffX: 0,
			expectedCoeffY: 1,
		},
		{
			name:           "GCD with one integer being zero",
			a:              0,
			b:              42,
			expectedGCD:    42,
			expectedCoeffX: 0,
			expectedCoeffY: 1,
		},
		{
			name:           "GCD with both integers being zero",
			a:              0,
			b:              0,
			expectedGCD:    0,
			expectedCoeffX: 0,
			expectedCoeffY: 0,
		},
		{
			name:           "GCD with negative integers",
			a:              -81,
			b:              -153,
			expectedGCD:    9,
			expectedCoeffX: -1,
			expectedCoeffY: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gcd, x, y := gcd.ExtendedIterative(tc.a, tc.b)

			if gcd != tc.expectedGCD {
				t.Errorf("Expected GCD: %d, Got: %d", tc.expectedGCD, gcd)
			} else {
				t.Logf("Success for GCD: Expected %d, Got %d", tc.expectedGCD, gcd)
			}

			if tc.a*x+tc.b*y != tc.expectedGCD {
				t.Errorf("Bézout coefficients do not satisfy the equation for a=%d and b=%d", tc.a, tc.b)
			} else {
				t.Logf("Success for Bézout coefficients: a=%d, b=%d, x=%d, y=%d", tc.a, tc.b, x, y)
			}
		})
	}
}

