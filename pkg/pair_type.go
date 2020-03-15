package gfoo

import (
	"io"
)

var TPair PairType

type PairType struct {
	ValTypeBase
}

func (_ *PairType) Compare(x, y Val) Order {
	return x.data.(Pair).Compare(y.data.(Pair))
}

func (_ *PairType) Dump(val Val, out io.Writer) error {
	return val.data.(Pair).Dump(out)
}

func (_ *PairType) For(val Val, action func(Val) error, scope *Scope, pos Pos) error {
	v := val.data.(Pair)

	if err := action(v.left); err != nil {
		return err
	}

	if err := action(v.right); err != nil {
		return err
	}

	return nil
}

func (_ *PairType) New(name string, parents...Type) ValType {
	t := new(PairType)
	t.Init(name, parents...)
	return t
}

func (_ *PairType) Negate(val *Val) {
	v := val.data.(Pair)
	v.left.Negate()
	v.right.Negate()
	val.data = v
}

func (_ *PairType) Print(val Val, out io.Writer) error {
	return val.data.(Pair).Print(out)
}

func (self *PairType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
