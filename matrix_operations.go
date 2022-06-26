package algobra

func (m *Matrix[N]) T() (ans *Matrix[N], err error) {
	ans = NewEmptyMatrix[N](m.ColumnsNumber, m.RowsNumber)
	for i := 0; i < ans.RowsNumber; i++ {
		for j := 0; j < ans.ColumnsNumber; j++ {
			ans.Matr[i][j] = m.Matr[j][i]
		}
	}
	return ans, nil
}

func (m *Matrix[N]) Det() (num N, err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}
	return nil, nil
}
