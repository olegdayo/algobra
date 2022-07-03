package algobra

// MakeCopy makes a new copy of matrix.
func MakeCopy[N Number](m *Matrix[N]) (ans *Matrix[N]) {
	ans = NewEmptyMatrix[N](m.RowsNumber, m.ColumnsNumber)
	for i, row := range m.Matr {
		for j, num := range row {
			ans.Matr[i][j] = num
		}
	}
	return ans
}

// Copy copies from matrix 1 to matrix 2.
func Copy[N Number](from *Matrix[N], to *Matrix[N]) {
	to = MakeCopy[N](from)
}

// Copy copies argument matrix to given one.
func (m *Matrix[N]) Copy(other *Matrix[N]) {
	Copy[N](other, m)
}
