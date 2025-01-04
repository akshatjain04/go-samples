package sqrt

import (
	"math"
	"sqrt"
	"testing"
)

func TestNewSqrtDecomposition(t *testing.T) {

	type testCase[E any, Q any] struct {
		name              string
		elements          []E
		querySingleElem   func(element E) Q
		unionQ            func(q1 Q, q2 Q) Q
		updateQ           func(oldQ Q, oldE E, newE E) (newQ Q)
		expectedNumBlocks int
		expectedBlockSize uint64
	}

	tests := []testCase[int, int]{
		{
			name:              "Initialization with a non-empty slice and valid functions",
			elements:          []int{1, 2, 3, 4, 5},
			querySingleElem:   func(element int) int { return element * element },
			unionQ:            func(q1, q2 int) int { return q1 + q2 },
			updateQ:           func(oldQ, oldE, newE int) int { return oldQ - oldE*oldE + newE*newE },
			expectedNumBlocks: 2,
			expectedBlockSize: uint64(math.Sqrt(5)),
		},
		{
			name:              "Initialization with an empty slice",
			elements:          []int{},
			querySingleElem:   func(element int) int { return element },
			unionQ:            func(q1, q2 int) int { return q1 + q2 },
			updateQ:           func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			expectedNumBlocks: 0,
			expectedBlockSize: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			decomposition := sqrt.NewSqrtDecomposition(tc.elements, tc.querySingleElem, tc.unionQ, tc.updateQ)

			if gotNumBlocks := len(decomposition.Blocks()); int(gotNumBlocks) != tc.expectedNumBlocks {
				t.Errorf("Expected number of blocks: %v, got: %v", tc.expectedNumBlocks, gotNumBlocks)
			}

			if decomposition.BlockSize() != tc.expectedBlockSize {
				t.Errorf("Expected block size: %v, got: %v", tc.expectedBlockSize, decomposition.BlockSize())
			}

			if len(tc.elements) > 0 {

				firstBlockValue := decomposition.Blocks()[0]
				expectedFirstBlockValue := tc.querySingleElem(tc.elements[0])
				for i := 1; i < int(tc.expectedBlockSize) && i < len(tc.elements); i++ {
					expectedFirstBlockValue = tc.unionQ(expectedFirstBlockValue, tc.querySingleElem(tc.elements[i]))
				}
				if firstBlockValue != expectedFirstBlockValue {
					t.Errorf("Expected first block value: %v, got: %v", expectedFirstBlockValue, firstBlockValue)
				}
			}
		})
	}
}


