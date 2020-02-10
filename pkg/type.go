package gfoo

import (
	"io"
)

type Type interface {
	Call(target Val, vm *VM, stack *Slice) error
	Compare(x, y interface{}) Order
	Dump(val interface{}, out io.Writer) error
	Name() string
	Unquote(pos Pos, val interface{}) Form
}

type TypeBase struct {
	name string
}

func (self *TypeBase) Init(name string) {
	self.name = name
}

func (_ *TypeBase) Call(target Val, vm *VM, stack *Slice) error {
	stack.Push(target)
	return nil
}

func (self *TypeBase) Name() string {
	return self.name
}
