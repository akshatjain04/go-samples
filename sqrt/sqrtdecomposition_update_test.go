package sqrt

import (
	"fmt"
	"math"
	"os"
	"sqrt"
	"testing"
)

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestSqrtDecomposition_5584_EQ_5589Update(t *testing.T) {

	type E = int
	type Q = int

	tests := []struct {
		name             string
		initialElements  []E
		blockSize        uint64
		updateIndex      uint64
		newValue         E
		expectedPanic    bool
		expectedElements []E
		expectedBlocks   []Q
	}{
		{
			name:             "Scenario 1: Successful update within range",
			initialElements:  []E{1, 2, 3, 4, 5},
			blockSize:        2,
			updateIndex:      2,
			newValue:         6,
			expectedElements: []E{1, 2, 6, 4, 5},
			expectedBlocks:   []Q{3, 10, 5},
		},
		{
			name:             "Scenario 2: Update at the boundary of a block",
			initialElements:  []E{1, 2, 3, 4, 5},
			blockSize:        2,
			updateIndex:      2,
			newValue:         7,
			expectedElements: []E{1, 2, 7, 4, 5},
			expectedBlocks:   []Q{3, 11, 5},
		},
		{
			name:             "Scenario 3: Update with index out of range",
			initialElements:  []E{1, 2, 3},
			blockSize:        2,
			updateIndex:      5,
			newValue:         9,
			expectedPanic:    true,
			expectedElements: []E{1, 2, 3},
			expectedBlocks:   []Q{3, 3},
		},
		{
			name:             "Scenario 4: Update with minimal possible index (zero)",
			initialElements:  []E{1, 2, 3, 4, 5},
			blockSize:        2,
			updateIndex:      0,
			newValue:         8,
			expectedElements: []E{8, 2, 3, 4, 5},
			expectedBlocks:   []Q{10, 7, 5},
		},
		{
			name:             "Scenario 5: Update with the maximum possible index (end of slice)",
			initialElements:  []E{1, 2, 3, 4, 5},
			blockSize:        2,
			updateIndex:      4,
			newValue:         10,
			expectedElements: []E{1, 2, 3, 4, 10},
			expectedBlocks:   []Q{3, 7, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			sqrtDecomp := sqrt.NewSqrtDecomposition[E, Q](tc.initialElements, tc.blockSize, math.Sum)

			if tc.expectedPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Update did not panic on out of range index")
					} else {
						t.Log("Update function correctly panicked on out of range index")
					}
				}()
			}
			sqrtDecomp.Update(tc.updateIndex, tc.newValue)

			if !tc.expectedPanic {

				for i, v := range tc.expectedElements {
					if sqrtDecomp.Elements()[i] != v {
						t.Errorf("Element at index %d is incorrect, got: %v, want: %v", i, sqrtDecomp.Elements()[i], v)
					}
				}

				for i, v := range tc.expectedBlocks {
					if sqrtDecomp.Blocks()[i] != v {
						t.Errorf("Block summary at index %d is incorrect, got: %v, want: %v", i, sqrtDecomp.Blocks()[i], v)
					}
				}
			}
		})
	}
}
