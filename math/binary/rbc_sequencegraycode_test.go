package binary

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestSequenceGrayCode(t *testing.T) {

	testCases := []struct {
		name     string
		input    uint
		expected []uint
	}{
		{
			name:     "Zero input value",
			input:    0,
			expected: []uint{},
		},

		{
			name:  "Large input value",
			input: 20,
		},
		{
			name:     "Typical input value",
			input:    3,
			expected: []uint{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			name:  "Verify first and last values",
			input: 4,
		},
		{
			name:  "Sequence uniqueness and length",
			input: 5,
		},
		{
			name:  "Check for consecutive elements differing by one bit",
			input: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SequenceGrayCode(tc.input)

			switch tc.name {
			case "Zero input value", "Typical input value":
				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("SequenceGrayCode(%d) = %v, want %v", tc.input, result, tc.expected)
				} else {
					t.Logf("Success: Expected result for input %d", tc.input)
				}
			case "Large input value":
				if len(result) != int(math.Pow(2, float64(tc.input))) {
					t.Errorf("SequenceGrayCode(%d): expected length %d, got %d", tc.input, 1<<tc.input, len(result))
				} else {
					t.Logf("Success: Correct length for large input value %d", tc.input)
				}
			case "Verify first and last values":
				if result[0] != 0 || result[len(result)-1] != (1<<tc.input)-1 {
					t.Errorf("SequenceGrayCode(%d): incorrect first or last values, got first: %d, last: %d", tc.input, result[0], result[len(result)-1])
				} else {
					t.Logf("Success: Correct first and last values for input %d", tc.input)
				}
			case "Sequence uniqueness and length":
				set := make(map[uint]bool)
				for _, num := range result {
					set[num] = true
				}
				if len(set) != len(result) || len(result) != int(math.Pow(2, float64(tc.input))) {
					t.Errorf("SequenceGrayCode(%d): sequence contains duplicates or incorrect length", tc.input)
				} else {
					t.Logf("Success: Unique sequence and correct length for input %d", tc.input)
				}
			case "Check for consecutive elements differing by one bit":
				for i := 0; i < len(result)-1; i++ {
					if bits.OnesCount(uint(result[i]^result[i+1])) != 1 {
						t.Errorf("SequenceGrayCode(%d): consecutive elements %d and %d do not differ by one bit", tc.input, result[i], result[i+1])
					}
				}
				t.Logf("Success: Consecutive elements differ by one bit for input %d", tc.input)
			default:
				t.Errorf("Unknown test case: %s", tc.name)
			}
		})
	}
}
