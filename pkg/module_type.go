package gfoo

import (
	"fmt"
	"io"
	"unsafe"
)

var TModule ModuleType

type ModuleType struct {
	ValTypeBase
}

func (_ *ModuleType) Compare(x, y Val) Order {
	return ComparePointer(unsafe.Pointer(x.data.(*Module)), unsafe.Pointer(y.data.(*Module)))
}

func (self *ModuleType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v(%v)", self.name, unsafe.Pointer(val.data.(*Module)))
	return err
}

func (_ *ModuleType) Get(source Val, key string, pos Pos) (Val, error) {
	found := source.data.(*Module).Get(key)
	
	if found == nil || found.val == Undefined {
		return Nil, Error(pos, "Unknown identifier: %v", key)
	}

	return found.val, nil
}

func (_ *ModuleType) Keys(val Val) []string {
	var out []string

	for k, _ := range val.data.(*Module).bindings {
		out = append(out, k)
	}
	
	return out
}

func (_ *ModuleType) New(name string, parents...Type) ValType {
	t := new(ModuleType)
	t.Init(name, parents...)
	return t
}

func (self *ModuleType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *ModuleType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
