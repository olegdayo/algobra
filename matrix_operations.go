package algobra

func (m *Matrix[N]) T() (t *Matrix[N], err error) {
	t = NewEmptyMatrix[N](m.ColumnsNumber, m.RowsNumber)
	for i := 0; i < t.RowsNumber; i++ {
		for j := 0; j < t.ColumnsNumber; j++ {
			t.Matr[i][j] = m.Matr[j][i]
		}
	}
	return t, nil
}

func (m *Matrix[N]) Det() (det N, err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}
	return nil, nil
}

func (m *Matrix[N]) Adjugate() (adj *Matrix[N]) {
	return nil
}

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
