package gfoo

import (
	"io"
)

var TStack StackType

type StackType struct {
	ValTypeBase
}

func (_ *StackType) Bool(val Val) bool {
	return val.data.(*Stack).Len() != 0
}

func (_ *StackType) Compare(x, y Val) Order {
	return x.data.(*Stack).Compare(*y.data.(*Stack))
}

func (self *StackType) Clone(val Val) interface{} {
	return val.data.(*Stack).Clone()
}

func (_ *StackType) Dump(val Val, out io.Writer) error {
	return val.data.(*Stack).Dump(out)
}

func (_ *StackType) Iter(val Val, pos Pos) (Iter, error) {
	in := val.data.(*Stack).items
	i := 0
	
	return func(thread *Thread, pos Pos) (Val, error) {
		if i < len(in) {
			v := in[i]
			i++
			return v, nil
		}

		return Nil, nil
	}, nil
}

func (_ *StackType) Negate(val *Val) {
	v := val.data.(*Stack).Clone()
	
	for i := 0; i < v.Len(); i++ {
		v.items[i].Negate()
	}

	val.data = v
}

func (_ *StackType) New(name string, parents...Type) ValType {
	t := new(StackType)
	t.Init(name, parents...)
	return t
}

func (_ *StackType) Print(val Val, out io.Writer) error {
	return val.data.(*Stack).Print(out)
}

func (_ *StackType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return val.data.(*Stack).Unquote(scope, pos)
}
