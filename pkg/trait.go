package gfoo

import (
	"io"
)

var TAny, TNumber, TOption, TSequence Trait

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
	panic("Abstract method")
}

func (_ *Trait) Call(target Val, scope *Scope, stack *Stack, pos Pos) error {
	panic("Abstract method")
}

func (self *Trait) Clone(val Val) interface{} {
	panic("Abstract method")
}

func (self *Trait) Compare(x, y Val) Order {
	panic("Abstract method")
}

func (self *Trait) Dump(val Val, out io.Writer) error {
	panic("Abstract method")
}

func (self *Trait) Unquote(val Val, scope *Scope, pos Pos) Form {
	panic("Abstract method")
}
