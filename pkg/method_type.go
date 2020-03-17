package gfoo

import (
	"io"
)

var TMethod MethodType

type MethodType struct {
	ValTypeBase
}

func (_ *MethodType) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	m := target.data.(*Method)
	
	if !m.Applicable(stack) {
		return scope.Error(pos, "Method not applicable: %v %v", m.name, stack)
	}

	return m.Call(scope, stack, pos)
}

func (_ *MethodType) Compare(x, y Val) Order {
	return CompareString(x.data.(*Method).name, y.data.(*Method).name)
}

func (_ *MethodType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*Method).name)
	return err
}

func (_ *MethodType) New(name string, parents...Type) ValType {
	t := new(MethodType)
	t.Init(name, parents...)
	return t
}

func (self *MethodType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *MethodType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
