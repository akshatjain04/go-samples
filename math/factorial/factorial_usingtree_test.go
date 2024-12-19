package factorial

import (
	"fmt"
	"math"
	"testing"
)

func TestUsingTree(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Factorial of negative number", -1, 0},
		{"Factorial of zero", 0, 1},
		{"Factorial of one", 1, 1},
		{"Factorial of two", 2, 2},
		{"Factorial of a small positive number", 5, 120},
		{"Factorial of a larger positive number", 10, 3628800},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UsingTree(tc.input)
			if result != tc.expected {
				t.Errorf("UsingTree(%d) = %d, want %d", tc.input, result, tc.expected)
			} else {
				t.Logf("UsingTree(%d) = %d, as expected", tc.input, result)
			}
		})
	}

	t.Run("Factorial of the maximum int value", func(t *testing.T) {
		maxInt := math.MaxInt64
		result := UsingTree(maxInt)

		t.Log("UsingTree(maxInt) = result", result)
	})
}
