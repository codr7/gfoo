package gfoo

import (
	"fmt"
	"io"
)

var TString StringType

func init() {
	TString.Init("String")
}

type StringType struct {
	TypeBase
}

func (_ *StringType) Compare(x, y Val) Order {
	return CompareString(x.data.(string), y.data.(string))
}

func (_ *StringType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "\"%v\"", val.data.(string))
	return err
}

func (self *StringType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
