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

func (typ *TimeType) Dump(val interface{}, out io.Writer) error {
	v, err := val.(time.Time).MarshalText()

	if err == nil {
		_, err = out.Write(v)
	}
	
	return err
}

func (typ *TimeType) Compare(x, y interface{}) Order {
	xv, yv := x.(time.Time), y.(time.Time)
	
	if xv.Before(yv) {
		return Lt
	}

	if xv.After(yv) {
		return Gt
	}
	
	return Eq
}
