package gfoo

import (
	"fmt"
	"io"
)

var Symbol SymbolType

func init() {
	Symbol.Init("Symbol")
}

type SymbolType struct {
	StringType
}

func (_ *SymbolType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "'%v", val.(string))
	return err
}

func (_ *SymbolType) Unquote(pos Position, val interface{}) Form {
	return NewId(pos, val.(string))
}

