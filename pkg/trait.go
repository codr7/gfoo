package gfoo

import (
	"io"
)

var TAny, TNumber Trait

type Trait struct {
	TypeBase
}

func NewTrait(name string, parents...Type) *Trait {
	return new(Trait).Init(name, parents...)
}

func (self *Trait) Init(name string, parents...Type) *Trait {
	self.TypeBase.Init(name, parents)
	return self
}

func (_ *Trait) Bool(val Val) bool {
	panic("Not implemented")
}

func (_ *Trait) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	panic("Not implemented")
}

func (self *Trait) Clone(val Val) interface{} {
	panic("Not implemented")
}

func (self *Trait) Compare(x, y Val) Order {
	panic("Not implemented")
}

func (self *Trait) Dump(val Val, out io.Writer) error {
	panic("Not implemented")	
}

func (self *Trait) Unquote(val Val, scope *Scope, pos Pos) Form {
	panic("Not implemented")
}
