package sqrt

import (
	"math"
	"testing"
	"github.com/your/project/sqrt"
)

func TestNewSqrtDecomposition(t *testing.T) {

	testCases := []struct {
		name             string
		elements         []int
		querySingle      func(int) int
		union            func(int, int) int
		update           func(int, int, int) int
		expectedBlocks   []int
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name:     "Initialization with a non-empty slice and valid functions",
			elements: []int{1, 2, 3, 4},
			querySingle: func(e int) int {
				return e * e
			},
			union: func(a, b int) int {
				return a + b
			},
			update: func(oldQ, oldE, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedBlocks: []int{1, 5, 14},
		},
		{
			name:             "Initialization with an empty slice",
			elements:         []int{},
			querySingle:      func(e int) int { return e },
			union:            func(a, b int) int { return a + b },
			update:           func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			expectedBlocks:   []int{},
			expectedError:    false,
			expectedErrorMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sqrtDec := sqrt.NewSqrtDecomposition(tc.elements, tc.querySingle, tc.union, tc.update)

			if len(sqrtDec.Blocks()) != len(tc.expectedBlocks) {
				t.Errorf("Incorrect number of blocks: got %v, want %v", len(sqrtDec.Blocks()), len(tc.expectedBlocks))
			}

			for i, block := range sqrtDec.Blocks() {
				if block != tc.expectedBlocks[i] {
					t.Errorf("Incorrect block value at index %d: got %v, want %v", i, block, tc.expectedBlocks[i])
				}
			}

			expectedBlockSize := uint64(math.Sqrt(float64(len(tc.elements))))
			if sqrtDec.BlockSize() != expectedBlockSize {
				t.Errorf("Incorrect blockSize: got %v, want %v", sqrtDec.BlockSize(), expectedBlockSize)
			}
		})
	}
}

