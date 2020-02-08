package gfoo

import (
	"io"
	"strconv"
)

var TInt64 Int64Type

func init() {
	TInt64.Init("Int64")
}

type Int64Type struct {
	TypeBase
}

func (_ *Int64Type) Compare(x, y interface{}) Order {
	xv, yv := x.(int64), y.(int64)

	if xv < yv {
		return Lt
	}

	if xv > yv {
		return Gt
	}

	return Eq
}

func (_ *Int64Type) Dump(val interface{}, out io.Writer) error {
	_, err := io.WriteString(out, strconv.FormatInt(val.(int64), 10))
	return err
}

func (self *Int64Type) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}
