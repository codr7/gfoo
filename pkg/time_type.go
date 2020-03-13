package gfoo

import (
	"io"
	"time"
)

var TTime TimeType

func init() {
	TTime.Init("Time")
}

type TimeType struct {
	ValTypeBase
}

func (_ *TimeType) Bool(val Val) bool {
	var zero time.Time
	return val.data.(time.Time) != zero
}

func (_ *TimeType) Compare(x, y Val) Order {
	xv, yv := x.data.(time.Time), y.data.(time.Time)
	
	if xv.Before(yv) {
		return Lt
	}

	if xv.After(yv) {
		return Gt
	}
	
	return Eq
}

func (_ *TimeType) Dump(val Val, out io.Writer) error {
	t, err := val.data.(time.Time).MarshalText()

	if err == nil {
		_, err = out.Write(t)
	}
	
	return err
}

func (_ *TimeType) New(name string, parents...Type) ValType {
	t := new(TimeType)
	t.Init(name, parents...)
	return t
}

func (self *TimeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *TimeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
