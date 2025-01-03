package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/TheAlgorithms/Go/constraints"
)

// TestisSorted is a unit test function for the isSorted function.
func TestisSorted(t *testing.T) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Test scenarios
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
		// Scenario 10 is not applicable as it requires a compile-time check, which cannot be done in a runtime test
	}

	// Run test scenarios
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

// createSortedSlice creates a large sorted slice of random numbers.
func createSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

// createUnsortedSlice creates a large unsorted slice of random numbers.
func createUnsortedSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}

// Note: Scenario 10 cannot be tested in a unit test as it is a compile-time error
// to pass a slice of non-number type to the isSorted function, which expects a slice
// of constraints.Number. This is enforced by the Go type system and does not need a
// runtime test.
