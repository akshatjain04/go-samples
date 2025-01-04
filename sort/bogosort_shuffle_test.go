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
		name        string
		input       []float64
		expectEmpty bool
	}{
		{
			name:        "Valid input with positive integers",
			input:       []float64{1, 2, 3, 4, 5},
			expectEmpty: false,
		},
		{
			name:        "Valid input with negative integers",
			input:       []float64{-1, -2, -3, -4, -5},
			expectEmpty: false,
		},
		{
			name:        "Valid input with floating-point numbers",
			input:       []float64{0.1, 0.2, 0.3, 0.4, 0.5},
			expectEmpty: false,
		},
		{
			name:        "Empty slice input",
			input:       []float64{},
			expectEmpty: true,
		},
		{
			name:        "Single element slice input",
			input:       []float64{42},
			expectEmpty: false,
		},
		{
			name:        "Large input slice",
			input:       makeLargeSlice(1000),
			expectEmpty: false,
		},
		{
			name:        "Input slice with duplicates",
			input:       []float64{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expectEmpty: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]float64, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if tc.expectEmpty {
				if len(tc.input) != 0 {
					t.Errorf("expected empty slice, got %v", tc.input)
				}
			} else {
				if !reflect.DeepEqual(sortFloat64Slice(original), sortFloat64Slice(tc.input)) {
					t.Errorf("shuffled slice %v does not contain the same elements as the original slice %v", tc.input, original)
				}
				if reflect.DeepEqual(original, tc.input) {
					t.Errorf("shuffled slice %v should not be in the same order as the original slice %v", tc.input, original)
				}
			}
			t.Logf("Success: %v", tc.name)
		})
	}
}
func makeLargeSlice(size int) []float64 {
	slice := make([]float64, size)
	for i := range slice {
		slice[i] = rand.Float64()
	}
	return slice
}


func sortFloat64Slice(slice []float64) []float64 {
	sorted := make([]float64, len(slice))
	copy(sorted, slice)
	sort.Float64s(sorted)
	return sorted
}

