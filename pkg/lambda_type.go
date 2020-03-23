package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TLambda LambdaType

type LambdaType struct {
	ValTypeBase
}

func (_ *LambdaType) Call(target Val, thread *Thread, stack *Slice, pos Pos) error {
	return target.data.(*Lambda).Call(thread, stack, pos)
}

func (_ *LambdaType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Lambda)), unsafe.Pointer(y.data.(*Lambda)))
}

func (self *LambdaType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Lambda)))
	return err
}

func (_ *LambdaType) New(name string, parents...Type) ValType {
	t := new(LambdaType)
	t.Init(name, parents...)
	return t
}

func (self *LambdaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *LambdaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
