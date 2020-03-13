package gfoo

import (
	"io"
)

var TMeta MetaType

type MetaType struct {
	ValTypeBase
}

func (_ *MetaType) Compare(x, y Val) Order {
	return CompareString(x.data.(Type).Name(), y.data.(Type).Name())
}

func (_ *MetaType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(Type).Name())
	return err
}

func (_ *MetaType) New(name string, parents...Type) ValType {
	t := new(MetaType)
	t.Init(name, parents...)
	return t
}

func (self *MetaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *MetaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
