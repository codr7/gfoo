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
	TypeBase
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

func (self *TimeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
