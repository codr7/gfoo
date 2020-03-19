package gfoo

import (
	"io"
	"unsafe"
)

var TScope ScopeType

type ScopeType struct {
	ValTypeBase
}

func (_ *ScopeType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*ScopeForm)), unsafe.Pointer(y.data.(*ScopeForm)))
}

func (_ *ScopeType) Dump(val Val, out io.Writer) error {
	io.WriteString(out, "'");
	return val.data.(*ScopeForm).Dump(out)
}

func (_ *ScopeType) New(name string, parents...Type) ValType {
	t := new(ScopeType)
	t.Init(name, parents...)
	return t
}

func (self *ScopeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ScopeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return val.data.(*ScopeForm)
}
