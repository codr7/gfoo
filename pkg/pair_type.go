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

func (self *PairType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
