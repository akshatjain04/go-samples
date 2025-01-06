package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"github.com/TheAlgorithms/Go/constraints"
)

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
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
			input: []int{1.5, 2.5, 3.5},
			expected: func(input []int) bool {
				return reflect.DeepEqual(input, []int{1.5, 2.5, 3.5})
			},
		},
		{
			name:     "Empty slice input",
			input:    []int{},
			expected: func(input []int) bool { return len(input) == 0 },
		},
		{
			name:     "Single element slice input",
			input:    []int{42},
			expected: func(input []int) bool { return len(input) == 1 && input[0] == 42 },
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
				duplicates := map[int]int{1: 1, 2: 2, 3: 3}
				for _, v := range input {
					duplicates[v]--
					if duplicates[v] < 0 {
						return false
					}
				}
				for _, count := range duplicates {
					if count != 0 {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if !tc.expected(tc.input) {
				t.Errorf("Testshuffle failed for %s: got %v, want %v", tc.name, tc.input, original)
			} else {
				t.Logf("Testshuffle passed for %s", tc.name)
			}

			if len(original) > 0 && len(tc.input) > 0 && reflect.DeepEqual(original, tc.input) {
				t.Errorf("Testshuffle failed for %s: slice was not shuffled", tc.name)
			} else {
				t.Logf("Testshuffle passed for %s: slice was shuffled", tc.name)
			}
		})
	}
}
