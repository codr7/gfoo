package gfoo

import (
	"fmt"
	"io"
)

var TId SymbolType

func init() {
	TId.Init("Id")
}

type SymbolType struct {
	StringType
}

func (_ *SymbolType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "'%v", val.(string))
	return err
}

func (_ *SymbolType) Unquote(pos Pos, val interface{}) Form {
	return NewId(pos, val.(string))
}
