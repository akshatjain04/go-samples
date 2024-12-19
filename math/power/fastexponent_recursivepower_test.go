package power

import "testing"

func TestRecursivePower(t *testing.T) {

	tests := []struct {
		name   string
		base   uint
		power  uint
		expect uint
	}{
		{"Base case of power zero", 5, 0, 1},
		{"Power of one", 5, 1, 5},
		{"Even power", 2, 4, 16},
		{"Odd power", 2, 3, 8},
		{"Large base and power", 10, 5, 100000},
		{"Minimum non-zero base", 1, 10, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := RecursivePower(tc.base, tc.power)

			if result != tc.expect {
				t.Errorf("RecursivePower(%d, %d) = %d; expected %d", tc.base, tc.power, result, tc.expect)
			} else {
				t.Logf("RecursivePower(%d, %d) = %d, as expected", tc.base, tc.power, result)
			}
		})
	}
}
