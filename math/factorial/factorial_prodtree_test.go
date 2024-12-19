package factorial

import (
	"fmt"
	"testing"
)

func TestprodTree(t *testing.T) {

	type testCase struct {
		name     string
		l        int
		r        int
		expected int
	}

	tests := []testCase{
		{"Equal bounds", 5, 5, 5},
		{"Consecutive numbers", 3, 4, 12},
		{"Range of numbers", 1, 4, 24},
		{"Left bound greater than right", 4, 3, 1},
		{"Large range", 1, 10, 3628800},
		{"Zero as left bound", 0, 2, 0},
		{"Negative bounds", -3, -1, -6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := prodTree(tc.l, tc.r)

			if result != tc.expected {
				t.Errorf("prodTree(%d, %d) = %d, want %d", tc.l, tc.r, result, tc.expected)
			} else {
				t.Logf("Success: prodTree(%d, %d) = %d", tc.l, tc.r, result)
			}
		})
	}
}
