package binary

import (
	"fmt"
	"os"
	"testing"
)

func TestMeanUsingRightShift(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Mean of two positive integers", 6, 8, 7},
		{"Mean of two negative integers", -6, -8, -7},
		{"Mean of a positive and a negative integer", 5, -3, 1},
		{"Mean of large integers", 2147483646, 2147483647, 2147483646},
		{"Mean of zero and another integer", 0, 10, 5},
		{"Mean of two identical integers", 9, 9, 9},
		{"Mean with one of the parameters as the maximum integer value", 2147483647, 100, 1073741873},
		{"Mean with both parameters as the minimum integer value", -2147483648, -2147483648, -2147483648},
	}

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := MeanUsingRightShift(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("MeanUsingRightShift(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			} else {
				t.Logf("Success: %s", tt.name)
			}

			w.Close()
			var buf [128]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			fmt.Fprintf(os.Stdout, "Captured Output: %s\n", output)
		})

		r, w, _ = os.Pipe()
		os.Stdout = w
	}
}
