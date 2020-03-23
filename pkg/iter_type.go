package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TIter IterType

type IterType struct {
	ValTypeBase
}

func (_ *IterType) Compare(x, y Val) Order {
	xi, yi := x.data.(Iter), y.data.(Iter)
	return ComparePointer(unsafe.Pointer(&xi), unsafe.Pointer(&yi))
}

func (self *IterType) Dump(val Val, out io.Writer) error {
	i := val.data.(Iter)
	_, err := fmt.Fprintf(out, "Iter(%v)", unsafe.Pointer(&i))
	return err
}

func (_ *IterType) Iter(val Val, pos Pos) (Iter, error) {
	return val.data.(Iter), nil
}

func (_ *IterType) New(name string, parents...Type) ValType {
	t := new(IterType)
	t.Init(name, parents...)
	return t
}

func (self *IterType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *IterType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
