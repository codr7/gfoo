package gfoo

import (
	"io"
)

var TType MetaType

func init() {
	TType.Init("Type")
}

type MetaType struct {
	TypeBase
}

func (_ *MetaType) Compare(x, y interface{}) Order {
	return CompareString(x.(Type).Name(), y.(Type).Name())
}

func (_ *MetaType) Dump(val interface{}, out io.Writer) error {
	_, err := io.WriteString(out, val.(Type).Name())
	return err
}

func (self *MetaType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}
