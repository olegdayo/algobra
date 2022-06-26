package algobra

func NewMatrix[N Number](matr [][]N) (m *Matrix[N]) {
	m.RowsNumber = len(matr)
	m.ColumnsNumber = len(matr[0])
	m.Matr = matr
	return m
}

func NewEmptyMatrix[N Number](rowsNumber int, columnsNumber int) (m *Matrix[N]) {
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = columnsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
	}
	return m
}

func NewIdentityMatrix[N Number](rowsNumber int) (m *Matrix[N]) {
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = rowsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
		m.Matr[i][i] = 1
	}
	return m
}
