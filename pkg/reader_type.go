package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TReader ReaderType

type ReaderType struct {
	ValTypeBase
}

func (_ *ReaderType) Compare(x, y Val) Order {
	xw, yw := x.data.(io.Reader), y.data.(io.Reader)
	return ComparePointer(unsafe.Pointer(&xw), unsafe.Pointer(&yw))
}

func (self *ReaderType) Dump(val Val, out io.Writer) error {
	w := val.data.(io.Reader)
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(&w))
	return err
}

func (_ *ReaderType) New(name string, parents...Type) ValType {
	t := new(ReaderType)
	t.Init(name, parents...)
	return t
}

func (self *ReaderType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ReaderType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
