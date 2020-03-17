package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TWriter WriterType

type WriterType struct {
	ValTypeBase
}

func (_ *WriterType) Compare(x, y Val) Order {
	xw, yw := x.data.(io.Writer), y.data.(io.Writer)
	return ComparePointer(unsafe.Pointer(&xw), unsafe.Pointer(&yw))
}

func (self *WriterType) Dump(val Val, out io.Writer) error {
	w := val.data.(io.Writer)
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(&w))
	return err
}

func (_ *WriterType) New(name string, parents...Type) ValType {
	t := new(WriterType)
	t.Init(name, parents...)
	return t
}

func (self *WriterType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *WriterType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
