package sqrt

import (
	"errors"
	"testing"
)




type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}

func TestUpdate(t *testing.T) {

	testCases := []struct {
		name          string
		elements      []int
		blocks        []int
		blockSize     uint64
		index         uint64
		newElement    int
		expectedError error
	}{
		{
			name:          "Updating element in middle of the array",
			elements:      []int{1, 2, 3, 4, 5},
			blocks:        []int{1, 2, 3, 4, 5},
			blockSize:     5,
			index:         2,
			newElement:    10,
			expectedError: nil,
		},
		{
			name:          "Updating element at the beginning of the array",
			elements:      []int{1, 2, 3, 4, 5},
			blocks:        []int{1, 2, 3, 4, 5},
			blockSize:     5,
			index:         0,
			newElement:    10,
			expectedError: nil,
		},
		{
			name:          "Updating element at the end of the array",
			elements:      []int{1, 2, 3, 4, 5},
			blocks:        []int{1, 2, 3, 4, 5},
			blockSize:     5,
			index:         4,
			newElement:    10,
			expectedError: nil,
		},
		{
			name:          "Updating with an index out of range",
			elements:      []int{1, 2, 3, 4, 5},
			blocks:        []int{1, 2, 3, 4, 5},
			blockSize:     5,
			index:         5,
			newElement:    10,
			expectedError: errors.New("index out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			sqrt := SqrtDecomposition{
				elements:  tc.elements,
				blocks:    tc.blocks,
				blockSize: tc.blockSize,
				updateQ: func(q, e, newE int) int {
					return newE
				},
			}

			sqrt.Update(tc.index, tc.newElement)

			if tc.index < uint64(len(sqrt.elements)) {
				if sqrt.elements[tc.index] != tc.newElement {
					t.Errorf("expected elements[%d] to be %d, got %d", tc.index, tc.newElement, sqrt.elements[tc.index])
				}
				if sqrt.blocks[tc.index/sqrt.blockSize] != tc.newElement {
					t.Errorf("expected blocks[%d] to be %d, got %d", tc.index/sqrt.blockSize, tc.newElement, sqrt.blocks[tc.index/sqrt.blockSize])
				}
			} else {
				t.Errorf("expected error %v, got no error", tc.expectedError)
			}
		})
	}
}
func (s *SqrtDecomposition) Update(index uint64, newElement int) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
func (s *SqrtDecomposition) Update(index uint64, newElement int) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
