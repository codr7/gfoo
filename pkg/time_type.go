package gfoo

import (
	"io"
	"time"
)

var TTime TimeType

var MinTime, MaxTime time.Time

func init() {
	TTime.Init("Time")

	MinTime = time.Date(0, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	MaxTime = time.Date(9999, time.Month(12), 31, 23, 59, 59, 999999999, time.UTC)
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

func (self *TimeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *TimeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
