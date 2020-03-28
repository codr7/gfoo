package gfoo

import (
	"fmt"
	"io"
)

var TFunction FunctionType

type FunctionType struct {
	ValTypeBase
}

func (_ *FunctionType) Call(target Val, thread *Thread, stack *Stack, pos Pos) error {
	return target.data.(*Function).Call(thread, stack, pos)
}

func (_ *FunctionType) Compare(x, y Val) Order {
	return CompareString(x.data.(*Function).name, y.data.(*Function).name)
}

func (_ *FunctionType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Function(%v)", val.data.(*Function).name)
	return err
}

func (_ *FunctionType) New(name string, parents...Type) ValType {
	t := new(FunctionType)
	t.Init(name, parents...)
	return t
}

func (self *FunctionType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *FunctionType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
