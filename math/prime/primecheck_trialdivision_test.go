package prime

import (
	"fmt"
	"testing"
)

func TestTrialDivision(t *testing.T) {

	tests := []struct {
		name     string
		input    int64
		expected bool
	}{
		{
			name:     "Negative input",
			input:    -1,
			expected: false,
		},
		{
			name:     "Zero input",
			input:    0,
			expected: false,
		},
		{
			name:     "Input is 2, smallest prime",
			input:    2,
			expected: true,
		},
		{
			name:     "Small prime number",
			input:    3,
			expected: true,
		},
		{
			name:     "Small composite number",
			input:    4,
			expected: false,
		},
		{
			name:     "Another prime number",
			input:    5,
			expected: true,
		},
		{
			name:     "Another composite number",
			input:    6,
			expected: false,
		},
		{
			name:     "Large prime number",
			input:    9223372036854775783,
			expected: true,
		},
		{
			name:     "Largest int64",
			input:    9223372036854775807,
			expected: false,
		},
		{
			name:     "Even number greater than 2",
			input:    10,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := TrialDivision(tc.input)

			if result != tc.expected {
				t.Errorf("TrialDivision(%d) = %v, expected %v", tc.input, result, tc.expected)
			} else {
				t.Logf("TrialDivision(%d) = %v, as expected", tc.input, result)
			}
		})
	}
}
