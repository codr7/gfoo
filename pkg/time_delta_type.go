package gfoo

import (
	"io"
)

var TTimeDelta TimeDeltaType

func init() {
	TTimeDelta.Init("TimeDelta")
}

type TimeDeltaType struct {
	ValTypeBase
}

func (_ *TimeDeltaType) Bool(val Val) bool {
	var zero TimeDelta
	return val.data.(TimeDelta) != zero
}

func (_ *TimeDeltaType) Compare(x, y Val) Order {
	return x.data.(TimeDelta).Compare(y.data.(TimeDelta))
}

func (_ *TimeDeltaType) Dump(val Val, out io.Writer) error {
	return val.data.(TimeDelta).Dump(out)
}

func (_ *TimeDeltaType) New(name string, parents...Type) ValType {
	t := new(TimeDeltaType)
	t.Init(name, parents...)
	return t
}

func (self *TimeDeltaType) Print(val Val, out io.Writer) error {
	return val.data.(TimeDelta).Dump(out)
}

func (_ *TimeDeltaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
