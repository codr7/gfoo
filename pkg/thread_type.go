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

func (_ *ThreadType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Thread)), unsafe.Pointer(y.data.(*Thread)))
}

func (self *ThreadType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Thread)))
	return err
}

func (_ *ThreadType) New(name string, parents...Type) ValType {
	t := new(ThreadType)
	t.Init(name, parents...)
	return t
}

func (self *ThreadType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ThreadType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
