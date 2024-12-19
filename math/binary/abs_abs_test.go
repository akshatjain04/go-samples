package binary

import (
	"fmt"
	"os"
	"testing"
)

func TestAbs(t *testing.T) {

	tests := []struct {
		scenario string
		base     int
		n        int
		expected int
	}{
		{
			scenario: "Positive base and n",
			base:     2,
			n:        5,
			expected: 5,
		},
		{
			scenario: "Negative n value",
			base:     2,
			n:        -3,
			expected: 3,
		},
		{
			scenario: "Zero n value",
			base:     2,
			n:        0,
			expected: 0,
		},
		{
			scenario: "Base equals 1",
			base:     1,
			n:        -1,
			expected: 1,
		},
		{
			scenario: "Large values for base and n",
			base:     30,
			n:        1<<29 - 1,
			expected: 1<<29 - 1,
		},
		{
			scenario: "Base greater than the bit size of n",
			base:     33,
			n:        10,
			expected: 10,
		},
		{
			scenario: "Negative base",
			base:     -2,
			n:        5,
			expected: 5,
		},
		{
			scenario: "Base and n both zero",
			base:     0,
			n:        0,
			expected: 0,
		},
	}

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {

			result := Abs(tt.base, tt.n)

			fmt.Fprintf(w, "Result: %d\n", result)

			if result != tt.expected {
				t.Errorf("Abs(%d, %d) = %d; want %d", tt.base, tt.n, result, tt.expected)
			} else {
				t.Logf("Success: Abs(%d, %d) correctly returned %d", tt.base, tt.n, tt.expected)
			}
		})
	}

	w.Close()
	os.Stdout = originalStdout

	var output string
	fmt.Fscanf(r, "%s", &output)
	t.Log("Captured Output:", output)
}
