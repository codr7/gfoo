package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TScope ScopeType

func init() {
	TScope.Init("Scope")
}

type ScopeType struct {
	TypeBase
}

func (_ *ScopeType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*ScopeForm)), unsafe.Pointer(y.data.(*ScopeForm)))
}

func (_ *ScopeType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Scope(%v)", unsafe.Pointer(val.data.(*ScopeForm)))
	return err
}

func (self *ScopeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewScopeForm(val.data.(*ScopeForm).body, pos)
}
