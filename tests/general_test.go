package tests

import (
	"github.com/OFFLUCK/algobra"
	"testing"
)

const deviation int = 10

func TestAt(t *testing.T) {
	matrix := algobra.NewMatrix[int](
		algobra.Elems[int]{
			{0, 1, 0},
			{2, 3, 5},
		},
	)

	for i := 0; i < matrix.RowsNumber; i++ {
		for j := 0; j < matrix.ColumnsNumber; j++ {
			elem, err := matrix.At(i, j)
			if err != nil {
				t.Errorf("Index error: i: %d, j: %d", i, j)
			}
			if elem != matrix.Matr[i][j] {
				t.Fail()
			}
		}
	}

	for i := -deviation; i < 0; i++ {
		for j := -deviation; j < 0; j++ {
			_, err := matrix.At(i, j)
			if err == nil {
				t.Fail()
			}
		}
	}

	for i := matrix.RowsNumber; i < matrix.RowsNumber+deviation; i++ {
		for j := matrix.ColumnsNumber; j < matrix.ColumnsNumber+deviation; j++ {
			_, err := matrix.At(i, j)
			if err == nil {
				t.Fail()
			}
		}
	}
}
