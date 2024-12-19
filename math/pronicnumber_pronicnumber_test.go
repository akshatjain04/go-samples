package math_test

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPronicNumber tests the PronicNumber function for various scenarios.
func TestPronicNumber(t *testing.T) {
	// Redirecting output to capture it during tests
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Table-driven tests
	testCases := []struct {
		description string
		input       int
		expected    bool
	}{
		{
			description: "known pronic number",
			input:       12, // 3 * 4
			expected:    true,
		},
		{
			description: "non-pronic number",
			input:       10,
			expected:    false,
		},
		{
			description: "zero",
			input:       0, // 0 * 1
			expected:    true,
		},
		{
			description: "negative number",
			input:       -4,
			expected:    false,
		},
		{
			description: "maximum integer value",
			input:       math.MaxInt,
			expected:    false, // TODO: Adjust expected result based on PronicNumber implementation for MaxInt
		},
		{
			description: "odd number",
			input:       7,
			expected:    false,
		},
		{
			description: "perfect square not pronic",
			input:       16, // 4 * 4
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := math.PronicNumber(tc.input)
			assert.Equal(t, tc.expected, result, "PronicNumber(%d) = %v, want %v", tc.input, result, tc.expected)
			if result == tc.expected {
				t.Logf("Success: %s - PronicNumber(%d) = %v", tc.description, tc.input, result)
			} else {
				t.Errorf("Failure: %s - PronicNumber(%d) = %v, want %v", tc.description, tc.input, result, tc.expected)
			}
		})
	}

	// Resetting Stdout
	w.Close()
	os.Stdout = oldStdout
	var buf []byte
	fmt.Fscanf(r, "%s", &buf)
	t.Log("Captured Output:", string(buf))

	// TODO: Additional tests can be written for concurrency if the PronicNumber function is used in a concurrent context.
}
