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

func (_ *SliceType) Compare(x, y Val) Order {
	return x.data.(*Slice).Compare(*y.data.(*Slice))
}

func (_ *SliceType) Dump(val Val, out io.Writer) error {
	return val.data.(*Slice).Dump(out)
}

func (_ *SliceType) Unquote(val Val, pos Pos) Form {
	return val.data.(*Slice).Unquote(pos)
}
