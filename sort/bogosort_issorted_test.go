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
		scenario      string
		input         []int
		expected      bool
		expectCompile bool
	}{
		{
			scenario:      "Empty slice",
			input:         []int{},
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Single element slice",
			input:         []int{42},
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Sorted slice with multiple elements",
			input:         []int{1, 2, 3, 4, 5},
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Unsorted slice with multiple elements",
			input:         []int{5, 3, 4, 1, 2},
			expected:      false,
			expectCompile: true,
		},
		{
			scenario:      "Large sorted slice",
			input:         generateSortedSlice(1000),
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Large unsorted slice",
			input:         generateRandomSlice(1000),
			expected:      false,
			expectCompile: true,
		},
		{
			scenario:      "Slice with duplicate elements",
			input:         []int{1, 2, 2, 3, 3, 4, 4, 5, 5},
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Slice with negative and positive numbers",
			input:         []int{-3, -2, -1, 0, 1, 2, 3},
			expected:      true,
			expectCompile: true,
		},
		{
			scenario:      "Slice with maximum and minimum values",
			input:         []int{constraints.MinInt, 0, constraints.MaxInt},
			expected:      true,
			expectCompile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {

			t.Log("Input slice:", tt.input)

			if !tt.expectCompile {
				t.Skip("Skipping test case that expects a compile-time error")
			}

			result := isSorted(tt.input)

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			} else {
				t.Logf("Success: expected %v, got %v", tt.expected, result)
			}
		})
	}
}
func generateRandomSlice(length int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = rand.Intn(length)
	}
	return slice
}
func generateSortedSlice(length int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

