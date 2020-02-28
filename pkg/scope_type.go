package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TScope ScopeType

type ScopeType struct {
	ValTypeBase
}

func (_ *ScopeType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Scope)), unsafe.Pointer(y.data.(*Scope)))
}

func (_ *ScopeType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Scope(%v)", unsafe.Pointer(val.data.(*Scope)))
	return err
}

func (_ *ScopeType) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	scope = source.data.(*Scope)
	found := scope.Get(key)
	
	if found == nil || found.val == Nil {
		return Nil, scope.Error(pos, "Unknown identifier: %v", key)
	}

	return found.val, nil
}

func (self *ScopeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *ScopeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
