package gfoo

import (
	"io"
)

var TMeta MetaType

type MetaType struct {
	ValTypeBase
}

func (_ *MetaType) Compare(x, y Val) Order {
	xt, yt := x.data.(Type), y.data.(Type)

	if xt.Isa(yt) != nil {
		return Gt
	}

	if yt.Isa(xt) != nil {
		return Lt
	}

	return Eq
}

func (_ *MetaType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(Type).Name())
	return err
}

func (self *MetaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *MetaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
