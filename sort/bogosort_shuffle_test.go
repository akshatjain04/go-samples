package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)


func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    []int
		expected func([]int) bool
	}{
		{
			name:  "Valid input with positive integers",
			input: []int{1, 2, 3, 4, 5},
			expected: func(shuffled []int) bool {
				return reflect.DeepEqual(shuffled, []int{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []int{-1, -2, -3, -4, -5},
			expected: func(shuffled []int) bool {
				return reflect.DeepEqual(shuffled, []int{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []int{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(shuffled []int) bool {
				return reflect.DeepEqual(shuffled, []int{1.1, 2.2, 3.3, 4.4, 5.5})
			},
		},
		{
			name:     "Empty slice input",
			input:    []int{},
			expected: func(shuffled []int) bool { return len(shuffled) == 0 },
		},
		{
			name:     "Single element slice input",
			input:    []int{42},
			expected: func(shuffled []int) bool { return len(shuffled) == 1 && shuffled[0] == 42 },
		},
		{
			name:  "Large input slice",
			input: make([]int, 1000),
			expected: func(shuffled []int) bool {
				return len(shuffled) == 1000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []int{5, 5, 3, 3, 1, 1},
			expected: func(shuffled []int) bool {
				return reflect.DeepEqual(shuffled, []int{5, 5, 3, 3, 1, 1})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if !tc.expected(tc.input) {
				t.Errorf("shuffle() = %v, want %v", tc.input, original)
			} else {
				t.Logf("shuffle() successfully shuffled the input slice")
			}

			if len(tc.input) != len(original) {
				t.Errorf("shuffled slice length = %v, want %v", len(tc.input), len(original))
			}

			if !containsSameElements(tc.input, original) {
				t.Errorf("shuffled slice does not contain the same elements as the original")
			}
		})
	}
}
func containsSameElements(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	counts := make(map[int]int)
	for _, item := range a {
		counts[item]++
	}
	for _, item := range b {
		if counts[item] == 0 {
			return false
		}
		counts[item]--
	}
	return true
}


