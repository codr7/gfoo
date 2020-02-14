package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TThread ThreadType

func init() {
	TThread.Init("Thread")
}

type ThreadType struct {
	TypeBase
}

func (_ *ThreadType) Call(target Val, vm *VM, stack *Slice) error {
	return target.data.(*Thread).Call(stack)
}

func (_ *ThreadType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Thread)), unsafe.Pointer(y.data.(*Thread)))
}

func (_ *ThreadType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Thread(%v)", unsafe.Pointer(val.data.(*Thread)))
	return err
}

func (self *ThreadType) Unquote(val Val, pos Pos) Form {
	return NewLiteral(pos, val)
}
