package prime

import (
	"fmt"
	"os"
	"testing"
)

func TestFactorize(t *testing.T) {

	originalStdout := os.Stdout

	tests := []struct {
		name          string
		input         int64
		expected      map[int64]int64
		expectError   bool
		errorContains string
	}{
		{
			name:     "Factorize a prime number",
			input:    11,
			expected: map[int64]int64{11: 1},
		},
		{
			name:     "Factorize a square of a prime number",
			input:    49,
			expected: map[int64]int64{7: 2},
		},
		{
			name:          "Factorize zero",
			input:         0,
			expected:      map[int64]int64{},
			expectError:   true,
			errorContains: "not factorizable",
		},
		{
			name:          "Factorize a negative number",
			input:         -17,
			expected:      map[int64]int64{},
			expectError:   true,
			errorContains: "negative number",
		},
		{
			name:     "Factorize a large composite number",
			input:    123456,
			expected: map[int64]int64{2: 6, 3: 1, 643: 1},
		},
		{
			name:     "Factorize the number one",
			input:    1,
			expected: map[int64]int64{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			r, w, _ := os.Pipe()
			os.Stdout = w

			result := Factorize(tc.input)

			w.Close()
			os.Stdout = originalStdout

			var output string
			fmt.Fscanf(r, "%s", &output)

			if !mapsEqual(result, tc.expected) {
				t.Errorf("Factorize(%d) got %v, want %v", tc.input, result, tc.expected)
			}

			if tc.expectError && !containsError(output, tc.errorContains) {
				t.Errorf("Expected an error containing '%s' but got '%s'", tc.errorContains, output)
			}

			t.Log(tc.name)
		})
	}
}
func containsError(output, errorContains string) bool {
	return errorContains != "" && output != "" && strings.Contains(output, errorContains)
}
func mapsEqual(a, b map[int64]int64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}
