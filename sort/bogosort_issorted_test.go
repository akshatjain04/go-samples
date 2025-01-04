package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)


func TestisSorted(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
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
			input:    []int{5, 3, 2, 4, 1},
			expected: false,
		},
		{
			name:     "Large sorted slice",
			input:    makeSortedSlice(1000),
			expected: true,
		},
		{
			name:     "Large unsorted slice",
			input:    makeRandomSlice(1000),
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sort.IsSorted(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Passed %s", tc.name)
			}
		})
	}
}
func makeRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}
func makeSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}


