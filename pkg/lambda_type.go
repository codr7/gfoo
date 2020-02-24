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

func (_ *LambdaType) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	return target.data.(*Lambda).Call(stack, pos)
}

func (_ *LambdaType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Lambda)), unsafe.Pointer(y.data.(*Lambda)))
}

func (_ *LambdaType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Lambda(%v)", unsafe.Pointer(val.data.(*Lambda)))
	return err
}

func (self *LambdaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *LambdaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
