package gfoo

import (
	"fmt"
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
	d := val.data.(TimeDelta)
	_, err := fmt.Fprintf(out, "TimeDelta(%v %v %v)", d.years, d.months, d.days) 
	return err
}

func (_ *TimeDeltaType) New(name string, parents...Type) ValType {
	t := new(TimeDeltaType)
	t.Init(name, parents...)
	return t
}

func (self *TimeDeltaType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *TimeDeltaType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
