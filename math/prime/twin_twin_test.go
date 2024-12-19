package prime

import "testing"

func TestTwin(t *testing.T) {

	tests := []struct {
		name         string
		input        int
		expectedNum  int
		expectedBool bool
	}{
		{"Valid twin primes found", 11, 13, true},
		{"No twin primes found", 14, -1, false},
		{"Input is a negative number", -5, -1, false},
		{"Input is zero", 0, -1, false},
		{"Large prime input", 2996863034895, 2996863034897, true},
		{"Input is the first prime number (2)", 2, 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Executing test:", tt.name)
			num, boolVal := Twin(tt.input)
			if num != tt.expectedNum || boolVal != tt.expectedBool {
				t.Errorf("Twin(%d) = (%d, %v), want (%d, %v)", tt.input, num, boolVal, tt.expectedNum, tt.expectedBool)
			} else {
				t.Logf("Success: Expected (%d, %v) and got (%d, %v)", tt.expectedNum, tt.expectedBool, num, boolVal)
			}
		})
	}
}
