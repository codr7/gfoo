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

func (_ *LambdaType) Compare(x, y interface{}) Order {
	return ComparePointer(unsafe.Pointer(x.(*Lambda)), unsafe.Pointer(y.(*Lambda)))
}

func (_ *LambdaType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Lambda(%v)", unsafe.Pointer(val.(*Lambda)))
	return err
}

func (self *LambdaType) Unquote(pos Pos, val interface{}) Form {
	return NewGroup(pos, val.(*Lambda).forms)
}
