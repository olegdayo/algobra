package algobra

func MakeCopy[N Number](m *Matrix[N]) (ans *Matrix[N]) {
	ans = NewEmptyMatrix[N](m.RowsNumber, m.ColumnsNumber)
	for i, row := range m.Matr {
		for j, num := range row {
			ans.Matr[i][j] = num
		}
	}
	return ans
}

func Copy[N Number](from *Matrix[N], to *Matrix[N]) {
	to = MakeCopy[N](from)
}

func (m *Matrix[N]) Copy(other *Matrix[N]) {
	Copy[N](m, other)
}
