package gfoo

import (
	"fmt"
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
	xv, yv := x.([]Val), y.([]Val)
	xl, yl := len(xv), len(yv)
	
	for i := 0; i < MinInt(xl, yl); i++ {
		if o := xv[i].Compare(yv[i]); o != Eq {
			return o
		}		
	}
	
	return CompareInt(xl, yl)
}

func (_ *SliceType) Dump(val interface{}, out io.Writer) error {
	return DumpSlice(val.([]Val), out)
}

func (_ *SliceType) Unquote(pos Pos, val interface{}) Form {
	in := val.([]Val)
	out := make([]Form, len(in))

	for i, v := range in {
		out[i] = v.Unquote(pos)
	}

	return NewGroup(pos, out)
}

func DumpSlice(in []Val, out io.Writer) error {
	if _, err := fmt.Fprint(out, "("); err != nil {
		return err
	}

	for i, v := range in {
		if i > 0 {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
		}
		
		if err := v.Dump(out); err != nil {
			return err
		}
	}
	
	if _, err := fmt.Fprint(out, ")"); err != nil {
		return err
	}
	
	return nil
}
