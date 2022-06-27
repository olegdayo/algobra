package tests

import (
	"github.com/OFFLUCK/algobra"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	ints := []struct {
		expected *algobra.Matrix[int]
		got      *algobra.Matrix[int]
	}{
		{
			expected: &algobra.Matrix[int]{
				RowsNumber:    2,
				ColumnsNumber: 3,
				Matr: [][]int{
					{0, 1, 0},
					{2, 3, 5},
				},
			},
			got: algobra.NewMatrix[int](
				algobra.Elems[int]{
					{0, 1, 0},
					{2, 3, 5},
				}),
		},
		{
			expected: &algobra.Matrix[int]{
				RowsNumber:    1,
				ColumnsNumber: 1,
				Matr: [][]int{
					{0},
				},
			},
			got: algobra.NewMatrix[int](
				algobra.Elems[int]{
					{0},
				}),
		},
		{
			expected: &algobra.Matrix[int]{
				RowsNumber:    0,
				ColumnsNumber: 0,
				Matr:          [][]int{},
			},
			got: algobra.NewMatrix[int](
				algobra.Elems[int]{},
			),
		},
	}
	for _, test := range ints {
		if !test.expected.Equals(test.got) {
			t.Fail()
		}
	}
	t.Log("Int passed")

	floats := []struct {
		expected *algobra.Matrix[float64]
		got      *algobra.Matrix[float64]
	}{
		{
			expected: &algobra.Matrix[float64]{
				RowsNumber:    2,
				ColumnsNumber: 3,
				Matr: [][]float64{
					{0.0, 0.1, 0.0},
					{0.2, 0.3, 0.5},
				},
			},
			got: algobra.NewMatrix[float64](
				algobra.Elems[float64]{
					{0, 1, 0},
					{2, 3, 5},
				}),
		},
		{
			expected: &algobra.Matrix[float64]{
				RowsNumber:    1,
				ColumnsNumber: 1,
				Matr: [][]float64{
					{0.0},
				},
			},
			got: algobra.NewMatrix[float64](
				algobra.Elems[float64]{
					{0.0},
				}),
		},
		{
			expected: &algobra.Matrix[float64]{
				RowsNumber:    0,
				ColumnsNumber: 0,
				Matr:          [][]float64{},
			},
			got: algobra.NewMatrix[float64](
				algobra.Elems[float64]{},
			),
		},
	}
	for _, test := range floats {
		if !test.expected.Equals(test.got) {
			t.Fail()
		}
	}
	t.Log("Float passed")
}
