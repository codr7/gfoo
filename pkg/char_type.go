package gfoo

import (
	"fmt"
	"io"
)

var TChar CharType

type CharType struct {
	ValTypeBase
}

func (_ *CharType) Bool(val Val) bool {
	return val.data.(rune) != 0
}

func (_ *CharType) Compare(x, y Val) Order {
	return CompareRune(x.data.(rune), y.data.(rune))
}

func (_ *CharType) Dump(val Val, out io.Writer) error {
	var err error
	
	switch c := val.data.(rune); c {
	case '\n':
		_, err = io.WriteString(out, "\\n")
	default:
		_, err = fmt.Fprintf(out, "\\'%c", c)
	}
	
	return err
}

func (_ *CharType) New(name string, parents...Type) ValType {
	t := new(CharType)
	t.Init(name, parents...)
	return t
}

func (self *CharType) Negate(val *Val) {
	val.Init(&TBool, !self.Bool(*val))
}

func (_ *CharType) Print(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "%c", val.data.(rune))
	return err
}

func (_ *CharType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
