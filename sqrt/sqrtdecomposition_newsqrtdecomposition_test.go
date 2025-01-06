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
		name         string
		elements     []int
		query        func(int) int
		union        func(int, int) int
		update       func(int, int, int) int
		expectedSize int
	}{
		{
			name:     "Scenario 1: Non-empty slice with valid functions",
			elements: []int{1, 2, 3, 4},
			query: func(e int) int {
				return e * e
			},
			union: func(a, b int) int {
				return a + b
			},
			update: func(oldQ, oldE, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedSize: 2,
		},
		{
			name:         "Scenario 2: Empty slice",
			elements:     []int{},
			query:        func(e int) int { return e * e },
			union:        func(a, b int) int { return a + b },
			update:       func(oldQ, oldE, newE int) int { return oldQ - oldE*oldE + newE*newE },
			expectedSize: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sqrt.NewSqrtDecomposition(tc.elements, tc.query, tc.union, tc.update)
			if len(result.Blocks()) != tc.expectedSize {
				t.Errorf("expected %d blocks, got %d", tc.expectedSize, len(result.Blocks()))
			}

			t.Log("Test case passed:", tc.name)
		})
	}
}
