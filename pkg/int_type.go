package gfoo

import (
	"io"
)

var TInt IntType

type IntType struct {
	ValTypeBase
}

func (_ *IntType) Bool(val Val) bool {
	v := val.data.(*Int)
	return !(v.IsInt64() && v.Int64() == 0)
}

func (_ *IntType) Compare(x, y Val) Order {
	return Order(x.data.(*Int).Cmp(y.data.(*Int)))
}

func (_ *IntType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(*Int).String())
	return err
}

func (_ *IntType) For(val Val, action func(Val) error, scope *Scope, pos Pos) error {
	var i Int
	max := val.data.(*Int)
	step := NewInt(1)
	
	for i.Cmp(max) == -1 {
		v := NewInt(0)
		v.Add(v, &i)
		
		if err := action(NewVal(&TInt, v)); err != nil {
			return err
		}

		i.Add(&i, step)
	}

	return nil
}

func (self *IntType) Is(x, y Val) bool {
	return self.Compare(x, y) == Eq 
}

func (_ *IntType) Negate(val *Val) {
	var v Int
	v.Neg(val.data.(*Int))
	val.data = &v
}

func (_ *IntType) New(name string, parents...Type) ValType {
	t := new(IntType)
	t.Init(name, parents...)
	return t
}

func (self *IntType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *IntType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
