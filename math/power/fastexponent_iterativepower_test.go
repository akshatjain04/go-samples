package power

import (
	"fmt"
	"os"
	"testing"
)

func TestIterativePower(t *testing.T) {

	type testCase struct {
		n        uint
		power    uint
		expected uint
		name     string
	}

	testCases := []testCase{
		{n: 5, power: 0, expected: 1, name: "Base case with power of zero"},
		{n: 7, power: 1, expected: 7, name: "Power of one"},
		{n: 2, power: 3, expected: 8, name: "Normal exponentiation"},
		{n: 2, power: 30, expected: 1073741824, name: "Large exponent value"},
		{n: 1, power: 10, expected: 1, name: "Power operation with one as the base"},
		{n: 0, power: 5, expected: 0, name: "Zero base with non-zero power"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := IterativePower(tc.n, tc.power)

			if result != tc.expected {
				t.Errorf("IterativePower(%d, %d) = %d; want %d", tc.n, tc.power, result, tc.expected)
			} else {
				t.Logf("IterativePower(%d, %d) = %d, as expected", tc.n, tc.power, result)
			}
		})
	}
}
