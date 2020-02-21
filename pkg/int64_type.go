package gfoo

import (
	"io"
	"strconv"
)

var TInt64 Int64Type

type Int64Type struct {
	ValTypeBase
}

func (_ *Int64Type) Bool(val Val) bool {
	return val.data.(int64) != 0
}

func (_ *Int64Type) Compare(x, y Val) Order {
	xv, yv := x.data.(int64), y.data.(int64)

	if xv < yv {
		return Lt
	}

	if xv > yv {
		return Gt
	}

	return Eq
}

func (_ *Int64Type) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, strconv.FormatInt(val.data.(int64), 10))
	return err
}

func (self *Int64Type) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
