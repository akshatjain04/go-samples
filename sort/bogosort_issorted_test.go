package sort

import (
	"math/rand"
	"testing"
	"time"
	"github.com/TheAlgorithms/Go/constraints"
)



func TestisSorted(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name     string
		input    []constraints.Number
		expected bool
	}{
		{
			name:     "Empty slice",
			input:    []constraints.Number{},
			expected: true,
		},
		{
			name:     "Single element slice",
			input:    []constraints.Number{1},
			expected: true,
		},
		{
			name:     "Sorted slice with multiple elements",
			input:    []constraints.Number{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Unsorted slice with multiple elements",
			input:    []constraints.Number{5, 3, 4, 1, 2},
			expected: false,
		},
		{
			name:     "Large sorted slice",
			input:    generateSortedSlice(1000),
			expected: true,
		},
		{
			name:     "Large unsorted slice",
			input:    generateUnsortedSlice(1000),
			expected: false,
		},
		{
			name:     "Slice with duplicate elements",
			input:    []constraints.Number{1, 2, 2, 3, 3, 4, 4, 5},
			expected: true,
		},
		{
			name:     "Slice with negative and positive numbers",
			input:    []constraints.Number{-3, -2, -1, 0, 1, 2, 3},
			expected: true,
		},
		{
			name:     "Slice with maximum and minimum values",
			input:    []constraints.Number{constraints.MinInt, 0, constraints.MaxInt},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isSorted(tc.input)
			if result != tc.expected {
				t.Errorf("For %v, expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Success: %v correctly identified as %v", tc.name, tc.expected)
			}
		})
	}
}
func generateSortedSlice(size int) []constraints.Number {
	slice := make([]constraints.Number, size)
	for i := range slice {
		slice[i] = constraints.Number(i)
	}
	return slice
}
func generateUnsortedSlice(size int) []constraints.Number {
	slice := make([]constraints.Number, size)
	for i := range slice {
		slice[i] = constraints.Number(rand.Intn(size))
	}
	return slice
}

