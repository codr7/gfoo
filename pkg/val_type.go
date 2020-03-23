package gfoo

import (
	"io"
)

type ValType interface {
	Type
	Bool(val Val) bool
	Clone(val Val) interface{}
	Compare(x, y Val) Order
	Dump(val Val, out io.Writer) error
	Get(source Val, key string, pos Pos) (Val, error)
	Print(val Val, out io.Writer) error
	Is(x, y Val) bool
	Keys(val Val) []string
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

func (_ *ValTypeBase) Clone(val Val) interface{} {
	return val.data
}

func (self *ValTypeBase) Get(source Val, key string, pos Pos) (Val, error) {
	return Nil, Error(pos, "Dot access not supported for type: %v", self.name)
}

func (_ *ValTypeBase) Is(x, y Val) bool {
	return x.data == y.data
}

func (_ *ValTypeBase) Keys(val Val) []string {
	return nil
}

func (self *ValTypeBase) Name() string {
	return self.name
}

func (_ *ValTypeBase) Negate(val *Val) {
	val.Init(&TBool, false)
}


