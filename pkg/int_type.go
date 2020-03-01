package gfoo

import (
	"io"
	"math/big"
)

var TInt IntType

type IntType struct {
	ValTypeBase
}

func (_ *IntType) Bool(val Val) bool {
	v := val.data.(*big.Int)
	return !(v.IsInt64() && v.Int64() == 0)
}

func (_ *IntType) Compare(x, y Val) Order {
	return Order(x.data.(*big.Int).Cmp(y.data.(*big.Int)))
}

func (_ *IntType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*big.Int).String())
	return err
}

func (_ *IntType) New(name string, parents...Type) ValType {
	t := new(IntType)
	t.Init(name, parents...)
	return t
}

func (self *IntType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *IntType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
