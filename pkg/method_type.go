package gfoo

import (
	"io"
)

var TMethod MethodType

func init() {
	TMethod.Init("Method")
}

type MethodType struct {
	TypeBase
}

func (_ *MethodType) Compare(x, y Val) Order {
	return CompareString(x.data.(*Method).Name(), y.data.(*Method).Name())
}

func (_ *MethodType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*Method).Name())
	return err
}

func (self *MethodType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
