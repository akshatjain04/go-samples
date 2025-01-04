package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"github.com/TheAlgorithms/Go/constraints"
)




func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    []constraints.Number
		expected func([]constraints.Number) bool
	}{
		{
			name:  "Valid input with positive integers",
			input: []constraints.Number{1, 2, 3, 4, 5},
			expected: func(input []constraints.Number) bool {

				return reflect.DeepEqual(input, []constraints.Number{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []constraints.Number{-1, -2, -3, -4, -5},
			expected: func(input []constraints.Number) bool {
				return reflect.DeepEqual(input, []constraints.Number{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []constraints.Number{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(input []constraints.Number) bool {
				return reflect.DeepEqual(input, []constraints.Number{1.1, 2.2, 3.3, 4.4, 5.5})
			},
		},
		{
			name:  "Empty slice input",
			input: []constraints.Number{},
			expected: func(input []constraints.Number) bool {
				return len(input) == 0
			},
		},
		{
			name:  "Single element slice input",
			input: []constraints.Number{42},
			expected: func(input []constraints.Number) bool {
				return reflect.DeepEqual(input, []constraints.Number{42})
			},
		},
		{
			name:  "Large input slice",
			input: makeLargeSlice(1000),
			expected: func(input []constraints.Number) bool {
				return len(input) == 1000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []constraints.Number{1, 1, 2, 2, 3, 3},
			expected: func(input []constraints.Number) bool {
				return reflect.DeepEqual(input, []constraints.Number{1, 1, 2, 2, 3, 3})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]constraints.Number, len(tt.input))
			copy(original, tt.input)

			shuffle(tt.input)

			if !tt.expected(original) {
				t.Errorf("shuffle() = %v, want %v", tt.input, original)
			} else {
				t.Logf("shuffle() test passed for %v", tt.name)
			}

			if !isPermutation(original, tt.input) {
				t.Errorf("shuffle() result %v is not a permutation of %v", tt.input, original)
			}
		})
	}
}
func isPermutation[T constraints.Number](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[T]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}
func makeLargeSlice(size int) []constraints.Number {
	largeSlice := make([]constraints.Number, size)
	for i := range largeSlice {
		largeSlice[i] = constraints.Number(i)
	}
	return largeSlice
}


