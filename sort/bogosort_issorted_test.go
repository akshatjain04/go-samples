package sort

import (
	"math/rand"
	"sort"
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
			input:    []int{1},
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
			input:    []int{1, 2, 2, 3, 4},
			expected: true,
		},
		{
			name:     "Slice with negative and positive numbers",
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
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
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed: expected %v, got %v", tc.name, tc.expected, result)
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
	rand.Seed(time.Now().UnixNano())
	slice := generateSortedSlice(size)
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}
