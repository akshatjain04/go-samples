package sqrt

import (
	"errors"
	"math"
	"testing"
)






func TestSqrtDecomposition_5409_EQ_5414Update(t *testing.T) {

	testCases := []struct {
		name        string
		sqrtDecomp  *SqrtDecomposition[int, int]
		index       uint64
		newElement  int
		expectedErr error
	}{
		{
			name: "Updating element in middle of the array",
			sqrtDecomp: &SqrtDecomposition{
				elements:  []int{1, 2, 3, 4, 5},
				blocks:    []int{1, 2, 3},
				blockSize: 2,
				updateQ:   mockUpdateQ,
			},
			index:      2,
			newElement: 6,
		},
		{
			name: "Updating element at the beginning of the array",
			sqrtDecomp: &SqrtDecomposition{
				elements:  []int{1, 2, 3, 4, 5},
				blocks:    []int{1, 2, 3},
				blockSize: 2,
				updateQ:   mockUpdateQ,
			},
			index:      0,
			newElement: 6,
		},
		{
			name: "Updating element at the end of the array",
			sqrtDecomp: &SqrtDecomposition{
				elements:  []int{1, 2, 3, 4, 5},
				blocks:    []int{1, 2, 3},
				blockSize: 2,
				updateQ:   mockUpdateQ,
			},
			index:      4,
			newElement: 6,
		},
		{
			name: "Updating with an index out of range",
			sqrtDecomp: &SqrtDecomposition{
				elements:  []int{1, 2, 3, 4, 5},
				blocks:    []int{1, 2, 3},
				blockSize: 2,
				updateQ:   mockUpdateQ,
			},
			index:       10,
			newElement:  6,
			expectedErr: errors.New("index out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectedErr == nil {
						t.Errorf("Test %q - got panic: %v, want no panic", tc.name, r)
					} else if r != tc.expectedErr {
						t.Errorf("Test %q - got panic: %v, want panic: %v", tc.name, r, tc.expectedErr)
					}
				}
			}()
			tc.sqrtDecomp.Update(tc.index, tc.newElement)
			if tc.sqrtDecomp.elements[tc.index] != tc.newElement {
				t.Errorf("Test %q failed - element not updated correctly", tc.name)
			}
			blockIndex := tc.index / tc.sqrtDecomp.blockSize
			if tc.sqrtDecomp.blocks[blockIndex] != tc.newElement {
				t.Errorf("Test %q failed - block not updated correctly", tc.name)
			}
		})
	}
}
func (s *SqrtDecomposition[E, Q]) Update(index uint64, newElement E) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
func mockUpdateQ[E any, Q any](oldBlock Q, oldElement E, newElement E) Q {

	return newElement.(Q)
}
