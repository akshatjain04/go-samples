package sqrt

import (
	"math"
	"sqrt"
	"testing"
)


type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestNewSqrtDecomposition(t *testing.T) {

	testCases := []struct {
		name           string
		elements       []int
		queryFunc      func(int) int
		unionFunc      func(int, int) int
		updateFunc     func(int, int, int) int
		expectedBlocks []int
	}{
		{
			name:     "Initialization with a non-empty slice and valid functions",
			elements: []int{1, 2, 3, 4, 5},
			queryFunc: func(element int) int {
				return element * element
			},
			unionFunc: func(q1, q2 int) int {
				return q1 + q2
			},
			updateFunc: func(oldQ, oldE, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedBlocks: []int{1, 5, 14, 30},
		},
		{
			name:           "Initialization with an empty slice",
			elements:       []int{},
			queryFunc:      func(element int) int { return element },
			unionFunc:      func(q1, q2 int) int { return q1 + q2 },
			updateFunc:     func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			expectedBlocks: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			sd := sqrt.NewSqrtDecomposition(tc.elements, tc.queryFunc, tc.unionFunc, tc.updateFunc)

			if !blocksMatch(sd.Blocks(), tc.expectedBlocks) {
				t.Errorf("Blocks do not match expected result. Got: %v, Want: %v", sd.Blocks(), tc.expectedBlocks)
			}

			if len(tc.elements) == 0 && len(sd.Blocks()) == 0 {
				t.Log("Successfully handled an empty slice of elements.")
			} else if len(tc.elements) > 0 && blocksMatch(sd.Blocks(), tc.expectedBlocks) {
				t.Log("Successfully handled a non-empty slice of elements.")
			} else {
				t.Log("Failed to handle the provided slice of elements.")
			}
		})
	}
}
func blocksMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

