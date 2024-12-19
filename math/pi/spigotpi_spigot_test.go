package pi

import (
	"testing"
	"github.com/TheAlgorithms/Go/math/pi"
)

func TestSpigot(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		want        string
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "Zero digits of Pi",
			input:   0,
			want:    "",
			wantErr: false,
		},
		{
			name:    "Negative digits of Pi",
			input:   -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Input requiring rounding up during carry-over",
			input:   5,
			want:    "31416",
			wantErr: false,
		},
		{
			name:    "Large number of digits of Pi",
			input:   10000,
			want:    "",
			wantErr: false,
		},
		{
			name:    "First digit of Pi",
			input:   1,
			want:    "3",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pi.Spigot(tt.input)
			if (got != tt.want) != tt.wantErr {
				t.Errorf("Spigot(%d) = %v, want %v", tt.input, got, tt.want)
			} else {
				t.Logf("Success: %s", tt.name)
			}
		})
	}
}
