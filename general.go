package algobra

import "errors"

// At function:
// m.Matr[i][j] is almost the same to m.At(i, j).
func (m *Matrix[N]) At(i int, j int) (N, error) {
	if i < 0 || j < 0 {
		return 0, boundariesError
	}

	if i > m.RowsNumber-1 || j > m.ColumnsNumber-1 {
		return 0, boundariesError
	}

	return m.Matr[i][j], nil
}

func (m *Matrix[N]) GetRowsNumber() int {
	return m.RowsNumber
}

func (m *Matrix[N]) GetColumnsNumber() int {
	return m.ColumnsNumber
}

// InsertRow inserts a row.
func (m *Matrix[N]) InsertRow(index int, row []N) error {
	if len(row) != m.ColumnsNumber {
		return errors.New("wrong number of elements in a row")
	}

	if index < 0 || index > m.RowsNumber-1 {
		return boundariesError
	}

	m.Matr = append(m.Matr[:index+1], m.Matr[index:]...)
	m.Matr[index] = row

	m.RowsNumber++
	return nil
}

// InsertColumn inserts a column.
func (m *Matrix[N]) InsertColumn(index int, column []N) error {
	if len(column) != m.RowsNumber {
		return errors.New("wrong number of elements in a column")
	}

	if index < 0 || index > m.ColumnsNumber-1 {
		return boundariesError
	}

	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = append(m.Matr[i][:index+1], m.Matr[i][index:]...)
		m.Matr[i][index] = column[i]
	}

	m.ColumnsNumber++
	return nil
}

// IsSquare checks if matrix is square.
func (m *Matrix[N]) IsSquare() bool {
	return m.RowsNumber == m.ColumnsNumber
}

// ToBinary function:
// |1 2 3|    |1 0 1|
// |4 5 6| -> |0 1 0|
// |7 8 9|    |1 0 1|
func ToBinary[I Int](m *Matrix[I]) (ans *Matrix[I], err error) {
	ans.Copy(m)
	for i := 0; i < ans.RowsNumber; i++ {
		for j := 0; j < ans.ColumnsNumber; j++ {
			ans.Matr[i][j] = m.Matr[i][j] % 2
		}
	}
	return ans, err
}
