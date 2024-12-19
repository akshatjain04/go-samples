package gcd

import "testing"

func TestExtendedRecursive(t *testing.T) {

	testCases := []struct {
		name        string
		a, b        int64
		expectedGCD int64
	}{
		{"Base Case with Coprime Integers", 35, 64, 1},
		{"Non-Coprime Integers", 48, 18, 6},
		{"Zero and Non-Zero Integer", 0, 5, 5},
		{"Both Integers Zero", 0, 0, 0},
		{"Negative Integers", -25, -10, 5},
		{"Large Integers", 9876543210, 1234567890, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			gcd, x, y := ExtendedRecursive(tc.a, tc.b)

			if gcd != tc.expectedGCD {
				t.Errorf("GCD(%d, %d) = %d; want %d", tc.a, tc.b, gcd, tc.expectedGCD)
			}

			if tc.a*x+tc.b*y != gcd {
				t.Errorf("Coefficients do not satisfy the equation: %d*%d + %d*%d = %d; want %d", tc.a, x, tc.b, y, tc.a*x+tc.b*y, gcd)
			} else {
				t.Logf("Coefficients satisfy the equation: %d*%d + %d*%d = %d", tc.a, x, tc.b, y, gcd)
			}
		})
	}
}
