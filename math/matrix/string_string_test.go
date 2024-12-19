package matrix

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMatrix105T107String(t *testing.T) {

	type MatrixElement interface {
	}

	testCases := []struct {
		name           string
		matrixElements [][]MatrixElement
		expectedOutput string
	}{
		{
			name:           "Empty Matrix",
			matrixElements: [][]MatrixElement{},
			expectedOutput: "\n",
		},
		{
			name:           "Single Element Matrix",
			matrixElements: [][]MatrixElement{{1}},
			expectedOutput: "1 \n",
		},
		{
			name:           "Multiple Elements in a Single Row",
			matrixElements: [][]MatrixElement{{1, 2, 3}},
			expectedOutput: "1 2 3 \n",
		},
		{
			name:           "Multiple Rows and Columns",
			matrixElements: [][]MatrixElement{{1, 2}, {3, 4}},
			expectedOutput: "1 2 \n3 4 \n",
		},
		{
			name:           "Non-String Element Types",
			matrixElements: [][]MatrixElement{{"a", "b"}, {true, false}},
			expectedOutput: "a b \ntrue false \n",
		},
		{
			name:           "Large Matrix",
			matrixElements: makeLargeMatrix(1000, 1000),
			expectedOutput: "",
		},
		{
			name:           "Matrix with Different Length Rows",
			matrixElements: [][]MatrixElement{{1, 2}, {3}},
			expectedOutput: "1 2 \n3 \n",
		},
	}

	originalStdout := os.Stdout

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r, w, _ := os.Pipe()
			os.Stdout = w

			m := Matrix[MatrixElement]{elements: tc.matrixElements}

			output := m.String()
			w.Close()
			os.Stdout = originalStdout

			var buf strings.Builder
			fmt.Fscanf(r, "%s", &buf)

			if output != tc.expectedOutput {
				t.Logf("Failed: %s\n", tc.name)
				t.Errorf("Expected output:\n%s\nGot:\n%s\n", tc.expectedOutput, output)
			} else {
				t.Logf("Success: %s\n", tc.name)
			}
		})
	}
}
func makeLargeMatrix(rows, cols int) [][]MatrixElement {

	return [][]MatrixElement{{1, 2}, {3, 4}}
}
