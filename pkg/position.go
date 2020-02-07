package gfoo

const (
	MIN_LINE = 1
	MIN_COLUMN = 0
)

type Position struct {
	filename string
	line, column uint
}

func NewPosition(filename string) Position {
	return Position{filename: filename, line: MIN_LINE, column: MIN_COLUMN}
}

