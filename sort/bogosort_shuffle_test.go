package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"github.com/TheAlgorithms/Go/constraints"
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	type test[T constraints.Number] struct {
		name     string
		input    []T
		expected func([]T) bool
	}

	tests := []test[float64]{
		{
			name:  "Valid input with positive integers",
			input: []float64{1, 2, 3, 4, 5},
			expected: func(input []float64) bool {
				return reflect.DeepEqual(input, []float64{1, 2, 3, 4, 5})
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []float64{-1, -2, -3, -4, -5},
			expected: func(input []float64) bool {
				return reflect.DeepEqual(input, []float64{-1, -2, -3, -4, -5})
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: func(input []float64) bool {
				return reflect.DeepEqual(input, []float64{1.1, 2.2, 3.3, 4.4, 5.5})
			},
		},
		{
			name:  "Empty slice input",
			input: []float64{},
			expected: func(input []float64) bool {
				return len(input) == 0
			},
		},
		{
			name:  "Single element slice input",
			input: []float64{42},
			expected: func(input []float64) bool {
				return reflect.DeepEqual(input, []float64{42})
			},
		},
		{
			name:  "Large input slice",
			input: make([]float64, 10000),
			expected: func(input []float64) bool {
				return len(input) == 10000
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []float64{1, 2, 2, 3, 3, 3},
			expected: func(input []float64) bool {
				return reflect.DeepEqual(input, []float64{1, 2, 2, 3, 3, 3})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]float64, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if !tc.expected(original) {
				t.Errorf("shuffle() failed for %s, expected a permutation of %v, got %v", tc.name, original, tc.input)
			} else if reflect.DeepEqual(tc.input, original) {
				t.Errorf("shuffle() failed for %s, the order of elements did not change", tc.name)
			}
			t.Logf("shuffle() success for %s", tc.name)
		})
	}
}
