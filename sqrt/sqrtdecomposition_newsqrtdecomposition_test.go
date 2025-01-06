package sqrt

import (
	"fmt"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"your_project/sqrt"
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
		expectedLen int
	}{
		{
			name:        "Scenario 1: Initialization with a non-empty slice and valid functions",
			elements:    []int{1, 2, 3, 4, 5},
			expectedLen: 3,
		},
		{
			name:        "Scenario 2: Initialization with an empty slice",
			elements:    []int{},
			expectedLen: 0,
		},
		{
			name:        "Scenario 6: Handling of large element slices",
			elements:    make([]int, 10000),
			expectedLen: 100,
		},
	}

	querySingleElement := func(element int) int {
		return element * element
	}
	unionQ := func(q1 int, q2 int) int {
		return q1 + q2
	}
	updateQ := func(oldQ int, oldE int, newE int) int {
		return oldQ - oldE*oldE + newE*newE
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			sd := sqrt.NewSqrtDecomposition(tc.elements, querySingleElement, unionQ, updateQ)

			require.NotNil(t, sd, "SqrtDecomposition should not be nil")
			assert.Equal(t, tc.expectedLen, len(sd.Blocks()), "Blocks length should match the expected value")
			if len(tc.elements) > 0 {
				assert.Equal(t, querySingleElement(tc.elements[0]), sd.Blocks()[0], "First block should match the query of the first element")
			}
		})
	}

}
