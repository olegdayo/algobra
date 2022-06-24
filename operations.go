package algobra

func Multiply[N Number](m1 *Matrix[N], m2 *Matrix[N]) (ans *Matrix[N], err error) {
	ans.Rows = m1.GetRowsNumber()
	ans.Columns = m2.Columns
	ans.Matr = make([][]N, ans.Rows)
	for i := 0; i < ans.Rows; i++ {
		ans.Matr[i] = make([]N, ans.Columns)
	}

	for h := 0; h < ans.Rows; h++ {
		for i := 0; i < ans.Columns; i++ {
			ans.Matr[h][i] = 0
			for j := 0; j < m1.Columns; j++ {
				ans.Matr[h][i] += m1.Matr[h][j] * m2.Matr[j][i]
			}
		}
	}

	return ans, nil
}

func (um *Matrix[N]) MultiplyTo(m2 *Matrix[N]) (*Matrix[N], error) {
	return Multiply(um, m2)
}

func Power[N Number](m *Matrix[N], n int) (ans *Matrix[N], err error) {
	// Rewrite to effective
	ans.Rows = m.Rows
	ans.Columns = m.Columns
	ans.Matr = m.Matr
	for i := 0; i < n-1; i++ {
		ans, _ = ans.MultiplyTo(m)
	}
	return ans, nil
}

func (um *Matrix[N]) ToPower(n int) (ans *Matrix[N], err error) {
	return Power(um, n)
}

func (um *Matrix[N]) N() (ans *Matrix[N], err error) {
	ans.Rows = um.Columns
	ans.Columns = um.Rows
	ans.Matr = make([][]N, ans.Rows)
	for i := 0; i < ans.Rows; i++ {
		ans.Matr[i] = make([]N, ans.Columns)
	}
	for i := 0; i < ans.Rows; i++ {
		for j := 0; j < ans.Columns; j++ {
			ans.Matr[i][j] = um.Matr[j][i]
		}
	}
	return ans, nil
}

func ToBinary[I Int](m *Matrix[I]) (ans *Matrix[I], err error) {
	ans.Rows = m.GetRowsNumber()
	ans.Columns = m.Columns
	ans.Matr = m.Matr
	for i := 0; i < ans.Rows; i++ {
		for j := 0; j < ans.Columns; j++ {
			ans.Matr[i][j] = m.Matr[i][j] % 2
		}
	}
	return ans, err
}
