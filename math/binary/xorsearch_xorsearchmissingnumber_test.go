package binary

import "testing"

func TestXorSearchMissingNumber(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single element array with 0",
			input:    []int{0},
			expected: 1,
		},
		{
			name:     "Single element array with 1",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Large array",
			input:    generateLargeArrayWithMissingNumber(100000, 54321),
			expected: 54321,
		},
		{
			name:     "Consecutive numbers starting from non-zero",
			input:    []int{1, 2, 4},
			expected: 3,
		},
		{
			name:     "Array with duplicate numbers",
			input:    []int{0, 1, 2, 2},
			expected: -1,
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-1, 0, 1},
			expected: -1,
		},
		{
			name:     "Array with non-sequential numbers",
			input:    []int{2, 3, 0, 4},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := XorSearchMissingNumber(tt.input)
			if result != tt.expected {
				t.Errorf("XorSearchMissingNumber(%v) = %v, want %v", tt.input, result, tt.expected)
			} else {
				t.Logf("XorSearchMissingNumber(%v) = %v, as expected", tt.input, result)
			}
		})
	}
}
func generateLargeArrayWithMissingNumber(n, missing int) []int {
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i != missing {
			arr = append(arr, i)
		}
	}
	return arr
}
