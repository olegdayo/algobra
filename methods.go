package algobra

import "errors"

func (m *Matrix[N]) At(i int, j int) (N, error) {
	if i < 0 || j < 0 {
		return nil, errors.New("index out of boundaries")
	}

	if i > m.RowsNumber-1 || j > m.ColumnsNumber-1 {
		return nil, errors.New("index out of boundaries")
	}

	return m.Matr[i][j], nil
}

func (m *Matrix[N]) Copy(other *Matrix[N]) {
	Copy[N](m, other)
}

func (m *Matrix[N]) GetRowsNumber() int {
	return m.RowsNumber
}

func (m *Matrix[N]) GetColumnsNumber() int {
	return m.ColumnsNumber
}

func (m *Matrix[N]) InsertRow(index int, row []N) error {
	if len(row) != m.ColumnsNumber {
		return errors.New("wrong number of elements in row")
	}
	m.Matr = append(m.Matr[:index+1], m.Matr[index:]...)
	m.Matr[index] = row
	return nil
}

func (m *Matrix[N]) InsertColumn(index int, column []N) error {
	if len(column) != m.RowsNumber {
		return errors.New("wrong number of elements in column")
	}
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = append(m.Matr[i][:index+1], m.Matr[i][index:]...)
		m.Matr[i][index] = column[i]
	}
	return nil
}

func (m *Matrix[N]) MultiplyTo(m2 *Matrix[N]) (*Matrix[N], error) {
	return Multiply(m, m2)
}

func (m *Matrix[N]) ToPower(n int) (ans *Matrix[N], err error) {
	return Power(m, n)
}

func (m *Matrix[N]) T() (ans *Matrix[N], err error) {
	ans = NewEmptyMatrix[N](m.ColumnsNumber, m.RowsNumber)
	for i := 0; i < ans.RowsNumber; i++ {
		for j := 0; j < ans.ColumnsNumber; j++ {
			ans.Matr[i][j] = m.Matr[j][i]
		}
	}
	return ans, nil
}
