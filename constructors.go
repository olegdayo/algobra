package algobra

func NewMatrix[N Number](elems Elems[N]) (m *Matrix[N]) {
	m = new(Matrix[N])
	m.RowsNumber = len(elems)
	if m.RowsNumber == 0 {
		m.ColumnsNumber = 0
	} else {
		m.ColumnsNumber = len(elems[0])
	}
	m.Matr = elems
	return m
}

func NewEmptyMatrix[N Number](rowsNumber int, columnsNumber int) (m *Matrix[N]) {
	m = new(Matrix[N])
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = columnsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
	}
	return m
}

func NewIdentityMatrix[N Number](rowsNumber int) (m *Matrix[N]) {
	m = new(Matrix[N])
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = rowsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
		m.Matr[i][i] = 1
	}
	return m
}

func (m1 *Matrix[N]) Equals(m2 *Matrix[N]) bool {
	if m1.RowsNumber != m2.RowsNumber {
		return false
	}

	if m1.ColumnsNumber != m2.ColumnsNumber {
		return false
	}

	for i := 0; i < m1.RowsNumber; i++ {
		for j := 0; j < m2.ColumnsNumber; j++ {
			if float64(m1.Matr[i][j]-m2.Matr[i][j]) > EPS {
				return false
			}
		}
	}

	return true
}
