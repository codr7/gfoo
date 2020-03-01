package gfoo

import (
	"io"
)

var TSlice SliceType

func init() {
	TSlice.Init("Slice")
}

type SliceType struct {
	ValTypeBase
}

func (_ *SliceType) Bool(val Val) bool {
	return val.data.(*Slice).Len() != 0
}

func (_ *SliceType) Compare(x, y Val) Order {
	return x.data.(*Slice).Compare(*y.data.(*Slice))
}

func (self *SliceType) Clone(val Val) interface{} {
	return val.data.(*Slice).Clone()
}

func (_ *SliceType) Dump(val Val, out io.Writer) error {
	return val.data.(*Slice).Dump(out)
}

func (_ *SliceType) New(name string, parents...Type) ValType {
	t := new(SliceType)
	t.Init(name, parents...)
	return t
}

func (_ *SliceType) Print(val Val, out io.Writer) error {
	return val.data.(*Slice).Print(out)
}

func (_ *SliceType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return val.data.(*Slice).Unquote(scope, pos)
}
