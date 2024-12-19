package gcd

import (
	"fmt"
	"os"
	"testing"
)

func TestIterative(t *testing.T) {

	tests := []struct {
		name     string
		a        int64
		b        int64
		expected int64
	}{
		{
			name:     "Coprime integers",
			a:        13,
			b:        17,
			expected: 1,
		},
		{
			name:     "Non-coprime integers",
			a:        12,
			b:        18,
			expected: 6,
		},
		{
			name:     "One integer is zero",
			a:        5,
			b:        0,
			expected: 5,
		},
		{
			name:     "Both integers are zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "Negative integers",
			a:        -8,
			b:        -12,
			expected: 4,
		},
		{
			name:     "One positive and one negative integer",
			a:        9,
			b:        -6,
			expected: 3,
		},
	}

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := Iterative(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Iterative(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			} else {
				t.Logf("Iterative(%d, %d) correctly returned %d", tc.a, tc.b, result)
			}

			w.Close()

			var buf [512]byte
			n, _ := r.Read(buf[:])

			t.Logf("Captured stdout: %s", string(buf[:n]))

			w, _ = os.Pipe()
			os.Stdout = w
		})
	}
}
