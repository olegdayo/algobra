package algobra

type DiagonalMatrix[num Number] struct {
	Rows    int
	Columns int
	Matr    []num
}

func NewDiagonalMatrix[num Number](matr []num) (dm *DiagonalMatrix[num]) {
	dm.Rows = len(matr)
	dm.Columns = len(matr)
	dm.Matr = matr
	return dm
}

func (dm *DiagonalMatrix[num]) At(i int, j int) num {
	if i != j {
		return 0
	}
	return dm.Matr[i]
}

func (dm *DiagonalMatrix[num]) GetRowsNumber() int {
	return dm.Rows
}

func (dm *DiagonalMatrix[num]) GetColumnsNumber() int {
	return dm.Columns
}
