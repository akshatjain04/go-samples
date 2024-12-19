package min

import (
	"fmt"
	"os"
	"testing"
)

func TestBitwise(t *testing.T) {

	tests := []struct {
		name      string
		base      int
		value     int
		values    []int
		want      int
		expectErr bool
	}{
		{
			name:   "Scenario 1: Single value input",
			base:   0,
			value:  42,
			values: []int{},
			want:   42,
		},
		{
			name:   "Scenario 2: All values are equal",
			base:   1,
			value:  10,
			values: []int{10, 10, 10},
			want:   10,
		},
		{
			name:   "Scenario 3: Negative values",
			base:   2,
			value:  -5,
			values: []int{-10, -3},
			want:   -10,
		},
		{
			name:   "Scenario 4: Large numbers",
			base:   2,
			value:  1000000,
			values: []int{1000000000, 2000000000},
			want:   1000000,
		},
		{
			name:   "Scenario 5: Zero as one of the values",
			base:   1,
			value:  0,
			values: []int{5, 10},
			want:   0,
		},
		{
			name:   "Scenario 6: Base is zero",
			base:   0,
			value:  10,
			values: []int{5, 15},
			want:   5,
		},
		{
			name:      "Scenario 7: Invalid base value (negative or too large)",
			base:      -1,
			value:     10,
			values:    []int{5, 15},
			want:      0,
			expectErr: true,
		},
		{
			name:   "Scenario 8: No additional values provided",
			base:   1,
			value:  10,
			values: []int{},
			want:   10,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := Bitwise(tt.base, tt.value, tt.values...)

			if got != tt.want {
				t.Errorf("Bitwise() = %v, want %v", got, tt.want)
			}

			t.Logf("Scenario: %s\n", tt.name)
			t.Logf("Base: %d, Value: %d, Values: %v\n", tt.base, tt.value, tt.values)
			t.Logf("Got: %d, Want: %d\n", got, tt.want)

			if tt.expectErr {

				t.Log("Expected an error or unexpected result due to invalid base value")
			}
		})
	}

	w.Close()
	os.Stdout = oldStdout
	output := make([]byte, 1024)
	n, _ := r.Read(output)
	fmt.Fprintf(os.Stdout, "%s", output[:n])
}
