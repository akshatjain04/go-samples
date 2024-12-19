package gcd

import "testing"

func TestExtended(t *testing.T) {

	testCases := []struct {
		name        string
		a           int64
		b           int64
		expectedGCD int64
		validate    func(t *testing.T, gcd, x, y int64)
	}{
		{
			name:        "Basic co-prime numbers",
			a:           17,
			b:           31,
			expectedGCD: 1,
			validate: func(t *testing.T, gcd, x, y int64) {
				if gcd != 1 || 17*x+31*y != 1 {
					t.Errorf("Expected GCD to be 1 and 17x+31y to equal 1, got GCD=%d, equation result=%d", gcd, 17*x+31*y)
				}
			},
		},
		{
			name:        "Zero and any number",
			a:           0,
			b:           42,
			expectedGCD: 42,
			validate: func(t *testing.T, gcd, x, y int64) {
				if gcd != 42 || x != 0 || y != 1 {
					t.Errorf("Expected GCD to be 42 and x=0, y=1, got GCD=%d, x=%d, y=%d", gcd, x, y)
				}
			},
		},
		{
			name:        "Negative numbers",
			a:           -25,
			b:           -10,
			expectedGCD: 5,
			validate: func(t *testing.T, gcd, x, y int64) {
				if gcd != 5 || -25*x-10*y != gcd {
					t.Errorf("Expected GCD to be positive and -25x-10y to equal GCD, got GCD=%d, equation result=%d", gcd, -25*x-10*y)
				}
			},
		},
		{
			name:        "Both inputs are zero",
			a:           0,
			b:           0,
			expectedGCD: 0,
			validate: func(t *testing.T, gcd, x, y int64) {
				if gcd != 0 {
					t.Errorf("Expected GCD to be 0, got GCD=%d", gcd)
				}
			},
		},
		{
			name:        "Large numbers",
			a:           9223372036854775807,
			b:           9223372036854775806,
			expectedGCD: 1,
			validate: func(t *testing.T, gcd, x, y int64) {
				if gcd != 1 || 9223372036854775807*x+9223372036854775806*y != 1 {
					t.Errorf("Expected GCD to be 1 and large numbers to satisfy the equation, got GCD=%d, equation result=%d", gcd, 9223372036854775807*x+9223372036854775806*y)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			gcd, x, y := Extended(tc.a, tc.b)

			t.Logf("Testing scenario: %s", tc.name)
			t.Logf("Input values: a=%d, b=%d", tc.a, tc.b)
			t.Logf("Expected GCD: %d", tc.expectedGCD)
			t.Logf("Received: GCD=%d, x=%d, y=%d", gcd, x, y)

			tc.validate(t, gcd, x, y)
		})
	}
}
