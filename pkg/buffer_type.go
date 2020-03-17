package gfoo

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

var TBuffer BufferType

type Buffer = bytes.Buffer

type BufferType struct {
	ValTypeBase
}

func (_ *BufferType) Bool(val Val) bool {
	return val.data.(*Buffer).Len() != 0
}

func (_ *BufferType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Buffer)), unsafe.Pointer(y.data.(*Buffer)))
}

func (self *BufferType) Clone(val Val) interface{} {
	return (*Buffer)(bytes.NewBuffer(val.data.(*Buffer).Bytes()))
}

func (self *BufferType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Buffer)))
	return err
}

func (_ *BufferType) New(name string, parents...Type) ValType {
	t := new(BufferType)
	t.Init(name, parents...)
	return t
}

func (_ *BufferType) Print(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*Buffer).String())
	return err
}

func (_ *BufferType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
