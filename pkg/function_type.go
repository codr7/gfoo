package gfoo

import (
	"fmt"
	"io"
)

var TFunction FunctionType

type FunctionType struct {
	ValTypeBase
}

func (_ *FunctionType) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	return target.data.(*Function).Call(scope, stack, pos)
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
