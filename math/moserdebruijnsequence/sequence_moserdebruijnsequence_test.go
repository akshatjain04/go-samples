package moserdebruijnsequence

import (
	"fmt"
	"os"
	"testing"
	"github.com/TheAlgorithms/Go/math/moserdebruijnsequence"
)

func TestMoserDeBruijnSequence(t *testing.T) {

	tests := []struct {
		name           string
		input          int
		expectedOutput []int
		expectError    bool
	}{
		{
			name:           "Zero as Input",
			input:          0,
			expectedOutput: []int{},
			expectError:    false,
		},
		{
			name:           "Single Element Sequence",
			input:          1,
			expectedOutput: []int{0},
			expectError:    false,
		},
		{
			name:           "Typical Input",
			input:          4,
			expectedOutput: []int{0, 1, 4, 5},
			expectError:    false,
		},
		{
			name:           "Negative Input",
			input:          -1,
			expectedOutput: []int{},
			expectError:    true,
		},
		{
			name:           "Large Input Value",
			input:          1000,
			expectedOutput: nil,
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			output := moserdebruijnsequence.MoserDeBruijnSequence(tc.input)

			if len(output) != len(tc.expectedOutput) && !tc.expectError {
				t.Errorf("Test %s failed: expected length %d, got %d", tc.name, len(tc.expectedOutput), len(output))
			}

			for i := range tc.expectedOutput {
				if output[i] != tc.expectedOutput[i] {
					t.Errorf("Test %s failed: at index %d, expected %d, got %d", tc.name, i, tc.expectedOutput[i], output[i])
				}
			}

			if tc.expectError && len(output) != 0 {
				t.Errorf("Test %s failed: expected error or empty output, got %v", tc.name, output)
			}

			t.Logf("Test %s passed", tc.name)
		})
	}
}
