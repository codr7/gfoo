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
	ValTypeBase
}

func (_ *ThreadType) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	return target.data.(*Thread).Call(stack, pos)
}

func (_ *ThreadType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Thread)), unsafe.Pointer(y.data.(*Thread)))
}

func (_ *ThreadType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Thread(%v)", unsafe.Pointer(val.data.(*Thread)))
	return err
}

func (self *ThreadType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
