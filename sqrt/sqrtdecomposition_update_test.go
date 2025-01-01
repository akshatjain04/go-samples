package sqrt

import (
	"fmt"
	"math"
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

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestSqrtDecomposition_5584_EQ_5589Update(t *testing.T) {

	mockUpdateQ := func(oldSummary, oldElement, newElement int) int {
		return int(math.Sqrt(float64(oldSummary - oldElement + newElement)))
	}

	testCases := []struct {
		name                string
		initialElements     []int
		initialSummaries    []int
		blockSize           uint64
		updateIndex         uint64
		newElement          int
		expectedElements    []int
		expectedSummaries   []int
		expectPanicOrError  bool
		panicOrErrorMessage string
	}{
		{
			name:               "Scenario 1: Successful update of an element within range",
			initialElements:    []int{1, 2, 3, 4, 5},
			initialSummaries:   []int{3, 7},
			blockSize:          3,
			updateIndex:        1,
			newElement:         8,
			expectedElements:   []int{1, 8, 3, 4, 5},
			expectedSummaries:  []int{12, 7},
			expectPanicOrError: false,
		},
		{
			name:               "Scenario 2: Update at the boundary of a block",
			initialElements:    []int{1, 2, 3, 4, 5},
			initialSummaries:   []int{3, 7},
			blockSize:          3,
			updateIndex:        3,
			newElement:         10,
			expectedElements:   []int{1, 2, 3, 10, 5},
			expectedSummaries:  []int{3, 13},
			expectPanicOrError: false,
		},
		{
			name:                "Scenario 3: Update with an index out of range",
			initialElements:     []int{1, 2, 3, 4, 5},
			initialSummaries:    []int{3, 7},
			blockSize:           3,
			updateIndex:         10,
			expectPanicOrError:  true,
			panicOrErrorMessage: "index out of range",
		},
		{
			name:               "Scenario 4: Update with minimal possible index (zero)",
			initialElements:    []int{1, 2, 3, 4, 5},
			initialSummaries:   []int{3, 7},
			blockSize:          3,
			updateIndex:        0,
			newElement:         6,
			expectedElements:   []int{6, 2, 3, 4, 5},
			expectedSummaries:  []int{11, 7},
			expectPanicOrError: false,
		},
		{
			name:               "Scenario 5: Update with the maximum possible index (end of slice)",
			initialElements:    []int{1, 2, 3, 4, 5},
			initialSummaries:   []int{3, 7},
			blockSize:          3,
			updateIndex:        4,
			newElement:         9,
			expectedElements:   []int{1, 2, 3, 4, 9},
			expectedSummaries:  []int{3, 11},
			expectPanicOrError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			sqrtDecomp := &SqrtDecomposition[int, int]{
				elements:  tc.initialElements,
				blocks:    tc.initialSummaries,
				blockSize: tc.blockSize,
				updateQ:   mockUpdateQ,
			}

			getCapturedOutput := func() string {
				w.Close()
				os.Stdout = oldStdout
				var buf [1024]byte
				n, _ := r.Read(buf[:])
				return string(buf[:n])
			}

			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanicOrError {
						output := getCapturedOutput()
						if output != tc.panicOrErrorMessage {
							t.Errorf("Expected panic/error message '%s', got '%s'", tc.panicOrErrorMessage, output)
						}
					} else {
						t.Errorf("Unexpected panic: %v", r)
					}
				}
			}()

			if !tc.expectPanicOrError {
				sqrtDecomp.Update(tc.updateIndex, tc.newElement)

				for i, elem := range tc.expectedElements {
					if sqrtDecomp.elements[i] != elem {
						t.Errorf("Element at index %d was not updated correctly, got %d, want %d", i, sqrtDecomp.elements[i], elem)
					}
				}
				for i, sum := range tc.expectedSummaries {
					if sqrtDecomp.blocks[i] != sum {
						t.Errorf("Block summary at index %d was not updated correctly, got %d, want %d", i, sqrtDecomp.blocks[i], sum)
					}
				}
			} else {
				sqrtDecomp.Update(tc.updateIndex, tc.newElement)
				output := getCapturedOutput()
				t.Log(output)
			}
		})
	}
}
