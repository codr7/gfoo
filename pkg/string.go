package gfoo

import (
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

func (typ *StringType) Dump(val interface{}, out io.Writer) error {
	_, err := io.WriteString(out, val.(string))
	return err
}

func (typ *StringType) Compare(x, y interface{}) Order {
	return Order(strings.Compare(x.(string), y.(string)))
}
