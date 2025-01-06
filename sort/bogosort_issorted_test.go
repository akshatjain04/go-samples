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
			input:    []int{1, 2, 2, 3, 4, 4, 5},
			expected: true,
		},
		{
			name:     "Slice with negative and positive numbers",
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: true,
		},
		{
			name:     "Slice with maximum and minimum values",
			input:    []int{int(^uint(0) >> 1), -int(^uint(0) >> 1)},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSorted(tc.input)
			if actual != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, actual)
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}
func generateSortedSlice(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i
	}
	return slice
}
func generateUnsortedSlice(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(length)
	}
	return slice
}
