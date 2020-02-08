package gfoo

import (
	"fmt"
	"io"
	"strings"
)

var String StringType

func init() {
	String.Init("String")
}

type StringType struct {
	TypeBase
}

func (_ *StringType) Compare(x, y interface{}) Order {
	return Order(strings.Compare(x.(string), y.(string)))
}

func (_ *StringType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "\"%v\"", val.(string))
	return err
}

func (_ *StringType) Unquote(pos Position, val interface{}) Form {
	return NewLiteral(pos, &String, val)
}
