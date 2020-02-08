package gfoo

import (
	"io"
)

var TBool BoolType

func init() {
	TBool.Init("Bool")
}

type BoolType struct {
	TypeBase
}

func (_ *BoolType) Compare(x, y interface{}) Order {
	xv, yv := x.(bool), y.(bool)

	if xv && !yv {
		return Lt
	}

	if !xv && yv {
		return Gt
	}

	return Eq
}

func (_ *BoolType) Dump(val interface{}, out io.Writer) error {
	var s string
	
	if val.(bool) {
		s = "T"
	} else {
		s = "F"
	}
	
	_, err := io.WriteString(out, s)
	return err
}

func (self *BoolType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}
