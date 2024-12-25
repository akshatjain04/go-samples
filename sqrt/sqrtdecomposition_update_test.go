package sqrt

import (
	"fmt"
	"os"
	"sqrt"
	"testing"
)

type File struct {
	*file // os specific
}

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
func TestSqrtDecomposition_5409_EQ_5414Update(t *testing.T) {

	testCases := []struct {
		name          string
		index         uint64
		newElement    int
		wantPanic     bool
		expectedError string
		expected      []int
	}{
		{
			name:       "Successful update within range",
			index:      1,
			newElement: 10,
			expected:   []int{1, 10, 3, 4},
		},
		{
			name:       "Update at the boundary of a block",
			index:      3,
			newElement: 20,
			expected:   []int{1, 2, 3, 20},
		},
		{
			name:          "Update with an index out of range",
			index:         10,
			newElement:    30,
			wantPanic:     true,
			expectedError: "index out of range",
		},
		{
			name:       "Update with minimal possible index (zero)",
			index:      0,
			newElement: 5,
			expected:   []int{5, 2, 3, 4},
		},
		{
			name:       "Update with the maximum possible index (end of slice)",
			index:      3,
			newElement: 40,
			expected:   []int{1, 2, 3, 40},
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			elements := []int{1, 2, 3, 4}
			blockSize := uint64(math.Sqrt(float64(len(elements))))
			s := sqrt.NewSqrtDecomposition[int, int](blockSize, elements)

			defer func() {

				if r := recover(); r != nil {
					if !tc.wantPanic {
						t.Errorf("Test '%s' failed, unexpected panic: %v", tc.name, r)
					} else if r != tc.expectedError {
						t.Errorf("Test '%s' failed, expected panic with '%s', got: %v", tc.name, tc.expectedError, r)
					}
				}
			}()

			if !tc.wantPanic {
				s.Update(tc.index, tc.newElement)
			} else {

				defer func() { _ = w.Close() }()
				s.Update(tc.index, tc.newElement)
				_ = w.Close()
				var buf []byte
				_, _ = r.Read(buf)
				output := string(buf)
				if output != tc.expectedError {
					t.Errorf("Test '%s' failed, expected output '%s', got: '%s'", tc.name, tc.expectedError, output)
				}
				return
			}

			for i, val := range elements {
				if val != tc.expected[i] {
					t.Errorf("Test '%s' failed, expected element at index %d to be %d, got: %d", tc.name, i, tc.expected[i], val)
				}
			}

			t.Log("Scenario:", tc.name, "executed successfully")
		})
	}

	_ = w.Close()
	os.Stdout = oldStdout
}
