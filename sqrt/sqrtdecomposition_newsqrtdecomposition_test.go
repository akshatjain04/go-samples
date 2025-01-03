package sqrt

import (
	"math"
	"testing"
	"github.com/your_project/sqrt"
)

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestNewSqrtDecomposition(t *testing.T) {

	tests := []struct {
		name           string
		elements       []int
		expectedBlocks []int
		expectedError  bool
	}{
		{
			name:           "Initialization with a non-empty slice and valid functions",
			elements:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedBlocks: []int{1, 5, 9},
			expectedError:  false,
		},
		{
			name:           "Initialization with an empty slice",
			elements:       []int{},
			expectedBlocks: []int{},
			expectedError:  false,
		},
	}

	querySingleElement := func(element int) int {

		return element
	}

	unionQ := func(q1 int, q2 int) int {

		return q2
	}

	updateQ := func(oldQ int, oldE int, newE int) int {

		return newE
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			sd := sqrt.NewSqrtDecomposition(tt.elements, querySingleElement, unionQ, updateQ)

			if len(sd.Blocks()) != len(tt.expectedBlocks) {
				t.Errorf("Incorrect number of blocks: got %v, want %v", len(sd.Blocks()), len(tt.expectedBlocks))
			}

			for i, block := range sd.Blocks() {
				if block != tt.expectedBlocks[i] {
					t.Errorf("Incorrect block value at index %d: got %v, want %v", i, block, tt.expectedBlocks[i])
				}
			}
		})
	}
}
