package matrix

import (
	"testing"
	"matrix"
)

func TestMatrix_215_T_217CheckEqual(t *testing.T) {

	testCases := []struct {
		name     string
		matrix1  matrix.Matrix[int]
		matrix2  matrix.Matrix[int]
		expected bool
	}{
		{
			name:     "Identical elements",
			matrix1:  matrix.NewMatrix([][]int{{1, 2}, {3, 4}}),
			matrix2:  matrix.NewMatrix([][]int{{1, 2}, {3, 4}}),
			expected: true,
		},
		{
			name:     "Different elements",
			matrix1:  matrix.NewMatrix([][]int{{1, 2}, {3, 4}}),
			matrix2:  matrix.NewMatrix([][]int{{5, 6}, {7, 8}}),
			expected: false,
		},
		{
			name:     "Different dimensions",
			matrix1:  matrix.NewMatrix([][]int{{1, 2}}),
			matrix2:  matrix.NewMatrix([][]int{{1, 2}, {3, 4}}),
			expected: false,
		},
		{
			name:     "Empty matrices",
			matrix1:  matrix.NewMatrix([][]int{}),
			matrix2:  matrix.NewMatrix([][]int{}),
			expected: true,
		},
		{
			name:     "Identical dimensions but one value differs",
			matrix1:  matrix.NewMatrix([][]int{{1, 2}, {3, 4}}),
			matrix2:  matrix.NewMatrix([][]int{{1, 2}, {3, 5}}),
			expected: false,
		},
		{
			name:     "Large matrices with identical elements",
			matrix1:  matrix.NewMatrix(generateLargeIdenticalMatrix(1000)),
			matrix2:  matrix.NewMatrix(generateLargeIdenticalMatrix(1000)),
			expected: true,
		},
		{
			name:     "Large matrices with one value differs",
			matrix1:  matrix.NewMatrix(generateLargeIdenticalMatrix(1000)),
			matrix2:  matrix.NewMatrix(generateLargeMatrixWithOneDifference(1000)),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.matrix1.CheckEqual(tc.matrix2)
			if result != tc.expected {
				t.Errorf("CheckEqual() test failed for '%s': expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("CheckEqual() test passed for '%s'", tc.name)
			}
		})
	}
}

