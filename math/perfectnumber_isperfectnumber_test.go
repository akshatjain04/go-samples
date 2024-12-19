package math

import (
	"testing"
)

// TestIsPerfectNumber tests the IsPerfectNumber function for various scenarios
func TestIsPerfectNumber(t *testing.T) {
	// Define test cases in a table-driven style
	testCases := []struct {
		name           string
		input          uint
		expectedResult bool
	}{
		{"Perfect number check for a known perfect number", 28, true},
		{"Non-perfect number check", 10, false},
		{"Zero input check", 0, false},
		// Scenario 4 is omitted as negative numbers cannot be represented by uint
		{"Large perfect number check", 8128, true},
		{"Large non-perfect number check", 10000, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPerfectNumber(tc.input)
			if result != tc.expectedResult {
				t.Errorf("IsPerfectNumber(%d) = %v; want %v", tc.input, result, tc.expectedResult)
			} else {
				t.Logf("IsPerfectNumber(%d) = %v; as expected", tc.input, result)
			}
		})
	}
}
