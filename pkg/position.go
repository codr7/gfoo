package gfoo

type Position struct {
	Line, Column uint
	filename string
}

func NewPosition(filename string) Position {
	return Position{Line: 1, Column: 0, filename: filename}
}

