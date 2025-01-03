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
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []int{-5, -4, -3, -2, -1},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{-5, -4, -3, -2, -1})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []int{1, 2, 3, 4, 5},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Empty slice input",
			input: []int{},
			expected: func(input []int) bool {
				return len(input) == 0
			},
		},
		{
			name:  "Single element slice input",
			input: []int{1},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1})
			},
		},
		{
			name:  "Large input slice",
			input: make([]int, 10000),
			expected: func(input []int) bool {
				return len(input) == 10000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.input))
			copy(original, tc.input)
			shuffle(tc.input)
			if !tc.expected(original) {
				t.Errorf("shuffle() with %s did not produce a valid permutation", tc.name)
			}
			if !reflect.DeepEqual(original, tc.input) {
				t.Log("The shuffled slice has the same length and contains the same elements as the original slice")
			} else {
				t.Errorf("shuffle() with %s failed, the elements are in the same order", tc.name)
			}
		})
	}
}


