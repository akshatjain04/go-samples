package sqrt

import (
	"fmt"
	"os"
	"testing"
)







type File struct {
	*file // os specific
}

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestSqrtDecomposition_5584_E_Q_5589Update(t *testing.T) {

	type Element int
	type Query int

	updateQ := func(current Query, oldValue, newValue Element) Query {

		return Query(int(current) - int(oldValue) + int(newValue))
	}

	tests := []struct {
		name          string
		initialArray  []Element
		blockSize     uint64
		index         uint64
		newValue      Element
		expectedArray []Element
		expectedQuery Query
		expectPanic   bool
	}{
		{
			name:          "Scenario 1: Successful update within range",
			initialArray:  []Element{1, 2, 3, 4, 5},
			blockSize:     2,
			index:         2,
			newValue:      10,
			expectedArray: []Element{1, 2, 10, 4, 5},
			expectedQuery: 15,
			expectPanic:   false,
		},
		{
			name:          "Scenario 2: Update at the boundary of a block",
			initialArray:  []Element{1, 2, 3, 4, 5},
			blockSize:     2,
			index:         1,
			newValue:      8,
			expectedArray: []Element{1, 8, 3, 4, 5},
			expectedQuery: 21,
			expectPanic:   false,
		},
		{
			name:          "Scenario 3: Update with an index out of range",
			initialArray:  []Element{1, 2, 3, 4, 5},
			blockSize:     2,
			index:         10,
			newValue:      10,
			expectedArray: nil,
			expectedQuery: 0,
			expectPanic:   true,
		},
		{
			name:          "Scenario 4: Update with minimal possible index (zero)",
			initialArray:  []Element{1, 2, 3, 4, 5},
			blockSize:     2,
			index:         0,
			newValue:      6,
			expectedArray: []Element{6, 2, 3, 4, 5},
			expectedQuery: 20,
			expectPanic:   false,
		},
		{
			name:          "Scenario 5: Update with the maximum possible index (end of slice)",
			initialArray:  []Element{1, 2, 3, 4, 5},
			blockSize:     2,
			index:         4,
			newValue:      9,
			expectedArray: []Element{1, 2, 3, 4, 9},
			expectedQuery: 19,
			expectPanic:   false,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fmt.Fprintf = func(w *os.File, format string, a ...interface{}) (int, error) {
				return fmt.Fprintf(w, format, a...)
			}

			fmt.Fscanf = func(r *os.File, format string, a ...interface{}) (int, error) {
				return fmt.Fscanf(r, format, a...)
			}

			s := SqrtDecomposition[E, Q]{
				elements:  tt.initialArray,
				blocks:    make([]Q, len(tt.initialArray)/int(tt.blockSize)+1),
				blockSize: tt.blockSize,
				updateQ:   updateQ,
			}

			for i, elem := range s.elements {
				blockIdx := uint64(i) / s.blockSize
				s.blocks[blockIdx] = updateQ(s.blocks[blockIdx], 0, elem)
			}

			defer func() {

				if r := recover(); r != nil && !tt.expectPanic {
					t.Errorf("Update panicked unexpectedly: %v", r)
				}
			}()

			if !tt.expectPanic {
				s.Update(tt.index, tt.newValue)
			} else {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Update did not panic when expected")
					}
				}()
				s.Update(tt.index, tt.newValue)
			}

			if !tt.expectPanic {
				for i, elem := range s.elements {
					if elem != tt.expectedArray[i] {
						t.Errorf("Update failed, elements[%d] = %v, want %v", i, elem, tt.expectedArray[i])
					}
				}

				blockIdx := tt.index / s.blockSize
				if s.blocks[blockIdx] != tt.expectedQuery {
					t.Errorf("Update failed, blocks[%d] = %v, want %v", blockIdx, s.blocks[blockIdx], tt.expectedQuery)
				}
			}

			t.Logf("Tested %s, with initial array: %v, updated index: %d, new value: %v, expected array: %v, expected query: %v",
				tt.name, tt.initialArray, tt.index, tt.newValue, tt.expectedArray, tt.expectedQuery)
		})
	}

	w.Close()
	os.Stdout = oldStdout
	fmt.Fprintf = fmt.Fprintf
	fmt.Fscanf = fmt.Fscanf
}
