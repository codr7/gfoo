package gfoo

import (
	"io"
)

type ValType interface {
	Type
	Bool(val Val) bool
	Call(target Val, scope *Scope, stack *Slice, pos Pos) error
	Clone(val Val) interface{}
	Compare(x, y Val) Order
	Dump(val Val, out io.Writer) error
	Print(val Val, out io.Writer) error
	Unquote(val Val, scope *Scope, pos Pos) Form
}

type ValTypeBase struct {
	TypeBase
}

func (self *ValTypeBase) Init(name string, parents...Type) {
	self.TypeBase.Init(name, parents)
}

func (_ *ValTypeBase) Bool(val Val) bool {
	return true
}

func (_ *ValTypeBase) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(target)
	return nil
}

func (self *ValTypeBase) Clone(val Val) interface{} {
	return val.data
}

func (self *ValTypeBase) Name() string {
	return self.name
}
