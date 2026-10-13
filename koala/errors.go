package koala

import "errors"

var (
	ErrReadExcel       = errors.New("failed to read excel file")
	ErrMissingColumn   = errors.New("found missing column")
	ErrExtraColumn     = errors.New("found extra column")
	ErrDuplicateColumn = errors.New("found duplicate column")
)
