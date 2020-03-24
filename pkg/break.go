package gfoo

type BreakType struct {}

func (e *BreakType) Error() string {
	return "Break"
}

var Break BreakType
