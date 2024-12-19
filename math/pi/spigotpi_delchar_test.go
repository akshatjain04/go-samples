package pi

import (
	"fmt"
	"pi"
	"testing"
)

func TestdelChar(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		index     int
		expected  string
		wantPanic bool
	}{
		{
			name:      "Delete middle character",
			input:     "hello",
			index:     2,
			expected:  "helo",
			wantPanic: false,
		},
		{
			name:      "Delete first character",
			input:     "hello",
			index:     0,
			expected:  "ello",
			wantPanic: false,
		},
		{
			name:      "Delete last character",
			input:     "hello",
			index:     4,
			expected:  "hell",
			wantPanic: false,
		},
		{
			name:      "Delete single character",
			input:     "a",
			index:     0,
			expected:  "",
			wantPanic: false,
		},
		{
			name:      "Negative index",
			input:     "hello",
			index:     -1,
			expected:  "",
			wantPanic: true,
		},
		{
			name:      "Out-of-bounds index",
			input:     "hello",
			index:     5,
			expected:  "",
			wantPanic: true,
		},
		{
			name:      "Delete from empty string",
			input:     "",
			index:     0,
			expected:  "",
			wantPanic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("delChar did not panic when expected")
					} else {
						t.Log("delChar panicked as expected")
					}
				}()
			}

			result := pi.delChar(tc.input, tc.index)

			if result != tc.expected {
				t.Errorf("delChar(%q, %d) = %q; want %q", tc.input, tc.index, result, tc.expected)
			} else {
				t.Logf("Success: delChar(%q, %d) correctly yielded %q", tc.input, tc.index, result)
			}
		})
	}
}

