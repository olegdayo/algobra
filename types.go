package algobra

type Int interface {
	int | int8 | int16 | int32 | int64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Int | UInt | Float
}

// Elems is a type for convenient way of passing matrix into function.
type Elems[N Number] [][]N

type Matrix[N Number] struct {
	RowsNumber    int
	ColumnsNumber int
	Matr          Elems[N]
}
