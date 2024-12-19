package math

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestAbs(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Positive Number",
			input:    42,
			expected: 42,
		},
		{
			name:     "Negative Number",
			input:    -42,
			expected: 42,
		},
		{
			name:     "Zero Value",
			input:    0,
			expected: 0,
		},
		{
			name:     "Maximum Integer Value",
			input:    math.MaxInt32,
			expected: math.MaxInt32,
		},
		{
			name:     "Minimum Integer Value",
			input:    math.MinInt32,
			expected: math.MaxInt32 + 1,
		},
	}

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := Abs(tc.input)

			if result != tc.expected {
				t.Errorf("Abs(%d) = %d; want %d", tc.input, result, tc.expected)
			} else {
				t.Logf("Success: Abs(%d) correctly returned %d", tc.input, result)
			}
		})
	}

	w.Close()

	var buf [512]byte
	n, _ := r.Read(buf[:])
	output := string(buf[:n])

	fmt.Fprintf(originalStdout, "Captured output: %s\n", output)

	os.Stdout = originalStdout
}
