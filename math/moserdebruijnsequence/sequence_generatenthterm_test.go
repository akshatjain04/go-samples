package moserdebruijnsequence

import "testing"

func TestgenerateNthTerm(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Base Case - First Term",
			input:    0,
			expected: 0,
		},
		{
			name:     "Base Case - Second Term",
			input:    1,
			expected: 1,
		},
		{
			name:     "Even Number Input",
			input:    4,
			expected: 16,
		},
		{
			name:     "Odd Number Input",
			input:    3,
			expected: 5,
		},
		{
			name:     "Large Number Input",
			input:    10,
			expected: 136,
		},

		{
			name:     "Negative Number Input",
			input:    -1,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := generateNthTerm(tc.input)
			if actual != tc.expected {
				t.Errorf("generateNthTerm(%d): expected %d, got %d", tc.input, tc.expected, actual)
			} else {
				t.Logf("generateNthTerm(%d): expected %d, got %d, passed", tc.input, tc.expected, actual)
			}
		})
	}
}
