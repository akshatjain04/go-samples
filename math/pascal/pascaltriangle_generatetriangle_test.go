package pascal

import (
	"fmt"
	"os"
	"testing"
	"time"
	"pascal"
)

func TestGenerateTriangle(t *testing.T) {

	var tests = []struct {
		name          string
		input         int
		expectedError bool
	}{
		{"NegativeLineNumber", -1, true},
		{"LargeLineNumber", 100, false},
		{"PotentialOverflow", 30, false},

		{"ExtremelyLargeInput", 1000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError {

				old := os.Stdout
				defer func() { os.Stdout = old }()
				_, w, _ := os.Pipe()
				os.Stdout = w

				start := time.Now()
				result := pascal.GenerateTriangle(tt.input)
				duration := time.Since(start)

				if len(result) != 0 {
					t.Errorf("GenerateTriangle(%d) expected an empty slice, got %v", tt.input, result)
				}
				t.Log("Duration:", duration)
			} else {
				start := time.Now()
				result := pascal.GenerateTriangle(tt.input)
				duration := time.Since(start)

				if len(result) != tt.input {
					t.Errorf("GenerateTriangle(%d) expected a slice with length %d, got length %d", tt.input, tt.input, len(result))
				}
				if tt.input > 0 && (len(result[0]) != 1 || len(result[tt.input-1]) != tt.input) {
					t.Errorf("GenerateTriangle(%d) expected proper Pascal's triangle dimensions, got incorrect dimensions", tt.input)
				}
				t.Log("Duration:", duration)

			}
		})
	}

}

