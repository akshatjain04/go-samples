package math

import (
	"fmt"
	"math"
	"os"
	"testing"
)

// TestIsPowOfTwoUseLog tests the IsPowOfTwoUseLog function using table-driven tests.
func TestIsPowOfTwoUseLog(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name   string
		number float64
		want   bool
	}{
		{"Power of two", 16, true},
		{"Not a power of two", 10, false},
		{"Zero", 0, false},
		{"Max float64", math.MaxFloat64, false},
		{"Negative power of two", -8, false},
		{"Very small non-zero", math.SmallestNonzeroFloat64, false},
		{"Non-integer close to power of two", 15.99, false},
		{"Integer close to but not power of two", 31, false},
	}

	// Capture the original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := IsPowOfTwoUseLog(tt.number)

			// Assert
			if got != tt.want {
				t.Errorf("IsPowOfTwoUseLog(%v) = %v, want %v", tt.number, got, tt.want)
			} else {
				t.Logf("Success: IsPowOfTwoUseLog(%v) correctly returned %v", tt.number, tt.want)
			}

			// Capture the output
			w.Close()
			var buf [512]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			// Log the output for diagnostic purposes
			t.Log("Captured Output:", output)

			// Reset stdout
			os.Stdout = originalStdout

			// Use fmt.Fscanf() and fmt.Fprintf() if needed for input/output operations
			// TODO: Implement the input/output operations if the function under test requires it.
		})
	}
}

// Note: The function IsPowOfTwoUseLog is expected to be imported from the "math" package.
// If any additional functions are needed, ensure they are imported properly.
