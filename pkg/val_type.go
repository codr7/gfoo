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
	Get(source Val, key string, scope *Scope, pos Pos) (Val, error)
	Print(val Val, out io.Writer) error
	Is(x, y Val) bool
	Iter(val Val, scope *Scope, pos Pos) (Iter, error)
	Negate(val *Val)
	New(name string, parents...Type) ValType
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

func (self *ValTypeBase) Call(target Val, scope *Scope, stack *Slice, pos Pos) error {
	return scope.Error(pos, "Call not supported for type: %v", self.name)
}

func (self *ValTypeBase) Clone(val Val) interface{} {
	return val.data
}

func (self *ValTypeBase) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	return Nil, scope.Error(pos, "Dot access not supported for type: %v", self.name)
}

func (self *ValTypeBase) Is(x, y Val) bool {
	return x.data == y.data
}

func (self *ValTypeBase) Iter(val Val, scope *Scope, pos Pos) (Iter, error) {
	return nil, scope.Error(pos, "Iteration not supported for type: %v", self.name)
}

func (self *ValTypeBase) Name() string {
	return self.name
}

func (self *ValTypeBase) Negate(val *Val) {
	val.Init(&TBool, false)
}


