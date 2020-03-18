package gfoo

import (
	"io"
)

var TRecord RecordType

type RecordType struct {
	ValTypeBase
}

func (_ *RecordType) Clone(val Val) interface{} {
	return val.data.(*Record).Clone()
}

func (_ *RecordType) Compare(x, y Val) Order {
	x.data.(*Record).Compare(y.data.(*Record))
	return Eq
}

func (self *RecordType) Dump(val Val, out io.Writer) error {
	if _, err := io.WriteString(out, self.name); err != nil {
		return err
	}

	return val.data.(*Record).Dump(out)
}

func (_ *RecordType) Get(source Val, key string, scope *Scope, pos Pos) (Val, error) {
	return source.data.(*Record).Get(key, Nil), nil
}

func (_ *RecordType) Iter(val Val, scope *Scope, pos Pos) (Iter, error) {
	in := val.data.(*Record)
	i := 0
	
	return func(scope *Scope, pos Pos) (Val, error) {
		if i < in.Len() {
			f := in.fields[i]
			v := NewVal(&TPair, NewPair(NewVal(&TId, f.key), f.val))
			return v, nil
		}

		return Nil, nil
	}, nil
}

func (_ *RecordType) Negate(val *Val) {
	v := val.data.(*Record).Clone()
	
	for i := 0; i < v.Len(); i++ {
		v.fields[i].val.Negate()
	}

	val.data = v
}

func (_ *RecordType) New(name string, parents...Type) ValType {
	t := new(RecordType)
	t.Init(name, parents...)
	return t
}

func (self *RecordType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (_ *RecordType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
