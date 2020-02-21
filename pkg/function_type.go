package gfoo

import (
	"fmt"
	"io"
)

var TFunction FunctionType

func init() {
	TFunction.Init("Function")
}

type FunctionType struct {
	TypeBase
}

func (_ *FunctionType) Compare(x, y Val) Order {
	return CompareString(x.data.(*Function).name, y.data.(*Function).name)
}

func (_ *FunctionType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Function(%v)", val.data.(*Function).name)
	return err
}

func (self *FunctionType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
