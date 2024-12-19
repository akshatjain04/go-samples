package binary

import "testing"

func TestReverseBits(t *testing.T) {

	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{
			name:     "All bits set to 1",
			input:    0xFFFFFFFF,
			expected: 0xFFFFFFFF,
		},
		{
			name:     "Alternating bit pattern",
			input:    0xAAAAAAAA,
			expected: 0x55555555,
		},
		{
			name:     "Zero value input",
			input:    0x0,
			expected: 0x0,
		},
		{
			name:     "Single set bit",
			input:    0x2,
			expected: 0x40000000,
		},
		{
			name:     "Palindrome bit pattern",
			input:    0x80000001,
			expected: 0x80000001,
		},
		{
			name:     "Maximum unsigned integer value",
			input:    0xFFFFFFFF,
			expected: 0xFFFFFFFF,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Executing test case:", tc.name)

			result := ReverseBits(tc.input)

			if result != tc.expected {
				t.Errorf("ReverseBits(%v) = %v, want %v", tc.input, result, tc.expected)
			} else {
				t.Logf("Success: Expected and actual values match. (%v)", tc.expected)
			}
		})
	}
}
