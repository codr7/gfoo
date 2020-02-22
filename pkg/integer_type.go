package gfoo

import (
	"io"
	"math/big"
)

var TInteger IntegerType

type IntegerType struct {
	ValTypeBase
}

func (_ *IntegerType) Bool(val Val) bool {
	v := val.data.(*big.Int)
	return !(v.IsInt64() && v.Int64() == 0)
}

func (_ *IntegerType) Compare(x, y Val) Order {
	return Order(x.data.(*big.Int).Cmp(y.data.(*big.Int)))
}

func (_ *IntegerType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*big.Int).String())
	return err
}

func (self *IntegerType) Is(x, y Val) bool {
	return self.Compare(x, y) == Eq
}

func (self *IntegerType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
