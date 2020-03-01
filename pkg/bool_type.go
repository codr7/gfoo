package gfoo

import (
	"io"
)

var TBool BoolType

type BoolType struct {
	ValTypeBase
}

func (_ *BoolType) Bool(val Val) bool {
	return val.data.(bool)
}

func (_ *BoolType) Compare(x, y Val) Order {
	xv, yv := x.data.(bool), y.data.(bool)

	if xv && !yv {
		return Lt
	}

	if !xv && yv {
		return Gt
	}

	return Eq
}

func (_ *BoolType) Dump(val Val, out io.Writer) error {
	var s string
	
	if val.data.(bool) {
		s = "T"
	} else {
		s = "F"
	}
	
	_, err := io.WriteString(out, s)
	return err
}

func (_ *BoolType) New(name string, parents...Type) ValType {
	t := new(BoolType)
	t.Init(name, parents...)
	return t
}

func (_ *BoolType) Print(val Val, out io.Writer) error {
	var s string
	
	if val.data.(bool) {
		s = "true"
	} else {
		s = "false"
	}
	
	_, err := io.WriteString(out, s)
	return err
}

func (self *BoolType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
