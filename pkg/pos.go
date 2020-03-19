package gfoo

const (
	MIN_LINE = 1
	MIN_COLUMN = 0
)

var NilPos = NewPos("n/a")

type Pos struct {
	source string
	line, column uint
}

func NewPos(source string) Pos {
	return Pos{source: source, line: MIN_LINE, column: MIN_COLUMN}
}
