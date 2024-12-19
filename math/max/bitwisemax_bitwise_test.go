package max

import (
	"fmt"
	"max"
	"os"
	"testing"
)

func TestBitwise(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		base     int
		expected int
	}{
		{
			name:     "Scenario 1: Equal positive numbers",
			a:        5,
			b:        5,
			base:     3,
			expected: 5,
		},
		{
			name:     "Scenario 2: Positive `a`, negative `b`",
			a:        10,
			b:        -5,
			base:     3,
			expected: 10,
		},
		{
			name:     "Scenario 3: Negative `a`, positive `b`",
			a:        -5,
			b:        10,
			base:     3,
			expected: 10,
		},
		{
			name:     "Scenario 4: Large numbers",
			a:        2147483647,
			b:        2147483646,
			base:     31,
			expected: 2147483647,
		},
		{
			name:     "Scenario 5: Zero values",
			a:        0,
			b:        10,
			base:     3,
			expected: 10,
		},
		{
			name:     "Scenario 6: Base value edge cases",
			a:        15,
			b:        10,
			base:     0,
			expected: 15,
		},
		{
			name:     "Scenario 7: Negative base value",
			a:        10,
			b:        20,
			base:     -1,
			expected: 20,
		},
		{
			name:     "Scenario 8: Base value equal to integer bit size",
			a:        5,
			b:        3,
			base:     31,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := max.Bitwise(tt.a, tt.b, tt.base)
			if result != tt.expected {
				t.Errorf("Bitwise(%d, %d, %d) = %d; expected %d", tt.a, tt.b, tt.base, result, tt.expected)
			} else {
				t.Logf("Success: Bitwise(%d, %d, %d) correctly returned %d", tt.a, tt.b, tt.base, result)
			}
		})
	}
}

