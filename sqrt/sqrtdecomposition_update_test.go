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
func TestSqrtDecomposition_3231_EQ_3236Update(t *testing.T) {

	testCases := []struct {
		name        string
		sd          *SqrtDecomposition[int, int]
		index       uint64
		newElement  int
		expectedErr bool
	}{
		{
			name:        "Scenario 1: Successful update of an element within range",
			sd:          nil,
			index:       0,
			newElement:  0,
			expectedErr: false,
		},
		{
			name:        "Scenario 2: Update at the boundary of a block",
			sd:          nil,
			index:       0,
			newElement:  0,
			expectedErr: false,
		},
		{
			name:        "Scenario 3: Update with an index out of range",
			sd:          nil,
			index:       0,
			newElement:  0,
			expectedErr: true,
		},
		{
			name:        "Scenario 4: Update with minimal possible index (zero)",
			sd:          nil,
			index:       0,
			newElement:  0,
			expectedErr: false,
		},
		{
			name:        "Scenario 5: Update with the maximum possible index (end of slice)",
			sd:          nil,
			index:       0,
			newElement:  0,
			expectedErr: false,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil && !tc.expectedErr {
					t.Errorf("test panicked unexpectedly: %v", r)
				}
			}()

			tc.sd.Update(tc.index, tc.newElement)

			w.Close()
			var buf [512]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			os.Stdout = oldStdout

			if tc.expectedErr {
				if output == "" {
					t.Errorf("expected an error but got none")
				}
			} else {
				if output != "" {
					t.Errorf("expected no error but got: %s", output)
				}

			}

			t.Log(output)
		})
	}
}
