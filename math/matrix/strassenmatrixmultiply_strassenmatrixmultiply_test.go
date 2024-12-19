package matrix

import (
	"errors"
	"reflect"
	"sync"
	"testing"
)

type Matrix[T any] struct {
	elements [][]T
	rows     int
	columns  int
}




func (m Matrix[T]) Add(n Matrix[T]) (Matrix[T], error) {

	return Matrix[T]{}, nil
}
func (m Matrix[T]) Get(row, column int) (T, error) {

	var zero T
	return zero, nil
}
func (m Matrix[T]) Set(row, column int, value T) error {

	return nil
}
func (m Matrix[T]) SubMatrix(row, column, numRows, numColumns int) (Matrix[T], error) {

	return Matrix[T]{}, nil
}
func (m Matrix[T]) Subtract(n Matrix[T]) (Matrix[T], error) {

	return Matrix[T]{}, nil
}
func TestMatrix_561_T_563StrassenMatrixMultiply(t *testing.T) {

	tests := []struct {
		name      string
		matrixA   Matrix[int]
		matrixB   Matrix[int]
		want      Matrix[int]
		wantErr   bool
		errString string
	}{
		{
			name:    "Scenario 1: Successful multiplication of two 1x1 matrices",
			matrixA: New(1, 1, 2),
			matrixB: New(1, 1, 3),
			want:    New(1, 1, 6),
			wantErr: false,
		},
		{
			name:      "Scenario 2: Matrices with incompatible dimensions",
			matrixA:   New(2, 3, 1),
			matrixB:   New(4, 2, 1),
			want:      Matrix[int]{},
			wantErr:   true,
			errString: "dimensions mismatch",
		},
		{
			name:    "Scenario 3: Multiplication of two 2x2 matrices",
			matrixA: New(2, 2, 1),
			matrixB: New(2, 2, 1),
			want:    New(2, 2, 2),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrixA.StrassenMatrixMultiply(tt.matrixB)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.StrassenMatrixMultiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errString != "" && !errors.Is(err, errors.New(tt.errString)) {
				t.Errorf("Matrix.StrassenMatrixMultiply() error = %v, wantErrString %v", err, tt.errString)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.StrassenMatrixMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
