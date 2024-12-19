package binary

import (
	"fmt"
	"os"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Test with Zero",
			input:    0,
			expected: false,
		},
		{
			name:     "Test with Negative Number",
			input:    -4,
			expected: false,
		},
		{
			name:     "Test with Positive Number that is Power of Two",
			input:    16,
			expected: true,
		},
		{
			name:     "Test with Positive Number that is Not Power of Two",
			input:    7,
			expected: false,
		},
		{
			name:     "Test with Maximum Integer Value",
			input:    int(^uint(0) >> 1),
			expected: false,
		},
		{
			name:     "Test with One",
			input:    1,
			expected: true,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				w.Close()
				os.Stdout = oldStdout
			}()

			result := IsPowerOfTwo(tc.input)

			if result != tc.expected {
				t.Errorf("IsPowerOfTwo(%d) = %v; want %v", tc.input, result, tc.expected)
			} else {
				t.Logf("IsPowerOfTwo(%d) = %v; success", tc.input, result)
			}

			w.Close()
			var output string
			_, err := fmt.Fscanf(r, "%s", &output)
			if err != nil {
				t.Errorf("Failed to read output: %v", err)
			}
			t.Log("Captured Output:", output)
			r, w, _ = os.Pipe()
			os.Stdout = w
		})
	}
}
