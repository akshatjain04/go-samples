package power

import "testing"

func TestRecursivePower1(t *testing.T) {

	testCases := []struct {
		name   string
		base   uint
		power  uint
		expect uint
	}{
		{"Base case with power zero", 2, 0, 1},
		{"Power of one", 3, 1, 3},
		{"Even power", 4, 2, 16},
		{"Odd power", 5, 3, 125},
		{"Large power", 2, 20, 1048576},
		{"Zero base, non-zero power", 0, 4, 0},
		{"Power of two", 2, 8, 256},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RecursivePower1(tc.base, tc.power)
			if result != tc.expect {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expect, result)
			} else {
				t.Logf("Passed %s: expected %v, got %v", tc.name, tc.expect, result)
			}
		})
	}
}
