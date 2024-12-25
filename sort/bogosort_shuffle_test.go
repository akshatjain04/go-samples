package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
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
			input: []int{-1, -2, -3, -4, -5},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []int{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1.1, 2.2, 3.3, 4.4, 5.5})
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
			input: []int{42},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{42})
			},
		},
		{
			name:  "Large input slice",
			input: make([]int, 1000),
			expected: func(input []int) bool {
				return len(input) == 1000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []int{1, 2, 2, 3, 3, 3},
			expected: func(input []int) bool {

				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			original := make([]int, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if !tc.expected(tc.input) || len(original) != len(tc.input) {
				t.Errorf("shuffle() failed for %s, expected a permutation of %v, got %v", tc.name, original, tc.input)
			} else {
				t.Logf("shuffle() success for %s", tc.name)
			}

			if !isPermutation(original, tc.input) {
				t.Errorf("shuffle() failed for %s, result %v is not a permutation of %v", tc.name, tc.input, original)
			}
		})
	}
}
func isPermutation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[int]int)
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
