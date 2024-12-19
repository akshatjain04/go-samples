package factorial

import (
	"fmt"
	"math"
	"testing"
)

func TestIterative(t *testing.T) {

	tests := []struct {
		name      string
		input     int
		want      int
		wantError bool
	}{
		{
			name:      "Zero value as input",
			input:     0,
			want:      1,
			wantError: false,
		},
		{
			name:      "Negative value as input",
			input:     -1,
			want:      0,
			wantError: true,
		},
		{
			name:      "Small positive value as input",
			input:     5,
			want:      120,
			wantError: false,
		},
		{
			name:      "Large positive value as input",
			input:     20,
			want:      2432902008176640000,
			wantError: false,
		},
		{
			name:      "Maximum int value as input",
			input:     math.MaxInt64,
			want:      0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := Iterative(tt.input)

			if (got != tt.want) != tt.wantError {
				t.Errorf("Iterative(%d) = %d, want %d", tt.input, got, tt.want)
			} else {
				t.Logf("Success: Iterative(%d) correctly returned %d", tt.input, got)
			}
		})
	}
}
