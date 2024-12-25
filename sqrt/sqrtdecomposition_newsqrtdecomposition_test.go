package sqrt

import (
	"math"
	"testing"
	"sqrt"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestNewSqrtDecomposition(t *testing.T) {

	type testCase struct {
		name        string
		elements    []int
		querySingle func(int) int
		unionQ      func(int, int) int
		updateQ     func(int, int, int) int
		expectedLen int
	}

	testCases := []testCase{
		{
			name:     "Scenario 1: Initialization with a non-empty slice and valid functions",
			elements: []int{1, 2, 3, 4, 5},
			querySingle: func(x int) int {
				return x * x
			},
			unionQ: func(a, b int) int {
				return a + b
			},
			updateQ: func(oldQ, oldE, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedLen: int(math.Ceil(math.Sqrt(5))),
		},
		{
			name:        "Scenario 2: Initialization with an empty slice",
			elements:    []int{},
			querySingle: func(x int) int { return x * x },
			unionQ:      func(a, b int) int { return a + b },
			updateQ:     func(oldQ, oldE, newE int) int { return oldQ - oldE*oldE + newE*newE },
			expectedLen: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			decomposition := sqrt.NewSqrtDecomposition(tc.elements, tc.querySingle, tc.unionQ, tc.updateQ)

			if len(decomposition.Blocks()) != tc.expectedLen {
				t.Errorf("expected blocks length %d, got %d", tc.expectedLen, len(decomposition.Blocks()))
			}

			t.Log("Success:", tc.name)
		})
	}
}

