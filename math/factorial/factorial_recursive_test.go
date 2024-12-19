package factorial

import "testing"

func TestRecursive(t *testing.T) {

	testCases := []struct {
		name      string
		input     int
		want      int
		expectErr bool
	}{
		{"Basic functionality with n=5", 5, 120, false},
		{"Input value of n=0", 0, 1, false},
		{"Negative input value", -1, 0, true},
		{"Large input value", 100, -1, false},
		{"Input value of n=1", 1, 1, false},
		{"Verifying recursion correctness with known values", 4, 24, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := Recursive(tc.input)

			if tc.expectErr {
				if got != 0 {
					t.Errorf("Expected an error or 0 for input %d, got %d", tc.input, got)
				}
			} else {
				if got != tc.want {
					t.Errorf("For input %d, expected %d, got %d", tc.input, tc.want, got)
				}
			}

			if tc.expectErr {
				t.Logf("For input %d, correctly handled invalid input", tc.input)
			} else {
				t.Logf("For input %d, result was correctly computed as %d", tc.input, got)
			}
		})
	}

}
