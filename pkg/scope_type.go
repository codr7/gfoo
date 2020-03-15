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

func (self *ScopeType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Scope)))
	return err
}

func (_ *ScopeType) For(val Val, action func(Val) error, scope *Scope, pos Pos) error {
	s := val.data.(*Scope)

	for k, b := range s.bindings {
		if err := action(NewVal(&TPair, NewPair(NewVal(&TId, k), b.val))); err != nil {
			return err
		}
	}

	return nil
}

func (_ *ScopeType) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	scope = source.data.(*Scope)
	found := scope.Get(key)
	
	if found == nil || found.val == Nil {
		return Nil, scope.Error(pos, "Unknown identifier: %v", key)
	}

	return found.val, nil
}

func (_ *ScopeType) New(name string, parents...Type) ValType {
	t := new(ScopeType)
	t.Init(name, parents...)
	return t
}

func (self *ScopeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *ScopeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
