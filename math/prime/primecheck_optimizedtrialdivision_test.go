package prime

import (
	"fmt"
	"os"
	"testing"
)

func TestOptimizedTrialDivision(t *testing.T) {

	testCases := []struct {
		name  string
		input int64
		want  bool
	}{
		{
			name:  "Negative number",
			input: -5,
			want:  false,
		},
		{
			name:  "Zero",
			input: 0,
			want:  false,
		},
		{
			name:  "One",
			input: 1,
			want:  false,
		},
		{
			name:  "Small prime number",
			input: 3,
			want:  true,
		},
		{
			name:  "Small non-prime number",
			input: 4,
			want:  false,
		},
		{
			name:  "Large prime number",
			input: 7919,
			want:  true,
		},
		{
			name:  "Large non-prime number",
			input: 8000,
			want:  false,
		},
		{
			name:  "Prime square",
			input: 9,
			want:  false,
		},
		{
			name:  "Very large number",
			input: 9223372036854775783,
			want:  false,
		},
	}

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r, w, _ := os.Pipe()
			os.Stdout = w

			got := OptimizedTrialDivision(tc.input)

			w.Close()

			var buf string
			fmt.Fscanf(r, "%s", &buf)

			if got != tc.want {
				t.Errorf("OptimizedTrialDivision(%v) = %v, want %v", tc.input, got, tc.want)
			} else {
				t.Logf("OptimizedTrialDivision(%v) = %v, as expected", tc.input, got)
			}
		})
	}
}
