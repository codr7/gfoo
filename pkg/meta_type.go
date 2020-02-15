package gfoo

import (
	"io"
)

var TMeta MetaType

func init() {
	TMeta.Init("Type")
}

type MetaType struct {
	TypeBase
}

func (_ *MetaType) Compare(x, y Val) Order {
	return CompareString(x.data.(Type).Name(), y.data.(Type).Name())
}

func (_ *MetaType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(Type).Name())
	return err
}

func (self *MetaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
