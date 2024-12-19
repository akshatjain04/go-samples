package gcd

import "testing"

func TestRecursive(t *testing.T) {

	tests := []struct {
		name   string
		a      int64
		b      int64
		expect int64
	}{
		{"Basic functionality with two positive integers", 48, 18, 6},
		{"Handling one zero value", 0, 25, 25},
		{"Both values are zero", 0, 0, 0},
		{"Negative integers input", -50, -20, 10},
		{"Large integers input", 9223372036854775807, 9223372036854775806, 1},
		{"Coprime integers input", 13, 28, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Recursive(tc.a, tc.b)
			if result != tc.expect {
				t.Errorf("Failed %s: Expected %v, got %v", tc.name, tc.expect, result)
			} else {
				t.Logf("Passed %s: Expected %v, got %v", tc.name, tc.expect, result)
			}
		})
	}
}
