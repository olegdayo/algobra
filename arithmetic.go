package algobra

import "errors"

// Negative function gets state of an object.
// Examples:
// 1 -> -1
// -0.5 -> 0.5
// A -> -1 * A
func Negative[N Number](n interface{}) (interface{}, error) {
	switch obj := n.(type) {
	case N:
		var ans N
		ans = -obj
		return ans, nil

	case *Matrix[N]:
		var ans Matrix[N]
		ans.Copy(obj)
		for i := 0; i < ans.RowsNumber; i++ {
			for j := 0; j < ans.ColumnsNumber; j++ {
				ans.Matr[i][j] = -ans.Matr[i][j]
			}
		}
		return ans, nil

	default:
		return nil, invalidTypeError
	}
}

// Add function adds an object (a number or matrix) to a new instance of given matrix.
func (m *Matrix[N]) Add(n interface{}) (ans *Matrix[N], err error) {
	switch obj := n.(type) {
	case N:
		ans.Copy(m)
		for i := 0; i < m.RowsNumber; i++ {
			for j := 0; j < m.ColumnsNumber; j++ {
				ans.Matr[i][j] += obj
			}
		}

	case *Matrix[N]:
		ans.Copy(m)
		for i := 0; i < m.RowsNumber; i++ {
			for j := 0; j < m.ColumnsNumber; j++ {
				ans.Matr[i][j] += obj.Matr[i][j]
			}
		}

	default:
		return nil, invalidTypeError
	}
	return ans, nil
}

// Subtract function subtracts an object (a number or matrix) from a new instance of given matrix.
func (m *Matrix[N]) Subtract(n interface{}) (*Matrix[N], error) {
	neg, err := Negative[N](n)
	if err != nil {
		return nil, err
	}
	return m.Add(neg)
}

// Multiply function multiplies a new instance of given matrix to an object (a number or matrix).
func (m *Matrix[N]) Multiply(n interface{}) (ans *Matrix[N], err error) {
	switch obj := n.(type) {
	case N:
		ans.Copy(m)
		for i := 0; i < ans.RowsNumber; i++ {
			for j := 0; j < ans.ColumnsNumber; j++ {
				ans.Matr[i][j] *= obj
			}
		}
		return ans, err

	case *Matrix[N]:
		ans.Copy(m)
		if m.ColumnsNumber != obj.RowsNumber {
			return nil, errors.New("number of columns of first matrix" +
				" " + "has to be equal to number of rows of second matrix")
		}

		ans = NewEmptyMatrix[N](m.GetRowsNumber(), obj.GetColumnsNumber())
		for h := 0; h < ans.RowsNumber; h++ {
			for i := 0; i < ans.ColumnsNumber; i++ {
				ans.Matr[h][i] = 0
				for j := 0; j < m.ColumnsNumber; j++ {
					ans.Matr[h][i] += m.Matr[h][j] * obj.Matr[j][i]
				}
			}
		}
		return ans, nil

	default:
		return nil, invalidTypeError
	}
}

// Divide function divides a new instance of given matrix on a number.
func (m *Matrix[N]) Divide(n N) (ans *Matrix[N]) {
	ans.Copy(m)
	for i := 0; i < ans.RowsNumber; i++ {
		for j := 0; j < ans.ColumnsNumber; j++ {
			ans.Matr[i][j] /= n
		}
	}
	return ans
}

// Power of matrix.
// An efficient (logN).
// A ^ -x == (A ^ -1) ^ x.
func (m *Matrix[N]) Power(pow int) (ans *Matrix[N], err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}

	if pow < 0 {
		inv, err := m.Inverse()
		if err != nil {
			return nil, err
		}
		return inv.Power(-pow)
	}

	ans = NewIdentityMatrix[N](m.RowsNumber)

	for pow > 0 {
		if pow%2 == 0 {
			pow /= 2
			ans, err = ans.Multiply(ans)
			if err != nil {
				return nil, err
			}
		} else {
			pow--
			ans, err = ans.Multiply(m)
			if err != nil {
				return nil, err
			}
		}
	}

	return ans, nil
}
