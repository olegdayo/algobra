package algobra

// T creates instance of transposed matrix.
// |1 2 3|    |1 4 7|
// |4 5 6| -> |2 5 8|
// |7 8 9|    |3 6 9|
func (m *Matrix[N]) T() (t *Matrix[N], err error) {
	t = NewEmptyMatrix[N](m.ColumnsNumber, m.RowsNumber)
	for i := 0; i < t.RowsNumber; i++ {
		for j := 0; j < t.ColumnsNumber; j++ {
			t.Matr[i][j] = m.Matr[j][i]
		}
	}
	return t, nil
}

// Det gets determinant of a square matrix.
func (m *Matrix[N]) Det() (det N, err error) {
	if !m.IsSquare() {
		return 0, squareMatrixError
	}
	return 0, nil
}

// Adjugate creates an adjugate matrix for given.
func (m *Matrix[N]) Adjugate() (adj *Matrix[N]) {
	return nil
}

// Inverse creates an inverse matrix for given.
func (m *Matrix[N]) Inverse() (inv *Matrix[N], err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}

	det, err := m.Det()
	if err != nil {
		return nil, singularMatrixError
	}

	return m.Adjugate().Divide(det), nil
}
