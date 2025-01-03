package sort

import (
	"reflect"
	"testing"
	"math/rand"
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

	testCases := []struct {
		name     string
		input    []float64
		expected func([]float64) bool
	}{
		{
			name:  "Valid input with positive integers",
			input: []float64{1, 2, 3, 4, 5},
			expected: func(output []float64) bool {
				return reflect.DeepEqual(output, []float64{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []float64{-1, -2, -3, -4, -5},
			expected: func(output []float64) bool {
				return reflect.DeepEqual(output, []float64{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(output []float64) bool {
				return reflect.DeepEqual(output, []float64{1.1, 2.2, 3.3, 4.4, 5.5})
			},
		},
		{
			name:  "Empty slice input",
			input: []float64{},
			expected: func(output []float64) bool {
				return len(output) == 0
			},
		},
		{
			name:  "Single element slice input",
			input: []float64{42},
			expected: func(output []float64) bool {
				return reflect.DeepEqual(output, []float64{42})
			},
		},
		{
			name:  "Large input slice",
			input: makeLargeSlice(10000),
			expected: func(output []float64) bool {
				return len(output) == 10000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []float64{7, 7, 7, 7, 7},
			expected: func(output []float64) bool {
				return reflect.DeepEqual(output, []float64{7, 7, 7, 7, 7})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]float64, len(tc.input))
			copy(original, tc.input)
			shuffle(tc.input)

			if !tc.expected(original) {
				t.Errorf("shuffle() failed for %v; got %v, want %v", tc.name, tc.input, original)
			} else {
				t.Logf("shuffle() success for %v; got %v", tc.name, tc.input)
			}
		})
	}
}
func makeLargeSlice(size int) []float64 {
	largeSlice := make([]float64, size)
	for i := range largeSlice {
		largeSlice[i] = float64(i)
	}
	return largeSlice
}
