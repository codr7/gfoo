package gfoo

import (
	"archive/zip"
	"fmt"
	"io"
	"unsafe"
)

var TZipWriter ZipWriterType

type ZipWriterType struct {
	ValTypeBase
}

func (_ *ZipWriterType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*zip.Writer)), unsafe.Pointer(y.data.(*zip.Writer)))
}

func (self *ZipWriterType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*zip.Writer)))
	return err
}

func (_ *ZipWriterType) New(name string, parents...Type) ValType {
	t := new(ZipWriterType)
	t.Init(name, parents...)
	return t
}

func (self *ZipWriterType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ZipWriterType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
