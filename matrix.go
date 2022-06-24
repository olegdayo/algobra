package algobra

type Matrix[num Number] struct {
	Rows    int
	Columns int
	Matr    [][]num
}

func NewMatrix[num Number](matr [][]num) (m *Matrix[num]) {
	m.Rows = len(matr)
	m.Columns = len(matr[0])
	m.Matr = matr
	return m
}

func (m *Matrix[num]) At(i int, j int) num {
	return m.Matr[i][j]
}

func (m *Matrix[num]) GetRowsNumber() int {
	return m.Rows
}

func (m *Matrix[num]) GetColumnsNumber() int {
	return m.Columns
}
