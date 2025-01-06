package sqrt

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"sqrt"
)


type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestSqrtDecomposition_Update(t *testing.T) {
	type updateTestCase struct {
		name         string
		index        uint64
		newElement   int
		expectPanic  bool
		expectedElem int
		expectedSum  int
	}

	tests := []updateTestCase{
		{
			name:         "Scenario 1: Successful update of an element within range",
			index:        2,
			newElement:   10,
			expectPanic:  false,
			expectedElem: 10,
			expectedSum:  10,
		},
		{
			name:         "Scenario 2: Update at the boundary of a block",
			index:        0,
			newElement:   5,
			expectPanic:  false,
			expectedElem: 5,
			expectedSum:  5,
		},
		{
			name:        "Scenario 3: Update with an index out of range",
			index:       100,
			newElement:  20,
			expectPanic: true,
		},
		{
			name:         "Scenario 4: Update with minimal possible index (zero)",
			index:        0,
			newElement:   3,
			expectPanic:  false,
			expectedElem: 3,
			expectedSum:  3,
		},
		{
			name:         "Scenario 5: Update with the maximum possible index (end of slice)",
			index:        9,
			newElement:   7,
			expectPanic:  false,
			expectedElem: 7,
			expectedSum:  7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			elements := make([]int, 10)
			for i := range elements {
				elements[i] = int(i)
			}
			blockSize := uint64(math.Ceil(math.Sqrt(float64(len(elements)))))
			s := sqrt.NewSqrtDecomposition[int, int](elements, blockSize, func(a, b int) int {
				return a + b
			})

			if tc.expectPanic {
				assert.Panics(t, func() { s.Update(tc.index, tc.newElement) }, "Update should panic when index is out of range")
			} else {
				assert.NotPanics(t, func() { s.Update(tc.index, tc.newElement) }, "Update should not panic when index is within range")
				require.Equal(t, tc.expectedElem, s.Elements()[tc.index], "Element at index should be updated to the new value")
				assert.Equal(t, tc.expectedSum, s.Blocks()[tc.index/s.BlockSize()], "Block summary should be updated correctly")
			}
		})
	}
}
