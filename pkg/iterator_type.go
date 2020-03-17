package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TIterator IteratorType

type IteratorType struct {
	ValTypeBase
}

func (_ *IteratorType) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	for {
		v, err := target.data.(Iterator)(scope, pos)

		if err != nil {
			return err
		}
		
		if v == nil || *v != Nil {
			break
		}
	}

	return nil
}

func (_ *IteratorType) Compare(x, y Val) Order {
	xi, yi := x.data.(Iterator), y.data.(Iterator)
	return ComparePointer(unsafe.Pointer(&xi), unsafe.Pointer(&yi))
}

func (self *IteratorType) Dump(val Val, out io.Writer) error {
	i := val.data.(Iterator)
	_, err := fmt.Fprintf(out, "Iterator(%v)", unsafe.Pointer(&i))
	return err
}

func (_ *IteratorType) Iterator(val Val, scope *Scope, pos Pos) (Iterator, error) {
	return val.data.(Iterator), nil
}

func (_ *IteratorType) New(name string, parents...Type) ValType {
	t := new(IteratorType)
	t.Init(name, parents...)
	return t
}

func (self *IteratorType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *IteratorType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
