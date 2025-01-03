package sort

import (
	"math/rand"
	"testing"
	"time"
	"github.com/TheAlgorithms/Go/constraints"
)


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
			input:    createSortedSlice(1000),
			expected: true,
		},
		{
			name:     "Large unsorted slice",
			input:    createUnsortedSlice(1000),
			expected: false,
		},
		{
			name:     "Slice with duplicate elements",
			input:    []int{1, 2, 2, 3, 3},
			expected: true,
		},
		{
			name:     "Slice with negative and positive numbers",
			input:    []int{-2, -1, 0, 1, 2},
			expected: true,
		},
		{
			name:     "Slice with maximum and minimum values",
			input:    []int{constraints.MinInt, 0, constraints.MaxInt},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSorted(tt.input)
			if result != tt.expected {
				t.Errorf("isSorted(%v) = %v; expected %v", tt.input, result, tt.expected)
			} else {
				t.Logf("Success: isSorted(%v) correctly returned %v", tt.input, result)
			}
		})
	}
}
func createSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = i
	}
	return slice
}
func createUnsortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}

