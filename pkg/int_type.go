package gfoo

import (
	"io"
	"strconv"
)

var TInt IntType

type Int = int64

type IntType struct {
	ValTypeBase
}

func (_ *IntType) Bool(val Val) bool {
	return val.data.(Int) != 0
}

func (_ *IntType) Compare(x, y Val) Order {
	return CompareInt64(x.data.(Int), y.data.(Int))
}

func (_ *IntType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, strconv.FormatInt(val.data.(Int), 10))
	return err
}

func (self *IntType) Is(x, y Val) bool {
	return x.data == y.data 
}

func (_ *IntType) Iter(val Val, scope *Scope, pos Pos) (Iter, error) {
	var i Int
	max := val.data.(Int)
	
	return Iter(func(scope *Scope, pos Pos) (Val, error) {
		if i < max {
			v := NewVal(&TInt, i)
			i++
			return v, nil
		}
		
		return Nil, nil
	}), nil
}

func (_ *IntType) Negate(val *Val) {
	val.data = -val.data.(Int)
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
