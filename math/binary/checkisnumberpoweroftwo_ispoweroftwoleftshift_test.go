package binary

import (
	"fmt"
	"math"
	"testing"
)

func TestIsPowerOfTwoLeftShift(t *testing.T) {

	testCases := []struct {
		name   string
		input  uint
		expect bool
	}{
		{
			name:   "Known power of two (64)",
			input:  64,
			expect: true,
		},
		{
			name:   "Non-power of two (35)",
			input:  35,
			expect: false,
		},
		{
			name:   "Zero input",
			input:  0,
			expect: false,
		},
		{
			name:   "Maximum uint value",
			input:  uint(math.MaxUint64),
			expect: false,
		},
		{
			name:   "One as input",
			input:  1,
			expect: true,
		},
		{
			name:   "Smallest non-zero non-power of two (3)",
			input:  3,
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPowerOfTwoLeftShift(tc.input)
			if result != tc.expect {
				t.Errorf("IsPowerOfTwoLeftShift(%v) = %v, want %v", tc.input, result, tc.expect)
			} else {
				t.Logf("IsPowerOfTwoLeftShift(%v) = %v, as expected", tc.input, result)
			}
		})
	}
}
