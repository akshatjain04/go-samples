package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    []float64
		expected func([]float64) bool
	}{
		{
			name:  "Valid input with positive integers",
			input: []float64{1, 2, 3, 4, 5},
			expected: func(input []float64) bool {
				return hasSameElements(input, []float64{1, 2, 3, 4, 5}) && !reflect.DeepEqual(input, []float64{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []float64{-1, -2, -3, -4, -5},
			expected: func(input []float64) bool {
				return hasSameElements(input, []float64{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(input []float64) bool {
				return hasSameElements(input, []float64{1.1, 2.2, 3.3, 4.4, 5.5})
			},
		},
		{
			name:     "Empty slice input",
			input:    []float64{},
			expected: func(input []float64) bool { return len(input) == 0 },
		},
		{
			name:     "Single element slice input",
			input:    []float64{42},
			expected: func(input []float64) bool { return reflect.DeepEqual(input, []float64{42}) },
		},
		{
			name:  "Large input slice",
			input: largeSlice(1000),
			expected: func(input []float64) bool {
				return hasSameElements(input, largeSlice(1000))
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []float64{1, 2, 2, 3, 3, 3},
			expected: func(input []float64) bool {
				return hasSameElements(input, []float64{1, 2, 2, 3, 3, 3})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			input := make([]float64, len(tc.input))
			copy(input, tc.input)

			shuffle(input)

			if !tc.expected(input) {
				t.Errorf("Testshuffle (%s): shuffled slice does not meet expected criteria", tc.name)
			} else {
				t.Logf("Testshuffle (%s): success", tc.name)
			}
		})
	}
}
func hasSameElements[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	elementCount := make(map[T]int)
	for _, item := range a {
		elementCount[item]++
	}

	for _, item := range b {
		if count, exists := elementCount[item]; !exists || count == 0 {
			return false
		}
		elementCount[item]--
	}

	return true
}
func largeSlice(size int) []float64 {
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice[i] = float64(i)
	}
	return slice
}
