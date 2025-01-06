package sort

import (
	"math/rand"
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
func TestisSorted(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Single element slice",
			input:    []int{5},
			expected: true,
		},
		{
			name:     "Sorted slice with multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Unsorted slice with multiple elements",
			input:    []int{5, 3, 4, 1, 2},
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
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 5},
			expected: true,
		},
		{
			name:     "Slice with negative and positive numbers",
			input:    []int{-3, -1, 0, 2, 4},
			expected: true,
		},
		{
			name:     "Slice with maximum and minimum values",
			input:    []int{constraints.MinInt, 0, constraints.MaxInt},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := isSorted(tc.input)

			if result != tc.expected {
				t.Errorf("For %s, expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Success for %s", tc.name)
			}
		})
	}
}
func generateSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = i
	}
	return slice
}
func generateUnsortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}
