package sqrt

import (
	"fmt"
	"math"
	"os"
	"sqrt"
	"testing"
)

func TestNewSqrtDecomposition(t *testing.T) {

	tests := []struct {
		name        string
		elements    []int
		queryFunc   func(element int) int
		unionFunc   func(q1 int, q2 int) int
		updateFunc  func(oldQ int, oldE int, newE int) int
		expectedLen int
	}{
		{
			name:     "Initialization with a non-empty slice and valid functions",
			elements: []int{1, 2, 3, 4, 5},
			queryFunc: func(element int) int {
				return element * element
			},
			unionFunc: func(q1 int, q2 int) int {
				return q1 + q2
			},
			updateFunc: func(oldQ int, oldE int, newE int) int {
				return oldQ - oldE*oldE + newE*newE
			},
			expectedLen: 3,
		},
		{
			name:        "Initialization with an empty slice",
			elements:    []int{},
			queryFunc:   func(element int) int { return element },
			unionFunc:   func(q1 int, q2 int) int { return q1 + q2 },
			updateFunc:  func(oldQ int, oldE int, newE int) int { return oldQ - oldE + newE },
			expectedLen: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			decomposition := sqrt.NewSqrtDecomposition(tc.elements, tc.queryFunc, tc.unionFunc, tc.updateFunc)

			w.Close()
			os.Stdout = oldStdout

			var output string
			fmt.Fscanf(r, "%s", &output)

			if len(decomposition.Blocks) != tc.expectedLen {
				t.Errorf("Expected blocks length to be %d, got %d", tc.expectedLen, len(decomposition.Blocks))
			}

			t.Logf("Scenario: %s, Blocks length: %d", tc.name, len(decomposition.Blocks))
		})
	}
}



