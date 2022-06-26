package algobra

import "errors"

var (
	boundariesError   = errors.New("index out of boundaries")
	squareMatrixError = errors.New("matrix has to be square")
	invalidTypeError  = errors.New("invalid type")
)
