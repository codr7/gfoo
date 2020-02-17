package gfoo

import (
	"io"
)

type Type interface {
	Bool(val Val) bool
	Call(target Val, scope *Scope, stack *Slice, pos Pos) error
	Compare(x, y Val) Order
	Dump(val Val, out io.Writer) error
	Name() string
	Unquote(val Val, scope *Scope, pos Pos) Form
}

type TypeBase struct {
	name string
}

func (self *TypeBase) Init(name string) {
	self.name = name
}

func (_ *TypeBase) Bool(val Val) bool {
	return true
}

func (_ *TypeBase) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(target)
	return nil
}

func (self *TypeBase) Name() string {
	return self.name
}
