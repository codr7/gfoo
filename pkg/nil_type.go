package gfoo

import (
	"io"
)

var TNil NilType

type NilType struct {
	ValTypeBase
}

func (_ *NilType) Bool(val Val) bool {
	return false
}

func (_ *NilType) Compare(x, y Val) Order {
	return Eq
}

func (_ *NilType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, "NIL")
	return err
}

func (_ *NilType) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	return Nil, nil
}

func (_ *NilType) Negate(val *Val) {}

func (_ *NilType) New(name string, parents...Type) ValType {
	t := new(NilType)
	t.Init(name, parents...)
	return t
}

func (self *NilType) Print(val Val, out io.Writer) error {
	return nil
}

func (self *NilType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
