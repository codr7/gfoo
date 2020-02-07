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

func (typ *SliceType) Dump(val interface{}, out io.Writer) error {
	return DumpSlice(val.([]Value), out)
}

func (typ *SliceType) Compare(x, y interface{}) Order {
	xv, yv := x.([]Value), y.([]Value)
	xl, yl := len(xv), len(yv)
	
	for i := 0; i < MinInt(xl, yl); i++ {
		if o := xv[i].Compare(yv[i]); o != Eq {
			return o
		}		
	}
	
	return CompareInt(xl, yl)
}
