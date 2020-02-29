package gfoo

import (
	"io"
)

var TRecord RecordType

type RecordType struct {
	ValTypeBase
}

func (_ *RecordType) Compare(x, y Val) Order {
	x.data.(Record).Compare(y.data.(Record))
	return Eq
}

func (_ *RecordType) Dump(val Val, out io.Writer) error {
	return val.data.(Record).Dump(out)
}

func (_ *RecordType) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	return source.data.(Record).Get(key, Nil), nil
}

func (self *RecordType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *RecordType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
