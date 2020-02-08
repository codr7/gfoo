package gfoo

import (
	"io"
	"time"
)

var Time TimeType

func init() {
	Time.Init("Time")
}

type TimeType struct {
	TypeBase
}

func (_ *TimeType) Compare(x, y interface{}) Order {
	xv, yv := x.(time.Time), y.(time.Time)
	
	if xv.Before(yv) {
		return Lt
	}

	if xv.After(yv) {
		return Gt
	}
	
	return Eq
}

func (_ *TimeType) Dump(val interface{}, out io.Writer) error {
	t, err := val.(time.Time).MarshalText()

	if err == nil {
		_, err = out.Write(t)
	}
	
	return err
}

func (_ *TimeType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, &Time, val)
}
