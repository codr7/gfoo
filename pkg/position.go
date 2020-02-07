package gfoo

type Position struct {
	filename string
	line, column uint
}

func NewPosition(filename string) Position {
	return Position{filename: filename, line: 1, column: 0}
}

