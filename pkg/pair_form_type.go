package gfoo

import (
	"io"
	"unsafe"
)

var TPairForm PairFormType

type PairFormType struct {
	ValTypeBase
}

func (_ *PairFormType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*PairForm)), unsafe.Pointer(y.data.(*PairForm)))
}

func (_ *PairFormType) Dump(val Val, out io.Writer) error {
	io.WriteString(out, "'");
	return val.data.(*PairForm).Dump(out)
}

func (_ *PairFormType) New(name string, parents...Type) ValType {
	t := new(PairFormType)
	t.Init(name, parents...)
	return t
}

func (self *PairFormType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *PairFormType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return val.data.(*PairForm)
}
