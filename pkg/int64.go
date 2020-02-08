package gfoo

import (
	"io"
	"strconv"
)

var Int64 Int64Type

func init() {
	Int64.Init("Int64")
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

func (_ *Int64Type) Unquote(pos Position, val interface{}) Form {
	return NewLiteral(pos, &Int64, val)
}
