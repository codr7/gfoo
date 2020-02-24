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

func (_ *StringType) Print(val Val, out io.Writer) error {
	_, err := io.WriteString(out, val.data.(string))
	return err
}

func (self *StringType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
