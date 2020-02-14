package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TLambda LambdaType

func init() {
	TLambda.Init("Lambda")
}

type LambdaType struct {
	TypeBase
}

func (_ *LambdaType) Call(target Val, vm *VM, stack *Slice) error {
	return target.data.(*Lambda).Call(vm, stack)
}

func (_ *LambdaType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Lambda)), unsafe.Pointer(y.data.(*Lambda)))
}

func (_ *LambdaType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Lambda(%v)", unsafe.Pointer(val.data.(*Lambda)))
	return err
}

func (self *LambdaType) Unquote(val Val, pos Pos) Form {
	return NewGroup(pos, val.data.(*Lambda).forms)
}