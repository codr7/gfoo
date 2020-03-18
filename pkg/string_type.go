package gfoo

import (
	"fmt"
	"io"
)

var TString StringType

type StringType struct {
	ValTypeBase
}

func (_ *StringType) Bool(val Val) bool {
	return len(val.data.(string)) != 0
}

func (_ *StringType) Compare(x, y Val) Order {
	return CompareString(x.data.(string), y.data.(string))
}

func (_ *StringType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "\"%v\"", val.data.(string))
	return err
}

func (_ *StringType) Iter(val Val, scope *Scope, pos Pos) (Iter, error) {
	in := []rune(val.data.(string))
	i := 0
	
	return func(scope *Scope, pos Pos) (Val, error) {
		if i < len(in) {
			v := NewVal(&TChar, in[i])
			i++
			return v, nil
		}

		return Nil, nil
	}, nil
}

func (self *StringType) Negate(val *Val) {
	val.Init(&TBool, !self.Bool(*val))
}

func (_ *StringType) New(name string, parents...Type) ValType {
	t := new(StringType)
	t.Init(name, parents...)
	return t
}

func (_ *StringType) Print(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(string))
	return err
}

func (self *StringType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
