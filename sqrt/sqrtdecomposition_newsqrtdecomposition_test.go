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
		name               string
		elements           []int
		querySingleElement func(element int) int
		unionQ             func(q1 int, q2 int) int
		updateQ            func(oldQ int, oldE int, newE int) int
		expectedBlocks     []int
		expectError        bool
	}{

		{
			name:               "Initialization with a non-empty slice and valid functions",
			elements:           []int{1, 2, 3, 4, 5},
			querySingleElement: func(element int) int { return element * element },
			unionQ:             func(q1 int, q2 int) int { return q1 + q2 },
			updateQ:            func(oldQ int, oldE int, newE int) int { return oldQ - oldE*oldE + newE*newE },
			expectedBlocks:     []int{14, 25},
			expectError:        false,
		},

		{
			name:               "Initialization with an empty slice",
			elements:           []int{},
			querySingleElement: func(element int) int { return element },
			unionQ:             func(q1 int, q2 int) int { return q1 + q2 },
			updateQ:            func(oldQ int, oldE int, newE int) int { return oldQ - oldE + newE },
			expectedBlocks:     []int{},
			expectError:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			sqrtDec := sqrt.NewSqrtDecomposition(tt.elements, tt.querySingleElement, tt.unionQ, tt.updateQ)

			w.Close()
			os.Stdout = oldStdout

			var output string
			fmt.Fscanf(r, "%s", &output)

			if sqrtDec == nil && !tt.expectError {
				t.Errorf("NewSqrtDecomposition() returned nil, expected non-nil")
			} else if sqrtDec != nil {
				if len(sqrtDec.Blocks()) != len(tt.expectedBlocks) {
					t.Errorf("NewSqrtDecomposition() returned wrong number of blocks, got %v, want %v", len(sqrtDec.Blocks()), len(tt.expectedBlocks))
				}
				for i, block := range sqrtDec.Blocks() {
					if block != tt.expectedBlocks[i] {
						t.Errorf("Block at index %d is incorrect, got %v, want %v", i, block, tt.expectedBlocks[i])
					}
				}
			}

		})
	}
}



