package fibonacci

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestFormula(t *testing.T) {

	tests := []struct {
		name        string
		input       uint
		want        uint
		wantErr     bool
		errorMargin uint
	}{
		{
			name:    "Zero input value",
			input:   0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "First Fibonacci number",
			input:   1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Large input value",
			input:   20,
			want:    6765,
			wantErr: false,
		},
		{
			name:        "Precision loss for very large input value",
			input:       93,
			want:        12200160415121876738,
			wantErr:     false,
			errorMargin: 1,
		},

		{
			name:    "Numerical stability near potential overflow",
			input:   94,
			want:    0,
			wantErr: true,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Formula(tt.input)
			if tt.wantErr {
				if got != tt.want {
					t.Errorf("Formula() with input %v; got %v, want %v", tt.input, got, tt.want)
				} else {
					t.Logf("Formula() with input %v correctly resulted in overflow", tt.input)
				}
			} else if got < tt.want-tt.errorMargin || got > tt.want+tt.errorMargin {
				t.Errorf("Formula() with input %v; got %v, want within %v of %v", tt.input, got, tt.errorMargin, tt.want)
			} else {
				t.Logf("Formula() with input %v; got %v, as expected", tt.input, got)
			}
		})
	}

	w.Close()
	os.Stdout = oldStdout

}
