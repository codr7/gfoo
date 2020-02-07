package gfoo

import (
	"io"
)

var Slice SliceType

func init() {
	Slice.Init("Slice")
}

type SliceType struct {
	TypeBase
}

func (_ *SliceType) Compare(x, y interface{}) Order {
	xv, yv := x.([]Value), y.([]Value)
	xl, yl := len(xv), len(yv)
	
	for i := 0; i < MinInt(xl, yl); i++ {
		if o := xv[i].Compare(yv[i]); o != Eq {
			return o
		}		
	}
	
	return CompareInt(xl, yl)
}

func (_ *SliceType) Dump(val interface{}, out io.Writer) error {
	return DumpSlice(val.([]Value), out)
}

func (_ *SliceType) Unquote(val interface{}) Form {
	in := val.([]Value)
	out := make([]Form, len(in))

	for i, v := range in {
		out[i] = v.Unquote()
	}

	return NewSliceForm(out)
}

type SliceForm struct {
	items []Form
}

func NewSliceForm(items []Form) *SliceForm {
	return &SliceForm{items: items}
}

func (self *SliceForm) Quote() Value {
	v := make([]Value, len(self.items))

	for i, f := range self.items {
		v[i] = f.Quote()
	}
	
	return NewValue(&Slice, v)
}
