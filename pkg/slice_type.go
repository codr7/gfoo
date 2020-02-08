package gfoo

import (
	"io"
)

var TSlice SliceType

func init() {
	TSlice.Init("Slice")
}

type SliceType struct {
	TypeBase
}

func (_ *SliceType) Compare(x, y interface{}) Order {
	return x.(*Slice).Compare(*y.(*Slice))
}

func (_ *SliceType) Dump(val interface{}, out io.Writer) error {
	return val.(*Slice).Dump(out)
}

func (_ *SliceType) Unquote(pos Pos, val interface{}) Form {
	return val.(*Slice).Unquote(pos)
}
