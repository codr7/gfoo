package gfoo

import (
	"io"
	"unsafe"
)

var TScopeForm ScopeFormType

type ScopeFormType struct {
	ValTypeBase
}

func (_ *ScopeFormType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*ScopeForm)), unsafe.Pointer(y.data.(*ScopeForm)))
}

func (_ *ScopeFormType) Dump(val Val, out io.Writer) error {
	io.WriteString(out, "'");
	return val.data.(*ScopeForm).Dump(out)
}

func (_ *ScopeFormType) New(name string, parents...Type) ValType {
	t := new(ScopeFormType)
	t.Init(name, parents...)
	return t
}

func (self *ScopeFormType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ScopeFormType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return val.data.(*ScopeForm)
}
