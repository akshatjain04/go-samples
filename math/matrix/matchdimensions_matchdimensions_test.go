package matrix

import (
	"matrix"
	"testing"
)

func TestMatrix_99_T_101MatchDimensions(t *testing.T) {

	tests := []struct {
		name    string
		matrixA matrix.Matrix[int]
		matrixB matrix.Matrix[int]
		want    bool
	}{
		{
			name:    "Matching dimensions",
			matrixA: matrix.Matrix[int]{Rows: 3, Columns: 3},
			matrixB: matrix.Matrix[int]{Rows: 3, Columns: 3},
			want:    true,
		},
		{
			name:    "Different number of rows",
			matrixA: matrix.Matrix[int]{Rows: 2, Columns: 3},
			matrixB: matrix.Matrix[int]{Rows: 3, Columns: 3},
			want:    false,
		},
		{
			name:    "Different number of columns",
			matrixA: matrix.Matrix[int]{Rows: 3, Columns: 2},
			matrixB: matrix.Matrix[int]{Rows: 3, Columns: 3},
			want:    false,
		},
		{
			name:    "One matrix is empty",
			matrixA: matrix.Matrix[int]{Rows: 0, Columns: 0},
			matrixB: matrix.Matrix[int]{Rows: 3, Columns: 3},
			want:    false,
		},
		{
			name:    "Both matrices are empty",
			matrixA: matrix.Matrix[int]{Rows: 0, Columns: 0},
			matrixB: matrix.Matrix[int]{Rows: 0, Columns: 0},
			want:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.matrixA.MatchDimensions(tc.matrixB)
			if got != tc.want {
				t.Errorf("MatchDimensions() = %v, want %v", got, tc.want)
			} else {
				t.Logf("MatchDimensions() success for test case: %s", tc.name)
			}
		})
	}
}
