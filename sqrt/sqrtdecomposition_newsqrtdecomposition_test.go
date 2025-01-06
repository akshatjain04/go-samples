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
		name        string
		elements    []int
		querySingle func(int) int
		union       func(int, int) int
		update      func(int, int, int) int
		expectedLen int
	}{
		{
			name:     "Scenario 1: Non-empty slice with valid functions",
			elements: []int{1, 2, 3, 4, 5},
			querySingle: func(e int) int {
				return e * e
			},
			union: func(q1, q2 int) int {
				return q1 + q2
			},
			update: func(oldQ, oldE, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedLen: 3,
		},
		{
			name:        "Scenario 2: Empty slice",
			elements:    []int{},
			querySingle: func(e int) int { return e },
			union:       func(q1, q2 int) int { return q1 + q2 },
			update:      func(oldQ, oldE, newE int) int { return newE },
			expectedLen: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			querySingleElement := tc.querySingle
			unionQ := tc.union
			updateQ := tc.update

			sd := sqrt.NewSqrtDecomposition(tc.elements, querySingleElement, unionQ, updateQ)

			if len(sd.Blocks()) != tc.expectedLen {
				t.Errorf("expected length of blocks to be %v, got %v", tc.expectedLen, len(sd.Blocks()))
			}

		})
	}
}
