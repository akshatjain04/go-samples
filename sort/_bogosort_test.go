package github.com/TheAlgorithms/Go/sort

import (
	rand "math/rand"
	testing "testing"
	time "time"
)








/*
ROOST_METHOD_HASH=isSorted_fa5110ab18
ROOST_METHOD_SIG_HASH=isSorted_ef8018aa66

FUNCTION_DEF=func isSorted[T constraints.Number](arr []T) bool 

*/
func TestIsSorted(t *testing.T) {

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
			name:     "Already sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Not sorted slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Slice with duplicate elements",
			input:    []int{1, 2, 2, 3, 4},
			expected: true,
		},
		{
			name:     "Slice with all identical elements",
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Large random slice",
			input:    generateLargeSortedSlice(1000),
			expected: true,
		},
		{
			name:     "Slice with negative numbers",
			input:    []int{-3, -2, -1, 0, 1, 2},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test '%s': %v", tc.name, r)
					t.Fail()
				}
			}()

			result := isSorted(tc.input)
			if result != tc.expected {
				t.Errorf("isSorted(%v) = %v, expected %v", tc.input, result, tc.expected)
			} else {
				t.Logf("Success: isSorted(%v) correctly returned %v", tc.input, result)
			}
		})
	}
}

func generateLargeSortedSlice(size int) []int {
	largeSlice := make([]int, size)
	for i := range largeSlice {
		largeSlice[i] = rand.Intn(size)
	}

	return largeSlice
}

