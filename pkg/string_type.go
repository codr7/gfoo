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

func (_ *StringType) Compare(x, y interface{}) Order {
	return CompareString(x.(string), y.(string))
}

func (_ *StringType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "\"%v\"", val.(string))
	return err
}

func (self *StringType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}
