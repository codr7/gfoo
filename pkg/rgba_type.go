package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TRgba RgbaType

type RgbaType struct {
	ValTypeBase
}

func (_ *RgbaType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Rgba)), unsafe.Pointer(y.data.(*Rgba)))
}

func (self *RgbaType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Rgba)))
	return err
}

func (_ *RgbaType) New(name string, parents...Type) ValType {
	t := new(RgbaType)
	t.Init(name, parents...)
	return t
}

func (self *RgbaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *RgbaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
